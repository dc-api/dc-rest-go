# MentionSpamRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GuildId** | **string** |  | 
**CreatorId** | **string** |  | 
**Name** | **string** |  | 
**EventType** | [**AutomodEventType**](AutomodEventType.md) |  | 
**Actions** | [**[]DefaultKeywordRuleResponseActionsInner**](DefaultKeywordRuleResponseActionsInner.md) |  | 
**TriggerType** | [**AutomodTriggerType**](AutomodTriggerType.md) |  | 
**Enabled** | **bool** |  | 
**ExemptRoles** | **[]string** |  | 
**ExemptChannels** | **[]string** |  | 
**TriggerMetadata** | [**MentionSpamTriggerMetadataResponse**](MentionSpamTriggerMetadataResponse.md) |  | 

## Methods

### NewMentionSpamRuleResponse

`func NewMentionSpamRuleResponse(id string, guildId string, creatorId string, name string, eventType AutomodEventType, actions []DefaultKeywordRuleResponseActionsInner, triggerType AutomodTriggerType, enabled bool, exemptRoles []string, exemptChannels []string, triggerMetadata MentionSpamTriggerMetadataResponse, ) *MentionSpamRuleResponse`

NewMentionSpamRuleResponse instantiates a new MentionSpamRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMentionSpamRuleResponseWithDefaults

`func NewMentionSpamRuleResponseWithDefaults() *MentionSpamRuleResponse`

NewMentionSpamRuleResponseWithDefaults instantiates a new MentionSpamRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MentionSpamRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MentionSpamRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MentionSpamRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGuildId

`func (o *MentionSpamRuleResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *MentionSpamRuleResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *MentionSpamRuleResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetCreatorId

`func (o *MentionSpamRuleResponse) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *MentionSpamRuleResponse) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *MentionSpamRuleResponse) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetName

`func (o *MentionSpamRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MentionSpamRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MentionSpamRuleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetEventType

`func (o *MentionSpamRuleResponse) GetEventType() AutomodEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MentionSpamRuleResponse) GetEventTypeOk() (*AutomodEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MentionSpamRuleResponse) SetEventType(v AutomodEventType)`

SetEventType sets EventType field to given value.


### GetActions

`func (o *MentionSpamRuleResponse) GetActions() []DefaultKeywordRuleResponseActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MentionSpamRuleResponse) GetActionsOk() (*[]DefaultKeywordRuleResponseActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MentionSpamRuleResponse) SetActions(v []DefaultKeywordRuleResponseActionsInner)`

SetActions sets Actions field to given value.


### GetTriggerType

`func (o *MentionSpamRuleResponse) GetTriggerType() AutomodTriggerType`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *MentionSpamRuleResponse) GetTriggerTypeOk() (*AutomodTriggerType, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *MentionSpamRuleResponse) SetTriggerType(v AutomodTriggerType)`

SetTriggerType sets TriggerType field to given value.


### GetEnabled

`func (o *MentionSpamRuleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MentionSpamRuleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MentionSpamRuleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExemptRoles

`func (o *MentionSpamRuleResponse) GetExemptRoles() []string`

GetExemptRoles returns the ExemptRoles field if non-nil, zero value otherwise.

### GetExemptRolesOk

`func (o *MentionSpamRuleResponse) GetExemptRolesOk() (*[]string, bool)`

GetExemptRolesOk returns a tuple with the ExemptRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptRoles

`func (o *MentionSpamRuleResponse) SetExemptRoles(v []string)`

SetExemptRoles sets ExemptRoles field to given value.


### GetExemptChannels

`func (o *MentionSpamRuleResponse) GetExemptChannels() []string`

GetExemptChannels returns the ExemptChannels field if non-nil, zero value otherwise.

### GetExemptChannelsOk

`func (o *MentionSpamRuleResponse) GetExemptChannelsOk() (*[]string, bool)`

GetExemptChannelsOk returns a tuple with the ExemptChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptChannels

`func (o *MentionSpamRuleResponse) SetExemptChannels(v []string)`

SetExemptChannels sets ExemptChannels field to given value.


### GetTriggerMetadata

`func (o *MentionSpamRuleResponse) GetTriggerMetadata() MentionSpamTriggerMetadataResponse`

GetTriggerMetadata returns the TriggerMetadata field if non-nil, zero value otherwise.

### GetTriggerMetadataOk

`func (o *MentionSpamRuleResponse) GetTriggerMetadataOk() (*MentionSpamTriggerMetadataResponse, bool)`

GetTriggerMetadataOk returns a tuple with the TriggerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetadata

`func (o *MentionSpamRuleResponse) SetTriggerMetadata(v MentionSpamTriggerMetadataResponse)`

SetTriggerMetadata sets TriggerMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


