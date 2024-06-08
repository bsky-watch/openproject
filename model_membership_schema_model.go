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

// checks if the MembershipSchemaModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipSchemaModel{}

// MembershipSchemaModel struct for MembershipSchemaModel
type MembershipSchemaModel struct {
	Type string `json:"_type"`
	// A list of dependencies between one property's value and another property
	Dependencies []string `json:"_dependencies,omitempty"`
	Links SchemaModelLinks `json:"_links"`
	Id SchemaPropertyModel `json:"id"`
	CreatedAt SchemaPropertyModel `json:"createdAt"`
	UpdatedAt SchemaPropertyModel `json:"updatedAt"`
	NotificationMessage SchemaPropertyModel `json:"notificationMessage"`
	Project SchemaPropertyModel `json:"project"`
	Principal SchemaPropertyModel `json:"principal"`
	Roles SchemaPropertyModel `json:"roles"`
	AdditionalProperties map[string]interface{}
}

type _MembershipSchemaModel MembershipSchemaModel

// NewMembershipSchemaModel instantiates a new MembershipSchemaModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipSchemaModel(type_ string, links SchemaModelLinks, id SchemaPropertyModel, createdAt SchemaPropertyModel, updatedAt SchemaPropertyModel, notificationMessage SchemaPropertyModel, project SchemaPropertyModel, principal SchemaPropertyModel, roles SchemaPropertyModel) *MembershipSchemaModel {
	this := MembershipSchemaModel{}
	this.Type = type_
	this.Links = links
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.NotificationMessage = notificationMessage
	this.Project = project
	this.Principal = principal
	this.Roles = roles
	return &this
}

// NewMembershipSchemaModelWithDefaults instantiates a new MembershipSchemaModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipSchemaModelWithDefaults() *MembershipSchemaModel {
	this := MembershipSchemaModel{}
	return &this
}

// GetType returns the Type field value
func (o *MembershipSchemaModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MembershipSchemaModel) SetType(v string) {
	o.Type = v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *MembershipSchemaModel) GetDependencies() []string {
	if o == nil || IsNil(o.Dependencies) {
		var ret []string
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetDependenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *MembershipSchemaModel) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []string and assigns it to the Dependencies field.
func (o *MembershipSchemaModel) SetDependencies(v []string) {
	o.Dependencies = v
}

// GetLinks returns the Links field value
func (o *MembershipSchemaModel) GetLinks() SchemaModelLinks {
	if o == nil {
		var ret SchemaModelLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetLinksOk() (*SchemaModelLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *MembershipSchemaModel) SetLinks(v SchemaModelLinks) {
	o.Links = v
}

// GetId returns the Id field value
func (o *MembershipSchemaModel) GetId() SchemaPropertyModel {
	if o == nil {
		var ret SchemaPropertyModel
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetIdOk() (*SchemaPropertyModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MembershipSchemaModel) SetId(v SchemaPropertyModel) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MembershipSchemaModel) GetCreatedAt() SchemaPropertyModel {
	if o == nil {
		var ret SchemaPropertyModel
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetCreatedAtOk() (*SchemaPropertyModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MembershipSchemaModel) SetCreatedAt(v SchemaPropertyModel) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *MembershipSchemaModel) GetUpdatedAt() SchemaPropertyModel {
	if o == nil {
		var ret SchemaPropertyModel
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetUpdatedAtOk() (*SchemaPropertyModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *MembershipSchemaModel) SetUpdatedAt(v SchemaPropertyModel) {
	o.UpdatedAt = v
}

// GetNotificationMessage returns the NotificationMessage field value
func (o *MembershipSchemaModel) GetNotificationMessage() SchemaPropertyModel {
	if o == nil {
		var ret SchemaPropertyModel
		return ret
	}

	return o.NotificationMessage
}

// GetNotificationMessageOk returns a tuple with the NotificationMessage field value
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetNotificationMessageOk() (*SchemaPropertyModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationMessage, true
}

// SetNotificationMessage sets field value
func (o *MembershipSchemaModel) SetNotificationMessage(v SchemaPropertyModel) {
	o.NotificationMessage = v
}

// GetProject returns the Project field value
func (o *MembershipSchemaModel) GetProject() SchemaPropertyModel {
	if o == nil {
		var ret SchemaPropertyModel
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetProjectOk() (*SchemaPropertyModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *MembershipSchemaModel) SetProject(v SchemaPropertyModel) {
	o.Project = v
}

// GetPrincipal returns the Principal field value
func (o *MembershipSchemaModel) GetPrincipal() SchemaPropertyModel {
	if o == nil {
		var ret SchemaPropertyModel
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetPrincipalOk() (*SchemaPropertyModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *MembershipSchemaModel) SetPrincipal(v SchemaPropertyModel) {
	o.Principal = v
}

// GetRoles returns the Roles field value
func (o *MembershipSchemaModel) GetRoles() SchemaPropertyModel {
	if o == nil {
		var ret SchemaPropertyModel
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *MembershipSchemaModel) GetRolesOk() (*SchemaPropertyModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *MembershipSchemaModel) SetRoles(v SchemaPropertyModel) {
	o.Roles = v
}

func (o MembershipSchemaModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipSchemaModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_type"] = o.Type
	if !IsNil(o.Dependencies) {
		toSerialize["_dependencies"] = o.Dependencies
	}
	toSerialize["_links"] = o.Links
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["notificationMessage"] = o.NotificationMessage
	toSerialize["project"] = o.Project
	toSerialize["principal"] = o.Principal
	toSerialize["roles"] = o.Roles

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MembershipSchemaModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_type",
		"_links",
		"id",
		"createdAt",
		"updatedAt",
		"notificationMessage",
		"project",
		"principal",
		"roles",
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

	varMembershipSchemaModel := _MembershipSchemaModel{}

	err = json.Unmarshal(data, &varMembershipSchemaModel)

	if err != nil {
		return err
	}

	*o = MembershipSchemaModel(varMembershipSchemaModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_type")
		delete(additionalProperties, "_dependencies")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "notificationMessage")
		delete(additionalProperties, "project")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "roles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMembershipSchemaModel struct {
	value *MembershipSchemaModel
	isSet bool
}

func (v NullableMembershipSchemaModel) Get() *MembershipSchemaModel {
	return v.value
}

func (v *NullableMembershipSchemaModel) Set(val *MembershipSchemaModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipSchemaModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipSchemaModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipSchemaModel(val *MembershipSchemaModel) *NullableMembershipSchemaModel {
	return &NullableMembershipSchemaModel{value: val, isSet: true}
}

func (v NullableMembershipSchemaModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipSchemaModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


