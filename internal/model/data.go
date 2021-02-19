package model

import "time"

type Quality struct {
	ID             string       `json:"_id"`
	Date           time.Time    `json:"date"`
	CityCode       int          `json:"cityCode"`
	CityName       string       `json:"cityName"`
	StatusCode     int          `json:"statusCode"`
	StatusName     string       `json:"statusName"`
	AreaCode       int          `json:"areaCode"`
	AreaName       string       `json:"areaName"`
	ApplyUserName  string       `json:"applyUserName"`
	ApplyTimestamp time.Time    `json:"applyTimestamp"`
	ApplyContent   ApplyContent `json:"applyContent"`
}
type Detail struct {
	Date             time.Time `json:"date"`
	PrimaryPollutant string    `json:"primaryPollutant"`
	AirIndexLevel    string    `json:"airIndexLevel"`
	AirQualityIndex  string    `json:"airQualityIndex"`
}
type ApplyContent struct {
	Detail []Detail `json:"detail"`
}
