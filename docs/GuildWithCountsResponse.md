# GuildWithCountsResponse

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
**ApproximateMemberCount** | Pointer to **int32** |  | [optional] 
**ApproximatePresenceCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewGuildWithCountsResponse

`func NewGuildWithCountsResponse(id string, name string, features []GuildFeatures, ownerId string, region string, afkTimeout AfkTimeouts, systemChannelFlags int32, widgetEnabled bool, verificationLevel VerificationLevels, roles []GuildRoleResponse, defaultMessageNotifications UserNotificationSettings, mfaLevel GuildMFALevel, explicitContentFilter GuildExplicitContentFilterTypes, maxMembers int32, maxStageVideoChannelUsers int32, maxVideoChannelUsers int32, premiumTier PremiumGuildTiers, premiumSubscriptionCount int32, preferredLocale AvailableLocalesEnum, premiumProgressBarEnabled bool, nsfw bool, nsfwLevel GuildNSFWContentLevel, emojis []EmojiResponse, stickers []GuildStickerResponse, ) *GuildWithCountsResponse`

NewGuildWithCountsResponse instantiates a new GuildWithCountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildWithCountsResponseWithDefaults

`func NewGuildWithCountsResponseWithDefaults() *GuildWithCountsResponse`

NewGuildWithCountsResponseWithDefaults instantiates a new GuildWithCountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildWithCountsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildWithCountsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildWithCountsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GuildWithCountsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildWithCountsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildWithCountsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *GuildWithCountsResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GuildWithCountsResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GuildWithCountsResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GuildWithCountsResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetDescription

`func (o *GuildWithCountsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildWithCountsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildWithCountsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildWithCountsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomeHeader

`func (o *GuildWithCountsResponse) GetHomeHeader() string`

GetHomeHeader returns the HomeHeader field if non-nil, zero value otherwise.

### GetHomeHeaderOk

`func (o *GuildWithCountsResponse) GetHomeHeaderOk() (*string, bool)`

GetHomeHeaderOk returns a tuple with the HomeHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeHeader

`func (o *GuildWithCountsResponse) SetHomeHeader(v string)`

SetHomeHeader sets HomeHeader field to given value.

### HasHomeHeader

`func (o *GuildWithCountsResponse) HasHomeHeader() bool`

HasHomeHeader returns a boolean if a field has been set.

### GetSplash

`func (o *GuildWithCountsResponse) GetSplash() string`

GetSplash returns the Splash field if non-nil, zero value otherwise.

### GetSplashOk

`func (o *GuildWithCountsResponse) GetSplashOk() (*string, bool)`

GetSplashOk returns a tuple with the Splash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplash

`func (o *GuildWithCountsResponse) SetSplash(v string)`

SetSplash sets Splash field to given value.

### HasSplash

`func (o *GuildWithCountsResponse) HasSplash() bool`

HasSplash returns a boolean if a field has been set.

### GetDiscoverySplash

`func (o *GuildWithCountsResponse) GetDiscoverySplash() string`

GetDiscoverySplash returns the DiscoverySplash field if non-nil, zero value otherwise.

### GetDiscoverySplashOk

`func (o *GuildWithCountsResponse) GetDiscoverySplashOk() (*string, bool)`

GetDiscoverySplashOk returns a tuple with the DiscoverySplash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySplash

`func (o *GuildWithCountsResponse) SetDiscoverySplash(v string)`

SetDiscoverySplash sets DiscoverySplash field to given value.

### HasDiscoverySplash

`func (o *GuildWithCountsResponse) HasDiscoverySplash() bool`

HasDiscoverySplash returns a boolean if a field has been set.

### GetFeatures

`func (o *GuildWithCountsResponse) GetFeatures() []GuildFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GuildWithCountsResponse) GetFeaturesOk() (*[]GuildFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GuildWithCountsResponse) SetFeatures(v []GuildFeatures)`

SetFeatures sets Features field to given value.


### GetBanner

`func (o *GuildWithCountsResponse) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *GuildWithCountsResponse) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *GuildWithCountsResponse) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *GuildWithCountsResponse) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetOwnerId

`func (o *GuildWithCountsResponse) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GuildWithCountsResponse) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GuildWithCountsResponse) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetApplicationId

`func (o *GuildWithCountsResponse) GetApplicationId() GetEntitlementsSkuIdsParameterOneOfInner`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GuildWithCountsResponse) GetApplicationIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GuildWithCountsResponse) SetApplicationId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GuildWithCountsResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetRegion

`func (o *GuildWithCountsResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GuildWithCountsResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GuildWithCountsResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAfkChannelId

`func (o *GuildWithCountsResponse) GetAfkChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetAfkChannelId returns the AfkChannelId field if non-nil, zero value otherwise.

### GetAfkChannelIdOk

`func (o *GuildWithCountsResponse) GetAfkChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetAfkChannelIdOk returns a tuple with the AfkChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkChannelId

`func (o *GuildWithCountsResponse) SetAfkChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetAfkChannelId sets AfkChannelId field to given value.

### HasAfkChannelId

`func (o *GuildWithCountsResponse) HasAfkChannelId() bool`

HasAfkChannelId returns a boolean if a field has been set.

### GetAfkTimeout

`func (o *GuildWithCountsResponse) GetAfkTimeout() AfkTimeouts`

GetAfkTimeout returns the AfkTimeout field if non-nil, zero value otherwise.

### GetAfkTimeoutOk

`func (o *GuildWithCountsResponse) GetAfkTimeoutOk() (*AfkTimeouts, bool)`

GetAfkTimeoutOk returns a tuple with the AfkTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkTimeout

`func (o *GuildWithCountsResponse) SetAfkTimeout(v AfkTimeouts)`

SetAfkTimeout sets AfkTimeout field to given value.


### GetSystemChannelId

`func (o *GuildWithCountsResponse) GetSystemChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetSystemChannelId returns the SystemChannelId field if non-nil, zero value otherwise.

### GetSystemChannelIdOk

`func (o *GuildWithCountsResponse) GetSystemChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetSystemChannelIdOk returns a tuple with the SystemChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelId

`func (o *GuildWithCountsResponse) SetSystemChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetSystemChannelId sets SystemChannelId field to given value.

### HasSystemChannelId

`func (o *GuildWithCountsResponse) HasSystemChannelId() bool`

HasSystemChannelId returns a boolean if a field has been set.

### GetSystemChannelFlags

`func (o *GuildWithCountsResponse) GetSystemChannelFlags() int32`

GetSystemChannelFlags returns the SystemChannelFlags field if non-nil, zero value otherwise.

### GetSystemChannelFlagsOk

`func (o *GuildWithCountsResponse) GetSystemChannelFlagsOk() (*int32, bool)`

GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelFlags

`func (o *GuildWithCountsResponse) SetSystemChannelFlags(v int32)`

SetSystemChannelFlags sets SystemChannelFlags field to given value.


### GetWidgetEnabled

`func (o *GuildWithCountsResponse) GetWidgetEnabled() bool`

GetWidgetEnabled returns the WidgetEnabled field if non-nil, zero value otherwise.

### GetWidgetEnabledOk

`func (o *GuildWithCountsResponse) GetWidgetEnabledOk() (*bool, bool)`

GetWidgetEnabledOk returns a tuple with the WidgetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetEnabled

`func (o *GuildWithCountsResponse) SetWidgetEnabled(v bool)`

SetWidgetEnabled sets WidgetEnabled field to given value.


### GetWidgetChannelId

`func (o *GuildWithCountsResponse) GetWidgetChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetWidgetChannelId returns the WidgetChannelId field if non-nil, zero value otherwise.

### GetWidgetChannelIdOk

`func (o *GuildWithCountsResponse) GetWidgetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetWidgetChannelIdOk returns a tuple with the WidgetChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetChannelId

`func (o *GuildWithCountsResponse) SetWidgetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetWidgetChannelId sets WidgetChannelId field to given value.

### HasWidgetChannelId

`func (o *GuildWithCountsResponse) HasWidgetChannelId() bool`

HasWidgetChannelId returns a boolean if a field has been set.

### GetVerificationLevel

`func (o *GuildWithCountsResponse) GetVerificationLevel() VerificationLevels`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *GuildWithCountsResponse) GetVerificationLevelOk() (*VerificationLevels, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *GuildWithCountsResponse) SetVerificationLevel(v VerificationLevels)`

SetVerificationLevel sets VerificationLevel field to given value.


### GetRoles

`func (o *GuildWithCountsResponse) GetRoles() []GuildRoleResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GuildWithCountsResponse) GetRolesOk() (*[]GuildRoleResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GuildWithCountsResponse) SetRoles(v []GuildRoleResponse)`

SetRoles sets Roles field to given value.


### GetDefaultMessageNotifications

`func (o *GuildWithCountsResponse) GetDefaultMessageNotifications() UserNotificationSettings`

GetDefaultMessageNotifications returns the DefaultMessageNotifications field if non-nil, zero value otherwise.

### GetDefaultMessageNotificationsOk

`func (o *GuildWithCountsResponse) GetDefaultMessageNotificationsOk() (*UserNotificationSettings, bool)`

GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessageNotifications

`func (o *GuildWithCountsResponse) SetDefaultMessageNotifications(v UserNotificationSettings)`

SetDefaultMessageNotifications sets DefaultMessageNotifications field to given value.


### GetMfaLevel

`func (o *GuildWithCountsResponse) GetMfaLevel() GuildMFALevel`

GetMfaLevel returns the MfaLevel field if non-nil, zero value otherwise.

### GetMfaLevelOk

`func (o *GuildWithCountsResponse) GetMfaLevelOk() (*GuildMFALevel, bool)`

GetMfaLevelOk returns a tuple with the MfaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaLevel

`func (o *GuildWithCountsResponse) SetMfaLevel(v GuildMFALevel)`

SetMfaLevel sets MfaLevel field to given value.


### GetExplicitContentFilter

`func (o *GuildWithCountsResponse) GetExplicitContentFilter() GuildExplicitContentFilterTypes`

GetExplicitContentFilter returns the ExplicitContentFilter field if non-nil, zero value otherwise.

### GetExplicitContentFilterOk

`func (o *GuildWithCountsResponse) GetExplicitContentFilterOk() (*GuildExplicitContentFilterTypes, bool)`

GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContentFilter

`func (o *GuildWithCountsResponse) SetExplicitContentFilter(v GuildExplicitContentFilterTypes)`

SetExplicitContentFilter sets ExplicitContentFilter field to given value.


### GetMaxPresences

`func (o *GuildWithCountsResponse) GetMaxPresences() int32`

GetMaxPresences returns the MaxPresences field if non-nil, zero value otherwise.

### GetMaxPresencesOk

`func (o *GuildWithCountsResponse) GetMaxPresencesOk() (*int32, bool)`

GetMaxPresencesOk returns a tuple with the MaxPresences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPresences

`func (o *GuildWithCountsResponse) SetMaxPresences(v int32)`

SetMaxPresences sets MaxPresences field to given value.

### HasMaxPresences

`func (o *GuildWithCountsResponse) HasMaxPresences() bool`

HasMaxPresences returns a boolean if a field has been set.

### GetMaxMembers

`func (o *GuildWithCountsResponse) GetMaxMembers() int32`

GetMaxMembers returns the MaxMembers field if non-nil, zero value otherwise.

### GetMaxMembersOk

`func (o *GuildWithCountsResponse) GetMaxMembersOk() (*int32, bool)`

GetMaxMembersOk returns a tuple with the MaxMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMembers

`func (o *GuildWithCountsResponse) SetMaxMembers(v int32)`

SetMaxMembers sets MaxMembers field to given value.


### GetMaxStageVideoChannelUsers

`func (o *GuildWithCountsResponse) GetMaxStageVideoChannelUsers() int32`

GetMaxStageVideoChannelUsers returns the MaxStageVideoChannelUsers field if non-nil, zero value otherwise.

### GetMaxStageVideoChannelUsersOk

`func (o *GuildWithCountsResponse) GetMaxStageVideoChannelUsersOk() (*int32, bool)`

GetMaxStageVideoChannelUsersOk returns a tuple with the MaxStageVideoChannelUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStageVideoChannelUsers

`func (o *GuildWithCountsResponse) SetMaxStageVideoChannelUsers(v int32)`

SetMaxStageVideoChannelUsers sets MaxStageVideoChannelUsers field to given value.


### GetMaxVideoChannelUsers

`func (o *GuildWithCountsResponse) GetMaxVideoChannelUsers() int32`

GetMaxVideoChannelUsers returns the MaxVideoChannelUsers field if non-nil, zero value otherwise.

### GetMaxVideoChannelUsersOk

`func (o *GuildWithCountsResponse) GetMaxVideoChannelUsersOk() (*int32, bool)`

GetMaxVideoChannelUsersOk returns a tuple with the MaxVideoChannelUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVideoChannelUsers

`func (o *GuildWithCountsResponse) SetMaxVideoChannelUsers(v int32)`

SetMaxVideoChannelUsers sets MaxVideoChannelUsers field to given value.


### GetVanityUrlCode

`func (o *GuildWithCountsResponse) GetVanityUrlCode() string`

GetVanityUrlCode returns the VanityUrlCode field if non-nil, zero value otherwise.

### GetVanityUrlCodeOk

`func (o *GuildWithCountsResponse) GetVanityUrlCodeOk() (*string, bool)`

GetVanityUrlCodeOk returns a tuple with the VanityUrlCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVanityUrlCode

`func (o *GuildWithCountsResponse) SetVanityUrlCode(v string)`

SetVanityUrlCode sets VanityUrlCode field to given value.

### HasVanityUrlCode

`func (o *GuildWithCountsResponse) HasVanityUrlCode() bool`

HasVanityUrlCode returns a boolean if a field has been set.

### GetPremiumTier

`func (o *GuildWithCountsResponse) GetPremiumTier() PremiumGuildTiers`

GetPremiumTier returns the PremiumTier field if non-nil, zero value otherwise.

### GetPremiumTierOk

`func (o *GuildWithCountsResponse) GetPremiumTierOk() (*PremiumGuildTiers, bool)`

GetPremiumTierOk returns a tuple with the PremiumTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumTier

`func (o *GuildWithCountsResponse) SetPremiumTier(v PremiumGuildTiers)`

SetPremiumTier sets PremiumTier field to given value.


### GetPremiumSubscriptionCount

`func (o *GuildWithCountsResponse) GetPremiumSubscriptionCount() int32`

GetPremiumSubscriptionCount returns the PremiumSubscriptionCount field if non-nil, zero value otherwise.

### GetPremiumSubscriptionCountOk

`func (o *GuildWithCountsResponse) GetPremiumSubscriptionCountOk() (*int32, bool)`

GetPremiumSubscriptionCountOk returns a tuple with the PremiumSubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSubscriptionCount

`func (o *GuildWithCountsResponse) SetPremiumSubscriptionCount(v int32)`

SetPremiumSubscriptionCount sets PremiumSubscriptionCount field to given value.


### GetPreferredLocale

`func (o *GuildWithCountsResponse) GetPreferredLocale() AvailableLocalesEnum`

GetPreferredLocale returns the PreferredLocale field if non-nil, zero value otherwise.

### GetPreferredLocaleOk

`func (o *GuildWithCountsResponse) GetPreferredLocaleOk() (*AvailableLocalesEnum, bool)`

GetPreferredLocaleOk returns a tuple with the PreferredLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLocale

`func (o *GuildWithCountsResponse) SetPreferredLocale(v AvailableLocalesEnum)`

SetPreferredLocale sets PreferredLocale field to given value.


### GetRulesChannelId

`func (o *GuildWithCountsResponse) GetRulesChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetRulesChannelId returns the RulesChannelId field if non-nil, zero value otherwise.

### GetRulesChannelIdOk

`func (o *GuildWithCountsResponse) GetRulesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetRulesChannelIdOk returns a tuple with the RulesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesChannelId

`func (o *GuildWithCountsResponse) SetRulesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetRulesChannelId sets RulesChannelId field to given value.

### HasRulesChannelId

`func (o *GuildWithCountsResponse) HasRulesChannelId() bool`

HasRulesChannelId returns a boolean if a field has been set.

### GetSafetyAlertsChannelId

`func (o *GuildWithCountsResponse) GetSafetyAlertsChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetSafetyAlertsChannelId returns the SafetyAlertsChannelId field if non-nil, zero value otherwise.

### GetSafetyAlertsChannelIdOk

`func (o *GuildWithCountsResponse) GetSafetyAlertsChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetSafetyAlertsChannelIdOk returns a tuple with the SafetyAlertsChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyAlertsChannelId

`func (o *GuildWithCountsResponse) SetSafetyAlertsChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetSafetyAlertsChannelId sets SafetyAlertsChannelId field to given value.

### HasSafetyAlertsChannelId

`func (o *GuildWithCountsResponse) HasSafetyAlertsChannelId() bool`

HasSafetyAlertsChannelId returns a boolean if a field has been set.

### GetPublicUpdatesChannelId

`func (o *GuildWithCountsResponse) GetPublicUpdatesChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetPublicUpdatesChannelId returns the PublicUpdatesChannelId field if non-nil, zero value otherwise.

### GetPublicUpdatesChannelIdOk

`func (o *GuildWithCountsResponse) GetPublicUpdatesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetPublicUpdatesChannelIdOk returns a tuple with the PublicUpdatesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUpdatesChannelId

`func (o *GuildWithCountsResponse) SetPublicUpdatesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetPublicUpdatesChannelId sets PublicUpdatesChannelId field to given value.

### HasPublicUpdatesChannelId

`func (o *GuildWithCountsResponse) HasPublicUpdatesChannelId() bool`

HasPublicUpdatesChannelId returns a boolean if a field has been set.

### GetPremiumProgressBarEnabled

`func (o *GuildWithCountsResponse) GetPremiumProgressBarEnabled() bool`

GetPremiumProgressBarEnabled returns the PremiumProgressBarEnabled field if non-nil, zero value otherwise.

### GetPremiumProgressBarEnabledOk

`func (o *GuildWithCountsResponse) GetPremiumProgressBarEnabledOk() (*bool, bool)`

GetPremiumProgressBarEnabledOk returns a tuple with the PremiumProgressBarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumProgressBarEnabled

`func (o *GuildWithCountsResponse) SetPremiumProgressBarEnabled(v bool)`

SetPremiumProgressBarEnabled sets PremiumProgressBarEnabled field to given value.


### GetNsfw

`func (o *GuildWithCountsResponse) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *GuildWithCountsResponse) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *GuildWithCountsResponse) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.


### GetNsfwLevel

`func (o *GuildWithCountsResponse) GetNsfwLevel() GuildNSFWContentLevel`

GetNsfwLevel returns the NsfwLevel field if non-nil, zero value otherwise.

### GetNsfwLevelOk

`func (o *GuildWithCountsResponse) GetNsfwLevelOk() (*GuildNSFWContentLevel, bool)`

GetNsfwLevelOk returns a tuple with the NsfwLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfwLevel

`func (o *GuildWithCountsResponse) SetNsfwLevel(v GuildNSFWContentLevel)`

SetNsfwLevel sets NsfwLevel field to given value.


### GetEmojis

`func (o *GuildWithCountsResponse) GetEmojis() []EmojiResponse`

GetEmojis returns the Emojis field if non-nil, zero value otherwise.

### GetEmojisOk

`func (o *GuildWithCountsResponse) GetEmojisOk() (*[]EmojiResponse, bool)`

GetEmojisOk returns a tuple with the Emojis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojis

`func (o *GuildWithCountsResponse) SetEmojis(v []EmojiResponse)`

SetEmojis sets Emojis field to given value.


### GetStickers

`func (o *GuildWithCountsResponse) GetStickers() []GuildStickerResponse`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *GuildWithCountsResponse) GetStickersOk() (*[]GuildStickerResponse, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *GuildWithCountsResponse) SetStickers(v []GuildStickerResponse)`

SetStickers sets Stickers field to given value.


### GetApproximateMemberCount

`func (o *GuildWithCountsResponse) GetApproximateMemberCount() int32`

GetApproximateMemberCount returns the ApproximateMemberCount field if non-nil, zero value otherwise.

### GetApproximateMemberCountOk

`func (o *GuildWithCountsResponse) GetApproximateMemberCountOk() (*int32, bool)`

GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMemberCount

`func (o *GuildWithCountsResponse) SetApproximateMemberCount(v int32)`

SetApproximateMemberCount sets ApproximateMemberCount field to given value.

### HasApproximateMemberCount

`func (o *GuildWithCountsResponse) HasApproximateMemberCount() bool`

HasApproximateMemberCount returns a boolean if a field has been set.

### GetApproximatePresenceCount

`func (o *GuildWithCountsResponse) GetApproximatePresenceCount() int32`

GetApproximatePresenceCount returns the ApproximatePresenceCount field if non-nil, zero value otherwise.

### GetApproximatePresenceCountOk

`func (o *GuildWithCountsResponse) GetApproximatePresenceCountOk() (*int32, bool)`

GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximatePresenceCount

`func (o *GuildWithCountsResponse) SetApproximatePresenceCount(v int32)`

SetApproximatePresenceCount sets ApproximatePresenceCount field to given value.

### HasApproximatePresenceCount

`func (o *GuildWithCountsResponse) HasApproximatePresenceCount() bool`

HasApproximatePresenceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


