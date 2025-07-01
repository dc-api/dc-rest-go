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

// checks if the MessageEmbedFooterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageEmbedFooterResponse{}

// MessageEmbedFooterResponse struct for MessageEmbedFooterResponse
type MessageEmbedFooterResponse struct {
	Text string `json:"text"`
	IconUrl NullableString `json:"icon_url,omitempty"`
	ProxyIconUrl NullableString `json:"proxy_icon_url,omitempty"`
}

type _MessageEmbedFooterResponse MessageEmbedFooterResponse

// NewMessageEmbedFooterResponse instantiates a new MessageEmbedFooterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEmbedFooterResponse(text string) *MessageEmbedFooterResponse {
	this := MessageEmbedFooterResponse{}
	this.Text = text
	return &this
}

// NewMessageEmbedFooterResponseWithDefaults instantiates a new MessageEmbedFooterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEmbedFooterResponseWithDefaults() *MessageEmbedFooterResponse {
	this := MessageEmbedFooterResponse{}
	return &this
}

// GetText returns the Text field value
func (o *MessageEmbedFooterResponse) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *MessageEmbedFooterResponse) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *MessageEmbedFooterResponse) SetText(v string) {
	o.Text = v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedFooterResponse) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedFooterResponse) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl.Get()) {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *MessageEmbedFooterResponse) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *MessageEmbedFooterResponse) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *MessageEmbedFooterResponse) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *MessageEmbedFooterResponse) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetProxyIconUrl returns the ProxyIconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedFooterResponse) GetProxyIconUrl() string {
	if o == nil || IsNil(o.ProxyIconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyIconUrl.Get()
}

// GetProxyIconUrlOk returns a tuple with the ProxyIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedFooterResponse) GetProxyIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyIconUrl.Get()) {
		return nil, false
	}
	return o.ProxyIconUrl.Get(), o.ProxyIconUrl.IsSet()
}

// HasProxyIconUrl returns a boolean if a field has been set.
func (o *MessageEmbedFooterResponse) HasProxyIconUrl() bool {
	if o != nil && o.ProxyIconUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyIconUrl gets a reference to the given NullableString and assigns it to the ProxyIconUrl field.
func (o *MessageEmbedFooterResponse) SetProxyIconUrl(v string) {
	o.ProxyIconUrl.Set(&v)
}

// SetProxyIconUrlNil sets the value for ProxyIconUrl to be an explicit nil
func (o *MessageEmbedFooterResponse) SetProxyIconUrlNil() {
	o.ProxyIconUrl.Set(nil)
}

// UnsetProxyIconUrl ensures that no value is present for ProxyIconUrl, not even an explicit nil
func (o *MessageEmbedFooterResponse) UnsetProxyIconUrl() {
	o.ProxyIconUrl.Unset()
}

func (o MessageEmbedFooterResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageEmbedFooterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	if o.ProxyIconUrl.IsSet() {
		toSerialize["proxy_icon_url"] = o.ProxyIconUrl.Get()
	}
	return toSerialize, nil
}

func (o *MessageEmbedFooterResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
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

	varMessageEmbedFooterResponse := _MessageEmbedFooterResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageEmbedFooterResponse)

	if err != nil {
		return err
	}

	*o = MessageEmbedFooterResponse(varMessageEmbedFooterResponse)

	return err
}

type NullableMessageEmbedFooterResponse struct {
	value *MessageEmbedFooterResponse
	isSet bool
}

func (v NullableMessageEmbedFooterResponse) Get() *MessageEmbedFooterResponse {
	return v.value
}

func (v *NullableMessageEmbedFooterResponse) Set(val *MessageEmbedFooterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEmbedFooterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEmbedFooterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEmbedFooterResponse(val *MessageEmbedFooterResponse) *NullableMessageEmbedFooterResponse {
	return &NullableMessageEmbedFooterResponse{value: val, isSet: true}
}

func (v NullableMessageEmbedFooterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEmbedFooterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


