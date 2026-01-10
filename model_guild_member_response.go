/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-10T02:53:50.800755799Z[Etc/UTC]
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
	"time"
	"bytes"
	"fmt"
)

// checks if the GuildMemberResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildMemberResponse{}

// GuildMemberResponse struct for GuildMemberResponse
type GuildMemberResponse struct {
	Avatar *string `json:"avatar,omitempty"`
	AvatarDecorationData *BasicGuildMemberResponseAvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	Banner *string `json:"banner,omitempty"`
	CommunicationDisabledUntil *time.Time `json:"communication_disabled_until,omitempty"`
	Flags int32 `json:"flags"`
	JoinedAt time.Time `json:"joined_at"`
	Nick *string `json:"nick,omitempty"`
	Pending bool `json:"pending"`
	PremiumSince *time.Time `json:"premium_since,omitempty"`
	Roles []string `json:"roles"`
	Collectibles *BasicGuildMemberResponseCollectibles `json:"collectibles,omitempty"`
	User UserResponse `json:"user"`
	Mute bool `json:"mute"`
	Deaf bool `json:"deaf"`
}

type _GuildMemberResponse GuildMemberResponse

// NewGuildMemberResponse instantiates a new GuildMemberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildMemberResponse(flags int32, joinedAt time.Time, pending bool, roles []string, user UserResponse, mute bool, deaf bool) *GuildMemberResponse {
	this := GuildMemberResponse{}
	this.Flags = flags
	this.JoinedAt = joinedAt
	this.Pending = pending
	this.Roles = roles
	this.User = user
	this.Mute = mute
	this.Deaf = deaf
	return &this
}

// NewGuildMemberResponseWithDefaults instantiates a new GuildMemberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildMemberResponseWithDefaults() *GuildMemberResponse {
	this := GuildMemberResponse{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *GuildMemberResponse) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *GuildMemberResponse) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *GuildMemberResponse) SetAvatar(v string) {
	o.Avatar = &v
}


// GetAvatarDecorationData returns the AvatarDecorationData field value if set, zero value otherwise.
func (o *GuildMemberResponse) GetAvatarDecorationData() BasicGuildMemberResponseAvatarDecorationData {
	if o == nil || IsNil(o.AvatarDecorationData) {
		var ret BasicGuildMemberResponseAvatarDecorationData
		return ret
	}
	return *o.AvatarDecorationData
}

// GetAvatarDecorationDataOk returns a tuple with the AvatarDecorationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetAvatarDecorationDataOk() (*BasicGuildMemberResponseAvatarDecorationData, bool) {
	if o == nil || IsNil(o.AvatarDecorationData) {
		return nil, false
	}
	return o.AvatarDecorationData, true
}

// HasAvatarDecorationData returns a boolean if a field has been set.
func (o *GuildMemberResponse) HasAvatarDecorationData() bool {
	if o != nil && !IsNil(o.AvatarDecorationData) {
		return true
	}

	return false
}

// SetAvatarDecorationData gets a reference to the given BasicGuildMemberResponseAvatarDecorationData and assigns it to the AvatarDecorationData field.
func (o *GuildMemberResponse) SetAvatarDecorationData(v BasicGuildMemberResponseAvatarDecorationData) {
	o.AvatarDecorationData = &v
}


// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *GuildMemberResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner) {
		var ret string
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *GuildMemberResponse) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given string and assigns it to the Banner field.
func (o *GuildMemberResponse) SetBanner(v string) {
	o.Banner = &v
}


// GetCommunicationDisabledUntil returns the CommunicationDisabledUntil field value if set, zero value otherwise.
func (o *GuildMemberResponse) GetCommunicationDisabledUntil() time.Time {
	if o == nil || IsNil(o.CommunicationDisabledUntil) {
		var ret time.Time
		return ret
	}
	return *o.CommunicationDisabledUntil
}

// GetCommunicationDisabledUntilOk returns a tuple with the CommunicationDisabledUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetCommunicationDisabledUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CommunicationDisabledUntil) {
		return nil, false
	}
	return o.CommunicationDisabledUntil, true
}

// HasCommunicationDisabledUntil returns a boolean if a field has been set.
func (o *GuildMemberResponse) HasCommunicationDisabledUntil() bool {
	if o != nil && !IsNil(o.CommunicationDisabledUntil) {
		return true
	}

	return false
}

// SetCommunicationDisabledUntil gets a reference to the given time.Time and assigns it to the CommunicationDisabledUntil field.
func (o *GuildMemberResponse) SetCommunicationDisabledUntil(v time.Time) {
	o.CommunicationDisabledUntil = &v
}


// GetFlags returns the Flags field value
func (o *GuildMemberResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *GuildMemberResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetJoinedAt returns the JoinedAt field value
func (o *GuildMemberResponse) GetJoinedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetJoinedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinedAt, true
}

// SetJoinedAt sets field value
func (o *GuildMemberResponse) SetJoinedAt(v time.Time) {
	o.JoinedAt = v
}

// GetNick returns the Nick field value if set, zero value otherwise.
func (o *GuildMemberResponse) GetNick() string {
	if o == nil || IsNil(o.Nick) {
		var ret string
		return ret
	}
	return *o.Nick
}

// GetNickOk returns a tuple with the Nick field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetNickOk() (*string, bool) {
	if o == nil || IsNil(o.Nick) {
		return nil, false
	}
	return o.Nick, true
}

// HasNick returns a boolean if a field has been set.
func (o *GuildMemberResponse) HasNick() bool {
	if o != nil && !IsNil(o.Nick) {
		return true
	}

	return false
}

// SetNick gets a reference to the given string and assigns it to the Nick field.
func (o *GuildMemberResponse) SetNick(v string) {
	o.Nick = &v
}


// GetPending returns the Pending field value
func (o *GuildMemberResponse) GetPending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *GuildMemberResponse) SetPending(v bool) {
	o.Pending = v
}

// GetPremiumSince returns the PremiumSince field value if set, zero value otherwise.
func (o *GuildMemberResponse) GetPremiumSince() time.Time {
	if o == nil || IsNil(o.PremiumSince) {
		var ret time.Time
		return ret
	}
	return *o.PremiumSince
}

// GetPremiumSinceOk returns a tuple with the PremiumSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetPremiumSinceOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PremiumSince) {
		return nil, false
	}
	return o.PremiumSince, true
}

// HasPremiumSince returns a boolean if a field has been set.
func (o *GuildMemberResponse) HasPremiumSince() bool {
	if o != nil && !IsNil(o.PremiumSince) {
		return true
	}

	return false
}

// SetPremiumSince gets a reference to the given time.Time and assigns it to the PremiumSince field.
func (o *GuildMemberResponse) SetPremiumSince(v time.Time) {
	o.PremiumSince = &v
}


// GetRoles returns the Roles field value
func (o *GuildMemberResponse) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *GuildMemberResponse) SetRoles(v []string) {
	o.Roles = v
}

// GetCollectibles returns the Collectibles field value if set, zero value otherwise.
func (o *GuildMemberResponse) GetCollectibles() BasicGuildMemberResponseCollectibles {
	if o == nil || IsNil(o.Collectibles) {
		var ret BasicGuildMemberResponseCollectibles
		return ret
	}
	return *o.Collectibles
}

// GetCollectiblesOk returns a tuple with the Collectibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetCollectiblesOk() (*BasicGuildMemberResponseCollectibles, bool) {
	if o == nil || IsNil(o.Collectibles) {
		return nil, false
	}
	return o.Collectibles, true
}

// HasCollectibles returns a boolean if a field has been set.
func (o *GuildMemberResponse) HasCollectibles() bool {
	if o != nil && !IsNil(o.Collectibles) {
		return true
	}

	return false
}

// SetCollectibles gets a reference to the given BasicGuildMemberResponseCollectibles and assigns it to the Collectibles field.
func (o *GuildMemberResponse) SetCollectibles(v BasicGuildMemberResponseCollectibles) {
	o.Collectibles = &v
}


// GetUser returns the User field value
func (o *GuildMemberResponse) GetUser() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GuildMemberResponse) SetUser(v UserResponse) {
	o.User = v
}

// GetMute returns the Mute field value
func (o *GuildMemberResponse) GetMute() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Mute
}

// GetMuteOk returns a tuple with the Mute field value
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetMuteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mute, true
}

// SetMute sets field value
func (o *GuildMemberResponse) SetMute(v bool) {
	o.Mute = v
}

// GetDeaf returns the Deaf field value
func (o *GuildMemberResponse) GetDeaf() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deaf
}

// GetDeafOk returns a tuple with the Deaf field value
// and a boolean to check if the value has been set.
func (o *GuildMemberResponse) GetDeafOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deaf, true
}

// SetDeaf sets field value
func (o *GuildMemberResponse) SetDeaf(v bool) {
	o.Deaf = v
}

func (o GuildMemberResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildMemberResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.AvatarDecorationData) {
		toSerialize["avatar_decoration_data"] = o.AvatarDecorationData
	}
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	if !IsNil(o.CommunicationDisabledUntil) {
		toSerialize["communication_disabled_until"] = o.CommunicationDisabledUntil
	}
	toSerialize["flags"] = o.Flags
	toSerialize["joined_at"] = o.JoinedAt
	if !IsNil(o.Nick) {
		toSerialize["nick"] = o.Nick
	}
	toSerialize["pending"] = o.Pending
	if !IsNil(o.PremiumSince) {
		toSerialize["premium_since"] = o.PremiumSince
	}
	toSerialize["roles"] = o.Roles
	if !IsNil(o.Collectibles) {
		toSerialize["collectibles"] = o.Collectibles
	}
	toSerialize["user"] = o.User
	toSerialize["mute"] = o.Mute
	toSerialize["deaf"] = o.Deaf
	return toSerialize, nil
}

func (o *GuildMemberResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"flags",
		"joined_at",
		"pending",
		"roles",
		"user",
		"mute",
		"deaf",
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

	varGuildMemberResponse := _GuildMemberResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildMemberResponse)

	if err != nil {
		return err
	}

	*o = GuildMemberResponse(varGuildMemberResponse)

	return err
}

type NullableGuildMemberResponse struct {
	value *GuildMemberResponse
	isSet bool
}

func (v NullableGuildMemberResponse) Get() *GuildMemberResponse {
	return v.value
}

func (v *NullableGuildMemberResponse) Set(val *GuildMemberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildMemberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildMemberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildMemberResponse(val *GuildMemberResponse) *NullableGuildMemberResponse {
	return &NullableGuildMemberResponse{value: val, isSet: true}
}

func (v NullableGuildMemberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildMemberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

