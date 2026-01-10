# PrivateApplicationResponseTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Icon** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OwnerUserId** | **string** |  | 
**Members** | [**[]TeamMemberResponse**](TeamMemberResponse.md) |  | 

## Methods

### NewPrivateApplicationResponseTeam

`func NewPrivateApplicationResponseTeam(id string, name string, ownerUserId string, members []TeamMemberResponse, ) *PrivateApplicationResponseTeam`

NewPrivateApplicationResponseTeam instantiates a new PrivateApplicationResponseTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateApplicationResponseTeamWithDefaults

`func NewPrivateApplicationResponseTeamWithDefaults() *PrivateApplicationResponseTeam`

NewPrivateApplicationResponseTeamWithDefaults instantiates a new PrivateApplicationResponseTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateApplicationResponseTeam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateApplicationResponseTeam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateApplicationResponseTeam) SetId(v string)`

SetId sets Id field to given value.


### GetIcon

`func (o *PrivateApplicationResponseTeam) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PrivateApplicationResponseTeam) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PrivateApplicationResponseTeam) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PrivateApplicationResponseTeam) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetName

`func (o *PrivateApplicationResponseTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateApplicationResponseTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateApplicationResponseTeam) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerUserId

`func (o *PrivateApplicationResponseTeam) GetOwnerUserId() string`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *PrivateApplicationResponseTeam) GetOwnerUserIdOk() (*string, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *PrivateApplicationResponseTeam) SetOwnerUserId(v string)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetMembers

`func (o *PrivateApplicationResponseTeam) GetMembers() []TeamMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *PrivateApplicationResponseTeam) GetMembersOk() (*[]TeamMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *PrivateApplicationResponseTeam) SetMembers(v []TeamMemberResponse)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


