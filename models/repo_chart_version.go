// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RepoChartVersion repo chart version
//
// swagger:model repo.ChartVersion
type RepoChartVersion struct {

	// Annotations are additional mappings uninterpreted by Helm,
	// made available for inspection by other applications.
	Annotations map[string]string `json:"annotations,omitempty"`

	// The API Version of this chart. Required.
	APIVersion string `json:"apiVersion,omitempty"`

	// The version of the application enclosed inside of this chart.
	AppVersion string `json:"appVersion,omitempty"`

	// ChecksumDeprecated is deprecated in Helm 3, and therefore ignored. Helm 3 replaced
	// this with Digest. However, with a strict YAML parser enabled, a field must be
	// present on the struct for backwards compatibility.
	Checksum string `json:"checksum,omitempty"`

	// The condition to check to enable chart
	Condition string `json:"condition,omitempty"`

	// created
	Created string `json:"created,omitempty"`

	// Dependencies are a list of dependencies for a chart.
	Dependencies []*ChartDependency `json:"dependencies"`

	// Whether or not this chart is deprecated
	Deprecated bool `json:"deprecated,omitempty"`

	// A one-sentence description of the chart
	Description string `json:"description,omitempty"`

	// digest
	Digest string `json:"digest,omitempty"`

	// EngineDeprecated is deprecated in Helm 3, and therefore ignored. However, with a strict
	// YAML parser enabled, this field must be present.
	Engine string `json:"engine,omitempty"`

	// The URL to a relevant project page, git repo, or contact person
	Home string `json:"home,omitempty"`

	// The URL to an icon file.
	Icon string `json:"icon,omitempty"`

	// A list of string keywords
	Keywords []string `json:"keywords"`

	// KubeVersion is a SemVer constraint specifying the version of Kubernetes required.
	KubeVersion string `json:"kubeVersion,omitempty"`

	// A list of name and URL/email address combinations for the maintainer(s)
	Maintainers []*ChartMaintainer `json:"maintainers"`

	// The name of the chart. Required.
	Name string `json:"name,omitempty"`

	// removed
	Removed bool `json:"removed,omitempty"`

	// Source is the URL to the source code of this chart
	Sources []string `json:"sources"`

	// The tags to check to enable chart
	Tags string `json:"tags,omitempty"`

	// TillerVersionDeprecated is deprecated in Helm 3, and therefore ignored. However, with a strict
	// YAML parser enabled, this field must be present.
	TillerVersion string `json:"tillerVersion,omitempty"`

	// Specifies the chart type: application or library
	Type string `json:"type,omitempty"`

	// URLDeprecated is deprecated in Helm 3, superseded by URLs. It is ignored. However,
	// with a strict YAML parser enabled, this must be present on the struct.
	URL string `json:"url,omitempty"`

	// urls
	Urls []string `json:"urls"`

	// A SemVer 2 conformant version string of the chart. Required.
	Version string `json:"version,omitempty"`
}

// Validate validates this repo chart version
func (m *RepoChartVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintainers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepoChartVersion) validateDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RepoChartVersion) validateMaintainers(formats strfmt.Registry) error {
	if swag.IsZero(m.Maintainers) { // not required
		return nil
	}

	for i := 0; i < len(m.Maintainers); i++ {
		if swag.IsZero(m.Maintainers[i]) { // not required
			continue
		}

		if m.Maintainers[i] != nil {
			if err := m.Maintainers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("maintainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("maintainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this repo chart version based on the context it is used
func (m *RepoChartVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaintainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepoChartVersion) contextValidateDependencies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dependencies); i++ {

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RepoChartVersion) contextValidateMaintainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Maintainers); i++ {

		if m.Maintainers[i] != nil {
			if err := m.Maintainers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("maintainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("maintainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepoChartVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepoChartVersion) UnmarshalBinary(b []byte) error {
	var res RepoChartVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
