# WidgetSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**ChannelId** | Pointer to **string** |  | [optional] 

## Methods

### NewWidgetSettingsResponse

`func NewWidgetSettingsResponse(enabled bool, ) *WidgetSettingsResponse`

NewWidgetSettingsResponse instantiates a new WidgetSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetSettingsResponseWithDefaults

`func NewWidgetSettingsResponseWithDefaults() *WidgetSettingsResponse`

NewWidgetSettingsResponseWithDefaults instantiates a new WidgetSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WidgetSettingsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WidgetSettingsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WidgetSettingsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetChannelId

`func (o *WidgetSettingsResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *WidgetSettingsResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *WidgetSettingsResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *WidgetSettingsResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


