# ScheduleApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**projectProjectIdSchedulesPost**](ScheduleApi.md#projectProjectIdSchedulesPost) | **POST** /project/{project_id}/schedules | create schedule |
| [**projectProjectIdSchedulesScheduleIdDelete**](ScheduleApi.md#projectProjectIdSchedulesScheduleIdDelete) | **DELETE** /project/{project_id}/schedules/{schedule_id} | Deletes schedule |
| [**projectProjectIdSchedulesScheduleIdGet**](ScheduleApi.md#projectProjectIdSchedulesScheduleIdGet) | **GET** /project/{project_id}/schedules/{schedule_id} | Get schedule |
| [**projectProjectIdSchedulesScheduleIdPut**](ScheduleApi.md#projectProjectIdSchedulesScheduleIdPut) | **PUT** /project/{project_id}/schedules/{schedule_id} | Updates schedule |


<a name="projectProjectIdSchedulesPost"></a>
# **projectProjectIdSchedulesPost**
> Schedule projectProjectIdSchedulesPost(projectId, schedule)

create schedule

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ScheduleApi;

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

    ScheduleApi apiInstance = new ScheduleApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    ScheduleRequest schedule = new ScheduleRequest(); // ScheduleRequest | 
    try {
      Schedule result = apiInstance.projectProjectIdSchedulesPost(projectId, schedule);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ScheduleApi#projectProjectIdSchedulesPost");
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
| **schedule** | [**ScheduleRequest**](ScheduleRequest.md)|  | |

### Return type

[**Schedule**](Schedule.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | schedule created |  -  |

<a name="projectProjectIdSchedulesScheduleIdDelete"></a>
# **projectProjectIdSchedulesScheduleIdDelete**
> projectProjectIdSchedulesScheduleIdDelete(projectId, scheduleId)

Deletes schedule

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ScheduleApi;

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

    ScheduleApi apiInstance = new ScheduleApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer scheduleId = 9; // Integer | schedule ID
    try {
      apiInstance.projectProjectIdSchedulesScheduleIdDelete(projectId, scheduleId);
    } catch (ApiException e) {
      System.err.println("Exception when calling ScheduleApi#projectProjectIdSchedulesScheduleIdDelete");
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
| **scheduleId** | **Integer**| schedule ID | |

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
| **204** | schedule deleted |  -  |

<a name="projectProjectIdSchedulesScheduleIdGet"></a>
# **projectProjectIdSchedulesScheduleIdGet**
> Schedule projectProjectIdSchedulesScheduleIdGet(projectId, scheduleId)

Get schedule

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ScheduleApi;

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

    ScheduleApi apiInstance = new ScheduleApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer scheduleId = 9; // Integer | schedule ID
    try {
      Schedule result = apiInstance.projectProjectIdSchedulesScheduleIdGet(projectId, scheduleId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ScheduleApi#projectProjectIdSchedulesScheduleIdGet");
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
| **scheduleId** | **Integer**| schedule ID | |

### Return type

[**Schedule**](Schedule.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Schedule |  -  |

<a name="projectProjectIdSchedulesScheduleIdPut"></a>
# **projectProjectIdSchedulesScheduleIdPut**
> projectProjectIdSchedulesScheduleIdPut(projectId, scheduleId, schedule)

Updates schedule

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ScheduleApi;

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

    ScheduleApi apiInstance = new ScheduleApi(defaultClient);
    Integer projectId = 1; // Integer | Project ID
    Integer scheduleId = 9; // Integer | schedule ID
    ScheduleRequest schedule = new ScheduleRequest(); // ScheduleRequest | 
    try {
      apiInstance.projectProjectIdSchedulesScheduleIdPut(projectId, scheduleId, schedule);
    } catch (ApiException e) {
      System.err.println("Exception when calling ScheduleApi#projectProjectIdSchedulesScheduleIdPut");
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
| **scheduleId** | **Integer**| schedule ID | |
| **schedule** | [**ScheduleRequest**](ScheduleRequest.md)|  | |

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
| **204** | schedule updated |  -  |

