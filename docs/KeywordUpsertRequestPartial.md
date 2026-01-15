# KeywordUpsertRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to [**AutomodEventType**](AutomodEventType.md) |  | [optional] 
**Actions** | Pointer to [**[]DefaultKeywordListUpsertRequestActionsInner**](DefaultKeywordListUpsertRequestActionsInner.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExemptRoles** | Pointer to **[]string** |  | [optional] 
**ExemptChannels** | Pointer to **[]string** |  | [optional] 
**TriggerType** | Pointer to [**AutomodTriggerType**](AutomodTriggerType.md) |  | [optional] 
**TriggerMetadata** | Pointer to [**KeywordUpsertRequestTriggerMetadata**](KeywordUpsertRequestTriggerMetadata.md) |  | [optional] 

## Methods

### NewKeywordUpsertRequestPartial

`func NewKeywordUpsertRequestPartial() *KeywordUpsertRequestPartial`

NewKeywordUpsertRequestPartial instantiates a new KeywordUpsertRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeywordUpsertRequestPartialWithDefaults

`func NewKeywordUpsertRequestPartialWithDefaults() *KeywordUpsertRequestPartial`

NewKeywordUpsertRequestPartialWithDefaults instantiates a new KeywordUpsertRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeywordUpsertRequestPartial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeywordUpsertRequestPartial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeywordUpsertRequestPartial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeywordUpsertRequestPartial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEventType

`func (o *KeywordUpsertRequestPartial) GetEventType() AutomodEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *KeywordUpsertRequestPartial) GetEventTypeOk() (*AutomodEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *KeywordUpsertRequestPartial) SetEventType(v AutomodEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *KeywordUpsertRequestPartial) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetActions

`func (o *KeywordUpsertRequestPartial) GetActions() []DefaultKeywordListUpsertRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *KeywordUpsertRequestPartial) GetActionsOk() (*[]DefaultKeywordListUpsertRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *KeywordUpsertRequestPartial) SetActions(v []DefaultKeywordListUpsertRequestActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *KeywordUpsertRequestPartial) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetEnabled

`func (o *KeywordUpsertRequestPartial) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeywordUpsertRequestPartial) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeywordUpsertRequestPartial) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeywordUpsertRequestPartial) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExemptRoles

`func (o *KeywordUpsertRequestPartial) GetExemptRoles() []string`

GetExemptRoles returns the ExemptRoles field if non-nil, zero value otherwise.

### GetExemptRolesOk

`func (o *KeywordUpsertRequestPartial) GetExemptRolesOk() (*[]string, bool)`

GetExemptRolesOk returns a tuple with the ExemptRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptRoles

`func (o *KeywordUpsertRequestPartial) SetExemptRoles(v []string)`

SetExemptRoles sets ExemptRoles field to given value.

### HasExemptRoles

`func (o *KeywordUpsertRequestPartial) HasExemptRoles() bool`

HasExemptRoles returns a boolean if a field has been set.

### GetExemptChannels

`func (o *KeywordUpsertRequestPartial) GetExemptChannels() []string`

GetExemptChannels returns the ExemptChannels field if non-nil, zero value otherwise.

### GetExemptChannelsOk

`func (o *KeywordUpsertRequestPartial) GetExemptChannelsOk() (*[]string, bool)`

GetExemptChannelsOk returns a tuple with the ExemptChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptChannels

`func (o *KeywordUpsertRequestPartial) SetExemptChannels(v []string)`

SetExemptChannels sets ExemptChannels field to given value.

### HasExemptChannels

`func (o *KeywordUpsertRequestPartial) HasExemptChannels() bool`

HasExemptChannels returns a boolean if a field has been set.

### GetTriggerType

`func (o *KeywordUpsertRequestPartial) GetTriggerType() AutomodTriggerType`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *KeywordUpsertRequestPartial) GetTriggerTypeOk() (*AutomodTriggerType, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *KeywordUpsertRequestPartial) SetTriggerType(v AutomodTriggerType)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *KeywordUpsertRequestPartial) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetTriggerMetadata

`func (o *KeywordUpsertRequestPartial) GetTriggerMetadata() KeywordUpsertRequestTriggerMetadata`

GetTriggerMetadata returns the TriggerMetadata field if non-nil, zero value otherwise.

### GetTriggerMetadataOk

`func (o *KeywordUpsertRequestPartial) GetTriggerMetadataOk() (*KeywordUpsertRequestTriggerMetadata, bool)`

GetTriggerMetadataOk returns a tuple with the TriggerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetadata

`func (o *KeywordUpsertRequestPartial) SetTriggerMetadata(v KeywordUpsertRequestTriggerMetadata)`

SetTriggerMetadata sets TriggerMetadata field to given value.

### HasTriggerMetadata

`func (o *KeywordUpsertRequestPartial) HasTriggerMetadata() bool`

HasTriggerMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


