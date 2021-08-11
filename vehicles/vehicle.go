package vehicles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/knexguy101/VeoRide"
)

type Location struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Type    int     `json:"type"`
	Created string  `json:"created"`
}

type Vehicle struct {
	VehicleNumber  int      `json:"vehicleNumber"`
	VehicleType    int      `json:"vehicleType"`
	VehicleVersion string   `json:"vehicleVersion"`
	IotBattery     int      `json:"iotBattery"`
	VehicleBattery int      `json:"vehicleBattery"`
	Location       Location `json:"location"`
}

type VehicleInfo struct {
	veoRide.VeoRideAPIResponse
	Data Data   `json:"data"`
}

type Price struct {
	Price     float64 `json:"price"`
	Frequency int     `json:"frequency"`
	UnlockFee float64 `json:"unlockFee"`
}

type Data struct {
	VehicleNumber  int         `json:"vehicleNumber"`
	VehicleType    int         `json:"vehicleType"`
	VehicleVersion string      `json:"vehicleVersion"`
	Locked         bool        `json:"locked"`
	Mac            string      `json:"mac"`
	Connected      bool        `json:"connected"`
	VehicleBattery int         `json:"vehicleBattery"`
	Price          Price       `json:"price"`
	ChainLock      interface{} `json:"chainLock"`
}

func (vh *Vehicle) GetInfo(vrc *veoRide.VeoRideClient) (*VehicleInfo, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://cluster-prod.veoride.com/api/customers/vehicles/number/%d", vh.VehicleNumber), nil)
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var vehicleInfo VehicleInfo

	err = json.Unmarshal(resData, &vehicleInfo)
	if err != nil {
		return nil, err
	}

	return &vehicleInfo, nil
}

func (vh *Vehicle) Ring(vrc *veoRide.VeoRideClient) (*veoRide.VeoRideAPIResponse, error) {
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://cluster-prod.veoride.com/api/customers/vehicles/number/%d/find", vh.VehicleNumber), nil)
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
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
