/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
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
	"time"
	"bytes"
	"fmt"
)

// checks if the PrivateGroupChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateGroupChannelResponse{}

// PrivateGroupChannelResponse struct for PrivateGroupChannelResponse
type PrivateGroupChannelResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type ChannelTypes `json:"type"`
	LastMessageId *GetEntitlementsSkuIdsParameterOneOfInner `json:"last_message_id,omitempty"`
	Flags int32 `json:"flags"`
	LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`
	Recipients []UserResponse `json:"recipients"`
	Name *string `json:"name,omitempty"`
	Icon *string `json:"icon,omitempty"`
	OwnerId string `json:"owner_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Managed *bool `json:"managed,omitempty"`
	ApplicationId *string `json:"application_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _PrivateGroupChannelResponse PrivateGroupChannelResponse

// NewPrivateGroupChannelResponse instantiates a new PrivateGroupChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateGroupChannelResponse(id string, type_ ChannelTypes, flags int32, recipients []UserResponse, ownerId string) *PrivateGroupChannelResponse {
	this := PrivateGroupChannelResponse{}
	this.Id = id
	this.Type = type_
	this.Flags = flags
	this.Recipients = recipients
	this.OwnerId = ownerId
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
func (o *PrivateGroupChannelResponse) GetType() ChannelTypes {
	if o == nil {
		var ret ChannelTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetTypeOk() (*ChannelTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PrivateGroupChannelResponse) SetType(v ChannelTypes) {
	o.Type = v
}

// GetLastMessageId returns the LastMessageId field value if set, zero value otherwise.
func (o *PrivateGroupChannelResponse) GetLastMessageId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.LastMessageId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.LastMessageId
}

// GetLastMessageIdOk returns a tuple with the LastMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetLastMessageIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetLastMessageId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the LastMessageId field.
func (o *PrivateGroupChannelResponse) SetLastMessageId(v GetEntitlementsSkuIdsParameterOneOfInner) {
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

// GetLastPinTimestamp returns the LastPinTimestamp field value if set, zero value otherwise.
func (o *PrivateGroupChannelResponse) GetLastPinTimestamp() time.Time {
	if o == nil || IsNil(o.LastPinTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastPinTimestamp
}

// GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetLastPinTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPinTimestamp) {
		return nil, false
	}
	return o.LastPinTimestamp, true
}

// HasLastPinTimestamp returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasLastPinTimestamp() bool {
	if o != nil && !IsNil(o.LastPinTimestamp) {
		return true
	}

	return false
}

// SetLastPinTimestamp gets a reference to the given time.Time and assigns it to the LastPinTimestamp field.
func (o *PrivateGroupChannelResponse) SetLastPinTimestamp(v time.Time) {
	o.LastPinTimestamp = &v
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

// GetName returns the Name field value if set, zero value otherwise.
func (o *PrivateGroupChannelResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrivateGroupChannelResponse) SetName(v string) {
	o.Name = &v
}


// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *PrivateGroupChannelResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *PrivateGroupChannelResponse) SetIcon(v string) {
	o.Icon = &v
}


// GetOwnerId returns the OwnerId field value
func (o *PrivateGroupChannelResponse) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *PrivateGroupChannelResponse) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *PrivateGroupChannelResponse) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGroupChannelResponse) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *PrivateGroupChannelResponse) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *PrivateGroupChannelResponse) SetManaged(v bool) {
	o.Managed = &v
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
	if !IsNil(o.LastPinTimestamp) {
		toSerialize["last_pin_timestamp"] = o.LastPinTimestamp
	}
	toSerialize["recipients"] = o.Recipients
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["owner_id"] = o.OwnerId
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
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
		"owner_id",
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

