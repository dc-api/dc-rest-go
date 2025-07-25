/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// checks if the LobbyMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LobbyMemberRequest{}

// LobbyMemberRequest struct for LobbyMemberRequest
type LobbyMemberRequest struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Flags NullableInt32 `json:"flags,omitempty"`
}

type _LobbyMemberRequest LobbyMemberRequest

// NewLobbyMemberRequest instantiates a new LobbyMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLobbyMemberRequest(id string) *LobbyMemberRequest {
	this := LobbyMemberRequest{}
	this.Id = id
	return &this
}

// NewLobbyMemberRequestWithDefaults instantiates a new LobbyMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLobbyMemberRequestWithDefaults() *LobbyMemberRequest {
	this := LobbyMemberRequest{}
	return &this
}

// GetId returns the Id field value
func (o *LobbyMemberRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LobbyMemberRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LobbyMemberRequest) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LobbyMemberRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LobbyMemberRequest) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LobbyMemberRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *LobbyMemberRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}


// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LobbyMemberRequest) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LobbyMemberRequest) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags.Get()) {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *LobbyMemberRequest) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *LobbyMemberRequest) SetFlags(v int32) {
	o.Flags.Set(&v)
}

// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *LobbyMemberRequest) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *LobbyMemberRequest) UnsetFlags() {
	o.Flags.Unset()
}

func (o LobbyMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LobbyMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	return toSerialize, nil
}

func (o *LobbyMemberRequest) UnmarshalJSON(data []byte) (err error) {
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

	varLobbyMemberRequest := _LobbyMemberRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLobbyMemberRequest)

	if err != nil {
		return err
	}

	*o = LobbyMemberRequest(varLobbyMemberRequest)

	return err
}

type NullableLobbyMemberRequest struct {
	value *LobbyMemberRequest
	isSet bool
}

func (v NullableLobbyMemberRequest) Get() *LobbyMemberRequest {
	return v.value
}

func (v *NullableLobbyMemberRequest) Set(val *LobbyMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLobbyMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLobbyMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLobbyMemberRequest(val *LobbyMemberRequest) *NullableLobbyMemberRequest {
	return &NullableLobbyMemberRequest{value: val, isSet: true}
}

func (v NullableLobbyMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLobbyMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


