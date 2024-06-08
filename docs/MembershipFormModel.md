# MembershipFormModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Embedded** | [**MembershipFormModelEmbedded**](MembershipFormModelEmbedded.md) |  | 
**Links** | [**MembershipFormModelLinks**](MembershipFormModelLinks.md) |  | 

## Methods

### NewMembershipFormModel

`func NewMembershipFormModel(type_ string, embedded MembershipFormModelEmbedded, links MembershipFormModelLinks, ) *MembershipFormModel`

NewMembershipFormModel instantiates a new MembershipFormModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipFormModelWithDefaults

`func NewMembershipFormModelWithDefaults() *MembershipFormModel`

NewMembershipFormModelWithDefaults instantiates a new MembershipFormModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MembershipFormModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MembershipFormModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MembershipFormModel) SetType(v string)`

SetType sets Type field to given value.


### GetEmbedded

`func (o *MembershipFormModel) GetEmbedded() MembershipFormModelEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *MembershipFormModel) GetEmbeddedOk() (*MembershipFormModelEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *MembershipFormModel) SetEmbedded(v MembershipFormModelEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetLinks

`func (o *MembershipFormModel) GetLinks() MembershipFormModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MembershipFormModel) GetLinksOk() (*MembershipFormModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MembershipFormModel) SetLinks(v MembershipFormModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


