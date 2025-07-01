/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// checks if the MessageEmbedAuthorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageEmbedAuthorResponse{}

// MessageEmbedAuthorResponse struct for MessageEmbedAuthorResponse
type MessageEmbedAuthorResponse struct {
	Name string `json:"name"`
	Url NullableString `json:"url,omitempty"`
	IconUrl NullableString `json:"icon_url,omitempty"`
	ProxyIconUrl NullableString `json:"proxy_icon_url,omitempty"`
}

type _MessageEmbedAuthorResponse MessageEmbedAuthorResponse

// NewMessageEmbedAuthorResponse instantiates a new MessageEmbedAuthorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEmbedAuthorResponse(name string) *MessageEmbedAuthorResponse {
	this := MessageEmbedAuthorResponse{}
	this.Name = name
	return &this
}

// NewMessageEmbedAuthorResponseWithDefaults instantiates a new MessageEmbedAuthorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEmbedAuthorResponseWithDefaults() *MessageEmbedAuthorResponse {
	this := MessageEmbedAuthorResponse{}
	return &this
}

// GetName returns the Name field value
func (o *MessageEmbedAuthorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MessageEmbedAuthorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MessageEmbedAuthorResponse) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedAuthorResponse) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedAuthorResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url.Get()) {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageEmbedAuthorResponse) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *MessageEmbedAuthorResponse) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *MessageEmbedAuthorResponse) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *MessageEmbedAuthorResponse) UnsetUrl() {
	o.Url.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedAuthorResponse) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedAuthorResponse) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl.Get()) {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *MessageEmbedAuthorResponse) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *MessageEmbedAuthorResponse) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *MessageEmbedAuthorResponse) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *MessageEmbedAuthorResponse) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetProxyIconUrl returns the ProxyIconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedAuthorResponse) GetProxyIconUrl() string {
	if o == nil || IsNil(o.ProxyIconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyIconUrl.Get()
}

// GetProxyIconUrlOk returns a tuple with the ProxyIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedAuthorResponse) GetProxyIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyIconUrl.Get()) {
		return nil, false
	}
	return o.ProxyIconUrl.Get(), o.ProxyIconUrl.IsSet()
}

// HasProxyIconUrl returns a boolean if a field has been set.
func (o *MessageEmbedAuthorResponse) HasProxyIconUrl() bool {
	if o != nil && o.ProxyIconUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyIconUrl gets a reference to the given NullableString and assigns it to the ProxyIconUrl field.
func (o *MessageEmbedAuthorResponse) SetProxyIconUrl(v string) {
	o.ProxyIconUrl.Set(&v)
}

// SetProxyIconUrlNil sets the value for ProxyIconUrl to be an explicit nil
func (o *MessageEmbedAuthorResponse) SetProxyIconUrlNil() {
	o.ProxyIconUrl.Set(nil)
}

// UnsetProxyIconUrl ensures that no value is present for ProxyIconUrl, not even an explicit nil
func (o *MessageEmbedAuthorResponse) UnsetProxyIconUrl() {
	o.ProxyIconUrl.Unset()
}

func (o MessageEmbedAuthorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageEmbedAuthorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	if o.ProxyIconUrl.IsSet() {
		toSerialize["proxy_icon_url"] = o.ProxyIconUrl.Get()
	}
	return toSerialize, nil
}

func (o *MessageEmbedAuthorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varMessageEmbedAuthorResponse := _MessageEmbedAuthorResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageEmbedAuthorResponse)

	if err != nil {
		return err
	}

	*o = MessageEmbedAuthorResponse(varMessageEmbedAuthorResponse)

	return err
}

type NullableMessageEmbedAuthorResponse struct {
	value *MessageEmbedAuthorResponse
	isSet bool
}

func (v NullableMessageEmbedAuthorResponse) Get() *MessageEmbedAuthorResponse {
	return v.value
}

func (v *NullableMessageEmbedAuthorResponse) Set(val *MessageEmbedAuthorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEmbedAuthorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEmbedAuthorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEmbedAuthorResponse(val *MessageEmbedAuthorResponse) *NullableMessageEmbedAuthorResponse {
	return &NullableMessageEmbedAuthorResponse{value: val, isSet: true}
}

func (v NullableMessageEmbedAuthorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEmbedAuthorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


