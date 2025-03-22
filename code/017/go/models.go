/* From Python to Go: Go: 017 - NETCONF : modules */

package main

// Import
import (
	"encoding/xml"
)

// Data types
type OpenConfigInterface struct {
	Name   string `xml:"name,omitempty"`
	Config struct {
		Name        string `xml:"name,omitempty"`
		Description string `xml:"description,omitempty"`
		Enabled     bool   `xml:"enabled,omitempty"`
	} `xml:"config,omitempty"`
}
type OpenConfigInterfaces struct {
	XMLName   xml.Name              `xml:"interfaces"`
	Interface []OpenConfigInterface `xml:"interface,omitempty"`
}
type NetconfConfig struct {
	XMLName    xml.Name `xml:"config,omitempty"`
	Interfaces OpenConfigInterfaces
}
type NetconfData struct {
	XMLName    xml.Name `xml:"data,omitempty"`
	Interfaces OpenConfigInterfaces
}

type RPCResponse struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Data    NetconfData
}
