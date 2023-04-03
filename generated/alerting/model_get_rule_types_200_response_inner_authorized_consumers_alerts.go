/*
Alerting

OpenAPI schema for alerting endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerting

import (
	"encoding/json"
)

// checks if the GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts{}

// GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts struct for GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts
type GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts struct {
	All  *bool `json:"all,omitempty"`
	Read *bool `json:"read,omitempty"`
}

// NewGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts instantiates a new GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts() *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts {
	this := GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts{}
	return &this
}

// NewGetRuleTypes200ResponseInnerAuthorizedConsumersAlertsWithDefaults instantiates a new GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRuleTypes200ResponseInnerAuthorizedConsumersAlertsWithDefaults() *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts {
	this := GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) GetAll() bool {
	if o == nil || IsNil(o.All) {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) GetAllOk() (*bool, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) SetAll(v bool) {
	o.All = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) GetRead() bool {
	if o == nil || IsNil(o.Read) {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) GetReadOk() (*bool, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) HasRead() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) SetRead(v bool) {
	o.Read = &v
}

func (o GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.All) {
		toSerialize["all"] = o.All
	}
	if !IsNil(o.Read) {
		toSerialize["read"] = o.Read
	}
	return toSerialize, nil
}

type NullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts struct {
	value *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts
	isSet bool
}

func (v NullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) Get() *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts {
	return v.value
}

func (v *NullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) Set(val *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts(val *GetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) *NullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts {
	return &NullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts{value: val, isSet: true}
}

func (v NullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRuleTypes200ResponseInnerAuthorizedConsumersAlerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}