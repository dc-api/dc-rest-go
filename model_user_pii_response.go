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
	"bytes"
	"fmt"
)

// checks if the UserPIIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPIIResponse{}

// UserPIIResponse struct for UserPIIResponse
type UserPIIResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Username string `json:"username"`
	Avatar *string `json:"avatar,omitempty"`
	Discriminator string `json:"discriminator"`
	PublicFlags int32 `json:"public_flags"`
	Flags int64 `json:"flags"`
	Bot *bool `json:"bot,omitempty"`
	System *bool `json:"system,omitempty"`
	Banner *string `json:"banner,omitempty"`
	AccentColor *int32 `json:"accent_color,omitempty"`
	GlobalName *string `json:"global_name,omitempty"`
	AvatarDecorationData *BasicGuildMemberResponseAvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	Collectibles *BasicGuildMemberResponseCollectibles `json:"collectibles,omitempty"`
	PrimaryGuild *UserPIIResponsePrimaryGuild `json:"primary_guild,omitempty"`
	MfaEnabled bool `json:"mfa_enabled"`
	Locale AvailableLocalesEnum `json:"locale"`
	PremiumType *PremiumTypes `json:"premium_type,omitempty"`
	Email *string `json:"email,omitempty"`
	Verified *bool `json:"verified,omitempty"`
}

type _UserPIIResponse UserPIIResponse

// NewUserPIIResponse instantiates a new UserPIIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPIIResponse(id string, username string, discriminator string, publicFlags int32, flags int64, mfaEnabled bool, locale AvailableLocalesEnum) *UserPIIResponse {
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

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *UserPIIResponse) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *UserPIIResponse) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *UserPIIResponse) SetAvatar(v string) {
	o.Avatar = &v
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

// GetBot returns the Bot field value if set, zero value otherwise.
func (o *UserPIIResponse) GetBot() bool {
	if o == nil || IsNil(o.Bot) {
		var ret bool
		return ret
	}
	return *o.Bot
}

// GetBotOk returns a tuple with the Bot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetBotOk() (*bool, bool) {
	if o == nil || IsNil(o.Bot) {
		return nil, false
	}
	return o.Bot, true
}

// HasBot returns a boolean if a field has been set.
func (o *UserPIIResponse) HasBot() bool {
	if o != nil && !IsNil(o.Bot) {
		return true
	}

	return false
}

// SetBot gets a reference to the given bool and assigns it to the Bot field.
func (o *UserPIIResponse) SetBot(v bool) {
	o.Bot = &v
}


// GetSystem returns the System field value if set, zero value otherwise.
func (o *UserPIIResponse) GetSystem() bool {
	if o == nil || IsNil(o.System) {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *UserPIIResponse) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *UserPIIResponse) SetSystem(v bool) {
	o.System = &v
}


// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *UserPIIResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner) {
		var ret string
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *UserPIIResponse) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given string and assigns it to the Banner field.
func (o *UserPIIResponse) SetBanner(v string) {
	o.Banner = &v
}


// GetAccentColor returns the AccentColor field value if set, zero value otherwise.
func (o *UserPIIResponse) GetAccentColor() int32 {
	if o == nil || IsNil(o.AccentColor) {
		var ret int32
		return ret
	}
	return *o.AccentColor
}

// GetAccentColorOk returns a tuple with the AccentColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetAccentColorOk() (*int32, bool) {
	if o == nil || IsNil(o.AccentColor) {
		return nil, false
	}
	return o.AccentColor, true
}

// HasAccentColor returns a boolean if a field has been set.
func (o *UserPIIResponse) HasAccentColor() bool {
	if o != nil && !IsNil(o.AccentColor) {
		return true
	}

	return false
}

// SetAccentColor gets a reference to the given int32 and assigns it to the AccentColor field.
func (o *UserPIIResponse) SetAccentColor(v int32) {
	o.AccentColor = &v
}


// GetGlobalName returns the GlobalName field value if set, zero value otherwise.
func (o *UserPIIResponse) GetGlobalName() string {
	if o == nil || IsNil(o.GlobalName) {
		var ret string
		return ret
	}
	return *o.GlobalName
}

// GetGlobalNameOk returns a tuple with the GlobalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetGlobalNameOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalName) {
		return nil, false
	}
	return o.GlobalName, true
}

// HasGlobalName returns a boolean if a field has been set.
func (o *UserPIIResponse) HasGlobalName() bool {
	if o != nil && !IsNil(o.GlobalName) {
		return true
	}

	return false
}

// SetGlobalName gets a reference to the given string and assigns it to the GlobalName field.
func (o *UserPIIResponse) SetGlobalName(v string) {
	o.GlobalName = &v
}


// GetAvatarDecorationData returns the AvatarDecorationData field value if set, zero value otherwise.
func (o *UserPIIResponse) GetAvatarDecorationData() BasicGuildMemberResponseAvatarDecorationData {
	if o == nil || IsNil(o.AvatarDecorationData) {
		var ret BasicGuildMemberResponseAvatarDecorationData
		return ret
	}
	return *o.AvatarDecorationData
}

// GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetAvatarDecorationDataOk() (*BasicGuildMemberResponseAvatarDecorationData, bool) {
	if o == nil || IsNil(o.AvatarDecorationData) {
		return nil, false
	}
	return o.AvatarDecorationData, true
}

// HasAvatarDecorationData returns a boolean if a field has been set.
func (o *UserPIIResponse) HasAvatarDecorationData() bool {
	if o != nil && !IsNil(o.AvatarDecorationData) {
		return true
	}

	return false
}

// SetAvatarDecorationData gets a reference to the given BasicGuildMemberResponseAvatarDecorationData and assigns it to the AvatarDecorationData field.
func (o *UserPIIResponse) SetAvatarDecorationData(v BasicGuildMemberResponseAvatarDecorationData) {
	o.AvatarDecorationData = &v
}


// GetCollectibles returns the Collectibles field value if set, zero value otherwise.
func (o *UserPIIResponse) GetCollectibles() BasicGuildMemberResponseCollectibles {
	if o == nil || IsNil(o.Collectibles) {
		var ret BasicGuildMemberResponseCollectibles
		return ret
	}
	return *o.Collectibles
}

// GetCollectiblesOk returns a tuple with the Collectibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetCollectiblesOk() (*BasicGuildMemberResponseCollectibles, bool) {
	if o == nil || IsNil(o.Collectibles) {
		return nil, false
	}
	return o.Collectibles, true
}

// HasCollectibles returns a boolean if a field has been set.
func (o *UserPIIResponse) HasCollectibles() bool {
	if o != nil && !IsNil(o.Collectibles) {
		return true
	}

	return false
}

// SetCollectibles gets a reference to the given BasicGuildMemberResponseCollectibles and assigns it to the Collectibles field.
func (o *UserPIIResponse) SetCollectibles(v BasicGuildMemberResponseCollectibles) {
	o.Collectibles = &v
}


// GetPrimaryGuild returns the PrimaryGuild field value if set, zero value otherwise.
func (o *UserPIIResponse) GetPrimaryGuild() UserPIIResponsePrimaryGuild {
	if o == nil || IsNil(o.PrimaryGuild) {
		var ret UserPIIResponsePrimaryGuild
		return ret
	}
	return *o.PrimaryGuild
}

// GetPrimaryGuildOk returns a tuple with the PrimaryGuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetPrimaryGuildOk() (*UserPIIResponsePrimaryGuild, bool) {
	if o == nil || IsNil(o.PrimaryGuild) {
		return nil, false
	}
	return o.PrimaryGuild, true
}

// HasPrimaryGuild returns a boolean if a field has been set.
func (o *UserPIIResponse) HasPrimaryGuild() bool {
	if o != nil && !IsNil(o.PrimaryGuild) {
		return true
	}

	return false
}

// SetPrimaryGuild gets a reference to the given UserPIIResponsePrimaryGuild and assigns it to the PrimaryGuild field.
func (o *UserPIIResponse) SetPrimaryGuild(v UserPIIResponsePrimaryGuild) {
	o.PrimaryGuild = &v
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
func (o *UserPIIResponse) GetLocale() AvailableLocalesEnum {
	if o == nil {
		var ret AvailableLocalesEnum
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetLocaleOk() (*AvailableLocalesEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *UserPIIResponse) SetLocale(v AvailableLocalesEnum) {
	o.Locale = v
}

// GetPremiumType returns the PremiumType field value if set, zero value otherwise.
func (o *UserPIIResponse) GetPremiumType() PremiumTypes {
	if o == nil || IsNil(o.PremiumType) {
		var ret PremiumTypes
		return ret
	}
	return *o.PremiumType
}

// GetPremiumTypeOk returns a tuple with the PremiumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetPremiumTypeOk() (*PremiumTypes, bool) {
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

// SetPremiumType gets a reference to the given PremiumTypes and assigns it to the PremiumType field.
func (o *UserPIIResponse) SetPremiumType(v PremiumTypes) {
	o.PremiumType = &v
}


// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserPIIResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserPIIResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserPIIResponse) SetEmail(v string) {
	o.Email = &v
}


// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *UserPIIResponse) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPIIResponse) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *UserPIIResponse) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *UserPIIResponse) SetVerified(v bool) {
	o.Verified = &v
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
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	toSerialize["discriminator"] = o.Discriminator
	toSerialize["public_flags"] = o.PublicFlags
	toSerialize["flags"] = o.Flags
	if !IsNil(o.Bot) {
		toSerialize["bot"] = o.Bot
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	if !IsNil(o.AccentColor) {
		toSerialize["accent_color"] = o.AccentColor
	}
	if !IsNil(o.GlobalName) {
		toSerialize["global_name"] = o.GlobalName
	}
	if !IsNil(o.AvatarDecorationData) {
		toSerialize["avatar_decoration_data"] = o.AvatarDecorationData
	}
	if !IsNil(o.Collectibles) {
		toSerialize["collectibles"] = o.Collectibles
	}
	if !IsNil(o.PrimaryGuild) {
		toSerialize["primary_guild"] = o.PrimaryGuild
	}
	toSerialize["mfa_enabled"] = o.MfaEnabled
	toSerialize["locale"] = o.Locale
	if !IsNil(o.PremiumType) {
		toSerialize["premium_type"] = o.PremiumType
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
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

