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

// NrLocation - Contains the NR user location.
type NrLocation struct {

	Tai Tai `json:"tai"`

	Ncgi Ncgi `json:"ncgi"`

	IgnoreNcgi bool `json:"ignoreNcgi,omitempty"`

	// The value represents the elapsed time in minutes since the last network contact of the mobile station. Value \"0\" indicates that the location information was obtained after a successful paging procedure for Active Location Retrieval when the UE is in idle mode or after a successful  NG-RAN location reporting procedure with the eNB when the UE is in connected mode. Any other value than \"0\" indicates that the location information is the last known one. See 3GPP TS 29.002 clause 17.7.8. 
	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`

	// Refer to geographical Information. See 3GPP TS 23.032 clause 7.3.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used. 
	GeographicalInformation string `json:"geographicalInformation,omitempty"`

	// Refers to Calling Geodetic Location. See ITU-T Recommendation Q.763 (1999) [24] clause 3.88.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used. 
	GeodeticInformation string `json:"geodeticInformation,omitempty"`

	GlobalGnbId GlobalRanNodeId `json:"globalGnbId,omitempty"`

	NtnTaiInfo NtnTaiInfo `json:"ntnTaiInfo,omitempty"`
}
