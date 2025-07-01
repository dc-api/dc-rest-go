# ScheduledEventUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuildScheduledEventId** | **string** |  | 
**UserId** | **string** |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 
**Member** | Pointer to [**NullableGuildMemberResponse**](GuildMemberResponse.md) |  | [optional] 

## Methods

### NewScheduledEventUserResponse

`func NewScheduledEventUserResponse(guildScheduledEventId string, userId string, ) *ScheduledEventUserResponse`

NewScheduledEventUserResponse instantiates a new ScheduledEventUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledEventUserResponseWithDefaults

`func NewScheduledEventUserResponseWithDefaults() *ScheduledEventUserResponse`

NewScheduledEventUserResponseWithDefaults instantiates a new ScheduledEventUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuildScheduledEventId

`func (o *ScheduledEventUserResponse) GetGuildScheduledEventId() string`

GetGuildScheduledEventId returns the GuildScheduledEventId field if non-nil, zero value otherwise.

### GetGuildScheduledEventIdOk

`func (o *ScheduledEventUserResponse) GetGuildScheduledEventIdOk() (*string, bool)`

GetGuildScheduledEventIdOk returns a tuple with the GuildScheduledEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildScheduledEventId

`func (o *ScheduledEventUserResponse) SetGuildScheduledEventId(v string)`

SetGuildScheduledEventId sets GuildScheduledEventId field to given value.


### GetUserId

`func (o *ScheduledEventUserResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ScheduledEventUserResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ScheduledEventUserResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *ScheduledEventUserResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ScheduledEventUserResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ScheduledEventUserResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ScheduledEventUserResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ScheduledEventUserResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ScheduledEventUserResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetMember

`func (o *ScheduledEventUserResponse) GetMember() GuildMemberResponse`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ScheduledEventUserResponse) GetMemberOk() (*GuildMemberResponse, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ScheduledEventUserResponse) SetMember(v GuildMemberResponse)`

SetMember sets Member field to given value.

### HasMember

`func (o *ScheduledEventUserResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### SetMemberNil

`func (o *ScheduledEventUserResponse) SetMemberNil(b bool)`

 SetMemberNil sets the value for Member to be an explicit nil

### UnsetMember
`func (o *ScheduledEventUserResponse) UnsetMember()`

UnsetMember ensures that no value is present for Member, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


