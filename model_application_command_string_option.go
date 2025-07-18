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

// checks if the ApplicationCommandStringOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandStringOption{}

// ApplicationCommandStringOption struct for ApplicationCommandStringOption
type ApplicationCommandStringOption struct {
	Type int32 `json:"type"`
	Name string `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description string `json:"description"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	Required NullableBool `json:"required,omitempty"`
	Autocomplete NullableBool `json:"autocomplete,omitempty"`
	MinLength NullableInt32 `json:"min_length,omitempty"`
	MaxLength NullableInt32 `json:"max_length,omitempty"`
	Choices []ApplicationCommandOptionStringChoice `json:"choices,omitempty"`
}

type _ApplicationCommandStringOption ApplicationCommandStringOption

// NewApplicationCommandStringOption instantiates a new ApplicationCommandStringOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandStringOption(type_ int32, name string, description string) *ApplicationCommandStringOption {
	this := ApplicationCommandStringOption{}
	this.Type = type_
	this.Name = name
	this.Description = description
	return &this
}

// NewApplicationCommandStringOptionWithDefaults instantiates a new ApplicationCommandStringOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandStringOptionWithDefaults() *ApplicationCommandStringOption {
	this := ApplicationCommandStringOption{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationCommandStringOption) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandStringOption) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationCommandStringOption) SetType(v int32) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ApplicationCommandStringOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandStringOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandStringOption) SetName(v string) {
	o.Name = v
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandStringOption) GetNameLocalizations() map[string]string {
	if o == nil || IsNil(o.NameLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandStringOption) GetNameLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return map[string]string{}, false
	}
	return o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandStringOption) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandStringOption) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}


// GetDescription returns the Description field value
func (o *ApplicationCommandStringOption) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandStringOption) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApplicationCommandStringOption) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandStringOption) GetDescriptionLocalizations() map[string]string {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandStringOption) GetDescriptionLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return map[string]string{}, false
	}
	return o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandStringOption) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationCommandStringOption) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}


// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandStringOption) GetRequired() bool {
	if o == nil || IsNil(o.Required.Get()) {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandStringOption) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required.Get()) {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *ApplicationCommandStringOption) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *ApplicationCommandStringOption) SetRequired(v bool) {
	o.Required.Set(&v)
}

// SetRequiredNil sets the value for Required to be an explicit nil
func (o *ApplicationCommandStringOption) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *ApplicationCommandStringOption) UnsetRequired() {
	o.Required.Unset()
}

// GetAutocomplete returns the Autocomplete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandStringOption) GetAutocomplete() bool {
	if o == nil || IsNil(o.Autocomplete.Get()) {
		var ret bool
		return ret
	}
	return *o.Autocomplete.Get()
}

// GetAutocompleteOk returns a tuple with the Autocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandStringOption) GetAutocompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Autocomplete.Get()) {
		return nil, false
	}
	return o.Autocomplete.Get(), o.Autocomplete.IsSet()
}

// HasAutocomplete returns a boolean if a field has been set.
func (o *ApplicationCommandStringOption) HasAutocomplete() bool {
	if o != nil && o.Autocomplete.IsSet() {
		return true
	}

	return false
}

// SetAutocomplete gets a reference to the given NullableBool and assigns it to the Autocomplete field.
func (o *ApplicationCommandStringOption) SetAutocomplete(v bool) {
	o.Autocomplete.Set(&v)
}

// SetAutocompleteNil sets the value for Autocomplete to be an explicit nil
func (o *ApplicationCommandStringOption) SetAutocompleteNil() {
	o.Autocomplete.Set(nil)
}

// UnsetAutocomplete ensures that no value is present for Autocomplete, not even an explicit nil
func (o *ApplicationCommandStringOption) UnsetAutocomplete() {
	o.Autocomplete.Unset()
}

// GetMinLength returns the MinLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandStringOption) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MinLength.Get()
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandStringOption) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength.Get()) {
		return nil, false
	}
	return o.MinLength.Get(), o.MinLength.IsSet()
}

// HasMinLength returns a boolean if a field has been set.
func (o *ApplicationCommandStringOption) HasMinLength() bool {
	if o != nil && o.MinLength.IsSet() {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given NullableInt32 and assigns it to the MinLength field.
func (o *ApplicationCommandStringOption) SetMinLength(v int32) {
	o.MinLength.Set(&v)
}

// SetMinLengthNil sets the value for MinLength to be an explicit nil
func (o *ApplicationCommandStringOption) SetMinLengthNil() {
	o.MinLength.Set(nil)
}

// UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
func (o *ApplicationCommandStringOption) UnsetMinLength() {
	o.MinLength.Unset()
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandStringOption) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxLength.Get()
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandStringOption) GetMaxLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLength.Get()) {
		return nil, false
	}
	return o.MaxLength.Get(), o.MaxLength.IsSet()
}

// HasMaxLength returns a boolean if a field has been set.
func (o *ApplicationCommandStringOption) HasMaxLength() bool {
	if o != nil && o.MaxLength.IsSet() {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given NullableInt32 and assigns it to the MaxLength field.
func (o *ApplicationCommandStringOption) SetMaxLength(v int32) {
	o.MaxLength.Set(&v)
}

// SetMaxLengthNil sets the value for MaxLength to be an explicit nil
func (o *ApplicationCommandStringOption) SetMaxLengthNil() {
	o.MaxLength.Set(nil)
}

// UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
func (o *ApplicationCommandStringOption) UnsetMaxLength() {
	o.MaxLength.Unset()
}

// GetChoices returns the Choices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandStringOption) GetChoices() []ApplicationCommandOptionStringChoice {
	if o == nil {
		var ret []ApplicationCommandOptionStringChoice
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandStringOption) GetChoicesOk() ([]ApplicationCommandOptionStringChoice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *ApplicationCommandStringOption) HasChoices() bool {
	if o != nil && !IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []ApplicationCommandOptionStringChoice and assigns it to the Choices field.
func (o *ApplicationCommandStringOption) SetChoices(v []ApplicationCommandOptionStringChoice) {
	o.Choices = v
}


func (o ApplicationCommandStringOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandStringOption) ToMap() (map[string]interface{}, error) {
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
	if o.MinLength.IsSet() {
		toSerialize["min_length"] = o.MinLength.Get()
	}
	if o.MaxLength.IsSet() {
		toSerialize["max_length"] = o.MaxLength.Get()
	}
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	return toSerialize, nil
}

func (o *ApplicationCommandStringOption) UnmarshalJSON(data []byte) (err error) {
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

	varApplicationCommandStringOption := _ApplicationCommandStringOption{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandStringOption)

	if err != nil {
		return err
	}

	*o = ApplicationCommandStringOption(varApplicationCommandStringOption)

	return err
}

type NullableApplicationCommandStringOption struct {
	value *ApplicationCommandStringOption
	isSet bool
}

func (v NullableApplicationCommandStringOption) Get() *ApplicationCommandStringOption {
	return v.value
}

func (v *NullableApplicationCommandStringOption) Set(val *ApplicationCommandStringOption) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandStringOption) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandStringOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandStringOption(val *ApplicationCommandStringOption) *NullableApplicationCommandStringOption {
	return &NullableApplicationCommandStringOption{value: val, isSet: true}
}

func (v NullableApplicationCommandStringOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandStringOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


