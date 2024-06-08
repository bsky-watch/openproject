# MembershipReadModelEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**ProjectModel**](ProjectModel.md) |  | [optional] 
**Principal** | Pointer to [**MembershipReadModelEmbeddedPrincipal**](MembershipReadModelEmbeddedPrincipal.md) |  | [optional] 
**Roles** | Pointer to [**[]RoleModel**](RoleModel.md) |  | [optional] 

## Methods

### NewMembershipReadModelEmbedded

`func NewMembershipReadModelEmbedded() *MembershipReadModelEmbedded`

NewMembershipReadModelEmbedded instantiates a new MembershipReadModelEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipReadModelEmbeddedWithDefaults

`func NewMembershipReadModelEmbeddedWithDefaults() *MembershipReadModelEmbedded`

NewMembershipReadModelEmbeddedWithDefaults instantiates a new MembershipReadModelEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *MembershipReadModelEmbedded) GetProject() ProjectModel`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MembershipReadModelEmbedded) GetProjectOk() (*ProjectModel, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MembershipReadModelEmbedded) SetProject(v ProjectModel)`

SetProject sets Project field to given value.

### HasProject

`func (o *MembershipReadModelEmbedded) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPrincipal

`func (o *MembershipReadModelEmbedded) GetPrincipal() MembershipReadModelEmbeddedPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *MembershipReadModelEmbedded) GetPrincipalOk() (*MembershipReadModelEmbeddedPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *MembershipReadModelEmbedded) SetPrincipal(v MembershipReadModelEmbeddedPrincipal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *MembershipReadModelEmbedded) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetRoles

`func (o *MembershipReadModelEmbedded) GetRoles() []RoleModel`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MembershipReadModelEmbedded) GetRolesOk() (*[]RoleModel, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MembershipReadModelEmbedded) SetRoles(v []RoleModel)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *MembershipReadModelEmbedded) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


