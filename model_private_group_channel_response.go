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
	"time"
	"bytes"
	"fmt"
)

// checks if the PrivateGroupChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateGroupChannelResponse{}

// PrivateGroupChannelResponse struct for PrivateGroupChannelResponse
type PrivateGroupChannelResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type int32 `json:"type"`
	LastMessageId *string `json:"last_message_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Flags int32 `json:"flags"`
	LastPinTimestamp NullableTime `json:"last_pin_timestamp,omitempty"`
	Recipients []UserResponse `json:"recipients"`
	Name NullableString `json:"name,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	OwnerId *string `json:"owner_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Managed NullableBool `json:"managed,omitempty"`
	ApplicationId *string `json:"application_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _PrivateGroupChannelResponse PrivateGroupChannelResponse

// NewPrivateGroupChannelResponse instantiates a new PrivateGroupChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateGroupChannelResponse(id string, type_ int32, flags int32, recipients []UserResponse) *PrivateGroupChannelResponse {
	this := PrivateGroupChannelResponse{}
	this.Id = id
	this.Type = type_
	this.Flags = flags
	this.Recipients = recipients
	return &this
}

// NewPrivateGroupChannelResponseWithDefaults instantiates a new PrivateGroupChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateGroupChannelResponseWithDefaults() *PrivateGroupChannelResponse {
	this := PrivateGroupChannelResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PrivateGroupChannelResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrivateGroupChannelResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *PrivateGroupChannelResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PrivateGroupChannelResponse) SetType(v int32) {
	o.Type = v
}

// GetLastMessageId returns the LastMessageId field value if set, zero value otherwise.
func (o *PrivateGroupChannelResponse) GetLastMessageId() string {
	if o == nil || IsNil(o.LastMessageId) {
		var ret string
		return ret
	}
	return *o.LastMessageId
}

// GetLastMessageIdOk returns a tuple with the LastMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetLastMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastMessageId) {
		return nil, false
	}
	return o.LastMessageId, true
}

// HasLastMessageId returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasLastMessageId() bool {
	if o != nil && !IsNil(o.LastMessageId) {
		return true
	}

	return false
}

// SetLastMessageId gets a reference to the given string and assigns it to the LastMessageId field.
func (o *PrivateGroupChannelResponse) SetLastMessageId(v string) {
	o.LastMessageId = &v
}


// GetFlags returns the Flags field value
func (o *PrivateGroupChannelResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *PrivateGroupChannelResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetLastPinTimestamp returns the LastPinTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateGroupChannelResponse) GetLastPinTimestamp() time.Time {
	if o == nil || IsNil(o.LastPinTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastPinTimestamp.Get()
}

// GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateGroupChannelResponse) GetLastPinTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPinTimestamp.Get()) {
		return nil, false
	}
	return o.LastPinTimestamp.Get(), o.LastPinTimestamp.IsSet()
}

// HasLastPinTimestamp returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasLastPinTimestamp() bool {
	if o != nil && o.LastPinTimestamp.IsSet() {
		return true
	}

	return false
}

// SetLastPinTimestamp gets a reference to the given NullableTime and assigns it to the LastPinTimestamp field.
func (o *PrivateGroupChannelResponse) SetLastPinTimestamp(v time.Time) {
	o.LastPinTimestamp.Set(&v)
}

// SetLastPinTimestampNil sets the value for LastPinTimestamp to be an explicit nil
func (o *PrivateGroupChannelResponse) SetLastPinTimestampNil() {
	o.LastPinTimestamp.Set(nil)
}

// UnsetLastPinTimestamp ensures that no value is present for LastPinTimestamp, not even an explicit nil
func (o *PrivateGroupChannelResponse) UnsetLastPinTimestamp() {
	o.LastPinTimestamp.Unset()
}

// GetRecipients returns the Recipients field value
func (o *PrivateGroupChannelResponse) GetRecipients() []UserResponse {
	if o == nil {
		var ret []UserResponse
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetRecipientsOk() ([]UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *PrivateGroupChannelResponse) SetRecipients(v []UserResponse) {
	o.Recipients = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateGroupChannelResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateGroupChannelResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PrivateGroupChannelResponse) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PrivateGroupChannelResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PrivateGroupChannelResponse) UnsetName() {
	o.Name.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateGroupChannelResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateGroupChannelResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *PrivateGroupChannelResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *PrivateGroupChannelResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *PrivateGroupChannelResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *PrivateGroupChannelResponse) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *PrivateGroupChannelResponse) SetOwnerId(v string) {
	o.OwnerId = &v
}


// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateGroupChannelResponse) GetManaged() bool {
	if o == nil || IsNil(o.Managed.Get()) {
		var ret bool
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateGroupChannelResponse) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed.Get()) {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableBool and assigns it to the Managed field.
func (o *PrivateGroupChannelResponse) SetManaged(v bool) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *PrivateGroupChannelResponse) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *PrivateGroupChannelResponse) UnsetManaged() {
	o.Managed.Unset()
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *PrivateGroupChannelResponse) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *PrivateGroupChannelResponse) SetApplicationId(v string) {
	o.ApplicationId = &v
}


func (o PrivateGroupChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateGroupChannelResponse) ToMap() (map[string]interface{}, error) {
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
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	return toSerialize, nil
}

func (o *PrivateGroupChannelResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPrivateGroupChannelResponse := _PrivateGroupChannelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrivateGroupChannelResponse)

	if err != nil {
		return err
	}

	*o = PrivateGroupChannelResponse(varPrivateGroupChannelResponse)

	return err
}

type NullablePrivateGroupChannelResponse struct {
	value *PrivateGroupChannelResponse
	isSet bool
}

func (v NullablePrivateGroupChannelResponse) Get() *PrivateGroupChannelResponse {
	return v.value
}

func (v *NullablePrivateGroupChannelResponse) Set(val *PrivateGroupChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateGroupChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateGroupChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateGroupChannelResponse(val *PrivateGroupChannelResponse) *NullablePrivateGroupChannelResponse {
	return &NullablePrivateGroupChannelResponse{value: val, isSet: true}
}

func (v NullablePrivateGroupChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateGroupChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


