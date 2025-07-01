/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// checks if the ButtonComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ButtonComponentResponse{}

// ButtonComponentResponse struct for ButtonComponentResponse
type ButtonComponentResponse struct {
	Type int32 `json:"type"`
	Id int32 `json:"id"`
	CustomId NullableString `json:"custom_id,omitempty"`
	Style NullableInt32 `json:"style"`
	Label NullableString `json:"label,omitempty"`
	Disabled NullableBool `json:"disabled,omitempty"`
	Emoji NullableComponentEmojiResponse `json:"emoji,omitempty"`
	Url NullableString `json:"url,omitempty"`
	SkuId *string `json:"sku_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _ButtonComponentResponse ButtonComponentResponse

// NewButtonComponentResponse instantiates a new ButtonComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewButtonComponentResponse(type_ int32, id int32, style NullableInt32) *ButtonComponentResponse {
	this := ButtonComponentResponse{}
	this.Type = type_
	this.Id = id
	this.Style = style
	return &this
}

// NewButtonComponentResponseWithDefaults instantiates a new ButtonComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewButtonComponentResponseWithDefaults() *ButtonComponentResponse {
	this := ButtonComponentResponse{}
	return &this
}

// GetType returns the Type field value
func (o *ButtonComponentResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ButtonComponentResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ButtonComponentResponse) SetType(v int32) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ButtonComponentResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ButtonComponentResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ButtonComponentResponse) SetId(v int32) {
	o.Id = v
}

// GetCustomId returns the CustomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentResponse) GetCustomId() string {
	if o == nil || IsNil(o.CustomId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomId.Get()
}

// GetCustomIdOk returns a tuple with the CustomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentResponse) GetCustomIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomId.Get()) {
		return nil, false
	}
	return o.CustomId.Get(), o.CustomId.IsSet()
}

// HasCustomId returns a boolean if a field has been set.
func (o *ButtonComponentResponse) HasCustomId() bool {
	if o != nil && o.CustomId.IsSet() {
		return true
	}

	return false
}

// SetCustomId gets a reference to the given NullableString and assigns it to the CustomId field.
func (o *ButtonComponentResponse) SetCustomId(v string) {
	o.CustomId.Set(&v)
}

// SetCustomIdNil sets the value for CustomId to be an explicit nil
func (o *ButtonComponentResponse) SetCustomIdNil() {
	o.CustomId.Set(nil)
}

// UnsetCustomId ensures that no value is present for CustomId, not even an explicit nil
func (o *ButtonComponentResponse) UnsetCustomId() {
	o.CustomId.Unset()
}

// GetStyle returns the Style field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ButtonComponentResponse) GetStyle() int32 {
	if o == nil || o.Style.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Style.Get()
}

// GetStyleOk returns a tuple with the Style field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentResponse) GetStyleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Style.Get(), o.Style.IsSet()
}

// SetStyle sets field value
func (o *ButtonComponentResponse) SetStyle(v int32) {
	o.Style.Set(&v)
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentResponse) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentResponse) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label.Get()) {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *ButtonComponentResponse) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *ButtonComponentResponse) SetLabel(v string) {
	o.Label.Set(&v)
}

// SetLabelNil sets the value for Label to be an explicit nil
func (o *ButtonComponentResponse) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *ButtonComponentResponse) UnsetLabel() {
	o.Label.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentResponse) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentResponse) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled.Get()) {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *ButtonComponentResponse) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *ButtonComponentResponse) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}

// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *ButtonComponentResponse) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *ButtonComponentResponse) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentResponse) GetEmoji() ComponentEmojiResponse {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret ComponentEmojiResponse
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentResponse) GetEmojiOk() (*ComponentEmojiResponse, bool) {
	if o == nil || IsNil(o.Emoji.Get()) {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *ButtonComponentResponse) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullableComponentEmojiResponse and assigns it to the Emoji field.
func (o *ButtonComponentResponse) SetEmoji(v ComponentEmojiResponse) {
	o.Emoji.Set(&v)
}

// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *ButtonComponentResponse) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *ButtonComponentResponse) UnsetEmoji() {
	o.Emoji.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ButtonComponentResponse) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ButtonComponentResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url.Get()) {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ButtonComponentResponse) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ButtonComponentResponse) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *ButtonComponentResponse) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ButtonComponentResponse) UnsetUrl() {
	o.Url.Unset()
}

// GetSkuId returns the SkuId field value if set, zero value otherwise.
func (o *ButtonComponentResponse) GetSkuId() string {
	if o == nil || IsNil(o.SkuId) {
		var ret string
		return ret
	}
	return *o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ButtonComponentResponse) GetSkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.SkuId) {
		return nil, false
	}
	return o.SkuId, true
}

// HasSkuId returns a boolean if a field has been set.
func (o *ButtonComponentResponse) HasSkuId() bool {
	if o != nil && !IsNil(o.SkuId) {
		return true
	}

	return false
}

// SetSkuId gets a reference to the given string and assigns it to the SkuId field.
func (o *ButtonComponentResponse) SetSkuId(v string) {
	o.SkuId = &v
}


func (o ButtonComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ButtonComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
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
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.SkuId) {
		toSerialize["sku_id"] = o.SkuId
	}
	return toSerialize, nil
}

func (o *ButtonComponentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varButtonComponentResponse := _ButtonComponentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varButtonComponentResponse)

	if err != nil {
		return err
	}

	*o = ButtonComponentResponse(varButtonComponentResponse)

	return err
}

type NullableButtonComponentResponse struct {
	value *ButtonComponentResponse
	isSet bool
}

func (v NullableButtonComponentResponse) Get() *ButtonComponentResponse {
	return v.value
}

func (v *NullableButtonComponentResponse) Set(val *ButtonComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableButtonComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableButtonComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableButtonComponentResponse(val *ButtonComponentResponse) *NullableButtonComponentResponse {
	return &NullableButtonComponentResponse{value: val, isSet: true}
}

func (v NullableButtonComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableButtonComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


