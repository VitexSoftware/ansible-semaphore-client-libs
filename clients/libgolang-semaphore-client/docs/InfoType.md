# InfoType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**UpdateBody** | Pointer to **string** |  | [optional] 
**Update** | Pointer to [**InfoTypeUpdate**](InfoTypeUpdate.md) |  | [optional] 

## Methods

### NewInfoType

`func NewInfoType() *InfoType`

NewInfoType instantiates a new InfoType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoTypeWithDefaults

`func NewInfoTypeWithDefaults() *InfoType`

NewInfoTypeWithDefaults instantiates a new InfoType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *InfoType) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InfoType) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InfoType) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InfoType) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUpdateBody

`func (o *InfoType) GetUpdateBody() string`

GetUpdateBody returns the UpdateBody field if non-nil, zero value otherwise.

### GetUpdateBodyOk

`func (o *InfoType) GetUpdateBodyOk() (*string, bool)`

GetUpdateBodyOk returns a tuple with the UpdateBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateBody

`func (o *InfoType) SetUpdateBody(v string)`

SetUpdateBody sets UpdateBody field to given value.

### HasUpdateBody

`func (o *InfoType) HasUpdateBody() bool`

HasUpdateBody returns a boolean if a field has been set.

### GetUpdate

`func (o *InfoType) GetUpdate() InfoTypeUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *InfoType) GetUpdateOk() (*InfoTypeUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *InfoType) SetUpdate(v InfoTypeUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *InfoType) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


