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

// checks if the WorkPackageModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkPackageModel{}

// WorkPackageModel struct for WorkPackageModel
type WorkPackageModel struct {
	// Work package id
	Id *int `json:"id,omitempty"`
	// The version of the item as used for optimistic locking
	LockVersion *int `json:"lockVersion,omitempty"`
	// Work package subject
	Subject *string `json:"subject,omitempty"`
	Type *string `json:"_type,omitempty"`
	Description *WorkPackageModelDescription `json:"description,omitempty"`
	// If false (default) schedule automatically.
	ScheduleManually *bool `json:"scheduleManually,omitempty"`
	// If true, the work package is in a readonly status so with the exception of the status, no other property can be altered.
	Readonly *bool `json:"readonly,omitempty"`
	// Scheduled beginning of a work package
	StartDate *string `json:"startDate,omitempty"`
	// Scheduled end of a work package
	DueDate *string `json:"dueDate,omitempty"`
	// Date on which a milestone is achieved
	Date *string `json:"date,omitempty"`
	// Similar to start date but is not set by a client but rather deduced by the work packages' descendants. If manual scheduleManually is active, the two dates can deviate.
	DerivedStartDate *string `json:"derivedStartDate,omitempty"`
	// Similar to due date but is not set by a client but rather deduced by the work packages' descendants. If manual scheduleManually is active, the two dates can deviate.
	DerivedDueDate *string `json:"derivedDueDate,omitempty"`
	// **(NOT IMPLEMENTED)** The amount of time in hours the work package needs to be completed. Not available for milestone type of work packages.
	Duration *string `json:"duration,omitempty"`
	// Time a work package likely needs to be completed excluding its descendants
	EstimatedTime *string `json:"estimatedTime,omitempty"`
	// Time a work package likely needs to be completed including its descendants
	DerivedEstimatedTime *string `json:"derivedEstimatedTime,omitempty"`
	// **(NOT IMPLEMENTED)** When scheduling, whether or not to ignore the non working days being defined. A work package with the flag set to true will be allowed to be scheduled to a non working day.
	IgnoreNonWorkingDays *bool `json:"ignoreNonWorkingDays,omitempty"`
	// The time booked for this work package by users working on it  # Conditions  **Permission** view time entries
	SpentTime *string `json:"spentTime,omitempty"`
	// Amount of total completion for a work package
	PercentageDone *int `json:"percentageDone,omitempty"`
	// Amount of total completion for a work package derived from itself and its descendant work packages
	DerivedPercentageDone *int `json:"derivedPercentageDone,omitempty"`
	// Time of creation. Can be writable by admins with the `apiv3_write_readonly_attributes` setting enabled.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Time of the most recent change to the work package.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Links *WorkPackageModelLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkPackageModel WorkPackageModel

// NewWorkPackageModel instantiates a new WorkPackageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkPackageModel() *WorkPackageModel {
	this := WorkPackageModel{}
	return &this
}

// NewWorkPackageModelWithDefaults instantiates a new WorkPackageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkPackageModelWithDefaults() *WorkPackageModel {
	this := WorkPackageModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkPackageModel) GetId() int {
	if o == nil || IsNil(o.Id) {
		var ret int
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetIdOk() (*int, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkPackageModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int and assigns it to the Id field.
func (o *WorkPackageModel) SetId(v int) {
	o.Id = &v
}

// GetLockVersion returns the LockVersion field value if set, zero value otherwise.
func (o *WorkPackageModel) GetLockVersion() int {
	if o == nil || IsNil(o.LockVersion) {
		var ret int
		return ret
	}
	return *o.LockVersion
}

// GetLockVersionOk returns a tuple with the LockVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetLockVersionOk() (*int, bool) {
	if o == nil || IsNil(o.LockVersion) {
		return nil, false
	}
	return o.LockVersion, true
}

// HasLockVersion returns a boolean if a field has been set.
func (o *WorkPackageModel) HasLockVersion() bool {
	if o != nil && !IsNil(o.LockVersion) {
		return true
	}

	return false
}

// SetLockVersion gets a reference to the given int and assigns it to the LockVersion field.
func (o *WorkPackageModel) SetLockVersion(v int) {
	o.LockVersion = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *WorkPackageModel) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *WorkPackageModel) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *WorkPackageModel) SetSubject(v string) {
	o.Subject = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkPackageModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkPackageModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkPackageModel) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkPackageModel) GetDescription() WorkPackageModelDescription {
	if o == nil || IsNil(o.Description) {
		var ret WorkPackageModelDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetDescriptionOk() (*WorkPackageModelDescription, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkPackageModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given WorkPackageModelDescription and assigns it to the Description field.
func (o *WorkPackageModel) SetDescription(v WorkPackageModelDescription) {
	o.Description = &v
}

// GetScheduleManually returns the ScheduleManually field value if set, zero value otherwise.
func (o *WorkPackageModel) GetScheduleManually() bool {
	if o == nil || IsNil(o.ScheduleManually) {
		var ret bool
		return ret
	}
	return *o.ScheduleManually
}

// GetScheduleManuallyOk returns a tuple with the ScheduleManually field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetScheduleManuallyOk() (*bool, bool) {
	if o == nil || IsNil(o.ScheduleManually) {
		return nil, false
	}
	return o.ScheduleManually, true
}

// HasScheduleManually returns a boolean if a field has been set.
func (o *WorkPackageModel) HasScheduleManually() bool {
	if o != nil && !IsNil(o.ScheduleManually) {
		return true
	}

	return false
}

// SetScheduleManually gets a reference to the given bool and assigns it to the ScheduleManually field.
func (o *WorkPackageModel) SetScheduleManually(v bool) {
	o.ScheduleManually = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *WorkPackageModel) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *WorkPackageModel) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *WorkPackageModel) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *WorkPackageModel) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *WorkPackageModel) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *WorkPackageModel) SetStartDate(v string) {
	o.StartDate = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *WorkPackageModel) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *WorkPackageModel) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *WorkPackageModel) SetDueDate(v string) {
	o.DueDate = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *WorkPackageModel) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *WorkPackageModel) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *WorkPackageModel) SetDate(v string) {
	o.Date = &v
}

// GetDerivedStartDate returns the DerivedStartDate field value if set, zero value otherwise.
func (o *WorkPackageModel) GetDerivedStartDate() string {
	if o == nil || IsNil(o.DerivedStartDate) {
		var ret string
		return ret
	}
	return *o.DerivedStartDate
}

// GetDerivedStartDateOk returns a tuple with the DerivedStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetDerivedStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.DerivedStartDate) {
		return nil, false
	}
	return o.DerivedStartDate, true
}

// HasDerivedStartDate returns a boolean if a field has been set.
func (o *WorkPackageModel) HasDerivedStartDate() bool {
	if o != nil && !IsNil(o.DerivedStartDate) {
		return true
	}

	return false
}

// SetDerivedStartDate gets a reference to the given string and assigns it to the DerivedStartDate field.
func (o *WorkPackageModel) SetDerivedStartDate(v string) {
	o.DerivedStartDate = &v
}

// GetDerivedDueDate returns the DerivedDueDate field value if set, zero value otherwise.
func (o *WorkPackageModel) GetDerivedDueDate() string {
	if o == nil || IsNil(o.DerivedDueDate) {
		var ret string
		return ret
	}
	return *o.DerivedDueDate
}

// GetDerivedDueDateOk returns a tuple with the DerivedDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetDerivedDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DerivedDueDate) {
		return nil, false
	}
	return o.DerivedDueDate, true
}

// HasDerivedDueDate returns a boolean if a field has been set.
func (o *WorkPackageModel) HasDerivedDueDate() bool {
	if o != nil && !IsNil(o.DerivedDueDate) {
		return true
	}

	return false
}

// SetDerivedDueDate gets a reference to the given string and assigns it to the DerivedDueDate field.
func (o *WorkPackageModel) SetDerivedDueDate(v string) {
	o.DerivedDueDate = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *WorkPackageModel) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *WorkPackageModel) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *WorkPackageModel) SetDuration(v string) {
	o.Duration = &v
}

// GetEstimatedTime returns the EstimatedTime field value if set, zero value otherwise.
func (o *WorkPackageModel) GetEstimatedTime() string {
	if o == nil || IsNil(o.EstimatedTime) {
		var ret string
		return ret
	}
	return *o.EstimatedTime
}

// GetEstimatedTimeOk returns a tuple with the EstimatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetEstimatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedTime) {
		return nil, false
	}
	return o.EstimatedTime, true
}

// HasEstimatedTime returns a boolean if a field has been set.
func (o *WorkPackageModel) HasEstimatedTime() bool {
	if o != nil && !IsNil(o.EstimatedTime) {
		return true
	}

	return false
}

// SetEstimatedTime gets a reference to the given string and assigns it to the EstimatedTime field.
func (o *WorkPackageModel) SetEstimatedTime(v string) {
	o.EstimatedTime = &v
}

// GetDerivedEstimatedTime returns the DerivedEstimatedTime field value if set, zero value otherwise.
func (o *WorkPackageModel) GetDerivedEstimatedTime() string {
	if o == nil || IsNil(o.DerivedEstimatedTime) {
		var ret string
		return ret
	}
	return *o.DerivedEstimatedTime
}

// GetDerivedEstimatedTimeOk returns a tuple with the DerivedEstimatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetDerivedEstimatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DerivedEstimatedTime) {
		return nil, false
	}
	return o.DerivedEstimatedTime, true
}

// HasDerivedEstimatedTime returns a boolean if a field has been set.
func (o *WorkPackageModel) HasDerivedEstimatedTime() bool {
	if o != nil && !IsNil(o.DerivedEstimatedTime) {
		return true
	}

	return false
}

// SetDerivedEstimatedTime gets a reference to the given string and assigns it to the DerivedEstimatedTime field.
func (o *WorkPackageModel) SetDerivedEstimatedTime(v string) {
	o.DerivedEstimatedTime = &v
}

// GetIgnoreNonWorkingDays returns the IgnoreNonWorkingDays field value if set, zero value otherwise.
func (o *WorkPackageModel) GetIgnoreNonWorkingDays() bool {
	if o == nil || IsNil(o.IgnoreNonWorkingDays) {
		var ret bool
		return ret
	}
	return *o.IgnoreNonWorkingDays
}

// GetIgnoreNonWorkingDaysOk returns a tuple with the IgnoreNonWorkingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetIgnoreNonWorkingDaysOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreNonWorkingDays) {
		return nil, false
	}
	return o.IgnoreNonWorkingDays, true
}

// HasIgnoreNonWorkingDays returns a boolean if a field has been set.
func (o *WorkPackageModel) HasIgnoreNonWorkingDays() bool {
	if o != nil && !IsNil(o.IgnoreNonWorkingDays) {
		return true
	}

	return false
}

// SetIgnoreNonWorkingDays gets a reference to the given bool and assigns it to the IgnoreNonWorkingDays field.
func (o *WorkPackageModel) SetIgnoreNonWorkingDays(v bool) {
	o.IgnoreNonWorkingDays = &v
}

// GetSpentTime returns the SpentTime field value if set, zero value otherwise.
func (o *WorkPackageModel) GetSpentTime() string {
	if o == nil || IsNil(o.SpentTime) {
		var ret string
		return ret
	}
	return *o.SpentTime
}

// GetSpentTimeOk returns a tuple with the SpentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetSpentTimeOk() (*string, bool) {
	if o == nil || IsNil(o.SpentTime) {
		return nil, false
	}
	return o.SpentTime, true
}

// HasSpentTime returns a boolean if a field has been set.
func (o *WorkPackageModel) HasSpentTime() bool {
	if o != nil && !IsNil(o.SpentTime) {
		return true
	}

	return false
}

// SetSpentTime gets a reference to the given string and assigns it to the SpentTime field.
func (o *WorkPackageModel) SetSpentTime(v string) {
	o.SpentTime = &v
}

// GetPercentageDone returns the PercentageDone field value if set, zero value otherwise.
func (o *WorkPackageModel) GetPercentageDone() int {
	if o == nil || IsNil(o.PercentageDone) {
		var ret int
		return ret
	}
	return *o.PercentageDone
}

// GetPercentageDoneOk returns a tuple with the PercentageDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetPercentageDoneOk() (*int, bool) {
	if o == nil || IsNil(o.PercentageDone) {
		return nil, false
	}
	return o.PercentageDone, true
}

// HasPercentageDone returns a boolean if a field has been set.
func (o *WorkPackageModel) HasPercentageDone() bool {
	if o != nil && !IsNil(o.PercentageDone) {
		return true
	}

	return false
}

// SetPercentageDone gets a reference to the given int and assigns it to the PercentageDone field.
func (o *WorkPackageModel) SetPercentageDone(v int) {
	o.PercentageDone = &v
}

// GetDerivedPercentageDone returns the DerivedPercentageDone field value if set, zero value otherwise.
func (o *WorkPackageModel) GetDerivedPercentageDone() int {
	if o == nil || IsNil(o.DerivedPercentageDone) {
		var ret int
		return ret
	}
	return *o.DerivedPercentageDone
}

// GetDerivedPercentageDoneOk returns a tuple with the DerivedPercentageDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetDerivedPercentageDoneOk() (*int, bool) {
	if o == nil || IsNil(o.DerivedPercentageDone) {
		return nil, false
	}
	return o.DerivedPercentageDone, true
}

// HasDerivedPercentageDone returns a boolean if a field has been set.
func (o *WorkPackageModel) HasDerivedPercentageDone() bool {
	if o != nil && !IsNil(o.DerivedPercentageDone) {
		return true
	}

	return false
}

// SetDerivedPercentageDone gets a reference to the given int and assigns it to the DerivedPercentageDone field.
func (o *WorkPackageModel) SetDerivedPercentageDone(v int) {
	o.DerivedPercentageDone = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WorkPackageModel) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WorkPackageModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WorkPackageModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkPackageModel) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkPackageModel) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WorkPackageModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WorkPackageModel) GetLinks() WorkPackageModelLinks {
	if o == nil || IsNil(o.Links) {
		var ret WorkPackageModelLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModel) GetLinksOk() (*WorkPackageModelLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WorkPackageModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WorkPackageModelLinks and assigns it to the Links field.
func (o *WorkPackageModel) SetLinks(v WorkPackageModelLinks) {
	o.Links = &v
}

func (o WorkPackageModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkPackageModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LockVersion) {
		toSerialize["lockVersion"] = o.LockVersion
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Type) {
		toSerialize["_type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ScheduleManually) {
		toSerialize["scheduleManually"] = o.ScheduleManually
	}
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
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
	if !IsNil(o.DerivedStartDate) {
		toSerialize["derivedStartDate"] = o.DerivedStartDate
	}
	if !IsNil(o.DerivedDueDate) {
		toSerialize["derivedDueDate"] = o.DerivedDueDate
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.EstimatedTime) {
		toSerialize["estimatedTime"] = o.EstimatedTime
	}
	if !IsNil(o.DerivedEstimatedTime) {
		toSerialize["derivedEstimatedTime"] = o.DerivedEstimatedTime
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
	if !IsNil(o.DerivedPercentageDone) {
		toSerialize["derivedPercentageDone"] = o.DerivedPercentageDone
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

func (o *WorkPackageModel) UnmarshalJSON(data []byte) (err error) {
	varWorkPackageModel := _WorkPackageModel{}

	err = json.Unmarshal(data, &varWorkPackageModel)

	if err != nil {
		return err
	}

	*o = WorkPackageModel(varWorkPackageModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "lockVersion")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "scheduleManually")
		delete(additionalProperties, "readonly")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "dueDate")
		delete(additionalProperties, "date")
		delete(additionalProperties, "derivedStartDate")
		delete(additionalProperties, "derivedDueDate")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "estimatedTime")
		delete(additionalProperties, "derivedEstimatedTime")
		delete(additionalProperties, "ignoreNonWorkingDays")
		delete(additionalProperties, "spentTime")
		delete(additionalProperties, "percentageDone")
		delete(additionalProperties, "derivedPercentageDone")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkPackageModel struct {
	value *WorkPackageModel
	isSet bool
}

func (v NullableWorkPackageModel) Get() *WorkPackageModel {
	return v.value
}

func (v *NullableWorkPackageModel) Set(val *WorkPackageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkPackageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkPackageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkPackageModel(val *WorkPackageModel) *NullableWorkPackageModel {
	return &NullableWorkPackageModel{value: val, isSet: true}
}

func (v NullableWorkPackageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkPackageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


