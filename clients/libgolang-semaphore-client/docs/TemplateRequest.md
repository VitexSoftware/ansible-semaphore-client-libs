# TemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**InventoryId** | Pointer to **int32** |  | [optional] 
**RepositoryId** | Pointer to **int32** |  | [optional] 
**EnvironmentId** | Pointer to **int32** |  | [optional] 
**ViewId** | Pointer to **int32** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Playbook** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OverrideArgs** | Pointer to **bool** |  | [optional] 

## Methods

### NewTemplateRequest

`func NewTemplateRequest() *TemplateRequest`

NewTemplateRequest instantiates a new TemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateRequestWithDefaults

`func NewTemplateRequestWithDefaults() *TemplateRequest`

NewTemplateRequestWithDefaults instantiates a new TemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *TemplateRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TemplateRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TemplateRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TemplateRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetInventoryId

`func (o *TemplateRequest) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *TemplateRequest) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *TemplateRequest) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *TemplateRequest) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *TemplateRequest) GetRepositoryId() int32`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *TemplateRequest) GetRepositoryIdOk() (*int32, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *TemplateRequest) SetRepositoryId(v int32)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *TemplateRequest) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *TemplateRequest) GetEnvironmentId() int32`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *TemplateRequest) GetEnvironmentIdOk() (*int32, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *TemplateRequest) SetEnvironmentId(v int32)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *TemplateRequest) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetViewId

`func (o *TemplateRequest) GetViewId() int32`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *TemplateRequest) GetViewIdOk() (*int32, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *TemplateRequest) SetViewId(v int32)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *TemplateRequest) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetAlias

`func (o *TemplateRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *TemplateRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *TemplateRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *TemplateRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetPlaybook

`func (o *TemplateRequest) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *TemplateRequest) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *TemplateRequest) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *TemplateRequest) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetArguments

`func (o *TemplateRequest) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *TemplateRequest) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *TemplateRequest) SetArguments(v string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *TemplateRequest) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOverrideArgs

`func (o *TemplateRequest) GetOverrideArgs() bool`

GetOverrideArgs returns the OverrideArgs field if non-nil, zero value otherwise.

### GetOverrideArgsOk

`func (o *TemplateRequest) GetOverrideArgsOk() (*bool, bool)`

GetOverrideArgsOk returns a tuple with the OverrideArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideArgs

`func (o *TemplateRequest) SetOverrideArgs(v bool)`

SetOverrideArgs sets OverrideArgs field to given value.

### HasOverrideArgs

`func (o *TemplateRequest) HasOverrideArgs() bool`

HasOverrideArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


