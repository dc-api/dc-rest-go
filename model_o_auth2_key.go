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

// checks if the OAuth2Key type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2Key{}

// OAuth2Key struct for OAuth2Key
type OAuth2Key struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	N string `json:"n"`
	E string `json:"e"`
	Alg string `json:"alg"`
}

type _OAuth2Key OAuth2Key

// NewOAuth2Key instantiates a new OAuth2Key object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Key(kty string, use string, kid string, n string, e string, alg string) *OAuth2Key {
	this := OAuth2Key{}
	this.Kty = kty
	this.Use = use
	this.Kid = kid
	this.N = n
	this.E = e
	this.Alg = alg
	return &this
}

// NewOAuth2KeyWithDefaults instantiates a new OAuth2Key object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2KeyWithDefaults() *OAuth2Key {
	this := OAuth2Key{}
	return &this
}

// GetKty returns the Kty field value
func (o *OAuth2Key) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *OAuth2Key) SetKty(v string) {
	o.Kty = v
}

// GetUse returns the Use field value
func (o *OAuth2Key) GetUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Use
}

// GetUseOk returns a tuple with the Use field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Use, true
}

// SetUse sets field value
func (o *OAuth2Key) SetUse(v string) {
	o.Use = v
}

// GetKid returns the Kid field value
func (o *OAuth2Key) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *OAuth2Key) SetKid(v string) {
	o.Kid = v
}

// GetN returns the N field value
func (o *OAuth2Key) GetN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N
}

// GetNOk returns a tuple with the N field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N, true
}

// SetN sets field value
func (o *OAuth2Key) SetN(v string) {
	o.N = v
}

// GetE returns the E field value
func (o *OAuth2Key) GetE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.E
}

// GetEOk returns a tuple with the E field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.E, true
}

// SetE sets field value
func (o *OAuth2Key) SetE(v string) {
	o.E = v
}

// GetAlg returns the Alg field value
func (o *OAuth2Key) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *OAuth2Key) SetAlg(v string) {
	o.Alg = v
}

func (o OAuth2Key) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2Key) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kty"] = o.Kty
	toSerialize["use"] = o.Use
	toSerialize["kid"] = o.Kid
	toSerialize["n"] = o.N
	toSerialize["e"] = o.E
	toSerialize["alg"] = o.Alg
	return toSerialize, nil
}

func (o *OAuth2Key) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kty",
		"use",
		"kid",
		"n",
		"e",
		"alg",
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

	varOAuth2Key := _OAuth2Key{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2Key)

	if err != nil {
		return err
	}

	*o = OAuth2Key(varOAuth2Key)

	return err
}

type NullableOAuth2Key struct {
	value *OAuth2Key
	isSet bool
}

func (v NullableOAuth2Key) Get() *OAuth2Key {
	return v.value
}

func (v *NullableOAuth2Key) Set(val *OAuth2Key) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Key) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Key) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Key(val *OAuth2Key) *NullableOAuth2Key {
	return &NullableOAuth2Key{value: val, isSet: true}
}

func (v NullableOAuth2Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Key) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


