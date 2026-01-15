/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
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
	"bytes"
	"fmt"
)

// checks if the MentionableSelectComponentForMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MentionableSelectComponentForMessageRequest{}

// MentionableSelectComponentForMessageRequest struct for MentionableSelectComponentForMessageRequest
type MentionableSelectComponentForMessageRequest struct {
	Type MessageComponentTypes `json:"type"`
	Id *int32 `json:"id,omitempty"`
	CustomId string `json:"custom_id"`
	Placeholder *string `json:"placeholder,omitempty"`
	MinValues *int32 `json:"min_values,omitempty"`
	MaxValues *int32 `json:"max_values,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Required *bool `json:"required,omitempty"`
	DefaultValues []MentionableSelectComponentForMessageRequestDefaultValuesInner `json:"default_values,omitempty"`
}

type _MentionableSelectComponentForMessageRequest MentionableSelectComponentForMessageRequest

// NewMentionableSelectComponentForMessageRequest instantiates a new MentionableSelectComponentForMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionableSelectComponentForMessageRequest(type_ MessageComponentTypes, customId string) *MentionableSelectComponentForMessageRequest {
	this := MentionableSelectComponentForMessageRequest{}
	this.Type = type_
	this.CustomId = customId
	return &this
}

// NewMentionableSelectComponentForMessageRequestWithDefaults instantiates a new MentionableSelectComponentForMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionableSelectComponentForMessageRequestWithDefaults() *MentionableSelectComponentForMessageRequest {
	this := MentionableSelectComponentForMessageRequest{}
	return &this
}

// GetType returns the Type field value
func (o *MentionableSelectComponentForMessageRequest) GetType() MessageComponentTypes {
	if o == nil {
		var ret MessageComponentTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MentionableSelectComponentForMessageRequest) GetTypeOk() (*MessageComponentTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MentionableSelectComponentForMessageRequest) SetType(v MessageComponentTypes) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MentionableSelectComponentForMessageRequest) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionableSelectComponentForMessageRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MentionableSelectComponentForMessageRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MentionableSelectComponentForMessageRequest) SetId(v int32) {
	o.Id = &v
}


// GetCustomId returns the CustomId field value
func (o *MentionableSelectComponentForMessageRequest) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *MentionableSelectComponentForMessageRequest) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *MentionableSelectComponentForMessageRequest) SetCustomId(v string) {
	o.CustomId = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *MentionableSelectComponentForMessageRequest) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionableSelectComponentForMessageRequest) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *MentionableSelectComponentForMessageRequest) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *MentionableSelectComponentForMessageRequest) SetPlaceholder(v string) {
	o.Placeholder = &v
}


// GetMinValues returns the MinValues field value if set, zero value otherwise.
func (o *MentionableSelectComponentForMessageRequest) GetMinValues() int32 {
	if o == nil || IsNil(o.MinValues) {
		var ret int32
		return ret
	}
	return *o.MinValues
}

// GetMinValuesOk returns a tuple with the MinValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionableSelectComponentForMessageRequest) GetMinValuesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinValues) {
		return nil, false
	}
	return o.MinValues, true
}

// HasMinValues returns a boolean if a field has been set.
func (o *MentionableSelectComponentForMessageRequest) HasMinValues() bool {
	if o != nil && !IsNil(o.MinValues) {
		return true
	}

	return false
}

// SetMinValues gets a reference to the given int32 and assigns it to the MinValues field.
func (o *MentionableSelectComponentForMessageRequest) SetMinValues(v int32) {
	o.MinValues = &v
}


// GetMaxValues returns the MaxValues field value if set, zero value otherwise.
func (o *MentionableSelectComponentForMessageRequest) GetMaxValues() int32 {
	if o == nil || IsNil(o.MaxValues) {
		var ret int32
		return ret
	}
	return *o.MaxValues
}

// GetMaxValuesOk returns a tuple with the MaxValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionableSelectComponentForMessageRequest) GetMaxValuesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxValues) {
		return nil, false
	}
	return o.MaxValues, true
}

// HasMaxValues returns a boolean if a field has been set.
func (o *MentionableSelectComponentForMessageRequest) HasMaxValues() bool {
	if o != nil && !IsNil(o.MaxValues) {
		return true
	}

	return false
}

// SetMaxValues gets a reference to the given int32 and assigns it to the MaxValues field.
func (o *MentionableSelectComponentForMessageRequest) SetMaxValues(v int32) {
	o.MaxValues = &v
}


// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *MentionableSelectComponentForMessageRequest) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionableSelectComponentForMessageRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *MentionableSelectComponentForMessageRequest) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *MentionableSelectComponentForMessageRequest) SetDisabled(v bool) {
	o.Disabled = &v
}


// GetRequired returns the Required field value if set, zero value otherwise.
func (o *MentionableSelectComponentForMessageRequest) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionableSelectComponentForMessageRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *MentionableSelectComponentForMessageRequest) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *MentionableSelectComponentForMessageRequest) SetRequired(v bool) {
	o.Required = &v
}


// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise.
func (o *MentionableSelectComponentForMessageRequest) GetDefaultValues() []MentionableSelectComponentForMessageRequestDefaultValuesInner {
	if o == nil || IsNil(o.DefaultValues) {
		var ret []MentionableSelectComponentForMessageRequestDefaultValuesInner
		return ret
	}
	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionableSelectComponentForMessageRequest) GetDefaultValuesOk() ([]MentionableSelectComponentForMessageRequestDefaultValuesInner, bool) {
	if o == nil || IsNil(o.DefaultValues) {
		return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *MentionableSelectComponentForMessageRequest) HasDefaultValues() bool {
	if o != nil && !IsNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given []MentionableSelectComponentForMessageRequestDefaultValuesInner and assigns it to the DefaultValues field.
func (o *MentionableSelectComponentForMessageRequest) SetDefaultValues(v []MentionableSelectComponentForMessageRequestDefaultValuesInner) {
	o.DefaultValues = v
}


func (o MentionableSelectComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MentionableSelectComponentForMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["custom_id"] = o.CustomId
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if !IsNil(o.MinValues) {
		toSerialize["min_values"] = o.MinValues
	}
	if !IsNil(o.MaxValues) {
		toSerialize["max_values"] = o.MaxValues
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.DefaultValues) {
		toSerialize["default_values"] = o.DefaultValues
	}
	return toSerialize, nil
}

func (o *MentionableSelectComponentForMessageRequest) UnmarshalJSON(data []byte) (err error) {
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

	varMentionableSelectComponentForMessageRequest := _MentionableSelectComponentForMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMentionableSelectComponentForMessageRequest)

	if err != nil {
		return err
	}

	*o = MentionableSelectComponentForMessageRequest(varMentionableSelectComponentForMessageRequest)

	return err
}

type NullableMentionableSelectComponentForMessageRequest struct {
	value *MentionableSelectComponentForMessageRequest
	isSet bool
}

func (v NullableMentionableSelectComponentForMessageRequest) Get() *MentionableSelectComponentForMessageRequest {
	return v.value
}

func (v *NullableMentionableSelectComponentForMessageRequest) Set(val *MentionableSelectComponentForMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionableSelectComponentForMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionableSelectComponentForMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionableSelectComponentForMessageRequest(val *MentionableSelectComponentForMessageRequest) *NullableMentionableSelectComponentForMessageRequest {
	return &NullableMentionableSelectComponentForMessageRequest{value: val, isSet: true}
}

func (v NullableMentionableSelectComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionableSelectComponentForMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

