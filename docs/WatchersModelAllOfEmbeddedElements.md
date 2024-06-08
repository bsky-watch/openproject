# WatchersModelAllOfEmbeddedElements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** | The principal&#39;s unique identifier. | 
**Name** | **string** | The principal&#39;s display name, layout depends on instance settings. | 
**CreatedAt** | Pointer to **time.Time** | Time of creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time of the most recent change to the user | [optional] 
**Links** | [**UserModelAllOfLinks**](UserModelAllOfLinks.md) |  | 
**Avatar** | **string** | URL to user&#39;s avatar | 
**Login** | Pointer to **string** | The user&#39;s login name  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**FirstName** | Pointer to **string** | The user&#39;s first name  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address  # Conditions  - E-Mail address not hidden - User is not a new record - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**Admin** | Pointer to **bool** | Flag indicating whether or not the user is an admin  # Conditions  - &#x60;admin&#x60; | [optional] 
**Status** | Pointer to **string** | The current activation status of the user.  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**Language** | Pointer to **string** | User&#39;s language | ISO 639-1 format  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 
**IdentityUrl** | Pointer to **string** | User&#39;s identity_url for OmniAuth authentication.  # Conditions  - User is self, or &#x60;create_user&#x60; or &#x60;manage_user&#x60; permission globally | [optional] 

## Methods

### NewWatchersModelAllOfEmbeddedElements

`func NewWatchersModelAllOfEmbeddedElements(type_ string, id int64, name string, links UserModelAllOfLinks, avatar string, ) *WatchersModelAllOfEmbeddedElements`

NewWatchersModelAllOfEmbeddedElements instantiates a new WatchersModelAllOfEmbeddedElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchersModelAllOfEmbeddedElementsWithDefaults

`func NewWatchersModelAllOfEmbeddedElementsWithDefaults() *WatchersModelAllOfEmbeddedElements`

NewWatchersModelAllOfEmbeddedElementsWithDefaults instantiates a new WatchersModelAllOfEmbeddedElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WatchersModelAllOfEmbeddedElements) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WatchersModelAllOfEmbeddedElements) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WatchersModelAllOfEmbeddedElements) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *WatchersModelAllOfEmbeddedElements) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WatchersModelAllOfEmbeddedElements) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WatchersModelAllOfEmbeddedElements) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *WatchersModelAllOfEmbeddedElements) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WatchersModelAllOfEmbeddedElements) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WatchersModelAllOfEmbeddedElements) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *WatchersModelAllOfEmbeddedElements) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WatchersModelAllOfEmbeddedElements) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WatchersModelAllOfEmbeddedElements) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WatchersModelAllOfEmbeddedElements) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WatchersModelAllOfEmbeddedElements) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WatchersModelAllOfEmbeddedElements) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WatchersModelAllOfEmbeddedElements) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WatchersModelAllOfEmbeddedElements) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *WatchersModelAllOfEmbeddedElements) GetLinks() UserModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WatchersModelAllOfEmbeddedElements) GetLinksOk() (*UserModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WatchersModelAllOfEmbeddedElements) SetLinks(v UserModelAllOfLinks)`

SetLinks sets Links field to given value.


### GetAvatar

`func (o *WatchersModelAllOfEmbeddedElements) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *WatchersModelAllOfEmbeddedElements) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *WatchersModelAllOfEmbeddedElements) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetLogin

`func (o *WatchersModelAllOfEmbeddedElements) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *WatchersModelAllOfEmbeddedElements) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *WatchersModelAllOfEmbeddedElements) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *WatchersModelAllOfEmbeddedElements) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetFirstName

`func (o *WatchersModelAllOfEmbeddedElements) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *WatchersModelAllOfEmbeddedElements) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *WatchersModelAllOfEmbeddedElements) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *WatchersModelAllOfEmbeddedElements) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *WatchersModelAllOfEmbeddedElements) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *WatchersModelAllOfEmbeddedElements) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *WatchersModelAllOfEmbeddedElements) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *WatchersModelAllOfEmbeddedElements) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *WatchersModelAllOfEmbeddedElements) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WatchersModelAllOfEmbeddedElements) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WatchersModelAllOfEmbeddedElements) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WatchersModelAllOfEmbeddedElements) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAdmin

`func (o *WatchersModelAllOfEmbeddedElements) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *WatchersModelAllOfEmbeddedElements) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *WatchersModelAllOfEmbeddedElements) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *WatchersModelAllOfEmbeddedElements) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetStatus

`func (o *WatchersModelAllOfEmbeddedElements) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WatchersModelAllOfEmbeddedElements) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WatchersModelAllOfEmbeddedElements) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WatchersModelAllOfEmbeddedElements) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLanguage

`func (o *WatchersModelAllOfEmbeddedElements) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WatchersModelAllOfEmbeddedElements) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WatchersModelAllOfEmbeddedElements) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *WatchersModelAllOfEmbeddedElements) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetIdentityUrl

`func (o *WatchersModelAllOfEmbeddedElements) GetIdentityUrl() string`

GetIdentityUrl returns the IdentityUrl field if non-nil, zero value otherwise.

### GetIdentityUrlOk

`func (o *WatchersModelAllOfEmbeddedElements) GetIdentityUrlOk() (*string, bool)`

GetIdentityUrlOk returns a tuple with the IdentityUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityUrl

`func (o *WatchersModelAllOfEmbeddedElements) SetIdentityUrl(v string)`

SetIdentityUrl sets IdentityUrl field to given value.

### HasIdentityUrl

`func (o *WatchersModelAllOfEmbeddedElements) HasIdentityUrl() bool`

HasIdentityUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


