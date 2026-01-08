# UpdateGuildOnboardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompts** | Pointer to [**[]UpdateOnboardingPromptRequest**](UpdateOnboardingPromptRequest.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DefaultChannelIds** | Pointer to **[]string** |  | [optional] 
**Mode** | Pointer to [**UpdateGuildOnboardingRequestMode**](UpdateGuildOnboardingRequestMode.md) |  | [optional] 

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

### GetMode

`func (o *UpdateGuildOnboardingRequest) GetMode() UpdateGuildOnboardingRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateGuildOnboardingRequest) GetModeOk() (*UpdateGuildOnboardingRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateGuildOnboardingRequest) SetMode(v UpdateGuildOnboardingRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateGuildOnboardingRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


