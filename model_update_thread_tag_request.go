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

// checks if the UpdateThreadTagRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateThreadTagRequest{}

// UpdateThreadTagRequest struct for UpdateThreadTagRequest
type UpdateThreadTagRequest struct {
	Name string `json:"name"`
	EmojiId *string `json:"emoji_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EmojiName NullableString `json:"emoji_name,omitempty"`
	Moderated NullableBool `json:"moderated,omitempty"`
	Id *string `json:"id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _UpdateThreadTagRequest UpdateThreadTagRequest

// NewUpdateThreadTagRequest instantiates a new UpdateThreadTagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateThreadTagRequest(name string) *UpdateThreadTagRequest {
	this := UpdateThreadTagRequest{}
	this.Name = name
	return &this
}

// NewUpdateThreadTagRequestWithDefaults instantiates a new UpdateThreadTagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateThreadTagRequestWithDefaults() *UpdateThreadTagRequest {
	this := UpdateThreadTagRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateThreadTagRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateThreadTagRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateThreadTagRequest) SetName(v string) {
	o.Name = v
}

// GetEmojiId returns the EmojiId field value if set, zero value otherwise.
func (o *UpdateThreadTagRequest) GetEmojiId() string {
	if o == nil || IsNil(o.EmojiId) {
		var ret string
		return ret
	}
	return *o.EmojiId
}

// GetEmojiIdOk returns a tuple with the EmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadTagRequest) GetEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiId) {
		return nil, false
	}
	return o.EmojiId, true
}

// HasEmojiId returns a boolean if a field has been set.
func (o *UpdateThreadTagRequest) HasEmojiId() bool {
	if o != nil && !IsNil(o.EmojiId) {
		return true
	}

	return false
}

// SetEmojiId gets a reference to the given string and assigns it to the EmojiId field.
func (o *UpdateThreadTagRequest) SetEmojiId(v string) {
	o.EmojiId = &v
}


// GetEmojiName returns the EmojiName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadTagRequest) GetEmojiName() string {
	if o == nil || IsNil(o.EmojiName.Get()) {
		var ret string
		return ret
	}
	return *o.EmojiName.Get()
}

// GetEmojiNameOk returns a tuple with the EmojiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadTagRequest) GetEmojiNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiName.Get()) {
		return nil, false
	}
	return o.EmojiName.Get(), o.EmojiName.IsSet()
}

// HasEmojiName returns a boolean if a field has been set.
func (o *UpdateThreadTagRequest) HasEmojiName() bool {
	if o != nil && o.EmojiName.IsSet() {
		return true
	}

	return false
}

// SetEmojiName gets a reference to the given NullableString and assigns it to the EmojiName field.
func (o *UpdateThreadTagRequest) SetEmojiName(v string) {
	o.EmojiName.Set(&v)
}

// SetEmojiNameNil sets the value for EmojiName to be an explicit nil
func (o *UpdateThreadTagRequest) SetEmojiNameNil() {
	o.EmojiName.Set(nil)
}

// UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
func (o *UpdateThreadTagRequest) UnsetEmojiName() {
	o.EmojiName.Unset()
}

// GetModerated returns the Moderated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadTagRequest) GetModerated() bool {
	if o == nil || IsNil(o.Moderated.Get()) {
		var ret bool
		return ret
	}
	return *o.Moderated.Get()
}

// GetModeratedOk returns a tuple with the Moderated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadTagRequest) GetModeratedOk() (*bool, bool) {
	if o == nil || IsNil(o.Moderated.Get()) {
		return nil, false
	}
	return o.Moderated.Get(), o.Moderated.IsSet()
}

// HasModerated returns a boolean if a field has been set.
func (o *UpdateThreadTagRequest) HasModerated() bool {
	if o != nil && o.Moderated.IsSet() {
		return true
	}

	return false
}

// SetModerated gets a reference to the given NullableBool and assigns it to the Moderated field.
func (o *UpdateThreadTagRequest) SetModerated(v bool) {
	o.Moderated.Set(&v)
}

// SetModeratedNil sets the value for Moderated to be an explicit nil
func (o *UpdateThreadTagRequest) SetModeratedNil() {
	o.Moderated.Set(nil)
}

// UnsetModerated ensures that no value is present for Moderated, not even an explicit nil
func (o *UpdateThreadTagRequest) UnsetModerated() {
	o.Moderated.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateThreadTagRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadTagRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateThreadTagRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateThreadTagRequest) SetId(v string) {
	o.Id = &v
}


func (o UpdateThreadTagRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateThreadTagRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.EmojiId) {
		toSerialize["emoji_id"] = o.EmojiId
	}
	if o.EmojiName.IsSet() {
		toSerialize["emoji_name"] = o.EmojiName.Get()
	}
	if o.Moderated.IsSet() {
		toSerialize["moderated"] = o.Moderated.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

func (o *UpdateThreadTagRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateThreadTagRequest := _UpdateThreadTagRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateThreadTagRequest)

	if err != nil {
		return err
	}

	*o = UpdateThreadTagRequest(varUpdateThreadTagRequest)

	return err
}

type NullableUpdateThreadTagRequest struct {
	value *UpdateThreadTagRequest
	isSet bool
}

func (v NullableUpdateThreadTagRequest) Get() *UpdateThreadTagRequest {
	return v.value
}

func (v *NullableUpdateThreadTagRequest) Set(val *UpdateThreadTagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateThreadTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateThreadTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateThreadTagRequest(val *UpdateThreadTagRequest) *NullableUpdateThreadTagRequest {
	return &NullableUpdateThreadTagRequest{value: val, isSet: true}
}

func (v NullableUpdateThreadTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateThreadTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


