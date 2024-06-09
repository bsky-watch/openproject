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
)

// checks if the TypesByProjectModelAllOfEmbeddedElements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypesByProjectModelAllOfEmbeddedElements{}

// TypesByProjectModelAllOfEmbeddedElements struct for TypesByProjectModelAllOfEmbeddedElements
type TypesByProjectModelAllOfEmbeddedElements struct {
	// Type id
	Id *int `json:"id,omitempty"`
	// Type name
	Name *string `json:"name,omitempty"`
	// The color used to represent this type
	Color *string `json:"color,omitempty"`
	// Sort index of the type
	Position *int `json:"position,omitempty"`
	// Is this type active by default in new projects?
	IsDefault *bool `json:"isDefault,omitempty"`
	// Do work packages of this type represent a milestone?
	IsMilestone *bool `json:"isMilestone,omitempty"`
	// Time of creation
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Time of the most recent change to the user
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Links *TypeModelLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TypesByProjectModelAllOfEmbeddedElements TypesByProjectModelAllOfEmbeddedElements

// NewTypesByProjectModelAllOfEmbeddedElements instantiates a new TypesByProjectModelAllOfEmbeddedElements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesByProjectModelAllOfEmbeddedElements() *TypesByProjectModelAllOfEmbeddedElements {
	this := TypesByProjectModelAllOfEmbeddedElements{}
	return &this
}

// NewTypesByProjectModelAllOfEmbeddedElementsWithDefaults instantiates a new TypesByProjectModelAllOfEmbeddedElements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesByProjectModelAllOfEmbeddedElementsWithDefaults() *TypesByProjectModelAllOfEmbeddedElements {
	this := TypesByProjectModelAllOfEmbeddedElements{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetId() int {
	if o == nil || IsNil(o.Id) {
		var ret int
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetIdOk() (*int, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int and assigns it to the Id field.
func (o *TypesByProjectModelAllOfEmbeddedElements) SetId(v int) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TypesByProjectModelAllOfEmbeddedElements) SetName(v string) {
	o.Name = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *TypesByProjectModelAllOfEmbeddedElements) SetColor(v string) {
	o.Color = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetPosition() int {
	if o == nil || IsNil(o.Position) {
		var ret int
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetPositionOk() (*int, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int and assigns it to the Position field.
func (o *TypesByProjectModelAllOfEmbeddedElements) SetPosition(v int) {
	o.Position = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *TypesByProjectModelAllOfEmbeddedElements) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsMilestone returns the IsMilestone field value if set, zero value otherwise.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetIsMilestone() bool {
	if o == nil || IsNil(o.IsMilestone) {
		var ret bool
		return ret
	}
	return *o.IsMilestone
}

// GetIsMilestoneOk returns a tuple with the IsMilestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetIsMilestoneOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMilestone) {
		return nil, false
	}
	return o.IsMilestone, true
}

// HasIsMilestone returns a boolean if a field has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) HasIsMilestone() bool {
	if o != nil && !IsNil(o.IsMilestone) {
		return true
	}

	return false
}

// SetIsMilestone gets a reference to the given bool and assigns it to the IsMilestone field.
func (o *TypesByProjectModelAllOfEmbeddedElements) SetIsMilestone(v bool) {
	o.IsMilestone = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TypesByProjectModelAllOfEmbeddedElements) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TypesByProjectModelAllOfEmbeddedElements) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetLinks() TypeModelLinks {
	if o == nil || IsNil(o.Links) {
		var ret TypeModelLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) GetLinksOk() (*TypeModelLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TypesByProjectModelAllOfEmbeddedElements) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TypeModelLinks and assigns it to the Links field.
func (o *TypesByProjectModelAllOfEmbeddedElements) SetLinks(v TypeModelLinks) {
	o.Links = &v
}

func (o TypesByProjectModelAllOfEmbeddedElements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypesByProjectModelAllOfEmbeddedElements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.IsMilestone) {
		toSerialize["isMilestone"] = o.IsMilestone
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TypesByProjectModelAllOfEmbeddedElements) UnmarshalJSON(data []byte) (err error) {
	varTypesByProjectModelAllOfEmbeddedElements := _TypesByProjectModelAllOfEmbeddedElements{}

	err = json.Unmarshal(data, &varTypesByProjectModelAllOfEmbeddedElements)

	if err != nil {
		return err
	}

	*o = TypesByProjectModelAllOfEmbeddedElements(varTypesByProjectModelAllOfEmbeddedElements)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "color")
		delete(additionalProperties, "position")
		delete(additionalProperties, "isDefault")
		delete(additionalProperties, "isMilestone")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTypesByProjectModelAllOfEmbeddedElements struct {
	value *TypesByProjectModelAllOfEmbeddedElements
	isSet bool
}

func (v NullableTypesByProjectModelAllOfEmbeddedElements) Get() *TypesByProjectModelAllOfEmbeddedElements {
	return v.value
}

func (v *NullableTypesByProjectModelAllOfEmbeddedElements) Set(val *TypesByProjectModelAllOfEmbeddedElements) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesByProjectModelAllOfEmbeddedElements) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesByProjectModelAllOfEmbeddedElements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesByProjectModelAllOfEmbeddedElements(val *TypesByProjectModelAllOfEmbeddedElements) *NullableTypesByProjectModelAllOfEmbeddedElements {
	return &NullableTypesByProjectModelAllOfEmbeddedElements{value: val, isSet: true}
}

func (v NullableTypesByProjectModelAllOfEmbeddedElements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesByProjectModelAllOfEmbeddedElements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


