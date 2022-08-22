# OpenapiClient::Task

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  | [optional] |
| **template_id** | **Integer** |  | [optional] |
| **status** | **String** |  | [optional] |
| **debug** | **Boolean** |  | [optional] |
| **playbook** | **String** |  | [optional] |
| **environment** | **String** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::Task.new(
  id: 23,
  template_id: null,
  status: null,
  debug: null,
  playbook: null,
  environment: null
)
```

