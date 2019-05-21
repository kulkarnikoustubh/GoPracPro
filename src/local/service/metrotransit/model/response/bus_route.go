package response

//BusRoute to hold response from rest api
type BusRoute struct {
	Description string `json:"Description"`
	ProviderID  string `json:"ProviderID"`
	Route       string `json:"Route"`
}
