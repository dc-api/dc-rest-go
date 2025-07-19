/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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

// checks if the MessageEmbedImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageEmbedImageResponse{}

// MessageEmbedImageResponse struct for MessageEmbedImageResponse
type MessageEmbedImageResponse struct {
	Url NullableString `json:"url,omitempty"`
	ProxyUrl NullableString `json:"proxy_url,omitempty"`
	Width *int64 `json:"width,omitempty"`
	Height *int64 `json:"height,omitempty"`
	ContentType NullableString `json:"content_type,omitempty"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	PlaceholderVersion *int64 `json:"placeholder_version,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Flags *int64 `json:"flags,omitempty"`
}

// NewMessageEmbedImageResponse instantiates a new MessageEmbedImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEmbedImageResponse() *MessageEmbedImageResponse {
	this := MessageEmbedImageResponse{}
	return &this
}

// NewMessageEmbedImageResponseWithDefaults instantiates a new MessageEmbedImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEmbedImageResponseWithDefaults() *MessageEmbedImageResponse {
	this := MessageEmbedImageResponse{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedImageResponse) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedImageResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url.Get()) {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageEmbedImageResponse) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *MessageEmbedImageResponse) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *MessageEmbedImageResponse) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *MessageEmbedImageResponse) UnsetUrl() {
	o.Url.Unset()
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedImageResponse) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUrl.Get()
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedImageResponse) GetProxyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyUrl.Get()) {
		return nil, false
	}
	return o.ProxyUrl.Get(), o.ProxyUrl.IsSet()
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *MessageEmbedImageResponse) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given NullableString and assigns it to the ProxyUrl field.
func (o *MessageEmbedImageResponse) SetProxyUrl(v string) {
	o.ProxyUrl.Set(&v)
}

// SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil
func (o *MessageEmbedImageResponse) SetProxyUrlNil() {
	o.ProxyUrl.Set(nil)
}

// UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
func (o *MessageEmbedImageResponse) UnsetProxyUrl() {
	o.ProxyUrl.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *MessageEmbedImageResponse) GetWidth() int64 {
	if o == nil || IsNil(o.Width) {
		var ret int64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedImageResponse) GetWidthOk() (*int64, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *MessageEmbedImageResponse) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int64 and assigns it to the Width field.
func (o *MessageEmbedImageResponse) SetWidth(v int64) {
	o.Width = &v
}


// GetHeight returns the Height field value if set, zero value otherwise.
func (o *MessageEmbedImageResponse) GetHeight() int64 {
	if o == nil || IsNil(o.Height) {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedImageResponse) GetHeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *MessageEmbedImageResponse) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *MessageEmbedImageResponse) SetHeight(v int64) {
	o.Height = &v
}


// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedImageResponse) GetContentType() string {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedImageResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType.Get()) {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *MessageEmbedImageResponse) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *MessageEmbedImageResponse) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *MessageEmbedImageResponse) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *MessageEmbedImageResponse) UnsetContentType() {
	o.ContentType.Unset()
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedImageResponse) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedImageResponse) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder.Get()) {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *MessageEmbedImageResponse) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *MessageEmbedImageResponse) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}

// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *MessageEmbedImageResponse) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *MessageEmbedImageResponse) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetPlaceholderVersion returns the PlaceholderVersion field value if set, zero value otherwise.
func (o *MessageEmbedImageResponse) GetPlaceholderVersion() int64 {
	if o == nil || IsNil(o.PlaceholderVersion) {
		var ret int64
		return ret
	}
	return *o.PlaceholderVersion
}

// GetPlaceholderVersionOk returns a tuple with the PlaceholderVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedImageResponse) GetPlaceholderVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.PlaceholderVersion) {
		return nil, false
	}
	return o.PlaceholderVersion, true
}

// HasPlaceholderVersion returns a boolean if a field has been set.
func (o *MessageEmbedImageResponse) HasPlaceholderVersion() bool {
	if o != nil && !IsNil(o.PlaceholderVersion) {
		return true
	}

	return false
}

// SetPlaceholderVersion gets a reference to the given int64 and assigns it to the PlaceholderVersion field.
func (o *MessageEmbedImageResponse) SetPlaceholderVersion(v int64) {
	o.PlaceholderVersion = &v
}


// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedImageResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedImageResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageEmbedImageResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MessageEmbedImageResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MessageEmbedImageResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MessageEmbedImageResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *MessageEmbedImageResponse) GetFlags() int64 {
	if o == nil || IsNil(o.Flags) {
		var ret int64
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedImageResponse) GetFlagsOk() (*int64, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *MessageEmbedImageResponse) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int64 and assigns it to the Flags field.
func (o *MessageEmbedImageResponse) SetFlags(v int64) {
	o.Flags = &v
}


func (o MessageEmbedImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageEmbedImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.ProxyUrl.IsSet() {
		toSerialize["proxy_url"] = o.ProxyUrl.Get()
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if !IsNil(o.PlaceholderVersion) {
		toSerialize["placeholder_version"] = o.PlaceholderVersion
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

type NullableMessageEmbedImageResponse struct {
	value *MessageEmbedImageResponse
	isSet bool
}

func (v NullableMessageEmbedImageResponse) Get() *MessageEmbedImageResponse {
	return v.value
}

func (v *NullableMessageEmbedImageResponse) Set(val *MessageEmbedImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEmbedImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEmbedImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEmbedImageResponse(val *MessageEmbedImageResponse) *NullableMessageEmbedImageResponse {
	return &NullableMessageEmbedImageResponse{value: val, isSet: true}
}

func (v NullableMessageEmbedImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEmbedImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


