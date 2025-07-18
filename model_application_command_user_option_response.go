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

// checks if the ApplicationCommandUserOptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandUserOptionResponse{}

// ApplicationCommandUserOptionResponse struct for ApplicationCommandUserOptionResponse
type ApplicationCommandUserOptionResponse struct {
	Type int32 `json:"type"`
	Name string `json:"name"`
	NameLocalized NullableString `json:"name_localized,omitempty"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description string `json:"description"`
	DescriptionLocalized NullableString `json:"description_localized,omitempty"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	Required NullableBool `json:"required,omitempty"`
}

type _ApplicationCommandUserOptionResponse ApplicationCommandUserOptionResponse

// NewApplicationCommandUserOptionResponse instantiates a new ApplicationCommandUserOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandUserOptionResponse(type_ int32, name string, description string) *ApplicationCommandUserOptionResponse {
	this := ApplicationCommandUserOptionResponse{}
	this.Type = type_
	this.Name = name
	this.Description = description
	return &this
}

// NewApplicationCommandUserOptionResponseWithDefaults instantiates a new ApplicationCommandUserOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandUserOptionResponseWithDefaults() *ApplicationCommandUserOptionResponse {
	this := ApplicationCommandUserOptionResponse{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationCommandUserOptionResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandUserOptionResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationCommandUserOptionResponse) SetType(v int32) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ApplicationCommandUserOptionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandUserOptionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandUserOptionResponse) SetName(v string) {
	o.Name = v
}

// GetNameLocalized returns the NameLocalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUserOptionResponse) GetNameLocalized() string {
	if o == nil || IsNil(o.NameLocalized.Get()) {
		var ret string
		return ret
	}
	return *o.NameLocalized.Get()
}

// GetNameLocalizedOk returns a tuple with the NameLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUserOptionResponse) GetNameLocalizedOk() (*string, bool) {
	if o == nil || IsNil(o.NameLocalized.Get()) {
		return nil, false
	}
	return o.NameLocalized.Get(), o.NameLocalized.IsSet()
}

// HasNameLocalized returns a boolean if a field has been set.
func (o *ApplicationCommandUserOptionResponse) HasNameLocalized() bool {
	if o != nil && o.NameLocalized.IsSet() {
		return true
	}

	return false
}

// SetNameLocalized gets a reference to the given NullableString and assigns it to the NameLocalized field.
func (o *ApplicationCommandUserOptionResponse) SetNameLocalized(v string) {
	o.NameLocalized.Set(&v)
}

// SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil
func (o *ApplicationCommandUserOptionResponse) SetNameLocalizedNil() {
	o.NameLocalized.Set(nil)
}

// UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil
func (o *ApplicationCommandUserOptionResponse) UnsetNameLocalized() {
	o.NameLocalized.Unset()
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandUserOptionResponse) GetNameLocalizations() map[string]string {
	if o == nil || IsNil(o.NameLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandUserOptionResponse) GetNameLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return map[string]string{}, false
	}
	return o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandUserOptionResponse) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandUserOptionResponse) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}


// GetDescription returns the Description field value
func (o *ApplicationCommandUserOptionResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandUserOptionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApplicationCommandUserOptionResponse) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionLocalized returns the DescriptionLocalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUserOptionResponse) GetDescriptionLocalized() string {
	if o == nil || IsNil(o.DescriptionLocalized.Get()) {
		var ret string
		return ret
	}
	return *o.DescriptionLocalized.Get()
}

// GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUserOptionResponse) GetDescriptionLocalizedOk() (*string, bool) {
	if o == nil || IsNil(o.DescriptionLocalized.Get()) {
		return nil, false
	}
	return o.DescriptionLocalized.Get(), o.DescriptionLocalized.IsSet()
}

// HasDescriptionLocalized returns a boolean if a field has been set.
func (o *ApplicationCommandUserOptionResponse) HasDescriptionLocalized() bool {
	if o != nil && o.DescriptionLocalized.IsSet() {
		return true
	}

	return false
}

// SetDescriptionLocalized gets a reference to the given NullableString and assigns it to the DescriptionLocalized field.
func (o *ApplicationCommandUserOptionResponse) SetDescriptionLocalized(v string) {
	o.DescriptionLocalized.Set(&v)
}

// SetDescriptionLocalizedNil sets the value for DescriptionLocalized to be an explicit nil
func (o *ApplicationCommandUserOptionResponse) SetDescriptionLocalizedNil() {
	o.DescriptionLocalized.Set(nil)
}

// UnsetDescriptionLocalized ensures that no value is present for DescriptionLocalized, not even an explicit nil
func (o *ApplicationCommandUserOptionResponse) UnsetDescriptionLocalized() {
	o.DescriptionLocalized.Unset()
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandUserOptionResponse) GetDescriptionLocalizations() map[string]string {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandUserOptionResponse) GetDescriptionLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return map[string]string{}, false
	}
	return o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandUserOptionResponse) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationCommandUserOptionResponse) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}


// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUserOptionResponse) GetRequired() bool {
	if o == nil || IsNil(o.Required.Get()) {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUserOptionResponse) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required.Get()) {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *ApplicationCommandUserOptionResponse) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *ApplicationCommandUserOptionResponse) SetRequired(v bool) {
	o.Required.Set(&v)
}

// SetRequiredNil sets the value for Required to be an explicit nil
func (o *ApplicationCommandUserOptionResponse) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *ApplicationCommandUserOptionResponse) UnsetRequired() {
	o.Required.Unset()
}

func (o ApplicationCommandUserOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandUserOptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if o.NameLocalized.IsSet() {
		toSerialize["name_localized"] = o.NameLocalized.Get()
	}
	if !IsNil(o.NameLocalizations) {
		toSerialize["name_localizations"] = o.NameLocalizations
	}
	toSerialize["description"] = o.Description
	if o.DescriptionLocalized.IsSet() {
		toSerialize["description_localized"] = o.DescriptionLocalized.Get()
	}
	if !IsNil(o.DescriptionLocalizations) {
		toSerialize["description_localizations"] = o.DescriptionLocalizations
	}
	if o.Required.IsSet() {
		toSerialize["required"] = o.Required.Get()
	}
	return toSerialize, nil
}

func (o *ApplicationCommandUserOptionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varApplicationCommandUserOptionResponse := _ApplicationCommandUserOptionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandUserOptionResponse)

	if err != nil {
		return err
	}

	*o = ApplicationCommandUserOptionResponse(varApplicationCommandUserOptionResponse)

	return err
}

type NullableApplicationCommandUserOptionResponse struct {
	value *ApplicationCommandUserOptionResponse
	isSet bool
}

func (v NullableApplicationCommandUserOptionResponse) Get() *ApplicationCommandUserOptionResponse {
	return v.value
}

func (v *NullableApplicationCommandUserOptionResponse) Set(val *ApplicationCommandUserOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandUserOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandUserOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandUserOptionResponse(val *ApplicationCommandUserOptionResponse) *NullableApplicationCommandUserOptionResponse {
	return &NullableApplicationCommandUserOptionResponse{value: val, isSet: true}
}

func (v NullableApplicationCommandUserOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandUserOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


