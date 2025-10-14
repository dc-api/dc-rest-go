/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-10-14T14:27:02.905060650Z[Etc/UTC]
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

// checks if the CustomClientThemeShareRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomClientThemeShareRequest{}

// CustomClientThemeShareRequest struct for CustomClientThemeShareRequest
type CustomClientThemeShareRequest struct {
	Colors []string `json:"colors"`
	GradientAngle int32 `json:"gradient_angle"`
	BaseMix int32 `json:"base_mix"`
	BaseTheme *int32 `json:"base_theme,omitempty"`
}

type _CustomClientThemeShareRequest CustomClientThemeShareRequest

// NewCustomClientThemeShareRequest instantiates a new CustomClientThemeShareRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomClientThemeShareRequest(colors []string, gradientAngle int32, baseMix int32) *CustomClientThemeShareRequest {
	this := CustomClientThemeShareRequest{}
	this.Colors = colors
	this.GradientAngle = gradientAngle
	this.BaseMix = baseMix
	return &this
}

// NewCustomClientThemeShareRequestWithDefaults instantiates a new CustomClientThemeShareRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomClientThemeShareRequestWithDefaults() *CustomClientThemeShareRequest {
	this := CustomClientThemeShareRequest{}
	return &this
}

// GetColors returns the Colors field value
func (o *CustomClientThemeShareRequest) GetColors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Colors
}

// GetColorsOk returns a tuple with the Colors field value
// and a boolean to check if the value has been set.
func (o *CustomClientThemeShareRequest) GetColorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Colors, true
}

// SetColors sets field value
func (o *CustomClientThemeShareRequest) SetColors(v []string) {
	o.Colors = v
}

// GetGradientAngle returns the GradientAngle field value
func (o *CustomClientThemeShareRequest) GetGradientAngle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GradientAngle
}

// GetGradientAngleOk returns a tuple with the GradientAngle field value
// and a boolean to check if the value has been set.
func (o *CustomClientThemeShareRequest) GetGradientAngleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GradientAngle, true
}

// SetGradientAngle sets field value
func (o *CustomClientThemeShareRequest) SetGradientAngle(v int32) {
	o.GradientAngle = v
}

// GetBaseMix returns the BaseMix field value
func (o *CustomClientThemeShareRequest) GetBaseMix() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BaseMix
}

// GetBaseMixOk returns a tuple with the BaseMix field value
// and a boolean to check if the value has been set.
func (o *CustomClientThemeShareRequest) GetBaseMixOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseMix, true
}

// SetBaseMix sets field value
func (o *CustomClientThemeShareRequest) SetBaseMix(v int32) {
	o.BaseMix = v
}

// GetBaseTheme returns the BaseTheme field value if set, zero value otherwise.
func (o *CustomClientThemeShareRequest) GetBaseTheme() int32 {
	if o == nil || IsNil(o.BaseTheme) {
		var ret int32
		return ret
	}
	return *o.BaseTheme
}

// GetBaseThemeOk returns a tuple with the BaseTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomClientThemeShareRequest) GetBaseThemeOk() (*int32, bool) {
	if o == nil || IsNil(o.BaseTheme) {
		return nil, false
	}
	return o.BaseTheme, true
}

// HasBaseTheme returns a boolean if a field has been set.
func (o *CustomClientThemeShareRequest) HasBaseTheme() bool {
	if o != nil && !IsNil(o.BaseTheme) {
		return true
	}

	return false
}

// SetBaseTheme gets a reference to the given int32 and assigns it to the BaseTheme field.
func (o *CustomClientThemeShareRequest) SetBaseTheme(v int32) {
	o.BaseTheme = &v
}


func (o CustomClientThemeShareRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomClientThemeShareRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["colors"] = o.Colors
	toSerialize["gradient_angle"] = o.GradientAngle
	toSerialize["base_mix"] = o.BaseMix
	if !IsNil(o.BaseTheme) {
		toSerialize["base_theme"] = o.BaseTheme
	}
	return toSerialize, nil
}

func (o *CustomClientThemeShareRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"colors",
		"gradient_angle",
		"base_mix",
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

	varCustomClientThemeShareRequest := _CustomClientThemeShareRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomClientThemeShareRequest)

	if err != nil {
		return err
	}

	*o = CustomClientThemeShareRequest(varCustomClientThemeShareRequest)

	return err
}

type NullableCustomClientThemeShareRequest struct {
	value *CustomClientThemeShareRequest
	isSet bool
}

func (v NullableCustomClientThemeShareRequest) Get() *CustomClientThemeShareRequest {
	return v.value
}

func (v *NullableCustomClientThemeShareRequest) Set(val *CustomClientThemeShareRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomClientThemeShareRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomClientThemeShareRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomClientThemeShareRequest(val *CustomClientThemeShareRequest) *NullableCustomClientThemeShareRequest {
	return &NullableCustomClientThemeShareRequest{value: val, isSet: true}
}

func (v NullableCustomClientThemeShareRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomClientThemeShareRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

