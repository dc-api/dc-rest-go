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
)

// checks if the CreateGuildInviteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGuildInviteRequest{}

// CreateGuildInviteRequest struct for CreateGuildInviteRequest
type CreateGuildInviteRequest struct {
	MaxAge *int32 `json:"max_age,omitempty"`
	Temporary *bool `json:"temporary,omitempty"`
	MaxUses *int32 `json:"max_uses,omitempty"`
	Unique *bool `json:"unique,omitempty"`
	TargetUserId *GetEntitlementsSkuIdsParameterOneOfInner `json:"target_user_id,omitempty"`
	TargetApplicationId *GetEntitlementsSkuIdsParameterOneOfInner `json:"target_application_id,omitempty"`
	TargetType *CreateGuildInviteRequestTargetType `json:"target_type,omitempty"`
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

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *CreateGuildInviteRequest) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge) {
		var ret int32
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildInviteRequest) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int32 and assigns it to the MaxAge field.
func (o *CreateGuildInviteRequest) SetMaxAge(v int32) {
	o.MaxAge = &v
}


// GetTemporary returns the Temporary field value if set, zero value otherwise.
func (o *CreateGuildInviteRequest) GetTemporary() bool {
	if o == nil || IsNil(o.Temporary) {
		var ret bool
		return ret
	}
	return *o.Temporary
}

// GetTemporaryOk returns a tuple with the Temporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildInviteRequest) GetTemporaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Temporary) {
		return nil, false
	}
	return o.Temporary, true
}

// HasTemporary returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasTemporary() bool {
	if o != nil && !IsNil(o.Temporary) {
		return true
	}

	return false
}

// SetTemporary gets a reference to the given bool and assigns it to the Temporary field.
func (o *CreateGuildInviteRequest) SetTemporary(v bool) {
	o.Temporary = &v
}


// GetMaxUses returns the MaxUses field value if set, zero value otherwise.
func (o *CreateGuildInviteRequest) GetMaxUses() int32 {
	if o == nil || IsNil(o.MaxUses) {
		var ret int32
		return ret
	}
	return *o.MaxUses
}

// GetMaxUsesOk returns a tuple with the MaxUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildInviteRequest) GetMaxUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUses) {
		return nil, false
	}
	return o.MaxUses, true
}

// HasMaxUses returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasMaxUses() bool {
	if o != nil && !IsNil(o.MaxUses) {
		return true
	}

	return false
}

// SetMaxUses gets a reference to the given int32 and assigns it to the MaxUses field.
func (o *CreateGuildInviteRequest) SetMaxUses(v int32) {
	o.MaxUses = &v
}


// GetUnique returns the Unique field value if set, zero value otherwise.
func (o *CreateGuildInviteRequest) GetUnique() bool {
	if o == nil || IsNil(o.Unique) {
		var ret bool
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildInviteRequest) GetUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.Unique) {
		return nil, false
	}
	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasUnique() bool {
	if o != nil && !IsNil(o.Unique) {
		return true
	}

	return false
}

// SetUnique gets a reference to the given bool and assigns it to the Unique field.
func (o *CreateGuildInviteRequest) SetUnique(v bool) {
	o.Unique = &v
}


// GetTargetUserId returns the TargetUserId field value if set, zero value otherwise.
func (o *CreateGuildInviteRequest) GetTargetUserId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.TargetUserId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.TargetUserId
}

// GetTargetUserIdOk returns a tuple with the TargetUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildInviteRequest) GetTargetUserIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetTargetUserId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the TargetUserId field.
func (o *CreateGuildInviteRequest) SetTargetUserId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.TargetUserId = &v
}


// GetTargetApplicationId returns the TargetApplicationId field value if set, zero value otherwise.
func (o *CreateGuildInviteRequest) GetTargetApplicationId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.TargetApplicationId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.TargetApplicationId
}

// GetTargetApplicationIdOk returns a tuple with the TargetApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildInviteRequest) GetTargetApplicationIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetTargetApplicationId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the TargetApplicationId field.
func (o *CreateGuildInviteRequest) SetTargetApplicationId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.TargetApplicationId = &v
}


// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *CreateGuildInviteRequest) GetTargetType() CreateGuildInviteRequestTargetType {
	if o == nil || IsNil(o.TargetType) {
		var ret CreateGuildInviteRequestTargetType
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildInviteRequest) GetTargetTypeOk() (*CreateGuildInviteRequestTargetType, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *CreateGuildInviteRequest) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given CreateGuildInviteRequestTargetType and assigns it to the TargetType field.
func (o *CreateGuildInviteRequest) SetTargetType(v CreateGuildInviteRequestTargetType) {
	o.TargetType = &v
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
	if !IsNil(o.MaxAge) {
		toSerialize["max_age"] = o.MaxAge
	}
	if !IsNil(o.Temporary) {
		toSerialize["temporary"] = o.Temporary
	}
	if !IsNil(o.MaxUses) {
		toSerialize["max_uses"] = o.MaxUses
	}
	if !IsNil(o.Unique) {
		toSerialize["unique"] = o.Unique
	}
	if !IsNil(o.TargetUserId) {
		toSerialize["target_user_id"] = o.TargetUserId
	}
	if !IsNil(o.TargetApplicationId) {
		toSerialize["target_application_id"] = o.TargetApplicationId
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
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

