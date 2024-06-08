# UserModel

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

### NewUserModel

`func NewUserModel(type_ string, id int64, name string, links UserModelAllOfLinks, avatar string, ) *UserModel`

NewUserModel instantiates a new UserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelWithDefaults

`func NewUserModelWithDefaults() *UserModel`

NewUserModelWithDefaults instantiates a new UserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserModel) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *UserModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserModel) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *UserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserModel) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *UserModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *UserModel) GetLinks() UserModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserModel) GetLinksOk() (*UserModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserModel) SetLinks(v UserModelAllOfLinks)`

SetLinks sets Links field to given value.


### GetAvatar

`func (o *UserModel) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserModel) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserModel) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetLogin

`func (o *UserModel) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserModel) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserModel) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *UserModel) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetFirstName

`func (o *UserModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UserModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAdmin

`func (o *UserModel) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserModel) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserModel) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserModel) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetStatus

`func (o *UserModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLanguage

`func (o *UserModel) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserModel) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserModel) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserModel) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetIdentityUrl

`func (o *UserModel) GetIdentityUrl() string`

GetIdentityUrl returns the IdentityUrl field if non-nil, zero value otherwise.

### GetIdentityUrlOk

`func (o *UserModel) GetIdentityUrlOk() (*string, bool)`

GetIdentityUrlOk returns a tuple with the IdentityUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityUrl

`func (o *UserModel) SetIdentityUrl(v string)`

SetIdentityUrl sets IdentityUrl field to given value.

### HasIdentityUrl

`func (o *UserModel) HasIdentityUrl() bool`

HasIdentityUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


