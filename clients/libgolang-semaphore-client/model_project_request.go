/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.8.34
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ProjectRequest struct for ProjectRequest
type ProjectRequest struct {
	Name *string `json:"name,omitempty"`
	Alert *bool `json:"alert,omitempty"`
}

// NewProjectRequest instantiates a new ProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRequest() *ProjectRequest {
	this := ProjectRequest{}
	return &this
}

// NewProjectRequestWithDefaults instantiates a new ProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRequestWithDefaults() *ProjectRequest {
	this := ProjectRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectRequest) SetName(v string) {
	o.Name = &v
}

// GetAlert returns the Alert field value if set, zero value otherwise.
func (o *ProjectRequest) GetAlert() bool {
	if o == nil || o.Alert == nil {
		var ret bool
		return ret
	}
	return *o.Alert
}

// GetAlertOk returns a tuple with the Alert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRequest) GetAlertOk() (*bool, bool) {
	if o == nil || o.Alert == nil {
		return nil, false
	}
	return o.Alert, true
}

// HasAlert returns a boolean if a field has been set.
func (o *ProjectRequest) HasAlert() bool {
	if o != nil && o.Alert != nil {
		return true
	}

	return false
}

// SetAlert gets a reference to the given bool and assigns it to the Alert field.
func (o *ProjectRequest) SetAlert(v bool) {
	o.Alert = &v
}

func (o ProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Alert != nil {
		toSerialize["alert"] = o.Alert
	}
	return json.Marshal(toSerialize)
}

type NullableProjectRequest struct {
	value *ProjectRequest
	isSet bool
}

func (v NullableProjectRequest) Get() *ProjectRequest {
	return v.value
}

func (v *NullableProjectRequest) Set(val *ProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRequest(val *ProjectRequest) *NullableProjectRequest {
	return &NullableProjectRequest{value: val, isSet: true}
}

func (v NullableProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


