# TaskOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **int32** |  | [optional] 
**Task** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskOutput

`func NewTaskOutput() *TaskOutput`

NewTaskOutput instantiates a new TaskOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskOutputWithDefaults

`func NewTaskOutputWithDefaults() *TaskOutput`

NewTaskOutputWithDefaults instantiates a new TaskOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskOutput) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskOutput) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskOutput) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TaskOutput) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTask

`func (o *TaskOutput) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TaskOutput) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TaskOutput) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *TaskOutput) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTime

`func (o *TaskOutput) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TaskOutput) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TaskOutput) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *TaskOutput) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOutput

`func (o *TaskOutput) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TaskOutput) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TaskOutput) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TaskOutput) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


