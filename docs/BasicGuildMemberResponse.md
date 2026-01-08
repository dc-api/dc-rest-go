# BasicGuildMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** |  | [optional] 
**AvatarDecorationData** | Pointer to [**BasicGuildMemberResponseAvatarDecorationData**](BasicGuildMemberResponseAvatarDecorationData.md) |  | [optional] 
**Banner** | Pointer to **string** |  | [optional] 
**CommunicationDisabledUntil** | Pointer to **time.Time** |  | [optional] 
**Flags** | **int32** |  | 
**JoinedAt** | **time.Time** |  | 
**Nick** | Pointer to **string** |  | [optional] 
**Pending** | **bool** |  | 
**PremiumSince** | Pointer to **time.Time** |  | [optional] 
**Roles** | **[]string** |  | 
**Collectibles** | Pointer to [**BasicGuildMemberResponseCollectibles**](BasicGuildMemberResponseCollectibles.md) |  | [optional] 

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

### GetAvatarDecorationData

`func (o *BasicGuildMemberResponse) GetAvatarDecorationData() BasicGuildMemberResponseAvatarDecorationData`

GetAvatarDecorationData returns the AvatarDecorationData field if non-nil, zero value otherwise.

### GetAvatarDecorationDataOk

`func (o *BasicGuildMemberResponse) GetAvatarDecorationDataOk() (*BasicGuildMemberResponseAvatarDecorationData, bool)`

GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarDecorationData

`func (o *BasicGuildMemberResponse) SetAvatarDecorationData(v BasicGuildMemberResponseAvatarDecorationData)`

SetAvatarDecorationData sets AvatarDecorationData field to given value.

### HasAvatarDecorationData

`func (o *BasicGuildMemberResponse) HasAvatarDecorationData() bool`

HasAvatarDecorationData returns a boolean if a field has been set.

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

`func (o *BasicGuildMemberResponse) GetCollectibles() BasicGuildMemberResponseCollectibles`

GetCollectibles returns the Collectibles field if non-nil, zero value otherwise.

### GetCollectiblesOk

`func (o *BasicGuildMemberResponse) GetCollectiblesOk() (*BasicGuildMemberResponseCollectibles, bool)`

GetCollectiblesOk returns a tuple with the Collectibles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectibles

`func (o *BasicGuildMemberResponse) SetCollectibles(v BasicGuildMemberResponseCollectibles)`

SetCollectibles sets Collectibles field to given value.

### HasCollectibles

`func (o *BasicGuildMemberResponse) HasCollectibles() bool`

HasCollectibles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


