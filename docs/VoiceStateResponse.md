# VoiceStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **string** |  | [optional] 
**Deaf** | **bool** |  | 
**GuildId** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**NullableGuildMemberResponse**](GuildMemberResponse.md) |  | [optional] 
**Mute** | **bool** |  | 
**RequestToSpeakTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Suppress** | **bool** |  | 
**SelfStream** | Pointer to **NullableBool** |  | [optional] 
**SelfDeaf** | **bool** |  | 
**SelfMute** | **bool** |  | 
**SelfVideo** | **bool** |  | 
**SessionId** | **string** |  | 
**UserId** | **string** |  | 

## Methods

### NewVoiceStateResponse

`func NewVoiceStateResponse(deaf bool, mute bool, suppress bool, selfDeaf bool, selfMute bool, selfVideo bool, sessionId string, userId string, ) *VoiceStateResponse`

NewVoiceStateResponse instantiates a new VoiceStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceStateResponseWithDefaults

`func NewVoiceStateResponseWithDefaults() *VoiceStateResponse`

NewVoiceStateResponseWithDefaults instantiates a new VoiceStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *VoiceStateResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *VoiceStateResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *VoiceStateResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *VoiceStateResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetDeaf

`func (o *VoiceStateResponse) GetDeaf() bool`

GetDeaf returns the Deaf field if non-nil, zero value otherwise.

### GetDeafOk

`func (o *VoiceStateResponse) GetDeafOk() (*bool, bool)`

GetDeafOk returns a tuple with the Deaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaf

`func (o *VoiceStateResponse) SetDeaf(v bool)`

SetDeaf sets Deaf field to given value.


### GetGuildId

`func (o *VoiceStateResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *VoiceStateResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *VoiceStateResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *VoiceStateResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetMember

`func (o *VoiceStateResponse) GetMember() GuildMemberResponse`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *VoiceStateResponse) GetMemberOk() (*GuildMemberResponse, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *VoiceStateResponse) SetMember(v GuildMemberResponse)`

SetMember sets Member field to given value.

### HasMember

`func (o *VoiceStateResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### SetMemberNil

`func (o *VoiceStateResponse) SetMemberNil(b bool)`

 SetMemberNil sets the value for Member to be an explicit nil

### UnsetMember
`func (o *VoiceStateResponse) UnsetMember()`

UnsetMember ensures that no value is present for Member, not even an explicit nil
### GetMute

`func (o *VoiceStateResponse) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *VoiceStateResponse) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *VoiceStateResponse) SetMute(v bool)`

SetMute sets Mute field to given value.


### GetRequestToSpeakTimestamp

`func (o *VoiceStateResponse) GetRequestToSpeakTimestamp() time.Time`

GetRequestToSpeakTimestamp returns the RequestToSpeakTimestamp field if non-nil, zero value otherwise.

### GetRequestToSpeakTimestampOk

`func (o *VoiceStateResponse) GetRequestToSpeakTimestampOk() (*time.Time, bool)`

GetRequestToSpeakTimestampOk returns a tuple with the RequestToSpeakTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToSpeakTimestamp

`func (o *VoiceStateResponse) SetRequestToSpeakTimestamp(v time.Time)`

SetRequestToSpeakTimestamp sets RequestToSpeakTimestamp field to given value.

### HasRequestToSpeakTimestamp

`func (o *VoiceStateResponse) HasRequestToSpeakTimestamp() bool`

HasRequestToSpeakTimestamp returns a boolean if a field has been set.

### SetRequestToSpeakTimestampNil

`func (o *VoiceStateResponse) SetRequestToSpeakTimestampNil(b bool)`

 SetRequestToSpeakTimestampNil sets the value for RequestToSpeakTimestamp to be an explicit nil

### UnsetRequestToSpeakTimestamp
`func (o *VoiceStateResponse) UnsetRequestToSpeakTimestamp()`

UnsetRequestToSpeakTimestamp ensures that no value is present for RequestToSpeakTimestamp, not even an explicit nil
### GetSuppress

`func (o *VoiceStateResponse) GetSuppress() bool`

GetSuppress returns the Suppress field if non-nil, zero value otherwise.

### GetSuppressOk

`func (o *VoiceStateResponse) GetSuppressOk() (*bool, bool)`

GetSuppressOk returns a tuple with the Suppress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppress

`func (o *VoiceStateResponse) SetSuppress(v bool)`

SetSuppress sets Suppress field to given value.


### GetSelfStream

`func (o *VoiceStateResponse) GetSelfStream() bool`

GetSelfStream returns the SelfStream field if non-nil, zero value otherwise.

### GetSelfStreamOk

`func (o *VoiceStateResponse) GetSelfStreamOk() (*bool, bool)`

GetSelfStreamOk returns a tuple with the SelfStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfStream

`func (o *VoiceStateResponse) SetSelfStream(v bool)`

SetSelfStream sets SelfStream field to given value.

### HasSelfStream

`func (o *VoiceStateResponse) HasSelfStream() bool`

HasSelfStream returns a boolean if a field has been set.

### SetSelfStreamNil

`func (o *VoiceStateResponse) SetSelfStreamNil(b bool)`

 SetSelfStreamNil sets the value for SelfStream to be an explicit nil

### UnsetSelfStream
`func (o *VoiceStateResponse) UnsetSelfStream()`

UnsetSelfStream ensures that no value is present for SelfStream, not even an explicit nil
### GetSelfDeaf

`func (o *VoiceStateResponse) GetSelfDeaf() bool`

GetSelfDeaf returns the SelfDeaf field if non-nil, zero value otherwise.

### GetSelfDeafOk

`func (o *VoiceStateResponse) GetSelfDeafOk() (*bool, bool)`

GetSelfDeafOk returns a tuple with the SelfDeaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfDeaf

`func (o *VoiceStateResponse) SetSelfDeaf(v bool)`

SetSelfDeaf sets SelfDeaf field to given value.


### GetSelfMute

`func (o *VoiceStateResponse) GetSelfMute() bool`

GetSelfMute returns the SelfMute field if non-nil, zero value otherwise.

### GetSelfMuteOk

`func (o *VoiceStateResponse) GetSelfMuteOk() (*bool, bool)`

GetSelfMuteOk returns a tuple with the SelfMute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfMute

`func (o *VoiceStateResponse) SetSelfMute(v bool)`

SetSelfMute sets SelfMute field to given value.


### GetSelfVideo

`func (o *VoiceStateResponse) GetSelfVideo() bool`

GetSelfVideo returns the SelfVideo field if non-nil, zero value otherwise.

### GetSelfVideoOk

`func (o *VoiceStateResponse) GetSelfVideoOk() (*bool, bool)`

GetSelfVideoOk returns a tuple with the SelfVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfVideo

`func (o *VoiceStateResponse) SetSelfVideo(v bool)`

SetSelfVideo sets SelfVideo field to given value.


### GetSessionId

`func (o *VoiceStateResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *VoiceStateResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *VoiceStateResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetUserId

`func (o *VoiceStateResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VoiceStateResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VoiceStateResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


