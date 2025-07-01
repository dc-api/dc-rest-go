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

// checks if the StageInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StageInstanceResponse{}

// StageInstanceResponse struct for StageInstanceResponse
type StageInstanceResponse struct {
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Topic string `json:"topic"`
	PrivacyLevel int32 `json:"privacy_level"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	DiscoverableDisabled NullableBool `json:"discoverable_disabled,omitempty"`
	GuildScheduledEventId *string `json:"guild_scheduled_event_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _StageInstanceResponse StageInstanceResponse

// NewStageInstanceResponse instantiates a new StageInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStageInstanceResponse(guildId string, channelId string, topic string, privacyLevel int32, id string) *StageInstanceResponse {
	this := StageInstanceResponse{}
	this.GuildId = guildId
	this.ChannelId = channelId
	this.Topic = topic
	this.PrivacyLevel = privacyLevel
	this.Id = id
	return &this
}

// NewStageInstanceResponseWithDefaults instantiates a new StageInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStageInstanceResponseWithDefaults() *StageInstanceResponse {
	this := StageInstanceResponse{}
	return &this
}

// GetGuildId returns the GuildId field value
func (o *StageInstanceResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *StageInstanceResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *StageInstanceResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetChannelId returns the ChannelId field value
func (o *StageInstanceResponse) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *StageInstanceResponse) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *StageInstanceResponse) SetChannelId(v string) {
	o.ChannelId = v
}

// GetTopic returns the Topic field value
func (o *StageInstanceResponse) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *StageInstanceResponse) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *StageInstanceResponse) SetTopic(v string) {
	o.Topic = v
}

// GetPrivacyLevel returns the PrivacyLevel field value
func (o *StageInstanceResponse) GetPrivacyLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value
// and a boolean to check if the value has been set.
func (o *StageInstanceResponse) GetPrivacyLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyLevel, true
}

// SetPrivacyLevel sets field value
func (o *StageInstanceResponse) SetPrivacyLevel(v int32) {
	o.PrivacyLevel = v
}

// GetId returns the Id field value
func (o *StageInstanceResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StageInstanceResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StageInstanceResponse) SetId(v string) {
	o.Id = v
}

// GetDiscoverableDisabled returns the DiscoverableDisabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StageInstanceResponse) GetDiscoverableDisabled() bool {
	if o == nil || IsNil(o.DiscoverableDisabled.Get()) {
		var ret bool
		return ret
	}
	return *o.DiscoverableDisabled.Get()
}

// GetDiscoverableDisabledOk returns a tuple with the DiscoverableDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageInstanceResponse) GetDiscoverableDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DiscoverableDisabled.Get()) {
		return nil, false
	}
	return o.DiscoverableDisabled.Get(), o.DiscoverableDisabled.IsSet()
}

// HasDiscoverableDisabled returns a boolean if a field has been set.
func (o *StageInstanceResponse) HasDiscoverableDisabled() bool {
	if o != nil && o.DiscoverableDisabled.IsSet() {
		return true
	}

	return false
}

// SetDiscoverableDisabled gets a reference to the given NullableBool and assigns it to the DiscoverableDisabled field.
func (o *StageInstanceResponse) SetDiscoverableDisabled(v bool) {
	o.DiscoverableDisabled.Set(&v)
}

// SetDiscoverableDisabledNil sets the value for DiscoverableDisabled to be an explicit nil
func (o *StageInstanceResponse) SetDiscoverableDisabledNil() {
	o.DiscoverableDisabled.Set(nil)
}

// UnsetDiscoverableDisabled ensures that no value is present for DiscoverableDisabled, not even an explicit nil
func (o *StageInstanceResponse) UnsetDiscoverableDisabled() {
	o.DiscoverableDisabled.Unset()
}

// GetGuildScheduledEventId returns the GuildScheduledEventId field value if set, zero value otherwise.
func (o *StageInstanceResponse) GetGuildScheduledEventId() string {
	if o == nil || IsNil(o.GuildScheduledEventId) {
		var ret string
		return ret
	}
	return *o.GuildScheduledEventId
}

// GetGuildScheduledEventIdOk returns a tuple with the GuildScheduledEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageInstanceResponse) GetGuildScheduledEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildScheduledEventId) {
		return nil, false
	}
	return o.GuildScheduledEventId, true
}

// HasGuildScheduledEventId returns a boolean if a field has been set.
func (o *StageInstanceResponse) HasGuildScheduledEventId() bool {
	if o != nil && !IsNil(o.GuildScheduledEventId) {
		return true
	}

	return false
}

// SetGuildScheduledEventId gets a reference to the given string and assigns it to the GuildScheduledEventId field.
func (o *StageInstanceResponse) SetGuildScheduledEventId(v string) {
	o.GuildScheduledEventId = &v
}


func (o StageInstanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StageInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["guild_id"] = o.GuildId
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["topic"] = o.Topic
	toSerialize["privacy_level"] = o.PrivacyLevel
	toSerialize["id"] = o.Id
	if o.DiscoverableDisabled.IsSet() {
		toSerialize["discoverable_disabled"] = o.DiscoverableDisabled.Get()
	}
	if !IsNil(o.GuildScheduledEventId) {
		toSerialize["guild_scheduled_event_id"] = o.GuildScheduledEventId
	}
	return toSerialize, nil
}

func (o *StageInstanceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"guild_id",
		"channel_id",
		"topic",
		"privacy_level",
		"id",
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

	varStageInstanceResponse := _StageInstanceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStageInstanceResponse)

	if err != nil {
		return err
	}

	*o = StageInstanceResponse(varStageInstanceResponse)

	return err
}

type NullableStageInstanceResponse struct {
	value *StageInstanceResponse
	isSet bool
}

func (v NullableStageInstanceResponse) Get() *StageInstanceResponse {
	return v.value
}

func (v *NullableStageInstanceResponse) Set(val *StageInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStageInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStageInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStageInstanceResponse(val *StageInstanceResponse) *NullableStageInstanceResponse {
	return &NullableStageInstanceResponse{value: val, isSet: true}
}

func (v NullableStageInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStageInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


