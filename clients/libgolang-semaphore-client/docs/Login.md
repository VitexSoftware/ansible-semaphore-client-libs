# Login

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to **string** | Username/Email address | [optional] 
**Password** | Pointer to **string** | Password | [optional] 

## Methods

### NewLogin

`func NewLogin() *Login`

NewLogin instantiates a new Login object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWithDefaults

`func NewLoginWithDefaults() *Login`

NewLoginWithDefaults instantiates a new Login object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *Login) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Login) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Login) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *Login) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetPassword

`func (o *Login) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Login) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Login) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Login) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


