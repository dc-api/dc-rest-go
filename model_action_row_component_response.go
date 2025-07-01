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

// checks if the ActionRowComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionRowComponentResponse{}

// ActionRowComponentResponse struct for ActionRowComponentResponse
type ActionRowComponentResponse struct {
	Type int32 `json:"type"`
	Id int32 `json:"id"`
	Components []ActionRowComponentResponseComponentsInner `json:"components,omitempty"`
}

type _ActionRowComponentResponse ActionRowComponentResponse

// NewActionRowComponentResponse instantiates a new ActionRowComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionRowComponentResponse(type_ int32, id int32) *ActionRowComponentResponse {
	this := ActionRowComponentResponse{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewActionRowComponentResponseWithDefaults instantiates a new ActionRowComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionRowComponentResponseWithDefaults() *ActionRowComponentResponse {
	this := ActionRowComponentResponse{}
	return &this
}

// GetType returns the Type field value
func (o *ActionRowComponentResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActionRowComponentResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActionRowComponentResponse) SetType(v int32) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ActionRowComponentResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActionRowComponentResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActionRowComponentResponse) SetId(v int32) {
	o.Id = v
}

// GetComponents returns the Components field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionRowComponentResponse) GetComponents() []ActionRowComponentResponseComponentsInner {
	if o == nil {
		var ret []ActionRowComponentResponseComponentsInner
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionRowComponentResponse) GetComponentsOk() ([]ActionRowComponentResponseComponentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ActionRowComponentResponse) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ActionRowComponentResponseComponentsInner and assigns it to the Components field.
func (o *ActionRowComponentResponse) SetComponents(v []ActionRowComponentResponseComponentsInner) {
	o.Components = v
}


func (o ActionRowComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionRowComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	return toSerialize, nil
}

func (o *ActionRowComponentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varActionRowComponentResponse := _ActionRowComponentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActionRowComponentResponse)

	if err != nil {
		return err
	}

	*o = ActionRowComponentResponse(varActionRowComponentResponse)

	return err
}

type NullableActionRowComponentResponse struct {
	value *ActionRowComponentResponse
	isSet bool
}

func (v NullableActionRowComponentResponse) Get() *ActionRowComponentResponse {
	return v.value
}

func (v *NullableActionRowComponentResponse) Set(val *ActionRowComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableActionRowComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableActionRowComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionRowComponentResponse(val *ActionRowComponentResponse) *NullableActionRowComponentResponse {
	return &NullableActionRowComponentResponse{value: val, isSet: true}
}

func (v NullableActionRowComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionRowComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


