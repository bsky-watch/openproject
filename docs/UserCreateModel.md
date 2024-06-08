# UserCreateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | **bool** |  | 
**Email** | **string** |  | 
**Login** | **string** |  | 
**Password** | Pointer to **string** | The users password.  *Conditions:*  Only writable on creation, not on update. | [optional] 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Status** | Pointer to **string** | The current activation status of the user.  *Conditions:*  Only writable on creation, not on update. | [optional] 
**Language** | **string** |  | 

## Methods

### NewUserCreateModel

`func NewUserCreateModel(admin bool, email string, login string, firstName string, lastName string, language string, ) *UserCreateModel`

NewUserCreateModel instantiates a new UserCreateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateModelWithDefaults

`func NewUserCreateModelWithDefaults() *UserCreateModel`

NewUserCreateModelWithDefaults instantiates a new UserCreateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *UserCreateModel) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserCreateModel) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserCreateModel) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.


### GetEmail

`func (o *UserCreateModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreateModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreateModel) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLogin

`func (o *UserCreateModel) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserCreateModel) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserCreateModel) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *UserCreateModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreateModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreateModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreateModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetFirstName

`func (o *UserCreateModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCreateModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCreateModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserCreateModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCreateModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCreateModel) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetStatus

`func (o *UserCreateModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserCreateModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserCreateModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserCreateModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLanguage

`func (o *UserCreateModel) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserCreateModel) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserCreateModel) SetLanguage(v string)`

SetLanguage sets Language field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


