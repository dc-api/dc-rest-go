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

// checks if the MinimalContentMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinimalContentMessageResponse{}

// MinimalContentMessageResponse struct for MinimalContentMessageResponse
type MinimalContentMessageResponse struct {
	Type int32 `json:"type"`
	Content string `json:"content"`
	Mentions []UserResponse `json:"mentions"`
	MentionRoles []string `json:"mention_roles"`
	Attachments []MessageAttachmentResponse `json:"attachments"`
	Embeds []MessageEmbedResponse `json:"embeds"`
	Timestamp time.Time `json:"timestamp"`
	EditedTimestamp NullableTime `json:"edited_timestamp,omitempty"`
	Flags int32 `json:"flags"`
	Components []BasicMessageResponseComponentsInner `json:"components"`
	Resolved NullableResolvedObjectsResponse `json:"resolved,omitempty"`
	Stickers []GetSticker200Response `json:"stickers,omitempty"`
	StickerItems []MessageStickerItemResponse `json:"sticker_items,omitempty"`
}

type _MinimalContentMessageResponse MinimalContentMessageResponse

// NewMinimalContentMessageResponse instantiates a new MinimalContentMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalContentMessageResponse(type_ int32, content string, mentions []UserResponse, mentionRoles []string, attachments []MessageAttachmentResponse, embeds []MessageEmbedResponse, timestamp time.Time, flags int32, components []BasicMessageResponseComponentsInner) *MinimalContentMessageResponse {
	this := MinimalContentMessageResponse{}
	this.Type = type_
	this.Content = content
	this.Mentions = mentions
	this.MentionRoles = mentionRoles
	this.Attachments = attachments
	this.Embeds = embeds
	this.Timestamp = timestamp
	this.Flags = flags
	this.Components = components
	return &this
}

// NewMinimalContentMessageResponseWithDefaults instantiates a new MinimalContentMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalContentMessageResponseWithDefaults() *MinimalContentMessageResponse {
	this := MinimalContentMessageResponse{}
	return &this
}

// GetType returns the Type field value
func (o *MinimalContentMessageResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MinimalContentMessageResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MinimalContentMessageResponse) SetType(v int32) {
	o.Type = v
}

// GetContent returns the Content field value
func (o *MinimalContentMessageResponse) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *MinimalContentMessageResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *MinimalContentMessageResponse) SetContent(v string) {
	o.Content = v
}

// GetMentions returns the Mentions field value
func (o *MinimalContentMessageResponse) GetMentions() []UserResponse {
	if o == nil {
		var ret []UserResponse
		return ret
	}

	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value
// and a boolean to check if the value has been set.
func (o *MinimalContentMessageResponse) GetMentionsOk() ([]UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mentions, true
}

// SetMentions sets field value
func (o *MinimalContentMessageResponse) SetMentions(v []UserResponse) {
	o.Mentions = v
}

// GetMentionRoles returns the MentionRoles field value
func (o *MinimalContentMessageResponse) GetMentionRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MentionRoles
}

// GetMentionRolesOk returns a tuple with the MentionRoles field value
// and a boolean to check if the value has been set.
func (o *MinimalContentMessageResponse) GetMentionRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionRoles, true
}

// SetMentionRoles sets field value
func (o *MinimalContentMessageResponse) SetMentionRoles(v []string) {
	o.MentionRoles = v
}

// GetAttachments returns the Attachments field value
func (o *MinimalContentMessageResponse) GetAttachments() []MessageAttachmentResponse {
	if o == nil {
		var ret []MessageAttachmentResponse
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *MinimalContentMessageResponse) GetAttachmentsOk() ([]MessageAttachmentResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// SetAttachments sets field value
func (o *MinimalContentMessageResponse) SetAttachments(v []MessageAttachmentResponse) {
	o.Attachments = v
}

// GetEmbeds returns the Embeds field value
func (o *MinimalContentMessageResponse) GetEmbeds() []MessageEmbedResponse {
	if o == nil {
		var ret []MessageEmbedResponse
		return ret
	}

	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value
// and a boolean to check if the value has been set.
func (o *MinimalContentMessageResponse) GetEmbedsOk() ([]MessageEmbedResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embeds, true
}

// SetEmbeds sets field value
func (o *MinimalContentMessageResponse) SetEmbeds(v []MessageEmbedResponse) {
	o.Embeds = v
}

// GetTimestamp returns the Timestamp field value
func (o *MinimalContentMessageResponse) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MinimalContentMessageResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MinimalContentMessageResponse) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetEditedTimestamp returns the EditedTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MinimalContentMessageResponse) GetEditedTimestamp() time.Time {
	if o == nil || IsNil(o.EditedTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EditedTimestamp.Get()
}

// GetEditedTimestampOk returns a tuple with the EditedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MinimalContentMessageResponse) GetEditedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EditedTimestamp.Get()) {
		return nil, false
	}
	return o.EditedTimestamp.Get(), o.EditedTimestamp.IsSet()
}

// HasEditedTimestamp returns a boolean if a field has been set.
func (o *MinimalContentMessageResponse) HasEditedTimestamp() bool {
	if o != nil && o.EditedTimestamp.IsSet() {
		return true
	}

	return false
}

// SetEditedTimestamp gets a reference to the given NullableTime and assigns it to the EditedTimestamp field.
func (o *MinimalContentMessageResponse) SetEditedTimestamp(v time.Time) {
	o.EditedTimestamp.Set(&v)
}

// SetEditedTimestampNil sets the value for EditedTimestamp to be an explicit nil
func (o *MinimalContentMessageResponse) SetEditedTimestampNil() {
	o.EditedTimestamp.Set(nil)
}

// UnsetEditedTimestamp ensures that no value is present for EditedTimestamp, not even an explicit nil
func (o *MinimalContentMessageResponse) UnsetEditedTimestamp() {
	o.EditedTimestamp.Unset()
}

// GetFlags returns the Flags field value
func (o *MinimalContentMessageResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *MinimalContentMessageResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *MinimalContentMessageResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetComponents returns the Components field value
func (o *MinimalContentMessageResponse) GetComponents() []BasicMessageResponseComponentsInner {
	if o == nil {
		var ret []BasicMessageResponseComponentsInner
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *MinimalContentMessageResponse) GetComponentsOk() ([]BasicMessageResponseComponentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *MinimalContentMessageResponse) SetComponents(v []BasicMessageResponseComponentsInner) {
	o.Components = v
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MinimalContentMessageResponse) GetResolved() ResolvedObjectsResponse {
	if o == nil || IsNil(o.Resolved.Get()) {
		var ret ResolvedObjectsResponse
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MinimalContentMessageResponse) GetResolvedOk() (*ResolvedObjectsResponse, bool) {
	if o == nil || IsNil(o.Resolved.Get()) {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *MinimalContentMessageResponse) HasResolved() bool {
	if o != nil && o.Resolved.IsSet() {
		return true
	}

	return false
}

// SetResolved gets a reference to the given NullableResolvedObjectsResponse and assigns it to the Resolved field.
func (o *MinimalContentMessageResponse) SetResolved(v ResolvedObjectsResponse) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil
func (o *MinimalContentMessageResponse) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
func (o *MinimalContentMessageResponse) UnsetResolved() {
	o.Resolved.Unset()
}

// GetStickers returns the Stickers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MinimalContentMessageResponse) GetStickers() []GetSticker200Response {
	if o == nil {
		var ret []GetSticker200Response
		return ret
	}
	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MinimalContentMessageResponse) GetStickersOk() ([]GetSticker200Response, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stickers, true
}

// HasStickers returns a boolean if a field has been set.
func (o *MinimalContentMessageResponse) HasStickers() bool {
	if o != nil && !IsNil(o.Stickers) {
		return true
	}

	return false
}

// SetStickers gets a reference to the given []GetSticker200Response and assigns it to the Stickers field.
func (o *MinimalContentMessageResponse) SetStickers(v []GetSticker200Response) {
	o.Stickers = v
}


// GetStickerItems returns the StickerItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MinimalContentMessageResponse) GetStickerItems() []MessageStickerItemResponse {
	if o == nil {
		var ret []MessageStickerItemResponse
		return ret
	}
	return o.StickerItems
}

// GetStickerItemsOk returns a tuple with the StickerItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MinimalContentMessageResponse) GetStickerItemsOk() ([]MessageStickerItemResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.StickerItems, true
}

// HasStickerItems returns a boolean if a field has been set.
func (o *MinimalContentMessageResponse) HasStickerItems() bool {
	if o != nil && !IsNil(o.StickerItems) {
		return true
	}

	return false
}

// SetStickerItems gets a reference to the given []MessageStickerItemResponse and assigns it to the StickerItems field.
func (o *MinimalContentMessageResponse) SetStickerItems(v []MessageStickerItemResponse) {
	o.StickerItems = v
}


func (o MinimalContentMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinimalContentMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["content"] = o.Content
	toSerialize["mentions"] = o.Mentions
	toSerialize["mention_roles"] = o.MentionRoles
	toSerialize["attachments"] = o.Attachments
	toSerialize["embeds"] = o.Embeds
	toSerialize["timestamp"] = o.Timestamp
	if o.EditedTimestamp.IsSet() {
		toSerialize["edited_timestamp"] = o.EditedTimestamp.Get()
	}
	toSerialize["flags"] = o.Flags
	toSerialize["components"] = o.Components
	if o.Resolved.IsSet() {
		toSerialize["resolved"] = o.Resolved.Get()
	}
	if o.Stickers != nil {
		toSerialize["stickers"] = o.Stickers
	}
	if o.StickerItems != nil {
		toSerialize["sticker_items"] = o.StickerItems
	}
	return toSerialize, nil
}

func (o *MinimalContentMessageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"content",
		"mentions",
		"mention_roles",
		"attachments",
		"embeds",
		"timestamp",
		"flags",
		"components",
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

	varMinimalContentMessageResponse := _MinimalContentMessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMinimalContentMessageResponse)

	if err != nil {
		return err
	}

	*o = MinimalContentMessageResponse(varMinimalContentMessageResponse)

	return err
}

type NullableMinimalContentMessageResponse struct {
	value *MinimalContentMessageResponse
	isSet bool
}

func (v NullableMinimalContentMessageResponse) Get() *MinimalContentMessageResponse {
	return v.value
}

func (v *NullableMinimalContentMessageResponse) Set(val *MinimalContentMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalContentMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalContentMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalContentMessageResponse(val *MinimalContentMessageResponse) *NullableMinimalContentMessageResponse {
	return &NullableMinimalContentMessageResponse{value: val, isSet: true}
}

func (v NullableMinimalContentMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalContentMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


