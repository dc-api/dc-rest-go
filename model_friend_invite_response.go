/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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
	"time"
	"bytes"
	"fmt"
)

// checks if the FriendInviteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FriendInviteResponse{}

// FriendInviteResponse struct for FriendInviteResponse
type FriendInviteResponse struct {
	Type NullableInt32 `json:"type,omitempty"`
	Code string `json:"code"`
	Inviter NullableUserResponse `json:"inviter,omitempty"`
	MaxAge NullableInt32 `json:"max_age,omitempty"`
	CreatedAt NullableTime `json:"created_at,omitempty"`
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	FriendsCount NullableInt32 `json:"friends_count,omitempty"`
	Channel NullableInviteChannelResponse `json:"channel,omitempty"`
	IsContact NullableBool `json:"is_contact,omitempty"`
	Uses NullableInt32 `json:"uses,omitempty"`
	MaxUses NullableInt32 `json:"max_uses,omitempty"`
	Flags NullableInt32 `json:"flags,omitempty"`
}

type _FriendInviteResponse FriendInviteResponse

// NewFriendInviteResponse instantiates a new FriendInviteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFriendInviteResponse(code string) *FriendInviteResponse {
	this := FriendInviteResponse{}
	this.Code = code
	return &this
}

// NewFriendInviteResponseWithDefaults instantiates a new FriendInviteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFriendInviteResponseWithDefaults() *FriendInviteResponse {
	this := FriendInviteResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *FriendInviteResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *FriendInviteResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *FriendInviteResponse) UnsetType() {
	o.Type.Unset()
}

// GetCode returns the Code field value
func (o *FriendInviteResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *FriendInviteResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *FriendInviteResponse) SetCode(v string) {
	o.Code = v
}

// GetInviter returns the Inviter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetInviter() UserResponse {
	if o == nil || IsNil(o.Inviter.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.Inviter.Get()
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetInviterOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Inviter.Get()) {
		return nil, false
	}
	return o.Inviter.Get(), o.Inviter.IsSet()
}

// HasInviter returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasInviter() bool {
	if o != nil && o.Inviter.IsSet() {
		return true
	}

	return false
}

// SetInviter gets a reference to the given NullableUserResponse and assigns it to the Inviter field.
func (o *FriendInviteResponse) SetInviter(v UserResponse) {
	o.Inviter.Set(&v)
}

// SetInviterNil sets the value for Inviter to be an explicit nil
func (o *FriendInviteResponse) SetInviterNil() {
	o.Inviter.Set(nil)
}

// UnsetInviter ensures that no value is present for Inviter, not even an explicit nil
func (o *FriendInviteResponse) UnsetInviter() {
	o.Inviter.Unset()
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAge.Get()
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge.Get()) {
		return nil, false
	}
	return o.MaxAge.Get(), o.MaxAge.IsSet()
}

// HasMaxAge returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasMaxAge() bool {
	if o != nil && o.MaxAge.IsSet() {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given NullableInt32 and assigns it to the MaxAge field.
func (o *FriendInviteResponse) SetMaxAge(v int32) {
	o.MaxAge.Set(&v)
}

// SetMaxAgeNil sets the value for MaxAge to be an explicit nil
func (o *FriendInviteResponse) SetMaxAgeNil() {
	o.MaxAge.Set(nil)
}

// UnsetMaxAge ensures that no value is present for MaxAge, not even an explicit nil
func (o *FriendInviteResponse) UnsetMaxAge() {
	o.MaxAge.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *FriendInviteResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *FriendInviteResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *FriendInviteResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *FriendInviteResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *FriendInviteResponse) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *FriendInviteResponse) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetFriendsCount returns the FriendsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetFriendsCount() int32 {
	if o == nil || IsNil(o.FriendsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FriendsCount.Get()
}

// GetFriendsCountOk returns a tuple with the FriendsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetFriendsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FriendsCount.Get()) {
		return nil, false
	}
	return o.FriendsCount.Get(), o.FriendsCount.IsSet()
}

// HasFriendsCount returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasFriendsCount() bool {
	if o != nil && o.FriendsCount.IsSet() {
		return true
	}

	return false
}

// SetFriendsCount gets a reference to the given NullableInt32 and assigns it to the FriendsCount field.
func (o *FriendInviteResponse) SetFriendsCount(v int32) {
	o.FriendsCount.Set(&v)
}

// SetFriendsCountNil sets the value for FriendsCount to be an explicit nil
func (o *FriendInviteResponse) SetFriendsCountNil() {
	o.FriendsCount.Set(nil)
}

// UnsetFriendsCount ensures that no value is present for FriendsCount, not even an explicit nil
func (o *FriendInviteResponse) UnsetFriendsCount() {
	o.FriendsCount.Unset()
}

// GetChannel returns the Channel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetChannel() InviteChannelResponse {
	if o == nil || IsNil(o.Channel.Get()) {
		var ret InviteChannelResponse
		return ret
	}
	return *o.Channel.Get()
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetChannelOk() (*InviteChannelResponse, bool) {
	if o == nil || IsNil(o.Channel.Get()) {
		return nil, false
	}
	return o.Channel.Get(), o.Channel.IsSet()
}

// HasChannel returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasChannel() bool {
	if o != nil && o.Channel.IsSet() {
		return true
	}

	return false
}

// SetChannel gets a reference to the given NullableInviteChannelResponse and assigns it to the Channel field.
func (o *FriendInviteResponse) SetChannel(v InviteChannelResponse) {
	o.Channel.Set(&v)
}

// SetChannelNil sets the value for Channel to be an explicit nil
func (o *FriendInviteResponse) SetChannelNil() {
	o.Channel.Set(nil)
}

// UnsetChannel ensures that no value is present for Channel, not even an explicit nil
func (o *FriendInviteResponse) UnsetChannel() {
	o.Channel.Unset()
}

// GetIsContact returns the IsContact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetIsContact() bool {
	if o == nil || IsNil(o.IsContact.Get()) {
		var ret bool
		return ret
	}
	return *o.IsContact.Get()
}

// GetIsContactOk returns a tuple with the IsContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetIsContactOk() (*bool, bool) {
	if o == nil || IsNil(o.IsContact.Get()) {
		return nil, false
	}
	return o.IsContact.Get(), o.IsContact.IsSet()
}

// HasIsContact returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasIsContact() bool {
	if o != nil && o.IsContact.IsSet() {
		return true
	}

	return false
}

// SetIsContact gets a reference to the given NullableBool and assigns it to the IsContact field.
func (o *FriendInviteResponse) SetIsContact(v bool) {
	o.IsContact.Set(&v)
}

// SetIsContactNil sets the value for IsContact to be an explicit nil
func (o *FriendInviteResponse) SetIsContactNil() {
	o.IsContact.Set(nil)
}

// UnsetIsContact ensures that no value is present for IsContact, not even an explicit nil
func (o *FriendInviteResponse) UnsetIsContact() {
	o.IsContact.Unset()
}

// GetUses returns the Uses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetUses() int32 {
	if o == nil || IsNil(o.Uses.Get()) {
		var ret int32
		return ret
	}
	return *o.Uses.Get()
}

// GetUsesOk returns a tuple with the Uses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.Uses.Get()) {
		return nil, false
	}
	return o.Uses.Get(), o.Uses.IsSet()
}

// HasUses returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasUses() bool {
	if o != nil && o.Uses.IsSet() {
		return true
	}

	return false
}

// SetUses gets a reference to the given NullableInt32 and assigns it to the Uses field.
func (o *FriendInviteResponse) SetUses(v int32) {
	o.Uses.Set(&v)
}

// SetUsesNil sets the value for Uses to be an explicit nil
func (o *FriendInviteResponse) SetUsesNil() {
	o.Uses.Set(nil)
}

// UnsetUses ensures that no value is present for Uses, not even an explicit nil
func (o *FriendInviteResponse) UnsetUses() {
	o.Uses.Unset()
}

// GetMaxUses returns the MaxUses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetMaxUses() int32 {
	if o == nil || IsNil(o.MaxUses.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxUses.Get()
}

// GetMaxUsesOk returns a tuple with the MaxUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetMaxUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUses.Get()) {
		return nil, false
	}
	return o.MaxUses.Get(), o.MaxUses.IsSet()
}

// HasMaxUses returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasMaxUses() bool {
	if o != nil && o.MaxUses.IsSet() {
		return true
	}

	return false
}

// SetMaxUses gets a reference to the given NullableInt32 and assigns it to the MaxUses field.
func (o *FriendInviteResponse) SetMaxUses(v int32) {
	o.MaxUses.Set(&v)
}

// SetMaxUsesNil sets the value for MaxUses to be an explicit nil
func (o *FriendInviteResponse) SetMaxUsesNil() {
	o.MaxUses.Set(nil)
}

// UnsetMaxUses ensures that no value is present for MaxUses, not even an explicit nil
func (o *FriendInviteResponse) UnsetMaxUses() {
	o.MaxUses.Unset()
}

// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendInviteResponse) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags.Get()) {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *FriendInviteResponse) SetFlags(v int32) {
	o.Flags.Set(&v)
}

// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *FriendInviteResponse) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *FriendInviteResponse) UnsetFlags() {
	o.Flags.Unset()
}

func (o FriendInviteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FriendInviteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	toSerialize["code"] = o.Code
	if o.Inviter.IsSet() {
		toSerialize["inviter"] = o.Inviter.Get()
	}
	if o.MaxAge.IsSet() {
		toSerialize["max_age"] = o.MaxAge.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.FriendsCount.IsSet() {
		toSerialize["friends_count"] = o.FriendsCount.Get()
	}
	if o.Channel.IsSet() {
		toSerialize["channel"] = o.Channel.Get()
	}
	if o.IsContact.IsSet() {
		toSerialize["is_contact"] = o.IsContact.Get()
	}
	if o.Uses.IsSet() {
		toSerialize["uses"] = o.Uses.Get()
	}
	if o.MaxUses.IsSet() {
		toSerialize["max_uses"] = o.MaxUses.Get()
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	return toSerialize, nil
}

func (o *FriendInviteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varFriendInviteResponse := _FriendInviteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFriendInviteResponse)

	if err != nil {
		return err
	}

	*o = FriendInviteResponse(varFriendInviteResponse)

	return err
}

type NullableFriendInviteResponse struct {
	value *FriendInviteResponse
	isSet bool
}

func (v NullableFriendInviteResponse) Get() *FriendInviteResponse {
	return v.value
}

func (v *NullableFriendInviteResponse) Set(val *FriendInviteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFriendInviteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFriendInviteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFriendInviteResponse(val *FriendInviteResponse) *NullableFriendInviteResponse {
	return &NullableFriendInviteResponse{value: val, isSet: true}
}

func (v NullableFriendInviteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFriendInviteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


