package areas

import (
	"encoding/json"
	"io"
	"net/http"
	"veoRide"
)

type Fences struct {
	veoRide.VeoRideAPIResponse
	Data []FenceData `json:"data"`
}

type FenceCoordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type FenceData struct {
	FenceType        int                `json:"fenceType"`
	FenceCoordinates []FenceCoordinates `json:"fenceCoordinates"`
}

func GetFences(vrc *veoRide.VeoRideClient) (*Fences, error) {
	req, _ := http.NewRequest("GET", "https://cluster-prod.veoride.com/api/customers/areas/fences", nil)
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var fences Fences

	err = json.Unmarshal(resData, &fences)
	if err != nil {
		return nil, err
	}

	return &fences, nil
}