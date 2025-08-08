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

// checks if the SoundboardSoundSendRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoundboardSoundSendRequest{}

// SoundboardSoundSendRequest struct for SoundboardSoundSendRequest
type SoundboardSoundSendRequest struct {
	SoundId string `json:"sound_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SourceGuildId *string `json:"source_guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _SoundboardSoundSendRequest SoundboardSoundSendRequest

// NewSoundboardSoundSendRequest instantiates a new SoundboardSoundSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoundboardSoundSendRequest(soundId string) *SoundboardSoundSendRequest {
	this := SoundboardSoundSendRequest{}
	this.SoundId = soundId
	return &this
}

// NewSoundboardSoundSendRequestWithDefaults instantiates a new SoundboardSoundSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoundboardSoundSendRequestWithDefaults() *SoundboardSoundSendRequest {
	this := SoundboardSoundSendRequest{}
	return &this
}

// GetSoundId returns the SoundId field value
func (o *SoundboardSoundSendRequest) GetSoundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SoundId
}

// GetSoundIdOk returns a tuple with the SoundId field value
// and a boolean to check if the value has been set.
func (o *SoundboardSoundSendRequest) GetSoundIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoundId, true
}

// SetSoundId sets field value
func (o *SoundboardSoundSendRequest) SetSoundId(v string) {
	o.SoundId = v
}

// GetSourceGuildId returns the SourceGuildId field value if set, zero value otherwise.
func (o *SoundboardSoundSendRequest) GetSourceGuildId() string {
	if o == nil || IsNil(o.SourceGuildId) {
		var ret string
		return ret
	}
	return *o.SourceGuildId
}

// GetSourceGuildIdOk returns a tuple with the SourceGuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundboardSoundSendRequest) GetSourceGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceGuildId) {
		return nil, false
	}
	return o.SourceGuildId, true
}

// HasSourceGuildId returns a boolean if a field has been set.
func (o *SoundboardSoundSendRequest) HasSourceGuildId() bool {
	if o != nil && !IsNil(o.SourceGuildId) {
		return true
	}

	return false
}

// SetSourceGuildId gets a reference to the given string and assigns it to the SourceGuildId field.
func (o *SoundboardSoundSendRequest) SetSourceGuildId(v string) {
	o.SourceGuildId = &v
}


func (o SoundboardSoundSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoundboardSoundSendRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sound_id"] = o.SoundId
	if !IsNil(o.SourceGuildId) {
		toSerialize["source_guild_id"] = o.SourceGuildId
	}
	return toSerialize, nil
}

func (o *SoundboardSoundSendRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sound_id",
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

	varSoundboardSoundSendRequest := _SoundboardSoundSendRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSoundboardSoundSendRequest)

	if err != nil {
		return err
	}

	*o = SoundboardSoundSendRequest(varSoundboardSoundSendRequest)

	return err
}

type NullableSoundboardSoundSendRequest struct {
	value *SoundboardSoundSendRequest
	isSet bool
}

func (v NullableSoundboardSoundSendRequest) Get() *SoundboardSoundSendRequest {
	return v.value
}

func (v *NullableSoundboardSoundSendRequest) Set(val *SoundboardSoundSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSoundboardSoundSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSoundboardSoundSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoundboardSoundSendRequest(val *SoundboardSoundSendRequest) *NullableSoundboardSoundSendRequest {
	return &NullableSoundboardSoundSendRequest{value: val, isSet: true}
}

func (v NullableSoundboardSoundSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoundboardSoundSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


