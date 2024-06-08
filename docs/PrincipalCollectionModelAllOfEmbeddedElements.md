# PrincipalCollectionModelAllOfEmbeddedElements

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
**Status** | Pointer to **string** | The current activation status of the placeholder user.  # Conditions  - User has &#x60;manage_placeholder_user&#x60; permission globally | [optional] 
**Language** | Pointer to **string** | User&#39;s language | ISO 639-1 format  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**IdentityUrl** | Pointer to **string** | User&#39;s identity_url for OmniAuth authentication.  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**Embedded** | [**GroupModelAllOfEmbedded**](GroupModelAllOfEmbedded.md) |  | 

## Methods

### NewPrincipalCollectionModelAllOfEmbeddedElements

`func NewPrincipalCollectionModelAllOfEmbeddedElements(type_ string, id int64, name string, links GroupModelAllOfLinks, avatar string, embedded GroupModelAllOfEmbedded, ) *PrincipalCollectionModelAllOfEmbeddedElements`

NewPrincipalCollectionModelAllOfEmbeddedElements instantiates a new PrincipalCollectionModelAllOfEmbeddedElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalCollectionModelAllOfEmbeddedElementsWithDefaults

`func NewPrincipalCollectionModelAllOfEmbeddedElementsWithDefaults() *PrincipalCollectionModelAllOfEmbeddedElements`

NewPrincipalCollectionModelAllOfEmbeddedElementsWithDefaults instantiates a new PrincipalCollectionModelAllOfEmbeddedElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetLinks() GroupModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetLinksOk() (*GroupModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetLinks(v GroupModelAllOfLinks)`

SetLinks sets Links field to given value.


### GetAvatar

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetLogin

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetFirstName

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAdmin

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLanguage

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetIdentityUrl

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetIdentityUrl() string`

GetIdentityUrl returns the IdentityUrl field if non-nil, zero value otherwise.

### GetIdentityUrlOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetIdentityUrlOk() (*string, bool)`

GetIdentityUrlOk returns a tuple with the IdentityUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityUrl

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetIdentityUrl(v string)`

SetIdentityUrl sets IdentityUrl field to given value.

### HasIdentityUrl

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) HasIdentityUrl() bool`

HasIdentityUrl returns a boolean if a field has been set.

### GetEmbedded

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetEmbedded() GroupModelAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) GetEmbeddedOk() (*GroupModelAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PrincipalCollectionModelAllOfEmbeddedElements) SetEmbedded(v GroupModelAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


