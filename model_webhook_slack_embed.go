/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
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

// checks if the WebhookSlackEmbed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSlackEmbed{}

// WebhookSlackEmbed struct for WebhookSlackEmbed
type WebhookSlackEmbed struct {
	Title NullableString `json:"title,omitempty"`
	TitleLink NullableString `json:"title_link,omitempty"`
	Text NullableString `json:"text,omitempty"`
	Color NullableString `json:"color,omitempty" validate:"regexp=^#(([0-9a-fA-F]{2}){3}|([0-9a-fA-F]){3})$"`
	Ts NullableInt32 `json:"ts,omitempty"`
	Pretext NullableString `json:"pretext,omitempty"`
	Footer NullableString `json:"footer,omitempty"`
	FooterIcon NullableString `json:"footer_icon,omitempty"`
	AuthorName NullableString `json:"author_name,omitempty"`
	AuthorLink NullableString `json:"author_link,omitempty"`
	AuthorIcon NullableString `json:"author_icon,omitempty"`
	ImageUrl NullableString `json:"image_url,omitempty"`
	ThumbUrl NullableString `json:"thumb_url,omitempty"`
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

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title.Get()) {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *WebhookSlackEmbed) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *WebhookSlackEmbed) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetTitle() {
	o.Title.Unset()
}

// GetTitleLink returns the TitleLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetTitleLink() string {
	if o == nil || IsNil(o.TitleLink.Get()) {
		var ret string
		return ret
	}
	return *o.TitleLink.Get()
}

// GetTitleLinkOk returns a tuple with the TitleLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetTitleLinkOk() (*string, bool) {
	if o == nil || IsNil(o.TitleLink.Get()) {
		return nil, false
	}
	return o.TitleLink.Get(), o.TitleLink.IsSet()
}

// HasTitleLink returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasTitleLink() bool {
	if o != nil && o.TitleLink.IsSet() {
		return true
	}

	return false
}

// SetTitleLink gets a reference to the given NullableString and assigns it to the TitleLink field.
func (o *WebhookSlackEmbed) SetTitleLink(v string) {
	o.TitleLink.Set(&v)
}

// SetTitleLinkNil sets the value for TitleLink to be an explicit nil
func (o *WebhookSlackEmbed) SetTitleLinkNil() {
	o.TitleLink.Set(nil)
}

// UnsetTitleLink ensures that no value is present for TitleLink, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetTitleLink() {
	o.TitleLink.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text.Get()) {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *WebhookSlackEmbed) SetText(v string) {
	o.Text.Set(&v)
}

// SetTextNil sets the value for Text to be an explicit nil
func (o *WebhookSlackEmbed) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetText() {
	o.Text.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color.Get()) {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *WebhookSlackEmbed) SetColor(v string) {
	o.Color.Set(&v)
}

// SetColorNil sets the value for Color to be an explicit nil
func (o *WebhookSlackEmbed) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetColor() {
	o.Color.Unset()
}

// GetTs returns the Ts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetTs() int32 {
	if o == nil || IsNil(o.Ts.Get()) {
		var ret int32
		return ret
	}
	return *o.Ts.Get()
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetTsOk() (*int32, bool) {
	if o == nil || IsNil(o.Ts.Get()) {
		return nil, false
	}
	return o.Ts.Get(), o.Ts.IsSet()
}

// HasTs returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasTs() bool {
	if o != nil && o.Ts.IsSet() {
		return true
	}

	return false
}

// SetTs gets a reference to the given NullableInt32 and assigns it to the Ts field.
func (o *WebhookSlackEmbed) SetTs(v int32) {
	o.Ts.Set(&v)
}

// SetTsNil sets the value for Ts to be an explicit nil
func (o *WebhookSlackEmbed) SetTsNil() {
	o.Ts.Set(nil)
}

// UnsetTs ensures that no value is present for Ts, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetTs() {
	o.Ts.Unset()
}

// GetPretext returns the Pretext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetPretext() string {
	if o == nil || IsNil(o.Pretext.Get()) {
		var ret string
		return ret
	}
	return *o.Pretext.Get()
}

// GetPretextOk returns a tuple with the Pretext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetPretextOk() (*string, bool) {
	if o == nil || IsNil(o.Pretext.Get()) {
		return nil, false
	}
	return o.Pretext.Get(), o.Pretext.IsSet()
}

// HasPretext returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasPretext() bool {
	if o != nil && o.Pretext.IsSet() {
		return true
	}

	return false
}

// SetPretext gets a reference to the given NullableString and assigns it to the Pretext field.
func (o *WebhookSlackEmbed) SetPretext(v string) {
	o.Pretext.Set(&v)
}

// SetPretextNil sets the value for Pretext to be an explicit nil
func (o *WebhookSlackEmbed) SetPretextNil() {
	o.Pretext.Set(nil)
}

// UnsetPretext ensures that no value is present for Pretext, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetPretext() {
	o.Pretext.Unset()
}

// GetFooter returns the Footer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetFooter() string {
	if o == nil || IsNil(o.Footer.Get()) {
		var ret string
		return ret
	}
	return *o.Footer.Get()
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetFooterOk() (*string, bool) {
	if o == nil || IsNil(o.Footer.Get()) {
		return nil, false
	}
	return o.Footer.Get(), o.Footer.IsSet()
}

// HasFooter returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasFooter() bool {
	if o != nil && o.Footer.IsSet() {
		return true
	}

	return false
}

// SetFooter gets a reference to the given NullableString and assigns it to the Footer field.
func (o *WebhookSlackEmbed) SetFooter(v string) {
	o.Footer.Set(&v)
}

// SetFooterNil sets the value for Footer to be an explicit nil
func (o *WebhookSlackEmbed) SetFooterNil() {
	o.Footer.Set(nil)
}

// UnsetFooter ensures that no value is present for Footer, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetFooter() {
	o.Footer.Unset()
}

// GetFooterIcon returns the FooterIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetFooterIcon() string {
	if o == nil || IsNil(o.FooterIcon.Get()) {
		var ret string
		return ret
	}
	return *o.FooterIcon.Get()
}

// GetFooterIconOk returns a tuple with the FooterIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetFooterIconOk() (*string, bool) {
	if o == nil || IsNil(o.FooterIcon.Get()) {
		return nil, false
	}
	return o.FooterIcon.Get(), o.FooterIcon.IsSet()
}

// HasFooterIcon returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasFooterIcon() bool {
	if o != nil && o.FooterIcon.IsSet() {
		return true
	}

	return false
}

// SetFooterIcon gets a reference to the given NullableString and assigns it to the FooterIcon field.
func (o *WebhookSlackEmbed) SetFooterIcon(v string) {
	o.FooterIcon.Set(&v)
}

// SetFooterIconNil sets the value for FooterIcon to be an explicit nil
func (o *WebhookSlackEmbed) SetFooterIconNil() {
	o.FooterIcon.Set(nil)
}

// UnsetFooterIcon ensures that no value is present for FooterIcon, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetFooterIcon() {
	o.FooterIcon.Unset()
}

// GetAuthorName returns the AuthorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetAuthorName() string {
	if o == nil || IsNil(o.AuthorName.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorName.Get()
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetAuthorNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorName.Get()) {
		return nil, false
	}
	return o.AuthorName.Get(), o.AuthorName.IsSet()
}

// HasAuthorName returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasAuthorName() bool {
	if o != nil && o.AuthorName.IsSet() {
		return true
	}

	return false
}

// SetAuthorName gets a reference to the given NullableString and assigns it to the AuthorName field.
func (o *WebhookSlackEmbed) SetAuthorName(v string) {
	o.AuthorName.Set(&v)
}

// SetAuthorNameNil sets the value for AuthorName to be an explicit nil
func (o *WebhookSlackEmbed) SetAuthorNameNil() {
	o.AuthorName.Set(nil)
}

// UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetAuthorName() {
	o.AuthorName.Unset()
}

// GetAuthorLink returns the AuthorLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetAuthorLink() string {
	if o == nil || IsNil(o.AuthorLink.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorLink.Get()
}

// GetAuthorLinkOk returns a tuple with the AuthorLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetAuthorLinkOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorLink.Get()) {
		return nil, false
	}
	return o.AuthorLink.Get(), o.AuthorLink.IsSet()
}

// HasAuthorLink returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasAuthorLink() bool {
	if o != nil && o.AuthorLink.IsSet() {
		return true
	}

	return false
}

// SetAuthorLink gets a reference to the given NullableString and assigns it to the AuthorLink field.
func (o *WebhookSlackEmbed) SetAuthorLink(v string) {
	o.AuthorLink.Set(&v)
}

// SetAuthorLinkNil sets the value for AuthorLink to be an explicit nil
func (o *WebhookSlackEmbed) SetAuthorLinkNil() {
	o.AuthorLink.Set(nil)
}

// UnsetAuthorLink ensures that no value is present for AuthorLink, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetAuthorLink() {
	o.AuthorLink.Unset()
}

// GetAuthorIcon returns the AuthorIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetAuthorIcon() string {
	if o == nil || IsNil(o.AuthorIcon.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorIcon.Get()
}

// GetAuthorIconOk returns a tuple with the AuthorIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetAuthorIconOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorIcon.Get()) {
		return nil, false
	}
	return o.AuthorIcon.Get(), o.AuthorIcon.IsSet()
}

// HasAuthorIcon returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasAuthorIcon() bool {
	if o != nil && o.AuthorIcon.IsSet() {
		return true
	}

	return false
}

// SetAuthorIcon gets a reference to the given NullableString and assigns it to the AuthorIcon field.
func (o *WebhookSlackEmbed) SetAuthorIcon(v string) {
	o.AuthorIcon.Set(&v)
}

// SetAuthorIconNil sets the value for AuthorIcon to be an explicit nil
func (o *WebhookSlackEmbed) SetAuthorIconNil() {
	o.AuthorIcon.Set(nil)
}

// UnsetAuthorIcon ensures that no value is present for AuthorIcon, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetAuthorIcon() {
	o.AuthorIcon.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *WebhookSlackEmbed) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *WebhookSlackEmbed) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetThumbUrl returns the ThumbUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetThumbUrl() string {
	if o == nil || IsNil(o.ThumbUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ThumbUrl.Get()
}

// GetThumbUrlOk returns a tuple with the ThumbUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetThumbUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbUrl.Get()) {
		return nil, false
	}
	return o.ThumbUrl.Get(), o.ThumbUrl.IsSet()
}

// HasThumbUrl returns a boolean if a field has been set.
func (o *WebhookSlackEmbed) HasThumbUrl() bool {
	if o != nil && o.ThumbUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbUrl gets a reference to the given NullableString and assigns it to the ThumbUrl field.
func (o *WebhookSlackEmbed) SetThumbUrl(v string) {
	o.ThumbUrl.Set(&v)
}

// SetThumbUrlNil sets the value for ThumbUrl to be an explicit nil
func (o *WebhookSlackEmbed) SetThumbUrlNil() {
	o.ThumbUrl.Set(nil)
}

// UnsetThumbUrl ensures that no value is present for ThumbUrl, not even an explicit nil
func (o *WebhookSlackEmbed) UnsetThumbUrl() {
	o.ThumbUrl.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbed) GetFields() []WebhookSlackEmbedField {
	if o == nil {
		var ret []WebhookSlackEmbedField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbed) GetFieldsOk() ([]WebhookSlackEmbedField, bool) {
	if o == nil {
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
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TitleLink.IsSet() {
		toSerialize["title_link"] = o.TitleLink.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Ts.IsSet() {
		toSerialize["ts"] = o.Ts.Get()
	}
	if o.Pretext.IsSet() {
		toSerialize["pretext"] = o.Pretext.Get()
	}
	if o.Footer.IsSet() {
		toSerialize["footer"] = o.Footer.Get()
	}
	if o.FooterIcon.IsSet() {
		toSerialize["footer_icon"] = o.FooterIcon.Get()
	}
	if o.AuthorName.IsSet() {
		toSerialize["author_name"] = o.AuthorName.Get()
	}
	if o.AuthorLink.IsSet() {
		toSerialize["author_link"] = o.AuthorLink.Get()
	}
	if o.AuthorIcon.IsSet() {
		toSerialize["author_icon"] = o.AuthorIcon.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.ThumbUrl.IsSet() {
		toSerialize["thumb_url"] = o.ThumbUrl.Get()
	}
	if o.Fields != nil {
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


