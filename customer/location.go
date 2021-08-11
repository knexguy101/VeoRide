package customer

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"github.com/knexguy101/VeoRide"
)

type VeoRideChangeLocationRequest struct {
	Location Location `json:"location"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func ChangeLocation(vrc *veoRide.VeoRideClient, lat, lng float64) (*veoRide.VeoRideAPIResponse, error) {
	resData, _ := json.Marshal(VeoRideChangeLocationRequest{
		Location: Location {
			Lat: lat,
			Lng: lng,
		},
	})
	req, _ := http.NewRequest("PATCH", "https://cluster-prod.veoride.com/api/customers", bytes.NewBuffer(resData))
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var veoRes veoRide.VeoRideAPIResponse

	err = json.Unmarshal(resData, &veoRes)
	if err != nil {
		return nil, err
	}

	return &veoRes, nil
}