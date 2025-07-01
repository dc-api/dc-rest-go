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

// checks if the EditLobbyChannelLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditLobbyChannelLinkRequest{}

// EditLobbyChannelLinkRequest struct for EditLobbyChannelLinkRequest
type EditLobbyChannelLinkRequest struct {
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

// NewEditLobbyChannelLinkRequest instantiates a new EditLobbyChannelLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditLobbyChannelLinkRequest() *EditLobbyChannelLinkRequest {
	this := EditLobbyChannelLinkRequest{}
	return &this
}

// NewEditLobbyChannelLinkRequestWithDefaults instantiates a new EditLobbyChannelLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditLobbyChannelLinkRequestWithDefaults() *EditLobbyChannelLinkRequest {
	this := EditLobbyChannelLinkRequest{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *EditLobbyChannelLinkRequest) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditLobbyChannelLinkRequest) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *EditLobbyChannelLinkRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *EditLobbyChannelLinkRequest) SetChannelId(v string) {
	o.ChannelId = &v
}


func (o EditLobbyChannelLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditLobbyChannelLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	return toSerialize, nil
}

type NullableEditLobbyChannelLinkRequest struct {
	value *EditLobbyChannelLinkRequest
	isSet bool
}

func (v NullableEditLobbyChannelLinkRequest) Get() *EditLobbyChannelLinkRequest {
	return v.value
}

func (v *NullableEditLobbyChannelLinkRequest) Set(val *EditLobbyChannelLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditLobbyChannelLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditLobbyChannelLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditLobbyChannelLinkRequest(val *EditLobbyChannelLinkRequest) *NullableEditLobbyChannelLinkRequest {
	return &NullableEditLobbyChannelLinkRequest{value: val, isSet: true}
}

func (v NullableEditLobbyChannelLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditLobbyChannelLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


