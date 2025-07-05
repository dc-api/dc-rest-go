/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
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

// checks if the GuildTemplateChannelTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildTemplateChannelTags{}

// GuildTemplateChannelTags struct for GuildTemplateChannelTags
type GuildTemplateChannelTags struct {
	Name string `json:"name"`
	EmojiId *string `json:"emoji_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EmojiName NullableString `json:"emoji_name,omitempty"`
	Moderated NullableBool `json:"moderated,omitempty"`
}

type _GuildTemplateChannelTags GuildTemplateChannelTags

// NewGuildTemplateChannelTags instantiates a new GuildTemplateChannelTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildTemplateChannelTags(name string) *GuildTemplateChannelTags {
	this := GuildTemplateChannelTags{}
	this.Name = name
	return &this
}

// NewGuildTemplateChannelTagsWithDefaults instantiates a new GuildTemplateChannelTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildTemplateChannelTagsWithDefaults() *GuildTemplateChannelTags {
	this := GuildTemplateChannelTags{}
	return &this
}

// GetName returns the Name field value
func (o *GuildTemplateChannelTags) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelTags) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildTemplateChannelTags) SetName(v string) {
	o.Name = v
}

// GetEmojiId returns the EmojiId field value if set, zero value otherwise.
func (o *GuildTemplateChannelTags) GetEmojiId() string {
	if o == nil || IsNil(o.EmojiId) {
		var ret string
		return ret
	}
	return *o.EmojiId
}

// GetEmojiIdOk returns a tuple with the EmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelTags) GetEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiId) {
		return nil, false
	}
	return o.EmojiId, true
}

// HasEmojiId returns a boolean if a field has been set.
func (o *GuildTemplateChannelTags) HasEmojiId() bool {
	if o != nil && !IsNil(o.EmojiId) {
		return true
	}

	return false
}

// SetEmojiId gets a reference to the given string and assigns it to the EmojiId field.
func (o *GuildTemplateChannelTags) SetEmojiId(v string) {
	o.EmojiId = &v
}


// GetEmojiName returns the EmojiName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelTags) GetEmojiName() string {
	if o == nil || IsNil(o.EmojiName.Get()) {
		var ret string
		return ret
	}
	return *o.EmojiName.Get()
}

// GetEmojiNameOk returns a tuple with the EmojiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelTags) GetEmojiNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiName.Get()) {
		return nil, false
	}
	return o.EmojiName.Get(), o.EmojiName.IsSet()
}

// HasEmojiName returns a boolean if a field has been set.
func (o *GuildTemplateChannelTags) HasEmojiName() bool {
	if o != nil && o.EmojiName.IsSet() {
		return true
	}

	return false
}

// SetEmojiName gets a reference to the given NullableString and assigns it to the EmojiName field.
func (o *GuildTemplateChannelTags) SetEmojiName(v string) {
	o.EmojiName.Set(&v)
}

// SetEmojiNameNil sets the value for EmojiName to be an explicit nil
func (o *GuildTemplateChannelTags) SetEmojiNameNil() {
	o.EmojiName.Set(nil)
}

// UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
func (o *GuildTemplateChannelTags) UnsetEmojiName() {
	o.EmojiName.Unset()
}

// GetModerated returns the Moderated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelTags) GetModerated() bool {
	if o == nil || IsNil(o.Moderated.Get()) {
		var ret bool
		return ret
	}
	return *o.Moderated.Get()
}

// GetModeratedOk returns a tuple with the Moderated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelTags) GetModeratedOk() (*bool, bool) {
	if o == nil || IsNil(o.Moderated.Get()) {
		return nil, false
	}
	return o.Moderated.Get(), o.Moderated.IsSet()
}

// HasModerated returns a boolean if a field has been set.
func (o *GuildTemplateChannelTags) HasModerated() bool {
	if o != nil && o.Moderated.IsSet() {
		return true
	}

	return false
}

// SetModerated gets a reference to the given NullableBool and assigns it to the Moderated field.
func (o *GuildTemplateChannelTags) SetModerated(v bool) {
	o.Moderated.Set(&v)
}

// SetModeratedNil sets the value for Moderated to be an explicit nil
func (o *GuildTemplateChannelTags) SetModeratedNil() {
	o.Moderated.Set(nil)
}

// UnsetModerated ensures that no value is present for Moderated, not even an explicit nil
func (o *GuildTemplateChannelTags) UnsetModerated() {
	o.Moderated.Unset()
}

func (o GuildTemplateChannelTags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildTemplateChannelTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.EmojiId) {
		toSerialize["emoji_id"] = o.EmojiId
	}
	if o.EmojiName.IsSet() {
		toSerialize["emoji_name"] = o.EmojiName.Get()
	}
	if o.Moderated.IsSet() {
		toSerialize["moderated"] = o.Moderated.Get()
	}
	return toSerialize, nil
}

func (o *GuildTemplateChannelTags) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varGuildTemplateChannelTags := _GuildTemplateChannelTags{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildTemplateChannelTags)

	if err != nil {
		return err
	}

	*o = GuildTemplateChannelTags(varGuildTemplateChannelTags)

	return err
}

type NullableGuildTemplateChannelTags struct {
	value *GuildTemplateChannelTags
	isSet bool
}

func (v NullableGuildTemplateChannelTags) Get() *GuildTemplateChannelTags {
	return v.value
}

func (v *NullableGuildTemplateChannelTags) Set(val *GuildTemplateChannelTags) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildTemplateChannelTags) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildTemplateChannelTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildTemplateChannelTags(val *GuildTemplateChannelTags) *NullableGuildTemplateChannelTags {
	return &NullableGuildTemplateChannelTags{value: val, isSet: true}
}

func (v NullableGuildTemplateChannelTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildTemplateChannelTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


