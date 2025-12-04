/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-12-04T03:54:15.906780824Z[Etc/UTC]
 *    * - **Generator Version**: 7.17.0
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

// checks if the BulkBanUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkBanUsersRequest{}

// BulkBanUsersRequest struct for BulkBanUsersRequest
type BulkBanUsersRequest struct {
	UserIds []string `json:"user_ids"`
	DeleteMessageSeconds NullableInt32 `json:"delete_message_seconds,omitempty"`
}

type _BulkBanUsersRequest BulkBanUsersRequest

// NewBulkBanUsersRequest instantiates a new BulkBanUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkBanUsersRequest(userIds []string) *BulkBanUsersRequest {
	this := BulkBanUsersRequest{}
	this.UserIds = userIds
	return &this
}

// NewBulkBanUsersRequestWithDefaults instantiates a new BulkBanUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkBanUsersRequestWithDefaults() *BulkBanUsersRequest {
	this := BulkBanUsersRequest{}
	return &this
}

// GetUserIds returns the UserIds field value
func (o *BulkBanUsersRequest) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value
// and a boolean to check if the value has been set.
func (o *BulkBanUsersRequest) GetUserIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserIds, true
}

// SetUserIds sets field value
func (o *BulkBanUsersRequest) SetUserIds(v []string) {
	o.UserIds = v
}

// GetDeleteMessageSeconds returns the DeleteMessageSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkBanUsersRequest) GetDeleteMessageSeconds() int32 {
	if o == nil || IsNil(o.DeleteMessageSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.DeleteMessageSeconds.Get()
}

// GetDeleteMessageSecondsOk returns a tuple with the DeleteMessageSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkBanUsersRequest) GetDeleteMessageSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.DeleteMessageSeconds.Get()) {
		return nil, false
	}
	return o.DeleteMessageSeconds.Get(), o.DeleteMessageSeconds.IsSet()
}

// HasDeleteMessageSeconds returns a boolean if a field has been set.
func (o *BulkBanUsersRequest) HasDeleteMessageSeconds() bool {
	if o != nil && o.DeleteMessageSeconds.IsSet() {
		return true
	}

	return false
}

// SetDeleteMessageSeconds gets a reference to the given NullableInt32 and assigns it to the DeleteMessageSeconds field.
func (o *BulkBanUsersRequest) SetDeleteMessageSeconds(v int32) {
	o.DeleteMessageSeconds.Set(&v)
}

// SetDeleteMessageSecondsNil sets the value for DeleteMessageSeconds to be an explicit nil
func (o *BulkBanUsersRequest) SetDeleteMessageSecondsNil() {
	o.DeleteMessageSeconds.Set(nil)
}

// UnsetDeleteMessageSeconds ensures that no value is present for DeleteMessageSeconds, not even an explicit nil
func (o *BulkBanUsersRequest) UnsetDeleteMessageSeconds() {
	o.DeleteMessageSeconds.Unset()
}

func (o BulkBanUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkBanUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_ids"] = o.UserIds
	if o.DeleteMessageSeconds.IsSet() {
		toSerialize["delete_message_seconds"] = o.DeleteMessageSeconds.Get()
	}
	return toSerialize, nil
}

func (o *BulkBanUsersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_ids",
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

	varBulkBanUsersRequest := _BulkBanUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkBanUsersRequest)

	if err != nil {
		return err
	}

	*o = BulkBanUsersRequest(varBulkBanUsersRequest)

	return err
}

type NullableBulkBanUsersRequest struct {
	value *BulkBanUsersRequest
	isSet bool
}

func (v NullableBulkBanUsersRequest) Get() *BulkBanUsersRequest {
	return v.value
}

func (v *NullableBulkBanUsersRequest) Set(val *BulkBanUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkBanUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkBanUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkBanUsersRequest(val *BulkBanUsersRequest) *NullableBulkBanUsersRequest {
	return &NullableBulkBanUsersRequest{value: val, isSet: true}
}

func (v NullableBulkBanUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkBanUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

