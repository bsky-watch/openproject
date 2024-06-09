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

// checks if the TimeEntryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeEntryModel{}

// TimeEntryModel struct for TimeEntryModel
type TimeEntryModel struct {
	// The id of the time entry
	Id *int `json:"id,omitempty"`
	Comment *TimeEntryModelComment `json:"comment,omitempty"`
	// The date the expenditure is booked for
	SpentOn *string `json:"spentOn,omitempty"`
	// The time quantifying the expenditure
	Hours *string `json:"hours,omitempty"`
	// Whether the time entry is actively tracking time
	Ongoing *bool `json:"ongoing,omitempty"`
	// The time the time entry was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the time entry was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Links *TimeEntryModelLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimeEntryModel TimeEntryModel

// NewTimeEntryModel instantiates a new TimeEntryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeEntryModel() *TimeEntryModel {
	this := TimeEntryModel{}
	return &this
}

// NewTimeEntryModelWithDefaults instantiates a new TimeEntryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeEntryModelWithDefaults() *TimeEntryModel {
	this := TimeEntryModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TimeEntryModel) GetId() int {
	if o == nil || IsNil(o.Id) {
		var ret int
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeEntryModel) GetIdOk() (*int, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TimeEntryModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int and assigns it to the Id field.
func (o *TimeEntryModel) SetId(v int) {
	o.Id = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TimeEntryModel) GetComment() TimeEntryModelComment {
	if o == nil || IsNil(o.Comment) {
		var ret TimeEntryModelComment
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeEntryModel) GetCommentOk() (*TimeEntryModelComment, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TimeEntryModel) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given TimeEntryModelComment and assigns it to the Comment field.
func (o *TimeEntryModel) SetComment(v TimeEntryModelComment) {
	o.Comment = &v
}

// GetSpentOn returns the SpentOn field value if set, zero value otherwise.
func (o *TimeEntryModel) GetSpentOn() string {
	if o == nil || IsNil(o.SpentOn) {
		var ret string
		return ret
	}
	return *o.SpentOn
}

// GetSpentOnOk returns a tuple with the SpentOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeEntryModel) GetSpentOnOk() (*string, bool) {
	if o == nil || IsNil(o.SpentOn) {
		return nil, false
	}
	return o.SpentOn, true
}

// HasSpentOn returns a boolean if a field has been set.
func (o *TimeEntryModel) HasSpentOn() bool {
	if o != nil && !IsNil(o.SpentOn) {
		return true
	}

	return false
}

// SetSpentOn gets a reference to the given string and assigns it to the SpentOn field.
func (o *TimeEntryModel) SetSpentOn(v string) {
	o.SpentOn = &v
}

// GetHours returns the Hours field value if set, zero value otherwise.
func (o *TimeEntryModel) GetHours() string {
	if o == nil || IsNil(o.Hours) {
		var ret string
		return ret
	}
	return *o.Hours
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeEntryModel) GetHoursOk() (*string, bool) {
	if o == nil || IsNil(o.Hours) {
		return nil, false
	}
	return o.Hours, true
}

// HasHours returns a boolean if a field has been set.
func (o *TimeEntryModel) HasHours() bool {
	if o != nil && !IsNil(o.Hours) {
		return true
	}

	return false
}

// SetHours gets a reference to the given string and assigns it to the Hours field.
func (o *TimeEntryModel) SetHours(v string) {
	o.Hours = &v
}

// GetOngoing returns the Ongoing field value if set, zero value otherwise.
func (o *TimeEntryModel) GetOngoing() bool {
	if o == nil || IsNil(o.Ongoing) {
		var ret bool
		return ret
	}
	return *o.Ongoing
}

// GetOngoingOk returns a tuple with the Ongoing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeEntryModel) GetOngoingOk() (*bool, bool) {
	if o == nil || IsNil(o.Ongoing) {
		return nil, false
	}
	return o.Ongoing, true
}

// HasOngoing returns a boolean if a field has been set.
func (o *TimeEntryModel) HasOngoing() bool {
	if o != nil && !IsNil(o.Ongoing) {
		return true
	}

	return false
}

// SetOngoing gets a reference to the given bool and assigns it to the Ongoing field.
func (o *TimeEntryModel) SetOngoing(v bool) {
	o.Ongoing = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TimeEntryModel) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeEntryModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TimeEntryModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TimeEntryModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TimeEntryModel) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeEntryModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TimeEntryModel) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TimeEntryModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TimeEntryModel) GetLinks() TimeEntryModelLinks {
	if o == nil || IsNil(o.Links) {
		var ret TimeEntryModelLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeEntryModel) GetLinksOk() (*TimeEntryModelLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TimeEntryModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TimeEntryModelLinks and assigns it to the Links field.
func (o *TimeEntryModel) SetLinks(v TimeEntryModelLinks) {
	o.Links = &v
}

func (o TimeEntryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeEntryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.SpentOn) {
		toSerialize["spentOn"] = o.SpentOn
	}
	if !IsNil(o.Hours) {
		toSerialize["hours"] = o.Hours
	}
	if !IsNil(o.Ongoing) {
		toSerialize["ongoing"] = o.Ongoing
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

func (o *TimeEntryModel) UnmarshalJSON(data []byte) (err error) {
	varTimeEntryModel := _TimeEntryModel{}

	err = json.Unmarshal(data, &varTimeEntryModel)

	if err != nil {
		return err
	}

	*o = TimeEntryModel(varTimeEntryModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "spentOn")
		delete(additionalProperties, "hours")
		delete(additionalProperties, "ongoing")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimeEntryModel struct {
	value *TimeEntryModel
	isSet bool
}

func (v NullableTimeEntryModel) Get() *TimeEntryModel {
	return v.value
}

func (v *NullableTimeEntryModel) Set(val *TimeEntryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeEntryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeEntryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeEntryModel(val *TimeEntryModel) *NullableTimeEntryModel {
	return &NullableTimeEntryModel{value: val, isSet: true}
}

func (v NullableTimeEntryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeEntryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


