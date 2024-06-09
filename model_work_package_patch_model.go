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

// checks if the WorkPackagePatchModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkPackagePatchModel{}

// WorkPackagePatchModel struct for WorkPackagePatchModel
type WorkPackagePatchModel struct {
	// The version of the item as used for optimistic locking
	LockVersion int `json:"lockVersion"`
	// Work package subject
	Subject *string `json:"subject,omitempty"`
	Description *WorkPackageModelDescription `json:"description,omitempty"`
	// If false (default) schedule automatically.
	ScheduleManually *bool `json:"scheduleManually,omitempty"`
	// Scheduled beginning of a work package
	StartDate *string `json:"startDate,omitempty"`
	// Scheduled end of a work package
	DueDate *string `json:"dueDate,omitempty"`
	// Date on which a milestone is achieved
	Date *string `json:"date,omitempty"`
	// Time a work package likely needs to be completed excluding its descendants
	EstimatedTime *string `json:"estimatedTime,omitempty"`
	// **(NOT IMPLEMENTED)** When scheduling, whether or not to ignore the non working days being defined. A work package with the flag set to true will be allowed to be scheduled to a non working day.
	IgnoreNonWorkingDays *bool `json:"ignoreNonWorkingDays,omitempty"`
	// The time booked for this work package by users working on it  # Conditions  **Permission** view time entries
	SpentTime *string `json:"spentTime,omitempty"`
	// Amount of total completion for a work package
	PercentageDone *int `json:"percentageDone,omitempty"`
	// Time of creation
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Time of the most recent change to the work package
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Links *WorkPackagePatchModelLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkPackagePatchModel WorkPackagePatchModel

// NewWorkPackagePatchModel instantiates a new WorkPackagePatchModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkPackagePatchModel(lockVersion int) *WorkPackagePatchModel {
	this := WorkPackagePatchModel{}
	this.LockVersion = lockVersion
	return &this
}

// NewWorkPackagePatchModelWithDefaults instantiates a new WorkPackagePatchModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkPackagePatchModelWithDefaults() *WorkPackagePatchModel {
	this := WorkPackagePatchModel{}
	return &this
}

// GetLockVersion returns the LockVersion field value
func (o *WorkPackagePatchModel) GetLockVersion() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.LockVersion
}

// GetLockVersionOk returns a tuple with the LockVersion field value
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetLockVersionOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockVersion, true
}

// SetLockVersion sets field value
func (o *WorkPackagePatchModel) SetLockVersion(v int) {
	o.LockVersion = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *WorkPackagePatchModel) SetSubject(v string) {
	o.Subject = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetDescription() WorkPackageModelDescription {
	if o == nil || IsNil(o.Description) {
		var ret WorkPackageModelDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetDescriptionOk() (*WorkPackageModelDescription, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given WorkPackageModelDescription and assigns it to the Description field.
func (o *WorkPackagePatchModel) SetDescription(v WorkPackageModelDescription) {
	o.Description = &v
}

// GetScheduleManually returns the ScheduleManually field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetScheduleManually() bool {
	if o == nil || IsNil(o.ScheduleManually) {
		var ret bool
		return ret
	}
	return *o.ScheduleManually
}

// GetScheduleManuallyOk returns a tuple with the ScheduleManually field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetScheduleManuallyOk() (*bool, bool) {
	if o == nil || IsNil(o.ScheduleManually) {
		return nil, false
	}
	return o.ScheduleManually, true
}

// HasScheduleManually returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasScheduleManually() bool {
	if o != nil && !IsNil(o.ScheduleManually) {
		return true
	}

	return false
}

// SetScheduleManually gets a reference to the given bool and assigns it to the ScheduleManually field.
func (o *WorkPackagePatchModel) SetScheduleManually(v bool) {
	o.ScheduleManually = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *WorkPackagePatchModel) SetStartDate(v string) {
	o.StartDate = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *WorkPackagePatchModel) SetDueDate(v string) {
	o.DueDate = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *WorkPackagePatchModel) SetDate(v string) {
	o.Date = &v
}

// GetEstimatedTime returns the EstimatedTime field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetEstimatedTime() string {
	if o == nil || IsNil(o.EstimatedTime) {
		var ret string
		return ret
	}
	return *o.EstimatedTime
}

// GetEstimatedTimeOk returns a tuple with the EstimatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetEstimatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedTime) {
		return nil, false
	}
	return o.EstimatedTime, true
}

// HasEstimatedTime returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasEstimatedTime() bool {
	if o != nil && !IsNil(o.EstimatedTime) {
		return true
	}

	return false
}

// SetEstimatedTime gets a reference to the given string and assigns it to the EstimatedTime field.
func (o *WorkPackagePatchModel) SetEstimatedTime(v string) {
	o.EstimatedTime = &v
}

// GetIgnoreNonWorkingDays returns the IgnoreNonWorkingDays field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetIgnoreNonWorkingDays() bool {
	if o == nil || IsNil(o.IgnoreNonWorkingDays) {
		var ret bool
		return ret
	}
	return *o.IgnoreNonWorkingDays
}

// GetIgnoreNonWorkingDaysOk returns a tuple with the IgnoreNonWorkingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetIgnoreNonWorkingDaysOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreNonWorkingDays) {
		return nil, false
	}
	return o.IgnoreNonWorkingDays, true
}

// HasIgnoreNonWorkingDays returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasIgnoreNonWorkingDays() bool {
	if o != nil && !IsNil(o.IgnoreNonWorkingDays) {
		return true
	}

	return false
}

// SetIgnoreNonWorkingDays gets a reference to the given bool and assigns it to the IgnoreNonWorkingDays field.
func (o *WorkPackagePatchModel) SetIgnoreNonWorkingDays(v bool) {
	o.IgnoreNonWorkingDays = &v
}

// GetSpentTime returns the SpentTime field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetSpentTime() string {
	if o == nil || IsNil(o.SpentTime) {
		var ret string
		return ret
	}
	return *o.SpentTime
}

// GetSpentTimeOk returns a tuple with the SpentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetSpentTimeOk() (*string, bool) {
	if o == nil || IsNil(o.SpentTime) {
		return nil, false
	}
	return o.SpentTime, true
}

// HasSpentTime returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasSpentTime() bool {
	if o != nil && !IsNil(o.SpentTime) {
		return true
	}

	return false
}

// SetSpentTime gets a reference to the given string and assigns it to the SpentTime field.
func (o *WorkPackagePatchModel) SetSpentTime(v string) {
	o.SpentTime = &v
}

// GetPercentageDone returns the PercentageDone field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetPercentageDone() int {
	if o == nil || IsNil(o.PercentageDone) {
		var ret int
		return ret
	}
	return *o.PercentageDone
}

// GetPercentageDoneOk returns a tuple with the PercentageDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetPercentageDoneOk() (*int, bool) {
	if o == nil || IsNil(o.PercentageDone) {
		return nil, false
	}
	return o.PercentageDone, true
}

// HasPercentageDone returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasPercentageDone() bool {
	if o != nil && !IsNil(o.PercentageDone) {
		return true
	}

	return false
}

// SetPercentageDone gets a reference to the given int and assigns it to the PercentageDone field.
func (o *WorkPackagePatchModel) SetPercentageDone(v int) {
	o.PercentageDone = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WorkPackagePatchModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WorkPackagePatchModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WorkPackagePatchModel) GetLinks() WorkPackagePatchModelLinks {
	if o == nil || IsNil(o.Links) {
		var ret WorkPackagePatchModelLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackagePatchModel) GetLinksOk() (*WorkPackagePatchModelLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WorkPackagePatchModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WorkPackagePatchModelLinks and assigns it to the Links field.
func (o *WorkPackagePatchModel) SetLinks(v WorkPackagePatchModelLinks) {
	o.Links = &v
}

func (o WorkPackagePatchModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkPackagePatchModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lockVersion"] = o.LockVersion
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ScheduleManually) {
		toSerialize["scheduleManually"] = o.ScheduleManually
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.EstimatedTime) {
		toSerialize["estimatedTime"] = o.EstimatedTime
	}
	if !IsNil(o.IgnoreNonWorkingDays) {
		toSerialize["ignoreNonWorkingDays"] = o.IgnoreNonWorkingDays
	}
	if !IsNil(o.SpentTime) {
		toSerialize["spentTime"] = o.SpentTime
	}
	if !IsNil(o.PercentageDone) {
		toSerialize["percentageDone"] = o.PercentageDone
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

func (o *WorkPackagePatchModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lockVersion",
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

	varWorkPackagePatchModel := _WorkPackagePatchModel{}

	err = json.Unmarshal(data, &varWorkPackagePatchModel)

	if err != nil {
		return err
	}

	*o = WorkPackagePatchModel(varWorkPackagePatchModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lockVersion")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "description")
		delete(additionalProperties, "scheduleManually")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "dueDate")
		delete(additionalProperties, "date")
		delete(additionalProperties, "estimatedTime")
		delete(additionalProperties, "ignoreNonWorkingDays")
		delete(additionalProperties, "spentTime")
		delete(additionalProperties, "percentageDone")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkPackagePatchModel struct {
	value *WorkPackagePatchModel
	isSet bool
}

func (v NullableWorkPackagePatchModel) Get() *WorkPackagePatchModel {
	return v.value
}

func (v *NullableWorkPackagePatchModel) Set(val *WorkPackagePatchModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkPackagePatchModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkPackagePatchModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkPackagePatchModel(val *WorkPackagePatchModel) *NullableWorkPackagePatchModel {
	return &NullableWorkPackagePatchModel{value: val, isSet: true}
}

func (v NullableWorkPackagePatchModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkPackagePatchModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


