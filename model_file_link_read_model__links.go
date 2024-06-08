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

// checks if the FileLinkReadModelLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileLinkReadModelLinks{}

// FileLinkReadModelLinks struct for FileLinkReadModelLinks
type FileLinkReadModelLinks struct {
	Self FileLinkReadModelLinksSelf `json:"self"`
	Storage FileLinkReadModelLinksStorage `json:"storage"`
	Container FileLinkReadModelLinksContainer `json:"container"`
	Creator FileLinkReadModelLinksCreator `json:"creator"`
	Delete *FileLinkReadModelLinksDelete `json:"delete,omitempty"`
	Status *FileLinkReadModelLinksStatus `json:"status,omitempty"`
	OriginOpen FileLinkReadModelLinksOriginOpen `json:"originOpen"`
	StaticOriginOpen FileLinkReadModelLinksStaticOriginOpen `json:"staticOriginOpen"`
	OriginOpenLocation FileLinkReadModelLinksOriginOpenLocation `json:"originOpenLocation"`
	StaticOriginOpenLocation FileLinkReadModelLinksStaticOriginOpenLocation `json:"staticOriginOpenLocation"`
	StaticOriginDownload FileLinkReadModelLinksStaticOriginDownload `json:"staticOriginDownload"`
	AdditionalProperties map[string]interface{}
}

type _FileLinkReadModelLinks FileLinkReadModelLinks

// NewFileLinkReadModelLinks instantiates a new FileLinkReadModelLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLinkReadModelLinks(self FileLinkReadModelLinksSelf, storage FileLinkReadModelLinksStorage, container FileLinkReadModelLinksContainer, creator FileLinkReadModelLinksCreator, originOpen FileLinkReadModelLinksOriginOpen, staticOriginOpen FileLinkReadModelLinksStaticOriginOpen, originOpenLocation FileLinkReadModelLinksOriginOpenLocation, staticOriginOpenLocation FileLinkReadModelLinksStaticOriginOpenLocation, staticOriginDownload FileLinkReadModelLinksStaticOriginDownload) *FileLinkReadModelLinks {
	this := FileLinkReadModelLinks{}
	this.Self = self
	this.Storage = storage
	this.Container = container
	this.Creator = creator
	this.OriginOpen = originOpen
	this.StaticOriginOpen = staticOriginOpen
	this.OriginOpenLocation = originOpenLocation
	this.StaticOriginOpenLocation = staticOriginOpenLocation
	this.StaticOriginDownload = staticOriginDownload
	return &this
}

// NewFileLinkReadModelLinksWithDefaults instantiates a new FileLinkReadModelLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLinkReadModelLinksWithDefaults() *FileLinkReadModelLinks {
	this := FileLinkReadModelLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *FileLinkReadModelLinks) GetSelf() FileLinkReadModelLinksSelf {
	if o == nil {
		var ret FileLinkReadModelLinksSelf
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetSelfOk() (*FileLinkReadModelLinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *FileLinkReadModelLinks) SetSelf(v FileLinkReadModelLinksSelf) {
	o.Self = v
}

// GetStorage returns the Storage field value
func (o *FileLinkReadModelLinks) GetStorage() FileLinkReadModelLinksStorage {
	if o == nil {
		var ret FileLinkReadModelLinksStorage
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetStorageOk() (*FileLinkReadModelLinksStorage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *FileLinkReadModelLinks) SetStorage(v FileLinkReadModelLinksStorage) {
	o.Storage = v
}

// GetContainer returns the Container field value
func (o *FileLinkReadModelLinks) GetContainer() FileLinkReadModelLinksContainer {
	if o == nil {
		var ret FileLinkReadModelLinksContainer
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetContainerOk() (*FileLinkReadModelLinksContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *FileLinkReadModelLinks) SetContainer(v FileLinkReadModelLinksContainer) {
	o.Container = v
}

// GetCreator returns the Creator field value
func (o *FileLinkReadModelLinks) GetCreator() FileLinkReadModelLinksCreator {
	if o == nil {
		var ret FileLinkReadModelLinksCreator
		return ret
	}

	return o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetCreatorOk() (*FileLinkReadModelLinksCreator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creator, true
}

// SetCreator sets field value
func (o *FileLinkReadModelLinks) SetCreator(v FileLinkReadModelLinksCreator) {
	o.Creator = v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *FileLinkReadModelLinks) GetDelete() FileLinkReadModelLinksDelete {
	if o == nil || IsNil(o.Delete) {
		var ret FileLinkReadModelLinksDelete
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetDeleteOk() (*FileLinkReadModelLinksDelete, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *FileLinkReadModelLinks) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given FileLinkReadModelLinksDelete and assigns it to the Delete field.
func (o *FileLinkReadModelLinks) SetDelete(v FileLinkReadModelLinksDelete) {
	o.Delete = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FileLinkReadModelLinks) GetStatus() FileLinkReadModelLinksStatus {
	if o == nil || IsNil(o.Status) {
		var ret FileLinkReadModelLinksStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetStatusOk() (*FileLinkReadModelLinksStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FileLinkReadModelLinks) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given FileLinkReadModelLinksStatus and assigns it to the Status field.
func (o *FileLinkReadModelLinks) SetStatus(v FileLinkReadModelLinksStatus) {
	o.Status = &v
}

// GetOriginOpen returns the OriginOpen field value
func (o *FileLinkReadModelLinks) GetOriginOpen() FileLinkReadModelLinksOriginOpen {
	if o == nil {
		var ret FileLinkReadModelLinksOriginOpen
		return ret
	}

	return o.OriginOpen
}

// GetOriginOpenOk returns a tuple with the OriginOpen field value
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetOriginOpenOk() (*FileLinkReadModelLinksOriginOpen, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginOpen, true
}

// SetOriginOpen sets field value
func (o *FileLinkReadModelLinks) SetOriginOpen(v FileLinkReadModelLinksOriginOpen) {
	o.OriginOpen = v
}

// GetStaticOriginOpen returns the StaticOriginOpen field value
func (o *FileLinkReadModelLinks) GetStaticOriginOpen() FileLinkReadModelLinksStaticOriginOpen {
	if o == nil {
		var ret FileLinkReadModelLinksStaticOriginOpen
		return ret
	}

	return o.StaticOriginOpen
}

// GetStaticOriginOpenOk returns a tuple with the StaticOriginOpen field value
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetStaticOriginOpenOk() (*FileLinkReadModelLinksStaticOriginOpen, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StaticOriginOpen, true
}

// SetStaticOriginOpen sets field value
func (o *FileLinkReadModelLinks) SetStaticOriginOpen(v FileLinkReadModelLinksStaticOriginOpen) {
	o.StaticOriginOpen = v
}

// GetOriginOpenLocation returns the OriginOpenLocation field value
func (o *FileLinkReadModelLinks) GetOriginOpenLocation() FileLinkReadModelLinksOriginOpenLocation {
	if o == nil {
		var ret FileLinkReadModelLinksOriginOpenLocation
		return ret
	}

	return o.OriginOpenLocation
}

// GetOriginOpenLocationOk returns a tuple with the OriginOpenLocation field value
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetOriginOpenLocationOk() (*FileLinkReadModelLinksOriginOpenLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginOpenLocation, true
}

// SetOriginOpenLocation sets field value
func (o *FileLinkReadModelLinks) SetOriginOpenLocation(v FileLinkReadModelLinksOriginOpenLocation) {
	o.OriginOpenLocation = v
}

// GetStaticOriginOpenLocation returns the StaticOriginOpenLocation field value
func (o *FileLinkReadModelLinks) GetStaticOriginOpenLocation() FileLinkReadModelLinksStaticOriginOpenLocation {
	if o == nil {
		var ret FileLinkReadModelLinksStaticOriginOpenLocation
		return ret
	}

	return o.StaticOriginOpenLocation
}

// GetStaticOriginOpenLocationOk returns a tuple with the StaticOriginOpenLocation field value
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetStaticOriginOpenLocationOk() (*FileLinkReadModelLinksStaticOriginOpenLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StaticOriginOpenLocation, true
}

// SetStaticOriginOpenLocation sets field value
func (o *FileLinkReadModelLinks) SetStaticOriginOpenLocation(v FileLinkReadModelLinksStaticOriginOpenLocation) {
	o.StaticOriginOpenLocation = v
}

// GetStaticOriginDownload returns the StaticOriginDownload field value
func (o *FileLinkReadModelLinks) GetStaticOriginDownload() FileLinkReadModelLinksStaticOriginDownload {
	if o == nil {
		var ret FileLinkReadModelLinksStaticOriginDownload
		return ret
	}

	return o.StaticOriginDownload
}

// GetStaticOriginDownloadOk returns a tuple with the StaticOriginDownload field value
// and a boolean to check if the value has been set.
func (o *FileLinkReadModelLinks) GetStaticOriginDownloadOk() (*FileLinkReadModelLinksStaticOriginDownload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StaticOriginDownload, true
}

// SetStaticOriginDownload sets field value
func (o *FileLinkReadModelLinks) SetStaticOriginDownload(v FileLinkReadModelLinksStaticOriginDownload) {
	o.StaticOriginDownload = v
}

func (o FileLinkReadModelLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileLinkReadModelLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["storage"] = o.Storage
	toSerialize["container"] = o.Container
	toSerialize["creator"] = o.Creator
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["originOpen"] = o.OriginOpen
	toSerialize["staticOriginOpen"] = o.StaticOriginOpen
	toSerialize["originOpenLocation"] = o.OriginOpenLocation
	toSerialize["staticOriginOpenLocation"] = o.StaticOriginOpenLocation
	toSerialize["staticOriginDownload"] = o.StaticOriginDownload

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileLinkReadModelLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
		"storage",
		"container",
		"creator",
		"originOpen",
		"staticOriginOpen",
		"originOpenLocation",
		"staticOriginOpenLocation",
		"staticOriginDownload",
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

	varFileLinkReadModelLinks := _FileLinkReadModelLinks{}

	err = json.Unmarshal(data, &varFileLinkReadModelLinks)

	if err != nil {
		return err
	}

	*o = FileLinkReadModelLinks(varFileLinkReadModelLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "storage")
		delete(additionalProperties, "container")
		delete(additionalProperties, "creator")
		delete(additionalProperties, "delete")
		delete(additionalProperties, "status")
		delete(additionalProperties, "originOpen")
		delete(additionalProperties, "staticOriginOpen")
		delete(additionalProperties, "originOpenLocation")
		delete(additionalProperties, "staticOriginOpenLocation")
		delete(additionalProperties, "staticOriginDownload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileLinkReadModelLinks struct {
	value *FileLinkReadModelLinks
	isSet bool
}

func (v NullableFileLinkReadModelLinks) Get() *FileLinkReadModelLinks {
	return v.value
}

func (v *NullableFileLinkReadModelLinks) Set(val *FileLinkReadModelLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLinkReadModelLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLinkReadModelLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLinkReadModelLinks(val *FileLinkReadModelLinks) *NullableFileLinkReadModelLinks {
	return &NullableFileLinkReadModelLinks{value: val, isSet: true}
}

func (v NullableFileLinkReadModelLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLinkReadModelLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


