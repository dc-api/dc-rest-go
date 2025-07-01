# GuildPatchRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**VerificationLevel** | Pointer to **int32** |  | [optional] 
**DefaultMessageNotifications** | Pointer to **int32** |  | [optional] 
**ExplicitContentFilter** | Pointer to **int32** |  | [optional] 
**PreferredLocale** | Pointer to **string** |  | [optional] 
**AfkTimeout** | Pointer to **NullableInt32** |  | [optional] 
**AfkChannelId** | Pointer to **string** |  | [optional] 
**SystemChannelId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Splash** | Pointer to **NullableString** |  | [optional] 
**Banner** | Pointer to **NullableString** |  | [optional] 
**SystemChannelFlags** | Pointer to **NullableInt32** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**DiscoverySplash** | Pointer to **NullableString** |  | [optional] 
**HomeHeader** | Pointer to **NullableString** |  | [optional] 
**RulesChannelId** | Pointer to **string** |  | [optional] 
**SafetyAlertsChannelId** | Pointer to **string** |  | [optional] 
**PublicUpdatesChannelId** | Pointer to **string** |  | [optional] 
**PremiumProgressBarEnabled** | Pointer to **NullableBool** |  | [optional] 

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

### SetDescriptionNil

`func (o *GuildPatchRequestPartial) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuildPatchRequestPartial) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### SetRegionNil

`func (o *GuildPatchRequestPartial) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *GuildPatchRequestPartial) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
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

### SetIconNil

`func (o *GuildPatchRequestPartial) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *GuildPatchRequestPartial) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetVerificationLevel

`func (o *GuildPatchRequestPartial) GetVerificationLevel() int32`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *GuildPatchRequestPartial) GetVerificationLevelOk() (*int32, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *GuildPatchRequestPartial) SetVerificationLevel(v int32)`

SetVerificationLevel sets VerificationLevel field to given value.

### HasVerificationLevel

`func (o *GuildPatchRequestPartial) HasVerificationLevel() bool`

HasVerificationLevel returns a boolean if a field has been set.

### GetDefaultMessageNotifications

`func (o *GuildPatchRequestPartial) GetDefaultMessageNotifications() int32`

GetDefaultMessageNotifications returns the DefaultMessageNotifications field if non-nil, zero value otherwise.

### GetDefaultMessageNotificationsOk

`func (o *GuildPatchRequestPartial) GetDefaultMessageNotificationsOk() (*int32, bool)`

GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessageNotifications

`func (o *GuildPatchRequestPartial) SetDefaultMessageNotifications(v int32)`

SetDefaultMessageNotifications sets DefaultMessageNotifications field to given value.

### HasDefaultMessageNotifications

`func (o *GuildPatchRequestPartial) HasDefaultMessageNotifications() bool`

HasDefaultMessageNotifications returns a boolean if a field has been set.

### GetExplicitContentFilter

`func (o *GuildPatchRequestPartial) GetExplicitContentFilter() int32`

GetExplicitContentFilter returns the ExplicitContentFilter field if non-nil, zero value otherwise.

### GetExplicitContentFilterOk

`func (o *GuildPatchRequestPartial) GetExplicitContentFilterOk() (*int32, bool)`

GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContentFilter

`func (o *GuildPatchRequestPartial) SetExplicitContentFilter(v int32)`

SetExplicitContentFilter sets ExplicitContentFilter field to given value.

### HasExplicitContentFilter

`func (o *GuildPatchRequestPartial) HasExplicitContentFilter() bool`

HasExplicitContentFilter returns a boolean if a field has been set.

### GetPreferredLocale

`func (o *GuildPatchRequestPartial) GetPreferredLocale() string`

GetPreferredLocale returns the PreferredLocale field if non-nil, zero value otherwise.

### GetPreferredLocaleOk

`func (o *GuildPatchRequestPartial) GetPreferredLocaleOk() (*string, bool)`

GetPreferredLocaleOk returns a tuple with the PreferredLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLocale

`func (o *GuildPatchRequestPartial) SetPreferredLocale(v string)`

SetPreferredLocale sets PreferredLocale field to given value.

### HasPreferredLocale

`func (o *GuildPatchRequestPartial) HasPreferredLocale() bool`

HasPreferredLocale returns a boolean if a field has been set.

### GetAfkTimeout

`func (o *GuildPatchRequestPartial) GetAfkTimeout() int32`

GetAfkTimeout returns the AfkTimeout field if non-nil, zero value otherwise.

### GetAfkTimeoutOk

`func (o *GuildPatchRequestPartial) GetAfkTimeoutOk() (*int32, bool)`

GetAfkTimeoutOk returns a tuple with the AfkTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkTimeout

`func (o *GuildPatchRequestPartial) SetAfkTimeout(v int32)`

SetAfkTimeout sets AfkTimeout field to given value.

### HasAfkTimeout

`func (o *GuildPatchRequestPartial) HasAfkTimeout() bool`

HasAfkTimeout returns a boolean if a field has been set.

### SetAfkTimeoutNil

`func (o *GuildPatchRequestPartial) SetAfkTimeoutNil(b bool)`

 SetAfkTimeoutNil sets the value for AfkTimeout to be an explicit nil

### UnsetAfkTimeout
`func (o *GuildPatchRequestPartial) UnsetAfkTimeout()`

UnsetAfkTimeout ensures that no value is present for AfkTimeout, not even an explicit nil
### GetAfkChannelId

`func (o *GuildPatchRequestPartial) GetAfkChannelId() string`

GetAfkChannelId returns the AfkChannelId field if non-nil, zero value otherwise.

### GetAfkChannelIdOk

`func (o *GuildPatchRequestPartial) GetAfkChannelIdOk() (*string, bool)`

GetAfkChannelIdOk returns a tuple with the AfkChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkChannelId

`func (o *GuildPatchRequestPartial) SetAfkChannelId(v string)`

SetAfkChannelId sets AfkChannelId field to given value.

### HasAfkChannelId

`func (o *GuildPatchRequestPartial) HasAfkChannelId() bool`

HasAfkChannelId returns a boolean if a field has been set.

### GetSystemChannelId

`func (o *GuildPatchRequestPartial) GetSystemChannelId() string`

GetSystemChannelId returns the SystemChannelId field if non-nil, zero value otherwise.

### GetSystemChannelIdOk

`func (o *GuildPatchRequestPartial) GetSystemChannelIdOk() (*string, bool)`

GetSystemChannelIdOk returns a tuple with the SystemChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelId

`func (o *GuildPatchRequestPartial) SetSystemChannelId(v string)`

SetSystemChannelId sets SystemChannelId field to given value.

### HasSystemChannelId

`func (o *GuildPatchRequestPartial) HasSystemChannelId() bool`

HasSystemChannelId returns a boolean if a field has been set.

### GetOwnerId

`func (o *GuildPatchRequestPartial) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GuildPatchRequestPartial) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GuildPatchRequestPartial) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *GuildPatchRequestPartial) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

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

### SetSplashNil

`func (o *GuildPatchRequestPartial) SetSplashNil(b bool)`

 SetSplashNil sets the value for Splash to be an explicit nil

### UnsetSplash
`func (o *GuildPatchRequestPartial) UnsetSplash()`

UnsetSplash ensures that no value is present for Splash, not even an explicit nil
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

### SetBannerNil

`func (o *GuildPatchRequestPartial) SetBannerNil(b bool)`

 SetBannerNil sets the value for Banner to be an explicit nil

### UnsetBanner
`func (o *GuildPatchRequestPartial) UnsetBanner()`

UnsetBanner ensures that no value is present for Banner, not even an explicit nil
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

### SetSystemChannelFlagsNil

`func (o *GuildPatchRequestPartial) SetSystemChannelFlagsNil(b bool)`

 SetSystemChannelFlagsNil sets the value for SystemChannelFlags to be an explicit nil

### UnsetSystemChannelFlags
`func (o *GuildPatchRequestPartial) UnsetSystemChannelFlags()`

UnsetSystemChannelFlags ensures that no value is present for SystemChannelFlags, not even an explicit nil
### GetFeatures

`func (o *GuildPatchRequestPartial) GetFeatures() []*string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GuildPatchRequestPartial) GetFeaturesOk() (*[]*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GuildPatchRequestPartial) SetFeatures(v []*string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GuildPatchRequestPartial) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *GuildPatchRequestPartial) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *GuildPatchRequestPartial) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
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

### SetDiscoverySplashNil

`func (o *GuildPatchRequestPartial) SetDiscoverySplashNil(b bool)`

 SetDiscoverySplashNil sets the value for DiscoverySplash to be an explicit nil

### UnsetDiscoverySplash
`func (o *GuildPatchRequestPartial) UnsetDiscoverySplash()`

UnsetDiscoverySplash ensures that no value is present for DiscoverySplash, not even an explicit nil
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

### SetHomeHeaderNil

`func (o *GuildPatchRequestPartial) SetHomeHeaderNil(b bool)`

 SetHomeHeaderNil sets the value for HomeHeader to be an explicit nil

### UnsetHomeHeader
`func (o *GuildPatchRequestPartial) UnsetHomeHeader()`

UnsetHomeHeader ensures that no value is present for HomeHeader, not even an explicit nil
### GetRulesChannelId

`func (o *GuildPatchRequestPartial) GetRulesChannelId() string`

GetRulesChannelId returns the RulesChannelId field if non-nil, zero value otherwise.

### GetRulesChannelIdOk

`func (o *GuildPatchRequestPartial) GetRulesChannelIdOk() (*string, bool)`

GetRulesChannelIdOk returns a tuple with the RulesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesChannelId

`func (o *GuildPatchRequestPartial) SetRulesChannelId(v string)`

SetRulesChannelId sets RulesChannelId field to given value.

### HasRulesChannelId

`func (o *GuildPatchRequestPartial) HasRulesChannelId() bool`

HasRulesChannelId returns a boolean if a field has been set.

### GetSafetyAlertsChannelId

`func (o *GuildPatchRequestPartial) GetSafetyAlertsChannelId() string`

GetSafetyAlertsChannelId returns the SafetyAlertsChannelId field if non-nil, zero value otherwise.

### GetSafetyAlertsChannelIdOk

`func (o *GuildPatchRequestPartial) GetSafetyAlertsChannelIdOk() (*string, bool)`

GetSafetyAlertsChannelIdOk returns a tuple with the SafetyAlertsChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyAlertsChannelId

`func (o *GuildPatchRequestPartial) SetSafetyAlertsChannelId(v string)`

SetSafetyAlertsChannelId sets SafetyAlertsChannelId field to given value.

### HasSafetyAlertsChannelId

`func (o *GuildPatchRequestPartial) HasSafetyAlertsChannelId() bool`

HasSafetyAlertsChannelId returns a boolean if a field has been set.

### GetPublicUpdatesChannelId

`func (o *GuildPatchRequestPartial) GetPublicUpdatesChannelId() string`

GetPublicUpdatesChannelId returns the PublicUpdatesChannelId field if non-nil, zero value otherwise.

### GetPublicUpdatesChannelIdOk

`func (o *GuildPatchRequestPartial) GetPublicUpdatesChannelIdOk() (*string, bool)`

GetPublicUpdatesChannelIdOk returns a tuple with the PublicUpdatesChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUpdatesChannelId

`func (o *GuildPatchRequestPartial) SetPublicUpdatesChannelId(v string)`

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

### SetPremiumProgressBarEnabledNil

`func (o *GuildPatchRequestPartial) SetPremiumProgressBarEnabledNil(b bool)`

 SetPremiumProgressBarEnabledNil sets the value for PremiumProgressBarEnabled to be an explicit nil

### UnsetPremiumProgressBarEnabled
`func (o *GuildPatchRequestPartial) UnsetPremiumProgressBarEnabled()`

UnsetPremiumProgressBarEnabled ensures that no value is present for PremiumProgressBarEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


