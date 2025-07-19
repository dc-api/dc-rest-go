/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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

// checks if the PollEmoji type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollEmoji{}

// PollEmoji struct for PollEmoji
type PollEmoji struct {
	Id *string `json:"id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name NullableString `json:"name,omitempty"`
	Animated NullableBool `json:"animated,omitempty"`
}

// NewPollEmoji instantiates a new PollEmoji object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollEmoji() *PollEmoji {
	this := PollEmoji{}
	return &this
}

// NewPollEmojiWithDefaults instantiates a new PollEmoji object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollEmojiWithDefaults() *PollEmoji {
	this := PollEmoji{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PollEmoji) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollEmoji) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PollEmoji) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PollEmoji) SetId(v string) {
	o.Id = &v
}


// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PollEmoji) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PollEmoji) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PollEmoji) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PollEmoji) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PollEmoji) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PollEmoji) UnsetName() {
	o.Name.Unset()
}

// GetAnimated returns the Animated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PollEmoji) GetAnimated() bool {
	if o == nil || IsNil(o.Animated.Get()) {
		var ret bool
		return ret
	}
	return *o.Animated.Get()
}

// GetAnimatedOk returns a tuple with the Animated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PollEmoji) GetAnimatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Animated.Get()) {
		return nil, false
	}
	return o.Animated.Get(), o.Animated.IsSet()
}

// HasAnimated returns a boolean if a field has been set.
func (o *PollEmoji) HasAnimated() bool {
	if o != nil && o.Animated.IsSet() {
		return true
	}

	return false
}

// SetAnimated gets a reference to the given NullableBool and assigns it to the Animated field.
func (o *PollEmoji) SetAnimated(v bool) {
	o.Animated.Set(&v)
}

// SetAnimatedNil sets the value for Animated to be an explicit nil
func (o *PollEmoji) SetAnimatedNil() {
	o.Animated.Set(nil)
}

// UnsetAnimated ensures that no value is present for Animated, not even an explicit nil
func (o *PollEmoji) UnsetAnimated() {
	o.Animated.Unset()
}

func (o PollEmoji) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollEmoji) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Animated.IsSet() {
		toSerialize["animated"] = o.Animated.Get()
	}
	return toSerialize, nil
}

type NullablePollEmoji struct {
	value *PollEmoji
	isSet bool
}

func (v NullablePollEmoji) Get() *PollEmoji {
	return v.value
}

func (v *NullablePollEmoji) Set(val *PollEmoji) {
	v.value = val
	v.isSet = true
}

func (v NullablePollEmoji) IsSet() bool {
	return v.isSet
}

func (v *NullablePollEmoji) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollEmoji(val *PollEmoji) *NullablePollEmoji {
	return &NullablePollEmoji{value: val, isSet: true}
}

func (v NullablePollEmoji) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollEmoji) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


