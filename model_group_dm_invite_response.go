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
	"time"
	"bytes"
	"fmt"
)

// checks if the GroupDMInviteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupDMInviteResponse{}

// GroupDMInviteResponse struct for GroupDMInviteResponse
type GroupDMInviteResponse struct {
	Type NullableInt32 `json:"type,omitempty"`
	Code string `json:"code"`
	Inviter NullableUserResponse `json:"inviter,omitempty"`
	MaxAge NullableInt32 `json:"max_age,omitempty"`
	CreatedAt NullableTime `json:"created_at,omitempty"`
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	Channel NullableInviteChannelResponse `json:"channel,omitempty"`
	ApproximateMemberCount NullableInt32 `json:"approximate_member_count,omitempty"`
}

type _GroupDMInviteResponse GroupDMInviteResponse

// NewGroupDMInviteResponse instantiates a new GroupDMInviteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDMInviteResponse(code string) *GroupDMInviteResponse {
	this := GroupDMInviteResponse{}
	this.Code = code
	return &this
}

// NewGroupDMInviteResponseWithDefaults instantiates a new GroupDMInviteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDMInviteResponseWithDefaults() *GroupDMInviteResponse {
	this := GroupDMInviteResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDMInviteResponse) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDMInviteResponse) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *GroupDMInviteResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *GroupDMInviteResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *GroupDMInviteResponse) UnsetType() {
	o.Type.Unset()
}

// GetCode returns the Code field value
func (o *GroupDMInviteResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GroupDMInviteResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GroupDMInviteResponse) SetCode(v string) {
	o.Code = v
}

// GetInviter returns the Inviter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDMInviteResponse) GetInviter() UserResponse {
	if o == nil || IsNil(o.Inviter.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.Inviter.Get()
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDMInviteResponse) GetInviterOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Inviter.Get()) {
		return nil, false
	}
	return o.Inviter.Get(), o.Inviter.IsSet()
}

// HasInviter returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasInviter() bool {
	if o != nil && o.Inviter.IsSet() {
		return true
	}

	return false
}

// SetInviter gets a reference to the given NullableUserResponse and assigns it to the Inviter field.
func (o *GroupDMInviteResponse) SetInviter(v UserResponse) {
	o.Inviter.Set(&v)
}

// SetInviterNil sets the value for Inviter to be an explicit nil
func (o *GroupDMInviteResponse) SetInviterNil() {
	o.Inviter.Set(nil)
}

// UnsetInviter ensures that no value is present for Inviter, not even an explicit nil
func (o *GroupDMInviteResponse) UnsetInviter() {
	o.Inviter.Unset()
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDMInviteResponse) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAge.Get()
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDMInviteResponse) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge.Get()) {
		return nil, false
	}
	return o.MaxAge.Get(), o.MaxAge.IsSet()
}

// HasMaxAge returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasMaxAge() bool {
	if o != nil && o.MaxAge.IsSet() {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given NullableInt32 and assigns it to the MaxAge field.
func (o *GroupDMInviteResponse) SetMaxAge(v int32) {
	o.MaxAge.Set(&v)
}

// SetMaxAgeNil sets the value for MaxAge to be an explicit nil
func (o *GroupDMInviteResponse) SetMaxAgeNil() {
	o.MaxAge.Set(nil)
}

// UnsetMaxAge ensures that no value is present for MaxAge, not even an explicit nil
func (o *GroupDMInviteResponse) UnsetMaxAge() {
	o.MaxAge.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDMInviteResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDMInviteResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *GroupDMInviteResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GroupDMInviteResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GroupDMInviteResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDMInviteResponse) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDMInviteResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *GroupDMInviteResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *GroupDMInviteResponse) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *GroupDMInviteResponse) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetChannel returns the Channel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDMInviteResponse) GetChannel() InviteChannelResponse {
	if o == nil || IsNil(o.Channel.Get()) {
		var ret InviteChannelResponse
		return ret
	}
	return *o.Channel.Get()
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDMInviteResponse) GetChannelOk() (*InviteChannelResponse, bool) {
	if o == nil || IsNil(o.Channel.Get()) {
		return nil, false
	}
	return o.Channel.Get(), o.Channel.IsSet()
}

// HasChannel returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasChannel() bool {
	if o != nil && o.Channel.IsSet() {
		return true
	}

	return false
}

// SetChannel gets a reference to the given NullableInviteChannelResponse and assigns it to the Channel field.
func (o *GroupDMInviteResponse) SetChannel(v InviteChannelResponse) {
	o.Channel.Set(&v)
}

// SetChannelNil sets the value for Channel to be an explicit nil
func (o *GroupDMInviteResponse) SetChannelNil() {
	o.Channel.Set(nil)
}

// UnsetChannel ensures that no value is present for Channel, not even an explicit nil
func (o *GroupDMInviteResponse) UnsetChannel() {
	o.Channel.Unset()
}

// GetApproximateMemberCount returns the ApproximateMemberCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDMInviteResponse) GetApproximateMemberCount() int32 {
	if o == nil || IsNil(o.ApproximateMemberCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ApproximateMemberCount.Get()
}

// GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDMInviteResponse) GetApproximateMemberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproximateMemberCount.Get()) {
		return nil, false
	}
	return o.ApproximateMemberCount.Get(), o.ApproximateMemberCount.IsSet()
}

// HasApproximateMemberCount returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasApproximateMemberCount() bool {
	if o != nil && o.ApproximateMemberCount.IsSet() {
		return true
	}

	return false
}

// SetApproximateMemberCount gets a reference to the given NullableInt32 and assigns it to the ApproximateMemberCount field.
func (o *GroupDMInviteResponse) SetApproximateMemberCount(v int32) {
	o.ApproximateMemberCount.Set(&v)
}

// SetApproximateMemberCountNil sets the value for ApproximateMemberCount to be an explicit nil
func (o *GroupDMInviteResponse) SetApproximateMemberCountNil() {
	o.ApproximateMemberCount.Set(nil)
}

// UnsetApproximateMemberCount ensures that no value is present for ApproximateMemberCount, not even an explicit nil
func (o *GroupDMInviteResponse) UnsetApproximateMemberCount() {
	o.ApproximateMemberCount.Unset()
}

func (o GroupDMInviteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupDMInviteResponse) ToMap() (map[string]interface{}, error) {
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
	if o.Channel.IsSet() {
		toSerialize["channel"] = o.Channel.Get()
	}
	if o.ApproximateMemberCount.IsSet() {
		toSerialize["approximate_member_count"] = o.ApproximateMemberCount.Get()
	}
	return toSerialize, nil
}

func (o *GroupDMInviteResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGroupDMInviteResponse := _GroupDMInviteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupDMInviteResponse)

	if err != nil {
		return err
	}

	*o = GroupDMInviteResponse(varGroupDMInviteResponse)

	return err
}

type NullableGroupDMInviteResponse struct {
	value *GroupDMInviteResponse
	isSet bool
}

func (v NullableGroupDMInviteResponse) Get() *GroupDMInviteResponse {
	return v.value
}

func (v *NullableGroupDMInviteResponse) Set(val *GroupDMInviteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupDMInviteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupDMInviteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupDMInviteResponse(val *GroupDMInviteResponse) *NullableGroupDMInviteResponse {
	return &NullableGroupDMInviteResponse{value: val, isSet: true}
}

func (v NullableGroupDMInviteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupDMInviteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


