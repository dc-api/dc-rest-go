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
	"bytes"
	"fmt"
)

// checks if the ContainerComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerComponentResponse{}

// ContainerComponentResponse struct for ContainerComponentResponse
type ContainerComponentResponse struct {
	Type int32 `json:"type"`
	Id int32 `json:"id"`
	AccentColor NullableInt32 `json:"accent_color,omitempty"`
	Components []ContainerComponentResponseComponentsInner `json:"components"`
	Spoiler bool `json:"spoiler"`
}

type _ContainerComponentResponse ContainerComponentResponse

// NewContainerComponentResponse instantiates a new ContainerComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerComponentResponse(type_ int32, id int32, components []ContainerComponentResponseComponentsInner, spoiler bool) *ContainerComponentResponse {
	this := ContainerComponentResponse{}
	this.Type = type_
	this.Id = id
	this.Components = components
	this.Spoiler = spoiler
	return &this
}

// NewContainerComponentResponseWithDefaults instantiates a new ContainerComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerComponentResponseWithDefaults() *ContainerComponentResponse {
	this := ContainerComponentResponse{}
	return &this
}

// GetType returns the Type field value
func (o *ContainerComponentResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContainerComponentResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContainerComponentResponse) SetType(v int32) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ContainerComponentResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContainerComponentResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContainerComponentResponse) SetId(v int32) {
	o.Id = v
}

// GetAccentColor returns the AccentColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerComponentResponse) GetAccentColor() int32 {
	if o == nil || IsNil(o.AccentColor.Get()) {
		var ret int32
		return ret
	}
	return *o.AccentColor.Get()
}

// GetAccentColorOk returns a tuple with the AccentColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerComponentResponse) GetAccentColorOk() (*int32, bool) {
	if o == nil || IsNil(o.AccentColor.Get()) {
		return nil, false
	}
	return o.AccentColor.Get(), o.AccentColor.IsSet()
}

// HasAccentColor returns a boolean if a field has been set.
func (o *ContainerComponentResponse) HasAccentColor() bool {
	if o != nil && o.AccentColor.IsSet() {
		return true
	}

	return false
}

// SetAccentColor gets a reference to the given NullableInt32 and assigns it to the AccentColor field.
func (o *ContainerComponentResponse) SetAccentColor(v int32) {
	o.AccentColor.Set(&v)
}

// SetAccentColorNil sets the value for AccentColor to be an explicit nil
func (o *ContainerComponentResponse) SetAccentColorNil() {
	o.AccentColor.Set(nil)
}

// UnsetAccentColor ensures that no value is present for AccentColor, not even an explicit nil
func (o *ContainerComponentResponse) UnsetAccentColor() {
	o.AccentColor.Unset()
}

// GetComponents returns the Components field value
func (o *ContainerComponentResponse) GetComponents() []ContainerComponentResponseComponentsInner {
	if o == nil {
		var ret []ContainerComponentResponseComponentsInner
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *ContainerComponentResponse) GetComponentsOk() ([]ContainerComponentResponseComponentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *ContainerComponentResponse) SetComponents(v []ContainerComponentResponseComponentsInner) {
	o.Components = v
}

// GetSpoiler returns the Spoiler field value
func (o *ContainerComponentResponse) GetSpoiler() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Spoiler
}

// GetSpoilerOk returns a tuple with the Spoiler field value
// and a boolean to check if the value has been set.
func (o *ContainerComponentResponse) GetSpoilerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spoiler, true
}

// SetSpoiler sets field value
func (o *ContainerComponentResponse) SetSpoiler(v bool) {
	o.Spoiler = v
}

func (o ContainerComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if o.AccentColor.IsSet() {
		toSerialize["accent_color"] = o.AccentColor.Get()
	}
	toSerialize["components"] = o.Components
	toSerialize["spoiler"] = o.Spoiler
	return toSerialize, nil
}

func (o *ContainerComponentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"components",
		"spoiler",
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

	varContainerComponentResponse := _ContainerComponentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerComponentResponse)

	if err != nil {
		return err
	}

	*o = ContainerComponentResponse(varContainerComponentResponse)

	return err
}

type NullableContainerComponentResponse struct {
	value *ContainerComponentResponse
	isSet bool
}

func (v NullableContainerComponentResponse) Get() *ContainerComponentResponse {
	return v.value
}

func (v *NullableContainerComponentResponse) Set(val *ContainerComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerComponentResponse(val *ContainerComponentResponse) *NullableContainerComponentResponse {
	return &NullableContainerComponentResponse{value: val, isSet: true}
}

func (v NullableContainerComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


