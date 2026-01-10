/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-10T02:53:50.800755799Z[Etc/UTC]
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

// checks if the MessageEditRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageEditRequestPartial{}

// MessageEditRequestPartial struct for MessageEditRequestPartial
type MessageEditRequestPartial struct {
	Content *string `json:"content,omitempty"`
	Embeds []RichEmbed `json:"embeds,omitempty"`
	Flags *int32 `json:"flags,omitempty"`
	AllowedMentions *BaseCreateMessageCreateRequestAllowedMentions `json:"allowed_mentions,omitempty"`
	StickerIds []string `json:"sticker_ids,omitempty"`
	Components []BaseCreateMessageCreateRequestComponentsInner `json:"components,omitempty"`
	Attachments []MessageAttachmentRequest `json:"attachments,omitempty"`
}

// NewMessageEditRequestPartial instantiates a new MessageEditRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEditRequestPartial() *MessageEditRequestPartial {
	this := MessageEditRequestPartial{}
	return &this
}

// NewMessageEditRequestPartialWithDefaults instantiates a new MessageEditRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEditRequestPartialWithDefaults() *MessageEditRequestPartial {
	this := MessageEditRequestPartial{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *MessageEditRequestPartial) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEditRequestPartial) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *MessageEditRequestPartial) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *MessageEditRequestPartial) SetContent(v string) {
	o.Content = &v
}


// GetEmbeds returns the Embeds field value if set, zero value otherwise.
func (o *MessageEditRequestPartial) GetEmbeds() []RichEmbed {
	if o == nil || IsNil(o.Embeds) {
		var ret []RichEmbed
		return ret
	}
	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEditRequestPartial) GetEmbedsOk() ([]RichEmbed, bool) {
	if o == nil || IsNil(o.Embeds) {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *MessageEditRequestPartial) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given []RichEmbed and assigns it to the Embeds field.
func (o *MessageEditRequestPartial) SetEmbeds(v []RichEmbed) {
	o.Embeds = v
}


// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *MessageEditRequestPartial) GetFlags() int32 {
	if o == nil || IsNil(o.Flags) {
		var ret int32
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEditRequestPartial) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *MessageEditRequestPartial) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int32 and assigns it to the Flags field.
func (o *MessageEditRequestPartial) SetFlags(v int32) {
	o.Flags = &v
}


// GetAllowedMentions returns the AllowedMentions field value if set, zero value otherwise.
func (o *MessageEditRequestPartial) GetAllowedMentions() BaseCreateMessageCreateRequestAllowedMentions {
	if o == nil || IsNil(o.AllowedMentions) {
		var ret BaseCreateMessageCreateRequestAllowedMentions
		return ret
	}
	return *o.AllowedMentions
}

// GetAllowedMentionsOk returns a tuple with the AllowedMentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEditRequestPartial) GetAllowedMentionsOk() (*BaseCreateMessageCreateRequestAllowedMentions, bool) {
	if o == nil || IsNil(o.AllowedMentions) {
		return nil, false
	}
	return o.AllowedMentions, true
}

// HasAllowedMentions returns a boolean if a field has been set.
func (o *MessageEditRequestPartial) HasAllowedMentions() bool {
	if o != nil && !IsNil(o.AllowedMentions) {
		return true
	}

	return false
}

// SetAllowedMentions gets a reference to the given BaseCreateMessageCreateRequestAllowedMentions and assigns it to the AllowedMentions field.
func (o *MessageEditRequestPartial) SetAllowedMentions(v BaseCreateMessageCreateRequestAllowedMentions) {
	o.AllowedMentions = &v
}


// GetStickerIds returns the StickerIds field value if set, zero value otherwise.
func (o *MessageEditRequestPartial) GetStickerIds() []string {
	if o == nil || IsNil(o.StickerIds) {
		var ret []string
		return ret
	}
	return o.StickerIds
}

// GetStickerIdsOk returns a tuple with the StickerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEditRequestPartial) GetStickerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StickerIds) {
		return nil, false
	}
	return o.StickerIds, true
}

// HasStickerIds returns a boolean if a field has been set.
func (o *MessageEditRequestPartial) HasStickerIds() bool {
	if o != nil && !IsNil(o.StickerIds) {
		return true
	}

	return false
}

// SetStickerIds gets a reference to the given []string and assigns it to the StickerIds field.
func (o *MessageEditRequestPartial) SetStickerIds(v []string) {
	o.StickerIds = v
}


// GetComponents returns the Components field value if set, zero value otherwise.
func (o *MessageEditRequestPartial) GetComponents() []BaseCreateMessageCreateRequestComponentsInner {
	if o == nil || IsNil(o.Components) {
		var ret []BaseCreateMessageCreateRequestComponentsInner
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEditRequestPartial) GetComponentsOk() ([]BaseCreateMessageCreateRequestComponentsInner, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *MessageEditRequestPartial) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []BaseCreateMessageCreateRequestComponentsInner and assigns it to the Components field.
func (o *MessageEditRequestPartial) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner) {
	o.Components = v
}


// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *MessageEditRequestPartial) GetAttachments() []MessageAttachmentRequest {
	if o == nil || IsNil(o.Attachments) {
		var ret []MessageAttachmentRequest
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEditRequestPartial) GetAttachmentsOk() ([]MessageAttachmentRequest, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *MessageEditRequestPartial) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []MessageAttachmentRequest and assigns it to the Attachments field.
func (o *MessageEditRequestPartial) SetAttachments(v []MessageAttachmentRequest) {
	o.Attachments = v
}


func (o MessageEditRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageEditRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Embeds) {
		toSerialize["embeds"] = o.Embeds
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.AllowedMentions) {
		toSerialize["allowed_mentions"] = o.AllowedMentions
	}
	if !IsNil(o.StickerIds) {
		toSerialize["sticker_ids"] = o.StickerIds
	}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableMessageEditRequestPartial struct {
	value *MessageEditRequestPartial
	isSet bool
}

func (v NullableMessageEditRequestPartial) Get() *MessageEditRequestPartial {
	return v.value
}

func (v *NullableMessageEditRequestPartial) Set(val *MessageEditRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEditRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEditRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEditRequestPartial(val *MessageEditRequestPartial) *NullableMessageEditRequestPartial {
	return &NullableMessageEditRequestPartial{value: val, isSet: true}
}

func (v NullableMessageEditRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEditRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

