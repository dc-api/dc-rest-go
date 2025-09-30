/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-09-30T01:39:44.165564220Z[Etc/UTC]
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
	"bytes"
	"fmt"
)

// checks if the GuildTemplateRoleColorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildTemplateRoleColorsResponse{}

// GuildTemplateRoleColorsResponse struct for GuildTemplateRoleColorsResponse
type GuildTemplateRoleColorsResponse struct {
	PrimaryColor int32 `json:"primary_color"`
	SecondaryColor NullableInt32 `json:"secondary_color,omitempty"`
	TertiaryColor NullableInt32 `json:"tertiary_color,omitempty"`
}

type _GuildTemplateRoleColorsResponse GuildTemplateRoleColorsResponse

// NewGuildTemplateRoleColorsResponse instantiates a new GuildTemplateRoleColorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildTemplateRoleColorsResponse(primaryColor int32) *GuildTemplateRoleColorsResponse {
	this := GuildTemplateRoleColorsResponse{}
	this.PrimaryColor = primaryColor
	return &this
}

// NewGuildTemplateRoleColorsResponseWithDefaults instantiates a new GuildTemplateRoleColorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildTemplateRoleColorsResponseWithDefaults() *GuildTemplateRoleColorsResponse {
	this := GuildTemplateRoleColorsResponse{}
	return &this
}

// GetPrimaryColor returns the PrimaryColor field value
func (o *GuildTemplateRoleColorsResponse) GetPrimaryColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateRoleColorsResponse) GetPrimaryColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryColor, true
}

// SetPrimaryColor sets field value
func (o *GuildTemplateRoleColorsResponse) SetPrimaryColor(v int32) {
	o.PrimaryColor = v
}

// GetSecondaryColor returns the SecondaryColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateRoleColorsResponse) GetSecondaryColor() int32 {
	if o == nil || IsNil(o.SecondaryColor.Get()) {
		var ret int32
		return ret
	}
	return *o.SecondaryColor.Get()
}

// GetSecondaryColorOk returns a tuple with the SecondaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateRoleColorsResponse) GetSecondaryColorOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondaryColor.Get()) {
		return nil, false
	}
	return o.SecondaryColor.Get(), o.SecondaryColor.IsSet()
}

// HasSecondaryColor returns a boolean if a field has been set.
func (o *GuildTemplateRoleColorsResponse) HasSecondaryColor() bool {
	if o != nil && o.SecondaryColor.IsSet() {
		return true
	}

	return false
}

// SetSecondaryColor gets a reference to the given NullableInt32 and assigns it to the SecondaryColor field.
func (o *GuildTemplateRoleColorsResponse) SetSecondaryColor(v int32) {
	o.SecondaryColor.Set(&v)
}

// SetSecondaryColorNil sets the value for SecondaryColor to be an explicit nil
func (o *GuildTemplateRoleColorsResponse) SetSecondaryColorNil() {
	o.SecondaryColor.Set(nil)
}

// UnsetSecondaryColor ensures that no value is present for SecondaryColor, not even an explicit nil
func (o *GuildTemplateRoleColorsResponse) UnsetSecondaryColor() {
	o.SecondaryColor.Unset()
}

// GetTertiaryColor returns the TertiaryColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateRoleColorsResponse) GetTertiaryColor() int32 {
	if o == nil || IsNil(o.TertiaryColor.Get()) {
		var ret int32
		return ret
	}
	return *o.TertiaryColor.Get()
}

// GetTertiaryColorOk returns a tuple with the TertiaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateRoleColorsResponse) GetTertiaryColorOk() (*int32, bool) {
	if o == nil || IsNil(o.TertiaryColor.Get()) {
		return nil, false
	}
	return o.TertiaryColor.Get(), o.TertiaryColor.IsSet()
}

// HasTertiaryColor returns a boolean if a field has been set.
func (o *GuildTemplateRoleColorsResponse) HasTertiaryColor() bool {
	if o != nil && o.TertiaryColor.IsSet() {
		return true
	}

	return false
}

// SetTertiaryColor gets a reference to the given NullableInt32 and assigns it to the TertiaryColor field.
func (o *GuildTemplateRoleColorsResponse) SetTertiaryColor(v int32) {
	o.TertiaryColor.Set(&v)
}

// SetTertiaryColorNil sets the value for TertiaryColor to be an explicit nil
func (o *GuildTemplateRoleColorsResponse) SetTertiaryColorNil() {
	o.TertiaryColor.Set(nil)
}

// UnsetTertiaryColor ensures that no value is present for TertiaryColor, not even an explicit nil
func (o *GuildTemplateRoleColorsResponse) UnsetTertiaryColor() {
	o.TertiaryColor.Unset()
}

func (o GuildTemplateRoleColorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildTemplateRoleColorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["primary_color"] = o.PrimaryColor
	if o.SecondaryColor.IsSet() {
		toSerialize["secondary_color"] = o.SecondaryColor.Get()
	}
	if o.TertiaryColor.IsSet() {
		toSerialize["tertiary_color"] = o.TertiaryColor.Get()
	}
	return toSerialize, nil
}

func (o *GuildTemplateRoleColorsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"primary_color",
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

	varGuildTemplateRoleColorsResponse := _GuildTemplateRoleColorsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildTemplateRoleColorsResponse)

	if err != nil {
		return err
	}

	*o = GuildTemplateRoleColorsResponse(varGuildTemplateRoleColorsResponse)

	return err
}

type NullableGuildTemplateRoleColorsResponse struct {
	value *GuildTemplateRoleColorsResponse
	isSet bool
}

func (v NullableGuildTemplateRoleColorsResponse) Get() *GuildTemplateRoleColorsResponse {
	return v.value
}

func (v *NullableGuildTemplateRoleColorsResponse) Set(val *GuildTemplateRoleColorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildTemplateRoleColorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildTemplateRoleColorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildTemplateRoleColorsResponse(val *GuildTemplateRoleColorsResponse) *NullableGuildTemplateRoleColorsResponse {
	return &NullableGuildTemplateRoleColorsResponse{value: val, isSet: true}
}

func (v NullableGuildTemplateRoleColorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildTemplateRoleColorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


