/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
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
)

// checks if the GuildPatchRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildPatchRequestPartial{}

// GuildPatchRequestPartial struct for GuildPatchRequestPartial
type GuildPatchRequestPartial struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Region *string `json:"region,omitempty"`
	Icon *string `json:"icon,omitempty"`
	VerificationLevel *GuildPatchRequestPartialVerificationLevel `json:"verification_level,omitempty"`
	DefaultMessageNotifications *GuildPatchRequestPartialDefaultMessageNotifications `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter *GuildPatchRequestPartialExplicitContentFilter `json:"explicit_content_filter,omitempty"`
	PreferredLocale *GuildPatchRequestPartialPreferredLocale `json:"preferred_locale,omitempty"`
	AfkTimeout *GuildPatchRequestPartialAfkTimeout `json:"afk_timeout,omitempty"`
	AfkChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"afk_channel_id,omitempty"`
	SystemChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"system_channel_id,omitempty"`
	Splash *string `json:"splash,omitempty"`
	Banner *string `json:"banner,omitempty"`
	SystemChannelFlags *int32 `json:"system_channel_flags,omitempty"`
	Features []string `json:"features,omitempty"`
	DiscoverySplash *string `json:"discovery_splash,omitempty"`
	HomeHeader *string `json:"home_header,omitempty"`
	RulesChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"rules_channel_id,omitempty"`
	SafetyAlertsChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"safety_alerts_channel_id,omitempty"`
	PublicUpdatesChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"public_updates_channel_id,omitempty"`
	PremiumProgressBarEnabled *bool `json:"premium_progress_bar_enabled,omitempty"`
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


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GuildPatchRequestPartial) SetDescription(v string) {
	o.Description = &v
}


// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GuildPatchRequestPartial) SetRegion(v string) {
	o.Region = &v
}


// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *GuildPatchRequestPartial) SetIcon(v string) {
	o.Icon = &v
}


// GetVerificationLevel returns the VerificationLevel field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetVerificationLevel() GuildPatchRequestPartialVerificationLevel {
	if o == nil || IsNil(o.VerificationLevel) {
		var ret GuildPatchRequestPartialVerificationLevel
		return ret
	}
	return *o.VerificationLevel
}

// GetVerificationLevelOk returns a tuple with the VerificationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetVerificationLevelOk() (*GuildPatchRequestPartialVerificationLevel, bool) {
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

// SetVerificationLevel gets a reference to the given GuildPatchRequestPartialVerificationLevel and assigns it to the VerificationLevel field.
func (o *GuildPatchRequestPartial) SetVerificationLevel(v GuildPatchRequestPartialVerificationLevel) {
	o.VerificationLevel = &v
}


// GetDefaultMessageNotifications returns the DefaultMessageNotifications field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetDefaultMessageNotifications() GuildPatchRequestPartialDefaultMessageNotifications {
	if o == nil || IsNil(o.DefaultMessageNotifications) {
		var ret GuildPatchRequestPartialDefaultMessageNotifications
		return ret
	}
	return *o.DefaultMessageNotifications
}

// GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetDefaultMessageNotificationsOk() (*GuildPatchRequestPartialDefaultMessageNotifications, bool) {
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

// SetDefaultMessageNotifications gets a reference to the given GuildPatchRequestPartialDefaultMessageNotifications and assigns it to the DefaultMessageNotifications field.
func (o *GuildPatchRequestPartial) SetDefaultMessageNotifications(v GuildPatchRequestPartialDefaultMessageNotifications) {
	o.DefaultMessageNotifications = &v
}


// GetExplicitContentFilter returns the ExplicitContentFilter field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetExplicitContentFilter() GuildPatchRequestPartialExplicitContentFilter {
	if o == nil || IsNil(o.ExplicitContentFilter) {
		var ret GuildPatchRequestPartialExplicitContentFilter
		return ret
	}
	return *o.ExplicitContentFilter
}

// GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetExplicitContentFilterOk() (*GuildPatchRequestPartialExplicitContentFilter, bool) {
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

// SetExplicitContentFilter gets a reference to the given GuildPatchRequestPartialExplicitContentFilter and assigns it to the ExplicitContentFilter field.
func (o *GuildPatchRequestPartial) SetExplicitContentFilter(v GuildPatchRequestPartialExplicitContentFilter) {
	o.ExplicitContentFilter = &v
}


// GetPreferredLocale returns the PreferredLocale field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetPreferredLocale() GuildPatchRequestPartialPreferredLocale {
	if o == nil || IsNil(o.PreferredLocale) {
		var ret GuildPatchRequestPartialPreferredLocale
		return ret
	}
	return *o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetPreferredLocaleOk() (*GuildPatchRequestPartialPreferredLocale, bool) {
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

// SetPreferredLocale gets a reference to the given GuildPatchRequestPartialPreferredLocale and assigns it to the PreferredLocale field.
func (o *GuildPatchRequestPartial) SetPreferredLocale(v GuildPatchRequestPartialPreferredLocale) {
	o.PreferredLocale = &v
}


// GetAfkTimeout returns the AfkTimeout field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetAfkTimeout() GuildPatchRequestPartialAfkTimeout {
	if o == nil || IsNil(o.AfkTimeout) {
		var ret GuildPatchRequestPartialAfkTimeout
		return ret
	}
	return *o.AfkTimeout
}

// GetAfkTimeoutOk returns a tuple with the AfkTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetAfkTimeoutOk() (*GuildPatchRequestPartialAfkTimeout, bool) {
	if o == nil || IsNil(o.AfkTimeout) {
		return nil, false
	}
	return o.AfkTimeout, true
}

// HasAfkTimeout returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasAfkTimeout() bool {
	if o != nil && !IsNil(o.AfkTimeout) {
		return true
	}

	return false
}

// SetAfkTimeout gets a reference to the given GuildPatchRequestPartialAfkTimeout and assigns it to the AfkTimeout field.
func (o *GuildPatchRequestPartial) SetAfkTimeout(v GuildPatchRequestPartialAfkTimeout) {
	o.AfkTimeout = &v
}


// GetAfkChannelId returns the AfkChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetAfkChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.AfkChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.AfkChannelId
}

// GetAfkChannelIdOk returns a tuple with the AfkChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetAfkChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetAfkChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the AfkChannelId field.
func (o *GuildPatchRequestPartial) SetAfkChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.AfkChannelId = &v
}


// GetSystemChannelId returns the SystemChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetSystemChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.SystemChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.SystemChannelId
}

// GetSystemChannelIdOk returns a tuple with the SystemChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetSystemChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetSystemChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the SystemChannelId field.
func (o *GuildPatchRequestPartial) SetSystemChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.SystemChannelId = &v
}


// GetSplash returns the Splash field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetSplash() string {
	if o == nil || IsNil(o.Splash) {
		var ret string
		return ret
	}
	return *o.Splash
}

// GetSplashOk returns a tuple with the Splash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetSplashOk() (*string, bool) {
	if o == nil || IsNil(o.Splash) {
		return nil, false
	}
	return o.Splash, true
}

// HasSplash returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasSplash() bool {
	if o != nil && !IsNil(o.Splash) {
		return true
	}

	return false
}

// SetSplash gets a reference to the given string and assigns it to the Splash field.
func (o *GuildPatchRequestPartial) SetSplash(v string) {
	o.Splash = &v
}


// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetBanner() string {
	if o == nil || IsNil(o.Banner) {
		var ret string
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given string and assigns it to the Banner field.
func (o *GuildPatchRequestPartial) SetBanner(v string) {
	o.Banner = &v
}


// GetSystemChannelFlags returns the SystemChannelFlags field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetSystemChannelFlags() int32 {
	if o == nil || IsNil(o.SystemChannelFlags) {
		var ret int32
		return ret
	}
	return *o.SystemChannelFlags
}

// GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetSystemChannelFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.SystemChannelFlags) {
		return nil, false
	}
	return o.SystemChannelFlags, true
}

// HasSystemChannelFlags returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasSystemChannelFlags() bool {
	if o != nil && !IsNil(o.SystemChannelFlags) {
		return true
	}

	return false
}

// SetSystemChannelFlags gets a reference to the given int32 and assigns it to the SystemChannelFlags field.
func (o *GuildPatchRequestPartial) SetSystemChannelFlags(v int32) {
	o.SystemChannelFlags = &v
}


// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetFeatures() []string {
	if o == nil || IsNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Features) {
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

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *GuildPatchRequestPartial) SetFeatures(v []string) {
	o.Features = v
}


// GetDiscoverySplash returns the DiscoverySplash field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetDiscoverySplash() string {
	if o == nil || IsNil(o.DiscoverySplash) {
		var ret string
		return ret
	}
	return *o.DiscoverySplash
}

// GetDiscoverySplashOk returns a tuple with the DiscoverySplash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetDiscoverySplashOk() (*string, bool) {
	if o == nil || IsNil(o.DiscoverySplash) {
		return nil, false
	}
	return o.DiscoverySplash, true
}

// HasDiscoverySplash returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasDiscoverySplash() bool {
	if o != nil && !IsNil(o.DiscoverySplash) {
		return true
	}

	return false
}

// SetDiscoverySplash gets a reference to the given string and assigns it to the DiscoverySplash field.
func (o *GuildPatchRequestPartial) SetDiscoverySplash(v string) {
	o.DiscoverySplash = &v
}


// GetHomeHeader returns the HomeHeader field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetHomeHeader() string {
	if o == nil || IsNil(o.HomeHeader) {
		var ret string
		return ret
	}
	return *o.HomeHeader
}

// GetHomeHeaderOk returns a tuple with the HomeHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetHomeHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.HomeHeader) {
		return nil, false
	}
	return o.HomeHeader, true
}

// HasHomeHeader returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasHomeHeader() bool {
	if o != nil && !IsNil(o.HomeHeader) {
		return true
	}

	return false
}

// SetHomeHeader gets a reference to the given string and assigns it to the HomeHeader field.
func (o *GuildPatchRequestPartial) SetHomeHeader(v string) {
	o.HomeHeader = &v
}


// GetRulesChannelId returns the RulesChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetRulesChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.RulesChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.RulesChannelId
}

// GetRulesChannelIdOk returns a tuple with the RulesChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetRulesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetRulesChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the RulesChannelId field.
func (o *GuildPatchRequestPartial) SetRulesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.RulesChannelId = &v
}


// GetSafetyAlertsChannelId returns the SafetyAlertsChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetSafetyAlertsChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.SafetyAlertsChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.SafetyAlertsChannelId
}

// GetSafetyAlertsChannelIdOk returns a tuple with the SafetyAlertsChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetSafetyAlertsChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetSafetyAlertsChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the SafetyAlertsChannelId field.
func (o *GuildPatchRequestPartial) SetSafetyAlertsChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.SafetyAlertsChannelId = &v
}


// GetPublicUpdatesChannelId returns the PublicUpdatesChannelId field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetPublicUpdatesChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.PublicUpdatesChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.PublicUpdatesChannelId
}

// GetPublicUpdatesChannelIdOk returns a tuple with the PublicUpdatesChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetPublicUpdatesChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetPublicUpdatesChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the PublicUpdatesChannelId field.
func (o *GuildPatchRequestPartial) SetPublicUpdatesChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.PublicUpdatesChannelId = &v
}


// GetPremiumProgressBarEnabled returns the PremiumProgressBarEnabled field value if set, zero value otherwise.
func (o *GuildPatchRequestPartial) GetPremiumProgressBarEnabled() bool {
	if o == nil || IsNil(o.PremiumProgressBarEnabled) {
		var ret bool
		return ret
	}
	return *o.PremiumProgressBarEnabled
}

// GetPremiumProgressBarEnabledOk returns a tuple with the PremiumProgressBarEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildPatchRequestPartial) GetPremiumProgressBarEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PremiumProgressBarEnabled) {
		return nil, false
	}
	return o.PremiumProgressBarEnabled, true
}

// HasPremiumProgressBarEnabled returns a boolean if a field has been set.
func (o *GuildPatchRequestPartial) HasPremiumProgressBarEnabled() bool {
	if o != nil && !IsNil(o.PremiumProgressBarEnabled) {
		return true
	}

	return false
}

// SetPremiumProgressBarEnabled gets a reference to the given bool and assigns it to the PremiumProgressBarEnabled field.
func (o *GuildPatchRequestPartial) SetPremiumProgressBarEnabled(v bool) {
	o.PremiumProgressBarEnabled = &v
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
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
	if !IsNil(o.AfkTimeout) {
		toSerialize["afk_timeout"] = o.AfkTimeout
	}
	if !IsNil(o.AfkChannelId) {
		toSerialize["afk_channel_id"] = o.AfkChannelId
	}
	if !IsNil(o.SystemChannelId) {
		toSerialize["system_channel_id"] = o.SystemChannelId
	}
	if !IsNil(o.Splash) {
		toSerialize["splash"] = o.Splash
	}
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	if !IsNil(o.SystemChannelFlags) {
		toSerialize["system_channel_flags"] = o.SystemChannelFlags
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.DiscoverySplash) {
		toSerialize["discovery_splash"] = o.DiscoverySplash
	}
	if !IsNil(o.HomeHeader) {
		toSerialize["home_header"] = o.HomeHeader
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
	if !IsNil(o.PremiumProgressBarEnabled) {
		toSerialize["premium_progress_bar_enabled"] = o.PremiumProgressBarEnabled
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

