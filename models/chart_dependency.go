// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChartDependency chart dependency
//
// swagger:model chart.Dependency
type ChartDependency struct {

	// Alias usable alias to be used for the chart
	Alias string `json:"alias,omitempty"`

	// A yaml path that resolves to a boolean, used for enabling/disabling charts (e.g. subchart1.enabled )
	Condition string `json:"condition,omitempty"`

	// Enabled bool determines if chart should be loaded
	Enabled bool `json:"enabled,omitempty"`

	// ImportValues holds the mapping of source values to parent key to be imported. Each item can be a
	// string or pair of child/parent sublist items.
	ImportValues []interface{} `json:"import-values"`

	// Name is the name of the dependency.
	//
	// This must mach the name in the dependency's Chart.yaml.
	Name string `json:"name,omitempty"`

	// The URL to the repository.
	//
	// Appending `index.yaml` to this string should result in a URL that can be
	// used to fetch the repository index.
	Repository string `json:"repository,omitempty"`

	// Tags can be used to group charts for enabling/disabling together
	Tags []string `json:"tags"`

	// Version is the version (range) of this chart.
	//
	// A lock file will always produce a single version, while a dependency
	// may contain a semantic version range.
	Version string `json:"version,omitempty"`
}

// Validate validates this chart dependency
func (m *ChartDependency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chart dependency based on context it is used
func (m *ChartDependency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChartDependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChartDependency) UnmarshalBinary(b []byte) error {
	var res ChartDependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
