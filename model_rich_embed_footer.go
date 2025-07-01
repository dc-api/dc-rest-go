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
)

// checks if the RichEmbedFooter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RichEmbedFooter{}

// RichEmbedFooter struct for RichEmbedFooter
type RichEmbedFooter struct {
	Text NullableString `json:"text,omitempty"`
	IconUrl NullableString `json:"icon_url,omitempty"`
}

// NewRichEmbedFooter instantiates a new RichEmbedFooter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRichEmbedFooter() *RichEmbedFooter {
	this := RichEmbedFooter{}
	return &this
}

// NewRichEmbedFooterWithDefaults instantiates a new RichEmbedFooter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRichEmbedFooterWithDefaults() *RichEmbedFooter {
	this := RichEmbedFooter{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedFooter) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedFooter) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text.Get()) {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *RichEmbedFooter) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *RichEmbedFooter) SetText(v string) {
	o.Text.Set(&v)
}

// SetTextNil sets the value for Text to be an explicit nil
func (o *RichEmbedFooter) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *RichEmbedFooter) UnsetText() {
	o.Text.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedFooter) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedFooter) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl.Get()) {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *RichEmbedFooter) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *RichEmbedFooter) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *RichEmbedFooter) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *RichEmbedFooter) UnsetIconUrl() {
	o.IconUrl.Unset()
}

func (o RichEmbedFooter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RichEmbedFooter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	return toSerialize, nil
}

type NullableRichEmbedFooter struct {
	value *RichEmbedFooter
	isSet bool
}

func (v NullableRichEmbedFooter) Get() *RichEmbedFooter {
	return v.value
}

func (v *NullableRichEmbedFooter) Set(val *RichEmbedFooter) {
	v.value = val
	v.isSet = true
}

func (v NullableRichEmbedFooter) IsSet() bool {
	return v.isSet
}

func (v *NullableRichEmbedFooter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRichEmbedFooter(val *RichEmbedFooter) *NullableRichEmbedFooter {
	return &NullableRichEmbedFooter{value: val, isSet: true}
}

func (v NullableRichEmbedFooter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRichEmbedFooter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


