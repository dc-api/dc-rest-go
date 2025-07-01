# ThreadMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**JoinTimestamp** | **time.Time** |  | 
**Flags** | **int32** |  | 
**Member** | Pointer to [**NullableGuildMemberResponse**](GuildMemberResponse.md) |  | [optional] 

## Methods

### NewThreadMemberResponse

`func NewThreadMemberResponse(id string, userId string, joinTimestamp time.Time, flags int32, ) *ThreadMemberResponse`

NewThreadMemberResponse instantiates a new ThreadMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadMemberResponseWithDefaults

`func NewThreadMemberResponseWithDefaults() *ThreadMemberResponse`

NewThreadMemberResponseWithDefaults instantiates a new ThreadMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThreadMemberResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreadMemberResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreadMemberResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ThreadMemberResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ThreadMemberResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ThreadMemberResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetJoinTimestamp

`func (o *ThreadMemberResponse) GetJoinTimestamp() time.Time`

GetJoinTimestamp returns the JoinTimestamp field if non-nil, zero value otherwise.

### GetJoinTimestampOk

`func (o *ThreadMemberResponse) GetJoinTimestampOk() (*time.Time, bool)`

GetJoinTimestampOk returns a tuple with the JoinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTimestamp

`func (o *ThreadMemberResponse) SetJoinTimestamp(v time.Time)`

SetJoinTimestamp sets JoinTimestamp field to given value.


### GetFlags

`func (o *ThreadMemberResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ThreadMemberResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ThreadMemberResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetMember

`func (o *ThreadMemberResponse) GetMember() GuildMemberResponse`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ThreadMemberResponse) GetMemberOk() (*GuildMemberResponse, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ThreadMemberResponse) SetMember(v GuildMemberResponse)`

SetMember sets Member field to given value.

### HasMember

`func (o *ThreadMemberResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### SetMemberNil

`func (o *ThreadMemberResponse) SetMemberNil(b bool)`

 SetMemberNil sets the value for Member to be an explicit nil

### UnsetMember
`func (o *ThreadMemberResponse) UnsetMember()`

UnsetMember ensures that no value is present for Member, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


