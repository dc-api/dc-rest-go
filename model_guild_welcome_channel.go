/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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

// checks if the GuildWelcomeChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildWelcomeChannel{}

// GuildWelcomeChannel struct for GuildWelcomeChannel
type GuildWelcomeChannel struct {
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Description string `json:"description"`
	EmojiId *string `json:"emoji_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EmojiName NullableString `json:"emoji_name,omitempty"`
}

type _GuildWelcomeChannel GuildWelcomeChannel

// NewGuildWelcomeChannel instantiates a new GuildWelcomeChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildWelcomeChannel(channelId string, description string) *GuildWelcomeChannel {
	this := GuildWelcomeChannel{}
	this.ChannelId = channelId
	this.Description = description
	return &this
}

// NewGuildWelcomeChannelWithDefaults instantiates a new GuildWelcomeChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildWelcomeChannelWithDefaults() *GuildWelcomeChannel {
	this := GuildWelcomeChannel{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *GuildWelcomeChannel) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *GuildWelcomeChannel) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *GuildWelcomeChannel) SetChannelId(v string) {
	o.ChannelId = v
}

// GetDescription returns the Description field value
func (o *GuildWelcomeChannel) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GuildWelcomeChannel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GuildWelcomeChannel) SetDescription(v string) {
	o.Description = v
}

// GetEmojiId returns the EmojiId field value if set, zero value otherwise.
func (o *GuildWelcomeChannel) GetEmojiId() string {
	if o == nil || IsNil(o.EmojiId) {
		var ret string
		return ret
	}
	return *o.EmojiId
}

// GetEmojiIdOk returns a tuple with the EmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildWelcomeChannel) GetEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiId) {
		return nil, false
	}
	return o.EmojiId, true
}

// HasEmojiId returns a boolean if a field has been set.
func (o *GuildWelcomeChannel) HasEmojiId() bool {
	if o != nil && !IsNil(o.EmojiId) {
		return true
	}

	return false
}

// SetEmojiId gets a reference to the given string and assigns it to the EmojiId field.
func (o *GuildWelcomeChannel) SetEmojiId(v string) {
	o.EmojiId = &v
}


// GetEmojiName returns the EmojiName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWelcomeChannel) GetEmojiName() string {
	if o == nil || IsNil(o.EmojiName.Get()) {
		var ret string
		return ret
	}
	return *o.EmojiName.Get()
}

// GetEmojiNameOk returns a tuple with the EmojiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWelcomeChannel) GetEmojiNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiName.Get()) {
		return nil, false
	}
	return o.EmojiName.Get(), o.EmojiName.IsSet()
}

// HasEmojiName returns a boolean if a field has been set.
func (o *GuildWelcomeChannel) HasEmojiName() bool {
	if o != nil && o.EmojiName.IsSet() {
		return true
	}

	return false
}

// SetEmojiName gets a reference to the given NullableString and assigns it to the EmojiName field.
func (o *GuildWelcomeChannel) SetEmojiName(v string) {
	o.EmojiName.Set(&v)
}

// SetEmojiNameNil sets the value for EmojiName to be an explicit nil
func (o *GuildWelcomeChannel) SetEmojiNameNil() {
	o.EmojiName.Set(nil)
}

// UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
func (o *GuildWelcomeChannel) UnsetEmojiName() {
	o.EmojiName.Unset()
}

func (o GuildWelcomeChannel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildWelcomeChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["description"] = o.Description
	if !IsNil(o.EmojiId) {
		toSerialize["emoji_id"] = o.EmojiId
	}
	if o.EmojiName.IsSet() {
		toSerialize["emoji_name"] = o.EmojiName.Get()
	}
	return toSerialize, nil
}

func (o *GuildWelcomeChannel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel_id",
		"description",
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

	varGuildWelcomeChannel := _GuildWelcomeChannel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildWelcomeChannel)

	if err != nil {
		return err
	}

	*o = GuildWelcomeChannel(varGuildWelcomeChannel)

	return err
}

type NullableGuildWelcomeChannel struct {
	value *GuildWelcomeChannel
	isSet bool
}

func (v NullableGuildWelcomeChannel) Get() *GuildWelcomeChannel {
	return v.value
}

func (v *NullableGuildWelcomeChannel) Set(val *GuildWelcomeChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildWelcomeChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildWelcomeChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildWelcomeChannel(val *GuildWelcomeChannel) *NullableGuildWelcomeChannel {
	return &NullableGuildWelcomeChannel{value: val, isSet: true}
}

func (v NullableGuildWelcomeChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildWelcomeChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


