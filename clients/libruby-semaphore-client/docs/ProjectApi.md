# OpenapiClient::ProjectApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**project_project_id_delete**](ProjectApi.md#project_project_id_delete) | **DELETE** /project/{project_id}/ | Delete project |
| [**project_project_id_environment_environment_id_delete**](ProjectApi.md#project_project_id_environment_environment_id_delete) | **DELETE** /project/{project_id}/environment/{environment_id} | Removes environment |
| [**project_project_id_environment_environment_id_put**](ProjectApi.md#project_project_id_environment_environment_id_put) | **PUT** /project/{project_id}/environment/{environment_id} | Update environment |
| [**project_project_id_environment_get**](ProjectApi.md#project_project_id_environment_get) | **GET** /project/{project_id}/environment | Get environment |
| [**project_project_id_environment_post**](ProjectApi.md#project_project_id_environment_post) | **POST** /project/{project_id}/environment | Add environment |
| [**project_project_id_events_get**](ProjectApi.md#project_project_id_events_get) | **GET** /project/{project_id}/events | Get Events related to this project |
| [**project_project_id_get**](ProjectApi.md#project_project_id_get) | **GET** /project/{project_id}/ | Fetch project |
| [**project_project_id_inventory_get**](ProjectApi.md#project_project_id_inventory_get) | **GET** /project/{project_id}/inventory | Get inventory |
| [**project_project_id_inventory_inventory_id_delete**](ProjectApi.md#project_project_id_inventory_inventory_id_delete) | **DELETE** /project/{project_id}/inventory/{inventory_id} | Removes inventory |
| [**project_project_id_inventory_inventory_id_put**](ProjectApi.md#project_project_id_inventory_inventory_id_put) | **PUT** /project/{project_id}/inventory/{inventory_id} | Updates inventory |
| [**project_project_id_inventory_post**](ProjectApi.md#project_project_id_inventory_post) | **POST** /project/{project_id}/inventory | create inventory |
| [**project_project_id_keys_get**](ProjectApi.md#project_project_id_keys_get) | **GET** /project/{project_id}/keys | Get access keys linked to project |
| [**project_project_id_keys_key_id_delete**](ProjectApi.md#project_project_id_keys_key_id_delete) | **DELETE** /project/{project_id}/keys/{key_id} | Removes access key |
| [**project_project_id_keys_key_id_put**](ProjectApi.md#project_project_id_keys_key_id_put) | **PUT** /project/{project_id}/keys/{key_id} | Updates access key |
| [**project_project_id_keys_post**](ProjectApi.md#project_project_id_keys_post) | **POST** /project/{project_id}/keys | Add access key |
| [**project_project_id_put**](ProjectApi.md#project_project_id_put) | **PUT** /project/{project_id}/ | Update project |
| [**project_project_id_repositories_get**](ProjectApi.md#project_project_id_repositories_get) | **GET** /project/{project_id}/repositories | Get repositories |
| [**project_project_id_repositories_post**](ProjectApi.md#project_project_id_repositories_post) | **POST** /project/{project_id}/repositories | Add repository |
| [**project_project_id_repositories_repository_id_delete**](ProjectApi.md#project_project_id_repositories_repository_id_delete) | **DELETE** /project/{project_id}/repositories/{repository_id} | Removes repository |
| [**project_project_id_tasks_get**](ProjectApi.md#project_project_id_tasks_get) | **GET** /project/{project_id}/tasks | Get Tasks related to current project |
| [**project_project_id_tasks_last_get**](ProjectApi.md#project_project_id_tasks_last_get) | **GET** /project/{project_id}/tasks/last | Get last 200 Tasks related to current project |
| [**project_project_id_tasks_post**](ProjectApi.md#project_project_id_tasks_post) | **POST** /project/{project_id}/tasks | Starts a job |
| [**project_project_id_tasks_task_id_delete**](ProjectApi.md#project_project_id_tasks_task_id_delete) | **DELETE** /project/{project_id}/tasks/{task_id} | Deletes task (including output) |
| [**project_project_id_tasks_task_id_get**](ProjectApi.md#project_project_id_tasks_task_id_get) | **GET** /project/{project_id}/tasks/{task_id} | Get a single task |
| [**project_project_id_tasks_task_id_output_get**](ProjectApi.md#project_project_id_tasks_task_id_output_get) | **GET** /project/{project_id}/tasks/{task_id}/output | Get task output |
| [**project_project_id_templates_get**](ProjectApi.md#project_project_id_templates_get) | **GET** /project/{project_id}/templates | Get template |
| [**project_project_id_templates_post**](ProjectApi.md#project_project_id_templates_post) | **POST** /project/{project_id}/templates | create template |
| [**project_project_id_templates_template_id_delete**](ProjectApi.md#project_project_id_templates_template_id_delete) | **DELETE** /project/{project_id}/templates/{template_id} | Removes template |
| [**project_project_id_templates_template_id_get**](ProjectApi.md#project_project_id_templates_template_id_get) | **GET** /project/{project_id}/templates/{template_id} | Get template |
| [**project_project_id_templates_template_id_put**](ProjectApi.md#project_project_id_templates_template_id_put) | **PUT** /project/{project_id}/templates/{template_id} | Updates template |
| [**project_project_id_users_get**](ProjectApi.md#project_project_id_users_get) | **GET** /project/{project_id}/users | Get users linked to project |
| [**project_project_id_users_post**](ProjectApi.md#project_project_id_users_post) | **POST** /project/{project_id}/users | Link user to project |
| [**project_project_id_users_user_id_admin_delete**](ProjectApi.md#project_project_id_users_user_id_admin_delete) | **DELETE** /project/{project_id}/users/{user_id}/admin | Revoke admin privileges |
| [**project_project_id_users_user_id_admin_post**](ProjectApi.md#project_project_id_users_user_id_admin_post) | **POST** /project/{project_id}/users/{user_id}/admin | Makes user admin |
| [**project_project_id_users_user_id_delete**](ProjectApi.md#project_project_id_users_user_id_delete) | **DELETE** /project/{project_id}/users/{user_id} | Removes user from project |
| [**project_project_id_views_get**](ProjectApi.md#project_project_id_views_get) | **GET** /project/{project_id}/views | Get view |
| [**project_project_id_views_post**](ProjectApi.md#project_project_id_views_post) | **POST** /project/{project_id}/views | create view |
| [**project_project_id_views_view_id_delete**](ProjectApi.md#project_project_id_views_view_id_delete) | **DELETE** /project/{project_id}/views/{view_id} | Removes view |
| [**project_project_id_views_view_id_get**](ProjectApi.md#project_project_id_views_view_id_get) | **GET** /project/{project_id}/views/{view_id} | Get view |
| [**project_project_id_views_view_id_put**](ProjectApi.md#project_project_id_views_view_id_put) | **PUT** /project/{project_id}/views/{view_id} | Updates view |


## project_project_id_delete

> project_project_id_delete(project_id)

Delete project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID

begin
  # Delete project
  api_instance.project_project_id_delete(project_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_delete: #{e}"
end
```

#### Using the project_project_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_delete_with_http_info(project_id)

```ruby
begin
  # Delete project
  data, status_code, headers = api_instance.project_project_id_delete_with_http_info(project_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_environment_environment_id_delete

> project_project_id_environment_environment_id_delete(project_id, environment_id)

Removes environment

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
environment_id = 6 # Integer | environment ID

begin
  # Removes environment
  api_instance.project_project_id_environment_environment_id_delete(project_id, environment_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_environment_environment_id_delete: #{e}"
end
```

#### Using the project_project_id_environment_environment_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_environment_environment_id_delete_with_http_info(project_id, environment_id)

```ruby
begin
  # Removes environment
  data, status_code, headers = api_instance.project_project_id_environment_environment_id_delete_with_http_info(project_id, environment_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_environment_environment_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **environment_id** | **Integer** | environment ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_environment_environment_id_put

> project_project_id_environment_environment_id_put(project_id, environment_id, environment)

Update environment

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
environment_id = 6 # Integer | environment ID
environment = OpenapiClient::EnvironmentRequest.new # EnvironmentRequest | 

begin
  # Update environment
  api_instance.project_project_id_environment_environment_id_put(project_id, environment_id, environment)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_environment_environment_id_put: #{e}"
end
```

#### Using the project_project_id_environment_environment_id_put_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_environment_environment_id_put_with_http_info(project_id, environment_id, environment)

```ruby
begin
  # Update environment
  data, status_code, headers = api_instance.project_project_id_environment_environment_id_put_with_http_info(project_id, environment_id, environment)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_environment_environment_id_put_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **environment_id** | **Integer** | environment ID |  |
| **environment** | [**EnvironmentRequest**](EnvironmentRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## project_project_id_environment_get

> <Array<Environment>> project_project_id_environment_get(project_id, sort, order)

Get environment

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
sort = 'db-deploy' # String | sorting name
order = 'desc' # String | ordering manner

begin
  # Get environment
  result = api_instance.project_project_id_environment_get(project_id, sort, order)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_environment_get: #{e}"
end
```

#### Using the project_project_id_environment_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Environment>>, Integer, Hash)> project_project_id_environment_get_with_http_info(project_id, sort, order)

```ruby
begin
  # Get environment
  data, status_code, headers = api_instance.project_project_id_environment_get_with_http_info(project_id, sort, order)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Environment>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_environment_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **sort** | **String** | sorting name |  |
| **order** | **String** | ordering manner |  |

### Return type

[**Array&lt;Environment&gt;**](Environment.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_environment_post

> project_project_id_environment_post(project_id, environment)

Add environment

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
environment = OpenapiClient::EnvironmentRequest.new # EnvironmentRequest | 

begin
  # Add environment
  api_instance.project_project_id_environment_post(project_id, environment)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_environment_post: #{e}"
end
```

#### Using the project_project_id_environment_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_environment_post_with_http_info(project_id, environment)

```ruby
begin
  # Add environment
  data, status_code, headers = api_instance.project_project_id_environment_post_with_http_info(project_id, environment)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_environment_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **environment** | [**EnvironmentRequest**](EnvironmentRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## project_project_id_events_get

> <Array<Event>> project_project_id_events_get(project_id)

Get Events related to this project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID

begin
  # Get Events related to this project
  result = api_instance.project_project_id_events_get(project_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_events_get: #{e}"
end
```

#### Using the project_project_id_events_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Event>>, Integer, Hash)> project_project_id_events_get_with_http_info(project_id)

```ruby
begin
  # Get Events related to this project
  data, status_code, headers = api_instance.project_project_id_events_get_with_http_info(project_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Event>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_events_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |

### Return type

[**Array&lt;Event&gt;**](Event.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_get

> <Project> project_project_id_get(project_id)

Fetch project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID

begin
  # Fetch project
  result = api_instance.project_project_id_get(project_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_get: #{e}"
end
```

#### Using the project_project_id_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Project>, Integer, Hash)> project_project_id_get_with_http_info(project_id)

```ruby
begin
  # Fetch project
  data, status_code, headers = api_instance.project_project_id_get_with_http_info(project_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Project>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |

### Return type

[**Project**](Project.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_inventory_get

> <Array<Inventory>> project_project_id_inventory_get(project_id, sort, order)

Get inventory

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
sort = 'name' # String | sorting name
order = 'asc' # String | ordering manner

begin
  # Get inventory
  result = api_instance.project_project_id_inventory_get(project_id, sort, order)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_inventory_get: #{e}"
end
```

#### Using the project_project_id_inventory_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Inventory>>, Integer, Hash)> project_project_id_inventory_get_with_http_info(project_id, sort, order)

```ruby
begin
  # Get inventory
  data, status_code, headers = api_instance.project_project_id_inventory_get_with_http_info(project_id, sort, order)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Inventory>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_inventory_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **sort** | **String** | sorting name |  |
| **order** | **String** | ordering manner |  |

### Return type

[**Array&lt;Inventory&gt;**](Inventory.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_inventory_inventory_id_delete

> project_project_id_inventory_inventory_id_delete(project_id, inventory_id)

Removes inventory

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
inventory_id = 5 # Integer | inventory ID

begin
  # Removes inventory
  api_instance.project_project_id_inventory_inventory_id_delete(project_id, inventory_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_inventory_inventory_id_delete: #{e}"
end
```

#### Using the project_project_id_inventory_inventory_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_inventory_inventory_id_delete_with_http_info(project_id, inventory_id)

```ruby
begin
  # Removes inventory
  data, status_code, headers = api_instance.project_project_id_inventory_inventory_id_delete_with_http_info(project_id, inventory_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_inventory_inventory_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **inventory_id** | **Integer** | inventory ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_inventory_inventory_id_put

> project_project_id_inventory_inventory_id_put(project_id, inventory_id, inventory)

Updates inventory

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
inventory_id = 5 # Integer | inventory ID
inventory = OpenapiClient::InventoryRequest.new # InventoryRequest | 

begin
  # Updates inventory
  api_instance.project_project_id_inventory_inventory_id_put(project_id, inventory_id, inventory)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_inventory_inventory_id_put: #{e}"
end
```

#### Using the project_project_id_inventory_inventory_id_put_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_inventory_inventory_id_put_with_http_info(project_id, inventory_id, inventory)

```ruby
begin
  # Updates inventory
  data, status_code, headers = api_instance.project_project_id_inventory_inventory_id_put_with_http_info(project_id, inventory_id, inventory)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_inventory_inventory_id_put_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **inventory_id** | **Integer** | inventory ID |  |
| **inventory** | [**InventoryRequest**](InventoryRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## project_project_id_inventory_post

> <Inventory> project_project_id_inventory_post(project_id, inventory)

create inventory

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
inventory = OpenapiClient::InventoryRequest.new # InventoryRequest | 

begin
  # create inventory
  result = api_instance.project_project_id_inventory_post(project_id, inventory)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_inventory_post: #{e}"
end
```

#### Using the project_project_id_inventory_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Inventory>, Integer, Hash)> project_project_id_inventory_post_with_http_info(project_id, inventory)

```ruby
begin
  # create inventory
  data, status_code, headers = api_instance.project_project_id_inventory_post_with_http_info(project_id, inventory)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Inventory>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_inventory_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **inventory** | [**InventoryRequest**](InventoryRequest.md) |  |  |

### Return type

[**Inventory**](Inventory.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_keys_get

> <Array<AccessKey>> project_project_id_keys_get(project_id, sort, order, opts)

Get access keys linked to project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
sort = 'name' # String | sorting name
order = 'asc' # String | ordering manner
opts = {
  key_type: 'none' # String | Filter by key type
}

begin
  # Get access keys linked to project
  result = api_instance.project_project_id_keys_get(project_id, sort, order, opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_keys_get: #{e}"
end
```

#### Using the project_project_id_keys_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<AccessKey>>, Integer, Hash)> project_project_id_keys_get_with_http_info(project_id, sort, order, opts)

```ruby
begin
  # Get access keys linked to project
  data, status_code, headers = api_instance.project_project_id_keys_get_with_http_info(project_id, sort, order, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<AccessKey>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_keys_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **sort** | **String** | sorting name |  |
| **order** | **String** | ordering manner |  |
| **key_type** | **String** | Filter by key type | [optional] |

### Return type

[**Array&lt;AccessKey&gt;**](AccessKey.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_keys_key_id_delete

> project_project_id_keys_key_id_delete(project_id, key_id)

Removes access key

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
key_id = 3 # Integer | key ID

begin
  # Removes access key
  api_instance.project_project_id_keys_key_id_delete(project_id, key_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_keys_key_id_delete: #{e}"
end
```

#### Using the project_project_id_keys_key_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_keys_key_id_delete_with_http_info(project_id, key_id)

```ruby
begin
  # Removes access key
  data, status_code, headers = api_instance.project_project_id_keys_key_id_delete_with_http_info(project_id, key_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_keys_key_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **key_id** | **Integer** | key ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_keys_key_id_put

> project_project_id_keys_key_id_put(project_id, key_id, access_key)

Updates access key

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
key_id = 3 # Integer | key ID
access_key = OpenapiClient::AccessKeyRequest.new # AccessKeyRequest | 

begin
  # Updates access key
  api_instance.project_project_id_keys_key_id_put(project_id, key_id, access_key)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_keys_key_id_put: #{e}"
end
```

#### Using the project_project_id_keys_key_id_put_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_keys_key_id_put_with_http_info(project_id, key_id, access_key)

```ruby
begin
  # Updates access key
  data, status_code, headers = api_instance.project_project_id_keys_key_id_put_with_http_info(project_id, key_id, access_key)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_keys_key_id_put_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **key_id** | **Integer** | key ID |  |
| **access_key** | [**AccessKeyRequest**](AccessKeyRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## project_project_id_keys_post

> project_project_id_keys_post(project_id, access_key)

Add access key

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
access_key = OpenapiClient::AccessKeyRequest.new # AccessKeyRequest | 

begin
  # Add access key
  api_instance.project_project_id_keys_post(project_id, access_key)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_keys_post: #{e}"
end
```

#### Using the project_project_id_keys_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_keys_post_with_http_info(project_id, access_key)

```ruby
begin
  # Add access key
  data, status_code, headers = api_instance.project_project_id_keys_post_with_http_info(project_id, access_key)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_keys_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **access_key** | [**AccessKeyRequest**](AccessKeyRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## project_project_id_put

> project_project_id_put(project_id, project)

Update project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
project = OpenapiClient::ProjectProjectIdDeleteRequest.new # ProjectProjectIdDeleteRequest | 

begin
  # Update project
  api_instance.project_project_id_put(project_id, project)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_put: #{e}"
end
```

#### Using the project_project_id_put_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_put_with_http_info(project_id, project)

```ruby
begin
  # Update project
  data, status_code, headers = api_instance.project_project_id_put_with_http_info(project_id, project)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_put_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **project** | [**ProjectProjectIdDeleteRequest**](ProjectProjectIdDeleteRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## project_project_id_repositories_get

> <Array<Repository>> project_project_id_repositories_get(project_id, sort, order)

Get repositories

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
sort = 'name' # String | sorting name
order = 'asc' # String | ordering manner

begin
  # Get repositories
  result = api_instance.project_project_id_repositories_get(project_id, sort, order)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_repositories_get: #{e}"
end
```

#### Using the project_project_id_repositories_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Repository>>, Integer, Hash)> project_project_id_repositories_get_with_http_info(project_id, sort, order)

```ruby
begin
  # Get repositories
  data, status_code, headers = api_instance.project_project_id_repositories_get_with_http_info(project_id, sort, order)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Repository>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_repositories_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **sort** | **String** | sorting name |  |
| **order** | **String** | ordering manner |  |

### Return type

[**Array&lt;Repository&gt;**](Repository.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_repositories_post

> project_project_id_repositories_post(project_id, repository)

Add repository

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
repository = OpenapiClient::RepositoryRequest.new # RepositoryRequest | 

begin
  # Add repository
  api_instance.project_project_id_repositories_post(project_id, repository)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_repositories_post: #{e}"
end
```

#### Using the project_project_id_repositories_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_repositories_post_with_http_info(project_id, repository)

```ruby
begin
  # Add repository
  data, status_code, headers = api_instance.project_project_id_repositories_post_with_http_info(project_id, repository)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_repositories_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **repository** | [**RepositoryRequest**](RepositoryRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## project_project_id_repositories_repository_id_delete

> project_project_id_repositories_repository_id_delete(project_id, repository_id)

Removes repository

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
repository_id = 4 # Integer | repository ID

begin
  # Removes repository
  api_instance.project_project_id_repositories_repository_id_delete(project_id, repository_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_repositories_repository_id_delete: #{e}"
end
```

#### Using the project_project_id_repositories_repository_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_repositories_repository_id_delete_with_http_info(project_id, repository_id)

```ruby
begin
  # Removes repository
  data, status_code, headers = api_instance.project_project_id_repositories_repository_id_delete_with_http_info(project_id, repository_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_repositories_repository_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **repository_id** | **Integer** | repository ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_tasks_get

> <Array<Task>> project_project_id_tasks_get(project_id)

Get Tasks related to current project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID

begin
  # Get Tasks related to current project
  result = api_instance.project_project_id_tasks_get(project_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_get: #{e}"
end
```

#### Using the project_project_id_tasks_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Task>>, Integer, Hash)> project_project_id_tasks_get_with_http_info(project_id)

```ruby
begin
  # Get Tasks related to current project
  data, status_code, headers = api_instance.project_project_id_tasks_get_with_http_info(project_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Task>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |

### Return type

[**Array&lt;Task&gt;**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_tasks_last_get

> <Array<Task>> project_project_id_tasks_last_get(project_id)

Get last 200 Tasks related to current project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID

begin
  # Get last 200 Tasks related to current project
  result = api_instance.project_project_id_tasks_last_get(project_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_last_get: #{e}"
end
```

#### Using the project_project_id_tasks_last_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Task>>, Integer, Hash)> project_project_id_tasks_last_get_with_http_info(project_id)

```ruby
begin
  # Get last 200 Tasks related to current project
  data, status_code, headers = api_instance.project_project_id_tasks_last_get_with_http_info(project_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Task>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_last_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |

### Return type

[**Array&lt;Task&gt;**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_tasks_post

> <Task> project_project_id_tasks_post(project_id, task)

Starts a job

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
task = OpenapiClient::ProjectProjectIdTasksGetRequest.new # ProjectProjectIdTasksGetRequest | 

begin
  # Starts a job
  result = api_instance.project_project_id_tasks_post(project_id, task)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_post: #{e}"
end
```

#### Using the project_project_id_tasks_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Task>, Integer, Hash)> project_project_id_tasks_post_with_http_info(project_id, task)

```ruby
begin
  # Starts a job
  data, status_code, headers = api_instance.project_project_id_tasks_post_with_http_info(project_id, task)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Task>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **task** | [**ProjectProjectIdTasksGetRequest**](ProjectProjectIdTasksGetRequest.md) |  |  |

### Return type

[**Task**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_tasks_task_id_delete

> project_project_id_tasks_task_id_delete(project_id, task_id)

Deletes task (including output)

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
task_id = 8 # Integer | task ID

begin
  # Deletes task (including output)
  api_instance.project_project_id_tasks_task_id_delete(project_id, task_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_task_id_delete: #{e}"
end
```

#### Using the project_project_id_tasks_task_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_tasks_task_id_delete_with_http_info(project_id, task_id)

```ruby
begin
  # Deletes task (including output)
  data, status_code, headers = api_instance.project_project_id_tasks_task_id_delete_with_http_info(project_id, task_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_task_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **task_id** | **Integer** | task ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_tasks_task_id_get

> <Task> project_project_id_tasks_task_id_get(project_id, task_id)

Get a single task

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
task_id = 8 # Integer | task ID

begin
  # Get a single task
  result = api_instance.project_project_id_tasks_task_id_get(project_id, task_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_task_id_get: #{e}"
end
```

#### Using the project_project_id_tasks_task_id_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Task>, Integer, Hash)> project_project_id_tasks_task_id_get_with_http_info(project_id, task_id)

```ruby
begin
  # Get a single task
  data, status_code, headers = api_instance.project_project_id_tasks_task_id_get_with_http_info(project_id, task_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Task>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_task_id_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **task_id** | **Integer** | task ID |  |

### Return type

[**Task**](Task.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_tasks_task_id_output_get

> <Array<TaskOutput>> project_project_id_tasks_task_id_output_get(project_id, task_id)

Get task output

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
task_id = 8 # Integer | task ID

begin
  # Get task output
  result = api_instance.project_project_id_tasks_task_id_output_get(project_id, task_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_task_id_output_get: #{e}"
end
```

#### Using the project_project_id_tasks_task_id_output_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<TaskOutput>>, Integer, Hash)> project_project_id_tasks_task_id_output_get_with_http_info(project_id, task_id)

```ruby
begin
  # Get task output
  data, status_code, headers = api_instance.project_project_id_tasks_task_id_output_get_with_http_info(project_id, task_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<TaskOutput>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_tasks_task_id_output_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **task_id** | **Integer** | task ID |  |

### Return type

[**Array&lt;TaskOutput&gt;**](TaskOutput.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_templates_get

> <Array<Template>> project_project_id_templates_get(project_id, sort, order)

Get template

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
sort = 'alias' # String | sorting name
order = 'asc' # String | ordering manner

begin
  # Get template
  result = api_instance.project_project_id_templates_get(project_id, sort, order)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_get: #{e}"
end
```

#### Using the project_project_id_templates_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Template>>, Integer, Hash)> project_project_id_templates_get_with_http_info(project_id, sort, order)

```ruby
begin
  # Get template
  data, status_code, headers = api_instance.project_project_id_templates_get_with_http_info(project_id, sort, order)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Template>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **sort** | **String** | sorting name |  |
| **order** | **String** | ordering manner |  |

### Return type

[**Array&lt;Template&gt;**](Template.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_templates_post

> <Template> project_project_id_templates_post(project_id, template)

create template

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
template = OpenapiClient::TemplateRequest.new # TemplateRequest | 

begin
  # create template
  result = api_instance.project_project_id_templates_post(project_id, template)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_post: #{e}"
end
```

#### Using the project_project_id_templates_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Template>, Integer, Hash)> project_project_id_templates_post_with_http_info(project_id, template)

```ruby
begin
  # create template
  data, status_code, headers = api_instance.project_project_id_templates_post_with_http_info(project_id, template)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Template>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **template** | [**TemplateRequest**](TemplateRequest.md) |  |  |

### Return type

[**Template**](Template.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_templates_template_id_delete

> project_project_id_templates_template_id_delete(project_id, template_id)

Removes template

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
template_id = 7 # Integer | template ID

begin
  # Removes template
  api_instance.project_project_id_templates_template_id_delete(project_id, template_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_template_id_delete: #{e}"
end
```

#### Using the project_project_id_templates_template_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_templates_template_id_delete_with_http_info(project_id, template_id)

```ruby
begin
  # Removes template
  data, status_code, headers = api_instance.project_project_id_templates_template_id_delete_with_http_info(project_id, template_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_template_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **template_id** | **Integer** | template ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_templates_template_id_get

> <Template> project_project_id_templates_template_id_get(project_id, template_id)

Get template

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
template_id = 7 # Integer | template ID

begin
  # Get template
  result = api_instance.project_project_id_templates_template_id_get(project_id, template_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_template_id_get: #{e}"
end
```

#### Using the project_project_id_templates_template_id_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Template>, Integer, Hash)> project_project_id_templates_template_id_get_with_http_info(project_id, template_id)

```ruby
begin
  # Get template
  data, status_code, headers = api_instance.project_project_id_templates_template_id_get_with_http_info(project_id, template_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Template>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_template_id_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **template_id** | **Integer** | template ID |  |

### Return type

[**Template**](Template.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_templates_template_id_put

> project_project_id_templates_template_id_put(project_id, template_id, template)

Updates template

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
template_id = 7 # Integer | template ID
template = OpenapiClient::TemplateRequest.new # TemplateRequest | 

begin
  # Updates template
  api_instance.project_project_id_templates_template_id_put(project_id, template_id, template)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_template_id_put: #{e}"
end
```

#### Using the project_project_id_templates_template_id_put_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_templates_template_id_put_with_http_info(project_id, template_id, template)

```ruby
begin
  # Updates template
  data, status_code, headers = api_instance.project_project_id_templates_template_id_put_with_http_info(project_id, template_id, template)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_templates_template_id_put_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **template_id** | **Integer** | template ID |  |
| **template** | [**TemplateRequest**](TemplateRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## project_project_id_users_get

> <Array<User>> project_project_id_users_get(project_id, sort, order)

Get users linked to project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
sort = 'name' # String | sorting name
order = 'asc' # String | ordering manner

begin
  # Get users linked to project
  result = api_instance.project_project_id_users_get(project_id, sort, order)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_get: #{e}"
end
```

#### Using the project_project_id_users_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<User>>, Integer, Hash)> project_project_id_users_get_with_http_info(project_id, sort, order)

```ruby
begin
  # Get users linked to project
  data, status_code, headers = api_instance.project_project_id_users_get_with_http_info(project_id, sort, order)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<User>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **sort** | **String** | sorting name |  |
| **order** | **String** | ordering manner |  |

### Return type

[**Array&lt;User&gt;**](User.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_users_post

> project_project_id_users_post(project_id, user)

Link user to project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
user = OpenapiClient::ProjectProjectIdUsersGetRequest.new # ProjectProjectIdUsersGetRequest | 

begin
  # Link user to project
  api_instance.project_project_id_users_post(project_id, user)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_post: #{e}"
end
```

#### Using the project_project_id_users_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_users_post_with_http_info(project_id, user)

```ruby
begin
  # Link user to project
  data, status_code, headers = api_instance.project_project_id_users_post_with_http_info(project_id, user)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **user** | [**ProjectProjectIdUsersGetRequest**](ProjectProjectIdUsersGetRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## project_project_id_users_user_id_admin_delete

> project_project_id_users_user_id_admin_delete(project_id, user_id)

Revoke admin privileges

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
user_id = 2 # Integer | User ID

begin
  # Revoke admin privileges
  api_instance.project_project_id_users_user_id_admin_delete(project_id, user_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_user_id_admin_delete: #{e}"
end
```

#### Using the project_project_id_users_user_id_admin_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_users_user_id_admin_delete_with_http_info(project_id, user_id)

```ruby
begin
  # Revoke admin privileges
  data, status_code, headers = api_instance.project_project_id_users_user_id_admin_delete_with_http_info(project_id, user_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_user_id_admin_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **user_id** | **Integer** | User ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_users_user_id_admin_post

> project_project_id_users_user_id_admin_post(project_id, user_id)

Makes user admin

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
user_id = 2 # Integer | User ID

begin
  # Makes user admin
  api_instance.project_project_id_users_user_id_admin_post(project_id, user_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_user_id_admin_post: #{e}"
end
```

#### Using the project_project_id_users_user_id_admin_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_users_user_id_admin_post_with_http_info(project_id, user_id)

```ruby
begin
  # Makes user admin
  data, status_code, headers = api_instance.project_project_id_users_user_id_admin_post_with_http_info(project_id, user_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_user_id_admin_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **user_id** | **Integer** | User ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_users_user_id_delete

> project_project_id_users_user_id_delete(project_id, user_id)

Removes user from project

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
user_id = 2 # Integer | User ID

begin
  # Removes user from project
  api_instance.project_project_id_users_user_id_delete(project_id, user_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_user_id_delete: #{e}"
end
```

#### Using the project_project_id_users_user_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_users_user_id_delete_with_http_info(project_id, user_id)

```ruby
begin
  # Removes user from project
  data, status_code, headers = api_instance.project_project_id_users_user_id_delete_with_http_info(project_id, user_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_users_user_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **user_id** | **Integer** | User ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_views_get

> <Array<View>> project_project_id_views_get(project_id)

Get view

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID

begin
  # Get view
  result = api_instance.project_project_id_views_get(project_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_get: #{e}"
end
```

#### Using the project_project_id_views_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<View>>, Integer, Hash)> project_project_id_views_get_with_http_info(project_id)

```ruby
begin
  # Get view
  data, status_code, headers = api_instance.project_project_id_views_get_with_http_info(project_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<View>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |

### Return type

[**Array&lt;View&gt;**](View.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_views_post

> <View> project_project_id_views_post(project_id, view)

create view

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
view = OpenapiClient::ViewRequest.new # ViewRequest | 

begin
  # create view
  result = api_instance.project_project_id_views_post(project_id, view)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_post: #{e}"
end
```

#### Using the project_project_id_views_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<View>, Integer, Hash)> project_project_id_views_post_with_http_info(project_id, view)

```ruby
begin
  # create view
  data, status_code, headers = api_instance.project_project_id_views_post_with_http_info(project_id, view)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <View>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **view** | [**ViewRequest**](ViewRequest.md) |  |  |

### Return type

[**View**](View.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_views_view_id_delete

> project_project_id_views_view_id_delete(project_id, view_id)

Removes view

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
view_id = 10 # Integer | view ID

begin
  # Removes view
  api_instance.project_project_id_views_view_id_delete(project_id, view_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_view_id_delete: #{e}"
end
```

#### Using the project_project_id_views_view_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_views_view_id_delete_with_http_info(project_id, view_id)

```ruby
begin
  # Removes view
  data, status_code, headers = api_instance.project_project_id_views_view_id_delete_with_http_info(project_id, view_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_view_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **view_id** | **Integer** | view ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## project_project_id_views_view_id_get

> <View> project_project_id_views_view_id_get(project_id, view_id)

Get view

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
view_id = 10 # Integer | view ID

begin
  # Get view
  result = api_instance.project_project_id_views_view_id_get(project_id, view_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_view_id_get: #{e}"
end
```

#### Using the project_project_id_views_view_id_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<View>, Integer, Hash)> project_project_id_views_view_id_get_with_http_info(project_id, view_id)

```ruby
begin
  # Get view
  data, status_code, headers = api_instance.project_project_id_views_view_id_get_with_http_info(project_id, view_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <View>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_view_id_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **view_id** | **Integer** | view ID |  |

### Return type

[**View**](View.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## project_project_id_views_view_id_put

> project_project_id_views_view_id_put(project_id, view_id, view)

Updates view

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

api_instance = OpenapiClient::ProjectApi.new
project_id = 1 # Integer | Project ID
view_id = 10 # Integer | view ID
view = OpenapiClient::ViewRequest.new # ViewRequest | 

begin
  # Updates view
  api_instance.project_project_id_views_view_id_put(project_id, view_id, view)
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_view_id_put: #{e}"
end
```

#### Using the project_project_id_views_view_id_put_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> project_project_id_views_view_id_put_with_http_info(project_id, view_id, view)

```ruby
begin
  # Updates view
  data, status_code, headers = api_instance.project_project_id_views_view_id_put_with_http_info(project_id, view_id, view)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProjectApi->project_project_id_views_view_id_put_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** | Project ID |  |
| **view_id** | **Integer** | view ID |  |
| **view** | [**ViewRequest**](ViewRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

