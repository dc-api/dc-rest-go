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

// checks if the ApplicationCommandAttachmentOptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandAttachmentOptionResponse{}

// ApplicationCommandAttachmentOptionResponse struct for ApplicationCommandAttachmentOptionResponse
type ApplicationCommandAttachmentOptionResponse struct {
	Type int32 `json:"type"`
	Name string `json:"name"`
	NameLocalized NullableString `json:"name_localized,omitempty"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description string `json:"description"`
	DescriptionLocalized NullableString `json:"description_localized,omitempty"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	Required NullableBool `json:"required,omitempty"`
}

type _ApplicationCommandAttachmentOptionResponse ApplicationCommandAttachmentOptionResponse

// NewApplicationCommandAttachmentOptionResponse instantiates a new ApplicationCommandAttachmentOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandAttachmentOptionResponse(type_ int32, name string, description string) *ApplicationCommandAttachmentOptionResponse {
	this := ApplicationCommandAttachmentOptionResponse{}
	this.Type = type_
	this.Name = name
	this.Description = description
	return &this
}

// NewApplicationCommandAttachmentOptionResponseWithDefaults instantiates a new ApplicationCommandAttachmentOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandAttachmentOptionResponseWithDefaults() *ApplicationCommandAttachmentOptionResponse {
	this := ApplicationCommandAttachmentOptionResponse{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationCommandAttachmentOptionResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandAttachmentOptionResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationCommandAttachmentOptionResponse) SetType(v int32) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ApplicationCommandAttachmentOptionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandAttachmentOptionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandAttachmentOptionResponse) SetName(v string) {
	o.Name = v
}

// GetNameLocalized returns the NameLocalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandAttachmentOptionResponse) GetNameLocalized() string {
	if o == nil || IsNil(o.NameLocalized.Get()) {
		var ret string
		return ret
	}
	return *o.NameLocalized.Get()
}

// GetNameLocalizedOk returns a tuple with the NameLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandAttachmentOptionResponse) GetNameLocalizedOk() (*string, bool) {
	if o == nil || IsNil(o.NameLocalized.Get()) {
		return nil, false
	}
	return o.NameLocalized.Get(), o.NameLocalized.IsSet()
}

// HasNameLocalized returns a boolean if a field has been set.
func (o *ApplicationCommandAttachmentOptionResponse) HasNameLocalized() bool {
	if o != nil && o.NameLocalized.IsSet() {
		return true
	}

	return false
}

// SetNameLocalized gets a reference to the given NullableString and assigns it to the NameLocalized field.
func (o *ApplicationCommandAttachmentOptionResponse) SetNameLocalized(v string) {
	o.NameLocalized.Set(&v)
}

// SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil
func (o *ApplicationCommandAttachmentOptionResponse) SetNameLocalizedNil() {
	o.NameLocalized.Set(nil)
}

// UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil
func (o *ApplicationCommandAttachmentOptionResponse) UnsetNameLocalized() {
	o.NameLocalized.Unset()
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandAttachmentOptionResponse) GetNameLocalizations() map[string]string {
	if o == nil || IsNil(o.NameLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandAttachmentOptionResponse) GetNameLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return map[string]string{}, false
	}
	return o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandAttachmentOptionResponse) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandAttachmentOptionResponse) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}


// GetDescription returns the Description field value
func (o *ApplicationCommandAttachmentOptionResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandAttachmentOptionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApplicationCommandAttachmentOptionResponse) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionLocalized returns the DescriptionLocalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandAttachmentOptionResponse) GetDescriptionLocalized() string {
	if o == nil || IsNil(o.DescriptionLocalized.Get()) {
		var ret string
		return ret
	}
	return *o.DescriptionLocalized.Get()
}

// GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandAttachmentOptionResponse) GetDescriptionLocalizedOk() (*string, bool) {
	if o == nil || IsNil(o.DescriptionLocalized.Get()) {
		return nil, false
	}
	return o.DescriptionLocalized.Get(), o.DescriptionLocalized.IsSet()
}

// HasDescriptionLocalized returns a boolean if a field has been set.
func (o *ApplicationCommandAttachmentOptionResponse) HasDescriptionLocalized() bool {
	if o != nil && o.DescriptionLocalized.IsSet() {
		return true
	}

	return false
}

// SetDescriptionLocalized gets a reference to the given NullableString and assigns it to the DescriptionLocalized field.
func (o *ApplicationCommandAttachmentOptionResponse) SetDescriptionLocalized(v string) {
	o.DescriptionLocalized.Set(&v)
}

// SetDescriptionLocalizedNil sets the value for DescriptionLocalized to be an explicit nil
func (o *ApplicationCommandAttachmentOptionResponse) SetDescriptionLocalizedNil() {
	o.DescriptionLocalized.Set(nil)
}

// UnsetDescriptionLocalized ensures that no value is present for DescriptionLocalized, not even an explicit nil
func (o *ApplicationCommandAttachmentOptionResponse) UnsetDescriptionLocalized() {
	o.DescriptionLocalized.Unset()
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandAttachmentOptionResponse) GetDescriptionLocalizations() map[string]string {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandAttachmentOptionResponse) GetDescriptionLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return map[string]string{}, false
	}
	return o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandAttachmentOptionResponse) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationCommandAttachmentOptionResponse) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}


// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandAttachmentOptionResponse) GetRequired() bool {
	if o == nil || IsNil(o.Required.Get()) {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandAttachmentOptionResponse) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required.Get()) {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *ApplicationCommandAttachmentOptionResponse) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *ApplicationCommandAttachmentOptionResponse) SetRequired(v bool) {
	o.Required.Set(&v)
}

// SetRequiredNil sets the value for Required to be an explicit nil
func (o *ApplicationCommandAttachmentOptionResponse) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *ApplicationCommandAttachmentOptionResponse) UnsetRequired() {
	o.Required.Unset()
}

func (o ApplicationCommandAttachmentOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandAttachmentOptionResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ApplicationCommandAttachmentOptionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varApplicationCommandAttachmentOptionResponse := _ApplicationCommandAttachmentOptionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandAttachmentOptionResponse)

	if err != nil {
		return err
	}

	*o = ApplicationCommandAttachmentOptionResponse(varApplicationCommandAttachmentOptionResponse)

	return err
}

type NullableApplicationCommandAttachmentOptionResponse struct {
	value *ApplicationCommandAttachmentOptionResponse
	isSet bool
}

func (v NullableApplicationCommandAttachmentOptionResponse) Get() *ApplicationCommandAttachmentOptionResponse {
	return v.value
}

func (v *NullableApplicationCommandAttachmentOptionResponse) Set(val *ApplicationCommandAttachmentOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandAttachmentOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandAttachmentOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandAttachmentOptionResponse(val *ApplicationCommandAttachmentOptionResponse) *NullableApplicationCommandAttachmentOptionResponse {
	return &NullableApplicationCommandAttachmentOptionResponse{value: val, isSet: true}
}

func (v NullableApplicationCommandAttachmentOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandAttachmentOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


