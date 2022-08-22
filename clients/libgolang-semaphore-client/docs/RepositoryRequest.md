# RepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**GitUrl** | Pointer to **string** |  | [optional] 
**GitBranch** | Pointer to **string** |  | [optional] 
**SshKeyId** | Pointer to **int32** |  | [optional] 

## Methods

### NewRepositoryRequest

`func NewRepositoryRequest() *RepositoryRequest`

NewRepositoryRequest instantiates a new RepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRequestWithDefaults

`func NewRepositoryRequestWithDefaults() *RepositoryRequest`

NewRepositoryRequestWithDefaults instantiates a new RepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositoryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepositoryRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *RepositoryRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RepositoryRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RepositoryRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RepositoryRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetGitUrl

`func (o *RepositoryRequest) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *RepositoryRequest) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *RepositoryRequest) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *RepositoryRequest) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### GetGitBranch

`func (o *RepositoryRequest) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *RepositoryRequest) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *RepositoryRequest) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *RepositoryRequest) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetSshKeyId

`func (o *RepositoryRequest) GetSshKeyId() int32`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *RepositoryRequest) GetSshKeyIdOk() (*int32, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *RepositoryRequest) SetSshKeyId(v int32)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *RepositoryRequest) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


