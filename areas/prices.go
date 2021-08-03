package areas

import (
	"encoding/json"
	"io"
	"net/http"
	"veoRide"
)

type Prices struct {
	veoRide.VeoRideAPIResponse
	Data []PriceData `json:"data"`
}

type PriceData struct {
	Price       float64 `json:"price"`
	Frequency   int     `json:"frequency"`
	UnlockFee   float64 `json:"unlockFee"`
	VehicleType int     `json:"vehicleType"`
}

func GetPrices(vrc *veoRide.VeoRideClient) (*Prices, error) {
	req, _ := http.NewRequest("POST", "https://cluster-prod.veoride.com/api/customers/areas/prices", nil)
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var prices Prices

	err = json.Unmarshal(resData, &prices)
	if err != nil {
		return nil, err
	}

	return &prices, nil
}