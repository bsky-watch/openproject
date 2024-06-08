# MembershipWriteModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | Pointer to [**MembershipWriteModelLinksPrincipal**](MembershipWriteModelLinksPrincipal.md) |  | [optional] 
**Roles** | Pointer to [**[]MembershipReadModelLinksRolesInner**](MembershipReadModelLinksRolesInner.md) |  | [optional] 
**Project** | Pointer to [**MembershipWriteModelLinksProject**](MembershipWriteModelLinksProject.md) |  | [optional] 

## Methods

### NewMembershipWriteModelLinks

`func NewMembershipWriteModelLinks() *MembershipWriteModelLinks`

NewMembershipWriteModelLinks instantiates a new MembershipWriteModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipWriteModelLinksWithDefaults

`func NewMembershipWriteModelLinksWithDefaults() *MembershipWriteModelLinks`

NewMembershipWriteModelLinksWithDefaults instantiates a new MembershipWriteModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *MembershipWriteModelLinks) GetPrincipal() MembershipWriteModelLinksPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *MembershipWriteModelLinks) GetPrincipalOk() (*MembershipWriteModelLinksPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *MembershipWriteModelLinks) SetPrincipal(v MembershipWriteModelLinksPrincipal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *MembershipWriteModelLinks) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetRoles

`func (o *MembershipWriteModelLinks) GetRoles() []MembershipReadModelLinksRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MembershipWriteModelLinks) GetRolesOk() (*[]MembershipReadModelLinksRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MembershipWriteModelLinks) SetRoles(v []MembershipReadModelLinksRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *MembershipWriteModelLinks) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetProject

`func (o *MembershipWriteModelLinks) GetProject() MembershipWriteModelLinksProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MembershipWriteModelLinks) GetProjectOk() (*MembershipWriteModelLinksProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MembershipWriteModelLinks) SetProject(v MembershipWriteModelLinksProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *MembershipWriteModelLinks) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


