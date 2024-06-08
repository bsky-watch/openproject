# MembershipReadModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**MembershipReadModelLinksSelf**](MembershipReadModelLinksSelf.md) |  | 
**Schema** | [**MembershipReadModelLinksSchema**](MembershipReadModelLinksSchema.md) |  | 
**Update** | Pointer to [**MembershipReadModelLinksUpdate**](MembershipReadModelLinksUpdate.md) |  | [optional] 
**UpdateImmediately** | Pointer to [**MembershipReadModelLinksUpdateImmediately**](MembershipReadModelLinksUpdateImmediately.md) |  | [optional] 
**Project** | [**MembershipReadModelLinksProject**](MembershipReadModelLinksProject.md) |  | 
**Principal** | [**MembershipReadModelLinksPrincipal**](MembershipReadModelLinksPrincipal.md) |  | 
**Roles** | [**[]MembershipReadModelLinksRolesInner**](MembershipReadModelLinksRolesInner.md) |  | 

## Methods

### NewMembershipReadModelLinks

`func NewMembershipReadModelLinks(self MembershipReadModelLinksSelf, schema MembershipReadModelLinksSchema, project MembershipReadModelLinksProject, principal MembershipReadModelLinksPrincipal, roles []MembershipReadModelLinksRolesInner, ) *MembershipReadModelLinks`

NewMembershipReadModelLinks instantiates a new MembershipReadModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipReadModelLinksWithDefaults

`func NewMembershipReadModelLinksWithDefaults() *MembershipReadModelLinks`

NewMembershipReadModelLinksWithDefaults instantiates a new MembershipReadModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *MembershipReadModelLinks) GetSelf() MembershipReadModelLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MembershipReadModelLinks) GetSelfOk() (*MembershipReadModelLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *MembershipReadModelLinks) SetSelf(v MembershipReadModelLinksSelf)`

SetSelf sets Self field to given value.


### GetSchema

`func (o *MembershipReadModelLinks) GetSchema() MembershipReadModelLinksSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *MembershipReadModelLinks) GetSchemaOk() (*MembershipReadModelLinksSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *MembershipReadModelLinks) SetSchema(v MembershipReadModelLinksSchema)`

SetSchema sets Schema field to given value.


### GetUpdate

`func (o *MembershipReadModelLinks) GetUpdate() MembershipReadModelLinksUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *MembershipReadModelLinks) GetUpdateOk() (*MembershipReadModelLinksUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *MembershipReadModelLinks) SetUpdate(v MembershipReadModelLinksUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *MembershipReadModelLinks) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUpdateImmediately

`func (o *MembershipReadModelLinks) GetUpdateImmediately() MembershipReadModelLinksUpdateImmediately`

GetUpdateImmediately returns the UpdateImmediately field if non-nil, zero value otherwise.

### GetUpdateImmediatelyOk

`func (o *MembershipReadModelLinks) GetUpdateImmediatelyOk() (*MembershipReadModelLinksUpdateImmediately, bool)`

GetUpdateImmediatelyOk returns a tuple with the UpdateImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateImmediately

`func (o *MembershipReadModelLinks) SetUpdateImmediately(v MembershipReadModelLinksUpdateImmediately)`

SetUpdateImmediately sets UpdateImmediately field to given value.

### HasUpdateImmediately

`func (o *MembershipReadModelLinks) HasUpdateImmediately() bool`

HasUpdateImmediately returns a boolean if a field has been set.

### GetProject

`func (o *MembershipReadModelLinks) GetProject() MembershipReadModelLinksProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MembershipReadModelLinks) GetProjectOk() (*MembershipReadModelLinksProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MembershipReadModelLinks) SetProject(v MembershipReadModelLinksProject)`

SetProject sets Project field to given value.


### GetPrincipal

`func (o *MembershipReadModelLinks) GetPrincipal() MembershipReadModelLinksPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *MembershipReadModelLinks) GetPrincipalOk() (*MembershipReadModelLinksPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *MembershipReadModelLinks) SetPrincipal(v MembershipReadModelLinksPrincipal)`

SetPrincipal sets Principal field to given value.


### GetRoles

`func (o *MembershipReadModelLinks) GetRoles() []MembershipReadModelLinksRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MembershipReadModelLinks) GetRolesOk() (*[]MembershipReadModelLinksRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MembershipReadModelLinks) SetRoles(v []MembershipReadModelLinksRolesInner)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


