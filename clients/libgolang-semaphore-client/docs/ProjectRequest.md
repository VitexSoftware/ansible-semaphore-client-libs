# ProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Alert** | Pointer to **bool** |  | [optional] 

## Methods

### NewProjectRequest

`func NewProjectRequest() *ProjectRequest`

NewProjectRequest instantiates a new ProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRequestWithDefaults

`func NewProjectRequestWithDefaults() *ProjectRequest`

NewProjectRequestWithDefaults instantiates a new ProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAlert

`func (o *ProjectRequest) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *ProjectRequest) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *ProjectRequest) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *ProjectRequest) HasAlert() bool`

HasAlert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


