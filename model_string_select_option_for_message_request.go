/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
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

// checks if the StringSelectOptionForMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringSelectOptionForMessageRequest{}

// StringSelectOptionForMessageRequest struct for StringSelectOptionForMessageRequest
type StringSelectOptionForMessageRequest struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Description NullableString `json:"description,omitempty"`
	Default NullableBool `json:"default,omitempty"`
	Emoji NullableComponentEmojiForMessageRequest `json:"emoji,omitempty"`
}

type _StringSelectOptionForMessageRequest StringSelectOptionForMessageRequest

// NewStringSelectOptionForMessageRequest instantiates a new StringSelectOptionForMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringSelectOptionForMessageRequest(label string, value string) *StringSelectOptionForMessageRequest {
	this := StringSelectOptionForMessageRequest{}
	this.Label = label
	this.Value = value
	return &this
}

// NewStringSelectOptionForMessageRequestWithDefaults instantiates a new StringSelectOptionForMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringSelectOptionForMessageRequestWithDefaults() *StringSelectOptionForMessageRequest {
	this := StringSelectOptionForMessageRequest{}
	return &this
}

// GetLabel returns the Label field value
func (o *StringSelectOptionForMessageRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *StringSelectOptionForMessageRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *StringSelectOptionForMessageRequest) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *StringSelectOptionForMessageRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StringSelectOptionForMessageRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StringSelectOptionForMessageRequest) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringSelectOptionForMessageRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringSelectOptionForMessageRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StringSelectOptionForMessageRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StringSelectOptionForMessageRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StringSelectOptionForMessageRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StringSelectOptionForMessageRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringSelectOptionForMessageRequest) GetDefault() bool {
	if o == nil || IsNil(o.Default.Get()) {
		var ret bool
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringSelectOptionForMessageRequest) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default.Get()) {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *StringSelectOptionForMessageRequest) HasDefault() bool {
	if o != nil && o.Default.IsSet() {
		return true
	}

	return false
}

// SetDefault gets a reference to the given NullableBool and assigns it to the Default field.
func (o *StringSelectOptionForMessageRequest) SetDefault(v bool) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil
func (o *StringSelectOptionForMessageRequest) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil
func (o *StringSelectOptionForMessageRequest) UnsetDefault() {
	o.Default.Unset()
}

// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringSelectOptionForMessageRequest) GetEmoji() ComponentEmojiForMessageRequest {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret ComponentEmojiForMessageRequest
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringSelectOptionForMessageRequest) GetEmojiOk() (*ComponentEmojiForMessageRequest, bool) {
	if o == nil || IsNil(o.Emoji.Get()) {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *StringSelectOptionForMessageRequest) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullableComponentEmojiForMessageRequest and assigns it to the Emoji field.
func (o *StringSelectOptionForMessageRequest) SetEmoji(v ComponentEmojiForMessageRequest) {
	o.Emoji.Set(&v)
}

// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *StringSelectOptionForMessageRequest) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *StringSelectOptionForMessageRequest) UnsetEmoji() {
	o.Emoji.Unset()
}

func (o StringSelectOptionForMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringSelectOptionForMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["value"] = o.Value
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	return toSerialize, nil
}

func (o *StringSelectOptionForMessageRequest) UnmarshalJSON(data []byte) (err error) {
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

	varStringSelectOptionForMessageRequest := _StringSelectOptionForMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStringSelectOptionForMessageRequest)

	if err != nil {
		return err
	}

	*o = StringSelectOptionForMessageRequest(varStringSelectOptionForMessageRequest)

	return err
}

type NullableStringSelectOptionForMessageRequest struct {
	value *StringSelectOptionForMessageRequest
	isSet bool
}

func (v NullableStringSelectOptionForMessageRequest) Get() *StringSelectOptionForMessageRequest {
	return v.value
}

func (v *NullableStringSelectOptionForMessageRequest) Set(val *StringSelectOptionForMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStringSelectOptionForMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStringSelectOptionForMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringSelectOptionForMessageRequest(val *StringSelectOptionForMessageRequest) *NullableStringSelectOptionForMessageRequest {
	return &NullableStringSelectOptionForMessageRequest{value: val, isSet: true}
}

func (v NullableStringSelectOptionForMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringSelectOptionForMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


