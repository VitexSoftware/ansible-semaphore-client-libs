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

// ProjectProjectIdUsersGetRequest struct for ProjectProjectIdUsersGetRequest
type ProjectProjectIdUsersGetRequest struct {
	UserId *int32 `json:"user_id,omitempty"`
	Admin *bool `json:"admin,omitempty"`
}

// NewProjectProjectIdUsersGetRequest instantiates a new ProjectProjectIdUsersGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectProjectIdUsersGetRequest() *ProjectProjectIdUsersGetRequest {
	this := ProjectProjectIdUsersGetRequest{}
	return &this
}

// NewProjectProjectIdUsersGetRequestWithDefaults instantiates a new ProjectProjectIdUsersGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectProjectIdUsersGetRequestWithDefaults() *ProjectProjectIdUsersGetRequest {
	this := ProjectProjectIdUsersGetRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ProjectProjectIdUsersGetRequest) GetUserId() int32 {
	if o == nil || o.UserId == nil {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdUsersGetRequest) GetUserIdOk() (*int32, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ProjectProjectIdUsersGetRequest) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ProjectProjectIdUsersGetRequest) SetUserId(v int32) {
	o.UserId = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *ProjectProjectIdUsersGetRequest) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdUsersGetRequest) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *ProjectProjectIdUsersGetRequest) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *ProjectProjectIdUsersGetRequest) SetAdmin(v bool) {
	o.Admin = &v
}

func (o ProjectProjectIdUsersGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	return json.Marshal(toSerialize)
}

type NullableProjectProjectIdUsersGetRequest struct {
	value *ProjectProjectIdUsersGetRequest
	isSet bool
}

func (v NullableProjectProjectIdUsersGetRequest) Get() *ProjectProjectIdUsersGetRequest {
	return v.value
}

func (v *NullableProjectProjectIdUsersGetRequest) Set(val *ProjectProjectIdUsersGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectIdUsersGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectIdUsersGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectIdUsersGetRequest(val *ProjectProjectIdUsersGetRequest) *NullableProjectProjectIdUsersGetRequest {
	return &NullableProjectProjectIdUsersGetRequest{value: val, isSet: true}
}

func (v NullableProjectProjectIdUsersGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectIdUsersGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

