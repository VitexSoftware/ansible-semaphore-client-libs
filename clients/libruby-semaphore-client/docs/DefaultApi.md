# OpenapiClient::DefaultApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**events_get**](DefaultApi.md#events_get) | **GET** /events | Get Events related to Semaphore and projects you are part of |
| [**events_last_get**](DefaultApi.md#events_last_get) | **GET** /events/last | Get last 200 Events related to Semaphore and projects you are part of |
| [**info_get**](DefaultApi.md#info_get) | **GET** /info | Fetches information about semaphore |
| [**ping_get**](DefaultApi.md#ping_get) | **GET** /ping | PING test |
| [**ws_get**](DefaultApi.md#ws_get) | **GET** /ws | Websocket handler |


## events_get

> <Array<Event>> events_get

Get Events related to Semaphore and projects you are part of

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

api_instance = OpenapiClient::DefaultApi.new

begin
  # Get Events related to Semaphore and projects you are part of
  result = api_instance.events_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->events_get: #{e}"
end
```

#### Using the events_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Event>>, Integer, Hash)> events_get_with_http_info

```ruby
begin
  # Get Events related to Semaphore and projects you are part of
  data, status_code, headers = api_instance.events_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Event>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->events_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Array&lt;Event&gt;**](Event.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## events_last_get

> <Array<Event>> events_last_get

Get last 200 Events related to Semaphore and projects you are part of

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

api_instance = OpenapiClient::DefaultApi.new

begin
  # Get last 200 Events related to Semaphore and projects you are part of
  result = api_instance.events_last_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->events_last_get: #{e}"
end
```

#### Using the events_last_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Event>>, Integer, Hash)> events_last_get_with_http_info

```ruby
begin
  # Get last 200 Events related to Semaphore and projects you are part of
  data, status_code, headers = api_instance.events_last_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Event>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->events_last_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Array&lt;Event&gt;**](Event.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## info_get

> <InfoType> info_get

Fetches information about semaphore

you must be authenticated to use this

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

api_instance = OpenapiClient::DefaultApi.new

begin
  # Fetches information about semaphore
  result = api_instance.info_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->info_get: #{e}"
end
```

#### Using the info_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InfoType>, Integer, Hash)> info_get_with_http_info

```ruby
begin
  # Fetches information about semaphore
  data, status_code, headers = api_instance.info_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InfoType>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->info_get_with_http_info: #{e}"
end
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


## ping_get

> String ping_get

PING test

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

api_instance = OpenapiClient::DefaultApi.new

begin
  # PING test
  result = api_instance.ping_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->ping_get: #{e}"
end
```

#### Using the ping_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(String, Integer, Hash)> ping_get_with_http_info

```ruby
begin
  # PING test
  data, status_code, headers = api_instance.ping_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => String
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->ping_get_with_http_info: #{e}"
end
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


## ws_get

> ws_get

Websocket handler

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

api_instance = OpenapiClient::DefaultApi.new

begin
  # Websocket handler
  api_instance.ws_get
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->ws_get: #{e}"
end
```

#### Using the ws_get_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> ws_get_with_http_info

```ruby
begin
  # Websocket handler
  data, status_code, headers = api_instance.ws_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->ws_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

