# OpenapiClient::AuthenticationApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**auth_login_post**](AuthenticationApi.md#auth_login_post) | **POST** /auth/login | Performs Login |
| [**auth_logout_post**](AuthenticationApi.md#auth_logout_post) | **POST** /auth/logout | Destroys current session |
| [**user_tokens_api_token_id_delete**](AuthenticationApi.md#user_tokens_api_token_id_delete) | **DELETE** /user/tokens/{api_token_id} | Expires API token |
| [**user_tokens_get**](AuthenticationApi.md#user_tokens_get) | **GET** /user/tokens | Fetch API tokens for user |
| [**user_tokens_post**](AuthenticationApi.md#user_tokens_post) | **POST** /user/tokens | Create an API token |


## auth_login_post

> auth_login_post(login_body)

Performs Login

Upon success you will be logged in

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

api_instance = OpenapiClient::AuthenticationApi.new
login_body = OpenapiClient::Login.new # Login | 

begin
  # Performs Login
  api_instance.auth_login_post(login_body)
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->auth_login_post: #{e}"
end
```

#### Using the auth_login_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> auth_login_post_with_http_info(login_body)

```ruby
begin
  # Performs Login
  data, status_code, headers = api_instance.auth_login_post_with_http_info(login_body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->auth_login_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **login_body** | [**Login**](Login.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## auth_logout_post

> auth_logout_post

Destroys current session

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

api_instance = OpenapiClient::AuthenticationApi.new

begin
  # Destroys current session
  api_instance.auth_logout_post
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->auth_logout_post: #{e}"
end
```

#### Using the auth_logout_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> auth_logout_post_with_http_info

```ruby
begin
  # Destroys current session
  data, status_code, headers = api_instance.auth_logout_post_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->auth_logout_post_with_http_info: #{e}"
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


## user_tokens_api_token_id_delete

> user_tokens_api_token_id_delete(api_token_id)

Expires API token

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

api_instance = OpenapiClient::AuthenticationApi.new
api_token_id = 'kwofd61g93-yuqvex8efmhjkgnbxlo8mp1tin6spyhu=' # String | 

begin
  # Expires API token
  api_instance.user_tokens_api_token_id_delete(api_token_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->user_tokens_api_token_id_delete: #{e}"
end
```

#### Using the user_tokens_api_token_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> user_tokens_api_token_id_delete_with_http_info(api_token_id)

```ruby
begin
  # Expires API token
  data, status_code, headers = api_instance.user_tokens_api_token_id_delete_with_http_info(api_token_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->user_tokens_api_token_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **api_token_id** | **String** |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## user_tokens_get

> <Array<APIToken>> user_tokens_get

Fetch API tokens for user

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

api_instance = OpenapiClient::AuthenticationApi.new

begin
  # Fetch API tokens for user
  result = api_instance.user_tokens_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->user_tokens_get: #{e}"
end
```

#### Using the user_tokens_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<APIToken>>, Integer, Hash)> user_tokens_get_with_http_info

```ruby
begin
  # Fetch API tokens for user
  data, status_code, headers = api_instance.user_tokens_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<APIToken>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->user_tokens_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Array&lt;APIToken&gt;**](APIToken.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## user_tokens_post

> <APIToken> user_tokens_post

Create an API token

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

api_instance = OpenapiClient::AuthenticationApi.new

begin
  # Create an API token
  result = api_instance.user_tokens_post
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->user_tokens_post: #{e}"
end
```

#### Using the user_tokens_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<APIToken>, Integer, Hash)> user_tokens_post_with_http_info

```ruby
begin
  # Create an API token
  data, status_code, headers = api_instance.user_tokens_post_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <APIToken>
rescue OpenapiClient::ApiError => e
  puts "Error when calling AuthenticationApi->user_tokens_post_with_http_info: #{e}"
end
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

