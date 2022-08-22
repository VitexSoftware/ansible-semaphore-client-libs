# OpenapiClient::ScheduleRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  | [optional] |
| **cron_format** | **String** |  | [optional] |
| **project_id** | **Integer** |  | [optional] |
| **template_id** | **Integer** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::ScheduleRequest.new(
  id: null,
  cron_format: * * * 1 *,
  project_id: null,
  template_id: null
)
```

