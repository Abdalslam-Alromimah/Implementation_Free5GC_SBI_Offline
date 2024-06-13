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

type QfiContainerInformation struct {

	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	QFI int32 `json:"qFI,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeofFirstUsage time.Time `json:"timeofFirstUsage,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeofLastUsage time.Time `json:"timeofLastUsage,omitempty"`

	QoSInformation *QosData `json:"qoSInformation,omitempty"`

	QoSCharacteristics QosCharacteristics `json:"qoSCharacteristics,omitempty"`

	UserLocationInformation UserLocation `json:"userLocationInformation,omitempty"`

	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UetimeZone string `json:"uetimeZone,omitempty"`

	PresenceReportingAreaInformation map[string]PresenceInfo `json:"presenceReportingAreaInformation,omitempty"`

	RATType RatType `json:"rATType,omitempty"`

	ServingNetworkFunctionID []ServingNetworkFunctionId `json:"servingNetworkFunctionID,omitempty"`

	Var3gppPSDataOffStatus Model3GpppsDataOffStatus `json:"3gppPSDataOffStatus,omitempty"`
}
