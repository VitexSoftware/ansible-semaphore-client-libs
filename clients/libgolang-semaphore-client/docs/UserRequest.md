# UserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Alert** | Pointer to **bool** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserRequest

`func NewUserRequest() *UserRequest`

NewUserRequest instantiates a new UserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRequestWithDefaults

`func NewUserRequestWithDefaults() *UserRequest`

NewUserRequestWithDefaults instantiates a new UserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *UserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAlert

`func (o *UserRequest) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *UserRequest) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *UserRequest) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *UserRequest) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAdmin

`func (o *UserRequest) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserRequest) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserRequest) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserRequest) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


