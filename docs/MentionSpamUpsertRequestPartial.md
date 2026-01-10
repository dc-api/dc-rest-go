# MentionSpamUpsertRequestPartial

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
**TriggerMetadata** | Pointer to [**MentionSpamUpsertRequestTriggerMetadata**](MentionSpamUpsertRequestTriggerMetadata.md) |  | [optional] 

## Methods

### NewMentionSpamUpsertRequestPartial

`func NewMentionSpamUpsertRequestPartial() *MentionSpamUpsertRequestPartial`

NewMentionSpamUpsertRequestPartial instantiates a new MentionSpamUpsertRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMentionSpamUpsertRequestPartialWithDefaults

`func NewMentionSpamUpsertRequestPartialWithDefaults() *MentionSpamUpsertRequestPartial`

NewMentionSpamUpsertRequestPartialWithDefaults instantiates a new MentionSpamUpsertRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MentionSpamUpsertRequestPartial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MentionSpamUpsertRequestPartial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MentionSpamUpsertRequestPartial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MentionSpamUpsertRequestPartial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEventType

`func (o *MentionSpamUpsertRequestPartial) GetEventType() AutomodEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MentionSpamUpsertRequestPartial) GetEventTypeOk() (*AutomodEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MentionSpamUpsertRequestPartial) SetEventType(v AutomodEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MentionSpamUpsertRequestPartial) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetActions

`func (o *MentionSpamUpsertRequestPartial) GetActions() []DefaultKeywordListUpsertRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MentionSpamUpsertRequestPartial) GetActionsOk() (*[]DefaultKeywordListUpsertRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MentionSpamUpsertRequestPartial) SetActions(v []DefaultKeywordListUpsertRequestActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MentionSpamUpsertRequestPartial) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetEnabled

`func (o *MentionSpamUpsertRequestPartial) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MentionSpamUpsertRequestPartial) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MentionSpamUpsertRequestPartial) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MentionSpamUpsertRequestPartial) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExemptRoles

`func (o *MentionSpamUpsertRequestPartial) GetExemptRoles() []string`

GetExemptRoles returns the ExemptRoles field if non-nil, zero value otherwise.

### GetExemptRolesOk

`func (o *MentionSpamUpsertRequestPartial) GetExemptRolesOk() (*[]string, bool)`

GetExemptRolesOk returns a tuple with the ExemptRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptRoles

`func (o *MentionSpamUpsertRequestPartial) SetExemptRoles(v []string)`

SetExemptRoles sets ExemptRoles field to given value.

### HasExemptRoles

`func (o *MentionSpamUpsertRequestPartial) HasExemptRoles() bool`

HasExemptRoles returns a boolean if a field has been set.

### GetExemptChannels

`func (o *MentionSpamUpsertRequestPartial) GetExemptChannels() []string`

GetExemptChannels returns the ExemptChannels field if non-nil, zero value otherwise.

### GetExemptChannelsOk

`func (o *MentionSpamUpsertRequestPartial) GetExemptChannelsOk() (*[]string, bool)`

GetExemptChannelsOk returns a tuple with the ExemptChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptChannels

`func (o *MentionSpamUpsertRequestPartial) SetExemptChannels(v []string)`

SetExemptChannels sets ExemptChannels field to given value.

### HasExemptChannels

`func (o *MentionSpamUpsertRequestPartial) HasExemptChannels() bool`

HasExemptChannels returns a boolean if a field has been set.

### GetTriggerType

`func (o *MentionSpamUpsertRequestPartial) GetTriggerType() AutomodTriggerType`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *MentionSpamUpsertRequestPartial) GetTriggerTypeOk() (*AutomodTriggerType, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *MentionSpamUpsertRequestPartial) SetTriggerType(v AutomodTriggerType)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *MentionSpamUpsertRequestPartial) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetTriggerMetadata

`func (o *MentionSpamUpsertRequestPartial) GetTriggerMetadata() MentionSpamUpsertRequestTriggerMetadata`

GetTriggerMetadata returns the TriggerMetadata field if non-nil, zero value otherwise.

### GetTriggerMetadataOk

`func (o *MentionSpamUpsertRequestPartial) GetTriggerMetadataOk() (*MentionSpamUpsertRequestTriggerMetadata, bool)`

GetTriggerMetadataOk returns a tuple with the TriggerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetadata

`func (o *MentionSpamUpsertRequestPartial) SetTriggerMetadata(v MentionSpamUpsertRequestTriggerMetadata)`

SetTriggerMetadata sets TriggerMetadata field to given value.

### HasTriggerMetadata

`func (o *MentionSpamUpsertRequestPartial) HasTriggerMetadata() bool`

HasTriggerMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


