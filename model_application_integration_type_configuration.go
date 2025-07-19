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
)

// checks if the ApplicationIntegrationTypeConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationIntegrationTypeConfiguration{}

// ApplicationIntegrationTypeConfiguration struct for ApplicationIntegrationTypeConfiguration
type ApplicationIntegrationTypeConfiguration struct {
	Oauth2InstallParams NullableApplicationOAuth2InstallParams `json:"oauth2_install_params,omitempty"`
}

// NewApplicationIntegrationTypeConfiguration instantiates a new ApplicationIntegrationTypeConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationIntegrationTypeConfiguration() *ApplicationIntegrationTypeConfiguration {
	this := ApplicationIntegrationTypeConfiguration{}
	return &this
}

// NewApplicationIntegrationTypeConfigurationWithDefaults instantiates a new ApplicationIntegrationTypeConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationIntegrationTypeConfigurationWithDefaults() *ApplicationIntegrationTypeConfiguration {
	this := ApplicationIntegrationTypeConfiguration{}
	return &this
}

// GetOauth2InstallParams returns the Oauth2InstallParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationIntegrationTypeConfiguration) GetOauth2InstallParams() ApplicationOAuth2InstallParams {
	if o == nil || IsNil(o.Oauth2InstallParams.Get()) {
		var ret ApplicationOAuth2InstallParams
		return ret
	}
	return *o.Oauth2InstallParams.Get()
}

// GetOauth2InstallParamsOk returns a tuple with the Oauth2InstallParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationIntegrationTypeConfiguration) GetOauth2InstallParamsOk() (*ApplicationOAuth2InstallParams, bool) {
	if o == nil || IsNil(o.Oauth2InstallParams.Get()) {
		return nil, false
	}
	return o.Oauth2InstallParams.Get(), o.Oauth2InstallParams.IsSet()
}

// HasOauth2InstallParams returns a boolean if a field has been set.
func (o *ApplicationIntegrationTypeConfiguration) HasOauth2InstallParams() bool {
	if o != nil && o.Oauth2InstallParams.IsSet() {
		return true
	}

	return false
}

// SetOauth2InstallParams gets a reference to the given NullableApplicationOAuth2InstallParams and assigns it to the Oauth2InstallParams field.
func (o *ApplicationIntegrationTypeConfiguration) SetOauth2InstallParams(v ApplicationOAuth2InstallParams) {
	o.Oauth2InstallParams.Set(&v)
}

// SetOauth2InstallParamsNil sets the value for Oauth2InstallParams to be an explicit nil
func (o *ApplicationIntegrationTypeConfiguration) SetOauth2InstallParamsNil() {
	o.Oauth2InstallParams.Set(nil)
}

// UnsetOauth2InstallParams ensures that no value is present for Oauth2InstallParams, not even an explicit nil
func (o *ApplicationIntegrationTypeConfiguration) UnsetOauth2InstallParams() {
	o.Oauth2InstallParams.Unset()
}

func (o ApplicationIntegrationTypeConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationIntegrationTypeConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Oauth2InstallParams.IsSet() {
		toSerialize["oauth2_install_params"] = o.Oauth2InstallParams.Get()
	}
	return toSerialize, nil
}

type NullableApplicationIntegrationTypeConfiguration struct {
	value *ApplicationIntegrationTypeConfiguration
	isSet bool
}

func (v NullableApplicationIntegrationTypeConfiguration) Get() *ApplicationIntegrationTypeConfiguration {
	return v.value
}

func (v *NullableApplicationIntegrationTypeConfiguration) Set(val *ApplicationIntegrationTypeConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationIntegrationTypeConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationIntegrationTypeConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationIntegrationTypeConfiguration(val *ApplicationIntegrationTypeConfiguration) *NullableApplicationIntegrationTypeConfiguration {
	return &NullableApplicationIntegrationTypeConfiguration{value: val, isSet: true}
}

func (v NullableApplicationIntegrationTypeConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationIntegrationTypeConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


