# GroupDMInviteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InviteTypes**](InviteTypes.md) |  | 
**Code** | **string** |  | 
**Inviter** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**MaxAge** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Channel** | [**InviteChannelResponse**](InviteChannelResponse.md) |  | 
**ApproximateMemberCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewGroupDMInviteResponse

`func NewGroupDMInviteResponse(type_ InviteTypes, code string, channel InviteChannelResponse, ) *GroupDMInviteResponse`

NewGroupDMInviteResponse instantiates a new GroupDMInviteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupDMInviteResponseWithDefaults

`func NewGroupDMInviteResponseWithDefaults() *GroupDMInviteResponse`

NewGroupDMInviteResponseWithDefaults instantiates a new GroupDMInviteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GroupDMInviteResponse) GetType() InviteTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupDMInviteResponse) GetTypeOk() (*InviteTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupDMInviteResponse) SetType(v InviteTypes)`

SetType sets Type field to given value.


### GetCode

`func (o *GroupDMInviteResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GroupDMInviteResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GroupDMInviteResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetInviter

`func (o *GroupDMInviteResponse) GetInviter() UserResponse`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *GroupDMInviteResponse) GetInviterOk() (*UserResponse, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *GroupDMInviteResponse) SetInviter(v UserResponse)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *GroupDMInviteResponse) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetMaxAge

`func (o *GroupDMInviteResponse) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *GroupDMInviteResponse) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *GroupDMInviteResponse) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *GroupDMInviteResponse) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupDMInviteResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupDMInviteResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupDMInviteResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupDMInviteResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GroupDMInviteResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GroupDMInviteResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GroupDMInviteResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GroupDMInviteResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetChannel

`func (o *GroupDMInviteResponse) GetChannel() InviteChannelResponse`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *GroupDMInviteResponse) GetChannelOk() (*InviteChannelResponse, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *GroupDMInviteResponse) SetChannel(v InviteChannelResponse)`

SetChannel sets Channel field to given value.


### GetApproximateMemberCount

`func (o *GroupDMInviteResponse) GetApproximateMemberCount() int32`

GetApproximateMemberCount returns the ApproximateMemberCount field if non-nil, zero value otherwise.

### GetApproximateMemberCountOk

`func (o *GroupDMInviteResponse) GetApproximateMemberCountOk() (*int32, bool)`

GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMemberCount

`func (o *GroupDMInviteResponse) SetApproximateMemberCount(v int32)`

SetApproximateMemberCount sets ApproximateMemberCount field to given value.

### HasApproximateMemberCount

`func (o *GroupDMInviteResponse) HasApproximateMemberCount() bool`

HasApproximateMemberCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


