# UpdateSelfVoiceStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestToSpeakTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Suppress** | Pointer to **NullableBool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateSelfVoiceStateRequest

`func NewUpdateSelfVoiceStateRequest() *UpdateSelfVoiceStateRequest`

NewUpdateSelfVoiceStateRequest instantiates a new UpdateSelfVoiceStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSelfVoiceStateRequestWithDefaults

`func NewUpdateSelfVoiceStateRequestWithDefaults() *UpdateSelfVoiceStateRequest`

NewUpdateSelfVoiceStateRequestWithDefaults instantiates a new UpdateSelfVoiceStateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestToSpeakTimestamp

`func (o *UpdateSelfVoiceStateRequest) GetRequestToSpeakTimestamp() time.Time`

GetRequestToSpeakTimestamp returns the RequestToSpeakTimestamp field if non-nil, zero value otherwise.

### GetRequestToSpeakTimestampOk

`func (o *UpdateSelfVoiceStateRequest) GetRequestToSpeakTimestampOk() (*time.Time, bool)`

GetRequestToSpeakTimestampOk returns a tuple with the RequestToSpeakTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToSpeakTimestamp

`func (o *UpdateSelfVoiceStateRequest) SetRequestToSpeakTimestamp(v time.Time)`

SetRequestToSpeakTimestamp sets RequestToSpeakTimestamp field to given value.

### HasRequestToSpeakTimestamp

`func (o *UpdateSelfVoiceStateRequest) HasRequestToSpeakTimestamp() bool`

HasRequestToSpeakTimestamp returns a boolean if a field has been set.

### SetRequestToSpeakTimestampNil

`func (o *UpdateSelfVoiceStateRequest) SetRequestToSpeakTimestampNil(b bool)`

 SetRequestToSpeakTimestampNil sets the value for RequestToSpeakTimestamp to be an explicit nil

### UnsetRequestToSpeakTimestamp
`func (o *UpdateSelfVoiceStateRequest) UnsetRequestToSpeakTimestamp()`

UnsetRequestToSpeakTimestamp ensures that no value is present for RequestToSpeakTimestamp, not even an explicit nil
### GetSuppress

`func (o *UpdateSelfVoiceStateRequest) GetSuppress() bool`

GetSuppress returns the Suppress field if non-nil, zero value otherwise.

### GetSuppressOk

`func (o *UpdateSelfVoiceStateRequest) GetSuppressOk() (*bool, bool)`

GetSuppressOk returns a tuple with the Suppress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppress

`func (o *UpdateSelfVoiceStateRequest) SetSuppress(v bool)`

SetSuppress sets Suppress field to given value.

### HasSuppress

`func (o *UpdateSelfVoiceStateRequest) HasSuppress() bool`

HasSuppress returns a boolean if a field has been set.

### SetSuppressNil

`func (o *UpdateSelfVoiceStateRequest) SetSuppressNil(b bool)`

 SetSuppressNil sets the value for Suppress to be an explicit nil

### UnsetSuppress
`func (o *UpdateSelfVoiceStateRequest) UnsetSuppress()`

UnsetSuppress ensures that no value is present for Suppress, not even an explicit nil
### GetChannelId

`func (o *UpdateSelfVoiceStateRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UpdateSelfVoiceStateRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UpdateSelfVoiceStateRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *UpdateSelfVoiceStateRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


