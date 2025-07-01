# GuildMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **NullableString** |  | [optional] 
**AvatarDecorationData** | Pointer to [**NullableUserAvatarDecorationResponse**](UserAvatarDecorationResponse.md) |  | [optional] 
**Banner** | Pointer to **NullableString** |  | [optional] 
**CommunicationDisabledUntil** | Pointer to **NullableTime** |  | [optional] 
**Flags** | **int32** |  | 
**JoinedAt** | **time.Time** |  | 
**Nick** | Pointer to **NullableString** |  | [optional] 
**Pending** | **bool** |  | 
**PremiumSince** | Pointer to **NullableTime** |  | [optional] 
**Roles** | **[]string** |  | 
**User** | [**UserResponse**](UserResponse.md) |  | 
**Mute** | **bool** |  | 
**Deaf** | **bool** |  | 

## Methods

### NewGuildMemberResponse

`func NewGuildMemberResponse(flags int32, joinedAt time.Time, pending bool, roles []string, user UserResponse, mute bool, deaf bool, ) *GuildMemberResponse`

NewGuildMemberResponse instantiates a new GuildMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildMemberResponseWithDefaults

`func NewGuildMemberResponseWithDefaults() *GuildMemberResponse`

NewGuildMemberResponseWithDefaults instantiates a new GuildMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *GuildMemberResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GuildMemberResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GuildMemberResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *GuildMemberResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *GuildMemberResponse) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *GuildMemberResponse) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarDecorationData

`func (o *GuildMemberResponse) GetAvatarDecorationData() UserAvatarDecorationResponse`

GetAvatarDecorationData returns the AvatarDecorationData field if non-nil, zero value otherwise.

### GetAvatarDecorationDataOk

`func (o *GuildMemberResponse) GetAvatarDecorationDataOk() (*UserAvatarDecorationResponse, bool)`

GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarDecorationData

`func (o *GuildMemberResponse) SetAvatarDecorationData(v UserAvatarDecorationResponse)`

SetAvatarDecorationData sets AvatarDecorationData field to given value.

### HasAvatarDecorationData

`func (o *GuildMemberResponse) HasAvatarDecorationData() bool`

HasAvatarDecorationData returns a boolean if a field has been set.

### SetAvatarDecorationDataNil

`func (o *GuildMemberResponse) SetAvatarDecorationDataNil(b bool)`

 SetAvatarDecorationDataNil sets the value for AvatarDecorationData to be an explicit nil

### UnsetAvatarDecorationData
`func (o *GuildMemberResponse) UnsetAvatarDecorationData()`

UnsetAvatarDecorationData ensures that no value is present for AvatarDecorationData, not even an explicit nil
### GetBanner

`func (o *GuildMemberResponse) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *GuildMemberResponse) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *GuildMemberResponse) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *GuildMemberResponse) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### SetBannerNil

`func (o *GuildMemberResponse) SetBannerNil(b bool)`

 SetBannerNil sets the value for Banner to be an explicit nil

### UnsetBanner
`func (o *GuildMemberResponse) UnsetBanner()`

UnsetBanner ensures that no value is present for Banner, not even an explicit nil
### GetCommunicationDisabledUntil

`func (o *GuildMemberResponse) GetCommunicationDisabledUntil() time.Time`

GetCommunicationDisabledUntil returns the CommunicationDisabledUntil field if non-nil, zero value otherwise.

### GetCommunicationDisabledUntilOk

`func (o *GuildMemberResponse) GetCommunicationDisabledUntilOk() (*time.Time, bool)`

GetCommunicationDisabledUntilOk returns a tuple with the CommunicationDisabledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDisabledUntil

`func (o *GuildMemberResponse) SetCommunicationDisabledUntil(v time.Time)`

SetCommunicationDisabledUntil sets CommunicationDisabledUntil field to given value.

### HasCommunicationDisabledUntil

`func (o *GuildMemberResponse) HasCommunicationDisabledUntil() bool`

HasCommunicationDisabledUntil returns a boolean if a field has been set.

### SetCommunicationDisabledUntilNil

`func (o *GuildMemberResponse) SetCommunicationDisabledUntilNil(b bool)`

 SetCommunicationDisabledUntilNil sets the value for CommunicationDisabledUntil to be an explicit nil

### UnsetCommunicationDisabledUntil
`func (o *GuildMemberResponse) UnsetCommunicationDisabledUntil()`

UnsetCommunicationDisabledUntil ensures that no value is present for CommunicationDisabledUntil, not even an explicit nil
### GetFlags

`func (o *GuildMemberResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GuildMemberResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GuildMemberResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetJoinedAt

`func (o *GuildMemberResponse) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *GuildMemberResponse) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *GuildMemberResponse) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.


### GetNick

`func (o *GuildMemberResponse) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *GuildMemberResponse) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *GuildMemberResponse) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *GuildMemberResponse) HasNick() bool`

HasNick returns a boolean if a field has been set.

### SetNickNil

`func (o *GuildMemberResponse) SetNickNil(b bool)`

 SetNickNil sets the value for Nick to be an explicit nil

### UnsetNick
`func (o *GuildMemberResponse) UnsetNick()`

UnsetNick ensures that no value is present for Nick, not even an explicit nil
### GetPending

`func (o *GuildMemberResponse) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GuildMemberResponse) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GuildMemberResponse) SetPending(v bool)`

SetPending sets Pending field to given value.


### GetPremiumSince

`func (o *GuildMemberResponse) GetPremiumSince() time.Time`

GetPremiumSince returns the PremiumSince field if non-nil, zero value otherwise.

### GetPremiumSinceOk

`func (o *GuildMemberResponse) GetPremiumSinceOk() (*time.Time, bool)`

GetPremiumSinceOk returns a tuple with the PremiumSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSince

`func (o *GuildMemberResponse) SetPremiumSince(v time.Time)`

SetPremiumSince sets PremiumSince field to given value.

### HasPremiumSince

`func (o *GuildMemberResponse) HasPremiumSince() bool`

HasPremiumSince returns a boolean if a field has been set.

### SetPremiumSinceNil

`func (o *GuildMemberResponse) SetPremiumSinceNil(b bool)`

 SetPremiumSinceNil sets the value for PremiumSince to be an explicit nil

### UnsetPremiumSince
`func (o *GuildMemberResponse) UnsetPremiumSince()`

UnsetPremiumSince ensures that no value is present for PremiumSince, not even an explicit nil
### GetRoles

`func (o *GuildMemberResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GuildMemberResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GuildMemberResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetUser

`func (o *GuildMemberResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GuildMemberResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GuildMemberResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.


### GetMute

`func (o *GuildMemberResponse) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *GuildMemberResponse) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *GuildMemberResponse) SetMute(v bool)`

SetMute sets Mute field to given value.


### GetDeaf

`func (o *GuildMemberResponse) GetDeaf() bool`

GetDeaf returns the Deaf field if non-nil, zero value otherwise.

### GetDeafOk

`func (o *GuildMemberResponse) GetDeafOk() (*bool, bool)`

GetDeafOk returns a tuple with the Deaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaf

`func (o *GuildMemberResponse) SetDeaf(v bool)`

SetDeaf sets Deaf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


