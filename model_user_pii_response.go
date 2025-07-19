/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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

// checks if the UserPIIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPIIResponse{}

// UserPIIResponse struct for UserPIIResponse
type UserPIIResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Username string `json:"username"`
	Avatar NullableString `json:"avatar,omitempty"`
	Discriminator string `json:"discriminator"`
	PublicFlags int32 `json:"public_flags"`
	Flags int64 `json:"flags"`
	Bot NullableBool `json:"bot,omitempty"`
	System NullableBool `json:"system,omitempty"`
	Banner NullableString `json:"banner,omitempty"`
	AccentColor NullableInt32 `json:"accent_color,omitempty"`
	GlobalName NullableString `json:"global_name,omitempty"`
	AvatarDecorationData NullableUserAvatarDecorationResponse `json:"avatar_decoration_data,omitempty"`
	Collectibles NullableUserCollectiblesResponse `json:"collectibles,omitempty"`
	PrimaryGuild NullableUserPrimaryGuildResponse `json:"primary_guild,omitempty"`
	MfaEnabled bool `json:"mfa_enabled"`
	Locale string `json:"locale"`
	PremiumType *int32 `json:"premium_type,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Verified NullableBool `json:"verified,omitempty"`
}

type _UserPIIResponse UserPIIResponse

// NewUserPIIResponse instantiates a new UserPIIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPIIResponse(id string, username string, discriminator string, publicFlags int32, flags int64, mfaEnabled bool, locale string) *UserPIIResponse {
	this := UserPIIResponse{}
	this.Id = id
	this.Username = username
	this.Discriminator = discriminator
	this.PublicFlags = publicFlags
	this.Flags = flags
	this.MfaEnabled = mfaEnabled
	this.Locale = locale
	return &this
}

// NewUserPIIResponseWithDefaults instantiates a new UserPIIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPIIResponseWithDefaults() *UserPIIResponse {
	this := UserPIIResponse{}
	return &this
}

// GetId returns the Id field value
func (o *UserPIIResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserPIIResponse) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *UserPIIResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserPIIResponse) SetUsername(v string) {
	o.Username = v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar.Get()) {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *UserPIIResponse) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *UserPIIResponse) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *UserPIIResponse) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *UserPIIResponse) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetDiscriminator returns the Discriminator field value
func (o *UserPIIResponse) GetDiscriminator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Discriminator
}

// GetDiscriminatorOk returns a tuple with the Discriminator field value
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetDiscriminatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discriminator, true
}

// SetDiscriminator sets field value
func (o *UserPIIResponse) SetDiscriminator(v string) {
	o.Discriminator = v
}

// GetPublicFlags returns the PublicFlags field value
func (o *UserPIIResponse) GetPublicFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PublicFlags
}

// GetPublicFlagsOk returns a tuple with the PublicFlags field value
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetPublicFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicFlags, true
}

// SetPublicFlags sets field value
func (o *UserPIIResponse) SetPublicFlags(v int32) {
	o.PublicFlags = v
}

// GetFlags returns the Flags field value
func (o *UserPIIResponse) GetFlags() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetFlagsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *UserPIIResponse) SetFlags(v int64) {
	o.Flags = v
}

// GetBot returns the Bot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetBot() bool {
	if o == nil || IsNil(o.Bot.Get()) {
		var ret bool
		return ret
	}
	return *o.Bot.Get()
}

// GetBotOk returns a tuple with the Bot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetBotOk() (*bool, bool) {
	if o == nil || IsNil(o.Bot.Get()) {
		return nil, false
	}
	return o.Bot.Get(), o.Bot.IsSet()
}

// HasBot returns a boolean if a field has been set.
func (o *UserPIIResponse) HasBot() bool {
	if o != nil && o.Bot.IsSet() {
		return true
	}

	return false
}

// SetBot gets a reference to the given NullableBool and assigns it to the Bot field.
func (o *UserPIIResponse) SetBot(v bool) {
	o.Bot.Set(&v)
}

// SetBotNil sets the value for Bot to be an explicit nil
func (o *UserPIIResponse) SetBotNil() {
	o.Bot.Set(nil)
}

// UnsetBot ensures that no value is present for Bot, not even an explicit nil
func (o *UserPIIResponse) UnsetBot() {
	o.Bot.Unset()
}

// GetSystem returns the System field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetSystem() bool {
	if o == nil || IsNil(o.System.Get()) {
		var ret bool
		return ret
	}
	return *o.System.Get()
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.System.Get()) {
		return nil, false
	}
	return o.System.Get(), o.System.IsSet()
}

// HasSystem returns a boolean if a field has been set.
func (o *UserPIIResponse) HasSystem() bool {
	if o != nil && o.System.IsSet() {
		return true
	}

	return false
}

// SetSystem gets a reference to the given NullableBool and assigns it to the System field.
func (o *UserPIIResponse) SetSystem(v bool) {
	o.System.Set(&v)
}

// SetSystemNil sets the value for System to be an explicit nil
func (o *UserPIIResponse) SetSystemNil() {
	o.System.Set(nil)
}

// UnsetSystem ensures that no value is present for System, not even an explicit nil
func (o *UserPIIResponse) UnsetSystem() {
	o.System.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner.Get()) {
		var ret string
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner.Get()) {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *UserPIIResponse) HasBanner() bool {
	if o != nil && o.Banner.IsSet() {
		return true
	}

	return false
}

// SetBanner gets a reference to the given NullableString and assigns it to the Banner field.
func (o *UserPIIResponse) SetBanner(v string) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil
func (o *UserPIIResponse) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil
func (o *UserPIIResponse) UnsetBanner() {
	o.Banner.Unset()
}

// GetAccentColor returns the AccentColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetAccentColor() int32 {
	if o == nil || IsNil(o.AccentColor.Get()) {
		var ret int32
		return ret
	}
	return *o.AccentColor.Get()
}

// GetAccentColorOk returns a tuple with the AccentColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetAccentColorOk() (*int32, bool) {
	if o == nil || IsNil(o.AccentColor.Get()) {
		return nil, false
	}
	return o.AccentColor.Get(), o.AccentColor.IsSet()
}

// HasAccentColor returns a boolean if a field has been set.
func (o *UserPIIResponse) HasAccentColor() bool {
	if o != nil && o.AccentColor.IsSet() {
		return true
	}

	return false
}

// SetAccentColor gets a reference to the given NullableInt32 and assigns it to the AccentColor field.
func (o *UserPIIResponse) SetAccentColor(v int32) {
	o.AccentColor.Set(&v)
}

// SetAccentColorNil sets the value for AccentColor to be an explicit nil
func (o *UserPIIResponse) SetAccentColorNil() {
	o.AccentColor.Set(nil)
}

// UnsetAccentColor ensures that no value is present for AccentColor, not even an explicit nil
func (o *UserPIIResponse) UnsetAccentColor() {
	o.AccentColor.Unset()
}

// GetGlobalName returns the GlobalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetGlobalName() string {
	if o == nil || IsNil(o.GlobalName.Get()) {
		var ret string
		return ret
	}
	return *o.GlobalName.Get()
}

// GetGlobalNameOk returns a tuple with the GlobalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetGlobalNameOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalName.Get()) {
		return nil, false
	}
	return o.GlobalName.Get(), o.GlobalName.IsSet()
}

// HasGlobalName returns a boolean if a field has been set.
func (o *UserPIIResponse) HasGlobalName() bool {
	if o != nil && o.GlobalName.IsSet() {
		return true
	}

	return false
}

// SetGlobalName gets a reference to the given NullableString and assigns it to the GlobalName field.
func (o *UserPIIResponse) SetGlobalName(v string) {
	o.GlobalName.Set(&v)
}

// SetGlobalNameNil sets the value for GlobalName to be an explicit nil
func (o *UserPIIResponse) SetGlobalNameNil() {
	o.GlobalName.Set(nil)
}

// UnsetGlobalName ensures that no value is present for GlobalName, not even an explicit nil
func (o *UserPIIResponse) UnsetGlobalName() {
	o.GlobalName.Unset()
}

// GetAvatarDecorationData returns the AvatarDecorationData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetAvatarDecorationData() UserAvatarDecorationResponse {
	if o == nil || IsNil(o.AvatarDecorationData.Get()) {
		var ret UserAvatarDecorationResponse
		return ret
	}
	return *o.AvatarDecorationData.Get()
}

// GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetAvatarDecorationDataOk() (*UserAvatarDecorationResponse, bool) {
	if o == nil || IsNil(o.AvatarDecorationData.Get()) {
		return nil, false
	}
	return o.AvatarDecorationData.Get(), o.AvatarDecorationData.IsSet()
}

// HasAvatarDecorationData returns a boolean if a field has been set.
func (o *UserPIIResponse) HasAvatarDecorationData() bool {
	if o != nil && o.AvatarDecorationData.IsSet() {
		return true
	}

	return false
}

// SetAvatarDecorationData gets a reference to the given NullableUserAvatarDecorationResponse and assigns it to the AvatarDecorationData field.
func (o *UserPIIResponse) SetAvatarDecorationData(v UserAvatarDecorationResponse) {
	o.AvatarDecorationData.Set(&v)
}

// SetAvatarDecorationDataNil sets the value for AvatarDecorationData to be an explicit nil
func (o *UserPIIResponse) SetAvatarDecorationDataNil() {
	o.AvatarDecorationData.Set(nil)
}

// UnsetAvatarDecorationData ensures that no value is present for AvatarDecorationData, not even an explicit nil
func (o *UserPIIResponse) UnsetAvatarDecorationData() {
	o.AvatarDecorationData.Unset()
}

// GetCollectibles returns the Collectibles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetCollectibles() UserCollectiblesResponse {
	if o == nil || IsNil(o.Collectibles.Get()) {
		var ret UserCollectiblesResponse
		return ret
	}
	return *o.Collectibles.Get()
}

// GetCollectiblesOk returns a tuple with the Collectibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetCollectiblesOk() (*UserCollectiblesResponse, bool) {
	if o == nil || IsNil(o.Collectibles.Get()) {
		return nil, false
	}
	return o.Collectibles.Get(), o.Collectibles.IsSet()
}

// HasCollectibles returns a boolean if a field has been set.
func (o *UserPIIResponse) HasCollectibles() bool {
	if o != nil && o.Collectibles.IsSet() {
		return true
	}

	return false
}

// SetCollectibles gets a reference to the given NullableUserCollectiblesResponse and assigns it to the Collectibles field.
func (o *UserPIIResponse) SetCollectibles(v UserCollectiblesResponse) {
	o.Collectibles.Set(&v)
}

// SetCollectiblesNil sets the value for Collectibles to be an explicit nil
func (o *UserPIIResponse) SetCollectiblesNil() {
	o.Collectibles.Set(nil)
}

// UnsetCollectibles ensures that no value is present for Collectibles, not even an explicit nil
func (o *UserPIIResponse) UnsetCollectibles() {
	o.Collectibles.Unset()
}

// GetPrimaryGuild returns the PrimaryGuild field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetPrimaryGuild() UserPrimaryGuildResponse {
	if o == nil || IsNil(o.PrimaryGuild.Get()) {
		var ret UserPrimaryGuildResponse
		return ret
	}
	return *o.PrimaryGuild.Get()
}

// GetPrimaryGuildOk returns a tuple with the PrimaryGuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetPrimaryGuildOk() (*UserPrimaryGuildResponse, bool) {
	if o == nil || IsNil(o.PrimaryGuild.Get()) {
		return nil, false
	}
	return o.PrimaryGuild.Get(), o.PrimaryGuild.IsSet()
}

// HasPrimaryGuild returns a boolean if a field has been set.
func (o *UserPIIResponse) HasPrimaryGuild() bool {
	if o != nil && o.PrimaryGuild.IsSet() {
		return true
	}

	return false
}

// SetPrimaryGuild gets a reference to the given NullableUserPrimaryGuildResponse and assigns it to the PrimaryGuild field.
func (o *UserPIIResponse) SetPrimaryGuild(v UserPrimaryGuildResponse) {
	o.PrimaryGuild.Set(&v)
}

// SetPrimaryGuildNil sets the value for PrimaryGuild to be an explicit nil
func (o *UserPIIResponse) SetPrimaryGuildNil() {
	o.PrimaryGuild.Set(nil)
}

// UnsetPrimaryGuild ensures that no value is present for PrimaryGuild, not even an explicit nil
func (o *UserPIIResponse) UnsetPrimaryGuild() {
	o.PrimaryGuild.Unset()
}

// GetMfaEnabled returns the MfaEnabled field value
func (o *UserPIIResponse) GetMfaEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MfaEnabled
}

// GetMfaEnabledOk returns a tuple with the MfaEnabled field value
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetMfaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MfaEnabled, true
}

// SetMfaEnabled sets field value
func (o *UserPIIResponse) SetMfaEnabled(v bool) {
	o.MfaEnabled = v
}

// GetLocale returns the Locale field value
func (o *UserPIIResponse) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *UserPIIResponse) SetLocale(v string) {
	o.Locale = v
}

// GetPremiumType returns the PremiumType field value if set, zero value otherwise.
func (o *UserPIIResponse) GetPremiumType() int32 {
	if o == nil || IsNil(o.PremiumType) {
		var ret int32
		return ret
	}
	return *o.PremiumType
}

// GetPremiumTypeOk returns a tuple with the PremiumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetPremiumTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.PremiumType) {
		return nil, false
	}
	return o.PremiumType, true
}

// HasPremiumType returns a boolean if a field has been set.
func (o *UserPIIResponse) HasPremiumType() bool {
	if o != nil && !IsNil(o.PremiumType) {
		return true
	}

	return false
}

// SetPremiumType gets a reference to the given int32 and assigns it to the PremiumType field.
func (o *UserPIIResponse) SetPremiumType(v int32) {
	o.PremiumType = &v
}


// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email.Get()) {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserPIIResponse) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserPIIResponse) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserPIIResponse) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserPIIResponse) UnsetEmail() {
	o.Email.Unset()
}

// GetVerified returns the Verified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPIIResponse) GetVerified() bool {
	if o == nil || IsNil(o.Verified.Get()) {
		var ret bool
		return ret
	}
	return *o.Verified.Get()
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPIIResponse) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified.Get()) {
		return nil, false
	}
	return o.Verified.Get(), o.Verified.IsSet()
}

// HasVerified returns a boolean if a field has been set.
func (o *UserPIIResponse) HasVerified() bool {
	if o != nil && o.Verified.IsSet() {
		return true
	}

	return false
}

// SetVerified gets a reference to the given NullableBool and assigns it to the Verified field.
func (o *UserPIIResponse) SetVerified(v bool) {
	o.Verified.Set(&v)
}

// SetVerifiedNil sets the value for Verified to be an explicit nil
func (o *UserPIIResponse) SetVerifiedNil() {
	o.Verified.Set(nil)
}

// UnsetVerified ensures that no value is present for Verified, not even an explicit nil
func (o *UserPIIResponse) UnsetVerified() {
	o.Verified.Unset()
}

func (o UserPIIResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPIIResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	toSerialize["discriminator"] = o.Discriminator
	toSerialize["public_flags"] = o.PublicFlags
	toSerialize["flags"] = o.Flags
	if o.Bot.IsSet() {
		toSerialize["bot"] = o.Bot.Get()
	}
	if o.System.IsSet() {
		toSerialize["system"] = o.System.Get()
	}
	if o.Banner.IsSet() {
		toSerialize["banner"] = o.Banner.Get()
	}
	if o.AccentColor.IsSet() {
		toSerialize["accent_color"] = o.AccentColor.Get()
	}
	if o.GlobalName.IsSet() {
		toSerialize["global_name"] = o.GlobalName.Get()
	}
	if o.AvatarDecorationData.IsSet() {
		toSerialize["avatar_decoration_data"] = o.AvatarDecorationData.Get()
	}
	if o.Collectibles.IsSet() {
		toSerialize["collectibles"] = o.Collectibles.Get()
	}
	if o.PrimaryGuild.IsSet() {
		toSerialize["primary_guild"] = o.PrimaryGuild.Get()
	}
	toSerialize["mfa_enabled"] = o.MfaEnabled
	toSerialize["locale"] = o.Locale
	if !IsNil(o.PremiumType) {
		toSerialize["premium_type"] = o.PremiumType
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Verified.IsSet() {
		toSerialize["verified"] = o.Verified.Get()
	}
	return toSerialize, nil
}

func (o *UserPIIResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"discriminator",
		"public_flags",
		"flags",
		"mfa_enabled",
		"locale",
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

	varUserPIIResponse := _UserPIIResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserPIIResponse)

	if err != nil {
		return err
	}

	*o = UserPIIResponse(varUserPIIResponse)

	return err
}

type NullableUserPIIResponse struct {
	value *UserPIIResponse
	isSet bool
}

func (v NullableUserPIIResponse) Get() *UserPIIResponse {
	return v.value
}

func (v *NullableUserPIIResponse) Set(val *UserPIIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPIIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPIIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPIIResponse(val *UserPIIResponse) *NullableUserPIIResponse {
	return &NullableUserPIIResponse{value: val, isSet: true}
}

func (v NullableUserPIIResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPIIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


