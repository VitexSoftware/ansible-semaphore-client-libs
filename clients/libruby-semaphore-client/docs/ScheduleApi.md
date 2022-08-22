# OpenapiClient::ScheduleApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**project_project_id_schedules_post**](ScheduleApi.md#project_project_id_schedules_post) | **POST** /project/{project_id}/schedules | create schedule |
| [**project_project_id_schedules_schedule_id_delete**](ScheduleApi.md#project_project_id_schedules_schedule_id_delete) | **DELETE** /project/{project_id}/schedules/{schedule_id} | Deletes schedule |
| [**project_project_id_schedules_schedule_id_get**](ScheduleApi.md#project_project_id_schedules_schedule_id_get) | **GET** /project/{project_id}/schedules/{schedule_id} | Get schedule |
| [**project_project_id_schedules_schedule_id_put**](ScheduleApi.md#project_project_id_schedules_schedule_id_put) | **PUT** /project/{project_id}/schedules/{schedule_id} | Updates schedule |


## project_project_id_schedules_post

> <Schedule> project_project_id_schedules_post(project_id, schedule)

create schedule

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure API key authorization: bearer
  config.api_key['bearer'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['bearer'] = 'Bearer'

  # Configure API key authorization: cookie
  config.api_key['cookie'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['cookie'] = 'Bearer'
end

api_instance = OpenapiClient::ScheduleApi.new
project_id = 1 # Integer | Project ID
schedule = OpenapiClient::ScheduleRequest.new # ScheduleRequest | 

begin
  # create schedule
  result = api_instance.project_project_id_schedules_post(project_id, schedule)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ScheduleApi->project_project_id_schedules_post: #{e}"
end
```

#### Using the project_project_id_schedules_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Schedule>, Integer, Hash)> project_project_id_schedules_post_with_http_info(project_id, schedule)

```ruby
begin
  # create schedule
  data, status_code, headers = api_instance.project_project_id_schedules_post_with_http_info(project_id, schedule)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Schedule>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ScheduleApi->project_project_id_schedules_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **schedule** | [**ScheduleRequest**](ScheduleRequest.md) |  |  |

### Return type

[**Schedule**](Schedule.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_schedules_schedule_id_delete

> project_project_id_schedules_schedule_id_delete(project_id, schedule_id)

Deletes schedule

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure API key authorization: bearer
  config.api_key['bearer'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['bearer'] = 'Bearer'

  # Configure API key authorization: cookie
  config.api_key['cookie'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['cookie'] = 'Bearer'
end

api_instance = OpenapiClient::ScheduleApi.new
project_id = 1 # Integer | Project ID
schedule_id = 9 # Integer | schedule ID

begin
  # Deletes schedule
  api_instance.project_project_id_schedules_schedule_id_delete(project_id, schedule_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ScheduleApi->project_project_id_schedules_schedule_id_delete: #{e}"
end
```

#### Using the project_project_id_schedules_schedule_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_schedules_schedule_id_delete_with_http_info(project_id, schedule_id)

```ruby
begin
  # Deletes schedule
  data, status_code, headers = api_instance.project_project_id_schedules_schedule_id_delete_with_http_info(project_id, schedule_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ScheduleApi->project_project_id_schedules_schedule_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **schedule_id** | **Integer** | schedule ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_schedules_schedule_id_get

> <Schedule> project_project_id_schedules_schedule_id_get(project_id, schedule_id)

Get schedule

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure API key authorization: bearer
  config.api_key['bearer'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['bearer'] = 'Bearer'

  # Configure API key authorization: cookie
  config.api_key['cookie'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['cookie'] = 'Bearer'
end

api_instance = OpenapiClient::ScheduleApi.new
project_id = 1 # Integer | Project ID
schedule_id = 9 # Integer | schedule ID

begin
  # Get schedule
  result = api_instance.project_project_id_schedules_schedule_id_get(project_id, schedule_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ScheduleApi->project_project_id_schedules_schedule_id_get: #{e}"
end
```

#### Using the project_project_id_schedules_schedule_id_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Schedule>, Integer, Hash)> project_project_id_schedules_schedule_id_get_with_http_info(project_id, schedule_id)

```ruby
begin
  # Get schedule
  data, status_code, headers = api_instance.project_project_id_schedules_schedule_id_get_with_http_info(project_id, schedule_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Schedule>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ScheduleApi->project_project_id_schedules_schedule_id_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **schedule_id** | **Integer** | schedule ID |  |

### Return type

[**Schedule**](Schedule.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_schedules_schedule_id_put

> project_project_id_schedules_schedule_id_put(project_id, schedule_id, schedule)

Updates schedule

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure API key authorization: bearer
  config.api_key['bearer'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['bearer'] = 'Bearer'

  # Configure API key authorization: cookie
  config.api_key['cookie'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['cookie'] = 'Bearer'
end

api_instance = OpenapiClient::ScheduleApi.new
project_id = 1 # Integer | Project ID
schedule_id = 9 # Integer | schedule ID
schedule = OpenapiClient::ScheduleRequest.new # ScheduleRequest | 

begin
  # Updates schedule
  api_instance.project_project_id_schedules_schedule_id_put(project_id, schedule_id, schedule)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ScheduleApi->project_project_id_schedules_schedule_id_put: #{e}"
end
```

#### Using the project_project_id_schedules_schedule_id_put_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_schedules_schedule_id_put_with_http_info(project_id, schedule_id, schedule)

```ruby
begin
  # Updates schedule
  data, status_code, headers = api_instance.project_project_id_schedules_schedule_id_put_with_http_info(project_id, schedule_id, schedule)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ScheduleApi->project_project_id_schedules_schedule_id_put_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **schedule_id** | **Integer** | schedule ID |  |
| **schedule** | [**ScheduleRequest**](ScheduleRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

