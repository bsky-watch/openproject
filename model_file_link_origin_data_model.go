/*
OpenProject API V3 (Stable)

You're looking at the current **stable** documentation of the OpenProject APIv3. If you're interested in the current development version, please go to [github.com/opf](https://github.com/opf/openproject/tree/dev/docs/api/apiv3).  ## Introduction  The documentation for the APIv3 is written according to the [OpenAPI 3.0 Specification](https://swagger.io/specification/). You can either view the static version of this documentation on the [website](https://www.openproject.org/docs/api/introduction/) or the interactive version, rendered with [OpenAPI Explorer](https://github.com/Rhosys/openapi-explorer/blob/main/README.md), in your OpenProject installation under `/api/docs`. In the latter you can try out the various API endpoints directly interacting with our OpenProject data. Moreover you can access the specification source itself under `/api/v3/spec.json` and `/api/v3/spec.yml` (e.g. [here](https://community.openproject.org/api/v3/spec.yml)).  The APIv3 is a hypermedia REST API, a shorthand for \"Hypermedia As The Engine Of Application State\" (HATEOAS). This means that each endpoint of this API will have links to other resources or actions defined in the resulting body.  These related resources and actions for any given resource will be context sensitive. For example, only actions that the authenticated user can take are being rendered. This can be used to dynamically identify actions that the user might take for any given response.  As an example, if you fetch a work package through the [Work Package endpoint](https://www.openproject.org/docs/api/endpoints/work-packages/), the `update` link will only be present when the user you authenticated has been granted a permission to update the work package in the assigned project.  ## HAL+JSON  HAL is a simple format that gives a consistent and easy way to hyperlink between resources in your API. Read more in the following specification: [https://tools.ietf.org/html/draft-kelly-json-hal-08](https://tools.ietf.org/html/draft-kelly-json-hal-08)  **OpenProject API implementation of HAL+JSON format** enriches JSON and introduces a few meta properties:  - `_type` - specifies the type of the resource (e.g.: WorkPackage, Project) - `_links` - contains all related resource and action links available for the resource - `_embedded` - contains all embedded objects  HAL does not guarantee that embedded resources are embedded in their full representation, they might as well be partially represented (e.g. some properties can be left out). However in this API you have the guarantee that whenever a resource is **embedded**, it is embedded in its **full representation**.  ## API response structure  All API responses contain a single HAL+JSON object, even collections of objects are technically represented by a single HAL+JSON object that itself contains its members. More details on collections can be found in the [Collections Section](https://www.openproject.org/docs/api/collections/).  ## Authentication  The API supports the following authentication schemes: OAuth2, session based authentication, and basic auth.  Depending on the settings of the OpenProject instance many resources can be accessed without being authenticated. In case the instance requires authentication on all requests the client will receive an **HTTP 401** status code in response to any request.  Otherwise unauthenticated clients have all the permissions of the anonymous user.  ### Session-based Authentication  This means you have to login to OpenProject via the Web-Interface to be authenticated in the API. This method is well-suited for clients acting within the browser, like the Angular-Client built into OpenProject.  In this case, you always need to pass the HTTP header `X-Requested-With \"XMLHttpRequest\"` for authentication.  ### API Key through Basic Auth  Users can authenticate towards the API v3 using basic auth with the user name `apikey` (NOT your login) and the API key as the password. Users can find their API key on their account page.  Example:  ```shell API_KEY=2519132cdf62dcf5a66fd96394672079f9e9cad1 curl -u apikey:$API_KEY https://community.openproject.org/api/v3/users/42 ```  ### OAuth2.0 authentication  OpenProject allows authentication and authorization with OAuth2 with *Authorization code flow*, as well as *Client credentials* operation modes.  To get started, you first need to register an application in the OpenProject OAuth administration section of your installation. This will save an entry for your application with a client unique identifier (`client_id`) and an accompanying secret key (`client_secret`).  You can then use one the following guides to perform the supported OAuth 2.0 flows:  - [Authorization code flow](https://oauth.net/2/grant-types/authorization-code)  - [Authorization code flow with PKCE](https://doorkeeper.gitbook.io/guides/ruby-on-rails/pkce-flow), recommended for clients unable to keep the client_secret confidential.  - [Client credentials](https://oauth.net/2/grant-types/client-credentials/) - Requires an application to be bound to an impersonating user for non-public access  ### Why not username and password?  The simplest way to do basic auth would be to use a user's username and password naturally. However, OpenProject already has supported API keys in the past for the API v2, though not through basic auth.  Using **username and password** directly would have some advantages:  * It is intuitive for the user who then just has to provide those just as they would when logging into OpenProject.  * No extra logic for token management necessary.  On the other hand using **API keys** has some advantages too, which is why we went for that:  * If compromised while saved on an insecure client the user only has to regenerate the API key instead of changing their password, too.  * They are naturally long and random which makes them invulnerable to dictionary attacks and harder to crack in general.  Most importantly users may not actually have a password to begin with. Specifically when they have registered through an OpenID Connect provider.  ## Cross-Origin Resource Sharing (CORS)  By default, the OpenProject API is _not_ responding with any CORS headers. If you want to allow cross-domain AJAX calls against your OpenProject instance, you need to enable CORS headers being returned.  Please see [our API settings documentation](https://www.openproject.org/docs/system-admin-guide/api-and-webhooks/) on how to selectively enable CORS.  ## Allowed HTTP methods  - `GET` - Get a single resource or collection of resources  - `POST` - Create a new resource or perform  - `PATCH` - Update a resource  - `DELETE` - Delete a resource  ## Compression  Responses are compressed if requested by the client. Currently [gzip](https://www.gzip.org/) and [deflate](https://tools.ietf.org/html/rfc1951) are supported. The client signals the desired compression by setting the [`Accept-Encoding` header](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.3). If no `Accept-Encoding` header is send, `Accept-Encoding: identity` is assumed which will result in the API responding uncompressed.

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openproject

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the FileLinkOriginDataModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileLinkOriginDataModel{}

// FileLinkOriginDataModel struct for FileLinkOriginDataModel
type FileLinkOriginDataModel struct {
	// Linked file's id on the origin
	Id string `json:"id"`
	// Linked file's name on the origin
	Name string `json:"name"`
	// MIME type of the linked file.  To link a folder entity, the custom MIME type `application/x-op-directory` MUST be provided. Otherwise it defaults back to an unknown MIME type.
	MimeType *string `json:"mimeType,omitempty"`
	// file size on origin in bytes
	Size *int64 `json:"size,omitempty"`
	// Timestamp of the creation datetime of the file on the origin
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp of the datetime of the last modification of the file on the origin
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
	// Display name of the author that created the file on the origin
	CreatedByName *string `json:"createdByName,omitempty"`
	// Display name of the author that modified the file on the origin last
	LastModifiedByName *string `json:"lastModifiedByName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FileLinkOriginDataModel FileLinkOriginDataModel

// NewFileLinkOriginDataModel instantiates a new FileLinkOriginDataModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLinkOriginDataModel(id string, name string) *FileLinkOriginDataModel {
	this := FileLinkOriginDataModel{}
	this.Id = id
	this.Name = name
	return &this
}

// NewFileLinkOriginDataModelWithDefaults instantiates a new FileLinkOriginDataModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLinkOriginDataModelWithDefaults() *FileLinkOriginDataModel {
	this := FileLinkOriginDataModel{}
	return &this
}

// GetId returns the Id field value
func (o *FileLinkOriginDataModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileLinkOriginDataModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileLinkOriginDataModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FileLinkOriginDataModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileLinkOriginDataModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileLinkOriginDataModel) SetName(v string) {
	o.Name = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *FileLinkOriginDataModel) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkOriginDataModel) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *FileLinkOriginDataModel) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *FileLinkOriginDataModel) SetMimeType(v string) {
	o.MimeType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FileLinkOriginDataModel) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkOriginDataModel) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FileLinkOriginDataModel) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *FileLinkOriginDataModel) SetSize(v int64) {
	o.Size = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FileLinkOriginDataModel) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkOriginDataModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FileLinkOriginDataModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FileLinkOriginDataModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastModifiedAt returns the LastModifiedAt field value if set, zero value otherwise.
func (o *FileLinkOriginDataModel) GetLastModifiedAt() time.Time {
	if o == nil || IsNil(o.LastModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedAt
}

// GetLastModifiedAtOk returns a tuple with the LastModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkOriginDataModel) GetLastModifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedAt) {
		return nil, false
	}
	return o.LastModifiedAt, true
}

// HasLastModifiedAt returns a boolean if a field has been set.
func (o *FileLinkOriginDataModel) HasLastModifiedAt() bool {
	if o != nil && !IsNil(o.LastModifiedAt) {
		return true
	}

	return false
}

// SetLastModifiedAt gets a reference to the given time.Time and assigns it to the LastModifiedAt field.
func (o *FileLinkOriginDataModel) SetLastModifiedAt(v time.Time) {
	o.LastModifiedAt = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *FileLinkOriginDataModel) GetCreatedByName() string {
	if o == nil || IsNil(o.CreatedByName) {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkOriginDataModel) GetCreatedByNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByName) {
		return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *FileLinkOriginDataModel) HasCreatedByName() bool {
	if o != nil && !IsNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *FileLinkOriginDataModel) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetLastModifiedByName returns the LastModifiedByName field value if set, zero value otherwise.
func (o *FileLinkOriginDataModel) GetLastModifiedByName() string {
	if o == nil || IsNil(o.LastModifiedByName) {
		var ret string
		return ret
	}
	return *o.LastModifiedByName
}

// GetLastModifiedByNameOk returns a tuple with the LastModifiedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkOriginDataModel) GetLastModifiedByNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByName) {
		return nil, false
	}
	return o.LastModifiedByName, true
}

// HasLastModifiedByName returns a boolean if a field has been set.
func (o *FileLinkOriginDataModel) HasLastModifiedByName() bool {
	if o != nil && !IsNil(o.LastModifiedByName) {
		return true
	}

	return false
}

// SetLastModifiedByName gets a reference to the given string and assigns it to the LastModifiedByName field.
func (o *FileLinkOriginDataModel) SetLastModifiedByName(v string) {
	o.LastModifiedByName = &v
}

func (o FileLinkOriginDataModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileLinkOriginDataModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.MimeType) {
		toSerialize["mimeType"] = o.MimeType
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.LastModifiedAt) {
		toSerialize["lastModifiedAt"] = o.LastModifiedAt
	}
	if !IsNil(o.CreatedByName) {
		toSerialize["createdByName"] = o.CreatedByName
	}
	if !IsNil(o.LastModifiedByName) {
		toSerialize["lastModifiedByName"] = o.LastModifiedByName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileLinkOriginDataModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varFileLinkOriginDataModel := _FileLinkOriginDataModel{}

	err = json.Unmarshal(data, &varFileLinkOriginDataModel)

	if err != nil {
		return err
	}

	*o = FileLinkOriginDataModel(varFileLinkOriginDataModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "mimeType")
		delete(additionalProperties, "size")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "lastModifiedAt")
		delete(additionalProperties, "createdByName")
		delete(additionalProperties, "lastModifiedByName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileLinkOriginDataModel struct {
	value *FileLinkOriginDataModel
	isSet bool
}

func (v NullableFileLinkOriginDataModel) Get() *FileLinkOriginDataModel {
	return v.value
}

func (v *NullableFileLinkOriginDataModel) Set(val *FileLinkOriginDataModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLinkOriginDataModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLinkOriginDataModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLinkOriginDataModel(val *FileLinkOriginDataModel) *NullableFileLinkOriginDataModel {
	return &NullableFileLinkOriginDataModel{value: val, isSet: true}
}

func (v NullableFileLinkOriginDataModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLinkOriginDataModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

