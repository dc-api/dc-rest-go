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

// checks if the DiscordIntegrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscordIntegrationResponse{}

// DiscordIntegrationResponse struct for DiscordIntegrationResponse
type DiscordIntegrationResponse struct {
	Type NullableString `json:"type"`
	Name NullableString `json:"name,omitempty"`
	Account NullableAccountResponse `json:"account,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Application IntegrationApplicationResponse `json:"application"`
	Scopes []string `json:"scopes"`
	User NullableUserResponse `json:"user,omitempty"`
}

type _DiscordIntegrationResponse DiscordIntegrationResponse

// NewDiscordIntegrationResponse instantiates a new DiscordIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscordIntegrationResponse(type_ NullableString, id string, application IntegrationApplicationResponse, scopes []string) *DiscordIntegrationResponse {
	this := DiscordIntegrationResponse{}
	this.Type = type_
	this.Id = id
	this.Application = application
	this.Scopes = scopes
	return &this
}

// NewDiscordIntegrationResponseWithDefaults instantiates a new DiscordIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscordIntegrationResponseWithDefaults() *DiscordIntegrationResponse {
	this := DiscordIntegrationResponse{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DiscordIntegrationResponse) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscordIntegrationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *DiscordIntegrationResponse) SetType(v string) {
	o.Type.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscordIntegrationResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscordIntegrationResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DiscordIntegrationResponse) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *DiscordIntegrationResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DiscordIntegrationResponse) UnsetName() {
	o.Name.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscordIntegrationResponse) GetAccount() AccountResponse {
	if o == nil || IsNil(o.Account.Get()) {
		var ret AccountResponse
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscordIntegrationResponse) GetAccountOk() (*AccountResponse, bool) {
	if o == nil || IsNil(o.Account.Get()) {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableAccountResponse and assigns it to the Account field.
func (o *DiscordIntegrationResponse) SetAccount(v AccountResponse) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *DiscordIntegrationResponse) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *DiscordIntegrationResponse) UnsetAccount() {
	o.Account.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscordIntegrationResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscordIntegrationResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled.Get()) {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *DiscordIntegrationResponse) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *DiscordIntegrationResponse) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *DiscordIntegrationResponse) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetId returns the Id field value
func (o *DiscordIntegrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscordIntegrationResponse) SetId(v string) {
	o.Id = v
}

// GetApplication returns the Application field value
func (o *DiscordIntegrationResponse) GetApplication() IntegrationApplicationResponse {
	if o == nil {
		var ret IntegrationApplicationResponse
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetApplicationOk() (*IntegrationApplicationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *DiscordIntegrationResponse) SetApplication(v IntegrationApplicationResponse) {
	o.Application = v
}

// GetScopes returns the Scopes field value
func (o *DiscordIntegrationResponse) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *DiscordIntegrationResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscordIntegrationResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscordIntegrationResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User.Get()) {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *DiscordIntegrationResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *DiscordIntegrationResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *DiscordIntegrationResponse) UnsetUser() {
	o.User.Unset()
}

func (o DiscordIntegrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscordIntegrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["application"] = o.Application
	toSerialize["scopes"] = o.Scopes
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

func (o *DiscordIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"application",
		"scopes",
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

	varDiscordIntegrationResponse := _DiscordIntegrationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscordIntegrationResponse)

	if err != nil {
		return err
	}

	*o = DiscordIntegrationResponse(varDiscordIntegrationResponse)

	return err
}

type NullableDiscordIntegrationResponse struct {
	value *DiscordIntegrationResponse
	isSet bool
}

func (v NullableDiscordIntegrationResponse) Get() *DiscordIntegrationResponse {
	return v.value
}

func (v *NullableDiscordIntegrationResponse) Set(val *DiscordIntegrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscordIntegrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscordIntegrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscordIntegrationResponse(val *DiscordIntegrationResponse) *NullableDiscordIntegrationResponse {
	return &NullableDiscordIntegrationResponse{value: val, isSet: true}
}

func (v NullableDiscordIntegrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscordIntegrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


