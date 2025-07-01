/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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

// checks if the ModalInteractionCallbackRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModalInteractionCallbackRequestData{}

// ModalInteractionCallbackRequestData struct for ModalInteractionCallbackRequestData
type ModalInteractionCallbackRequestData struct {
	CustomId string `json:"custom_id"`
	Title string `json:"title"`
	Components []ActionRowComponentForModalRequest `json:"components"`
}

type _ModalInteractionCallbackRequestData ModalInteractionCallbackRequestData

// NewModalInteractionCallbackRequestData instantiates a new ModalInteractionCallbackRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModalInteractionCallbackRequestData(customId string, title string, components []ActionRowComponentForModalRequest) *ModalInteractionCallbackRequestData {
	this := ModalInteractionCallbackRequestData{}
	this.CustomId = customId
	this.Title = title
	this.Components = components
	return &this
}

// NewModalInteractionCallbackRequestDataWithDefaults instantiates a new ModalInteractionCallbackRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModalInteractionCallbackRequestDataWithDefaults() *ModalInteractionCallbackRequestData {
	this := ModalInteractionCallbackRequestData{}
	return &this
}

// GetCustomId returns the CustomId field value
func (o *ModalInteractionCallbackRequestData) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *ModalInteractionCallbackRequestData) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *ModalInteractionCallbackRequestData) SetCustomId(v string) {
	o.CustomId = v
}

// GetTitle returns the Title field value
func (o *ModalInteractionCallbackRequestData) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ModalInteractionCallbackRequestData) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ModalInteractionCallbackRequestData) SetTitle(v string) {
	o.Title = v
}

// GetComponents returns the Components field value
func (o *ModalInteractionCallbackRequestData) GetComponents() []ActionRowComponentForModalRequest {
	if o == nil {
		var ret []ActionRowComponentForModalRequest
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *ModalInteractionCallbackRequestData) GetComponentsOk() ([]ActionRowComponentForModalRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *ModalInteractionCallbackRequestData) SetComponents(v []ActionRowComponentForModalRequest) {
	o.Components = v
}

func (o ModalInteractionCallbackRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModalInteractionCallbackRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["custom_id"] = o.CustomId
	toSerialize["title"] = o.Title
	toSerialize["components"] = o.Components
	return toSerialize, nil
}

func (o *ModalInteractionCallbackRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"custom_id",
		"title",
		"components",
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

	varModalInteractionCallbackRequestData := _ModalInteractionCallbackRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModalInteractionCallbackRequestData)

	if err != nil {
		return err
	}

	*o = ModalInteractionCallbackRequestData(varModalInteractionCallbackRequestData)

	return err
}

type NullableModalInteractionCallbackRequestData struct {
	value *ModalInteractionCallbackRequestData
	isSet bool
}

func (v NullableModalInteractionCallbackRequestData) Get() *ModalInteractionCallbackRequestData {
	return v.value
}

func (v *NullableModalInteractionCallbackRequestData) Set(val *ModalInteractionCallbackRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableModalInteractionCallbackRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableModalInteractionCallbackRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModalInteractionCallbackRequestData(val *ModalInteractionCallbackRequestData) *NullableModalInteractionCallbackRequestData {
	return &NullableModalInteractionCallbackRequestData{value: val, isSet: true}
}

func (v NullableModalInteractionCallbackRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModalInteractionCallbackRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


