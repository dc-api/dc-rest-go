/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
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

// checks if the MessageReactionEmojiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReactionEmojiResponse{}

// MessageReactionEmojiResponse struct for MessageReactionEmojiResponse
type MessageReactionEmojiResponse struct {
	Id *GetEntitlementsSkuIdsParameterOneOfInner `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Animated *bool `json:"animated,omitempty"`
}

// NewMessageReactionEmojiResponse instantiates a new MessageReactionEmojiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReactionEmojiResponse() *MessageReactionEmojiResponse {
	this := MessageReactionEmojiResponse{}
	return &this
}

// NewMessageReactionEmojiResponseWithDefaults instantiates a new MessageReactionEmojiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReactionEmojiResponseWithDefaults() *MessageReactionEmojiResponse {
	this := MessageReactionEmojiResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MessageReactionEmojiResponse) GetId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.Id) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReactionEmojiResponse) GetIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MessageReactionEmojiResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the Id field.
func (o *MessageReactionEmojiResponse) SetId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.Id = &v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *MessageReactionEmojiResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReactionEmojiResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MessageReactionEmojiResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MessageReactionEmojiResponse) SetName(v string) {
	o.Name = &v
}


// GetAnimated returns the Animated field value if set, zero value otherwise.
func (o *MessageReactionEmojiResponse) GetAnimated() bool {
	if o == nil || IsNil(o.Animated) {
		var ret bool
		return ret
	}
	return *o.Animated
}

// GetAnimatedOk returns a tuple with the Animated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReactionEmojiResponse) GetAnimatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Animated) {
		return nil, false
	}
	return o.Animated, true
}

// HasAnimated returns a boolean if a field has been set.
func (o *MessageReactionEmojiResponse) HasAnimated() bool {
	if o != nil && !IsNil(o.Animated) {
		return true
	}

	return false
}

// SetAnimated gets a reference to the given bool and assigns it to the Animated field.
func (o *MessageReactionEmojiResponse) SetAnimated(v bool) {
	o.Animated = &v
}


func (o MessageReactionEmojiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReactionEmojiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Animated) {
		toSerialize["animated"] = o.Animated
	}
	return toSerialize, nil
}

type NullableMessageReactionEmojiResponse struct {
	value *MessageReactionEmojiResponse
	isSet bool
}

func (v NullableMessageReactionEmojiResponse) Get() *MessageReactionEmojiResponse {
	return v.value
}

func (v *NullableMessageReactionEmojiResponse) Set(val *MessageReactionEmojiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReactionEmojiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReactionEmojiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReactionEmojiResponse(val *MessageReactionEmojiResponse) *NullableMessageReactionEmojiResponse {
	return &NullableMessageReactionEmojiResponse{value: val, isSet: true}
}

func (v NullableMessageReactionEmojiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReactionEmojiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

