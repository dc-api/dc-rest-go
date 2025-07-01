# UpdateGuildWidgetSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateGuildWidgetSettingsRequest

`func NewUpdateGuildWidgetSettingsRequest() *UpdateGuildWidgetSettingsRequest`

NewUpdateGuildWidgetSettingsRequest instantiates a new UpdateGuildWidgetSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGuildWidgetSettingsRequestWithDefaults

`func NewUpdateGuildWidgetSettingsRequestWithDefaults() *UpdateGuildWidgetSettingsRequest`

NewUpdateGuildWidgetSettingsRequestWithDefaults instantiates a new UpdateGuildWidgetSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *UpdateGuildWidgetSettingsRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UpdateGuildWidgetSettingsRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UpdateGuildWidgetSettingsRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *UpdateGuildWidgetSettingsRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateGuildWidgetSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateGuildWidgetSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateGuildWidgetSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateGuildWidgetSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UpdateGuildWidgetSettingsRequest) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateGuildWidgetSettingsRequest) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


