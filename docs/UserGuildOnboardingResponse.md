# UserGuildOnboardingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuildId** | **string** |  | 
**Prompts** | [**[]OnboardingPromptResponse**](OnboardingPromptResponse.md) |  | 
**DefaultChannelIds** | **[]string** |  | 
**Enabled** | **bool** |  | 

## Methods

### NewUserGuildOnboardingResponse

`func NewUserGuildOnboardingResponse(guildId string, prompts []OnboardingPromptResponse, defaultChannelIds []string, enabled bool, ) *UserGuildOnboardingResponse`

NewUserGuildOnboardingResponse instantiates a new UserGuildOnboardingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGuildOnboardingResponseWithDefaults

`func NewUserGuildOnboardingResponseWithDefaults() *UserGuildOnboardingResponse`

NewUserGuildOnboardingResponseWithDefaults instantiates a new UserGuildOnboardingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuildId

`func (o *UserGuildOnboardingResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *UserGuildOnboardingResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *UserGuildOnboardingResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetPrompts

`func (o *UserGuildOnboardingResponse) GetPrompts() []OnboardingPromptResponse`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *UserGuildOnboardingResponse) GetPromptsOk() (*[]OnboardingPromptResponse, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *UserGuildOnboardingResponse) SetPrompts(v []OnboardingPromptResponse)`

SetPrompts sets Prompts field to given value.


### GetDefaultChannelIds

`func (o *UserGuildOnboardingResponse) GetDefaultChannelIds() []string`

GetDefaultChannelIds returns the DefaultChannelIds field if non-nil, zero value otherwise.

### GetDefaultChannelIdsOk

`func (o *UserGuildOnboardingResponse) GetDefaultChannelIdsOk() (*[]string, bool)`

GetDefaultChannelIdsOk returns a tuple with the DefaultChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannelIds

`func (o *UserGuildOnboardingResponse) SetDefaultChannelIds(v []string)`

SetDefaultChannelIds sets DefaultChannelIds field to given value.


### GetEnabled

`func (o *UserGuildOnboardingResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserGuildOnboardingResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserGuildOnboardingResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


