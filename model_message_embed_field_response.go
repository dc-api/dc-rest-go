/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// checks if the MessageEmbedFieldResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageEmbedFieldResponse{}

// MessageEmbedFieldResponse struct for MessageEmbedFieldResponse
type MessageEmbedFieldResponse struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Inline bool `json:"inline"`
}

type _MessageEmbedFieldResponse MessageEmbedFieldResponse

// NewMessageEmbedFieldResponse instantiates a new MessageEmbedFieldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEmbedFieldResponse(name string, value string, inline bool) *MessageEmbedFieldResponse {
	this := MessageEmbedFieldResponse{}
	this.Name = name
	this.Value = value
	this.Inline = inline
	return &this
}

// NewMessageEmbedFieldResponseWithDefaults instantiates a new MessageEmbedFieldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEmbedFieldResponseWithDefaults() *MessageEmbedFieldResponse {
	this := MessageEmbedFieldResponse{}
	return &this
}

// GetName returns the Name field value
func (o *MessageEmbedFieldResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MessageEmbedFieldResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MessageEmbedFieldResponse) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *MessageEmbedFieldResponse) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MessageEmbedFieldResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MessageEmbedFieldResponse) SetValue(v string) {
	o.Value = v
}

// GetInline returns the Inline field value
func (o *MessageEmbedFieldResponse) GetInline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Inline
}

// GetInlineOk returns a tuple with the Inline field value
// and a boolean to check if the value has been set.
func (o *MessageEmbedFieldResponse) GetInlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inline, true
}

// SetInline sets field value
func (o *MessageEmbedFieldResponse) SetInline(v bool) {
	o.Inline = v
}

func (o MessageEmbedFieldResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageEmbedFieldResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	toSerialize["inline"] = o.Inline
	return toSerialize, nil
}

func (o *MessageEmbedFieldResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
		"inline",
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

	varMessageEmbedFieldResponse := _MessageEmbedFieldResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageEmbedFieldResponse)

	if err != nil {
		return err
	}

	*o = MessageEmbedFieldResponse(varMessageEmbedFieldResponse)

	return err
}

type NullableMessageEmbedFieldResponse struct {
	value *MessageEmbedFieldResponse
	isSet bool
}

func (v NullableMessageEmbedFieldResponse) Get() *MessageEmbedFieldResponse {
	return v.value
}

func (v *NullableMessageEmbedFieldResponse) Set(val *MessageEmbedFieldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEmbedFieldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEmbedFieldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEmbedFieldResponse(val *MessageEmbedFieldResponse) *NullableMessageEmbedFieldResponse {
	return &NullableMessageEmbedFieldResponse{value: val, isSet: true}
}

func (v NullableMessageEmbedFieldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEmbedFieldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


