/*
OpenProject API V3 (Stable)

You're looking at the current **stable** documentation of the OpenProject APIv3. If you're interested in the current development version, please go to [github.com/opf](https://github.com/opf/openproject/tree/dev/docs/api/apiv3).  ## Introduction  The documentation for the APIv3 is written according to the [OpenAPI 3.0 Specification](https://swagger.io/specification/). You can either view the static version of this documentation on the [website](https://www.openproject.org/docs/api/introduction/) or the interactive version, rendered with [OpenAPI Explorer](https://github.com/Rhosys/openapi-explorer/blob/main/README.md), in your OpenProject installation under `/api/docs`. In the latter you can try out the various API endpoints directly interacting with our OpenProject data. Moreover you can access the specification source itself under `/api/v3/spec.json` and `/api/v3/spec.yml` (e.g. [here](https://community.openproject.org/api/v3/spec.yml)).  The APIv3 is a hypermedia REST API, a shorthand for \"Hypermedia As The Engine Of Application State\" (HATEOAS). This means that each endpoint of this API will have links to other resources or actions defined in the resulting body.  These related resources and actions for any given resource will be context sensitive. For example, only actions that the authenticated user can take are being rendered. This can be used to dynamically identify actions that the user might take for any given response.  As an example, if you fetch a work package through the [Work Package endpoint](https://www.openproject.org/docs/api/endpoints/work-packages/), the `update` link will only be present when the user you authenticated has been granted a permission to update the work package in the assigned project.  ## HAL+JSON  HAL is a simple format that gives a consistent and easy way to hyperlink between resources in your API. Read more in the following specification: [https://tools.ietf.org/html/draft-kelly-json-hal-08](https://tools.ietf.org/html/draft-kelly-json-hal-08)  **OpenProject API implementation of HAL+JSON format** enriches JSON and introduces a few meta properties:  - `_type` - specifies the type of the resource (e.g.: WorkPackage, Project) - `_links` - contains all related resource and action links available for the resource - `_embedded` - contains all embedded objects  HAL does not guarantee that embedded resources are embedded in their full representation, they might as well be partially represented (e.g. some properties can be left out). However in this API you have the guarantee that whenever a resource is **embedded**, it is embedded in its **full representation**.  ## API response structure  All API responses contain a single HAL+JSON object, even collections of objects are technically represented by a single HAL+JSON object that itself contains its members. More details on collections can be found in the [Collections Section](https://www.openproject.org/docs/api/collections/).  ## Authentication  The API supports the following authentication schemes: OAuth2, session based authentication, and basic auth.  Depending on the settings of the OpenProject instance many resources can be accessed without being authenticated. In case the instance requires authentication on all requests the client will receive an **HTTP 401** status code in response to any request.  Otherwise unauthenticated clients have all the permissions of the anonymous user.  ### Session-based Authentication  This means you have to login to OpenProject via the Web-Interface to be authenticated in the API. This method is well-suited for clients acting within the browser, like the Angular-Client built into OpenProject.  In this case, you always need to pass the HTTP header `X-Requested-With \"XMLHttpRequest\"` for authentication.  ### API Key through Basic Auth  Users can authenticate towards the API v3 using basic auth with the user name `apikey` (NOT your login) and the API key as the password. Users can find their API key on their account page.  Example:  ```shell API_KEY=2519132cdf62dcf5a66fd96394672079f9e9cad1 curl -u apikey:$API_KEY https://community.openproject.org/api/v3/users/42 ```  ### OAuth2.0 authentication  OpenProject allows authentication and authorization with OAuth2 with *Authorization code flow*, as well as *Client credentials* operation modes.  To get started, you first need to register an application in the OpenProject OAuth administration section of your installation. This will save an entry for your application with a client unique identifier (`client_id`) and an accompanying secret key (`client_secret`).  You can then use one the following guides to perform the supported OAuth 2.0 flows:  - [Authorization code flow](https://oauth.net/2/grant-types/authorization-code)  - [Authorization code flow with PKCE](https://doorkeeper.gitbook.io/guides/ruby-on-rails/pkce-flow), recommended for clients unable to keep the client_secret confidential.  - [Client credentials](https://oauth.net/2/grant-types/client-credentials/) - Requires an application to be bound to an impersonating user for non-public access  ### Why not username and password?  The simplest way to do basic auth would be to use a user's username and password naturally. However, OpenProject already has supported API keys in the past for the API v2, though not through basic auth.  Using **username and password** directly would have some advantages:  * It is intuitive for the user who then just has to provide those just as they would when logging into OpenProject.  * No extra logic for token management necessary.  On the other hand using **API keys** has some advantages too, which is why we went for that:  * If compromised while saved on an insecure client the user only has to regenerate the API key instead of changing their password, too.  * They are naturally long and random which makes them invulnerable to dictionary attacks and harder to crack in general.  Most importantly users may not actually have a password to begin with. Specifically when they have registered through an OpenID Connect provider.  ## Cross-Origin Resource Sharing (CORS)  By default, the OpenProject API is _not_ responding with any CORS headers. If you want to allow cross-domain AJAX calls against your OpenProject instance, you need to enable CORS headers being returned.  Please see [our API settings documentation](https://www.openproject.org/docs/system-admin-guide/api-and-webhooks/) on how to selectively enable CORS.  ## Allowed HTTP methods  - `GET` - Get a single resource or collection of resources  - `POST` - Create a new resource or perform  - `PATCH` - Update a resource  - `DELETE` - Delete a resource  ## Compression  Responses are compressed if requested by the client. Currently [gzip](https://www.gzip.org/) and [deflate](https://tools.ietf.org/html/rfc1951) are supported. The client signals the desired compression by setting the [`Accept-Encoding` header](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.3). If no `Accept-Encoding` header is send, `Accept-Encoding: identity` is assumed which will result in the API responding uncompressed.

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openproject

import (
	"encoding/json"
)

// checks if the WorkPackagePatchModelLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkPackagePatchModelLinks{}

// WorkPackagePatchModelLinks struct for WorkPackagePatchModelLinks
type WorkPackagePatchModelLinks struct {
	Assignee *WorkPackageModelLinksAssignee `json:"assignee,omitempty"`
	Budget *WorkPackageModelLinksBudget `json:"budget,omitempty"`
	Category *WorkPackageModelLinksCategory `json:"category,omitempty"`
	Parent *WorkPackageModelLinksParent `json:"parent,omitempty"`
	Priority *WorkPackageModelLinksPriority `json:"priority,omitempty"`
	Project *WorkPackageModelLinksProject `json:"project,omitempty"`
	Responsible *WorkPackageModelLinksResponsible `json:"responsible,omitempty"`
	Status *WorkPackageModelLinksStatus `json:"status,omitempty"`
	Type *WorkPackageModelLinksType `json:"type,omitempty"`
	Version *WorkPackageModelLinksVersion `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkPackagePatchModelLinks WorkPackagePatchModelLinks

// NewWorkPackagePatchModelLinks instantiates a new WorkPackagePatchModelLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkPackagePatchModelLinks() *WorkPackagePatchModelLinks {
	this := WorkPackagePatchModelLinks{}
	return &this
}

// NewWorkPackagePatchModelLinksWithDefaults instantiates a new WorkPackagePatchModelLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkPackagePatchModelLinksWithDefaults() *WorkPackagePatchModelLinks {
	this := WorkPackagePatchModelLinks{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetAssignee() WorkPackageModelLinksAssignee {
	if o == nil || IsNil(o.Assignee) {
		var ret WorkPackageModelLinksAssignee
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetAssigneeOk() (*WorkPackageModelLinksAssignee, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given WorkPackageModelLinksAssignee and assigns it to the Assignee field.
func (o *WorkPackagePatchModelLinks) SetAssignee(v WorkPackageModelLinksAssignee) {
	o.Assignee = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetBudget() WorkPackageModelLinksBudget {
	if o == nil || IsNil(o.Budget) {
		var ret WorkPackageModelLinksBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetBudgetOk() (*WorkPackageModelLinksBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given WorkPackageModelLinksBudget and assigns it to the Budget field.
func (o *WorkPackagePatchModelLinks) SetBudget(v WorkPackageModelLinksBudget) {
	o.Budget = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetCategory() WorkPackageModelLinksCategory {
	if o == nil || IsNil(o.Category) {
		var ret WorkPackageModelLinksCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetCategoryOk() (*WorkPackageModelLinksCategory, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given WorkPackageModelLinksCategory and assigns it to the Category field.
func (o *WorkPackagePatchModelLinks) SetCategory(v WorkPackageModelLinksCategory) {
	o.Category = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetParent() WorkPackageModelLinksParent {
	if o == nil || IsNil(o.Parent) {
		var ret WorkPackageModelLinksParent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetParentOk() (*WorkPackageModelLinksParent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given WorkPackageModelLinksParent and assigns it to the Parent field.
func (o *WorkPackagePatchModelLinks) SetParent(v WorkPackageModelLinksParent) {
	o.Parent = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetPriority() WorkPackageModelLinksPriority {
	if o == nil || IsNil(o.Priority) {
		var ret WorkPackageModelLinksPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetPriorityOk() (*WorkPackageModelLinksPriority, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given WorkPackageModelLinksPriority and assigns it to the Priority field.
func (o *WorkPackagePatchModelLinks) SetPriority(v WorkPackageModelLinksPriority) {
	o.Priority = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetProject() WorkPackageModelLinksProject {
	if o == nil || IsNil(o.Project) {
		var ret WorkPackageModelLinksProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetProjectOk() (*WorkPackageModelLinksProject, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given WorkPackageModelLinksProject and assigns it to the Project field.
func (o *WorkPackagePatchModelLinks) SetProject(v WorkPackageModelLinksProject) {
	o.Project = &v
}

// GetResponsible returns the Responsible field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetResponsible() WorkPackageModelLinksResponsible {
	if o == nil || IsNil(o.Responsible) {
		var ret WorkPackageModelLinksResponsible
		return ret
	}
	return *o.Responsible
}

// GetResponsibleOk returns a tuple with the Responsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetResponsibleOk() (*WorkPackageModelLinksResponsible, bool) {
	if o == nil || IsNil(o.Responsible) {
		return nil, false
	}
	return o.Responsible, true
}

// HasResponsible returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasResponsible() bool {
	if o != nil && !IsNil(o.Responsible) {
		return true
	}

	return false
}

// SetResponsible gets a reference to the given WorkPackageModelLinksResponsible and assigns it to the Responsible field.
func (o *WorkPackagePatchModelLinks) SetResponsible(v WorkPackageModelLinksResponsible) {
	o.Responsible = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetStatus() WorkPackageModelLinksStatus {
	if o == nil || IsNil(o.Status) {
		var ret WorkPackageModelLinksStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetStatusOk() (*WorkPackageModelLinksStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WorkPackageModelLinksStatus and assigns it to the Status field.
func (o *WorkPackagePatchModelLinks) SetStatus(v WorkPackageModelLinksStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetType() WorkPackageModelLinksType {
	if o == nil || IsNil(o.Type) {
		var ret WorkPackageModelLinksType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetTypeOk() (*WorkPackageModelLinksType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given WorkPackageModelLinksType and assigns it to the Type field.
func (o *WorkPackagePatchModelLinks) SetType(v WorkPackageModelLinksType) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkPackagePatchModelLinks) GetVersion() WorkPackageModelLinksVersion {
	if o == nil || IsNil(o.Version) {
		var ret WorkPackageModelLinksVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModelLinks) GetVersionOk() (*WorkPackageModelLinksVersion, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkPackagePatchModelLinks) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given WorkPackageModelLinksVersion and assigns it to the Version field.
func (o *WorkPackagePatchModelLinks) SetVersion(v WorkPackageModelLinksVersion) {
	o.Version = &v
}

func (o WorkPackagePatchModelLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkPackagePatchModelLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Responsible) {
		toSerialize["responsible"] = o.Responsible
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkPackagePatchModelLinks) UnmarshalJSON(data []byte) (err error) {
	varWorkPackagePatchModelLinks := _WorkPackagePatchModelLinks{}

	err = json.Unmarshal(data, &varWorkPackagePatchModelLinks)

	if err != nil {
		return err
	}

	*o = WorkPackagePatchModelLinks(varWorkPackagePatchModelLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assignee")
		delete(additionalProperties, "budget")
		delete(additionalProperties, "category")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "project")
		delete(additionalProperties, "responsible")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkPackagePatchModelLinks struct {
	value *WorkPackagePatchModelLinks
	isSet bool
}

func (v NullableWorkPackagePatchModelLinks) Get() *WorkPackagePatchModelLinks {
	return v.value
}

func (v *NullableWorkPackagePatchModelLinks) Set(val *WorkPackagePatchModelLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkPackagePatchModelLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkPackagePatchModelLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkPackagePatchModelLinks(val *WorkPackagePatchModelLinks) *NullableWorkPackagePatchModelLinks {
	return &NullableWorkPackagePatchModelLinks{value: val, isSet: true}
}

func (v NullableWorkPackagePatchModelLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkPackagePatchModelLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


