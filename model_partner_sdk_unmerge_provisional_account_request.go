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

// checks if the PartnerSdkUnmergeProvisionalAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerSdkUnmergeProvisionalAccountRequest{}

// PartnerSdkUnmergeProvisionalAccountRequest struct for PartnerSdkUnmergeProvisionalAccountRequest
type PartnerSdkUnmergeProvisionalAccountRequest struct {
	ClientId string `json:"client_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ClientSecret NullableString `json:"client_secret,omitempty"`
	ExternalAuthToken string `json:"external_auth_token"`
	ExternalAuthType NullableString `json:"external_auth_type"`
}

type _PartnerSdkUnmergeProvisionalAccountRequest PartnerSdkUnmergeProvisionalAccountRequest

// NewPartnerSdkUnmergeProvisionalAccountRequest instantiates a new PartnerSdkUnmergeProvisionalAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerSdkUnmergeProvisionalAccountRequest(clientId string, externalAuthToken string, externalAuthType NullableString) *PartnerSdkUnmergeProvisionalAccountRequest {
	this := PartnerSdkUnmergeProvisionalAccountRequest{}
	this.ClientId = clientId
	this.ExternalAuthToken = externalAuthToken
	this.ExternalAuthType = externalAuthType
	return &this
}

// NewPartnerSdkUnmergeProvisionalAccountRequestWithDefaults instantiates a new PartnerSdkUnmergeProvisionalAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerSdkUnmergeProvisionalAccountRequestWithDefaults() *PartnerSdkUnmergeProvisionalAccountRequest {
	this := PartnerSdkUnmergeProvisionalAccountRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *PartnerSdkUnmergeProvisionalAccountRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PartnerSdkUnmergeProvisionalAccountRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *PartnerSdkUnmergeProvisionalAccountRequest) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}

// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *PartnerSdkUnmergeProvisionalAccountRequest) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *PartnerSdkUnmergeProvisionalAccountRequest) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetExternalAuthToken returns the ExternalAuthToken field value
func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetExternalAuthToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalAuthToken
}

// GetExternalAuthTokenOk returns a tuple with the ExternalAuthToken field value
// and a boolean to check if the value has been set.
func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetExternalAuthTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalAuthToken, true
}

// SetExternalAuthToken sets field value
func (o *PartnerSdkUnmergeProvisionalAccountRequest) SetExternalAuthToken(v string) {
	o.ExternalAuthToken = v
}

// GetExternalAuthType returns the ExternalAuthType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetExternalAuthType() string {
	if o == nil || o.ExternalAuthType.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalAuthType.Get()
}

// GetExternalAuthTypeOk returns a tuple with the ExternalAuthType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetExternalAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalAuthType.Get(), o.ExternalAuthType.IsSet()
}

// SetExternalAuthType sets field value
func (o *PartnerSdkUnmergeProvisionalAccountRequest) SetExternalAuthType(v string) {
	o.ExternalAuthType.Set(&v)
}

func (o PartnerSdkUnmergeProvisionalAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerSdkUnmergeProvisionalAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client_id"] = o.ClientId
	if o.ClientSecret.IsSet() {
		toSerialize["client_secret"] = o.ClientSecret.Get()
	}
	toSerialize["external_auth_token"] = o.ExternalAuthToken
	toSerialize["external_auth_type"] = o.ExternalAuthType.Get()
	return toSerialize, nil
}

func (o *PartnerSdkUnmergeProvisionalAccountRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_id",
		"external_auth_token",
		"external_auth_type",
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

	varPartnerSdkUnmergeProvisionalAccountRequest := _PartnerSdkUnmergeProvisionalAccountRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPartnerSdkUnmergeProvisionalAccountRequest)

	if err != nil {
		return err
	}

	*o = PartnerSdkUnmergeProvisionalAccountRequest(varPartnerSdkUnmergeProvisionalAccountRequest)

	return err
}

type NullablePartnerSdkUnmergeProvisionalAccountRequest struct {
	value *PartnerSdkUnmergeProvisionalAccountRequest
	isSet bool
}

func (v NullablePartnerSdkUnmergeProvisionalAccountRequest) Get() *PartnerSdkUnmergeProvisionalAccountRequest {
	return v.value
}

func (v *NullablePartnerSdkUnmergeProvisionalAccountRequest) Set(val *PartnerSdkUnmergeProvisionalAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerSdkUnmergeProvisionalAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerSdkUnmergeProvisionalAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerSdkUnmergeProvisionalAccountRequest(val *PartnerSdkUnmergeProvisionalAccountRequest) *NullablePartnerSdkUnmergeProvisionalAccountRequest {
	return &NullablePartnerSdkUnmergeProvisionalAccountRequest{value: val, isSet: true}
}

func (v NullablePartnerSdkUnmergeProvisionalAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerSdkUnmergeProvisionalAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


