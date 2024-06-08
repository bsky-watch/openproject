# MembershipReadModelEmbeddedPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** | The principal&#39;s unique identifier. | 
**Name** | **string** | The principal&#39;s display name, layout depends on instance settings. | 
**CreatedAt** | Pointer to **time.Time** | Time of creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time of the most recent change to the principal | [optional] 
**Links** | [**GroupModelAllOfLinks**](GroupModelAllOfLinks.md) |  | 
**Avatar** | **string** | URL to user&#39;s avatar | 
**Login** | Pointer to **string** | The user&#39;s login name  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**FirstName** | Pointer to **string** | The user&#39;s first name  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address  # Conditions  - E-Mail address not hidden - User is not a new record - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**Admin** | Pointer to **bool** | Flag indicating whether or not the user is an admin  # Conditions  - &#x60;admin&#x60; | [optional] 
**Status** | Pointer to **string** | The current activation status of the user.  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**Language** | Pointer to **string** | User&#39;s language | ISO 639-1 format  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**IdentityUrl** | Pointer to **string** | User&#39;s identity_url for OmniAuth authentication.  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**Embedded** | [**GroupModelAllOfEmbedded**](GroupModelAllOfEmbedded.md) |  | 

## Methods

### NewMembershipReadModelEmbeddedPrincipal

`func NewMembershipReadModelEmbeddedPrincipal(type_ string, id int64, name string, links GroupModelAllOfLinks, avatar string, embedded GroupModelAllOfEmbedded, ) *MembershipReadModelEmbeddedPrincipal`

NewMembershipReadModelEmbeddedPrincipal instantiates a new MembershipReadModelEmbeddedPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipReadModelEmbeddedPrincipalWithDefaults

`func NewMembershipReadModelEmbeddedPrincipalWithDefaults() *MembershipReadModelEmbeddedPrincipal`

NewMembershipReadModelEmbeddedPrincipalWithDefaults instantiates a new MembershipReadModelEmbeddedPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MembershipReadModelEmbeddedPrincipal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MembershipReadModelEmbeddedPrincipal) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *MembershipReadModelEmbeddedPrincipal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MembershipReadModelEmbeddedPrincipal) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *MembershipReadModelEmbeddedPrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MembershipReadModelEmbeddedPrincipal) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *MembershipReadModelEmbeddedPrincipal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MembershipReadModelEmbeddedPrincipal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MembershipReadModelEmbeddedPrincipal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MembershipReadModelEmbeddedPrincipal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MembershipReadModelEmbeddedPrincipal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MembershipReadModelEmbeddedPrincipal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *MembershipReadModelEmbeddedPrincipal) GetLinks() GroupModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetLinksOk() (*GroupModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MembershipReadModelEmbeddedPrincipal) SetLinks(v GroupModelAllOfLinks)`

SetLinks sets Links field to given value.


### GetAvatar

`func (o *MembershipReadModelEmbeddedPrincipal) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *MembershipReadModelEmbeddedPrincipal) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetLogin

`func (o *MembershipReadModelEmbeddedPrincipal) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *MembershipReadModelEmbeddedPrincipal) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *MembershipReadModelEmbeddedPrincipal) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetFirstName

`func (o *MembershipReadModelEmbeddedPrincipal) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MembershipReadModelEmbeddedPrincipal) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MembershipReadModelEmbeddedPrincipal) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *MembershipReadModelEmbeddedPrincipal) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MembershipReadModelEmbeddedPrincipal) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MembershipReadModelEmbeddedPrincipal) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *MembershipReadModelEmbeddedPrincipal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MembershipReadModelEmbeddedPrincipal) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MembershipReadModelEmbeddedPrincipal) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAdmin

`func (o *MembershipReadModelEmbeddedPrincipal) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *MembershipReadModelEmbeddedPrincipal) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *MembershipReadModelEmbeddedPrincipal) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetStatus

`func (o *MembershipReadModelEmbeddedPrincipal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MembershipReadModelEmbeddedPrincipal) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MembershipReadModelEmbeddedPrincipal) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLanguage

`func (o *MembershipReadModelEmbeddedPrincipal) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MembershipReadModelEmbeddedPrincipal) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MembershipReadModelEmbeddedPrincipal) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetIdentityUrl

`func (o *MembershipReadModelEmbeddedPrincipal) GetIdentityUrl() string`

GetIdentityUrl returns the IdentityUrl field if non-nil, zero value otherwise.

### GetIdentityUrlOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetIdentityUrlOk() (*string, bool)`

GetIdentityUrlOk returns a tuple with the IdentityUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityUrl

`func (o *MembershipReadModelEmbeddedPrincipal) SetIdentityUrl(v string)`

SetIdentityUrl sets IdentityUrl field to given value.

### HasIdentityUrl

`func (o *MembershipReadModelEmbeddedPrincipal) HasIdentityUrl() bool`

HasIdentityUrl returns a boolean if a field has been set.

### GetEmbedded

`func (o *MembershipReadModelEmbeddedPrincipal) GetEmbedded() GroupModelAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *MembershipReadModelEmbeddedPrincipal) GetEmbeddedOk() (*GroupModelAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *MembershipReadModelEmbeddedPrincipal) SetEmbedded(v GroupModelAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


