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

// checks if the WorkPackageModelLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkPackageModelLinks{}

// WorkPackageModelLinks struct for WorkPackageModelLinks
type WorkPackageModelLinks struct {
	AddAttachment *WorkPackageModelLinksAddAttachment `json:"addAttachment,omitempty"`
	AddComment *WorkPackageModelLinksAddComment `json:"addComment,omitempty"`
	AddRelation *WorkPackageModelLinksAddRelation `json:"addRelation,omitempty"`
	AddWatcher *WorkPackageModelLinksAddWatcher `json:"addWatcher,omitempty"`
	CustomActions []WorkPackageModelLinksCustomActionsInner `json:"customActions,omitempty"`
	PreviewMarkup *WorkPackageModelLinksPreviewMarkup `json:"previewMarkup,omitempty"`
	RemoveWatcher *WorkPackageModelLinksRemoveWatcher `json:"removeWatcher,omitempty"`
	Unwatch *WorkPackageModelLinksUnwatch `json:"unwatch,omitempty"`
	Update *WorkPackageModelLinksUpdate `json:"update,omitempty"`
	UpdateImmediately *WorkPackageModelLinksUpdateImmediately `json:"updateImmediately,omitempty"`
	Watch *WorkPackageModelLinksWatch `json:"watch,omitempty"`
	Self *WorkPackageModelLinksSelf `json:"self,omitempty"`
	Schema *WorkPackageModelLinksSchema `json:"schema,omitempty"`
	Ancestors []WorkPackageModelLinksAncestorsInner `json:"ancestors,omitempty"`
	Attachments *WorkPackageModelLinksAttachments `json:"attachments,omitempty"`
	Author *WorkPackageModelLinksAuthor `json:"author,omitempty"`
	Assignee *WorkPackageModelLinksAssignee `json:"assignee,omitempty"`
	AvailableWatchers *WorkPackageModelLinksAvailableWatchers `json:"availableWatchers,omitempty"`
	Budget *WorkPackageModelLinksBudget `json:"budget,omitempty"`
	Category *WorkPackageModelLinksCategory `json:"category,omitempty"`
	Children []WorkPackageModelLinksChildrenInner `json:"children,omitempty"`
	AddFileLink *WorkPackageModelLinksAddFileLink `json:"addFileLink,omitempty"`
	FileLinks *WorkPackageModelLinksFileLinks `json:"fileLinks,omitempty"`
	Parent *WorkPackageModelLinksParent `json:"parent,omitempty"`
	Priority *WorkPackageModelLinksPriority `json:"priority,omitempty"`
	Project *WorkPackageModelLinksProject `json:"project,omitempty"`
	Responsible *WorkPackageModelLinksResponsible `json:"responsible,omitempty"`
	Relations *WorkPackageModelLinksRelations `json:"relations,omitempty"`
	Revisions *WorkPackageModelLinksRevisions `json:"revisions,omitempty"`
	Status *WorkPackageModelLinksStatus `json:"status,omitempty"`
	TimeEntries *WorkPackageModelLinksTimeEntries `json:"timeEntries,omitempty"`
	Type *WorkPackageModelLinksType `json:"type,omitempty"`
	Version *WorkPackageModelLinksVersion `json:"version,omitempty"`
	Watchers *WorkPackageModelLinksWatchers `json:"watchers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkPackageModelLinks WorkPackageModelLinks

// NewWorkPackageModelLinks instantiates a new WorkPackageModelLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkPackageModelLinks() *WorkPackageModelLinks {
	this := WorkPackageModelLinks{}
	return &this
}

// NewWorkPackageModelLinksWithDefaults instantiates a new WorkPackageModelLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkPackageModelLinksWithDefaults() *WorkPackageModelLinks {
	this := WorkPackageModelLinks{}
	return &this
}

// GetAddAttachment returns the AddAttachment field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAddAttachment() WorkPackageModelLinksAddAttachment {
	if o == nil || IsNil(o.AddAttachment) {
		var ret WorkPackageModelLinksAddAttachment
		return ret
	}
	return *o.AddAttachment
}

// GetAddAttachmentOk returns a tuple with the AddAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAddAttachmentOk() (*WorkPackageModelLinksAddAttachment, bool) {
	if o == nil || IsNil(o.AddAttachment) {
		return nil, false
	}
	return o.AddAttachment, true
}

// HasAddAttachment returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAddAttachment() bool {
	if o != nil && !IsNil(o.AddAttachment) {
		return true
	}

	return false
}

// SetAddAttachment gets a reference to the given WorkPackageModelLinksAddAttachment and assigns it to the AddAttachment field.
func (o *WorkPackageModelLinks) SetAddAttachment(v WorkPackageModelLinksAddAttachment) {
	o.AddAttachment = &v
}

// GetAddComment returns the AddComment field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAddComment() WorkPackageModelLinksAddComment {
	if o == nil || IsNil(o.AddComment) {
		var ret WorkPackageModelLinksAddComment
		return ret
	}
	return *o.AddComment
}

// GetAddCommentOk returns a tuple with the AddComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAddCommentOk() (*WorkPackageModelLinksAddComment, bool) {
	if o == nil || IsNil(o.AddComment) {
		return nil, false
	}
	return o.AddComment, true
}

// HasAddComment returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAddComment() bool {
	if o != nil && !IsNil(o.AddComment) {
		return true
	}

	return false
}

// SetAddComment gets a reference to the given WorkPackageModelLinksAddComment and assigns it to the AddComment field.
func (o *WorkPackageModelLinks) SetAddComment(v WorkPackageModelLinksAddComment) {
	o.AddComment = &v
}

// GetAddRelation returns the AddRelation field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAddRelation() WorkPackageModelLinksAddRelation {
	if o == nil || IsNil(o.AddRelation) {
		var ret WorkPackageModelLinksAddRelation
		return ret
	}
	return *o.AddRelation
}

// GetAddRelationOk returns a tuple with the AddRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAddRelationOk() (*WorkPackageModelLinksAddRelation, bool) {
	if o == nil || IsNil(o.AddRelation) {
		return nil, false
	}
	return o.AddRelation, true
}

// HasAddRelation returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAddRelation() bool {
	if o != nil && !IsNil(o.AddRelation) {
		return true
	}

	return false
}

// SetAddRelation gets a reference to the given WorkPackageModelLinksAddRelation and assigns it to the AddRelation field.
func (o *WorkPackageModelLinks) SetAddRelation(v WorkPackageModelLinksAddRelation) {
	o.AddRelation = &v
}

// GetAddWatcher returns the AddWatcher field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAddWatcher() WorkPackageModelLinksAddWatcher {
	if o == nil || IsNil(o.AddWatcher) {
		var ret WorkPackageModelLinksAddWatcher
		return ret
	}
	return *o.AddWatcher
}

// GetAddWatcherOk returns a tuple with the AddWatcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAddWatcherOk() (*WorkPackageModelLinksAddWatcher, bool) {
	if o == nil || IsNil(o.AddWatcher) {
		return nil, false
	}
	return o.AddWatcher, true
}

// HasAddWatcher returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAddWatcher() bool {
	if o != nil && !IsNil(o.AddWatcher) {
		return true
	}

	return false
}

// SetAddWatcher gets a reference to the given WorkPackageModelLinksAddWatcher and assigns it to the AddWatcher field.
func (o *WorkPackageModelLinks) SetAddWatcher(v WorkPackageModelLinksAddWatcher) {
	o.AddWatcher = &v
}

// GetCustomActions returns the CustomActions field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetCustomActions() []WorkPackageModelLinksCustomActionsInner {
	if o == nil || IsNil(o.CustomActions) {
		var ret []WorkPackageModelLinksCustomActionsInner
		return ret
	}
	return o.CustomActions
}

// GetCustomActionsOk returns a tuple with the CustomActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetCustomActionsOk() ([]WorkPackageModelLinksCustomActionsInner, bool) {
	if o == nil || IsNil(o.CustomActions) {
		return nil, false
	}
	return o.CustomActions, true
}

// HasCustomActions returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasCustomActions() bool {
	if o != nil && !IsNil(o.CustomActions) {
		return true
	}

	return false
}

// SetCustomActions gets a reference to the given []WorkPackageModelLinksCustomActionsInner and assigns it to the CustomActions field.
func (o *WorkPackageModelLinks) SetCustomActions(v []WorkPackageModelLinksCustomActionsInner) {
	o.CustomActions = v
}

// GetPreviewMarkup returns the PreviewMarkup field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetPreviewMarkup() WorkPackageModelLinksPreviewMarkup {
	if o == nil || IsNil(o.PreviewMarkup) {
		var ret WorkPackageModelLinksPreviewMarkup
		return ret
	}
	return *o.PreviewMarkup
}

// GetPreviewMarkupOk returns a tuple with the PreviewMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetPreviewMarkupOk() (*WorkPackageModelLinksPreviewMarkup, bool) {
	if o == nil || IsNil(o.PreviewMarkup) {
		return nil, false
	}
	return o.PreviewMarkup, true
}

// HasPreviewMarkup returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasPreviewMarkup() bool {
	if o != nil && !IsNil(o.PreviewMarkup) {
		return true
	}

	return false
}

// SetPreviewMarkup gets a reference to the given WorkPackageModelLinksPreviewMarkup and assigns it to the PreviewMarkup field.
func (o *WorkPackageModelLinks) SetPreviewMarkup(v WorkPackageModelLinksPreviewMarkup) {
	o.PreviewMarkup = &v
}

// GetRemoveWatcher returns the RemoveWatcher field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetRemoveWatcher() WorkPackageModelLinksRemoveWatcher {
	if o == nil || IsNil(o.RemoveWatcher) {
		var ret WorkPackageModelLinksRemoveWatcher
		return ret
	}
	return *o.RemoveWatcher
}

// GetRemoveWatcherOk returns a tuple with the RemoveWatcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetRemoveWatcherOk() (*WorkPackageModelLinksRemoveWatcher, bool) {
	if o == nil || IsNil(o.RemoveWatcher) {
		return nil, false
	}
	return o.RemoveWatcher, true
}

// HasRemoveWatcher returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasRemoveWatcher() bool {
	if o != nil && !IsNil(o.RemoveWatcher) {
		return true
	}

	return false
}

// SetRemoveWatcher gets a reference to the given WorkPackageModelLinksRemoveWatcher and assigns it to the RemoveWatcher field.
func (o *WorkPackageModelLinks) SetRemoveWatcher(v WorkPackageModelLinksRemoveWatcher) {
	o.RemoveWatcher = &v
}

// GetUnwatch returns the Unwatch field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetUnwatch() WorkPackageModelLinksUnwatch {
	if o == nil || IsNil(o.Unwatch) {
		var ret WorkPackageModelLinksUnwatch
		return ret
	}
	return *o.Unwatch
}

// GetUnwatchOk returns a tuple with the Unwatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetUnwatchOk() (*WorkPackageModelLinksUnwatch, bool) {
	if o == nil || IsNil(o.Unwatch) {
		return nil, false
	}
	return o.Unwatch, true
}

// HasUnwatch returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasUnwatch() bool {
	if o != nil && !IsNil(o.Unwatch) {
		return true
	}

	return false
}

// SetUnwatch gets a reference to the given WorkPackageModelLinksUnwatch and assigns it to the Unwatch field.
func (o *WorkPackageModelLinks) SetUnwatch(v WorkPackageModelLinksUnwatch) {
	o.Unwatch = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetUpdate() WorkPackageModelLinksUpdate {
	if o == nil || IsNil(o.Update) {
		var ret WorkPackageModelLinksUpdate
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetUpdateOk() (*WorkPackageModelLinksUpdate, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given WorkPackageModelLinksUpdate and assigns it to the Update field.
func (o *WorkPackageModelLinks) SetUpdate(v WorkPackageModelLinksUpdate) {
	o.Update = &v
}

// GetUpdateImmediately returns the UpdateImmediately field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetUpdateImmediately() WorkPackageModelLinksUpdateImmediately {
	if o == nil || IsNil(o.UpdateImmediately) {
		var ret WorkPackageModelLinksUpdateImmediately
		return ret
	}
	return *o.UpdateImmediately
}

// GetUpdateImmediatelyOk returns a tuple with the UpdateImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetUpdateImmediatelyOk() (*WorkPackageModelLinksUpdateImmediately, bool) {
	if o == nil || IsNil(o.UpdateImmediately) {
		return nil, false
	}
	return o.UpdateImmediately, true
}

// HasUpdateImmediately returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasUpdateImmediately() bool {
	if o != nil && !IsNil(o.UpdateImmediately) {
		return true
	}

	return false
}

// SetUpdateImmediately gets a reference to the given WorkPackageModelLinksUpdateImmediately and assigns it to the UpdateImmediately field.
func (o *WorkPackageModelLinks) SetUpdateImmediately(v WorkPackageModelLinksUpdateImmediately) {
	o.UpdateImmediately = &v
}

// GetWatch returns the Watch field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetWatch() WorkPackageModelLinksWatch {
	if o == nil || IsNil(o.Watch) {
		var ret WorkPackageModelLinksWatch
		return ret
	}
	return *o.Watch
}

// GetWatchOk returns a tuple with the Watch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetWatchOk() (*WorkPackageModelLinksWatch, bool) {
	if o == nil || IsNil(o.Watch) {
		return nil, false
	}
	return o.Watch, true
}

// HasWatch returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasWatch() bool {
	if o != nil && !IsNil(o.Watch) {
		return true
	}

	return false
}

// SetWatch gets a reference to the given WorkPackageModelLinksWatch and assigns it to the Watch field.
func (o *WorkPackageModelLinks) SetWatch(v WorkPackageModelLinksWatch) {
	o.Watch = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetSelf() WorkPackageModelLinksSelf {
	if o == nil || IsNil(o.Self) {
		var ret WorkPackageModelLinksSelf
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetSelfOk() (*WorkPackageModelLinksSelf, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given WorkPackageModelLinksSelf and assigns it to the Self field.
func (o *WorkPackageModelLinks) SetSelf(v WorkPackageModelLinksSelf) {
	o.Self = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetSchema() WorkPackageModelLinksSchema {
	if o == nil || IsNil(o.Schema) {
		var ret WorkPackageModelLinksSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetSchemaOk() (*WorkPackageModelLinksSchema, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given WorkPackageModelLinksSchema and assigns it to the Schema field.
func (o *WorkPackageModelLinks) SetSchema(v WorkPackageModelLinksSchema) {
	o.Schema = &v
}

// GetAncestors returns the Ancestors field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAncestors() []WorkPackageModelLinksAncestorsInner {
	if o == nil || IsNil(o.Ancestors) {
		var ret []WorkPackageModelLinksAncestorsInner
		return ret
	}
	return o.Ancestors
}

// GetAncestorsOk returns a tuple with the Ancestors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAncestorsOk() ([]WorkPackageModelLinksAncestorsInner, bool) {
	if o == nil || IsNil(o.Ancestors) {
		return nil, false
	}
	return o.Ancestors, true
}

// HasAncestors returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAncestors() bool {
	if o != nil && !IsNil(o.Ancestors) {
		return true
	}

	return false
}

// SetAncestors gets a reference to the given []WorkPackageModelLinksAncestorsInner and assigns it to the Ancestors field.
func (o *WorkPackageModelLinks) SetAncestors(v []WorkPackageModelLinksAncestorsInner) {
	o.Ancestors = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAttachments() WorkPackageModelLinksAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret WorkPackageModelLinksAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAttachmentsOk() (*WorkPackageModelLinksAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given WorkPackageModelLinksAttachments and assigns it to the Attachments field.
func (o *WorkPackageModelLinks) SetAttachments(v WorkPackageModelLinksAttachments) {
	o.Attachments = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAuthor() WorkPackageModelLinksAuthor {
	if o == nil || IsNil(o.Author) {
		var ret WorkPackageModelLinksAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAuthorOk() (*WorkPackageModelLinksAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given WorkPackageModelLinksAuthor and assigns it to the Author field.
func (o *WorkPackageModelLinks) SetAuthor(v WorkPackageModelLinksAuthor) {
	o.Author = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAssignee() WorkPackageModelLinksAssignee {
	if o == nil || IsNil(o.Assignee) {
		var ret WorkPackageModelLinksAssignee
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAssigneeOk() (*WorkPackageModelLinksAssignee, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given WorkPackageModelLinksAssignee and assigns it to the Assignee field.
func (o *WorkPackageModelLinks) SetAssignee(v WorkPackageModelLinksAssignee) {
	o.Assignee = &v
}

// GetAvailableWatchers returns the AvailableWatchers field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAvailableWatchers() WorkPackageModelLinksAvailableWatchers {
	if o == nil || IsNil(o.AvailableWatchers) {
		var ret WorkPackageModelLinksAvailableWatchers
		return ret
	}
	return *o.AvailableWatchers
}

// GetAvailableWatchersOk returns a tuple with the AvailableWatchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAvailableWatchersOk() (*WorkPackageModelLinksAvailableWatchers, bool) {
	if o == nil || IsNil(o.AvailableWatchers) {
		return nil, false
	}
	return o.AvailableWatchers, true
}

// HasAvailableWatchers returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAvailableWatchers() bool {
	if o != nil && !IsNil(o.AvailableWatchers) {
		return true
	}

	return false
}

// SetAvailableWatchers gets a reference to the given WorkPackageModelLinksAvailableWatchers and assigns it to the AvailableWatchers field.
func (o *WorkPackageModelLinks) SetAvailableWatchers(v WorkPackageModelLinksAvailableWatchers) {
	o.AvailableWatchers = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetBudget() WorkPackageModelLinksBudget {
	if o == nil || IsNil(o.Budget) {
		var ret WorkPackageModelLinksBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetBudgetOk() (*WorkPackageModelLinksBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given WorkPackageModelLinksBudget and assigns it to the Budget field.
func (o *WorkPackageModelLinks) SetBudget(v WorkPackageModelLinksBudget) {
	o.Budget = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetCategory() WorkPackageModelLinksCategory {
	if o == nil || IsNil(o.Category) {
		var ret WorkPackageModelLinksCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetCategoryOk() (*WorkPackageModelLinksCategory, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given WorkPackageModelLinksCategory and assigns it to the Category field.
func (o *WorkPackageModelLinks) SetCategory(v WorkPackageModelLinksCategory) {
	o.Category = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetChildren() []WorkPackageModelLinksChildrenInner {
	if o == nil || IsNil(o.Children) {
		var ret []WorkPackageModelLinksChildrenInner
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetChildrenOk() ([]WorkPackageModelLinksChildrenInner, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []WorkPackageModelLinksChildrenInner and assigns it to the Children field.
func (o *WorkPackageModelLinks) SetChildren(v []WorkPackageModelLinksChildrenInner) {
	o.Children = v
}

// GetAddFileLink returns the AddFileLink field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetAddFileLink() WorkPackageModelLinksAddFileLink {
	if o == nil || IsNil(o.AddFileLink) {
		var ret WorkPackageModelLinksAddFileLink
		return ret
	}
	return *o.AddFileLink
}

// GetAddFileLinkOk returns a tuple with the AddFileLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetAddFileLinkOk() (*WorkPackageModelLinksAddFileLink, bool) {
	if o == nil || IsNil(o.AddFileLink) {
		return nil, false
	}
	return o.AddFileLink, true
}

// HasAddFileLink returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasAddFileLink() bool {
	if o != nil && !IsNil(o.AddFileLink) {
		return true
	}

	return false
}

// SetAddFileLink gets a reference to the given WorkPackageModelLinksAddFileLink and assigns it to the AddFileLink field.
func (o *WorkPackageModelLinks) SetAddFileLink(v WorkPackageModelLinksAddFileLink) {
	o.AddFileLink = &v
}

// GetFileLinks returns the FileLinks field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetFileLinks() WorkPackageModelLinksFileLinks {
	if o == nil || IsNil(o.FileLinks) {
		var ret WorkPackageModelLinksFileLinks
		return ret
	}
	return *o.FileLinks
}

// GetFileLinksOk returns a tuple with the FileLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetFileLinksOk() (*WorkPackageModelLinksFileLinks, bool) {
	if o == nil || IsNil(o.FileLinks) {
		return nil, false
	}
	return o.FileLinks, true
}

// HasFileLinks returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasFileLinks() bool {
	if o != nil && !IsNil(o.FileLinks) {
		return true
	}

	return false
}

// SetFileLinks gets a reference to the given WorkPackageModelLinksFileLinks and assigns it to the FileLinks field.
func (o *WorkPackageModelLinks) SetFileLinks(v WorkPackageModelLinksFileLinks) {
	o.FileLinks = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetParent() WorkPackageModelLinksParent {
	if o == nil || IsNil(o.Parent) {
		var ret WorkPackageModelLinksParent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetParentOk() (*WorkPackageModelLinksParent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given WorkPackageModelLinksParent and assigns it to the Parent field.
func (o *WorkPackageModelLinks) SetParent(v WorkPackageModelLinksParent) {
	o.Parent = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetPriority() WorkPackageModelLinksPriority {
	if o == nil || IsNil(o.Priority) {
		var ret WorkPackageModelLinksPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetPriorityOk() (*WorkPackageModelLinksPriority, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given WorkPackageModelLinksPriority and assigns it to the Priority field.
func (o *WorkPackageModelLinks) SetPriority(v WorkPackageModelLinksPriority) {
	o.Priority = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetProject() WorkPackageModelLinksProject {
	if o == nil || IsNil(o.Project) {
		var ret WorkPackageModelLinksProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetProjectOk() (*WorkPackageModelLinksProject, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given WorkPackageModelLinksProject and assigns it to the Project field.
func (o *WorkPackageModelLinks) SetProject(v WorkPackageModelLinksProject) {
	o.Project = &v
}

// GetResponsible returns the Responsible field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetResponsible() WorkPackageModelLinksResponsible {
	if o == nil || IsNil(o.Responsible) {
		var ret WorkPackageModelLinksResponsible
		return ret
	}
	return *o.Responsible
}

// GetResponsibleOk returns a tuple with the Responsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetResponsibleOk() (*WorkPackageModelLinksResponsible, bool) {
	if o == nil || IsNil(o.Responsible) {
		return nil, false
	}
	return o.Responsible, true
}

// HasResponsible returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasResponsible() bool {
	if o != nil && !IsNil(o.Responsible) {
		return true
	}

	return false
}

// SetResponsible gets a reference to the given WorkPackageModelLinksResponsible and assigns it to the Responsible field.
func (o *WorkPackageModelLinks) SetResponsible(v WorkPackageModelLinksResponsible) {
	o.Responsible = &v
}

// GetRelations returns the Relations field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetRelations() WorkPackageModelLinksRelations {
	if o == nil || IsNil(o.Relations) {
		var ret WorkPackageModelLinksRelations
		return ret
	}
	return *o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetRelationsOk() (*WorkPackageModelLinksRelations, bool) {
	if o == nil || IsNil(o.Relations) {
		return nil, false
	}
	return o.Relations, true
}

// HasRelations returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasRelations() bool {
	if o != nil && !IsNil(o.Relations) {
		return true
	}

	return false
}

// SetRelations gets a reference to the given WorkPackageModelLinksRelations and assigns it to the Relations field.
func (o *WorkPackageModelLinks) SetRelations(v WorkPackageModelLinksRelations) {
	o.Relations = &v
}

// GetRevisions returns the Revisions field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetRevisions() WorkPackageModelLinksRevisions {
	if o == nil || IsNil(o.Revisions) {
		var ret WorkPackageModelLinksRevisions
		return ret
	}
	return *o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetRevisionsOk() (*WorkPackageModelLinksRevisions, bool) {
	if o == nil || IsNil(o.Revisions) {
		return nil, false
	}
	return o.Revisions, true
}

// HasRevisions returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasRevisions() bool {
	if o != nil && !IsNil(o.Revisions) {
		return true
	}

	return false
}

// SetRevisions gets a reference to the given WorkPackageModelLinksRevisions and assigns it to the Revisions field.
func (o *WorkPackageModelLinks) SetRevisions(v WorkPackageModelLinksRevisions) {
	o.Revisions = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetStatus() WorkPackageModelLinksStatus {
	if o == nil || IsNil(o.Status) {
		var ret WorkPackageModelLinksStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetStatusOk() (*WorkPackageModelLinksStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WorkPackageModelLinksStatus and assigns it to the Status field.
func (o *WorkPackageModelLinks) SetStatus(v WorkPackageModelLinksStatus) {
	o.Status = &v
}

// GetTimeEntries returns the TimeEntries field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetTimeEntries() WorkPackageModelLinksTimeEntries {
	if o == nil || IsNil(o.TimeEntries) {
		var ret WorkPackageModelLinksTimeEntries
		return ret
	}
	return *o.TimeEntries
}

// GetTimeEntriesOk returns a tuple with the TimeEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetTimeEntriesOk() (*WorkPackageModelLinksTimeEntries, bool) {
	if o == nil || IsNil(o.TimeEntries) {
		return nil, false
	}
	return o.TimeEntries, true
}

// HasTimeEntries returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasTimeEntries() bool {
	if o != nil && !IsNil(o.TimeEntries) {
		return true
	}

	return false
}

// SetTimeEntries gets a reference to the given WorkPackageModelLinksTimeEntries and assigns it to the TimeEntries field.
func (o *WorkPackageModelLinks) SetTimeEntries(v WorkPackageModelLinksTimeEntries) {
	o.TimeEntries = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetType() WorkPackageModelLinksType {
	if o == nil || IsNil(o.Type) {
		var ret WorkPackageModelLinksType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetTypeOk() (*WorkPackageModelLinksType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given WorkPackageModelLinksType and assigns it to the Type field.
func (o *WorkPackageModelLinks) SetType(v WorkPackageModelLinksType) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetVersion() WorkPackageModelLinksVersion {
	if o == nil || IsNil(o.Version) {
		var ret WorkPackageModelLinksVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetVersionOk() (*WorkPackageModelLinksVersion, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given WorkPackageModelLinksVersion and assigns it to the Version field.
func (o *WorkPackageModelLinks) SetVersion(v WorkPackageModelLinksVersion) {
	o.Version = &v
}

// GetWatchers returns the Watchers field value if set, zero value otherwise.
func (o *WorkPackageModelLinks) GetWatchers() WorkPackageModelLinksWatchers {
	if o == nil || IsNil(o.Watchers) {
		var ret WorkPackageModelLinksWatchers
		return ret
	}
	return *o.Watchers
}

// GetWatchersOk returns a tuple with the Watchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkPackageModelLinks) GetWatchersOk() (*WorkPackageModelLinksWatchers, bool) {
	if o == nil || IsNil(o.Watchers) {
		return nil, false
	}
	return o.Watchers, true
}

// HasWatchers returns a boolean if a field has been set.
func (o *WorkPackageModelLinks) HasWatchers() bool {
	if o != nil && !IsNil(o.Watchers) {
		return true
	}

	return false
}

// SetWatchers gets a reference to the given WorkPackageModelLinksWatchers and assigns it to the Watchers field.
func (o *WorkPackageModelLinks) SetWatchers(v WorkPackageModelLinksWatchers) {
	o.Watchers = &v
}

func (o WorkPackageModelLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkPackageModelLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddAttachment) {
		toSerialize["addAttachment"] = o.AddAttachment
	}
	if !IsNil(o.AddComment) {
		toSerialize["addComment"] = o.AddComment
	}
	if !IsNil(o.AddRelation) {
		toSerialize["addRelation"] = o.AddRelation
	}
	if !IsNil(o.AddWatcher) {
		toSerialize["addWatcher"] = o.AddWatcher
	}
	if !IsNil(o.CustomActions) {
		toSerialize["customActions"] = o.CustomActions
	}
	if !IsNil(o.PreviewMarkup) {
		toSerialize["previewMarkup"] = o.PreviewMarkup
	}
	if !IsNil(o.RemoveWatcher) {
		toSerialize["removeWatcher"] = o.RemoveWatcher
	}
	if !IsNil(o.Unwatch) {
		toSerialize["unwatch"] = o.Unwatch
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.UpdateImmediately) {
		toSerialize["updateImmediately"] = o.UpdateImmediately
	}
	if !IsNil(o.Watch) {
		toSerialize["watch"] = o.Watch
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Ancestors) {
		toSerialize["ancestors"] = o.Ancestors
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.AvailableWatchers) {
		toSerialize["availableWatchers"] = o.AvailableWatchers
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.AddFileLink) {
		toSerialize["addFileLink"] = o.AddFileLink
	}
	if !IsNil(o.FileLinks) {
		toSerialize["fileLinks"] = o.FileLinks
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
	if !IsNil(o.Relations) {
		toSerialize["relations"] = o.Relations
	}
	if !IsNil(o.Revisions) {
		toSerialize["revisions"] = o.Revisions
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TimeEntries) {
		toSerialize["timeEntries"] = o.TimeEntries
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Watchers) {
		toSerialize["watchers"] = o.Watchers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkPackageModelLinks) UnmarshalJSON(data []byte) (err error) {
	varWorkPackageModelLinks := _WorkPackageModelLinks{}

	err = json.Unmarshal(data, &varWorkPackageModelLinks)

	if err != nil {
		return err
	}

	*o = WorkPackageModelLinks(varWorkPackageModelLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "addAttachment")
		delete(additionalProperties, "addComment")
		delete(additionalProperties, "addRelation")
		delete(additionalProperties, "addWatcher")
		delete(additionalProperties, "customActions")
		delete(additionalProperties, "previewMarkup")
		delete(additionalProperties, "removeWatcher")
		delete(additionalProperties, "unwatch")
		delete(additionalProperties, "update")
		delete(additionalProperties, "updateImmediately")
		delete(additionalProperties, "watch")
		delete(additionalProperties, "self")
		delete(additionalProperties, "schema")
		delete(additionalProperties, "ancestors")
		delete(additionalProperties, "attachments")
		delete(additionalProperties, "author")
		delete(additionalProperties, "assignee")
		delete(additionalProperties, "availableWatchers")
		delete(additionalProperties, "budget")
		delete(additionalProperties, "category")
		delete(additionalProperties, "children")
		delete(additionalProperties, "addFileLink")
		delete(additionalProperties, "fileLinks")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "project")
		delete(additionalProperties, "responsible")
		delete(additionalProperties, "relations")
		delete(additionalProperties, "revisions")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeEntries")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		delete(additionalProperties, "watchers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkPackageModelLinks struct {
	value *WorkPackageModelLinks
	isSet bool
}

func (v NullableWorkPackageModelLinks) Get() *WorkPackageModelLinks {
	return v.value
}

func (v *NullableWorkPackageModelLinks) Set(val *WorkPackageModelLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkPackageModelLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkPackageModelLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkPackageModelLinks(val *WorkPackageModelLinks) *NullableWorkPackageModelLinks {
	return &NullableWorkPackageModelLinks{value: val, isSet: true}
}

func (v NullableWorkPackageModelLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkPackageModelLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


