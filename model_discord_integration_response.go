/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
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
	"bytes"
	"fmt"
)

// checks if the DiscordIntegrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscordIntegrationResponse{}

// DiscordIntegrationResponse struct for DiscordIntegrationResponse
type DiscordIntegrationResponse struct {
	Type IntegrationTypes `json:"type"`
	Name *string `json:"name,omitempty"`
	Account AccountResponse `json:"account"`
	Enabled bool `json:"enabled"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Application IntegrationApplicationResponse `json:"application"`
	Scopes []OAuth2Scopes `json:"scopes"`
	User *UserResponse `json:"user,omitempty"`
}

type _DiscordIntegrationResponse DiscordIntegrationResponse

// NewDiscordIntegrationResponse instantiates a new DiscordIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscordIntegrationResponse(type_ IntegrationTypes, account AccountResponse, enabled bool, id string, application IntegrationApplicationResponse, scopes []OAuth2Scopes) *DiscordIntegrationResponse {
	this := DiscordIntegrationResponse{}
	this.Type = type_
	this.Account = account
	this.Enabled = enabled
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
func (o *DiscordIntegrationResponse) GetType() IntegrationTypes {
	if o == nil {
		var ret IntegrationTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetTypeOk() (*IntegrationTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DiscordIntegrationResponse) SetType(v IntegrationTypes) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DiscordIntegrationResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DiscordIntegrationResponse) SetName(v string) {
	o.Name = &v
}


// GetAccount returns the Account field value
func (o *DiscordIntegrationResponse) GetAccount() AccountResponse {
	if o == nil {
		var ret AccountResponse
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetAccountOk() (*AccountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *DiscordIntegrationResponse) SetAccount(v AccountResponse) {
	o.Account = v
}

// GetEnabled returns the Enabled field value
func (o *DiscordIntegrationResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DiscordIntegrationResponse) SetEnabled(v bool) {
	o.Enabled = v
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
func (o *DiscordIntegrationResponse) GetScopes() []OAuth2Scopes {
	if o == nil {
		var ret []OAuth2Scopes
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetScopesOk() ([]OAuth2Scopes, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *DiscordIntegrationResponse) SetScopes(v []OAuth2Scopes) {
	o.Scopes = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DiscordIntegrationResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User) {
		var ret UserResponse
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserResponse and assigns it to the User field.
func (o *DiscordIntegrationResponse) SetUser(v UserResponse) {
	o.User = &v
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
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["account"] = o.Account
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["application"] = o.Application
	toSerialize["scopes"] = o.Scopes
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

func (o *DiscordIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"account",
		"enabled",
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

