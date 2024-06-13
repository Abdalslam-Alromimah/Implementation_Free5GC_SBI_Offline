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

// NodeFunctionality struct for NodeFunctionality
type NodeFunctionality struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NodeFunctionality) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NodeFunctionality)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NodeFunctionality) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNodeFunctionality struct {
	value *NodeFunctionality
	isSet bool
}

func (v NullableNodeFunctionality) Get() *NodeFunctionality {
	return v.value
}

func (v *NullableNodeFunctionality) Set(val *NodeFunctionality) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeFunctionality) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeFunctionality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeFunctionality(val *NodeFunctionality) *NullableNodeFunctionality {
	return &NullableNodeFunctionality{value: val, isSet: true}
}

func (v NullableNodeFunctionality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeFunctionality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

