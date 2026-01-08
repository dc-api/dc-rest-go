# GuildResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Icon** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HomeHeader** | Pointer to **string** |  | [optional] 
**Splash** | Pointer to **string** |  | [optional] 
**DiscoverySplash** | Pointer to **string** |  | [optional] 
**Features** | [**[]GuildFeatures**](GuildFeatures.md) |  | 
**Banner** | Pointer to **string** |  | [optional] 
**OwnerId** | **string** |  | 
**ApplicationId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Region** | **string** |  | 
**AfkChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**AfkTimeout** | [**AfkTimeouts**](AfkTimeouts.md) |  | 
**SystemChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**SystemChannelFlags** | **int32** |  | 
**WidgetEnabled** | **bool** |  | 
**WidgetChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**VerificationLevel** | [**VerificationLevels**](VerificationLevels.md) |  | 
**Roles** | [**[]GuildRoleResponse**](GuildRoleResponse.md) |  | 
**DefaultMessageNotifications** | [**UserNotificationSettings**](UserNotificationSettings.md) |  | 
**MfaLevel** | [**GuildMFALevel**](GuildMFALevel.md) |  | 
**ExplicitContentFilter** | [**GuildExplicitContentFilterTypes**](GuildExplicitContentFilterTypes.md) |  | 
**MaxPresences** | Pointer to **int32** |  | [optional] 
**MaxMembers** | **int32** |  | 
**MaxStageVideoChannelUsers** | **int32** |  | 
**MaxVideoChannelUsers** | **int32** |  | 
**VanityUrlCode** | Pointer to **string** |  | [optional] 
**PremiumTier** | [**PremiumGuildTiers**](PremiumGuildTiers.md) |  | 
**PremiumSubscriptionCount** | **int32** |  | 
**PreferredLocale** | [**AvailableLocalesEnum**](AvailableLocalesEnum.md) |  | 
**RulesChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**SafetyAlertsChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**PublicUpdatesChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**PremiumProgressBarEnabled** | **bool** |  | 
**Nsfw** | **bool** |  | 
**NsfwLevel** | [**GuildNSFWContentLevel**](GuildNSFWContentLevel.md) |  | 
**Emojis** | [**[]EmojiResponse**](EmojiResponse.md) |  | 
**Stickers** | [**[]GuildStickerResponse**](GuildStickerResponse.md) |  | 

## Methods

### NewGuildResponse

`func NewGuildResponse(id string, name string, features []GuildFeatures, ownerId string, region string, afkTimeout AfkTimeouts, systemChannelFlags int32, widgetEnabled bool, verificationLevel VerificationLevels, roles []GuildRoleResponse, defaultMessageNotifications UserNotificationSettings, mfaLevel GuildMFALevel, explicitContentFilter GuildExplicitContentFilterTypes, maxMembers int32, maxStageVideoChannelUsers int32, maxVideoChannelUsers int32, premiumTier PremiumGuildTiers, premiumSubscriptionCount int32, preferredLocale AvailableLocalesEnum, premiumProgressBarEnabled bool, nsfw bool, nsfwLevel GuildNSFWContentLevel, emojis []EmojiResponse, stickers []GuildStickerResponse, ) *GuildResponse`

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

### GetFeatures

`func (o *GuildResponse) GetFeatures() []GuildFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GuildResponse) GetFeaturesOk() (*[]GuildFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GuildResponse) SetFeatures(v []GuildFeatures)`

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

`func (o *GuildResponse) GetApplicationId() GetEntitlementsSkuIdsParameterOneOfInner`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GuildResponse) GetApplicationIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GuildResponse) SetApplicationId(v GetEntitlementsSkuIdsParameterOneOfInner)`

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

`func (o *GuildResponse) GetAfkChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetAfkChannelId returns the AfkChannelId field if non-nil, zero value otherwise.

### GetAfkChannelIdOk

`func (o *GuildResponse) GetAfkChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetAfkChannelIdOk returns a tuple with the AfkChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkChannelId

`func (o *GuildResponse) SetAfkChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetAfkChannelId sets AfkChannelId field to given value.

### HasAfkChannelId

`func (o *GuildResponse) HasAfkChannelId() bool`

HasAfkChannelId returns a boolean if a field has been set.

### GetAfkTimeout

`func (o *GuildResponse) GetAfkTimeout() AfkTimeouts`

GetAfkTimeout returns the AfkTimeout field if non-nil, zero value otherwise.

### GetAfkTimeoutOk

`func (o *GuildResponse) GetAfkTimeoutOk() (*AfkTimeouts, bool)`

GetAfkTimeoutOk returns a tuple with the AfkTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkTimeout

`func (o *GuildResponse) SetAfkTimeout(v AfkTimeouts)`

SetAfkTimeout sets AfkTimeout field to given value.


### GetSystemChannelId

`func (o *GuildResponse) GetSystemChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetSystemChannelId returns the SystemChannelId field if non-nil, zero value otherwise.

### GetSystemChannelIdOk

`func (o *GuildResponse) GetSystemChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetSystemChannelIdOk returns a tuple with the SystemChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelId

`func (o *GuildResponse) SetSystemChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

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

`func (o *GuildResponse) GetWidgetChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetWidgetChannelId returns the WidgetChannelId field if non-nil, zero value otherwise.

### GetWidgetChannelIdOk

`func (o *GuildResponse) GetWidgetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetWidgetChannelIdOk returns a tuple with the WidgetChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetChannelId

`func (o *GuildResponse) SetWidgetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetWidgetChannelId sets WidgetChannelId field to given value.

### HasWidgetChannelId

`func (o *GuildResponse) HasWidgetChannelId() bool`

HasWidgetChannelId returns a boolean if a field has been set.

### GetVerificationLevel

`func (o *GuildResponse) GetVerificationLevel() VerificationLevels`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *GuildResponse) GetVerificationLevelOk() (*VerificationLevels, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *GuildResponse) SetVerificationLevel(v VerificationLevels)`

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

`func (o *GuildResponse) GetDefaultMessageNotifications() UserNotificationSettings`

GetDefaultMessageNotifications returns the DefaultMessageNotifications field if non-nil, zero value otherwise.

### GetDefaultMessageNotificationsOk

`func (o *GuildResponse) GetDefaultMessageNotificationsOk() (*UserNotificationSettings, bool)`

GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessageNotifications

`func (o *GuildResponse) SetDefaultMessageNotifications(v UserNotificationSettings)`

SetDefaultMessageNotifications sets DefaultMessageNotifications field to given value.


### GetMfaLevel

`func (o *GuildResponse) GetMfaLevel() GuildMFALevel`

GetMfaLevel returns the MfaLevel field if non-nil, zero value otherwise.

### GetMfaLevelOk

`func (o *GuildResponse) GetMfaLevelOk() (*GuildMFALevel, bool)`

GetMfaLevelOk returns a tuple with the MfaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaLevel

`func (o *GuildResponse) SetMfaLevel(v GuildMFALevel)`

SetMfaLevel sets MfaLevel field to given value.


### GetExplicitContentFilter

`func (o *GuildResponse) GetExplicitContentFilter() GuildExplicitContentFilterTypes`

GetExplicitContentFilter returns the ExplicitContentFilter field if non-nil, zero value otherwise.

### GetExplicitContentFilterOk

`func (o *GuildResponse) GetExplicitContentFilterOk() (*GuildExplicitContentFilterTypes, bool)`

GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContentFilter

`func (o *GuildResponse) SetExplicitContentFilter(v GuildExplicitContentFilterTypes)`

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

### GetPremiumTier

`func (o *GuildResponse) GetPremiumTier() PremiumGuildTiers`

GetPremiumTier returns the PremiumTier field if non-nil, zero value otherwise.

### GetPremiumTierOk

`func (o *GuildResponse) GetPremiumTierOk() (*PremiumGuildTiers, bool)`

GetPremiumTierOk returns a tuple with the PremiumTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumTier

`func (o *GuildResponse) SetPremiumTier(v PremiumGuildTiers)`

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

`func (o *GuildResponse) GetPreferredLocale() AvailableLocalesEnum`

GetPreferredLocale returns the PreferredLocale field if non-nil, zero value otherwise.

### GetPreferredLocaleOk

`func (o *GuildResponse) GetPreferredLocaleOk() (*AvailableLocalesEnum, bool)`

GetPreferredLocaleOk returns a tuple with the PreferredLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLocale

`func (o *GuildResponse) SetPreferredLocale(v AvailableLocalesEnum)`

SetPreferredLocale sets PreferredLocale field to given value.


### GetRulesChannelId

`func (o *GuildResponse) GetRulesChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetRulesChannelId returns the RulesChannelId field if non-nil, zero value otherwise.

### GetRulesChannelIdOk

`func (o *GuildResponse) GetRulesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetRulesChannelIdOk returns a tuple with the RulesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesChannelId

`func (o *GuildResponse) SetRulesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetRulesChannelId sets RulesChannelId field to given value.

### HasRulesChannelId

`func (o *GuildResponse) HasRulesChannelId() bool`

HasRulesChannelId returns a boolean if a field has been set.

### GetSafetyAlertsChannelId

`func (o *GuildResponse) GetSafetyAlertsChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetSafetyAlertsChannelId returns the SafetyAlertsChannelId field if non-nil, zero value otherwise.

### GetSafetyAlertsChannelIdOk

`func (o *GuildResponse) GetSafetyAlertsChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetSafetyAlertsChannelIdOk returns a tuple with the SafetyAlertsChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyAlertsChannelId

`func (o *GuildResponse) SetSafetyAlertsChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetSafetyAlertsChannelId sets SafetyAlertsChannelId field to given value.

### HasSafetyAlertsChannelId

`func (o *GuildResponse) HasSafetyAlertsChannelId() bool`

HasSafetyAlertsChannelId returns a boolean if a field has been set.

### GetPublicUpdatesChannelId

`func (o *GuildResponse) GetPublicUpdatesChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetPublicUpdatesChannelId returns the PublicUpdatesChannelId field if non-nil, zero value otherwise.

### GetPublicUpdatesChannelIdOk

`func (o *GuildResponse) GetPublicUpdatesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetPublicUpdatesChannelIdOk returns a tuple with the PublicUpdatesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUpdatesChannelId

`func (o *GuildResponse) SetPublicUpdatesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

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

`func (o *GuildResponse) GetNsfwLevel() GuildNSFWContentLevel`

GetNsfwLevel returns the NsfwLevel field if non-nil, zero value otherwise.

### GetNsfwLevelOk

`func (o *GuildResponse) GetNsfwLevelOk() (*GuildNSFWContentLevel, bool)`

GetNsfwLevelOk returns a tuple with the NsfwLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfwLevel

`func (o *GuildResponse) SetNsfwLevel(v GuildNSFWContentLevel)`

SetNsfwLevel sets NsfwLevel field to given value.


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


