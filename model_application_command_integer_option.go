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

// checks if the ApplicationCommandIntegerOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandIntegerOption{}

// ApplicationCommandIntegerOption struct for ApplicationCommandIntegerOption
type ApplicationCommandIntegerOption struct {
	Type int32 `json:"type"`
	Name string `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description string `json:"description"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	Required NullableBool `json:"required,omitempty"`
	Autocomplete NullableBool `json:"autocomplete,omitempty"`
	Choices []ApplicationCommandOptionIntegerChoice `json:"choices,omitempty"`
	MinValue *int64 `json:"min_value,omitempty"`
	MaxValue *int64 `json:"max_value,omitempty"`
}

type _ApplicationCommandIntegerOption ApplicationCommandIntegerOption

// NewApplicationCommandIntegerOption instantiates a new ApplicationCommandIntegerOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandIntegerOption(type_ int32, name string, description string) *ApplicationCommandIntegerOption {
	this := ApplicationCommandIntegerOption{}
	this.Type = type_
	this.Name = name
	this.Description = description
	return &this
}

// NewApplicationCommandIntegerOptionWithDefaults instantiates a new ApplicationCommandIntegerOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandIntegerOptionWithDefaults() *ApplicationCommandIntegerOption {
	this := ApplicationCommandIntegerOption{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationCommandIntegerOption) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandIntegerOption) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationCommandIntegerOption) SetType(v int32) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ApplicationCommandIntegerOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandIntegerOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandIntegerOption) SetName(v string) {
	o.Name = v
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandIntegerOption) GetNameLocalizations() map[string]string {
	if o == nil || IsNil(o.NameLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandIntegerOption) GetNameLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return map[string]string{}, false
	}
	return o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandIntegerOption) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandIntegerOption) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}


// GetDescription returns the Description field value
func (o *ApplicationCommandIntegerOption) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandIntegerOption) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApplicationCommandIntegerOption) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandIntegerOption) GetDescriptionLocalizations() map[string]string {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandIntegerOption) GetDescriptionLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return map[string]string{}, false
	}
	return o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandIntegerOption) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationCommandIntegerOption) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}


// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandIntegerOption) GetRequired() bool {
	if o == nil || IsNil(o.Required.Get()) {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandIntegerOption) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required.Get()) {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *ApplicationCommandIntegerOption) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *ApplicationCommandIntegerOption) SetRequired(v bool) {
	o.Required.Set(&v)
}

// SetRequiredNil sets the value for Required to be an explicit nil
func (o *ApplicationCommandIntegerOption) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *ApplicationCommandIntegerOption) UnsetRequired() {
	o.Required.Unset()
}

// GetAutocomplete returns the Autocomplete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandIntegerOption) GetAutocomplete() bool {
	if o == nil || IsNil(o.Autocomplete.Get()) {
		var ret bool
		return ret
	}
	return *o.Autocomplete.Get()
}

// GetAutocompleteOk returns a tuple with the Autocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandIntegerOption) GetAutocompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Autocomplete.Get()) {
		return nil, false
	}
	return o.Autocomplete.Get(), o.Autocomplete.IsSet()
}

// HasAutocomplete returns a boolean if a field has been set.
func (o *ApplicationCommandIntegerOption) HasAutocomplete() bool {
	if o != nil && o.Autocomplete.IsSet() {
		return true
	}

	return false
}

// SetAutocomplete gets a reference to the given NullableBool and assigns it to the Autocomplete field.
func (o *ApplicationCommandIntegerOption) SetAutocomplete(v bool) {
	o.Autocomplete.Set(&v)
}

// SetAutocompleteNil sets the value for Autocomplete to be an explicit nil
func (o *ApplicationCommandIntegerOption) SetAutocompleteNil() {
	o.Autocomplete.Set(nil)
}

// UnsetAutocomplete ensures that no value is present for Autocomplete, not even an explicit nil
func (o *ApplicationCommandIntegerOption) UnsetAutocomplete() {
	o.Autocomplete.Unset()
}

// GetChoices returns the Choices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandIntegerOption) GetChoices() []ApplicationCommandOptionIntegerChoice {
	if o == nil {
		var ret []ApplicationCommandOptionIntegerChoice
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandIntegerOption) GetChoicesOk() ([]ApplicationCommandOptionIntegerChoice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *ApplicationCommandIntegerOption) HasChoices() bool {
	if o != nil && !IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []ApplicationCommandOptionIntegerChoice and assigns it to the Choices field.
func (o *ApplicationCommandIntegerOption) SetChoices(v []ApplicationCommandOptionIntegerChoice) {
	o.Choices = v
}


// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *ApplicationCommandIntegerOption) GetMinValue() int64 {
	if o == nil || IsNil(o.MinValue) {
		var ret int64
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandIntegerOption) GetMinValueOk() (*int64, bool) {
	if o == nil || IsNil(o.MinValue) {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *ApplicationCommandIntegerOption) HasMinValue() bool {
	if o != nil && !IsNil(o.MinValue) {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given int64 and assigns it to the MinValue field.
func (o *ApplicationCommandIntegerOption) SetMinValue(v int64) {
	o.MinValue = &v
}


// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *ApplicationCommandIntegerOption) GetMaxValue() int64 {
	if o == nil || IsNil(o.MaxValue) {
		var ret int64
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandIntegerOption) GetMaxValueOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxValue) {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *ApplicationCommandIntegerOption) HasMaxValue() bool {
	if o != nil && !IsNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given int64 and assigns it to the MaxValue field.
func (o *ApplicationCommandIntegerOption) SetMaxValue(v int64) {
	o.MaxValue = &v
}


func (o ApplicationCommandIntegerOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandIntegerOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.NameLocalizations) {
		toSerialize["name_localizations"] = o.NameLocalizations
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.DescriptionLocalizations) {
		toSerialize["description_localizations"] = o.DescriptionLocalizations
	}
	if o.Required.IsSet() {
		toSerialize["required"] = o.Required.Get()
	}
	if o.Autocomplete.IsSet() {
		toSerialize["autocomplete"] = o.Autocomplete.Get()
	}
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	if !IsNil(o.MinValue) {
		toSerialize["min_value"] = o.MinValue
	}
	if !IsNil(o.MaxValue) {
		toSerialize["max_value"] = o.MaxValue
	}
	return toSerialize, nil
}

func (o *ApplicationCommandIntegerOption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"description",
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

	varApplicationCommandIntegerOption := _ApplicationCommandIntegerOption{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandIntegerOption)

	if err != nil {
		return err
	}

	*o = ApplicationCommandIntegerOption(varApplicationCommandIntegerOption)

	return err
}

type NullableApplicationCommandIntegerOption struct {
	value *ApplicationCommandIntegerOption
	isSet bool
}

func (v NullableApplicationCommandIntegerOption) Get() *ApplicationCommandIntegerOption {
	return v.value
}

func (v *NullableApplicationCommandIntegerOption) Set(val *ApplicationCommandIntegerOption) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandIntegerOption) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandIntegerOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandIntegerOption(val *ApplicationCommandIntegerOption) *NullableApplicationCommandIntegerOption {
	return &NullableApplicationCommandIntegerOption{value: val, isSet: true}
}

func (v NullableApplicationCommandIntegerOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandIntegerOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


