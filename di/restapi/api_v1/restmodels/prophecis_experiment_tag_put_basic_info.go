// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProphecisExperimentTagPutBasicInfo prophecis experiment tag put basic info
// swagger:model ProphecisExperimentTagPutBasicInfo
type ProphecisExperimentTagPutBasicInfo struct {

	// Experiment Id
	ExpID float64 `json:"exp_id,omitempty"`

	// Experiment's Tag
	ExpTag string `json:"exp_tag,omitempty"`
}

// Validate validates this prophecis experiment tag put basic info
func (m *ProphecisExperimentTagPutBasicInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProphecisExperimentTagPutBasicInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProphecisExperimentTagPutBasicInfo) UnmarshalBinary(b []byte) error {
	var res ProphecisExperimentTagPutBasicInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}