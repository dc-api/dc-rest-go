/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
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

// checks if the CreateGroupDMInviteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupDMInviteRequest{}

// CreateGroupDMInviteRequest struct for CreateGroupDMInviteRequest
type CreateGroupDMInviteRequest struct {
	MaxAge NullableInt32 `json:"max_age,omitempty"`
}

// NewCreateGroupDMInviteRequest instantiates a new CreateGroupDMInviteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupDMInviteRequest() *CreateGroupDMInviteRequest {
	this := CreateGroupDMInviteRequest{}
	return &this
}

// NewCreateGroupDMInviteRequestWithDefaults instantiates a new CreateGroupDMInviteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupDMInviteRequestWithDefaults() *CreateGroupDMInviteRequest {
	this := CreateGroupDMInviteRequest{}
	return &this
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupDMInviteRequest) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAge.Get()
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupDMInviteRequest) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge.Get()) {
		return nil, false
	}
	return o.MaxAge.Get(), o.MaxAge.IsSet()
}

// HasMaxAge returns a boolean if a field has been set.
func (o *CreateGroupDMInviteRequest) HasMaxAge() bool {
	if o != nil && o.MaxAge.IsSet() {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given NullableInt32 and assigns it to the MaxAge field.
func (o *CreateGroupDMInviteRequest) SetMaxAge(v int32) {
	o.MaxAge.Set(&v)
}

// SetMaxAgeNil sets the value for MaxAge to be an explicit nil
func (o *CreateGroupDMInviteRequest) SetMaxAgeNil() {
	o.MaxAge.Set(nil)
}

// UnsetMaxAge ensures that no value is present for MaxAge, not even an explicit nil
func (o *CreateGroupDMInviteRequest) UnsetMaxAge() {
	o.MaxAge.Unset()
}

func (o CreateGroupDMInviteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupDMInviteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxAge.IsSet() {
		toSerialize["max_age"] = o.MaxAge.Get()
	}
	return toSerialize, nil
}

type NullableCreateGroupDMInviteRequest struct {
	value *CreateGroupDMInviteRequest
	isSet bool
}

func (v NullableCreateGroupDMInviteRequest) Get() *CreateGroupDMInviteRequest {
	return v.value
}

func (v *NullableCreateGroupDMInviteRequest) Set(val *CreateGroupDMInviteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupDMInviteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupDMInviteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupDMInviteRequest(val *CreateGroupDMInviteRequest) *NullableCreateGroupDMInviteRequest {
	return &NullableCreateGroupDMInviteRequest{value: val, isSet: true}
}

func (v NullableCreateGroupDMInviteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupDMInviteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


