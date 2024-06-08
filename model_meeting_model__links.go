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

// checks if the MeetingModelLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeetingModelLinks{}

// MeetingModelLinks struct for MeetingModelLinks
type MeetingModelLinks struct {
	Self *MeetingModelLinksSelf `json:"self,omitempty"`
	Author *MeetingModelLinksAuthor `json:"author,omitempty"`
	Project *MeetingModelLinksProject `json:"project,omitempty"`
	Attachments *MeetingModelLinksAttachments `json:"attachments,omitempty"`
	AddAttachment *MeetingModelLinksAddAttachment `json:"addAttachment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MeetingModelLinks MeetingModelLinks

// NewMeetingModelLinks instantiates a new MeetingModelLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeetingModelLinks() *MeetingModelLinks {
	this := MeetingModelLinks{}
	return &this
}

// NewMeetingModelLinksWithDefaults instantiates a new MeetingModelLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeetingModelLinksWithDefaults() *MeetingModelLinks {
	this := MeetingModelLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *MeetingModelLinks) GetSelf() MeetingModelLinksSelf {
	if o == nil || IsNil(o.Self) {
		var ret MeetingModelLinksSelf
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingModelLinks) GetSelfOk() (*MeetingModelLinksSelf, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *MeetingModelLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given MeetingModelLinksSelf and assigns it to the Self field.
func (o *MeetingModelLinks) SetSelf(v MeetingModelLinksSelf) {
	o.Self = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *MeetingModelLinks) GetAuthor() MeetingModelLinksAuthor {
	if o == nil || IsNil(o.Author) {
		var ret MeetingModelLinksAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingModelLinks) GetAuthorOk() (*MeetingModelLinksAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *MeetingModelLinks) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given MeetingModelLinksAuthor and assigns it to the Author field.
func (o *MeetingModelLinks) SetAuthor(v MeetingModelLinksAuthor) {
	o.Author = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *MeetingModelLinks) GetProject() MeetingModelLinksProject {
	if o == nil || IsNil(o.Project) {
		var ret MeetingModelLinksProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingModelLinks) GetProjectOk() (*MeetingModelLinksProject, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *MeetingModelLinks) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given MeetingModelLinksProject and assigns it to the Project field.
func (o *MeetingModelLinks) SetProject(v MeetingModelLinksProject) {
	o.Project = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *MeetingModelLinks) GetAttachments() MeetingModelLinksAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret MeetingModelLinksAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingModelLinks) GetAttachmentsOk() (*MeetingModelLinksAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *MeetingModelLinks) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given MeetingModelLinksAttachments and assigns it to the Attachments field.
func (o *MeetingModelLinks) SetAttachments(v MeetingModelLinksAttachments) {
	o.Attachments = &v
}

// GetAddAttachment returns the AddAttachment field value if set, zero value otherwise.
func (o *MeetingModelLinks) GetAddAttachment() MeetingModelLinksAddAttachment {
	if o == nil || IsNil(o.AddAttachment) {
		var ret MeetingModelLinksAddAttachment
		return ret
	}
	return *o.AddAttachment
}

// GetAddAttachmentOk returns a tuple with the AddAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingModelLinks) GetAddAttachmentOk() (*MeetingModelLinksAddAttachment, bool) {
	if o == nil || IsNil(o.AddAttachment) {
		return nil, false
	}
	return o.AddAttachment, true
}

// HasAddAttachment returns a boolean if a field has been set.
func (o *MeetingModelLinks) HasAddAttachment() bool {
	if o != nil && !IsNil(o.AddAttachment) {
		return true
	}

	return false
}

// SetAddAttachment gets a reference to the given MeetingModelLinksAddAttachment and assigns it to the AddAttachment field.
func (o *MeetingModelLinks) SetAddAttachment(v MeetingModelLinksAddAttachment) {
	o.AddAttachment = &v
}

func (o MeetingModelLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeetingModelLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.AddAttachment) {
		toSerialize["addAttachment"] = o.AddAttachment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MeetingModelLinks) UnmarshalJSON(data []byte) (err error) {
	varMeetingModelLinks := _MeetingModelLinks{}

	err = json.Unmarshal(data, &varMeetingModelLinks)

	if err != nil {
		return err
	}

	*o = MeetingModelLinks(varMeetingModelLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "author")
		delete(additionalProperties, "project")
		delete(additionalProperties, "attachments")
		delete(additionalProperties, "addAttachment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMeetingModelLinks struct {
	value *MeetingModelLinks
	isSet bool
}

func (v NullableMeetingModelLinks) Get() *MeetingModelLinks {
	return v.value
}

func (v *NullableMeetingModelLinks) Set(val *MeetingModelLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableMeetingModelLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableMeetingModelLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeetingModelLinks(val *MeetingModelLinks) *NullableMeetingModelLinks {
	return &NullableMeetingModelLinks{value: val, isSet: true}
}

func (v NullableMeetingModelLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeetingModelLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


