# TeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**OwnerUserId** | **string** |  | 
**Members** | [**[]TeamMemberResponse**](TeamMemberResponse.md) |  | 

## Methods

### NewTeamResponse

`func NewTeamResponse(id string, name string, ownerUserId string, members []TeamMemberResponse, ) *TeamResponse`

NewTeamResponse instantiates a new TeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamResponseWithDefaults

`func NewTeamResponseWithDefaults() *TeamResponse`

NewTeamResponseWithDefaults instantiates a new TeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIcon

`func (o *TeamResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *TeamResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *TeamResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *TeamResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *TeamResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *TeamResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetName

`func (o *TeamResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerUserId

`func (o *TeamResponse) GetOwnerUserId() string`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *TeamResponse) GetOwnerUserIdOk() (*string, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *TeamResponse) SetOwnerUserId(v string)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetMembers

`func (o *TeamResponse) GetMembers() []TeamMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamResponse) GetMembersOk() (*[]TeamMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamResponse) SetMembers(v []TeamMemberResponse)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


