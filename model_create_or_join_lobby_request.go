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

// checks if the CreateOrJoinLobbyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrJoinLobbyRequest{}

// CreateOrJoinLobbyRequest struct for CreateOrJoinLobbyRequest
type CreateOrJoinLobbyRequest struct {
	IdleTimeoutSeconds NullableInt32 `json:"idle_timeout_seconds,omitempty"`
	LobbyMetadata map[string]string `json:"lobby_metadata,omitempty"`
	MemberMetadata map[string]string `json:"member_metadata,omitempty"`
	Secret string `json:"secret"`
}

type _CreateOrJoinLobbyRequest CreateOrJoinLobbyRequest

// NewCreateOrJoinLobbyRequest instantiates a new CreateOrJoinLobbyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrJoinLobbyRequest(secret string) *CreateOrJoinLobbyRequest {
	this := CreateOrJoinLobbyRequest{}
	this.Secret = secret
	return &this
}

// NewCreateOrJoinLobbyRequestWithDefaults instantiates a new CreateOrJoinLobbyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrJoinLobbyRequestWithDefaults() *CreateOrJoinLobbyRequest {
	this := CreateOrJoinLobbyRequest{}
	return &this
}

// GetIdleTimeoutSeconds returns the IdleTimeoutSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrJoinLobbyRequest) GetIdleTimeoutSeconds() int32 {
	if o == nil || IsNil(o.IdleTimeoutSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.IdleTimeoutSeconds.Get()
}

// GetIdleTimeoutSecondsOk returns a tuple with the IdleTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrJoinLobbyRequest) GetIdleTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.IdleTimeoutSeconds.Get()) {
		return nil, false
	}
	return o.IdleTimeoutSeconds.Get(), o.IdleTimeoutSeconds.IsSet()
}

// HasIdleTimeoutSeconds returns a boolean if a field has been set.
func (o *CreateOrJoinLobbyRequest) HasIdleTimeoutSeconds() bool {
	if o != nil && o.IdleTimeoutSeconds.IsSet() {
		return true
	}

	return false
}

// SetIdleTimeoutSeconds gets a reference to the given NullableInt32 and assigns it to the IdleTimeoutSeconds field.
func (o *CreateOrJoinLobbyRequest) SetIdleTimeoutSeconds(v int32) {
	o.IdleTimeoutSeconds.Set(&v)
}

// SetIdleTimeoutSecondsNil sets the value for IdleTimeoutSeconds to be an explicit nil
func (o *CreateOrJoinLobbyRequest) SetIdleTimeoutSecondsNil() {
	o.IdleTimeoutSeconds.Set(nil)
}

// UnsetIdleTimeoutSeconds ensures that no value is present for IdleTimeoutSeconds, not even an explicit nil
func (o *CreateOrJoinLobbyRequest) UnsetIdleTimeoutSeconds() {
	o.IdleTimeoutSeconds.Unset()
}

// GetLobbyMetadata returns the LobbyMetadata field value if set, zero value otherwise.
func (o *CreateOrJoinLobbyRequest) GetLobbyMetadata() map[string]string {
	if o == nil || IsNil(o.LobbyMetadata) {
		var ret map[string]string
		return ret
	}
	return o.LobbyMetadata
}

// GetLobbyMetadataOk returns a tuple with the LobbyMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrJoinLobbyRequest) GetLobbyMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.LobbyMetadata) {
		return map[string]string{}, false
	}
	return o.LobbyMetadata, true
}

// HasLobbyMetadata returns a boolean if a field has been set.
func (o *CreateOrJoinLobbyRequest) HasLobbyMetadata() bool {
	if o != nil && !IsNil(o.LobbyMetadata) {
		return true
	}

	return false
}

// SetLobbyMetadata gets a reference to the given map[string]string and assigns it to the LobbyMetadata field.
func (o *CreateOrJoinLobbyRequest) SetLobbyMetadata(v map[string]string) {
	o.LobbyMetadata = v
}


// GetMemberMetadata returns the MemberMetadata field value if set, zero value otherwise.
func (o *CreateOrJoinLobbyRequest) GetMemberMetadata() map[string]string {
	if o == nil || IsNil(o.MemberMetadata) {
		var ret map[string]string
		return ret
	}
	return o.MemberMetadata
}

// GetMemberMetadataOk returns a tuple with the MemberMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrJoinLobbyRequest) GetMemberMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.MemberMetadata) {
		return map[string]string{}, false
	}
	return o.MemberMetadata, true
}

// HasMemberMetadata returns a boolean if a field has been set.
func (o *CreateOrJoinLobbyRequest) HasMemberMetadata() bool {
	if o != nil && !IsNil(o.MemberMetadata) {
		return true
	}

	return false
}

// SetMemberMetadata gets a reference to the given map[string]string and assigns it to the MemberMetadata field.
func (o *CreateOrJoinLobbyRequest) SetMemberMetadata(v map[string]string) {
	o.MemberMetadata = v
}


// GetSecret returns the Secret field value
func (o *CreateOrJoinLobbyRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateOrJoinLobbyRequest) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateOrJoinLobbyRequest) SetSecret(v string) {
	o.Secret = v
}

func (o CreateOrJoinLobbyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrJoinLobbyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IdleTimeoutSeconds.IsSet() {
		toSerialize["idle_timeout_seconds"] = o.IdleTimeoutSeconds.Get()
	}
	if !IsNil(o.LobbyMetadata) {
		toSerialize["lobby_metadata"] = o.LobbyMetadata
	}
	if !IsNil(o.MemberMetadata) {
		toSerialize["member_metadata"] = o.MemberMetadata
	}
	toSerialize["secret"] = o.Secret
	return toSerialize, nil
}

func (o *CreateOrJoinLobbyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"secret",
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

	varCreateOrJoinLobbyRequest := _CreateOrJoinLobbyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrJoinLobbyRequest)

	if err != nil {
		return err
	}

	*o = CreateOrJoinLobbyRequest(varCreateOrJoinLobbyRequest)

	return err
}

type NullableCreateOrJoinLobbyRequest struct {
	value *CreateOrJoinLobbyRequest
	isSet bool
}

func (v NullableCreateOrJoinLobbyRequest) Get() *CreateOrJoinLobbyRequest {
	return v.value
}

func (v *NullableCreateOrJoinLobbyRequest) Set(val *CreateOrJoinLobbyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrJoinLobbyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrJoinLobbyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrJoinLobbyRequest(val *CreateOrJoinLobbyRequest) *NullableCreateOrJoinLobbyRequest {
	return &NullableCreateOrJoinLobbyRequest{value: val, isSet: true}
}

func (v NullableCreateOrJoinLobbyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrJoinLobbyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


