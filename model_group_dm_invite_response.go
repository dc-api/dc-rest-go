/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-28T06:15:04.313783603Z[Etc/UTC]
 *    * - **Generator Version**: 7.19.0
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
	Type NullableInt32 `json:"type"`
	Code string `json:"code"`
	Inviter *UserResponse `json:"inviter,omitempty"`
	MaxAge *int32 `json:"max_age,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	Channel InviteChannelResponse `json:"channel"`
	ApproximateMemberCount NullableInt32 `json:"approximate_member_count,omitempty"`
}

type _GroupDMInviteResponse GroupDMInviteResponse

// NewGroupDMInviteResponse instantiates a new GroupDMInviteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDMInviteResponse(type_ NullableInt32, code string, channel InviteChannelResponse) *GroupDMInviteResponse {
	this := GroupDMInviteResponse{}
	this.Type = type_
	this.Code = code
	this.Channel = channel
	return &this
}

// NewGroupDMInviteResponseWithDefaults instantiates a new GroupDMInviteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDMInviteResponseWithDefaults() *GroupDMInviteResponse {
	this := GroupDMInviteResponse{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *GroupDMInviteResponse) GetType() int32 {
	if o == nil || o.Type.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDMInviteResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *GroupDMInviteResponse) SetType(v int32) {
	o.Type.Set(&v)
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

// GetInviter returns the Inviter field value if set, zero value otherwise.
func (o *GroupDMInviteResponse) GetInviter() UserResponse {
	if o == nil || IsNil(o.Inviter) {
		var ret UserResponse
		return ret
	}
	return *o.Inviter
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDMInviteResponse) GetInviterOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Inviter) {
		return nil, false
	}
	return o.Inviter, true
}

// HasInviter returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasInviter() bool {
	if o != nil && !IsNil(o.Inviter) {
		return true
	}

	return false
}

// SetInviter gets a reference to the given UserResponse and assigns it to the Inviter field.
func (o *GroupDMInviteResponse) SetInviter(v UserResponse) {
	o.Inviter = &v
}


// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *GroupDMInviteResponse) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge) {
		var ret int32
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDMInviteResponse) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int32 and assigns it to the MaxAge field.
func (o *GroupDMInviteResponse) SetMaxAge(v int32) {
	o.MaxAge = &v
}


// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GroupDMInviteResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDMInviteResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupDMInviteResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GroupDMInviteResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
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

// GetChannel returns the Channel field value
func (o *GroupDMInviteResponse) GetChannel() InviteChannelResponse {
	if o == nil {
		var ret InviteChannelResponse
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *GroupDMInviteResponse) GetChannelOk() (*InviteChannelResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *GroupDMInviteResponse) SetChannel(v InviteChannelResponse) {
	o.Channel = v
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
	toSerialize["channel"] = o.Channel
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
		"type",
		"code",
		"channel",
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

