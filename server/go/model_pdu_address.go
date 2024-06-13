/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PduAddress struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	PduIPv4Address string `json:"pduIPv4Address,omitempty"`

	PduIPv6AddresswithPrefix Ipv6Addr `json:"pduIPv6AddresswithPrefix,omitempty"`

	PduAddressprefixlength int32 `json:"pduAddressprefixlength,omitempty"`

	IPv4dynamicAddressFlag bool `json:"iPv4dynamicAddressFlag,omitempty"`

	IPv6dynamicPrefixFlag bool `json:"iPv6dynamicPrefixFlag,omitempty"`
}