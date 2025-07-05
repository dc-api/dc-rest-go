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

// checks if the MessageComponentInteractionMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageComponentInteractionMetadataResponse{}

// MessageComponentInteractionMetadataResponse struct for MessageComponentInteractionMetadataResponse
type MessageComponentInteractionMetadataResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type int32 `json:"type"`
	User NullableUserResponse `json:"user,omitempty"`
	AuthorizingIntegrationOwners map[string]string `json:"authorizing_integration_owners"`
	OriginalResponseMessageId *string `json:"original_response_message_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	InteractedMessageId string `json:"interacted_message_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _MessageComponentInteractionMetadataResponse MessageComponentInteractionMetadataResponse

// NewMessageComponentInteractionMetadataResponse instantiates a new MessageComponentInteractionMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageComponentInteractionMetadataResponse(id string, type_ int32, authorizingIntegrationOwners map[string]string, interactedMessageId string) *MessageComponentInteractionMetadataResponse {
	this := MessageComponentInteractionMetadataResponse{}
	this.Id = id
	this.Type = type_
	this.AuthorizingIntegrationOwners = authorizingIntegrationOwners
	this.InteractedMessageId = interactedMessageId
	return &this
}

// NewMessageComponentInteractionMetadataResponseWithDefaults instantiates a new MessageComponentInteractionMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageComponentInteractionMetadataResponseWithDefaults() *MessageComponentInteractionMetadataResponse {
	this := MessageComponentInteractionMetadataResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MessageComponentInteractionMetadataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageComponentInteractionMetadataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageComponentInteractionMetadataResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *MessageComponentInteractionMetadataResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageComponentInteractionMetadataResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageComponentInteractionMetadataResponse) SetType(v int32) {
	o.Type = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentInteractionMetadataResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentInteractionMetadataResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User.Get()) {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *MessageComponentInteractionMetadataResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *MessageComponentInteractionMetadataResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *MessageComponentInteractionMetadataResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *MessageComponentInteractionMetadataResponse) UnsetUser() {
	o.User.Unset()
}

// GetAuthorizingIntegrationOwners returns the AuthorizingIntegrationOwners field value
func (o *MessageComponentInteractionMetadataResponse) GetAuthorizingIntegrationOwners() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.AuthorizingIntegrationOwners
}

// GetAuthorizingIntegrationOwnersOk returns a tuple with the AuthorizingIntegrationOwners field value
// and a boolean to check if the value has been set.
func (o *MessageComponentInteractionMetadataResponse) GetAuthorizingIntegrationOwnersOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.AuthorizingIntegrationOwners, true
}

// SetAuthorizingIntegrationOwners sets field value
func (o *MessageComponentInteractionMetadataResponse) SetAuthorizingIntegrationOwners(v map[string]string) {
	o.AuthorizingIntegrationOwners = v
}

// GetOriginalResponseMessageId returns the OriginalResponseMessageId field value if set, zero value otherwise.
func (o *MessageComponentInteractionMetadataResponse) GetOriginalResponseMessageId() string {
	if o == nil || IsNil(o.OriginalResponseMessageId) {
		var ret string
		return ret
	}
	return *o.OriginalResponseMessageId
}

// GetOriginalResponseMessageIdOk returns a tuple with the OriginalResponseMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageComponentInteractionMetadataResponse) GetOriginalResponseMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalResponseMessageId) {
		return nil, false
	}
	return o.OriginalResponseMessageId, true
}

// HasOriginalResponseMessageId returns a boolean if a field has been set.
func (o *MessageComponentInteractionMetadataResponse) HasOriginalResponseMessageId() bool {
	if o != nil && !IsNil(o.OriginalResponseMessageId) {
		return true
	}

	return false
}

// SetOriginalResponseMessageId gets a reference to the given string and assigns it to the OriginalResponseMessageId field.
func (o *MessageComponentInteractionMetadataResponse) SetOriginalResponseMessageId(v string) {
	o.OriginalResponseMessageId = &v
}


// GetInteractedMessageId returns the InteractedMessageId field value
func (o *MessageComponentInteractionMetadataResponse) GetInteractedMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InteractedMessageId
}

// GetInteractedMessageIdOk returns a tuple with the InteractedMessageId field value
// and a boolean to check if the value has been set.
func (o *MessageComponentInteractionMetadataResponse) GetInteractedMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InteractedMessageId, true
}

// SetInteractedMessageId sets field value
func (o *MessageComponentInteractionMetadataResponse) SetInteractedMessageId(v string) {
	o.InteractedMessageId = v
}

func (o MessageComponentInteractionMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageComponentInteractionMetadataResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["interacted_message_id"] = o.InteractedMessageId
	return toSerialize, nil
}

func (o *MessageComponentInteractionMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"authorizing_integration_owners",
		"interacted_message_id",
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

	varMessageComponentInteractionMetadataResponse := _MessageComponentInteractionMetadataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageComponentInteractionMetadataResponse)

	if err != nil {
		return err
	}

	*o = MessageComponentInteractionMetadataResponse(varMessageComponentInteractionMetadataResponse)

	return err
}

type NullableMessageComponentInteractionMetadataResponse struct {
	value *MessageComponentInteractionMetadataResponse
	isSet bool
}

func (v NullableMessageComponentInteractionMetadataResponse) Get() *MessageComponentInteractionMetadataResponse {
	return v.value
}

func (v *NullableMessageComponentInteractionMetadataResponse) Set(val *MessageComponentInteractionMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageComponentInteractionMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageComponentInteractionMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageComponentInteractionMetadataResponse(val *MessageComponentInteractionMetadataResponse) *NullableMessageComponentInteractionMetadataResponse {
	return &NullableMessageComponentInteractionMetadataResponse{value: val, isSet: true}
}

func (v NullableMessageComponentInteractionMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageComponentInteractionMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


