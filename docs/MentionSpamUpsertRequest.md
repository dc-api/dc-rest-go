# MentionSpamUpsertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EventType** | **int32** |  | 
**Actions** | Pointer to [**[]DefaultKeywordListUpsertRequestActionsInner**](DefaultKeywordListUpsertRequestActionsInner.md) |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**ExemptRoles** | Pointer to **[]string** |  | [optional] 
**ExemptChannels** | Pointer to **[]string** |  | [optional] 
**TriggerType** | **int32** |  | 
**TriggerMetadata** | Pointer to [**NullableMentionSpamTriggerMetadata**](MentionSpamTriggerMetadata.md) |  | [optional] 

## Methods

### NewMentionSpamUpsertRequest

`func NewMentionSpamUpsertRequest(name string, eventType int32, triggerType int32, ) *MentionSpamUpsertRequest`

NewMentionSpamUpsertRequest instantiates a new MentionSpamUpsertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMentionSpamUpsertRequestWithDefaults

`func NewMentionSpamUpsertRequestWithDefaults() *MentionSpamUpsertRequest`

NewMentionSpamUpsertRequestWithDefaults instantiates a new MentionSpamUpsertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MentionSpamUpsertRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MentionSpamUpsertRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MentionSpamUpsertRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEventType

`func (o *MentionSpamUpsertRequest) GetEventType() int32`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MentionSpamUpsertRequest) GetEventTypeOk() (*int32, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MentionSpamUpsertRequest) SetEventType(v int32)`

SetEventType sets EventType field to given value.


### GetActions

`func (o *MentionSpamUpsertRequest) GetActions() []DefaultKeywordListUpsertRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MentionSpamUpsertRequest) GetActionsOk() (*[]DefaultKeywordListUpsertRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MentionSpamUpsertRequest) SetActions(v []DefaultKeywordListUpsertRequestActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MentionSpamUpsertRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *MentionSpamUpsertRequest) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *MentionSpamUpsertRequest) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetEnabled

`func (o *MentionSpamUpsertRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MentionSpamUpsertRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MentionSpamUpsertRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MentionSpamUpsertRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *MentionSpamUpsertRequest) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *MentionSpamUpsertRequest) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetExemptRoles

`func (o *MentionSpamUpsertRequest) GetExemptRoles() []string`

GetExemptRoles returns the ExemptRoles field if non-nil, zero value otherwise.

### GetExemptRolesOk

`func (o *MentionSpamUpsertRequest) GetExemptRolesOk() (*[]string, bool)`

GetExemptRolesOk returns a tuple with the ExemptRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptRoles

`func (o *MentionSpamUpsertRequest) SetExemptRoles(v []string)`

SetExemptRoles sets ExemptRoles field to given value.

### HasExemptRoles

`func (o *MentionSpamUpsertRequest) HasExemptRoles() bool`

HasExemptRoles returns a boolean if a field has been set.

### SetExemptRolesNil

`func (o *MentionSpamUpsertRequest) SetExemptRolesNil(b bool)`

 SetExemptRolesNil sets the value for ExemptRoles to be an explicit nil

### UnsetExemptRoles
`func (o *MentionSpamUpsertRequest) UnsetExemptRoles()`

UnsetExemptRoles ensures that no value is present for ExemptRoles, not even an explicit nil
### GetExemptChannels

`func (o *MentionSpamUpsertRequest) GetExemptChannels() []string`

GetExemptChannels returns the ExemptChannels field if non-nil, zero value otherwise.

### GetExemptChannelsOk

`func (o *MentionSpamUpsertRequest) GetExemptChannelsOk() (*[]string, bool)`

GetExemptChannelsOk returns a tuple with the ExemptChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptChannels

`func (o *MentionSpamUpsertRequest) SetExemptChannels(v []string)`

SetExemptChannels sets ExemptChannels field to given value.

### HasExemptChannels

`func (o *MentionSpamUpsertRequest) HasExemptChannels() bool`

HasExemptChannels returns a boolean if a field has been set.

### SetExemptChannelsNil

`func (o *MentionSpamUpsertRequest) SetExemptChannelsNil(b bool)`

 SetExemptChannelsNil sets the value for ExemptChannels to be an explicit nil

### UnsetExemptChannels
`func (o *MentionSpamUpsertRequest) UnsetExemptChannels()`

UnsetExemptChannels ensures that no value is present for ExemptChannels, not even an explicit nil
### GetTriggerType

`func (o *MentionSpamUpsertRequest) GetTriggerType() int32`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *MentionSpamUpsertRequest) GetTriggerTypeOk() (*int32, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *MentionSpamUpsertRequest) SetTriggerType(v int32)`

SetTriggerType sets TriggerType field to given value.


### GetTriggerMetadata

`func (o *MentionSpamUpsertRequest) GetTriggerMetadata() MentionSpamTriggerMetadata`

GetTriggerMetadata returns the TriggerMetadata field if non-nil, zero value otherwise.

### GetTriggerMetadataOk

`func (o *MentionSpamUpsertRequest) GetTriggerMetadataOk() (*MentionSpamTriggerMetadata, bool)`

GetTriggerMetadataOk returns a tuple with the TriggerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetadata

`func (o *MentionSpamUpsertRequest) SetTriggerMetadata(v MentionSpamTriggerMetadata)`

SetTriggerMetadata sets TriggerMetadata field to given value.

### HasTriggerMetadata

`func (o *MentionSpamUpsertRequest) HasTriggerMetadata() bool`

HasTriggerMetadata returns a boolean if a field has been set.

### SetTriggerMetadataNil

`func (o *MentionSpamUpsertRequest) SetTriggerMetadataNil(b bool)`

 SetTriggerMetadataNil sets the value for TriggerMetadata to be an explicit nil

### UnsetTriggerMetadata
`func (o *MentionSpamUpsertRequest) UnsetTriggerMetadata()`

UnsetTriggerMetadata ensures that no value is present for TriggerMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


