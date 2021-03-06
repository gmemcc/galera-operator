// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// S3 s3
// swagger:model s3
type S3 struct {

	// access key
	// Required: true
	AccessKey *string `json:"accessKey"`

	// backup dir
	// Required: true
	BackupDir *string `json:"backupDir"`

	// backup name
	// Required: true
	BackupName *string `json:"backupName"`

	// bucket
	// Required: true
	Bucket *string `json:"bucket"`

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`

	// force path style
	ForcePathStyle *bool `json:"forcePathStyle,omitempty"`

	// region
	Region *string `json:"region,omitempty"`

	// secret key
	// Required: true
	SecretKey *string `json:"secretKey"`
}

// Validate validates this s3
func (m *S3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("accessKey", "body", m.AccessKey); err != nil {
		return err
	}

	return nil
}

func (m *S3) validateBackupDir(formats strfmt.Registry) error {

	if err := validate.Required("backupDir", "body", m.BackupDir); err != nil {
		return err
	}

	return nil
}

func (m *S3) validateBackupName(formats strfmt.Registry) error {

	if err := validate.Required("backupName", "body", m.BackupName); err != nil {
		return err
	}

	return nil
}

func (m *S3) validateBucket(formats strfmt.Registry) error {

	if err := validate.Required("bucket", "body", m.Bucket); err != nil {
		return err
	}

	return nil
}

func (m *S3) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

func (m *S3) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secretKey", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3) UnmarshalBinary(b []byte) error {
	var res S3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
