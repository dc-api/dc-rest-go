# TeamMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserResponse**](UserResponse.md) |  | 
**TeamId** | **string** |  | 
**MembershipState** | [**TeamMembershipStates**](TeamMembershipStates.md) |  | 

## Methods

### NewTeamMemberResponse

`func NewTeamMemberResponse(user UserResponse, teamId string, membershipState TeamMembershipStates, ) *TeamMemberResponse`

NewTeamMemberResponse instantiates a new TeamMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberResponseWithDefaults

`func NewTeamMemberResponseWithDefaults() *TeamMemberResponse`

NewTeamMemberResponseWithDefaults instantiates a new TeamMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *TeamMemberResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeamMemberResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeamMemberResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.


### GetTeamId

`func (o *TeamMemberResponse) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamMemberResponse) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamMemberResponse) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetMembershipState

`func (o *TeamMemberResponse) GetMembershipState() TeamMembershipStates`

GetMembershipState returns the MembershipState field if non-nil, zero value otherwise.

### GetMembershipStateOk

`func (o *TeamMemberResponse) GetMembershipStateOk() (*TeamMembershipStates, bool)`

GetMembershipStateOk returns a tuple with the MembershipState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipState

`func (o *TeamMemberResponse) SetMembershipState(v TeamMembershipStates)`

SetMembershipState sets MembershipState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


