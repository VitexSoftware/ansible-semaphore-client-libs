# OpenapiClient::UserRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **name** | **String** |  | [optional] |
| **username** | **String** |  | [optional] |
| **email** | **String** |  | [optional] |
| **alert** | **Boolean** |  | [optional] |
| **admin** | **Boolean** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::UserRequest.new(
  name: Integration Test User,
  username: test-user,
  email: test@ansiblesemaphore.test,
  alert: null,
  admin: null
)
```

