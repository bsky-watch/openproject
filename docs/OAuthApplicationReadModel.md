# OAuthApplicationReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Type** | **string** |  | 
**Name** | **string** | The name of the OAuth 2 application | 
**ClientId** | **string** | OAuth 2 client id | 
**ClientSecret** | Pointer to **string** | OAuth 2 client secret. This is only returned when creating a new OAuth application. | [optional] 
**Confidential** | **bool** | true, if OAuth 2 credentials are confidential, false, if no secret is stored | 
**CreatedAt** | Pointer to **time.Time** | The time the OAuth 2 Application was created at | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the OAuth 2 Application was last updated | [optional] 
**Scopes** | Pointer to **[]string** | An array of the scopes of the OAuth 2 Application | [optional] 
**Links** | Pointer to [**OAuthApplicationReadModelLinks**](OAuthApplicationReadModelLinks.md) |  | [optional] 

## Methods

### NewOAuthApplicationReadModel

`func NewOAuthApplicationReadModel(id int64, type_ string, name string, clientId string, confidential bool, ) *OAuthApplicationReadModel`

NewOAuthApplicationReadModel instantiates a new OAuthApplicationReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthApplicationReadModelWithDefaults

`func NewOAuthApplicationReadModelWithDefaults() *OAuthApplicationReadModel`

NewOAuthApplicationReadModelWithDefaults instantiates a new OAuthApplicationReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OAuthApplicationReadModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuthApplicationReadModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuthApplicationReadModel) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *OAuthApplicationReadModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuthApplicationReadModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuthApplicationReadModel) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *OAuthApplicationReadModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuthApplicationReadModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuthApplicationReadModel) SetName(v string)`

SetName sets Name field to given value.


### GetClientId

`func (o *OAuthApplicationReadModel) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthApplicationReadModel) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthApplicationReadModel) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuthApplicationReadModel) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthApplicationReadModel) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthApplicationReadModel) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthApplicationReadModel) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetConfidential

`func (o *OAuthApplicationReadModel) GetConfidential() bool`

GetConfidential returns the Confidential field if non-nil, zero value otherwise.

### GetConfidentialOk

`func (o *OAuthApplicationReadModel) GetConfidentialOk() (*bool, bool)`

GetConfidentialOk returns a tuple with the Confidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidential

`func (o *OAuthApplicationReadModel) SetConfidential(v bool)`

SetConfidential sets Confidential field to given value.


### GetCreatedAt

`func (o *OAuthApplicationReadModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OAuthApplicationReadModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OAuthApplicationReadModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OAuthApplicationReadModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OAuthApplicationReadModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OAuthApplicationReadModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OAuthApplicationReadModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OAuthApplicationReadModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetScopes

`func (o *OAuthApplicationReadModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuthApplicationReadModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuthApplicationReadModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuthApplicationReadModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetLinks

`func (o *OAuthApplicationReadModel) GetLinks() OAuthApplicationReadModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuthApplicationReadModel) GetLinksOk() (*OAuthApplicationReadModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuthApplicationReadModel) SetLinks(v OAuthApplicationReadModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuthApplicationReadModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


