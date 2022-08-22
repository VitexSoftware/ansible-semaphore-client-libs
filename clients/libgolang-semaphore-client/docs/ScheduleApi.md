# \ScheduleApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectProjectIdSchedulesPost**](ScheduleApi.md#ProjectProjectIdSchedulesPost) | **Post** /project/{project_id}/schedules | create schedule
[**ProjectProjectIdSchedulesScheduleIdDelete**](ScheduleApi.md#ProjectProjectIdSchedulesScheduleIdDelete) | **Delete** /project/{project_id}/schedules/{schedule_id} | Deletes schedule
[**ProjectProjectIdSchedulesScheduleIdGet**](ScheduleApi.md#ProjectProjectIdSchedulesScheduleIdGet) | **Get** /project/{project_id}/schedules/{schedule_id} | Get schedule
[**ProjectProjectIdSchedulesScheduleIdPut**](ScheduleApi.md#ProjectProjectIdSchedulesScheduleIdPut) | **Put** /project/{project_id}/schedules/{schedule_id} | Updates schedule



## ProjectProjectIdSchedulesPost

> Schedule ProjectProjectIdSchedulesPost(ctx, projectId).Schedule(schedule).Execute()

create schedule

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := int32(1) // int32 | Project ID
    schedule := *openapiclient.NewScheduleRequest() // ScheduleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleApi.ProjectProjectIdSchedulesPost(context.Background(), projectId).Schedule(schedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.ProjectProjectIdSchedulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdSchedulesPost`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.ProjectProjectIdSchedulesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdSchedulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schedule** | [**ScheduleRequest**](ScheduleRequest.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdSchedulesScheduleIdDelete

> ProjectProjectIdSchedulesScheduleIdDelete(ctx, projectId, scheduleId).Execute()

Deletes schedule

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := int32(1) // int32 | Project ID
    scheduleId := int32(9) // int32 | schedule ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleApi.ProjectProjectIdSchedulesScheduleIdDelete(context.Background(), projectId, scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.ProjectProjectIdSchedulesScheduleIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**scheduleId** | **int32** | schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdSchedulesScheduleIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdSchedulesScheduleIdGet

> Schedule ProjectProjectIdSchedulesScheduleIdGet(ctx, projectId, scheduleId).Execute()

Get schedule

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := int32(1) // int32 | Project ID
    scheduleId := int32(9) // int32 | schedule ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleApi.ProjectProjectIdSchedulesScheduleIdGet(context.Background(), projectId, scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.ProjectProjectIdSchedulesScheduleIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdSchedulesScheduleIdGet`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.ProjectProjectIdSchedulesScheduleIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**scheduleId** | **int32** | schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdSchedulesScheduleIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Schedule**](Schedule.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdSchedulesScheduleIdPut

> ProjectProjectIdSchedulesScheduleIdPut(ctx, projectId, scheduleId).Schedule(schedule).Execute()

Updates schedule

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := int32(1) // int32 | Project ID
    scheduleId := int32(9) // int32 | schedule ID
    schedule := *openapiclient.NewScheduleRequest() // ScheduleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleApi.ProjectProjectIdSchedulesScheduleIdPut(context.Background(), projectId, scheduleId).Schedule(schedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.ProjectProjectIdSchedulesScheduleIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**scheduleId** | **int32** | schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdSchedulesScheduleIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **schedule** | [**ScheduleRequest**](ScheduleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

