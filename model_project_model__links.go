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

// checks if the ProjectModelLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectModelLinks{}

// ProjectModelLinks struct for ProjectModelLinks
type ProjectModelLinks struct {
	Update *ProjectModelLinksUpdate `json:"update,omitempty"`
	UpdateImmediately *ProjectModelLinksUpdateImmediately `json:"updateImmediately,omitempty"`
	Delete *ProjectModelLinksDelete `json:"delete,omitempty"`
	CreateWorkPackage *ProjectModelLinksCreateWorkPackage `json:"createWorkPackage,omitempty"`
	CreateWorkPackageImmediately *ProjectModelLinksCreateWorkPackageImmediately `json:"createWorkPackageImmediately,omitempty"`
	Self ProjectModelLinksSelf `json:"self"`
	Categories ProjectModelLinksCategories `json:"categories"`
	Types ProjectModelLinksTypes `json:"types"`
	Versions ProjectModelLinksVersions `json:"versions"`
	Memberships ProjectModelLinksMemberships `json:"memberships"`
	WorkPackages ProjectModelLinksWorkPackages `json:"workPackages"`
	Parent *ProjectModelLinksParent `json:"parent,omitempty"`
	Status *ProjectModelLinksStatus `json:"status,omitempty"`
	Storages []ProjectModelLinksStoragesInner `json:"storages,omitempty"`
	ProjectStorages *ProjectModelLinksProjectStorages `json:"projectStorages,omitempty"`
	Ancestors []ProjectModelLinksAncestorsInner `json:"ancestors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectModelLinks ProjectModelLinks

// NewProjectModelLinks instantiates a new ProjectModelLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectModelLinks(self ProjectModelLinksSelf, categories ProjectModelLinksCategories, types ProjectModelLinksTypes, versions ProjectModelLinksVersions, memberships ProjectModelLinksMemberships, workPackages ProjectModelLinksWorkPackages) *ProjectModelLinks {
	this := ProjectModelLinks{}
	this.Self = self
	this.Categories = categories
	this.Types = types
	this.Versions = versions
	this.Memberships = memberships
	this.WorkPackages = workPackages
	return &this
}

// NewProjectModelLinksWithDefaults instantiates a new ProjectModelLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectModelLinksWithDefaults() *ProjectModelLinks {
	this := ProjectModelLinks{}
	return &this
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetUpdate() ProjectModelLinksUpdate {
	if o == nil || IsNil(o.Update) {
		var ret ProjectModelLinksUpdate
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetUpdateOk() (*ProjectModelLinksUpdate, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given ProjectModelLinksUpdate and assigns it to the Update field.
func (o *ProjectModelLinks) SetUpdate(v ProjectModelLinksUpdate) {
	o.Update = &v
}

// GetUpdateImmediately returns the UpdateImmediately field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetUpdateImmediately() ProjectModelLinksUpdateImmediately {
	if o == nil || IsNil(o.UpdateImmediately) {
		var ret ProjectModelLinksUpdateImmediately
		return ret
	}
	return *o.UpdateImmediately
}

// GetUpdateImmediatelyOk returns a tuple with the UpdateImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetUpdateImmediatelyOk() (*ProjectModelLinksUpdateImmediately, bool) {
	if o == nil || IsNil(o.UpdateImmediately) {
		return nil, false
	}
	return o.UpdateImmediately, true
}

// HasUpdateImmediately returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasUpdateImmediately() bool {
	if o != nil && !IsNil(o.UpdateImmediately) {
		return true
	}

	return false
}

// SetUpdateImmediately gets a reference to the given ProjectModelLinksUpdateImmediately and assigns it to the UpdateImmediately field.
func (o *ProjectModelLinks) SetUpdateImmediately(v ProjectModelLinksUpdateImmediately) {
	o.UpdateImmediately = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetDelete() ProjectModelLinksDelete {
	if o == nil || IsNil(o.Delete) {
		var ret ProjectModelLinksDelete
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetDeleteOk() (*ProjectModelLinksDelete, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given ProjectModelLinksDelete and assigns it to the Delete field.
func (o *ProjectModelLinks) SetDelete(v ProjectModelLinksDelete) {
	o.Delete = &v
}

// GetCreateWorkPackage returns the CreateWorkPackage field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetCreateWorkPackage() ProjectModelLinksCreateWorkPackage {
	if o == nil || IsNil(o.CreateWorkPackage) {
		var ret ProjectModelLinksCreateWorkPackage
		return ret
	}
	return *o.CreateWorkPackage
}

// GetCreateWorkPackageOk returns a tuple with the CreateWorkPackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetCreateWorkPackageOk() (*ProjectModelLinksCreateWorkPackage, bool) {
	if o == nil || IsNil(o.CreateWorkPackage) {
		return nil, false
	}
	return o.CreateWorkPackage, true
}

// HasCreateWorkPackage returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasCreateWorkPackage() bool {
	if o != nil && !IsNil(o.CreateWorkPackage) {
		return true
	}

	return false
}

// SetCreateWorkPackage gets a reference to the given ProjectModelLinksCreateWorkPackage and assigns it to the CreateWorkPackage field.
func (o *ProjectModelLinks) SetCreateWorkPackage(v ProjectModelLinksCreateWorkPackage) {
	o.CreateWorkPackage = &v
}

// GetCreateWorkPackageImmediately returns the CreateWorkPackageImmediately field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetCreateWorkPackageImmediately() ProjectModelLinksCreateWorkPackageImmediately {
	if o == nil || IsNil(o.CreateWorkPackageImmediately) {
		var ret ProjectModelLinksCreateWorkPackageImmediately
		return ret
	}
	return *o.CreateWorkPackageImmediately
}

// GetCreateWorkPackageImmediatelyOk returns a tuple with the CreateWorkPackageImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetCreateWorkPackageImmediatelyOk() (*ProjectModelLinksCreateWorkPackageImmediately, bool) {
	if o == nil || IsNil(o.CreateWorkPackageImmediately) {
		return nil, false
	}
	return o.CreateWorkPackageImmediately, true
}

// HasCreateWorkPackageImmediately returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasCreateWorkPackageImmediately() bool {
	if o != nil && !IsNil(o.CreateWorkPackageImmediately) {
		return true
	}

	return false
}

// SetCreateWorkPackageImmediately gets a reference to the given ProjectModelLinksCreateWorkPackageImmediately and assigns it to the CreateWorkPackageImmediately field.
func (o *ProjectModelLinks) SetCreateWorkPackageImmediately(v ProjectModelLinksCreateWorkPackageImmediately) {
	o.CreateWorkPackageImmediately = &v
}

// GetSelf returns the Self field value
func (o *ProjectModelLinks) GetSelf() ProjectModelLinksSelf {
	if o == nil {
		var ret ProjectModelLinksSelf
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetSelfOk() (*ProjectModelLinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ProjectModelLinks) SetSelf(v ProjectModelLinksSelf) {
	o.Self = v
}

// GetCategories returns the Categories field value
func (o *ProjectModelLinks) GetCategories() ProjectModelLinksCategories {
	if o == nil {
		var ret ProjectModelLinksCategories
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetCategoriesOk() (*ProjectModelLinksCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value
func (o *ProjectModelLinks) SetCategories(v ProjectModelLinksCategories) {
	o.Categories = v
}

// GetTypes returns the Types field value
func (o *ProjectModelLinks) GetTypes() ProjectModelLinksTypes {
	if o == nil {
		var ret ProjectModelLinksTypes
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetTypesOk() (*ProjectModelLinksTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Types, true
}

// SetTypes sets field value
func (o *ProjectModelLinks) SetTypes(v ProjectModelLinksTypes) {
	o.Types = v
}

// GetVersions returns the Versions field value
func (o *ProjectModelLinks) GetVersions() ProjectModelLinksVersions {
	if o == nil {
		var ret ProjectModelLinksVersions
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetVersionsOk() (*ProjectModelLinksVersions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Versions, true
}

// SetVersions sets field value
func (o *ProjectModelLinks) SetVersions(v ProjectModelLinksVersions) {
	o.Versions = v
}

// GetMemberships returns the Memberships field value
func (o *ProjectModelLinks) GetMemberships() ProjectModelLinksMemberships {
	if o == nil {
		var ret ProjectModelLinksMemberships
		return ret
	}

	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetMembershipsOk() (*ProjectModelLinksMemberships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memberships, true
}

// SetMemberships sets field value
func (o *ProjectModelLinks) SetMemberships(v ProjectModelLinksMemberships) {
	o.Memberships = v
}

// GetWorkPackages returns the WorkPackages field value
func (o *ProjectModelLinks) GetWorkPackages() ProjectModelLinksWorkPackages {
	if o == nil {
		var ret ProjectModelLinksWorkPackages
		return ret
	}

	return o.WorkPackages
}

// GetWorkPackagesOk returns a tuple with the WorkPackages field value
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetWorkPackagesOk() (*ProjectModelLinksWorkPackages, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkPackages, true
}

// SetWorkPackages sets field value
func (o *ProjectModelLinks) SetWorkPackages(v ProjectModelLinksWorkPackages) {
	o.WorkPackages = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetParent() ProjectModelLinksParent {
	if o == nil || IsNil(o.Parent) {
		var ret ProjectModelLinksParent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetParentOk() (*ProjectModelLinksParent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given ProjectModelLinksParent and assigns it to the Parent field.
func (o *ProjectModelLinks) SetParent(v ProjectModelLinksParent) {
	o.Parent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetStatus() ProjectModelLinksStatus {
	if o == nil || IsNil(o.Status) {
		var ret ProjectModelLinksStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetStatusOk() (*ProjectModelLinksStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ProjectModelLinksStatus and assigns it to the Status field.
func (o *ProjectModelLinks) SetStatus(v ProjectModelLinksStatus) {
	o.Status = &v
}

// GetStorages returns the Storages field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetStorages() []ProjectModelLinksStoragesInner {
	if o == nil || IsNil(o.Storages) {
		var ret []ProjectModelLinksStoragesInner
		return ret
	}
	return o.Storages
}

// GetStoragesOk returns a tuple with the Storages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetStoragesOk() ([]ProjectModelLinksStoragesInner, bool) {
	if o == nil || IsNil(o.Storages) {
		return nil, false
	}
	return o.Storages, true
}

// HasStorages returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasStorages() bool {
	if o != nil && !IsNil(o.Storages) {
		return true
	}

	return false
}

// SetStorages gets a reference to the given []ProjectModelLinksStoragesInner and assigns it to the Storages field.
func (o *ProjectModelLinks) SetStorages(v []ProjectModelLinksStoragesInner) {
	o.Storages = v
}

// GetProjectStorages returns the ProjectStorages field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetProjectStorages() ProjectModelLinksProjectStorages {
	if o == nil || IsNil(o.ProjectStorages) {
		var ret ProjectModelLinksProjectStorages
		return ret
	}
	return *o.ProjectStorages
}

// GetProjectStoragesOk returns a tuple with the ProjectStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetProjectStoragesOk() (*ProjectModelLinksProjectStorages, bool) {
	if o == nil || IsNil(o.ProjectStorages) {
		return nil, false
	}
	return o.ProjectStorages, true
}

// HasProjectStorages returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasProjectStorages() bool {
	if o != nil && !IsNil(o.ProjectStorages) {
		return true
	}

	return false
}

// SetProjectStorages gets a reference to the given ProjectModelLinksProjectStorages and assigns it to the ProjectStorages field.
func (o *ProjectModelLinks) SetProjectStorages(v ProjectModelLinksProjectStorages) {
	o.ProjectStorages = &v
}

// GetAncestors returns the Ancestors field value if set, zero value otherwise.
func (o *ProjectModelLinks) GetAncestors() []ProjectModelLinksAncestorsInner {
	if o == nil || IsNil(o.Ancestors) {
		var ret []ProjectModelLinksAncestorsInner
		return ret
	}
	return o.Ancestors
}

// GetAncestorsOk returns a tuple with the Ancestors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectModelLinks) GetAncestorsOk() ([]ProjectModelLinksAncestorsInner, bool) {
	if o == nil || IsNil(o.Ancestors) {
		return nil, false
	}
	return o.Ancestors, true
}

// HasAncestors returns a boolean if a field has been set.
func (o *ProjectModelLinks) HasAncestors() bool {
	if o != nil && !IsNil(o.Ancestors) {
		return true
	}

	return false
}

// SetAncestors gets a reference to the given []ProjectModelLinksAncestorsInner and assigns it to the Ancestors field.
func (o *ProjectModelLinks) SetAncestors(v []ProjectModelLinksAncestorsInner) {
	o.Ancestors = v
}

func (o ProjectModelLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectModelLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.UpdateImmediately) {
		toSerialize["updateImmediately"] = o.UpdateImmediately
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.CreateWorkPackage) {
		toSerialize["createWorkPackage"] = o.CreateWorkPackage
	}
	if !IsNil(o.CreateWorkPackageImmediately) {
		toSerialize["createWorkPackageImmediately"] = o.CreateWorkPackageImmediately
	}
	toSerialize["self"] = o.Self
	toSerialize["categories"] = o.Categories
	toSerialize["types"] = o.Types
	toSerialize["versions"] = o.Versions
	toSerialize["memberships"] = o.Memberships
	toSerialize["workPackages"] = o.WorkPackages
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Storages) {
		toSerialize["storages"] = o.Storages
	}
	if !IsNil(o.ProjectStorages) {
		toSerialize["projectStorages"] = o.ProjectStorages
	}
	if !IsNil(o.Ancestors) {
		toSerialize["ancestors"] = o.Ancestors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectModelLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
		"categories",
		"types",
		"versions",
		"memberships",
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

	varProjectModelLinks := _ProjectModelLinks{}

	err = json.Unmarshal(data, &varProjectModelLinks)

	if err != nil {
		return err
	}

	*o = ProjectModelLinks(varProjectModelLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "update")
		delete(additionalProperties, "updateImmediately")
		delete(additionalProperties, "delete")
		delete(additionalProperties, "createWorkPackage")
		delete(additionalProperties, "createWorkPackageImmediately")
		delete(additionalProperties, "self")
		delete(additionalProperties, "categories")
		delete(additionalProperties, "types")
		delete(additionalProperties, "versions")
		delete(additionalProperties, "memberships")
		delete(additionalProperties, "workPackages")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "status")
		delete(additionalProperties, "storages")
		delete(additionalProperties, "projectStorages")
		delete(additionalProperties, "ancestors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectModelLinks struct {
	value *ProjectModelLinks
	isSet bool
}

func (v NullableProjectModelLinks) Get() *ProjectModelLinks {
	return v.value
}

func (v *NullableProjectModelLinks) Set(val *ProjectModelLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectModelLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectModelLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectModelLinks(val *ProjectModelLinks) *NullableProjectModelLinks {
	return &NullableProjectModelLinks{value: val, isSet: true}
}

func (v NullableProjectModelLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectModelLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


