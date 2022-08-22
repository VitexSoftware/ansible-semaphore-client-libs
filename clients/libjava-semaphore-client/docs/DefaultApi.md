# DefaultApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**eventsGet**](DefaultApi.md#eventsGet) | **GET** /events | Get Events related to Semaphore and projects you are part of |
| [**eventsLastGet**](DefaultApi.md#eventsLastGet) | **GET** /events/last | Get last 200 Events related to Semaphore and projects you are part of |
| [**infoGet**](DefaultApi.md#infoGet) | **GET** /info | Fetches information about semaphore |
| [**pingGet**](DefaultApi.md#pingGet) | **GET** /ping | PING test |
| [**wsGet**](DefaultApi.md#wsGet) | **GET** /ws | Websocket handler |


<a name="eventsGet"></a>
# **eventsGet**
> List&lt;Event&gt; eventsGet()

Get Events related to Semaphore and projects you are part of

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

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

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      List<Event> result = apiInstance.eventsGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#eventsGet");
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

<a name="eventsLastGet"></a>
# **eventsLastGet**
> List&lt;Event&gt; eventsLastGet()

Get last 200 Events related to Semaphore and projects you are part of

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

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

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      List<Event> result = apiInstance.eventsLastGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#eventsLastGet");
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

<a name="infoGet"></a>
# **infoGet**
> InfoType infoGet()

Fetches information about semaphore

you must be authenticated to use this

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

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

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      InfoType result = apiInstance.infoGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#infoGet");
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

[**InfoType**](InfoType.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | ok |  -  |

<a name="pingGet"></a>
# **pingGet**
> String pingGet()

PING test

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

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

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      String result = apiInstance.pingGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#pingGet");
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

**String**

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful \&quot;PONG\&quot; reply |  * content-type -  <br>  |

<a name="wsGet"></a>
# **wsGet**
> wsGet()

Websocket handler

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

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

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      apiInstance.wsGet();
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#wsGet");
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

null (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **401** | not authenticated |  -  |

