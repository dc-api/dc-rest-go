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
	"bytes"
	"fmt"
)

// checks if the ComponentEmojiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentEmojiResponse{}

// ComponentEmojiResponse struct for ComponentEmojiResponse
type ComponentEmojiResponse struct {
	Id *string `json:"id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Animated NullableBool `json:"animated,omitempty"`
}

type _ComponentEmojiResponse ComponentEmojiResponse

// NewComponentEmojiResponse instantiates a new ComponentEmojiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentEmojiResponse(name string) *ComponentEmojiResponse {
	this := ComponentEmojiResponse{}
	this.Name = name
	return &this
}

// NewComponentEmojiResponseWithDefaults instantiates a new ComponentEmojiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentEmojiResponseWithDefaults() *ComponentEmojiResponse {
	this := ComponentEmojiResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentEmojiResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentEmojiResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentEmojiResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComponentEmojiResponse) SetId(v string) {
	o.Id = &v
}


// GetName returns the Name field value
func (o *ComponentEmojiResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ComponentEmojiResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ComponentEmojiResponse) SetName(v string) {
	o.Name = v
}

// GetAnimated returns the Animated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentEmojiResponse) GetAnimated() bool {
	if o == nil || IsNil(o.Animated.Get()) {
		var ret bool
		return ret
	}
	return *o.Animated.Get()
}

// GetAnimatedOk returns a tuple with the Animated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentEmojiResponse) GetAnimatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Animated.Get()) {
		return nil, false
	}
	return o.Animated.Get(), o.Animated.IsSet()
}

// HasAnimated returns a boolean if a field has been set.
func (o *ComponentEmojiResponse) HasAnimated() bool {
	if o != nil && o.Animated.IsSet() {
		return true
	}

	return false
}

// SetAnimated gets a reference to the given NullableBool and assigns it to the Animated field.
func (o *ComponentEmojiResponse) SetAnimated(v bool) {
	o.Animated.Set(&v)
}

// SetAnimatedNil sets the value for Animated to be an explicit nil
func (o *ComponentEmojiResponse) SetAnimatedNil() {
	o.Animated.Set(nil)
}

// UnsetAnimated ensures that no value is present for Animated, not even an explicit nil
func (o *ComponentEmojiResponse) UnsetAnimated() {
	o.Animated.Unset()
}

func (o ComponentEmojiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentEmojiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.Animated.IsSet() {
		toSerialize["animated"] = o.Animated.Get()
	}
	return toSerialize, nil
}

func (o *ComponentEmojiResponse) UnmarshalJSON(data []byte) (err error) {
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

	varComponentEmojiResponse := _ComponentEmojiResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varComponentEmojiResponse)

	if err != nil {
		return err
	}

	*o = ComponentEmojiResponse(varComponentEmojiResponse)

	return err
}

type NullableComponentEmojiResponse struct {
	value *ComponentEmojiResponse
	isSet bool
}

func (v NullableComponentEmojiResponse) Get() *ComponentEmojiResponse {
	return v.value
}

func (v *NullableComponentEmojiResponse) Set(val *ComponentEmojiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentEmojiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentEmojiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentEmojiResponse(val *ComponentEmojiResponse) *NullableComponentEmojiResponse {
	return &NullableComponentEmojiResponse{value: val, isSet: true}
}

func (v NullableComponentEmojiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentEmojiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


