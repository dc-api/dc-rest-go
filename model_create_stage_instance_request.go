/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-08-08T14:09:23.736426080Z[Etc/UTC]
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

// checks if the CreateStageInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStageInstanceRequest{}

// CreateStageInstanceRequest struct for CreateStageInstanceRequest
type CreateStageInstanceRequest struct {
	Topic string `json:"topic"`
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	PrivacyLevel *int32 `json:"privacy_level,omitempty"`
	GuildScheduledEventId *string `json:"guild_scheduled_event_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SendStartNotification NullableBool `json:"send_start_notification,omitempty"`
}

type _CreateStageInstanceRequest CreateStageInstanceRequest

// NewCreateStageInstanceRequest instantiates a new CreateStageInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStageInstanceRequest(topic string, channelId string) *CreateStageInstanceRequest {
	this := CreateStageInstanceRequest{}
	this.Topic = topic
	this.ChannelId = channelId
	return &this
}

// NewCreateStageInstanceRequestWithDefaults instantiates a new CreateStageInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStageInstanceRequestWithDefaults() *CreateStageInstanceRequest {
	this := CreateStageInstanceRequest{}
	return &this
}

// GetTopic returns the Topic field value
func (o *CreateStageInstanceRequest) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *CreateStageInstanceRequest) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *CreateStageInstanceRequest) SetTopic(v string) {
	o.Topic = v
}

// GetChannelId returns the ChannelId field value
func (o *CreateStageInstanceRequest) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *CreateStageInstanceRequest) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *CreateStageInstanceRequest) SetChannelId(v string) {
	o.ChannelId = v
}

// GetPrivacyLevel returns the PrivacyLevel field value if set, zero value otherwise.
func (o *CreateStageInstanceRequest) GetPrivacyLevel() int32 {
	if o == nil || IsNil(o.PrivacyLevel) {
		var ret int32
		return ret
	}
	return *o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStageInstanceRequest) GetPrivacyLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.PrivacyLevel) {
		return nil, false
	}
	return o.PrivacyLevel, true
}

// HasPrivacyLevel returns a boolean if a field has been set.
func (o *CreateStageInstanceRequest) HasPrivacyLevel() bool {
	if o != nil && !IsNil(o.PrivacyLevel) {
		return true
	}

	return false
}

// SetPrivacyLevel gets a reference to the given int32 and assigns it to the PrivacyLevel field.
func (o *CreateStageInstanceRequest) SetPrivacyLevel(v int32) {
	o.PrivacyLevel = &v
}


// GetGuildScheduledEventId returns the GuildScheduledEventId field value if set, zero value otherwise.
func (o *CreateStageInstanceRequest) GetGuildScheduledEventId() string {
	if o == nil || IsNil(o.GuildScheduledEventId) {
		var ret string
		return ret
	}
	return *o.GuildScheduledEventId
}

// GetGuildScheduledEventIdOk returns a tuple with the GuildScheduledEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStageInstanceRequest) GetGuildScheduledEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildScheduledEventId) {
		return nil, false
	}
	return o.GuildScheduledEventId, true
}

// HasGuildScheduledEventId returns a boolean if a field has been set.
func (o *CreateStageInstanceRequest) HasGuildScheduledEventId() bool {
	if o != nil && !IsNil(o.GuildScheduledEventId) {
		return true
	}

	return false
}

// SetGuildScheduledEventId gets a reference to the given string and assigns it to the GuildScheduledEventId field.
func (o *CreateStageInstanceRequest) SetGuildScheduledEventId(v string) {
	o.GuildScheduledEventId = &v
}


// GetSendStartNotification returns the SendStartNotification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateStageInstanceRequest) GetSendStartNotification() bool {
	if o == nil || IsNil(o.SendStartNotification.Get()) {
		var ret bool
		return ret
	}
	return *o.SendStartNotification.Get()
}

// GetSendStartNotificationOk returns a tuple with the SendStartNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateStageInstanceRequest) GetSendStartNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.SendStartNotification.Get()) {
		return nil, false
	}
	return o.SendStartNotification.Get(), o.SendStartNotification.IsSet()
}

// HasSendStartNotification returns a boolean if a field has been set.
func (o *CreateStageInstanceRequest) HasSendStartNotification() bool {
	if o != nil && o.SendStartNotification.IsSet() {
		return true
	}

	return false
}

// SetSendStartNotification gets a reference to the given NullableBool and assigns it to the SendStartNotification field.
func (o *CreateStageInstanceRequest) SetSendStartNotification(v bool) {
	o.SendStartNotification.Set(&v)
}

// SetSendStartNotificationNil sets the value for SendStartNotification to be an explicit nil
func (o *CreateStageInstanceRequest) SetSendStartNotificationNil() {
	o.SendStartNotification.Set(nil)
}

// UnsetSendStartNotification ensures that no value is present for SendStartNotification, not even an explicit nil
func (o *CreateStageInstanceRequest) UnsetSendStartNotification() {
	o.SendStartNotification.Unset()
}

func (o CreateStageInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStageInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["topic"] = o.Topic
	toSerialize["channel_id"] = o.ChannelId
	if !IsNil(o.PrivacyLevel) {
		toSerialize["privacy_level"] = o.PrivacyLevel
	}
	if !IsNil(o.GuildScheduledEventId) {
		toSerialize["guild_scheduled_event_id"] = o.GuildScheduledEventId
	}
	if o.SendStartNotification.IsSet() {
		toSerialize["send_start_notification"] = o.SendStartNotification.Get()
	}
	return toSerialize, nil
}

func (o *CreateStageInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"topic",
		"channel_id",
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

	varCreateStageInstanceRequest := _CreateStageInstanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateStageInstanceRequest)

	if err != nil {
		return err
	}

	*o = CreateStageInstanceRequest(varCreateStageInstanceRequest)

	return err
}

type NullableCreateStageInstanceRequest struct {
	value *CreateStageInstanceRequest
	isSet bool
}

func (v NullableCreateStageInstanceRequest) Get() *CreateStageInstanceRequest {
	return v.value
}

func (v *NullableCreateStageInstanceRequest) Set(val *CreateStageInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStageInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStageInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStageInstanceRequest(val *CreateStageInstanceRequest) *NullableCreateStageInstanceRequest {
	return &NullableCreateStageInstanceRequest{value: val, isSet: true}
}

func (v NullableCreateStageInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStageInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


