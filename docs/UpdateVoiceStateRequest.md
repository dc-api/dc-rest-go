# UpdateVoiceStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suppress** | Pointer to **NullableBool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateVoiceStateRequest

`func NewUpdateVoiceStateRequest() *UpdateVoiceStateRequest`

NewUpdateVoiceStateRequest instantiates a new UpdateVoiceStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVoiceStateRequestWithDefaults

`func NewUpdateVoiceStateRequestWithDefaults() *UpdateVoiceStateRequest`

NewUpdateVoiceStateRequestWithDefaults instantiates a new UpdateVoiceStateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuppress

`func (o *UpdateVoiceStateRequest) GetSuppress() bool`

GetSuppress returns the Suppress field if non-nil, zero value otherwise.

### GetSuppressOk

`func (o *UpdateVoiceStateRequest) GetSuppressOk() (*bool, bool)`

GetSuppressOk returns a tuple with the Suppress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppress

`func (o *UpdateVoiceStateRequest) SetSuppress(v bool)`

SetSuppress sets Suppress field to given value.

### HasSuppress

`func (o *UpdateVoiceStateRequest) HasSuppress() bool`

HasSuppress returns a boolean if a field has been set.

### SetSuppressNil

`func (o *UpdateVoiceStateRequest) SetSuppressNil(b bool)`

 SetSuppressNil sets the value for Suppress to be an explicit nil

### UnsetSuppress
`func (o *UpdateVoiceStateRequest) UnsetSuppress()`

UnsetSuppress ensures that no value is present for Suppress, not even an explicit nil
### GetChannelId

`func (o *UpdateVoiceStateRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UpdateVoiceStateRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UpdateVoiceStateRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *UpdateVoiceStateRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


