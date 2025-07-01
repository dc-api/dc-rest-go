# KeywordRuleResponse

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
**TriggerMetadata** | [**KeywordTriggerMetadataResponse**](KeywordTriggerMetadataResponse.md) |  | 

## Methods

### NewKeywordRuleResponse

`func NewKeywordRuleResponse(id string, guildId string, creatorId string, name string, eventType int32, actions []DefaultKeywordRuleResponseActionsInner, triggerType int32, triggerMetadata KeywordTriggerMetadataResponse, ) *KeywordRuleResponse`

NewKeywordRuleResponse instantiates a new KeywordRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeywordRuleResponseWithDefaults

`func NewKeywordRuleResponseWithDefaults() *KeywordRuleResponse`

NewKeywordRuleResponseWithDefaults instantiates a new KeywordRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeywordRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeywordRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeywordRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGuildId

`func (o *KeywordRuleResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *KeywordRuleResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *KeywordRuleResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetCreatorId

`func (o *KeywordRuleResponse) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *KeywordRuleResponse) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *KeywordRuleResponse) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetName

`func (o *KeywordRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeywordRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeywordRuleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetEventType

`func (o *KeywordRuleResponse) GetEventType() int32`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *KeywordRuleResponse) GetEventTypeOk() (*int32, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *KeywordRuleResponse) SetEventType(v int32)`

SetEventType sets EventType field to given value.


### GetActions

`func (o *KeywordRuleResponse) GetActions() []DefaultKeywordRuleResponseActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *KeywordRuleResponse) GetActionsOk() (*[]DefaultKeywordRuleResponseActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *KeywordRuleResponse) SetActions(v []DefaultKeywordRuleResponseActionsInner)`

SetActions sets Actions field to given value.


### GetTriggerType

`func (o *KeywordRuleResponse) GetTriggerType() int32`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *KeywordRuleResponse) GetTriggerTypeOk() (*int32, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *KeywordRuleResponse) SetTriggerType(v int32)`

SetTriggerType sets TriggerType field to given value.


### GetEnabled

`func (o *KeywordRuleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeywordRuleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeywordRuleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeywordRuleResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *KeywordRuleResponse) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *KeywordRuleResponse) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetExemptRoles

`func (o *KeywordRuleResponse) GetExemptRoles() []string`

GetExemptRoles returns the ExemptRoles field if non-nil, zero value otherwise.

### GetExemptRolesOk

`func (o *KeywordRuleResponse) GetExemptRolesOk() (*[]string, bool)`

GetExemptRolesOk returns a tuple with the ExemptRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptRoles

`func (o *KeywordRuleResponse) SetExemptRoles(v []string)`

SetExemptRoles sets ExemptRoles field to given value.

### HasExemptRoles

`func (o *KeywordRuleResponse) HasExemptRoles() bool`

HasExemptRoles returns a boolean if a field has been set.

### SetExemptRolesNil

`func (o *KeywordRuleResponse) SetExemptRolesNil(b bool)`

 SetExemptRolesNil sets the value for ExemptRoles to be an explicit nil

### UnsetExemptRoles
`func (o *KeywordRuleResponse) UnsetExemptRoles()`

UnsetExemptRoles ensures that no value is present for ExemptRoles, not even an explicit nil
### GetExemptChannels

`func (o *KeywordRuleResponse) GetExemptChannels() []string`

GetExemptChannels returns the ExemptChannels field if non-nil, zero value otherwise.

### GetExemptChannelsOk

`func (o *KeywordRuleResponse) GetExemptChannelsOk() (*[]string, bool)`

GetExemptChannelsOk returns a tuple with the ExemptChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptChannels

`func (o *KeywordRuleResponse) SetExemptChannels(v []string)`

SetExemptChannels sets ExemptChannels field to given value.

### HasExemptChannels

`func (o *KeywordRuleResponse) HasExemptChannels() bool`

HasExemptChannels returns a boolean if a field has been set.

### SetExemptChannelsNil

`func (o *KeywordRuleResponse) SetExemptChannelsNil(b bool)`

 SetExemptChannelsNil sets the value for ExemptChannels to be an explicit nil

### UnsetExemptChannels
`func (o *KeywordRuleResponse) UnsetExemptChannels()`

UnsetExemptChannels ensures that no value is present for ExemptChannels, not even an explicit nil
### GetTriggerMetadata

`func (o *KeywordRuleResponse) GetTriggerMetadata() KeywordTriggerMetadataResponse`

GetTriggerMetadata returns the TriggerMetadata field if non-nil, zero value otherwise.

### GetTriggerMetadataOk

`func (o *KeywordRuleResponse) GetTriggerMetadataOk() (*KeywordTriggerMetadataResponse, bool)`

GetTriggerMetadataOk returns a tuple with the TriggerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetadata

`func (o *KeywordRuleResponse) SetTriggerMetadata(v KeywordTriggerMetadataResponse)`

SetTriggerMetadata sets TriggerMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


