/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-12-21T03:28:05.066589019Z[Etc/UTC]
 *    * - **Generator Version**: 7.17.0
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
	"time"
	"bytes"
	"fmt"
)

// checks if the BasicGuildMemberResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicGuildMemberResponse{}

// BasicGuildMemberResponse struct for BasicGuildMemberResponse
type BasicGuildMemberResponse struct {
	Avatar NullableString `json:"avatar,omitempty"`
	AvatarDecorationData NullableUserAvatarDecorationResponse `json:"avatar_decoration_data,omitempty"`
	Banner NullableString `json:"banner,omitempty"`
	CommunicationDisabledUntil NullableTime `json:"communication_disabled_until,omitempty"`
	Flags int32 `json:"flags"`
	JoinedAt time.Time `json:"joined_at"`
	Nick NullableString `json:"nick,omitempty"`
	Pending bool `json:"pending"`
	PremiumSince NullableTime `json:"premium_since,omitempty"`
	Roles []string `json:"roles"`
	Collectibles NullableUserCollectiblesResponse `json:"collectibles,omitempty"`
}

type _BasicGuildMemberResponse BasicGuildMemberResponse

// NewBasicGuildMemberResponse instantiates a new BasicGuildMemberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicGuildMemberResponse(flags int32, joinedAt time.Time, pending bool, roles []string) *BasicGuildMemberResponse {
	this := BasicGuildMemberResponse{}
	this.Flags = flags
	this.JoinedAt = joinedAt
	this.Pending = pending
	this.Roles = roles
	return &this
}

// NewBasicGuildMemberResponseWithDefaults instantiates a new BasicGuildMemberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicGuildMemberResponseWithDefaults() *BasicGuildMemberResponse {
	this := BasicGuildMemberResponse{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BasicGuildMemberResponse) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicGuildMemberResponse) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar.Get()) {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *BasicGuildMemberResponse) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *BasicGuildMemberResponse) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *BasicGuildMemberResponse) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *BasicGuildMemberResponse) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetAvatarDecorationData returns the AvatarDecorationData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BasicGuildMemberResponse) GetAvatarDecorationData() UserAvatarDecorationResponse {
	if o == nil || IsNil(o.AvatarDecorationData.Get()) {
		var ret UserAvatarDecorationResponse
		return ret
	}
	return *o.AvatarDecorationData.Get()
}

// GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicGuildMemberResponse) GetAvatarDecorationDataOk() (*UserAvatarDecorationResponse, bool) {
	if o == nil || IsNil(o.AvatarDecorationData.Get()) {
		return nil, false
	}
	return o.AvatarDecorationData.Get(), o.AvatarDecorationData.IsSet()
}

// HasAvatarDecorationData returns a boolean if a field has been set.
func (o *BasicGuildMemberResponse) HasAvatarDecorationData() bool {
	if o != nil && o.AvatarDecorationData.IsSet() {
		return true
	}

	return false
}

// SetAvatarDecorationData gets a reference to the given NullableUserAvatarDecorationResponse and assigns it to the AvatarDecorationData field.
func (o *BasicGuildMemberResponse) SetAvatarDecorationData(v UserAvatarDecorationResponse) {
	o.AvatarDecorationData.Set(&v)
}

// SetAvatarDecorationDataNil sets the value for AvatarDecorationData to be an explicit nil
func (o *BasicGuildMemberResponse) SetAvatarDecorationDataNil() {
	o.AvatarDecorationData.Set(nil)
}

// UnsetAvatarDecorationData ensures that no value is present for AvatarDecorationData, not even an explicit nil
func (o *BasicGuildMemberResponse) UnsetAvatarDecorationData() {
	o.AvatarDecorationData.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BasicGuildMemberResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner.Get()) {
		var ret string
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicGuildMemberResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner.Get()) {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *BasicGuildMemberResponse) HasBanner() bool {
	if o != nil && o.Banner.IsSet() {
		return true
	}

	return false
}

// SetBanner gets a reference to the given NullableString and assigns it to the Banner field.
func (o *BasicGuildMemberResponse) SetBanner(v string) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil
func (o *BasicGuildMemberResponse) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil
func (o *BasicGuildMemberResponse) UnsetBanner() {
	o.Banner.Unset()
}

// GetCommunicationDisabledUntil returns the CommunicationDisabledUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BasicGuildMemberResponse) GetCommunicationDisabledUntil() time.Time {
	if o == nil || IsNil(o.CommunicationDisabledUntil.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CommunicationDisabledUntil.Get()
}

// GetCommunicationDisabledUntilOk returns a tuple with the CommunicationDisabledUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicGuildMemberResponse) GetCommunicationDisabledUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CommunicationDisabledUntil.Get()) {
		return nil, false
	}
	return o.CommunicationDisabledUntil.Get(), o.CommunicationDisabledUntil.IsSet()
}

// HasCommunicationDisabledUntil returns a boolean if a field has been set.
func (o *BasicGuildMemberResponse) HasCommunicationDisabledUntil() bool {
	if o != nil && o.CommunicationDisabledUntil.IsSet() {
		return true
	}

	return false
}

// SetCommunicationDisabledUntil gets a reference to the given NullableTime and assigns it to the CommunicationDisabledUntil field.
func (o *BasicGuildMemberResponse) SetCommunicationDisabledUntil(v time.Time) {
	o.CommunicationDisabledUntil.Set(&v)
}

// SetCommunicationDisabledUntilNil sets the value for CommunicationDisabledUntil to be an explicit nil
func (o *BasicGuildMemberResponse) SetCommunicationDisabledUntilNil() {
	o.CommunicationDisabledUntil.Set(nil)
}

// UnsetCommunicationDisabledUntil ensures that no value is present for CommunicationDisabledUntil, not even an explicit nil
func (o *BasicGuildMemberResponse) UnsetCommunicationDisabledUntil() {
	o.CommunicationDisabledUntil.Unset()
}

// GetFlags returns the Flags field value
func (o *BasicGuildMemberResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *BasicGuildMemberResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *BasicGuildMemberResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetJoinedAt returns the JoinedAt field value
func (o *BasicGuildMemberResponse) GetJoinedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value
// and a boolean to check if the value has been set.
func (o *BasicGuildMemberResponse) GetJoinedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinedAt, true
}

// SetJoinedAt sets field value
func (o *BasicGuildMemberResponse) SetJoinedAt(v time.Time) {
	o.JoinedAt = v
}

// GetNick returns the Nick field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BasicGuildMemberResponse) GetNick() string {
	if o == nil || IsNil(o.Nick.Get()) {
		var ret string
		return ret
	}
	return *o.Nick.Get()
}

// GetNickOk returns a tuple with the Nick field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicGuildMemberResponse) GetNickOk() (*string, bool) {
	if o == nil || IsNil(o.Nick.Get()) {
		return nil, false
	}
	return o.Nick.Get(), o.Nick.IsSet()
}

// HasNick returns a boolean if a field has been set.
func (o *BasicGuildMemberResponse) HasNick() bool {
	if o != nil && o.Nick.IsSet() {
		return true
	}

	return false
}

// SetNick gets a reference to the given NullableString and assigns it to the Nick field.
func (o *BasicGuildMemberResponse) SetNick(v string) {
	o.Nick.Set(&v)
}

// SetNickNil sets the value for Nick to be an explicit nil
func (o *BasicGuildMemberResponse) SetNickNil() {
	o.Nick.Set(nil)
}

// UnsetNick ensures that no value is present for Nick, not even an explicit nil
func (o *BasicGuildMemberResponse) UnsetNick() {
	o.Nick.Unset()
}

// GetPending returns the Pending field value
func (o *BasicGuildMemberResponse) GetPending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *BasicGuildMemberResponse) GetPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *BasicGuildMemberResponse) SetPending(v bool) {
	o.Pending = v
}

// GetPremiumSince returns the PremiumSince field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BasicGuildMemberResponse) GetPremiumSince() time.Time {
	if o == nil || IsNil(o.PremiumSince.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PremiumSince.Get()
}

// GetPremiumSinceOk returns a tuple with the PremiumSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicGuildMemberResponse) GetPremiumSinceOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PremiumSince.Get()) {
		return nil, false
	}
	return o.PremiumSince.Get(), o.PremiumSince.IsSet()
}

// HasPremiumSince returns a boolean if a field has been set.
func (o *BasicGuildMemberResponse) HasPremiumSince() bool {
	if o != nil && o.PremiumSince.IsSet() {
		return true
	}

	return false
}

// SetPremiumSince gets a reference to the given NullableTime and assigns it to the PremiumSince field.
func (o *BasicGuildMemberResponse) SetPremiumSince(v time.Time) {
	o.PremiumSince.Set(&v)
}

// SetPremiumSinceNil sets the value for PremiumSince to be an explicit nil
func (o *BasicGuildMemberResponse) SetPremiumSinceNil() {
	o.PremiumSince.Set(nil)
}

// UnsetPremiumSince ensures that no value is present for PremiumSince, not even an explicit nil
func (o *BasicGuildMemberResponse) UnsetPremiumSince() {
	o.PremiumSince.Unset()
}

// GetRoles returns the Roles field value
func (o *BasicGuildMemberResponse) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *BasicGuildMemberResponse) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *BasicGuildMemberResponse) SetRoles(v []string) {
	o.Roles = v
}

// GetCollectibles returns the Collectibles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BasicGuildMemberResponse) GetCollectibles() UserCollectiblesResponse {
	if o == nil || IsNil(o.Collectibles.Get()) {
		var ret UserCollectiblesResponse
		return ret
	}
	return *o.Collectibles.Get()
}

// GetCollectiblesOk returns a tuple with the Collectibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicGuildMemberResponse) GetCollectiblesOk() (*UserCollectiblesResponse, bool) {
	if o == nil || IsNil(o.Collectibles.Get()) {
		return nil, false
	}
	return o.Collectibles.Get(), o.Collectibles.IsSet()
}

// HasCollectibles returns a boolean if a field has been set.
func (o *BasicGuildMemberResponse) HasCollectibles() bool {
	if o != nil && o.Collectibles.IsSet() {
		return true
	}

	return false
}

// SetCollectibles gets a reference to the given NullableUserCollectiblesResponse and assigns it to the Collectibles field.
func (o *BasicGuildMemberResponse) SetCollectibles(v UserCollectiblesResponse) {
	o.Collectibles.Set(&v)
}

// SetCollectiblesNil sets the value for Collectibles to be an explicit nil
func (o *BasicGuildMemberResponse) SetCollectiblesNil() {
	o.Collectibles.Set(nil)
}

// UnsetCollectibles ensures that no value is present for Collectibles, not even an explicit nil
func (o *BasicGuildMemberResponse) UnsetCollectibles() {
	o.Collectibles.Unset()
}

func (o BasicGuildMemberResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicGuildMemberResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.AvatarDecorationData.IsSet() {
		toSerialize["avatar_decoration_data"] = o.AvatarDecorationData.Get()
	}
	if o.Banner.IsSet() {
		toSerialize["banner"] = o.Banner.Get()
	}
	if o.CommunicationDisabledUntil.IsSet() {
		toSerialize["communication_disabled_until"] = o.CommunicationDisabledUntil.Get()
	}
	toSerialize["flags"] = o.Flags
	toSerialize["joined_at"] = o.JoinedAt
	if o.Nick.IsSet() {
		toSerialize["nick"] = o.Nick.Get()
	}
	toSerialize["pending"] = o.Pending
	if o.PremiumSince.IsSet() {
		toSerialize["premium_since"] = o.PremiumSince.Get()
	}
	toSerialize["roles"] = o.Roles
	if o.Collectibles.IsSet() {
		toSerialize["collectibles"] = o.Collectibles.Get()
	}
	return toSerialize, nil
}

func (o *BasicGuildMemberResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"flags",
		"joined_at",
		"pending",
		"roles",
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

	varBasicGuildMemberResponse := _BasicGuildMemberResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBasicGuildMemberResponse)

	if err != nil {
		return err
	}

	*o = BasicGuildMemberResponse(varBasicGuildMemberResponse)

	return err
}

type NullableBasicGuildMemberResponse struct {
	value *BasicGuildMemberResponse
	isSet bool
}

func (v NullableBasicGuildMemberResponse) Get() *BasicGuildMemberResponse {
	return v.value
}

func (v *NullableBasicGuildMemberResponse) Set(val *BasicGuildMemberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicGuildMemberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicGuildMemberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicGuildMemberResponse(val *BasicGuildMemberResponse) *NullableBasicGuildMemberResponse {
	return &NullableBasicGuildMemberResponse{value: val, isSet: true}
}

func (v NullableBasicGuildMemberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicGuildMemberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

