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

// checks if the MembershipWriteModelMetaNotificationMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipWriteModelMetaNotificationMessage{}

// MembershipWriteModelMetaNotificationMessage struct for MembershipWriteModelMetaNotificationMessage
type MembershipWriteModelMetaNotificationMessage struct {
	// Indicates the formatting language of the raw text
	Format string `json:"format"`
	// The raw text, as entered by the user
	Raw *string `json:"raw,omitempty"`
	// The text converted to HTML according to the format
	Html *string `json:"html,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MembershipWriteModelMetaNotificationMessage MembershipWriteModelMetaNotificationMessage

// NewMembershipWriteModelMetaNotificationMessage instantiates a new MembershipWriteModelMetaNotificationMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipWriteModelMetaNotificationMessage(format string) *MembershipWriteModelMetaNotificationMessage {
	this := MembershipWriteModelMetaNotificationMessage{}
	this.Format = format
	return &this
}

// NewMembershipWriteModelMetaNotificationMessageWithDefaults instantiates a new MembershipWriteModelMetaNotificationMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipWriteModelMetaNotificationMessageWithDefaults() *MembershipWriteModelMetaNotificationMessage {
	this := MembershipWriteModelMetaNotificationMessage{}
	return &this
}

// GetFormat returns the Format field value
func (o *MembershipWriteModelMetaNotificationMessage) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *MembershipWriteModelMetaNotificationMessage) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *MembershipWriteModelMetaNotificationMessage) SetFormat(v string) {
	o.Format = v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *MembershipWriteModelMetaNotificationMessage) GetRaw() string {
	if o == nil || IsNil(o.Raw) {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipWriteModelMetaNotificationMessage) GetRawOk() (*string, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *MembershipWriteModelMetaNotificationMessage) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *MembershipWriteModelMetaNotificationMessage) SetRaw(v string) {
	o.Raw = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *MembershipWriteModelMetaNotificationMessage) GetHtml() string {
	if o == nil || IsNil(o.Html) {
		var ret string
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipWriteModelMetaNotificationMessage) GetHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.Html) {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *MembershipWriteModelMetaNotificationMessage) HasHtml() bool {
	if o != nil && !IsNil(o.Html) {
		return true
	}

	return false
}

// SetHtml gets a reference to the given string and assigns it to the Html field.
func (o *MembershipWriteModelMetaNotificationMessage) SetHtml(v string) {
	o.Html = &v
}

func (o MembershipWriteModelMetaNotificationMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipWriteModelMetaNotificationMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["format"] = o.Format
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	if !IsNil(o.Html) {
		toSerialize["html"] = o.Html
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MembershipWriteModelMetaNotificationMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"format",
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

	varMembershipWriteModelMetaNotificationMessage := _MembershipWriteModelMetaNotificationMessage{}

	err = json.Unmarshal(data, &varMembershipWriteModelMetaNotificationMessage)

	if err != nil {
		return err
	}

	*o = MembershipWriteModelMetaNotificationMessage(varMembershipWriteModelMetaNotificationMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "raw")
		delete(additionalProperties, "html")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMembershipWriteModelMetaNotificationMessage struct {
	value *MembershipWriteModelMetaNotificationMessage
	isSet bool
}

func (v NullableMembershipWriteModelMetaNotificationMessage) Get() *MembershipWriteModelMetaNotificationMessage {
	return v.value
}

func (v *NullableMembershipWriteModelMetaNotificationMessage) Set(val *MembershipWriteModelMetaNotificationMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipWriteModelMetaNotificationMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipWriteModelMetaNotificationMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipWriteModelMetaNotificationMessage(val *MembershipWriteModelMetaNotificationMessage) *NullableMembershipWriteModelMetaNotificationMessage {
	return &NullableMembershipWriteModelMetaNotificationMessage{value: val, isSet: true}
}

func (v NullableMembershipWriteModelMetaNotificationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipWriteModelMetaNotificationMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


