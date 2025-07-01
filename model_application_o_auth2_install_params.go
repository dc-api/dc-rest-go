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

// checks if the ApplicationOAuth2InstallParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationOAuth2InstallParams{}

// ApplicationOAuth2InstallParams struct for ApplicationOAuth2InstallParams
type ApplicationOAuth2InstallParams struct {
	Scopes []string `json:"scopes,omitempty"`
	Permissions NullableInt32 `json:"permissions,omitempty"`
}

// NewApplicationOAuth2InstallParams instantiates a new ApplicationOAuth2InstallParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOAuth2InstallParams() *ApplicationOAuth2InstallParams {
	this := ApplicationOAuth2InstallParams{}
	return &this
}

// NewApplicationOAuth2InstallParamsWithDefaults instantiates a new ApplicationOAuth2InstallParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOAuth2InstallParamsWithDefaults() *ApplicationOAuth2InstallParams {
	this := ApplicationOAuth2InstallParams{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationOAuth2InstallParams) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationOAuth2InstallParams) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ApplicationOAuth2InstallParams) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ApplicationOAuth2InstallParams) SetScopes(v []string) {
	o.Scopes = v
}


// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationOAuth2InstallParams) GetPermissions() int32 {
	if o == nil || IsNil(o.Permissions.Get()) {
		var ret int32
		return ret
	}
	return *o.Permissions.Get()
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationOAuth2InstallParams) GetPermissionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Permissions.Get()) {
		return nil, false
	}
	return o.Permissions.Get(), o.Permissions.IsSet()
}

// HasPermissions returns a boolean if a field has been set.
func (o *ApplicationOAuth2InstallParams) HasPermissions() bool {
	if o != nil && o.Permissions.IsSet() {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given NullableInt32 and assigns it to the Permissions field.
func (o *ApplicationOAuth2InstallParams) SetPermissions(v int32) {
	o.Permissions.Set(&v)
}

// SetPermissionsNil sets the value for Permissions to be an explicit nil
func (o *ApplicationOAuth2InstallParams) SetPermissionsNil() {
	o.Permissions.Set(nil)
}

// UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
func (o *ApplicationOAuth2InstallParams) UnsetPermissions() {
	o.Permissions.Unset()
}

func (o ApplicationOAuth2InstallParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationOAuth2InstallParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Permissions.IsSet() {
		toSerialize["permissions"] = o.Permissions.Get()
	}
	return toSerialize, nil
}

type NullableApplicationOAuth2InstallParams struct {
	value *ApplicationOAuth2InstallParams
	isSet bool
}

func (v NullableApplicationOAuth2InstallParams) Get() *ApplicationOAuth2InstallParams {
	return v.value
}

func (v *NullableApplicationOAuth2InstallParams) Set(val *ApplicationOAuth2InstallParams) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOAuth2InstallParams) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOAuth2InstallParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOAuth2InstallParams(val *ApplicationOAuth2InstallParams) *NullableApplicationOAuth2InstallParams {
	return &NullableApplicationOAuth2InstallParams{value: val, isSet: true}
}

func (v NullableApplicationOAuth2InstallParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOAuth2InstallParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


