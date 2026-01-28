# UpdateSelfVoiceStateRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestToSpeakTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Suppress** | Pointer to **NullableBool** |  | [optional] 
**ChannelId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateSelfVoiceStateRequestPartial

`func NewUpdateSelfVoiceStateRequestPartial() *UpdateSelfVoiceStateRequestPartial`

NewUpdateSelfVoiceStateRequestPartial instantiates a new UpdateSelfVoiceStateRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSelfVoiceStateRequestPartialWithDefaults

`func NewUpdateSelfVoiceStateRequestPartialWithDefaults() *UpdateSelfVoiceStateRequestPartial`

NewUpdateSelfVoiceStateRequestPartialWithDefaults instantiates a new UpdateSelfVoiceStateRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestToSpeakTimestamp

`func (o *UpdateSelfVoiceStateRequestPartial) GetRequestToSpeakTimestamp() time.Time`

GetRequestToSpeakTimestamp returns the RequestToSpeakTimestamp field if non-nil, zero value otherwise.

### GetRequestToSpeakTimestampOk

`func (o *UpdateSelfVoiceStateRequestPartial) GetRequestToSpeakTimestampOk() (*time.Time, bool)`

GetRequestToSpeakTimestampOk returns a tuple with the RequestToSpeakTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToSpeakTimestamp

`func (o *UpdateSelfVoiceStateRequestPartial) SetRequestToSpeakTimestamp(v time.Time)`

SetRequestToSpeakTimestamp sets RequestToSpeakTimestamp field to given value.

### HasRequestToSpeakTimestamp

`func (o *UpdateSelfVoiceStateRequestPartial) HasRequestToSpeakTimestamp() bool`

HasRequestToSpeakTimestamp returns a boolean if a field has been set.

### SetRequestToSpeakTimestampNil

`func (o *UpdateSelfVoiceStateRequestPartial) SetRequestToSpeakTimestampNil(b bool)`

 SetRequestToSpeakTimestampNil sets the value for RequestToSpeakTimestamp to be an explicit nil

### UnsetRequestToSpeakTimestamp
`func (o *UpdateSelfVoiceStateRequestPartial) UnsetRequestToSpeakTimestamp()`

UnsetRequestToSpeakTimestamp ensures that no value is present for RequestToSpeakTimestamp, not even an explicit nil
### GetSuppress

`func (o *UpdateSelfVoiceStateRequestPartial) GetSuppress() bool`

GetSuppress returns the Suppress field if non-nil, zero value otherwise.

### GetSuppressOk

`func (o *UpdateSelfVoiceStateRequestPartial) GetSuppressOk() (*bool, bool)`

GetSuppressOk returns a tuple with the Suppress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppress

`func (o *UpdateSelfVoiceStateRequestPartial) SetSuppress(v bool)`

SetSuppress sets Suppress field to given value.

### HasSuppress

`func (o *UpdateSelfVoiceStateRequestPartial) HasSuppress() bool`

HasSuppress returns a boolean if a field has been set.

### SetSuppressNil

`func (o *UpdateSelfVoiceStateRequestPartial) SetSuppressNil(b bool)`

 SetSuppressNil sets the value for Suppress to be an explicit nil

### UnsetSuppress
`func (o *UpdateSelfVoiceStateRequestPartial) UnsetSuppress()`

UnsetSuppress ensures that no value is present for Suppress, not even an explicit nil
### GetChannelId

`func (o *UpdateSelfVoiceStateRequestPartial) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UpdateSelfVoiceStateRequestPartial) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UpdateSelfVoiceStateRequestPartial) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *UpdateSelfVoiceStateRequestPartial) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *UpdateSelfVoiceStateRequestPartial) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *UpdateSelfVoiceStateRequestPartial) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


