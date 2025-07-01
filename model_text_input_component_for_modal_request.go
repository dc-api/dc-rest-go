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

// checks if the TextInputComponentForModalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextInputComponentForModalRequest{}

// TextInputComponentForModalRequest struct for TextInputComponentForModalRequest
type TextInputComponentForModalRequest struct {
	Type int32 `json:"type"`
	CustomId string `json:"custom_id"`
	Style int32 `json:"style"`
	Label string `json:"label"`
	Value NullableString `json:"value,omitempty"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	Required NullableBool `json:"required,omitempty"`
	MinLength NullableInt32 `json:"min_length,omitempty"`
	MaxLength NullableInt32 `json:"max_length,omitempty"`
}

type _TextInputComponentForModalRequest TextInputComponentForModalRequest

// NewTextInputComponentForModalRequest instantiates a new TextInputComponentForModalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextInputComponentForModalRequest(type_ int32, customId string, style int32, label string) *TextInputComponentForModalRequest {
	this := TextInputComponentForModalRequest{}
	this.Type = type_
	this.CustomId = customId
	this.Style = style
	this.Label = label
	return &this
}

// NewTextInputComponentForModalRequestWithDefaults instantiates a new TextInputComponentForModalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextInputComponentForModalRequestWithDefaults() *TextInputComponentForModalRequest {
	this := TextInputComponentForModalRequest{}
	return &this
}

// GetType returns the Type field value
func (o *TextInputComponentForModalRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextInputComponentForModalRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TextInputComponentForModalRequest) SetType(v int32) {
	o.Type = v
}

// GetCustomId returns the CustomId field value
func (o *TextInputComponentForModalRequest) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *TextInputComponentForModalRequest) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *TextInputComponentForModalRequest) SetCustomId(v string) {
	o.CustomId = v
}

// GetStyle returns the Style field value
func (o *TextInputComponentForModalRequest) GetStyle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Style
}

// GetStyleOk returns a tuple with the Style field value
// and a boolean to check if the value has been set.
func (o *TextInputComponentForModalRequest) GetStyleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Style, true
}

// SetStyle sets field value
func (o *TextInputComponentForModalRequest) SetStyle(v int32) {
	o.Style = v
}

// GetLabel returns the Label field value
func (o *TextInputComponentForModalRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TextInputComponentForModalRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *TextInputComponentForModalRequest) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextInputComponentForModalRequest) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextInputComponentForModalRequest) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value.Get()) {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *TextInputComponentForModalRequest) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *TextInputComponentForModalRequest) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *TextInputComponentForModalRequest) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *TextInputComponentForModalRequest) UnsetValue() {
	o.Value.Unset()
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextInputComponentForModalRequest) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextInputComponentForModalRequest) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder.Get()) {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *TextInputComponentForModalRequest) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *TextInputComponentForModalRequest) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}

// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *TextInputComponentForModalRequest) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *TextInputComponentForModalRequest) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextInputComponentForModalRequest) GetRequired() bool {
	if o == nil || IsNil(o.Required.Get()) {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextInputComponentForModalRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required.Get()) {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *TextInputComponentForModalRequest) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *TextInputComponentForModalRequest) SetRequired(v bool) {
	o.Required.Set(&v)
}

// SetRequiredNil sets the value for Required to be an explicit nil
func (o *TextInputComponentForModalRequest) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *TextInputComponentForModalRequest) UnsetRequired() {
	o.Required.Unset()
}

// GetMinLength returns the MinLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextInputComponentForModalRequest) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MinLength.Get()
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextInputComponentForModalRequest) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength.Get()) {
		return nil, false
	}
	return o.MinLength.Get(), o.MinLength.IsSet()
}

// HasMinLength returns a boolean if a field has been set.
func (o *TextInputComponentForModalRequest) HasMinLength() bool {
	if o != nil && o.MinLength.IsSet() {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given NullableInt32 and assigns it to the MinLength field.
func (o *TextInputComponentForModalRequest) SetMinLength(v int32) {
	o.MinLength.Set(&v)
}

// SetMinLengthNil sets the value for MinLength to be an explicit nil
func (o *TextInputComponentForModalRequest) SetMinLengthNil() {
	o.MinLength.Set(nil)
}

// UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
func (o *TextInputComponentForModalRequest) UnsetMinLength() {
	o.MinLength.Unset()
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextInputComponentForModalRequest) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxLength.Get()
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextInputComponentForModalRequest) GetMaxLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLength.Get()) {
		return nil, false
	}
	return o.MaxLength.Get(), o.MaxLength.IsSet()
}

// HasMaxLength returns a boolean if a field has been set.
func (o *TextInputComponentForModalRequest) HasMaxLength() bool {
	if o != nil && o.MaxLength.IsSet() {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given NullableInt32 and assigns it to the MaxLength field.
func (o *TextInputComponentForModalRequest) SetMaxLength(v int32) {
	o.MaxLength.Set(&v)
}

// SetMaxLengthNil sets the value for MaxLength to be an explicit nil
func (o *TextInputComponentForModalRequest) SetMaxLengthNil() {
	o.MaxLength.Set(nil)
}

// UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
func (o *TextInputComponentForModalRequest) UnsetMaxLength() {
	o.MaxLength.Unset()
}

func (o TextInputComponentForModalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextInputComponentForModalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["custom_id"] = o.CustomId
	toSerialize["style"] = o.Style
	toSerialize["label"] = o.Label
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if o.Required.IsSet() {
		toSerialize["required"] = o.Required.Get()
	}
	if o.MinLength.IsSet() {
		toSerialize["min_length"] = o.MinLength.Get()
	}
	if o.MaxLength.IsSet() {
		toSerialize["max_length"] = o.MaxLength.Get()
	}
	return toSerialize, nil
}

func (o *TextInputComponentForModalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"custom_id",
		"style",
		"label",
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

	varTextInputComponentForModalRequest := _TextInputComponentForModalRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTextInputComponentForModalRequest)

	if err != nil {
		return err
	}

	*o = TextInputComponentForModalRequest(varTextInputComponentForModalRequest)

	return err
}

type NullableTextInputComponentForModalRequest struct {
	value *TextInputComponentForModalRequest
	isSet bool
}

func (v NullableTextInputComponentForModalRequest) Get() *TextInputComponentForModalRequest {
	return v.value
}

func (v *NullableTextInputComponentForModalRequest) Set(val *TextInputComponentForModalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTextInputComponentForModalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTextInputComponentForModalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextInputComponentForModalRequest(val *TextInputComponentForModalRequest) *NullableTextInputComponentForModalRequest {
	return &NullableTextInputComponentForModalRequest{value: val, isSet: true}
}

func (v NullableTextInputComponentForModalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextInputComponentForModalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


