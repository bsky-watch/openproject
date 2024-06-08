# MembershipCollectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Total** | **int64** | The total amount of elements available in the collection. | 
**Count** | **int64** | Actual amount of elements in this response. | 
**Links** | [**PaginatedCollectionModelAllOfLinks**](PaginatedCollectionModelAllOfLinks.md) |  | 
**PageSize** | **int64** | Amount of elements that a response will hold. | 
**Offset** | **int64** | The page number that is requested from paginated collection. | 
**Embedded** | [**MembershipCollectionModelAllOfEmbedded**](MembershipCollectionModelAllOfEmbedded.md) |  | 

## Methods

### NewMembershipCollectionModel

`func NewMembershipCollectionModel(type_ string, total int64, count int64, links PaginatedCollectionModelAllOfLinks, pageSize int64, offset int64, embedded MembershipCollectionModelAllOfEmbedded, ) *MembershipCollectionModel`

NewMembershipCollectionModel instantiates a new MembershipCollectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipCollectionModelWithDefaults

`func NewMembershipCollectionModelWithDefaults() *MembershipCollectionModel`

NewMembershipCollectionModelWithDefaults instantiates a new MembershipCollectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MembershipCollectionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MembershipCollectionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MembershipCollectionModel) SetType(v string)`

SetType sets Type field to given value.


### GetTotal

`func (o *MembershipCollectionModel) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MembershipCollectionModel) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MembershipCollectionModel) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCount

`func (o *MembershipCollectionModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MembershipCollectionModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MembershipCollectionModel) SetCount(v int64)`

SetCount sets Count field to given value.


### GetLinks

`func (o *MembershipCollectionModel) GetLinks() PaginatedCollectionModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MembershipCollectionModel) GetLinksOk() (*PaginatedCollectionModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MembershipCollectionModel) SetLinks(v PaginatedCollectionModelAllOfLinks)`

SetLinks sets Links field to given value.


### GetPageSize

`func (o *MembershipCollectionModel) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *MembershipCollectionModel) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *MembershipCollectionModel) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetOffset

`func (o *MembershipCollectionModel) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MembershipCollectionModel) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MembershipCollectionModel) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetEmbedded

`func (o *MembershipCollectionModel) GetEmbedded() MembershipCollectionModelAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *MembershipCollectionModel) GetEmbeddedOk() (*MembershipCollectionModelAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *MembershipCollectionModel) SetEmbedded(v MembershipCollectionModelAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


