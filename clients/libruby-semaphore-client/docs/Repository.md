# OpenapiClient::Repository

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **project_id** | **Integer** |  | [optional] |
| **git_url** | **String** |  | [optional] |
| **git_branch** | **String** |  | [optional] |
| **ssh_key_id** | **Integer** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::Repository.new(
  id: null,
  name: Test,
  project_id: null,
  git_url: git@example.com,
  git_branch: master,
  ssh_key_id: null
)
```

