package payments

import (
	"encoding/json"
	veoRide "github.com/knexguy101/VeoRide"
	"io"
	"net/http"
	"strings"
)

func RemovePayment(cardId string, vrc *veoRide.VeoRideClient) (*veoRide.VeoRideAPIResponse, error) {

	req, _ := http.NewRequest("DELETE", "https://cluster-prod.veoride.com/api/customers/payments", strings.NewReader(cardId))
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
