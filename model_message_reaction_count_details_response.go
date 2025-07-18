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

// checks if the MessageReactionCountDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReactionCountDetailsResponse{}

// MessageReactionCountDetailsResponse struct for MessageReactionCountDetailsResponse
type MessageReactionCountDetailsResponse struct {
	Burst int32 `json:"burst"`
	Normal int32 `json:"normal"`
}

type _MessageReactionCountDetailsResponse MessageReactionCountDetailsResponse

// NewMessageReactionCountDetailsResponse instantiates a new MessageReactionCountDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReactionCountDetailsResponse(burst int32, normal int32) *MessageReactionCountDetailsResponse {
	this := MessageReactionCountDetailsResponse{}
	this.Burst = burst
	this.Normal = normal
	return &this
}

// NewMessageReactionCountDetailsResponseWithDefaults instantiates a new MessageReactionCountDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReactionCountDetailsResponseWithDefaults() *MessageReactionCountDetailsResponse {
	this := MessageReactionCountDetailsResponse{}
	return &this
}

// GetBurst returns the Burst field value
func (o *MessageReactionCountDetailsResponse) GetBurst() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Burst
}

// GetBurstOk returns a tuple with the Burst field value
// and a boolean to check if the value has been set.
func (o *MessageReactionCountDetailsResponse) GetBurstOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Burst, true
}

// SetBurst sets field value
func (o *MessageReactionCountDetailsResponse) SetBurst(v int32) {
	o.Burst = v
}

// GetNormal returns the Normal field value
func (o *MessageReactionCountDetailsResponse) GetNormal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Normal
}

// GetNormalOk returns a tuple with the Normal field value
// and a boolean to check if the value has been set.
func (o *MessageReactionCountDetailsResponse) GetNormalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Normal, true
}

// SetNormal sets field value
func (o *MessageReactionCountDetailsResponse) SetNormal(v int32) {
	o.Normal = v
}

func (o MessageReactionCountDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReactionCountDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["burst"] = o.Burst
	toSerialize["normal"] = o.Normal
	return toSerialize, nil
}

func (o *MessageReactionCountDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"burst",
		"normal",
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

	varMessageReactionCountDetailsResponse := _MessageReactionCountDetailsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageReactionCountDetailsResponse)

	if err != nil {
		return err
	}

	*o = MessageReactionCountDetailsResponse(varMessageReactionCountDetailsResponse)

	return err
}

type NullableMessageReactionCountDetailsResponse struct {
	value *MessageReactionCountDetailsResponse
	isSet bool
}

func (v NullableMessageReactionCountDetailsResponse) Get() *MessageReactionCountDetailsResponse {
	return v.value
}

func (v *NullableMessageReactionCountDetailsResponse) Set(val *MessageReactionCountDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReactionCountDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReactionCountDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReactionCountDetailsResponse(val *MessageReactionCountDetailsResponse) *NullableMessageReactionCountDetailsResponse {
	return &NullableMessageReactionCountDetailsResponse{value: val, isSet: true}
}

func (v NullableMessageReactionCountDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReactionCountDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


