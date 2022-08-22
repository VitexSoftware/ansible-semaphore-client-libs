# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
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

### NewTemplate

`func NewTemplate() *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Template) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Template) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Template) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Template) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *Template) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Template) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Template) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Template) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetInventoryId

`func (o *Template) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *Template) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *Template) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *Template) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *Template) GetRepositoryId() int32`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *Template) GetRepositoryIdOk() (*int32, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *Template) SetRepositoryId(v int32)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *Template) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *Template) GetEnvironmentId() int32`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Template) GetEnvironmentIdOk() (*int32, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Template) SetEnvironmentId(v int32)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *Template) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetViewId

`func (o *Template) GetViewId() int32`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *Template) GetViewIdOk() (*int32, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *Template) SetViewId(v int32)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *Template) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetAlias

`func (o *Template) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Template) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Template) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Template) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetPlaybook

`func (o *Template) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *Template) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *Template) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *Template) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetArguments

`func (o *Template) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Template) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Template) SetArguments(v string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *Template) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetDescription

`func (o *Template) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Template) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Template) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Template) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOverrideArgs

`func (o *Template) GetOverrideArgs() bool`

GetOverrideArgs returns the OverrideArgs field if non-nil, zero value otherwise.

### GetOverrideArgsOk

`func (o *Template) GetOverrideArgsOk() (*bool, bool)`

GetOverrideArgsOk returns a tuple with the OverrideArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideArgs

`func (o *Template) SetOverrideArgs(v bool)`

SetOverrideArgs sets OverrideArgs field to given value.

### HasOverrideArgs

`func (o *Template) HasOverrideArgs() bool`

HasOverrideArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


