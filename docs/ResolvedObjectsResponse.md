# ResolvedObjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**map[string]UserResponse**](UserResponse.md) |  | [optional] 
**Members** | Pointer to [**map[string]BasicGuildMemberResponse**](BasicGuildMemberResponse.md) |  | [optional] 
**Channels** | Pointer to [**map[string]GetChannel200Response**](GetChannel200Response.md) |  | [optional] 
**Roles** | Pointer to [**map[string]GuildRoleResponse**](GuildRoleResponse.md) |  | [optional] 

## Methods

### NewResolvedObjectsResponse

`func NewResolvedObjectsResponse() *ResolvedObjectsResponse`

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

### HasUsers

`func (o *ResolvedObjectsResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetMembers

`func (o *ResolvedObjectsResponse) GetMembers() map[string]BasicGuildMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ResolvedObjectsResponse) GetMembersOk() (*map[string]BasicGuildMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ResolvedObjectsResponse) SetMembers(v map[string]BasicGuildMemberResponse)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ResolvedObjectsResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

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

### HasChannels

`func (o *ResolvedObjectsResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

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

### HasRoles

`func (o *ResolvedObjectsResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


