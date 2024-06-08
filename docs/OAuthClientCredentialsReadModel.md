# OAuthClientCredentialsReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Type** | **string** |  | 
**ClientId** | **string** | OAuth 2 client id | 
**Confidential** | **bool** | true, if OAuth 2 credentials are confidential, false, if no secret is stored | 
**CreatedAt** | Pointer to **time.Time** | The time the OAuth client credentials were created at | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the OAuth client credentials were last updated | [optional] 
**Links** | [**OAuthClientCredentialsReadModelLinks**](OAuthClientCredentialsReadModelLinks.md) |  | 

## Methods

### NewOAuthClientCredentialsReadModel

`func NewOAuthClientCredentialsReadModel(id int64, type_ string, clientId string, confidential bool, links OAuthClientCredentialsReadModelLinks, ) *OAuthClientCredentialsReadModel`

NewOAuthClientCredentialsReadModel instantiates a new OAuthClientCredentialsReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientCredentialsReadModelWithDefaults

`func NewOAuthClientCredentialsReadModelWithDefaults() *OAuthClientCredentialsReadModel`

NewOAuthClientCredentialsReadModelWithDefaults instantiates a new OAuthClientCredentialsReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OAuthClientCredentialsReadModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuthClientCredentialsReadModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuthClientCredentialsReadModel) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *OAuthClientCredentialsReadModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuthClientCredentialsReadModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuthClientCredentialsReadModel) SetType(v string)`

SetType sets Type field to given value.


### GetClientId

`func (o *OAuthClientCredentialsReadModel) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthClientCredentialsReadModel) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthClientCredentialsReadModel) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetConfidential

`func (o *OAuthClientCredentialsReadModel) GetConfidential() bool`

GetConfidential returns the Confidential field if non-nil, zero value otherwise.

### GetConfidentialOk

`func (o *OAuthClientCredentialsReadModel) GetConfidentialOk() (*bool, bool)`

GetConfidentialOk returns a tuple with the Confidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidential

`func (o *OAuthClientCredentialsReadModel) SetConfidential(v bool)`

SetConfidential sets Confidential field to given value.


### GetCreatedAt

`func (o *OAuthClientCredentialsReadModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OAuthClientCredentialsReadModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OAuthClientCredentialsReadModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OAuthClientCredentialsReadModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OAuthClientCredentialsReadModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OAuthClientCredentialsReadModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OAuthClientCredentialsReadModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OAuthClientCredentialsReadModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *OAuthClientCredentialsReadModel) GetLinks() OAuthClientCredentialsReadModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuthClientCredentialsReadModel) GetLinksOk() (*OAuthClientCredentialsReadModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuthClientCredentialsReadModel) SetLinks(v OAuthClientCredentialsReadModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


