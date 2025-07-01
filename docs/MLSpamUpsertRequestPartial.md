# MLSpamUpsertRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **int32** |  | [optional] 
**Actions** | Pointer to [**[]DefaultKeywordListUpsertRequestActionsInner**](DefaultKeywordListUpsertRequestActionsInner.md) |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**ExemptRoles** | Pointer to **[]string** |  | [optional] 
**ExemptChannels** | Pointer to **[]string** |  | [optional] 
**TriggerType** | Pointer to **int32** |  | [optional] 
**TriggerMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMLSpamUpsertRequestPartial

`func NewMLSpamUpsertRequestPartial() *MLSpamUpsertRequestPartial`

NewMLSpamUpsertRequestPartial instantiates a new MLSpamUpsertRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMLSpamUpsertRequestPartialWithDefaults

`func NewMLSpamUpsertRequestPartialWithDefaults() *MLSpamUpsertRequestPartial`

NewMLSpamUpsertRequestPartialWithDefaults instantiates a new MLSpamUpsertRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MLSpamUpsertRequestPartial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MLSpamUpsertRequestPartial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MLSpamUpsertRequestPartial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MLSpamUpsertRequestPartial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEventType

`func (o *MLSpamUpsertRequestPartial) GetEventType() int32`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MLSpamUpsertRequestPartial) GetEventTypeOk() (*int32, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MLSpamUpsertRequestPartial) SetEventType(v int32)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MLSpamUpsertRequestPartial) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetActions

`func (o *MLSpamUpsertRequestPartial) GetActions() []DefaultKeywordListUpsertRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MLSpamUpsertRequestPartial) GetActionsOk() (*[]DefaultKeywordListUpsertRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MLSpamUpsertRequestPartial) SetActions(v []DefaultKeywordListUpsertRequestActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MLSpamUpsertRequestPartial) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *MLSpamUpsertRequestPartial) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *MLSpamUpsertRequestPartial) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetEnabled

`func (o *MLSpamUpsertRequestPartial) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MLSpamUpsertRequestPartial) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MLSpamUpsertRequestPartial) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MLSpamUpsertRequestPartial) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *MLSpamUpsertRequestPartial) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *MLSpamUpsertRequestPartial) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetExemptRoles

`func (o *MLSpamUpsertRequestPartial) GetExemptRoles() []string`

GetExemptRoles returns the ExemptRoles field if non-nil, zero value otherwise.

### GetExemptRolesOk

`func (o *MLSpamUpsertRequestPartial) GetExemptRolesOk() (*[]string, bool)`

GetExemptRolesOk returns a tuple with the ExemptRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptRoles

`func (o *MLSpamUpsertRequestPartial) SetExemptRoles(v []string)`

SetExemptRoles sets ExemptRoles field to given value.

### HasExemptRoles

`func (o *MLSpamUpsertRequestPartial) HasExemptRoles() bool`

HasExemptRoles returns a boolean if a field has been set.

### SetExemptRolesNil

`func (o *MLSpamUpsertRequestPartial) SetExemptRolesNil(b bool)`

 SetExemptRolesNil sets the value for ExemptRoles to be an explicit nil

### UnsetExemptRoles
`func (o *MLSpamUpsertRequestPartial) UnsetExemptRoles()`

UnsetExemptRoles ensures that no value is present for ExemptRoles, not even an explicit nil
### GetExemptChannels

`func (o *MLSpamUpsertRequestPartial) GetExemptChannels() []string`

GetExemptChannels returns the ExemptChannels field if non-nil, zero value otherwise.

### GetExemptChannelsOk

`func (o *MLSpamUpsertRequestPartial) GetExemptChannelsOk() (*[]string, bool)`

GetExemptChannelsOk returns a tuple with the ExemptChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptChannels

`func (o *MLSpamUpsertRequestPartial) SetExemptChannels(v []string)`

SetExemptChannels sets ExemptChannels field to given value.

### HasExemptChannels

`func (o *MLSpamUpsertRequestPartial) HasExemptChannels() bool`

HasExemptChannels returns a boolean if a field has been set.

### SetExemptChannelsNil

`func (o *MLSpamUpsertRequestPartial) SetExemptChannelsNil(b bool)`

 SetExemptChannelsNil sets the value for ExemptChannels to be an explicit nil

### UnsetExemptChannels
`func (o *MLSpamUpsertRequestPartial) UnsetExemptChannels()`

UnsetExemptChannels ensures that no value is present for ExemptChannels, not even an explicit nil
### GetTriggerType

`func (o *MLSpamUpsertRequestPartial) GetTriggerType() int32`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *MLSpamUpsertRequestPartial) GetTriggerTypeOk() (*int32, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *MLSpamUpsertRequestPartial) SetTriggerType(v int32)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *MLSpamUpsertRequestPartial) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetTriggerMetadata

`func (o *MLSpamUpsertRequestPartial) GetTriggerMetadata() map[string]interface{}`

GetTriggerMetadata returns the TriggerMetadata field if non-nil, zero value otherwise.

### GetTriggerMetadataOk

`func (o *MLSpamUpsertRequestPartial) GetTriggerMetadataOk() (*map[string]interface{}, bool)`

GetTriggerMetadataOk returns a tuple with the TriggerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetadata

`func (o *MLSpamUpsertRequestPartial) SetTriggerMetadata(v map[string]interface{})`

SetTriggerMetadata sets TriggerMetadata field to given value.

### HasTriggerMetadata

`func (o *MLSpamUpsertRequestPartial) HasTriggerMetadata() bool`

HasTriggerMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


