# ProjectsApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**projectsGet**](ProjectsApi.md#projectsGet) | **GET** /projects | Get projects |
| [**projectsPost**](ProjectsApi.md#projectsPost) | **POST** /projects | Create a new project |


<a name="projectsGet"></a>
# **projectsGet**
> List&lt;Project&gt; projectsGet()

Get projects

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectsApi;

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

    ProjectsApi apiInstance = new ProjectsApi(defaultClient);
    try {
      List<Project> result = apiInstance.projectsGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectsApi#projectsGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List&lt;Project&gt;**](Project.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | List of projects |  -  |

<a name="projectsPost"></a>
# **projectsPost**
> projectsPost(project)

Create a new project

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProjectsApi;

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

    ProjectsApi apiInstance = new ProjectsApi(defaultClient);
    ProjectRequest project = new ProjectRequest(); // ProjectRequest | 
    try {
      apiInstance.projectsPost(project);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProjectsApi#projectsPost");
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
| **project** | [**ProjectRequest**](ProjectRequest.md)|  | |

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
| **201** | Created project |  -  |

