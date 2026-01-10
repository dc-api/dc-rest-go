/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-10T02:53:50.800755799Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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

// checks if the FileUploadComponentForModalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileUploadComponentForModalRequest{}

// FileUploadComponentForModalRequest struct for FileUploadComponentForModalRequest
type FileUploadComponentForModalRequest struct {
	Type MessageComponentTypes `json:"type"`
	Id *int32 `json:"id,omitempty"`
	CustomId string `json:"custom_id"`
	MinValues *int32 `json:"min_values,omitempty"`
	MaxValues *int32 `json:"max_values,omitempty"`
	Required *bool `json:"required,omitempty"`
}

type _FileUploadComponentForModalRequest FileUploadComponentForModalRequest

// NewFileUploadComponentForModalRequest instantiates a new FileUploadComponentForModalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileUploadComponentForModalRequest(type_ MessageComponentTypes, customId string) *FileUploadComponentForModalRequest {
	this := FileUploadComponentForModalRequest{}
	this.Type = type_
	this.CustomId = customId
	return &this
}

// NewFileUploadComponentForModalRequestWithDefaults instantiates a new FileUploadComponentForModalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileUploadComponentForModalRequestWithDefaults() *FileUploadComponentForModalRequest {
	this := FileUploadComponentForModalRequest{}
	return &this
}

// GetType returns the Type field value
func (o *FileUploadComponentForModalRequest) GetType() MessageComponentTypes {
	if o == nil {
		var ret MessageComponentTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FileUploadComponentForModalRequest) GetTypeOk() (*MessageComponentTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FileUploadComponentForModalRequest) SetType(v MessageComponentTypes) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileUploadComponentForModalRequest) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadComponentForModalRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileUploadComponentForModalRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FileUploadComponentForModalRequest) SetId(v int32) {
	o.Id = &v
}


// GetCustomId returns the CustomId field value
func (o *FileUploadComponentForModalRequest) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *FileUploadComponentForModalRequest) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *FileUploadComponentForModalRequest) SetCustomId(v string) {
	o.CustomId = v
}

// GetMinValues returns the MinValues field value if set, zero value otherwise.
func (o *FileUploadComponentForModalRequest) GetMinValues() int32 {
	if o == nil || IsNil(o.MinValues) {
		var ret int32
		return ret
	}
	return *o.MinValues
}

// GetMinValuesOk returns a tuple with the MinValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadComponentForModalRequest) GetMinValuesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinValues) {
		return nil, false
	}
	return o.MinValues, true
}

// HasMinValues returns a boolean if a field has been set.
func (o *FileUploadComponentForModalRequest) HasMinValues() bool {
	if o != nil && !IsNil(o.MinValues) {
		return true
	}

	return false
}

// SetMinValues gets a reference to the given int32 and assigns it to the MinValues field.
func (o *FileUploadComponentForModalRequest) SetMinValues(v int32) {
	o.MinValues = &v
}


// GetMaxValues returns the MaxValues field value if set, zero value otherwise.
func (o *FileUploadComponentForModalRequest) GetMaxValues() int32 {
	if o == nil || IsNil(o.MaxValues) {
		var ret int32
		return ret
	}
	return *o.MaxValues
}

// GetMaxValuesOk returns a tuple with the MaxValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadComponentForModalRequest) GetMaxValuesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxValues) {
		return nil, false
	}
	return o.MaxValues, true
}

// HasMaxValues returns a boolean if a field has been set.
func (o *FileUploadComponentForModalRequest) HasMaxValues() bool {
	if o != nil && !IsNil(o.MaxValues) {
		return true
	}

	return false
}

// SetMaxValues gets a reference to the given int32 and assigns it to the MaxValues field.
func (o *FileUploadComponentForModalRequest) SetMaxValues(v int32) {
	o.MaxValues = &v
}


// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FileUploadComponentForModalRequest) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadComponentForModalRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FileUploadComponentForModalRequest) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FileUploadComponentForModalRequest) SetRequired(v bool) {
	o.Required = &v
}


func (o FileUploadComponentForModalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileUploadComponentForModalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["custom_id"] = o.CustomId
	if !IsNil(o.MinValues) {
		toSerialize["min_values"] = o.MinValues
	}
	if !IsNil(o.MaxValues) {
		toSerialize["max_values"] = o.MaxValues
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

func (o *FileUploadComponentForModalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"custom_id",
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

	varFileUploadComponentForModalRequest := _FileUploadComponentForModalRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileUploadComponentForModalRequest)

	if err != nil {
		return err
	}

	*o = FileUploadComponentForModalRequest(varFileUploadComponentForModalRequest)

	return err
}

type NullableFileUploadComponentForModalRequest struct {
	value *FileUploadComponentForModalRequest
	isSet bool
}

func (v NullableFileUploadComponentForModalRequest) Get() *FileUploadComponentForModalRequest {
	return v.value
}

func (v *NullableFileUploadComponentForModalRequest) Set(val *FileUploadComponentForModalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileUploadComponentForModalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileUploadComponentForModalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileUploadComponentForModalRequest(val *FileUploadComponentForModalRequest) *NullableFileUploadComponentForModalRequest {
	return &NullableFileUploadComponentForModalRequest{value: val, isSet: true}
}

func (v NullableFileUploadComponentForModalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileUploadComponentForModalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

