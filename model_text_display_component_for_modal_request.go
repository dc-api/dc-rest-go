/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-09-30T01:39:44.165564220Z[Etc/UTC]
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

// checks if the TextDisplayComponentForModalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextDisplayComponentForModalRequest{}

// TextDisplayComponentForModalRequest struct for TextDisplayComponentForModalRequest
type TextDisplayComponentForModalRequest struct {
	Type int32 `json:"type"`
	Id NullableInt32 `json:"id,omitempty"`
	Content string `json:"content"`
}

type _TextDisplayComponentForModalRequest TextDisplayComponentForModalRequest

// NewTextDisplayComponentForModalRequest instantiates a new TextDisplayComponentForModalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextDisplayComponentForModalRequest(type_ int32, content string) *TextDisplayComponentForModalRequest {
	this := TextDisplayComponentForModalRequest{}
	this.Type = type_
	this.Content = content
	return &this
}

// NewTextDisplayComponentForModalRequestWithDefaults instantiates a new TextDisplayComponentForModalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextDisplayComponentForModalRequestWithDefaults() *TextDisplayComponentForModalRequest {
	this := TextDisplayComponentForModalRequest{}
	return &this
}

// GetType returns the Type field value
func (o *TextDisplayComponentForModalRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextDisplayComponentForModalRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TextDisplayComponentForModalRequest) SetType(v int32) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextDisplayComponentForModalRequest) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextDisplayComponentForModalRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id.Get()) {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TextDisplayComponentForModalRequest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *TextDisplayComponentForModalRequest) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *TextDisplayComponentForModalRequest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TextDisplayComponentForModalRequest) UnsetId() {
	o.Id.Unset()
}

// GetContent returns the Content field value
func (o *TextDisplayComponentForModalRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *TextDisplayComponentForModalRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *TextDisplayComponentForModalRequest) SetContent(v string) {
	o.Content = v
}

func (o TextDisplayComponentForModalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextDisplayComponentForModalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

func (o *TextDisplayComponentForModalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"content",
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

	varTextDisplayComponentForModalRequest := _TextDisplayComponentForModalRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTextDisplayComponentForModalRequest)

	if err != nil {
		return err
	}

	*o = TextDisplayComponentForModalRequest(varTextDisplayComponentForModalRequest)

	return err
}

type NullableTextDisplayComponentForModalRequest struct {
	value *TextDisplayComponentForModalRequest
	isSet bool
}

func (v NullableTextDisplayComponentForModalRequest) Get() *TextDisplayComponentForModalRequest {
	return v.value
}

func (v *NullableTextDisplayComponentForModalRequest) Set(val *TextDisplayComponentForModalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTextDisplayComponentForModalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTextDisplayComponentForModalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextDisplayComponentForModalRequest(val *TextDisplayComponentForModalRequest) *NullableTextDisplayComponentForModalRequest {
	return &NullableTextDisplayComponentForModalRequest{value: val, isSet: true}
}

func (v NullableTextDisplayComponentForModalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextDisplayComponentForModalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


