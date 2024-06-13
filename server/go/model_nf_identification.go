/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type NfIdentification struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NFName string `json:"nFName,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	NFIPv4Address string `json:"nFIPv4Address,omitempty"`

	NFIPv6Address Ipv6Addr `json:"nFIPv6Address,omitempty"`

	NFPLMNID PlmnId `json:"nFPLMNID,omitempty"`

	NodeFunctionality NodeFunctionality `json:"nodeFunctionality"`

	NFFqdn string `json:"nFFqdn,omitempty"`
}
