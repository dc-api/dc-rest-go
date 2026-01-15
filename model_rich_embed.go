/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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
)

// checks if the RichEmbed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RichEmbed{}

// RichEmbed struct for RichEmbed
type RichEmbed struct {
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
	Title *string `json:"title,omitempty"`
	Color *int32 `json:"color,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Description *string `json:"description,omitempty"`
	Author *RichEmbedAuthor `json:"author,omitempty"`
	Image *RichEmbedImage `json:"image,omitempty"`
	Thumbnail *RichEmbedThumbnail `json:"thumbnail,omitempty"`
	Footer *RichEmbedFooter `json:"footer,omitempty"`
	Fields []RichEmbedField `json:"fields,omitempty"`
	Provider *RichEmbedProvider `json:"provider,omitempty"`
	Video *RichEmbedVideo `json:"video,omitempty"`
}

// NewRichEmbed instantiates a new RichEmbed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRichEmbed() *RichEmbed {
	this := RichEmbed{}
	return &this
}

// NewRichEmbedWithDefaults instantiates a new RichEmbed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRichEmbedWithDefaults() *RichEmbed {
	this := RichEmbed{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RichEmbed) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RichEmbed) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RichEmbed) SetType(v string) {
	o.Type = &v
}


// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RichEmbed) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RichEmbed) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RichEmbed) SetUrl(v string) {
	o.Url = &v
}


// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RichEmbed) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RichEmbed) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RichEmbed) SetTitle(v string) {
	o.Title = &v
}


// GetColor returns the Color field value if set, zero value otherwise.
func (o *RichEmbed) GetColor() int32 {
	if o == nil || IsNil(o.Color) {
		var ret int32
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetColorOk() (*int32, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *RichEmbed) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given int32 and assigns it to the Color field.
func (o *RichEmbed) SetColor(v int32) {
	o.Color = &v
}


// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *RichEmbed) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *RichEmbed) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *RichEmbed) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RichEmbed) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RichEmbed) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RichEmbed) SetDescription(v string) {
	o.Description = &v
}


// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *RichEmbed) GetAuthor() RichEmbedAuthor {
	if o == nil || IsNil(o.Author) {
		var ret RichEmbedAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetAuthorOk() (*RichEmbedAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *RichEmbed) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given RichEmbedAuthor and assigns it to the Author field.
func (o *RichEmbed) SetAuthor(v RichEmbedAuthor) {
	o.Author = &v
}


// GetImage returns the Image field value if set, zero value otherwise.
func (o *RichEmbed) GetImage() RichEmbedImage {
	if o == nil || IsNil(o.Image) {
		var ret RichEmbedImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetImageOk() (*RichEmbedImage, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *RichEmbed) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given RichEmbedImage and assigns it to the Image field.
func (o *RichEmbed) SetImage(v RichEmbedImage) {
	o.Image = &v
}


// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *RichEmbed) GetThumbnail() RichEmbedThumbnail {
	if o == nil || IsNil(o.Thumbnail) {
		var ret RichEmbedThumbnail
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetThumbnailOk() (*RichEmbedThumbnail, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *RichEmbed) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given RichEmbedThumbnail and assigns it to the Thumbnail field.
func (o *RichEmbed) SetThumbnail(v RichEmbedThumbnail) {
	o.Thumbnail = &v
}


// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *RichEmbed) GetFooter() RichEmbedFooter {
	if o == nil || IsNil(o.Footer) {
		var ret RichEmbedFooter
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetFooterOk() (*RichEmbedFooter, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *RichEmbed) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given RichEmbedFooter and assigns it to the Footer field.
func (o *RichEmbed) SetFooter(v RichEmbedFooter) {
	o.Footer = &v
}


// GetFields returns the Fields field value if set, zero value otherwise.
func (o *RichEmbed) GetFields() []RichEmbedField {
	if o == nil || IsNil(o.Fields) {
		var ret []RichEmbedField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetFieldsOk() ([]RichEmbedField, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *RichEmbed) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []RichEmbedField and assigns it to the Fields field.
func (o *RichEmbed) SetFields(v []RichEmbedField) {
	o.Fields = v
}


// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *RichEmbed) GetProvider() RichEmbedProvider {
	if o == nil || IsNil(o.Provider) {
		var ret RichEmbedProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetProviderOk() (*RichEmbedProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *RichEmbed) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given RichEmbedProvider and assigns it to the Provider field.
func (o *RichEmbed) SetProvider(v RichEmbedProvider) {
	o.Provider = &v
}


// GetVideo returns the Video field value if set, zero value otherwise.
func (o *RichEmbed) GetVideo() RichEmbedVideo {
	if o == nil || IsNil(o.Video) {
		var ret RichEmbedVideo
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichEmbed) GetVideoOk() (*RichEmbedVideo, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *RichEmbed) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given RichEmbedVideo and assigns it to the Video field.
func (o *RichEmbed) SetVideo(v RichEmbedVideo) {
	o.Video = &v
}


func (o RichEmbed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RichEmbed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.Footer) {
		toSerialize["footer"] = o.Footer
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	return toSerialize, nil
}

type NullableRichEmbed struct {
	value *RichEmbed
	isSet bool
}

func (v NullableRichEmbed) Get() *RichEmbed {
	return v.value
}

func (v *NullableRichEmbed) Set(val *RichEmbed) {
	v.value = val
	v.isSet = true
}

func (v NullableRichEmbed) IsSet() bool {
	return v.isSet
}

func (v *NullableRichEmbed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRichEmbed(val *RichEmbed) *NullableRichEmbed {
	return &NullableRichEmbed{value: val, isSet: true}
}

func (v NullableRichEmbed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRichEmbed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

