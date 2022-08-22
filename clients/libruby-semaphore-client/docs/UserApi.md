# OpenapiClient::UserApi

All URIs are relative to *https://demo.ansible-semaphore.com/api*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**user_get**](UserApi.md#user_get) | **GET** /user/ | Fetch logged in user |
| [**user_tokens_api_token_id_delete**](UserApi.md#user_tokens_api_token_id_delete) | **DELETE** /user/tokens/{api_token_id} | Expires API token |
| [**user_tokens_get**](UserApi.md#user_tokens_get) | **GET** /user/tokens | Fetch API tokens for user |
| [**user_tokens_post**](UserApi.md#user_tokens_post) | **POST** /user/tokens | Create an API token |
| [**users_get**](UserApi.md#users_get) | **GET** /users | Fetches all users |
| [**users_post**](UserApi.md#users_post) | **POST** /users | Creates a user |
| [**users_user_id_delete**](UserApi.md#users_user_id_delete) | **DELETE** /users/{user_id}/ | Deletes user |
| [**users_user_id_get**](UserApi.md#users_user_id_get) | **GET** /users/{user_id}/ | Fetches a user profile |
| [**users_user_id_password_post**](UserApi.md#users_user_id_password_post) | **POST** /users/{user_id}/password | Updates user password |
| [**users_user_id_put**](UserApi.md#users_user_id_put) | **PUT** /users/{user_id}/ | Updates user details |


## user_get

> <User> user_get

Fetch logged in user

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

api_instance = OpenapiClient::UserApi.new

begin
  # Fetch logged in user
  result = api_instance.user_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->user_get: #{e}"
end
```

#### Using the user_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<User>, Integer, Hash)> user_get_with_http_info

```ruby
begin
  # Fetch logged in user
  data, status_code, headers = api_instance.user_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <User>
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->user_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


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

api_instance = OpenapiClient::UserApi.new
api_token_id = 'kwofd61g93-yuqvex8efmhjkgnbxlo8mp1tin6spyhu=' # String | 

begin
  # Expires API token
  api_instance.user_tokens_api_token_id_delete(api_token_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->user_tokens_api_token_id_delete: #{e}"
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
  puts "Error when calling UserApi->user_tokens_api_token_id_delete_with_http_info: #{e}"
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

api_instance = OpenapiClient::UserApi.new

begin
  # Fetch API tokens for user
  result = api_instance.user_tokens_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->user_tokens_get: #{e}"
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
  puts "Error when calling UserApi->user_tokens_get_with_http_info: #{e}"
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

api_instance = OpenapiClient::UserApi.new

begin
  # Create an API token
  result = api_instance.user_tokens_post
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->user_tokens_post: #{e}"
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
  puts "Error when calling UserApi->user_tokens_post_with_http_info: #{e}"
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


## users_get

> <Array<User>> users_get

Fetches all users

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

api_instance = OpenapiClient::UserApi.new

begin
  # Fetches all users
  result = api_instance.users_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_get: #{e}"
end
```

#### Using the users_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<User>>, Integer, Hash)> users_get_with_http_info

```ruby
begin
  # Fetches all users
  data, status_code, headers = api_instance.users_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<User>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Array&lt;User&gt;**](User.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## users_post

> <User> users_post(user)

Creates a user

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

api_instance = OpenapiClient::UserApi.new
user = OpenapiClient::UserRequest.new # UserRequest | 

begin
  # Creates a user
  result = api_instance.users_post(user)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_post: #{e}"
end
```

#### Using the users_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<User>, Integer, Hash)> users_post_with_http_info(user)

```ruby
begin
  # Creates a user
  data, status_code, headers = api_instance.users_post_with_http_info(user)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <User>
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **user** | [**UserRequest**](UserRequest.md) |  |  |

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8


## users_user_id_delete

> users_user_id_delete(user_id)

Deletes user

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

api_instance = OpenapiClient::UserApi.new
user_id = 2 # Integer | User ID

begin
  # Deletes user
  api_instance.users_user_id_delete(user_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_user_id_delete: #{e}"
end
```

#### Using the users_user_id_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> users_user_id_delete_with_http_info(user_id)

```ruby
begin
  # Deletes user
  data, status_code, headers = api_instance.users_user_id_delete_with_http_info(user_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_user_id_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **user_id** | **Integer** | User ID |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## users_user_id_get

> <User> users_user_id_get(user_id)

Fetches a user profile

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

api_instance = OpenapiClient::UserApi.new
user_id = 2 # Integer | User ID

begin
  # Fetches a user profile
  result = api_instance.users_user_id_get(user_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_user_id_get: #{e}"
end
```

#### Using the users_user_id_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<User>, Integer, Hash)> users_user_id_get_with_http_info(user_id)

```ruby
begin
  # Fetches a user profile
  data, status_code, headers = api_instance.users_user_id_get_with_http_info(user_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <User>
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_user_id_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **user_id** | **Integer** | User ID |  |

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8


## users_user_id_password_post

> users_user_id_password_post(user_id, password)

Updates user password

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

api_instance = OpenapiClient::UserApi.new
user_id = 2 # Integer | User ID
password = OpenapiClient::UsersUserIdPasswordPostRequest.new # UsersUserIdPasswordPostRequest | 

begin
  # Updates user password
  api_instance.users_user_id_password_post(user_id, password)
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_user_id_password_post: #{e}"
end
```

#### Using the users_user_id_password_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> users_user_id_password_post_with_http_info(user_id, password)

```ruby
begin
  # Updates user password
  data, status_code, headers = api_instance.users_user_id_password_post_with_http_info(user_id, password)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_user_id_password_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **user_id** | **Integer** | User ID |  |
| **password** | [**UsersUserIdPasswordPostRequest**](UsersUserIdPasswordPostRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## users_user_id_put

> users_user_id_put(user_id, user)

Updates user details

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

api_instance = OpenapiClient::UserApi.new
user_id = 2 # Integer | User ID
user = OpenapiClient::UserPutRequest.new # UserPutRequest | 

begin
  # Updates user details
  api_instance.users_user_id_put(user_id, user)
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_user_id_put: #{e}"
end
```

#### Using the users_user_id_put_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> users_user_id_put_with_http_info(user_id, user)

```ruby
begin
  # Updates user details
  data, status_code, headers = api_instance.users_user_id_put_with_http_info(user_id, user)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling UserApi->users_user_id_put_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **user_id** | **Integer** | User ID |  |
| **user** | [**UserPutRequest**](UserPutRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

