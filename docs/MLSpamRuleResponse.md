# MLSpamRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GuildId** | **string** |  | 
**CreatorId** | **string** |  | 
**Name** | **string** |  | 
**EventType** | **int32** |  | 
**Actions** | [**[]DefaultKeywordRuleResponseActionsInner**](DefaultKeywordRuleResponseActionsInner.md) |  | 
**TriggerType** | **int32** |  | 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**ExemptRoles** | Pointer to **[]string** |  | [optional] 
**ExemptChannels** | Pointer to **[]string** |  | [optional] 
**TriggerMetadata** | **map[string]interface{}** |  | 

## Methods

### NewMLSpamRuleResponse

`func NewMLSpamRuleResponse(id string, guildId string, creatorId string, name string, eventType int32, actions []DefaultKeywordRuleResponseActionsInner, triggerType int32, triggerMetadata map[string]interface{}, ) *MLSpamRuleResponse`

NewMLSpamRuleResponse instantiates a new MLSpamRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMLSpamRuleResponseWithDefaults

`func NewMLSpamRuleResponseWithDefaults() *MLSpamRuleResponse`

NewMLSpamRuleResponseWithDefaults instantiates a new MLSpamRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MLSpamRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MLSpamRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MLSpamRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGuildId

`func (o *MLSpamRuleResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *MLSpamRuleResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *MLSpamRuleResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetCreatorId

`func (o *MLSpamRuleResponse) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *MLSpamRuleResponse) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *MLSpamRuleResponse) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetName

`func (o *MLSpamRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MLSpamRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MLSpamRuleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetEventType

`func (o *MLSpamRuleResponse) GetEventType() int32`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MLSpamRuleResponse) GetEventTypeOk() (*int32, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MLSpamRuleResponse) SetEventType(v int32)`

SetEventType sets EventType field to given value.


### GetActions

`func (o *MLSpamRuleResponse) GetActions() []DefaultKeywordRuleResponseActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MLSpamRuleResponse) GetActionsOk() (*[]DefaultKeywordRuleResponseActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MLSpamRuleResponse) SetActions(v []DefaultKeywordRuleResponseActionsInner)`

SetActions sets Actions field to given value.


### GetTriggerType

`func (o *MLSpamRuleResponse) GetTriggerType() int32`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *MLSpamRuleResponse) GetTriggerTypeOk() (*int32, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *MLSpamRuleResponse) SetTriggerType(v int32)`

SetTriggerType sets TriggerType field to given value.


### GetEnabled

`func (o *MLSpamRuleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MLSpamRuleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MLSpamRuleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MLSpamRuleResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *MLSpamRuleResponse) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *MLSpamRuleResponse) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetExemptRoles

`func (o *MLSpamRuleResponse) GetExemptRoles() []string`

GetExemptRoles returns the ExemptRoles field if non-nil, zero value otherwise.

### GetExemptRolesOk

`func (o *MLSpamRuleResponse) GetExemptRolesOk() (*[]string, bool)`

GetExemptRolesOk returns a tuple with the ExemptRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptRoles

`func (o *MLSpamRuleResponse) SetExemptRoles(v []string)`

SetExemptRoles sets ExemptRoles field to given value.

### HasExemptRoles

`func (o *MLSpamRuleResponse) HasExemptRoles() bool`

HasExemptRoles returns a boolean if a field has been set.

### SetExemptRolesNil

`func (o *MLSpamRuleResponse) SetExemptRolesNil(b bool)`

 SetExemptRolesNil sets the value for ExemptRoles to be an explicit nil

### UnsetExemptRoles
`func (o *MLSpamRuleResponse) UnsetExemptRoles()`

UnsetExemptRoles ensures that no value is present for ExemptRoles, not even an explicit nil
### GetExemptChannels

`func (o *MLSpamRuleResponse) GetExemptChannels() []string`

GetExemptChannels returns the ExemptChannels field if non-nil, zero value otherwise.

### GetExemptChannelsOk

`func (o *MLSpamRuleResponse) GetExemptChannelsOk() (*[]string, bool)`

GetExemptChannelsOk returns a tuple with the ExemptChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptChannels

`func (o *MLSpamRuleResponse) SetExemptChannels(v []string)`

SetExemptChannels sets ExemptChannels field to given value.

### HasExemptChannels

`func (o *MLSpamRuleResponse) HasExemptChannels() bool`

HasExemptChannels returns a boolean if a field has been set.

### SetExemptChannelsNil

`func (o *MLSpamRuleResponse) SetExemptChannelsNil(b bool)`

 SetExemptChannelsNil sets the value for ExemptChannels to be an explicit nil

### UnsetExemptChannels
`func (o *MLSpamRuleResponse) UnsetExemptChannels()`

UnsetExemptChannels ensures that no value is present for ExemptChannels, not even an explicit nil
### GetTriggerMetadata

`func (o *MLSpamRuleResponse) GetTriggerMetadata() map[string]interface{}`

GetTriggerMetadata returns the TriggerMetadata field if non-nil, zero value otherwise.

### GetTriggerMetadataOk

`func (o *MLSpamRuleResponse) GetTriggerMetadataOk() (*map[string]interface{}, bool)`

GetTriggerMetadataOk returns a tuple with the TriggerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetadata

`func (o *MLSpamRuleResponse) SetTriggerMetadata(v map[string]interface{})`

SetTriggerMetadata sets TriggerMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


