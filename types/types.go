package types

import "time"

type StumbleResponse struct {
	User struct {
		ID           int    `json:"Id"`
		DeviceID     string `json:"DeviceId"`
		Username     string `json:"Username"`
		Country      string `json:"Country"`
		Crowns       int    `json:"Crowns"`
		HiddenRating int    `json:"HiddenRating"`
		IsBanned     bool   `json:"IsBanned"`
	} `json:"User"`
	Timestamp time.Time `json:"Timestamp"`
}
