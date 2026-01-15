# FriendInviteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InviteTypes**](InviteTypes.md) |  | 
**Code** | **string** |  | 
**Inviter** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**MaxAge** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**FriendsCount** | Pointer to **int32** |  | [optional] 
**Channel** | Pointer to [**FriendInviteResponseChannel**](FriendInviteResponseChannel.md) |  | [optional] 
**IsContact** | Pointer to **bool** |  | [optional] 
**Uses** | Pointer to **int32** |  | [optional] 
**MaxUses** | Pointer to **int32** |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 

## Methods

### NewFriendInviteResponse

`func NewFriendInviteResponse(type_ InviteTypes, code string, ) *FriendInviteResponse`

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

`func (o *FriendInviteResponse) GetType() InviteTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FriendInviteResponse) GetTypeOk() (*InviteTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FriendInviteResponse) SetType(v InviteTypes)`

SetType sets Type field to given value.


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

### GetChannel

`func (o *FriendInviteResponse) GetChannel() FriendInviteResponseChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *FriendInviteResponse) GetChannelOk() (*FriendInviteResponseChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *FriendInviteResponse) SetChannel(v FriendInviteResponseChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *FriendInviteResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


