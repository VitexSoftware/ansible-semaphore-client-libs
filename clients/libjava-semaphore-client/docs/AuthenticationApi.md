# AuthenticationApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**authLoginPost**](AuthenticationApi.md#authLoginPost) | **POST** /auth/login | Performs Login |
| [**authLogoutPost**](AuthenticationApi.md#authLogoutPost) | **POST** /auth/logout | Destroys current session |
| [**userTokensApiTokenIdDelete**](AuthenticationApi.md#userTokensApiTokenIdDelete) | **DELETE** /user/tokens/{api_token_id} | Expires API token |
| [**userTokensGet**](AuthenticationApi.md#userTokensGet) | **GET** /user/tokens | Fetch API tokens for user |
| [**userTokensPost**](AuthenticationApi.md#userTokensPost) | **POST** /user/tokens | Create an API token |


<a name="authLoginPost"></a>
# **authLoginPost**
> authLoginPost(loginBody)

Performs Login

Upon success you will be logged in

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

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

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    Login loginBody = new Login(); // Login | 
    try {
      apiInstance.authLoginPost(loginBody);
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#authLoginPost");
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
| **loginBody** | [**Login**](Login.md)|  | |

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
| **204** | You are logged in |  -  |
| **400** | something in body is missing / is invalid |  -  |

<a name="authLogoutPost"></a>
# **authLogoutPost**
> authLogoutPost()

Destroys current session

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

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

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    try {
      apiInstance.authLogoutPost();
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#authLogoutPost");
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
| **204** | Your session was successfully nuked |  -  |

<a name="userTokensApiTokenIdDelete"></a>
# **userTokensApiTokenIdDelete**
> userTokensApiTokenIdDelete(apiTokenId)

Expires API token

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

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

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    String apiTokenId = "kwofd61g93-yuqvex8efmhjkgnbxlo8mp1tin6spyhu="; // String | 
    try {
      apiInstance.userTokensApiTokenIdDelete(apiTokenId);
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#userTokensApiTokenIdDelete");
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
| **apiTokenId** | **String**|  | |

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
| **204** | Expired API Token |  -  |

<a name="userTokensGet"></a>
# **userTokensGet**
> List&lt;APIToken&gt; userTokensGet()

Fetch API tokens for user

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

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

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    try {
      List<APIToken> result = apiInstance.userTokensGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#userTokensGet");
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

[**List&lt;APIToken&gt;**](APIToken.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | API Tokens |  -  |

<a name="userTokensPost"></a>
# **userTokensPost**
> APIToken userTokensPost()

Create an API token

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

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

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    try {
      APIToken result = apiInstance.userTokensPost();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#userTokensPost");
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

[**APIToken**](APIToken.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain; charset=utf-8

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | API Token |  -  |

