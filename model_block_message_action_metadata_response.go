/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-08-08T14:09:23.736426080Z[Etc/UTC]
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

// checks if the BlockMessageActionMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockMessageActionMetadataResponse{}

// BlockMessageActionMetadataResponse struct for BlockMessageActionMetadataResponse
type BlockMessageActionMetadataResponse struct {
	CustomMessage NullableString `json:"custom_message,omitempty"`
}

// NewBlockMessageActionMetadataResponse instantiates a new BlockMessageActionMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockMessageActionMetadataResponse() *BlockMessageActionMetadataResponse {
	this := BlockMessageActionMetadataResponse{}
	return &this
}

// NewBlockMessageActionMetadataResponseWithDefaults instantiates a new BlockMessageActionMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockMessageActionMetadataResponseWithDefaults() *BlockMessageActionMetadataResponse {
	this := BlockMessageActionMetadataResponse{}
	return &this
}

// GetCustomMessage returns the CustomMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockMessageActionMetadataResponse) GetCustomMessage() string {
	if o == nil || IsNil(o.CustomMessage.Get()) {
		var ret string
		return ret
	}
	return *o.CustomMessage.Get()
}

// GetCustomMessageOk returns a tuple with the CustomMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockMessageActionMetadataResponse) GetCustomMessageOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMessage.Get()) {
		return nil, false
	}
	return o.CustomMessage.Get(), o.CustomMessage.IsSet()
}

// HasCustomMessage returns a boolean if a field has been set.
func (o *BlockMessageActionMetadataResponse) HasCustomMessage() bool {
	if o != nil && o.CustomMessage.IsSet() {
		return true
	}

	return false
}

// SetCustomMessage gets a reference to the given NullableString and assigns it to the CustomMessage field.
func (o *BlockMessageActionMetadataResponse) SetCustomMessage(v string) {
	o.CustomMessage.Set(&v)
}

// SetCustomMessageNil sets the value for CustomMessage to be an explicit nil
func (o *BlockMessageActionMetadataResponse) SetCustomMessageNil() {
	o.CustomMessage.Set(nil)
}

// UnsetCustomMessage ensures that no value is present for CustomMessage, not even an explicit nil
func (o *BlockMessageActionMetadataResponse) UnsetCustomMessage() {
	o.CustomMessage.Unset()
}

func (o BlockMessageActionMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockMessageActionMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomMessage.IsSet() {
		toSerialize["custom_message"] = o.CustomMessage.Get()
	}
	return toSerialize, nil
}

type NullableBlockMessageActionMetadataResponse struct {
	value *BlockMessageActionMetadataResponse
	isSet bool
}

func (v NullableBlockMessageActionMetadataResponse) Get() *BlockMessageActionMetadataResponse {
	return v.value
}

func (v *NullableBlockMessageActionMetadataResponse) Set(val *BlockMessageActionMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockMessageActionMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockMessageActionMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockMessageActionMetadataResponse(val *BlockMessageActionMetadataResponse) *NullableBlockMessageActionMetadataResponse {
	return &NullableBlockMessageActionMetadataResponse{value: val, isSet: true}
}

func (v NullableBlockMessageActionMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockMessageActionMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


