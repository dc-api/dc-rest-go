/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
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
)

// checks if the WebhookSlackEmbed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSlackEmbed{}

// WebhookSlackEmbed struct for WebhookSlackEmbed
type WebhookSlackEmbed struct {
	Title *string `json:"title,omitempty"`
	TitleLink *string `json:"title_link,omitempty"`
	Text *string `json:"text,omitempty"`
	Color *string `json:"color,omitempty" validate:"regexp=^#(([0-9a-fA-F]{2}){3}|([0-9a-fA-F]){3})$"`
	Ts *int32 `json:"ts,omitempty"`
	Pretext *string `json:"pretext,omitempty"`
	Footer *string `json:"footer,omitempty"`
	FooterIcon *string `json:"footer_icon,omitempty"`
	AuthorName *string `json:"author_name,omitempty"`
	AuthorLink *string `json:"author_link,omitempty"`
	AuthorIcon *string `json:"author_icon,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	ThumbUrl *string `json:"thumb_url,omitempty"`
	Fields []WebhookSlackEmbedField `json:"fields,omitempty"`
}

// NewWebhookSlackEmbed instantiates a new WebhookSlackEmbed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSlackEmbed() *WebhookSlackEmbed {
	this := WebhookSlackEmbed{}
	return &this
}

// NewWebhookSlackEmbedWithDefaults instantiates a new WebhookSlackEmbed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSlackEmbedWithDefaults() *WebhookSlackEmbed {
	this := WebhookSlackEmbed{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WebhookSlackEmbed) SetTitle(v string) {
	o.Title = &v
}


// GetTitleLink returns the TitleLink field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetTitleLink() string {
	if o == nil || IsNil(o.TitleLink) {
		var ret string
		return ret
	}
	return *o.TitleLink
}

// GetTitleLinkOk returns a tuple with the TitleLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetTitleLinkOk() (*string, bool) {
	if o == nil || IsNil(o.TitleLink) {
		return nil, false
	}
	return o.TitleLink, true
}

// HasTitleLink returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasTitleLink() bool {
	if o != nil && !IsNil(o.TitleLink) {
		return true
	}

	return false
}

// SetTitleLink gets a reference to the given string and assigns it to the TitleLink field.
func (o *WebhookSlackEmbed) SetTitleLink(v string) {
	o.TitleLink = &v
}


// GetText returns the Text field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *WebhookSlackEmbed) SetText(v string) {
	o.Text = &v
}


// GetColor returns the Color field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *WebhookSlackEmbed) SetColor(v string) {
	o.Color = &v
}


// GetTs returns the Ts field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetTs() int32 {
	if o == nil || IsNil(o.Ts) {
		var ret int32
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetTsOk() (*int32, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given int32 and assigns it to the Ts field.
func (o *WebhookSlackEmbed) SetTs(v int32) {
	o.Ts = &v
}


// GetPretext returns the Pretext field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetPretext() string {
	if o == nil || IsNil(o.Pretext) {
		var ret string
		return ret
	}
	return *o.Pretext
}

// GetPretextOk returns a tuple with the Pretext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetPretextOk() (*string, bool) {
	if o == nil || IsNil(o.Pretext) {
		return nil, false
	}
	return o.Pretext, true
}

// HasPretext returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasPretext() bool {
	if o != nil && !IsNil(o.Pretext) {
		return true
	}

	return false
}

// SetPretext gets a reference to the given string and assigns it to the Pretext field.
func (o *WebhookSlackEmbed) SetPretext(v string) {
	o.Pretext = &v
}


// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetFooter() string {
	if o == nil || IsNil(o.Footer) {
		var ret string
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetFooterOk() (*string, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given string and assigns it to the Footer field.
func (o *WebhookSlackEmbed) SetFooter(v string) {
	o.Footer = &v
}


// GetFooterIcon returns the FooterIcon field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetFooterIcon() string {
	if o == nil || IsNil(o.FooterIcon) {
		var ret string
		return ret
	}
	return *o.FooterIcon
}

// GetFooterIconOk returns a tuple with the FooterIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetFooterIconOk() (*string, bool) {
	if o == nil || IsNil(o.FooterIcon) {
		return nil, false
	}
	return o.FooterIcon, true
}

// HasFooterIcon returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasFooterIcon() bool {
	if o != nil && !IsNil(o.FooterIcon) {
		return true
	}

	return false
}

// SetFooterIcon gets a reference to the given string and assigns it to the FooterIcon field.
func (o *WebhookSlackEmbed) SetFooterIcon(v string) {
	o.FooterIcon = &v
}


// GetAuthorName returns the AuthorName field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetAuthorName() string {
	if o == nil || IsNil(o.AuthorName) {
		var ret string
		return ret
	}
	return *o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetAuthorNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorName) {
		return nil, false
	}
	return o.AuthorName, true
}

// HasAuthorName returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasAuthorName() bool {
	if o != nil && !IsNil(o.AuthorName) {
		return true
	}

	return false
}

// SetAuthorName gets a reference to the given string and assigns it to the AuthorName field.
func (o *WebhookSlackEmbed) SetAuthorName(v string) {
	o.AuthorName = &v
}


// GetAuthorLink returns the AuthorLink field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetAuthorLink() string {
	if o == nil || IsNil(o.AuthorLink) {
		var ret string
		return ret
	}
	return *o.AuthorLink
}

// GetAuthorLinkOk returns a tuple with the AuthorLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetAuthorLinkOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorLink) {
		return nil, false
	}
	return o.AuthorLink, true
}

// HasAuthorLink returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasAuthorLink() bool {
	if o != nil && !IsNil(o.AuthorLink) {
		return true
	}

	return false
}

// SetAuthorLink gets a reference to the given string and assigns it to the AuthorLink field.
func (o *WebhookSlackEmbed) SetAuthorLink(v string) {
	o.AuthorLink = &v
}


// GetAuthorIcon returns the AuthorIcon field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetAuthorIcon() string {
	if o == nil || IsNil(o.AuthorIcon) {
		var ret string
		return ret
	}
	return *o.AuthorIcon
}

// GetAuthorIconOk returns a tuple with the AuthorIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetAuthorIconOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorIcon) {
		return nil, false
	}
	return o.AuthorIcon, true
}

// HasAuthorIcon returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasAuthorIcon() bool {
	if o != nil && !IsNil(o.AuthorIcon) {
		return true
	}

	return false
}

// SetAuthorIcon gets a reference to the given string and assigns it to the AuthorIcon field.
func (o *WebhookSlackEmbed) SetAuthorIcon(v string) {
	o.AuthorIcon = &v
}


// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *WebhookSlackEmbed) SetImageUrl(v string) {
	o.ImageUrl = &v
}


// GetThumbUrl returns the ThumbUrl field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetThumbUrl() string {
	if o == nil || IsNil(o.ThumbUrl) {
		var ret string
		return ret
	}
	return *o.ThumbUrl
}

// GetThumbUrlOk returns a tuple with the ThumbUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetThumbUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbUrl) {
		return nil, false
	}
	return o.ThumbUrl, true
}

// HasThumbUrl returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasThumbUrl() bool {
	if o != nil && !IsNil(o.ThumbUrl) {
		return true
	}

	return false
}

// SetThumbUrl gets a reference to the given string and assigns it to the ThumbUrl field.
func (o *WebhookSlackEmbed) SetThumbUrl(v string) {
	o.ThumbUrl = &v
}


// GetFields returns the Fields field value if set, zero value otherwise.
func (o *WebhookSlackEmbed) GetFields() []WebhookSlackEmbedField {
	if o == nil || IsNil(o.Fields) {
		var ret []WebhookSlackEmbedField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbed) GetFieldsOk() ([]WebhookSlackEmbedField, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []WebhookSlackEmbedField and assigns it to the Fields field.
func (o *WebhookSlackEmbed) SetFields(v []WebhookSlackEmbedField) {
	o.Fields = v
}


func (o WebhookSlackEmbed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSlackEmbed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.TitleLink) {
		toSerialize["title_link"] = o.TitleLink
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Pretext) {
		toSerialize["pretext"] = o.Pretext
	}
	if !IsNil(o.Footer) {
		toSerialize["footer"] = o.Footer
	}
	if !IsNil(o.FooterIcon) {
		toSerialize["footer_icon"] = o.FooterIcon
	}
	if !IsNil(o.AuthorName) {
		toSerialize["author_name"] = o.AuthorName
	}
	if !IsNil(o.AuthorLink) {
		toSerialize["author_link"] = o.AuthorLink
	}
	if !IsNil(o.AuthorIcon) {
		toSerialize["author_icon"] = o.AuthorIcon
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.ThumbUrl) {
		toSerialize["thumb_url"] = o.ThumbUrl
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableWebhookSlackEmbed struct {
	value *WebhookSlackEmbed
	isSet bool
}

func (v NullableWebhookSlackEmbed) Get() *WebhookSlackEmbed {
	return v.value
}

func (v *NullableWebhookSlackEmbed) Set(val *WebhookSlackEmbed) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSlackEmbed) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSlackEmbed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSlackEmbed(val *WebhookSlackEmbed) *NullableWebhookSlackEmbed {
	return &NullableWebhookSlackEmbed{value: val, isSet: true}
}

func (v NullableWebhookSlackEmbed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSlackEmbed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

