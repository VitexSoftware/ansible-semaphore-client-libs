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

// AccessKeyRequest struct for AccessKeyRequest
type AccessKeyRequest struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	ProjectId *int32 `json:"project_id,omitempty"`
}

// NewAccessKeyRequest instantiates a new AccessKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKeyRequest() *AccessKeyRequest {
	this := AccessKeyRequest{}
	return &this
}

// NewAccessKeyRequestWithDefaults instantiates a new AccessKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyRequestWithDefaults() *AccessKeyRequest {
	this := AccessKeyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessKeyRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessKeyRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessKeyRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessKeyRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessKeyRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccessKeyRequest) SetType(v string) {
	o.Type = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AccessKeyRequest) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AccessKeyRequest) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *AccessKeyRequest) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o AccessKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableAccessKeyRequest struct {
	value *AccessKeyRequest
	isSet bool
}

func (v NullableAccessKeyRequest) Get() *AccessKeyRequest {
	return v.value
}

func (v *NullableAccessKeyRequest) Set(val *AccessKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKeyRequest(val *AccessKeyRequest) *NullableAccessKeyRequest {
	return &NullableAccessKeyRequest{value: val, isSet: true}
}

func (v NullableAccessKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


