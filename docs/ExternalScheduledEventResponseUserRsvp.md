# ExternalScheduledEventResponseUserRsvp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuildScheduledEventId** | **string** |  | 
**UserId** | **string** |  | 
**User** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**Member** | Pointer to [**GuildMemberResponse**](GuildMemberResponse.md) |  | [optional] 

## Methods

### NewExternalScheduledEventResponseUserRsvp

`func NewExternalScheduledEventResponseUserRsvp(guildScheduledEventId string, userId string, ) *ExternalScheduledEventResponseUserRsvp`

NewExternalScheduledEventResponseUserRsvp instantiates a new ExternalScheduledEventResponseUserRsvp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalScheduledEventResponseUserRsvpWithDefaults

`func NewExternalScheduledEventResponseUserRsvpWithDefaults() *ExternalScheduledEventResponseUserRsvp`

NewExternalScheduledEventResponseUserRsvpWithDefaults instantiates a new ExternalScheduledEventResponseUserRsvp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuildScheduledEventId

`func (o *ExternalScheduledEventResponseUserRsvp) GetGuildScheduledEventId() string`

GetGuildScheduledEventId returns the GuildScheduledEventId field if non-nil, zero value otherwise.

### GetGuildScheduledEventIdOk

`func (o *ExternalScheduledEventResponseUserRsvp) GetGuildScheduledEventIdOk() (*string, bool)`

GetGuildScheduledEventIdOk returns a tuple with the GuildScheduledEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildScheduledEventId

`func (o *ExternalScheduledEventResponseUserRsvp) SetGuildScheduledEventId(v string)`

SetGuildScheduledEventId sets GuildScheduledEventId field to given value.


### GetUserId

`func (o *ExternalScheduledEventResponseUserRsvp) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalScheduledEventResponseUserRsvp) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalScheduledEventResponseUserRsvp) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *ExternalScheduledEventResponseUserRsvp) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExternalScheduledEventResponseUserRsvp) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExternalScheduledEventResponseUserRsvp) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ExternalScheduledEventResponseUserRsvp) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetMember

`func (o *ExternalScheduledEventResponseUserRsvp) GetMember() GuildMemberResponse`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ExternalScheduledEventResponseUserRsvp) GetMemberOk() (*GuildMemberResponse, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ExternalScheduledEventResponseUserRsvp) SetMember(v GuildMemberResponse)`

SetMember sets Member field to given value.

### HasMember

`func (o *ExternalScheduledEventResponseUserRsvp) HasMember() bool`

HasMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


