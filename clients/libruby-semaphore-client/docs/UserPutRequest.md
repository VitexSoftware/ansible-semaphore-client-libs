# OpenapiClient::UserPutRequest

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

instance = OpenapiClient::UserPutRequest.new(
  name: Integration Test User2,
  username: test-user2,
  email: test2@ansiblesemaphore.test,
  alert: null,
  admin: null
)
```

