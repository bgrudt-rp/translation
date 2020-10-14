package model

type SourceSystem struct {
	UuID        string                  `json:"uuid"`
	Description string                  `json:"description"`
	Application SourceSystemApplication `json:"application"`
}

type SourceSystemApplication struct {
	UuID   string `json:"uuid"`
	Name   string `json:"app_name"`
	Vendor string `json:"vendor"`
}
