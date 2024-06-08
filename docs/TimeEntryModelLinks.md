# TimeEntryModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**TimeEntryModelLinksSelf**](TimeEntryModelLinksSelf.md) |  | 
**UpdateImmediately** | Pointer to [**TimeEntryModelLinksUpdateImmediately**](TimeEntryModelLinksUpdateImmediately.md) |  | [optional] 
**Update** | Pointer to [**TimeEntryModelLinksUpdate**](TimeEntryModelLinksUpdate.md) |  | [optional] 
**Delete** | Pointer to [**TimeEntryModelLinksDelete**](TimeEntryModelLinksDelete.md) |  | [optional] 
**Schema** | Pointer to [**TimeEntryModelLinksSchema**](TimeEntryModelLinksSchema.md) |  | [optional] 
**Project** | [**TimeEntryModelLinksProject**](TimeEntryModelLinksProject.md) |  | 
**WorkPackage** | Pointer to [**TimeEntryModelLinksWorkPackage**](TimeEntryModelLinksWorkPackage.md) |  | [optional] 
**User** | [**TimeEntryModelLinksUser**](TimeEntryModelLinksUser.md) |  | 
**Activity** | [**TimeEntryModelLinksActivity**](TimeEntryModelLinksActivity.md) |  | 

## Methods

### NewTimeEntryModelLinks

`func NewTimeEntryModelLinks(self TimeEntryModelLinksSelf, project TimeEntryModelLinksProject, user TimeEntryModelLinksUser, activity TimeEntryModelLinksActivity, ) *TimeEntryModelLinks`

NewTimeEntryModelLinks instantiates a new TimeEntryModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeEntryModelLinksWithDefaults

`func NewTimeEntryModelLinksWithDefaults() *TimeEntryModelLinks`

NewTimeEntryModelLinksWithDefaults instantiates a new TimeEntryModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *TimeEntryModelLinks) GetSelf() TimeEntryModelLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *TimeEntryModelLinks) GetSelfOk() (*TimeEntryModelLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *TimeEntryModelLinks) SetSelf(v TimeEntryModelLinksSelf)`

SetSelf sets Self field to given value.


### GetUpdateImmediately

`func (o *TimeEntryModelLinks) GetUpdateImmediately() TimeEntryModelLinksUpdateImmediately`

GetUpdateImmediately returns the UpdateImmediately field if non-nil, zero value otherwise.

### GetUpdateImmediatelyOk

`func (o *TimeEntryModelLinks) GetUpdateImmediatelyOk() (*TimeEntryModelLinksUpdateImmediately, bool)`

GetUpdateImmediatelyOk returns a tuple with the UpdateImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateImmediately

`func (o *TimeEntryModelLinks) SetUpdateImmediately(v TimeEntryModelLinksUpdateImmediately)`

SetUpdateImmediately sets UpdateImmediately field to given value.

### HasUpdateImmediately

`func (o *TimeEntryModelLinks) HasUpdateImmediately() bool`

HasUpdateImmediately returns a boolean if a field has been set.

### GetUpdate

`func (o *TimeEntryModelLinks) GetUpdate() TimeEntryModelLinksUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *TimeEntryModelLinks) GetUpdateOk() (*TimeEntryModelLinksUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *TimeEntryModelLinks) SetUpdate(v TimeEntryModelLinksUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *TimeEntryModelLinks) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *TimeEntryModelLinks) GetDelete() TimeEntryModelLinksDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *TimeEntryModelLinks) GetDeleteOk() (*TimeEntryModelLinksDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *TimeEntryModelLinks) SetDelete(v TimeEntryModelLinksDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *TimeEntryModelLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetSchema

`func (o *TimeEntryModelLinks) GetSchema() TimeEntryModelLinksSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *TimeEntryModelLinks) GetSchemaOk() (*TimeEntryModelLinksSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *TimeEntryModelLinks) SetSchema(v TimeEntryModelLinksSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *TimeEntryModelLinks) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetProject

`func (o *TimeEntryModelLinks) GetProject() TimeEntryModelLinksProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TimeEntryModelLinks) GetProjectOk() (*TimeEntryModelLinksProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TimeEntryModelLinks) SetProject(v TimeEntryModelLinksProject)`

SetProject sets Project field to given value.


### GetWorkPackage

`func (o *TimeEntryModelLinks) GetWorkPackage() TimeEntryModelLinksWorkPackage`

GetWorkPackage returns the WorkPackage field if non-nil, zero value otherwise.

### GetWorkPackageOk

`func (o *TimeEntryModelLinks) GetWorkPackageOk() (*TimeEntryModelLinksWorkPackage, bool)`

GetWorkPackageOk returns a tuple with the WorkPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPackage

`func (o *TimeEntryModelLinks) SetWorkPackage(v TimeEntryModelLinksWorkPackage)`

SetWorkPackage sets WorkPackage field to given value.

### HasWorkPackage

`func (o *TimeEntryModelLinks) HasWorkPackage() bool`

HasWorkPackage returns a boolean if a field has been set.

### GetUser

`func (o *TimeEntryModelLinks) GetUser() TimeEntryModelLinksUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TimeEntryModelLinks) GetUserOk() (*TimeEntryModelLinksUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TimeEntryModelLinks) SetUser(v TimeEntryModelLinksUser)`

SetUser sets User field to given value.


### GetActivity

`func (o *TimeEntryModelLinks) GetActivity() TimeEntryModelLinksActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *TimeEntryModelLinks) GetActivityOk() (*TimeEntryModelLinksActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *TimeEntryModelLinks) SetActivity(v TimeEntryModelLinksActivity)`

SetActivity sets Activity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


