# GuildPatchRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**VerificationLevel** | Pointer to [**GuildPatchRequestPartialVerificationLevel**](GuildPatchRequestPartialVerificationLevel.md) |  | [optional] 
**DefaultMessageNotifications** | Pointer to [**GuildPatchRequestPartialDefaultMessageNotifications**](GuildPatchRequestPartialDefaultMessageNotifications.md) |  | [optional] 
**ExplicitContentFilter** | Pointer to [**GuildPatchRequestPartialExplicitContentFilter**](GuildPatchRequestPartialExplicitContentFilter.md) |  | [optional] 
**PreferredLocale** | Pointer to [**GuildPatchRequestPartialPreferredLocale**](GuildPatchRequestPartialPreferredLocale.md) |  | [optional] 
**AfkTimeout** | Pointer to [**GuildPatchRequestPartialAfkTimeout**](GuildPatchRequestPartialAfkTimeout.md) |  | [optional] 
**AfkChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**SystemChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Splash** | Pointer to **string** |  | [optional] 
**Banner** | Pointer to **string** |  | [optional] 
**SystemChannelFlags** | Pointer to **int32** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**DiscoverySplash** | Pointer to **string** |  | [optional] 
**HomeHeader** | Pointer to **string** |  | [optional] 
**RulesChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**SafetyAlertsChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**PublicUpdatesChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**PremiumProgressBarEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewGuildPatchRequestPartial

`func NewGuildPatchRequestPartial() *GuildPatchRequestPartial`

NewGuildPatchRequestPartial instantiates a new GuildPatchRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildPatchRequestPartialWithDefaults

`func NewGuildPatchRequestPartialWithDefaults() *GuildPatchRequestPartial`

NewGuildPatchRequestPartialWithDefaults instantiates a new GuildPatchRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GuildPatchRequestPartial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildPatchRequestPartial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildPatchRequestPartial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GuildPatchRequestPartial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GuildPatchRequestPartial) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildPatchRequestPartial) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildPatchRequestPartial) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildPatchRequestPartial) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRegion

`func (o *GuildPatchRequestPartial) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GuildPatchRequestPartial) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GuildPatchRequestPartial) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GuildPatchRequestPartial) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetIcon

`func (o *GuildPatchRequestPartial) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GuildPatchRequestPartial) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GuildPatchRequestPartial) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GuildPatchRequestPartial) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetVerificationLevel

`func (o *GuildPatchRequestPartial) GetVerificationLevel() GuildPatchRequestPartialVerificationLevel`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *GuildPatchRequestPartial) GetVerificationLevelOk() (*GuildPatchRequestPartialVerificationLevel, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *GuildPatchRequestPartial) SetVerificationLevel(v GuildPatchRequestPartialVerificationLevel)`

SetVerificationLevel sets VerificationLevel field to given value.

### HasVerificationLevel

`func (o *GuildPatchRequestPartial) HasVerificationLevel() bool`

HasVerificationLevel returns a boolean if a field has been set.

### GetDefaultMessageNotifications

`func (o *GuildPatchRequestPartial) GetDefaultMessageNotifications() GuildPatchRequestPartialDefaultMessageNotifications`

GetDefaultMessageNotifications returns the DefaultMessageNotifications field if non-nil, zero value otherwise.

### GetDefaultMessageNotificationsOk

`func (o *GuildPatchRequestPartial) GetDefaultMessageNotificationsOk() (*GuildPatchRequestPartialDefaultMessageNotifications, bool)`

GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessageNotifications

`func (o *GuildPatchRequestPartial) SetDefaultMessageNotifications(v GuildPatchRequestPartialDefaultMessageNotifications)`

SetDefaultMessageNotifications sets DefaultMessageNotifications field to given value.

### HasDefaultMessageNotifications

`func (o *GuildPatchRequestPartial) HasDefaultMessageNotifications() bool`

HasDefaultMessageNotifications returns a boolean if a field has been set.

### GetExplicitContentFilter

`func (o *GuildPatchRequestPartial) GetExplicitContentFilter() GuildPatchRequestPartialExplicitContentFilter`

GetExplicitContentFilter returns the ExplicitContentFilter field if non-nil, zero value otherwise.

### GetExplicitContentFilterOk

`func (o *GuildPatchRequestPartial) GetExplicitContentFilterOk() (*GuildPatchRequestPartialExplicitContentFilter, bool)`

GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContentFilter

`func (o *GuildPatchRequestPartial) SetExplicitContentFilter(v GuildPatchRequestPartialExplicitContentFilter)`

SetExplicitContentFilter sets ExplicitContentFilter field to given value.

### HasExplicitContentFilter

`func (o *GuildPatchRequestPartial) HasExplicitContentFilter() bool`

HasExplicitContentFilter returns a boolean if a field has been set.

### GetPreferredLocale

`func (o *GuildPatchRequestPartial) GetPreferredLocale() GuildPatchRequestPartialPreferredLocale`

GetPreferredLocale returns the PreferredLocale field if non-nil, zero value otherwise.

### GetPreferredLocaleOk

`func (o *GuildPatchRequestPartial) GetPreferredLocaleOk() (*GuildPatchRequestPartialPreferredLocale, bool)`

GetPreferredLocaleOk returns a tuple with the PreferredLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLocale

`func (o *GuildPatchRequestPartial) SetPreferredLocale(v GuildPatchRequestPartialPreferredLocale)`

SetPreferredLocale sets PreferredLocale field to given value.

### HasPreferredLocale

`func (o *GuildPatchRequestPartial) HasPreferredLocale() bool`

HasPreferredLocale returns a boolean if a field has been set.

### GetAfkTimeout

`func (o *GuildPatchRequestPartial) GetAfkTimeout() GuildPatchRequestPartialAfkTimeout`

GetAfkTimeout returns the AfkTimeout field if non-nil, zero value otherwise.

### GetAfkTimeoutOk

`func (o *GuildPatchRequestPartial) GetAfkTimeoutOk() (*GuildPatchRequestPartialAfkTimeout, bool)`

GetAfkTimeoutOk returns a tuple with the AfkTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkTimeout

`func (o *GuildPatchRequestPartial) SetAfkTimeout(v GuildPatchRequestPartialAfkTimeout)`

SetAfkTimeout sets AfkTimeout field to given value.

### HasAfkTimeout

`func (o *GuildPatchRequestPartial) HasAfkTimeout() bool`

HasAfkTimeout returns a boolean if a field has been set.

### GetAfkChannelId

`func (o *GuildPatchRequestPartial) GetAfkChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetAfkChannelId returns the AfkChannelId field if non-nil, zero value otherwise.

### GetAfkChannelIdOk

`func (o *GuildPatchRequestPartial) GetAfkChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetAfkChannelIdOk returns a tuple with the AfkChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkChannelId

`func (o *GuildPatchRequestPartial) SetAfkChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetAfkChannelId sets AfkChannelId field to given value.

### HasAfkChannelId

`func (o *GuildPatchRequestPartial) HasAfkChannelId() bool`

HasAfkChannelId returns a boolean if a field has been set.

### GetSystemChannelId

`func (o *GuildPatchRequestPartial) GetSystemChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetSystemChannelId returns the SystemChannelId field if non-nil, zero value otherwise.

### GetSystemChannelIdOk

`func (o *GuildPatchRequestPartial) GetSystemChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetSystemChannelIdOk returns a tuple with the SystemChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelId

`func (o *GuildPatchRequestPartial) SetSystemChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetSystemChannelId sets SystemChannelId field to given value.

### HasSystemChannelId

`func (o *GuildPatchRequestPartial) HasSystemChannelId() bool`

HasSystemChannelId returns a boolean if a field has been set.

### GetSplash

`func (o *GuildPatchRequestPartial) GetSplash() string`

GetSplash returns the Splash field if non-nil, zero value otherwise.

### GetSplashOk

`func (o *GuildPatchRequestPartial) GetSplashOk() (*string, bool)`

GetSplashOk returns a tuple with the Splash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplash

`func (o *GuildPatchRequestPartial) SetSplash(v string)`

SetSplash sets Splash field to given value.

### HasSplash

`func (o *GuildPatchRequestPartial) HasSplash() bool`

HasSplash returns a boolean if a field has been set.

### GetBanner

`func (o *GuildPatchRequestPartial) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *GuildPatchRequestPartial) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *GuildPatchRequestPartial) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *GuildPatchRequestPartial) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetSystemChannelFlags

`func (o *GuildPatchRequestPartial) GetSystemChannelFlags() int32`

GetSystemChannelFlags returns the SystemChannelFlags field if non-nil, zero value otherwise.

### GetSystemChannelFlagsOk

`func (o *GuildPatchRequestPartial) GetSystemChannelFlagsOk() (*int32, bool)`

GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelFlags

`func (o *GuildPatchRequestPartial) SetSystemChannelFlags(v int32)`

SetSystemChannelFlags sets SystemChannelFlags field to given value.

### HasSystemChannelFlags

`func (o *GuildPatchRequestPartial) HasSystemChannelFlags() bool`

HasSystemChannelFlags returns a boolean if a field has been set.

### GetFeatures

`func (o *GuildPatchRequestPartial) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GuildPatchRequestPartial) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GuildPatchRequestPartial) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GuildPatchRequestPartial) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetDiscoverySplash

`func (o *GuildPatchRequestPartial) GetDiscoverySplash() string`

GetDiscoverySplash returns the DiscoverySplash field if non-nil, zero value otherwise.

### GetDiscoverySplashOk

`func (o *GuildPatchRequestPartial) GetDiscoverySplashOk() (*string, bool)`

GetDiscoverySplashOk returns a tuple with the DiscoverySplash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySplash

`func (o *GuildPatchRequestPartial) SetDiscoverySplash(v string)`

SetDiscoverySplash sets DiscoverySplash field to given value.

### HasDiscoverySplash

`func (o *GuildPatchRequestPartial) HasDiscoverySplash() bool`

HasDiscoverySplash returns a boolean if a field has been set.

### GetHomeHeader

`func (o *GuildPatchRequestPartial) GetHomeHeader() string`

GetHomeHeader returns the HomeHeader field if non-nil, zero value otherwise.

### GetHomeHeaderOk

`func (o *GuildPatchRequestPartial) GetHomeHeaderOk() (*string, bool)`

GetHomeHeaderOk returns a tuple with the HomeHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeHeader

`func (o *GuildPatchRequestPartial) SetHomeHeader(v string)`

SetHomeHeader sets HomeHeader field to given value.

### HasHomeHeader

`func (o *GuildPatchRequestPartial) HasHomeHeader() bool`

HasHomeHeader returns a boolean if a field has been set.

### GetRulesChannelId

`func (o *GuildPatchRequestPartial) GetRulesChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetRulesChannelId returns the RulesChannelId field if non-nil, zero value otherwise.

### GetRulesChannelIdOk

`func (o *GuildPatchRequestPartial) GetRulesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetRulesChannelIdOk returns a tuple with the RulesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesChannelId

`func (o *GuildPatchRequestPartial) SetRulesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetRulesChannelId sets RulesChannelId field to given value.

### HasRulesChannelId

`func (o *GuildPatchRequestPartial) HasRulesChannelId() bool`

HasRulesChannelId returns a boolean if a field has been set.

### GetSafetyAlertsChannelId

`func (o *GuildPatchRequestPartial) GetSafetyAlertsChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetSafetyAlertsChannelId returns the SafetyAlertsChannelId field if non-nil, zero value otherwise.

### GetSafetyAlertsChannelIdOk

`func (o *GuildPatchRequestPartial) GetSafetyAlertsChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetSafetyAlertsChannelIdOk returns a tuple with the SafetyAlertsChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyAlertsChannelId

`func (o *GuildPatchRequestPartial) SetSafetyAlertsChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetSafetyAlertsChannelId sets SafetyAlertsChannelId field to given value.

### HasSafetyAlertsChannelId

`func (o *GuildPatchRequestPartial) HasSafetyAlertsChannelId() bool`

HasSafetyAlertsChannelId returns a boolean if a field has been set.

### GetPublicUpdatesChannelId

`func (o *GuildPatchRequestPartial) GetPublicUpdatesChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetPublicUpdatesChannelId returns the PublicUpdatesChannelId field if non-nil, zero value otherwise.

### GetPublicUpdatesChannelIdOk

`func (o *GuildPatchRequestPartial) GetPublicUpdatesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetPublicUpdatesChannelIdOk returns a tuple with the PublicUpdatesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUpdatesChannelId

`func (o *GuildPatchRequestPartial) SetPublicUpdatesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetPublicUpdatesChannelId sets PublicUpdatesChannelId field to given value.

### HasPublicUpdatesChannelId

`func (o *GuildPatchRequestPartial) HasPublicUpdatesChannelId() bool`

HasPublicUpdatesChannelId returns a boolean if a field has been set.

### GetPremiumProgressBarEnabled

`func (o *GuildPatchRequestPartial) GetPremiumProgressBarEnabled() bool`

GetPremiumProgressBarEnabled returns the PremiumProgressBarEnabled field if non-nil, zero value otherwise.

### GetPremiumProgressBarEnabledOk

`func (o *GuildPatchRequestPartial) GetPremiumProgressBarEnabledOk() (*bool, bool)`

GetPremiumProgressBarEnabledOk returns a tuple with the PremiumProgressBarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumProgressBarEnabled

`func (o *GuildPatchRequestPartial) SetPremiumProgressBarEnabled(v bool)`

SetPremiumProgressBarEnabled sets PremiumProgressBarEnabled field to given value.

### HasPremiumProgressBarEnabled

`func (o *GuildPatchRequestPartial) HasPremiumProgressBarEnabled() bool`

HasPremiumProgressBarEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


