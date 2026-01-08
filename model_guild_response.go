/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
 * 
 * <details>
 * <summary><strong>⚠️ Important Disclaimer & Limitation of Liability</strong></summary>
 * <br>
 * > **IMPORTANT**: This software is provided "as is" without any warranties, express or implied, including but not limited
 * > to warranties of merchantability, fitness for a particular purpose, or non-infringement. The developers, contributors,
 * > and licensors (collectively, "Developers") make no representations regarding the accuracy, completeness, or reliability
 * > of this software or its outputs.
 * > 
 * > This client is not intended to provide financial, investment, tax, or legal advice. It facilitates interaction with the
 * > Discord HTTP API (Preview) service but does not endorse or recommend any financial actions, including the purchase, sale, or holding of
 * > financial instruments (e.g., stocks, bonds, derivatives, cryptocurrencies). Users must consult qualified financial or
 * > legal professionals before making decisions based on this software's outputs.
 * > 
 * > Financial markets are inherently speculative and carry significant risks. Using this software in trading, analysis, or
 * > other financial activities may result in substantial losses, including total loss of capital. The Developers are not
 * > liable for any losses or damages arising from such use. Users assume full responsibility for validating the software's
 * > outputs and ensuring their suitability for intended purposes.
 * > 
 * > This client may rely on third-party data or services (e.g., market feeds, APIs). The Developers do not control or verify
 * > the accuracy of these services and are not liable for any errors, delays, or losses resulting from their use. Users must
 * > comply with third-party terms and conditions.
 * > 
 * > Users are solely responsible for ensuring compliance with all applicable financial, tax, and regulatory requirements in
 * > their jurisdiction. This includes obtaining necessary licenses or approvals for trading or investment activities. The
 * > Developers disclaim liability for any legal consequences arising from non-compliance.
 * > 
 * > To the fullest extent permitted by law, the Developers shall not be liable for any direct, indirect, incidental,
 * > consequential, or punitive damages arising from the use or inability to use this software, including but not limited to
 * > loss of profits, data, or business opportunities.
 * 
 * </details>
 */

package dc_rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GuildResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildResponse{}

// GuildResponse struct for GuildResponse
type GuildResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Icon *string `json:"icon,omitempty"`
	Description *string `json:"description,omitempty"`
	HomeHeader *string `json:"home_header,omitempty"`
	Splash *string `json:"splash,omitempty"`
	DiscoverySplash *string `json:"discovery_splash,omitempty"`
	Features []GuildFeatures `json:"features"`
	Banner *string `json:"banner,omitempty"`
	OwnerId string `json:"owner_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ApplicationId *GetEntitlementsSkuIdsParameterOneOfInner `json:"application_id,omitempty"`
	Region string `json:"region"`
	AfkChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"afk_channel_id,omitempty"`
	AfkTimeout AfkTimeouts `json:"afk_timeout"`
	SystemChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"system_channel_id,omitempty"`
	SystemChannelFlags int32 `json:"system_channel_flags"`
	WidgetEnabled bool `json:"widget_enabled"`
	WidgetChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"widget_channel_id,omitempty"`
	VerificationLevel VerificationLevels `json:"verification_level"`
	Roles []GuildRoleResponse `json:"roles"`
	DefaultMessageNotifications UserNotificationSettings `json:"default_message_notifications"`
	MfaLevel GuildMFALevel `json:"mfa_level"`
	ExplicitContentFilter GuildExplicitContentFilterTypes `json:"explicit_content_filter"`
	MaxPresences *int32 `json:"max_presences,omitempty"`
	MaxMembers int32 `json:"max_members"`
	MaxStageVideoChannelUsers int32 `json:"max_stage_video_channel_users"`
	MaxVideoChannelUsers int32 `json:"max_video_channel_users"`
	VanityUrlCode *string `json:"vanity_url_code,omitempty"`
	PremiumTier PremiumGuildTiers `json:"premium_tier"`
	PremiumSubscriptionCount int32 `json:"premium_subscription_count"`
	PreferredLocale AvailableLocalesEnum `json:"preferred_locale"`
	RulesChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"rules_channel_id,omitempty"`
	SafetyAlertsChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"safety_alerts_channel_id,omitempty"`
	PublicUpdatesChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"public_updates_channel_id,omitempty"`
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`
	Nsfw bool `json:"nsfw"`
	NsfwLevel GuildNSFWContentLevel `json:"nsfw_level"`
	Emojis []EmojiResponse `json:"emojis"`
	Stickers []GuildStickerResponse `json:"stickers"`
}

type _GuildResponse GuildResponse

// NewGuildResponse instantiates a new GuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildResponse(id string, name string, features []GuildFeatures, ownerId string, region string, afkTimeout AfkTimeouts, systemChannelFlags int32, widgetEnabled bool, verificationLevel VerificationLevels, roles []GuildRoleResponse, defaultMessageNotifications UserNotificationSettings, mfaLevel GuildMFALevel, explicitContentFilter GuildExplicitContentFilterTypes, maxMembers int32, maxStageVideoChannelUsers int32, maxVideoChannelUsers int32, premiumTier PremiumGuildTiers, premiumSubscriptionCount int32, preferredLocale AvailableLocalesEnum, premiumProgressBarEnabled bool, nsfw bool, nsfwLevel GuildNSFWContentLevel, emojis []EmojiResponse, stickers []GuildStickerResponse) *GuildResponse {
	this := GuildResponse{}
	this.Id = id
	this.Name = name
	this.Features = features
	this.OwnerId = ownerId
	this.Region = region
	this.AfkTimeout = afkTimeout
	this.SystemChannelFlags = systemChannelFlags
	this.WidgetEnabled = widgetEnabled
	this.VerificationLevel = verificationLevel
	this.Roles = roles
	this.DefaultMessageNotifications = defaultMessageNotifications
	this.MfaLevel = mfaLevel
	this.ExplicitContentFilter = explicitContentFilter
	this.MaxMembers = maxMembers
	this.MaxStageVideoChannelUsers = maxStageVideoChannelUsers
	this.MaxVideoChannelUsers = maxVideoChannelUsers
	this.PremiumTier = premiumTier
	this.PremiumSubscriptionCount = premiumSubscriptionCount
	this.PreferredLocale = preferredLocale
	this.PremiumProgressBarEnabled = premiumProgressBarEnabled
	this.Nsfw = nsfw
	this.NsfwLevel = nsfwLevel
	this.Emojis = emojis
	this.Stickers = stickers
	return &this
}

// NewGuildResponseWithDefaults instantiates a new GuildResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildResponseWithDefaults() *GuildResponse {
	this := GuildResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GuildResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GuildResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GuildResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildResponse) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *GuildResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *GuildResponse) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *GuildResponse) SetIcon(v string) {
	o.Icon = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GuildResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GuildResponse) SetDescription(v string) {
	o.Description = &v
}


// GetHomeHeader returns the HomeHeader field value if set, zero value otherwise.
func (o *GuildResponse) GetHomeHeader() string {
	if o == nil || IsNil(o.HomeHeader) {
		var ret string
		return ret
	}
	return *o.HomeHeader
}

// GetHomeHeaderOk returns a tuple with the HomeHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetHomeHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.HomeHeader) {
		return nil, false
	}
	return o.HomeHeader, true
}

// HasHomeHeader returns a boolean if a field has been set.
func (o *GuildResponse) HasHomeHeader() bool {
	if o != nil && !IsNil(o.HomeHeader) {
		return true
	}

	return false
}

// SetHomeHeader gets a reference to the given string and assigns it to the HomeHeader field.
func (o *GuildResponse) SetHomeHeader(v string) {
	o.HomeHeader = &v
}


// GetSplash returns the Splash field value if set, zero value otherwise.
func (o *GuildResponse) GetSplash() string {
	if o == nil || IsNil(o.Splash) {
		var ret string
		return ret
	}
	return *o.Splash
}

// GetSplashOk returns a tuple with the Splash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetSplashOk() (*string, bool) {
	if o == nil || IsNil(o.Splash) {
		return nil, false
	}
	return o.Splash, true
}

// HasSplash returns a boolean if a field has been set.
func (o *GuildResponse) HasSplash() bool {
	if o != nil && !IsNil(o.Splash) {
		return true
	}

	return false
}

// SetSplash gets a reference to the given string and assigns it to the Splash field.
func (o *GuildResponse) SetSplash(v string) {
	o.Splash = &v
}


// GetDiscoverySplash returns the DiscoverySplash field value if set, zero value otherwise.
func (o *GuildResponse) GetDiscoverySplash() string {
	if o == nil || IsNil(o.DiscoverySplash) {
		var ret string
		return ret
	}
	return *o.DiscoverySplash
}

// GetDiscoverySplashOk returns a tuple with the DiscoverySplash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetDiscoverySplashOk() (*string, bool) {
	if o == nil || IsNil(o.DiscoverySplash) {
		return nil, false
	}
	return o.DiscoverySplash, true
}

// HasDiscoverySplash returns a boolean if a field has been set.
func (o *GuildResponse) HasDiscoverySplash() bool {
	if o != nil && !IsNil(o.DiscoverySplash) {
		return true
	}

	return false
}

// SetDiscoverySplash gets a reference to the given string and assigns it to the DiscoverySplash field.
func (o *GuildResponse) SetDiscoverySplash(v string) {
	o.DiscoverySplash = &v
}


// GetFeatures returns the Features field value
func (o *GuildResponse) GetFeatures() []GuildFeatures {
	if o == nil {
		var ret []GuildFeatures
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetFeaturesOk() ([]GuildFeatures, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *GuildResponse) SetFeatures(v []GuildFeatures) {
	o.Features = v
}

// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *GuildResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner) {
		var ret string
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *GuildResponse) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given string and assigns it to the Banner field.
func (o *GuildResponse) SetBanner(v string) {
	o.Banner = &v
}


// GetOwnerId returns the OwnerId field value
func (o *GuildResponse) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *GuildResponse) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *GuildResponse) GetApplicationId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ApplicationId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetApplicationIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *GuildResponse) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ApplicationId field.
func (o *GuildResponse) SetApplicationId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ApplicationId = &v
}


// GetRegion returns the Region field value
func (o *GuildResponse) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *GuildResponse) SetRegion(v string) {
	o.Region = v
}

// GetAfkChannelId returns the AfkChannelId field value if set, zero value otherwise.
func (o *GuildResponse) GetAfkChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.AfkChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.AfkChannelId
}

// GetAfkChannelIdOk returns a tuple with the AfkChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetAfkChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.AfkChannelId) {
		return nil, false
	}
	return o.AfkChannelId, true
}

// HasAfkChannelId returns a boolean if a field has been set.
func (o *GuildResponse) HasAfkChannelId() bool {
	if o != nil && !IsNil(o.AfkChannelId) {
		return true
	}

	return false
}

// SetAfkChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the AfkChannelId field.
func (o *GuildResponse) SetAfkChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.AfkChannelId = &v
}


// GetAfkTimeout returns the AfkTimeout field value
func (o *GuildResponse) GetAfkTimeout() AfkTimeouts {
	if o == nil {
		var ret AfkTimeouts
		return ret
	}

	return o.AfkTimeout
}

// GetAfkTimeoutOk returns a tuple with the AfkTimeout field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetAfkTimeoutOk() (*AfkTimeouts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfkTimeout, true
}

// SetAfkTimeout sets field value
func (o *GuildResponse) SetAfkTimeout(v AfkTimeouts) {
	o.AfkTimeout = v
}

// GetSystemChannelId returns the SystemChannelId field value if set, zero value otherwise.
func (o *GuildResponse) GetSystemChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.SystemChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.SystemChannelId
}

// GetSystemChannelIdOk returns a tuple with the SystemChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetSystemChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.SystemChannelId) {
		return nil, false
	}
	return o.SystemChannelId, true
}

// HasSystemChannelId returns a boolean if a field has been set.
func (o *GuildResponse) HasSystemChannelId() bool {
	if o != nil && !IsNil(o.SystemChannelId) {
		return true
	}

	return false
}

// SetSystemChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the SystemChannelId field.
func (o *GuildResponse) SetSystemChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.SystemChannelId = &v
}


// GetSystemChannelFlags returns the SystemChannelFlags field value
func (o *GuildResponse) GetSystemChannelFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SystemChannelFlags
}

// GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetSystemChannelFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemChannelFlags, true
}

// SetSystemChannelFlags sets field value
func (o *GuildResponse) SetSystemChannelFlags(v int32) {
	o.SystemChannelFlags = v
}

// GetWidgetEnabled returns the WidgetEnabled field value
func (o *GuildResponse) GetWidgetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WidgetEnabled
}

// GetWidgetEnabledOk returns a tuple with the WidgetEnabled field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetWidgetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WidgetEnabled, true
}

// SetWidgetEnabled sets field value
func (o *GuildResponse) SetWidgetEnabled(v bool) {
	o.WidgetEnabled = v
}

// GetWidgetChannelId returns the WidgetChannelId field value if set, zero value otherwise.
func (o *GuildResponse) GetWidgetChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.WidgetChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.WidgetChannelId
}

// GetWidgetChannelIdOk returns a tuple with the WidgetChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetWidgetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.WidgetChannelId) {
		return nil, false
	}
	return o.WidgetChannelId, true
}

// HasWidgetChannelId returns a boolean if a field has been set.
func (o *GuildResponse) HasWidgetChannelId() bool {
	if o != nil && !IsNil(o.WidgetChannelId) {
		return true
	}

	return false
}

// SetWidgetChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the WidgetChannelId field.
func (o *GuildResponse) SetWidgetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.WidgetChannelId = &v
}


// GetVerificationLevel returns the VerificationLevel field value
func (o *GuildResponse) GetVerificationLevel() VerificationLevels {
	if o == nil {
		var ret VerificationLevels
		return ret
	}

	return o.VerificationLevel
}

// GetVerificationLevelOk returns a tuple with the VerificationLevel field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetVerificationLevelOk() (*VerificationLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationLevel, true
}

// SetVerificationLevel sets field value
func (o *GuildResponse) SetVerificationLevel(v VerificationLevels) {
	o.VerificationLevel = v
}

// GetRoles returns the Roles field value
func (o *GuildResponse) GetRoles() []GuildRoleResponse {
	if o == nil {
		var ret []GuildRoleResponse
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetRolesOk() ([]GuildRoleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *GuildResponse) SetRoles(v []GuildRoleResponse) {
	o.Roles = v
}

// GetDefaultMessageNotifications returns the DefaultMessageNotifications field value
func (o *GuildResponse) GetDefaultMessageNotifications() UserNotificationSettings {
	if o == nil {
		var ret UserNotificationSettings
		return ret
	}

	return o.DefaultMessageNotifications
}

// GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetDefaultMessageNotificationsOk() (*UserNotificationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMessageNotifications, true
}

// SetDefaultMessageNotifications sets field value
func (o *GuildResponse) SetDefaultMessageNotifications(v UserNotificationSettings) {
	o.DefaultMessageNotifications = v
}

// GetMfaLevel returns the MfaLevel field value
func (o *GuildResponse) GetMfaLevel() GuildMFALevel {
	if o == nil {
		var ret GuildMFALevel
		return ret
	}

	return o.MfaLevel
}

// GetMfaLevelOk returns a tuple with the MfaLevel field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetMfaLevelOk() (*GuildMFALevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MfaLevel, true
}

// SetMfaLevel sets field value
func (o *GuildResponse) SetMfaLevel(v GuildMFALevel) {
	o.MfaLevel = v
}

// GetExplicitContentFilter returns the ExplicitContentFilter field value
func (o *GuildResponse) GetExplicitContentFilter() GuildExplicitContentFilterTypes {
	if o == nil {
		var ret GuildExplicitContentFilterTypes
		return ret
	}

	return o.ExplicitContentFilter
}

// GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetExplicitContentFilterOk() (*GuildExplicitContentFilterTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExplicitContentFilter, true
}

// SetExplicitContentFilter sets field value
func (o *GuildResponse) SetExplicitContentFilter(v GuildExplicitContentFilterTypes) {
	o.ExplicitContentFilter = v
}

// GetMaxPresences returns the MaxPresences field value if set, zero value otherwise.
func (o *GuildResponse) GetMaxPresences() int32 {
	if o == nil || IsNil(o.MaxPresences) {
		var ret int32
		return ret
	}
	return *o.MaxPresences
}

// GetMaxPresencesOk returns a tuple with the MaxPresences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetMaxPresencesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPresences) {
		return nil, false
	}
	return o.MaxPresences, true
}

// HasMaxPresences returns a boolean if a field has been set.
func (o *GuildResponse) HasMaxPresences() bool {
	if o != nil && !IsNil(o.MaxPresences) {
		return true
	}

	return false
}

// SetMaxPresences gets a reference to the given int32 and assigns it to the MaxPresences field.
func (o *GuildResponse) SetMaxPresences(v int32) {
	o.MaxPresences = &v
}


// GetMaxMembers returns the MaxMembers field value
func (o *GuildResponse) GetMaxMembers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxMembers
}

// GetMaxMembersOk returns a tuple with the MaxMembers field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetMaxMembersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMembers, true
}

// SetMaxMembers sets field value
func (o *GuildResponse) SetMaxMembers(v int32) {
	o.MaxMembers = v
}

// GetMaxStageVideoChannelUsers returns the MaxStageVideoChannelUsers field value
func (o *GuildResponse) GetMaxStageVideoChannelUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxStageVideoChannelUsers
}

// GetMaxStageVideoChannelUsersOk returns a tuple with the MaxStageVideoChannelUsers field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetMaxStageVideoChannelUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStageVideoChannelUsers, true
}

// SetMaxStageVideoChannelUsers sets field value
func (o *GuildResponse) SetMaxStageVideoChannelUsers(v int32) {
	o.MaxStageVideoChannelUsers = v
}

// GetMaxVideoChannelUsers returns the MaxVideoChannelUsers field value
func (o *GuildResponse) GetMaxVideoChannelUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxVideoChannelUsers
}

// GetMaxVideoChannelUsersOk returns a tuple with the MaxVideoChannelUsers field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetMaxVideoChannelUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxVideoChannelUsers, true
}

// SetMaxVideoChannelUsers sets field value
func (o *GuildResponse) SetMaxVideoChannelUsers(v int32) {
	o.MaxVideoChannelUsers = v
}

// GetVanityUrlCode returns the VanityUrlCode field value if set, zero value otherwise.
func (o *GuildResponse) GetVanityUrlCode() string {
	if o == nil || IsNil(o.VanityUrlCode) {
		var ret string
		return ret
	}
	return *o.VanityUrlCode
}

// GetVanityUrlCodeOk returns a tuple with the VanityUrlCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetVanityUrlCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VanityUrlCode) {
		return nil, false
	}
	return o.VanityUrlCode, true
}

// HasVanityUrlCode returns a boolean if a field has been set.
func (o *GuildResponse) HasVanityUrlCode() bool {
	if o != nil && !IsNil(o.VanityUrlCode) {
		return true
	}

	return false
}

// SetVanityUrlCode gets a reference to the given string and assigns it to the VanityUrlCode field.
func (o *GuildResponse) SetVanityUrlCode(v string) {
	o.VanityUrlCode = &v
}


// GetPremiumTier returns the PremiumTier field value
func (o *GuildResponse) GetPremiumTier() PremiumGuildTiers {
	if o == nil {
		var ret PremiumGuildTiers
		return ret
	}

	return o.PremiumTier
}

// GetPremiumTierOk returns a tuple with the PremiumTier field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetPremiumTierOk() (*PremiumGuildTiers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PremiumTier, true
}

// SetPremiumTier sets field value
func (o *GuildResponse) SetPremiumTier(v PremiumGuildTiers) {
	o.PremiumTier = v
}

// GetPremiumSubscriptionCount returns the PremiumSubscriptionCount field value
func (o *GuildResponse) GetPremiumSubscriptionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PremiumSubscriptionCount
}

// GetPremiumSubscriptionCountOk returns a tuple with the PremiumSubscriptionCount field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetPremiumSubscriptionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PremiumSubscriptionCount, true
}

// SetPremiumSubscriptionCount sets field value
func (o *GuildResponse) SetPremiumSubscriptionCount(v int32) {
	o.PremiumSubscriptionCount = v
}

// GetPreferredLocale returns the PreferredLocale field value
func (o *GuildResponse) GetPreferredLocale() AvailableLocalesEnum {
	if o == nil {
		var ret AvailableLocalesEnum
		return ret
	}

	return o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetPreferredLocaleOk() (*AvailableLocalesEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreferredLocale, true
}

// SetPreferredLocale sets field value
func (o *GuildResponse) SetPreferredLocale(v AvailableLocalesEnum) {
	o.PreferredLocale = v
}

// GetRulesChannelId returns the RulesChannelId field value if set, zero value otherwise.
func (o *GuildResponse) GetRulesChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.RulesChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.RulesChannelId
}

// GetRulesChannelIdOk returns a tuple with the RulesChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetRulesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.RulesChannelId) {
		return nil, false
	}
	return o.RulesChannelId, true
}

// HasRulesChannelId returns a boolean if a field has been set.
func (o *GuildResponse) HasRulesChannelId() bool {
	if o != nil && !IsNil(o.RulesChannelId) {
		return true
	}

	return false
}

// SetRulesChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the RulesChannelId field.
func (o *GuildResponse) SetRulesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.RulesChannelId = &v
}


// GetSafetyAlertsChannelId returns the SafetyAlertsChannelId field value if set, zero value otherwise.
func (o *GuildResponse) GetSafetyAlertsChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.SafetyAlertsChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.SafetyAlertsChannelId
}

// GetSafetyAlertsChannelIdOk returns a tuple with the SafetyAlertsChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetSafetyAlertsChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.SafetyAlertsChannelId) {
		return nil, false
	}
	return o.SafetyAlertsChannelId, true
}

// HasSafetyAlertsChannelId returns a boolean if a field has been set.
func (o *GuildResponse) HasSafetyAlertsChannelId() bool {
	if o != nil && !IsNil(o.SafetyAlertsChannelId) {
		return true
	}

	return false
}

// SetSafetyAlertsChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the SafetyAlertsChannelId field.
func (o *GuildResponse) SetSafetyAlertsChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.SafetyAlertsChannelId = &v
}


// GetPublicUpdatesChannelId returns the PublicUpdatesChannelId field value if set, zero value otherwise.
func (o *GuildResponse) GetPublicUpdatesChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.PublicUpdatesChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.PublicUpdatesChannelId
}

// GetPublicUpdatesChannelIdOk returns a tuple with the PublicUpdatesChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetPublicUpdatesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.PublicUpdatesChannelId) {
		return nil, false
	}
	return o.PublicUpdatesChannelId, true
}

// HasPublicUpdatesChannelId returns a boolean if a field has been set.
func (o *GuildResponse) HasPublicUpdatesChannelId() bool {
	if o != nil && !IsNil(o.PublicUpdatesChannelId) {
		return true
	}

	return false
}

// SetPublicUpdatesChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the PublicUpdatesChannelId field.
func (o *GuildResponse) SetPublicUpdatesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.PublicUpdatesChannelId = &v
}


// GetPremiumProgressBarEnabled returns the PremiumProgressBarEnabled field value
func (o *GuildResponse) GetPremiumProgressBarEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PremiumProgressBarEnabled
}

// GetPremiumProgressBarEnabledOk returns a tuple with the PremiumProgressBarEnabled field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetPremiumProgressBarEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PremiumProgressBarEnabled, true
}

// SetPremiumProgressBarEnabled sets field value
func (o *GuildResponse) SetPremiumProgressBarEnabled(v bool) {
	o.PremiumProgressBarEnabled = v
}

// GetNsfw returns the Nsfw field value
func (o *GuildResponse) GetNsfw() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetNsfwOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nsfw, true
}

// SetNsfw sets field value
func (o *GuildResponse) SetNsfw(v bool) {
	o.Nsfw = v
}

// GetNsfwLevel returns the NsfwLevel field value
func (o *GuildResponse) GetNsfwLevel() GuildNSFWContentLevel {
	if o == nil {
		var ret GuildNSFWContentLevel
		return ret
	}

	return o.NsfwLevel
}

// GetNsfwLevelOk returns a tuple with the NsfwLevel field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetNsfwLevelOk() (*GuildNSFWContentLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NsfwLevel, true
}

// SetNsfwLevel sets field value
func (o *GuildResponse) SetNsfwLevel(v GuildNSFWContentLevel) {
	o.NsfwLevel = v
}

// GetEmojis returns the Emojis field value
func (o *GuildResponse) GetEmojis() []EmojiResponse {
	if o == nil {
		var ret []EmojiResponse
		return ret
	}

	return o.Emojis
}

// GetEmojisOk returns a tuple with the Emojis field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetEmojisOk() ([]EmojiResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emojis, true
}

// SetEmojis sets field value
func (o *GuildResponse) SetEmojis(v []EmojiResponse) {
	o.Emojis = v
}

// GetStickers returns the Stickers field value
func (o *GuildResponse) GetStickers() []GuildStickerResponse {
	if o == nil {
		var ret []GuildStickerResponse
		return ret
	}

	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value
// and a boolean to check if the value has been set.
func (o *GuildResponse) GetStickersOk() ([]GuildStickerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stickers, true
}

// SetStickers sets field value
func (o *GuildResponse) SetStickers(v []GuildStickerResponse) {
	o.Stickers = v
}

func (o GuildResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.HomeHeader) {
		toSerialize["home_header"] = o.HomeHeader
	}
	if !IsNil(o.Splash) {
		toSerialize["splash"] = o.Splash
	}
	if !IsNil(o.DiscoverySplash) {
		toSerialize["discovery_splash"] = o.DiscoverySplash
	}
	toSerialize["features"] = o.Features
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	toSerialize["owner_id"] = o.OwnerId
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	toSerialize["region"] = o.Region
	if !IsNil(o.AfkChannelId) {
		toSerialize["afk_channel_id"] = o.AfkChannelId
	}
	toSerialize["afk_timeout"] = o.AfkTimeout
	if !IsNil(o.SystemChannelId) {
		toSerialize["system_channel_id"] = o.SystemChannelId
	}
	toSerialize["system_channel_flags"] = o.SystemChannelFlags
	toSerialize["widget_enabled"] = o.WidgetEnabled
	if !IsNil(o.WidgetChannelId) {
		toSerialize["widget_channel_id"] = o.WidgetChannelId
	}
	toSerialize["verification_level"] = o.VerificationLevel
	toSerialize["roles"] = o.Roles
	toSerialize["default_message_notifications"] = o.DefaultMessageNotifications
	toSerialize["mfa_level"] = o.MfaLevel
	toSerialize["explicit_content_filter"] = o.ExplicitContentFilter
	if !IsNil(o.MaxPresences) {
		toSerialize["max_presences"] = o.MaxPresences
	}
	toSerialize["max_members"] = o.MaxMembers
	toSerialize["max_stage_video_channel_users"] = o.MaxStageVideoChannelUsers
	toSerialize["max_video_channel_users"] = o.MaxVideoChannelUsers
	if !IsNil(o.VanityUrlCode) {
		toSerialize["vanity_url_code"] = o.VanityUrlCode
	}
	toSerialize["premium_tier"] = o.PremiumTier
	toSerialize["premium_subscription_count"] = o.PremiumSubscriptionCount
	toSerialize["preferred_locale"] = o.PreferredLocale
	if !IsNil(o.RulesChannelId) {
		toSerialize["rules_channel_id"] = o.RulesChannelId
	}
	if !IsNil(o.SafetyAlertsChannelId) {
		toSerialize["safety_alerts_channel_id"] = o.SafetyAlertsChannelId
	}
	if !IsNil(o.PublicUpdatesChannelId) {
		toSerialize["public_updates_channel_id"] = o.PublicUpdatesChannelId
	}
	toSerialize["premium_progress_bar_enabled"] = o.PremiumProgressBarEnabled
	toSerialize["nsfw"] = o.Nsfw
	toSerialize["nsfw_level"] = o.NsfwLevel
	toSerialize["emojis"] = o.Emojis
	toSerialize["stickers"] = o.Stickers
	return toSerialize, nil
}

func (o *GuildResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"features",
		"owner_id",
		"region",
		"afk_timeout",
		"system_channel_flags",
		"widget_enabled",
		"verification_level",
		"roles",
		"default_message_notifications",
		"mfa_level",
		"explicit_content_filter",
		"max_members",
		"max_stage_video_channel_users",
		"max_video_channel_users",
		"premium_tier",
		"premium_subscription_count",
		"preferred_locale",
		"premium_progress_bar_enabled",
		"nsfw",
		"nsfw_level",
		"emojis",
		"stickers",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGuildResponse := _GuildResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildResponse)

	if err != nil {
		return err
	}

	*o = GuildResponse(varGuildResponse)

	return err
}

type NullableGuildResponse struct {
	value *GuildResponse
	isSet bool
}

func (v NullableGuildResponse) Get() *GuildResponse {
	return v.value
}

func (v *NullableGuildResponse) Set(val *GuildResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildResponse(val *GuildResponse) *NullableGuildResponse {
	return &NullableGuildResponse{value: val, isSet: true}
}

func (v NullableGuildResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

