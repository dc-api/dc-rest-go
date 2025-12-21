/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-12-21T03:28:05.066589019Z[Etc/UTC]
 *    * - **Generator Version**: 7.17.0
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
	Url *string `json:"url,omitempty"`
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Color *int32 `json:"color,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Fields []MessageEmbedFieldResponse `json:"fields,omitempty"`
	Author *MessageEmbedAuthorResponse `json:"author,omitempty"`
	Provider *MessageEmbedProviderResponse `json:"provider,omitempty"`
	Image *MessageEmbedImageResponse `json:"image,omitempty"`
	Thumbnail *MessageEmbedImageResponse `json:"thumbnail,omitempty"`
	Video *MessageEmbedVideoResponse `json:"video,omitempty"`
	Footer *MessageEmbedFooterResponse `json:"footer,omitempty"`
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

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MessageEmbedResponse) SetUrl(v string) {
	o.Url = &v
}


// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MessageEmbedResponse) SetTitle(v string) {
	o.Title = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MessageEmbedResponse) SetDescription(v string) {
	o.Description = &v
}


// GetColor returns the Color field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetColor() int32 {
	if o == nil || IsNil(o.Color) {
		var ret int32
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetColorOk() (*int32, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given int32 and assigns it to the Color field.
func (o *MessageEmbedResponse) SetColor(v int32) {
	o.Color = &v
}


// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MessageEmbedResponse) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}


// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetFields() []MessageEmbedFieldResponse {
	if o == nil || IsNil(o.Fields) {
		var ret []MessageEmbedFieldResponse
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetFieldsOk() ([]MessageEmbedFieldResponse, bool) {
	if o == nil || IsNil(o.Fields) {
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


// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetAuthor() MessageEmbedAuthorResponse {
	if o == nil || IsNil(o.Author) {
		var ret MessageEmbedAuthorResponse
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetAuthorOk() (*MessageEmbedAuthorResponse, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given MessageEmbedAuthorResponse and assigns it to the Author field.
func (o *MessageEmbedResponse) SetAuthor(v MessageEmbedAuthorResponse) {
	o.Author = &v
}


// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetProvider() MessageEmbedProviderResponse {
	if o == nil || IsNil(o.Provider) {
		var ret MessageEmbedProviderResponse
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetProviderOk() (*MessageEmbedProviderResponse, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given MessageEmbedProviderResponse and assigns it to the Provider field.
func (o *MessageEmbedResponse) SetProvider(v MessageEmbedProviderResponse) {
	o.Provider = &v
}


// GetImage returns the Image field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetImage() MessageEmbedImageResponse {
	if o == nil || IsNil(o.Image) {
		var ret MessageEmbedImageResponse
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetImageOk() (*MessageEmbedImageResponse, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given MessageEmbedImageResponse and assigns it to the Image field.
func (o *MessageEmbedResponse) SetImage(v MessageEmbedImageResponse) {
	o.Image = &v
}


// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetThumbnail() MessageEmbedImageResponse {
	if o == nil || IsNil(o.Thumbnail) {
		var ret MessageEmbedImageResponse
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetThumbnailOk() (*MessageEmbedImageResponse, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given MessageEmbedImageResponse and assigns it to the Thumbnail field.
func (o *MessageEmbedResponse) SetThumbnail(v MessageEmbedImageResponse) {
	o.Thumbnail = &v
}


// GetVideo returns the Video field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetVideo() MessageEmbedVideoResponse {
	if o == nil || IsNil(o.Video) {
		var ret MessageEmbedVideoResponse
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetVideoOk() (*MessageEmbedVideoResponse, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given MessageEmbedVideoResponse and assigns it to the Video field.
func (o *MessageEmbedResponse) SetVideo(v MessageEmbedVideoResponse) {
	o.Video = &v
}


// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *MessageEmbedResponse) GetFooter() MessageEmbedFooterResponse {
	if o == nil || IsNil(o.Footer) {
		var ret MessageEmbedFooterResponse
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEmbedResponse) GetFooterOk() (*MessageEmbedFooterResponse, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *MessageEmbedResponse) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given MessageEmbedFooterResponse and assigns it to the Footer field.
func (o *MessageEmbedResponse) SetFooter(v MessageEmbedFooterResponse) {
	o.Footer = &v
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
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !IsNil(o.Footer) {
		toSerialize["footer"] = o.Footer
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

