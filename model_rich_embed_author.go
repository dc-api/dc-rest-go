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

// checks if the RichEmbedAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RichEmbedAuthor{}

// RichEmbedAuthor struct for RichEmbedAuthor
type RichEmbedAuthor struct {
	Name NullableString `json:"name,omitempty"`
	Url NullableString `json:"url,omitempty"`
	IconUrl NullableString `json:"icon_url,omitempty"`
}

// NewRichEmbedAuthor instantiates a new RichEmbedAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRichEmbedAuthor() *RichEmbedAuthor {
	this := RichEmbedAuthor{}
	return &this
}

// NewRichEmbedAuthorWithDefaults instantiates a new RichEmbedAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRichEmbedAuthorWithDefaults() *RichEmbedAuthor {
	this := RichEmbedAuthor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedAuthor) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedAuthor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RichEmbedAuthor) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RichEmbedAuthor) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *RichEmbedAuthor) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RichEmbedAuthor) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedAuthor) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedAuthor) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url.Get()) {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *RichEmbedAuthor) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *RichEmbedAuthor) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *RichEmbedAuthor) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *RichEmbedAuthor) UnsetUrl() {
	o.Url.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedAuthor) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedAuthor) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl.Get()) {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *RichEmbedAuthor) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *RichEmbedAuthor) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *RichEmbedAuthor) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *RichEmbedAuthor) UnsetIconUrl() {
	o.IconUrl.Unset()
}

func (o RichEmbedAuthor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RichEmbedAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	return toSerialize, nil
}

type NullableRichEmbedAuthor struct {
	value *RichEmbedAuthor
	isSet bool
}

func (v NullableRichEmbedAuthor) Get() *RichEmbedAuthor {
	return v.value
}

func (v *NullableRichEmbedAuthor) Set(val *RichEmbedAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableRichEmbedAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableRichEmbedAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRichEmbedAuthor(val *RichEmbedAuthor) *NullableRichEmbedAuthor {
	return &NullableRichEmbedAuthor{value: val, isSet: true}
}

func (v NullableRichEmbedAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRichEmbedAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


