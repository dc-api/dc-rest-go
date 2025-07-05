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
	"bytes"
	"fmt"
)

// checks if the ButtonComponentForMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ButtonComponentForMessageRequest{}

// ButtonComponentForMessageRequest struct for ButtonComponentForMessageRequest
type ButtonComponentForMessageRequest struct {
	Type int32 `json:"type"`
	CustomId NullableString `json:"custom_id,omitempty"`
	Style NullableInt32 `json:"style"`
	Label NullableString `json:"label,omitempty"`
	Disabled NullableBool `json:"disabled,omitempty"`
	Url NullableString `json:"url,omitempty"`
	SkuId *string `json:"sku_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Emoji NullableComponentEmojiForMessageRequest `json:"emoji,omitempty"`
}

type _ButtonComponentForMessageRequest ButtonComponentForMessageRequest

// NewButtonComponentForMessageRequest instantiates a new ButtonComponentForMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewButtonComponentForMessageRequest(type_ int32, style NullableInt32) *ButtonComponentForMessageRequest {
	this := ButtonComponentForMessageRequest{}
	this.Type = type_
	this.Style = style
	return &this
}

// NewButtonComponentForMessageRequestWithDefaults instantiates a new ButtonComponentForMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewButtonComponentForMessageRequestWithDefaults() *ButtonComponentForMessageRequest {
	this := ButtonComponentForMessageRequest{}
	return &this
}

// GetType returns the Type field value
func (o *ButtonComponentForMessageRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ButtonComponentForMessageRequest) SetType(v int32) {
	o.Type = v
}

// GetCustomId returns the CustomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentForMessageRequest) GetCustomId() string {
	if o == nil || IsNil(o.CustomId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomId.Get()
}

// GetCustomIdOk returns a tuple with the CustomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentForMessageRequest) GetCustomIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomId.Get()) {
		return nil, false
	}
	return o.CustomId.Get(), o.CustomId.IsSet()
}

// HasCustomId returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasCustomId() bool {
	if o != nil && o.CustomId.IsSet() {
		return true
	}

	return false
}

// SetCustomId gets a reference to the given NullableString and assigns it to the CustomId field.
func (o *ButtonComponentForMessageRequest) SetCustomId(v string) {
	o.CustomId.Set(&v)
}

// SetCustomIdNil sets the value for CustomId to be an explicit nil
func (o *ButtonComponentForMessageRequest) SetCustomIdNil() {
	o.CustomId.Set(nil)
}

// UnsetCustomId ensures that no value is present for CustomId, not even an explicit nil
func (o *ButtonComponentForMessageRequest) UnsetCustomId() {
	o.CustomId.Unset()
}

// GetStyle returns the Style field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ButtonComponentForMessageRequest) GetStyle() int32 {
	if o == nil || o.Style.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Style.Get()
}

// GetStyleOk returns a tuple with the Style field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentForMessageRequest) GetStyleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Style.Get(), o.Style.IsSet()
}

// SetStyle sets field value
func (o *ButtonComponentForMessageRequest) SetStyle(v int32) {
	o.Style.Set(&v)
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentForMessageRequest) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentForMessageRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label.Get()) {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *ButtonComponentForMessageRequest) SetLabel(v string) {
	o.Label.Set(&v)
}

// SetLabelNil sets the value for Label to be an explicit nil
func (o *ButtonComponentForMessageRequest) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *ButtonComponentForMessageRequest) UnsetLabel() {
	o.Label.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentForMessageRequest) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentForMessageRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled.Get()) {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *ButtonComponentForMessageRequest) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}

// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *ButtonComponentForMessageRequest) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *ButtonComponentForMessageRequest) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentForMessageRequest) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentForMessageRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url.Get()) {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ButtonComponentForMessageRequest) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *ButtonComponentForMessageRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ButtonComponentForMessageRequest) UnsetUrl() {
	o.Url.Unset()
}

// GetSkuId returns the SkuId field value if set, zero value otherwise.
func (o *ButtonComponentForMessageRequest) GetSkuId() string {
	if o == nil || IsNil(o.SkuId) {
		var ret string
		return ret
	}
	return *o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetSkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.SkuId) {
		return nil, false
	}
	return o.SkuId, true
}

// HasSkuId returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasSkuId() bool {
	if o != nil && !IsNil(o.SkuId) {
		return true
	}

	return false
}

// SetSkuId gets a reference to the given string and assigns it to the SkuId field.
func (o *ButtonComponentForMessageRequest) SetSkuId(v string) {
	o.SkuId = &v
}


// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentForMessageRequest) GetEmoji() ComponentEmojiForMessageRequest {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret ComponentEmojiForMessageRequest
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentForMessageRequest) GetEmojiOk() (*ComponentEmojiForMessageRequest, bool) {
	if o == nil || IsNil(o.Emoji.Get()) {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullableComponentEmojiForMessageRequest and assigns it to the Emoji field.
func (o *ButtonComponentForMessageRequest) SetEmoji(v ComponentEmojiForMessageRequest) {
	o.Emoji.Set(&v)
}

// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *ButtonComponentForMessageRequest) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *ButtonComponentForMessageRequest) UnsetEmoji() {
	o.Emoji.Unset()
}

func (o ButtonComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ButtonComponentForMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.CustomId.IsSet() {
		toSerialize["custom_id"] = o.CustomId.Get()
	}
	toSerialize["style"] = o.Style.Get()
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Disabled.IsSet() {
		toSerialize["disabled"] = o.Disabled.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.SkuId) {
		toSerialize["sku_id"] = o.SkuId
	}
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	return toSerialize, nil
}

func (o *ButtonComponentForMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"style",
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

	varButtonComponentForMessageRequest := _ButtonComponentForMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varButtonComponentForMessageRequest)

	if err != nil {
		return err
	}

	*o = ButtonComponentForMessageRequest(varButtonComponentForMessageRequest)

	return err
}

type NullableButtonComponentForMessageRequest struct {
	value *ButtonComponentForMessageRequest
	isSet bool
}

func (v NullableButtonComponentForMessageRequest) Get() *ButtonComponentForMessageRequest {
	return v.value
}

func (v *NullableButtonComponentForMessageRequest) Set(val *ButtonComponentForMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableButtonComponentForMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableButtonComponentForMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableButtonComponentForMessageRequest(val *ButtonComponentForMessageRequest) *NullableButtonComponentForMessageRequest {
	return &NullableButtonComponentForMessageRequest{value: val, isSet: true}
}

func (v NullableButtonComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableButtonComponentForMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


