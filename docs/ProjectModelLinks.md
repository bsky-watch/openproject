# ProjectModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Update** | Pointer to [**ProjectModelLinksUpdate**](ProjectModelLinksUpdate.md) |  | [optional] 
**UpdateImmediately** | Pointer to [**ProjectModelLinksUpdateImmediately**](ProjectModelLinksUpdateImmediately.md) |  | [optional] 
**Delete** | Pointer to [**ProjectModelLinksDelete**](ProjectModelLinksDelete.md) |  | [optional] 
**CreateWorkPackage** | Pointer to [**ProjectModelLinksCreateWorkPackage**](ProjectModelLinksCreateWorkPackage.md) |  | [optional] 
**CreateWorkPackageImmediately** | Pointer to [**ProjectModelLinksCreateWorkPackageImmediately**](ProjectModelLinksCreateWorkPackageImmediately.md) |  | [optional] 
**Self** | [**ProjectModelLinksSelf**](ProjectModelLinksSelf.md) |  | 
**Categories** | [**ProjectModelLinksCategories**](ProjectModelLinksCategories.md) |  | 
**Types** | [**ProjectModelLinksTypes**](ProjectModelLinksTypes.md) |  | 
**Versions** | [**ProjectModelLinksVersions**](ProjectModelLinksVersions.md) |  | 
**Memberships** | [**ProjectModelLinksMemberships**](ProjectModelLinksMemberships.md) |  | 
**WorkPackages** | [**ProjectModelLinksWorkPackages**](ProjectModelLinksWorkPackages.md) |  | 
**Parent** | Pointer to [**ProjectModelLinksParent**](ProjectModelLinksParent.md) |  | [optional] 
**Status** | Pointer to [**ProjectModelLinksStatus**](ProjectModelLinksStatus.md) |  | [optional] 
**Storages** | Pointer to [**[]ProjectModelLinksStoragesInner**](ProjectModelLinksStoragesInner.md) |  | [optional] 
**ProjectStorages** | Pointer to [**ProjectModelLinksProjectStorages**](ProjectModelLinksProjectStorages.md) |  | [optional] 
**Ancestors** | Pointer to [**[]ProjectModelLinksAncestorsInner**](ProjectModelLinksAncestorsInner.md) |  | [optional] 

## Methods

### NewProjectModelLinks

`func NewProjectModelLinks(self ProjectModelLinksSelf, categories ProjectModelLinksCategories, types ProjectModelLinksTypes, versions ProjectModelLinksVersions, memberships ProjectModelLinksMemberships, workPackages ProjectModelLinksWorkPackages, ) *ProjectModelLinks`

NewProjectModelLinks instantiates a new ProjectModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectModelLinksWithDefaults

`func NewProjectModelLinksWithDefaults() *ProjectModelLinks`

NewProjectModelLinksWithDefaults instantiates a new ProjectModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdate

`func (o *ProjectModelLinks) GetUpdate() ProjectModelLinksUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ProjectModelLinks) GetUpdateOk() (*ProjectModelLinksUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ProjectModelLinks) SetUpdate(v ProjectModelLinksUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ProjectModelLinks) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUpdateImmediately

`func (o *ProjectModelLinks) GetUpdateImmediately() ProjectModelLinksUpdateImmediately`

GetUpdateImmediately returns the UpdateImmediately field if non-nil, zero value otherwise.

### GetUpdateImmediatelyOk

`func (o *ProjectModelLinks) GetUpdateImmediatelyOk() (*ProjectModelLinksUpdateImmediately, bool)`

GetUpdateImmediatelyOk returns a tuple with the UpdateImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateImmediately

`func (o *ProjectModelLinks) SetUpdateImmediately(v ProjectModelLinksUpdateImmediately)`

SetUpdateImmediately sets UpdateImmediately field to given value.

### HasUpdateImmediately

`func (o *ProjectModelLinks) HasUpdateImmediately() bool`

HasUpdateImmediately returns a boolean if a field has been set.

### GetDelete

`func (o *ProjectModelLinks) GetDelete() ProjectModelLinksDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ProjectModelLinks) GetDeleteOk() (*ProjectModelLinksDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ProjectModelLinks) SetDelete(v ProjectModelLinksDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ProjectModelLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetCreateWorkPackage

`func (o *ProjectModelLinks) GetCreateWorkPackage() ProjectModelLinksCreateWorkPackage`

GetCreateWorkPackage returns the CreateWorkPackage field if non-nil, zero value otherwise.

### GetCreateWorkPackageOk

`func (o *ProjectModelLinks) GetCreateWorkPackageOk() (*ProjectModelLinksCreateWorkPackage, bool)`

GetCreateWorkPackageOk returns a tuple with the CreateWorkPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateWorkPackage

`func (o *ProjectModelLinks) SetCreateWorkPackage(v ProjectModelLinksCreateWorkPackage)`

SetCreateWorkPackage sets CreateWorkPackage field to given value.

### HasCreateWorkPackage

`func (o *ProjectModelLinks) HasCreateWorkPackage() bool`

HasCreateWorkPackage returns a boolean if a field has been set.

### GetCreateWorkPackageImmediately

`func (o *ProjectModelLinks) GetCreateWorkPackageImmediately() ProjectModelLinksCreateWorkPackageImmediately`

GetCreateWorkPackageImmediately returns the CreateWorkPackageImmediately field if non-nil, zero value otherwise.

### GetCreateWorkPackageImmediatelyOk

`func (o *ProjectModelLinks) GetCreateWorkPackageImmediatelyOk() (*ProjectModelLinksCreateWorkPackageImmediately, bool)`

GetCreateWorkPackageImmediatelyOk returns a tuple with the CreateWorkPackageImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateWorkPackageImmediately

`func (o *ProjectModelLinks) SetCreateWorkPackageImmediately(v ProjectModelLinksCreateWorkPackageImmediately)`

SetCreateWorkPackageImmediately sets CreateWorkPackageImmediately field to given value.

### HasCreateWorkPackageImmediately

`func (o *ProjectModelLinks) HasCreateWorkPackageImmediately() bool`

HasCreateWorkPackageImmediately returns a boolean if a field has been set.

### GetSelf

`func (o *ProjectModelLinks) GetSelf() ProjectModelLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ProjectModelLinks) GetSelfOk() (*ProjectModelLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ProjectModelLinks) SetSelf(v ProjectModelLinksSelf)`

SetSelf sets Self field to given value.


### GetCategories

`func (o *ProjectModelLinks) GetCategories() ProjectModelLinksCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ProjectModelLinks) GetCategoriesOk() (*ProjectModelLinksCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ProjectModelLinks) SetCategories(v ProjectModelLinksCategories)`

SetCategories sets Categories field to given value.


### GetTypes

`func (o *ProjectModelLinks) GetTypes() ProjectModelLinksTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ProjectModelLinks) GetTypesOk() (*ProjectModelLinksTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ProjectModelLinks) SetTypes(v ProjectModelLinksTypes)`

SetTypes sets Types field to given value.


### GetVersions

`func (o *ProjectModelLinks) GetVersions() ProjectModelLinksVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ProjectModelLinks) GetVersionsOk() (*ProjectModelLinksVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ProjectModelLinks) SetVersions(v ProjectModelLinksVersions)`

SetVersions sets Versions field to given value.


### GetMemberships

`func (o *ProjectModelLinks) GetMemberships() ProjectModelLinksMemberships`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *ProjectModelLinks) GetMembershipsOk() (*ProjectModelLinksMemberships, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *ProjectModelLinks) SetMemberships(v ProjectModelLinksMemberships)`

SetMemberships sets Memberships field to given value.


### GetWorkPackages

`func (o *ProjectModelLinks) GetWorkPackages() ProjectModelLinksWorkPackages`

GetWorkPackages returns the WorkPackages field if non-nil, zero value otherwise.

### GetWorkPackagesOk

`func (o *ProjectModelLinks) GetWorkPackagesOk() (*ProjectModelLinksWorkPackages, bool)`

GetWorkPackagesOk returns a tuple with the WorkPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPackages

`func (o *ProjectModelLinks) SetWorkPackages(v ProjectModelLinksWorkPackages)`

SetWorkPackages sets WorkPackages field to given value.


### GetParent

`func (o *ProjectModelLinks) GetParent() ProjectModelLinksParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ProjectModelLinks) GetParentOk() (*ProjectModelLinksParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ProjectModelLinks) SetParent(v ProjectModelLinksParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ProjectModelLinks) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectModelLinks) GetStatus() ProjectModelLinksStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectModelLinks) GetStatusOk() (*ProjectModelLinksStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectModelLinks) SetStatus(v ProjectModelLinksStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectModelLinks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorages

`func (o *ProjectModelLinks) GetStorages() []ProjectModelLinksStoragesInner`

GetStorages returns the Storages field if non-nil, zero value otherwise.

### GetStoragesOk

`func (o *ProjectModelLinks) GetStoragesOk() (*[]ProjectModelLinksStoragesInner, bool)`

GetStoragesOk returns a tuple with the Storages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorages

`func (o *ProjectModelLinks) SetStorages(v []ProjectModelLinksStoragesInner)`

SetStorages sets Storages field to given value.

### HasStorages

`func (o *ProjectModelLinks) HasStorages() bool`

HasStorages returns a boolean if a field has been set.

### GetProjectStorages

`func (o *ProjectModelLinks) GetProjectStorages() ProjectModelLinksProjectStorages`

GetProjectStorages returns the ProjectStorages field if non-nil, zero value otherwise.

### GetProjectStoragesOk

`func (o *ProjectModelLinks) GetProjectStoragesOk() (*ProjectModelLinksProjectStorages, bool)`

GetProjectStoragesOk returns a tuple with the ProjectStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStorages

`func (o *ProjectModelLinks) SetProjectStorages(v ProjectModelLinksProjectStorages)`

SetProjectStorages sets ProjectStorages field to given value.

### HasProjectStorages

`func (o *ProjectModelLinks) HasProjectStorages() bool`

HasProjectStorages returns a boolean if a field has been set.

### GetAncestors

`func (o *ProjectModelLinks) GetAncestors() []ProjectModelLinksAncestorsInner`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ProjectModelLinks) GetAncestorsOk() (*[]ProjectModelLinksAncestorsInner, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ProjectModelLinks) SetAncestors(v []ProjectModelLinksAncestorsInner)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ProjectModelLinks) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


