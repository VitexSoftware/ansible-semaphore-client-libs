=begin
#API

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 2.8.34

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

# Common files
require 'openapi_client/api_client'
require 'openapi_client/api_error'
require 'openapi_client/version'
require 'openapi_client/configuration'

# Models
require 'openapi_client/models/api_token'
require 'openapi_client/models/access_key'
require 'openapi_client/models/access_key_request'
require 'openapi_client/models/environment'
require 'openapi_client/models/environment_request'
require 'openapi_client/models/event'
require 'openapi_client/models/info_type'
require 'openapi_client/models/info_type_update'
require 'openapi_client/models/inventory'
require 'openapi_client/models/inventory_request'
require 'openapi_client/models/login'
require 'openapi_client/models/project'
require 'openapi_client/models/project_project_id_delete_request'
require 'openapi_client/models/project_project_id_tasks_get_request'
require 'openapi_client/models/project_project_id_users_get_request'
require 'openapi_client/models/project_request'
require 'openapi_client/models/repository'
require 'openapi_client/models/repository_request'
require 'openapi_client/models/schedule'
require 'openapi_client/models/schedule_request'
require 'openapi_client/models/task'
require 'openapi_client/models/task_output'
require 'openapi_client/models/template'
require 'openapi_client/models/template_request'
require 'openapi_client/models/user'
require 'openapi_client/models/user_put_request'
require 'openapi_client/models/user_request'
require 'openapi_client/models/users_user_id_password_post_request'
require 'openapi_client/models/view'
require 'openapi_client/models/view_request'

# APIs
require 'openapi_client/api/authentication_api'
require 'openapi_client/api/default_api'
require 'openapi_client/api/project_api'
require 'openapi_client/api/projects_api'
require 'openapi_client/api/schedule_api'
require 'openapi_client/api/user_api'

module OpenapiClient
  class << self
    # Customize default settings for the SDK using block.
    #   OpenapiClient.configure do |config|
    #     config.username = "xxx"
    #     config.password = "xxx"
    #   end
    # If no block given, return the default Configuration object.
    def configure
      if block_given?
        yield(Configuration.default)
      else
        Configuration.default
      end
    end
  end
end