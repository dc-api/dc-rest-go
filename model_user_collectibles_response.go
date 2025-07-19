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
)

// checks if the UserCollectiblesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCollectiblesResponse{}

// UserCollectiblesResponse struct for UserCollectiblesResponse
type UserCollectiblesResponse struct {
	Nameplate NullableUserNameplateResponse `json:"nameplate,omitempty"`
}

// NewUserCollectiblesResponse instantiates a new UserCollectiblesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCollectiblesResponse() *UserCollectiblesResponse {
	this := UserCollectiblesResponse{}
	return &this
}

// NewUserCollectiblesResponseWithDefaults instantiates a new UserCollectiblesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCollectiblesResponseWithDefaults() *UserCollectiblesResponse {
	this := UserCollectiblesResponse{}
	return &this
}

// GetNameplate returns the Nameplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCollectiblesResponse) GetNameplate() UserNameplateResponse {
	if o == nil || IsNil(o.Nameplate.Get()) {
		var ret UserNameplateResponse
		return ret
	}
	return *o.Nameplate.Get()
}

// GetNameplateOk returns a tuple with the Nameplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCollectiblesResponse) GetNameplateOk() (*UserNameplateResponse, bool) {
	if o == nil || IsNil(o.Nameplate.Get()) {
		return nil, false
	}
	return o.Nameplate.Get(), o.Nameplate.IsSet()
}

// HasNameplate returns a boolean if a field has been set.
func (o *UserCollectiblesResponse) HasNameplate() bool {
	if o != nil && o.Nameplate.IsSet() {
		return true
	}

	return false
}

// SetNameplate gets a reference to the given NullableUserNameplateResponse and assigns it to the Nameplate field.
func (o *UserCollectiblesResponse) SetNameplate(v UserNameplateResponse) {
	o.Nameplate.Set(&v)
}

// SetNameplateNil sets the value for Nameplate to be an explicit nil
func (o *UserCollectiblesResponse) SetNameplateNil() {
	o.Nameplate.Set(nil)
}

// UnsetNameplate ensures that no value is present for Nameplate, not even an explicit nil
func (o *UserCollectiblesResponse) UnsetNameplate() {
	o.Nameplate.Unset()
}

func (o UserCollectiblesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCollectiblesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Nameplate.IsSet() {
		toSerialize["nameplate"] = o.Nameplate.Get()
	}
	return toSerialize, nil
}

type NullableUserCollectiblesResponse struct {
	value *UserCollectiblesResponse
	isSet bool
}

func (v NullableUserCollectiblesResponse) Get() *UserCollectiblesResponse {
	return v.value
}

func (v *NullableUserCollectiblesResponse) Set(val *UserCollectiblesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCollectiblesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCollectiblesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCollectiblesResponse(val *UserCollectiblesResponse) *NullableUserCollectiblesResponse {
	return &NullableUserCollectiblesResponse{value: val, isSet: true}
}

func (v NullableUserCollectiblesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCollectiblesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


