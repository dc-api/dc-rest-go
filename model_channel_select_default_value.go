/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// checks if the ChannelSelectDefaultValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelSelectDefaultValue{}

// ChannelSelectDefaultValue struct for ChannelSelectDefaultValue
type ChannelSelectDefaultValue struct {
	Type NullableString `json:"type"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _ChannelSelectDefaultValue ChannelSelectDefaultValue

// NewChannelSelectDefaultValue instantiates a new ChannelSelectDefaultValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelSelectDefaultValue(type_ NullableString, id string) *ChannelSelectDefaultValue {
	this := ChannelSelectDefaultValue{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewChannelSelectDefaultValueWithDefaults instantiates a new ChannelSelectDefaultValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelSelectDefaultValueWithDefaults() *ChannelSelectDefaultValue {
	this := ChannelSelectDefaultValue{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChannelSelectDefaultValue) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelSelectDefaultValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *ChannelSelectDefaultValue) SetType(v string) {
	o.Type.Set(&v)
}

// GetId returns the Id field value
func (o *ChannelSelectDefaultValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChannelSelectDefaultValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChannelSelectDefaultValue) SetId(v string) {
	o.Id = v
}

func (o ChannelSelectDefaultValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelSelectDefaultValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type.Get()
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ChannelSelectDefaultValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varChannelSelectDefaultValue := _ChannelSelectDefaultValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelSelectDefaultValue)

	if err != nil {
		return err
	}

	*o = ChannelSelectDefaultValue(varChannelSelectDefaultValue)

	return err
}

type NullableChannelSelectDefaultValue struct {
	value *ChannelSelectDefaultValue
	isSet bool
}

func (v NullableChannelSelectDefaultValue) Get() *ChannelSelectDefaultValue {
	return v.value
}

func (v *NullableChannelSelectDefaultValue) Set(val *ChannelSelectDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelSelectDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelSelectDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelSelectDefaultValue(val *ChannelSelectDefaultValue) *NullableChannelSelectDefaultValue {
	return &NullableChannelSelectDefaultValue{value: val, isSet: true}
}

func (v NullableChannelSelectDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelSelectDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


