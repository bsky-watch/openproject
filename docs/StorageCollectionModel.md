# StorageCollectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Total** | **int64** | The total amount of elements available in the collection. | 
**Count** | **int64** | Actual amount of elements in this response. | 
**Links** | [**StorageCollectionModelAllOfLinks**](StorageCollectionModelAllOfLinks.md) |  | 
**Embedded** | [**StorageCollectionModelAllOfEmbedded**](StorageCollectionModelAllOfEmbedded.md) |  | 

## Methods

### NewStorageCollectionModel

`func NewStorageCollectionModel(type_ string, total int64, count int64, links StorageCollectionModelAllOfLinks, embedded StorageCollectionModelAllOfEmbedded, ) *StorageCollectionModel`

NewStorageCollectionModel instantiates a new StorageCollectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageCollectionModelWithDefaults

`func NewStorageCollectionModelWithDefaults() *StorageCollectionModel`

NewStorageCollectionModelWithDefaults instantiates a new StorageCollectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StorageCollectionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageCollectionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageCollectionModel) SetType(v string)`

SetType sets Type field to given value.


### GetTotal

`func (o *StorageCollectionModel) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StorageCollectionModel) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StorageCollectionModel) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCount

`func (o *StorageCollectionModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StorageCollectionModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StorageCollectionModel) SetCount(v int64)`

SetCount sets Count field to given value.


### GetLinks

`func (o *StorageCollectionModel) GetLinks() StorageCollectionModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageCollectionModel) GetLinksOk() (*StorageCollectionModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageCollectionModel) SetLinks(v StorageCollectionModelAllOfLinks)`

SetLinks sets Links field to given value.


### GetEmbedded

`func (o *StorageCollectionModel) GetEmbedded() StorageCollectionModelAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *StorageCollectionModel) GetEmbeddedOk() (*StorageCollectionModelAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *StorageCollectionModel) SetEmbedded(v StorageCollectionModelAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


