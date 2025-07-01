/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
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
	"time"
	"bytes"
	"fmt"
)

// checks if the VoiceStateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoiceStateResponse{}

// VoiceStateResponse struct for VoiceStateResponse
type VoiceStateResponse struct {
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Deaf bool `json:"deaf"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Member NullableGuildMemberResponse `json:"member,omitempty"`
	Mute bool `json:"mute"`
	RequestToSpeakTimestamp NullableTime `json:"request_to_speak_timestamp,omitempty"`
	Suppress bool `json:"suppress"`
	SelfStream NullableBool `json:"self_stream,omitempty"`
	SelfDeaf bool `json:"self_deaf"`
	SelfMute bool `json:"self_mute"`
	SelfVideo bool `json:"self_video"`
	SessionId string `json:"session_id"`
	UserId string `json:"user_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _VoiceStateResponse VoiceStateResponse

// NewVoiceStateResponse instantiates a new VoiceStateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceStateResponse(deaf bool, mute bool, suppress bool, selfDeaf bool, selfMute bool, selfVideo bool, sessionId string, userId string) *VoiceStateResponse {
	this := VoiceStateResponse{}
	this.Deaf = deaf
	this.Mute = mute
	this.Suppress = suppress
	this.SelfDeaf = selfDeaf
	this.SelfMute = selfMute
	this.SelfVideo = selfVideo
	this.SessionId = sessionId
	this.UserId = userId
	return &this
}

// NewVoiceStateResponseWithDefaults instantiates a new VoiceStateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceStateResponseWithDefaults() *VoiceStateResponse {
	this := VoiceStateResponse{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *VoiceStateResponse) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *VoiceStateResponse) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *VoiceStateResponse) SetChannelId(v string) {
	o.ChannelId = &v
}


// GetDeaf returns the Deaf field value
func (o *VoiceStateResponse) GetDeaf() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deaf
}

// GetDeafOk returns a tuple with the Deaf field value
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetDeafOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deaf, true
}

// SetDeaf sets field value
func (o *VoiceStateResponse) SetDeaf(v bool) {
	o.Deaf = v
}

// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *VoiceStateResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *VoiceStateResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *VoiceStateResponse) SetGuildId(v string) {
	o.GuildId = &v
}


// GetMember returns the Member field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceStateResponse) GetMember() GuildMemberResponse {
	if o == nil || IsNil(o.Member.Get()) {
		var ret GuildMemberResponse
		return ret
	}
	return *o.Member.Get()
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceStateResponse) GetMemberOk() (*GuildMemberResponse, bool) {
	if o == nil || IsNil(o.Member.Get()) {
		return nil, false
	}
	return o.Member.Get(), o.Member.IsSet()
}

// HasMember returns a boolean if a field has been set.
func (o *VoiceStateResponse) HasMember() bool {
	if o != nil && o.Member.IsSet() {
		return true
	}

	return false
}

// SetMember gets a reference to the given NullableGuildMemberResponse and assigns it to the Member field.
func (o *VoiceStateResponse) SetMember(v GuildMemberResponse) {
	o.Member.Set(&v)
}

// SetMemberNil sets the value for Member to be an explicit nil
func (o *VoiceStateResponse) SetMemberNil() {
	o.Member.Set(nil)
}

// UnsetMember ensures that no value is present for Member, not even an explicit nil
func (o *VoiceStateResponse) UnsetMember() {
	o.Member.Unset()
}

// GetMute returns the Mute field value
func (o *VoiceStateResponse) GetMute() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Mute
}

// GetMuteOk returns a tuple with the Mute field value
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetMuteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mute, true
}

// SetMute sets field value
func (o *VoiceStateResponse) SetMute(v bool) {
	o.Mute = v
}

// GetRequestToSpeakTimestamp returns the RequestToSpeakTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceStateResponse) GetRequestToSpeakTimestamp() time.Time {
	if o == nil || IsNil(o.RequestToSpeakTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RequestToSpeakTimestamp.Get()
}

// GetRequestToSpeakTimestampOk returns a tuple with the RequestToSpeakTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceStateResponse) GetRequestToSpeakTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestToSpeakTimestamp.Get()) {
		return nil, false
	}
	return o.RequestToSpeakTimestamp.Get(), o.RequestToSpeakTimestamp.IsSet()
}

// HasRequestToSpeakTimestamp returns a boolean if a field has been set.
func (o *VoiceStateResponse) HasRequestToSpeakTimestamp() bool {
	if o != nil && o.RequestToSpeakTimestamp.IsSet() {
		return true
	}

	return false
}

// SetRequestToSpeakTimestamp gets a reference to the given NullableTime and assigns it to the RequestToSpeakTimestamp field.
func (o *VoiceStateResponse) SetRequestToSpeakTimestamp(v time.Time) {
	o.RequestToSpeakTimestamp.Set(&v)
}

// SetRequestToSpeakTimestampNil sets the value for RequestToSpeakTimestamp to be an explicit nil
func (o *VoiceStateResponse) SetRequestToSpeakTimestampNil() {
	o.RequestToSpeakTimestamp.Set(nil)
}

// UnsetRequestToSpeakTimestamp ensures that no value is present for RequestToSpeakTimestamp, not even an explicit nil
func (o *VoiceStateResponse) UnsetRequestToSpeakTimestamp() {
	o.RequestToSpeakTimestamp.Unset()
}

// GetSuppress returns the Suppress field value
func (o *VoiceStateResponse) GetSuppress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Suppress
}

// GetSuppressOk returns a tuple with the Suppress field value
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetSuppressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suppress, true
}

// SetSuppress sets field value
func (o *VoiceStateResponse) SetSuppress(v bool) {
	o.Suppress = v
}

// GetSelfStream returns the SelfStream field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceStateResponse) GetSelfStream() bool {
	if o == nil || IsNil(o.SelfStream.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfStream.Get()
}

// GetSelfStreamOk returns a tuple with the SelfStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceStateResponse) GetSelfStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfStream.Get()) {
		return nil, false
	}
	return o.SelfStream.Get(), o.SelfStream.IsSet()
}

// HasSelfStream returns a boolean if a field has been set.
func (o *VoiceStateResponse) HasSelfStream() bool {
	if o != nil && o.SelfStream.IsSet() {
		return true
	}

	return false
}

// SetSelfStream gets a reference to the given NullableBool and assigns it to the SelfStream field.
func (o *VoiceStateResponse) SetSelfStream(v bool) {
	o.SelfStream.Set(&v)
}

// SetSelfStreamNil sets the value for SelfStream to be an explicit nil
func (o *VoiceStateResponse) SetSelfStreamNil() {
	o.SelfStream.Set(nil)
}

// UnsetSelfStream ensures that no value is present for SelfStream, not even an explicit nil
func (o *VoiceStateResponse) UnsetSelfStream() {
	o.SelfStream.Unset()
}

// GetSelfDeaf returns the SelfDeaf field value
func (o *VoiceStateResponse) GetSelfDeaf() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelfDeaf
}

// GetSelfDeafOk returns a tuple with the SelfDeaf field value
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetSelfDeafOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfDeaf, true
}

// SetSelfDeaf sets field value
func (o *VoiceStateResponse) SetSelfDeaf(v bool) {
	o.SelfDeaf = v
}

// GetSelfMute returns the SelfMute field value
func (o *VoiceStateResponse) GetSelfMute() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelfMute
}

// GetSelfMuteOk returns a tuple with the SelfMute field value
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetSelfMuteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfMute, true
}

// SetSelfMute sets field value
func (o *VoiceStateResponse) SetSelfMute(v bool) {
	o.SelfMute = v
}

// GetSelfVideo returns the SelfVideo field value
func (o *VoiceStateResponse) GetSelfVideo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelfVideo
}

// GetSelfVideoOk returns a tuple with the SelfVideo field value
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetSelfVideoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfVideo, true
}

// SetSelfVideo sets field value
func (o *VoiceStateResponse) SetSelfVideo(v bool) {
	o.SelfVideo = v
}

// GetSessionId returns the SessionId field value
func (o *VoiceStateResponse) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *VoiceStateResponse) SetSessionId(v string) {
	o.SessionId = v
}

// GetUserId returns the UserId field value
func (o *VoiceStateResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *VoiceStateResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *VoiceStateResponse) SetUserId(v string) {
	o.UserId = v
}

func (o VoiceStateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceStateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	toSerialize["deaf"] = o.Deaf
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	if o.Member.IsSet() {
		toSerialize["member"] = o.Member.Get()
	}
	toSerialize["mute"] = o.Mute
	if o.RequestToSpeakTimestamp.IsSet() {
		toSerialize["request_to_speak_timestamp"] = o.RequestToSpeakTimestamp.Get()
	}
	toSerialize["suppress"] = o.Suppress
	if o.SelfStream.IsSet() {
		toSerialize["self_stream"] = o.SelfStream.Get()
	}
	toSerialize["self_deaf"] = o.SelfDeaf
	toSerialize["self_mute"] = o.SelfMute
	toSerialize["self_video"] = o.SelfVideo
	toSerialize["session_id"] = o.SessionId
	toSerialize["user_id"] = o.UserId
	return toSerialize, nil
}

func (o *VoiceStateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deaf",
		"mute",
		"suppress",
		"self_deaf",
		"self_mute",
		"self_video",
		"session_id",
		"user_id",
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

	varVoiceStateResponse := _VoiceStateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVoiceStateResponse)

	if err != nil {
		return err
	}

	*o = VoiceStateResponse(varVoiceStateResponse)

	return err
}

type NullableVoiceStateResponse struct {
	value *VoiceStateResponse
	isSet bool
}

func (v NullableVoiceStateResponse) Get() *VoiceStateResponse {
	return v.value
}

func (v *NullableVoiceStateResponse) Set(val *VoiceStateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceStateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceStateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceStateResponse(val *VoiceStateResponse) *NullableVoiceStateResponse {
	return &NullableVoiceStateResponse{value: val, isSet: true}
}

func (v NullableVoiceStateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceStateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


