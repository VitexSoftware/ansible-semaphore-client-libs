# OpenAPI\Client\ScheduleApi

All URIs are relative to https://demo.ansible-semaphore.com/api.

Method | HTTP request | Description
------------- | ------------- | -------------
[**projectProjectIdSchedulesPost()**](ScheduleApi.md#projectProjectIdSchedulesPost) | **POST** /project/{project_id}/schedules | create schedule
[**projectProjectIdSchedulesScheduleIdDelete()**](ScheduleApi.md#projectProjectIdSchedulesScheduleIdDelete) | **DELETE** /project/{project_id}/schedules/{schedule_id} | Deletes schedule
[**projectProjectIdSchedulesScheduleIdGet()**](ScheduleApi.md#projectProjectIdSchedulesScheduleIdGet) | **GET** /project/{project_id}/schedules/{schedule_id} | Get schedule
[**projectProjectIdSchedulesScheduleIdPut()**](ScheduleApi.md#projectProjectIdSchedulesScheduleIdPut) | **PUT** /project/{project_id}/schedules/{schedule_id} | Updates schedule


## `projectProjectIdSchedulesPost()`

```php
projectProjectIdSchedulesPost($project_id, $schedule): \OpenAPI\Client\Model\Schedule
```

create schedule

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


$apiInstance = new OpenAPI\Client\Api\ScheduleApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$schedule = new \OpenAPI\Client\Model\ScheduleRequest(); // \OpenAPI\Client\Model\ScheduleRequest

try {
    $result = $apiInstance->projectProjectIdSchedulesPost($project_id, $schedule);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ScheduleApi->projectProjectIdSchedulesPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **schedule** | [**\OpenAPI\Client\Model\ScheduleRequest**](../Model/ScheduleRequest.md)|  |

### Return type

[**\OpenAPI\Client\Model\Schedule**](../Model/Schedule.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdSchedulesScheduleIdDelete()`

```php
projectProjectIdSchedulesScheduleIdDelete($project_id, $schedule_id)
```

Deletes schedule

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


$apiInstance = new OpenAPI\Client\Api\ScheduleApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$schedule_id = 9; // int | schedule ID

try {
    $apiInstance->projectProjectIdSchedulesScheduleIdDelete($project_id, $schedule_id);
} catch (Exception $e) {
    echo 'Exception when calling ScheduleApi->projectProjectIdSchedulesScheduleIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **schedule_id** | **int**| schedule ID |

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

## `projectProjectIdSchedulesScheduleIdGet()`

```php
projectProjectIdSchedulesScheduleIdGet($project_id, $schedule_id): \OpenAPI\Client\Model\Schedule
```

Get schedule

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


$apiInstance = new OpenAPI\Client\Api\ScheduleApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$schedule_id = 9; // int | schedule ID

try {
    $result = $apiInstance->projectProjectIdSchedulesScheduleIdGet($project_id, $schedule_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ScheduleApi->projectProjectIdSchedulesScheduleIdGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **schedule_id** | **int**| schedule ID |

### Return type

[**\OpenAPI\Client\Model\Schedule**](../Model/Schedule.md)

### Authorization

[bearer](../../README.md#bearer), [cookie](../../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `text/plain; charset=utf-8`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `projectProjectIdSchedulesScheduleIdPut()`

```php
projectProjectIdSchedulesScheduleIdPut($project_id, $schedule_id, $schedule)
```

Updates schedule

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


$apiInstance = new OpenAPI\Client\Api\ScheduleApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$project_id = 1; // int | Project ID
$schedule_id = 9; // int | schedule ID
$schedule = new \OpenAPI\Client\Model\ScheduleRequest(); // \OpenAPI\Client\Model\ScheduleRequest

try {
    $apiInstance->projectProjectIdSchedulesScheduleIdPut($project_id, $schedule_id, $schedule);
} catch (Exception $e) {
    echo 'Exception when calling ScheduleApi->projectProjectIdSchedulesScheduleIdPut: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **int**| Project ID |
 **schedule_id** | **int**| schedule ID |
 **schedule** | [**\OpenAPI\Client\Model\ScheduleRequest**](../Model/ScheduleRequest.md)|  |

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
