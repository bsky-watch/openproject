# MembershipSchemaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dependencies** | Pointer to **[]string** | A list of dependencies between one property&#39;s value and another property | [optional] 
**Links** | [**SchemaModelLinks**](SchemaModelLinks.md) |  | 
**Id** | [**SchemaPropertyModel**](SchemaPropertyModel.md) |  | 
**CreatedAt** | [**SchemaPropertyModel**](SchemaPropertyModel.md) |  | 
**UpdatedAt** | [**SchemaPropertyModel**](SchemaPropertyModel.md) |  | 
**NotificationMessage** | [**SchemaPropertyModel**](SchemaPropertyModel.md) |  | 
**Project** | [**SchemaPropertyModel**](SchemaPropertyModel.md) |  | 
**Principal** | [**SchemaPropertyModel**](SchemaPropertyModel.md) |  | 
**Roles** | [**SchemaPropertyModel**](SchemaPropertyModel.md) |  | 

## Methods

### NewMembershipSchemaModel

`func NewMembershipSchemaModel(type_ string, links SchemaModelLinks, id SchemaPropertyModel, createdAt SchemaPropertyModel, updatedAt SchemaPropertyModel, notificationMessage SchemaPropertyModel, project SchemaPropertyModel, principal SchemaPropertyModel, roles SchemaPropertyModel, ) *MembershipSchemaModel`

NewMembershipSchemaModel instantiates a new MembershipSchemaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipSchemaModelWithDefaults

`func NewMembershipSchemaModelWithDefaults() *MembershipSchemaModel`

NewMembershipSchemaModelWithDefaults instantiates a new MembershipSchemaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MembershipSchemaModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MembershipSchemaModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MembershipSchemaModel) SetType(v string)`

SetType sets Type field to given value.


### GetDependencies

`func (o *MembershipSchemaModel) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *MembershipSchemaModel) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *MembershipSchemaModel) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *MembershipSchemaModel) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetLinks

`func (o *MembershipSchemaModel) GetLinks() SchemaModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MembershipSchemaModel) GetLinksOk() (*SchemaModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MembershipSchemaModel) SetLinks(v SchemaModelLinks)`

SetLinks sets Links field to given value.


### GetId

`func (o *MembershipSchemaModel) GetId() SchemaPropertyModel`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MembershipSchemaModel) GetIdOk() (*SchemaPropertyModel, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MembershipSchemaModel) SetId(v SchemaPropertyModel)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *MembershipSchemaModel) GetCreatedAt() SchemaPropertyModel`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MembershipSchemaModel) GetCreatedAtOk() (*SchemaPropertyModel, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MembershipSchemaModel) SetCreatedAt(v SchemaPropertyModel)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MembershipSchemaModel) GetUpdatedAt() SchemaPropertyModel`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MembershipSchemaModel) GetUpdatedAtOk() (*SchemaPropertyModel, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MembershipSchemaModel) SetUpdatedAt(v SchemaPropertyModel)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetNotificationMessage

`func (o *MembershipSchemaModel) GetNotificationMessage() SchemaPropertyModel`

GetNotificationMessage returns the NotificationMessage field if non-nil, zero value otherwise.

### GetNotificationMessageOk

`func (o *MembershipSchemaModel) GetNotificationMessageOk() (*SchemaPropertyModel, bool)`

GetNotificationMessageOk returns a tuple with the NotificationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMessage

`func (o *MembershipSchemaModel) SetNotificationMessage(v SchemaPropertyModel)`

SetNotificationMessage sets NotificationMessage field to given value.


### GetProject

`func (o *MembershipSchemaModel) GetProject() SchemaPropertyModel`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MembershipSchemaModel) GetProjectOk() (*SchemaPropertyModel, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MembershipSchemaModel) SetProject(v SchemaPropertyModel)`

SetProject sets Project field to given value.


### GetPrincipal

`func (o *MembershipSchemaModel) GetPrincipal() SchemaPropertyModel`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *MembershipSchemaModel) GetPrincipalOk() (*SchemaPropertyModel, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *MembershipSchemaModel) SetPrincipal(v SchemaPropertyModel)`

SetPrincipal sets Principal field to given value.


### GetRoles

`func (o *MembershipSchemaModel) GetRoles() SchemaPropertyModel`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MembershipSchemaModel) GetRolesOk() (*SchemaPropertyModel, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MembershipSchemaModel) SetRoles(v SchemaPropertyModel)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


