# ConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumAttachmentFileSize** | Pointer to **int64** | The maximum allowed size of an attachment in Bytes | [optional] [readonly] 
**HostName** | Pointer to **string** | The host name configured for the system | [optional] [readonly] 
**PerPageOptions** | Pointer to **[]int64** | Page size steps to be offered in paginated list UI | [optional] 
**ActiveFeatureFlags** | Pointer to **[]string** | The list of all feature flags that are active | [optional] 

## Methods

### NewConfigurationModel

`func NewConfigurationModel() *ConfigurationModel`

NewConfigurationModel instantiates a new ConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationModelWithDefaults

`func NewConfigurationModelWithDefaults() *ConfigurationModel`

NewConfigurationModelWithDefaults instantiates a new ConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumAttachmentFileSize

`func (o *ConfigurationModel) GetMaximumAttachmentFileSize() int64`

GetMaximumAttachmentFileSize returns the MaximumAttachmentFileSize field if non-nil, zero value otherwise.

### GetMaximumAttachmentFileSizeOk

`func (o *ConfigurationModel) GetMaximumAttachmentFileSizeOk() (*int64, bool)`

GetMaximumAttachmentFileSizeOk returns a tuple with the MaximumAttachmentFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAttachmentFileSize

`func (o *ConfigurationModel) SetMaximumAttachmentFileSize(v int64)`

SetMaximumAttachmentFileSize sets MaximumAttachmentFileSize field to given value.

### HasMaximumAttachmentFileSize

`func (o *ConfigurationModel) HasMaximumAttachmentFileSize() bool`

HasMaximumAttachmentFileSize returns a boolean if a field has been set.

### GetHostName

`func (o *ConfigurationModel) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ConfigurationModel) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ConfigurationModel) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ConfigurationModel) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetPerPageOptions

`func (o *ConfigurationModel) GetPerPageOptions() []int64`

GetPerPageOptions returns the PerPageOptions field if non-nil, zero value otherwise.

### GetPerPageOptionsOk

`func (o *ConfigurationModel) GetPerPageOptionsOk() (*[]int64, bool)`

GetPerPageOptionsOk returns a tuple with the PerPageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPageOptions

`func (o *ConfigurationModel) SetPerPageOptions(v []int64)`

SetPerPageOptions sets PerPageOptions field to given value.

### HasPerPageOptions

`func (o *ConfigurationModel) HasPerPageOptions() bool`

HasPerPageOptions returns a boolean if a field has been set.

### GetActiveFeatureFlags

`func (o *ConfigurationModel) GetActiveFeatureFlags() []string`

GetActiveFeatureFlags returns the ActiveFeatureFlags field if non-nil, zero value otherwise.

### GetActiveFeatureFlagsOk

`func (o *ConfigurationModel) GetActiveFeatureFlagsOk() (*[]string, bool)`

GetActiveFeatureFlagsOk returns a tuple with the ActiveFeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFeatureFlags

`func (o *ConfigurationModel) SetActiveFeatureFlags(v []string)`

SetActiveFeatureFlags sets ActiveFeatureFlags field to given value.

### HasActiveFeatureFlags

`func (o *ConfigurationModel) HasActiveFeatureFlags() bool`

HasActiveFeatureFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


