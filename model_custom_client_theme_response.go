/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-10-25T01:29:55.152026602Z[Etc/UTC]
 *    * - **Generator Version**: 7.16.0
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

// checks if the CustomClientThemeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomClientThemeResponse{}

// CustomClientThemeResponse struct for CustomClientThemeResponse
type CustomClientThemeResponse struct {
	Colors []string `json:"colors"`
	GradientAngle int32 `json:"gradient_angle"`
	BaseMix int32 `json:"base_mix"`
	BaseTheme int32 `json:"base_theme"`
}

type _CustomClientThemeResponse CustomClientThemeResponse

// NewCustomClientThemeResponse instantiates a new CustomClientThemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomClientThemeResponse(colors []string, gradientAngle int32, baseMix int32, baseTheme int32) *CustomClientThemeResponse {
	this := CustomClientThemeResponse{}
	this.Colors = colors
	this.GradientAngle = gradientAngle
	this.BaseMix = baseMix
	this.BaseTheme = baseTheme
	return &this
}

// NewCustomClientThemeResponseWithDefaults instantiates a new CustomClientThemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomClientThemeResponseWithDefaults() *CustomClientThemeResponse {
	this := CustomClientThemeResponse{}
	return &this
}

// GetColors returns the Colors field value
func (o *CustomClientThemeResponse) GetColors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Colors
}

// GetColorsOk returns a tuple with the Colors field value
// and a boolean to check if the value has been set.
func (o *CustomClientThemeResponse) GetColorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Colors, true
}

// SetColors sets field value
func (o *CustomClientThemeResponse) SetColors(v []string) {
	o.Colors = v
}

// GetGradientAngle returns the GradientAngle field value
func (o *CustomClientThemeResponse) GetGradientAngle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GradientAngle
}

// GetGradientAngleOk returns a tuple with the GradientAngle field value
// and a boolean to check if the value has been set.
func (o *CustomClientThemeResponse) GetGradientAngleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GradientAngle, true
}

// SetGradientAngle sets field value
func (o *CustomClientThemeResponse) SetGradientAngle(v int32) {
	o.GradientAngle = v
}

// GetBaseMix returns the BaseMix field value
func (o *CustomClientThemeResponse) GetBaseMix() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BaseMix
}

// GetBaseMixOk returns a tuple with the BaseMix field value
// and a boolean to check if the value has been set.
func (o *CustomClientThemeResponse) GetBaseMixOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseMix, true
}

// SetBaseMix sets field value
func (o *CustomClientThemeResponse) SetBaseMix(v int32) {
	o.BaseMix = v
}

// GetBaseTheme returns the BaseTheme field value
func (o *CustomClientThemeResponse) GetBaseTheme() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BaseTheme
}

// GetBaseThemeOk returns a tuple with the BaseTheme field value
// and a boolean to check if the value has been set.
func (o *CustomClientThemeResponse) GetBaseThemeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseTheme, true
}

// SetBaseTheme sets field value
func (o *CustomClientThemeResponse) SetBaseTheme(v int32) {
	o.BaseTheme = v
}

func (o CustomClientThemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomClientThemeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["colors"] = o.Colors
	toSerialize["gradient_angle"] = o.GradientAngle
	toSerialize["base_mix"] = o.BaseMix
	toSerialize["base_theme"] = o.BaseTheme
	return toSerialize, nil
}

func (o *CustomClientThemeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"colors",
		"gradient_angle",
		"base_mix",
		"base_theme",
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

	varCustomClientThemeResponse := _CustomClientThemeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomClientThemeResponse)

	if err != nil {
		return err
	}

	*o = CustomClientThemeResponse(varCustomClientThemeResponse)

	return err
}

type NullableCustomClientThemeResponse struct {
	value *CustomClientThemeResponse
	isSet bool
}

func (v NullableCustomClientThemeResponse) Get() *CustomClientThemeResponse {
	return v.value
}

func (v *NullableCustomClientThemeResponse) Set(val *CustomClientThemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomClientThemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomClientThemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomClientThemeResponse(val *CustomClientThemeResponse) *NullableCustomClientThemeResponse {
	return &NullableCustomClientThemeResponse{value: val, isSet: true}
}

func (v NullableCustomClientThemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomClientThemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

