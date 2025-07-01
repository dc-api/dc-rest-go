/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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

// checks if the ProvisionalTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionalTokenResponse{}

// ProvisionalTokenResponse struct for ProvisionalTokenResponse
type ProvisionalTokenResponse struct {
	TokenType string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn int32 `json:"expires_in"`
	Scope string `json:"scope"`
	IdToken string `json:"id_token"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	ExpiresAtS NullableInt32 `json:"expires_at_s,omitempty"`
}

type _ProvisionalTokenResponse ProvisionalTokenResponse

// NewProvisionalTokenResponse instantiates a new ProvisionalTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionalTokenResponse(tokenType string, accessToken string, expiresIn int32, scope string, idToken string) *ProvisionalTokenResponse {
	this := ProvisionalTokenResponse{}
	this.TokenType = tokenType
	this.AccessToken = accessToken
	this.ExpiresIn = expiresIn
	this.Scope = scope
	this.IdToken = idToken
	return &this
}

// NewProvisionalTokenResponseWithDefaults instantiates a new ProvisionalTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionalTokenResponseWithDefaults() *ProvisionalTokenResponse {
	this := ProvisionalTokenResponse{}
	return &this
}

// GetTokenType returns the TokenType field value
func (o *ProvisionalTokenResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *ProvisionalTokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *ProvisionalTokenResponse) SetTokenType(v string) {
	o.TokenType = v
}

// GetAccessToken returns the AccessToken field value
func (o *ProvisionalTokenResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ProvisionalTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ProvisionalTokenResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *ProvisionalTokenResponse) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *ProvisionalTokenResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *ProvisionalTokenResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetScope returns the Scope field value
func (o *ProvisionalTokenResponse) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ProvisionalTokenResponse) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *ProvisionalTokenResponse) SetScope(v string) {
	o.Scope = v
}

// GetIdToken returns the IdToken field value
func (o *ProvisionalTokenResponse) GetIdToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value
// and a boolean to check if the value has been set.
func (o *ProvisionalTokenResponse) GetIdTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdToken, true
}

// SetIdToken sets field value
func (o *ProvisionalTokenResponse) SetIdToken(v string) {
	o.IdToken = v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionalTokenResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionalTokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *ProvisionalTokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *ProvisionalTokenResponse) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}

// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *ProvisionalTokenResponse) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *ProvisionalTokenResponse) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionalTokenResponse) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionalTokenResponse) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ProvisionalTokenResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ProvisionalTokenResponse) SetScopes(v []string) {
	o.Scopes = v
}


// GetExpiresAtS returns the ExpiresAtS field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionalTokenResponse) GetExpiresAtS() int32 {
	if o == nil || IsNil(o.ExpiresAtS.Get()) {
		var ret int32
		return ret
	}
	return *o.ExpiresAtS.Get()
}

// GetExpiresAtSOk returns a tuple with the ExpiresAtS field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionalTokenResponse) GetExpiresAtSOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresAtS.Get()) {
		return nil, false
	}
	return o.ExpiresAtS.Get(), o.ExpiresAtS.IsSet()
}

// HasExpiresAtS returns a boolean if a field has been set.
func (o *ProvisionalTokenResponse) HasExpiresAtS() bool {
	if o != nil && o.ExpiresAtS.IsSet() {
		return true
	}

	return false
}

// SetExpiresAtS gets a reference to the given NullableInt32 and assigns it to the ExpiresAtS field.
func (o *ProvisionalTokenResponse) SetExpiresAtS(v int32) {
	o.ExpiresAtS.Set(&v)
}

// SetExpiresAtSNil sets the value for ExpiresAtS to be an explicit nil
func (o *ProvisionalTokenResponse) SetExpiresAtSNil() {
	o.ExpiresAtS.Set(nil)
}

// UnsetExpiresAtS ensures that no value is present for ExpiresAtS, not even an explicit nil
func (o *ProvisionalTokenResponse) UnsetExpiresAtS() {
	o.ExpiresAtS.Unset()
}

func (o ProvisionalTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionalTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_type"] = o.TokenType
	toSerialize["access_token"] = o.AccessToken
	toSerialize["expires_in"] = o.ExpiresIn
	toSerialize["scope"] = o.Scope
	toSerialize["id_token"] = o.IdToken
	if o.RefreshToken.IsSet() {
		toSerialize["refresh_token"] = o.RefreshToken.Get()
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.ExpiresAtS.IsSet() {
		toSerialize["expires_at_s"] = o.ExpiresAtS.Get()
	}
	return toSerialize, nil
}

func (o *ProvisionalTokenResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token_type",
		"access_token",
		"expires_in",
		"scope",
		"id_token",
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

	varProvisionalTokenResponse := _ProvisionalTokenResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProvisionalTokenResponse)

	if err != nil {
		return err
	}

	*o = ProvisionalTokenResponse(varProvisionalTokenResponse)

	return err
}

type NullableProvisionalTokenResponse struct {
	value *ProvisionalTokenResponse
	isSet bool
}

func (v NullableProvisionalTokenResponse) Get() *ProvisionalTokenResponse {
	return v.value
}

func (v *NullableProvisionalTokenResponse) Set(val *ProvisionalTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionalTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionalTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionalTokenResponse(val *ProvisionalTokenResponse) *NullableProvisionalTokenResponse {
	return &NullableProvisionalTokenResponse{value: val, isSet: true}
}

func (v NullableProvisionalTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionalTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


