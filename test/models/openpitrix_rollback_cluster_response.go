// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRollbackClusterResponse openpitrix rollback cluster response
// swagger:model openpitrixRollbackClusterResponse
type OpenpitrixRollbackClusterResponse struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// job id
	JobID string `json:"job_id,omitempty"`
}

// Validate validates this openpitrix rollback cluster response
func (m *OpenpitrixRollbackClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRollbackClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRollbackClusterResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRollbackClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}