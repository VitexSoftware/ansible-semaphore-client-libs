# APIToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewAPIToken

`func NewAPIToken() *APIToken`

NewAPIToken instantiates a new APIToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITokenWithDefaults

`func NewAPITokenWithDefaults() *APIToken`

NewAPITokenWithDefaults instantiates a new APIToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APIToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *APIToken) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *APIToken) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *APIToken) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *APIToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpired

`func (o *APIToken) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *APIToken) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *APIToken) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *APIToken) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetUserId

`func (o *APIToken) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APIToken) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APIToken) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *APIToken) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


