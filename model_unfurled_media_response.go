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

// checks if the UnfurledMediaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnfurledMediaResponse{}

// UnfurledMediaResponse struct for UnfurledMediaResponse
type UnfurledMediaResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Url string `json:"url"`
	ProxyUrl string `json:"proxy_url"`
	Width NullableInt32 `json:"width,omitempty"`
	Height NullableInt32 `json:"height,omitempty"`
	ContentType NullableString `json:"content_type,omitempty"`
	AttachmentId *string `json:"attachment_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _UnfurledMediaResponse UnfurledMediaResponse

// NewUnfurledMediaResponse instantiates a new UnfurledMediaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnfurledMediaResponse(id string, url string, proxyUrl string) *UnfurledMediaResponse {
	this := UnfurledMediaResponse{}
	this.Id = id
	this.Url = url
	this.ProxyUrl = proxyUrl
	return &this
}

// NewUnfurledMediaResponseWithDefaults instantiates a new UnfurledMediaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnfurledMediaResponseWithDefaults() *UnfurledMediaResponse {
	this := UnfurledMediaResponse{}
	return &this
}

// GetId returns the Id field value
func (o *UnfurledMediaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnfurledMediaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnfurledMediaResponse) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *UnfurledMediaResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UnfurledMediaResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UnfurledMediaResponse) SetUrl(v string) {
	o.Url = v
}

// GetProxyUrl returns the ProxyUrl field value
func (o *UnfurledMediaResponse) GetProxyUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProxyUrl
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value
// and a boolean to check if the value has been set.
func (o *UnfurledMediaResponse) GetProxyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProxyUrl, true
}

// SetProxyUrl sets field value
func (o *UnfurledMediaResponse) SetProxyUrl(v string) {
	o.ProxyUrl = v
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnfurledMediaResponse) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnfurledMediaResponse) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width.Get()) {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *UnfurledMediaResponse) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *UnfurledMediaResponse) SetWidth(v int32) {
	o.Width.Set(&v)
}

// SetWidthNil sets the value for Width to be an explicit nil
func (o *UnfurledMediaResponse) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *UnfurledMediaResponse) UnsetWidth() {
	o.Width.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnfurledMediaResponse) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnfurledMediaResponse) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height.Get()) {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *UnfurledMediaResponse) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *UnfurledMediaResponse) SetHeight(v int32) {
	o.Height.Set(&v)
}

// SetHeightNil sets the value for Height to be an explicit nil
func (o *UnfurledMediaResponse) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *UnfurledMediaResponse) UnsetHeight() {
	o.Height.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnfurledMediaResponse) GetContentType() string {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnfurledMediaResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType.Get()) {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *UnfurledMediaResponse) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *UnfurledMediaResponse) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *UnfurledMediaResponse) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *UnfurledMediaResponse) UnsetContentType() {
	o.ContentType.Unset()
}

// GetAttachmentId returns the AttachmentId field value if set, zero value otherwise.
func (o *UnfurledMediaResponse) GetAttachmentId() string {
	if o == nil || IsNil(o.AttachmentId) {
		var ret string
		return ret
	}
	return *o.AttachmentId
}

// GetAttachmentIdOk returns a tuple with the AttachmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnfurledMediaResponse) GetAttachmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AttachmentId) {
		return nil, false
	}
	return o.AttachmentId, true
}

// HasAttachmentId returns a boolean if a field has been set.
func (o *UnfurledMediaResponse) HasAttachmentId() bool {
	if o != nil && !IsNil(o.AttachmentId) {
		return true
	}

	return false
}

// SetAttachmentId gets a reference to the given string and assigns it to the AttachmentId field.
func (o *UnfurledMediaResponse) SetAttachmentId(v string) {
	o.AttachmentId = &v
}


func (o UnfurledMediaResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnfurledMediaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["proxy_url"] = o.ProxyUrl
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if !IsNil(o.AttachmentId) {
		toSerialize["attachment_id"] = o.AttachmentId
	}
	return toSerialize, nil
}

func (o *UnfurledMediaResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"proxy_url",
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

	varUnfurledMediaResponse := _UnfurledMediaResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnfurledMediaResponse)

	if err != nil {
		return err
	}

	*o = UnfurledMediaResponse(varUnfurledMediaResponse)

	return err
}

type NullableUnfurledMediaResponse struct {
	value *UnfurledMediaResponse
	isSet bool
}

func (v NullableUnfurledMediaResponse) Get() *UnfurledMediaResponse {
	return v.value
}

func (v *NullableUnfurledMediaResponse) Set(val *UnfurledMediaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnfurledMediaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnfurledMediaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnfurledMediaResponse(val *UnfurledMediaResponse) *NullableUnfurledMediaResponse {
	return &NullableUnfurledMediaResponse{value: val, isSet: true}
}

func (v NullableUnfurledMediaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnfurledMediaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


