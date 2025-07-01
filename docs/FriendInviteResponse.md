# FriendInviteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableInt32** |  | [optional] 
**Code** | **string** |  | 
**Inviter** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 
**MaxAge** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**FriendsCount** | Pointer to **NullableInt32** |  | [optional] 
**Channel** | Pointer to [**NullableInviteChannelResponse**](InviteChannelResponse.md) |  | [optional] 
**IsContact** | Pointer to **NullableBool** |  | [optional] 
**Uses** | Pointer to **NullableInt32** |  | [optional] 
**MaxUses** | Pointer to **NullableInt32** |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewFriendInviteResponse

`func NewFriendInviteResponse(code string, ) *FriendInviteResponse`

NewFriendInviteResponse instantiates a new FriendInviteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFriendInviteResponseWithDefaults

`func NewFriendInviteResponseWithDefaults() *FriendInviteResponse`

NewFriendInviteResponseWithDefaults instantiates a new FriendInviteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FriendInviteResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FriendInviteResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FriendInviteResponse) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *FriendInviteResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FriendInviteResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FriendInviteResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCode

`func (o *FriendInviteResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FriendInviteResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FriendInviteResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetInviter

`func (o *FriendInviteResponse) GetInviter() UserResponse`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *FriendInviteResponse) GetInviterOk() (*UserResponse, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *FriendInviteResponse) SetInviter(v UserResponse)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *FriendInviteResponse) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### SetInviterNil

`func (o *FriendInviteResponse) SetInviterNil(b bool)`

 SetInviterNil sets the value for Inviter to be an explicit nil

### UnsetInviter
`func (o *FriendInviteResponse) UnsetInviter()`

UnsetInviter ensures that no value is present for Inviter, not even an explicit nil
### GetMaxAge

`func (o *FriendInviteResponse) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *FriendInviteResponse) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *FriendInviteResponse) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *FriendInviteResponse) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### SetMaxAgeNil

`func (o *FriendInviteResponse) SetMaxAgeNil(b bool)`

 SetMaxAgeNil sets the value for MaxAge to be an explicit nil

### UnsetMaxAge
`func (o *FriendInviteResponse) UnsetMaxAge()`

UnsetMaxAge ensures that no value is present for MaxAge, not even an explicit nil
### GetCreatedAt

`func (o *FriendInviteResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FriendInviteResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FriendInviteResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FriendInviteResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *FriendInviteResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *FriendInviteResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpiresAt

`func (o *FriendInviteResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FriendInviteResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FriendInviteResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FriendInviteResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *FriendInviteResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *FriendInviteResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFriendsCount

`func (o *FriendInviteResponse) GetFriendsCount() int32`

GetFriendsCount returns the FriendsCount field if non-nil, zero value otherwise.

### GetFriendsCountOk

`func (o *FriendInviteResponse) GetFriendsCountOk() (*int32, bool)`

GetFriendsCountOk returns a tuple with the FriendsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendsCount

`func (o *FriendInviteResponse) SetFriendsCount(v int32)`

SetFriendsCount sets FriendsCount field to given value.

### HasFriendsCount

`func (o *FriendInviteResponse) HasFriendsCount() bool`

HasFriendsCount returns a boolean if a field has been set.

### SetFriendsCountNil

`func (o *FriendInviteResponse) SetFriendsCountNil(b bool)`

 SetFriendsCountNil sets the value for FriendsCount to be an explicit nil

### UnsetFriendsCount
`func (o *FriendInviteResponse) UnsetFriendsCount()`

UnsetFriendsCount ensures that no value is present for FriendsCount, not even an explicit nil
### GetChannel

`func (o *FriendInviteResponse) GetChannel() InviteChannelResponse`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *FriendInviteResponse) GetChannelOk() (*InviteChannelResponse, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *FriendInviteResponse) SetChannel(v InviteChannelResponse)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *FriendInviteResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *FriendInviteResponse) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *FriendInviteResponse) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetIsContact

`func (o *FriendInviteResponse) GetIsContact() bool`

GetIsContact returns the IsContact field if non-nil, zero value otherwise.

### GetIsContactOk

`func (o *FriendInviteResponse) GetIsContactOk() (*bool, bool)`

GetIsContactOk returns a tuple with the IsContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContact

`func (o *FriendInviteResponse) SetIsContact(v bool)`

SetIsContact sets IsContact field to given value.

### HasIsContact

`func (o *FriendInviteResponse) HasIsContact() bool`

HasIsContact returns a boolean if a field has been set.

### SetIsContactNil

`func (o *FriendInviteResponse) SetIsContactNil(b bool)`

 SetIsContactNil sets the value for IsContact to be an explicit nil

### UnsetIsContact
`func (o *FriendInviteResponse) UnsetIsContact()`

UnsetIsContact ensures that no value is present for IsContact, not even an explicit nil
### GetUses

`func (o *FriendInviteResponse) GetUses() int32`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *FriendInviteResponse) GetUsesOk() (*int32, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *FriendInviteResponse) SetUses(v int32)`

SetUses sets Uses field to given value.

### HasUses

`func (o *FriendInviteResponse) HasUses() bool`

HasUses returns a boolean if a field has been set.

### SetUsesNil

`func (o *FriendInviteResponse) SetUsesNil(b bool)`

 SetUsesNil sets the value for Uses to be an explicit nil

### UnsetUses
`func (o *FriendInviteResponse) UnsetUses()`

UnsetUses ensures that no value is present for Uses, not even an explicit nil
### GetMaxUses

`func (o *FriendInviteResponse) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *FriendInviteResponse) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *FriendInviteResponse) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *FriendInviteResponse) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### SetMaxUsesNil

`func (o *FriendInviteResponse) SetMaxUsesNil(b bool)`

 SetMaxUsesNil sets the value for MaxUses to be an explicit nil

### UnsetMaxUses
`func (o *FriendInviteResponse) UnsetMaxUses()`

UnsetMaxUses ensures that no value is present for MaxUses, not even an explicit nil
### GetFlags

`func (o *FriendInviteResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *FriendInviteResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *FriendInviteResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *FriendInviteResponse) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *FriendInviteResponse) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *FriendInviteResponse) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


