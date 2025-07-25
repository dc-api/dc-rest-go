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

// checks if the WelcomeMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WelcomeMessageResponse{}

// WelcomeMessageResponse struct for WelcomeMessageResponse
type WelcomeMessageResponse struct {
	AuthorIds []string `json:"author_ids"`
	Message string `json:"message"`
}

type _WelcomeMessageResponse WelcomeMessageResponse

// NewWelcomeMessageResponse instantiates a new WelcomeMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWelcomeMessageResponse(authorIds []string, message string) *WelcomeMessageResponse {
	this := WelcomeMessageResponse{}
	this.AuthorIds = authorIds
	this.Message = message
	return &this
}

// NewWelcomeMessageResponseWithDefaults instantiates a new WelcomeMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWelcomeMessageResponseWithDefaults() *WelcomeMessageResponse {
	this := WelcomeMessageResponse{}
	return &this
}

// GetAuthorIds returns the AuthorIds field value
func (o *WelcomeMessageResponse) GetAuthorIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorIds
}

// GetAuthorIdsOk returns a tuple with the AuthorIds field value
// and a boolean to check if the value has been set.
func (o *WelcomeMessageResponse) GetAuthorIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorIds, true
}

// SetAuthorIds sets field value
func (o *WelcomeMessageResponse) SetAuthorIds(v []string) {
	o.AuthorIds = v
}

// GetMessage returns the Message field value
func (o *WelcomeMessageResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *WelcomeMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *WelcomeMessageResponse) SetMessage(v string) {
	o.Message = v
}

func (o WelcomeMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WelcomeMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["author_ids"] = o.AuthorIds
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *WelcomeMessageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"author_ids",
		"message",
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

	varWelcomeMessageResponse := _WelcomeMessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWelcomeMessageResponse)

	if err != nil {
		return err
	}

	*o = WelcomeMessageResponse(varWelcomeMessageResponse)

	return err
}

type NullableWelcomeMessageResponse struct {
	value *WelcomeMessageResponse
	isSet bool
}

func (v NullableWelcomeMessageResponse) Get() *WelcomeMessageResponse {
	return v.value
}

func (v *NullableWelcomeMessageResponse) Set(val *WelcomeMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWelcomeMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWelcomeMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWelcomeMessageResponse(val *WelcomeMessageResponse) *NullableWelcomeMessageResponse {
	return &NullableWelcomeMessageResponse{value: val, isSet: true}
}

func (v NullableWelcomeMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWelcomeMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


