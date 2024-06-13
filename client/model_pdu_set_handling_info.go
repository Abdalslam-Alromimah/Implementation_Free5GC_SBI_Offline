/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PduSetHandlingInfo Possible values are: - \"ALL_PDUS_NEEDED\": All PDUs of the PDU Set are needed - \"ALL_PDUS_NOT_NEEDED\": All PDUs of the PDU Set are not needed 
type PduSetHandlingInfo struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PduSetHandlingInfo) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PduSetHandlingInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PduSetHandlingInfo) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePduSetHandlingInfo struct {
	value *PduSetHandlingInfo
	isSet bool
}

func (v NullablePduSetHandlingInfo) Get() *PduSetHandlingInfo {
	return v.value
}

func (v *NullablePduSetHandlingInfo) Set(val *PduSetHandlingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSetHandlingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSetHandlingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSetHandlingInfo(val *PduSetHandlingInfo) *NullablePduSetHandlingInfo {
	return &NullablePduSetHandlingInfo{value: val, isSet: true}
}

func (v NullablePduSetHandlingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSetHandlingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

