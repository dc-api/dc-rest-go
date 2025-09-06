/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-09-06T05:32:22.317073992Z[Etc/UTC]
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

// checks if the StringSelectOptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringSelectOptionResponse{}

// StringSelectOptionResponse struct for StringSelectOptionResponse
type StringSelectOptionResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Description *string `json:"description,omitempty"`
	Emoji *ComponentEmojiResponse `json:"emoji,omitempty"`
	Default *bool `json:"default,omitempty"`
}

type _StringSelectOptionResponse StringSelectOptionResponse

// NewStringSelectOptionResponse instantiates a new StringSelectOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringSelectOptionResponse(label string, value string) *StringSelectOptionResponse {
	this := StringSelectOptionResponse{}
	this.Label = label
	this.Value = value
	return &this
}

// NewStringSelectOptionResponseWithDefaults instantiates a new StringSelectOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringSelectOptionResponseWithDefaults() *StringSelectOptionResponse {
	this := StringSelectOptionResponse{}
	return &this
}

// GetLabel returns the Label field value
func (o *StringSelectOptionResponse) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *StringSelectOptionResponse) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *StringSelectOptionResponse) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *StringSelectOptionResponse) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StringSelectOptionResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StringSelectOptionResponse) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StringSelectOptionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSelectOptionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StringSelectOptionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StringSelectOptionResponse) SetDescription(v string) {
	o.Description = &v
}


// GetEmoji returns the Emoji field value if set, zero value otherwise.
func (o *StringSelectOptionResponse) GetEmoji() ComponentEmojiResponse {
	if o == nil || IsNil(o.Emoji) {
		var ret ComponentEmojiResponse
		return ret
	}
	return *o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSelectOptionResponse) GetEmojiOk() (*ComponentEmojiResponse, bool) {
	if o == nil || IsNil(o.Emoji) {
		return nil, false
	}
	return o.Emoji, true
}

// HasEmoji returns a boolean if a field has been set.
func (o *StringSelectOptionResponse) HasEmoji() bool {
	if o != nil && !IsNil(o.Emoji) {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given ComponentEmojiResponse and assigns it to the Emoji field.
func (o *StringSelectOptionResponse) SetEmoji(v ComponentEmojiResponse) {
	o.Emoji = &v
}


// GetDefault returns the Default field value if set, zero value otherwise.
func (o *StringSelectOptionResponse) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSelectOptionResponse) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *StringSelectOptionResponse) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *StringSelectOptionResponse) SetDefault(v bool) {
	o.Default = &v
}


func (o StringSelectOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringSelectOptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["value"] = o.Value
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Emoji) {
		toSerialize["emoji"] = o.Emoji
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

func (o *StringSelectOptionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"value",
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

	varStringSelectOptionResponse := _StringSelectOptionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStringSelectOptionResponse)

	if err != nil {
		return err
	}

	*o = StringSelectOptionResponse(varStringSelectOptionResponse)

	return err
}

type NullableStringSelectOptionResponse struct {
	value *StringSelectOptionResponse
	isSet bool
}

func (v NullableStringSelectOptionResponse) Get() *StringSelectOptionResponse {
	return v.value
}

func (v *NullableStringSelectOptionResponse) Set(val *StringSelectOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStringSelectOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStringSelectOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringSelectOptionResponse(val *StringSelectOptionResponse) *NullableStringSelectOptionResponse {
	return &NullableStringSelectOptionResponse{value: val, isSet: true}
}

func (v NullableStringSelectOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringSelectOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


