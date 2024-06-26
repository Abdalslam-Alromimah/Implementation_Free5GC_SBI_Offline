/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MAPDUSessionInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MAPDUSessionInformation{}

// MAPDUSessionInformation struct for MAPDUSessionInformation
type MAPDUSessionInformation struct {
	MAPDUSessionIndicator *MaPduIndication `json:"mAPDUSessionIndicator,omitempty"`
	ATSSSCapability *AtsssCapability `json:"aTSSSCapability,omitempty"`
}

// NewMAPDUSessionInformation instantiates a new MAPDUSessionInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMAPDUSessionInformation() *MAPDUSessionInformation {
	this := MAPDUSessionInformation{}
	return &this
}

// NewMAPDUSessionInformationWithDefaults instantiates a new MAPDUSessionInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMAPDUSessionInformationWithDefaults() *MAPDUSessionInformation {
	this := MAPDUSessionInformation{}
	return &this
}

// GetMAPDUSessionIndicator returns the MAPDUSessionIndicator field value if set, zero value otherwise.
func (o *MAPDUSessionInformation) GetMAPDUSessionIndicator() MaPduIndication {
	if o == nil || IsNil(o.MAPDUSessionIndicator) {
		var ret MaPduIndication
		return ret
	}
	return *o.MAPDUSessionIndicator
}

// GetMAPDUSessionIndicatorOk returns a tuple with the MAPDUSessionIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MAPDUSessionInformation) GetMAPDUSessionIndicatorOk() (*MaPduIndication, bool) {
	if o == nil || IsNil(o.MAPDUSessionIndicator) {
		return nil, false
	}
	return o.MAPDUSessionIndicator, true
}

// HasMAPDUSessionIndicator returns a boolean if a field has been set.
func (o *MAPDUSessionInformation) HasMAPDUSessionIndicator() bool {
	if o != nil && !IsNil(o.MAPDUSessionIndicator) {
		return true
	}

	return false
}

// SetMAPDUSessionIndicator gets a reference to the given MaPduIndication and assigns it to the MAPDUSessionIndicator field.
func (o *MAPDUSessionInformation) SetMAPDUSessionIndicator(v MaPduIndication) {
	o.MAPDUSessionIndicator = &v
}

// GetATSSSCapability returns the ATSSSCapability field value if set, zero value otherwise.
func (o *MAPDUSessionInformation) GetATSSSCapability() AtsssCapability {
	if o == nil || IsNil(o.ATSSSCapability) {
		var ret AtsssCapability
		return ret
	}
	return *o.ATSSSCapability
}

// GetATSSSCapabilityOk returns a tuple with the ATSSSCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MAPDUSessionInformation) GetATSSSCapabilityOk() (*AtsssCapability, bool) {
	if o == nil || IsNil(o.ATSSSCapability) {
		return nil, false
	}
	return o.ATSSSCapability, true
}

// HasATSSSCapability returns a boolean if a field has been set.
func (o *MAPDUSessionInformation) HasATSSSCapability() bool {
	if o != nil && !IsNil(o.ATSSSCapability) {
		return true
	}

	return false
}

// SetATSSSCapability gets a reference to the given AtsssCapability and assigns it to the ATSSSCapability field.
func (o *MAPDUSessionInformation) SetATSSSCapability(v AtsssCapability) {
	o.ATSSSCapability = &v
}

func (o MAPDUSessionInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MAPDUSessionInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MAPDUSessionIndicator) {
		toSerialize["mAPDUSessionIndicator"] = o.MAPDUSessionIndicator
	}
	if !IsNil(o.ATSSSCapability) {
		toSerialize["aTSSSCapability"] = o.ATSSSCapability
	}
	return toSerialize, nil
}

type NullableMAPDUSessionInformation struct {
	value *MAPDUSessionInformation
	isSet bool
}

func (v NullableMAPDUSessionInformation) Get() *MAPDUSessionInformation {
	return v.value
}

func (v *NullableMAPDUSessionInformation) Set(val *MAPDUSessionInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableMAPDUSessionInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableMAPDUSessionInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMAPDUSessionInformation(val *MAPDUSessionInformation) *NullableMAPDUSessionInformation {
	return &NullableMAPDUSessionInformation{value: val, isSet: true}
}

func (v NullableMAPDUSessionInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMAPDUSessionInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


