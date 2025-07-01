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

// checks if the ApplicationOAuth2InstallParamsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationOAuth2InstallParamsResponse{}

// ApplicationOAuth2InstallParamsResponse struct for ApplicationOAuth2InstallParamsResponse
type ApplicationOAuth2InstallParamsResponse struct {
	Scopes []string `json:"scopes"`
	Permissions string `json:"permissions"`
}

type _ApplicationOAuth2InstallParamsResponse ApplicationOAuth2InstallParamsResponse

// NewApplicationOAuth2InstallParamsResponse instantiates a new ApplicationOAuth2InstallParamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOAuth2InstallParamsResponse(scopes []string, permissions string) *ApplicationOAuth2InstallParamsResponse {
	this := ApplicationOAuth2InstallParamsResponse{}
	this.Scopes = scopes
	this.Permissions = permissions
	return &this
}

// NewApplicationOAuth2InstallParamsResponseWithDefaults instantiates a new ApplicationOAuth2InstallParamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOAuth2InstallParamsResponseWithDefaults() *ApplicationOAuth2InstallParamsResponse {
	this := ApplicationOAuth2InstallParamsResponse{}
	return &this
}

// GetScopes returns the Scopes field value
func (o *ApplicationOAuth2InstallParamsResponse) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *ApplicationOAuth2InstallParamsResponse) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *ApplicationOAuth2InstallParamsResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetPermissions returns the Permissions field value
func (o *ApplicationOAuth2InstallParamsResponse) GetPermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ApplicationOAuth2InstallParamsResponse) GetPermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *ApplicationOAuth2InstallParamsResponse) SetPermissions(v string) {
	o.Permissions = v
}

func (o ApplicationOAuth2InstallParamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationOAuth2InstallParamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scopes"] = o.Scopes
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

func (o *ApplicationOAuth2InstallParamsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scopes",
		"permissions",
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

	varApplicationOAuth2InstallParamsResponse := _ApplicationOAuth2InstallParamsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationOAuth2InstallParamsResponse)

	if err != nil {
		return err
	}

	*o = ApplicationOAuth2InstallParamsResponse(varApplicationOAuth2InstallParamsResponse)

	return err
}

type NullableApplicationOAuth2InstallParamsResponse struct {
	value *ApplicationOAuth2InstallParamsResponse
	isSet bool
}

func (v NullableApplicationOAuth2InstallParamsResponse) Get() *ApplicationOAuth2InstallParamsResponse {
	return v.value
}

func (v *NullableApplicationOAuth2InstallParamsResponse) Set(val *ApplicationOAuth2InstallParamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOAuth2InstallParamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOAuth2InstallParamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOAuth2InstallParamsResponse(val *ApplicationOAuth2InstallParamsResponse) *NullableApplicationOAuth2InstallParamsResponse {
	return &NullableApplicationOAuth2InstallParamsResponse{value: val, isSet: true}
}

func (v NullableApplicationOAuth2InstallParamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOAuth2InstallParamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


