package veoRide

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type VeoRideLoginRequest struct {
	Code string `json:"code"`
	Phone int64 `json:"phone"`
	AppVersion string `json:"appVersion"`
	PhoneModel string `json:"phoneModel"`
}

type VeoRideLoginResponse struct {
	VeoRideAPIResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

func (vrc *VeoRideClient) Login(phone int64, smsCodeChan chan string) error {
	err := vrc.getCode(phone)
	if err != nil {
		return err
	}

	code, ok := <- smsCodeChan
	if !ok {
		return errors.New("channel did not recieve sms code")
	}

	return vrc.sendLogin(phone, code)
}

func (vrc *VeoRideClient) sendLogin(phone int64, code string) error {

	//send request to login
	resDat, _ := json.Marshal(VeoRideLoginRequest{
		Code: code,
		Phone: phone,
		AppVersion: AppVersion,
		PhoneModel: PhoneModel,
	})
	req, _ := http.NewRequest("POST", "https://cluster-prod.veoride.com/api/customers/auth/auth-code/verification", bytes.NewBuffer(resDat))
	req.Header = map[string][]string {
		"Accept": {"*/*"},
		"User-Agent": {"Veo/5 CFNetwork/1237 Darwin/20.4.0"},
		"Content-Type": {"application/json"}, //yeah this gets sent for some reason on the app
	}
	res, err := vrc.RoundTrip(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resDat, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var apiRes VeoRideLoginResponse

	err = json.Unmarshal(resDat, &apiRes)
	if err != nil {
		return err
	}

	if !apiRes.IsSuccess() {
		return VeoRideNonSuccessResponse
	}

	vrc.AuthToken = apiRes.Data.Token

	return nil
}

func (vrc *VeoRideClient) getCode(phone int64) error {

	//send request for sms code
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://cluster-prod.veoride.com/api/customers/auth/auth-code?phone=%d", phone), nil)
	req.Header = map[string][]string {
		"Accept": {"*/*"},
		"User-Agent": {"Veo/5 CFNetwork/1237 Darwin/20.4.0"},
		"Content-Type": {"application/json"}, //yeah this gets sent for some reason on the app
	}
	res, err := vrc.RoundTrip(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resDat, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var apiRes VeoRideAPIResponse

	err = json.Unmarshal(resDat, &apiRes)
	if err != nil {
		return err
	}

	if !apiRes.IsSuccess() {
		return VeoRideNonSuccessResponse
	}

	return nil
}
