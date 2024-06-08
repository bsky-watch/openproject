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

// checks if the PaginatedCollectionModelAllOfLinksChangeSize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedCollectionModelAllOfLinksChangeSize{}

// PaginatedCollectionModelAllOfLinksChangeSize struct for PaginatedCollectionModelAllOfLinksChangeSize
type PaginatedCollectionModelAllOfLinksChangeSize struct {
	// URL to the referenced resource (might be relative)
	Href string `json:"href"`
	// Representative label for the resource
	Title *string `json:"title,omitempty"`
	// If true the href contains parts that need to be replaced by the client
	Templated *bool `json:"templated,omitempty"`
	// The HTTP verb to use when requesting the resource
	Method *string `json:"method,omitempty"`
	// The payload to send in the request to achieve the desired result
	Payload map[string]interface{} `json:"payload,omitempty"`
	// An optional unique identifier to the link object
	Identifier *string `json:"identifier,omitempty"`
	// The MIME-Type of the returned resource.
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedCollectionModelAllOfLinksChangeSize PaginatedCollectionModelAllOfLinksChangeSize

// NewPaginatedCollectionModelAllOfLinksChangeSize instantiates a new PaginatedCollectionModelAllOfLinksChangeSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCollectionModelAllOfLinksChangeSize(href string) *PaginatedCollectionModelAllOfLinksChangeSize {
	this := PaginatedCollectionModelAllOfLinksChangeSize{}
	this.Href = href
	var templated bool = false
	this.Templated = &templated
	var method string = "GET"
	this.Method = &method
	return &this
}

// NewPaginatedCollectionModelAllOfLinksChangeSizeWithDefaults instantiates a new PaginatedCollectionModelAllOfLinksChangeSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCollectionModelAllOfLinksChangeSizeWithDefaults() *PaginatedCollectionModelAllOfLinksChangeSize {
	this := PaginatedCollectionModelAllOfLinksChangeSize{}
	var templated bool = false
	this.Templated = &templated
	var method string = "GET"
	this.Method = &method
	return &this
}

// GetHref returns the Href field value
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *PaginatedCollectionModelAllOfLinksChangeSize) SetHref(v string) {
	o.Href = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) SetTitle(v string) {
	o.Title = &v
}

// GetTemplated returns the Templated field value if set, zero value otherwise.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetTemplated() bool {
	if o == nil || IsNil(o.Templated) {
		var ret bool
		return ret
	}
	return *o.Templated
}

// GetTemplatedOk returns a tuple with the Templated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetTemplatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Templated) {
		return nil, false
	}
	return o.Templated, true
}

// HasTemplated returns a boolean if a field has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) HasTemplated() bool {
	if o != nil && !IsNil(o.Templated) {
		return true
	}

	return false
}

// SetTemplated gets a reference to the given bool and assigns it to the Templated field.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) SetTemplated(v bool) {
	o.Templated = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) SetMethod(v string) {
	o.Method = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetPayload() map[string]interface{} {
	if o == nil || IsNil(o.Payload) {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Payload) {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaginatedCollectionModelAllOfLinksChangeSize) SetType(v string) {
	o.Type = &v
}

func (o PaginatedCollectionModelAllOfLinksChangeSize) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedCollectionModelAllOfLinksChangeSize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Templated) {
		toSerialize["templated"] = o.Templated
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedCollectionModelAllOfLinksChangeSize) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
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

	varPaginatedCollectionModelAllOfLinksChangeSize := _PaginatedCollectionModelAllOfLinksChangeSize{}

	err = json.Unmarshal(data, &varPaginatedCollectionModelAllOfLinksChangeSize)

	if err != nil {
		return err
	}

	*o = PaginatedCollectionModelAllOfLinksChangeSize(varPaginatedCollectionModelAllOfLinksChangeSize)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "title")
		delete(additionalProperties, "templated")
		delete(additionalProperties, "method")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedCollectionModelAllOfLinksChangeSize struct {
	value *PaginatedCollectionModelAllOfLinksChangeSize
	isSet bool
}

func (v NullablePaginatedCollectionModelAllOfLinksChangeSize) Get() *PaginatedCollectionModelAllOfLinksChangeSize {
	return v.value
}

func (v *NullablePaginatedCollectionModelAllOfLinksChangeSize) Set(val *PaginatedCollectionModelAllOfLinksChangeSize) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCollectionModelAllOfLinksChangeSize) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCollectionModelAllOfLinksChangeSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCollectionModelAllOfLinksChangeSize(val *PaginatedCollectionModelAllOfLinksChangeSize) *NullablePaginatedCollectionModelAllOfLinksChangeSize {
	return &NullablePaginatedCollectionModelAllOfLinksChangeSize{value: val, isSet: true}
}

func (v NullablePaginatedCollectionModelAllOfLinksChangeSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCollectionModelAllOfLinksChangeSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


