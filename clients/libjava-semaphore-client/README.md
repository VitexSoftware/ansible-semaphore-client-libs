# semaphore

API
- API version: 2.8.34
  - Build date: 2022-08-22T22:16:56.905602471+02:00[Europe/Prague]

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)


*Automatically generated by the [OpenAPI Generator](https://openapi-generator.tech)*


## Requirements

Building the API client library requires:
1. Java 1.8+
2. Maven (3.8.3+)/Gradle (7.2+)

## Installation

To install the API client library to your local Maven repository, simply execute:

```shell
mvn clean install
```

To deploy it to a remote Maven repository instead, configure the settings of the repository and execute:

```shell
mvn clean deploy
```

Refer to the [OSSRH Guide](http://central.sonatype.org/pages/ossrh-guide.html) for more information.

### Maven users

Add this dependency to your project's POM:

```xml
<dependency>
  <groupId>org.openapitools</groupId>
  <artifactId>semaphore</artifactId>
  <version>2.8.34</version>
  <scope>compile</scope>
</dependency>
```

### Gradle users

Add this dependency to your project's build file:

```groovy
  repositories {
    mavenCentral()     // Needed if the 'semaphore' jar has been published to maven central.
    mavenLocal()       // Needed if the 'semaphore' jar has been published to the local maven repo.
  }

  dependencies {
     implementation "org.openapitools:semaphore:2.8.34"
  }
```

### Others

At first generate the JAR by executing:

```shell
mvn clean package
```

Then manually install the following JARs:

* `target/semaphore-2.8.34.jar`
* `target/lib/*.jar`

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Java code:

```java

// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://demo.ansible-semaphore.com/api");
    
    // Configure API key authorization: bearer
    ApiKeyAuth bearer = (ApiKeyAuth) defaultClient.getAuthentication("bearer");
    bearer.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //bearer.setApiKeyPrefix("Token");

    // Configure API key authorization: cookie
    ApiKeyAuth cookie = (ApiKeyAuth) defaultClient.getAuthentication("cookie");
    cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //cookie.setApiKeyPrefix("Token");

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    Login loginBody = new Login(); // Login | 
    try {
      apiInstance.authLoginPost(loginBody);
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#authLoginPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

## Documentation for API Endpoints

All URIs are relative to *https://demo.ansible-semaphore.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthenticationApi* | [**authLoginPost**](docs/AuthenticationApi.md#authLoginPost) | **POST** /auth/login | Performs Login
*AuthenticationApi* | [**authLogoutPost**](docs/AuthenticationApi.md#authLogoutPost) | **POST** /auth/logout | Destroys current session
*AuthenticationApi* | [**userTokensApiTokenIdDelete**](docs/AuthenticationApi.md#userTokensApiTokenIdDelete) | **DELETE** /user/tokens/{api_token_id} | Expires API token
*AuthenticationApi* | [**userTokensGet**](docs/AuthenticationApi.md#userTokensGet) | **GET** /user/tokens | Fetch API tokens for user
*AuthenticationApi* | [**userTokensPost**](docs/AuthenticationApi.md#userTokensPost) | **POST** /user/tokens | Create an API token
*DefaultApi* | [**eventsGet**](docs/DefaultApi.md#eventsGet) | **GET** /events | Get Events related to Semaphore and projects you are part of
*DefaultApi* | [**eventsLastGet**](docs/DefaultApi.md#eventsLastGet) | **GET** /events/last | Get last 200 Events related to Semaphore and projects you are part of
*DefaultApi* | [**infoGet**](docs/DefaultApi.md#infoGet) | **GET** /info | Fetches information about semaphore
*DefaultApi* | [**pingGet**](docs/DefaultApi.md#pingGet) | **GET** /ping | PING test
*DefaultApi* | [**wsGet**](docs/DefaultApi.md#wsGet) | **GET** /ws | Websocket handler
*ProjectApi* | [**projectProjectIdDelete**](docs/ProjectApi.md#projectProjectIdDelete) | **DELETE** /project/{project_id}/ | Delete project
*ProjectApi* | [**projectProjectIdEnvironmentEnvironmentIdDelete**](docs/ProjectApi.md#projectProjectIdEnvironmentEnvironmentIdDelete) | **DELETE** /project/{project_id}/environment/{environment_id} | Removes environment
*ProjectApi* | [**projectProjectIdEnvironmentEnvironmentIdPut**](docs/ProjectApi.md#projectProjectIdEnvironmentEnvironmentIdPut) | **PUT** /project/{project_id}/environment/{environment_id} | Update environment
*ProjectApi* | [**projectProjectIdEnvironmentGet**](docs/ProjectApi.md#projectProjectIdEnvironmentGet) | **GET** /project/{project_id}/environment | Get environment
*ProjectApi* | [**projectProjectIdEnvironmentPost**](docs/ProjectApi.md#projectProjectIdEnvironmentPost) | **POST** /project/{project_id}/environment | Add environment
*ProjectApi* | [**projectProjectIdEventsGet**](docs/ProjectApi.md#projectProjectIdEventsGet) | **GET** /project/{project_id}/events | Get Events related to this project
*ProjectApi* | [**projectProjectIdGet**](docs/ProjectApi.md#projectProjectIdGet) | **GET** /project/{project_id}/ | Fetch project
*ProjectApi* | [**projectProjectIdInventoryGet**](docs/ProjectApi.md#projectProjectIdInventoryGet) | **GET** /project/{project_id}/inventory | Get inventory
*ProjectApi* | [**projectProjectIdInventoryInventoryIdDelete**](docs/ProjectApi.md#projectProjectIdInventoryInventoryIdDelete) | **DELETE** /project/{project_id}/inventory/{inventory_id} | Removes inventory
*ProjectApi* | [**projectProjectIdInventoryInventoryIdPut**](docs/ProjectApi.md#projectProjectIdInventoryInventoryIdPut) | **PUT** /project/{project_id}/inventory/{inventory_id} | Updates inventory
*ProjectApi* | [**projectProjectIdInventoryPost**](docs/ProjectApi.md#projectProjectIdInventoryPost) | **POST** /project/{project_id}/inventory | create inventory
*ProjectApi* | [**projectProjectIdKeysGet**](docs/ProjectApi.md#projectProjectIdKeysGet) | **GET** /project/{project_id}/keys | Get access keys linked to project
*ProjectApi* | [**projectProjectIdKeysKeyIdDelete**](docs/ProjectApi.md#projectProjectIdKeysKeyIdDelete) | **DELETE** /project/{project_id}/keys/{key_id} | Removes access key
*ProjectApi* | [**projectProjectIdKeysKeyIdPut**](docs/ProjectApi.md#projectProjectIdKeysKeyIdPut) | **PUT** /project/{project_id}/keys/{key_id} | Updates access key
*ProjectApi* | [**projectProjectIdKeysPost**](docs/ProjectApi.md#projectProjectIdKeysPost) | **POST** /project/{project_id}/keys | Add access key
*ProjectApi* | [**projectProjectIdPut**](docs/ProjectApi.md#projectProjectIdPut) | **PUT** /project/{project_id}/ | Update project
*ProjectApi* | [**projectProjectIdRepositoriesGet**](docs/ProjectApi.md#projectProjectIdRepositoriesGet) | **GET** /project/{project_id}/repositories | Get repositories
*ProjectApi* | [**projectProjectIdRepositoriesPost**](docs/ProjectApi.md#projectProjectIdRepositoriesPost) | **POST** /project/{project_id}/repositories | Add repository
*ProjectApi* | [**projectProjectIdRepositoriesRepositoryIdDelete**](docs/ProjectApi.md#projectProjectIdRepositoriesRepositoryIdDelete) | **DELETE** /project/{project_id}/repositories/{repository_id} | Removes repository
*ProjectApi* | [**projectProjectIdTasksGet**](docs/ProjectApi.md#projectProjectIdTasksGet) | **GET** /project/{project_id}/tasks | Get Tasks related to current project
*ProjectApi* | [**projectProjectIdTasksLastGet**](docs/ProjectApi.md#projectProjectIdTasksLastGet) | **GET** /project/{project_id}/tasks/last | Get last 200 Tasks related to current project
*ProjectApi* | [**projectProjectIdTasksPost**](docs/ProjectApi.md#projectProjectIdTasksPost) | **POST** /project/{project_id}/tasks | Starts a job
*ProjectApi* | [**projectProjectIdTasksTaskIdDelete**](docs/ProjectApi.md#projectProjectIdTasksTaskIdDelete) | **DELETE** /project/{project_id}/tasks/{task_id} | Deletes task (including output)
*ProjectApi* | [**projectProjectIdTasksTaskIdGet**](docs/ProjectApi.md#projectProjectIdTasksTaskIdGet) | **GET** /project/{project_id}/tasks/{task_id} | Get a single task
*ProjectApi* | [**projectProjectIdTasksTaskIdOutputGet**](docs/ProjectApi.md#projectProjectIdTasksTaskIdOutputGet) | **GET** /project/{project_id}/tasks/{task_id}/output | Get task output
*ProjectApi* | [**projectProjectIdTemplatesGet**](docs/ProjectApi.md#projectProjectIdTemplatesGet) | **GET** /project/{project_id}/templates | Get template
*ProjectApi* | [**projectProjectIdTemplatesPost**](docs/ProjectApi.md#projectProjectIdTemplatesPost) | **POST** /project/{project_id}/templates | create template
*ProjectApi* | [**projectProjectIdTemplatesTemplateIdDelete**](docs/ProjectApi.md#projectProjectIdTemplatesTemplateIdDelete) | **DELETE** /project/{project_id}/templates/{template_id} | Removes template
*ProjectApi* | [**projectProjectIdTemplatesTemplateIdGet**](docs/ProjectApi.md#projectProjectIdTemplatesTemplateIdGet) | **GET** /project/{project_id}/templates/{template_id} | Get template
*ProjectApi* | [**projectProjectIdTemplatesTemplateIdPut**](docs/ProjectApi.md#projectProjectIdTemplatesTemplateIdPut) | **PUT** /project/{project_id}/templates/{template_id} | Updates template
*ProjectApi* | [**projectProjectIdUsersGet**](docs/ProjectApi.md#projectProjectIdUsersGet) | **GET** /project/{project_id}/users | Get users linked to project
*ProjectApi* | [**projectProjectIdUsersPost**](docs/ProjectApi.md#projectProjectIdUsersPost) | **POST** /project/{project_id}/users | Link user to project
*ProjectApi* | [**projectProjectIdUsersUserIdAdminDelete**](docs/ProjectApi.md#projectProjectIdUsersUserIdAdminDelete) | **DELETE** /project/{project_id}/users/{user_id}/admin | Revoke admin privileges
*ProjectApi* | [**projectProjectIdUsersUserIdAdminPost**](docs/ProjectApi.md#projectProjectIdUsersUserIdAdminPost) | **POST** /project/{project_id}/users/{user_id}/admin | Makes user admin
*ProjectApi* | [**projectProjectIdUsersUserIdDelete**](docs/ProjectApi.md#projectProjectIdUsersUserIdDelete) | **DELETE** /project/{project_id}/users/{user_id} | Removes user from project
*ProjectApi* | [**projectProjectIdViewsGet**](docs/ProjectApi.md#projectProjectIdViewsGet) | **GET** /project/{project_id}/views | Get view
*ProjectApi* | [**projectProjectIdViewsPost**](docs/ProjectApi.md#projectProjectIdViewsPost) | **POST** /project/{project_id}/views | create view
*ProjectApi* | [**projectProjectIdViewsViewIdDelete**](docs/ProjectApi.md#projectProjectIdViewsViewIdDelete) | **DELETE** /project/{project_id}/views/{view_id} | Removes view
*ProjectApi* | [**projectProjectIdViewsViewIdGet**](docs/ProjectApi.md#projectProjectIdViewsViewIdGet) | **GET** /project/{project_id}/views/{view_id} | Get view
*ProjectApi* | [**projectProjectIdViewsViewIdPut**](docs/ProjectApi.md#projectProjectIdViewsViewIdPut) | **PUT** /project/{project_id}/views/{view_id} | Updates view
*ProjectsApi* | [**projectsGet**](docs/ProjectsApi.md#projectsGet) | **GET** /projects | Get projects
*ProjectsApi* | [**projectsPost**](docs/ProjectsApi.md#projectsPost) | **POST** /projects | Create a new project
*ScheduleApi* | [**projectProjectIdSchedulesPost**](docs/ScheduleApi.md#projectProjectIdSchedulesPost) | **POST** /project/{project_id}/schedules | create schedule
*ScheduleApi* | [**projectProjectIdSchedulesScheduleIdDelete**](docs/ScheduleApi.md#projectProjectIdSchedulesScheduleIdDelete) | **DELETE** /project/{project_id}/schedules/{schedule_id} | Deletes schedule
*ScheduleApi* | [**projectProjectIdSchedulesScheduleIdGet**](docs/ScheduleApi.md#projectProjectIdSchedulesScheduleIdGet) | **GET** /project/{project_id}/schedules/{schedule_id} | Get schedule
*ScheduleApi* | [**projectProjectIdSchedulesScheduleIdPut**](docs/ScheduleApi.md#projectProjectIdSchedulesScheduleIdPut) | **PUT** /project/{project_id}/schedules/{schedule_id} | Updates schedule
*UserApi* | [**userGet**](docs/UserApi.md#userGet) | **GET** /user/ | Fetch logged in user
*UserApi* | [**userTokensApiTokenIdDelete**](docs/UserApi.md#userTokensApiTokenIdDelete) | **DELETE** /user/tokens/{api_token_id} | Expires API token
*UserApi* | [**userTokensGet**](docs/UserApi.md#userTokensGet) | **GET** /user/tokens | Fetch API tokens for user
*UserApi* | [**userTokensPost**](docs/UserApi.md#userTokensPost) | **POST** /user/tokens | Create an API token
*UserApi* | [**usersGet**](docs/UserApi.md#usersGet) | **GET** /users | Fetches all users
*UserApi* | [**usersPost**](docs/UserApi.md#usersPost) | **POST** /users | Creates a user
*UserApi* | [**usersUserIdDelete**](docs/UserApi.md#usersUserIdDelete) | **DELETE** /users/{user_id}/ | Deletes user
*UserApi* | [**usersUserIdGet**](docs/UserApi.md#usersUserIdGet) | **GET** /users/{user_id}/ | Fetches a user profile
*UserApi* | [**usersUserIdPasswordPost**](docs/UserApi.md#usersUserIdPasswordPost) | **POST** /users/{user_id}/password | Updates user password
*UserApi* | [**usersUserIdPut**](docs/UserApi.md#usersUserIdPut) | **PUT** /users/{user_id}/ | Updates user details


## Documentation for Models

 - [APIToken](docs/APIToken.md)
 - [AccessKey](docs/AccessKey.md)
 - [AccessKeyRequest](docs/AccessKeyRequest.md)
 - [Environment](docs/Environment.md)
 - [EnvironmentRequest](docs/EnvironmentRequest.md)
 - [Event](docs/Event.md)
 - [InfoType](docs/InfoType.md)
 - [InfoTypeUpdate](docs/InfoTypeUpdate.md)
 - [Inventory](docs/Inventory.md)
 - [InventoryRequest](docs/InventoryRequest.md)
 - [Login](docs/Login.md)
 - [Project](docs/Project.md)
 - [ProjectProjectIdDeleteRequest](docs/ProjectProjectIdDeleteRequest.md)
 - [ProjectProjectIdTasksGetRequest](docs/ProjectProjectIdTasksGetRequest.md)
 - [ProjectProjectIdUsersGetRequest](docs/ProjectProjectIdUsersGetRequest.md)
 - [ProjectRequest](docs/ProjectRequest.md)
 - [Repository](docs/Repository.md)
 - [RepositoryRequest](docs/RepositoryRequest.md)
 - [Schedule](docs/Schedule.md)
 - [ScheduleRequest](docs/ScheduleRequest.md)
 - [Task](docs/Task.md)
 - [TaskOutput](docs/TaskOutput.md)
 - [Template](docs/Template.md)
 - [TemplateRequest](docs/TemplateRequest.md)
 - [User](docs/User.md)
 - [UserPutRequest](docs/UserPutRequest.md)
 - [UserRequest](docs/UserRequest.md)
 - [UsersUserIdPasswordPostRequest](docs/UsersUserIdPasswordPostRequest.md)
 - [View](docs/View.md)
 - [ViewRequest](docs/ViewRequest.md)


## Documentation for Authorization

Authentication schemes defined for the API:
### bearer

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

### cookie

- **Type**: API key
- **API key parameter name**: Cookie
- **Location**: HTTP header


## Recommendation

It's recommended to create an instance of `ApiClient` per thread in a multithreaded environment to avoid any potential issues.

## Author


