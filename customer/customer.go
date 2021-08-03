package customer

import (
	"encoding/json"
	"io"
	"net/http"
	"veoRide"
)

type Customer struct {
	veoRide.VeoRideAPIResponse
	Data struct {
		Phone           int64       `json:"phone"`
		CountryCode     int         `json:"countryCode"`
		Deposit         float64     `json:"deposit"`
		RideCredit      float64     `json:"rideCredit"`
		Status          int         `json:"status"`
		FullName        interface{} `json:"fullName"`
		InviteCode      string      `json:"inviteCode"`
		Email           interface{} `json:"email"`
		EmailType       int         `json:"emailType"`
		RideCount       int         `json:"rideCount"`
		IsAgeValidated  bool        `json:"isAgeValidated"`
		IsLowIncome     bool        `json:"isLowIncome"`
		IsAutoReloaded  bool        `json:"isAutoReloaded"`
		IsEducationMode bool        `json:"isEducationMode"`
	} `json:"data"`
}

func GetCustomer(vrc *veoRide.VeoRideClient) (*Customer, error) {
	req, _ := http.NewRequest("GET", "https://cluster-prod.veoride.com/api/customers", nil)
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var customer Customer

	err = json.Unmarshal(resData, &customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}