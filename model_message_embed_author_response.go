/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-28T06:15:04.313783603Z[Etc/UTC]
 *    * - **Generator Version**: 7.19.0
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
	Url *string `json:"url,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
	ProxyIconUrl *string `json:"proxy_icon_url,omitempty"`
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

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MessageEmbedAuthorResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedAuthorResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageEmbedAuthorResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MessageEmbedAuthorResponse) SetUrl(v string) {
	o.Url = &v
}


// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *MessageEmbedAuthorResponse) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedAuthorResponse) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *MessageEmbedAuthorResponse) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *MessageEmbedAuthorResponse) SetIconUrl(v string) {
	o.IconUrl = &v
}


// GetProxyIconUrl returns the ProxyIconUrl field value if set, zero value otherwise.
func (o *MessageEmbedAuthorResponse) GetProxyIconUrl() string {
	if o == nil || IsNil(o.ProxyIconUrl) {
		var ret string
		return ret
	}
	return *o.ProxyIconUrl
}

// GetProxyIconUrlOk returns a tuple with the ProxyIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedAuthorResponse) GetProxyIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyIconUrl) {
		return nil, false
	}
	return o.ProxyIconUrl, true
}

// HasProxyIconUrl returns a boolean if a field has been set.
func (o *MessageEmbedAuthorResponse) HasProxyIconUrl() bool {
	if o != nil && !IsNil(o.ProxyIconUrl) {
		return true
	}

	return false
}

// SetProxyIconUrl gets a reference to the given string and assigns it to the ProxyIconUrl field.
func (o *MessageEmbedAuthorResponse) SetProxyIconUrl(v string) {
	o.ProxyIconUrl = &v
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
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.IconUrl) {
		toSerialize["icon_url"] = o.IconUrl
	}
	if !IsNil(o.ProxyIconUrl) {
		toSerialize["proxy_icon_url"] = o.ProxyIconUrl
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

