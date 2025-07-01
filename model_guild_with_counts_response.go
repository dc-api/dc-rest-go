/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
 *    * - **Generator Version**: 7.14.0
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

// checks if the GuildWithCountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildWithCountsResponse{}

// GuildWithCountsResponse struct for GuildWithCountsResponse
type GuildWithCountsResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Icon NullableString `json:"icon,omitempty"`
	Description NullableString `json:"description,omitempty"`
	HomeHeader NullableString `json:"home_header,omitempty"`
	Splash NullableString `json:"splash,omitempty"`
	DiscoverySplash NullableString `json:"discovery_splash,omitempty"`
	Features []string `json:"features"`
	Banner NullableString `json:"banner,omitempty"`
	OwnerId string `json:"owner_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ApplicationId *string `json:"application_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Region string `json:"region"`
	AfkChannelId *string `json:"afk_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	AfkTimeout NullableInt32 `json:"afk_timeout"`
	SystemChannelId *string `json:"system_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SystemChannelFlags int32 `json:"system_channel_flags"`
	WidgetEnabled bool `json:"widget_enabled"`
	WidgetChannelId *string `json:"widget_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	VerificationLevel int32 `json:"verification_level"`
	Roles []GuildRoleResponse `json:"roles"`
	DefaultMessageNotifications int32 `json:"default_message_notifications"`
	MfaLevel int32 `json:"mfa_level"`
	ExplicitContentFilter int32 `json:"explicit_content_filter"`
	MaxPresences NullableInt32 `json:"max_presences,omitempty"`
	MaxMembers NullableInt32 `json:"max_members,omitempty"`
	MaxStageVideoChannelUsers NullableInt32 `json:"max_stage_video_channel_users,omitempty"`
	MaxVideoChannelUsers NullableInt32 `json:"max_video_channel_users,omitempty"`
	VanityUrlCode NullableString `json:"vanity_url_code,omitempty"`
	PremiumTier int32 `json:"premium_tier"`
	PremiumSubscriptionCount int32 `json:"premium_subscription_count"`
	PreferredLocale string `json:"preferred_locale"`
	RulesChannelId *string `json:"rules_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SafetyAlertsChannelId *string `json:"safety_alerts_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	PublicUpdatesChannelId *string `json:"public_updates_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`
	Nsfw bool `json:"nsfw"`
	NsfwLevel NullableInt32 `json:"nsfw_level"`
	Emojis []EmojiResponse `json:"emojis"`
	Stickers []GuildStickerResponse `json:"stickers"`
	ApproximateMemberCount NullableInt32 `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount NullableInt32 `json:"approximate_presence_count,omitempty"`
}

type _GuildWithCountsResponse GuildWithCountsResponse

// NewGuildWithCountsResponse instantiates a new GuildWithCountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildWithCountsResponse(id string, name string, features []string, ownerId string, region string, afkTimeout NullableInt32, systemChannelFlags int32, widgetEnabled bool, verificationLevel int32, roles []GuildRoleResponse, defaultMessageNotifications int32, mfaLevel int32, explicitContentFilter int32, premiumTier int32, premiumSubscriptionCount int32, preferredLocale string, premiumProgressBarEnabled bool, nsfw bool, nsfwLevel NullableInt32, emojis []EmojiResponse, stickers []GuildStickerResponse) *GuildWithCountsResponse {
	this := GuildWithCountsResponse{}
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

// NewGuildWithCountsResponseWithDefaults instantiates a new GuildWithCountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildWithCountsResponseWithDefaults() *GuildWithCountsResponse {
	this := GuildWithCountsResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GuildWithCountsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GuildWithCountsResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GuildWithCountsResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildWithCountsResponse) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *GuildWithCountsResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *GuildWithCountsResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuildWithCountsResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuildWithCountsResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetHomeHeader returns the HomeHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetHomeHeader() string {
	if o == nil || IsNil(o.HomeHeader.Get()) {
		var ret string
		return ret
	}
	return *o.HomeHeader.Get()
}

// GetHomeHeaderOk returns a tuple with the HomeHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetHomeHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.HomeHeader.Get()) {
		return nil, false
	}
	return o.HomeHeader.Get(), o.HomeHeader.IsSet()
}

// HasHomeHeader returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasHomeHeader() bool {
	if o != nil && o.HomeHeader.IsSet() {
		return true
	}

	return false
}

// SetHomeHeader gets a reference to the given NullableString and assigns it to the HomeHeader field.
func (o *GuildWithCountsResponse) SetHomeHeader(v string) {
	o.HomeHeader.Set(&v)
}

// SetHomeHeaderNil sets the value for HomeHeader to be an explicit nil
func (o *GuildWithCountsResponse) SetHomeHeaderNil() {
	o.HomeHeader.Set(nil)
}

// UnsetHomeHeader ensures that no value is present for HomeHeader, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetHomeHeader() {
	o.HomeHeader.Unset()
}

// GetSplash returns the Splash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetSplash() string {
	if o == nil || IsNil(o.Splash.Get()) {
		var ret string
		return ret
	}
	return *o.Splash.Get()
}

// GetSplashOk returns a tuple with the Splash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetSplashOk() (*string, bool) {
	if o == nil || IsNil(o.Splash.Get()) {
		return nil, false
	}
	return o.Splash.Get(), o.Splash.IsSet()
}

// HasSplash returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasSplash() bool {
	if o != nil && o.Splash.IsSet() {
		return true
	}

	return false
}

// SetSplash gets a reference to the given NullableString and assigns it to the Splash field.
func (o *GuildWithCountsResponse) SetSplash(v string) {
	o.Splash.Set(&v)
}

// SetSplashNil sets the value for Splash to be an explicit nil
func (o *GuildWithCountsResponse) SetSplashNil() {
	o.Splash.Set(nil)
}

// UnsetSplash ensures that no value is present for Splash, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetSplash() {
	o.Splash.Unset()
}

// GetDiscoverySplash returns the DiscoverySplash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetDiscoverySplash() string {
	if o == nil || IsNil(o.DiscoverySplash.Get()) {
		var ret string
		return ret
	}
	return *o.DiscoverySplash.Get()
}

// GetDiscoverySplashOk returns a tuple with the DiscoverySplash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetDiscoverySplashOk() (*string, bool) {
	if o == nil || IsNil(o.DiscoverySplash.Get()) {
		return nil, false
	}
	return o.DiscoverySplash.Get(), o.DiscoverySplash.IsSet()
}

// HasDiscoverySplash returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasDiscoverySplash() bool {
	if o != nil && o.DiscoverySplash.IsSet() {
		return true
	}

	return false
}

// SetDiscoverySplash gets a reference to the given NullableString and assigns it to the DiscoverySplash field.
func (o *GuildWithCountsResponse) SetDiscoverySplash(v string) {
	o.DiscoverySplash.Set(&v)
}

// SetDiscoverySplashNil sets the value for DiscoverySplash to be an explicit nil
func (o *GuildWithCountsResponse) SetDiscoverySplashNil() {
	o.DiscoverySplash.Set(nil)
}

// UnsetDiscoverySplash ensures that no value is present for DiscoverySplash, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetDiscoverySplash() {
	o.DiscoverySplash.Unset()
}

// GetFeatures returns the Features field value
func (o *GuildWithCountsResponse) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *GuildWithCountsResponse) SetFeatures(v []string) {
	o.Features = v
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner.Get()) {
		var ret string
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner.Get()) {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasBanner() bool {
	if o != nil && o.Banner.IsSet() {
		return true
	}

	return false
}

// SetBanner gets a reference to the given NullableString and assigns it to the Banner field.
func (o *GuildWithCountsResponse) SetBanner(v string) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil
func (o *GuildWithCountsResponse) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetBanner() {
	o.Banner.Unset()
}

// GetOwnerId returns the OwnerId field value
func (o *GuildWithCountsResponse) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *GuildWithCountsResponse) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *GuildWithCountsResponse) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *GuildWithCountsResponse) SetApplicationId(v string) {
	o.ApplicationId = &v
}


// GetRegion returns the Region field value
func (o *GuildWithCountsResponse) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *GuildWithCountsResponse) SetRegion(v string) {
	o.Region = v
}

// GetAfkChannelId returns the AfkChannelId field value if set, zero value otherwise.
func (o *GuildWithCountsResponse) GetAfkChannelId() string {
	if o == nil || IsNil(o.AfkChannelId) {
		var ret string
		return ret
	}
	return *o.AfkChannelId
}

// GetAfkChannelIdOk returns a tuple with the AfkChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetAfkChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfkChannelId) {
		return nil, false
	}
	return o.AfkChannelId, true
}

// HasAfkChannelId returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasAfkChannelId() bool {
	if o != nil && !IsNil(o.AfkChannelId) {
		return true
	}

	return false
}

// SetAfkChannelId gets a reference to the given string and assigns it to the AfkChannelId field.
func (o *GuildWithCountsResponse) SetAfkChannelId(v string) {
	o.AfkChannelId = &v
}


// GetAfkTimeout returns the AfkTimeout field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *GuildWithCountsResponse) GetAfkTimeout() int32 {
	if o == nil || o.AfkTimeout.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AfkTimeout.Get()
}

// GetAfkTimeoutOk returns a tuple with the AfkTimeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetAfkTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfkTimeout.Get(), o.AfkTimeout.IsSet()
}

// SetAfkTimeout sets field value
func (o *GuildWithCountsResponse) SetAfkTimeout(v int32) {
	o.AfkTimeout.Set(&v)
}

// GetSystemChannelId returns the SystemChannelId field value if set, zero value otherwise.
func (o *GuildWithCountsResponse) GetSystemChannelId() string {
	if o == nil || IsNil(o.SystemChannelId) {
		var ret string
		return ret
	}
	return *o.SystemChannelId
}

// GetSystemChannelIdOk returns a tuple with the SystemChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetSystemChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemChannelId) {
		return nil, false
	}
	return o.SystemChannelId, true
}

// HasSystemChannelId returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasSystemChannelId() bool {
	if o != nil && !IsNil(o.SystemChannelId) {
		return true
	}

	return false
}

// SetSystemChannelId gets a reference to the given string and assigns it to the SystemChannelId field.
func (o *GuildWithCountsResponse) SetSystemChannelId(v string) {
	o.SystemChannelId = &v
}


// GetSystemChannelFlags returns the SystemChannelFlags field value
func (o *GuildWithCountsResponse) GetSystemChannelFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SystemChannelFlags
}

// GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetSystemChannelFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemChannelFlags, true
}

// SetSystemChannelFlags sets field value
func (o *GuildWithCountsResponse) SetSystemChannelFlags(v int32) {
	o.SystemChannelFlags = v
}

// GetWidgetEnabled returns the WidgetEnabled field value
func (o *GuildWithCountsResponse) GetWidgetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WidgetEnabled
}

// GetWidgetEnabledOk returns a tuple with the WidgetEnabled field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetWidgetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WidgetEnabled, true
}

// SetWidgetEnabled sets field value
func (o *GuildWithCountsResponse) SetWidgetEnabled(v bool) {
	o.WidgetEnabled = v
}

// GetWidgetChannelId returns the WidgetChannelId field value if set, zero value otherwise.
func (o *GuildWithCountsResponse) GetWidgetChannelId() string {
	if o == nil || IsNil(o.WidgetChannelId) {
		var ret string
		return ret
	}
	return *o.WidgetChannelId
}

// GetWidgetChannelIdOk returns a tuple with the WidgetChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetWidgetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.WidgetChannelId) {
		return nil, false
	}
	return o.WidgetChannelId, true
}

// HasWidgetChannelId returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasWidgetChannelId() bool {
	if o != nil && !IsNil(o.WidgetChannelId) {
		return true
	}

	return false
}

// SetWidgetChannelId gets a reference to the given string and assigns it to the WidgetChannelId field.
func (o *GuildWithCountsResponse) SetWidgetChannelId(v string) {
	o.WidgetChannelId = &v
}


// GetVerificationLevel returns the VerificationLevel field value
func (o *GuildWithCountsResponse) GetVerificationLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VerificationLevel
}

// GetVerificationLevelOk returns a tuple with the VerificationLevel field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetVerificationLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationLevel, true
}

// SetVerificationLevel sets field value
func (o *GuildWithCountsResponse) SetVerificationLevel(v int32) {
	o.VerificationLevel = v
}

// GetRoles returns the Roles field value
func (o *GuildWithCountsResponse) GetRoles() []GuildRoleResponse {
	if o == nil {
		var ret []GuildRoleResponse
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetRolesOk() ([]GuildRoleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *GuildWithCountsResponse) SetRoles(v []GuildRoleResponse) {
	o.Roles = v
}

// GetDefaultMessageNotifications returns the DefaultMessageNotifications field value
func (o *GuildWithCountsResponse) GetDefaultMessageNotifications() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefaultMessageNotifications
}

// GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetDefaultMessageNotificationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMessageNotifications, true
}

// SetDefaultMessageNotifications sets field value
func (o *GuildWithCountsResponse) SetDefaultMessageNotifications(v int32) {
	o.DefaultMessageNotifications = v
}

// GetMfaLevel returns the MfaLevel field value
func (o *GuildWithCountsResponse) GetMfaLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MfaLevel
}

// GetMfaLevelOk returns a tuple with the MfaLevel field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetMfaLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MfaLevel, true
}

// SetMfaLevel sets field value
func (o *GuildWithCountsResponse) SetMfaLevel(v int32) {
	o.MfaLevel = v
}

// GetExplicitContentFilter returns the ExplicitContentFilter field value
func (o *GuildWithCountsResponse) GetExplicitContentFilter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExplicitContentFilter
}

// GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetExplicitContentFilterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExplicitContentFilter, true
}

// SetExplicitContentFilter sets field value
func (o *GuildWithCountsResponse) SetExplicitContentFilter(v int32) {
	o.ExplicitContentFilter = v
}

// GetMaxPresences returns the MaxPresences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetMaxPresences() int32 {
	if o == nil || IsNil(o.MaxPresences.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPresences.Get()
}

// GetMaxPresencesOk returns a tuple with the MaxPresences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetMaxPresencesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPresences.Get()) {
		return nil, false
	}
	return o.MaxPresences.Get(), o.MaxPresences.IsSet()
}

// HasMaxPresences returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasMaxPresences() bool {
	if o != nil && o.MaxPresences.IsSet() {
		return true
	}

	return false
}

// SetMaxPresences gets a reference to the given NullableInt32 and assigns it to the MaxPresences field.
func (o *GuildWithCountsResponse) SetMaxPresences(v int32) {
	o.MaxPresences.Set(&v)
}

// SetMaxPresencesNil sets the value for MaxPresences to be an explicit nil
func (o *GuildWithCountsResponse) SetMaxPresencesNil() {
	o.MaxPresences.Set(nil)
}

// UnsetMaxPresences ensures that no value is present for MaxPresences, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetMaxPresences() {
	o.MaxPresences.Unset()
}

// GetMaxMembers returns the MaxMembers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetMaxMembers() int32 {
	if o == nil || IsNil(o.MaxMembers.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxMembers.Get()
}

// GetMaxMembersOk returns a tuple with the MaxMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetMaxMembersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxMembers.Get()) {
		return nil, false
	}
	return o.MaxMembers.Get(), o.MaxMembers.IsSet()
}

// HasMaxMembers returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasMaxMembers() bool {
	if o != nil && o.MaxMembers.IsSet() {
		return true
	}

	return false
}

// SetMaxMembers gets a reference to the given NullableInt32 and assigns it to the MaxMembers field.
func (o *GuildWithCountsResponse) SetMaxMembers(v int32) {
	o.MaxMembers.Set(&v)
}

// SetMaxMembersNil sets the value for MaxMembers to be an explicit nil
func (o *GuildWithCountsResponse) SetMaxMembersNil() {
	o.MaxMembers.Set(nil)
}

// UnsetMaxMembers ensures that no value is present for MaxMembers, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetMaxMembers() {
	o.MaxMembers.Unset()
}

// GetMaxStageVideoChannelUsers returns the MaxStageVideoChannelUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetMaxStageVideoChannelUsers() int32 {
	if o == nil || IsNil(o.MaxStageVideoChannelUsers.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxStageVideoChannelUsers.Get()
}

// GetMaxStageVideoChannelUsersOk returns a tuple with the MaxStageVideoChannelUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetMaxStageVideoChannelUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxStageVideoChannelUsers.Get()) {
		return nil, false
	}
	return o.MaxStageVideoChannelUsers.Get(), o.MaxStageVideoChannelUsers.IsSet()
}

// HasMaxStageVideoChannelUsers returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasMaxStageVideoChannelUsers() bool {
	if o != nil && o.MaxStageVideoChannelUsers.IsSet() {
		return true
	}

	return false
}

// SetMaxStageVideoChannelUsers gets a reference to the given NullableInt32 and assigns it to the MaxStageVideoChannelUsers field.
func (o *GuildWithCountsResponse) SetMaxStageVideoChannelUsers(v int32) {
	o.MaxStageVideoChannelUsers.Set(&v)
}

// SetMaxStageVideoChannelUsersNil sets the value for MaxStageVideoChannelUsers to be an explicit nil
func (o *GuildWithCountsResponse) SetMaxStageVideoChannelUsersNil() {
	o.MaxStageVideoChannelUsers.Set(nil)
}

// UnsetMaxStageVideoChannelUsers ensures that no value is present for MaxStageVideoChannelUsers, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetMaxStageVideoChannelUsers() {
	o.MaxStageVideoChannelUsers.Unset()
}

// GetMaxVideoChannelUsers returns the MaxVideoChannelUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetMaxVideoChannelUsers() int32 {
	if o == nil || IsNil(o.MaxVideoChannelUsers.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxVideoChannelUsers.Get()
}

// GetMaxVideoChannelUsersOk returns a tuple with the MaxVideoChannelUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetMaxVideoChannelUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxVideoChannelUsers.Get()) {
		return nil, false
	}
	return o.MaxVideoChannelUsers.Get(), o.MaxVideoChannelUsers.IsSet()
}

// HasMaxVideoChannelUsers returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasMaxVideoChannelUsers() bool {
	if o != nil && o.MaxVideoChannelUsers.IsSet() {
		return true
	}

	return false
}

// SetMaxVideoChannelUsers gets a reference to the given NullableInt32 and assigns it to the MaxVideoChannelUsers field.
func (o *GuildWithCountsResponse) SetMaxVideoChannelUsers(v int32) {
	o.MaxVideoChannelUsers.Set(&v)
}

// SetMaxVideoChannelUsersNil sets the value for MaxVideoChannelUsers to be an explicit nil
func (o *GuildWithCountsResponse) SetMaxVideoChannelUsersNil() {
	o.MaxVideoChannelUsers.Set(nil)
}

// UnsetMaxVideoChannelUsers ensures that no value is present for MaxVideoChannelUsers, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetMaxVideoChannelUsers() {
	o.MaxVideoChannelUsers.Unset()
}

// GetVanityUrlCode returns the VanityUrlCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetVanityUrlCode() string {
	if o == nil || IsNil(o.VanityUrlCode.Get()) {
		var ret string
		return ret
	}
	return *o.VanityUrlCode.Get()
}

// GetVanityUrlCodeOk returns a tuple with the VanityUrlCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetVanityUrlCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VanityUrlCode.Get()) {
		return nil, false
	}
	return o.VanityUrlCode.Get(), o.VanityUrlCode.IsSet()
}

// HasVanityUrlCode returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasVanityUrlCode() bool {
	if o != nil && o.VanityUrlCode.IsSet() {
		return true
	}

	return false
}

// SetVanityUrlCode gets a reference to the given NullableString and assigns it to the VanityUrlCode field.
func (o *GuildWithCountsResponse) SetVanityUrlCode(v string) {
	o.VanityUrlCode.Set(&v)
}

// SetVanityUrlCodeNil sets the value for VanityUrlCode to be an explicit nil
func (o *GuildWithCountsResponse) SetVanityUrlCodeNil() {
	o.VanityUrlCode.Set(nil)
}

// UnsetVanityUrlCode ensures that no value is present for VanityUrlCode, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetVanityUrlCode() {
	o.VanityUrlCode.Unset()
}

// GetPremiumTier returns the PremiumTier field value
func (o *GuildWithCountsResponse) GetPremiumTier() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PremiumTier
}

// GetPremiumTierOk returns a tuple with the PremiumTier field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetPremiumTierOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PremiumTier, true
}

// SetPremiumTier sets field value
func (o *GuildWithCountsResponse) SetPremiumTier(v int32) {
	o.PremiumTier = v
}

// GetPremiumSubscriptionCount returns the PremiumSubscriptionCount field value
func (o *GuildWithCountsResponse) GetPremiumSubscriptionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PremiumSubscriptionCount
}

// GetPremiumSubscriptionCountOk returns a tuple with the PremiumSubscriptionCount field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetPremiumSubscriptionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PremiumSubscriptionCount, true
}

// SetPremiumSubscriptionCount sets field value
func (o *GuildWithCountsResponse) SetPremiumSubscriptionCount(v int32) {
	o.PremiumSubscriptionCount = v
}

// GetPreferredLocale returns the PreferredLocale field value
func (o *GuildWithCountsResponse) GetPreferredLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetPreferredLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreferredLocale, true
}

// SetPreferredLocale sets field value
func (o *GuildWithCountsResponse) SetPreferredLocale(v string) {
	o.PreferredLocale = v
}

// GetRulesChannelId returns the RulesChannelId field value if set, zero value otherwise.
func (o *GuildWithCountsResponse) GetRulesChannelId() string {
	if o == nil || IsNil(o.RulesChannelId) {
		var ret string
		return ret
	}
	return *o.RulesChannelId
}

// GetRulesChannelIdOk returns a tuple with the RulesChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetRulesChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.RulesChannelId) {
		return nil, false
	}
	return o.RulesChannelId, true
}

// HasRulesChannelId returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasRulesChannelId() bool {
	if o != nil && !IsNil(o.RulesChannelId) {
		return true
	}

	return false
}

// SetRulesChannelId gets a reference to the given string and assigns it to the RulesChannelId field.
func (o *GuildWithCountsResponse) SetRulesChannelId(v string) {
	o.RulesChannelId = &v
}


// GetSafetyAlertsChannelId returns the SafetyAlertsChannelId field value if set, zero value otherwise.
func (o *GuildWithCountsResponse) GetSafetyAlertsChannelId() string {
	if o == nil || IsNil(o.SafetyAlertsChannelId) {
		var ret string
		return ret
	}
	return *o.SafetyAlertsChannelId
}

// GetSafetyAlertsChannelIdOk returns a tuple with the SafetyAlertsChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetSafetyAlertsChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.SafetyAlertsChannelId) {
		return nil, false
	}
	return o.SafetyAlertsChannelId, true
}

// HasSafetyAlertsChannelId returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasSafetyAlertsChannelId() bool {
	if o != nil && !IsNil(o.SafetyAlertsChannelId) {
		return true
	}

	return false
}

// SetSafetyAlertsChannelId gets a reference to the given string and assigns it to the SafetyAlertsChannelId field.
func (o *GuildWithCountsResponse) SetSafetyAlertsChannelId(v string) {
	o.SafetyAlertsChannelId = &v
}


// GetPublicUpdatesChannelId returns the PublicUpdatesChannelId field value if set, zero value otherwise.
func (o *GuildWithCountsResponse) GetPublicUpdatesChannelId() string {
	if o == nil || IsNil(o.PublicUpdatesChannelId) {
		var ret string
		return ret
	}
	return *o.PublicUpdatesChannelId
}

// GetPublicUpdatesChannelIdOk returns a tuple with the PublicUpdatesChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetPublicUpdatesChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicUpdatesChannelId) {
		return nil, false
	}
	return o.PublicUpdatesChannelId, true
}

// HasPublicUpdatesChannelId returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasPublicUpdatesChannelId() bool {
	if o != nil && !IsNil(o.PublicUpdatesChannelId) {
		return true
	}

	return false
}

// SetPublicUpdatesChannelId gets a reference to the given string and assigns it to the PublicUpdatesChannelId field.
func (o *GuildWithCountsResponse) SetPublicUpdatesChannelId(v string) {
	o.PublicUpdatesChannelId = &v
}


// GetPremiumProgressBarEnabled returns the PremiumProgressBarEnabled field value
func (o *GuildWithCountsResponse) GetPremiumProgressBarEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PremiumProgressBarEnabled
}

// GetPremiumProgressBarEnabledOk returns a tuple with the PremiumProgressBarEnabled field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetPremiumProgressBarEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PremiumProgressBarEnabled, true
}

// SetPremiumProgressBarEnabled sets field value
func (o *GuildWithCountsResponse) SetPremiumProgressBarEnabled(v bool) {
	o.PremiumProgressBarEnabled = v
}

// GetNsfw returns the Nsfw field value
func (o *GuildWithCountsResponse) GetNsfw() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetNsfwOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nsfw, true
}

// SetNsfw sets field value
func (o *GuildWithCountsResponse) SetNsfw(v bool) {
	o.Nsfw = v
}

// GetNsfwLevel returns the NsfwLevel field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *GuildWithCountsResponse) GetNsfwLevel() int32 {
	if o == nil || o.NsfwLevel.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NsfwLevel.Get()
}

// GetNsfwLevelOk returns a tuple with the NsfwLevel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetNsfwLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NsfwLevel.Get(), o.NsfwLevel.IsSet()
}

// SetNsfwLevel sets field value
func (o *GuildWithCountsResponse) SetNsfwLevel(v int32) {
	o.NsfwLevel.Set(&v)
}

// GetEmojis returns the Emojis field value
func (o *GuildWithCountsResponse) GetEmojis() []EmojiResponse {
	if o == nil {
		var ret []EmojiResponse
		return ret
	}

	return o.Emojis
}

// GetEmojisOk returns a tuple with the Emojis field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetEmojisOk() ([]EmojiResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emojis, true
}

// SetEmojis sets field value
func (o *GuildWithCountsResponse) SetEmojis(v []EmojiResponse) {
	o.Emojis = v
}

// GetStickers returns the Stickers field value
func (o *GuildWithCountsResponse) GetStickers() []GuildStickerResponse {
	if o == nil {
		var ret []GuildStickerResponse
		return ret
	}

	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value
// and a boolean to check if the value has been set.
func (o *GuildWithCountsResponse) GetStickersOk() ([]GuildStickerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stickers, true
}

// SetStickers sets field value
func (o *GuildWithCountsResponse) SetStickers(v []GuildStickerResponse) {
	o.Stickers = v
}

// GetApproximateMemberCount returns the ApproximateMemberCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetApproximateMemberCount() int32 {
	if o == nil || IsNil(o.ApproximateMemberCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ApproximateMemberCount.Get()
}

// GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetApproximateMemberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproximateMemberCount.Get()) {
		return nil, false
	}
	return o.ApproximateMemberCount.Get(), o.ApproximateMemberCount.IsSet()
}

// HasApproximateMemberCount returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasApproximateMemberCount() bool {
	if o != nil && o.ApproximateMemberCount.IsSet() {
		return true
	}

	return false
}

// SetApproximateMemberCount gets a reference to the given NullableInt32 and assigns it to the ApproximateMemberCount field.
func (o *GuildWithCountsResponse) SetApproximateMemberCount(v int32) {
	o.ApproximateMemberCount.Set(&v)
}

// SetApproximateMemberCountNil sets the value for ApproximateMemberCount to be an explicit nil
func (o *GuildWithCountsResponse) SetApproximateMemberCountNil() {
	o.ApproximateMemberCount.Set(nil)
}

// UnsetApproximateMemberCount ensures that no value is present for ApproximateMemberCount, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetApproximateMemberCount() {
	o.ApproximateMemberCount.Unset()
}

// GetApproximatePresenceCount returns the ApproximatePresenceCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWithCountsResponse) GetApproximatePresenceCount() int32 {
	if o == nil || IsNil(o.ApproximatePresenceCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ApproximatePresenceCount.Get()
}

// GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWithCountsResponse) GetApproximatePresenceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproximatePresenceCount.Get()) {
		return nil, false
	}
	return o.ApproximatePresenceCount.Get(), o.ApproximatePresenceCount.IsSet()
}

// HasApproximatePresenceCount returns a boolean if a field has been set.
func (o *GuildWithCountsResponse) HasApproximatePresenceCount() bool {
	if o != nil && o.ApproximatePresenceCount.IsSet() {
		return true
	}

	return false
}

// SetApproximatePresenceCount gets a reference to the given NullableInt32 and assigns it to the ApproximatePresenceCount field.
func (o *GuildWithCountsResponse) SetApproximatePresenceCount(v int32) {
	o.ApproximatePresenceCount.Set(&v)
}

// SetApproximatePresenceCountNil sets the value for ApproximatePresenceCount to be an explicit nil
func (o *GuildWithCountsResponse) SetApproximatePresenceCountNil() {
	o.ApproximatePresenceCount.Set(nil)
}

// UnsetApproximatePresenceCount ensures that no value is present for ApproximatePresenceCount, not even an explicit nil
func (o *GuildWithCountsResponse) UnsetApproximatePresenceCount() {
	o.ApproximatePresenceCount.Unset()
}

func (o GuildWithCountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildWithCountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.HomeHeader.IsSet() {
		toSerialize["home_header"] = o.HomeHeader.Get()
	}
	if o.Splash.IsSet() {
		toSerialize["splash"] = o.Splash.Get()
	}
	if o.DiscoverySplash.IsSet() {
		toSerialize["discovery_splash"] = o.DiscoverySplash.Get()
	}
	toSerialize["features"] = o.Features
	if o.Banner.IsSet() {
		toSerialize["banner"] = o.Banner.Get()
	}
	toSerialize["owner_id"] = o.OwnerId
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	toSerialize["region"] = o.Region
	if !IsNil(o.AfkChannelId) {
		toSerialize["afk_channel_id"] = o.AfkChannelId
	}
	toSerialize["afk_timeout"] = o.AfkTimeout.Get()
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
	if o.MaxPresences.IsSet() {
		toSerialize["max_presences"] = o.MaxPresences.Get()
	}
	if o.MaxMembers.IsSet() {
		toSerialize["max_members"] = o.MaxMembers.Get()
	}
	if o.MaxStageVideoChannelUsers.IsSet() {
		toSerialize["max_stage_video_channel_users"] = o.MaxStageVideoChannelUsers.Get()
	}
	if o.MaxVideoChannelUsers.IsSet() {
		toSerialize["max_video_channel_users"] = o.MaxVideoChannelUsers.Get()
	}
	if o.VanityUrlCode.IsSet() {
		toSerialize["vanity_url_code"] = o.VanityUrlCode.Get()
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
	toSerialize["nsfw_level"] = o.NsfwLevel.Get()
	toSerialize["emojis"] = o.Emojis
	toSerialize["stickers"] = o.Stickers
	if o.ApproximateMemberCount.IsSet() {
		toSerialize["approximate_member_count"] = o.ApproximateMemberCount.Get()
	}
	if o.ApproximatePresenceCount.IsSet() {
		toSerialize["approximate_presence_count"] = o.ApproximatePresenceCount.Get()
	}
	return toSerialize, nil
}

func (o *GuildWithCountsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGuildWithCountsResponse := _GuildWithCountsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildWithCountsResponse)

	if err != nil {
		return err
	}

	*o = GuildWithCountsResponse(varGuildWithCountsResponse)

	return err
}

type NullableGuildWithCountsResponse struct {
	value *GuildWithCountsResponse
	isSet bool
}

func (v NullableGuildWithCountsResponse) Get() *GuildWithCountsResponse {
	return v.value
}

func (v *NullableGuildWithCountsResponse) Set(val *GuildWithCountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildWithCountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildWithCountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildWithCountsResponse(val *GuildWithCountsResponse) *NullableGuildWithCountsResponse {
	return &NullableGuildWithCountsResponse{value: val, isSet: true}
}

func (v NullableGuildWithCountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildWithCountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


