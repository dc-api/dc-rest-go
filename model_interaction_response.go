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

// checks if the InteractionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InteractionResponse{}

// InteractionResponse struct for InteractionResponse
type InteractionResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type int32 `json:"type"`
	ResponseMessageId *string `json:"response_message_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ResponseMessageLoading NullableBool `json:"response_message_loading,omitempty"`
	ResponseMessageEphemeral NullableBool `json:"response_message_ephemeral,omitempty"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _InteractionResponse InteractionResponse

// NewInteractionResponse instantiates a new InteractionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractionResponse(id string, type_ int32) *InteractionResponse {
	this := InteractionResponse{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewInteractionResponseWithDefaults instantiates a new InteractionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractionResponseWithDefaults() *InteractionResponse {
	this := InteractionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *InteractionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InteractionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InteractionResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *InteractionResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InteractionResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InteractionResponse) SetType(v int32) {
	o.Type = v
}

// GetResponseMessageId returns the ResponseMessageId field value if set, zero value otherwise.
func (o *InteractionResponse) GetResponseMessageId() string {
	if o == nil || IsNil(o.ResponseMessageId) {
		var ret string
		return ret
	}
	return *o.ResponseMessageId
}

// GetResponseMessageIdOk returns a tuple with the ResponseMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InteractionResponse) GetResponseMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseMessageId) {
		return nil, false
	}
	return o.ResponseMessageId, true
}

// HasResponseMessageId returns a boolean if a field has been set.
func (o *InteractionResponse) HasResponseMessageId() bool {
	if o != nil && !IsNil(o.ResponseMessageId) {
		return true
	}

	return false
}

// SetResponseMessageId gets a reference to the given string and assigns it to the ResponseMessageId field.
func (o *InteractionResponse) SetResponseMessageId(v string) {
	o.ResponseMessageId = &v
}


// GetResponseMessageLoading returns the ResponseMessageLoading field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InteractionResponse) GetResponseMessageLoading() bool {
	if o == nil || IsNil(o.ResponseMessageLoading.Get()) {
		var ret bool
		return ret
	}
	return *o.ResponseMessageLoading.Get()
}

// GetResponseMessageLoadingOk returns a tuple with the ResponseMessageLoading field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InteractionResponse) GetResponseMessageLoadingOk() (*bool, bool) {
	if o == nil || IsNil(o.ResponseMessageLoading.Get()) {
		return nil, false
	}
	return o.ResponseMessageLoading.Get(), o.ResponseMessageLoading.IsSet()
}

// HasResponseMessageLoading returns a boolean if a field has been set.
func (o *InteractionResponse) HasResponseMessageLoading() bool {
	if o != nil && o.ResponseMessageLoading.IsSet() {
		return true
	}

	return false
}

// SetResponseMessageLoading gets a reference to the given NullableBool and assigns it to the ResponseMessageLoading field.
func (o *InteractionResponse) SetResponseMessageLoading(v bool) {
	o.ResponseMessageLoading.Set(&v)
}

// SetResponseMessageLoadingNil sets the value for ResponseMessageLoading to be an explicit nil
func (o *InteractionResponse) SetResponseMessageLoadingNil() {
	o.ResponseMessageLoading.Set(nil)
}

// UnsetResponseMessageLoading ensures that no value is present for ResponseMessageLoading, not even an explicit nil
func (o *InteractionResponse) UnsetResponseMessageLoading() {
	o.ResponseMessageLoading.Unset()
}

// GetResponseMessageEphemeral returns the ResponseMessageEphemeral field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InteractionResponse) GetResponseMessageEphemeral() bool {
	if o == nil || IsNil(o.ResponseMessageEphemeral.Get()) {
		var ret bool
		return ret
	}
	return *o.ResponseMessageEphemeral.Get()
}

// GetResponseMessageEphemeralOk returns a tuple with the ResponseMessageEphemeral field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InteractionResponse) GetResponseMessageEphemeralOk() (*bool, bool) {
	if o == nil || IsNil(o.ResponseMessageEphemeral.Get()) {
		return nil, false
	}
	return o.ResponseMessageEphemeral.Get(), o.ResponseMessageEphemeral.IsSet()
}

// HasResponseMessageEphemeral returns a boolean if a field has been set.
func (o *InteractionResponse) HasResponseMessageEphemeral() bool {
	if o != nil && o.ResponseMessageEphemeral.IsSet() {
		return true
	}

	return false
}

// SetResponseMessageEphemeral gets a reference to the given NullableBool and assigns it to the ResponseMessageEphemeral field.
func (o *InteractionResponse) SetResponseMessageEphemeral(v bool) {
	o.ResponseMessageEphemeral.Set(&v)
}

// SetResponseMessageEphemeralNil sets the value for ResponseMessageEphemeral to be an explicit nil
func (o *InteractionResponse) SetResponseMessageEphemeralNil() {
	o.ResponseMessageEphemeral.Set(nil)
}

// UnsetResponseMessageEphemeral ensures that no value is present for ResponseMessageEphemeral, not even an explicit nil
func (o *InteractionResponse) UnsetResponseMessageEphemeral() {
	o.ResponseMessageEphemeral.Unset()
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *InteractionResponse) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InteractionResponse) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *InteractionResponse) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *InteractionResponse) SetChannelId(v string) {
	o.ChannelId = &v
}


// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *InteractionResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InteractionResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *InteractionResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *InteractionResponse) SetGuildId(v string) {
	o.GuildId = &v
}


func (o InteractionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InteractionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.ResponseMessageId) {
		toSerialize["response_message_id"] = o.ResponseMessageId
	}
	if o.ResponseMessageLoading.IsSet() {
		toSerialize["response_message_loading"] = o.ResponseMessageLoading.Get()
	}
	if o.ResponseMessageEphemeral.IsSet() {
		toSerialize["response_message_ephemeral"] = o.ResponseMessageEphemeral.Get()
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	return toSerialize, nil
}

func (o *InteractionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
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

	varInteractionResponse := _InteractionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInteractionResponse)

	if err != nil {
		return err
	}

	*o = InteractionResponse(varInteractionResponse)

	return err
}

type NullableInteractionResponse struct {
	value *InteractionResponse
	isSet bool
}

func (v NullableInteractionResponse) Get() *InteractionResponse {
	return v.value
}

func (v *NullableInteractionResponse) Set(val *InteractionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractionResponse(val *InteractionResponse) *NullableInteractionResponse {
	return &NullableInteractionResponse{value: val, isSet: true}
}

func (v NullableInteractionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


