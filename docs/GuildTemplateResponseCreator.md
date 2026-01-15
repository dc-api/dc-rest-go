# GuildTemplateResponseCreator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Avatar** | Pointer to **string** |  | [optional] 
**Discriminator** | **string** |  | 
**PublicFlags** | **int32** |  | 
**Flags** | **int64** |  | 
**Bot** | Pointer to **bool** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**Banner** | Pointer to **string** |  | [optional] 
**AccentColor** | Pointer to **int32** |  | [optional] 
**GlobalName** | Pointer to **string** |  | [optional] 
**AvatarDecorationData** | Pointer to [**BasicGuildMemberResponseAvatarDecorationData**](BasicGuildMemberResponseAvatarDecorationData.md) |  | [optional] 
**Collectibles** | Pointer to [**BasicGuildMemberResponseCollectibles**](BasicGuildMemberResponseCollectibles.md) |  | [optional] 
**PrimaryGuild** | Pointer to [**UserPIIResponsePrimaryGuild**](UserPIIResponsePrimaryGuild.md) |  | [optional] 

## Methods

### NewGuildTemplateResponseCreator

`func NewGuildTemplateResponseCreator(id string, username string, discriminator string, publicFlags int32, flags int64, ) *GuildTemplateResponseCreator`

NewGuildTemplateResponseCreator instantiates a new GuildTemplateResponseCreator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildTemplateResponseCreatorWithDefaults

`func NewGuildTemplateResponseCreatorWithDefaults() *GuildTemplateResponseCreator`

NewGuildTemplateResponseCreatorWithDefaults instantiates a new GuildTemplateResponseCreator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildTemplateResponseCreator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildTemplateResponseCreator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildTemplateResponseCreator) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *GuildTemplateResponseCreator) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GuildTemplateResponseCreator) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GuildTemplateResponseCreator) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAvatar

`func (o *GuildTemplateResponseCreator) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GuildTemplateResponseCreator) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GuildTemplateResponseCreator) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *GuildTemplateResponseCreator) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetDiscriminator

`func (o *GuildTemplateResponseCreator) GetDiscriminator() string`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *GuildTemplateResponseCreator) GetDiscriminatorOk() (*string, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *GuildTemplateResponseCreator) SetDiscriminator(v string)`

SetDiscriminator sets Discriminator field to given value.


### GetPublicFlags

`func (o *GuildTemplateResponseCreator) GetPublicFlags() int32`

GetPublicFlags returns the PublicFlags field if non-nil, zero value otherwise.

### GetPublicFlagsOk

`func (o *GuildTemplateResponseCreator) GetPublicFlagsOk() (*int32, bool)`

GetPublicFlagsOk returns a tuple with the PublicFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFlags

`func (o *GuildTemplateResponseCreator) SetPublicFlags(v int32)`

SetPublicFlags sets PublicFlags field to given value.


### GetFlags

`func (o *GuildTemplateResponseCreator) GetFlags() int64`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GuildTemplateResponseCreator) GetFlagsOk() (*int64, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GuildTemplateResponseCreator) SetFlags(v int64)`

SetFlags sets Flags field to given value.


### GetBot

`func (o *GuildTemplateResponseCreator) GetBot() bool`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *GuildTemplateResponseCreator) GetBotOk() (*bool, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *GuildTemplateResponseCreator) SetBot(v bool)`

SetBot sets Bot field to given value.

### HasBot

`func (o *GuildTemplateResponseCreator) HasBot() bool`

HasBot returns a boolean if a field has been set.

### GetSystem

`func (o *GuildTemplateResponseCreator) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *GuildTemplateResponseCreator) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *GuildTemplateResponseCreator) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *GuildTemplateResponseCreator) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetBanner

`func (o *GuildTemplateResponseCreator) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *GuildTemplateResponseCreator) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *GuildTemplateResponseCreator) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *GuildTemplateResponseCreator) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetAccentColor

`func (o *GuildTemplateResponseCreator) GetAccentColor() int32`

GetAccentColor returns the AccentColor field if non-nil, zero value otherwise.

### GetAccentColorOk

`func (o *GuildTemplateResponseCreator) GetAccentColorOk() (*int32, bool)`

GetAccentColorOk returns a tuple with the AccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentColor

`func (o *GuildTemplateResponseCreator) SetAccentColor(v int32)`

SetAccentColor sets AccentColor field to given value.

### HasAccentColor

`func (o *GuildTemplateResponseCreator) HasAccentColor() bool`

HasAccentColor returns a boolean if a field has been set.

### GetGlobalName

`func (o *GuildTemplateResponseCreator) GetGlobalName() string`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *GuildTemplateResponseCreator) GetGlobalNameOk() (*string, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *GuildTemplateResponseCreator) SetGlobalName(v string)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *GuildTemplateResponseCreator) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### GetAvatarDecorationData

`func (o *GuildTemplateResponseCreator) GetAvatarDecorationData() BasicGuildMemberResponseAvatarDecorationData`

GetAvatarDecorationData returns the AvatarDecorationData field if non-nil, zero value otherwise.

### GetAvatarDecorationDataOk

`func (o *GuildTemplateResponseCreator) GetAvatarDecorationDataOk() (*BasicGuildMemberResponseAvatarDecorationData, bool)`

GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarDecorationData

`func (o *GuildTemplateResponseCreator) SetAvatarDecorationData(v BasicGuildMemberResponseAvatarDecorationData)`

SetAvatarDecorationData sets AvatarDecorationData field to given value.

### HasAvatarDecorationData

`func (o *GuildTemplateResponseCreator) HasAvatarDecorationData() bool`

HasAvatarDecorationData returns a boolean if a field has been set.

### GetCollectibles

`func (o *GuildTemplateResponseCreator) GetCollectibles() BasicGuildMemberResponseCollectibles`

GetCollectibles returns the Collectibles field if non-nil, zero value otherwise.

### GetCollectiblesOk

`func (o *GuildTemplateResponseCreator) GetCollectiblesOk() (*BasicGuildMemberResponseCollectibles, bool)`

GetCollectiblesOk returns a tuple with the Collectibles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectibles

`func (o *GuildTemplateResponseCreator) SetCollectibles(v BasicGuildMemberResponseCollectibles)`

SetCollectibles sets Collectibles field to given value.

### HasCollectibles

`func (o *GuildTemplateResponseCreator) HasCollectibles() bool`

HasCollectibles returns a boolean if a field has been set.

### GetPrimaryGuild

`func (o *GuildTemplateResponseCreator) GetPrimaryGuild() UserPIIResponsePrimaryGuild`

GetPrimaryGuild returns the PrimaryGuild field if non-nil, zero value otherwise.

### GetPrimaryGuildOk

`func (o *GuildTemplateResponseCreator) GetPrimaryGuildOk() (*UserPIIResponsePrimaryGuild, bool)`

GetPrimaryGuildOk returns a tuple with the PrimaryGuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGuild

`func (o *GuildTemplateResponseCreator) SetPrimaryGuild(v UserPIIResponsePrimaryGuild)`

SetPrimaryGuild sets PrimaryGuild field to given value.

### HasPrimaryGuild

`func (o *GuildTemplateResponseCreator) HasPrimaryGuild() bool`

HasPrimaryGuild returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


