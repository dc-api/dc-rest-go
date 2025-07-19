/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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

// checks if the ApplicationCommandInteractionMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandInteractionMetadataResponse{}

// ApplicationCommandInteractionMetadataResponse struct for ApplicationCommandInteractionMetadataResponse
type ApplicationCommandInteractionMetadataResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type int32 `json:"type"`
	User NullableUserResponse `json:"user,omitempty"`
	AuthorizingIntegrationOwners map[string]string `json:"authorizing_integration_owners"`
	OriginalResponseMessageId *string `json:"original_response_message_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	TargetUser NullableUserResponse `json:"target_user,omitempty"`
	TargetMessageId *string `json:"target_message_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _ApplicationCommandInteractionMetadataResponse ApplicationCommandInteractionMetadataResponse

// NewApplicationCommandInteractionMetadataResponse instantiates a new ApplicationCommandInteractionMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandInteractionMetadataResponse(id string, type_ int32, authorizingIntegrationOwners map[string]string) *ApplicationCommandInteractionMetadataResponse {
	this := ApplicationCommandInteractionMetadataResponse{}
	this.Id = id
	this.Type = type_
	this.AuthorizingIntegrationOwners = authorizingIntegrationOwners
	return &this
}

// NewApplicationCommandInteractionMetadataResponseWithDefaults instantiates a new ApplicationCommandInteractionMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandInteractionMetadataResponseWithDefaults() *ApplicationCommandInteractionMetadataResponse {
	this := ApplicationCommandInteractionMetadataResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationCommandInteractionMetadataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandInteractionMetadataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationCommandInteractionMetadataResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ApplicationCommandInteractionMetadataResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandInteractionMetadataResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationCommandInteractionMetadataResponse) SetType(v int32) {
	o.Type = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandInteractionMetadataResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandInteractionMetadataResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User.Get()) {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ApplicationCommandInteractionMetadataResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *ApplicationCommandInteractionMetadataResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *ApplicationCommandInteractionMetadataResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ApplicationCommandInteractionMetadataResponse) UnsetUser() {
	o.User.Unset()
}

// GetAuthorizingIntegrationOwners returns the AuthorizingIntegrationOwners field value
func (o *ApplicationCommandInteractionMetadataResponse) GetAuthorizingIntegrationOwners() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.AuthorizingIntegrationOwners
}

// GetAuthorizingIntegrationOwnersOk returns a tuple with the AuthorizingIntegrationOwners field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandInteractionMetadataResponse) GetAuthorizingIntegrationOwnersOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.AuthorizingIntegrationOwners, true
}

// SetAuthorizingIntegrationOwners sets field value
func (o *ApplicationCommandInteractionMetadataResponse) SetAuthorizingIntegrationOwners(v map[string]string) {
	o.AuthorizingIntegrationOwners = v
}

// GetOriginalResponseMessageId returns the OriginalResponseMessageId field value if set, zero value otherwise.
func (o *ApplicationCommandInteractionMetadataResponse) GetOriginalResponseMessageId() string {
	if o == nil || IsNil(o.OriginalResponseMessageId) {
		var ret string
		return ret
	}
	return *o.OriginalResponseMessageId
}

// GetOriginalResponseMessageIdOk returns a tuple with the OriginalResponseMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandInteractionMetadataResponse) GetOriginalResponseMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalResponseMessageId) {
		return nil, false
	}
	return o.OriginalResponseMessageId, true
}

// HasOriginalResponseMessageId returns a boolean if a field has been set.
func (o *ApplicationCommandInteractionMetadataResponse) HasOriginalResponseMessageId() bool {
	if o != nil && !IsNil(o.OriginalResponseMessageId) {
		return true
	}

	return false
}

// SetOriginalResponseMessageId gets a reference to the given string and assigns it to the OriginalResponseMessageId field.
func (o *ApplicationCommandInteractionMetadataResponse) SetOriginalResponseMessageId(v string) {
	o.OriginalResponseMessageId = &v
}


// GetTargetUser returns the TargetUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandInteractionMetadataResponse) GetTargetUser() UserResponse {
	if o == nil || IsNil(o.TargetUser.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.TargetUser.Get()
}

// GetTargetUserOk returns a tuple with the TargetUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandInteractionMetadataResponse) GetTargetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.TargetUser.Get()) {
		return nil, false
	}
	return o.TargetUser.Get(), o.TargetUser.IsSet()
}

// HasTargetUser returns a boolean if a field has been set.
func (o *ApplicationCommandInteractionMetadataResponse) HasTargetUser() bool {
	if o != nil && o.TargetUser.IsSet() {
		return true
	}

	return false
}

// SetTargetUser gets a reference to the given NullableUserResponse and assigns it to the TargetUser field.
func (o *ApplicationCommandInteractionMetadataResponse) SetTargetUser(v UserResponse) {
	o.TargetUser.Set(&v)
}

// SetTargetUserNil sets the value for TargetUser to be an explicit nil
func (o *ApplicationCommandInteractionMetadataResponse) SetTargetUserNil() {
	o.TargetUser.Set(nil)
}

// UnsetTargetUser ensures that no value is present for TargetUser, not even an explicit nil
func (o *ApplicationCommandInteractionMetadataResponse) UnsetTargetUser() {
	o.TargetUser.Unset()
}

// GetTargetMessageId returns the TargetMessageId field value if set, zero value otherwise.
func (o *ApplicationCommandInteractionMetadataResponse) GetTargetMessageId() string {
	if o == nil || IsNil(o.TargetMessageId) {
		var ret string
		return ret
	}
	return *o.TargetMessageId
}

// GetTargetMessageIdOk returns a tuple with the TargetMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandInteractionMetadataResponse) GetTargetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetMessageId) {
		return nil, false
	}
	return o.TargetMessageId, true
}

// HasTargetMessageId returns a boolean if a field has been set.
func (o *ApplicationCommandInteractionMetadataResponse) HasTargetMessageId() bool {
	if o != nil && !IsNil(o.TargetMessageId) {
		return true
	}

	return false
}

// SetTargetMessageId gets a reference to the given string and assigns it to the TargetMessageId field.
func (o *ApplicationCommandInteractionMetadataResponse) SetTargetMessageId(v string) {
	o.TargetMessageId = &v
}


func (o ApplicationCommandInteractionMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandInteractionMetadataResponse) ToMap() (map[string]interface{}, error) {
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
	if o.TargetUser.IsSet() {
		toSerialize["target_user"] = o.TargetUser.Get()
	}
	if !IsNil(o.TargetMessageId) {
		toSerialize["target_message_id"] = o.TargetMessageId
	}
	return toSerialize, nil
}

func (o *ApplicationCommandInteractionMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"authorizing_integration_owners",
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

	varApplicationCommandInteractionMetadataResponse := _ApplicationCommandInteractionMetadataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandInteractionMetadataResponse)

	if err != nil {
		return err
	}

	*o = ApplicationCommandInteractionMetadataResponse(varApplicationCommandInteractionMetadataResponse)

	return err
}

type NullableApplicationCommandInteractionMetadataResponse struct {
	value *ApplicationCommandInteractionMetadataResponse
	isSet bool
}

func (v NullableApplicationCommandInteractionMetadataResponse) Get() *ApplicationCommandInteractionMetadataResponse {
	return v.value
}

func (v *NullableApplicationCommandInteractionMetadataResponse) Set(val *ApplicationCommandInteractionMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandInteractionMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandInteractionMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandInteractionMetadataResponse(val *ApplicationCommandInteractionMetadataResponse) *NullableApplicationCommandInteractionMetadataResponse {
	return &NullableApplicationCommandInteractionMetadataResponse{value: val, isSet: true}
}

func (v NullableApplicationCommandInteractionMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandInteractionMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


