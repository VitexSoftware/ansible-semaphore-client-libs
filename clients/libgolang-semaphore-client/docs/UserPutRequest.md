# UserPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Alert** | Pointer to **bool** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserPutRequest

`func NewUserPutRequest() *UserPutRequest`

NewUserPutRequest instantiates a new UserPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPutRequestWithDefaults

`func NewUserPutRequestWithDefaults() *UserPutRequest`

NewUserPutRequestWithDefaults instantiates a new UserPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *UserPutRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPutRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPutRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserPutRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UserPutRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPutRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPutRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserPutRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAlert

`func (o *UserPutRequest) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *UserPutRequest) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *UserPutRequest) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *UserPutRequest) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAdmin

`func (o *UserPutRequest) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserPutRequest) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserPutRequest) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserPutRequest) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


