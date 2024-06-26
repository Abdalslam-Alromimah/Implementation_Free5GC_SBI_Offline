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

type PduSessionChargingInformation struct {

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	// Deprecated
	ChargingId int32 `json:"chargingId,omitempty"`

	UserInformation UserInformation `json:"userInformation,omitempty"`

	UserLocationinfo UserLocation `json:"userLocationinfo,omitempty"`

	MAPDUNon3GPPUserLocationInfo UserLocation `json:"mAPDUNon3GPPUserLocationInfo,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	UserLocationTime time.Time `json:"userLocationTime,omitempty"`

	PresenceReportingAreaInformation map[string]PresenceInfo `json:"presenceReportingAreaInformation,omitempty"`

	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UetimeZone string `json:"uetimeZone,omitempty"`

	PduSessionInformation PduSessionInformation `json:"pduSessionInformation"`

	// indicating a time in seconds.
	UnitCountInactivityTimer int32 `json:"unitCountInactivityTimer,omitempty"`

	RANSecondaryRATUsageReport RanSecondaryRatUsageReport `json:"rANSecondaryRATUsageReport,omitempty"`
}
