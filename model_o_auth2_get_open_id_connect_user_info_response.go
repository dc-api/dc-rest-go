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

// checks if the OAuth2GetOpenIDConnectUserInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2GetOpenIDConnectUserInfoResponse{}

// OAuth2GetOpenIDConnectUserInfoResponse struct for OAuth2GetOpenIDConnectUserInfoResponse
type OAuth2GetOpenIDConnectUserInfoResponse struct {
	Sub string `json:"sub"`
	Email NullableString `json:"email,omitempty"`
	EmailVerified NullableBool `json:"email_verified,omitempty"`
	PreferredUsername NullableString `json:"preferred_username,omitempty"`
	Nickname NullableString `json:"nickname,omitempty"`
	Picture NullableString `json:"picture,omitempty"`
	Locale NullableString `json:"locale,omitempty"`
}

type _OAuth2GetOpenIDConnectUserInfoResponse OAuth2GetOpenIDConnectUserInfoResponse

// NewOAuth2GetOpenIDConnectUserInfoResponse instantiates a new OAuth2GetOpenIDConnectUserInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2GetOpenIDConnectUserInfoResponse(sub string) *OAuth2GetOpenIDConnectUserInfoResponse {
	this := OAuth2GetOpenIDConnectUserInfoResponse{}
	this.Sub = sub
	return &this
}

// NewOAuth2GetOpenIDConnectUserInfoResponseWithDefaults instantiates a new OAuth2GetOpenIDConnectUserInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2GetOpenIDConnectUserInfoResponseWithDefaults() *OAuth2GetOpenIDConnectUserInfoResponse {
	this := OAuth2GetOpenIDConnectUserInfoResponse{}
	return &this
}

// GetSub returns the Sub field value
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetSub(v string) {
	o.Sub = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email.Get()) {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetEmail() {
	o.Email.Unset()
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetEmailVerified() bool {
	if o == nil || IsNil(o.EmailVerified.Get()) {
		var ret bool
		return ret
	}
	return *o.EmailVerified.Get()
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailVerified.Get()) {
		return nil, false
	}
	return o.EmailVerified.Get(), o.EmailVerified.IsSet()
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasEmailVerified() bool {
	if o != nil && o.EmailVerified.IsSet() {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given NullableBool and assigns it to the EmailVerified field.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetEmailVerified(v bool) {
	o.EmailVerified.Set(&v)
}

// SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetEmailVerifiedNil() {
	o.EmailVerified.Set(nil)
}

// UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetEmailVerified() {
	o.EmailVerified.Unset()
}

// GetPreferredUsername returns the PreferredUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetPreferredUsername() string {
	if o == nil || IsNil(o.PreferredUsername.Get()) {
		var ret string
		return ret
	}
	return *o.PreferredUsername.Get()
}

// GetPreferredUsernameOk returns a tuple with the PreferredUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetPreferredUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredUsername.Get()) {
		return nil, false
	}
	return o.PreferredUsername.Get(), o.PreferredUsername.IsSet()
}

// HasPreferredUsername returns a boolean if a field has been set.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasPreferredUsername() bool {
	if o != nil && o.PreferredUsername.IsSet() {
		return true
	}

	return false
}

// SetPreferredUsername gets a reference to the given NullableString and assigns it to the PreferredUsername field.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetPreferredUsername(v string) {
	o.PreferredUsername.Set(&v)
}

// SetPreferredUsernameNil sets the value for PreferredUsername to be an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetPreferredUsernameNil() {
	o.PreferredUsername.Set(nil)
}

// UnsetPreferredUsername ensures that no value is present for PreferredUsername, not even an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetPreferredUsername() {
	o.PreferredUsername.Unset()
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname.Get()) {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetNickname(v string) {
	o.Nickname.Set(&v)
}

// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetNickname() {
	o.Nickname.Unset()
}

// GetPicture returns the Picture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetPicture() string {
	if o == nil || IsNil(o.Picture.Get()) {
		var ret string
		return ret
	}
	return *o.Picture.Get()
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetPictureOk() (*string, bool) {
	if o == nil || IsNil(o.Picture.Get()) {
		return nil, false
	}
	return o.Picture.Get(), o.Picture.IsSet()
}

// HasPicture returns a boolean if a field has been set.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasPicture() bool {
	if o != nil && o.Picture.IsSet() {
		return true
	}

	return false
}

// SetPicture gets a reference to the given NullableString and assigns it to the Picture field.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetPicture(v string) {
	o.Picture.Set(&v)
}

// SetPictureNil sets the value for Picture to be an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetPictureNil() {
	o.Picture.Set(nil)
}

// UnsetPicture ensures that no value is present for Picture, not even an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetPicture() {
	o.Picture.Unset()
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetLocale() string {
	if o == nil || IsNil(o.Locale.Get()) {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale.Get()) {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// HasLocale returns a boolean if a field has been set.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasLocale() bool {
	if o != nil && o.Locale.IsSet() {
		return true
	}

	return false
}

// SetLocale gets a reference to the given NullableString and assigns it to the Locale field.
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetLocale(v string) {
	o.Locale.Set(&v)
}

// SetLocaleNil sets the value for Locale to be an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetLocaleNil() {
	o.Locale.Set(nil)
}

// UnsetLocale ensures that no value is present for Locale, not even an explicit nil
func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetLocale() {
	o.Locale.Unset()
}

func (o OAuth2GetOpenIDConnectUserInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2GetOpenIDConnectUserInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sub"] = o.Sub
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.EmailVerified.IsSet() {
		toSerialize["email_verified"] = o.EmailVerified.Get()
	}
	if o.PreferredUsername.IsSet() {
		toSerialize["preferred_username"] = o.PreferredUsername.Get()
	}
	if o.Nickname.IsSet() {
		toSerialize["nickname"] = o.Nickname.Get()
	}
	if o.Picture.IsSet() {
		toSerialize["picture"] = o.Picture.Get()
	}
	if o.Locale.IsSet() {
		toSerialize["locale"] = o.Locale.Get()
	}
	return toSerialize, nil
}

func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sub",
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

	varOAuth2GetOpenIDConnectUserInfoResponse := _OAuth2GetOpenIDConnectUserInfoResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2GetOpenIDConnectUserInfoResponse)

	if err != nil {
		return err
	}

	*o = OAuth2GetOpenIDConnectUserInfoResponse(varOAuth2GetOpenIDConnectUserInfoResponse)

	return err
}

type NullableOAuth2GetOpenIDConnectUserInfoResponse struct {
	value *OAuth2GetOpenIDConnectUserInfoResponse
	isSet bool
}

func (v NullableOAuth2GetOpenIDConnectUserInfoResponse) Get() *OAuth2GetOpenIDConnectUserInfoResponse {
	return v.value
}

func (v *NullableOAuth2GetOpenIDConnectUserInfoResponse) Set(val *OAuth2GetOpenIDConnectUserInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2GetOpenIDConnectUserInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2GetOpenIDConnectUserInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2GetOpenIDConnectUserInfoResponse(val *OAuth2GetOpenIDConnectUserInfoResponse) *NullableOAuth2GetOpenIDConnectUserInfoResponse {
	return &NullableOAuth2GetOpenIDConnectUserInfoResponse{value: val, isSet: true}
}

func (v NullableOAuth2GetOpenIDConnectUserInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2GetOpenIDConnectUserInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


