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

// checks if the StandardStickerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardStickerResponse{}

// StandardStickerResponse struct for StandardStickerResponse
type StandardStickerResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Tags string `json:"tags"`
	Type int32 `json:"type"`
	FormatType NullableInt32 `json:"format_type,omitempty"`
	Description NullableString `json:"description,omitempty"`
	PackId string `json:"pack_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SortValue int32 `json:"sort_value"`
}

type _StandardStickerResponse StandardStickerResponse

// NewStandardStickerResponse instantiates a new StandardStickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardStickerResponse(id string, name string, tags string, type_ int32, packId string, sortValue int32) *StandardStickerResponse {
	this := StandardStickerResponse{}
	this.Id = id
	this.Name = name
	this.Tags = tags
	this.Type = type_
	this.PackId = packId
	this.SortValue = sortValue
	return &this
}

// NewStandardStickerResponseWithDefaults instantiates a new StandardStickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardStickerResponseWithDefaults() *StandardStickerResponse {
	this := StandardStickerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *StandardStickerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StandardStickerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StandardStickerResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *StandardStickerResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StandardStickerResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StandardStickerResponse) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value
func (o *StandardStickerResponse) GetTags() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *StandardStickerResponse) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *StandardStickerResponse) SetTags(v string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *StandardStickerResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StandardStickerResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StandardStickerResponse) SetType(v int32) {
	o.Type = v
}

// GetFormatType returns the FormatType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandardStickerResponse) GetFormatType() int32 {
	if o == nil || IsNil(o.FormatType.Get()) {
		var ret int32
		return ret
	}
	return *o.FormatType.Get()
}

// GetFormatTypeOk returns a tuple with the FormatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandardStickerResponse) GetFormatTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.FormatType.Get()) {
		return nil, false
	}
	return o.FormatType.Get(), o.FormatType.IsSet()
}

// HasFormatType returns a boolean if a field has been set.
func (o *StandardStickerResponse) HasFormatType() bool {
	if o != nil && o.FormatType.IsSet() {
		return true
	}

	return false
}

// SetFormatType gets a reference to the given NullableInt32 and assigns it to the FormatType field.
func (o *StandardStickerResponse) SetFormatType(v int32) {
	o.FormatType.Set(&v)
}

// SetFormatTypeNil sets the value for FormatType to be an explicit nil
func (o *StandardStickerResponse) SetFormatTypeNil() {
	o.FormatType.Set(nil)
}

// UnsetFormatType ensures that no value is present for FormatType, not even an explicit nil
func (o *StandardStickerResponse) UnsetFormatType() {
	o.FormatType.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandardStickerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandardStickerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StandardStickerResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StandardStickerResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StandardStickerResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StandardStickerResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetPackId returns the PackId field value
func (o *StandardStickerResponse) GetPackId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackId
}

// GetPackIdOk returns a tuple with the PackId field value
// and a boolean to check if the value has been set.
func (o *StandardStickerResponse) GetPackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackId, true
}

// SetPackId sets field value
func (o *StandardStickerResponse) SetPackId(v string) {
	o.PackId = v
}

// GetSortValue returns the SortValue field value
func (o *StandardStickerResponse) GetSortValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortValue
}

// GetSortValueOk returns a tuple with the SortValue field value
// and a boolean to check if the value has been set.
func (o *StandardStickerResponse) GetSortValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortValue, true
}

// SetSortValue sets field value
func (o *StandardStickerResponse) SetSortValue(v int32) {
	o.SortValue = v
}

func (o StandardStickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardStickerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type
	if o.FormatType.IsSet() {
		toSerialize["format_type"] = o.FormatType.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["pack_id"] = o.PackId
	toSerialize["sort_value"] = o.SortValue
	return toSerialize, nil
}

func (o *StandardStickerResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"tags",
		"type",
		"pack_id",
		"sort_value",
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

	varStandardStickerResponse := _StandardStickerResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStandardStickerResponse)

	if err != nil {
		return err
	}

	*o = StandardStickerResponse(varStandardStickerResponse)

	return err
}

type NullableStandardStickerResponse struct {
	value *StandardStickerResponse
	isSet bool
}

func (v NullableStandardStickerResponse) Get() *StandardStickerResponse {
	return v.value
}

func (v *NullableStandardStickerResponse) Set(val *StandardStickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardStickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardStickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardStickerResponse(val *StandardStickerResponse) *NullableStandardStickerResponse {
	return &NullableStandardStickerResponse{value: val, isSet: true}
}

func (v NullableStandardStickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardStickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


