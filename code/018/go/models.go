/* From Python to Go: Go: 018 - GNMI: modules */

package main

// Import

// Data types
type OpenConfigInterface struct {
	Name   string `xml:"name,omitempty"`
	Config struct {
		Name        string `xml:"name,omitempty" json:"name,omitempty"`
		Description string `xml:"description,omitempty" json:"description,omitempty"`
		Enabled     bool   `xml:"enabled,omitempty" json:"enabled,omitempty"`
	} `xml:"config" json:"config"`
}
type OpenConfigInterfaces struct {
	Interface []OpenConfigInterface `xml:"openconfig-interfaces:interface,omitempty" json:"openconfig-interfaces:interface,omitempty"`
}
