# UserPIIResponse

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
**MfaEnabled** | **bool** |  | 
**Locale** | **string** |  | 
**PremiumType** | Pointer to **int32** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Verified** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUserPIIResponse

`func NewUserPIIResponse(id string, username string, discriminator string, publicFlags int32, flags int64, mfaEnabled bool, locale string, ) *UserPIIResponse`

NewUserPIIResponse instantiates a new UserPIIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPIIResponseWithDefaults

`func NewUserPIIResponseWithDefaults() *UserPIIResponse`

NewUserPIIResponseWithDefaults instantiates a new UserPIIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserPIIResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPIIResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPIIResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *UserPIIResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPIIResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPIIResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAvatar

`func (o *UserPIIResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserPIIResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserPIIResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserPIIResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *UserPIIResponse) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *UserPIIResponse) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetDiscriminator

`func (o *UserPIIResponse) GetDiscriminator() string`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *UserPIIResponse) GetDiscriminatorOk() (*string, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *UserPIIResponse) SetDiscriminator(v string)`

SetDiscriminator sets Discriminator field to given value.


### GetPublicFlags

`func (o *UserPIIResponse) GetPublicFlags() int32`

GetPublicFlags returns the PublicFlags field if non-nil, zero value otherwise.

### GetPublicFlagsOk

`func (o *UserPIIResponse) GetPublicFlagsOk() (*int32, bool)`

GetPublicFlagsOk returns a tuple with the PublicFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFlags

`func (o *UserPIIResponse) SetPublicFlags(v int32)`

SetPublicFlags sets PublicFlags field to given value.


### GetFlags

`func (o *UserPIIResponse) GetFlags() int64`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UserPIIResponse) GetFlagsOk() (*int64, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UserPIIResponse) SetFlags(v int64)`

SetFlags sets Flags field to given value.


### GetBot

`func (o *UserPIIResponse) GetBot() bool`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *UserPIIResponse) GetBotOk() (*bool, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *UserPIIResponse) SetBot(v bool)`

SetBot sets Bot field to given value.

### HasBot

`func (o *UserPIIResponse) HasBot() bool`

HasBot returns a boolean if a field has been set.

### SetBotNil

`func (o *UserPIIResponse) SetBotNil(b bool)`

 SetBotNil sets the value for Bot to be an explicit nil

### UnsetBot
`func (o *UserPIIResponse) UnsetBot()`

UnsetBot ensures that no value is present for Bot, not even an explicit nil
### GetSystem

`func (o *UserPIIResponse) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *UserPIIResponse) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *UserPIIResponse) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *UserPIIResponse) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *UserPIIResponse) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *UserPIIResponse) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetBanner

`func (o *UserPIIResponse) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *UserPIIResponse) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *UserPIIResponse) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *UserPIIResponse) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### SetBannerNil

`func (o *UserPIIResponse) SetBannerNil(b bool)`

 SetBannerNil sets the value for Banner to be an explicit nil

### UnsetBanner
`func (o *UserPIIResponse) UnsetBanner()`

UnsetBanner ensures that no value is present for Banner, not even an explicit nil
### GetAccentColor

`func (o *UserPIIResponse) GetAccentColor() int32`

GetAccentColor returns the AccentColor field if non-nil, zero value otherwise.

### GetAccentColorOk

`func (o *UserPIIResponse) GetAccentColorOk() (*int32, bool)`

GetAccentColorOk returns a tuple with the AccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentColor

`func (o *UserPIIResponse) SetAccentColor(v int32)`

SetAccentColor sets AccentColor field to given value.

### HasAccentColor

`func (o *UserPIIResponse) HasAccentColor() bool`

HasAccentColor returns a boolean if a field has been set.

### SetAccentColorNil

`func (o *UserPIIResponse) SetAccentColorNil(b bool)`

 SetAccentColorNil sets the value for AccentColor to be an explicit nil

### UnsetAccentColor
`func (o *UserPIIResponse) UnsetAccentColor()`

UnsetAccentColor ensures that no value is present for AccentColor, not even an explicit nil
### GetGlobalName

`func (o *UserPIIResponse) GetGlobalName() string`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *UserPIIResponse) GetGlobalNameOk() (*string, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *UserPIIResponse) SetGlobalName(v string)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *UserPIIResponse) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### SetGlobalNameNil

`func (o *UserPIIResponse) SetGlobalNameNil(b bool)`

 SetGlobalNameNil sets the value for GlobalName to be an explicit nil

### UnsetGlobalName
`func (o *UserPIIResponse) UnsetGlobalName()`

UnsetGlobalName ensures that no value is present for GlobalName, not even an explicit nil
### GetAvatarDecorationData

`func (o *UserPIIResponse) GetAvatarDecorationData() UserAvatarDecorationResponse`

GetAvatarDecorationData returns the AvatarDecorationData field if non-nil, zero value otherwise.

### GetAvatarDecorationDataOk

`func (o *UserPIIResponse) GetAvatarDecorationDataOk() (*UserAvatarDecorationResponse, bool)`

GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarDecorationData

`func (o *UserPIIResponse) SetAvatarDecorationData(v UserAvatarDecorationResponse)`

SetAvatarDecorationData sets AvatarDecorationData field to given value.

### HasAvatarDecorationData

`func (o *UserPIIResponse) HasAvatarDecorationData() bool`

HasAvatarDecorationData returns a boolean if a field has been set.

### SetAvatarDecorationDataNil

`func (o *UserPIIResponse) SetAvatarDecorationDataNil(b bool)`

 SetAvatarDecorationDataNil sets the value for AvatarDecorationData to be an explicit nil

### UnsetAvatarDecorationData
`func (o *UserPIIResponse) UnsetAvatarDecorationData()`

UnsetAvatarDecorationData ensures that no value is present for AvatarDecorationData, not even an explicit nil
### GetCollectibles

`func (o *UserPIIResponse) GetCollectibles() UserCollectiblesResponse`

GetCollectibles returns the Collectibles field if non-nil, zero value otherwise.

### GetCollectiblesOk

`func (o *UserPIIResponse) GetCollectiblesOk() (*UserCollectiblesResponse, bool)`

GetCollectiblesOk returns a tuple with the Collectibles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectibles

`func (o *UserPIIResponse) SetCollectibles(v UserCollectiblesResponse)`

SetCollectibles sets Collectibles field to given value.

### HasCollectibles

`func (o *UserPIIResponse) HasCollectibles() bool`

HasCollectibles returns a boolean if a field has been set.

### SetCollectiblesNil

`func (o *UserPIIResponse) SetCollectiblesNil(b bool)`

 SetCollectiblesNil sets the value for Collectibles to be an explicit nil

### UnsetCollectibles
`func (o *UserPIIResponse) UnsetCollectibles()`

UnsetCollectibles ensures that no value is present for Collectibles, not even an explicit nil
### GetPrimaryGuild

`func (o *UserPIIResponse) GetPrimaryGuild() UserPrimaryGuildResponse`

GetPrimaryGuild returns the PrimaryGuild field if non-nil, zero value otherwise.

### GetPrimaryGuildOk

`func (o *UserPIIResponse) GetPrimaryGuildOk() (*UserPrimaryGuildResponse, bool)`

GetPrimaryGuildOk returns a tuple with the PrimaryGuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGuild

`func (o *UserPIIResponse) SetPrimaryGuild(v UserPrimaryGuildResponse)`

SetPrimaryGuild sets PrimaryGuild field to given value.

### HasPrimaryGuild

`func (o *UserPIIResponse) HasPrimaryGuild() bool`

HasPrimaryGuild returns a boolean if a field has been set.

### SetPrimaryGuildNil

`func (o *UserPIIResponse) SetPrimaryGuildNil(b bool)`

 SetPrimaryGuildNil sets the value for PrimaryGuild to be an explicit nil

### UnsetPrimaryGuild
`func (o *UserPIIResponse) UnsetPrimaryGuild()`

UnsetPrimaryGuild ensures that no value is present for PrimaryGuild, not even an explicit nil
### GetMfaEnabled

`func (o *UserPIIResponse) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *UserPIIResponse) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *UserPIIResponse) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.


### GetLocale

`func (o *UserPIIResponse) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserPIIResponse) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserPIIResponse) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetPremiumType

`func (o *UserPIIResponse) GetPremiumType() int32`

GetPremiumType returns the PremiumType field if non-nil, zero value otherwise.

### GetPremiumTypeOk

`func (o *UserPIIResponse) GetPremiumTypeOk() (*int32, bool)`

GetPremiumTypeOk returns a tuple with the PremiumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumType

`func (o *UserPIIResponse) SetPremiumType(v int32)`

SetPremiumType sets PremiumType field to given value.

### HasPremiumType

`func (o *UserPIIResponse) HasPremiumType() bool`

HasPremiumType returns a boolean if a field has been set.

### GetEmail

`func (o *UserPIIResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPIIResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPIIResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserPIIResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserPIIResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserPIIResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetVerified

`func (o *UserPIIResponse) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserPIIResponse) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserPIIResponse) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UserPIIResponse) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### SetVerifiedNil

`func (o *UserPIIResponse) SetVerifiedNil(b bool)`

 SetVerifiedNil sets the value for Verified to be an explicit nil

### UnsetVerified
`func (o *UserPIIResponse) UnsetVerified()`

UnsetVerified ensures that no value is present for Verified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


