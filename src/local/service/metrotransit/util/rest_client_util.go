package util

import (
	"fmt"
	"net/http"

	"github.com/GoPracPro/src/local/platform/rest"
	"github.com/GoPracPro/src/local/platform/rest/util"
)

//ExecuteGET for input URL and attach reponse to input responseModel which should be
//reference of the response model which needs to be type case as api reponse specifications
func ExecuteGET(extURL string, responseModel interface{}) error {
	req, err := http.NewRequest(http.MethodGet, extURL, nil)
	if err != nil {
		return err
	}
	params := make(map[string]string, 1)
	params["format"] = "json"
	resp, err := rest.GetRestClient().Execute(req, params, nil)
	if err != nil {
		return err
	}
	success := util.CheckIsSuccess(resp)
	if !success {
		return fmt.Errorf("Response is not success from api url : %s. Status : %s", extURL, resp.Status)
	}
	return util.DecodeResponseForStruct(resp, responseModel)
}
