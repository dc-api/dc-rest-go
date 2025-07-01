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

// checks if the ApplicationCommandOptionNumberChoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandOptionNumberChoice{}

// ApplicationCommandOptionNumberChoice struct for ApplicationCommandOptionNumberChoice
type ApplicationCommandOptionNumberChoice struct {
	Name string `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Value float64 `json:"value"`
}

type _ApplicationCommandOptionNumberChoice ApplicationCommandOptionNumberChoice

// NewApplicationCommandOptionNumberChoice instantiates a new ApplicationCommandOptionNumberChoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandOptionNumberChoice(name string, value float64) *ApplicationCommandOptionNumberChoice {
	this := ApplicationCommandOptionNumberChoice{}
	this.Name = name
	this.Value = value
	return &this
}

// NewApplicationCommandOptionNumberChoiceWithDefaults instantiates a new ApplicationCommandOptionNumberChoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandOptionNumberChoiceWithDefaults() *ApplicationCommandOptionNumberChoice {
	this := ApplicationCommandOptionNumberChoice{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationCommandOptionNumberChoice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandOptionNumberChoice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandOptionNumberChoice) SetName(v string) {
	o.Name = v
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandOptionNumberChoice) GetNameLocalizations() map[string]string {
	if o == nil || IsNil(o.NameLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandOptionNumberChoice) GetNameLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return map[string]string{}, false
	}
	return o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandOptionNumberChoice) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandOptionNumberChoice) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}


// GetValue returns the Value field value
func (o *ApplicationCommandOptionNumberChoice) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandOptionNumberChoice) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ApplicationCommandOptionNumberChoice) SetValue(v float64) {
	o.Value = v
}

func (o ApplicationCommandOptionNumberChoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandOptionNumberChoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.NameLocalizations) {
		toSerialize["name_localizations"] = o.NameLocalizations
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ApplicationCommandOptionNumberChoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varApplicationCommandOptionNumberChoice := _ApplicationCommandOptionNumberChoice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandOptionNumberChoice)

	if err != nil {
		return err
	}

	*o = ApplicationCommandOptionNumberChoice(varApplicationCommandOptionNumberChoice)

	return err
}

type NullableApplicationCommandOptionNumberChoice struct {
	value *ApplicationCommandOptionNumberChoice
	isSet bool
}

func (v NullableApplicationCommandOptionNumberChoice) Get() *ApplicationCommandOptionNumberChoice {
	return v.value
}

func (v *NullableApplicationCommandOptionNumberChoice) Set(val *ApplicationCommandOptionNumberChoice) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandOptionNumberChoice) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandOptionNumberChoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandOptionNumberChoice(val *ApplicationCommandOptionNumberChoice) *NullableApplicationCommandOptionNumberChoice {
	return &NullableApplicationCommandOptionNumberChoice{value: val, isSet: true}
}

func (v NullableApplicationCommandOptionNumberChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandOptionNumberChoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


