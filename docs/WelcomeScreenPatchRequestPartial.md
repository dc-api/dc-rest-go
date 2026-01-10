# WelcomeScreenPatchRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**WelcomeChannels** | Pointer to [**[]GuildWelcomeChannel**](GuildWelcomeChannel.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewWelcomeScreenPatchRequestPartial

`func NewWelcomeScreenPatchRequestPartial() *WelcomeScreenPatchRequestPartial`

NewWelcomeScreenPatchRequestPartial instantiates a new WelcomeScreenPatchRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWelcomeScreenPatchRequestPartialWithDefaults

`func NewWelcomeScreenPatchRequestPartialWithDefaults() *WelcomeScreenPatchRequestPartial`

NewWelcomeScreenPatchRequestPartialWithDefaults instantiates a new WelcomeScreenPatchRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WelcomeScreenPatchRequestPartial) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WelcomeScreenPatchRequestPartial) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WelcomeScreenPatchRequestPartial) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WelcomeScreenPatchRequestPartial) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWelcomeChannels

`func (o *WelcomeScreenPatchRequestPartial) GetWelcomeChannels() []GuildWelcomeChannel`

GetWelcomeChannels returns the WelcomeChannels field if non-nil, zero value otherwise.

### GetWelcomeChannelsOk

`func (o *WelcomeScreenPatchRequestPartial) GetWelcomeChannelsOk() (*[]GuildWelcomeChannel, bool)`

GetWelcomeChannelsOk returns a tuple with the WelcomeChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeChannels

`func (o *WelcomeScreenPatchRequestPartial) SetWelcomeChannels(v []GuildWelcomeChannel)`

SetWelcomeChannels sets WelcomeChannels field to given value.

### HasWelcomeChannels

`func (o *WelcomeScreenPatchRequestPartial) HasWelcomeChannels() bool`

HasWelcomeChannels returns a boolean if a field has been set.

### GetEnabled

`func (o *WelcomeScreenPatchRequestPartial) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WelcomeScreenPatchRequestPartial) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WelcomeScreenPatchRequestPartial) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WelcomeScreenPatchRequestPartial) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


