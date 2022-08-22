# ScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CronFormat** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 

## Methods

### NewScheduleRequest

`func NewScheduleRequest() *ScheduleRequest`

NewScheduleRequest instantiates a new ScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRequestWithDefaults

`func NewScheduleRequestWithDefaults() *ScheduleRequest`

NewScheduleRequestWithDefaults instantiates a new ScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCronFormat

`func (o *ScheduleRequest) GetCronFormat() string`

GetCronFormat returns the CronFormat field if non-nil, zero value otherwise.

### GetCronFormatOk

`func (o *ScheduleRequest) GetCronFormatOk() (*string, bool)`

GetCronFormatOk returns a tuple with the CronFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronFormat

`func (o *ScheduleRequest) SetCronFormat(v string)`

SetCronFormat sets CronFormat field to given value.

### HasCronFormat

`func (o *ScheduleRequest) HasCronFormat() bool`

HasCronFormat returns a boolean if a field has been set.

### GetProjectId

`func (o *ScheduleRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ScheduleRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ScheduleRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ScheduleRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTemplateId

`func (o *ScheduleRequest) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ScheduleRequest) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ScheduleRequest) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ScheduleRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


