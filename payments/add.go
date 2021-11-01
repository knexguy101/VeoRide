package payments

import (
	"bytes"
	"encoding/json"
	veoRide "github.com/knexguy101/VeoRide"
	"io"
	"net/http"
)

type CreditCardInfo struct {
	CC string
	Month string
	Year string
	CVV string
}

type AddPaymentPayload struct {
	TokenId string `json:"tokenId"`
}

func AddPayment(cci *CreditCardInfo, vrc *veoRide.VeoRideClient) (*veoRide.VeoRideAPIResponse, error) {

	token, err := TokenizePayment(cci.CC, cci.Month, cci.Year, cci.CVV)
	if err != nil {
		return nil, err
	}

	payload, _ := json.Marshal(AddPaymentPayload{
		TokenId: token.ID,
	})

	req, _ := http.NewRequest("POST", "https://cluster-prod.veoride.com/api/customers/payments", bytes.NewBuffer(payload))
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var payment veoRide.VeoRideAPIResponse

	err = json.Unmarshal(resData, &payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
