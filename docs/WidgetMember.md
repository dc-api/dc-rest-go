# WidgetMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Discriminator** | [**WidgetUserDiscriminator**](WidgetUserDiscriminator.md) |  | 
**Avatar** | Pointer to [**nil**](nil.md) |  | [optional] 
**Status** | **string** |  | 
**AvatarUrl** | **string** |  | 
**Activity** | Pointer to [**WidgetActivity**](WidgetActivity.md) |  | [optional] 
**Deaf** | Pointer to **bool** |  | [optional] 
**Mute** | Pointer to **bool** |  | [optional] 
**SelfDeaf** | Pointer to **bool** |  | [optional] 
**SelfMute** | Pointer to **bool** |  | [optional] 
**Suppress** | Pointer to **bool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 

## Methods

### NewWidgetMember

`func NewWidgetMember(id string, username string, discriminator WidgetUserDiscriminator, status string, avatarUrl string, ) *WidgetMember`

NewWidgetMember instantiates a new WidgetMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetMemberWithDefaults

`func NewWidgetMemberWithDefaults() *WidgetMember`

NewWidgetMemberWithDefaults instantiates a new WidgetMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WidgetMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WidgetMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WidgetMember) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *WidgetMember) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WidgetMember) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WidgetMember) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDiscriminator

`func (o *WidgetMember) GetDiscriminator() WidgetUserDiscriminator`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *WidgetMember) GetDiscriminatorOk() (*WidgetUserDiscriminator, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *WidgetMember) SetDiscriminator(v WidgetUserDiscriminator)`

SetDiscriminator sets Discriminator field to given value.


### GetAvatar

`func (o *WidgetMember) GetAvatar() nil`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *WidgetMember) GetAvatarOk() (*nil, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *WidgetMember) SetAvatar(v nil)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *WidgetMember) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetStatus

`func (o *WidgetMember) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WidgetMember) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WidgetMember) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAvatarUrl

`func (o *WidgetMember) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *WidgetMember) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *WidgetMember) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetActivity

`func (o *WidgetMember) GetActivity() WidgetActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *WidgetMember) GetActivityOk() (*WidgetActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *WidgetMember) SetActivity(v WidgetActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *WidgetMember) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetDeaf

`func (o *WidgetMember) GetDeaf() bool`

GetDeaf returns the Deaf field if non-nil, zero value otherwise.

### GetDeafOk

`func (o *WidgetMember) GetDeafOk() (*bool, bool)`

GetDeafOk returns a tuple with the Deaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaf

`func (o *WidgetMember) SetDeaf(v bool)`

SetDeaf sets Deaf field to given value.

### HasDeaf

`func (o *WidgetMember) HasDeaf() bool`

HasDeaf returns a boolean if a field has been set.

### GetMute

`func (o *WidgetMember) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *WidgetMember) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *WidgetMember) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *WidgetMember) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetSelfDeaf

`func (o *WidgetMember) GetSelfDeaf() bool`

GetSelfDeaf returns the SelfDeaf field if non-nil, zero value otherwise.

### GetSelfDeafOk

`func (o *WidgetMember) GetSelfDeafOk() (*bool, bool)`

GetSelfDeafOk returns a tuple with the SelfDeaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfDeaf

`func (o *WidgetMember) SetSelfDeaf(v bool)`

SetSelfDeaf sets SelfDeaf field to given value.

### HasSelfDeaf

`func (o *WidgetMember) HasSelfDeaf() bool`

HasSelfDeaf returns a boolean if a field has been set.

### GetSelfMute

`func (o *WidgetMember) GetSelfMute() bool`

GetSelfMute returns the SelfMute field if non-nil, zero value otherwise.

### GetSelfMuteOk

`func (o *WidgetMember) GetSelfMuteOk() (*bool, bool)`

GetSelfMuteOk returns a tuple with the SelfMute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfMute

`func (o *WidgetMember) SetSelfMute(v bool)`

SetSelfMute sets SelfMute field to given value.

### HasSelfMute

`func (o *WidgetMember) HasSelfMute() bool`

HasSelfMute returns a boolean if a field has been set.

### GetSuppress

`func (o *WidgetMember) GetSuppress() bool`

GetSuppress returns the Suppress field if non-nil, zero value otherwise.

### GetSuppressOk

`func (o *WidgetMember) GetSuppressOk() (*bool, bool)`

GetSuppressOk returns a tuple with the Suppress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppress

`func (o *WidgetMember) SetSuppress(v bool)`

SetSuppress sets Suppress field to given value.

### HasSuppress

`func (o *WidgetMember) HasSuppress() bool`

HasSuppress returns a boolean if a field has been set.

### GetChannelId

`func (o *WidgetMember) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *WidgetMember) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *WidgetMember) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *WidgetMember) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


