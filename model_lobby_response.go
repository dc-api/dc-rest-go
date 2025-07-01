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

// checks if the LobbyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LobbyResponse{}

// LobbyResponse struct for LobbyResponse
type LobbyResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ApplicationId string `json:"application_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Members []LobbyMemberResponse `json:"members,omitempty"`
	LinkedChannel NullableGuildChannelResponse `json:"linked_channel,omitempty"`
}

type _LobbyResponse LobbyResponse

// NewLobbyResponse instantiates a new LobbyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLobbyResponse(id string, applicationId string) *LobbyResponse {
	this := LobbyResponse{}
	this.Id = id
	this.ApplicationId = applicationId
	return &this
}

// NewLobbyResponseWithDefaults instantiates a new LobbyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLobbyResponseWithDefaults() *LobbyResponse {
	this := LobbyResponse{}
	return &this
}

// GetId returns the Id field value
func (o *LobbyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LobbyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LobbyResponse) SetId(v string) {
	o.Id = v
}

// GetApplicationId returns the ApplicationId field value
func (o *LobbyResponse) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *LobbyResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *LobbyResponse) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LobbyResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LobbyResponse) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LobbyResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *LobbyResponse) SetMetadata(v map[string]string) {
	o.Metadata = v
}


// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LobbyResponse) GetMembers() []LobbyMemberResponse {
	if o == nil {
		var ret []LobbyMemberResponse
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LobbyResponse) GetMembersOk() ([]LobbyMemberResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *LobbyResponse) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []LobbyMemberResponse and assigns it to the Members field.
func (o *LobbyResponse) SetMembers(v []LobbyMemberResponse) {
	o.Members = v
}


// GetLinkedChannel returns the LinkedChannel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LobbyResponse) GetLinkedChannel() GuildChannelResponse {
	if o == nil || IsNil(o.LinkedChannel.Get()) {
		var ret GuildChannelResponse
		return ret
	}
	return *o.LinkedChannel.Get()
}

// GetLinkedChannelOk returns a tuple with the LinkedChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LobbyResponse) GetLinkedChannelOk() (*GuildChannelResponse, bool) {
	if o == nil || IsNil(o.LinkedChannel.Get()) {
		return nil, false
	}
	return o.LinkedChannel.Get(), o.LinkedChannel.IsSet()
}

// HasLinkedChannel returns a boolean if a field has been set.
func (o *LobbyResponse) HasLinkedChannel() bool {
	if o != nil && o.LinkedChannel.IsSet() {
		return true
	}

	return false
}

// SetLinkedChannel gets a reference to the given NullableGuildChannelResponse and assigns it to the LinkedChannel field.
func (o *LobbyResponse) SetLinkedChannel(v GuildChannelResponse) {
	o.LinkedChannel.Set(&v)
}

// SetLinkedChannelNil sets the value for LinkedChannel to be an explicit nil
func (o *LobbyResponse) SetLinkedChannelNil() {
	o.LinkedChannel.Set(nil)
}

// UnsetLinkedChannel ensures that no value is present for LinkedChannel, not even an explicit nil
func (o *LobbyResponse) UnsetLinkedChannel() {
	o.LinkedChannel.Unset()
}

func (o LobbyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LobbyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["application_id"] = o.ApplicationId
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.LinkedChannel.IsSet() {
		toSerialize["linked_channel"] = o.LinkedChannel.Get()
	}
	return toSerialize, nil
}

func (o *LobbyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"application_id",
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

	varLobbyResponse := _LobbyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLobbyResponse)

	if err != nil {
		return err
	}

	*o = LobbyResponse(varLobbyResponse)

	return err
}

type NullableLobbyResponse struct {
	value *LobbyResponse
	isSet bool
}

func (v NullableLobbyResponse) Get() *LobbyResponse {
	return v.value
}

func (v *NullableLobbyResponse) Set(val *LobbyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLobbyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLobbyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLobbyResponse(val *LobbyResponse) *NullableLobbyResponse {
	return &NullableLobbyResponse{value: val, isSet: true}
}

func (v NullableLobbyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLobbyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


