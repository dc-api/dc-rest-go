# UpdateGuildOnboardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompts** | Pointer to [**[]UpdateOnboardingPromptRequest**](UpdateOnboardingPromptRequest.md) |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**DefaultChannelIds** | Pointer to **[]string** |  | [optional] 
**Mode** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateGuildOnboardingRequest

`func NewUpdateGuildOnboardingRequest() *UpdateGuildOnboardingRequest`

NewUpdateGuildOnboardingRequest instantiates a new UpdateGuildOnboardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGuildOnboardingRequestWithDefaults

`func NewUpdateGuildOnboardingRequestWithDefaults() *UpdateGuildOnboardingRequest`

NewUpdateGuildOnboardingRequestWithDefaults instantiates a new UpdateGuildOnboardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompts

`func (o *UpdateGuildOnboardingRequest) GetPrompts() []UpdateOnboardingPromptRequest`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *UpdateGuildOnboardingRequest) GetPromptsOk() (*[]UpdateOnboardingPromptRequest, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *UpdateGuildOnboardingRequest) SetPrompts(v []UpdateOnboardingPromptRequest)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *UpdateGuildOnboardingRequest) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### SetPromptsNil

`func (o *UpdateGuildOnboardingRequest) SetPromptsNil(b bool)`

 SetPromptsNil sets the value for Prompts to be an explicit nil

### UnsetPrompts
`func (o *UpdateGuildOnboardingRequest) UnsetPrompts()`

UnsetPrompts ensures that no value is present for Prompts, not even an explicit nil
### GetEnabled

`func (o *UpdateGuildOnboardingRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateGuildOnboardingRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateGuildOnboardingRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateGuildOnboardingRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UpdateGuildOnboardingRequest) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateGuildOnboardingRequest) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetDefaultChannelIds

`func (o *UpdateGuildOnboardingRequest) GetDefaultChannelIds() []string`

GetDefaultChannelIds returns the DefaultChannelIds field if non-nil, zero value otherwise.

### GetDefaultChannelIdsOk

`func (o *UpdateGuildOnboardingRequest) GetDefaultChannelIdsOk() (*[]string, bool)`

GetDefaultChannelIdsOk returns a tuple with the DefaultChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannelIds

`func (o *UpdateGuildOnboardingRequest) SetDefaultChannelIds(v []string)`

SetDefaultChannelIds sets DefaultChannelIds field to given value.

### HasDefaultChannelIds

`func (o *UpdateGuildOnboardingRequest) HasDefaultChannelIds() bool`

HasDefaultChannelIds returns a boolean if a field has been set.

### SetDefaultChannelIdsNil

`func (o *UpdateGuildOnboardingRequest) SetDefaultChannelIdsNil(b bool)`

 SetDefaultChannelIdsNil sets the value for DefaultChannelIds to be an explicit nil

### UnsetDefaultChannelIds
`func (o *UpdateGuildOnboardingRequest) UnsetDefaultChannelIds()`

UnsetDefaultChannelIds ensures that no value is present for DefaultChannelIds, not even an explicit nil
### GetMode

`func (o *UpdateGuildOnboardingRequest) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateGuildOnboardingRequest) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateGuildOnboardingRequest) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateGuildOnboardingRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


