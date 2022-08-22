# OpenapiClient::TemplateRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **Integer** |  | [optional] |
| **inventory_id** | **Integer** |  | [optional] |
| **repository_id** | **Integer** |  | [optional] |
| **environment_id** | **Integer** |  | [optional] |
| **view_id** | **Integer** |  | [optional] |
| **_alias** | **String** |  | [optional] |
| **playbook** | **String** |  | [optional] |
| **arguments** | **String** |  | [optional] |
| **description** | **String** |  | [optional] |
| **override_args** | **Boolean** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::TemplateRequest.new(
  project_id: null,
  inventory_id: null,
  repository_id: null,
  environment_id: null,
  view_id: null,
  _alias: Test,
  playbook: test.yml,
  arguments: [],
  description: Hello, World!,
  override_args: null
)
```

