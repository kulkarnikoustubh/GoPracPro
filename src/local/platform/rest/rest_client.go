package rest

import (
	"net/http"
	"sync"
	"time"

	"github.com/GoPracPro/src/local/platform/rest/errors"
	"github.com/GoPracPro/src/local/platform/rest/util"
)

var once sync.Once
var instance ClientWrapper

//ClientWrapper for rest client implementaion
type ClientWrapper interface {
	/**Execute method will invoke underlying framework for invoking rest calls */
	Execute(req *http.Request, parameters map[string]string, payload interface{}) (*http.Response, errors.RestClientError)
}

type restClientWrapperImpl struct {
	client *http.Client
}

//GetRestClient gives rest client
func GetRestClient() ClientWrapper {
	once.Do(func() {
		instance = &restClientWrapperImpl{createClient()}
	})
	return instance
}

//Execute the request
func (r *restClientWrapperImpl) Execute(req *http.Request, parameters map[string]string, payload interface{}) (*http.Response, errors.RestClientError) {
	addQueryParams(req, parameters)
	if payload != nil {
		e := util.EncodePayload(req, payload)
		if e != nil {
			return nil, errors.Errorf("EncodePayload of payload : %v failed. Reason : %v ", payload, e)
		}
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, errors.Errorf("Execution of the request : %v failed. Reason : %v ", req, err)
	}
	return resp, nil
}

func addQueryParams(req *http.Request, parameters map[string]string) {
	if parameters != nil {
		qry := req.URL.Query()
		for key, value := range parameters {
			qry.Add(key, value)
		}
		req.URL.RawQuery = qry.Encode()
	}
}

func createClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	return client
}
