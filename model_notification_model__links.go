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

// checks if the NotificationModelLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationModelLinks{}

// NotificationModelLinks struct for NotificationModelLinks
type NotificationModelLinks struct {
	Self NotificationModelLinksSelf `json:"self"`
	ReadIAN *NotificationModelLinksReadIAN `json:"readIAN,omitempty"`
	UnreadIAN *NotificationModelLinksUnreadIAN `json:"unreadIAN,omitempty"`
	Project NotificationModelLinksProject `json:"project"`
	Actor NotificationModelLinksActor `json:"actor"`
	Resource NotificationModelLinksResource `json:"resource"`
	Activity NotificationModelLinksActivity `json:"activity"`
	Details []NotificationModelLinksDetailsInner `json:"details"`
	AdditionalProperties map[string]interface{}
}

type _NotificationModelLinks NotificationModelLinks

// NewNotificationModelLinks instantiates a new NotificationModelLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationModelLinks(self NotificationModelLinksSelf, project NotificationModelLinksProject, actor NotificationModelLinksActor, resource NotificationModelLinksResource, activity NotificationModelLinksActivity, details []NotificationModelLinksDetailsInner) *NotificationModelLinks {
	this := NotificationModelLinks{}
	this.Self = self
	this.Project = project
	this.Actor = actor
	this.Resource = resource
	this.Activity = activity
	this.Details = details
	return &this
}

// NewNotificationModelLinksWithDefaults instantiates a new NotificationModelLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationModelLinksWithDefaults() *NotificationModelLinks {
	this := NotificationModelLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *NotificationModelLinks) GetSelf() NotificationModelLinksSelf {
	if o == nil {
		var ret NotificationModelLinksSelf
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *NotificationModelLinks) GetSelfOk() (*NotificationModelLinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *NotificationModelLinks) SetSelf(v NotificationModelLinksSelf) {
	o.Self = v
}

// GetReadIAN returns the ReadIAN field value if set, zero value otherwise.
func (o *NotificationModelLinks) GetReadIAN() NotificationModelLinksReadIAN {
	if o == nil || IsNil(o.ReadIAN) {
		var ret NotificationModelLinksReadIAN
		return ret
	}
	return *o.ReadIAN
}

// GetReadIANOk returns a tuple with the ReadIAN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModelLinks) GetReadIANOk() (*NotificationModelLinksReadIAN, bool) {
	if o == nil || IsNil(o.ReadIAN) {
		return nil, false
	}
	return o.ReadIAN, true
}

// HasReadIAN returns a boolean if a field has been set.
func (o *NotificationModelLinks) HasReadIAN() bool {
	if o != nil && !IsNil(o.ReadIAN) {
		return true
	}

	return false
}

// SetReadIAN gets a reference to the given NotificationModelLinksReadIAN and assigns it to the ReadIAN field.
func (o *NotificationModelLinks) SetReadIAN(v NotificationModelLinksReadIAN) {
	o.ReadIAN = &v
}

// GetUnreadIAN returns the UnreadIAN field value if set, zero value otherwise.
func (o *NotificationModelLinks) GetUnreadIAN() NotificationModelLinksUnreadIAN {
	if o == nil || IsNil(o.UnreadIAN) {
		var ret NotificationModelLinksUnreadIAN
		return ret
	}
	return *o.UnreadIAN
}

// GetUnreadIANOk returns a tuple with the UnreadIAN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModelLinks) GetUnreadIANOk() (*NotificationModelLinksUnreadIAN, bool) {
	if o == nil || IsNil(o.UnreadIAN) {
		return nil, false
	}
	return o.UnreadIAN, true
}

// HasUnreadIAN returns a boolean if a field has been set.
func (o *NotificationModelLinks) HasUnreadIAN() bool {
	if o != nil && !IsNil(o.UnreadIAN) {
		return true
	}

	return false
}

// SetUnreadIAN gets a reference to the given NotificationModelLinksUnreadIAN and assigns it to the UnreadIAN field.
func (o *NotificationModelLinks) SetUnreadIAN(v NotificationModelLinksUnreadIAN) {
	o.UnreadIAN = &v
}

// GetProject returns the Project field value
func (o *NotificationModelLinks) GetProject() NotificationModelLinksProject {
	if o == nil {
		var ret NotificationModelLinksProject
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *NotificationModelLinks) GetProjectOk() (*NotificationModelLinksProject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *NotificationModelLinks) SetProject(v NotificationModelLinksProject) {
	o.Project = v
}

// GetActor returns the Actor field value
func (o *NotificationModelLinks) GetActor() NotificationModelLinksActor {
	if o == nil {
		var ret NotificationModelLinksActor
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *NotificationModelLinks) GetActorOk() (*NotificationModelLinksActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *NotificationModelLinks) SetActor(v NotificationModelLinksActor) {
	o.Actor = v
}

// GetResource returns the Resource field value
func (o *NotificationModelLinks) GetResource() NotificationModelLinksResource {
	if o == nil {
		var ret NotificationModelLinksResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *NotificationModelLinks) GetResourceOk() (*NotificationModelLinksResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *NotificationModelLinks) SetResource(v NotificationModelLinksResource) {
	o.Resource = v
}

// GetActivity returns the Activity field value
func (o *NotificationModelLinks) GetActivity() NotificationModelLinksActivity {
	if o == nil {
		var ret NotificationModelLinksActivity
		return ret
	}

	return o.Activity
}

// GetActivityOk returns a tuple with the Activity field value
// and a boolean to check if the value has been set.
func (o *NotificationModelLinks) GetActivityOk() (*NotificationModelLinksActivity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Activity, true
}

// SetActivity sets field value
func (o *NotificationModelLinks) SetActivity(v NotificationModelLinksActivity) {
	o.Activity = v
}

// GetDetails returns the Details field value
func (o *NotificationModelLinks) GetDetails() []NotificationModelLinksDetailsInner {
	if o == nil {
		var ret []NotificationModelLinksDetailsInner
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *NotificationModelLinks) GetDetailsOk() ([]NotificationModelLinksDetailsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details, true
}

// SetDetails sets field value
func (o *NotificationModelLinks) SetDetails(v []NotificationModelLinksDetailsInner) {
	o.Details = v
}

func (o NotificationModelLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationModelLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.ReadIAN) {
		toSerialize["readIAN"] = o.ReadIAN
	}
	if !IsNil(o.UnreadIAN) {
		toSerialize["unreadIAN"] = o.UnreadIAN
	}
	toSerialize["project"] = o.Project
	toSerialize["actor"] = o.Actor
	toSerialize["resource"] = o.Resource
	toSerialize["activity"] = o.Activity
	toSerialize["details"] = o.Details

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationModelLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
		"project",
		"actor",
		"resource",
		"activity",
		"details",
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

	varNotificationModelLinks := _NotificationModelLinks{}

	err = json.Unmarshal(data, &varNotificationModelLinks)

	if err != nil {
		return err
	}

	*o = NotificationModelLinks(varNotificationModelLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "readIAN")
		delete(additionalProperties, "unreadIAN")
		delete(additionalProperties, "project")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "activity")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationModelLinks struct {
	value *NotificationModelLinks
	isSet bool
}

func (v NullableNotificationModelLinks) Get() *NotificationModelLinks {
	return v.value
}

func (v *NullableNotificationModelLinks) Set(val *NotificationModelLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationModelLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationModelLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationModelLinks(val *NotificationModelLinks) *NullableNotificationModelLinks {
	return &NullableNotificationModelLinks{value: val, isSet: true}
}

func (v NullableNotificationModelLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationModelLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


