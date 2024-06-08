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

// checks if the QueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryModel{}

// QueryModel struct for QueryModel
type QueryModel struct {
	// Query id
	Id *int64 `json:"id,omitempty"`
	// Query name
	Name *string `json:"name,omitempty"`
	// A set of QueryFilters which will be applied to the work packages to determine the resulting work packages
	Filters []QueryFilterInstanceSchemaModel `json:"filters,omitempty"`
	// Should sums (of supported properties) be shown?
	Sums *bool `json:"sums,omitempty"`
	// Should the timeline mode be shown?
	// Deprecated
	TimelineVisible *bool `json:"timelineVisible,omitempty"`
	// Which labels are shown in the timeline, empty when default
	TimelineLabels []string `json:"timelineLabels,omitempty"`
	// Which zoom level should the timeline be rendered in?
	// Deprecated
	TimelineZoomLevel *string `json:"timelineZoomLevel,omitempty"`
	// Timestamps to filter by when showing changed attributes on work packages.
	Timestamps []string `json:"timestamps,omitempty"`
	// Which highlighting mode should the table have?
	// Deprecated
	HighlightingMode *string `json:"highlightingMode,omitempty"`
	// Should the hierarchy mode be enabled?
	// Deprecated
	ShowHierarchies *bool `json:"showHierarchies,omitempty"`
	// Should the query be hidden from the query list?
	// Deprecated
	Hidden *bool `json:"hidden,omitempty"`
	// Can users besides the owner see the query?
	Public *bool `json:"public,omitempty"`
	// Should the query be highlighted to the user?
	Starred *bool `json:"starred,omitempty"`
	// Time of creation
	CreatedAt time.Time `json:"createdAt"`
	// Time of the most recent change to the query
	UpdatedAt time.Time `json:"updatedAt"`
	Links *QueryModelLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryModel QueryModel

// NewQueryModel instantiates a new QueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryModel(createdAt time.Time, updatedAt time.Time) *QueryModel {
	this := QueryModel{}
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewQueryModelWithDefaults instantiates a new QueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryModelWithDefaults() *QueryModel {
	this := QueryModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QueryModel) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryModel) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QueryModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *QueryModel) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueryModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QueryModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueryModel) SetName(v string) {
	o.Name = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *QueryModel) GetFilters() []QueryFilterInstanceSchemaModel {
	if o == nil || IsNil(o.Filters) {
		var ret []QueryFilterInstanceSchemaModel
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryModel) GetFiltersOk() ([]QueryFilterInstanceSchemaModel, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *QueryModel) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []QueryFilterInstanceSchemaModel and assigns it to the Filters field.
func (o *QueryModel) SetFilters(v []QueryFilterInstanceSchemaModel) {
	o.Filters = v
}

// GetSums returns the Sums field value if set, zero value otherwise.
func (o *QueryModel) GetSums() bool {
	if o == nil || IsNil(o.Sums) {
		var ret bool
		return ret
	}
	return *o.Sums
}

// GetSumsOk returns a tuple with the Sums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryModel) GetSumsOk() (*bool, bool) {
	if o == nil || IsNil(o.Sums) {
		return nil, false
	}
	return o.Sums, true
}

// HasSums returns a boolean if a field has been set.
func (o *QueryModel) HasSums() bool {
	if o != nil && !IsNil(o.Sums) {
		return true
	}

	return false
}

// SetSums gets a reference to the given bool and assigns it to the Sums field.
func (o *QueryModel) SetSums(v bool) {
	o.Sums = &v
}

// GetTimelineVisible returns the TimelineVisible field value if set, zero value otherwise.
// Deprecated
func (o *QueryModel) GetTimelineVisible() bool {
	if o == nil || IsNil(o.TimelineVisible) {
		var ret bool
		return ret
	}
	return *o.TimelineVisible
}

// GetTimelineVisibleOk returns a tuple with the TimelineVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryModel) GetTimelineVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.TimelineVisible) {
		return nil, false
	}
	return o.TimelineVisible, true
}

// HasTimelineVisible returns a boolean if a field has been set.
func (o *QueryModel) HasTimelineVisible() bool {
	if o != nil && !IsNil(o.TimelineVisible) {
		return true
	}

	return false
}

// SetTimelineVisible gets a reference to the given bool and assigns it to the TimelineVisible field.
// Deprecated
func (o *QueryModel) SetTimelineVisible(v bool) {
	o.TimelineVisible = &v
}

// GetTimelineLabels returns the TimelineLabels field value if set, zero value otherwise.
func (o *QueryModel) GetTimelineLabels() []string {
	if o == nil || IsNil(o.TimelineLabels) {
		var ret []string
		return ret
	}
	return o.TimelineLabels
}

// GetTimelineLabelsOk returns a tuple with the TimelineLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryModel) GetTimelineLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.TimelineLabels) {
		return nil, false
	}
	return o.TimelineLabels, true
}

// HasTimelineLabels returns a boolean if a field has been set.
func (o *QueryModel) HasTimelineLabels() bool {
	if o != nil && !IsNil(o.TimelineLabels) {
		return true
	}

	return false
}

// SetTimelineLabels gets a reference to the given []string and assigns it to the TimelineLabels field.
func (o *QueryModel) SetTimelineLabels(v []string) {
	o.TimelineLabels = v
}

// GetTimelineZoomLevel returns the TimelineZoomLevel field value if set, zero value otherwise.
// Deprecated
func (o *QueryModel) GetTimelineZoomLevel() string {
	if o == nil || IsNil(o.TimelineZoomLevel) {
		var ret string
		return ret
	}
	return *o.TimelineZoomLevel
}

// GetTimelineZoomLevelOk returns a tuple with the TimelineZoomLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryModel) GetTimelineZoomLevelOk() (*string, bool) {
	if o == nil || IsNil(o.TimelineZoomLevel) {
		return nil, false
	}
	return o.TimelineZoomLevel, true
}

// HasTimelineZoomLevel returns a boolean if a field has been set.
func (o *QueryModel) HasTimelineZoomLevel() bool {
	if o != nil && !IsNil(o.TimelineZoomLevel) {
		return true
	}

	return false
}

// SetTimelineZoomLevel gets a reference to the given string and assigns it to the TimelineZoomLevel field.
// Deprecated
func (o *QueryModel) SetTimelineZoomLevel(v string) {
	o.TimelineZoomLevel = &v
}

// GetTimestamps returns the Timestamps field value if set, zero value otherwise.
func (o *QueryModel) GetTimestamps() []string {
	if o == nil || IsNil(o.Timestamps) {
		var ret []string
		return ret
	}
	return o.Timestamps
}

// GetTimestampsOk returns a tuple with the Timestamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryModel) GetTimestampsOk() ([]string, bool) {
	if o == nil || IsNil(o.Timestamps) {
		return nil, false
	}
	return o.Timestamps, true
}

// HasTimestamps returns a boolean if a field has been set.
func (o *QueryModel) HasTimestamps() bool {
	if o != nil && !IsNil(o.Timestamps) {
		return true
	}

	return false
}

// SetTimestamps gets a reference to the given []string and assigns it to the Timestamps field.
func (o *QueryModel) SetTimestamps(v []string) {
	o.Timestamps = v
}

// GetHighlightingMode returns the HighlightingMode field value if set, zero value otherwise.
// Deprecated
func (o *QueryModel) GetHighlightingMode() string {
	if o == nil || IsNil(o.HighlightingMode) {
		var ret string
		return ret
	}
	return *o.HighlightingMode
}

// GetHighlightingModeOk returns a tuple with the HighlightingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryModel) GetHighlightingModeOk() (*string, bool) {
	if o == nil || IsNil(o.HighlightingMode) {
		return nil, false
	}
	return o.HighlightingMode, true
}

// HasHighlightingMode returns a boolean if a field has been set.
func (o *QueryModel) HasHighlightingMode() bool {
	if o != nil && !IsNil(o.HighlightingMode) {
		return true
	}

	return false
}

// SetHighlightingMode gets a reference to the given string and assigns it to the HighlightingMode field.
// Deprecated
func (o *QueryModel) SetHighlightingMode(v string) {
	o.HighlightingMode = &v
}

// GetShowHierarchies returns the ShowHierarchies field value if set, zero value otherwise.
// Deprecated
func (o *QueryModel) GetShowHierarchies() bool {
	if o == nil || IsNil(o.ShowHierarchies) {
		var ret bool
		return ret
	}
	return *o.ShowHierarchies
}

// GetShowHierarchiesOk returns a tuple with the ShowHierarchies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryModel) GetShowHierarchiesOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowHierarchies) {
		return nil, false
	}
	return o.ShowHierarchies, true
}

// HasShowHierarchies returns a boolean if a field has been set.
func (o *QueryModel) HasShowHierarchies() bool {
	if o != nil && !IsNil(o.ShowHierarchies) {
		return true
	}

	return false
}

// SetShowHierarchies gets a reference to the given bool and assigns it to the ShowHierarchies field.
// Deprecated
func (o *QueryModel) SetShowHierarchies(v bool) {
	o.ShowHierarchies = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
// Deprecated
func (o *QueryModel) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryModel) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *QueryModel) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
// Deprecated
func (o *QueryModel) SetHidden(v bool) {
	o.Hidden = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *QueryModel) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryModel) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *QueryModel) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *QueryModel) SetPublic(v bool) {
	o.Public = &v
}

// GetStarred returns the Starred field value if set, zero value otherwise.
func (o *QueryModel) GetStarred() bool {
	if o == nil || IsNil(o.Starred) {
		var ret bool
		return ret
	}
	return *o.Starred
}

// GetStarredOk returns a tuple with the Starred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryModel) GetStarredOk() (*bool, bool) {
	if o == nil || IsNil(o.Starred) {
		return nil, false
	}
	return o.Starred, true
}

// HasStarred returns a boolean if a field has been set.
func (o *QueryModel) HasStarred() bool {
	if o != nil && !IsNil(o.Starred) {
		return true
	}

	return false
}

// SetStarred gets a reference to the given bool and assigns it to the Starred field.
func (o *QueryModel) SetStarred(v bool) {
	o.Starred = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *QueryModel) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *QueryModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *QueryModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *QueryModel) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *QueryModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *QueryModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *QueryModel) GetLinks() QueryModelLinks {
	if o == nil || IsNil(o.Links) {
		var ret QueryModelLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryModel) GetLinksOk() (*QueryModelLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *QueryModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given QueryModelLinks and assigns it to the Links field.
func (o *QueryModel) SetLinks(v QueryModelLinks) {
	o.Links = &v
}

func (o QueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Sums) {
		toSerialize["sums"] = o.Sums
	}
	if !IsNil(o.TimelineVisible) {
		toSerialize["timelineVisible"] = o.TimelineVisible
	}
	if !IsNil(o.TimelineLabels) {
		toSerialize["timelineLabels"] = o.TimelineLabels
	}
	if !IsNil(o.TimelineZoomLevel) {
		toSerialize["timelineZoomLevel"] = o.TimelineZoomLevel
	}
	if !IsNil(o.Timestamps) {
		toSerialize["timestamps"] = o.Timestamps
	}
	if !IsNil(o.HighlightingMode) {
		toSerialize["highlightingMode"] = o.HighlightingMode
	}
	if !IsNil(o.ShowHierarchies) {
		toSerialize["showHierarchies"] = o.ShowHierarchies
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.Starred) {
		toSerialize["starred"] = o.Starred
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"updatedAt",
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

	varQueryModel := _QueryModel{}

	err = json.Unmarshal(data, &varQueryModel)

	if err != nil {
		return err
	}

	*o = QueryModel(varQueryModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "sums")
		delete(additionalProperties, "timelineVisible")
		delete(additionalProperties, "timelineLabels")
		delete(additionalProperties, "timelineZoomLevel")
		delete(additionalProperties, "timestamps")
		delete(additionalProperties, "highlightingMode")
		delete(additionalProperties, "showHierarchies")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "public")
		delete(additionalProperties, "starred")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryModel struct {
	value *QueryModel
	isSet bool
}

func (v NullableQueryModel) Get() *QueryModel {
	return v.value
}

func (v *NullableQueryModel) Set(val *QueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryModel(val *QueryModel) *NullableQueryModel {
	return &NullableQueryModel{value: val, isSet: true}
}

func (v NullableQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


