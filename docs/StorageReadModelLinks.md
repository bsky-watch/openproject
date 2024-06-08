# StorageReadModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**StorageReadModelLinksSelf**](StorageReadModelLinksSelf.md) |  | 
**Type** | [**StorageReadModelLinksType**](StorageReadModelLinksType.md) |  | 
**Origin** | [**StorageReadModelLinksOrigin**](StorageReadModelLinksOrigin.md) |  | 
**Open** | [**StorageReadModelLinksOpen**](StorageReadModelLinksOpen.md) |  | 
**AuthorizationState** | [**StorageReadModelLinksAuthorizationState**](StorageReadModelLinksAuthorizationState.md) |  | 
**Authorize** | Pointer to [**StorageReadModelLinksAuthorize**](StorageReadModelLinksAuthorize.md) |  | [optional] 
**OauthApplication** | Pointer to [**StorageReadModelLinksOauthApplication**](StorageReadModelLinksOauthApplication.md) |  | [optional] 
**OauthClientCredentials** | Pointer to [**StorageReadModelLinksOauthClientCredentials**](StorageReadModelLinksOauthClientCredentials.md) |  | [optional] 

## Methods

### NewStorageReadModelLinks

`func NewStorageReadModelLinks(self StorageReadModelLinksSelf, type_ StorageReadModelLinksType, origin StorageReadModelLinksOrigin, open StorageReadModelLinksOpen, authorizationState StorageReadModelLinksAuthorizationState, ) *StorageReadModelLinks`

NewStorageReadModelLinks instantiates a new StorageReadModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageReadModelLinksWithDefaults

`func NewStorageReadModelLinksWithDefaults() *StorageReadModelLinks`

NewStorageReadModelLinksWithDefaults instantiates a new StorageReadModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *StorageReadModelLinks) GetSelf() StorageReadModelLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *StorageReadModelLinks) GetSelfOk() (*StorageReadModelLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *StorageReadModelLinks) SetSelf(v StorageReadModelLinksSelf)`

SetSelf sets Self field to given value.


### GetType

`func (o *StorageReadModelLinks) GetType() StorageReadModelLinksType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageReadModelLinks) GetTypeOk() (*StorageReadModelLinksType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageReadModelLinks) SetType(v StorageReadModelLinksType)`

SetType sets Type field to given value.


### GetOrigin

`func (o *StorageReadModelLinks) GetOrigin() StorageReadModelLinksOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *StorageReadModelLinks) GetOriginOk() (*StorageReadModelLinksOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *StorageReadModelLinks) SetOrigin(v StorageReadModelLinksOrigin)`

SetOrigin sets Origin field to given value.


### GetOpen

`func (o *StorageReadModelLinks) GetOpen() StorageReadModelLinksOpen`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *StorageReadModelLinks) GetOpenOk() (*StorageReadModelLinksOpen, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *StorageReadModelLinks) SetOpen(v StorageReadModelLinksOpen)`

SetOpen sets Open field to given value.


### GetAuthorizationState

`func (o *StorageReadModelLinks) GetAuthorizationState() StorageReadModelLinksAuthorizationState`

GetAuthorizationState returns the AuthorizationState field if non-nil, zero value otherwise.

### GetAuthorizationStateOk

`func (o *StorageReadModelLinks) GetAuthorizationStateOk() (*StorageReadModelLinksAuthorizationState, bool)`

GetAuthorizationStateOk returns a tuple with the AuthorizationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationState

`func (o *StorageReadModelLinks) SetAuthorizationState(v StorageReadModelLinksAuthorizationState)`

SetAuthorizationState sets AuthorizationState field to given value.


### GetAuthorize

`func (o *StorageReadModelLinks) GetAuthorize() StorageReadModelLinksAuthorize`

GetAuthorize returns the Authorize field if non-nil, zero value otherwise.

### GetAuthorizeOk

`func (o *StorageReadModelLinks) GetAuthorizeOk() (*StorageReadModelLinksAuthorize, bool)`

GetAuthorizeOk returns a tuple with the Authorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorize

`func (o *StorageReadModelLinks) SetAuthorize(v StorageReadModelLinksAuthorize)`

SetAuthorize sets Authorize field to given value.

### HasAuthorize

`func (o *StorageReadModelLinks) HasAuthorize() bool`

HasAuthorize returns a boolean if a field has been set.

### GetOauthApplication

`func (o *StorageReadModelLinks) GetOauthApplication() StorageReadModelLinksOauthApplication`

GetOauthApplication returns the OauthApplication field if non-nil, zero value otherwise.

### GetOauthApplicationOk

`func (o *StorageReadModelLinks) GetOauthApplicationOk() (*StorageReadModelLinksOauthApplication, bool)`

GetOauthApplicationOk returns a tuple with the OauthApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthApplication

`func (o *StorageReadModelLinks) SetOauthApplication(v StorageReadModelLinksOauthApplication)`

SetOauthApplication sets OauthApplication field to given value.

### HasOauthApplication

`func (o *StorageReadModelLinks) HasOauthApplication() bool`

HasOauthApplication returns a boolean if a field has been set.

### GetOauthClientCredentials

`func (o *StorageReadModelLinks) GetOauthClientCredentials() StorageReadModelLinksOauthClientCredentials`

GetOauthClientCredentials returns the OauthClientCredentials field if non-nil, zero value otherwise.

### GetOauthClientCredentialsOk

`func (o *StorageReadModelLinks) GetOauthClientCredentialsOk() (*StorageReadModelLinksOauthClientCredentials, bool)`

GetOauthClientCredentialsOk returns a tuple with the OauthClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientCredentials

`func (o *StorageReadModelLinks) SetOauthClientCredentials(v StorageReadModelLinksOauthClientCredentials)`

SetOauthClientCredentials sets OauthClientCredentials field to given value.

### HasOauthClientCredentials

`func (o *StorageReadModelLinks) HasOauthClientCredentials() bool`

HasOauthClientCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


