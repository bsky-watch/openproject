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

// checks if the AvailableAssigneesModelAllOfEmbeddedElements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableAssigneesModelAllOfEmbeddedElements{}

// AvailableAssigneesModelAllOfEmbeddedElements struct for AvailableAssigneesModelAllOfEmbeddedElements
type AvailableAssigneesModelAllOfEmbeddedElements struct {
	Type string `json:"_type"`
	// The principal's unique identifier.
	Id int `json:"id"`
	// The principal's display name, layout depends on instance settings.
	Name string `json:"name"`
	// Time of creation
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Time of the most recent change to the user
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Links UserModelAllOfLinks `json:"_links"`
	// URL to user's avatar
	Avatar string `json:"avatar"`
	// The user's login name  # Conditions  - User is self, or `create_user` or `manage_user` permission globally
	Login *string `json:"login,omitempty"`
	// The user's first name  # Conditions  - User is self, or `create_user` or `manage_user` permission globally
	FirstName *string `json:"firstName,omitempty"`
	// The user's last name  # Conditions  - User is self, or `create_user` or `manage_user` permission globally
	LastName *string `json:"lastName,omitempty"`
	// The user's email address  # Conditions  - E-Mail address not hidden - User is not a new record - User is self, or `create_user` or `manage_user` permission globally
	Email *string `json:"email,omitempty"`
	// Flag indicating whether or not the user is an admin  # Conditions  - `admin`
	Admin *bool `json:"admin,omitempty"`
	// The current activation status of the user.  # Conditions  - User is self, or `create_user` or `manage_user` permission globally
	Status *string `json:"status,omitempty"`
	// User's language | ISO 639-1 format  # Conditions  - User is self, or `create_user` or `manage_user` permission globally
	Language *string `json:"language,omitempty"`
	// User's identity_url for OmniAuth authentication.  # Conditions  - User is self, or `create_user` or `manage_user` permission globally
	IdentityUrl NullableString `json:"identityUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AvailableAssigneesModelAllOfEmbeddedElements AvailableAssigneesModelAllOfEmbeddedElements

// NewAvailableAssigneesModelAllOfEmbeddedElements instantiates a new AvailableAssigneesModelAllOfEmbeddedElements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableAssigneesModelAllOfEmbeddedElements(type_ string, id int, name string, links UserModelAllOfLinks, avatar string) *AvailableAssigneesModelAllOfEmbeddedElements {
	this := AvailableAssigneesModelAllOfEmbeddedElements{}
	this.Type = type_
	this.Id = id
	this.Name = name
	this.Links = links
	this.Avatar = avatar
	return &this
}

// NewAvailableAssigneesModelAllOfEmbeddedElementsWithDefaults instantiates a new AvailableAssigneesModelAllOfEmbeddedElements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableAssigneesModelAllOfEmbeddedElementsWithDefaults() *AvailableAssigneesModelAllOfEmbeddedElements {
	this := AvailableAssigneesModelAllOfEmbeddedElements{}
	return &this
}

// GetType returns the Type field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetId() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetIdOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetId(v int) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLinks returns the Links field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetLinks() UserModelAllOfLinks {
	if o == nil {
		var ret UserModelAllOfLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetLinksOk() (*UserModelAllOfLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetLinks(v UserModelAllOfLinks) {
	o.Links = v
}

// GetAvatar returns the Avatar field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetAvatar(v string) {
	o.Avatar = v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetLogin(v string) {
	o.Login = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetEmail(v string) {
	o.Email = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetAdmin() bool {
	if o == nil || IsNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetAdmin(v bool) {
	o.Admin = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetStatus(v string) {
	o.Status = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetLanguage(v string) {
	o.Language = &v
}

// GetIdentityUrl returns the IdentityUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetIdentityUrl() string {
	if o == nil || IsNil(o.IdentityUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityUrl.Get()
}

// GetIdentityUrlOk returns a tuple with the IdentityUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailableAssigneesModelAllOfEmbeddedElements) GetIdentityUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityUrl.Get(), o.IdentityUrl.IsSet()
}

// HasIdentityUrl returns a boolean if a field has been set.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) HasIdentityUrl() bool {
	if o != nil && o.IdentityUrl.IsSet() {
		return true
	}

	return false
}

// SetIdentityUrl gets a reference to the given NullableString and assigns it to the IdentityUrl field.
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetIdentityUrl(v string) {
	o.IdentityUrl.Set(&v)
}
// SetIdentityUrlNil sets the value for IdentityUrl to be an explicit nil
func (o *AvailableAssigneesModelAllOfEmbeddedElements) SetIdentityUrlNil() {
	o.IdentityUrl.Set(nil)
}

// UnsetIdentityUrl ensures that no value is present for IdentityUrl, not even an explicit nil
func (o *AvailableAssigneesModelAllOfEmbeddedElements) UnsetIdentityUrl() {
	o.IdentityUrl.Unset()
}

func (o AvailableAssigneesModelAllOfEmbeddedElements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailableAssigneesModelAllOfEmbeddedElements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	toSerialize["_links"] = o.Links
	toSerialize["avatar"] = o.Avatar
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if o.IdentityUrl.IsSet() {
		toSerialize["identityUrl"] = o.IdentityUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvailableAssigneesModelAllOfEmbeddedElements) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_type",
		"id",
		"name",
		"_links",
		"avatar",
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

	varAvailableAssigneesModelAllOfEmbeddedElements := _AvailableAssigneesModelAllOfEmbeddedElements{}

	err = json.Unmarshal(data, &varAvailableAssigneesModelAllOfEmbeddedElements)

	if err != nil {
		return err
	}

	*o = AvailableAssigneesModelAllOfEmbeddedElements(varAvailableAssigneesModelAllOfEmbeddedElements)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "avatar")
		delete(additionalProperties, "login")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "admin")
		delete(additionalProperties, "status")
		delete(additionalProperties, "language")
		delete(additionalProperties, "identityUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvailableAssigneesModelAllOfEmbeddedElements struct {
	value *AvailableAssigneesModelAllOfEmbeddedElements
	isSet bool
}

func (v NullableAvailableAssigneesModelAllOfEmbeddedElements) Get() *AvailableAssigneesModelAllOfEmbeddedElements {
	return v.value
}

func (v *NullableAvailableAssigneesModelAllOfEmbeddedElements) Set(val *AvailableAssigneesModelAllOfEmbeddedElements) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableAssigneesModelAllOfEmbeddedElements) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableAssigneesModelAllOfEmbeddedElements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableAssigneesModelAllOfEmbeddedElements(val *AvailableAssigneesModelAllOfEmbeddedElements) *NullableAvailableAssigneesModelAllOfEmbeddedElements {
	return &NullableAvailableAssigneesModelAllOfEmbeddedElements{value: val, isSet: true}
}

func (v NullableAvailableAssigneesModelAllOfEmbeddedElements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableAssigneesModelAllOfEmbeddedElements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


