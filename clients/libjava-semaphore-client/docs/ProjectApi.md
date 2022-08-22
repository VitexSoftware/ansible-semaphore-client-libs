# ProjectApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**projectProjectIdDelete**](ProjectApi.md#projectProjectIdDelete) | **DELETE** /project/{project_id}/ | Delete project |
| [**projectProjectIdEnvironmentEnvironmentIdDelete**](ProjectApi.md#projectProjectIdEnvironmentEnvironmentIdDelete) | **DELETE** /project/{project_id}/environment/{environment_id} | Removes environment |
| [**projectProjectIdEnvironmentEnvironmentIdPut**](ProjectApi.md#projectProjectIdEnvironmentEnvironmentIdPut) | **PUT** /project/{project_id}/environment/{environment_id} | Update environment |
| [**projectProjectIdEnvironmentGet**](ProjectApi.md#projectProjectIdEnvironmentGet) | **GET** /project/{project_id}/environment | Get environment |
| [**projectProjectIdEnvironmentPost**](ProjectApi.md#projectProjectIdEnvironmentPost) | **POST** /project/{project_id}/environment | Add environment |
| [**projectProjectIdEventsGet**](ProjectApi.md#projectProjectIdEventsGet) | **GET** /project/{project_id}/events | Get Events related to this project |
| [**projectProjectIdGet**](ProjectApi.md#projectProjectIdGet) | **GET** /project/{project_id}/ | Fetch project |
| [**projectProjectIdInventoryGet**](ProjectApi.md#projectProjectIdInventoryGet) | **GET** /project/{project_id}/inventory | Get inventory |
| [**projectProjectIdInventoryInventoryIdDelete**](ProjectApi.md#projectProjectIdInventoryInventoryIdDelete) | **DELETE** /project/{project_id}/inventory/{inventory_id} | Removes inventory |
| [**projectProjectIdInventoryInventoryIdPut**](ProjectApi.md#projectProjectIdInventoryInventoryIdPut) | **PUT** /project/{project_id}/inventory/{inventory_id} | Updates inventory |
| [**projectProjectIdInventoryPost**](ProjectApi.md#projectProjectIdInventoryPost) | **POST** /project/{project_id}/inventory | create inventory |
| [**projectProjectIdKeysGet**](ProjectApi.md#projectProjectIdKeysGet) | **GET** /project/{project_id}/keys | Get access keys linked to project |
| [**projectProjectIdKeysKeyIdDelete**](ProjectApi.md#projectProjectIdKeysKeyIdDelete) | **DELETE** /project/{project_id}/keys/{key_id} | Removes access key |
| [**projectProjectIdKeysKeyIdPut**](ProjectApi.md#projectProjectIdKeysKeyIdPut) | **PUT** /project/{project_id}/keys/{key_id} | Updates access key |
| [**projectProjectIdKeysPost**](ProjectApi.md#projectProjectIdKeysPost) | **POST** /project/{project_id}/keys | Add access key |
| [**projectProjectIdPut**](ProjectApi.md#projectProjectIdPut) | **PUT** /project/{project_id}/ | Update project |
| [**projectProjectIdRepositoriesGet**](ProjectApi.md#projectProjectIdRepositoriesGet) | **GET** /project/{project_id}/repositories | Get repositories |
| [**projectProjectIdRepositoriesPost**](ProjectApi.md#projectProjectIdRepositoriesPost) | **POST** /project/{project_id}/repositories | Add repository |
| [**projectProjectIdRepositoriesRepositoryIdDelete**](ProjectApi.md#projectProjectIdRepositoriesRepositoryIdDelete) | **DELETE** /project/{project_id}/repositories/{repository_id} | Removes repository |
| [**projectProjectIdTasksGet**](ProjectApi.md#projectProjectIdTasksGet) | **GET** /project/{project_id}/tasks | Get Tasks related to current project |
| [**projectProjectIdTasksLastGet**](ProjectApi.md#projectProjectIdTasksLastGet) | **GET** /project/{project_id}/tasks/last | Get last 200 Tasks related to current project |
| [**projectProjectIdTasksPost**](ProjectApi.md#projectProjectIdTasksPost) | **POST** /project/{project_id}/tasks | Starts a job |
| [**projectProjectIdTasksTaskIdDelete**](ProjectApi.md#projectProjectIdTasksTaskIdDelete) | **DELETE** /project/{project_id}/tasks/{task_id} | Deletes task (including output) |
| [**projectProjectIdTasksTaskIdGet**](ProjectApi.md#projectProjectIdTasksTaskIdGet) | **GET** /project/{project_id}/tasks/{task_id} | Get a single task |
| [**projectProjectIdTasksTaskIdOutputGet**](ProjectApi.md#projectProjectIdTasksTaskIdOutputGet) | **GET** /project/{project_id}/tasks/{task_id}/output | Get task output |
| [**projectProjectIdTemplatesGet**](ProjectApi.md#projectProjectIdTemplatesGet) | **GET** /project/{project_id}/templates | Get template |
| [**projectProjectIdTemplatesPost**](ProjectApi.md#projectProjectIdTemplatesPost) | **POST** /project/{project_id}/templates | create template |
| [**projectProjectIdTemplatesTemplateIdDelete**](ProjectApi.md#projectProjectIdTemplatesTemplateIdDelete) | **DELETE** /project/{project_id}/templates/{template_id} | Removes template |
| [**projectProjectIdTemplatesTemplateIdGet**](ProjectApi.md#projectProjectIdTemplatesTemplateIdGet) | **GET** /project/{project_id}/templates/{template_id} | Get template |
| [**projectProjectIdTemplatesTemplateIdPut**](ProjectApi.md#projectProjectIdTemplatesTemplateIdPut) | **PUT** /project/{project_id}/templates/{template_id} | Updates template |
| [**projectProjectIdUsersGet**](ProjectApi.md#projectProjectIdUsersGet) | **GET** /project/{project_id}/users | Get users linked to project |
| [**projectProjectIdUsersPost**](ProjectApi.md#projectProjectIdUsersPost) | **POST** /project/{project_id}/users | Link user to project |
| [**projectProjectIdUsersUserIdAdminDelete**](ProjectApi.md#projectProjectIdUsersUserIdAdminDelete) | **DELETE** /project/{project_id}/users/{user_id}/admin | Revoke admin privileges |
| [**projectProjectIdUsersUserIdAdminPost**](ProjectApi.md#projectProjectIdUsersUserIdAdminPost) | **POST** /project/{project_id}/users/{user_id}/admin | Makes user admin |
| [**projectProjectIdUsersUserIdDelete**](ProjectApi.md#projectProjectIdUsersUserIdDelete) | **DELETE** /project/{project_id}/users/{user_id} | Removes user from project |
| [**projectProjectIdViewsGet**](ProjectApi.md#projectProjectIdViewsGet) | **GET** /project/{project_id}/views | Get view |
| [**projectProjectIdViewsPost**](ProjectApi.md#projectProjectIdViewsPost) | **POST** /project/{project_id}/views | create view |
| [**projectProjectIdViewsViewIdDelete**](ProjectApi.md#projectProjectIdViewsViewIdDelete) | **DELETE** /project/{project_id}/views/{view_id} | Removes view |
| [**projectProjectIdViewsViewIdGet**](ProjectApi.md#projectProjectIdViewsViewIdGet) | **GET** /project/{project_id}/views/{view_id} | Get view |
| [**projectProjectIdViewsViewIdPut**](ProjectApi.md#projectProjectIdViewsViewIdPut) | **PUT** /project/{project_id}/views/{view_id} | Updates view |


<a name="projectProjectIdDelete"></a>
# **projectProjectIdDelete**
> projectProjectIdDelete(projectId)

Delete project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    try {
      apiInstance.projectProjectIdDelete(projectId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Project deleted |  -  |

<a name="projectProjectIdEnvironmentEnvironmentIdDelete"></a>
# **projectProjectIdEnvironmentEnvironmentIdDelete**
> projectProjectIdEnvironmentEnvironmentIdDelete(projectId, environmentId)

Removes environment

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer environmentId = 6; // Integer | environment ID
    try {
      apiInstance.projectProjectIdEnvironmentEnvironmentIdDelete(projectId, environmentId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdEnvironmentEnvironmentIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **environmentId** | **Integer**| environment ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | environment removed |  -  |

<a name="projectProjectIdEnvironmentEnvironmentIdPut"></a>
# **projectProjectIdEnvironmentEnvironmentIdPut**
> projectProjectIdEnvironmentEnvironmentIdPut(projectId, environmentId, environment)

Update environment

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer environmentId = 6; // Integer | environment ID
    EnvironmentRequest environment = new EnvironmentRequest(); // EnvironmentRequest | 
    try {
      apiInstance.projectProjectIdEnvironmentEnvironmentIdPut(projectId, environmentId, environment);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdEnvironmentEnvironmentIdPut");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **environmentId** | **Integer**| environment ID | |
| **environment** | [**EnvironmentRequest**](EnvironmentRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Environment Updated |  -  |

<a name="projectProjectIdEnvironmentGet"></a>
# **projectProjectIdEnvironmentGet**
> List&lt;Environment&gt; projectProjectIdEnvironmentGet(projectId, sort, order)

Get environment

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    String sort = "db-deploy"; // String | sorting name
    String order = "desc"; // String | ordering manner
    try {
      List<Environment> result = apiInstance.projectProjectIdEnvironmentGet(projectId, sort, order);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdEnvironmentGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **sort** | **String**| sorting name | |
| **order** | **String**| ordering manner | |

### Return type

[**List&lt;Environment&gt;**](Environment.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | environment |  -  |

<a name="projectProjectIdEnvironmentPost"></a>
# **projectProjectIdEnvironmentPost**
> projectProjectIdEnvironmentPost(projectId, environment)

Add environment

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    EnvironmentRequest environment = new EnvironmentRequest(); // EnvironmentRequest | 
    try {
      apiInstance.projectProjectIdEnvironmentPost(projectId, environment);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdEnvironmentPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **environment** | [**EnvironmentRequest**](EnvironmentRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Environment created |  -  |

<a name="projectProjectIdEventsGet"></a>
# **projectProjectIdEventsGet**
> List&lt;Event&gt; projectProjectIdEventsGet(projectId)

Get Events related to this project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    try {
      List<Event> result = apiInstance.projectProjectIdEventsGet(projectId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdEventsGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |

### Return type

[**List&lt;Event&gt;**](Event.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of events in chronological order |  -  |

<a name="projectProjectIdGet"></a>
# **projectProjectIdGet**
> Project projectProjectIdGet(projectId)

Fetch project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    try {
      Project result = apiInstance.projectProjectIdGet(projectId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |

### Return type

[**Project**](Project.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Project |  -  |

<a name="projectProjectIdInventoryGet"></a>
# **projectProjectIdInventoryGet**
> List&lt;Inventory&gt; projectProjectIdInventoryGet(projectId, sort, order)

Get inventory

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    String sort = "name"; // String | sorting name
    String order = "asc"; // String | ordering manner
    try {
      List<Inventory> result = apiInstance.projectProjectIdInventoryGet(projectId, sort, order);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdInventoryGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **sort** | **String**| sorting name | [enum: name, type] |
| **order** | **String**| ordering manner | [enum: asc, desc] |

### Return type

[**List&lt;Inventory&gt;**](Inventory.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | inventory |  -  |

<a name="projectProjectIdInventoryInventoryIdDelete"></a>
# **projectProjectIdInventoryInventoryIdDelete**
> projectProjectIdInventoryInventoryIdDelete(projectId, inventoryId)

Removes inventory

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer inventoryId = 5; // Integer | inventory ID
    try {
      apiInstance.projectProjectIdInventoryInventoryIdDelete(projectId, inventoryId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdInventoryInventoryIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **inventoryId** | **Integer**| inventory ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | inventory removed |  -  |

<a name="projectProjectIdInventoryInventoryIdPut"></a>
# **projectProjectIdInventoryInventoryIdPut**
> projectProjectIdInventoryInventoryIdPut(projectId, inventoryId, inventory)

Updates inventory

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer inventoryId = 5; // Integer | inventory ID
    InventoryRequest inventory = new InventoryRequest(); // InventoryRequest | 
    try {
      apiInstance.projectProjectIdInventoryInventoryIdPut(projectId, inventoryId, inventory);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdInventoryInventoryIdPut");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **inventoryId** | **Integer**| inventory ID | |
| **inventory** | [**InventoryRequest**](InventoryRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Inventory updated |  -  |

<a name="projectProjectIdInventoryPost"></a>
# **projectProjectIdInventoryPost**
> Inventory projectProjectIdInventoryPost(projectId, inventory)

create inventory

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    InventoryRequest inventory = new InventoryRequest(); // InventoryRequest | 
    try {
      Inventory result = apiInstance.projectProjectIdInventoryPost(projectId, inventory);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdInventoryPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **inventory** | [**InventoryRequest**](InventoryRequest.md)|  | |

### Return type

[**Inventory**](Inventory.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | inventory created |  -  |

<a name="projectProjectIdKeysGet"></a>
# **projectProjectIdKeysGet**
> List&lt;AccessKey&gt; projectProjectIdKeysGet(projectId, sort, order, keyType)

Get access keys linked to project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    String sort = "name"; // String | sorting name
    String order = "asc"; // String | ordering manner
    String keyType = "none"; // String | Filter by key type
    try {
      List<AccessKey> result = apiInstance.projectProjectIdKeysGet(projectId, sort, order, keyType);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdKeysGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **sort** | **String**| sorting name | [enum: name, type] |
| **order** | **String**| ordering manner | [enum: asc, desc] |
| **keyType** | **String**| Filter by key type | [optional] [enum: none, ssh, login_password] |

### Return type

[**List&lt;AccessKey&gt;**](AccessKey.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Access Keys |  -  |

<a name="projectProjectIdKeysKeyIdDelete"></a>
# **projectProjectIdKeysKeyIdDelete**
> projectProjectIdKeysKeyIdDelete(projectId, keyId)

Removes access key

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer keyId = 3; // Integer | key ID
    try {
      apiInstance.projectProjectIdKeysKeyIdDelete(projectId, keyId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdKeysKeyIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **keyId** | **Integer**| key ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | access key removed |  -  |

<a name="projectProjectIdKeysKeyIdPut"></a>
# **projectProjectIdKeysKeyIdPut**
> projectProjectIdKeysKeyIdPut(projectId, keyId, accessKey)

Updates access key

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer keyId = 3; // Integer | key ID
    AccessKeyRequest accessKey = new AccessKeyRequest(); // AccessKeyRequest | 
    try {
      apiInstance.projectProjectIdKeysKeyIdPut(projectId, keyId, accessKey);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdKeysKeyIdPut");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **keyId** | **Integer**| key ID | |
| **accessKey** | [**AccessKeyRequest**](AccessKeyRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Key updated |  -  |
| **400** | Bad type |  -  |

<a name="projectProjectIdKeysPost"></a>
# **projectProjectIdKeysPost**
> projectProjectIdKeysPost(projectId, accessKey)

Add access key

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    AccessKeyRequest accessKey = new AccessKeyRequest(); // AccessKeyRequest | 
    try {
      apiInstance.projectProjectIdKeysPost(projectId, accessKey);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdKeysPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **accessKey** | [**AccessKeyRequest**](AccessKeyRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Access Key created |  -  |
| **400** | Bad type |  -  |

<a name="projectProjectIdPut"></a>
# **projectProjectIdPut**
> projectProjectIdPut(projectId, project)

Update project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    ProjectProjectIdDeleteRequest project = new ProjectProjectIdDeleteRequest(); // ProjectProjectIdDeleteRequest | 
    try {
      apiInstance.projectProjectIdPut(projectId, project);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdPut");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **project** | [**ProjectProjectIdDeleteRequest**](ProjectProjectIdDeleteRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Project saved |  -  |

<a name="projectProjectIdRepositoriesGet"></a>
# **projectProjectIdRepositoriesGet**
> List&lt;Repository&gt; projectProjectIdRepositoriesGet(projectId, sort, order)

Get repositories

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    String sort = "name"; // String | sorting name
    String order = "asc"; // String | ordering manner
    try {
      List<Repository> result = apiInstance.projectProjectIdRepositoriesGet(projectId, sort, order);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdRepositoriesGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **sort** | **String**| sorting name | [enum: name, git_url, ssh_key] |
| **order** | **String**| ordering manner | [enum: asc, desc] |

### Return type

[**List&lt;Repository&gt;**](Repository.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | repositories |  -  |

<a name="projectProjectIdRepositoriesPost"></a>
# **projectProjectIdRepositoriesPost**
> projectProjectIdRepositoriesPost(projectId, repository)

Add repository

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    RepositoryRequest repository = new RepositoryRequest(); // RepositoryRequest | 
    try {
      apiInstance.projectProjectIdRepositoriesPost(projectId, repository);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdRepositoriesPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **repository** | [**RepositoryRequest**](RepositoryRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Repository created |  -  |

<a name="projectProjectIdRepositoriesRepositoryIdDelete"></a>
# **projectProjectIdRepositoriesRepositoryIdDelete**
> projectProjectIdRepositoriesRepositoryIdDelete(projectId, repositoryId)

Removes repository

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer repositoryId = 4; // Integer | repository ID
    try {
      apiInstance.projectProjectIdRepositoriesRepositoryIdDelete(projectId, repositoryId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdRepositoriesRepositoryIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **repositoryId** | **Integer**| repository ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | repository removed |  -  |

<a name="projectProjectIdTasksGet"></a>
# **projectProjectIdTasksGet**
> List&lt;Task&gt; projectProjectIdTasksGet(projectId)

Get Tasks related to current project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    try {
      List<Task> result = apiInstance.projectProjectIdTasksGet(projectId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTasksGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |

### Return type

[**List&lt;Task&gt;**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of tasks in chronological order |  -  |

<a name="projectProjectIdTasksLastGet"></a>
# **projectProjectIdTasksLastGet**
> List&lt;Task&gt; projectProjectIdTasksLastGet(projectId)

Get last 200 Tasks related to current project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    try {
      List<Task> result = apiInstance.projectProjectIdTasksLastGet(projectId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTasksLastGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |

### Return type

[**List&lt;Task&gt;**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of tasks in chronological order |  -  |

<a name="projectProjectIdTasksPost"></a>
# **projectProjectIdTasksPost**
> Task projectProjectIdTasksPost(projectId, task)

Starts a job

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    ProjectProjectIdTasksGetRequest task = new ProjectProjectIdTasksGetRequest(); // ProjectProjectIdTasksGetRequest | 
    try {
      Task result = apiInstance.projectProjectIdTasksPost(projectId, task);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTasksPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **task** | [**ProjectProjectIdTasksGetRequest**](ProjectProjectIdTasksGetRequest.md)|  | |

### Return type

[**Task**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Task queued |  -  |

<a name="projectProjectIdTasksTaskIdDelete"></a>
# **projectProjectIdTasksTaskIdDelete**
> projectProjectIdTasksTaskIdDelete(projectId, taskId)

Deletes task (including output)

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer taskId = 8; // Integer | task ID
    try {
      apiInstance.projectProjectIdTasksTaskIdDelete(projectId, taskId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTasksTaskIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **taskId** | **Integer**| task ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | task deleted |  -  |

<a name="projectProjectIdTasksTaskIdGet"></a>
# **projectProjectIdTasksTaskIdGet**
> Task projectProjectIdTasksTaskIdGet(projectId, taskId)

Get a single task

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer taskId = 8; // Integer | task ID
    try {
      Task result = apiInstance.projectProjectIdTasksTaskIdGet(projectId, taskId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTasksTaskIdGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **taskId** | **Integer**| task ID | |

### Return type

[**Task**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Task |  -  |

<a name="projectProjectIdTasksTaskIdOutputGet"></a>
# **projectProjectIdTasksTaskIdOutputGet**
> List&lt;TaskOutput&gt; projectProjectIdTasksTaskIdOutputGet(projectId, taskId)

Get task output

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer taskId = 8; // Integer | task ID
    try {
      List<TaskOutput> result = apiInstance.projectProjectIdTasksTaskIdOutputGet(projectId, taskId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTasksTaskIdOutputGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **taskId** | **Integer**| task ID | |

### Return type

[**List&lt;TaskOutput&gt;**](TaskOutput.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | output |  -  |

<a name="projectProjectIdTemplatesGet"></a>
# **projectProjectIdTemplatesGet**
> List&lt;Template&gt; projectProjectIdTemplatesGet(projectId, sort, order)

Get template

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    String sort = "alias"; // String | sorting name
    String order = "asc"; // String | ordering manner
    try {
      List<Template> result = apiInstance.projectProjectIdTemplatesGet(projectId, sort, order);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTemplatesGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **sort** | **String**| sorting name | [enum: alias, playbook, ssh_key, inventory, environment, repository] |
| **order** | **String**| ordering manner | [enum: asc, desc] |

### Return type

[**List&lt;Template&gt;**](Template.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | template |  -  |

<a name="projectProjectIdTemplatesPost"></a>
# **projectProjectIdTemplatesPost**
> Template projectProjectIdTemplatesPost(projectId, template)

create template

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    TemplateRequest template = new TemplateRequest(); // TemplateRequest | 
    try {
      Template result = apiInstance.projectProjectIdTemplatesPost(projectId, template);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTemplatesPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **template** | [**TemplateRequest**](TemplateRequest.md)|  | |

### Return type

[**Template**](Template.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | template created |  -  |

<a name="projectProjectIdTemplatesTemplateIdDelete"></a>
# **projectProjectIdTemplatesTemplateIdDelete**
> projectProjectIdTemplatesTemplateIdDelete(projectId, templateId)

Removes template

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer templateId = 7; // Integer | template ID
    try {
      apiInstance.projectProjectIdTemplatesTemplateIdDelete(projectId, templateId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTemplatesTemplateIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **templateId** | **Integer**| template ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | template removed |  -  |

<a name="projectProjectIdTemplatesTemplateIdGet"></a>
# **projectProjectIdTemplatesTemplateIdGet**
> Template projectProjectIdTemplatesTemplateIdGet(projectId, templateId)

Get template

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer templateId = 7; // Integer | template ID
    try {
      Template result = apiInstance.projectProjectIdTemplatesTemplateIdGet(projectId, templateId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTemplatesTemplateIdGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **templateId** | **Integer**| template ID | |

### Return type

[**Template**](Template.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | template object |  -  |

<a name="projectProjectIdTemplatesTemplateIdPut"></a>
# **projectProjectIdTemplatesTemplateIdPut**
> projectProjectIdTemplatesTemplateIdPut(projectId, templateId, template)

Updates template

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer templateId = 7; // Integer | template ID
    TemplateRequest template = new TemplateRequest(); // TemplateRequest | 
    try {
      apiInstance.projectProjectIdTemplatesTemplateIdPut(projectId, templateId, template);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdTemplatesTemplateIdPut");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **templateId** | **Integer**| template ID | |
| **template** | [**TemplateRequest**](TemplateRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | template updated |  -  |

<a name="projectProjectIdUsersGet"></a>
# **projectProjectIdUsersGet**
> List&lt;User&gt; projectProjectIdUsersGet(projectId, sort, order)

Get users linked to project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    String sort = "name"; // String | sorting name
    String order = "asc"; // String | ordering manner
    try {
      List<User> result = apiInstance.projectProjectIdUsersGet(projectId, sort, order);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdUsersGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **sort** | **String**| sorting name | [enum: name, username, email, admin] |
| **order** | **String**| ordering manner | [enum: asc, desc] |

### Return type

[**List&lt;User&gt;**](User.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Users |  -  |

<a name="projectProjectIdUsersPost"></a>
# **projectProjectIdUsersPost**
> projectProjectIdUsersPost(projectId, user)

Link user to project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    ProjectProjectIdUsersGetRequest user = new ProjectProjectIdUsersGetRequest(); // ProjectProjectIdUsersGetRequest | 
    try {
      apiInstance.projectProjectIdUsersPost(projectId, user);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdUsersPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **user** | [**ProjectProjectIdUsersGetRequest**](ProjectProjectIdUsersGetRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | User added |  -  |

<a name="projectProjectIdUsersUserIdAdminDelete"></a>
# **projectProjectIdUsersUserIdAdminDelete**
> projectProjectIdUsersUserIdAdminDelete(projectId, userId)

Revoke admin privileges

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer userId = 2; // Integer | User ID
    try {
      apiInstance.projectProjectIdUsersUserIdAdminDelete(projectId, userId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdUsersUserIdAdminDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **userId** | **Integer**| User ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | User admin privileges revoked |  -  |

<a name="projectProjectIdUsersUserIdAdminPost"></a>
# **projectProjectIdUsersUserIdAdminPost**
> projectProjectIdUsersUserIdAdminPost(projectId, userId)

Makes user admin

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer userId = 2; // Integer | User ID
    try {
      apiInstance.projectProjectIdUsersUserIdAdminPost(projectId, userId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdUsersUserIdAdminPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **userId** | **Integer**| User ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | User made administrator |  -  |

<a name="projectProjectIdUsersUserIdDelete"></a>
# **projectProjectIdUsersUserIdDelete**
> projectProjectIdUsersUserIdDelete(projectId, userId)

Removes user from project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer userId = 2; // Integer | User ID
    try {
      apiInstance.projectProjectIdUsersUserIdDelete(projectId, userId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdUsersUserIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **userId** | **Integer**| User ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | User removed |  -  |

<a name="projectProjectIdViewsGet"></a>
# **projectProjectIdViewsGet**
> List&lt;View&gt; projectProjectIdViewsGet(projectId)

Get view

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    try {
      List<View> result = apiInstance.projectProjectIdViewsGet(projectId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdViewsGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |

### Return type

[**List&lt;View&gt;**](View.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | view |  -  |

<a name="projectProjectIdViewsPost"></a>
# **projectProjectIdViewsPost**
> View projectProjectIdViewsPost(projectId, view)

create view

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    ViewRequest view = new ViewRequest(); // ViewRequest | 
    try {
      View result = apiInstance.projectProjectIdViewsPost(projectId, view);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdViewsPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **view** | [**ViewRequest**](ViewRequest.md)|  | |

### Return type

[**View**](View.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | view created |  -  |

<a name="projectProjectIdViewsViewIdDelete"></a>
# **projectProjectIdViewsViewIdDelete**
> projectProjectIdViewsViewIdDelete(projectId, viewId)

Removes view

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer viewId = 10; // Integer | view ID
    try {
      apiInstance.projectProjectIdViewsViewIdDelete(projectId, viewId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdViewsViewIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **viewId** | **Integer**| view ID | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | view removed |  -  |

<a name="projectProjectIdViewsViewIdGet"></a>
# **projectProjectIdViewsViewIdGet**
> View projectProjectIdViewsViewIdGet(projectId, viewId)

Get view

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer viewId = 10; // Integer | view ID
    try {
      View result = apiInstance.projectProjectIdViewsViewIdGet(projectId, viewId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdViewsViewIdGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **viewId** | **Integer**| view ID | |

### Return type

[**View**](View.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | view object |  -  |

<a name="projectProjectIdViewsViewIdPut"></a>
# **projectProjectIdViewsViewIdPut**
> projectProjectIdViewsViewIdPut(projectId, viewId, view)

Updates view

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    ProjectApi apiInstance = new ProjectApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer viewId = 10; // Integer | view ID
    ViewRequest view = new ViewRequest(); // ViewRequest | 
    try {
      apiInstance.projectProjectIdViewsViewIdPut(projectId, viewId, view);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectApi#projectProjectIdViewsViewIdPut");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **Integer**| Project ID | |
| **viewId** | **Integer**| view ID | |
| **view** | [**ViewRequest**](ViewRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | view updated |  -  |

