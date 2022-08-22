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
 * ProjectProjectIdTasksGetRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-08-22T22:16:56.905602471+02:00[Europe/Prague]")
public class ProjectProjectIdTasksGetRequest {
  public static final String SERIALIZED_NAME_TEMPLATE_ID = "template_id";
  @SerializedName(SERIALIZED_NAME_TEMPLATE_ID)
  private Integer templateId;

  public static final String SERIALIZED_NAME_DEBUG = "debug";
  @SerializedName(SERIALIZED_NAME_DEBUG)
  private Boolean debug;

  public static final String SERIALIZED_NAME_DRY_RUN = "dry_run";
  @SerializedName(SERIALIZED_NAME_DRY_RUN)
  private Boolean dryRun;

  public static final String SERIALIZED_NAME_PLAYBOOK = "playbook";
  @SerializedName(SERIALIZED_NAME_PLAYBOOK)
  private String playbook;

  public static final String SERIALIZED_NAME_ENVIRONMENT = "environment";
  @SerializedName(SERIALIZED_NAME_ENVIRONMENT)
  private String environment;

  public ProjectProjectIdTasksGetRequest() { 
  }

  public ProjectProjectIdTasksGetRequest templateId(Integer templateId) {
    
    this.templateId = templateId;
    return this;
  }

   /**
   * Get templateId
   * @return templateId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getTemplateId() {
    return templateId;
  }


  public void setTemplateId(Integer templateId) {
    this.templateId = templateId;
  }


  public ProjectProjectIdTasksGetRequest debug(Boolean debug) {
    
    this.debug = debug;
    return this;
  }

   /**
   * Get debug
   * @return debug
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDebug() {
    return debug;
  }


  public void setDebug(Boolean debug) {
    this.debug = debug;
  }


  public ProjectProjectIdTasksGetRequest dryRun(Boolean dryRun) {
    
    this.dryRun = dryRun;
    return this;
  }

   /**
   * Get dryRun
   * @return dryRun
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDryRun() {
    return dryRun;
  }


  public void setDryRun(Boolean dryRun) {
    this.dryRun = dryRun;
  }


  public ProjectProjectIdTasksGetRequest playbook(String playbook) {
    
    this.playbook = playbook;
    return this;
  }

   /**
   * Get playbook
   * @return playbook
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPlaybook() {
    return playbook;
  }


  public void setPlaybook(String playbook) {
    this.playbook = playbook;
  }


  public ProjectProjectIdTasksGetRequest environment(String environment) {
    
    this.environment = environment;
    return this;
  }

   /**
   * Get environment
   * @return environment
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEnvironment() {
    return environment;
  }


  public void setEnvironment(String environment) {
    this.environment = environment;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ProjectProjectIdTasksGetRequest projectProjectIdTasksGetRequest = (ProjectProjectIdTasksGetRequest) o;
    return Objects.equals(this.templateId, projectProjectIdTasksGetRequest.templateId) &&
        Objects.equals(this.debug, projectProjectIdTasksGetRequest.debug) &&
        Objects.equals(this.dryRun, projectProjectIdTasksGetRequest.dryRun) &&
        Objects.equals(this.playbook, projectProjectIdTasksGetRequest.playbook) &&
        Objects.equals(this.environment, projectProjectIdTasksGetRequest.environment);
  }

  @Override
  public int hashCode() {
    return Objects.hash(templateId, debug, dryRun, playbook, environment);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ProjectProjectIdTasksGetRequest {\n");
    sb.append("    templateId: ").append(toIndentedString(templateId)).append("\n");
    sb.append("    debug: ").append(toIndentedString(debug)).append("\n");
    sb.append("    dryRun: ").append(toIndentedString(dryRun)).append("\n");
    sb.append("    playbook: ").append(toIndentedString(playbook)).append("\n");
    sb.append("    environment: ").append(toIndentedString(environment)).append("\n");
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
    openapiFields.add("template_id");
    openapiFields.add("debug");
    openapiFields.add("dry_run");
    openapiFields.add("playbook");
    openapiFields.add("environment");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to ProjectProjectIdTasksGetRequest
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (ProjectProjectIdTasksGetRequest.openapiRequiredFields.isEmpty()) {
          return;
        } else { // has required fields
          throw new IllegalArgumentException(String.format("The required field(s) %s in ProjectProjectIdTasksGetRequest is not found in the empty JSON string", ProjectProjectIdTasksGetRequest.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!ProjectProjectIdTasksGetRequest.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `ProjectProjectIdTasksGetRequest` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if (jsonObj.get("playbook") != null && !jsonObj.get("playbook").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `playbook` to be a primitive type in the JSON string but got `%s`", jsonObj.get("playbook").toString()));
      }
      if (jsonObj.get("environment") != null && !jsonObj.get("environment").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `environment` to be a primitive type in the JSON string but got `%s`", jsonObj.get("environment").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ProjectProjectIdTasksGetRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ProjectProjectIdTasksGetRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ProjectProjectIdTasksGetRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ProjectProjectIdTasksGetRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<ProjectProjectIdTasksGetRequest>() {
           @Override
           public void write(JsonWriter out, ProjectProjectIdTasksGetRequest value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public ProjectProjectIdTasksGetRequest read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of ProjectProjectIdTasksGetRequest given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ProjectProjectIdTasksGetRequest
  * @throws IOException if the JSON string is invalid with respect to ProjectProjectIdTasksGetRequest
  */
  public static ProjectProjectIdTasksGetRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ProjectProjectIdTasksGetRequest.class);
  }

 /**
  * Convert an instance of ProjectProjectIdTasksGetRequest to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}
