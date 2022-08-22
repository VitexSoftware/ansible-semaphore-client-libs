# OpenAPI\Client\ProjectApi

All URIs are relative to https://demo.ansible-semaphore.com/api.

Method | HTTP request | Description
------------- | ------------- | -------------
[**projectProjectIdDelete()**](ProjectApi.md#projectProjectIdDelete) | **DELETE** /project/{project_id}/ | Delete project
[**projectProjectIdEnvironmentEnvironmentIdDelete()**](ProjectApi.md#projectProjectIdEnvironmentEnvironmentIdDelete) | **DELETE** /project/{project_id}/environment/{environment_id} | Removes environment
[**projectProjectIdEnvironmentEnvironmentIdPut()**](ProjectApi.md#projectProjectIdEnvironmentEnvironmentIdPut) | **PUT** /project/{project_id}/environment/{environment_id} | Update environment
[**projectProjectIdEnvironmentGet()**](ProjectApi.md#projectProjectIdEnvironmentGet) | **GET** /project/{project_id}/environment | Get environment
[**projectProjectIdEnvironmentPost()**](ProjectApi.md#projectProjectIdEnvironmentPost) | **POST** /project/{project_id}/environment | Add environment
[**projectProjectIdEventsGet()**](ProjectApi.md#projectProjectIdEventsGet) | **GET** /project/{project_id}/events | Get Events related to this project
[**projectProjectIdGet()**](ProjectApi.md#projectProjectIdGet) | **GET** /project/{project_id}/ | Fetch project
[**projectProjectIdInventoryGet()**](ProjectApi.md#projectProjectIdInventoryGet) | **GET** /project/{project_id}/inventory | Get inventory
[**projectProjectIdInventoryInventoryIdDelete()**](ProjectApi.md#projectProjectIdInventoryInventoryIdDelete) | **DELETE** /project/{project_id}/inventory/{inventory_id} | Removes inventory
[**projectProjectIdInventoryInventoryIdPut()**](ProjectApi.md#projectProjectIdInventoryInventoryIdPut) | **PUT** /project/{project_id}/inventory/{inventory_id} | Updates inventory
[**projectProjectIdInventoryPost()**](ProjectApi.md#projectProjectIdInventoryPost) | **POST** /project/{project_id}/inventory | create inventory
[**projectProjectIdKeysGet()**](ProjectApi.md#projectProjectIdKeysGet) | **GET** /project/{project_id}/keys | Get access keys linked to project
[**projectProjectIdKeysKeyIdDelete()**](ProjectApi.md#projectProjectIdKeysKeyIdDelete) | **DELETE** /project/{project_id}/keys/{key_id} | Removes access key
[**projectProjectIdKeysKeyIdPut()**](ProjectApi.md#projectProjectIdKeysKeyIdPut) | **PUT** /project/{project_id}/keys/{key_id} | Updates access key
[**projectProjectIdKeysPost()**](ProjectApi.md#projectProjectIdKeysPost) | **POST** /project/{project_id}/keys | Add access key
[**projectProjectIdPut()**](ProjectApi.md#projectProjectIdPut) | **PUT** /project/{project_id}/ | Update project
[**projectProjectIdRepositoriesGet()**](ProjectApi.md#projectProjectIdRepositoriesGet) | **GET** /project/{project_id}/repositories | Get repositories
[**projectProjectIdRepositoriesPost()**](ProjectApi.md#projectProjectIdRepositoriesPost) | **POST** /project/{project_id}/repositories | Add repository
[**projectProjectIdRepositoriesRepositoryIdDelete()**](ProjectApi.md#projectProjectIdRepositoriesRepositoryIdDelete) | **DELETE** /project/{project_id}/repositories/{repository_id} | Removes repository
[**projectProjectIdTasksGet()**](ProjectApi.md#projectProjectIdTasksGet) | **GET** /project/{project_id}/tasks | Get Tasks related to current project
[**projectProjectIdTasksLastGet()**](ProjectApi.md#projectProjectIdTasksLastGet) | **GET** /project/{project_id}/tasks/last | Get last 200 Tasks related to current project
[**projectProjectIdTasksPost()**](ProjectApi.md#projectProjectIdTasksPost) | **POST** /project/{project_id}/tasks | Starts a job
[**projectProjectIdTasksTaskIdDelete()**](ProjectApi.md#projectProjectIdTasksTaskIdDelete) | **DELETE** /project/{project_id}/tasks/{task_id} | Deletes task (including output)
[**projectProjectIdTasksTaskIdGet()**](ProjectApi.md#projectProjectIdTasksTaskIdGet) | **GET** /project/{project_id}/tasks/{task_id} | Get a single task
[**projectProjectIdTasksTaskIdOutputGet()**](ProjectApi.md#projectProjectIdTasksTaskIdOutputGet) | **GET** /project/{project_id}/tasks/{task_id}/output | Get task output
[**projectProjectIdTemplatesGet()**](ProjectApi.md#projectProjectIdTemplatesGet) | **GET** /project/{project_id}/templates | Get template
[**projectProjectIdTemplatesPost()**](ProjectApi.md#projectProjectIdTemplatesPost) | **POST** /project/{project_id}/templates | create template
[**projectProjectIdTemplatesTemplateIdDelete()**](ProjectApi.md#projectProjectIdTemplatesTemplateIdDelete) | **DELETE** /project/{project_id}/templates/{template_id} | Removes template
[**projectProjectIdTemplatesTemplateIdGet()**](ProjectApi.md#projectProjectIdTemplatesTemplateIdGet) | **GET** /project/{project_id}/templates/{template_id} | Get template
[**projectProjectIdTemplatesTemplateIdPut()**](ProjectApi.md#projectProjectIdTemplatesTemplateIdPut) | **PUT** /project/{project_id}/templates/{template_id} | Updates template
[**projectProjectIdUsersGet()**](ProjectApi.md#projectProjectIdUsersGet) | **GET** /project/{project_id}/users | Get users linked to project
[**projectProjectIdUsersPost()**](ProjectApi.md#projectProjectIdUsersPost) | **POST** /project/{project_id}/users | Link user to project
[**projectProjectIdUsersUserIdAdminDelete()**](ProjectApi.md#projectProjectIdUsersUserIdAdminDelete) | **DELETE** /project/{project_id}/users/{user_id}/admin | Revoke admin privileges
[**projectProjectIdUsersUserIdAdminPost()**](ProjectApi.md#projectProjectIdUsersUserIdAdminPost) | **POST** /project/{project_id}/users/{user_id}/admin | Makes user admin
[**projectProjectIdUsersUserIdDelete()**](ProjectApi.md#projectProjectIdUsersUserIdDelete) | **DELETE** /project/{project_id}/users/{user_id} | Removes user from project
[**projectProjectIdViewsGet()**](ProjectApi.md#projectProjectIdViewsGet) | **GET** /project/{project_id}/views | Get view
[**projectProjectIdViewsPost()**](ProjectApi.md#projectProjectIdViewsPost) | **POST** /project/{project_id}/views | create view
[**projectProjectIdViewsViewIdDelete()**](ProjectApi.md#projectProjectIdViewsViewIdDelete) | **DELETE** /project/{project_id}/views/{view_id} | Removes view
[**projectProjectIdViewsViewIdGet()**](ProjectApi.md#projectProjectIdViewsViewIdGet) | **GET** /project/{project_id}/views/{view_id} | Get view
[**projectProjectIdViewsViewIdPut()**](ProjectApi.md#projectProjectIdViewsViewIdPut) | **PUT** /project/{project_id}/views/{view_id} | Updates view


## `projectProjectIdDelete()`

```php
projectProjectIdDelete($project_id)
```

Delete project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID

try {
    $apiInstance->projectProjectIdDelete($project_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdEnvironmentEnvironmentIdDelete()`

```php
projectProjectIdEnvironmentEnvironmentIdDelete($project_id, $environment_id)
```

Removes environment

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$environment_id = 6; // int | environment ID

try {
    $apiInstance->projectProjectIdEnvironmentEnvironmentIdDelete($project_id, $environment_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdEnvironmentEnvironmentIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **environment_id** | **int**| environment ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdEnvironmentEnvironmentIdPut()`

```php
projectProjectIdEnvironmentEnvironmentIdPut($project_id, $environment_id, $environment)
```

Update environment

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$environment_id = 6; // int | environment ID
$environment = new \OpenAPI\Client\Model\EnvironmentRequest(); // \OpenAPI\Client\Model\EnvironmentRequest

try {
    $apiInstance->projectProjectIdEnvironmentEnvironmentIdPut($project_id, $environment_id, $environment);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdEnvironmentEnvironmentIdPut: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **environment_id** | **int**| environment ID |
 **environment** | [**\OpenAPI\Client\Model\EnvironmentRequest**](../Model/EnvironmentRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdEnvironmentGet()`

```php
projectProjectIdEnvironmentGet($project_id, $sort, $order): \OpenAPI\Client\Model\Environment[]
```

Get environment

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$sort = db-deploy; // string | sorting name
$order = desc; // string | ordering manner

try {
    $result = $apiInstance->projectProjectIdEnvironmentGet($project_id, $sort, $order);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdEnvironmentGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **sort** | **string**| sorting name |
 **order** | **string**| ordering manner |

### Return type

[**\OpenAPI\Client\Model\Environment[]**](../Model/Environment.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdEnvironmentPost()`

```php
projectProjectIdEnvironmentPost($project_id, $environment)
```

Add environment

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$environment = new \OpenAPI\Client\Model\EnvironmentRequest(); // \OpenAPI\Client\Model\EnvironmentRequest

try {
    $apiInstance->projectProjectIdEnvironmentPost($project_id, $environment);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdEnvironmentPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **environment** | [**\OpenAPI\Client\Model\EnvironmentRequest**](../Model/EnvironmentRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdEventsGet()`

```php
projectProjectIdEventsGet($project_id): \OpenAPI\Client\Model\Event[]
```

Get Events related to this project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID

try {
    $result = $apiInstance->projectProjectIdEventsGet($project_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdEventsGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |

### Return type

[**\OpenAPI\Client\Model\Event[]**](../Model/Event.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdGet()`

```php
projectProjectIdGet($project_id): \OpenAPI\Client\Model\Project
```

Fetch project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID

try {
    $result = $apiInstance->projectProjectIdGet($project_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |

### Return type

[**\OpenAPI\Client\Model\Project**](../Model/Project.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdInventoryGet()`

```php
projectProjectIdInventoryGet($project_id, $sort, $order): \OpenAPI\Client\Model\Inventory[]
```

Get inventory

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$sort = 'sort_example'; // string | sorting name
$order = 'order_example'; // string | ordering manner

try {
    $result = $apiInstance->projectProjectIdInventoryGet($project_id, $sort, $order);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdInventoryGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **sort** | **string**| sorting name |
 **order** | **string**| ordering manner |

### Return type

[**\OpenAPI\Client\Model\Inventory[]**](../Model/Inventory.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdInventoryInventoryIdDelete()`

```php
projectProjectIdInventoryInventoryIdDelete($project_id, $inventory_id)
```

Removes inventory

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$inventory_id = 5; // int | inventory ID

try {
    $apiInstance->projectProjectIdInventoryInventoryIdDelete($project_id, $inventory_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdInventoryInventoryIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **inventory_id** | **int**| inventory ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdInventoryInventoryIdPut()`

```php
projectProjectIdInventoryInventoryIdPut($project_id, $inventory_id, $inventory)
```

Updates inventory

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$inventory_id = 5; // int | inventory ID
$inventory = new \OpenAPI\Client\Model\InventoryRequest(); // \OpenAPI\Client\Model\InventoryRequest

try {
    $apiInstance->projectProjectIdInventoryInventoryIdPut($project_id, $inventory_id, $inventory);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdInventoryInventoryIdPut: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **inventory_id** | **int**| inventory ID |
 **inventory** | [**\OpenAPI\Client\Model\InventoryRequest**](../Model/InventoryRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdInventoryPost()`

```php
projectProjectIdInventoryPost($project_id, $inventory): \OpenAPI\Client\Model\Inventory
```

create inventory

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$inventory = new \OpenAPI\Client\Model\InventoryRequest(); // \OpenAPI\Client\Model\InventoryRequest

try {
    $result = $apiInstance->projectProjectIdInventoryPost($project_id, $inventory);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdInventoryPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **inventory** | [**\OpenAPI\Client\Model\InventoryRequest**](../Model/InventoryRequest.md)|  |

### Return type

[**\OpenAPI\Client\Model\Inventory**](../Model/Inventory.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdKeysGet()`

```php
projectProjectIdKeysGet($project_id, $sort, $order, $key_type): \OpenAPI\Client\Model\AccessKey[]
```

Get access keys linked to project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$sort = type; // string | sorting name
$order = asc; // string | ordering manner
$key_type = none; // string | Filter by key type

try {
    $result = $apiInstance->projectProjectIdKeysGet($project_id, $sort, $order, $key_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdKeysGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **sort** | **string**| sorting name |
 **order** | **string**| ordering manner |
 **key_type** | **string**| Filter by key type | [optional]

### Return type

[**\OpenAPI\Client\Model\AccessKey[]**](../Model/AccessKey.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdKeysKeyIdDelete()`

```php
projectProjectIdKeysKeyIdDelete($project_id, $key_id)
```

Removes access key

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$key_id = 3; // int | key ID

try {
    $apiInstance->projectProjectIdKeysKeyIdDelete($project_id, $key_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdKeysKeyIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **key_id** | **int**| key ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdKeysKeyIdPut()`

```php
projectProjectIdKeysKeyIdPut($project_id, $key_id, $access_key)
```

Updates access key

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$key_id = 3; // int | key ID
$access_key = new \OpenAPI\Client\Model\AccessKeyRequest(); // \OpenAPI\Client\Model\AccessKeyRequest

try {
    $apiInstance->projectProjectIdKeysKeyIdPut($project_id, $key_id, $access_key);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdKeysKeyIdPut: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **key_id** | **int**| key ID |
 **access_key** | [**\OpenAPI\Client\Model\AccessKeyRequest**](../Model/AccessKeyRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdKeysPost()`

```php
projectProjectIdKeysPost($project_id, $access_key)
```

Add access key

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$access_key = new \OpenAPI\Client\Model\AccessKeyRequest(); // \OpenAPI\Client\Model\AccessKeyRequest

try {
    $apiInstance->projectProjectIdKeysPost($project_id, $access_key);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdKeysPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **access_key** | [**\OpenAPI\Client\Model\AccessKeyRequest**](../Model/AccessKeyRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdPut()`

```php
projectProjectIdPut($project_id, $project)
```

Update project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$project = new \OpenAPI\Client\Model\ProjectProjectIdDeleteRequest(); // \OpenAPI\Client\Model\ProjectProjectIdDeleteRequest

try {
    $apiInstance->projectProjectIdPut($project_id, $project);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdPut: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **project** | [**\OpenAPI\Client\Model\ProjectProjectIdDeleteRequest**](../Model/ProjectProjectIdDeleteRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdRepositoriesGet()`

```php
projectProjectIdRepositoriesGet($project_id, $sort, $order): \OpenAPI\Client\Model\Repository[]
```

Get repositories

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$sort = 'sort_example'; // string | sorting name
$order = 'order_example'; // string | ordering manner

try {
    $result = $apiInstance->projectProjectIdRepositoriesGet($project_id, $sort, $order);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdRepositoriesGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **sort** | **string**| sorting name |
 **order** | **string**| ordering manner |

### Return type

[**\OpenAPI\Client\Model\Repository[]**](../Model/Repository.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdRepositoriesPost()`

```php
projectProjectIdRepositoriesPost($project_id, $repository)
```

Add repository

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$repository = new \OpenAPI\Client\Model\RepositoryRequest(); // \OpenAPI\Client\Model\RepositoryRequest

try {
    $apiInstance->projectProjectIdRepositoriesPost($project_id, $repository);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdRepositoriesPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **repository** | [**\OpenAPI\Client\Model\RepositoryRequest**](../Model/RepositoryRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdRepositoriesRepositoryIdDelete()`

```php
projectProjectIdRepositoriesRepositoryIdDelete($project_id, $repository_id)
```

Removes repository

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$repository_id = 4; // int | repository ID

try {
    $apiInstance->projectProjectIdRepositoriesRepositoryIdDelete($project_id, $repository_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdRepositoriesRepositoryIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **repository_id** | **int**| repository ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTasksGet()`

```php
projectProjectIdTasksGet($project_id): \OpenAPI\Client\Model\Task[]
```

Get Tasks related to current project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID

try {
    $result = $apiInstance->projectProjectIdTasksGet($project_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTasksGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |

### Return type

[**\OpenAPI\Client\Model\Task[]**](../Model/Task.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTasksLastGet()`

```php
projectProjectIdTasksLastGet($project_id): \OpenAPI\Client\Model\Task[]
```

Get last 200 Tasks related to current project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID

try {
    $result = $apiInstance->projectProjectIdTasksLastGet($project_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTasksLastGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |

### Return type

[**\OpenAPI\Client\Model\Task[]**](../Model/Task.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTasksPost()`

```php
projectProjectIdTasksPost($project_id, $task): \OpenAPI\Client\Model\Task
```

Starts a job

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$task = new \OpenAPI\Client\Model\ProjectProjectIdTasksGetRequest(); // \OpenAPI\Client\Model\ProjectProjectIdTasksGetRequest

try {
    $result = $apiInstance->projectProjectIdTasksPost($project_id, $task);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTasksPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **task** | [**\OpenAPI\Client\Model\ProjectProjectIdTasksGetRequest**](../Model/ProjectProjectIdTasksGetRequest.md)|  |

### Return type

[**\OpenAPI\Client\Model\Task**](../Model/Task.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTasksTaskIdDelete()`

```php
projectProjectIdTasksTaskIdDelete($project_id, $task_id)
```

Deletes task (including output)

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$task_id = 8; // int | task ID

try {
    $apiInstance->projectProjectIdTasksTaskIdDelete($project_id, $task_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTasksTaskIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **task_id** | **int**| task ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTasksTaskIdGet()`

```php
projectProjectIdTasksTaskIdGet($project_id, $task_id): \OpenAPI\Client\Model\Task
```

Get a single task

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$task_id = 8; // int | task ID

try {
    $result = $apiInstance->projectProjectIdTasksTaskIdGet($project_id, $task_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTasksTaskIdGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **task_id** | **int**| task ID |

### Return type

[**\OpenAPI\Client\Model\Task**](../Model/Task.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTasksTaskIdOutputGet()`

```php
projectProjectIdTasksTaskIdOutputGet($project_id, $task_id): \OpenAPI\Client\Model\TaskOutput[]
```

Get task output

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$task_id = 8; // int | task ID

try {
    $result = $apiInstance->projectProjectIdTasksTaskIdOutputGet($project_id, $task_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTasksTaskIdOutputGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **task_id** | **int**| task ID |

### Return type

[**\OpenAPI\Client\Model\TaskOutput[]**](../Model/TaskOutput.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTemplatesGet()`

```php
projectProjectIdTemplatesGet($project_id, $sort, $order): \OpenAPI\Client\Model\Template[]
```

Get template

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$sort = 'sort_example'; // string | sorting name
$order = 'order_example'; // string | ordering manner

try {
    $result = $apiInstance->projectProjectIdTemplatesGet($project_id, $sort, $order);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTemplatesGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **sort** | **string**| sorting name |
 **order** | **string**| ordering manner |

### Return type

[**\OpenAPI\Client\Model\Template[]**](../Model/Template.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTemplatesPost()`

```php
projectProjectIdTemplatesPost($project_id, $template): \OpenAPI\Client\Model\Template
```

create template

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$template = new \OpenAPI\Client\Model\TemplateRequest(); // \OpenAPI\Client\Model\TemplateRequest

try {
    $result = $apiInstance->projectProjectIdTemplatesPost($project_id, $template);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTemplatesPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **template** | [**\OpenAPI\Client\Model\TemplateRequest**](../Model/TemplateRequest.md)|  |

### Return type

[**\OpenAPI\Client\Model\Template**](../Model/Template.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTemplatesTemplateIdDelete()`

```php
projectProjectIdTemplatesTemplateIdDelete($project_id, $template_id)
```

Removes template

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$template_id = 7; // int | template ID

try {
    $apiInstance->projectProjectIdTemplatesTemplateIdDelete($project_id, $template_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTemplatesTemplateIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **template_id** | **int**| template ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTemplatesTemplateIdGet()`

```php
projectProjectIdTemplatesTemplateIdGet($project_id, $template_id): \OpenAPI\Client\Model\Template
```

Get template

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$template_id = 7; // int | template ID

try {
    $result = $apiInstance->projectProjectIdTemplatesTemplateIdGet($project_id, $template_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTemplatesTemplateIdGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **template_id** | **int**| template ID |

### Return type

[**\OpenAPI\Client\Model\Template**](../Model/Template.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdTemplatesTemplateIdPut()`

```php
projectProjectIdTemplatesTemplateIdPut($project_id, $template_id, $template)
```

Updates template

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$template_id = 7; // int | template ID
$template = new \OpenAPI\Client\Model\TemplateRequest(); // \OpenAPI\Client\Model\TemplateRequest

try {
    $apiInstance->projectProjectIdTemplatesTemplateIdPut($project_id, $template_id, $template);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdTemplatesTemplateIdPut: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **template_id** | **int**| template ID |
 **template** | [**\OpenAPI\Client\Model\TemplateRequest**](../Model/TemplateRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdUsersGet()`

```php
projectProjectIdUsersGet($project_id, $sort, $order): \OpenAPI\Client\Model\User[]
```

Get users linked to project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$sort = email; // string | sorting name
$order = desc; // string | ordering manner

try {
    $result = $apiInstance->projectProjectIdUsersGet($project_id, $sort, $order);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdUsersGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **sort** | **string**| sorting name |
 **order** | **string**| ordering manner |

### Return type

[**\OpenAPI\Client\Model\User[]**](../Model/User.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdUsersPost()`

```php
projectProjectIdUsersPost($project_id, $user)
```

Link user to project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$user = new \OpenAPI\Client\Model\ProjectProjectIdUsersGetRequest(); // \OpenAPI\Client\Model\ProjectProjectIdUsersGetRequest

try {
    $apiInstance->projectProjectIdUsersPost($project_id, $user);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdUsersPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **user** | [**\OpenAPI\Client\Model\ProjectProjectIdUsersGetRequest**](../Model/ProjectProjectIdUsersGetRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdUsersUserIdAdminDelete()`

```php
projectProjectIdUsersUserIdAdminDelete($project_id, $user_id)
```

Revoke admin privileges

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$user_id = 2; // int | User ID

try {
    $apiInstance->projectProjectIdUsersUserIdAdminDelete($project_id, $user_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdUsersUserIdAdminDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **user_id** | **int**| User ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdUsersUserIdAdminPost()`

```php
projectProjectIdUsersUserIdAdminPost($project_id, $user_id)
```

Makes user admin

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$user_id = 2; // int | User ID

try {
    $apiInstance->projectProjectIdUsersUserIdAdminPost($project_id, $user_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdUsersUserIdAdminPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **user_id** | **int**| User ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdUsersUserIdDelete()`

```php
projectProjectIdUsersUserIdDelete($project_id, $user_id)
```

Removes user from project

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$user_id = 2; // int | User ID

try {
    $apiInstance->projectProjectIdUsersUserIdDelete($project_id, $user_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdUsersUserIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **user_id** | **int**| User ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdViewsGet()`

```php
projectProjectIdViewsGet($project_id): \OpenAPI\Client\Model\View[]
```

Get view

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID

try {
    $result = $apiInstance->projectProjectIdViewsGet($project_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdViewsGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |

### Return type

[**\OpenAPI\Client\Model\View[]**](../Model/View.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdViewsPost()`

```php
projectProjectIdViewsPost($project_id, $view): \OpenAPI\Client\Model\View
```

create view

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$view = new \OpenAPI\Client\Model\ViewRequest(); // \OpenAPI\Client\Model\ViewRequest

try {
    $result = $apiInstance->projectProjectIdViewsPost($project_id, $view);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdViewsPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **view** | [**\OpenAPI\Client\Model\ViewRequest**](../Model/ViewRequest.md)|  |

### Return type

[**\OpenAPI\Client\Model\View**](../Model/View.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdViewsViewIdDelete()`

```php
projectProjectIdViewsViewIdDelete($project_id, $view_id)
```

Removes view

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$view_id = 10; // int | view ID

try {
    $apiInstance->projectProjectIdViewsViewIdDelete($project_id, $view_id);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdViewsViewIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **view_id** | **int**| view ID |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdViewsViewIdGet()`

```php
projectProjectIdViewsViewIdGet($project_id, $view_id): \OpenAPI\Client\Model\View
```

Get view

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$view_id = 10; // int | view ID

try {
    $result = $apiInstance->projectProjectIdViewsViewIdGet($project_id, $view_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdViewsViewIdGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **view_id** | **int**| view ID |

### Return type

[**\OpenAPI\Client\Model\View**](../Model/View.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdViewsViewIdPut()`

```php
projectProjectIdViewsViewIdPut($project_id, $view_id, $view)
```

Updates view

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: bearer
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

// Configure API key authorization: cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Cookie', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Cookie', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\ProjectApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$view_id = 10; // int | view ID
$view = new \OpenAPI\Client\Model\ViewRequest(); // \OpenAPI\Client\Model\ViewRequest

try {
    $apiInstance->projectProjectIdViewsViewIdPut($project_id, $view_id, $view);
} catch (Exception $e) {
    echo 'Exception when calling ProjectApi->projectProjectIdViewsViewIdPut: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **view_id** | **int**| view ID |
 **view** | [**\OpenAPI\Client\Model\ViewRequest**](../Model/ViewRequest.md)|  |

### Return type

void (empty response body)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
