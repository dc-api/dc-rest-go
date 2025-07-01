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
	"bytes"
	"fmt"
)

// checks if the SoundboardCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoundboardCreateRequest{}

// SoundboardCreateRequest struct for SoundboardCreateRequest
type SoundboardCreateRequest struct {
	Name string `json:"name"`
	Volume NullableFloat64 `json:"volume,omitempty"`
	EmojiId *string `json:"emoji_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EmojiName NullableString `json:"emoji_name,omitempty"`
	Sound string `json:"sound"`
}

type _SoundboardCreateRequest SoundboardCreateRequest

// NewSoundboardCreateRequest instantiates a new SoundboardCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoundboardCreateRequest(name string, sound string) *SoundboardCreateRequest {
	this := SoundboardCreateRequest{}
	this.Name = name
	this.Sound = sound
	return &this
}

// NewSoundboardCreateRequestWithDefaults instantiates a new SoundboardCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoundboardCreateRequestWithDefaults() *SoundboardCreateRequest {
	this := SoundboardCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SoundboardCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SoundboardCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SoundboardCreateRequest) SetName(v string) {
	o.Name = v
}

// GetVolume returns the Volume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoundboardCreateRequest) GetVolume() float64 {
	if o == nil || IsNil(o.Volume.Get()) {
		var ret float64
		return ret
	}
	return *o.Volume.Get()
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoundboardCreateRequest) GetVolumeOk() (*float64, bool) {
	if o == nil || IsNil(o.Volume.Get()) {
		return nil, false
	}
	return o.Volume.Get(), o.Volume.IsSet()
}

// HasVolume returns a boolean if a field has been set.
func (o *SoundboardCreateRequest) HasVolume() bool {
	if o != nil && o.Volume.IsSet() {
		return true
	}

	return false
}

// SetVolume gets a reference to the given NullableFloat64 and assigns it to the Volume field.
func (o *SoundboardCreateRequest) SetVolume(v float64) {
	o.Volume.Set(&v)
}

// SetVolumeNil sets the value for Volume to be an explicit nil
func (o *SoundboardCreateRequest) SetVolumeNil() {
	o.Volume.Set(nil)
}

// UnsetVolume ensures that no value is present for Volume, not even an explicit nil
func (o *SoundboardCreateRequest) UnsetVolume() {
	o.Volume.Unset()
}

// GetEmojiId returns the EmojiId field value if set, zero value otherwise.
func (o *SoundboardCreateRequest) GetEmojiId() string {
	if o == nil || IsNil(o.EmojiId) {
		var ret string
		return ret
	}
	return *o.EmojiId
}

// GetEmojiIdOk returns a tuple with the EmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundboardCreateRequest) GetEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiId) {
		return nil, false
	}
	return o.EmojiId, true
}

// HasEmojiId returns a boolean if a field has been set.
func (o *SoundboardCreateRequest) HasEmojiId() bool {
	if o != nil && !IsNil(o.EmojiId) {
		return true
	}

	return false
}

// SetEmojiId gets a reference to the given string and assigns it to the EmojiId field.
func (o *SoundboardCreateRequest) SetEmojiId(v string) {
	o.EmojiId = &v
}


// GetEmojiName returns the EmojiName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoundboardCreateRequest) GetEmojiName() string {
	if o == nil || IsNil(o.EmojiName.Get()) {
		var ret string
		return ret
	}
	return *o.EmojiName.Get()
}

// GetEmojiNameOk returns a tuple with the EmojiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoundboardCreateRequest) GetEmojiNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiName.Get()) {
		return nil, false
	}
	return o.EmojiName.Get(), o.EmojiName.IsSet()
}

// HasEmojiName returns a boolean if a field has been set.
func (o *SoundboardCreateRequest) HasEmojiName() bool {
	if o != nil && o.EmojiName.IsSet() {
		return true
	}

	return false
}

// SetEmojiName gets a reference to the given NullableString and assigns it to the EmojiName field.
func (o *SoundboardCreateRequest) SetEmojiName(v string) {
	o.EmojiName.Set(&v)
}

// SetEmojiNameNil sets the value for EmojiName to be an explicit nil
func (o *SoundboardCreateRequest) SetEmojiNameNil() {
	o.EmojiName.Set(nil)
}

// UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
func (o *SoundboardCreateRequest) UnsetEmojiName() {
	o.EmojiName.Unset()
}

// GetSound returns the Sound field value
func (o *SoundboardCreateRequest) GetSound() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sound
}

// GetSoundOk returns a tuple with the Sound field value
// and a boolean to check if the value has been set.
func (o *SoundboardCreateRequest) GetSoundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sound, true
}

// SetSound sets field value
func (o *SoundboardCreateRequest) SetSound(v string) {
	o.Sound = v
}

func (o SoundboardCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoundboardCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Volume.IsSet() {
		toSerialize["volume"] = o.Volume.Get()
	}
	if !IsNil(o.EmojiId) {
		toSerialize["emoji_id"] = o.EmojiId
	}
	if o.EmojiName.IsSet() {
		toSerialize["emoji_name"] = o.EmojiName.Get()
	}
	toSerialize["sound"] = o.Sound
	return toSerialize, nil
}

func (o *SoundboardCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"sound",
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

	varSoundboardCreateRequest := _SoundboardCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSoundboardCreateRequest)

	if err != nil {
		return err
	}

	*o = SoundboardCreateRequest(varSoundboardCreateRequest)

	return err
}

type NullableSoundboardCreateRequest struct {
	value *SoundboardCreateRequest
	isSet bool
}

func (v NullableSoundboardCreateRequest) Get() *SoundboardCreateRequest {
	return v.value
}

func (v *NullableSoundboardCreateRequest) Set(val *SoundboardCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSoundboardCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSoundboardCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoundboardCreateRequest(val *SoundboardCreateRequest) *NullableSoundboardCreateRequest {
	return &NullableSoundboardCreateRequest{value: val, isSet: true}
}

func (v NullableSoundboardCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoundboardCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


