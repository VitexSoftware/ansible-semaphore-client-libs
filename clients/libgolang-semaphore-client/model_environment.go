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

// Environment struct for Environment
type Environment struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ProjectId *int32 `json:"project_id,omitempty"`
	Password *string `json:"password,omitempty"`
	Json *string `json:"json,omitempty"`
}

// NewEnvironment instantiates a new Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironment() *Environment {
	this := Environment{}
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Environment) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Environment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Environment) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Environment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Environment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Environment) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Environment) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Environment) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *Environment) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Environment) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Environment) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Environment) SetPassword(v string) {
	o.Password = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *Environment) GetJson() string {
	if o == nil || o.Json == nil {
		var ret string
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetJsonOk() (*string, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *Environment) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given string and assigns it to the Json field.
func (o *Environment) SetJson(v string) {
	o.Json = &v
}

func (o Environment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironment struct {
	value *Environment
	isSet bool
}

func (v NullableEnvironment) Get() *Environment {
	return v.value
}

func (v *NullableEnvironment) Set(val *Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironment(val *Environment) *NullableEnvironment {
	return &NullableEnvironment{value: val, isSet: true}
}

func (v NullableEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


