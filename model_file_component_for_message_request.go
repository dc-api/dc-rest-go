/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
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

// checks if the FileComponentForMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileComponentForMessageRequest{}

// FileComponentForMessageRequest struct for FileComponentForMessageRequest
type FileComponentForMessageRequest struct {
	Type int32 `json:"type"`
	Spoiler NullableBool `json:"spoiler,omitempty"`
	File UnfurledMediaRequestWithAttachmentReferenceRequired `json:"file"`
}

type _FileComponentForMessageRequest FileComponentForMessageRequest

// NewFileComponentForMessageRequest instantiates a new FileComponentForMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileComponentForMessageRequest(type_ int32, file UnfurledMediaRequestWithAttachmentReferenceRequired) *FileComponentForMessageRequest {
	this := FileComponentForMessageRequest{}
	this.Type = type_
	this.File = file
	return &this
}

// NewFileComponentForMessageRequestWithDefaults instantiates a new FileComponentForMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileComponentForMessageRequestWithDefaults() *FileComponentForMessageRequest {
	this := FileComponentForMessageRequest{}
	return &this
}

// GetType returns the Type field value
func (o *FileComponentForMessageRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FileComponentForMessageRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FileComponentForMessageRequest) SetType(v int32) {
	o.Type = v
}

// GetSpoiler returns the Spoiler field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileComponentForMessageRequest) GetSpoiler() bool {
	if o == nil || IsNil(o.Spoiler.Get()) {
		var ret bool
		return ret
	}
	return *o.Spoiler.Get()
}

// GetSpoilerOk returns a tuple with the Spoiler field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileComponentForMessageRequest) GetSpoilerOk() (*bool, bool) {
	if o == nil || IsNil(o.Spoiler.Get()) {
		return nil, false
	}
	return o.Spoiler.Get(), o.Spoiler.IsSet()
}

// HasSpoiler returns a boolean if a field has been set.
func (o *FileComponentForMessageRequest) HasSpoiler() bool {
	if o != nil && o.Spoiler.IsSet() {
		return true
	}

	return false
}

// SetSpoiler gets a reference to the given NullableBool and assigns it to the Spoiler field.
func (o *FileComponentForMessageRequest) SetSpoiler(v bool) {
	o.Spoiler.Set(&v)
}

// SetSpoilerNil sets the value for Spoiler to be an explicit nil
func (o *FileComponentForMessageRequest) SetSpoilerNil() {
	o.Spoiler.Set(nil)
}

// UnsetSpoiler ensures that no value is present for Spoiler, not even an explicit nil
func (o *FileComponentForMessageRequest) UnsetSpoiler() {
	o.Spoiler.Unset()
}

// GetFile returns the File field value
func (o *FileComponentForMessageRequest) GetFile() UnfurledMediaRequestWithAttachmentReferenceRequired {
	if o == nil {
		var ret UnfurledMediaRequestWithAttachmentReferenceRequired
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *FileComponentForMessageRequest) GetFileOk() (*UnfurledMediaRequestWithAttachmentReferenceRequired, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *FileComponentForMessageRequest) SetFile(v UnfurledMediaRequestWithAttachmentReferenceRequired) {
	o.File = v
}

func (o FileComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileComponentForMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Spoiler.IsSet() {
		toSerialize["spoiler"] = o.Spoiler.Get()
	}
	toSerialize["file"] = o.File
	return toSerialize, nil
}

func (o *FileComponentForMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"file",
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

	varFileComponentForMessageRequest := _FileComponentForMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileComponentForMessageRequest)

	if err != nil {
		return err
	}

	*o = FileComponentForMessageRequest(varFileComponentForMessageRequest)

	return err
}

type NullableFileComponentForMessageRequest struct {
	value *FileComponentForMessageRequest
	isSet bool
}

func (v NullableFileComponentForMessageRequest) Get() *FileComponentForMessageRequest {
	return v.value
}

func (v *NullableFileComponentForMessageRequest) Set(val *FileComponentForMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileComponentForMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileComponentForMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileComponentForMessageRequest(val *FileComponentForMessageRequest) *NullableFileComponentForMessageRequest {
	return &NullableFileComponentForMessageRequest{value: val, isSet: true}
}

func (v NullableFileComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileComponentForMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


