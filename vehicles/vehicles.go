package vehicles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/knexguy101/VeoRide"
)

type Vehicles struct {
	veoRide.VeoRideAPIResponse
	Vehicles []Vehicle `json:"data"`
}

func GetVehicles(vrc *veoRide.VeoRideClient, lat, lng float64) (*Vehicles, error)  {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://cluster-prod.veoride.com/api/customers/vehicles?lat=%f&lng=%f", lat, lng), nil)
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var vehicles Vehicles

	err = json.Unmarshal(resData, &vehicles)
	if err != nil {
		return nil, err
	}

	return &vehicles, nil
}
