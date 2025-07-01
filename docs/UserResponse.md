# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**Discriminator** | **string** |  | 
**PublicFlags** | **int32** |  | 
**Flags** | **int64** |  | 
**Bot** | Pointer to **NullableBool** |  | [optional] 
**System** | Pointer to **NullableBool** |  | [optional] 
**Banner** | Pointer to **NullableString** |  | [optional] 
**AccentColor** | Pointer to **NullableInt32** |  | [optional] 
**GlobalName** | Pointer to **NullableString** |  | [optional] 
**AvatarDecorationData** | Pointer to [**NullableUserAvatarDecorationResponse**](UserAvatarDecorationResponse.md) |  | [optional] 
**Collectibles** | Pointer to [**NullableUserCollectiblesResponse**](UserCollectiblesResponse.md) |  | [optional] 
**PrimaryGuild** | Pointer to [**NullableUserPrimaryGuildResponse**](UserPrimaryGuildResponse.md) |  | [optional] 

## Methods

### NewUserResponse

`func NewUserResponse(id string, username string, discriminator string, publicFlags int32, flags int64, ) *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *UserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAvatar

`func (o *UserResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *UserResponse) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *UserResponse) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetDiscriminator

`func (o *UserResponse) GetDiscriminator() string`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *UserResponse) GetDiscriminatorOk() (*string, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *UserResponse) SetDiscriminator(v string)`

SetDiscriminator sets Discriminator field to given value.


### GetPublicFlags

`func (o *UserResponse) GetPublicFlags() int32`

GetPublicFlags returns the PublicFlags field if non-nil, zero value otherwise.

### GetPublicFlagsOk

`func (o *UserResponse) GetPublicFlagsOk() (*int32, bool)`

GetPublicFlagsOk returns a tuple with the PublicFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFlags

`func (o *UserResponse) SetPublicFlags(v int32)`

SetPublicFlags sets PublicFlags field to given value.


### GetFlags

`func (o *UserResponse) GetFlags() int64`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UserResponse) GetFlagsOk() (*int64, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UserResponse) SetFlags(v int64)`

SetFlags sets Flags field to given value.


### GetBot

`func (o *UserResponse) GetBot() bool`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *UserResponse) GetBotOk() (*bool, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *UserResponse) SetBot(v bool)`

SetBot sets Bot field to given value.

### HasBot

`func (o *UserResponse) HasBot() bool`

HasBot returns a boolean if a field has been set.

### SetBotNil

`func (o *UserResponse) SetBotNil(b bool)`

 SetBotNil sets the value for Bot to be an explicit nil

### UnsetBot
`func (o *UserResponse) UnsetBot()`

UnsetBot ensures that no value is present for Bot, not even an explicit nil
### GetSystem

`func (o *UserResponse) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *UserResponse) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *UserResponse) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *UserResponse) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *UserResponse) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *UserResponse) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetBanner

`func (o *UserResponse) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *UserResponse) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *UserResponse) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *UserResponse) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### SetBannerNil

`func (o *UserResponse) SetBannerNil(b bool)`

 SetBannerNil sets the value for Banner to be an explicit nil

### UnsetBanner
`func (o *UserResponse) UnsetBanner()`

UnsetBanner ensures that no value is present for Banner, not even an explicit nil
### GetAccentColor

`func (o *UserResponse) GetAccentColor() int32`

GetAccentColor returns the AccentColor field if non-nil, zero value otherwise.

### GetAccentColorOk

`func (o *UserResponse) GetAccentColorOk() (*int32, bool)`

GetAccentColorOk returns a tuple with the AccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentColor

`func (o *UserResponse) SetAccentColor(v int32)`

SetAccentColor sets AccentColor field to given value.

### HasAccentColor

`func (o *UserResponse) HasAccentColor() bool`

HasAccentColor returns a boolean if a field has been set.

### SetAccentColorNil

`func (o *UserResponse) SetAccentColorNil(b bool)`

 SetAccentColorNil sets the value for AccentColor to be an explicit nil

### UnsetAccentColor
`func (o *UserResponse) UnsetAccentColor()`

UnsetAccentColor ensures that no value is present for AccentColor, not even an explicit nil
### GetGlobalName

`func (o *UserResponse) GetGlobalName() string`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *UserResponse) GetGlobalNameOk() (*string, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *UserResponse) SetGlobalName(v string)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *UserResponse) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### SetGlobalNameNil

`func (o *UserResponse) SetGlobalNameNil(b bool)`

 SetGlobalNameNil sets the value for GlobalName to be an explicit nil

### UnsetGlobalName
`func (o *UserResponse) UnsetGlobalName()`

UnsetGlobalName ensures that no value is present for GlobalName, not even an explicit nil
### GetAvatarDecorationData

`func (o *UserResponse) GetAvatarDecorationData() UserAvatarDecorationResponse`

GetAvatarDecorationData returns the AvatarDecorationData field if non-nil, zero value otherwise.

### GetAvatarDecorationDataOk

`func (o *UserResponse) GetAvatarDecorationDataOk() (*UserAvatarDecorationResponse, bool)`

GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarDecorationData

`func (o *UserResponse) SetAvatarDecorationData(v UserAvatarDecorationResponse)`

SetAvatarDecorationData sets AvatarDecorationData field to given value.

### HasAvatarDecorationData

`func (o *UserResponse) HasAvatarDecorationData() bool`

HasAvatarDecorationData returns a boolean if a field has been set.

### SetAvatarDecorationDataNil

`func (o *UserResponse) SetAvatarDecorationDataNil(b bool)`

 SetAvatarDecorationDataNil sets the value for AvatarDecorationData to be an explicit nil

### UnsetAvatarDecorationData
`func (o *UserResponse) UnsetAvatarDecorationData()`

UnsetAvatarDecorationData ensures that no value is present for AvatarDecorationData, not even an explicit nil
### GetCollectibles

`func (o *UserResponse) GetCollectibles() UserCollectiblesResponse`

GetCollectibles returns the Collectibles field if non-nil, zero value otherwise.

### GetCollectiblesOk

`func (o *UserResponse) GetCollectiblesOk() (*UserCollectiblesResponse, bool)`

GetCollectiblesOk returns a tuple with the Collectibles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectibles

`func (o *UserResponse) SetCollectibles(v UserCollectiblesResponse)`

SetCollectibles sets Collectibles field to given value.

### HasCollectibles

`func (o *UserResponse) HasCollectibles() bool`

HasCollectibles returns a boolean if a field has been set.

### SetCollectiblesNil

`func (o *UserResponse) SetCollectiblesNil(b bool)`

 SetCollectiblesNil sets the value for Collectibles to be an explicit nil

### UnsetCollectibles
`func (o *UserResponse) UnsetCollectibles()`

UnsetCollectibles ensures that no value is present for Collectibles, not even an explicit nil
### GetPrimaryGuild

`func (o *UserResponse) GetPrimaryGuild() UserPrimaryGuildResponse`

GetPrimaryGuild returns the PrimaryGuild field if non-nil, zero value otherwise.

### GetPrimaryGuildOk

`func (o *UserResponse) GetPrimaryGuildOk() (*UserPrimaryGuildResponse, bool)`

GetPrimaryGuildOk returns a tuple with the PrimaryGuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGuild

`func (o *UserResponse) SetPrimaryGuild(v UserPrimaryGuildResponse)`

SetPrimaryGuild sets PrimaryGuild field to given value.

### HasPrimaryGuild

`func (o *UserResponse) HasPrimaryGuild() bool`

HasPrimaryGuild returns a boolean if a field has been set.

### SetPrimaryGuildNil

`func (o *UserResponse) SetPrimaryGuildNil(b bool)`

 SetPrimaryGuildNil sets the value for PrimaryGuild to be an explicit nil

### UnsetPrimaryGuild
`func (o *UserResponse) UnsetPrimaryGuild()`

UnsetPrimaryGuild ensures that no value is present for PrimaryGuild, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


