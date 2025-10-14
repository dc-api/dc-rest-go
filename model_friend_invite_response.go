/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-10-14T14:27:02.905060650Z[Etc/UTC]
 *    * - **Generator Version**: 7.16.0
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
	Type NullableInt32 `json:"type"`
	Code string `json:"code"`
	Inviter *UserResponse `json:"inviter,omitempty"`
	MaxAge *int32 `json:"max_age,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	FriendsCount *int32 `json:"friends_count,omitempty"`
	Channel NullableInviteChannelResponse `json:"channel,omitempty"`
	IsContact *bool `json:"is_contact,omitempty"`
	Uses *int32 `json:"uses,omitempty"`
	MaxUses *int32 `json:"max_uses,omitempty"`
	Flags *int32 `json:"flags,omitempty"`
}

type _FriendInviteResponse FriendInviteResponse

// NewFriendInviteResponse instantiates a new FriendInviteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFriendInviteResponse(type_ NullableInt32, code string) *FriendInviteResponse {
	this := FriendInviteResponse{}
	this.Type = type_
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

// GetType returns the Type field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FriendInviteResponse) GetType() int32 {
	if o == nil || o.Type.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendInviteResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *FriendInviteResponse) SetType(v int32) {
	o.Type.Set(&v)
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

// GetInviter returns the Inviter field value if set, zero value otherwise.
func (o *FriendInviteResponse) GetInviter() UserResponse {
	if o == nil || IsNil(o.Inviter) {
		var ret UserResponse
		return ret
	}
	return *o.Inviter
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FriendInviteResponse) GetInviterOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Inviter) {
		return nil, false
	}
	return o.Inviter, true
}

// HasInviter returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasInviter() bool {
	if o != nil && !IsNil(o.Inviter) {
		return true
	}

	return false
}

// SetInviter gets a reference to the given UserResponse and assigns it to the Inviter field.
func (o *FriendInviteResponse) SetInviter(v UserResponse) {
	o.Inviter = &v
}


// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *FriendInviteResponse) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge) {
		var ret int32
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FriendInviteResponse) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int32 and assigns it to the MaxAge field.
func (o *FriendInviteResponse) SetMaxAge(v int32) {
	o.MaxAge = &v
}


// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FriendInviteResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FriendInviteResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FriendInviteResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
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

// GetFriendsCount returns the FriendsCount field value if set, zero value otherwise.
func (o *FriendInviteResponse) GetFriendsCount() int32 {
	if o == nil || IsNil(o.FriendsCount) {
		var ret int32
		return ret
	}
	return *o.FriendsCount
}

// GetFriendsCountOk returns a tuple with the FriendsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FriendInviteResponse) GetFriendsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FriendsCount) {
		return nil, false
	}
	return o.FriendsCount, true
}

// HasFriendsCount returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasFriendsCount() bool {
	if o != nil && !IsNil(o.FriendsCount) {
		return true
	}

	return false
}

// SetFriendsCount gets a reference to the given int32 and assigns it to the FriendsCount field.
func (o *FriendInviteResponse) SetFriendsCount(v int32) {
	o.FriendsCount = &v
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

// GetIsContact returns the IsContact field value if set, zero value otherwise.
func (o *FriendInviteResponse) GetIsContact() bool {
	if o == nil || IsNil(o.IsContact) {
		var ret bool
		return ret
	}
	return *o.IsContact
}

// GetIsContactOk returns a tuple with the IsContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FriendInviteResponse) GetIsContactOk() (*bool, bool) {
	if o == nil || IsNil(o.IsContact) {
		return nil, false
	}
	return o.IsContact, true
}

// HasIsContact returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasIsContact() bool {
	if o != nil && !IsNil(o.IsContact) {
		return true
	}

	return false
}

// SetIsContact gets a reference to the given bool and assigns it to the IsContact field.
func (o *FriendInviteResponse) SetIsContact(v bool) {
	o.IsContact = &v
}


// GetUses returns the Uses field value if set, zero value otherwise.
func (o *FriendInviteResponse) GetUses() int32 {
	if o == nil || IsNil(o.Uses) {
		var ret int32
		return ret
	}
	return *o.Uses
}

// GetUsesOk returns a tuple with the Uses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FriendInviteResponse) GetUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.Uses) {
		return nil, false
	}
	return o.Uses, true
}

// HasUses returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasUses() bool {
	if o != nil && !IsNil(o.Uses) {
		return true
	}

	return false
}

// SetUses gets a reference to the given int32 and assigns it to the Uses field.
func (o *FriendInviteResponse) SetUses(v int32) {
	o.Uses = &v
}


// GetMaxUses returns the MaxUses field value if set, zero value otherwise.
func (o *FriendInviteResponse) GetMaxUses() int32 {
	if o == nil || IsNil(o.MaxUses) {
		var ret int32
		return ret
	}
	return *o.MaxUses
}

// GetMaxUsesOk returns a tuple with the MaxUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FriendInviteResponse) GetMaxUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUses) {
		return nil, false
	}
	return o.MaxUses, true
}

// HasMaxUses returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasMaxUses() bool {
	if o != nil && !IsNil(o.MaxUses) {
		return true
	}

	return false
}

// SetMaxUses gets a reference to the given int32 and assigns it to the MaxUses field.
func (o *FriendInviteResponse) SetMaxUses(v int32) {
	o.MaxUses = &v
}


// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *FriendInviteResponse) GetFlags() int32 {
	if o == nil || IsNil(o.Flags) {
		var ret int32
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FriendInviteResponse) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *FriendInviteResponse) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int32 and assigns it to the Flags field.
func (o *FriendInviteResponse) SetFlags(v int32) {
	o.Flags = &v
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
	toSerialize["type"] = o.Type.Get()
	toSerialize["code"] = o.Code
	if !IsNil(o.Inviter) {
		toSerialize["inviter"] = o.Inviter
	}
	if !IsNil(o.MaxAge) {
		toSerialize["max_age"] = o.MaxAge
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if !IsNil(o.FriendsCount) {
		toSerialize["friends_count"] = o.FriendsCount
	}
	if o.Channel.IsSet() {
		toSerialize["channel"] = o.Channel.Get()
	}
	if !IsNil(o.IsContact) {
		toSerialize["is_contact"] = o.IsContact
	}
	if !IsNil(o.Uses) {
		toSerialize["uses"] = o.Uses
	}
	if !IsNil(o.MaxUses) {
		toSerialize["max_uses"] = o.MaxUses
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

func (o *FriendInviteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

