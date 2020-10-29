package model

import "time"

type ClientCode struct {
	ID                 string       `json:"id,omitempty"`
	UuID               string       `json:"uu_id"`
	Code               string       `json:"code"`
	Description        string       `json:"description"`
	AutoMapInt         int          `json:"automap_int"`
	ValidatedFlag      bool         `json:"validated_flag"`
	PrimaryMappingFlag bool         `json:"primary_flag"`
	CodeType           CodeType     `json:"code_type"`
	SourceSystem       SourceSystem `json:"source_system"`
	StandardCode       StandardCode `json:"standard_code,omitempty"`
	Metadata           Metadata     `json:"metadata,omitempty"`
}

type CodeMapping struct {
	CodeType     string `json:"code_type"`
	ClientCode   string `json:"client_code"`
	StandardCode string `json:"standard_code"`
}

type CodeType struct {
	ID          string   `json:"id,omitempty"`
	UuID        string   `json:"uuid"`
	Description string   `json:"description"`
	Metadata    Metadata `json:"metadata,omitempty"`
}

type Mappings struct {
	Mapping []CodeMapping `json:"mappings"`
}

type Metadata struct {
	CreatedBy  string    `json:"created_by,omitempty"`
	CreatedDT  time.Time `json:"created_dt,omitempty"`
	ModifiedBy string    `json:"modified_by,omitempty"`
	ModifiedDT time.Time `json:"modified_dt,omitempty"`
}

type SourceSystem struct {
	ID          string                  `json:"id,omitempty"`
	UuID        string                  `json:"uuid"`
	Description string                  `json:"description"`
	Metadata    Metadata                `json:"metadata,omitempty"`
	Application SourceSystemApplication `json:"application"`
}

type SourceSystemApplication struct {
	ID       string   `json:"id,omitempty"`
	UuID     string   `json:"uuid"`
	Name     string   `json:"app_name"`
	Vendor   string   `json:"vendor"`
	Metadata Metadata `json:"metadata,omitempty"`
}

type StandardCode struct {
	ID          string   `json:"id,omitempty"`
	UuID        string   `json:"uu_id"`
	Code        string   `json:"code"`
	Description string   `json:"description"`
	CodeType    CodeType `json:"code_type"`
	Metadata    Metadata `json:"metadata,omitempty"`
}

type StandardCodeByEMR struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Number      int    `json:"number"`
}

type StandardCodeList struct {
	ID            string          `json:"id,omitempty"`
	UuID          string          `json:"uuid"`
	Description   string          `json:"description"`
	Metadata      Metadata        `json:"metadata,omitempty"`
	StandardCodes []StandardCodes `json:"standard_codes,omitempty"`
}

type StandardCodes struct {
	ID          string   `json:"id,omitempty"`
	UuID        string   `json:"uu_id"`
	TypeID      string   `json:"type,omitempty"`
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Metadata    Metadata `json:"metadata,omitempty"`
}
