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
)

// checks if the MessageCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageCreateRequest{}

// MessageCreateRequest struct for MessageCreateRequest
type MessageCreateRequest struct {
	Content NullableString `json:"content,omitempty"`
	Embeds []RichEmbed `json:"embeds,omitempty"`
	AllowedMentions NullableMessageAllowedMentionsRequest `json:"allowed_mentions,omitempty"`
	StickerIds []string `json:"sticker_ids,omitempty"`
	Components []BaseCreateMessageCreateRequestComponentsInner `json:"components,omitempty"`
	Flags NullableInt32 `json:"flags,omitempty"`
	Attachments []MessageAttachmentRequest `json:"attachments,omitempty"`
	Poll NullablePollCreateRequest `json:"poll,omitempty"`
	ConfettiPotion map[string]interface{} `json:"confetti_potion,omitempty"`
	MessageReference NullableMessageReferenceRequest `json:"message_reference,omitempty"`
	Nonce NullableBasicMessageResponseNonce `json:"nonce,omitempty"`
	EnforceNonce NullableBool `json:"enforce_nonce,omitempty"`
	Tts NullableBool `json:"tts,omitempty"`
}

// NewMessageCreateRequest instantiates a new MessageCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageCreateRequest() *MessageCreateRequest {
	this := MessageCreateRequest{}
	return &this
}

// NewMessageCreateRequestWithDefaults instantiates a new MessageCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageCreateRequestWithDefaults() *MessageCreateRequest {
	this := MessageCreateRequest{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content.Get()) {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *MessageCreateRequest) SetContent(v string) {
	o.Content.Set(&v)
}

// SetContentNil sets the value for Content to be an explicit nil
func (o *MessageCreateRequest) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *MessageCreateRequest) UnsetContent() {
	o.Content.Unset()
}

// GetEmbeds returns the Embeds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetEmbeds() []RichEmbed {
	if o == nil {
		var ret []RichEmbed
		return ret
	}
	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetEmbedsOk() ([]RichEmbed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given []RichEmbed and assigns it to the Embeds field.
func (o *MessageCreateRequest) SetEmbeds(v []RichEmbed) {
	o.Embeds = v
}


// GetAllowedMentions returns the AllowedMentions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetAllowedMentions() MessageAllowedMentionsRequest {
	if o == nil || IsNil(o.AllowedMentions.Get()) {
		var ret MessageAllowedMentionsRequest
		return ret
	}
	return *o.AllowedMentions.Get()
}

// GetAllowedMentionsOk returns a tuple with the AllowedMentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool) {
	if o == nil || IsNil(o.AllowedMentions.Get()) {
		return nil, false
	}
	return o.AllowedMentions.Get(), o.AllowedMentions.IsSet()
}

// HasAllowedMentions returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasAllowedMentions() bool {
	if o != nil && o.AllowedMentions.IsSet() {
		return true
	}

	return false
}

// SetAllowedMentions gets a reference to the given NullableMessageAllowedMentionsRequest and assigns it to the AllowedMentions field.
func (o *MessageCreateRequest) SetAllowedMentions(v MessageAllowedMentionsRequest) {
	o.AllowedMentions.Set(&v)
}

// SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil
func (o *MessageCreateRequest) SetAllowedMentionsNil() {
	o.AllowedMentions.Set(nil)
}

// UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
func (o *MessageCreateRequest) UnsetAllowedMentions() {
	o.AllowedMentions.Unset()
}

// GetStickerIds returns the StickerIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetStickerIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.StickerIds
}

// GetStickerIdsOk returns a tuple with the StickerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetStickerIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StickerIds, true
}

// HasStickerIds returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasStickerIds() bool {
	if o != nil && !IsNil(o.StickerIds) {
		return true
	}

	return false
}

// SetStickerIds gets a reference to the given []string and assigns it to the StickerIds field.
func (o *MessageCreateRequest) SetStickerIds(v []string) {
	o.StickerIds = v
}


// GetComponents returns the Components field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetComponents() []BaseCreateMessageCreateRequestComponentsInner {
	if o == nil {
		var ret []BaseCreateMessageCreateRequestComponentsInner
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetComponentsOk() ([]BaseCreateMessageCreateRequestComponentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []BaseCreateMessageCreateRequestComponentsInner and assigns it to the Components field.
func (o *MessageCreateRequest) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner) {
	o.Components = v
}


// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags.Get()) {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *MessageCreateRequest) SetFlags(v int32) {
	o.Flags.Set(&v)
}

// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *MessageCreateRequest) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *MessageCreateRequest) UnsetFlags() {
	o.Flags.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetAttachments() []MessageAttachmentRequest {
	if o == nil {
		var ret []MessageAttachmentRequest
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetAttachmentsOk() ([]MessageAttachmentRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []MessageAttachmentRequest and assigns it to the Attachments field.
func (o *MessageCreateRequest) SetAttachments(v []MessageAttachmentRequest) {
	o.Attachments = v
}


// GetPoll returns the Poll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetPoll() PollCreateRequest {
	if o == nil || IsNil(o.Poll.Get()) {
		var ret PollCreateRequest
		return ret
	}
	return *o.Poll.Get()
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetPollOk() (*PollCreateRequest, bool) {
	if o == nil || IsNil(o.Poll.Get()) {
		return nil, false
	}
	return o.Poll.Get(), o.Poll.IsSet()
}

// HasPoll returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasPoll() bool {
	if o != nil && o.Poll.IsSet() {
		return true
	}

	return false
}

// SetPoll gets a reference to the given NullablePollCreateRequest and assigns it to the Poll field.
func (o *MessageCreateRequest) SetPoll(v PollCreateRequest) {
	o.Poll.Set(&v)
}

// SetPollNil sets the value for Poll to be an explicit nil
func (o *MessageCreateRequest) SetPollNil() {
	o.Poll.Set(nil)
}

// UnsetPoll ensures that no value is present for Poll, not even an explicit nil
func (o *MessageCreateRequest) UnsetPoll() {
	o.Poll.Unset()
}

// GetConfettiPotion returns the ConfettiPotion field value if set, zero value otherwise.
func (o *MessageCreateRequest) GetConfettiPotion() map[string]interface{} {
	if o == nil || IsNil(o.ConfettiPotion) {
		var ret map[string]interface{}
		return ret
	}
	return o.ConfettiPotion
}

// GetConfettiPotionOk returns a tuple with the ConfettiPotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageCreateRequest) GetConfettiPotionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ConfettiPotion) {
		return map[string]interface{}{}, false
	}
	return o.ConfettiPotion, true
}

// HasConfettiPotion returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasConfettiPotion() bool {
	if o != nil && !IsNil(o.ConfettiPotion) {
		return true
	}

	return false
}

// SetConfettiPotion gets a reference to the given map[string]interface{} and assigns it to the ConfettiPotion field.
func (o *MessageCreateRequest) SetConfettiPotion(v map[string]interface{}) {
	o.ConfettiPotion = v
}


// GetMessageReference returns the MessageReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetMessageReference() MessageReferenceRequest {
	if o == nil || IsNil(o.MessageReference.Get()) {
		var ret MessageReferenceRequest
		return ret
	}
	return *o.MessageReference.Get()
}

// GetMessageReferenceOk returns a tuple with the MessageReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetMessageReferenceOk() (*MessageReferenceRequest, bool) {
	if o == nil || IsNil(o.MessageReference.Get()) {
		return nil, false
	}
	return o.MessageReference.Get(), o.MessageReference.IsSet()
}

// HasMessageReference returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasMessageReference() bool {
	if o != nil && o.MessageReference.IsSet() {
		return true
	}

	return false
}

// SetMessageReference gets a reference to the given NullableMessageReferenceRequest and assigns it to the MessageReference field.
func (o *MessageCreateRequest) SetMessageReference(v MessageReferenceRequest) {
	o.MessageReference.Set(&v)
}

// SetMessageReferenceNil sets the value for MessageReference to be an explicit nil
func (o *MessageCreateRequest) SetMessageReferenceNil() {
	o.MessageReference.Set(nil)
}

// UnsetMessageReference ensures that no value is present for MessageReference, not even an explicit nil
func (o *MessageCreateRequest) UnsetMessageReference() {
	o.MessageReference.Unset()
}

// GetNonce returns the Nonce field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetNonce() BasicMessageResponseNonce {
	if o == nil || IsNil(o.Nonce.Get()) {
		var ret BasicMessageResponseNonce
		return ret
	}
	return *o.Nonce.Get()
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetNonceOk() (*BasicMessageResponseNonce, bool) {
	if o == nil || IsNil(o.Nonce.Get()) {
		return nil, false
	}
	return o.Nonce.Get(), o.Nonce.IsSet()
}

// HasNonce returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasNonce() bool {
	if o != nil && o.Nonce.IsSet() {
		return true
	}

	return false
}

// SetNonce gets a reference to the given NullableBasicMessageResponseNonce and assigns it to the Nonce field.
func (o *MessageCreateRequest) SetNonce(v BasicMessageResponseNonce) {
	o.Nonce.Set(&v)
}

// SetNonceNil sets the value for Nonce to be an explicit nil
func (o *MessageCreateRequest) SetNonceNil() {
	o.Nonce.Set(nil)
}

// UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
func (o *MessageCreateRequest) UnsetNonce() {
	o.Nonce.Unset()
}

// GetEnforceNonce returns the EnforceNonce field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetEnforceNonce() bool {
	if o == nil || IsNil(o.EnforceNonce.Get()) {
		var ret bool
		return ret
	}
	return *o.EnforceNonce.Get()
}

// GetEnforceNonceOk returns a tuple with the EnforceNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetEnforceNonceOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceNonce.Get()) {
		return nil, false
	}
	return o.EnforceNonce.Get(), o.EnforceNonce.IsSet()
}

// HasEnforceNonce returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasEnforceNonce() bool {
	if o != nil && o.EnforceNonce.IsSet() {
		return true
	}

	return false
}

// SetEnforceNonce gets a reference to the given NullableBool and assigns it to the EnforceNonce field.
func (o *MessageCreateRequest) SetEnforceNonce(v bool) {
	o.EnforceNonce.Set(&v)
}

// SetEnforceNonceNil sets the value for EnforceNonce to be an explicit nil
func (o *MessageCreateRequest) SetEnforceNonceNil() {
	o.EnforceNonce.Set(nil)
}

// UnsetEnforceNonce ensures that no value is present for EnforceNonce, not even an explicit nil
func (o *MessageCreateRequest) UnsetEnforceNonce() {
	o.EnforceNonce.Unset()
}

// GetTts returns the Tts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCreateRequest) GetTts() bool {
	if o == nil || IsNil(o.Tts.Get()) {
		var ret bool
		return ret
	}
	return *o.Tts.Get()
}

// GetTtsOk returns a tuple with the Tts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCreateRequest) GetTtsOk() (*bool, bool) {
	if o == nil || IsNil(o.Tts.Get()) {
		return nil, false
	}
	return o.Tts.Get(), o.Tts.IsSet()
}

// HasTts returns a boolean if a field has been set.
func (o *MessageCreateRequest) HasTts() bool {
	if o != nil && o.Tts.IsSet() {
		return true
	}

	return false
}

// SetTts gets a reference to the given NullableBool and assigns it to the Tts field.
func (o *MessageCreateRequest) SetTts(v bool) {
	o.Tts.Set(&v)
}

// SetTtsNil sets the value for Tts to be an explicit nil
func (o *MessageCreateRequest) SetTtsNil() {
	o.Tts.Set(nil)
}

// UnsetTts ensures that no value is present for Tts, not even an explicit nil
func (o *MessageCreateRequest) UnsetTts() {
	o.Tts.Unset()
}

func (o MessageCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.Embeds != nil {
		toSerialize["embeds"] = o.Embeds
	}
	if o.AllowedMentions.IsSet() {
		toSerialize["allowed_mentions"] = o.AllowedMentions.Get()
	}
	if o.StickerIds != nil {
		toSerialize["sticker_ids"] = o.StickerIds
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Poll.IsSet() {
		toSerialize["poll"] = o.Poll.Get()
	}
	if !IsNil(o.ConfettiPotion) {
		toSerialize["confetti_potion"] = o.ConfettiPotion
	}
	if o.MessageReference.IsSet() {
		toSerialize["message_reference"] = o.MessageReference.Get()
	}
	if o.Nonce.IsSet() {
		toSerialize["nonce"] = o.Nonce.Get()
	}
	if o.EnforceNonce.IsSet() {
		toSerialize["enforce_nonce"] = o.EnforceNonce.Get()
	}
	if o.Tts.IsSet() {
		toSerialize["tts"] = o.Tts.Get()
	}
	return toSerialize, nil
}

type NullableMessageCreateRequest struct {
	value *MessageCreateRequest
	isSet bool
}

func (v NullableMessageCreateRequest) Get() *MessageCreateRequest {
	return v.value
}

func (v *NullableMessageCreateRequest) Set(val *MessageCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCreateRequest(val *MessageCreateRequest) *NullableMessageCreateRequest {
	return &NullableMessageCreateRequest{value: val, isSet: true}
}

func (v NullableMessageCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


