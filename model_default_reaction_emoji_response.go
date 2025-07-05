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

// checks if the DefaultReactionEmojiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultReactionEmojiResponse{}

// DefaultReactionEmojiResponse struct for DefaultReactionEmojiResponse
type DefaultReactionEmojiResponse struct {
	EmojiId *string `json:"emoji_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EmojiName NullableString `json:"emoji_name,omitempty"`
}

// NewDefaultReactionEmojiResponse instantiates a new DefaultReactionEmojiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultReactionEmojiResponse() *DefaultReactionEmojiResponse {
	this := DefaultReactionEmojiResponse{}
	return &this
}

// NewDefaultReactionEmojiResponseWithDefaults instantiates a new DefaultReactionEmojiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultReactionEmojiResponseWithDefaults() *DefaultReactionEmojiResponse {
	this := DefaultReactionEmojiResponse{}
	return &this
}

// GetEmojiId returns the EmojiId field value if set, zero value otherwise.
func (o *DefaultReactionEmojiResponse) GetEmojiId() string {
	if o == nil || IsNil(o.EmojiId) {
		var ret string
		return ret
	}
	return *o.EmojiId
}

// GetEmojiIdOk returns a tuple with the EmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultReactionEmojiResponse) GetEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiId) {
		return nil, false
	}
	return o.EmojiId, true
}

// HasEmojiId returns a boolean if a field has been set.
func (o *DefaultReactionEmojiResponse) HasEmojiId() bool {
	if o != nil && !IsNil(o.EmojiId) {
		return true
	}

	return false
}

// SetEmojiId gets a reference to the given string and assigns it to the EmojiId field.
func (o *DefaultReactionEmojiResponse) SetEmojiId(v string) {
	o.EmojiId = &v
}


// GetEmojiName returns the EmojiName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultReactionEmojiResponse) GetEmojiName() string {
	if o == nil || IsNil(o.EmojiName.Get()) {
		var ret string
		return ret
	}
	return *o.EmojiName.Get()
}

// GetEmojiNameOk returns a tuple with the EmojiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultReactionEmojiResponse) GetEmojiNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiName.Get()) {
		return nil, false
	}
	return o.EmojiName.Get(), o.EmojiName.IsSet()
}

// HasEmojiName returns a boolean if a field has been set.
func (o *DefaultReactionEmojiResponse) HasEmojiName() bool {
	if o != nil && o.EmojiName.IsSet() {
		return true
	}

	return false
}

// SetEmojiName gets a reference to the given NullableString and assigns it to the EmojiName field.
func (o *DefaultReactionEmojiResponse) SetEmojiName(v string) {
	o.EmojiName.Set(&v)
}

// SetEmojiNameNil sets the value for EmojiName to be an explicit nil
func (o *DefaultReactionEmojiResponse) SetEmojiNameNil() {
	o.EmojiName.Set(nil)
}

// UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
func (o *DefaultReactionEmojiResponse) UnsetEmojiName() {
	o.EmojiName.Unset()
}

func (o DefaultReactionEmojiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultReactionEmojiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmojiId) {
		toSerialize["emoji_id"] = o.EmojiId
	}
	if o.EmojiName.IsSet() {
		toSerialize["emoji_name"] = o.EmojiName.Get()
	}
	return toSerialize, nil
}

type NullableDefaultReactionEmojiResponse struct {
	value *DefaultReactionEmojiResponse
	isSet bool
}

func (v NullableDefaultReactionEmojiResponse) Get() *DefaultReactionEmojiResponse {
	return v.value
}

func (v *NullableDefaultReactionEmojiResponse) Set(val *DefaultReactionEmojiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultReactionEmojiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultReactionEmojiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultReactionEmojiResponse(val *DefaultReactionEmojiResponse) *NullableDefaultReactionEmojiResponse {
	return &NullableDefaultReactionEmojiResponse{value: val, isSet: true}
}

func (v NullableDefaultReactionEmojiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultReactionEmojiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


