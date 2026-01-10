/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-10T02:53:50.800755799Z[Etc/UTC]
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

// checks if the MessageReferenceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReferenceRequest{}

// MessageReferenceRequest struct for MessageReferenceRequest
type MessageReferenceRequest struct {
	GuildId *GetEntitlementsSkuIdsParameterOneOfInner `json:"guild_id,omitempty"`
	ChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"channel_id,omitempty"`
	MessageId string `json:"message_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	FailIfNotExists *bool `json:"fail_if_not_exists,omitempty"`
	Type *MessageReferenceRequestType `json:"type,omitempty"`
}

type _MessageReferenceRequest MessageReferenceRequest

// NewMessageReferenceRequest instantiates a new MessageReferenceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReferenceRequest(messageId string) *MessageReferenceRequest {
	this := MessageReferenceRequest{}
	this.MessageId = messageId
	return &this
}

// NewMessageReferenceRequestWithDefaults instantiates a new MessageReferenceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReferenceRequestWithDefaults() *MessageReferenceRequest {
	this := MessageReferenceRequest{}
	return &this
}

// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *MessageReferenceRequest) GetGuildId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.GuildId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReferenceRequest) GetGuildIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *MessageReferenceRequest) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the GuildId field.
func (o *MessageReferenceRequest) SetGuildId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.GuildId = &v
}


// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *MessageReferenceRequest) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReferenceRequest) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *MessageReferenceRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ChannelId field.
func (o *MessageReferenceRequest) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ChannelId = &v
}


// GetMessageId returns the MessageId field value
func (o *MessageReferenceRequest) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *MessageReferenceRequest) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *MessageReferenceRequest) SetMessageId(v string) {
	o.MessageId = v
}

// GetFailIfNotExists returns the FailIfNotExists field value if set, zero value otherwise.
func (o *MessageReferenceRequest) GetFailIfNotExists() bool {
	if o == nil || IsNil(o.FailIfNotExists) {
		var ret bool
		return ret
	}
	return *o.FailIfNotExists
}

// GetFailIfNotExistsOk returns a tuple with the FailIfNotExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReferenceRequest) GetFailIfNotExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.FailIfNotExists) {
		return nil, false
	}
	return o.FailIfNotExists, true
}

// HasFailIfNotExists returns a boolean if a field has been set.
func (o *MessageReferenceRequest) HasFailIfNotExists() bool {
	if o != nil && !IsNil(o.FailIfNotExists) {
		return true
	}

	return false
}

// SetFailIfNotExists gets a reference to the given bool and assigns it to the FailIfNotExists field.
func (o *MessageReferenceRequest) SetFailIfNotExists(v bool) {
	o.FailIfNotExists = &v
}


// GetType returns the Type field value if set, zero value otherwise.
func (o *MessageReferenceRequest) GetType() MessageReferenceRequestType {
	if o == nil || IsNil(o.Type) {
		var ret MessageReferenceRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReferenceRequest) GetTypeOk() (*MessageReferenceRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MessageReferenceRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given MessageReferenceRequestType and assigns it to the Type field.
func (o *MessageReferenceRequest) SetType(v MessageReferenceRequestType) {
	o.Type = &v
}


func (o MessageReferenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReferenceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	toSerialize["message_id"] = o.MessageId
	if !IsNil(o.FailIfNotExists) {
		toSerialize["fail_if_not_exists"] = o.FailIfNotExists
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *MessageReferenceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_id",
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

	varMessageReferenceRequest := _MessageReferenceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageReferenceRequest)

	if err != nil {
		return err
	}

	*o = MessageReferenceRequest(varMessageReferenceRequest)

	return err
}

type NullableMessageReferenceRequest struct {
	value *MessageReferenceRequest
	isSet bool
}

func (v NullableMessageReferenceRequest) Get() *MessageReferenceRequest {
	return v.value
}

func (v *NullableMessageReferenceRequest) Set(val *MessageReferenceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReferenceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReferenceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReferenceRequest(val *MessageReferenceRequest) *NullableMessageReferenceRequest {
	return &NullableMessageReferenceRequest{value: val, isSet: true}
}

func (v NullableMessageReferenceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReferenceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

