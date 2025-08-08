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

// checks if the SoundboardSoundResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoundboardSoundResponse{}

// SoundboardSoundResponse struct for SoundboardSoundResponse
type SoundboardSoundResponse struct {
	Name string `json:"name"`
	SoundId string `json:"sound_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Volume float64 `json:"volume"`
	EmojiId *string `json:"emoji_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EmojiName NullableString `json:"emoji_name,omitempty"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Available bool `json:"available"`
	User NullableUserResponse `json:"user,omitempty"`
}

type _SoundboardSoundResponse SoundboardSoundResponse

// NewSoundboardSoundResponse instantiates a new SoundboardSoundResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoundboardSoundResponse(name string, soundId string, volume float64, available bool) *SoundboardSoundResponse {
	this := SoundboardSoundResponse{}
	this.Name = name
	this.SoundId = soundId
	this.Volume = volume
	this.Available = available
	return &this
}

// NewSoundboardSoundResponseWithDefaults instantiates a new SoundboardSoundResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoundboardSoundResponseWithDefaults() *SoundboardSoundResponse {
	this := SoundboardSoundResponse{}
	return &this
}

// GetName returns the Name field value
func (o *SoundboardSoundResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SoundboardSoundResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SoundboardSoundResponse) SetName(v string) {
	o.Name = v
}

// GetSoundId returns the SoundId field value
func (o *SoundboardSoundResponse) GetSoundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SoundId
}

// GetSoundIdOk returns a tuple with the SoundId field value
// and a boolean to check if the value has been set.
func (o *SoundboardSoundResponse) GetSoundIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoundId, true
}

// SetSoundId sets field value
func (o *SoundboardSoundResponse) SetSoundId(v string) {
	o.SoundId = v
}

// GetVolume returns the Volume field value
func (o *SoundboardSoundResponse) GetVolume() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *SoundboardSoundResponse) GetVolumeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *SoundboardSoundResponse) SetVolume(v float64) {
	o.Volume = v
}

// GetEmojiId returns the EmojiId field value if set, zero value otherwise.
func (o *SoundboardSoundResponse) GetEmojiId() string {
	if o == nil || IsNil(o.EmojiId) {
		var ret string
		return ret
	}
	return *o.EmojiId
}

// GetEmojiIdOk returns a tuple with the EmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundboardSoundResponse) GetEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiId) {
		return nil, false
	}
	return o.EmojiId, true
}

// HasEmojiId returns a boolean if a field has been set.
func (o *SoundboardSoundResponse) HasEmojiId() bool {
	if o != nil && !IsNil(o.EmojiId) {
		return true
	}

	return false
}

// SetEmojiId gets a reference to the given string and assigns it to the EmojiId field.
func (o *SoundboardSoundResponse) SetEmojiId(v string) {
	o.EmojiId = &v
}


// GetEmojiName returns the EmojiName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoundboardSoundResponse) GetEmojiName() string {
	if o == nil || IsNil(o.EmojiName.Get()) {
		var ret string
		return ret
	}
	return *o.EmojiName.Get()
}

// GetEmojiNameOk returns a tuple with the EmojiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoundboardSoundResponse) GetEmojiNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiName.Get()) {
		return nil, false
	}
	return o.EmojiName.Get(), o.EmojiName.IsSet()
}

// HasEmojiName returns a boolean if a field has been set.
func (o *SoundboardSoundResponse) HasEmojiName() bool {
	if o != nil && o.EmojiName.IsSet() {
		return true
	}

	return false
}

// SetEmojiName gets a reference to the given NullableString and assigns it to the EmojiName field.
func (o *SoundboardSoundResponse) SetEmojiName(v string) {
	o.EmojiName.Set(&v)
}

// SetEmojiNameNil sets the value for EmojiName to be an explicit nil
func (o *SoundboardSoundResponse) SetEmojiNameNil() {
	o.EmojiName.Set(nil)
}

// UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
func (o *SoundboardSoundResponse) UnsetEmojiName() {
	o.EmojiName.Unset()
}

// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *SoundboardSoundResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundboardSoundResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *SoundboardSoundResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *SoundboardSoundResponse) SetGuildId(v string) {
	o.GuildId = &v
}


// GetAvailable returns the Available field value
func (o *SoundboardSoundResponse) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *SoundboardSoundResponse) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *SoundboardSoundResponse) SetAvailable(v bool) {
	o.Available = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoundboardSoundResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoundboardSoundResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User.Get()) {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *SoundboardSoundResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *SoundboardSoundResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *SoundboardSoundResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *SoundboardSoundResponse) UnsetUser() {
	o.User.Unset()
}

func (o SoundboardSoundResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoundboardSoundResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["sound_id"] = o.SoundId
	toSerialize["volume"] = o.Volume
	if !IsNil(o.EmojiId) {
		toSerialize["emoji_id"] = o.EmojiId
	}
	if o.EmojiName.IsSet() {
		toSerialize["emoji_name"] = o.EmojiName.Get()
	}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	toSerialize["available"] = o.Available
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

func (o *SoundboardSoundResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"sound_id",
		"volume",
		"available",
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

	varSoundboardSoundResponse := _SoundboardSoundResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSoundboardSoundResponse)

	if err != nil {
		return err
	}

	*o = SoundboardSoundResponse(varSoundboardSoundResponse)

	return err
}

type NullableSoundboardSoundResponse struct {
	value *SoundboardSoundResponse
	isSet bool
}

func (v NullableSoundboardSoundResponse) Get() *SoundboardSoundResponse {
	return v.value
}

func (v *NullableSoundboardSoundResponse) Set(val *SoundboardSoundResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSoundboardSoundResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSoundboardSoundResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoundboardSoundResponse(val *SoundboardSoundResponse) *NullableSoundboardSoundResponse {
	return &NullableSoundboardSoundResponse{value: val, isSet: true}
}

func (v NullableSoundboardSoundResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoundboardSoundResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


