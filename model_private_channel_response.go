/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-08-08T14:09:23.736426080Z[Etc/UTC]
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
	"time"
	"bytes"
	"fmt"
)

// checks if the PrivateChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateChannelResponse{}

// PrivateChannelResponse struct for PrivateChannelResponse
type PrivateChannelResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type int32 `json:"type"`
	LastMessageId *string `json:"last_message_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Flags int32 `json:"flags"`
	LastPinTimestamp NullableTime `json:"last_pin_timestamp,omitempty"`
	Recipients []UserResponse `json:"recipients"`
}

type _PrivateChannelResponse PrivateChannelResponse

// NewPrivateChannelResponse instantiates a new PrivateChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateChannelResponse(id string, type_ int32, flags int32, recipients []UserResponse) *PrivateChannelResponse {
	this := PrivateChannelResponse{}
	this.Id = id
	this.Type = type_
	this.Flags = flags
	this.Recipients = recipients
	return &this
}

// NewPrivateChannelResponseWithDefaults instantiates a new PrivateChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateChannelResponseWithDefaults() *PrivateChannelResponse {
	this := PrivateChannelResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PrivateChannelResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrivateChannelResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrivateChannelResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *PrivateChannelResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PrivateChannelResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PrivateChannelResponse) SetType(v int32) {
	o.Type = v
}

// GetLastMessageId returns the LastMessageId field value if set, zero value otherwise.
func (o *PrivateChannelResponse) GetLastMessageId() string {
	if o == nil || IsNil(o.LastMessageId) {
		var ret string
		return ret
	}
	return *o.LastMessageId
}

// GetLastMessageIdOk returns a tuple with the LastMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateChannelResponse) GetLastMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastMessageId) {
		return nil, false
	}
	return o.LastMessageId, true
}

// HasLastMessageId returns a boolean if a field has been set.
func (o *PrivateChannelResponse) HasLastMessageId() bool {
	if o != nil && !IsNil(o.LastMessageId) {
		return true
	}

	return false
}

// SetLastMessageId gets a reference to the given string and assigns it to the LastMessageId field.
func (o *PrivateChannelResponse) SetLastMessageId(v string) {
	o.LastMessageId = &v
}


// GetFlags returns the Flags field value
func (o *PrivateChannelResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *PrivateChannelResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *PrivateChannelResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetLastPinTimestamp returns the LastPinTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateChannelResponse) GetLastPinTimestamp() time.Time {
	if o == nil || IsNil(o.LastPinTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastPinTimestamp.Get()
}

// GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateChannelResponse) GetLastPinTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPinTimestamp.Get()) {
		return nil, false
	}
	return o.LastPinTimestamp.Get(), o.LastPinTimestamp.IsSet()
}

// HasLastPinTimestamp returns a boolean if a field has been set.
func (o *PrivateChannelResponse) HasLastPinTimestamp() bool {
	if o != nil && o.LastPinTimestamp.IsSet() {
		return true
	}

	return false
}

// SetLastPinTimestamp gets a reference to the given NullableTime and assigns it to the LastPinTimestamp field.
func (o *PrivateChannelResponse) SetLastPinTimestamp(v time.Time) {
	o.LastPinTimestamp.Set(&v)
}

// SetLastPinTimestampNil sets the value for LastPinTimestamp to be an explicit nil
func (o *PrivateChannelResponse) SetLastPinTimestampNil() {
	o.LastPinTimestamp.Set(nil)
}

// UnsetLastPinTimestamp ensures that no value is present for LastPinTimestamp, not even an explicit nil
func (o *PrivateChannelResponse) UnsetLastPinTimestamp() {
	o.LastPinTimestamp.Unset()
}

// GetRecipients returns the Recipients field value
func (o *PrivateChannelResponse) GetRecipients() []UserResponse {
	if o == nil {
		var ret []UserResponse
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *PrivateChannelResponse) GetRecipientsOk() ([]UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *PrivateChannelResponse) SetRecipients(v []UserResponse) {
	o.Recipients = v
}

func (o PrivateChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.LastMessageId) {
		toSerialize["last_message_id"] = o.LastMessageId
	}
	toSerialize["flags"] = o.Flags
	if o.LastPinTimestamp.IsSet() {
		toSerialize["last_pin_timestamp"] = o.LastPinTimestamp.Get()
	}
	toSerialize["recipients"] = o.Recipients
	return toSerialize, nil
}

func (o *PrivateChannelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"flags",
		"recipients",
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

	varPrivateChannelResponse := _PrivateChannelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrivateChannelResponse)

	if err != nil {
		return err
	}

	*o = PrivateChannelResponse(varPrivateChannelResponse)

	return err
}

type NullablePrivateChannelResponse struct {
	value *PrivateChannelResponse
	isSet bool
}

func (v NullablePrivateChannelResponse) Get() *PrivateChannelResponse {
	return v.value
}

func (v *NullablePrivateChannelResponse) Set(val *PrivateChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateChannelResponse(val *PrivateChannelResponse) *NullablePrivateChannelResponse {
	return &NullablePrivateChannelResponse{value: val, isSet: true}
}

func (v NullablePrivateChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


