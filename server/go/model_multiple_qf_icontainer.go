/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type MultipleQfIcontainer struct {

	Triggers []Trigger `json:"triggers,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TriggerTimestamp time.Time `json:"triggerTimestamp,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	Time int32 `json:"time,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	TotalVolume int32 `json:"totalVolume,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UplinkVolume int32 `json:"uplinkVolume,omitempty"`

	LocalSequenceNumber int32 `json:"localSequenceNumber"`

	QFIContainerInformation QfiContainerInformation `json:"qFIContainerInformation,omitempty"`
}
