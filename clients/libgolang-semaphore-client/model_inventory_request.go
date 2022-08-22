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

// InventoryRequest struct for InventoryRequest
type InventoryRequest struct {
	Name *string `json:"name,omitempty"`
	ProjectId *int32 `json:"project_id,omitempty"`
	Inventory *string `json:"inventory,omitempty"`
	SshKeyId *int32 `json:"ssh_key_id,omitempty"`
	BecomeKeyId *int32 `json:"become_key_id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewInventoryRequest instantiates a new InventoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryRequest() *InventoryRequest {
	this := InventoryRequest{}
	return &this
}

// NewInventoryRequestWithDefaults instantiates a new InventoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryRequestWithDefaults() *InventoryRequest {
	this := InventoryRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InventoryRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InventoryRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InventoryRequest) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InventoryRequest) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InventoryRequest) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *InventoryRequest) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *InventoryRequest) GetInventory() string {
	if o == nil || o.Inventory == nil {
		var ret string
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRequest) GetInventoryOk() (*string, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *InventoryRequest) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given string and assigns it to the Inventory field.
func (o *InventoryRequest) SetInventory(v string) {
	o.Inventory = &v
}

// GetSshKeyId returns the SshKeyId field value if set, zero value otherwise.
func (o *InventoryRequest) GetSshKeyId() int32 {
	if o == nil || o.SshKeyId == nil {
		var ret int32
		return ret
	}
	return *o.SshKeyId
}

// GetSshKeyIdOk returns a tuple with the SshKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRequest) GetSshKeyIdOk() (*int32, bool) {
	if o == nil || o.SshKeyId == nil {
		return nil, false
	}
	return o.SshKeyId, true
}

// HasSshKeyId returns a boolean if a field has been set.
func (o *InventoryRequest) HasSshKeyId() bool {
	if o != nil && o.SshKeyId != nil {
		return true
	}

	return false
}

// SetSshKeyId gets a reference to the given int32 and assigns it to the SshKeyId field.
func (o *InventoryRequest) SetSshKeyId(v int32) {
	o.SshKeyId = &v
}

// GetBecomeKeyId returns the BecomeKeyId field value if set, zero value otherwise.
func (o *InventoryRequest) GetBecomeKeyId() int32 {
	if o == nil || o.BecomeKeyId == nil {
		var ret int32
		return ret
	}
	return *o.BecomeKeyId
}

// GetBecomeKeyIdOk returns a tuple with the BecomeKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRequest) GetBecomeKeyIdOk() (*int32, bool) {
	if o == nil || o.BecomeKeyId == nil {
		return nil, false
	}
	return o.BecomeKeyId, true
}

// HasBecomeKeyId returns a boolean if a field has been set.
func (o *InventoryRequest) HasBecomeKeyId() bool {
	if o != nil && o.BecomeKeyId != nil {
		return true
	}

	return false
}

// SetBecomeKeyId gets a reference to the given int32 and assigns it to the BecomeKeyId field.
func (o *InventoryRequest) SetBecomeKeyId(v int32) {
	o.BecomeKeyId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InventoryRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InventoryRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InventoryRequest) SetType(v string) {
	o.Type = &v
}

func (o InventoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.Inventory != nil {
		toSerialize["inventory"] = o.Inventory
	}
	if o.SshKeyId != nil {
		toSerialize["ssh_key_id"] = o.SshKeyId
	}
	if o.BecomeKeyId != nil {
		toSerialize["become_key_id"] = o.BecomeKeyId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryRequest struct {
	value *InventoryRequest
	isSet bool
}

func (v NullableInventoryRequest) Get() *InventoryRequest {
	return v.value
}

func (v *NullableInventoryRequest) Set(val *InventoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryRequest(val *InventoryRequest) *NullableInventoryRequest {
	return &NullableInventoryRequest{value: val, isSet: true}
}

func (v NullableInventoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


