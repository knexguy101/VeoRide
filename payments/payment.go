package payments

import (
	"encoding/json"
	veoRide "github.com/knexguy101/VeoRide"
	"io"
	"net/http"
)

type Payment struct {
	veoRide.VeoRideAPIResponse
	Data []Card `json:"data"`
}

type Card struct {
	Id string `json:"card"`
	Last4 string `json:"last4"`
	Brand string `json:"brand"`
	IsDefault bool `json:"isDefault"`
	Funding string `json:"funding"`
}

func GetPayments(vrc *veoRide.VeoRideClient) (*Payment, error) {
	req, _ := http.NewRequest("GET", "https://cluster-prod.veoride.com/api/customers/payments", nil)
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var payment Payment

	err = json.Unmarshal(resData, &payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

