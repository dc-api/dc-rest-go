/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
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
	"time"
)

// checks if the UpdateSelfVoiceStateRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSelfVoiceStateRequestPartial{}

// UpdateSelfVoiceStateRequestPartial struct for UpdateSelfVoiceStateRequestPartial
type UpdateSelfVoiceStateRequestPartial struct {
	RequestToSpeakTimestamp *time.Time `json:"request_to_speak_timestamp,omitempty"`
	Suppress *bool `json:"suppress,omitempty"`
	ChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"channel_id,omitempty"`
}

// NewUpdateSelfVoiceStateRequestPartial instantiates a new UpdateSelfVoiceStateRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSelfVoiceStateRequestPartial() *UpdateSelfVoiceStateRequestPartial {
	this := UpdateSelfVoiceStateRequestPartial{}
	return &this
}

// NewUpdateSelfVoiceStateRequestPartialWithDefaults instantiates a new UpdateSelfVoiceStateRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSelfVoiceStateRequestPartialWithDefaults() *UpdateSelfVoiceStateRequestPartial {
	this := UpdateSelfVoiceStateRequestPartial{}
	return &this
}

// GetRequestToSpeakTimestamp returns the RequestToSpeakTimestamp field value if set, zero value otherwise.
func (o *UpdateSelfVoiceStateRequestPartial) GetRequestToSpeakTimestamp() time.Time {
	if o == nil || IsNil(o.RequestToSpeakTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.RequestToSpeakTimestamp
}

// GetRequestToSpeakTimestampOk returns a tuple with the RequestToSpeakTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSelfVoiceStateRequestPartial) GetRequestToSpeakTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestToSpeakTimestamp) {
		return nil, false
	}
	return o.RequestToSpeakTimestamp, true
}

// HasRequestToSpeakTimestamp returns a boolean if a field has been set.
func (o *UpdateSelfVoiceStateRequestPartial) HasRequestToSpeakTimestamp() bool {
	if o != nil && !IsNil(o.RequestToSpeakTimestamp) {
		return true
	}

	return false
}

// SetRequestToSpeakTimestamp gets a reference to the given time.Time and assigns it to the RequestToSpeakTimestamp field.
func (o *UpdateSelfVoiceStateRequestPartial) SetRequestToSpeakTimestamp(v time.Time) {
	o.RequestToSpeakTimestamp = &v
}


// GetSuppress returns the Suppress field value if set, zero value otherwise.
func (o *UpdateSelfVoiceStateRequestPartial) GetSuppress() bool {
	if o == nil || IsNil(o.Suppress) {
		var ret bool
		return ret
	}
	return *o.Suppress
}

// GetSuppressOk returns a tuple with the Suppress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSelfVoiceStateRequestPartial) GetSuppressOk() (*bool, bool) {
	if o == nil || IsNil(o.Suppress) {
		return nil, false
	}
	return o.Suppress, true
}

// HasSuppress returns a boolean if a field has been set.
func (o *UpdateSelfVoiceStateRequestPartial) HasSuppress() bool {
	if o != nil && !IsNil(o.Suppress) {
		return true
	}

	return false
}

// SetSuppress gets a reference to the given bool and assigns it to the Suppress field.
func (o *UpdateSelfVoiceStateRequestPartial) SetSuppress(v bool) {
	o.Suppress = &v
}


// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *UpdateSelfVoiceStateRequestPartial) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSelfVoiceStateRequestPartial) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *UpdateSelfVoiceStateRequestPartial) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ChannelId field.
func (o *UpdateSelfVoiceStateRequestPartial) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ChannelId = &v
}


func (o UpdateSelfVoiceStateRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSelfVoiceStateRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestToSpeakTimestamp) {
		toSerialize["request_to_speak_timestamp"] = o.RequestToSpeakTimestamp
	}
	if !IsNil(o.Suppress) {
		toSerialize["suppress"] = o.Suppress
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	return toSerialize, nil
}

type NullableUpdateSelfVoiceStateRequestPartial struct {
	value *UpdateSelfVoiceStateRequestPartial
	isSet bool
}

func (v NullableUpdateSelfVoiceStateRequestPartial) Get() *UpdateSelfVoiceStateRequestPartial {
	return v.value
}

func (v *NullableUpdateSelfVoiceStateRequestPartial) Set(val *UpdateSelfVoiceStateRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSelfVoiceStateRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSelfVoiceStateRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSelfVoiceStateRequestPartial(val *UpdateSelfVoiceStateRequestPartial) *NullableUpdateSelfVoiceStateRequestPartial {
	return &NullableUpdateSelfVoiceStateRequestPartial{value: val, isSet: true}
}

func (v NullableUpdateSelfVoiceStateRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSelfVoiceStateRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

