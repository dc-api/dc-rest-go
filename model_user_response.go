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

// checks if the UserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponse{}

// UserResponse struct for UserResponse
type UserResponse struct {
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
}

type _UserResponse UserResponse

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse(id string, username string, discriminator string, publicFlags int32, flags int64) *UserResponse {
	this := UserResponse{}
	this.Id = id
	this.Username = username
	this.Discriminator = discriminator
	this.PublicFlags = publicFlags
	this.Flags = flags
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	return &this
}

// GetId returns the Id field value
func (o *UserResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserResponse) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *UserResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserResponse) SetUsername(v string) {
	o.Username = v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar.Get()) {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *UserResponse) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *UserResponse) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *UserResponse) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *UserResponse) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetDiscriminator returns the Discriminator field value
func (o *UserResponse) GetDiscriminator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Discriminator
}

// GetDiscriminatorOk returns a tuple with the Discriminator field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetDiscriminatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discriminator, true
}

// SetDiscriminator sets field value
func (o *UserResponse) SetDiscriminator(v string) {
	o.Discriminator = v
}

// GetPublicFlags returns the PublicFlags field value
func (o *UserResponse) GetPublicFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PublicFlags
}

// GetPublicFlagsOk returns a tuple with the PublicFlags field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetPublicFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicFlags, true
}

// SetPublicFlags sets field value
func (o *UserResponse) SetPublicFlags(v int32) {
	o.PublicFlags = v
}

// GetFlags returns the Flags field value
func (o *UserResponse) GetFlags() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFlagsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *UserResponse) SetFlags(v int64) {
	o.Flags = v
}

// GetBot returns the Bot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetBot() bool {
	if o == nil || IsNil(o.Bot.Get()) {
		var ret bool
		return ret
	}
	return *o.Bot.Get()
}

// GetBotOk returns a tuple with the Bot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetBotOk() (*bool, bool) {
	if o == nil || IsNil(o.Bot.Get()) {
		return nil, false
	}
	return o.Bot.Get(), o.Bot.IsSet()
}

// HasBot returns a boolean if a field has been set.
func (o *UserResponse) HasBot() bool {
	if o != nil && o.Bot.IsSet() {
		return true
	}

	return false
}

// SetBot gets a reference to the given NullableBool and assigns it to the Bot field.
func (o *UserResponse) SetBot(v bool) {
	o.Bot.Set(&v)
}

// SetBotNil sets the value for Bot to be an explicit nil
func (o *UserResponse) SetBotNil() {
	o.Bot.Set(nil)
}

// UnsetBot ensures that no value is present for Bot, not even an explicit nil
func (o *UserResponse) UnsetBot() {
	o.Bot.Unset()
}

// GetSystem returns the System field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetSystem() bool {
	if o == nil || IsNil(o.System.Get()) {
		var ret bool
		return ret
	}
	return *o.System.Get()
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.System.Get()) {
		return nil, false
	}
	return o.System.Get(), o.System.IsSet()
}

// HasSystem returns a boolean if a field has been set.
func (o *UserResponse) HasSystem() bool {
	if o != nil && o.System.IsSet() {
		return true
	}

	return false
}

// SetSystem gets a reference to the given NullableBool and assigns it to the System field.
func (o *UserResponse) SetSystem(v bool) {
	o.System.Set(&v)
}

// SetSystemNil sets the value for System to be an explicit nil
func (o *UserResponse) SetSystemNil() {
	o.System.Set(nil)
}

// UnsetSystem ensures that no value is present for System, not even an explicit nil
func (o *UserResponse) UnsetSystem() {
	o.System.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner.Get()) {
		var ret string
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner.Get()) {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *UserResponse) HasBanner() bool {
	if o != nil && o.Banner.IsSet() {
		return true
	}

	return false
}

// SetBanner gets a reference to the given NullableString and assigns it to the Banner field.
func (o *UserResponse) SetBanner(v string) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil
func (o *UserResponse) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil
func (o *UserResponse) UnsetBanner() {
	o.Banner.Unset()
}

// GetAccentColor returns the AccentColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetAccentColor() int32 {
	if o == nil || IsNil(o.AccentColor.Get()) {
		var ret int32
		return ret
	}
	return *o.AccentColor.Get()
}

// GetAccentColorOk returns a tuple with the AccentColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetAccentColorOk() (*int32, bool) {
	if o == nil || IsNil(o.AccentColor.Get()) {
		return nil, false
	}
	return o.AccentColor.Get(), o.AccentColor.IsSet()
}

// HasAccentColor returns a boolean if a field has been set.
func (o *UserResponse) HasAccentColor() bool {
	if o != nil && o.AccentColor.IsSet() {
		return true
	}

	return false
}

// SetAccentColor gets a reference to the given NullableInt32 and assigns it to the AccentColor field.
func (o *UserResponse) SetAccentColor(v int32) {
	o.AccentColor.Set(&v)
}

// SetAccentColorNil sets the value for AccentColor to be an explicit nil
func (o *UserResponse) SetAccentColorNil() {
	o.AccentColor.Set(nil)
}

// UnsetAccentColor ensures that no value is present for AccentColor, not even an explicit nil
func (o *UserResponse) UnsetAccentColor() {
	o.AccentColor.Unset()
}

// GetGlobalName returns the GlobalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetGlobalName() string {
	if o == nil || IsNil(o.GlobalName.Get()) {
		var ret string
		return ret
	}
	return *o.GlobalName.Get()
}

// GetGlobalNameOk returns a tuple with the GlobalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetGlobalNameOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalName.Get()) {
		return nil, false
	}
	return o.GlobalName.Get(), o.GlobalName.IsSet()
}

// HasGlobalName returns a boolean if a field has been set.
func (o *UserResponse) HasGlobalName() bool {
	if o != nil && o.GlobalName.IsSet() {
		return true
	}

	return false
}

// SetGlobalName gets a reference to the given NullableString and assigns it to the GlobalName field.
func (o *UserResponse) SetGlobalName(v string) {
	o.GlobalName.Set(&v)
}

// SetGlobalNameNil sets the value for GlobalName to be an explicit nil
func (o *UserResponse) SetGlobalNameNil() {
	o.GlobalName.Set(nil)
}

// UnsetGlobalName ensures that no value is present for GlobalName, not even an explicit nil
func (o *UserResponse) UnsetGlobalName() {
	o.GlobalName.Unset()
}

// GetAvatarDecorationData returns the AvatarDecorationData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetAvatarDecorationData() UserAvatarDecorationResponse {
	if o == nil || IsNil(o.AvatarDecorationData.Get()) {
		var ret UserAvatarDecorationResponse
		return ret
	}
	return *o.AvatarDecorationData.Get()
}

// GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetAvatarDecorationDataOk() (*UserAvatarDecorationResponse, bool) {
	if o == nil || IsNil(o.AvatarDecorationData.Get()) {
		return nil, false
	}
	return o.AvatarDecorationData.Get(), o.AvatarDecorationData.IsSet()
}

// HasAvatarDecorationData returns a boolean if a field has been set.
func (o *UserResponse) HasAvatarDecorationData() bool {
	if o != nil && o.AvatarDecorationData.IsSet() {
		return true
	}

	return false
}

// SetAvatarDecorationData gets a reference to the given NullableUserAvatarDecorationResponse and assigns it to the AvatarDecorationData field.
func (o *UserResponse) SetAvatarDecorationData(v UserAvatarDecorationResponse) {
	o.AvatarDecorationData.Set(&v)
}

// SetAvatarDecorationDataNil sets the value for AvatarDecorationData to be an explicit nil
func (o *UserResponse) SetAvatarDecorationDataNil() {
	o.AvatarDecorationData.Set(nil)
}

// UnsetAvatarDecorationData ensures that no value is present for AvatarDecorationData, not even an explicit nil
func (o *UserResponse) UnsetAvatarDecorationData() {
	o.AvatarDecorationData.Unset()
}

// GetCollectibles returns the Collectibles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetCollectibles() UserCollectiblesResponse {
	if o == nil || IsNil(o.Collectibles.Get()) {
		var ret UserCollectiblesResponse
		return ret
	}
	return *o.Collectibles.Get()
}

// GetCollectiblesOk returns a tuple with the Collectibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetCollectiblesOk() (*UserCollectiblesResponse, bool) {
	if o == nil || IsNil(o.Collectibles.Get()) {
		return nil, false
	}
	return o.Collectibles.Get(), o.Collectibles.IsSet()
}

// HasCollectibles returns a boolean if a field has been set.
func (o *UserResponse) HasCollectibles() bool {
	if o != nil && o.Collectibles.IsSet() {
		return true
	}

	return false
}

// SetCollectibles gets a reference to the given NullableUserCollectiblesResponse and assigns it to the Collectibles field.
func (o *UserResponse) SetCollectibles(v UserCollectiblesResponse) {
	o.Collectibles.Set(&v)
}

// SetCollectiblesNil sets the value for Collectibles to be an explicit nil
func (o *UserResponse) SetCollectiblesNil() {
	o.Collectibles.Set(nil)
}

// UnsetCollectibles ensures that no value is present for Collectibles, not even an explicit nil
func (o *UserResponse) UnsetCollectibles() {
	o.Collectibles.Unset()
}

// GetPrimaryGuild returns the PrimaryGuild field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetPrimaryGuild() UserPrimaryGuildResponse {
	if o == nil || IsNil(o.PrimaryGuild.Get()) {
		var ret UserPrimaryGuildResponse
		return ret
	}
	return *o.PrimaryGuild.Get()
}

// GetPrimaryGuildOk returns a tuple with the PrimaryGuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetPrimaryGuildOk() (*UserPrimaryGuildResponse, bool) {
	if o == nil || IsNil(o.PrimaryGuild.Get()) {
		return nil, false
	}
	return o.PrimaryGuild.Get(), o.PrimaryGuild.IsSet()
}

// HasPrimaryGuild returns a boolean if a field has been set.
func (o *UserResponse) HasPrimaryGuild() bool {
	if o != nil && o.PrimaryGuild.IsSet() {
		return true
	}

	return false
}

// SetPrimaryGuild gets a reference to the given NullableUserPrimaryGuildResponse and assigns it to the PrimaryGuild field.
func (o *UserResponse) SetPrimaryGuild(v UserPrimaryGuildResponse) {
	o.PrimaryGuild.Set(&v)
}

// SetPrimaryGuildNil sets the value for PrimaryGuild to be an explicit nil
func (o *UserResponse) SetPrimaryGuildNil() {
	o.PrimaryGuild.Set(nil)
}

// UnsetPrimaryGuild ensures that no value is present for PrimaryGuild, not even an explicit nil
func (o *UserResponse) UnsetPrimaryGuild() {
	o.PrimaryGuild.Unset()
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResponse) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *UserResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"discriminator",
		"public_flags",
		"flags",
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

	varUserResponse := _UserResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserResponse)

	if err != nil {
		return err
	}

	*o = UserResponse(varUserResponse)

	return err
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


