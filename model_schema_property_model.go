/*
OpenProject API V3 (Stable)

You're looking at the current **stable** documentation of the OpenProject APIv3. If you're interested in the current development version, please go to [github.com/opf](https://github.com/opf/openproject/tree/dev/docs/api/apiv3).  ## Introduction  The documentation for the APIv3 is written according to the [OpenAPI 3.0 Specification](https://swagger.io/specification/). You can either view the static version of this documentation on the [website](https://www.openproject.org/docs/api/introduction/) or the interactive version, rendered with [OpenAPI Explorer](https://github.com/Rhosys/openapi-explorer/blob/main/README.md), in your OpenProject installation under `/api/docs`. In the latter you can try out the various API endpoints directly interacting with our OpenProject data. Moreover you can access the specification source itself under `/api/v3/spec.json` and `/api/v3/spec.yml` (e.g. [here](https://community.openproject.org/api/v3/spec.yml)).  The APIv3 is a hypermedia REST API, a shorthand for \"Hypermedia As The Engine Of Application State\" (HATEOAS). This means that each endpoint of this API will have links to other resources or actions defined in the resulting body.  These related resources and actions for any given resource will be context sensitive. For example, only actions that the authenticated user can take are being rendered. This can be used to dynamically identify actions that the user might take for any given response.  As an example, if you fetch a work package through the [Work Package endpoint](https://www.openproject.org/docs/api/endpoints/work-packages/), the `update` link will only be present when the user you authenticated has been granted a permission to update the work package in the assigned project.  ## HAL+JSON  HAL is a simple format that gives a consistent and easy way to hyperlink between resources in your API. Read more in the following specification: [https://tools.ietf.org/html/draft-kelly-json-hal-08](https://tools.ietf.org/html/draft-kelly-json-hal-08)  **OpenProject API implementation of HAL+JSON format** enriches JSON and introduces a few meta properties:  - `_type` - specifies the type of the resource (e.g.: WorkPackage, Project) - `_links` - contains all related resource and action links available for the resource - `_embedded` - contains all embedded objects  HAL does not guarantee that embedded resources are embedded in their full representation, they might as well be partially represented (e.g. some properties can be left out). However in this API you have the guarantee that whenever a resource is **embedded**, it is embedded in its **full representation**.  ## API response structure  All API responses contain a single HAL+JSON object, even collections of objects are technically represented by a single HAL+JSON object that itself contains its members. More details on collections can be found in the [Collections Section](https://www.openproject.org/docs/api/collections/).  ## Authentication  The API supports the following authentication schemes: OAuth2, session based authentication, and basic auth.  Depending on the settings of the OpenProject instance many resources can be accessed without being authenticated. In case the instance requires authentication on all requests the client will receive an **HTTP 401** status code in response to any request.  Otherwise unauthenticated clients have all the permissions of the anonymous user.  ### Session-based Authentication  This means you have to login to OpenProject via the Web-Interface to be authenticated in the API. This method is well-suited for clients acting within the browser, like the Angular-Client built into OpenProject.  In this case, you always need to pass the HTTP header `X-Requested-With \"XMLHttpRequest\"` for authentication.  ### API Key through Basic Auth  Users can authenticate towards the API v3 using basic auth with the user name `apikey` (NOT your login) and the API key as the password. Users can find their API key on their account page.  Example:  ```shell API_KEY=2519132cdf62dcf5a66fd96394672079f9e9cad1 curl -u apikey:$API_KEY https://community.openproject.org/api/v3/users/42 ```  ### OAuth2.0 authentication  OpenProject allows authentication and authorization with OAuth2 with *Authorization code flow*, as well as *Client credentials* operation modes.  To get started, you first need to register an application in the OpenProject OAuth administration section of your installation. This will save an entry for your application with a client unique identifier (`client_id`) and an accompanying secret key (`client_secret`).  You can then use one the following guides to perform the supported OAuth 2.0 flows:  - [Authorization code flow](https://oauth.net/2/grant-types/authorization-code)  - [Authorization code flow with PKCE](https://doorkeeper.gitbook.io/guides/ruby-on-rails/pkce-flow), recommended for clients unable to keep the client_secret confidential.  - [Client credentials](https://oauth.net/2/grant-types/client-credentials/) - Requires an application to be bound to an impersonating user for non-public access  ### Why not username and password?  The simplest way to do basic auth would be to use a user's username and password naturally. However, OpenProject already has supported API keys in the past for the API v2, though not through basic auth.  Using **username and password** directly would have some advantages:  * It is intuitive for the user who then just has to provide those just as they would when logging into OpenProject.  * No extra logic for token management necessary.  On the other hand using **API keys** has some advantages too, which is why we went for that:  * If compromised while saved on an insecure client the user only has to regenerate the API key instead of changing their password, too.  * They are naturally long and random which makes them invulnerable to dictionary attacks and harder to crack in general.  Most importantly users may not actually have a password to begin with. Specifically when they have registered through an OpenID Connect provider.  ## Cross-Origin Resource Sharing (CORS)  By default, the OpenProject API is _not_ responding with any CORS headers. If you want to allow cross-domain AJAX calls against your OpenProject instance, you need to enable CORS headers being returned.  Please see [our API settings documentation](https://www.openproject.org/docs/system-admin-guide/api-and-webhooks/) on how to selectively enable CORS.  ## Allowed HTTP methods  - `GET` - Get a single resource or collection of resources  - `POST` - Create a new resource or perform  - `PATCH` - Update a resource  - `DELETE` - Delete a resource  ## Compression  Responses are compressed if requested by the client. Currently [gzip](https://www.gzip.org/) and [deflate](https://tools.ietf.org/html/rfc1951) are supported. The client signals the desired compression by setting the [`Accept-Encoding` header](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.3). If no `Accept-Encoding` header is send, `Accept-Encoding: identity` is assumed which will result in the API responding uncompressed.

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openproject

import (
	"encoding/json"
	"fmt"
)

// checks if the SchemaPropertyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaPropertyModel{}

// SchemaPropertyModel struct for SchemaPropertyModel
type SchemaPropertyModel struct {
	// The resource type for this property.
	Type string `json:"type"`
	// The name of the property.
	Name string `json:"name"`
	// Indicates, if the property is required for submitting a request of this schema.
	Required bool `json:"required"`
	// Indicates, if the property has a default.
	HasDefault bool `json:"hasDefault"`
	// Indicates, if the property is writable when sending a request of this schema.
	Writable bool `json:"writable"`
	// Additional options for the property.
	Object map[string]interface{} `json:"object,omitempty"`
	// Defines the json path where the property is located in the payload.
	Location *string `json:"location,omitempty"`
	// Useful links for this property (e.g. an endpoint to fetch allowed values)
	Links map[string]interface{} `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchemaPropertyModel SchemaPropertyModel

// NewSchemaPropertyModel instantiates a new SchemaPropertyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaPropertyModel(type_ string, name string, required bool, hasDefault bool, writable bool) *SchemaPropertyModel {
	this := SchemaPropertyModel{}
	this.Type = type_
	this.Name = name
	this.Required = required
	this.HasDefault = hasDefault
	this.Writable = writable
	var location string = ""
	this.Location = &location
	return &this
}

// NewSchemaPropertyModelWithDefaults instantiates a new SchemaPropertyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaPropertyModelWithDefaults() *SchemaPropertyModel {
	this := SchemaPropertyModel{}
	var location string = ""
	this.Location = &location
	return &this
}

// GetType returns the Type field value
func (o *SchemaPropertyModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SchemaPropertyModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SchemaPropertyModel) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *SchemaPropertyModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemaPropertyModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemaPropertyModel) SetName(v string) {
	o.Name = v
}

// GetRequired returns the Required field value
func (o *SchemaPropertyModel) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *SchemaPropertyModel) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *SchemaPropertyModel) SetRequired(v bool) {
	o.Required = v
}

// GetHasDefault returns the HasDefault field value
func (o *SchemaPropertyModel) GetHasDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasDefault
}

// GetHasDefaultOk returns a tuple with the HasDefault field value
// and a boolean to check if the value has been set.
func (o *SchemaPropertyModel) GetHasDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasDefault, true
}

// SetHasDefault sets field value
func (o *SchemaPropertyModel) SetHasDefault(v bool) {
	o.HasDefault = v
}

// GetWritable returns the Writable field value
func (o *SchemaPropertyModel) GetWritable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Writable
}

// GetWritableOk returns a tuple with the Writable field value
// and a boolean to check if the value has been set.
func (o *SchemaPropertyModel) GetWritableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Writable, true
}

// SetWritable sets field value
func (o *SchemaPropertyModel) SetWritable(v bool) {
	o.Writable = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *SchemaPropertyModel) GetObject() map[string]interface{} {
	if o == nil || IsNil(o.Object) {
		var ret map[string]interface{}
		return ret
	}
	return o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaPropertyModel) GetObjectOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return map[string]interface{}{}, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *SchemaPropertyModel) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given map[string]interface{} and assigns it to the Object field.
func (o *SchemaPropertyModel) SetObject(v map[string]interface{}) {
	o.Object = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SchemaPropertyModel) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaPropertyModel) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SchemaPropertyModel) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SchemaPropertyModel) SetLocation(v string) {
	o.Location = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SchemaPropertyModel) GetLinks() map[string]interface{} {
	if o == nil || IsNil(o.Links) {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaPropertyModel) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]interface{}{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SchemaPropertyModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *SchemaPropertyModel) SetLinks(v map[string]interface{}) {
	o.Links = v
}

func (o SchemaPropertyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaPropertyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["required"] = o.Required
	toSerialize["hasDefault"] = o.HasDefault
	toSerialize["writable"] = o.Writable
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchemaPropertyModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"required",
		"hasDefault",
		"writable",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSchemaPropertyModel := _SchemaPropertyModel{}

	err = json.Unmarshal(data, &varSchemaPropertyModel)

	if err != nil {
		return err
	}

	*o = SchemaPropertyModel(varSchemaPropertyModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "required")
		delete(additionalProperties, "hasDefault")
		delete(additionalProperties, "writable")
		delete(additionalProperties, "object")
		delete(additionalProperties, "location")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchemaPropertyModel struct {
	value *SchemaPropertyModel
	isSet bool
}

func (v NullableSchemaPropertyModel) Get() *SchemaPropertyModel {
	return v.value
}

func (v *NullableSchemaPropertyModel) Set(val *SchemaPropertyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaPropertyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaPropertyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaPropertyModel(val *SchemaPropertyModel) *NullableSchemaPropertyModel {
	return &NullableSchemaPropertyModel{value: val, isSet: true}
}

func (v NullableSchemaPropertyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaPropertyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


