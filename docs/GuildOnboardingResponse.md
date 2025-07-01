# GuildOnboardingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuildId** | **string** |  | 
**Prompts** | [**[]OnboardingPromptResponse**](OnboardingPromptResponse.md) |  | 
**DefaultChannelIds** | **[]string** |  | 
**Enabled** | **bool** |  | 

## Methods

### NewGuildOnboardingResponse

`func NewGuildOnboardingResponse(guildId string, prompts []OnboardingPromptResponse, defaultChannelIds []string, enabled bool, ) *GuildOnboardingResponse`

NewGuildOnboardingResponse instantiates a new GuildOnboardingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildOnboardingResponseWithDefaults

`func NewGuildOnboardingResponseWithDefaults() *GuildOnboardingResponse`

NewGuildOnboardingResponseWithDefaults instantiates a new GuildOnboardingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuildId

`func (o *GuildOnboardingResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GuildOnboardingResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GuildOnboardingResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetPrompts

`func (o *GuildOnboardingResponse) GetPrompts() []OnboardingPromptResponse`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *GuildOnboardingResponse) GetPromptsOk() (*[]OnboardingPromptResponse, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *GuildOnboardingResponse) SetPrompts(v []OnboardingPromptResponse)`

SetPrompts sets Prompts field to given value.


### GetDefaultChannelIds

`func (o *GuildOnboardingResponse) GetDefaultChannelIds() []string`

GetDefaultChannelIds returns the DefaultChannelIds field if non-nil, zero value otherwise.

### GetDefaultChannelIdsOk

`func (o *GuildOnboardingResponse) GetDefaultChannelIdsOk() (*[]string, bool)`

GetDefaultChannelIdsOk returns a tuple with the DefaultChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannelIds

`func (o *GuildOnboardingResponse) SetDefaultChannelIds(v []string)`

SetDefaultChannelIds sets DefaultChannelIds field to given value.


### GetEnabled

`func (o *GuildOnboardingResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GuildOnboardingResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GuildOnboardingResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


