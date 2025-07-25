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
)

// checks if the GuildRoleColorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildRoleColorsResponse{}

// GuildRoleColorsResponse struct for GuildRoleColorsResponse
type GuildRoleColorsResponse struct {
	PrimaryColor NullableInt32 `json:"primary_color,omitempty"`
	SecondaryColor NullableInt32 `json:"secondary_color,omitempty"`
	TertiaryColor NullableInt32 `json:"tertiary_color,omitempty"`
}

// NewGuildRoleColorsResponse instantiates a new GuildRoleColorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildRoleColorsResponse() *GuildRoleColorsResponse {
	this := GuildRoleColorsResponse{}
	return &this
}

// NewGuildRoleColorsResponseWithDefaults instantiates a new GuildRoleColorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildRoleColorsResponseWithDefaults() *GuildRoleColorsResponse {
	this := GuildRoleColorsResponse{}
	return &this
}

// GetPrimaryColor returns the PrimaryColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleColorsResponse) GetPrimaryColor() int32 {
	if o == nil || IsNil(o.PrimaryColor.Get()) {
		var ret int32
		return ret
	}
	return *o.PrimaryColor.Get()
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleColorsResponse) GetPrimaryColorOk() (*int32, bool) {
	if o == nil || IsNil(o.PrimaryColor.Get()) {
		return nil, false
	}
	return o.PrimaryColor.Get(), o.PrimaryColor.IsSet()
}

// HasPrimaryColor returns a boolean if a field has been set.
func (o *GuildRoleColorsResponse) HasPrimaryColor() bool {
	if o != nil && o.PrimaryColor.IsSet() {
		return true
	}

	return false
}

// SetPrimaryColor gets a reference to the given NullableInt32 and assigns it to the PrimaryColor field.
func (o *GuildRoleColorsResponse) SetPrimaryColor(v int32) {
	o.PrimaryColor.Set(&v)
}

// SetPrimaryColorNil sets the value for PrimaryColor to be an explicit nil
func (o *GuildRoleColorsResponse) SetPrimaryColorNil() {
	o.PrimaryColor.Set(nil)
}

// UnsetPrimaryColor ensures that no value is present for PrimaryColor, not even an explicit nil
func (o *GuildRoleColorsResponse) UnsetPrimaryColor() {
	o.PrimaryColor.Unset()
}

// GetSecondaryColor returns the SecondaryColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleColorsResponse) GetSecondaryColor() int32 {
	if o == nil || IsNil(o.SecondaryColor.Get()) {
		var ret int32
		return ret
	}
	return *o.SecondaryColor.Get()
}

// GetSecondaryColorOk returns a tuple with the SecondaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleColorsResponse) GetSecondaryColorOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondaryColor.Get()) {
		return nil, false
	}
	return o.SecondaryColor.Get(), o.SecondaryColor.IsSet()
}

// HasSecondaryColor returns a boolean if a field has been set.
func (o *GuildRoleColorsResponse) HasSecondaryColor() bool {
	if o != nil && o.SecondaryColor.IsSet() {
		return true
	}

	return false
}

// SetSecondaryColor gets a reference to the given NullableInt32 and assigns it to the SecondaryColor field.
func (o *GuildRoleColorsResponse) SetSecondaryColor(v int32) {
	o.SecondaryColor.Set(&v)
}

// SetSecondaryColorNil sets the value for SecondaryColor to be an explicit nil
func (o *GuildRoleColorsResponse) SetSecondaryColorNil() {
	o.SecondaryColor.Set(nil)
}

// UnsetSecondaryColor ensures that no value is present for SecondaryColor, not even an explicit nil
func (o *GuildRoleColorsResponse) UnsetSecondaryColor() {
	o.SecondaryColor.Unset()
}

// GetTertiaryColor returns the TertiaryColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleColorsResponse) GetTertiaryColor() int32 {
	if o == nil || IsNil(o.TertiaryColor.Get()) {
		var ret int32
		return ret
	}
	return *o.TertiaryColor.Get()
}

// GetTertiaryColorOk returns a tuple with the TertiaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleColorsResponse) GetTertiaryColorOk() (*int32, bool) {
	if o == nil || IsNil(o.TertiaryColor.Get()) {
		return nil, false
	}
	return o.TertiaryColor.Get(), o.TertiaryColor.IsSet()
}

// HasTertiaryColor returns a boolean if a field has been set.
func (o *GuildRoleColorsResponse) HasTertiaryColor() bool {
	if o != nil && o.TertiaryColor.IsSet() {
		return true
	}

	return false
}

// SetTertiaryColor gets a reference to the given NullableInt32 and assigns it to the TertiaryColor field.
func (o *GuildRoleColorsResponse) SetTertiaryColor(v int32) {
	o.TertiaryColor.Set(&v)
}

// SetTertiaryColorNil sets the value for TertiaryColor to be an explicit nil
func (o *GuildRoleColorsResponse) SetTertiaryColorNil() {
	o.TertiaryColor.Set(nil)
}

// UnsetTertiaryColor ensures that no value is present for TertiaryColor, not even an explicit nil
func (o *GuildRoleColorsResponse) UnsetTertiaryColor() {
	o.TertiaryColor.Unset()
}

func (o GuildRoleColorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildRoleColorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PrimaryColor.IsSet() {
		toSerialize["primary_color"] = o.PrimaryColor.Get()
	}
	if o.SecondaryColor.IsSet() {
		toSerialize["secondary_color"] = o.SecondaryColor.Get()
	}
	if o.TertiaryColor.IsSet() {
		toSerialize["tertiary_color"] = o.TertiaryColor.Get()
	}
	return toSerialize, nil
}

type NullableGuildRoleColorsResponse struct {
	value *GuildRoleColorsResponse
	isSet bool
}

func (v NullableGuildRoleColorsResponse) Get() *GuildRoleColorsResponse {
	return v.value
}

func (v *NullableGuildRoleColorsResponse) Set(val *GuildRoleColorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildRoleColorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildRoleColorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildRoleColorsResponse(val *GuildRoleColorsResponse) *NullableGuildRoleColorsResponse {
	return &NullableGuildRoleColorsResponse{value: val, isSet: true}
}

func (v NullableGuildRoleColorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildRoleColorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


