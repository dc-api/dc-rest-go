/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-10-14T14:27:02.905060650Z[Etc/UTC]
 *    * - **Generator Version**: 7.16.0
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

// checks if the UserSelectComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSelectComponentResponse{}

// UserSelectComponentResponse struct for UserSelectComponentResponse
type UserSelectComponentResponse struct {
	Type int32 `json:"type"`
	Id int32 `json:"id"`
	CustomId string `json:"custom_id"`
	Placeholder *string `json:"placeholder,omitempty"`
	MinValues NullableInt32 `json:"min_values,omitempty"`
	MaxValues NullableInt32 `json:"max_values,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	DefaultValues []UserSelectDefaultValueResponse `json:"default_values,omitempty"`
}

type _UserSelectComponentResponse UserSelectComponentResponse

// NewUserSelectComponentResponse instantiates a new UserSelectComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSelectComponentResponse(type_ int32, id int32, customId string) *UserSelectComponentResponse {
	this := UserSelectComponentResponse{}
	this.Type = type_
	this.Id = id
	this.CustomId = customId
	return &this
}

// NewUserSelectComponentResponseWithDefaults instantiates a new UserSelectComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSelectComponentResponseWithDefaults() *UserSelectComponentResponse {
	this := UserSelectComponentResponse{}
	return &this
}

// GetType returns the Type field value
func (o *UserSelectComponentResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserSelectComponentResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserSelectComponentResponse) SetType(v int32) {
	o.Type = v
}

// GetId returns the Id field value
func (o *UserSelectComponentResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSelectComponentResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSelectComponentResponse) SetId(v int32) {
	o.Id = v
}

// GetCustomId returns the CustomId field value
func (o *UserSelectComponentResponse) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *UserSelectComponentResponse) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *UserSelectComponentResponse) SetCustomId(v string) {
	o.CustomId = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *UserSelectComponentResponse) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSelectComponentResponse) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *UserSelectComponentResponse) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *UserSelectComponentResponse) SetPlaceholder(v string) {
	o.Placeholder = &v
}


// GetMinValues returns the MinValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSelectComponentResponse) GetMinValues() int32 {
	if o == nil || IsNil(o.MinValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MinValues.Get()
}

// GetMinValuesOk returns a tuple with the MinValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSelectComponentResponse) GetMinValuesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinValues.Get()) {
		return nil, false
	}
	return o.MinValues.Get(), o.MinValues.IsSet()
}

// HasMinValues returns a boolean if a field has been set.
func (o *UserSelectComponentResponse) HasMinValues() bool {
	if o != nil && o.MinValues.IsSet() {
		return true
	}

	return false
}

// SetMinValues gets a reference to the given NullableInt32 and assigns it to the MinValues field.
func (o *UserSelectComponentResponse) SetMinValues(v int32) {
	o.MinValues.Set(&v)
}

// SetMinValuesNil sets the value for MinValues to be an explicit nil
func (o *UserSelectComponentResponse) SetMinValuesNil() {
	o.MinValues.Set(nil)
}

// UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
func (o *UserSelectComponentResponse) UnsetMinValues() {
	o.MinValues.Unset()
}

// GetMaxValues returns the MaxValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSelectComponentResponse) GetMaxValues() int32 {
	if o == nil || IsNil(o.MaxValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxValues.Get()
}

// GetMaxValuesOk returns a tuple with the MaxValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSelectComponentResponse) GetMaxValuesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxValues.Get()) {
		return nil, false
	}
	return o.MaxValues.Get(), o.MaxValues.IsSet()
}

// HasMaxValues returns a boolean if a field has been set.
func (o *UserSelectComponentResponse) HasMaxValues() bool {
	if o != nil && o.MaxValues.IsSet() {
		return true
	}

	return false
}

// SetMaxValues gets a reference to the given NullableInt32 and assigns it to the MaxValues field.
func (o *UserSelectComponentResponse) SetMaxValues(v int32) {
	o.MaxValues.Set(&v)
}

// SetMaxValuesNil sets the value for MaxValues to be an explicit nil
func (o *UserSelectComponentResponse) SetMaxValuesNil() {
	o.MaxValues.Set(nil)
}

// UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
func (o *UserSelectComponentResponse) UnsetMaxValues() {
	o.MaxValues.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *UserSelectComponentResponse) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSelectComponentResponse) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *UserSelectComponentResponse) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *UserSelectComponentResponse) SetDisabled(v bool) {
	o.Disabled = &v
}


// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise.
func (o *UserSelectComponentResponse) GetDefaultValues() []UserSelectDefaultValueResponse {
	if o == nil || IsNil(o.DefaultValues) {
		var ret []UserSelectDefaultValueResponse
		return ret
	}
	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSelectComponentResponse) GetDefaultValuesOk() ([]UserSelectDefaultValueResponse, bool) {
	if o == nil || IsNil(o.DefaultValues) {
		return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *UserSelectComponentResponse) HasDefaultValues() bool {
	if o != nil && !IsNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given []UserSelectDefaultValueResponse and assigns it to the DefaultValues field.
func (o *UserSelectComponentResponse) SetDefaultValues(v []UserSelectDefaultValueResponse) {
	o.DefaultValues = v
}


func (o UserSelectComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSelectComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["custom_id"] = o.CustomId
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if o.MinValues.IsSet() {
		toSerialize["min_values"] = o.MinValues.Get()
	}
	if o.MaxValues.IsSet() {
		toSerialize["max_values"] = o.MaxValues.Get()
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.DefaultValues) {
		toSerialize["default_values"] = o.DefaultValues
	}
	return toSerialize, nil
}

func (o *UserSelectComponentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varUserSelectComponentResponse := _UserSelectComponentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSelectComponentResponse)

	if err != nil {
		return err
	}

	*o = UserSelectComponentResponse(varUserSelectComponentResponse)

	return err
}

type NullableUserSelectComponentResponse struct {
	value *UserSelectComponentResponse
	isSet bool
}

func (v NullableUserSelectComponentResponse) Get() *UserSelectComponentResponse {
	return v.value
}

func (v *NullableUserSelectComponentResponse) Set(val *UserSelectComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSelectComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSelectComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSelectComponentResponse(val *UserSelectComponentResponse) *NullableUserSelectComponentResponse {
	return &NullableUserSelectComponentResponse{value: val, isSet: true}
}

func (v NullableUserSelectComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSelectComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

