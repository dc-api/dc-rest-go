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

// checks if the SectionComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SectionComponentResponse{}

// SectionComponentResponse struct for SectionComponentResponse
type SectionComponentResponse struct {
	Type int32 `json:"type"`
	Id int32 `json:"id"`
	Components []TextDisplayComponentResponse `json:"components"`
	Accessory SectionComponentResponseAccessory `json:"accessory"`
}

type _SectionComponentResponse SectionComponentResponse

// NewSectionComponentResponse instantiates a new SectionComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSectionComponentResponse(type_ int32, id int32, components []TextDisplayComponentResponse, accessory SectionComponentResponseAccessory) *SectionComponentResponse {
	this := SectionComponentResponse{}
	this.Type = type_
	this.Id = id
	this.Components = components
	this.Accessory = accessory
	return &this
}

// NewSectionComponentResponseWithDefaults instantiates a new SectionComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSectionComponentResponseWithDefaults() *SectionComponentResponse {
	this := SectionComponentResponse{}
	return &this
}

// GetType returns the Type field value
func (o *SectionComponentResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SectionComponentResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SectionComponentResponse) SetType(v int32) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SectionComponentResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SectionComponentResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SectionComponentResponse) SetId(v int32) {
	o.Id = v
}

// GetComponents returns the Components field value
func (o *SectionComponentResponse) GetComponents() []TextDisplayComponentResponse {
	if o == nil {
		var ret []TextDisplayComponentResponse
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *SectionComponentResponse) GetComponentsOk() ([]TextDisplayComponentResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *SectionComponentResponse) SetComponents(v []TextDisplayComponentResponse) {
	o.Components = v
}

// GetAccessory returns the Accessory field value
func (o *SectionComponentResponse) GetAccessory() SectionComponentResponseAccessory {
	if o == nil {
		var ret SectionComponentResponseAccessory
		return ret
	}

	return o.Accessory
}

// GetAccessoryOk returns a tuple with the Accessory field value
// and a boolean to check if the value has been set.
func (o *SectionComponentResponse) GetAccessoryOk() (*SectionComponentResponseAccessory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accessory, true
}

// SetAccessory sets field value
func (o *SectionComponentResponse) SetAccessory(v SectionComponentResponseAccessory) {
	o.Accessory = v
}

func (o SectionComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SectionComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["components"] = o.Components
	toSerialize["accessory"] = o.Accessory
	return toSerialize, nil
}

func (o *SectionComponentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"components",
		"accessory",
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

	varSectionComponentResponse := _SectionComponentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSectionComponentResponse)

	if err != nil {
		return err
	}

	*o = SectionComponentResponse(varSectionComponentResponse)

	return err
}

type NullableSectionComponentResponse struct {
	value *SectionComponentResponse
	isSet bool
}

func (v NullableSectionComponentResponse) Get() *SectionComponentResponse {
	return v.value
}

func (v *NullableSectionComponentResponse) Set(val *SectionComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSectionComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSectionComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectionComponentResponse(val *SectionComponentResponse) *NullableSectionComponentResponse {
	return &NullableSectionComponentResponse{value: val, isSet: true}
}

func (v NullableSectionComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectionComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


