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

// checks if the RevisionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevisionModel{}

// RevisionModel struct for RevisionModel
type RevisionModel struct {
	// Revision's id, assigned by OpenProject
	Id *int `json:"id,omitempty"`
	// The raw SCM identifier of the revision (e.g. full SHA hash)
	Identifier string `json:"identifier"`
	// The SCM identifier of the revision, formatted (e.g. shortened unambiguous SHA hash). May be identical to identifier in many cases
	FormattedIdentifier string `json:"formattedIdentifier"`
	// The name of the author that committed this revision. Note that this name is retrieved from the repository and does not identify a user in OpenProject.
	AuthorName string `json:"authorName"`
	Message RevisionModelMessage `json:"message"`
	// The time this revision was committed to the repository
	CreatedAt time.Time `json:"createdAt"`
	Links *RevisionModelLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RevisionModel RevisionModel

// NewRevisionModel instantiates a new RevisionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionModel(identifier string, formattedIdentifier string, authorName string, message RevisionModelMessage, createdAt time.Time) *RevisionModel {
	this := RevisionModel{}
	this.Identifier = identifier
	this.FormattedIdentifier = formattedIdentifier
	this.AuthorName = authorName
	this.Message = message
	this.CreatedAt = createdAt
	return &this
}

// NewRevisionModelWithDefaults instantiates a new RevisionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionModelWithDefaults() *RevisionModel {
	this := RevisionModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RevisionModel) GetId() int {
	if o == nil || IsNil(o.Id) {
		var ret int
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionModel) GetIdOk() (*int, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RevisionModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int and assigns it to the Id field.
func (o *RevisionModel) SetId(v int) {
	o.Id = &v
}

// GetIdentifier returns the Identifier field value
func (o *RevisionModel) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *RevisionModel) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *RevisionModel) SetIdentifier(v string) {
	o.Identifier = v
}

// GetFormattedIdentifier returns the FormattedIdentifier field value
func (o *RevisionModel) GetFormattedIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormattedIdentifier
}

// GetFormattedIdentifierOk returns a tuple with the FormattedIdentifier field value
// and a boolean to check if the value has been set.
func (o *RevisionModel) GetFormattedIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormattedIdentifier, true
}

// SetFormattedIdentifier sets field value
func (o *RevisionModel) SetFormattedIdentifier(v string) {
	o.FormattedIdentifier = v
}

// GetAuthorName returns the AuthorName field value
func (o *RevisionModel) GetAuthorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value
// and a boolean to check if the value has been set.
func (o *RevisionModel) GetAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorName, true
}

// SetAuthorName sets field value
func (o *RevisionModel) SetAuthorName(v string) {
	o.AuthorName = v
}

// GetMessage returns the Message field value
func (o *RevisionModel) GetMessage() RevisionModelMessage {
	if o == nil {
		var ret RevisionModelMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RevisionModel) GetMessageOk() (*RevisionModelMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RevisionModel) SetMessage(v RevisionModelMessage) {
	o.Message = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RevisionModel) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RevisionModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RevisionModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RevisionModel) GetLinks() RevisionModelLinks {
	if o == nil || IsNil(o.Links) {
		var ret RevisionModelLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionModel) GetLinksOk() (*RevisionModelLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RevisionModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RevisionModelLinks and assigns it to the Links field.
func (o *RevisionModel) SetLinks(v RevisionModelLinks) {
	o.Links = &v
}

func (o RevisionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevisionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["identifier"] = o.Identifier
	toSerialize["formattedIdentifier"] = o.FormattedIdentifier
	toSerialize["authorName"] = o.AuthorName
	toSerialize["message"] = o.Message
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RevisionModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifier",
		"formattedIdentifier",
		"authorName",
		"message",
		"createdAt",
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

	varRevisionModel := _RevisionModel{}

	err = json.Unmarshal(data, &varRevisionModel)

	if err != nil {
		return err
	}

	*o = RevisionModel(varRevisionModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "formattedIdentifier")
		delete(additionalProperties, "authorName")
		delete(additionalProperties, "message")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevisionModel struct {
	value *RevisionModel
	isSet bool
}

func (v NullableRevisionModel) Get() *RevisionModel {
	return v.value
}

func (v *NullableRevisionModel) Set(val *RevisionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionModel(val *RevisionModel) *NullableRevisionModel {
	return &NullableRevisionModel{value: val, isSet: true}
}

func (v NullableRevisionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


