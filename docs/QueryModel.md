# QueryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Query id | [optional] [readonly] 
**Name** | Pointer to **string** | Query name | [optional] [readonly] 
**Filters** | Pointer to [**[]QueryFilterInstanceSchemaModel**](QueryFilterInstanceSchemaModel.md) | A set of QueryFilters which will be applied to the work packages to determine the resulting work packages | [optional] 
**Sums** | Pointer to **bool** | Should sums (of supported properties) be shown? | [optional] [readonly] 
**TimelineVisible** | Pointer to **bool** | Should the timeline mode be shown? | [optional] [readonly] 
**TimelineLabels** | Pointer to **[]string** | Which labels are shown in the timeline, empty when default | [optional] 
**TimelineZoomLevel** | Pointer to **string** | Which zoom level should the timeline be rendered in? | [optional] [readonly] 
**Timestamps** | Pointer to **[]string** | Timestamps to filter by when showing changed attributes on work packages. | [optional] 
**HighlightingMode** | Pointer to **string** | Which highlighting mode should the table have? | [optional] [readonly] 
**ShowHierarchies** | Pointer to **bool** | Should the hierarchy mode be enabled? | [optional] [readonly] 
**Hidden** | Pointer to **bool** | Should the query be hidden from the query list? | [optional] [readonly] 
**Public** | Pointer to **bool** | Can users besides the owner see the query? | [optional] [readonly] 
**Starred** | Pointer to **bool** | Should the query be highlighted to the user? | [optional] [readonly] 
**CreatedAt** | **time.Time** | Time of creation | [readonly] 
**UpdatedAt** | **time.Time** | Time of the most recent change to the query | [readonly] 
**Links** | Pointer to [**QueryModelLinks**](QueryModelLinks.md) |  | [optional] 

## Methods

### NewQueryModel

`func NewQueryModel(createdAt time.Time, updatedAt time.Time, ) *QueryModel`

NewQueryModel instantiates a new QueryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryModelWithDefaults

`func NewQueryModelWithDefaults() *QueryModel`

NewQueryModelWithDefaults instantiates a new QueryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueryModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QueryModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *QueryModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueryModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilters

`func (o *QueryModel) GetFilters() []QueryFilterInstanceSchemaModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryModel) GetFiltersOk() (*[]QueryFilterInstanceSchemaModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryModel) SetFilters(v []QueryFilterInstanceSchemaModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryModel) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSums

`func (o *QueryModel) GetSums() bool`

GetSums returns the Sums field if non-nil, zero value otherwise.

### GetSumsOk

`func (o *QueryModel) GetSumsOk() (*bool, bool)`

GetSumsOk returns a tuple with the Sums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSums

`func (o *QueryModel) SetSums(v bool)`

SetSums sets Sums field to given value.

### HasSums

`func (o *QueryModel) HasSums() bool`

HasSums returns a boolean if a field has been set.

### GetTimelineVisible

`func (o *QueryModel) GetTimelineVisible() bool`

GetTimelineVisible returns the TimelineVisible field if non-nil, zero value otherwise.

### GetTimelineVisibleOk

`func (o *QueryModel) GetTimelineVisibleOk() (*bool, bool)`

GetTimelineVisibleOk returns a tuple with the TimelineVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineVisible

`func (o *QueryModel) SetTimelineVisible(v bool)`

SetTimelineVisible sets TimelineVisible field to given value.

### HasTimelineVisible

`func (o *QueryModel) HasTimelineVisible() bool`

HasTimelineVisible returns a boolean if a field has been set.

### GetTimelineLabels

`func (o *QueryModel) GetTimelineLabels() []string`

GetTimelineLabels returns the TimelineLabels field if non-nil, zero value otherwise.

### GetTimelineLabelsOk

`func (o *QueryModel) GetTimelineLabelsOk() (*[]string, bool)`

GetTimelineLabelsOk returns a tuple with the TimelineLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineLabels

`func (o *QueryModel) SetTimelineLabels(v []string)`

SetTimelineLabels sets TimelineLabels field to given value.

### HasTimelineLabels

`func (o *QueryModel) HasTimelineLabels() bool`

HasTimelineLabels returns a boolean if a field has been set.

### GetTimelineZoomLevel

`func (o *QueryModel) GetTimelineZoomLevel() string`

GetTimelineZoomLevel returns the TimelineZoomLevel field if non-nil, zero value otherwise.

### GetTimelineZoomLevelOk

`func (o *QueryModel) GetTimelineZoomLevelOk() (*string, bool)`

GetTimelineZoomLevelOk returns a tuple with the TimelineZoomLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineZoomLevel

`func (o *QueryModel) SetTimelineZoomLevel(v string)`

SetTimelineZoomLevel sets TimelineZoomLevel field to given value.

### HasTimelineZoomLevel

`func (o *QueryModel) HasTimelineZoomLevel() bool`

HasTimelineZoomLevel returns a boolean if a field has been set.

### GetTimestamps

`func (o *QueryModel) GetTimestamps() []string`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *QueryModel) GetTimestampsOk() (*[]string, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *QueryModel) SetTimestamps(v []string)`

SetTimestamps sets Timestamps field to given value.

### HasTimestamps

`func (o *QueryModel) HasTimestamps() bool`

HasTimestamps returns a boolean if a field has been set.

### GetHighlightingMode

`func (o *QueryModel) GetHighlightingMode() string`

GetHighlightingMode returns the HighlightingMode field if non-nil, zero value otherwise.

### GetHighlightingModeOk

`func (o *QueryModel) GetHighlightingModeOk() (*string, bool)`

GetHighlightingModeOk returns a tuple with the HighlightingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightingMode

`func (o *QueryModel) SetHighlightingMode(v string)`

SetHighlightingMode sets HighlightingMode field to given value.

### HasHighlightingMode

`func (o *QueryModel) HasHighlightingMode() bool`

HasHighlightingMode returns a boolean if a field has been set.

### GetShowHierarchies

`func (o *QueryModel) GetShowHierarchies() bool`

GetShowHierarchies returns the ShowHierarchies field if non-nil, zero value otherwise.

### GetShowHierarchiesOk

`func (o *QueryModel) GetShowHierarchiesOk() (*bool, bool)`

GetShowHierarchiesOk returns a tuple with the ShowHierarchies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHierarchies

`func (o *QueryModel) SetShowHierarchies(v bool)`

SetShowHierarchies sets ShowHierarchies field to given value.

### HasShowHierarchies

`func (o *QueryModel) HasShowHierarchies() bool`

HasShowHierarchies returns a boolean if a field has been set.

### GetHidden

`func (o *QueryModel) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *QueryModel) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *QueryModel) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *QueryModel) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPublic

`func (o *QueryModel) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *QueryModel) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *QueryModel) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *QueryModel) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetStarred

`func (o *QueryModel) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *QueryModel) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *QueryModel) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *QueryModel) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QueryModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QueryModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QueryModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QueryModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QueryModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QueryModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *QueryModel) GetLinks() QueryModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *QueryModel) GetLinksOk() (*QueryModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *QueryModel) SetLinks(v QueryModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *QueryModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


