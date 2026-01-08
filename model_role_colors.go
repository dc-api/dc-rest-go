/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
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

// checks if the RoleColors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleColors{}

// RoleColors struct for RoleColors
type RoleColors struct {
	PrimaryColor *int32 `json:"primary_color,omitempty"`
	SecondaryColor *int32 `json:"secondary_color,omitempty"`
	TertiaryColor *int32 `json:"tertiary_color,omitempty"`
}

// NewRoleColors instantiates a new RoleColors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleColors() *RoleColors {
	this := RoleColors{}
	return &this
}

// NewRoleColorsWithDefaults instantiates a new RoleColors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleColorsWithDefaults() *RoleColors {
	this := RoleColors{}
	return &this
}

// GetPrimaryColor returns the PrimaryColor field value if set, zero value otherwise.
func (o *RoleColors) GetPrimaryColor() int32 {
	if o == nil || IsNil(o.PrimaryColor) {
		var ret int32
		return ret
	}
	return *o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleColors) GetPrimaryColorOk() (*int32, bool) {
	if o == nil || IsNil(o.PrimaryColor) {
		return nil, false
	}
	return o.PrimaryColor, true
}

// HasPrimaryColor returns a boolean if a field has been set.
func (o *RoleColors) HasPrimaryColor() bool {
	if o != nil && !IsNil(o.PrimaryColor) {
		return true
	}

	return false
}

// SetPrimaryColor gets a reference to the given int32 and assigns it to the PrimaryColor field.
func (o *RoleColors) SetPrimaryColor(v int32) {
	o.PrimaryColor = &v
}


// GetSecondaryColor returns the SecondaryColor field value if set, zero value otherwise.
func (o *RoleColors) GetSecondaryColor() int32 {
	if o == nil || IsNil(o.SecondaryColor) {
		var ret int32
		return ret
	}
	return *o.SecondaryColor
}

// GetSecondaryColorOk returns a tuple with the SecondaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleColors) GetSecondaryColorOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondaryColor) {
		return nil, false
	}
	return o.SecondaryColor, true
}

// HasSecondaryColor returns a boolean if a field has been set.
func (o *RoleColors) HasSecondaryColor() bool {
	if o != nil && !IsNil(o.SecondaryColor) {
		return true
	}

	return false
}

// SetSecondaryColor gets a reference to the given int32 and assigns it to the SecondaryColor field.
func (o *RoleColors) SetSecondaryColor(v int32) {
	o.SecondaryColor = &v
}


// GetTertiaryColor returns the TertiaryColor field value if set, zero value otherwise.
func (o *RoleColors) GetTertiaryColor() int32 {
	if o == nil || IsNil(o.TertiaryColor) {
		var ret int32
		return ret
	}
	return *o.TertiaryColor
}

// GetTertiaryColorOk returns a tuple with the TertiaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleColors) GetTertiaryColorOk() (*int32, bool) {
	if o == nil || IsNil(o.TertiaryColor) {
		return nil, false
	}
	return o.TertiaryColor, true
}

// HasTertiaryColor returns a boolean if a field has been set.
func (o *RoleColors) HasTertiaryColor() bool {
	if o != nil && !IsNil(o.TertiaryColor) {
		return true
	}

	return false
}

// SetTertiaryColor gets a reference to the given int32 and assigns it to the TertiaryColor field.
func (o *RoleColors) SetTertiaryColor(v int32) {
	o.TertiaryColor = &v
}


func (o RoleColors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleColors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrimaryColor) {
		toSerialize["primary_color"] = o.PrimaryColor
	}
	if !IsNil(o.SecondaryColor) {
		toSerialize["secondary_color"] = o.SecondaryColor
	}
	if !IsNil(o.TertiaryColor) {
		toSerialize["tertiary_color"] = o.TertiaryColor
	}
	return toSerialize, nil
}

type NullableRoleColors struct {
	value *RoleColors
	isSet bool
}

func (v NullableRoleColors) Get() *RoleColors {
	return v.value
}

func (v *NullableRoleColors) Set(val *RoleColors) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleColors) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleColors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleColors(val *RoleColors) *NullableRoleColors {
	return &NullableRoleColors{value: val, isSet: true}
}

func (v NullableRoleColors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleColors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

