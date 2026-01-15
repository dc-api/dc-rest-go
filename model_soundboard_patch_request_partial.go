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
)

// checks if the SoundboardPatchRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoundboardPatchRequestPartial{}

// SoundboardPatchRequestPartial struct for SoundboardPatchRequestPartial
type SoundboardPatchRequestPartial struct {
	Name *string `json:"name,omitempty"`
	Volume *float64 `json:"volume,omitempty"`
	EmojiId *GetEntitlementsSkuIdsParameterOneOfInner `json:"emoji_id,omitempty"`
	EmojiName *string `json:"emoji_name,omitempty"`
}

// NewSoundboardPatchRequestPartial instantiates a new SoundboardPatchRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoundboardPatchRequestPartial() *SoundboardPatchRequestPartial {
	this := SoundboardPatchRequestPartial{}
	return &this
}

// NewSoundboardPatchRequestPartialWithDefaults instantiates a new SoundboardPatchRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoundboardPatchRequestPartialWithDefaults() *SoundboardPatchRequestPartial {
	this := SoundboardPatchRequestPartial{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoundboardPatchRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundboardPatchRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoundboardPatchRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoundboardPatchRequestPartial) SetName(v string) {
	o.Name = &v
}


// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *SoundboardPatchRequestPartial) GetVolume() float64 {
	if o == nil || IsNil(o.Volume) {
		var ret float64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundboardPatchRequestPartial) GetVolumeOk() (*float64, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *SoundboardPatchRequestPartial) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float64 and assigns it to the Volume field.
func (o *SoundboardPatchRequestPartial) SetVolume(v float64) {
	o.Volume = &v
}


// GetEmojiId returns the EmojiId field value if set, zero value otherwise.
func (o *SoundboardPatchRequestPartial) GetEmojiId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.EmojiId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.EmojiId
}

// GetEmojiIdOk returns a tuple with the EmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundboardPatchRequestPartial) GetEmojiIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.EmojiId) {
		return nil, false
	}
	return o.EmojiId, true
}

// HasEmojiId returns a boolean if a field has been set.
func (o *SoundboardPatchRequestPartial) HasEmojiId() bool {
	if o != nil && !IsNil(o.EmojiId) {
		return true
	}

	return false
}

// SetEmojiId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the EmojiId field.
func (o *SoundboardPatchRequestPartial) SetEmojiId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.EmojiId = &v
}


// GetEmojiName returns the EmojiName field value if set, zero value otherwise.
func (o *SoundboardPatchRequestPartial) GetEmojiName() string {
	if o == nil || IsNil(o.EmojiName) {
		var ret string
		return ret
	}
	return *o.EmojiName
}

// GetEmojiNameOk returns a tuple with the EmojiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundboardPatchRequestPartial) GetEmojiNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiName) {
		return nil, false
	}
	return o.EmojiName, true
}

// HasEmojiName returns a boolean if a field has been set.
func (o *SoundboardPatchRequestPartial) HasEmojiName() bool {
	if o != nil && !IsNil(o.EmojiName) {
		return true
	}

	return false
}

// SetEmojiName gets a reference to the given string and assigns it to the EmojiName field.
func (o *SoundboardPatchRequestPartial) SetEmojiName(v string) {
	o.EmojiName = &v
}


func (o SoundboardPatchRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoundboardPatchRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.EmojiId) {
		toSerialize["emoji_id"] = o.EmojiId
	}
	if !IsNil(o.EmojiName) {
		toSerialize["emoji_name"] = o.EmojiName
	}
	return toSerialize, nil
}

type NullableSoundboardPatchRequestPartial struct {
	value *SoundboardPatchRequestPartial
	isSet bool
}

func (v NullableSoundboardPatchRequestPartial) Get() *SoundboardPatchRequestPartial {
	return v.value
}

func (v *NullableSoundboardPatchRequestPartial) Set(val *SoundboardPatchRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableSoundboardPatchRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableSoundboardPatchRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoundboardPatchRequestPartial(val *SoundboardPatchRequestPartial) *NullableSoundboardPatchRequestPartial {
	return &NullableSoundboardPatchRequestPartial{value: val, isSet: true}
}

func (v NullableSoundboardPatchRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoundboardPatchRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

