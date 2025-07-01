/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// checks if the ModalSubmitInteractionMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModalSubmitInteractionMetadataResponse{}

// ModalSubmitInteractionMetadataResponse struct for ModalSubmitInteractionMetadataResponse
type ModalSubmitInteractionMetadataResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type int32 `json:"type"`
	User NullableUserResponse `json:"user,omitempty"`
	AuthorizingIntegrationOwners map[string]string `json:"authorizing_integration_owners"`
	OriginalResponseMessageId *string `json:"original_response_message_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	TriggeringInteractionMetadata ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata `json:"triggering_interaction_metadata"`
}

type _ModalSubmitInteractionMetadataResponse ModalSubmitInteractionMetadataResponse

// NewModalSubmitInteractionMetadataResponse instantiates a new ModalSubmitInteractionMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModalSubmitInteractionMetadataResponse(id string, type_ int32, authorizingIntegrationOwners map[string]string, triggeringInteractionMetadata ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) *ModalSubmitInteractionMetadataResponse {
	this := ModalSubmitInteractionMetadataResponse{}
	this.Id = id
	this.Type = type_
	this.AuthorizingIntegrationOwners = authorizingIntegrationOwners
	this.TriggeringInteractionMetadata = triggeringInteractionMetadata
	return &this
}

// NewModalSubmitInteractionMetadataResponseWithDefaults instantiates a new ModalSubmitInteractionMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModalSubmitInteractionMetadataResponseWithDefaults() *ModalSubmitInteractionMetadataResponse {
	this := ModalSubmitInteractionMetadataResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ModalSubmitInteractionMetadataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModalSubmitInteractionMetadataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModalSubmitInteractionMetadataResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ModalSubmitInteractionMetadataResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ModalSubmitInteractionMetadataResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ModalSubmitInteractionMetadataResponse) SetType(v int32) {
	o.Type = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModalSubmitInteractionMetadataResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModalSubmitInteractionMetadataResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User.Get()) {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ModalSubmitInteractionMetadataResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *ModalSubmitInteractionMetadataResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *ModalSubmitInteractionMetadataResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ModalSubmitInteractionMetadataResponse) UnsetUser() {
	o.User.Unset()
}

// GetAuthorizingIntegrationOwners returns the AuthorizingIntegrationOwners field value
func (o *ModalSubmitInteractionMetadataResponse) GetAuthorizingIntegrationOwners() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.AuthorizingIntegrationOwners
}

// GetAuthorizingIntegrationOwnersOk returns a tuple with the AuthorizingIntegrationOwners field value
// and a boolean to check if the value has been set.
func (o *ModalSubmitInteractionMetadataResponse) GetAuthorizingIntegrationOwnersOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.AuthorizingIntegrationOwners, true
}

// SetAuthorizingIntegrationOwners sets field value
func (o *ModalSubmitInteractionMetadataResponse) SetAuthorizingIntegrationOwners(v map[string]string) {
	o.AuthorizingIntegrationOwners = v
}

// GetOriginalResponseMessageId returns the OriginalResponseMessageId field value if set, zero value otherwise.
func (o *ModalSubmitInteractionMetadataResponse) GetOriginalResponseMessageId() string {
	if o == nil || IsNil(o.OriginalResponseMessageId) {
		var ret string
		return ret
	}
	return *o.OriginalResponseMessageId
}

// GetOriginalResponseMessageIdOk returns a tuple with the OriginalResponseMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModalSubmitInteractionMetadataResponse) GetOriginalResponseMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalResponseMessageId) {
		return nil, false
	}
	return o.OriginalResponseMessageId, true
}

// HasOriginalResponseMessageId returns a boolean if a field has been set.
func (o *ModalSubmitInteractionMetadataResponse) HasOriginalResponseMessageId() bool {
	if o != nil && !IsNil(o.OriginalResponseMessageId) {
		return true
	}

	return false
}

// SetOriginalResponseMessageId gets a reference to the given string and assigns it to the OriginalResponseMessageId field.
func (o *ModalSubmitInteractionMetadataResponse) SetOriginalResponseMessageId(v string) {
	o.OriginalResponseMessageId = &v
}


// GetTriggeringInteractionMetadata returns the TriggeringInteractionMetadata field value
func (o *ModalSubmitInteractionMetadataResponse) GetTriggeringInteractionMetadata() ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata {
	if o == nil {
		var ret ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata
		return ret
	}

	return o.TriggeringInteractionMetadata
}

// GetTriggeringInteractionMetadataOk returns a tuple with the TriggeringInteractionMetadata field value
// and a boolean to check if the value has been set.
func (o *ModalSubmitInteractionMetadataResponse) GetTriggeringInteractionMetadataOk() (*ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggeringInteractionMetadata, true
}

// SetTriggeringInteractionMetadata sets field value
func (o *ModalSubmitInteractionMetadataResponse) SetTriggeringInteractionMetadata(v ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) {
	o.TriggeringInteractionMetadata = v
}

func (o ModalSubmitInteractionMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModalSubmitInteractionMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	toSerialize["authorizing_integration_owners"] = o.AuthorizingIntegrationOwners
	if !IsNil(o.OriginalResponseMessageId) {
		toSerialize["original_response_message_id"] = o.OriginalResponseMessageId
	}
	toSerialize["triggering_interaction_metadata"] = o.TriggeringInteractionMetadata
	return toSerialize, nil
}

func (o *ModalSubmitInteractionMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"authorizing_integration_owners",
		"triggering_interaction_metadata",
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

	varModalSubmitInteractionMetadataResponse := _ModalSubmitInteractionMetadataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModalSubmitInteractionMetadataResponse)

	if err != nil {
		return err
	}

	*o = ModalSubmitInteractionMetadataResponse(varModalSubmitInteractionMetadataResponse)

	return err
}

type NullableModalSubmitInteractionMetadataResponse struct {
	value *ModalSubmitInteractionMetadataResponse
	isSet bool
}

func (v NullableModalSubmitInteractionMetadataResponse) Get() *ModalSubmitInteractionMetadataResponse {
	return v.value
}

func (v *NullableModalSubmitInteractionMetadataResponse) Set(val *ModalSubmitInteractionMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModalSubmitInteractionMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModalSubmitInteractionMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModalSubmitInteractionMetadataResponse(val *ModalSubmitInteractionMetadataResponse) *NullableModalSubmitInteractionMetadataResponse {
	return &NullableModalSubmitInteractionMetadataResponse{value: val, isSet: true}
}

func (v NullableModalSubmitInteractionMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModalSubmitInteractionMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


