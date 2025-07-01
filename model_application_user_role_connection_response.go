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
)

// checks if the ApplicationUserRoleConnectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationUserRoleConnectionResponse{}

// ApplicationUserRoleConnectionResponse struct for ApplicationUserRoleConnectionResponse
type ApplicationUserRoleConnectionResponse struct {
	PlatformName NullableString `json:"platform_name,omitempty"`
	PlatformUsername NullableString `json:"platform_username,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// NewApplicationUserRoleConnectionResponse instantiates a new ApplicationUserRoleConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationUserRoleConnectionResponse() *ApplicationUserRoleConnectionResponse {
	this := ApplicationUserRoleConnectionResponse{}
	return &this
}

// NewApplicationUserRoleConnectionResponseWithDefaults instantiates a new ApplicationUserRoleConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationUserRoleConnectionResponseWithDefaults() *ApplicationUserRoleConnectionResponse {
	this := ApplicationUserRoleConnectionResponse{}
	return &this
}

// GetPlatformName returns the PlatformName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationUserRoleConnectionResponse) GetPlatformName() string {
	if o == nil || IsNil(o.PlatformName.Get()) {
		var ret string
		return ret
	}
	return *o.PlatformName.Get()
}

// GetPlatformNameOk returns a tuple with the PlatformName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationUserRoleConnectionResponse) GetPlatformNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformName.Get()) {
		return nil, false
	}
	return o.PlatformName.Get(), o.PlatformName.IsSet()
}

// HasPlatformName returns a boolean if a field has been set.
func (o *ApplicationUserRoleConnectionResponse) HasPlatformName() bool {
	if o != nil && o.PlatformName.IsSet() {
		return true
	}

	return false
}

// SetPlatformName gets a reference to the given NullableString and assigns it to the PlatformName field.
func (o *ApplicationUserRoleConnectionResponse) SetPlatformName(v string) {
	o.PlatformName.Set(&v)
}

// SetPlatformNameNil sets the value for PlatformName to be an explicit nil
func (o *ApplicationUserRoleConnectionResponse) SetPlatformNameNil() {
	o.PlatformName.Set(nil)
}

// UnsetPlatformName ensures that no value is present for PlatformName, not even an explicit nil
func (o *ApplicationUserRoleConnectionResponse) UnsetPlatformName() {
	o.PlatformName.Unset()
}

// GetPlatformUsername returns the PlatformUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationUserRoleConnectionResponse) GetPlatformUsername() string {
	if o == nil || IsNil(o.PlatformUsername.Get()) {
		var ret string
		return ret
	}
	return *o.PlatformUsername.Get()
}

// GetPlatformUsernameOk returns a tuple with the PlatformUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationUserRoleConnectionResponse) GetPlatformUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformUsername.Get()) {
		return nil, false
	}
	return o.PlatformUsername.Get(), o.PlatformUsername.IsSet()
}

// HasPlatformUsername returns a boolean if a field has been set.
func (o *ApplicationUserRoleConnectionResponse) HasPlatformUsername() bool {
	if o != nil && o.PlatformUsername.IsSet() {
		return true
	}

	return false
}

// SetPlatformUsername gets a reference to the given NullableString and assigns it to the PlatformUsername field.
func (o *ApplicationUserRoleConnectionResponse) SetPlatformUsername(v string) {
	o.PlatformUsername.Set(&v)
}

// SetPlatformUsernameNil sets the value for PlatformUsername to be an explicit nil
func (o *ApplicationUserRoleConnectionResponse) SetPlatformUsernameNil() {
	o.PlatformUsername.Set(nil)
}

// UnsetPlatformUsername ensures that no value is present for PlatformUsername, not even an explicit nil
func (o *ApplicationUserRoleConnectionResponse) UnsetPlatformUsername() {
	o.PlatformUsername.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApplicationUserRoleConnectionResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUserRoleConnectionResponse) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApplicationUserRoleConnectionResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ApplicationUserRoleConnectionResponse) SetMetadata(v map[string]string) {
	o.Metadata = v
}


func (o ApplicationUserRoleConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationUserRoleConnectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PlatformName.IsSet() {
		toSerialize["platform_name"] = o.PlatformName.Get()
	}
	if o.PlatformUsername.IsSet() {
		toSerialize["platform_username"] = o.PlatformUsername.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableApplicationUserRoleConnectionResponse struct {
	value *ApplicationUserRoleConnectionResponse
	isSet bool
}

func (v NullableApplicationUserRoleConnectionResponse) Get() *ApplicationUserRoleConnectionResponse {
	return v.value
}

func (v *NullableApplicationUserRoleConnectionResponse) Set(val *ApplicationUserRoleConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationUserRoleConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationUserRoleConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationUserRoleConnectionResponse(val *ApplicationUserRoleConnectionResponse) *NullableApplicationUserRoleConnectionResponse {
	return &NullableApplicationUserRoleConnectionResponse{value: val, isSet: true}
}

func (v NullableApplicationUserRoleConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationUserRoleConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


