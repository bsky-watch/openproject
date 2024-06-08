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

// checks if the ProjectStorageModelLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectStorageModelLinks{}

// ProjectStorageModelLinks struct for ProjectStorageModelLinks
type ProjectStorageModelLinks struct {
	Self ProjectStorageModelLinksSelf `json:"self"`
	Creator ProjectStorageModelLinksCreator `json:"creator"`
	Storage ProjectStorageModelLinksStorage `json:"storage"`
	Project ProjectStorageModelLinksProject `json:"project"`
	ProjectFolder *ProjectStorageModelLinksProjectFolder `json:"projectFolder,omitempty"`
	Open *ProjectStorageModelLinksOpen `json:"open,omitempty"`
	OpenWithConnectionEnsured *ProjectStorageModelLinksOpenWithConnectionEnsured `json:"openWithConnectionEnsured,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectStorageModelLinks ProjectStorageModelLinks

// NewProjectStorageModelLinks instantiates a new ProjectStorageModelLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectStorageModelLinks(self ProjectStorageModelLinksSelf, creator ProjectStorageModelLinksCreator, storage ProjectStorageModelLinksStorage, project ProjectStorageModelLinksProject) *ProjectStorageModelLinks {
	this := ProjectStorageModelLinks{}
	this.Self = self
	this.Creator = creator
	this.Storage = storage
	this.Project = project
	return &this
}

// NewProjectStorageModelLinksWithDefaults instantiates a new ProjectStorageModelLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectStorageModelLinksWithDefaults() *ProjectStorageModelLinks {
	this := ProjectStorageModelLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *ProjectStorageModelLinks) GetSelf() ProjectStorageModelLinksSelf {
	if o == nil {
		var ret ProjectStorageModelLinksSelf
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ProjectStorageModelLinks) GetSelfOk() (*ProjectStorageModelLinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ProjectStorageModelLinks) SetSelf(v ProjectStorageModelLinksSelf) {
	o.Self = v
}

// GetCreator returns the Creator field value
func (o *ProjectStorageModelLinks) GetCreator() ProjectStorageModelLinksCreator {
	if o == nil {
		var ret ProjectStorageModelLinksCreator
		return ret
	}

	return o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
func (o *ProjectStorageModelLinks) GetCreatorOk() (*ProjectStorageModelLinksCreator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creator, true
}

// SetCreator sets field value
func (o *ProjectStorageModelLinks) SetCreator(v ProjectStorageModelLinksCreator) {
	o.Creator = v
}

// GetStorage returns the Storage field value
func (o *ProjectStorageModelLinks) GetStorage() ProjectStorageModelLinksStorage {
	if o == nil {
		var ret ProjectStorageModelLinksStorage
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *ProjectStorageModelLinks) GetStorageOk() (*ProjectStorageModelLinksStorage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *ProjectStorageModelLinks) SetStorage(v ProjectStorageModelLinksStorage) {
	o.Storage = v
}

// GetProject returns the Project field value
func (o *ProjectStorageModelLinks) GetProject() ProjectStorageModelLinksProject {
	if o == nil {
		var ret ProjectStorageModelLinksProject
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ProjectStorageModelLinks) GetProjectOk() (*ProjectStorageModelLinksProject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ProjectStorageModelLinks) SetProject(v ProjectStorageModelLinksProject) {
	o.Project = v
}

// GetProjectFolder returns the ProjectFolder field value if set, zero value otherwise.
func (o *ProjectStorageModelLinks) GetProjectFolder() ProjectStorageModelLinksProjectFolder {
	if o == nil || IsNil(o.ProjectFolder) {
		var ret ProjectStorageModelLinksProjectFolder
		return ret
	}
	return *o.ProjectFolder
}

// GetProjectFolderOk returns a tuple with the ProjectFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectStorageModelLinks) GetProjectFolderOk() (*ProjectStorageModelLinksProjectFolder, bool) {
	if o == nil || IsNil(o.ProjectFolder) {
		return nil, false
	}
	return o.ProjectFolder, true
}

// HasProjectFolder returns a boolean if a field has been set.
func (o *ProjectStorageModelLinks) HasProjectFolder() bool {
	if o != nil && !IsNil(o.ProjectFolder) {
		return true
	}

	return false
}

// SetProjectFolder gets a reference to the given ProjectStorageModelLinksProjectFolder and assigns it to the ProjectFolder field.
func (o *ProjectStorageModelLinks) SetProjectFolder(v ProjectStorageModelLinksProjectFolder) {
	o.ProjectFolder = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *ProjectStorageModelLinks) GetOpen() ProjectStorageModelLinksOpen {
	if o == nil || IsNil(o.Open) {
		var ret ProjectStorageModelLinksOpen
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectStorageModelLinks) GetOpenOk() (*ProjectStorageModelLinksOpen, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *ProjectStorageModelLinks) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given ProjectStorageModelLinksOpen and assigns it to the Open field.
func (o *ProjectStorageModelLinks) SetOpen(v ProjectStorageModelLinksOpen) {
	o.Open = &v
}

// GetOpenWithConnectionEnsured returns the OpenWithConnectionEnsured field value if set, zero value otherwise.
func (o *ProjectStorageModelLinks) GetOpenWithConnectionEnsured() ProjectStorageModelLinksOpenWithConnectionEnsured {
	if o == nil || IsNil(o.OpenWithConnectionEnsured) {
		var ret ProjectStorageModelLinksOpenWithConnectionEnsured
		return ret
	}
	return *o.OpenWithConnectionEnsured
}

// GetOpenWithConnectionEnsuredOk returns a tuple with the OpenWithConnectionEnsured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectStorageModelLinks) GetOpenWithConnectionEnsuredOk() (*ProjectStorageModelLinksOpenWithConnectionEnsured, bool) {
	if o == nil || IsNil(o.OpenWithConnectionEnsured) {
		return nil, false
	}
	return o.OpenWithConnectionEnsured, true
}

// HasOpenWithConnectionEnsured returns a boolean if a field has been set.
func (o *ProjectStorageModelLinks) HasOpenWithConnectionEnsured() bool {
	if o != nil && !IsNil(o.OpenWithConnectionEnsured) {
		return true
	}

	return false
}

// SetOpenWithConnectionEnsured gets a reference to the given ProjectStorageModelLinksOpenWithConnectionEnsured and assigns it to the OpenWithConnectionEnsured field.
func (o *ProjectStorageModelLinks) SetOpenWithConnectionEnsured(v ProjectStorageModelLinksOpenWithConnectionEnsured) {
	o.OpenWithConnectionEnsured = &v
}

func (o ProjectStorageModelLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectStorageModelLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["creator"] = o.Creator
	toSerialize["storage"] = o.Storage
	toSerialize["project"] = o.Project
	if !IsNil(o.ProjectFolder) {
		toSerialize["projectFolder"] = o.ProjectFolder
	}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.OpenWithConnectionEnsured) {
		toSerialize["openWithConnectionEnsured"] = o.OpenWithConnectionEnsured
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectStorageModelLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
		"creator",
		"storage",
		"project",
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

	varProjectStorageModelLinks := _ProjectStorageModelLinks{}

	err = json.Unmarshal(data, &varProjectStorageModelLinks)

	if err != nil {
		return err
	}

	*o = ProjectStorageModelLinks(varProjectStorageModelLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "creator")
		delete(additionalProperties, "storage")
		delete(additionalProperties, "project")
		delete(additionalProperties, "projectFolder")
		delete(additionalProperties, "open")
		delete(additionalProperties, "openWithConnectionEnsured")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectStorageModelLinks struct {
	value *ProjectStorageModelLinks
	isSet bool
}

func (v NullableProjectStorageModelLinks) Get() *ProjectStorageModelLinks {
	return v.value
}

func (v *NullableProjectStorageModelLinks) Set(val *ProjectStorageModelLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectStorageModelLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectStorageModelLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectStorageModelLinks(val *ProjectStorageModelLinks) *NullableProjectStorageModelLinks {
	return &NullableProjectStorageModelLinks{value: val, isSet: true}
}

func (v NullableProjectStorageModelLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectStorageModelLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


