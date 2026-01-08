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
	"bytes"
	"fmt"
)

// checks if the ButtonComponentForMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ButtonComponentForMessageRequest{}

// ButtonComponentForMessageRequest struct for ButtonComponentForMessageRequest
type ButtonComponentForMessageRequest struct {
	Type MessageComponentTypes `json:"type"`
	Id *int32 `json:"id,omitempty"`
	CustomId *string `json:"custom_id,omitempty"`
	Style ButtonStyleTypes `json:"style"`
	Label *string `json:"label,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Url *string `json:"url,omitempty"`
	SkuId *GetEntitlementsSkuIdsParameterOneOfInner `json:"sku_id,omitempty"`
	Emoji *ButtonComponentForMessageRequestEmoji `json:"emoji,omitempty"`
}

type _ButtonComponentForMessageRequest ButtonComponentForMessageRequest

// NewButtonComponentForMessageRequest instantiates a new ButtonComponentForMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewButtonComponentForMessageRequest(type_ MessageComponentTypes, style ButtonStyleTypes) *ButtonComponentForMessageRequest {
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
func (o *ButtonComponentForMessageRequest) GetType() MessageComponentTypes {
	if o == nil {
		var ret MessageComponentTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetTypeOk() (*MessageComponentTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ButtonComponentForMessageRequest) SetType(v MessageComponentTypes) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ButtonComponentForMessageRequest) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ButtonComponentForMessageRequest) SetId(v int32) {
	o.Id = &v
}


// GetCustomId returns the CustomId field value if set, zero value otherwise.
func (o *ButtonComponentForMessageRequest) GetCustomId() string {
	if o == nil || IsNil(o.CustomId) {
		var ret string
		return ret
	}
	return *o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetCustomIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomId) {
		return nil, false
	}
	return o.CustomId, true
}

// HasCustomId returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasCustomId() bool {
	if o != nil && !IsNil(o.CustomId) {
		return true
	}

	return false
}

// SetCustomId gets a reference to the given string and assigns it to the CustomId field.
func (o *ButtonComponentForMessageRequest) SetCustomId(v string) {
	o.CustomId = &v
}


// GetStyle returns the Style field value
func (o *ButtonComponentForMessageRequest) GetStyle() ButtonStyleTypes {
	if o == nil {
		var ret ButtonStyleTypes
		return ret
	}

	return o.Style
}

// GetStyleOk returns a tuple with the Style field value
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetStyleOk() (*ButtonStyleTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Style, true
}

// SetStyle sets field value
func (o *ButtonComponentForMessageRequest) SetStyle(v ButtonStyleTypes) {
	o.Style = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ButtonComponentForMessageRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ButtonComponentForMessageRequest) SetLabel(v string) {
	o.Label = &v
}


// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ButtonComponentForMessageRequest) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ButtonComponentForMessageRequest) SetDisabled(v bool) {
	o.Disabled = &v
}


// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ButtonComponentForMessageRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ButtonComponentForMessageRequest) SetUrl(v string) {
	o.Url = &v
}


// GetSkuId returns the SkuId field value if set, zero value otherwise.
func (o *ButtonComponentForMessageRequest) GetSkuId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.SkuId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetSkuIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetSkuId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the SkuId field.
func (o *ButtonComponentForMessageRequest) SetSkuId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.SkuId = &v
}


// GetEmoji returns the Emoji field value if set, zero value otherwise.
func (o *ButtonComponentForMessageRequest) GetEmoji() ButtonComponentForMessageRequestEmoji {
	if o == nil || IsNil(o.Emoji) {
		var ret ButtonComponentForMessageRequestEmoji
		return ret
	}
	return *o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ButtonComponentForMessageRequest) GetEmojiOk() (*ButtonComponentForMessageRequestEmoji, bool) {
	if o == nil || IsNil(o.Emoji) {
		return nil, false
	}
	return o.Emoji, true
}

// HasEmoji returns a boolean if a field has been set.
func (o *ButtonComponentForMessageRequest) HasEmoji() bool {
	if o != nil && !IsNil(o.Emoji) {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given ButtonComponentForMessageRequestEmoji and assigns it to the Emoji field.
func (o *ButtonComponentForMessageRequest) SetEmoji(v ButtonComponentForMessageRequestEmoji) {
	o.Emoji = &v
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CustomId) {
		toSerialize["custom_id"] = o.CustomId
	}
	toSerialize["style"] = o.Style
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.SkuId) {
		toSerialize["sku_id"] = o.SkuId
	}
	if !IsNil(o.Emoji) {
		toSerialize["emoji"] = o.Emoji
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

