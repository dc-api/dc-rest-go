/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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
)

// checks if the PollMedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollMedia{}

// PollMedia struct for PollMedia
type PollMedia struct {
	Text NullableString `json:"text,omitempty"`
	Emoji NullablePollEmoji `json:"emoji,omitempty"`
}

// NewPollMedia instantiates a new PollMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollMedia() *PollMedia {
	this := PollMedia{}
	return &this
}

// NewPollMediaWithDefaults instantiates a new PollMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollMediaWithDefaults() *PollMedia {
	this := PollMedia{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PollMedia) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PollMedia) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text.Get()) {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *PollMedia) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *PollMedia) SetText(v string) {
	o.Text.Set(&v)
}

// SetTextNil sets the value for Text to be an explicit nil
func (o *PollMedia) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *PollMedia) UnsetText() {
	o.Text.Unset()
}

// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PollMedia) GetEmoji() PollEmoji {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret PollEmoji
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PollMedia) GetEmojiOk() (*PollEmoji, bool) {
	if o == nil || IsNil(o.Emoji.Get()) {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *PollMedia) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullablePollEmoji and assigns it to the Emoji field.
func (o *PollMedia) SetEmoji(v PollEmoji) {
	o.Emoji.Set(&v)
}

// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *PollMedia) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *PollMedia) UnsetEmoji() {
	o.Emoji.Unset()
}

func (o PollMedia) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollMedia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	return toSerialize, nil
}

type NullablePollMedia struct {
	value *PollMedia
	isSet bool
}

func (v NullablePollMedia) Get() *PollMedia {
	return v.value
}

func (v *NullablePollMedia) Set(val *PollMedia) {
	v.value = val
	v.isSet = true
}

func (v NullablePollMedia) IsSet() bool {
	return v.isSet
}

func (v *NullablePollMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollMedia(val *PollMedia) *NullablePollMedia {
	return &NullablePollMedia{value: val, isSet: true}
}

func (v NullablePollMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


