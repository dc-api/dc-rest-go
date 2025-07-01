# GuildResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**HomeHeader** | Pointer to **NullableString** |  | [optional] 
**Splash** | Pointer to **NullableString** |  | [optional] 
**DiscoverySplash** | Pointer to **NullableString** |  | [optional] 
**Features** | **[]string** |  | 
**Banner** | Pointer to **NullableString** |  | [optional] 
**OwnerId** | **string** |  | 
**ApplicationId** | Pointer to **string** |  | [optional] 
**Region** | **string** |  | 
**AfkChannelId** | Pointer to **string** |  | [optional] 
**AfkTimeout** | **NullableInt32** |  | 
**SystemChannelId** | Pointer to **string** |  | [optional] 
**SystemChannelFlags** | **int32** |  | 
**WidgetEnabled** | **bool** |  | 
**WidgetChannelId** | Pointer to **string** |  | [optional] 
**VerificationLevel** | **int32** |  | 
**Roles** | [**[]GuildRoleResponse**](GuildRoleResponse.md) |  | 
**DefaultMessageNotifications** | **int32** |  | 
**MfaLevel** | **int32** |  | 
**ExplicitContentFilter** | **int32** |  | 
**MaxPresences** | Pointer to **NullableInt32** |  | [optional] 
**MaxMembers** | Pointer to **NullableInt32** |  | [optional] 
**MaxStageVideoChannelUsers** | Pointer to **NullableInt32** |  | [optional] 
**MaxVideoChannelUsers** | Pointer to **NullableInt32** |  | [optional] 
**VanityUrlCode** | Pointer to **NullableString** |  | [optional] 
**PremiumTier** | **int32** |  | 
**PremiumSubscriptionCount** | **int32** |  | 
**PreferredLocale** | **string** |  | 
**RulesChannelId** | Pointer to **string** |  | [optional] 
**SafetyAlertsChannelId** | Pointer to **string** |  | [optional] 
**PublicUpdatesChannelId** | Pointer to **string** |  | [optional] 
**PremiumProgressBarEnabled** | **bool** |  | 
**Nsfw** | **bool** |  | 
**NsfwLevel** | **NullableInt32** |  | 
**Emojis** | [**[]EmojiResponse**](EmojiResponse.md) |  | 
**Stickers** | [**[]GuildStickerResponse**](GuildStickerResponse.md) |  | 

## Methods

### NewGuildResponse

`func NewGuildResponse(id string, name string, features []string, ownerId string, region string, afkTimeout NullableInt32, systemChannelFlags int32, widgetEnabled bool, verificationLevel int32, roles []GuildRoleResponse, defaultMessageNotifications int32, mfaLevel int32, explicitContentFilter int32, premiumTier int32, premiumSubscriptionCount int32, preferredLocale string, premiumProgressBarEnabled bool, nsfw bool, nsfwLevel NullableInt32, emojis []EmojiResponse, stickers []GuildStickerResponse, ) *GuildResponse`

NewGuildResponse instantiates a new GuildResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildResponseWithDefaults

`func NewGuildResponseWithDefaults() *GuildResponse`

NewGuildResponseWithDefaults instantiates a new GuildResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GuildResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *GuildResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GuildResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GuildResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GuildResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *GuildResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *GuildResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *GuildResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GuildResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuildResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHomeHeader

`func (o *GuildResponse) GetHomeHeader() string`

GetHomeHeader returns the HomeHeader field if non-nil, zero value otherwise.

### GetHomeHeaderOk

`func (o *GuildResponse) GetHomeHeaderOk() (*string, bool)`

GetHomeHeaderOk returns a tuple with the HomeHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeHeader

`func (o *GuildResponse) SetHomeHeader(v string)`

SetHomeHeader sets HomeHeader field to given value.

### HasHomeHeader

`func (o *GuildResponse) HasHomeHeader() bool`

HasHomeHeader returns a boolean if a field has been set.

### SetHomeHeaderNil

`func (o *GuildResponse) SetHomeHeaderNil(b bool)`

 SetHomeHeaderNil sets the value for HomeHeader to be an explicit nil

### UnsetHomeHeader
`func (o *GuildResponse) UnsetHomeHeader()`

UnsetHomeHeader ensures that no value is present for HomeHeader, not even an explicit nil
### GetSplash

`func (o *GuildResponse) GetSplash() string`

GetSplash returns the Splash field if non-nil, zero value otherwise.

### GetSplashOk

`func (o *GuildResponse) GetSplashOk() (*string, bool)`

GetSplashOk returns a tuple with the Splash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplash

`func (o *GuildResponse) SetSplash(v string)`

SetSplash sets Splash field to given value.

### HasSplash

`func (o *GuildResponse) HasSplash() bool`

HasSplash returns a boolean if a field has been set.

### SetSplashNil

`func (o *GuildResponse) SetSplashNil(b bool)`

 SetSplashNil sets the value for Splash to be an explicit nil

### UnsetSplash
`func (o *GuildResponse) UnsetSplash()`

UnsetSplash ensures that no value is present for Splash, not even an explicit nil
### GetDiscoverySplash

`func (o *GuildResponse) GetDiscoverySplash() string`

GetDiscoverySplash returns the DiscoverySplash field if non-nil, zero value otherwise.

### GetDiscoverySplashOk

`func (o *GuildResponse) GetDiscoverySplashOk() (*string, bool)`

GetDiscoverySplashOk returns a tuple with the DiscoverySplash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySplash

`func (o *GuildResponse) SetDiscoverySplash(v string)`

SetDiscoverySplash sets DiscoverySplash field to given value.

### HasDiscoverySplash

`func (o *GuildResponse) HasDiscoverySplash() bool`

HasDiscoverySplash returns a boolean if a field has been set.

### SetDiscoverySplashNil

`func (o *GuildResponse) SetDiscoverySplashNil(b bool)`

 SetDiscoverySplashNil sets the value for DiscoverySplash to be an explicit nil

### UnsetDiscoverySplash
`func (o *GuildResponse) UnsetDiscoverySplash()`

UnsetDiscoverySplash ensures that no value is present for DiscoverySplash, not even an explicit nil
### GetFeatures

`func (o *GuildResponse) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GuildResponse) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GuildResponse) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetBanner

`func (o *GuildResponse) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *GuildResponse) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *GuildResponse) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *GuildResponse) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### SetBannerNil

`func (o *GuildResponse) SetBannerNil(b bool)`

 SetBannerNil sets the value for Banner to be an explicit nil

### UnsetBanner
`func (o *GuildResponse) UnsetBanner()`

UnsetBanner ensures that no value is present for Banner, not even an explicit nil
### GetOwnerId

`func (o *GuildResponse) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GuildResponse) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GuildResponse) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetApplicationId

`func (o *GuildResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GuildResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GuildResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GuildResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetRegion

`func (o *GuildResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GuildResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GuildResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAfkChannelId

`func (o *GuildResponse) GetAfkChannelId() string`

GetAfkChannelId returns the AfkChannelId field if non-nil, zero value otherwise.

### GetAfkChannelIdOk

`func (o *GuildResponse) GetAfkChannelIdOk() (*string, bool)`

GetAfkChannelIdOk returns a tuple with the AfkChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkChannelId

`func (o *GuildResponse) SetAfkChannelId(v string)`

SetAfkChannelId sets AfkChannelId field to given value.

### HasAfkChannelId

`func (o *GuildResponse) HasAfkChannelId() bool`

HasAfkChannelId returns a boolean if a field has been set.

### GetAfkTimeout

`func (o *GuildResponse) GetAfkTimeout() int32`

GetAfkTimeout returns the AfkTimeout field if non-nil, zero value otherwise.

### GetAfkTimeoutOk

`func (o *GuildResponse) GetAfkTimeoutOk() (*int32, bool)`

GetAfkTimeoutOk returns a tuple with the AfkTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkTimeout

`func (o *GuildResponse) SetAfkTimeout(v int32)`

SetAfkTimeout sets AfkTimeout field to given value.


### SetAfkTimeoutNil

`func (o *GuildResponse) SetAfkTimeoutNil(b bool)`

 SetAfkTimeoutNil sets the value for AfkTimeout to be an explicit nil

### UnsetAfkTimeout
`func (o *GuildResponse) UnsetAfkTimeout()`

UnsetAfkTimeout ensures that no value is present for AfkTimeout, not even an explicit nil
### GetSystemChannelId

`func (o *GuildResponse) GetSystemChannelId() string`

GetSystemChannelId returns the SystemChannelId field if non-nil, zero value otherwise.

### GetSystemChannelIdOk

`func (o *GuildResponse) GetSystemChannelIdOk() (*string, bool)`

GetSystemChannelIdOk returns a tuple with the SystemChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelId

`func (o *GuildResponse) SetSystemChannelId(v string)`

SetSystemChannelId sets SystemChannelId field to given value.

### HasSystemChannelId

`func (o *GuildResponse) HasSystemChannelId() bool`

HasSystemChannelId returns a boolean if a field has been set.

### GetSystemChannelFlags

`func (o *GuildResponse) GetSystemChannelFlags() int32`

GetSystemChannelFlags returns the SystemChannelFlags field if non-nil, zero value otherwise.

### GetSystemChannelFlagsOk

`func (o *GuildResponse) GetSystemChannelFlagsOk() (*int32, bool)`

GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelFlags

`func (o *GuildResponse) SetSystemChannelFlags(v int32)`

SetSystemChannelFlags sets SystemChannelFlags field to given value.


### GetWidgetEnabled

`func (o *GuildResponse) GetWidgetEnabled() bool`

GetWidgetEnabled returns the WidgetEnabled field if non-nil, zero value otherwise.

### GetWidgetEnabledOk

`func (o *GuildResponse) GetWidgetEnabledOk() (*bool, bool)`

GetWidgetEnabledOk returns a tuple with the WidgetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetEnabled

`func (o *GuildResponse) SetWidgetEnabled(v bool)`

SetWidgetEnabled sets WidgetEnabled field to given value.


### GetWidgetChannelId

`func (o *GuildResponse) GetWidgetChannelId() string`

GetWidgetChannelId returns the WidgetChannelId field if non-nil, zero value otherwise.

### GetWidgetChannelIdOk

`func (o *GuildResponse) GetWidgetChannelIdOk() (*string, bool)`

GetWidgetChannelIdOk returns a tuple with the WidgetChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetChannelId

`func (o *GuildResponse) SetWidgetChannelId(v string)`

SetWidgetChannelId sets WidgetChannelId field to given value.

### HasWidgetChannelId

`func (o *GuildResponse) HasWidgetChannelId() bool`

HasWidgetChannelId returns a boolean if a field has been set.

### GetVerificationLevel

`func (o *GuildResponse) GetVerificationLevel() int32`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *GuildResponse) GetVerificationLevelOk() (*int32, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *GuildResponse) SetVerificationLevel(v int32)`

SetVerificationLevel sets VerificationLevel field to given value.


### GetRoles

`func (o *GuildResponse) GetRoles() []GuildRoleResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GuildResponse) GetRolesOk() (*[]GuildRoleResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GuildResponse) SetRoles(v []GuildRoleResponse)`

SetRoles sets Roles field to given value.


### GetDefaultMessageNotifications

`func (o *GuildResponse) GetDefaultMessageNotifications() int32`

GetDefaultMessageNotifications returns the DefaultMessageNotifications field if non-nil, zero value otherwise.

### GetDefaultMessageNotificationsOk

`func (o *GuildResponse) GetDefaultMessageNotificationsOk() (*int32, bool)`

GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessageNotifications

`func (o *GuildResponse) SetDefaultMessageNotifications(v int32)`

SetDefaultMessageNotifications sets DefaultMessageNotifications field to given value.


### GetMfaLevel

`func (o *GuildResponse) GetMfaLevel() int32`

GetMfaLevel returns the MfaLevel field if non-nil, zero value otherwise.

### GetMfaLevelOk

`func (o *GuildResponse) GetMfaLevelOk() (*int32, bool)`

GetMfaLevelOk returns a tuple with the MfaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaLevel

`func (o *GuildResponse) SetMfaLevel(v int32)`

SetMfaLevel sets MfaLevel field to given value.


### GetExplicitContentFilter

`func (o *GuildResponse) GetExplicitContentFilter() int32`

GetExplicitContentFilter returns the ExplicitContentFilter field if non-nil, zero value otherwise.

### GetExplicitContentFilterOk

`func (o *GuildResponse) GetExplicitContentFilterOk() (*int32, bool)`

GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContentFilter

`func (o *GuildResponse) SetExplicitContentFilter(v int32)`

SetExplicitContentFilter sets ExplicitContentFilter field to given value.


### GetMaxPresences

`func (o *GuildResponse) GetMaxPresences() int32`

GetMaxPresences returns the MaxPresences field if non-nil, zero value otherwise.

### GetMaxPresencesOk

`func (o *GuildResponse) GetMaxPresencesOk() (*int32, bool)`

GetMaxPresencesOk returns a tuple with the MaxPresences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPresences

`func (o *GuildResponse) SetMaxPresences(v int32)`

SetMaxPresences sets MaxPresences field to given value.

### HasMaxPresences

`func (o *GuildResponse) HasMaxPresences() bool`

HasMaxPresences returns a boolean if a field has been set.

### SetMaxPresencesNil

`func (o *GuildResponse) SetMaxPresencesNil(b bool)`

 SetMaxPresencesNil sets the value for MaxPresences to be an explicit nil

### UnsetMaxPresences
`func (o *GuildResponse) UnsetMaxPresences()`

UnsetMaxPresences ensures that no value is present for MaxPresences, not even an explicit nil
### GetMaxMembers

`func (o *GuildResponse) GetMaxMembers() int32`

GetMaxMembers returns the MaxMembers field if non-nil, zero value otherwise.

### GetMaxMembersOk

`func (o *GuildResponse) GetMaxMembersOk() (*int32, bool)`

GetMaxMembersOk returns a tuple with the MaxMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMembers

`func (o *GuildResponse) SetMaxMembers(v int32)`

SetMaxMembers sets MaxMembers field to given value.

### HasMaxMembers

`func (o *GuildResponse) HasMaxMembers() bool`

HasMaxMembers returns a boolean if a field has been set.

### SetMaxMembersNil

`func (o *GuildResponse) SetMaxMembersNil(b bool)`

 SetMaxMembersNil sets the value for MaxMembers to be an explicit nil

### UnsetMaxMembers
`func (o *GuildResponse) UnsetMaxMembers()`

UnsetMaxMembers ensures that no value is present for MaxMembers, not even an explicit nil
### GetMaxStageVideoChannelUsers

`func (o *GuildResponse) GetMaxStageVideoChannelUsers() int32`

GetMaxStageVideoChannelUsers returns the MaxStageVideoChannelUsers field if non-nil, zero value otherwise.

### GetMaxStageVideoChannelUsersOk

`func (o *GuildResponse) GetMaxStageVideoChannelUsersOk() (*int32, bool)`

GetMaxStageVideoChannelUsersOk returns a tuple with the MaxStageVideoChannelUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStageVideoChannelUsers

`func (o *GuildResponse) SetMaxStageVideoChannelUsers(v int32)`

SetMaxStageVideoChannelUsers sets MaxStageVideoChannelUsers field to given value.

### HasMaxStageVideoChannelUsers

`func (o *GuildResponse) HasMaxStageVideoChannelUsers() bool`

HasMaxStageVideoChannelUsers returns a boolean if a field has been set.

### SetMaxStageVideoChannelUsersNil

`func (o *GuildResponse) SetMaxStageVideoChannelUsersNil(b bool)`

 SetMaxStageVideoChannelUsersNil sets the value for MaxStageVideoChannelUsers to be an explicit nil

### UnsetMaxStageVideoChannelUsers
`func (o *GuildResponse) UnsetMaxStageVideoChannelUsers()`

UnsetMaxStageVideoChannelUsers ensures that no value is present for MaxStageVideoChannelUsers, not even an explicit nil
### GetMaxVideoChannelUsers

`func (o *GuildResponse) GetMaxVideoChannelUsers() int32`

GetMaxVideoChannelUsers returns the MaxVideoChannelUsers field if non-nil, zero value otherwise.

### GetMaxVideoChannelUsersOk

`func (o *GuildResponse) GetMaxVideoChannelUsersOk() (*int32, bool)`

GetMaxVideoChannelUsersOk returns a tuple with the MaxVideoChannelUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVideoChannelUsers

`func (o *GuildResponse) SetMaxVideoChannelUsers(v int32)`

SetMaxVideoChannelUsers sets MaxVideoChannelUsers field to given value.

### HasMaxVideoChannelUsers

`func (o *GuildResponse) HasMaxVideoChannelUsers() bool`

HasMaxVideoChannelUsers returns a boolean if a field has been set.

### SetMaxVideoChannelUsersNil

`func (o *GuildResponse) SetMaxVideoChannelUsersNil(b bool)`

 SetMaxVideoChannelUsersNil sets the value for MaxVideoChannelUsers to be an explicit nil

### UnsetMaxVideoChannelUsers
`func (o *GuildResponse) UnsetMaxVideoChannelUsers()`

UnsetMaxVideoChannelUsers ensures that no value is present for MaxVideoChannelUsers, not even an explicit nil
### GetVanityUrlCode

`func (o *GuildResponse) GetVanityUrlCode() string`

GetVanityUrlCode returns the VanityUrlCode field if non-nil, zero value otherwise.

### GetVanityUrlCodeOk

`func (o *GuildResponse) GetVanityUrlCodeOk() (*string, bool)`

GetVanityUrlCodeOk returns a tuple with the VanityUrlCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVanityUrlCode

`func (o *GuildResponse) SetVanityUrlCode(v string)`

SetVanityUrlCode sets VanityUrlCode field to given value.

### HasVanityUrlCode

`func (o *GuildResponse) HasVanityUrlCode() bool`

HasVanityUrlCode returns a boolean if a field has been set.

### SetVanityUrlCodeNil

`func (o *GuildResponse) SetVanityUrlCodeNil(b bool)`

 SetVanityUrlCodeNil sets the value for VanityUrlCode to be an explicit nil

### UnsetVanityUrlCode
`func (o *GuildResponse) UnsetVanityUrlCode()`

UnsetVanityUrlCode ensures that no value is present for VanityUrlCode, not even an explicit nil
### GetPremiumTier

`func (o *GuildResponse) GetPremiumTier() int32`

GetPremiumTier returns the PremiumTier field if non-nil, zero value otherwise.

### GetPremiumTierOk

`func (o *GuildResponse) GetPremiumTierOk() (*int32, bool)`

GetPremiumTierOk returns a tuple with the PremiumTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumTier

`func (o *GuildResponse) SetPremiumTier(v int32)`

SetPremiumTier sets PremiumTier field to given value.


### GetPremiumSubscriptionCount

`func (o *GuildResponse) GetPremiumSubscriptionCount() int32`

GetPremiumSubscriptionCount returns the PremiumSubscriptionCount field if non-nil, zero value otherwise.

### GetPremiumSubscriptionCountOk

`func (o *GuildResponse) GetPremiumSubscriptionCountOk() (*int32, bool)`

GetPremiumSubscriptionCountOk returns a tuple with the PremiumSubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSubscriptionCount

`func (o *GuildResponse) SetPremiumSubscriptionCount(v int32)`

SetPremiumSubscriptionCount sets PremiumSubscriptionCount field to given value.


### GetPreferredLocale

`func (o *GuildResponse) GetPreferredLocale() string`

GetPreferredLocale returns the PreferredLocale field if non-nil, zero value otherwise.

### GetPreferredLocaleOk

`func (o *GuildResponse) GetPreferredLocaleOk() (*string, bool)`

GetPreferredLocaleOk returns a tuple with the PreferredLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLocale

`func (o *GuildResponse) SetPreferredLocale(v string)`

SetPreferredLocale sets PreferredLocale field to given value.


### GetRulesChannelId

`func (o *GuildResponse) GetRulesChannelId() string`

GetRulesChannelId returns the RulesChannelId field if non-nil, zero value otherwise.

### GetRulesChannelIdOk

`func (o *GuildResponse) GetRulesChannelIdOk() (*string, bool)`

GetRulesChannelIdOk returns a tuple with the RulesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesChannelId

`func (o *GuildResponse) SetRulesChannelId(v string)`

SetRulesChannelId sets RulesChannelId field to given value.

### HasRulesChannelId

`func (o *GuildResponse) HasRulesChannelId() bool`

HasRulesChannelId returns a boolean if a field has been set.

### GetSafetyAlertsChannelId

`func (o *GuildResponse) GetSafetyAlertsChannelId() string`

GetSafetyAlertsChannelId returns the SafetyAlertsChannelId field if non-nil, zero value otherwise.

### GetSafetyAlertsChannelIdOk

`func (o *GuildResponse) GetSafetyAlertsChannelIdOk() (*string, bool)`

GetSafetyAlertsChannelIdOk returns a tuple with the SafetyAlertsChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyAlertsChannelId

`func (o *GuildResponse) SetSafetyAlertsChannelId(v string)`

SetSafetyAlertsChannelId sets SafetyAlertsChannelId field to given value.

### HasSafetyAlertsChannelId

`func (o *GuildResponse) HasSafetyAlertsChannelId() bool`

HasSafetyAlertsChannelId returns a boolean if a field has been set.

### GetPublicUpdatesChannelId

`func (o *GuildResponse) GetPublicUpdatesChannelId() string`

GetPublicUpdatesChannelId returns the PublicUpdatesChannelId field if non-nil, zero value otherwise.

### GetPublicUpdatesChannelIdOk

`func (o *GuildResponse) GetPublicUpdatesChannelIdOk() (*string, bool)`

GetPublicUpdatesChannelIdOk returns a tuple with the PublicUpdatesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUpdatesChannelId

`func (o *GuildResponse) SetPublicUpdatesChannelId(v string)`

SetPublicUpdatesChannelId sets PublicUpdatesChannelId field to given value.

### HasPublicUpdatesChannelId

`func (o *GuildResponse) HasPublicUpdatesChannelId() bool`

HasPublicUpdatesChannelId returns a boolean if a field has been set.

### GetPremiumProgressBarEnabled

`func (o *GuildResponse) GetPremiumProgressBarEnabled() bool`

GetPremiumProgressBarEnabled returns the PremiumProgressBarEnabled field if non-nil, zero value otherwise.

### GetPremiumProgressBarEnabledOk

`func (o *GuildResponse) GetPremiumProgressBarEnabledOk() (*bool, bool)`

GetPremiumProgressBarEnabledOk returns a tuple with the PremiumProgressBarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumProgressBarEnabled

`func (o *GuildResponse) SetPremiumProgressBarEnabled(v bool)`

SetPremiumProgressBarEnabled sets PremiumProgressBarEnabled field to given value.


### GetNsfw

`func (o *GuildResponse) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *GuildResponse) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *GuildResponse) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.


### GetNsfwLevel

`func (o *GuildResponse) GetNsfwLevel() int32`

GetNsfwLevel returns the NsfwLevel field if non-nil, zero value otherwise.

### GetNsfwLevelOk

`func (o *GuildResponse) GetNsfwLevelOk() (*int32, bool)`

GetNsfwLevelOk returns a tuple with the NsfwLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfwLevel

`func (o *GuildResponse) SetNsfwLevel(v int32)`

SetNsfwLevel sets NsfwLevel field to given value.


### SetNsfwLevelNil

`func (o *GuildResponse) SetNsfwLevelNil(b bool)`

 SetNsfwLevelNil sets the value for NsfwLevel to be an explicit nil

### UnsetNsfwLevel
`func (o *GuildResponse) UnsetNsfwLevel()`

UnsetNsfwLevel ensures that no value is present for NsfwLevel, not even an explicit nil
### GetEmojis

`func (o *GuildResponse) GetEmojis() []EmojiResponse`

GetEmojis returns the Emojis field if non-nil, zero value otherwise.

### GetEmojisOk

`func (o *GuildResponse) GetEmojisOk() (*[]EmojiResponse, bool)`

GetEmojisOk returns a tuple with the Emojis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojis

`func (o *GuildResponse) SetEmojis(v []EmojiResponse)`

SetEmojis sets Emojis field to given value.


### GetStickers

`func (o *GuildResponse) GetStickers() []GuildStickerResponse`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *GuildResponse) GetStickersOk() (*[]GuildStickerResponse, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *GuildResponse) SetStickers(v []GuildStickerResponse)`

SetStickers sets Stickers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


