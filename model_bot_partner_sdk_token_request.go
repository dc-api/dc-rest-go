/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// checks if the BotPartnerSdkTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BotPartnerSdkTokenRequest{}

// BotPartnerSdkTokenRequest struct for BotPartnerSdkTokenRequest
type BotPartnerSdkTokenRequest struct {
	ExternalUserId string `json:"external_user_id"`
	PreferredGlobalName NullableString `json:"preferred_global_name,omitempty"`
}

type _BotPartnerSdkTokenRequest BotPartnerSdkTokenRequest

// NewBotPartnerSdkTokenRequest instantiates a new BotPartnerSdkTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotPartnerSdkTokenRequest(externalUserId string) *BotPartnerSdkTokenRequest {
	this := BotPartnerSdkTokenRequest{}
	this.ExternalUserId = externalUserId
	return &this
}

// NewBotPartnerSdkTokenRequestWithDefaults instantiates a new BotPartnerSdkTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotPartnerSdkTokenRequestWithDefaults() *BotPartnerSdkTokenRequest {
	this := BotPartnerSdkTokenRequest{}
	return &this
}

// GetExternalUserId returns the ExternalUserId field value
func (o *BotPartnerSdkTokenRequest) GetExternalUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value
// and a boolean to check if the value has been set.
func (o *BotPartnerSdkTokenRequest) GetExternalUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUserId, true
}

// SetExternalUserId sets field value
func (o *BotPartnerSdkTokenRequest) SetExternalUserId(v string) {
	o.ExternalUserId = v
}

// GetPreferredGlobalName returns the PreferredGlobalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BotPartnerSdkTokenRequest) GetPreferredGlobalName() string {
	if o == nil || IsNil(o.PreferredGlobalName.Get()) {
		var ret string
		return ret
	}
	return *o.PreferredGlobalName.Get()
}

// GetPreferredGlobalNameOk returns a tuple with the PreferredGlobalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BotPartnerSdkTokenRequest) GetPreferredGlobalNameOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredGlobalName.Get()) {
		return nil, false
	}
	return o.PreferredGlobalName.Get(), o.PreferredGlobalName.IsSet()
}

// HasPreferredGlobalName returns a boolean if a field has been set.
func (o *BotPartnerSdkTokenRequest) HasPreferredGlobalName() bool {
	if o != nil && o.PreferredGlobalName.IsSet() {
		return true
	}

	return false
}

// SetPreferredGlobalName gets a reference to the given NullableString and assigns it to the PreferredGlobalName field.
func (o *BotPartnerSdkTokenRequest) SetPreferredGlobalName(v string) {
	o.PreferredGlobalName.Set(&v)
}

// SetPreferredGlobalNameNil sets the value for PreferredGlobalName to be an explicit nil
func (o *BotPartnerSdkTokenRequest) SetPreferredGlobalNameNil() {
	o.PreferredGlobalName.Set(nil)
}

// UnsetPreferredGlobalName ensures that no value is present for PreferredGlobalName, not even an explicit nil
func (o *BotPartnerSdkTokenRequest) UnsetPreferredGlobalName() {
	o.PreferredGlobalName.Unset()
}

func (o BotPartnerSdkTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BotPartnerSdkTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["external_user_id"] = o.ExternalUserId
	if o.PreferredGlobalName.IsSet() {
		toSerialize["preferred_global_name"] = o.PreferredGlobalName.Get()
	}
	return toSerialize, nil
}

func (o *BotPartnerSdkTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"external_user_id",
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

	varBotPartnerSdkTokenRequest := _BotPartnerSdkTokenRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBotPartnerSdkTokenRequest)

	if err != nil {
		return err
	}

	*o = BotPartnerSdkTokenRequest(varBotPartnerSdkTokenRequest)

	return err
}

type NullableBotPartnerSdkTokenRequest struct {
	value *BotPartnerSdkTokenRequest
	isSet bool
}

func (v NullableBotPartnerSdkTokenRequest) Get() *BotPartnerSdkTokenRequest {
	return v.value
}

func (v *NullableBotPartnerSdkTokenRequest) Set(val *BotPartnerSdkTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBotPartnerSdkTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBotPartnerSdkTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBotPartnerSdkTokenRequest(val *BotPartnerSdkTokenRequest) *NullableBotPartnerSdkTokenRequest {
	return &NullableBotPartnerSdkTokenRequest{value: val, isSet: true}
}

func (v NullableBotPartnerSdkTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBotPartnerSdkTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


