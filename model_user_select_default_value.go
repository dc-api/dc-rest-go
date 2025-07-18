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
	"bytes"
	"fmt"
)

// checks if the UserSelectDefaultValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSelectDefaultValue{}

// UserSelectDefaultValue struct for UserSelectDefaultValue
type UserSelectDefaultValue struct {
	Type NullableString `json:"type"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _UserSelectDefaultValue UserSelectDefaultValue

// NewUserSelectDefaultValue instantiates a new UserSelectDefaultValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSelectDefaultValue(type_ NullableString, id string) *UserSelectDefaultValue {
	this := UserSelectDefaultValue{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewUserSelectDefaultValueWithDefaults instantiates a new UserSelectDefaultValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSelectDefaultValueWithDefaults() *UserSelectDefaultValue {
	this := UserSelectDefaultValue{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserSelectDefaultValue) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSelectDefaultValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *UserSelectDefaultValue) SetType(v string) {
	o.Type.Set(&v)
}

// GetId returns the Id field value
func (o *UserSelectDefaultValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSelectDefaultValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSelectDefaultValue) SetId(v string) {
	o.Id = v
}

func (o UserSelectDefaultValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSelectDefaultValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type.Get()
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UserSelectDefaultValue) UnmarshalJSON(data []byte) (err error) {
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

	varUserSelectDefaultValue := _UserSelectDefaultValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSelectDefaultValue)

	if err != nil {
		return err
	}

	*o = UserSelectDefaultValue(varUserSelectDefaultValue)

	return err
}

type NullableUserSelectDefaultValue struct {
	value *UserSelectDefaultValue
	isSet bool
}

func (v NullableUserSelectDefaultValue) Get() *UserSelectDefaultValue {
	return v.value
}

func (v *NullableUserSelectDefaultValue) Set(val *UserSelectDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSelectDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSelectDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSelectDefaultValue(val *UserSelectDefaultValue) *NullableUserSelectDefaultValue {
	return &NullableUserSelectDefaultValue{value: val, isSet: true}
}

func (v NullableUserSelectDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSelectDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


