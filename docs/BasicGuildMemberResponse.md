# BasicGuildMemberResponse

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
**Collectibles** | Pointer to [**NullableUserCollectiblesResponse**](UserCollectiblesResponse.md) |  | [optional] 

## Methods

### NewBasicGuildMemberResponse

`func NewBasicGuildMemberResponse(flags int32, joinedAt time.Time, pending bool, roles []string, ) *BasicGuildMemberResponse`

NewBasicGuildMemberResponse instantiates a new BasicGuildMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicGuildMemberResponseWithDefaults

`func NewBasicGuildMemberResponseWithDefaults() *BasicGuildMemberResponse`

NewBasicGuildMemberResponseWithDefaults instantiates a new BasicGuildMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *BasicGuildMemberResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *BasicGuildMemberResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *BasicGuildMemberResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *BasicGuildMemberResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *BasicGuildMemberResponse) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *BasicGuildMemberResponse) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarDecorationData

`func (o *BasicGuildMemberResponse) GetAvatarDecorationData() UserAvatarDecorationResponse`

GetAvatarDecorationData returns the AvatarDecorationData field if non-nil, zero value otherwise.

### GetAvatarDecorationDataOk

`func (o *BasicGuildMemberResponse) GetAvatarDecorationDataOk() (*UserAvatarDecorationResponse, bool)`

GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarDecorationData

`func (o *BasicGuildMemberResponse) SetAvatarDecorationData(v UserAvatarDecorationResponse)`

SetAvatarDecorationData sets AvatarDecorationData field to given value.

### HasAvatarDecorationData

`func (o *BasicGuildMemberResponse) HasAvatarDecorationData() bool`

HasAvatarDecorationData returns a boolean if a field has been set.

### SetAvatarDecorationDataNil

`func (o *BasicGuildMemberResponse) SetAvatarDecorationDataNil(b bool)`

 SetAvatarDecorationDataNil sets the value for AvatarDecorationData to be an explicit nil

### UnsetAvatarDecorationData
`func (o *BasicGuildMemberResponse) UnsetAvatarDecorationData()`

UnsetAvatarDecorationData ensures that no value is present for AvatarDecorationData, not even an explicit nil
### GetBanner

`func (o *BasicGuildMemberResponse) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *BasicGuildMemberResponse) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *BasicGuildMemberResponse) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *BasicGuildMemberResponse) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### SetBannerNil

`func (o *BasicGuildMemberResponse) SetBannerNil(b bool)`

 SetBannerNil sets the value for Banner to be an explicit nil

### UnsetBanner
`func (o *BasicGuildMemberResponse) UnsetBanner()`

UnsetBanner ensures that no value is present for Banner, not even an explicit nil
### GetCommunicationDisabledUntil

`func (o *BasicGuildMemberResponse) GetCommunicationDisabledUntil() time.Time`

GetCommunicationDisabledUntil returns the CommunicationDisabledUntil field if non-nil, zero value otherwise.

### GetCommunicationDisabledUntilOk

`func (o *BasicGuildMemberResponse) GetCommunicationDisabledUntilOk() (*time.Time, bool)`

GetCommunicationDisabledUntilOk returns a tuple with the CommunicationDisabledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDisabledUntil

`func (o *BasicGuildMemberResponse) SetCommunicationDisabledUntil(v time.Time)`

SetCommunicationDisabledUntil sets CommunicationDisabledUntil field to given value.

### HasCommunicationDisabledUntil

`func (o *BasicGuildMemberResponse) HasCommunicationDisabledUntil() bool`

HasCommunicationDisabledUntil returns a boolean if a field has been set.

### SetCommunicationDisabledUntilNil

`func (o *BasicGuildMemberResponse) SetCommunicationDisabledUntilNil(b bool)`

 SetCommunicationDisabledUntilNil sets the value for CommunicationDisabledUntil to be an explicit nil

### UnsetCommunicationDisabledUntil
`func (o *BasicGuildMemberResponse) UnsetCommunicationDisabledUntil()`

UnsetCommunicationDisabledUntil ensures that no value is present for CommunicationDisabledUntil, not even an explicit nil
### GetFlags

`func (o *BasicGuildMemberResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *BasicGuildMemberResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *BasicGuildMemberResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetJoinedAt

`func (o *BasicGuildMemberResponse) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *BasicGuildMemberResponse) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *BasicGuildMemberResponse) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.


### GetNick

`func (o *BasicGuildMemberResponse) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *BasicGuildMemberResponse) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *BasicGuildMemberResponse) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *BasicGuildMemberResponse) HasNick() bool`

HasNick returns a boolean if a field has been set.

### SetNickNil

`func (o *BasicGuildMemberResponse) SetNickNil(b bool)`

 SetNickNil sets the value for Nick to be an explicit nil

### UnsetNick
`func (o *BasicGuildMemberResponse) UnsetNick()`

UnsetNick ensures that no value is present for Nick, not even an explicit nil
### GetPending

`func (o *BasicGuildMemberResponse) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *BasicGuildMemberResponse) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *BasicGuildMemberResponse) SetPending(v bool)`

SetPending sets Pending field to given value.


### GetPremiumSince

`func (o *BasicGuildMemberResponse) GetPremiumSince() time.Time`

GetPremiumSince returns the PremiumSince field if non-nil, zero value otherwise.

### GetPremiumSinceOk

`func (o *BasicGuildMemberResponse) GetPremiumSinceOk() (*time.Time, bool)`

GetPremiumSinceOk returns a tuple with the PremiumSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSince

`func (o *BasicGuildMemberResponse) SetPremiumSince(v time.Time)`

SetPremiumSince sets PremiumSince field to given value.

### HasPremiumSince

`func (o *BasicGuildMemberResponse) HasPremiumSince() bool`

HasPremiumSince returns a boolean if a field has been set.

### SetPremiumSinceNil

`func (o *BasicGuildMemberResponse) SetPremiumSinceNil(b bool)`

 SetPremiumSinceNil sets the value for PremiumSince to be an explicit nil

### UnsetPremiumSince
`func (o *BasicGuildMemberResponse) UnsetPremiumSince()`

UnsetPremiumSince ensures that no value is present for PremiumSince, not even an explicit nil
### GetRoles

`func (o *BasicGuildMemberResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BasicGuildMemberResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BasicGuildMemberResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetCollectibles

`func (o *BasicGuildMemberResponse) GetCollectibles() UserCollectiblesResponse`

GetCollectibles returns the Collectibles field if non-nil, zero value otherwise.

### GetCollectiblesOk

`func (o *BasicGuildMemberResponse) GetCollectiblesOk() (*UserCollectiblesResponse, bool)`

GetCollectiblesOk returns a tuple with the Collectibles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectibles

`func (o *BasicGuildMemberResponse) SetCollectibles(v UserCollectiblesResponse)`

SetCollectibles sets Collectibles field to given value.

### HasCollectibles

`func (o *BasicGuildMemberResponse) HasCollectibles() bool`

HasCollectibles returns a boolean if a field has been set.

### SetCollectiblesNil

`func (o *BasicGuildMemberResponse) SetCollectiblesNil(b bool)`

 SetCollectiblesNil sets the value for Collectibles to be an explicit nil

### UnsetCollectibles
`func (o *BasicGuildMemberResponse) UnsetCollectibles()`

UnsetCollectibles ensures that no value is present for Collectibles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


