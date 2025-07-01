# GuildHomeSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuildId** | **string** |  | 
**Enabled** | **bool** |  | 
**WelcomeMessage** | Pointer to [**NullableWelcomeMessageResponse**](WelcomeMessageResponse.md) |  | [optional] 
**NewMemberActions** | Pointer to [**[]NewMemberActionResponse**](NewMemberActionResponse.md) |  | [optional] 
**ResourceChannels** | Pointer to [**[]ResourceChannelResponse**](ResourceChannelResponse.md) |  | [optional] 

## Methods

### NewGuildHomeSettingsResponse

`func NewGuildHomeSettingsResponse(guildId string, enabled bool, ) *GuildHomeSettingsResponse`

NewGuildHomeSettingsResponse instantiates a new GuildHomeSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildHomeSettingsResponseWithDefaults

`func NewGuildHomeSettingsResponseWithDefaults() *GuildHomeSettingsResponse`

NewGuildHomeSettingsResponseWithDefaults instantiates a new GuildHomeSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuildId

`func (o *GuildHomeSettingsResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GuildHomeSettingsResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GuildHomeSettingsResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetEnabled

`func (o *GuildHomeSettingsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GuildHomeSettingsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GuildHomeSettingsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetWelcomeMessage

`func (o *GuildHomeSettingsResponse) GetWelcomeMessage() WelcomeMessageResponse`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *GuildHomeSettingsResponse) GetWelcomeMessageOk() (*WelcomeMessageResponse, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *GuildHomeSettingsResponse) SetWelcomeMessage(v WelcomeMessageResponse)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *GuildHomeSettingsResponse) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.

### SetWelcomeMessageNil

`func (o *GuildHomeSettingsResponse) SetWelcomeMessageNil(b bool)`

 SetWelcomeMessageNil sets the value for WelcomeMessage to be an explicit nil

### UnsetWelcomeMessage
`func (o *GuildHomeSettingsResponse) UnsetWelcomeMessage()`

UnsetWelcomeMessage ensures that no value is present for WelcomeMessage, not even an explicit nil
### GetNewMemberActions

`func (o *GuildHomeSettingsResponse) GetNewMemberActions() []*NewMemberActionResponse`

GetNewMemberActions returns the NewMemberActions field if non-nil, zero value otherwise.

### GetNewMemberActionsOk

`func (o *GuildHomeSettingsResponse) GetNewMemberActionsOk() (*[]*NewMemberActionResponse, bool)`

GetNewMemberActionsOk returns a tuple with the NewMemberActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMemberActions

`func (o *GuildHomeSettingsResponse) SetNewMemberActions(v []*NewMemberActionResponse)`

SetNewMemberActions sets NewMemberActions field to given value.

### HasNewMemberActions

`func (o *GuildHomeSettingsResponse) HasNewMemberActions() bool`

HasNewMemberActions returns a boolean if a field has been set.

### SetNewMemberActionsNil

`func (o *GuildHomeSettingsResponse) SetNewMemberActionsNil(b bool)`

 SetNewMemberActionsNil sets the value for NewMemberActions to be an explicit nil

### UnsetNewMemberActions
`func (o *GuildHomeSettingsResponse) UnsetNewMemberActions()`

UnsetNewMemberActions ensures that no value is present for NewMemberActions, not even an explicit nil
### GetResourceChannels

`func (o *GuildHomeSettingsResponse) GetResourceChannels() []*ResourceChannelResponse`

GetResourceChannels returns the ResourceChannels field if non-nil, zero value otherwise.

### GetResourceChannelsOk

`func (o *GuildHomeSettingsResponse) GetResourceChannelsOk() (*[]*ResourceChannelResponse, bool)`

GetResourceChannelsOk returns a tuple with the ResourceChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceChannels

`func (o *GuildHomeSettingsResponse) SetResourceChannels(v []*ResourceChannelResponse)`

SetResourceChannels sets ResourceChannels field to given value.

### HasResourceChannels

`func (o *GuildHomeSettingsResponse) HasResourceChannels() bool`

HasResourceChannels returns a boolean if a field has been set.

### SetResourceChannelsNil

`func (o *GuildHomeSettingsResponse) SetResourceChannelsNil(b bool)`

 SetResourceChannelsNil sets the value for ResourceChannels to be an explicit nil

### UnsetResourceChannels
`func (o *GuildHomeSettingsResponse) UnsetResourceChannels()`

UnsetResourceChannels ensures that no value is present for ResourceChannels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


