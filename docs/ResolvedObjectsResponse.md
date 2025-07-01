# ResolvedObjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**map[string]UserResponse**](UserResponse.md) |  | 
**Members** | [**map[string]GuildMemberResponse**](GuildMemberResponse.md) |  | 
**Channels** | [**map[string]GetChannel200Response**](GetChannel200Response.md) |  | 
**Roles** | [**map[string]GuildRoleResponse**](GuildRoleResponse.md) |  | 

## Methods

### NewResolvedObjectsResponse

`func NewResolvedObjectsResponse(users map[string]UserResponse, members map[string]GuildMemberResponse, channels map[string]GetChannel200Response, roles map[string]GuildRoleResponse, ) *ResolvedObjectsResponse`

NewResolvedObjectsResponse instantiates a new ResolvedObjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolvedObjectsResponseWithDefaults

`func NewResolvedObjectsResponseWithDefaults() *ResolvedObjectsResponse`

NewResolvedObjectsResponseWithDefaults instantiates a new ResolvedObjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ResolvedObjectsResponse) GetUsers() map[string]UserResponse`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ResolvedObjectsResponse) GetUsersOk() (*map[string]UserResponse, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ResolvedObjectsResponse) SetUsers(v map[string]UserResponse)`

SetUsers sets Users field to given value.


### GetMembers

`func (o *ResolvedObjectsResponse) GetMembers() map[string]GuildMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ResolvedObjectsResponse) GetMembersOk() (*map[string]GuildMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ResolvedObjectsResponse) SetMembers(v map[string]GuildMemberResponse)`

SetMembers sets Members field to given value.


### GetChannels

`func (o *ResolvedObjectsResponse) GetChannels() map[string]GetChannel200Response`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ResolvedObjectsResponse) GetChannelsOk() (*map[string]GetChannel200Response, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ResolvedObjectsResponse) SetChannels(v map[string]GetChannel200Response)`

SetChannels sets Channels field to given value.


### GetRoles

`func (o *ResolvedObjectsResponse) GetRoles() map[string]GuildRoleResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ResolvedObjectsResponse) GetRolesOk() (*map[string]GuildRoleResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ResolvedObjectsResponse) SetRoles(v map[string]GuildRoleResponse)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


