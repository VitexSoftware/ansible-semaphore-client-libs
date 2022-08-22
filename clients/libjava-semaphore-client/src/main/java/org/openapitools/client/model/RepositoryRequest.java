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


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * RepositoryRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-08-22T22:16:56.905602471+02:00[Europe/Prague]")
public class RepositoryRequest {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_PROJECT_ID = "project_id";
  @SerializedName(SERIALIZED_NAME_PROJECT_ID)
  private Integer projectId;

  public static final String SERIALIZED_NAME_GIT_URL = "git_url";
  @SerializedName(SERIALIZED_NAME_GIT_URL)
  private String gitUrl;

  public static final String SERIALIZED_NAME_GIT_BRANCH = "git_branch";
  @SerializedName(SERIALIZED_NAME_GIT_BRANCH)
  private String gitBranch;

  public static final String SERIALIZED_NAME_SSH_KEY_ID = "ssh_key_id";
  @SerializedName(SERIALIZED_NAME_SSH_KEY_ID)
  private Integer sshKeyId;

  public RepositoryRequest() { 
  }

  public RepositoryRequest name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "Test", value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public RepositoryRequest projectId(Integer projectId) {
    
    this.projectId = projectId;
    return this;
  }

   /**
   * Get projectId
   * @return projectId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getProjectId() {
    return projectId;
  }


  public void setProjectId(Integer projectId) {
    this.projectId = projectId;
  }


  public RepositoryRequest gitUrl(String gitUrl) {
    
    this.gitUrl = gitUrl;
    return this;
  }

   /**
   * Get gitUrl
   * @return gitUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "git@example.com", value = "")

  public String getGitUrl() {
    return gitUrl;
  }


  public void setGitUrl(String gitUrl) {
    this.gitUrl = gitUrl;
  }


  public RepositoryRequest gitBranch(String gitBranch) {
    
    this.gitBranch = gitBranch;
    return this;
  }

   /**
   * Get gitBranch
   * @return gitBranch
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "master", value = "")

  public String getGitBranch() {
    return gitBranch;
  }


  public void setGitBranch(String gitBranch) {
    this.gitBranch = gitBranch;
  }


  public RepositoryRequest sshKeyId(Integer sshKeyId) {
    
    this.sshKeyId = sshKeyId;
    return this;
  }

   /**
   * Get sshKeyId
   * @return sshKeyId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getSshKeyId() {
    return sshKeyId;
  }


  public void setSshKeyId(Integer sshKeyId) {
    this.sshKeyId = sshKeyId;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RepositoryRequest repositoryRequest = (RepositoryRequest) o;
    return Objects.equals(this.name, repositoryRequest.name) &&
        Objects.equals(this.projectId, repositoryRequest.projectId) &&
        Objects.equals(this.gitUrl, repositoryRequest.gitUrl) &&
        Objects.equals(this.gitBranch, repositoryRequest.gitBranch) &&
        Objects.equals(this.sshKeyId, repositoryRequest.sshKeyId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, projectId, gitUrl, gitBranch, sshKeyId);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RepositoryRequest {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    projectId: ").append(toIndentedString(projectId)).append("\n");
    sb.append("    gitUrl: ").append(toIndentedString(gitUrl)).append("\n");
    sb.append("    gitBranch: ").append(toIndentedString(gitBranch)).append("\n");
    sb.append("    sshKeyId: ").append(toIndentedString(sshKeyId)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("name");
    openapiFields.add("project_id");
    openapiFields.add("git_url");
    openapiFields.add("git_branch");
    openapiFields.add("ssh_key_id");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to RepositoryRequest
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (RepositoryRequest.openapiRequiredFields.isEmpty()) {
          return;
        } else { // has required fields
          throw new IllegalArgumentException(String.format("The required field(s) %s in RepositoryRequest is not found in the empty JSON string", RepositoryRequest.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!RepositoryRequest.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `RepositoryRequest` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if (jsonObj.get("name") != null && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if (jsonObj.get("git_url") != null && !jsonObj.get("git_url").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `git_url` to be a primitive type in the JSON string but got `%s`", jsonObj.get("git_url").toString()));
      }
      if (jsonObj.get("git_branch") != null && !jsonObj.get("git_branch").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `git_branch` to be a primitive type in the JSON string but got `%s`", jsonObj.get("git_branch").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!RepositoryRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'RepositoryRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<RepositoryRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(RepositoryRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<RepositoryRequest>() {
           @Override
           public void write(JsonWriter out, RepositoryRequest value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public RepositoryRequest read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of RepositoryRequest given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of RepositoryRequest
  * @throws IOException if the JSON string is invalid with respect to RepositoryRequest
  */
  public static RepositoryRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, RepositoryRequest.class);
  }

 /**
  * Convert an instance of RepositoryRequest to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}
