# WorkPackageActivitiesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Total** | **int64** | The total amount of elements available in the collection. | 
**Count** | **int64** | Actual amount of elements in this response. | 
**Links** | [**CollectionModelLinks**](CollectionModelLinks.md) |  | 
**Embedded** | Pointer to [**WorkPackageActivitiesModelAllOfEmbedded**](WorkPackageActivitiesModelAllOfEmbedded.md) |  | [optional] 

## Methods

### NewWorkPackageActivitiesModel

`func NewWorkPackageActivitiesModel(type_ string, total int64, count int64, links CollectionModelLinks, ) *WorkPackageActivitiesModel`

NewWorkPackageActivitiesModel instantiates a new WorkPackageActivitiesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkPackageActivitiesModelWithDefaults

`func NewWorkPackageActivitiesModelWithDefaults() *WorkPackageActivitiesModel`

NewWorkPackageActivitiesModelWithDefaults instantiates a new WorkPackageActivitiesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WorkPackageActivitiesModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkPackageActivitiesModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkPackageActivitiesModel) SetType(v string)`

SetType sets Type field to given value.


### GetTotal

`func (o *WorkPackageActivitiesModel) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WorkPackageActivitiesModel) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WorkPackageActivitiesModel) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCount

`func (o *WorkPackageActivitiesModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WorkPackageActivitiesModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WorkPackageActivitiesModel) SetCount(v int64)`

SetCount sets Count field to given value.


### GetLinks

`func (o *WorkPackageActivitiesModel) GetLinks() CollectionModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkPackageActivitiesModel) GetLinksOk() (*CollectionModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkPackageActivitiesModel) SetLinks(v CollectionModelLinks)`

SetLinks sets Links field to given value.


### GetEmbedded

`func (o *WorkPackageActivitiesModel) GetEmbedded() WorkPackageActivitiesModelAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *WorkPackageActivitiesModel) GetEmbeddedOk() (*WorkPackageActivitiesModelAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *WorkPackageActivitiesModel) SetEmbedded(v WorkPackageActivitiesModelAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *WorkPackageActivitiesModel) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


