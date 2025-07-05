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

// checks if the FileComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileComponentResponse{}

// FileComponentResponse struct for FileComponentResponse
type FileComponentResponse struct {
	Type int32 `json:"type"`
	Id int32 `json:"id"`
	File UnfurledMediaResponse `json:"file"`
	Name NullableString `json:"name,omitempty"`
	Size NullableInt32 `json:"size,omitempty"`
	Spoiler bool `json:"spoiler"`
}

type _FileComponentResponse FileComponentResponse

// NewFileComponentResponse instantiates a new FileComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileComponentResponse(type_ int32, id int32, file UnfurledMediaResponse, spoiler bool) *FileComponentResponse {
	this := FileComponentResponse{}
	this.Type = type_
	this.Id = id
	this.File = file
	this.Spoiler = spoiler
	return &this
}

// NewFileComponentResponseWithDefaults instantiates a new FileComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileComponentResponseWithDefaults() *FileComponentResponse {
	this := FileComponentResponse{}
	return &this
}

// GetType returns the Type field value
func (o *FileComponentResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FileComponentResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FileComponentResponse) SetType(v int32) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FileComponentResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileComponentResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileComponentResponse) SetId(v int32) {
	o.Id = v
}

// GetFile returns the File field value
func (o *FileComponentResponse) GetFile() UnfurledMediaResponse {
	if o == nil {
		var ret UnfurledMediaResponse
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *FileComponentResponse) GetFileOk() (*UnfurledMediaResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *FileComponentResponse) SetFile(v UnfurledMediaResponse) {
	o.File = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileComponentResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileComponentResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FileComponentResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FileComponentResponse) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *FileComponentResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FileComponentResponse) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileComponentResponse) GetSize() int32 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileComponentResponse) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size.Get()) {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *FileComponentResponse) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *FileComponentResponse) SetSize(v int32) {
	o.Size.Set(&v)
}

// SetSizeNil sets the value for Size to be an explicit nil
func (o *FileComponentResponse) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *FileComponentResponse) UnsetSize() {
	o.Size.Unset()
}

// GetSpoiler returns the Spoiler field value
func (o *FileComponentResponse) GetSpoiler() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Spoiler
}

// GetSpoilerOk returns a tuple with the Spoiler field value
// and a boolean to check if the value has been set.
func (o *FileComponentResponse) GetSpoilerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spoiler, true
}

// SetSpoiler sets field value
func (o *FileComponentResponse) SetSpoiler(v bool) {
	o.Spoiler = v
}

func (o FileComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["file"] = o.File
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	toSerialize["spoiler"] = o.Spoiler
	return toSerialize, nil
}

func (o *FileComponentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"file",
		"spoiler",
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

	varFileComponentResponse := _FileComponentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileComponentResponse)

	if err != nil {
		return err
	}

	*o = FileComponentResponse(varFileComponentResponse)

	return err
}

type NullableFileComponentResponse struct {
	value *FileComponentResponse
	isSet bool
}

func (v NullableFileComponentResponse) Get() *FileComponentResponse {
	return v.value
}

func (v *NullableFileComponentResponse) Set(val *FileComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileComponentResponse(val *FileComponentResponse) *NullableFileComponentResponse {
	return &NullableFileComponentResponse{value: val, isSet: true}
}

func (v NullableFileComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


