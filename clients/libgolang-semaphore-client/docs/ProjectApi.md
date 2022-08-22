# \ProjectApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectProjectIdDelete**](ProjectApi.md#ProjectProjectIdDelete) | **Delete** /project/{project_id}/ | Delete project
[**ProjectProjectIdEnvironmentEnvironmentIdDelete**](ProjectApi.md#ProjectProjectIdEnvironmentEnvironmentIdDelete) | **Delete** /project/{project_id}/environment/{environment_id} | Removes environment
[**ProjectProjectIdEnvironmentEnvironmentIdPut**](ProjectApi.md#ProjectProjectIdEnvironmentEnvironmentIdPut) | **Put** /project/{project_id}/environment/{environment_id} | Update environment
[**ProjectProjectIdEnvironmentGet**](ProjectApi.md#ProjectProjectIdEnvironmentGet) | **Get** /project/{project_id}/environment | Get environment
[**ProjectProjectIdEnvironmentPost**](ProjectApi.md#ProjectProjectIdEnvironmentPost) | **Post** /project/{project_id}/environment | Add environment
[**ProjectProjectIdEventsGet**](ProjectApi.md#ProjectProjectIdEventsGet) | **Get** /project/{project_id}/events | Get Events related to this project
[**ProjectProjectIdGet**](ProjectApi.md#ProjectProjectIdGet) | **Get** /project/{project_id}/ | Fetch project
[**ProjectProjectIdInventoryGet**](ProjectApi.md#ProjectProjectIdInventoryGet) | **Get** /project/{project_id}/inventory | Get inventory
[**ProjectProjectIdInventoryInventoryIdDelete**](ProjectApi.md#ProjectProjectIdInventoryInventoryIdDelete) | **Delete** /project/{project_id}/inventory/{inventory_id} | Removes inventory
[**ProjectProjectIdInventoryInventoryIdPut**](ProjectApi.md#ProjectProjectIdInventoryInventoryIdPut) | **Put** /project/{project_id}/inventory/{inventory_id} | Updates inventory
[**ProjectProjectIdInventoryPost**](ProjectApi.md#ProjectProjectIdInventoryPost) | **Post** /project/{project_id}/inventory | create inventory
[**ProjectProjectIdKeysGet**](ProjectApi.md#ProjectProjectIdKeysGet) | **Get** /project/{project_id}/keys | Get access keys linked to project
[**ProjectProjectIdKeysKeyIdDelete**](ProjectApi.md#ProjectProjectIdKeysKeyIdDelete) | **Delete** /project/{project_id}/keys/{key_id} | Removes access key
[**ProjectProjectIdKeysKeyIdPut**](ProjectApi.md#ProjectProjectIdKeysKeyIdPut) | **Put** /project/{project_id}/keys/{key_id} | Updates access key
[**ProjectProjectIdKeysPost**](ProjectApi.md#ProjectProjectIdKeysPost) | **Post** /project/{project_id}/keys | Add access key
[**ProjectProjectIdPut**](ProjectApi.md#ProjectProjectIdPut) | **Put** /project/{project_id}/ | Update project
[**ProjectProjectIdRepositoriesGet**](ProjectApi.md#ProjectProjectIdRepositoriesGet) | **Get** /project/{project_id}/repositories | Get repositories
[**ProjectProjectIdRepositoriesPost**](ProjectApi.md#ProjectProjectIdRepositoriesPost) | **Post** /project/{project_id}/repositories | Add repository
[**ProjectProjectIdRepositoriesRepositoryIdDelete**](ProjectApi.md#ProjectProjectIdRepositoriesRepositoryIdDelete) | **Delete** /project/{project_id}/repositories/{repository_id} | Removes repository
[**ProjectProjectIdTasksGet**](ProjectApi.md#ProjectProjectIdTasksGet) | **Get** /project/{project_id}/tasks | Get Tasks related to current project
[**ProjectProjectIdTasksLastGet**](ProjectApi.md#ProjectProjectIdTasksLastGet) | **Get** /project/{project_id}/tasks/last | Get last 200 Tasks related to current project
[**ProjectProjectIdTasksPost**](ProjectApi.md#ProjectProjectIdTasksPost) | **Post** /project/{project_id}/tasks | Starts a job
[**ProjectProjectIdTasksTaskIdDelete**](ProjectApi.md#ProjectProjectIdTasksTaskIdDelete) | **Delete** /project/{project_id}/tasks/{task_id} | Deletes task (including output)
[**ProjectProjectIdTasksTaskIdGet**](ProjectApi.md#ProjectProjectIdTasksTaskIdGet) | **Get** /project/{project_id}/tasks/{task_id} | Get a single task
[**ProjectProjectIdTasksTaskIdOutputGet**](ProjectApi.md#ProjectProjectIdTasksTaskIdOutputGet) | **Get** /project/{project_id}/tasks/{task_id}/output | Get task output
[**ProjectProjectIdTemplatesGet**](ProjectApi.md#ProjectProjectIdTemplatesGet) | **Get** /project/{project_id}/templates | Get template
[**ProjectProjectIdTemplatesPost**](ProjectApi.md#ProjectProjectIdTemplatesPost) | **Post** /project/{project_id}/templates | create template
[**ProjectProjectIdTemplatesTemplateIdDelete**](ProjectApi.md#ProjectProjectIdTemplatesTemplateIdDelete) | **Delete** /project/{project_id}/templates/{template_id} | Removes template
[**ProjectProjectIdTemplatesTemplateIdGet**](ProjectApi.md#ProjectProjectIdTemplatesTemplateIdGet) | **Get** /project/{project_id}/templates/{template_id} | Get template
[**ProjectProjectIdTemplatesTemplateIdPut**](ProjectApi.md#ProjectProjectIdTemplatesTemplateIdPut) | **Put** /project/{project_id}/templates/{template_id} | Updates template
[**ProjectProjectIdUsersGet**](ProjectApi.md#ProjectProjectIdUsersGet) | **Get** /project/{project_id}/users | Get users linked to project
[**ProjectProjectIdUsersPost**](ProjectApi.md#ProjectProjectIdUsersPost) | **Post** /project/{project_id}/users | Link user to project
[**ProjectProjectIdUsersUserIdAdminDelete**](ProjectApi.md#ProjectProjectIdUsersUserIdAdminDelete) | **Delete** /project/{project_id}/users/{user_id}/admin | Revoke admin privileges
[**ProjectProjectIdUsersUserIdAdminPost**](ProjectApi.md#ProjectProjectIdUsersUserIdAdminPost) | **Post** /project/{project_id}/users/{user_id}/admin | Makes user admin
[**ProjectProjectIdUsersUserIdDelete**](ProjectApi.md#ProjectProjectIdUsersUserIdDelete) | **Delete** /project/{project_id}/users/{user_id} | Removes user from project
[**ProjectProjectIdViewsGet**](ProjectApi.md#ProjectProjectIdViewsGet) | **Get** /project/{project_id}/views | Get view
[**ProjectProjectIdViewsPost**](ProjectApi.md#ProjectProjectIdViewsPost) | **Post** /project/{project_id}/views | create view
[**ProjectProjectIdViewsViewIdDelete**](ProjectApi.md#ProjectProjectIdViewsViewIdDelete) | **Delete** /project/{project_id}/views/{view_id} | Removes view
[**ProjectProjectIdViewsViewIdGet**](ProjectApi.md#ProjectProjectIdViewsViewIdGet) | **Get** /project/{project_id}/views/{view_id} | Get view
[**ProjectProjectIdViewsViewIdPut**](ProjectApi.md#ProjectProjectIdViewsViewIdPut) | **Put** /project/{project_id}/views/{view_id} | Updates view



## ProjectProjectIdDelete

> ProjectProjectIdDelete(ctx, projectId).Execute()

Delete project

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdDelete(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdEnvironmentEnvironmentIdDelete

> ProjectProjectIdEnvironmentEnvironmentIdDelete(ctx, projectId, environmentId).Execute()

Removes environment

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
    environmentId := int32(6) // int32 | environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdEnvironmentEnvironmentIdDelete(context.Background(), projectId, environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdEnvironmentEnvironmentIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**environmentId** | **int32** | environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEnvironmentEnvironmentIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdEnvironmentEnvironmentIdPut

> ProjectProjectIdEnvironmentEnvironmentIdPut(ctx, projectId, environmentId).Environment(environment).Execute()

Update environment

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
    environmentId := int32(6) // int32 | environment ID
    environment := *openapiclient.NewEnvironmentRequest() // EnvironmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdEnvironmentEnvironmentIdPut(context.Background(), projectId, environmentId).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdEnvironmentEnvironmentIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**environmentId** | **int32** | environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEnvironmentEnvironmentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environment** | [**EnvironmentRequest**](EnvironmentRequest.md) |  | 

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


## ProjectProjectIdEnvironmentGet

> []Environment ProjectProjectIdEnvironmentGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get environment

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
    sort := "db-deploy" // string | sorting name
    order := "desc" // string | ordering manner

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdEnvironmentGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdEnvironmentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdEnvironmentGet`: []Environment
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdEnvironmentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEnvironmentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]Environment**](Environment.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdEnvironmentPost

> ProjectProjectIdEnvironmentPost(ctx, projectId).Environment(environment).Execute()

Add environment

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
    environment := *openapiclient.NewEnvironmentRequest() // EnvironmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdEnvironmentPost(context.Background(), projectId).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdEnvironmentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEnvironmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | [**EnvironmentRequest**](EnvironmentRequest.md) |  | 

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


## ProjectProjectIdEventsGet

> []Event ProjectProjectIdEventsGet(ctx, projectId).Execute()

Get Events related to this project

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdEventsGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdEventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdEventsGet`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Event**](Event.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdGet

> Project ProjectProjectIdGet(ctx, projectId).Execute()

Fetch project

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdGet`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdInventoryGet

> []Inventory ProjectProjectIdInventoryGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get inventory

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
    sort := "sort_example" // string | sorting name
    order := "order_example" // string | ordering manner

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdInventoryGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdInventoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdInventoryGet`: []Inventory
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdInventoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]Inventory**](Inventory.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdInventoryInventoryIdDelete

> ProjectProjectIdInventoryInventoryIdDelete(ctx, projectId, inventoryId).Execute()

Removes inventory

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
    inventoryId := int32(5) // int32 | inventory ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdInventoryInventoryIdDelete(context.Background(), projectId, inventoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdInventoryInventoryIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**inventoryId** | **int32** | inventory ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdInventoryInventoryIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdInventoryInventoryIdPut

> ProjectProjectIdInventoryInventoryIdPut(ctx, projectId, inventoryId).Inventory(inventory).Execute()

Updates inventory

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
    inventoryId := int32(5) // int32 | inventory ID
    inventory := *openapiclient.NewInventoryRequest() // InventoryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdInventoryInventoryIdPut(context.Background(), projectId, inventoryId).Inventory(inventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdInventoryInventoryIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**inventoryId** | **int32** | inventory ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdInventoryInventoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inventory** | [**InventoryRequest**](InventoryRequest.md) |  | 

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


## ProjectProjectIdInventoryPost

> Inventory ProjectProjectIdInventoryPost(ctx, projectId).Inventory(inventory).Execute()

create inventory

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
    inventory := *openapiclient.NewInventoryRequest() // InventoryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdInventoryPost(context.Background(), projectId).Inventory(inventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdInventoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdInventoryPost`: Inventory
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdInventoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdInventoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inventory** | [**InventoryRequest**](InventoryRequest.md) |  | 

### Return type

[**Inventory**](Inventory.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdKeysGet

> []AccessKey ProjectProjectIdKeysGet(ctx, projectId).Sort(sort).Order(order).KeyType(keyType).Execute()

Get access keys linked to project

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
    sort := "type" // string | sorting name
    order := "asc" // string | ordering manner
    keyType := "none" // string | Filter by key type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdKeysGet(context.Background(), projectId).Sort(sort).Order(order).KeyType(keyType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdKeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdKeysGet`: []AccessKey
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 
 **keyType** | **string** | Filter by key type | 

### Return type

[**[]AccessKey**](AccessKey.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdKeysKeyIdDelete

> ProjectProjectIdKeysKeyIdDelete(ctx, projectId, keyId).Execute()

Removes access key

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
    keyId := int32(3) // int32 | key ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdKeysKeyIdDelete(context.Background(), projectId, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdKeysKeyIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**keyId** | **int32** | key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdKeysKeyIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdKeysKeyIdPut

> ProjectProjectIdKeysKeyIdPut(ctx, projectId, keyId).AccessKey(accessKey).Execute()

Updates access key

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
    keyId := int32(3) // int32 | key ID
    accessKey := *openapiclient.NewAccessKeyRequest() // AccessKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdKeysKeyIdPut(context.Background(), projectId, keyId).AccessKey(accessKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdKeysKeyIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**keyId** | **int32** | key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdKeysKeyIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessKey** | [**AccessKeyRequest**](AccessKeyRequest.md) |  | 

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


## ProjectProjectIdKeysPost

> ProjectProjectIdKeysPost(ctx, projectId).AccessKey(accessKey).Execute()

Add access key

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
    accessKey := *openapiclient.NewAccessKeyRequest() // AccessKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdKeysPost(context.Background(), projectId).AccessKey(accessKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdKeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessKey** | [**AccessKeyRequest**](AccessKeyRequest.md) |  | 

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


## ProjectProjectIdPut

> ProjectProjectIdPut(ctx, projectId).Project(project).Execute()

Update project

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
    project := *openapiclient.NewProjectProjectIdDeleteRequest() // ProjectProjectIdDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdPut(context.Background(), projectId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**ProjectProjectIdDeleteRequest**](ProjectProjectIdDeleteRequest.md) |  | 

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


## ProjectProjectIdRepositoriesGet

> []Repository ProjectProjectIdRepositoriesGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get repositories

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
    sort := "sort_example" // string | sorting name
    order := "order_example" // string | ordering manner

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdRepositoriesGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdRepositoriesGet`: []Repository
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdRepositoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdRepositoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdRepositoriesPost

> ProjectProjectIdRepositoriesPost(ctx, projectId).Repository(repository).Execute()

Add repository

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
    repository := *openapiclient.NewRepositoryRequest() // RepositoryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdRepositoriesPost(context.Background(), projectId).Repository(repository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdRepositoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdRepositoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | [**RepositoryRequest**](RepositoryRequest.md) |  | 

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


## ProjectProjectIdRepositoriesRepositoryIdDelete

> ProjectProjectIdRepositoriesRepositoryIdDelete(ctx, projectId, repositoryId).Execute()

Removes repository

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
    repositoryId := int32(4) // int32 | repository ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdRepositoriesRepositoryIdDelete(context.Background(), projectId, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdRepositoriesRepositoryIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**repositoryId** | **int32** | repository ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdRepositoriesRepositoryIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdTasksGet

> []Task ProjectProjectIdTasksGet(ctx, projectId).Execute()

Get Tasks related to current project

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTasksGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTasksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdTasksGet`: []Task
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Task**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksLastGet

> []Task ProjectProjectIdTasksLastGet(ctx, projectId).Execute()

Get last 200 Tasks related to current project

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTasksLastGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTasksLastGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdTasksLastGet`: []Task
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdTasksLastGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksLastGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Task**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksPost

> Task ProjectProjectIdTasksPost(ctx, projectId).Task(task).Execute()

Starts a job

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
    task := *openapiclient.NewProjectProjectIdTasksGetRequest() // ProjectProjectIdTasksGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTasksPost(context.Background(), projectId).Task(task).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTasksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdTasksPost`: Task
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdTasksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | [**ProjectProjectIdTasksGetRequest**](ProjectProjectIdTasksGetRequest.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksTaskIdDelete

> ProjectProjectIdTasksTaskIdDelete(ctx, projectId, taskId).Execute()

Deletes task (including output)

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
    taskId := int32(8) // int32 | task ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTasksTaskIdDelete(context.Background(), projectId, taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTasksTaskIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**taskId** | **int32** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksTaskIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdTasksTaskIdGet

> Task ProjectProjectIdTasksTaskIdGet(ctx, projectId, taskId).Execute()

Get a single task

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
    taskId := int32(8) // int32 | task ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTasksTaskIdGet(context.Background(), projectId, taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTasksTaskIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdTasksTaskIdGet`: Task
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdTasksTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**taskId** | **int32** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Task**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksTaskIdOutputGet

> []TaskOutput ProjectProjectIdTasksTaskIdOutputGet(ctx, projectId, taskId).Execute()

Get task output

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
    taskId := int32(8) // int32 | task ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTasksTaskIdOutputGet(context.Background(), projectId, taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTasksTaskIdOutputGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdTasksTaskIdOutputGet`: []TaskOutput
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdTasksTaskIdOutputGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**taskId** | **int32** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksTaskIdOutputGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TaskOutput**](TaskOutput.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTemplatesGet

> []Template ProjectProjectIdTemplatesGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get template

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
    sort := "sort_example" // string | sorting name
    order := "order_example" // string | ordering manner

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTemplatesGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdTemplatesGet`: []Template
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdTemplatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]Template**](Template.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTemplatesPost

> Template ProjectProjectIdTemplatesPost(ctx, projectId).Template(template).Execute()

create template

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
    template := *openapiclient.NewTemplateRequest() // TemplateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTemplatesPost(context.Background(), projectId).Template(template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdTemplatesPost`: Template
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdTemplatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **template** | [**TemplateRequest**](TemplateRequest.md) |  | 

### Return type

[**Template**](Template.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTemplatesTemplateIdDelete

> ProjectProjectIdTemplatesTemplateIdDelete(ctx, projectId, templateId).Execute()

Removes template

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
    templateId := int32(7) // int32 | template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTemplatesTemplateIdDelete(context.Background(), projectId, templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTemplatesTemplateIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**templateId** | **int32** | template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesTemplateIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdTemplatesTemplateIdGet

> Template ProjectProjectIdTemplatesTemplateIdGet(ctx, projectId, templateId).Execute()

Get template

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
    templateId := int32(7) // int32 | template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTemplatesTemplateIdGet(context.Background(), projectId, templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTemplatesTemplateIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdTemplatesTemplateIdGet`: Template
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdTemplatesTemplateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**templateId** | **int32** | template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesTemplateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Template**](Template.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTemplatesTemplateIdPut

> ProjectProjectIdTemplatesTemplateIdPut(ctx, projectId, templateId).Template(template).Execute()

Updates template

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
    templateId := int32(7) // int32 | template ID
    template := *openapiclient.NewTemplateRequest() // TemplateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdTemplatesTemplateIdPut(context.Background(), projectId, templateId).Template(template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdTemplatesTemplateIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**templateId** | **int32** | template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesTemplateIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **template** | [**TemplateRequest**](TemplateRequest.md) |  | 

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


## ProjectProjectIdUsersGet

> []User ProjectProjectIdUsersGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get users linked to project

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
    sort := "email" // string | sorting name
    order := "desc" // string | ordering manner

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdUsersGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdUsersGet`: []User
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]User**](User.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdUsersPost

> ProjectProjectIdUsersPost(ctx, projectId).User(user).Execute()

Link user to project

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
    user := *openapiclient.NewProjectProjectIdUsersGetRequest() // ProjectProjectIdUsersGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdUsersPost(context.Background(), projectId).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdUsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**ProjectProjectIdUsersGetRequest**](ProjectProjectIdUsersGetRequest.md) |  | 

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


## ProjectProjectIdUsersUserIdAdminDelete

> ProjectProjectIdUsersUserIdAdminDelete(ctx, projectId, userId).Execute()

Revoke admin privileges

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
    userId := int32(2) // int32 | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdUsersUserIdAdminDelete(context.Background(), projectId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdUsersUserIdAdminDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdUsersUserIdAdminDeleteRequest struct via the builder pattern


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


## ProjectProjectIdUsersUserIdAdminPost

> ProjectProjectIdUsersUserIdAdminPost(ctx, projectId, userId).Execute()

Makes user admin

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
    userId := int32(2) // int32 | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdUsersUserIdAdminPost(context.Background(), projectId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdUsersUserIdAdminPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdUsersUserIdAdminPostRequest struct via the builder pattern


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


## ProjectProjectIdUsersUserIdDelete

> ProjectProjectIdUsersUserIdDelete(ctx, projectId, userId).Execute()

Removes user from project

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
    userId := int32(2) // int32 | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdUsersUserIdDelete(context.Background(), projectId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdUsersUserIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdUsersUserIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdViewsGet

> []View ProjectProjectIdViewsGet(ctx, projectId).Execute()

Get view

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdViewsGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdViewsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdViewsGet`: []View
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdViewsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]View**](View.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdViewsPost

> View ProjectProjectIdViewsPost(ctx, projectId).View(view).Execute()

create view

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
    view := *openapiclient.NewViewRequest() // ViewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdViewsPost(context.Background(), projectId).View(view).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdViewsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdViewsPost`: View
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdViewsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | [**ViewRequest**](ViewRequest.md) |  | 

### Return type

[**View**](View.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdViewsViewIdDelete

> ProjectProjectIdViewsViewIdDelete(ctx, projectId, viewId).Execute()

Removes view

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
    viewId := int32(10) // int32 | view ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdViewsViewIdDelete(context.Background(), projectId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdViewsViewIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**viewId** | **int32** | view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsViewIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdViewsViewIdGet

> View ProjectProjectIdViewsViewIdGet(ctx, projectId, viewId).Execute()

Get view

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
    viewId := int32(10) // int32 | view ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdViewsViewIdGet(context.Background(), projectId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdViewsViewIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectProjectIdViewsViewIdGet`: View
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ProjectProjectIdViewsViewIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**viewId** | **int32** | view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsViewIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**View**](View.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdViewsViewIdPut

> ProjectProjectIdViewsViewIdPut(ctx, projectId, viewId).View(view).Execute()

Updates view

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
    viewId := int32(10) // int32 | view ID
    view := *openapiclient.NewViewRequest() // ViewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ProjectProjectIdViewsViewIdPut(context.Background(), projectId, viewId).View(view).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ProjectProjectIdViewsViewIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**viewId** | **int32** | view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsViewIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | [**ViewRequest**](ViewRequest.md) |  | 

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

