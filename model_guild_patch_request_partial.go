/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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
)

// checks if the GuildPatchRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildPatchRequestPartial{}

// GuildPatchRequestPartial struct for GuildPatchRequestPartial
type GuildPatchRequestPartial struct {
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Region NullableString `json:"region,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	VerificationLevel *int32 `json:"verification_level,omitempty"`
	DefaultMessageNotifications *int32 `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter *int32 `json:"explicit_content_filter,omitempty"`
	PreferredLocale *string `json:"preferred_locale,omitempty"`
	AfkTimeout NullableInt32 `json:"afk_timeout,omitempty"`
	AfkChannelId *string `json:"afk_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SystemChannelId *string `json:"system_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	OwnerId *string `json:"owner_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Splash NullableString `json:"splash,omitempty"`
	Banner NullableString `json:"banner,omitempty"`
	SystemChannelFlags NullableInt32 `json:"system_channel_flags,omitempty"`
	Features []*string `json:"features,omitempty"`
	DiscoverySplash NullableString `json:"discovery_splash,omitempty"`
	HomeHeader NullableString `json:"home_header,omitempty"`
	RulesChannelId *string `json:"rules_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SafetyAlertsChannelId *string `json:"safety_alerts_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	PublicUpdatesChannelId *string `json:"public_updates_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	PremiumProgressBarEnabled NullableBool `json:"premium_progress_bar_enabled,omitempty"`
}

// NewGuildPatchRequestPartial instantiates a new GuildPatchRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildPatchRequestPartial() *GuildPatchRequestPartial {
	this := GuildPatchRequestPartial{}
	return &this
}

// NewGuildPatchRequestPartialWithDefaults instantiates a new GuildPatchRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildPatchRequestPartialWithDefaults() *GuildPatchRequestPartial {
	this := GuildPatchRequestPartial{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GuildPatchRequestPartial) SetName(v string) {
	o.Name = &v
}


// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuildPatchRequestPartial) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuildPatchRequestPartial) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetDescription() {
	o.Description.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region.Get()) {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *GuildPatchRequestPartial) SetRegion(v string) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *GuildPatchRequestPartial) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetRegion() {
	o.Region.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *GuildPatchRequestPartial) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *GuildPatchRequestPartial) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetIcon() {
	o.Icon.Unset()
}

// GetVerificationLevel returns the VerificationLevel field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetVerificationLevel() int32 {
	if o == nil || IsNil(o.VerificationLevel) {
		var ret int32
		return ret
	}
	return *o.VerificationLevel
}

// GetVerificationLevelOk returns a tuple with the VerificationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetVerificationLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.VerificationLevel) {
		return nil, false
	}
	return o.VerificationLevel, true
}

// HasVerificationLevel returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasVerificationLevel() bool {
	if o != nil && !IsNil(o.VerificationLevel) {
		return true
	}

	return false
}

// SetVerificationLevel gets a reference to the given int32 and assigns it to the VerificationLevel field.
func (o *GuildPatchRequestPartial) SetVerificationLevel(v int32) {
	o.VerificationLevel = &v
}


// GetDefaultMessageNotifications returns the DefaultMessageNotifications field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetDefaultMessageNotifications() int32 {
	if o == nil || IsNil(o.DefaultMessageNotifications) {
		var ret int32
		return ret
	}
	return *o.DefaultMessageNotifications
}

// GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetDefaultMessageNotificationsOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultMessageNotifications) {
		return nil, false
	}
	return o.DefaultMessageNotifications, true
}

// HasDefaultMessageNotifications returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasDefaultMessageNotifications() bool {
	if o != nil && !IsNil(o.DefaultMessageNotifications) {
		return true
	}

	return false
}

// SetDefaultMessageNotifications gets a reference to the given int32 and assigns it to the DefaultMessageNotifications field.
func (o *GuildPatchRequestPartial) SetDefaultMessageNotifications(v int32) {
	o.DefaultMessageNotifications = &v
}


// GetExplicitContentFilter returns the ExplicitContentFilter field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetExplicitContentFilter() int32 {
	if o == nil || IsNil(o.ExplicitContentFilter) {
		var ret int32
		return ret
	}
	return *o.ExplicitContentFilter
}

// GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetExplicitContentFilterOk() (*int32, bool) {
	if o == nil || IsNil(o.ExplicitContentFilter) {
		return nil, false
	}
	return o.ExplicitContentFilter, true
}

// HasExplicitContentFilter returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasExplicitContentFilter() bool {
	if o != nil && !IsNil(o.ExplicitContentFilter) {
		return true
	}

	return false
}

// SetExplicitContentFilter gets a reference to the given int32 and assigns it to the ExplicitContentFilter field.
func (o *GuildPatchRequestPartial) SetExplicitContentFilter(v int32) {
	o.ExplicitContentFilter = &v
}


// GetPreferredLocale returns the PreferredLocale field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetPreferredLocale() string {
	if o == nil || IsNil(o.PreferredLocale) {
		var ret string
		return ret
	}
	return *o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetPreferredLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredLocale) {
		return nil, false
	}
	return o.PreferredLocale, true
}

// HasPreferredLocale returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasPreferredLocale() bool {
	if o != nil && !IsNil(o.PreferredLocale) {
		return true
	}

	return false
}

// SetPreferredLocale gets a reference to the given string and assigns it to the PreferredLocale field.
func (o *GuildPatchRequestPartial) SetPreferredLocale(v string) {
	o.PreferredLocale = &v
}


// GetAfkTimeout returns the AfkTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetAfkTimeout() int32 {
	if o == nil || IsNil(o.AfkTimeout.Get()) {
		var ret int32
		return ret
	}
	return *o.AfkTimeout.Get()
}

// GetAfkTimeoutOk returns a tuple with the AfkTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetAfkTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.AfkTimeout.Get()) {
		return nil, false
	}
	return o.AfkTimeout.Get(), o.AfkTimeout.IsSet()
}

// HasAfkTimeout returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasAfkTimeout() bool {
	if o != nil && o.AfkTimeout.IsSet() {
		return true
	}

	return false
}

// SetAfkTimeout gets a reference to the given NullableInt32 and assigns it to the AfkTimeout field.
func (o *GuildPatchRequestPartial) SetAfkTimeout(v int32) {
	o.AfkTimeout.Set(&v)
}

// SetAfkTimeoutNil sets the value for AfkTimeout to be an explicit nil
func (o *GuildPatchRequestPartial) SetAfkTimeoutNil() {
	o.AfkTimeout.Set(nil)
}

// UnsetAfkTimeout ensures that no value is present for AfkTimeout, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetAfkTimeout() {
	o.AfkTimeout.Unset()
}

// GetAfkChannelId returns the AfkChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetAfkChannelId() string {
	if o == nil || IsNil(o.AfkChannelId) {
		var ret string
		return ret
	}
	return *o.AfkChannelId
}

// GetAfkChannelIdOk returns a tuple with the AfkChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetAfkChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfkChannelId) {
		return nil, false
	}
	return o.AfkChannelId, true
}

// HasAfkChannelId returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasAfkChannelId() bool {
	if o != nil && !IsNil(o.AfkChannelId) {
		return true
	}

	return false
}

// SetAfkChannelId gets a reference to the given string and assigns it to the AfkChannelId field.
func (o *GuildPatchRequestPartial) SetAfkChannelId(v string) {
	o.AfkChannelId = &v
}


// GetSystemChannelId returns the SystemChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetSystemChannelId() string {
	if o == nil || IsNil(o.SystemChannelId) {
		var ret string
		return ret
	}
	return *o.SystemChannelId
}

// GetSystemChannelIdOk returns a tuple with the SystemChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetSystemChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemChannelId) {
		return nil, false
	}
	return o.SystemChannelId, true
}

// HasSystemChannelId returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasSystemChannelId() bool {
	if o != nil && !IsNil(o.SystemChannelId) {
		return true
	}

	return false
}

// SetSystemChannelId gets a reference to the given string and assigns it to the SystemChannelId field.
func (o *GuildPatchRequestPartial) SetSystemChannelId(v string) {
	o.SystemChannelId = &v
}


// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *GuildPatchRequestPartial) SetOwnerId(v string) {
	o.OwnerId = &v
}


// GetSplash returns the Splash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetSplash() string {
	if o == nil || IsNil(o.Splash.Get()) {
		var ret string
		return ret
	}
	return *o.Splash.Get()
}

// GetSplashOk returns a tuple with the Splash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetSplashOk() (*string, bool) {
	if o == nil || IsNil(o.Splash.Get()) {
		return nil, false
	}
	return o.Splash.Get(), o.Splash.IsSet()
}

// HasSplash returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasSplash() bool {
	if o != nil && o.Splash.IsSet() {
		return true
	}

	return false
}

// SetSplash gets a reference to the given NullableString and assigns it to the Splash field.
func (o *GuildPatchRequestPartial) SetSplash(v string) {
	o.Splash.Set(&v)
}

// SetSplashNil sets the value for Splash to be an explicit nil
func (o *GuildPatchRequestPartial) SetSplashNil() {
	o.Splash.Set(nil)
}

// UnsetSplash ensures that no value is present for Splash, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetSplash() {
	o.Splash.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetBanner() string {
	if o == nil || IsNil(o.Banner.Get()) {
		var ret string
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner.Get()) {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasBanner() bool {
	if o != nil && o.Banner.IsSet() {
		return true
	}

	return false
}

// SetBanner gets a reference to the given NullableString and assigns it to the Banner field.
func (o *GuildPatchRequestPartial) SetBanner(v string) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil
func (o *GuildPatchRequestPartial) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetBanner() {
	o.Banner.Unset()
}

// GetSystemChannelFlags returns the SystemChannelFlags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetSystemChannelFlags() int32 {
	if o == nil || IsNil(o.SystemChannelFlags.Get()) {
		var ret int32
		return ret
	}
	return *o.SystemChannelFlags.Get()
}

// GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetSystemChannelFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.SystemChannelFlags.Get()) {
		return nil, false
	}
	return o.SystemChannelFlags.Get(), o.SystemChannelFlags.IsSet()
}

// HasSystemChannelFlags returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasSystemChannelFlags() bool {
	if o != nil && o.SystemChannelFlags.IsSet() {
		return true
	}

	return false
}

// SetSystemChannelFlags gets a reference to the given NullableInt32 and assigns it to the SystemChannelFlags field.
func (o *GuildPatchRequestPartial) SetSystemChannelFlags(v int32) {
	o.SystemChannelFlags.Set(&v)
}

// SetSystemChannelFlagsNil sets the value for SystemChannelFlags to be an explicit nil
func (o *GuildPatchRequestPartial) SetSystemChannelFlagsNil() {
	o.SystemChannelFlags.Set(nil)
}

// UnsetSystemChannelFlags ensures that no value is present for SystemChannelFlags, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetSystemChannelFlags() {
	o.SystemChannelFlags.Unset()
}

// GetFeatures returns the Features field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetFeatures() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetFeaturesOk() ([]*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []*string and assigns it to the Features field.
func (o *GuildPatchRequestPartial) SetFeatures(v []*string) {
	o.Features = v
}


// GetDiscoverySplash returns the DiscoverySplash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetDiscoverySplash() string {
	if o == nil || IsNil(o.DiscoverySplash.Get()) {
		var ret string
		return ret
	}
	return *o.DiscoverySplash.Get()
}

// GetDiscoverySplashOk returns a tuple with the DiscoverySplash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetDiscoverySplashOk() (*string, bool) {
	if o == nil || IsNil(o.DiscoverySplash.Get()) {
		return nil, false
	}
	return o.DiscoverySplash.Get(), o.DiscoverySplash.IsSet()
}

// HasDiscoverySplash returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasDiscoverySplash() bool {
	if o != nil && o.DiscoverySplash.IsSet() {
		return true
	}

	return false
}

// SetDiscoverySplash gets a reference to the given NullableString and assigns it to the DiscoverySplash field.
func (o *GuildPatchRequestPartial) SetDiscoverySplash(v string) {
	o.DiscoverySplash.Set(&v)
}

// SetDiscoverySplashNil sets the value for DiscoverySplash to be an explicit nil
func (o *GuildPatchRequestPartial) SetDiscoverySplashNil() {
	o.DiscoverySplash.Set(nil)
}

// UnsetDiscoverySplash ensures that no value is present for DiscoverySplash, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetDiscoverySplash() {
	o.DiscoverySplash.Unset()
}

// GetHomeHeader returns the HomeHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetHomeHeader() string {
	if o == nil || IsNil(o.HomeHeader.Get()) {
		var ret string
		return ret
	}
	return *o.HomeHeader.Get()
}

// GetHomeHeaderOk returns a tuple with the HomeHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetHomeHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.HomeHeader.Get()) {
		return nil, false
	}
	return o.HomeHeader.Get(), o.HomeHeader.IsSet()
}

// HasHomeHeader returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasHomeHeader() bool {
	if o != nil && o.HomeHeader.IsSet() {
		return true
	}

	return false
}

// SetHomeHeader gets a reference to the given NullableString and assigns it to the HomeHeader field.
func (o *GuildPatchRequestPartial) SetHomeHeader(v string) {
	o.HomeHeader.Set(&v)
}

// SetHomeHeaderNil sets the value for HomeHeader to be an explicit nil
func (o *GuildPatchRequestPartial) SetHomeHeaderNil() {
	o.HomeHeader.Set(nil)
}

// UnsetHomeHeader ensures that no value is present for HomeHeader, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetHomeHeader() {
	o.HomeHeader.Unset()
}

// GetRulesChannelId returns the RulesChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetRulesChannelId() string {
	if o == nil || IsNil(o.RulesChannelId) {
		var ret string
		return ret
	}
	return *o.RulesChannelId
}

// GetRulesChannelIdOk returns a tuple with the RulesChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetRulesChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.RulesChannelId) {
		return nil, false
	}
	return o.RulesChannelId, true
}

// HasRulesChannelId returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasRulesChannelId() bool {
	if o != nil && !IsNil(o.RulesChannelId) {
		return true
	}

	return false
}

// SetRulesChannelId gets a reference to the given string and assigns it to the RulesChannelId field.
func (o *GuildPatchRequestPartial) SetRulesChannelId(v string) {
	o.RulesChannelId = &v
}


// GetSafetyAlertsChannelId returns the SafetyAlertsChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetSafetyAlertsChannelId() string {
	if o == nil || IsNil(o.SafetyAlertsChannelId) {
		var ret string
		return ret
	}
	return *o.SafetyAlertsChannelId
}

// GetSafetyAlertsChannelIdOk returns a tuple with the SafetyAlertsChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetSafetyAlertsChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.SafetyAlertsChannelId) {
		return nil, false
	}
	return o.SafetyAlertsChannelId, true
}

// HasSafetyAlertsChannelId returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasSafetyAlertsChannelId() bool {
	if o != nil && !IsNil(o.SafetyAlertsChannelId) {
		return true
	}

	return false
}

// SetSafetyAlertsChannelId gets a reference to the given string and assigns it to the SafetyAlertsChannelId field.
func (o *GuildPatchRequestPartial) SetSafetyAlertsChannelId(v string) {
	o.SafetyAlertsChannelId = &v
}


// GetPublicUpdatesChannelId returns the PublicUpdatesChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetPublicUpdatesChannelId() string {
	if o == nil || IsNil(o.PublicUpdatesChannelId) {
		var ret string
		return ret
	}
	return *o.PublicUpdatesChannelId
}

// GetPublicUpdatesChannelIdOk returns a tuple with the PublicUpdatesChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetPublicUpdatesChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicUpdatesChannelId) {
		return nil, false
	}
	return o.PublicUpdatesChannelId, true
}

// HasPublicUpdatesChannelId returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasPublicUpdatesChannelId() bool {
	if o != nil && !IsNil(o.PublicUpdatesChannelId) {
		return true
	}

	return false
}

// SetPublicUpdatesChannelId gets a reference to the given string and assigns it to the PublicUpdatesChannelId field.
func (o *GuildPatchRequestPartial) SetPublicUpdatesChannelId(v string) {
	o.PublicUpdatesChannelId = &v
}


// GetPremiumProgressBarEnabled returns the PremiumProgressBarEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPatchRequestPartial) GetPremiumProgressBarEnabled() bool {
	if o == nil || IsNil(o.PremiumProgressBarEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.PremiumProgressBarEnabled.Get()
}

// GetPremiumProgressBarEnabledOk returns a tuple with the PremiumProgressBarEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPatchRequestPartial) GetPremiumProgressBarEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PremiumProgressBarEnabled.Get()) {
		return nil, false
	}
	return o.PremiumProgressBarEnabled.Get(), o.PremiumProgressBarEnabled.IsSet()
}

// HasPremiumProgressBarEnabled returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasPremiumProgressBarEnabled() bool {
	if o != nil && o.PremiumProgressBarEnabled.IsSet() {
		return true
	}

	return false
}

// SetPremiumProgressBarEnabled gets a reference to the given NullableBool and assigns it to the PremiumProgressBarEnabled field.
func (o *GuildPatchRequestPartial) SetPremiumProgressBarEnabled(v bool) {
	o.PremiumProgressBarEnabled.Set(&v)
}

// SetPremiumProgressBarEnabledNil sets the value for PremiumProgressBarEnabled to be an explicit nil
func (o *GuildPatchRequestPartial) SetPremiumProgressBarEnabledNil() {
	o.PremiumProgressBarEnabled.Set(nil)
}

// UnsetPremiumProgressBarEnabled ensures that no value is present for PremiumProgressBarEnabled, not even an explicit nil
func (o *GuildPatchRequestPartial) UnsetPremiumProgressBarEnabled() {
	o.PremiumProgressBarEnabled.Unset()
}

func (o GuildPatchRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildPatchRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if !IsNil(o.VerificationLevel) {
		toSerialize["verification_level"] = o.VerificationLevel
	}
	if !IsNil(o.DefaultMessageNotifications) {
		toSerialize["default_message_notifications"] = o.DefaultMessageNotifications
	}
	if !IsNil(o.ExplicitContentFilter) {
		toSerialize["explicit_content_filter"] = o.ExplicitContentFilter
	}
	if !IsNil(o.PreferredLocale) {
		toSerialize["preferred_locale"] = o.PreferredLocale
	}
	if o.AfkTimeout.IsSet() {
		toSerialize["afk_timeout"] = o.AfkTimeout.Get()
	}
	if !IsNil(o.AfkChannelId) {
		toSerialize["afk_channel_id"] = o.AfkChannelId
	}
	if !IsNil(o.SystemChannelId) {
		toSerialize["system_channel_id"] = o.SystemChannelId
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if o.Splash.IsSet() {
		toSerialize["splash"] = o.Splash.Get()
	}
	if o.Banner.IsSet() {
		toSerialize["banner"] = o.Banner.Get()
	}
	if o.SystemChannelFlags.IsSet() {
		toSerialize["system_channel_flags"] = o.SystemChannelFlags.Get()
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.DiscoverySplash.IsSet() {
		toSerialize["discovery_splash"] = o.DiscoverySplash.Get()
	}
	if o.HomeHeader.IsSet() {
		toSerialize["home_header"] = o.HomeHeader.Get()
	}
	if !IsNil(o.RulesChannelId) {
		toSerialize["rules_channel_id"] = o.RulesChannelId
	}
	if !IsNil(o.SafetyAlertsChannelId) {
		toSerialize["safety_alerts_channel_id"] = o.SafetyAlertsChannelId
	}
	if !IsNil(o.PublicUpdatesChannelId) {
		toSerialize["public_updates_channel_id"] = o.PublicUpdatesChannelId
	}
	if o.PremiumProgressBarEnabled.IsSet() {
		toSerialize["premium_progress_bar_enabled"] = o.PremiumProgressBarEnabled.Get()
	}
	return toSerialize, nil
}

type NullableGuildPatchRequestPartial struct {
	value *GuildPatchRequestPartial
	isSet bool
}

func (v NullableGuildPatchRequestPartial) Get() *GuildPatchRequestPartial {
	return v.value
}

func (v *NullableGuildPatchRequestPartial) Set(val *GuildPatchRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildPatchRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildPatchRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildPatchRequestPartial(val *GuildPatchRequestPartial) *NullableGuildPatchRequestPartial {
	return &NullableGuildPatchRequestPartial{value: val, isSet: true}
}

func (v NullableGuildPatchRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildPatchRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


