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

// checks if the UserSelectComponentForMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSelectComponentForMessageRequest{}

// UserSelectComponentForMessageRequest struct for UserSelectComponentForMessageRequest
type UserSelectComponentForMessageRequest struct {
	Type int32 `json:"type"`
	CustomId string `json:"custom_id"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	MinValues NullableInt32 `json:"min_values,omitempty"`
	MaxValues NullableInt32 `json:"max_values,omitempty"`
	Disabled NullableBool `json:"disabled,omitempty"`
	DefaultValues []UserSelectDefaultValue `json:"default_values,omitempty"`
}

type _UserSelectComponentForMessageRequest UserSelectComponentForMessageRequest

// NewUserSelectComponentForMessageRequest instantiates a new UserSelectComponentForMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSelectComponentForMessageRequest(type_ int32, customId string) *UserSelectComponentForMessageRequest {
	this := UserSelectComponentForMessageRequest{}
	this.Type = type_
	this.CustomId = customId
	return &this
}

// NewUserSelectComponentForMessageRequestWithDefaults instantiates a new UserSelectComponentForMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSelectComponentForMessageRequestWithDefaults() *UserSelectComponentForMessageRequest {
	this := UserSelectComponentForMessageRequest{}
	return &this
}

// GetType returns the Type field value
func (o *UserSelectComponentForMessageRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserSelectComponentForMessageRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserSelectComponentForMessageRequest) SetType(v int32) {
	o.Type = v
}

// GetCustomId returns the CustomId field value
func (o *UserSelectComponentForMessageRequest) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *UserSelectComponentForMessageRequest) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *UserSelectComponentForMessageRequest) SetCustomId(v string) {
	o.CustomId = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSelectComponentForMessageRequest) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSelectComponentForMessageRequest) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder.Get()) {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *UserSelectComponentForMessageRequest) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *UserSelectComponentForMessageRequest) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}

// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *UserSelectComponentForMessageRequest) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *UserSelectComponentForMessageRequest) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetMinValues returns the MinValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSelectComponentForMessageRequest) GetMinValues() int32 {
	if o == nil || IsNil(o.MinValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MinValues.Get()
}

// GetMinValuesOk returns a tuple with the MinValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSelectComponentForMessageRequest) GetMinValuesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinValues.Get()) {
		return nil, false
	}
	return o.MinValues.Get(), o.MinValues.IsSet()
}

// HasMinValues returns a boolean if a field has been set.
func (o *UserSelectComponentForMessageRequest) HasMinValues() bool {
	if o != nil && o.MinValues.IsSet() {
		return true
	}

	return false
}

// SetMinValues gets a reference to the given NullableInt32 and assigns it to the MinValues field.
func (o *UserSelectComponentForMessageRequest) SetMinValues(v int32) {
	o.MinValues.Set(&v)
}

// SetMinValuesNil sets the value for MinValues to be an explicit nil
func (o *UserSelectComponentForMessageRequest) SetMinValuesNil() {
	o.MinValues.Set(nil)
}

// UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
func (o *UserSelectComponentForMessageRequest) UnsetMinValues() {
	o.MinValues.Unset()
}

// GetMaxValues returns the MaxValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSelectComponentForMessageRequest) GetMaxValues() int32 {
	if o == nil || IsNil(o.MaxValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxValues.Get()
}

// GetMaxValuesOk returns a tuple with the MaxValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSelectComponentForMessageRequest) GetMaxValuesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxValues.Get()) {
		return nil, false
	}
	return o.MaxValues.Get(), o.MaxValues.IsSet()
}

// HasMaxValues returns a boolean if a field has been set.
func (o *UserSelectComponentForMessageRequest) HasMaxValues() bool {
	if o != nil && o.MaxValues.IsSet() {
		return true
	}

	return false
}

// SetMaxValues gets a reference to the given NullableInt32 and assigns it to the MaxValues field.
func (o *UserSelectComponentForMessageRequest) SetMaxValues(v int32) {
	o.MaxValues.Set(&v)
}

// SetMaxValuesNil sets the value for MaxValues to be an explicit nil
func (o *UserSelectComponentForMessageRequest) SetMaxValuesNil() {
	o.MaxValues.Set(nil)
}

// UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
func (o *UserSelectComponentForMessageRequest) UnsetMaxValues() {
	o.MaxValues.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSelectComponentForMessageRequest) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSelectComponentForMessageRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled.Get()) {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *UserSelectComponentForMessageRequest) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *UserSelectComponentForMessageRequest) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}

// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *UserSelectComponentForMessageRequest) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *UserSelectComponentForMessageRequest) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSelectComponentForMessageRequest) GetDefaultValues() []UserSelectDefaultValue {
	if o == nil {
		var ret []UserSelectDefaultValue
		return ret
	}
	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSelectComponentForMessageRequest) GetDefaultValuesOk() ([]UserSelectDefaultValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *UserSelectComponentForMessageRequest) HasDefaultValues() bool {
	if o != nil && !IsNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given []UserSelectDefaultValue and assigns it to the DefaultValues field.
func (o *UserSelectComponentForMessageRequest) SetDefaultValues(v []UserSelectDefaultValue) {
	o.DefaultValues = v
}


func (o UserSelectComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSelectComponentForMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["custom_id"] = o.CustomId
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if o.MinValues.IsSet() {
		toSerialize["min_values"] = o.MinValues.Get()
	}
	if o.MaxValues.IsSet() {
		toSerialize["max_values"] = o.MaxValues.Get()
	}
	if o.Disabled.IsSet() {
		toSerialize["disabled"] = o.Disabled.Get()
	}
	if o.DefaultValues != nil {
		toSerialize["default_values"] = o.DefaultValues
	}
	return toSerialize, nil
}

func (o *UserSelectComponentForMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"custom_id",
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

	varUserSelectComponentForMessageRequest := _UserSelectComponentForMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSelectComponentForMessageRequest)

	if err != nil {
		return err
	}

	*o = UserSelectComponentForMessageRequest(varUserSelectComponentForMessageRequest)

	return err
}

type NullableUserSelectComponentForMessageRequest struct {
	value *UserSelectComponentForMessageRequest
	isSet bool
}

func (v NullableUserSelectComponentForMessageRequest) Get() *UserSelectComponentForMessageRequest {
	return v.value
}

func (v *NullableUserSelectComponentForMessageRequest) Set(val *UserSelectComponentForMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSelectComponentForMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSelectComponentForMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSelectComponentForMessageRequest(val *UserSelectComponentForMessageRequest) *NullableUserSelectComponentForMessageRequest {
	return &NullableUserSelectComponentForMessageRequest{value: val, isSet: true}
}

func (v NullableUserSelectComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSelectComponentForMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


