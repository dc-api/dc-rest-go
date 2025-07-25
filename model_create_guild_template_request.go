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

// checks if the CreateGuildTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGuildTemplateRequest{}

// CreateGuildTemplateRequest struct for CreateGuildTemplateRequest
type CreateGuildTemplateRequest struct {
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
}

type _CreateGuildTemplateRequest CreateGuildTemplateRequest

// NewCreateGuildTemplateRequest instantiates a new CreateGuildTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGuildTemplateRequest(name string) *CreateGuildTemplateRequest {
	this := CreateGuildTemplateRequest{}
	this.Name = name
	return &this
}

// NewCreateGuildTemplateRequestWithDefaults instantiates a new CreateGuildTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGuildTemplateRequestWithDefaults() *CreateGuildTemplateRequest {
	this := CreateGuildTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateGuildTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGuildTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGuildTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGuildTemplateRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateGuildTemplateRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateGuildTemplateRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateGuildTemplateRequest) UnsetDescription() {
	o.Description.Unset()
}

func (o CreateGuildTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGuildTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

func (o *CreateGuildTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateGuildTemplateRequest := _CreateGuildTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGuildTemplateRequest)

	if err != nil {
		return err
	}

	*o = CreateGuildTemplateRequest(varCreateGuildTemplateRequest)

	return err
}

type NullableCreateGuildTemplateRequest struct {
	value *CreateGuildTemplateRequest
	isSet bool
}

func (v NullableCreateGuildTemplateRequest) Get() *CreateGuildTemplateRequest {
	return v.value
}

func (v *NullableCreateGuildTemplateRequest) Set(val *CreateGuildTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGuildTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGuildTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGuildTemplateRequest(val *CreateGuildTemplateRequest) *NullableCreateGuildTemplateRequest {
	return &NullableCreateGuildTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateGuildTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGuildTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


