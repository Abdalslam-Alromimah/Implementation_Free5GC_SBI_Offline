/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TnapId - Contain the TNAP Identifier see clause5.6.2 of 3GPP TS 23.501.
type TnapId struct {

	// This IE shall be present if the UE is accessing the 5GC via a trusted WLAN access network.When present, it shall contain the SSID of the access point to which the UE is attached, that is received over NGAP,  see IEEE Std 802.11-2012.  
	SsId string `json:"ssId,omitempty"`

	// When present, it shall contain the BSSID of the access point to which the UE is attached, that is received over NGAP, see IEEE Std 802.11-2012.  
	BssId string `json:"bssId,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	CivicAddress string `json:"civicAddress,omitempty"`
}
