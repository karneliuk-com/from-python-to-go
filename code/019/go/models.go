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

type NetboxDcimDevices struct {
	/* Struct to store data from NetBox */
	Count   uint64 `json:"count"`
	Results []struct {
		Name       string `json:"name"`
		PrimaryIp4 struct {
			Address string `json:"address"`
		} `json:"primary_ip4"`
		Platform struct {
			Slug string `json:"slug"`
		} `json:"platform"`
		CustomFields struct {
			GnmiPort uint64 `json:"gnmi_port"`
		} `json:"custom_fields"`
	} `json:"results"`
}
