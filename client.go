package veoRide

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	AppVersion = "3.11.2"
	PhoneModel = "iPhone 12"
)

var VeoRideNonSuccessResponse = errors.New("the api returned a response code other than 0 (unsuccessful response code)")

type VeoRideAPIResponse struct {
	MSG string `json:"msg"`
	Code int `json:"code"`
}

func (ar *VeoRideAPIResponse) IsSuccess() bool {
	return ar.Code == 0
}

type VeoRideClient struct {
	http.Transport
	AuthToken string
}

func NewVeoRideClient() *VeoRideClient {
	vrc := VeoRideClient{}
	vrc.DisableKeepAlives = true
	return &vrc
}

func (vrc *VeoRideClient) MakeRequest(req *http.Request) (*http.Response, error) {
	req.Header = map[string][]string {
		"Accept": {"*/*"},
		"User-Agent": {"Veo/5 CFNetwork/1237 Darwin/20.4.0"},
		"Content-Type": {"application/json"}, //yeah this gets sent for some reason on the app
		"Authorization": {fmt.Sprintf("Bearer %s", vrc.AuthToken) },
	}
	return vrc.RoundTrip(req)
}
