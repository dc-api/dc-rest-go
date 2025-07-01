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

// checks if the MessageReactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReactionResponse{}

// MessageReactionResponse struct for MessageReactionResponse
type MessageReactionResponse struct {
	Emoji MessageReactionEmojiResponse `json:"emoji"`
	Count int32 `json:"count"`
	CountDetails MessageReactionCountDetailsResponse `json:"count_details"`
	BurstColors []string `json:"burst_colors"`
	MeBurst bool `json:"me_burst"`
	Me bool `json:"me"`
}

type _MessageReactionResponse MessageReactionResponse

// NewMessageReactionResponse instantiates a new MessageReactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReactionResponse(emoji MessageReactionEmojiResponse, count int32, countDetails MessageReactionCountDetailsResponse, burstColors []string, meBurst bool, me bool) *MessageReactionResponse {
	this := MessageReactionResponse{}
	this.Emoji = emoji
	this.Count = count
	this.CountDetails = countDetails
	this.BurstColors = burstColors
	this.MeBurst = meBurst
	this.Me = me
	return &this
}

// NewMessageReactionResponseWithDefaults instantiates a new MessageReactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReactionResponseWithDefaults() *MessageReactionResponse {
	this := MessageReactionResponse{}
	return &this
}

// GetEmoji returns the Emoji field value
func (o *MessageReactionResponse) GetEmoji() MessageReactionEmojiResponse {
	if o == nil {
		var ret MessageReactionEmojiResponse
		return ret
	}

	return o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value
// and a boolean to check if the value has been set.
func (o *MessageReactionResponse) GetEmojiOk() (*MessageReactionEmojiResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Emoji, true
}

// SetEmoji sets field value
func (o *MessageReactionResponse) SetEmoji(v MessageReactionEmojiResponse) {
	o.Emoji = v
}

// GetCount returns the Count field value
func (o *MessageReactionResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *MessageReactionResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *MessageReactionResponse) SetCount(v int32) {
	o.Count = v
}

// GetCountDetails returns the CountDetails field value
func (o *MessageReactionResponse) GetCountDetails() MessageReactionCountDetailsResponse {
	if o == nil {
		var ret MessageReactionCountDetailsResponse
		return ret
	}

	return o.CountDetails
}

// GetCountDetailsOk returns a tuple with the CountDetails field value
// and a boolean to check if the value has been set.
func (o *MessageReactionResponse) GetCountDetailsOk() (*MessageReactionCountDetailsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountDetails, true
}

// SetCountDetails sets field value
func (o *MessageReactionResponse) SetCountDetails(v MessageReactionCountDetailsResponse) {
	o.CountDetails = v
}

// GetBurstColors returns the BurstColors field value
func (o *MessageReactionResponse) GetBurstColors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BurstColors
}

// GetBurstColorsOk returns a tuple with the BurstColors field value
// and a boolean to check if the value has been set.
func (o *MessageReactionResponse) GetBurstColorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BurstColors, true
}

// SetBurstColors sets field value
func (o *MessageReactionResponse) SetBurstColors(v []string) {
	o.BurstColors = v
}

// GetMeBurst returns the MeBurst field value
func (o *MessageReactionResponse) GetMeBurst() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MeBurst
}

// GetMeBurstOk returns a tuple with the MeBurst field value
// and a boolean to check if the value has been set.
func (o *MessageReactionResponse) GetMeBurstOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeBurst, true
}

// SetMeBurst sets field value
func (o *MessageReactionResponse) SetMeBurst(v bool) {
	o.MeBurst = v
}

// GetMe returns the Me field value
func (o *MessageReactionResponse) GetMe() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Me
}

// GetMeOk returns a tuple with the Me field value
// and a boolean to check if the value has been set.
func (o *MessageReactionResponse) GetMeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Me, true
}

// SetMe sets field value
func (o *MessageReactionResponse) SetMe(v bool) {
	o.Me = v
}

func (o MessageReactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emoji"] = o.Emoji
	toSerialize["count"] = o.Count
	toSerialize["count_details"] = o.CountDetails
	toSerialize["burst_colors"] = o.BurstColors
	toSerialize["me_burst"] = o.MeBurst
	toSerialize["me"] = o.Me
	return toSerialize, nil
}

func (o *MessageReactionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emoji",
		"count",
		"count_details",
		"burst_colors",
		"me_burst",
		"me",
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

	varMessageReactionResponse := _MessageReactionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageReactionResponse)

	if err != nil {
		return err
	}

	*o = MessageReactionResponse(varMessageReactionResponse)

	return err
}

type NullableMessageReactionResponse struct {
	value *MessageReactionResponse
	isSet bool
}

func (v NullableMessageReactionResponse) Get() *MessageReactionResponse {
	return v.value
}

func (v *NullableMessageReactionResponse) Set(val *MessageReactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReactionResponse(val *MessageReactionResponse) *NullableMessageReactionResponse {
	return &NullableMessageReactionResponse{value: val, isSet: true}
}

func (v NullableMessageReactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


