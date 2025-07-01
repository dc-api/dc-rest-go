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

// checks if the BulkLobbyMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkLobbyMemberRequest{}

// BulkLobbyMemberRequest struct for BulkLobbyMemberRequest
type BulkLobbyMemberRequest struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Flags NullableInt32 `json:"flags,omitempty"`
	RemoveMember NullableBool `json:"remove_member,omitempty"`
}

type _BulkLobbyMemberRequest BulkLobbyMemberRequest

// NewBulkLobbyMemberRequest instantiates a new BulkLobbyMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkLobbyMemberRequest(id string) *BulkLobbyMemberRequest {
	this := BulkLobbyMemberRequest{}
	this.Id = id
	return &this
}

// NewBulkLobbyMemberRequestWithDefaults instantiates a new BulkLobbyMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkLobbyMemberRequestWithDefaults() *BulkLobbyMemberRequest {
	this := BulkLobbyMemberRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkLobbyMemberRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkLobbyMemberRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkLobbyMemberRequest) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BulkLobbyMemberRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkLobbyMemberRequest) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BulkLobbyMemberRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *BulkLobbyMemberRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}


// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkLobbyMemberRequest) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkLobbyMemberRequest) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags.Get()) {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *BulkLobbyMemberRequest) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *BulkLobbyMemberRequest) SetFlags(v int32) {
	o.Flags.Set(&v)
}

// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *BulkLobbyMemberRequest) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *BulkLobbyMemberRequest) UnsetFlags() {
	o.Flags.Unset()
}

// GetRemoveMember returns the RemoveMember field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkLobbyMemberRequest) GetRemoveMember() bool {
	if o == nil || IsNil(o.RemoveMember.Get()) {
		var ret bool
		return ret
	}
	return *o.RemoveMember.Get()
}

// GetRemoveMemberOk returns a tuple with the RemoveMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkLobbyMemberRequest) GetRemoveMemberOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveMember.Get()) {
		return nil, false
	}
	return o.RemoveMember.Get(), o.RemoveMember.IsSet()
}

// HasRemoveMember returns a boolean if a field has been set.
func (o *BulkLobbyMemberRequest) HasRemoveMember() bool {
	if o != nil && o.RemoveMember.IsSet() {
		return true
	}

	return false
}

// SetRemoveMember gets a reference to the given NullableBool and assigns it to the RemoveMember field.
func (o *BulkLobbyMemberRequest) SetRemoveMember(v bool) {
	o.RemoveMember.Set(&v)
}

// SetRemoveMemberNil sets the value for RemoveMember to be an explicit nil
func (o *BulkLobbyMemberRequest) SetRemoveMemberNil() {
	o.RemoveMember.Set(nil)
}

// UnsetRemoveMember ensures that no value is present for RemoveMember, not even an explicit nil
func (o *BulkLobbyMemberRequest) UnsetRemoveMember() {
	o.RemoveMember.Unset()
}

func (o BulkLobbyMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkLobbyMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	if o.RemoveMember.IsSet() {
		toSerialize["remove_member"] = o.RemoveMember.Get()
	}
	return toSerialize, nil
}

func (o *BulkLobbyMemberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varBulkLobbyMemberRequest := _BulkLobbyMemberRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkLobbyMemberRequest)

	if err != nil {
		return err
	}

	*o = BulkLobbyMemberRequest(varBulkLobbyMemberRequest)

	return err
}

type NullableBulkLobbyMemberRequest struct {
	value *BulkLobbyMemberRequest
	isSet bool
}

func (v NullableBulkLobbyMemberRequest) Get() *BulkLobbyMemberRequest {
	return v.value
}

func (v *NullableBulkLobbyMemberRequest) Set(val *BulkLobbyMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkLobbyMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkLobbyMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkLobbyMemberRequest(val *BulkLobbyMemberRequest) *NullableBulkLobbyMemberRequest {
	return &NullableBulkLobbyMemberRequest{value: val, isSet: true}
}

func (v NullableBulkLobbyMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkLobbyMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


