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

// checks if the CreateGuildInviteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGuildInviteRequest{}

// CreateGuildInviteRequest struct for CreateGuildInviteRequest
type CreateGuildInviteRequest struct {
	MaxAge NullableInt32 `json:"max_age,omitempty"`
	Temporary NullableBool `json:"temporary,omitempty"`
	MaxUses NullableInt32 `json:"max_uses,omitempty"`
	Unique NullableBool `json:"unique,omitempty"`
	TargetUserId *string `json:"target_user_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	TargetApplicationId *string `json:"target_application_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	TargetType NullableInt32 `json:"target_type,omitempty"`
}

// NewCreateGuildInviteRequest instantiates a new CreateGuildInviteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGuildInviteRequest() *CreateGuildInviteRequest {
	this := CreateGuildInviteRequest{}
	return &this
}

// NewCreateGuildInviteRequestWithDefaults instantiates a new CreateGuildInviteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGuildInviteRequestWithDefaults() *CreateGuildInviteRequest {
	this := CreateGuildInviteRequest{}
	return &this
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildInviteRequest) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAge.Get()
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildInviteRequest) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge.Get()) {
		return nil, false
	}
	return o.MaxAge.Get(), o.MaxAge.IsSet()
}

// HasMaxAge returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasMaxAge() bool {
	if o != nil && o.MaxAge.IsSet() {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given NullableInt32 and assigns it to the MaxAge field.
func (o *CreateGuildInviteRequest) SetMaxAge(v int32) {
	o.MaxAge.Set(&v)
}

// SetMaxAgeNil sets the value for MaxAge to be an explicit nil
func (o *CreateGuildInviteRequest) SetMaxAgeNil() {
	o.MaxAge.Set(nil)
}

// UnsetMaxAge ensures that no value is present for MaxAge, not even an explicit nil
func (o *CreateGuildInviteRequest) UnsetMaxAge() {
	o.MaxAge.Unset()
}

// GetTemporary returns the Temporary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildInviteRequest) GetTemporary() bool {
	if o == nil || IsNil(o.Temporary.Get()) {
		var ret bool
		return ret
	}
	return *o.Temporary.Get()
}

// GetTemporaryOk returns a tuple with the Temporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildInviteRequest) GetTemporaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Temporary.Get()) {
		return nil, false
	}
	return o.Temporary.Get(), o.Temporary.IsSet()
}

// HasTemporary returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasTemporary() bool {
	if o != nil && o.Temporary.IsSet() {
		return true
	}

	return false
}

// SetTemporary gets a reference to the given NullableBool and assigns it to the Temporary field.
func (o *CreateGuildInviteRequest) SetTemporary(v bool) {
	o.Temporary.Set(&v)
}

// SetTemporaryNil sets the value for Temporary to be an explicit nil
func (o *CreateGuildInviteRequest) SetTemporaryNil() {
	o.Temporary.Set(nil)
}

// UnsetTemporary ensures that no value is present for Temporary, not even an explicit nil
func (o *CreateGuildInviteRequest) UnsetTemporary() {
	o.Temporary.Unset()
}

// GetMaxUses returns the MaxUses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildInviteRequest) GetMaxUses() int32 {
	if o == nil || IsNil(o.MaxUses.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxUses.Get()
}

// GetMaxUsesOk returns a tuple with the MaxUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildInviteRequest) GetMaxUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUses.Get()) {
		return nil, false
	}
	return o.MaxUses.Get(), o.MaxUses.IsSet()
}

// HasMaxUses returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasMaxUses() bool {
	if o != nil && o.MaxUses.IsSet() {
		return true
	}

	return false
}

// SetMaxUses gets a reference to the given NullableInt32 and assigns it to the MaxUses field.
func (o *CreateGuildInviteRequest) SetMaxUses(v int32) {
	o.MaxUses.Set(&v)
}

// SetMaxUsesNil sets the value for MaxUses to be an explicit nil
func (o *CreateGuildInviteRequest) SetMaxUsesNil() {
	o.MaxUses.Set(nil)
}

// UnsetMaxUses ensures that no value is present for MaxUses, not even an explicit nil
func (o *CreateGuildInviteRequest) UnsetMaxUses() {
	o.MaxUses.Unset()
}

// GetUnique returns the Unique field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildInviteRequest) GetUnique() bool {
	if o == nil || IsNil(o.Unique.Get()) {
		var ret bool
		return ret
	}
	return *o.Unique.Get()
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildInviteRequest) GetUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.Unique.Get()) {
		return nil, false
	}
	return o.Unique.Get(), o.Unique.IsSet()
}

// HasUnique returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasUnique() bool {
	if o != nil && o.Unique.IsSet() {
		return true
	}

	return false
}

// SetUnique gets a reference to the given NullableBool and assigns it to the Unique field.
func (o *CreateGuildInviteRequest) SetUnique(v bool) {
	o.Unique.Set(&v)
}

// SetUniqueNil sets the value for Unique to be an explicit nil
func (o *CreateGuildInviteRequest) SetUniqueNil() {
	o.Unique.Set(nil)
}

// UnsetUnique ensures that no value is present for Unique, not even an explicit nil
func (o *CreateGuildInviteRequest) UnsetUnique() {
	o.Unique.Unset()
}

// GetTargetUserId returns the TargetUserId field value if set, zero value otherwise.
func (o *CreateGuildInviteRequest) GetTargetUserId() string {
	if o == nil || IsNil(o.TargetUserId) {
		var ret string
		return ret
	}
	return *o.TargetUserId
}

// GetTargetUserIdOk returns a tuple with the TargetUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildInviteRequest) GetTargetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUserId) {
		return nil, false
	}
	return o.TargetUserId, true
}

// HasTargetUserId returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasTargetUserId() bool {
	if o != nil && !IsNil(o.TargetUserId) {
		return true
	}

	return false
}

// SetTargetUserId gets a reference to the given string and assigns it to the TargetUserId field.
func (o *CreateGuildInviteRequest) SetTargetUserId(v string) {
	o.TargetUserId = &v
}


// GetTargetApplicationId returns the TargetApplicationId field value if set, zero value otherwise.
func (o *CreateGuildInviteRequest) GetTargetApplicationId() string {
	if o == nil || IsNil(o.TargetApplicationId) {
		var ret string
		return ret
	}
	return *o.TargetApplicationId
}

// GetTargetApplicationIdOk returns a tuple with the TargetApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildInviteRequest) GetTargetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetApplicationId) {
		return nil, false
	}
	return o.TargetApplicationId, true
}

// HasTargetApplicationId returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasTargetApplicationId() bool {
	if o != nil && !IsNil(o.TargetApplicationId) {
		return true
	}

	return false
}

// SetTargetApplicationId gets a reference to the given string and assigns it to the TargetApplicationId field.
func (o *CreateGuildInviteRequest) SetTargetApplicationId(v string) {
	o.TargetApplicationId = &v
}


// GetTargetType returns the TargetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildInviteRequest) GetTargetType() int32 {
	if o == nil || IsNil(o.TargetType.Get()) {
		var ret int32
		return ret
	}
	return *o.TargetType.Get()
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildInviteRequest) GetTargetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetType.Get()) {
		return nil, false
	}
	return o.TargetType.Get(), o.TargetType.IsSet()
}

// HasTargetType returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasTargetType() bool {
	if o != nil && o.TargetType.IsSet() {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given NullableInt32 and assigns it to the TargetType field.
func (o *CreateGuildInviteRequest) SetTargetType(v int32) {
	o.TargetType.Set(&v)
}

// SetTargetTypeNil sets the value for TargetType to be an explicit nil
func (o *CreateGuildInviteRequest) SetTargetTypeNil() {
	o.TargetType.Set(nil)
}

// UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
func (o *CreateGuildInviteRequest) UnsetTargetType() {
	o.TargetType.Unset()
}

func (o CreateGuildInviteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGuildInviteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxAge.IsSet() {
		toSerialize["max_age"] = o.MaxAge.Get()
	}
	if o.Temporary.IsSet() {
		toSerialize["temporary"] = o.Temporary.Get()
	}
	if o.MaxUses.IsSet() {
		toSerialize["max_uses"] = o.MaxUses.Get()
	}
	if o.Unique.IsSet() {
		toSerialize["unique"] = o.Unique.Get()
	}
	if !IsNil(o.TargetUserId) {
		toSerialize["target_user_id"] = o.TargetUserId
	}
	if !IsNil(o.TargetApplicationId) {
		toSerialize["target_application_id"] = o.TargetApplicationId
	}
	if o.TargetType.IsSet() {
		toSerialize["target_type"] = o.TargetType.Get()
	}
	return toSerialize, nil
}

type NullableCreateGuildInviteRequest struct {
	value *CreateGuildInviteRequest
	isSet bool
}

func (v NullableCreateGuildInviteRequest) Get() *CreateGuildInviteRequest {
	return v.value
}

func (v *NullableCreateGuildInviteRequest) Set(val *CreateGuildInviteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGuildInviteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGuildInviteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGuildInviteRequest(val *CreateGuildInviteRequest) *NullableCreateGuildInviteRequest {
	return &NullableCreateGuildInviteRequest{value: val, isSet: true}
}

func (v NullableCreateGuildInviteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGuildInviteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


