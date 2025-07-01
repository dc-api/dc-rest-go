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
)

// checks if the CreateLobbyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLobbyRequest{}

// CreateLobbyRequest struct for CreateLobbyRequest
type CreateLobbyRequest struct {
	IdleTimeoutSeconds NullableInt32 `json:"idle_timeout_seconds,omitempty"`
	Members []LobbyMemberRequest `json:"members,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// NewCreateLobbyRequest instantiates a new CreateLobbyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLobbyRequest() *CreateLobbyRequest {
	this := CreateLobbyRequest{}
	return &this
}

// NewCreateLobbyRequestWithDefaults instantiates a new CreateLobbyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLobbyRequestWithDefaults() *CreateLobbyRequest {
	this := CreateLobbyRequest{}
	return &this
}

// GetIdleTimeoutSeconds returns the IdleTimeoutSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateLobbyRequest) GetIdleTimeoutSeconds() int32 {
	if o == nil || IsNil(o.IdleTimeoutSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.IdleTimeoutSeconds.Get()
}

// GetIdleTimeoutSecondsOk returns a tuple with the IdleTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateLobbyRequest) GetIdleTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.IdleTimeoutSeconds.Get()) {
		return nil, false
	}
	return o.IdleTimeoutSeconds.Get(), o.IdleTimeoutSeconds.IsSet()
}

// HasIdleTimeoutSeconds returns a boolean if a field has been set.
func (o *CreateLobbyRequest) HasIdleTimeoutSeconds() bool {
	if o != nil && o.IdleTimeoutSeconds.IsSet() {
		return true
	}

	return false
}

// SetIdleTimeoutSeconds gets a reference to the given NullableInt32 and assigns it to the IdleTimeoutSeconds field.
func (o *CreateLobbyRequest) SetIdleTimeoutSeconds(v int32) {
	o.IdleTimeoutSeconds.Set(&v)
}

// SetIdleTimeoutSecondsNil sets the value for IdleTimeoutSeconds to be an explicit nil
func (o *CreateLobbyRequest) SetIdleTimeoutSecondsNil() {
	o.IdleTimeoutSeconds.Set(nil)
}

// UnsetIdleTimeoutSeconds ensures that no value is present for IdleTimeoutSeconds, not even an explicit nil
func (o *CreateLobbyRequest) UnsetIdleTimeoutSeconds() {
	o.IdleTimeoutSeconds.Unset()
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateLobbyRequest) GetMembers() []LobbyMemberRequest {
	if o == nil {
		var ret []LobbyMemberRequest
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateLobbyRequest) GetMembersOk() ([]LobbyMemberRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *CreateLobbyRequest) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []LobbyMemberRequest and assigns it to the Members field.
func (o *CreateLobbyRequest) SetMembers(v []LobbyMemberRequest) {
	o.Members = v
}


// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateLobbyRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLobbyRequest) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateLobbyRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateLobbyRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}


func (o CreateLobbyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLobbyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IdleTimeoutSeconds.IsSet() {
		toSerialize["idle_timeout_seconds"] = o.IdleTimeoutSeconds.Get()
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableCreateLobbyRequest struct {
	value *CreateLobbyRequest
	isSet bool
}

func (v NullableCreateLobbyRequest) Get() *CreateLobbyRequest {
	return v.value
}

func (v *NullableCreateLobbyRequest) Set(val *CreateLobbyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLobbyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLobbyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLobbyRequest(val *CreateLobbyRequest) *NullableCreateLobbyRequest {
	return &NullableCreateLobbyRequest{value: val, isSet: true}
}

func (v NullableCreateLobbyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLobbyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


