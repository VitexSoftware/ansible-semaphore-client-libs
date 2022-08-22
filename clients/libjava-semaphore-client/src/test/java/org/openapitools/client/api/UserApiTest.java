/*
 * API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.8.34
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.APIToken;
import org.openapitools.client.model.User;
import org.openapitools.client.model.UserPutRequest;
import org.openapitools.client.model.UserRequest;
import org.openapitools.client.model.UsersUserIdPasswordPostRequest;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for UserApi
 */
@Disabled
public class UserApiTest {

    private final UserApi api = new UserApi();

    /**
     * Fetch logged in user
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void userGetTest() throws ApiException {
        User response = api.userGet();
        // TODO: test validations
    }

    /**
     * Expires API token
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void userTokensApiTokenIdDeleteTest() throws ApiException {
        String apiTokenId = null;
        api.userTokensApiTokenIdDelete(apiTokenId);
        // TODO: test validations
    }

    /**
     * Fetch API tokens for user
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void userTokensGetTest() throws ApiException {
        List<APIToken> response = api.userTokensGet();
        // TODO: test validations
    }

    /**
     * Create an API token
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void userTokensPostTest() throws ApiException {
        APIToken response = api.userTokensPost();
        // TODO: test validations
    }

    /**
     * Fetches all users
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void usersGetTest() throws ApiException {
        List<User> response = api.usersGet();
        // TODO: test validations
    }

    /**
     * Creates a user
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void usersPostTest() throws ApiException {
        UserRequest user = null;
        User response = api.usersPost(user);
        // TODO: test validations
    }

    /**
     * Deletes user
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void usersUserIdDeleteTest() throws ApiException {
        Integer userId = null;
        api.usersUserIdDelete(userId);
        // TODO: test validations
    }

    /**
     * Fetches a user profile
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void usersUserIdGetTest() throws ApiException {
        Integer userId = null;
        User response = api.usersUserIdGet(userId);
        // TODO: test validations
    }

    /**
     * Updates user password
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void usersUserIdPasswordPostTest() throws ApiException {
        Integer userId = null;
        UsersUserIdPasswordPostRequest password = null;
        api.usersUserIdPasswordPost(userId, password);
        // TODO: test validations
    }

    /**
     * Updates user details
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void usersUserIdPutTest() throws ApiException {
        Integer userId = null;
        UserPutRequest user = null;
        api.usersUserIdPut(userId, user);
        // TODO: test validations
    }

}
