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
)

// checks if the PollMediaCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollMediaCreateRequest{}

// PollMediaCreateRequest struct for PollMediaCreateRequest
type PollMediaCreateRequest struct {
	Text NullableString `json:"text,omitempty"`
	Emoji NullablePollEmojiCreateRequest `json:"emoji,omitempty"`
}

// NewPollMediaCreateRequest instantiates a new PollMediaCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollMediaCreateRequest() *PollMediaCreateRequest {
	this := PollMediaCreateRequest{}
	return &this
}

// NewPollMediaCreateRequestWithDefaults instantiates a new PollMediaCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollMediaCreateRequestWithDefaults() *PollMediaCreateRequest {
	this := PollMediaCreateRequest{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PollMediaCreateRequest) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PollMediaCreateRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text.Get()) {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *PollMediaCreateRequest) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *PollMediaCreateRequest) SetText(v string) {
	o.Text.Set(&v)
}

// SetTextNil sets the value for Text to be an explicit nil
func (o *PollMediaCreateRequest) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *PollMediaCreateRequest) UnsetText() {
	o.Text.Unset()
}

// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PollMediaCreateRequest) GetEmoji() PollEmojiCreateRequest {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret PollEmojiCreateRequest
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PollMediaCreateRequest) GetEmojiOk() (*PollEmojiCreateRequest, bool) {
	if o == nil || IsNil(o.Emoji.Get()) {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *PollMediaCreateRequest) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullablePollEmojiCreateRequest and assigns it to the Emoji field.
func (o *PollMediaCreateRequest) SetEmoji(v PollEmojiCreateRequest) {
	o.Emoji.Set(&v)
}

// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *PollMediaCreateRequest) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *PollMediaCreateRequest) UnsetEmoji() {
	o.Emoji.Unset()
}

func (o PollMediaCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollMediaCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	return toSerialize, nil
}

type NullablePollMediaCreateRequest struct {
	value *PollMediaCreateRequest
	isSet bool
}

func (v NullablePollMediaCreateRequest) Get() *PollMediaCreateRequest {
	return v.value
}

func (v *NullablePollMediaCreateRequest) Set(val *PollMediaCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePollMediaCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePollMediaCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollMediaCreateRequest(val *PollMediaCreateRequest) *NullablePollMediaCreateRequest {
	return &NullablePollMediaCreateRequest{value: val, isSet: true}
}

func (v NullablePollMediaCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollMediaCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


