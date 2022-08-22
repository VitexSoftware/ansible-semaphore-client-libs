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
import org.openapitools.client.model.Login;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for AuthenticationApi
 */
@Disabled
public class AuthenticationApiTest {

    private final AuthenticationApi api = new AuthenticationApi();

    /**
     * Performs Login
     *
     * Upon success you will be logged in
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void authLoginPostTest() throws ApiException {
        Login loginBody = null;
        api.authLoginPost(loginBody);
        // TODO: test validations
    }

    /**
     * Destroys current session
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void authLogoutPostTest() throws ApiException {
        api.authLogoutPost();
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

}