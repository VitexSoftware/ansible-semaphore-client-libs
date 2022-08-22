# OpenapiClient::ProjectsApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**projects_get**](ProjectsApi.md#projects_get) | **GET** /projects | Get projects |
| [**projects_post**](ProjectsApi.md#projects_post) | **POST** /projects | Create a new project |


## projects_get

> <Array<Project>> projects_get

Get projects

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

api_instance = OpenapiClient::ProjectsApi.new

begin
  # Get projects
  result = api_instance.projects_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectsApi->projects_get: #{e}"
end
```

#### Using the projects_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Project>>, Integer, Hash)> projects_get_with_http_info

```ruby
begin
  # Get projects
  data, status_code, headers = api_instance.projects_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Project>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectsApi->projects_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Array&lt;Project&gt;**](Project.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## projects_post

> projects_post(project)

Create a new project

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

api_instance = OpenapiClient::ProjectsApi.new
project = OpenapiClient::ProjectRequest.new # ProjectRequest | 

begin
  # Create a new project
  api_instance.projects_post(project)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectsApi->projects_post: #{e}"
end
```

#### Using the projects_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> projects_post_with_http_info(project)

```ruby
begin
  # Create a new project
  data, status_code, headers = api_instance.projects_post_with_http_info(project)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectsApi->projects_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project** | [**ProjectRequest**](ProjectRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

