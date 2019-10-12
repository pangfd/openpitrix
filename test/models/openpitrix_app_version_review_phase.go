// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixAppVersionReviewPhase openpitrix app version review phase
// swagger:model openpitrixAppVersionReviewPhase
type OpenpitrixAppVersionReviewPhase struct {

	// review message
	Message string `json:"message,omitempty"`

	// user of reviewer
	Operator string `json:"operator,omitempty"`

	// operator type of reviewer eg.[global_admin|developer|business|technical|isv]
	OperatorType string `json:"operator_type,omitempty"`

	// app version review time
	ReviewTime strfmt.DateTime `json:"review_time,omitempty"`

	// review status of app version eg.[isv-in-review|isv-passed|isv-rejected|isv-draft|business-in-review|business-passed|business-rejected|develop-draft|develop-in-review|develop-passed|develop-rejected|develop-draft]
	Status string `json:"status,omitempty"`

	// record status changed time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`
}

// Validate validates this openpitrix app version review phase
func (m *OpenpitrixAppVersionReviewPhase) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixAppVersionReviewPhase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixAppVersionReviewPhase) UnmarshalBinary(b []byte) error {
	var res OpenpitrixAppVersionReviewPhase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}