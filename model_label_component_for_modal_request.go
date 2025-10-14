/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-10-14T14:27:02.905060650Z[Etc/UTC]
 *    * - **Generator Version**: 7.16.0
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

// checks if the LabelComponentForModalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelComponentForModalRequest{}

// LabelComponentForModalRequest struct for LabelComponentForModalRequest
type LabelComponentForModalRequest struct {
	Type int32 `json:"type"`
	Id NullableInt32 `json:"id,omitempty"`
	Label string `json:"label"`
	Description NullableString `json:"description,omitempty"`
	Component LabelComponentForModalRequestComponent `json:"component"`
}

type _LabelComponentForModalRequest LabelComponentForModalRequest

// NewLabelComponentForModalRequest instantiates a new LabelComponentForModalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelComponentForModalRequest(type_ int32, label string, component LabelComponentForModalRequestComponent) *LabelComponentForModalRequest {
	this := LabelComponentForModalRequest{}
	this.Type = type_
	this.Label = label
	this.Component = component
	return &this
}

// NewLabelComponentForModalRequestWithDefaults instantiates a new LabelComponentForModalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelComponentForModalRequestWithDefaults() *LabelComponentForModalRequest {
	this := LabelComponentForModalRequest{}
	return &this
}

// GetType returns the Type field value
func (o *LabelComponentForModalRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LabelComponentForModalRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LabelComponentForModalRequest) SetType(v int32) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LabelComponentForModalRequest) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelComponentForModalRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id.Get()) {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *LabelComponentForModalRequest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *LabelComponentForModalRequest) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *LabelComponentForModalRequest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *LabelComponentForModalRequest) UnsetId() {
	o.Id.Unset()
}

// GetLabel returns the Label field value
func (o *LabelComponentForModalRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *LabelComponentForModalRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *LabelComponentForModalRequest) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LabelComponentForModalRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelComponentForModalRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *LabelComponentForModalRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *LabelComponentForModalRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *LabelComponentForModalRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *LabelComponentForModalRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetComponent returns the Component field value
func (o *LabelComponentForModalRequest) GetComponent() LabelComponentForModalRequestComponent {
	if o == nil {
		var ret LabelComponentForModalRequestComponent
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *LabelComponentForModalRequest) GetComponentOk() (*LabelComponentForModalRequestComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *LabelComponentForModalRequest) SetComponent(v LabelComponentForModalRequestComponent) {
	o.Component = v
}

func (o LabelComponentForModalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelComponentForModalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["label"] = o.Label
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["component"] = o.Component
	return toSerialize, nil
}

func (o *LabelComponentForModalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"label",
		"component",
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

	varLabelComponentForModalRequest := _LabelComponentForModalRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLabelComponentForModalRequest)

	if err != nil {
		return err
	}

	*o = LabelComponentForModalRequest(varLabelComponentForModalRequest)

	return err
}

type NullableLabelComponentForModalRequest struct {
	value *LabelComponentForModalRequest
	isSet bool
}

func (v NullableLabelComponentForModalRequest) Get() *LabelComponentForModalRequest {
	return v.value
}

func (v *NullableLabelComponentForModalRequest) Set(val *LabelComponentForModalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelComponentForModalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelComponentForModalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelComponentForModalRequest(val *LabelComponentForModalRequest) *NullableLabelComponentForModalRequest {
	return &NullableLabelComponentForModalRequest{value: val, isSet: true}
}

func (v NullableLabelComponentForModalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelComponentForModalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

