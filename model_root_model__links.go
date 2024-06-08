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

// checks if the RootModelLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RootModelLinks{}

// RootModelLinks struct for RootModelLinks
type RootModelLinks struct {
	Self RootModelLinksSelf `json:"self"`
	Configuration RootModelLinksConfiguration `json:"configuration"`
	Memberships RootModelLinksMemberships `json:"memberships"`
	Priorities RootModelLinksPriorities `json:"priorities"`
	Relations RootModelLinksRelations `json:"relations"`
	Statuses RootModelLinksStatuses `json:"statuses"`
	TimeEntries RootModelLinksTimeEntries `json:"time_entries"`
	Types RootModelLinksTypes `json:"types"`
	User RootModelLinksUser `json:"user"`
	UserPreferences RootModelLinksUserPreferences `json:"userPreferences"`
	WorkPackages RootModelLinksWorkPackages `json:"workPackages"`
	AdditionalProperties map[string]interface{}
}

type _RootModelLinks RootModelLinks

// NewRootModelLinks instantiates a new RootModelLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootModelLinks(self RootModelLinksSelf, configuration RootModelLinksConfiguration, memberships RootModelLinksMemberships, priorities RootModelLinksPriorities, relations RootModelLinksRelations, statuses RootModelLinksStatuses, timeEntries RootModelLinksTimeEntries, types RootModelLinksTypes, user RootModelLinksUser, userPreferences RootModelLinksUserPreferences, workPackages RootModelLinksWorkPackages) *RootModelLinks {
	this := RootModelLinks{}
	this.Self = self
	this.Configuration = configuration
	this.Memberships = memberships
	this.Priorities = priorities
	this.Relations = relations
	this.Statuses = statuses
	this.TimeEntries = timeEntries
	this.Types = types
	this.User = user
	this.UserPreferences = userPreferences
	this.WorkPackages = workPackages
	return &this
}

// NewRootModelLinksWithDefaults instantiates a new RootModelLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootModelLinksWithDefaults() *RootModelLinks {
	this := RootModelLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *RootModelLinks) GetSelf() RootModelLinksSelf {
	if o == nil {
		var ret RootModelLinksSelf
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetSelfOk() (*RootModelLinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RootModelLinks) SetSelf(v RootModelLinksSelf) {
	o.Self = v
}

// GetConfiguration returns the Configuration field value
func (o *RootModelLinks) GetConfiguration() RootModelLinksConfiguration {
	if o == nil {
		var ret RootModelLinksConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetConfigurationOk() (*RootModelLinksConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *RootModelLinks) SetConfiguration(v RootModelLinksConfiguration) {
	o.Configuration = v
}

// GetMemberships returns the Memberships field value
func (o *RootModelLinks) GetMemberships() RootModelLinksMemberships {
	if o == nil {
		var ret RootModelLinksMemberships
		return ret
	}

	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetMembershipsOk() (*RootModelLinksMemberships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memberships, true
}

// SetMemberships sets field value
func (o *RootModelLinks) SetMemberships(v RootModelLinksMemberships) {
	o.Memberships = v
}

// GetPriorities returns the Priorities field value
func (o *RootModelLinks) GetPriorities() RootModelLinksPriorities {
	if o == nil {
		var ret RootModelLinksPriorities
		return ret
	}

	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetPrioritiesOk() (*RootModelLinksPriorities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priorities, true
}

// SetPriorities sets field value
func (o *RootModelLinks) SetPriorities(v RootModelLinksPriorities) {
	o.Priorities = v
}

// GetRelations returns the Relations field value
func (o *RootModelLinks) GetRelations() RootModelLinksRelations {
	if o == nil {
		var ret RootModelLinksRelations
		return ret
	}

	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetRelationsOk() (*RootModelLinksRelations, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relations, true
}

// SetRelations sets field value
func (o *RootModelLinks) SetRelations(v RootModelLinksRelations) {
	o.Relations = v
}

// GetStatuses returns the Statuses field value
func (o *RootModelLinks) GetStatuses() RootModelLinksStatuses {
	if o == nil {
		var ret RootModelLinksStatuses
		return ret
	}

	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetStatusesOk() (*RootModelLinksStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statuses, true
}

// SetStatuses sets field value
func (o *RootModelLinks) SetStatuses(v RootModelLinksStatuses) {
	o.Statuses = v
}

// GetTimeEntries returns the TimeEntries field value
func (o *RootModelLinks) GetTimeEntries() RootModelLinksTimeEntries {
	if o == nil {
		var ret RootModelLinksTimeEntries
		return ret
	}

	return o.TimeEntries
}

// GetTimeEntriesOk returns a tuple with the TimeEntries field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetTimeEntriesOk() (*RootModelLinksTimeEntries, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeEntries, true
}

// SetTimeEntries sets field value
func (o *RootModelLinks) SetTimeEntries(v RootModelLinksTimeEntries) {
	o.TimeEntries = v
}

// GetTypes returns the Types field value
func (o *RootModelLinks) GetTypes() RootModelLinksTypes {
	if o == nil {
		var ret RootModelLinksTypes
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetTypesOk() (*RootModelLinksTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Types, true
}

// SetTypes sets field value
func (o *RootModelLinks) SetTypes(v RootModelLinksTypes) {
	o.Types = v
}

// GetUser returns the User field value
func (o *RootModelLinks) GetUser() RootModelLinksUser {
	if o == nil {
		var ret RootModelLinksUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetUserOk() (*RootModelLinksUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *RootModelLinks) SetUser(v RootModelLinksUser) {
	o.User = v
}

// GetUserPreferences returns the UserPreferences field value
func (o *RootModelLinks) GetUserPreferences() RootModelLinksUserPreferences {
	if o == nil {
		var ret RootModelLinksUserPreferences
		return ret
	}

	return o.UserPreferences
}

// GetUserPreferencesOk returns a tuple with the UserPreferences field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetUserPreferencesOk() (*RootModelLinksUserPreferences, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPreferences, true
}

// SetUserPreferences sets field value
func (o *RootModelLinks) SetUserPreferences(v RootModelLinksUserPreferences) {
	o.UserPreferences = v
}

// GetWorkPackages returns the WorkPackages field value
func (o *RootModelLinks) GetWorkPackages() RootModelLinksWorkPackages {
	if o == nil {
		var ret RootModelLinksWorkPackages
		return ret
	}

	return o.WorkPackages
}

// GetWorkPackagesOk returns a tuple with the WorkPackages field value
// and a boolean to check if the value has been set.
func (o *RootModelLinks) GetWorkPackagesOk() (*RootModelLinksWorkPackages, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkPackages, true
}

// SetWorkPackages sets field value
func (o *RootModelLinks) SetWorkPackages(v RootModelLinksWorkPackages) {
	o.WorkPackages = v
}

func (o RootModelLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RootModelLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["configuration"] = o.Configuration
	toSerialize["memberships"] = o.Memberships
	toSerialize["priorities"] = o.Priorities
	toSerialize["relations"] = o.Relations
	toSerialize["statuses"] = o.Statuses
	toSerialize["time_entries"] = o.TimeEntries
	toSerialize["types"] = o.Types
	toSerialize["user"] = o.User
	toSerialize["userPreferences"] = o.UserPreferences
	toSerialize["workPackages"] = o.WorkPackages

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RootModelLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
		"configuration",
		"memberships",
		"priorities",
		"relations",
		"statuses",
		"time_entries",
		"types",
		"user",
		"userPreferences",
		"workPackages",
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

	varRootModelLinks := _RootModelLinks{}

	err = json.Unmarshal(data, &varRootModelLinks)

	if err != nil {
		return err
	}

	*o = RootModelLinks(varRootModelLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "memberships")
		delete(additionalProperties, "priorities")
		delete(additionalProperties, "relations")
		delete(additionalProperties, "statuses")
		delete(additionalProperties, "time_entries")
		delete(additionalProperties, "types")
		delete(additionalProperties, "user")
		delete(additionalProperties, "userPreferences")
		delete(additionalProperties, "workPackages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRootModelLinks struct {
	value *RootModelLinks
	isSet bool
}

func (v NullableRootModelLinks) Get() *RootModelLinks {
	return v.value
}

func (v *NullableRootModelLinks) Set(val *RootModelLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRootModelLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRootModelLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootModelLinks(val *RootModelLinks) *NullableRootModelLinks {
	return &NullableRootModelLinks{value: val, isSet: true}
}

func (v NullableRootModelLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootModelLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


