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
	} `xml:"config,omitempty"`
}
type OpenConfigInterfaces struct {
	XMLName   xml.Name              `xml:"interfaces"`
	Interface []OpenConfigInterface `xml:"interface,omitempty"`
}

type RPCResponse struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Data    struct {
		Interfaces OpenConfigInterfaces
	} `xml:"data,omitempty"`
}
