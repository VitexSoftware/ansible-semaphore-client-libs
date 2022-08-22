# AccessKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccessKeyRequest

`func NewAccessKeyRequest() *AccessKeyRequest`

NewAccessKeyRequest instantiates a new AccessKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyRequestWithDefaults

`func NewAccessKeyRequestWithDefaults() *AccessKeyRequest`

NewAccessKeyRequestWithDefaults instantiates a new AccessKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccessKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AccessKeyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessKeyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessKeyRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccessKeyRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProjectId

`func (o *AccessKeyRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AccessKeyRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AccessKeyRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AccessKeyRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


