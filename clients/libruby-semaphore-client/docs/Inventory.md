# OpenapiClient::Inventory

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **project_id** | **Integer** |  | [optional] |
| **inventory** | **String** |  | [optional] |
| **ssh_key_id** | **Integer** |  | [optional] |
| **become_key_id** | **Integer** |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::Inventory.new(
  id: null,
  name: Test,
  project_id: null,
  inventory: null,
  ssh_key_id: null,
  become_key_id: null,
  type: null
)
```

