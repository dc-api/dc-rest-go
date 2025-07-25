/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// checks if the ApplicationFormPartialDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationFormPartialDescription{}

// ApplicationFormPartialDescription struct for ApplicationFormPartialDescription
type ApplicationFormPartialDescription struct {
	Default string `json:"default"`
	Localizations map[string]string `json:"localizations,omitempty"`
}

type _ApplicationFormPartialDescription ApplicationFormPartialDescription

// NewApplicationFormPartialDescription instantiates a new ApplicationFormPartialDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFormPartialDescription(default_ string) *ApplicationFormPartialDescription {
	this := ApplicationFormPartialDescription{}
	this.Default = default_
	return &this
}

// NewApplicationFormPartialDescriptionWithDefaults instantiates a new ApplicationFormPartialDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFormPartialDescriptionWithDefaults() *ApplicationFormPartialDescription {
	this := ApplicationFormPartialDescription{}
	return &this
}

// GetDefault returns the Default field value
func (o *ApplicationFormPartialDescription) GetDefault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartialDescription) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *ApplicationFormPartialDescription) SetDefault(v string) {
	o.Default = v
}

// GetLocalizations returns the Localizations field value if set, zero value otherwise.
func (o *ApplicationFormPartialDescription) GetLocalizations() map[string]string {
	if o == nil || IsNil(o.Localizations) {
		var ret map[string]string
		return ret
	}
	return o.Localizations
}

// GetLocalizationsOk returns a tuple with the Localizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartialDescription) GetLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Localizations) {
		return map[string]string{}, false
	}
	return o.Localizations, true
}

// HasLocalizations returns a boolean if a field has been set.
func (o *ApplicationFormPartialDescription) HasLocalizations() bool {
	if o != nil && !IsNil(o.Localizations) {
		return true
	}

	return false
}

// SetLocalizations gets a reference to the given map[string]string and assigns it to the Localizations field.
func (o *ApplicationFormPartialDescription) SetLocalizations(v map[string]string) {
	o.Localizations = v
}


func (o ApplicationFormPartialDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationFormPartialDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default"] = o.Default
	if !IsNil(o.Localizations) {
		toSerialize["localizations"] = o.Localizations
	}
	return toSerialize, nil
}

func (o *ApplicationFormPartialDescription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default",
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

	varApplicationFormPartialDescription := _ApplicationFormPartialDescription{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationFormPartialDescription)

	if err != nil {
		return err
	}

	*o = ApplicationFormPartialDescription(varApplicationFormPartialDescription)

	return err
}

type NullableApplicationFormPartialDescription struct {
	value *ApplicationFormPartialDescription
	isSet bool
}

func (v NullableApplicationFormPartialDescription) Get() *ApplicationFormPartialDescription {
	return v.value
}

func (v *NullableApplicationFormPartialDescription) Set(val *ApplicationFormPartialDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFormPartialDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFormPartialDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFormPartialDescription(val *ApplicationFormPartialDescription) *NullableApplicationFormPartialDescription {
	return &NullableApplicationFormPartialDescription{value: val, isSet: true}
}

func (v NullableApplicationFormPartialDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFormPartialDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


