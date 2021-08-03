package areas

import (
	"encoding/json"
	"io"
	"net/http"
	"veoRide"
)

type Features struct {
	veoRide.VeoRideAPIResponse
	Data FeatureData   `json:"data"`
}

type AgeVerification struct {
	AgeVerificationEnabled bool `json:"ageVerificationEnabled"`
	Age                    int  `json:"age"`
}

type Regulations struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Position int    `json:"position"`
}

type Regulation struct {
	DisplayDuringOnBoarding bool          `json:"displayDuringOnBoarding"`
	Regulations             []Regulations `json:"regulations"`
}

type Availability struct {
	IsOpen      bool        `json:"isOpen"`
	Description interface{} `json:"description"`
}

type FeatureData struct {
	AgeVerification       AgeVerification `json:"ageVerification"`
	MembershipEnabled     bool            `json:"membershipEnabled"`
	RidePhotoEnabled      bool            `json:"ridePhotoEnabled"`
	HoldRideEnabled       bool            `json:"holdRideEnabled"`
	BillingAddressEnabled bool            `json:"billingAddressEnabled"`
	Regulation            Regulation      `json:"regulation"`
	Availability          Availability    `json:"availability"`
	Prompts               []string        `json:"prompts"`
	SurveyURL             interface{}     `json:"surveyUrl"`
	BikeLane              interface{}     `json:"bikeLane"`
}

func GetFeatures(vrc *veoRide.VeoRideClient) (*Features, error) {
	req, _ := http.NewRequest("POST", "https://cluster-prod.veoride.com/api/customers/areas/features", nil)
	res, err := vrc.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var features Features

	err = json.Unmarshal(resData, &features)
	if err != nil {
		return nil, err
	}

	return &features, nil
}