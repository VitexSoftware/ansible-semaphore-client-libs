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

// RepositoryRequest struct for RepositoryRequest
type RepositoryRequest struct {
	Name *string `json:"name,omitempty"`
	ProjectId *int32 `json:"project_id,omitempty"`
	GitUrl *string `json:"git_url,omitempty"`
	GitBranch *string `json:"git_branch,omitempty"`
	SshKeyId *int32 `json:"ssh_key_id,omitempty"`
}

// NewRepositoryRequest instantiates a new RepositoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryRequest() *RepositoryRequest {
	this := RepositoryRequest{}
	return &this
}

// NewRepositoryRequestWithDefaults instantiates a new RepositoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryRequestWithDefaults() *RepositoryRequest {
	this := RepositoryRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RepositoryRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RepositoryRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RepositoryRequest) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *RepositoryRequest) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *RepositoryRequest) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *RepositoryRequest) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetGitUrl returns the GitUrl field value if set, zero value otherwise.
func (o *RepositoryRequest) GetGitUrl() string {
	if o == nil || o.GitUrl == nil {
		var ret string
		return ret
	}
	return *o.GitUrl
}

// GetGitUrlOk returns a tuple with the GitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRequest) GetGitUrlOk() (*string, bool) {
	if o == nil || o.GitUrl == nil {
		return nil, false
	}
	return o.GitUrl, true
}

// HasGitUrl returns a boolean if a field has been set.
func (o *RepositoryRequest) HasGitUrl() bool {
	if o != nil && o.GitUrl != nil {
		return true
	}

	return false
}

// SetGitUrl gets a reference to the given string and assigns it to the GitUrl field.
func (o *RepositoryRequest) SetGitUrl(v string) {
	o.GitUrl = &v
}

// GetGitBranch returns the GitBranch field value if set, zero value otherwise.
func (o *RepositoryRequest) GetGitBranch() string {
	if o == nil || o.GitBranch == nil {
		var ret string
		return ret
	}
	return *o.GitBranch
}

// GetGitBranchOk returns a tuple with the GitBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRequest) GetGitBranchOk() (*string, bool) {
	if o == nil || o.GitBranch == nil {
		return nil, false
	}
	return o.GitBranch, true
}

// HasGitBranch returns a boolean if a field has been set.
func (o *RepositoryRequest) HasGitBranch() bool {
	if o != nil && o.GitBranch != nil {
		return true
	}

	return false
}

// SetGitBranch gets a reference to the given string and assigns it to the GitBranch field.
func (o *RepositoryRequest) SetGitBranch(v string) {
	o.GitBranch = &v
}

// GetSshKeyId returns the SshKeyId field value if set, zero value otherwise.
func (o *RepositoryRequest) GetSshKeyId() int32 {
	if o == nil || o.SshKeyId == nil {
		var ret int32
		return ret
	}
	return *o.SshKeyId
}

// GetSshKeyIdOk returns a tuple with the SshKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRequest) GetSshKeyIdOk() (*int32, bool) {
	if o == nil || o.SshKeyId == nil {
		return nil, false
	}
	return o.SshKeyId, true
}

// HasSshKeyId returns a boolean if a field has been set.
func (o *RepositoryRequest) HasSshKeyId() bool {
	if o != nil && o.SshKeyId != nil {
		return true
	}

	return false
}

// SetSshKeyId gets a reference to the given int32 and assigns it to the SshKeyId field.
func (o *RepositoryRequest) SetSshKeyId(v int32) {
	o.SshKeyId = &v
}

func (o RepositoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.GitUrl != nil {
		toSerialize["git_url"] = o.GitUrl
	}
	if o.GitBranch != nil {
		toSerialize["git_branch"] = o.GitBranch
	}
	if o.SshKeyId != nil {
		toSerialize["ssh_key_id"] = o.SshKeyId
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryRequest struct {
	value *RepositoryRequest
	isSet bool
}

func (v NullableRepositoryRequest) Get() *RepositoryRequest {
	return v.value
}

func (v *NullableRepositoryRequest) Set(val *RepositoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryRequest(val *RepositoryRequest) *NullableRepositoryRequest {
	return &NullableRepositoryRequest{value: val, isSet: true}
}

func (v NullableRepositoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


