package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

const (
	protocolSeparator string = "://"
	separator         string = "/"
)

var successStatuses = map[int]bool{
	http.StatusOK:                   true,
	http.StatusCreated:              true,
	http.StatusAccepted:             true,
	http.StatusNonAuthoritativeInfo: true,
	http.StatusNoContent:            true,
	http.StatusResetContent:         true,
	http.StatusPartialContent:       true,
	http.StatusMultiStatus:          true,
	http.StatusAlreadyReported:      true,
	http.StatusIMUsed:               true,
}

//CheckIsSuccess as per http success codes list
func CheckIsSuccess(response *http.Response) bool {
	return successStatuses[response.StatusCode] == true
}

//ExtractRootLevelParentData extract the value of the rootLevelParentKey in the response
func ExtractRootLevelParentData(response *http.Response, rootLevelParentKey string) (interface{}, error) {
	data, jsonErr := DecodeResponse(response)
	if jsonErr != nil {
		return nil, jsonErr
	}
	rootLevelData, isMap := data.(map[string]interface{})
	if isMap {
		return rootLevelData[rootLevelParentKey], nil
	}
	return nil, errors.New("Unable to find value")
}

//DecodeResponse decodes the response to
func DecodeResponse(response *http.Response) (interface{}, error) {
	rawData, ioErr := Decode(response.Body)
	if ioErr == nil && rawData != nil && len(rawData) > 1 {
		return DeserializeBytes(rawData)
	}
	return nil, ioErr
}

//DecodeResponseForStruct decodes response into provided struct
func DecodeResponseForStruct(response *http.Response, output interface{}) error {
	rawData, ioErr := Decode(response.Body)
	if ioErr == nil && rawData != nil && len(rawData) > 1 {
		return DeserializeBytesToStruct(rawData, output)
	}
	return ioErr
}

//Decode converts input body to a byte array
func Decode(body io.ReadCloser) ([]byte, error) {
	//this may be the case as there are some endpoints which respond in detail
	//body message incase of errors as well
	if body == nil {
		return nil, errors.New("Deserialize response is not a valid operation for empty response")
	}
	defer body.Close()
	return ioutil.ReadAll(body)
}

//EncodePayload encode incoming struct and set to the body of the nput request
func EncodePayload(req *http.Request, payload interface{}) error {
	buf, err := Encode(payload)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(buf)
	return nil
}

//Encode function gets the input data and encode to byte buffer
func Encode(data interface{}) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&data)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

//ExtractRequestBody extract the body from the request
func ExtractRequestBody(r *http.Request) ([]byte, error) {
	var jsonData interface{}
	err := extractBody(r, &jsonData)
	if err == nil {
		resp, e := transform(jsonData)
		if e == nil {
			return resp.Bytes(), nil
		}
		err = e
	}
	return []byte{}, err
}

func transform(jsonData interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	if jsonData == nil {
		return buf, nil
	}
	err := json.NewEncoder(&buf).Encode(jsonData)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

func extractBody(r *http.Request, data interface{}) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(data)
}

//TranfornsFormToStruct transform from input map to destination struct
func TranfornsFormToStruct(valueMap map[string]string, dest interface{}) error {
	return mapstructure.Decode(valueMap, dest)
}

/**
* DeserializeBytes function expects inputJson as transformable json data of map
* return error : incase input is not a json transformable
 */
func DeserializeBytes(inputJson []byte) (interface{}, error) {
	var output interface{}
	err := DeserializeBytesToStruct(inputJson, &output)
	return output, err
}

/**
* DeserializeBytesToStruct function expects inputJson as transformable json data of map
* return error : incase input is not a json transformable
 */
func DeserializeBytesToStruct(inputJson []byte, output interface{}) error {
	err := json.Unmarshal(inputJson, output)
	if err != nil {
		return err
	}
	return nil
}
