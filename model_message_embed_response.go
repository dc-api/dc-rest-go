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
	"time"
	"bytes"
	"fmt"
)

// checks if the MessageEmbedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageEmbedResponse{}

// MessageEmbedResponse struct for MessageEmbedResponse
type MessageEmbedResponse struct {
	Type string `json:"type"`
	Url NullableString `json:"url,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Color NullableInt32 `json:"color,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	Fields []MessageEmbedFieldResponse `json:"fields,omitempty"`
	Author NullableMessageEmbedAuthorResponse `json:"author,omitempty"`
	Provider NullableMessageEmbedProviderResponse `json:"provider,omitempty"`
	Image NullableMessageEmbedImageResponse `json:"image,omitempty"`
	Thumbnail NullableMessageEmbedImageResponse `json:"thumbnail,omitempty"`
	Video NullableMessageEmbedVideoResponse `json:"video,omitempty"`
	Footer NullableMessageEmbedFooterResponse `json:"footer,omitempty"`
}

type _MessageEmbedResponse MessageEmbedResponse

// NewMessageEmbedResponse instantiates a new MessageEmbedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEmbedResponse(type_ string) *MessageEmbedResponse {
	this := MessageEmbedResponse{}
	this.Type = type_
	return &this
}

// NewMessageEmbedResponseWithDefaults instantiates a new MessageEmbedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEmbedResponseWithDefaults() *MessageEmbedResponse {
	this := MessageEmbedResponse{}
	return &this
}

// GetType returns the Type field value
func (o *MessageEmbedResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageEmbedResponse) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url.Get()) {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *MessageEmbedResponse) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *MessageEmbedResponse) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *MessageEmbedResponse) UnsetUrl() {
	o.Url.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title.Get()) {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MessageEmbedResponse) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *MessageEmbedResponse) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MessageEmbedResponse) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MessageEmbedResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MessageEmbedResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MessageEmbedResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetColor() int32 {
	if o == nil || IsNil(o.Color.Get()) {
		var ret int32
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetColorOk() (*int32, bool) {
	if o == nil || IsNil(o.Color.Get()) {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableInt32 and assigns it to the Color field.
func (o *MessageEmbedResponse) SetColor(v int32) {
	o.Color.Set(&v)
}

// SetColorNil sets the value for Color to be an explicit nil
func (o *MessageEmbedResponse) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *MessageEmbedResponse) UnsetColor() {
	o.Color.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp.Get()) {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *MessageEmbedResponse) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *MessageEmbedResponse) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *MessageEmbedResponse) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetFields() []MessageEmbedFieldResponse {
	if o == nil {
		var ret []MessageEmbedFieldResponse
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetFieldsOk() ([]MessageEmbedFieldResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []MessageEmbedFieldResponse and assigns it to the Fields field.
func (o *MessageEmbedResponse) SetFields(v []MessageEmbedFieldResponse) {
	o.Fields = v
}


// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetAuthor() MessageEmbedAuthorResponse {
	if o == nil || IsNil(o.Author.Get()) {
		var ret MessageEmbedAuthorResponse
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetAuthorOk() (*MessageEmbedAuthorResponse, bool) {
	if o == nil || IsNil(o.Author.Get()) {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasAuthor() bool {
	if o != nil && o.Author.IsSet() {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given NullableMessageEmbedAuthorResponse and assigns it to the Author field.
func (o *MessageEmbedResponse) SetAuthor(v MessageEmbedAuthorResponse) {
	o.Author.Set(&v)
}

// SetAuthorNil sets the value for Author to be an explicit nil
func (o *MessageEmbedResponse) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnsetAuthor ensures that no value is present for Author, not even an explicit nil
func (o *MessageEmbedResponse) UnsetAuthor() {
	o.Author.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetProvider() MessageEmbedProviderResponse {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret MessageEmbedProviderResponse
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetProviderOk() (*MessageEmbedProviderResponse, bool) {
	if o == nil || IsNil(o.Provider.Get()) {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableMessageEmbedProviderResponse and assigns it to the Provider field.
func (o *MessageEmbedResponse) SetProvider(v MessageEmbedProviderResponse) {
	o.Provider.Set(&v)
}

// SetProviderNil sets the value for Provider to be an explicit nil
func (o *MessageEmbedResponse) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *MessageEmbedResponse) UnsetProvider() {
	o.Provider.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetImage() MessageEmbedImageResponse {
	if o == nil || IsNil(o.Image.Get()) {
		var ret MessageEmbedImageResponse
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetImageOk() (*MessageEmbedImageResponse, bool) {
	if o == nil || IsNil(o.Image.Get()) {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableMessageEmbedImageResponse and assigns it to the Image field.
func (o *MessageEmbedResponse) SetImage(v MessageEmbedImageResponse) {
	o.Image.Set(&v)
}

// SetImageNil sets the value for Image to be an explicit nil
func (o *MessageEmbedResponse) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *MessageEmbedResponse) UnsetImage() {
	o.Image.Unset()
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetThumbnail() MessageEmbedImageResponse {
	if o == nil || IsNil(o.Thumbnail.Get()) {
		var ret MessageEmbedImageResponse
		return ret
	}
	return *o.Thumbnail.Get()
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetThumbnailOk() (*MessageEmbedImageResponse, bool) {
	if o == nil || IsNil(o.Thumbnail.Get()) {
		return nil, false
	}
	return o.Thumbnail.Get(), o.Thumbnail.IsSet()
}

// HasThumbnail returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasThumbnail() bool {
	if o != nil && o.Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given NullableMessageEmbedImageResponse and assigns it to the Thumbnail field.
func (o *MessageEmbedResponse) SetThumbnail(v MessageEmbedImageResponse) {
	o.Thumbnail.Set(&v)
}

// SetThumbnailNil sets the value for Thumbnail to be an explicit nil
func (o *MessageEmbedResponse) SetThumbnailNil() {
	o.Thumbnail.Set(nil)
}

// UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
func (o *MessageEmbedResponse) UnsetThumbnail() {
	o.Thumbnail.Unset()
}

// GetVideo returns the Video field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetVideo() MessageEmbedVideoResponse {
	if o == nil || IsNil(o.Video.Get()) {
		var ret MessageEmbedVideoResponse
		return ret
	}
	return *o.Video.Get()
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetVideoOk() (*MessageEmbedVideoResponse, bool) {
	if o == nil || IsNil(o.Video.Get()) {
		return nil, false
	}
	return o.Video.Get(), o.Video.IsSet()
}

// HasVideo returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasVideo() bool {
	if o != nil && o.Video.IsSet() {
		return true
	}

	return false
}

// SetVideo gets a reference to the given NullableMessageEmbedVideoResponse and assigns it to the Video field.
func (o *MessageEmbedResponse) SetVideo(v MessageEmbedVideoResponse) {
	o.Video.Set(&v)
}

// SetVideoNil sets the value for Video to be an explicit nil
func (o *MessageEmbedResponse) SetVideoNil() {
	o.Video.Set(nil)
}

// UnsetVideo ensures that no value is present for Video, not even an explicit nil
func (o *MessageEmbedResponse) UnsetVideo() {
	o.Video.Unset()
}

// GetFooter returns the Footer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedResponse) GetFooter() MessageEmbedFooterResponse {
	if o == nil || IsNil(o.Footer.Get()) {
		var ret MessageEmbedFooterResponse
		return ret
	}
	return *o.Footer.Get()
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedResponse) GetFooterOk() (*MessageEmbedFooterResponse, bool) {
	if o == nil || IsNil(o.Footer.Get()) {
		return nil, false
	}
	return o.Footer.Get(), o.Footer.IsSet()
}

// HasFooter returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasFooter() bool {
	if o != nil && o.Footer.IsSet() {
		return true
	}

	return false
}

// SetFooter gets a reference to the given NullableMessageEmbedFooterResponse and assigns it to the Footer field.
func (o *MessageEmbedResponse) SetFooter(v MessageEmbedFooterResponse) {
	o.Footer.Set(&v)
}

// SetFooterNil sets the value for Footer to be an explicit nil
func (o *MessageEmbedResponse) SetFooterNil() {
	o.Footer.Set(nil)
}

// UnsetFooter ensures that no value is present for Footer, not even an explicit nil
func (o *MessageEmbedResponse) UnsetFooter() {
	o.Footer.Unset()
}

func (o MessageEmbedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageEmbedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Author.IsSet() {
		toSerialize["author"] = o.Author.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.Thumbnail.IsSet() {
		toSerialize["thumbnail"] = o.Thumbnail.Get()
	}
	if o.Video.IsSet() {
		toSerialize["video"] = o.Video.Get()
	}
	if o.Footer.IsSet() {
		toSerialize["footer"] = o.Footer.Get()
	}
	return toSerialize, nil
}

func (o *MessageEmbedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varMessageEmbedResponse := _MessageEmbedResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageEmbedResponse)

	if err != nil {
		return err
	}

	*o = MessageEmbedResponse(varMessageEmbedResponse)

	return err
}

type NullableMessageEmbedResponse struct {
	value *MessageEmbedResponse
	isSet bool
}

func (v NullableMessageEmbedResponse) Get() *MessageEmbedResponse {
	return v.value
}

func (v *NullableMessageEmbedResponse) Set(val *MessageEmbedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEmbedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEmbedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEmbedResponse(val *MessageEmbedResponse) *NullableMessageEmbedResponse {
	return &NullableMessageEmbedResponse{value: val, isSet: true}
}

func (v NullableMessageEmbedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEmbedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


