# UpdateGuildMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nick** | Pointer to **NullableString** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Mute** | Pointer to **NullableBool** |  | [optional] 
**Deaf** | Pointer to **NullableBool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CommunicationDisabledUntil** | Pointer to **NullableTime** |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUpdateGuildMemberRequest

`func NewUpdateGuildMemberRequest() *UpdateGuildMemberRequest`

NewUpdateGuildMemberRequest instantiates a new UpdateGuildMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGuildMemberRequestWithDefaults

`func NewUpdateGuildMemberRequestWithDefaults() *UpdateGuildMemberRequest`

NewUpdateGuildMemberRequestWithDefaults instantiates a new UpdateGuildMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNick

`func (o *UpdateGuildMemberRequest) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *UpdateGuildMemberRequest) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *UpdateGuildMemberRequest) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *UpdateGuildMemberRequest) HasNick() bool`

HasNick returns a boolean if a field has been set.

### SetNickNil

`func (o *UpdateGuildMemberRequest) SetNickNil(b bool)`

 SetNickNil sets the value for Nick to be an explicit nil

### UnsetNick
`func (o *UpdateGuildMemberRequest) UnsetNick()`

UnsetNick ensures that no value is present for Nick, not even an explicit nil
### GetRoles

`func (o *UpdateGuildMemberRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateGuildMemberRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateGuildMemberRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateGuildMemberRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UpdateGuildMemberRequest) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UpdateGuildMemberRequest) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetMute

`func (o *UpdateGuildMemberRequest) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *UpdateGuildMemberRequest) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *UpdateGuildMemberRequest) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *UpdateGuildMemberRequest) HasMute() bool`

HasMute returns a boolean if a field has been set.

### SetMuteNil

`func (o *UpdateGuildMemberRequest) SetMuteNil(b bool)`

 SetMuteNil sets the value for Mute to be an explicit nil

### UnsetMute
`func (o *UpdateGuildMemberRequest) UnsetMute()`

UnsetMute ensures that no value is present for Mute, not even an explicit nil
### GetDeaf

`func (o *UpdateGuildMemberRequest) GetDeaf() bool`

GetDeaf returns the Deaf field if non-nil, zero value otherwise.

### GetDeafOk

`func (o *UpdateGuildMemberRequest) GetDeafOk() (*bool, bool)`

GetDeafOk returns a tuple with the Deaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaf

`func (o *UpdateGuildMemberRequest) SetDeaf(v bool)`

SetDeaf sets Deaf field to given value.

### HasDeaf

`func (o *UpdateGuildMemberRequest) HasDeaf() bool`

HasDeaf returns a boolean if a field has been set.

### SetDeafNil

`func (o *UpdateGuildMemberRequest) SetDeafNil(b bool)`

 SetDeafNil sets the value for Deaf to be an explicit nil

### UnsetDeaf
`func (o *UpdateGuildMemberRequest) UnsetDeaf()`

UnsetDeaf ensures that no value is present for Deaf, not even an explicit nil
### GetChannelId

`func (o *UpdateGuildMemberRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UpdateGuildMemberRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UpdateGuildMemberRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *UpdateGuildMemberRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCommunicationDisabledUntil

`func (o *UpdateGuildMemberRequest) GetCommunicationDisabledUntil() time.Time`

GetCommunicationDisabledUntil returns the CommunicationDisabledUntil field if non-nil, zero value otherwise.

### GetCommunicationDisabledUntilOk

`func (o *UpdateGuildMemberRequest) GetCommunicationDisabledUntilOk() (*time.Time, bool)`

GetCommunicationDisabledUntilOk returns a tuple with the CommunicationDisabledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDisabledUntil

`func (o *UpdateGuildMemberRequest) SetCommunicationDisabledUntil(v time.Time)`

SetCommunicationDisabledUntil sets CommunicationDisabledUntil field to given value.

### HasCommunicationDisabledUntil

`func (o *UpdateGuildMemberRequest) HasCommunicationDisabledUntil() bool`

HasCommunicationDisabledUntil returns a boolean if a field has been set.

### SetCommunicationDisabledUntilNil

`func (o *UpdateGuildMemberRequest) SetCommunicationDisabledUntilNil(b bool)`

 SetCommunicationDisabledUntilNil sets the value for CommunicationDisabledUntil to be an explicit nil

### UnsetCommunicationDisabledUntil
`func (o *UpdateGuildMemberRequest) UnsetCommunicationDisabledUntil()`

UnsetCommunicationDisabledUntil ensures that no value is present for CommunicationDisabledUntil, not even an explicit nil
### GetFlags

`func (o *UpdateGuildMemberRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UpdateGuildMemberRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UpdateGuildMemberRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *UpdateGuildMemberRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *UpdateGuildMemberRequest) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *UpdateGuildMemberRequest) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


