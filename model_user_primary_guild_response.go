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

// checks if the UserPrimaryGuildResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPrimaryGuildResponse{}

// UserPrimaryGuildResponse struct for UserPrimaryGuildResponse
type UserPrimaryGuildResponse struct {
	IdentityGuildId *string `json:"identity_guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	IdentityEnabled NullableBool `json:"identity_enabled,omitempty"`
	Tag NullableString `json:"tag,omitempty"`
	Badge NullableString `json:"badge,omitempty"`
}

// NewUserPrimaryGuildResponse instantiates a new UserPrimaryGuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPrimaryGuildResponse() *UserPrimaryGuildResponse {
	this := UserPrimaryGuildResponse{}
	return &this
}

// NewUserPrimaryGuildResponseWithDefaults instantiates a new UserPrimaryGuildResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPrimaryGuildResponseWithDefaults() *UserPrimaryGuildResponse {
	this := UserPrimaryGuildResponse{}
	return &this
}

// GetIdentityGuildId returns the IdentityGuildId field value if set, zero value otherwise.
func (o *UserPrimaryGuildResponse) GetIdentityGuildId() string {
	if o == nil || IsNil(o.IdentityGuildId) {
		var ret string
		return ret
	}
	return *o.IdentityGuildId
}

// GetIdentityGuildIdOk returns a tuple with the IdentityGuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrimaryGuildResponse) GetIdentityGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityGuildId) {
		return nil, false
	}
	return o.IdentityGuildId, true
}

// HasIdentityGuildId returns a boolean if a field has been set.
func (o *UserPrimaryGuildResponse) HasIdentityGuildId() bool {
	if o != nil && !IsNil(o.IdentityGuildId) {
		return true
	}

	return false
}

// SetIdentityGuildId gets a reference to the given string and assigns it to the IdentityGuildId field.
func (o *UserPrimaryGuildResponse) SetIdentityGuildId(v string) {
	o.IdentityGuildId = &v
}


// GetIdentityEnabled returns the IdentityEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPrimaryGuildResponse) GetIdentityEnabled() bool {
	if o == nil || IsNil(o.IdentityEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IdentityEnabled.Get()
}

// GetIdentityEnabledOk returns a tuple with the IdentityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPrimaryGuildResponse) GetIdentityEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IdentityEnabled.Get()) {
		return nil, false
	}
	return o.IdentityEnabled.Get(), o.IdentityEnabled.IsSet()
}

// HasIdentityEnabled returns a boolean if a field has been set.
func (o *UserPrimaryGuildResponse) HasIdentityEnabled() bool {
	if o != nil && o.IdentityEnabled.IsSet() {
		return true
	}

	return false
}

// SetIdentityEnabled gets a reference to the given NullableBool and assigns it to the IdentityEnabled field.
func (o *UserPrimaryGuildResponse) SetIdentityEnabled(v bool) {
	o.IdentityEnabled.Set(&v)
}

// SetIdentityEnabledNil sets the value for IdentityEnabled to be an explicit nil
func (o *UserPrimaryGuildResponse) SetIdentityEnabledNil() {
	o.IdentityEnabled.Set(nil)
}

// UnsetIdentityEnabled ensures that no value is present for IdentityEnabled, not even an explicit nil
func (o *UserPrimaryGuildResponse) UnsetIdentityEnabled() {
	o.IdentityEnabled.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPrimaryGuildResponse) GetTag() string {
	if o == nil || IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPrimaryGuildResponse) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag.Get()) {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *UserPrimaryGuildResponse) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableString and assigns it to the Tag field.
func (o *UserPrimaryGuildResponse) SetTag(v string) {
	o.Tag.Set(&v)
}

// SetTagNil sets the value for Tag to be an explicit nil
func (o *UserPrimaryGuildResponse) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *UserPrimaryGuildResponse) UnsetTag() {
	o.Tag.Unset()
}

// GetBadge returns the Badge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPrimaryGuildResponse) GetBadge() string {
	if o == nil || IsNil(o.Badge.Get()) {
		var ret string
		return ret
	}
	return *o.Badge.Get()
}

// GetBadgeOk returns a tuple with the Badge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPrimaryGuildResponse) GetBadgeOk() (*string, bool) {
	if o == nil || IsNil(o.Badge.Get()) {
		return nil, false
	}
	return o.Badge.Get(), o.Badge.IsSet()
}

// HasBadge returns a boolean if a field has been set.
func (o *UserPrimaryGuildResponse) HasBadge() bool {
	if o != nil && o.Badge.IsSet() {
		return true
	}

	return false
}

// SetBadge gets a reference to the given NullableString and assigns it to the Badge field.
func (o *UserPrimaryGuildResponse) SetBadge(v string) {
	o.Badge.Set(&v)
}

// SetBadgeNil sets the value for Badge to be an explicit nil
func (o *UserPrimaryGuildResponse) SetBadgeNil() {
	o.Badge.Set(nil)
}

// UnsetBadge ensures that no value is present for Badge, not even an explicit nil
func (o *UserPrimaryGuildResponse) UnsetBadge() {
	o.Badge.Unset()
}

func (o UserPrimaryGuildResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPrimaryGuildResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdentityGuildId) {
		toSerialize["identity_guild_id"] = o.IdentityGuildId
	}
	if o.IdentityEnabled.IsSet() {
		toSerialize["identity_enabled"] = o.IdentityEnabled.Get()
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	if o.Badge.IsSet() {
		toSerialize["badge"] = o.Badge.Get()
	}
	return toSerialize, nil
}

type NullableUserPrimaryGuildResponse struct {
	value *UserPrimaryGuildResponse
	isSet bool
}

func (v NullableUserPrimaryGuildResponse) Get() *UserPrimaryGuildResponse {
	return v.value
}

func (v *NullableUserPrimaryGuildResponse) Set(val *UserPrimaryGuildResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPrimaryGuildResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPrimaryGuildResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPrimaryGuildResponse(val *UserPrimaryGuildResponse) *NullableUserPrimaryGuildResponse {
	return &NullableUserPrimaryGuildResponse{value: val, isSet: true}
}

func (v NullableUserPrimaryGuildResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPrimaryGuildResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


