# WidgetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**InstantInvite** | Pointer to **NullableString** |  | [optional] 
**Channels** | [**[]WidgetChannel**](WidgetChannel.md) |  | 
**Members** | [**[]WidgetMember**](WidgetMember.md) |  | 
**PresenceCount** | **int32** |  | 

## Methods

### NewWidgetResponse

`func NewWidgetResponse(id string, name string, channels []WidgetChannel, members []WidgetMember, presenceCount int32, ) *WidgetResponse`

NewWidgetResponse instantiates a new WidgetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetResponseWithDefaults

`func NewWidgetResponseWithDefaults() *WidgetResponse`

NewWidgetResponseWithDefaults instantiates a new WidgetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WidgetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WidgetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WidgetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WidgetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WidgetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WidgetResponse) SetName(v string)`

SetName sets Name field to given value.


### GetInstantInvite

`func (o *WidgetResponse) GetInstantInvite() string`

GetInstantInvite returns the InstantInvite field if non-nil, zero value otherwise.

### GetInstantInviteOk

`func (o *WidgetResponse) GetInstantInviteOk() (*string, bool)`

GetInstantInviteOk returns a tuple with the InstantInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantInvite

`func (o *WidgetResponse) SetInstantInvite(v string)`

SetInstantInvite sets InstantInvite field to given value.

### HasInstantInvite

`func (o *WidgetResponse) HasInstantInvite() bool`

HasInstantInvite returns a boolean if a field has been set.

### SetInstantInviteNil

`func (o *WidgetResponse) SetInstantInviteNil(b bool)`

 SetInstantInviteNil sets the value for InstantInvite to be an explicit nil

### UnsetInstantInvite
`func (o *WidgetResponse) UnsetInstantInvite()`

UnsetInstantInvite ensures that no value is present for InstantInvite, not even an explicit nil
### GetChannels

`func (o *WidgetResponse) GetChannels() []WidgetChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *WidgetResponse) GetChannelsOk() (*[]WidgetChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *WidgetResponse) SetChannels(v []WidgetChannel)`

SetChannels sets Channels field to given value.


### GetMembers

`func (o *WidgetResponse) GetMembers() []WidgetMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *WidgetResponse) GetMembersOk() (*[]WidgetMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *WidgetResponse) SetMembers(v []WidgetMember)`

SetMembers sets Members field to given value.


### GetPresenceCount

`func (o *WidgetResponse) GetPresenceCount() int32`

GetPresenceCount returns the PresenceCount field if non-nil, zero value otherwise.

### GetPresenceCountOk

`func (o *WidgetResponse) GetPresenceCountOk() (*int32, bool)`

GetPresenceCountOk returns a tuple with the PresenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceCount

`func (o *WidgetResponse) SetPresenceCount(v int32)`

SetPresenceCount sets PresenceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


