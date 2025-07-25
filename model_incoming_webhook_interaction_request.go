/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// checks if the IncomingWebhookInteractionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomingWebhookInteractionRequest{}

// IncomingWebhookInteractionRequest struct for IncomingWebhookInteractionRequest
type IncomingWebhookInteractionRequest struct {
	Content NullableString `json:"content,omitempty"`
	Embeds []RichEmbed `json:"embeds,omitempty"`
	AllowedMentions NullableMessageAllowedMentionsRequest `json:"allowed_mentions,omitempty"`
	Components []BaseCreateMessageCreateRequestComponentsInner `json:"components,omitempty"`
	Attachments []MessageAttachmentRequest `json:"attachments,omitempty"`
	Poll NullablePollCreateRequest `json:"poll,omitempty"`
	Tts NullableBool `json:"tts,omitempty"`
	Flags NullableInt32 `json:"flags,omitempty"`
}

// NewIncomingWebhookInteractionRequest instantiates a new IncomingWebhookInteractionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomingWebhookInteractionRequest() *IncomingWebhookInteractionRequest {
	this := IncomingWebhookInteractionRequest{}
	return &this
}

// NewIncomingWebhookInteractionRequestWithDefaults instantiates a new IncomingWebhookInteractionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomingWebhookInteractionRequestWithDefaults() *IncomingWebhookInteractionRequest {
	this := IncomingWebhookInteractionRequest{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookInteractionRequest) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookInteractionRequest) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content.Get()) {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *IncomingWebhookInteractionRequest) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *IncomingWebhookInteractionRequest) SetContent(v string) {
	o.Content.Set(&v)
}

// SetContentNil sets the value for Content to be an explicit nil
func (o *IncomingWebhookInteractionRequest) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *IncomingWebhookInteractionRequest) UnsetContent() {
	o.Content.Unset()
}

// GetEmbeds returns the Embeds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookInteractionRequest) GetEmbeds() []RichEmbed {
	if o == nil {
		var ret []RichEmbed
		return ret
	}
	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookInteractionRequest) GetEmbedsOk() ([]RichEmbed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *IncomingWebhookInteractionRequest) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given []RichEmbed and assigns it to the Embeds field.
func (o *IncomingWebhookInteractionRequest) SetEmbeds(v []RichEmbed) {
	o.Embeds = v
}


// GetAllowedMentions returns the AllowedMentions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookInteractionRequest) GetAllowedMentions() MessageAllowedMentionsRequest {
	if o == nil || IsNil(o.AllowedMentions.Get()) {
		var ret MessageAllowedMentionsRequest
		return ret
	}
	return *o.AllowedMentions.Get()
}

// GetAllowedMentionsOk returns a tuple with the AllowedMentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookInteractionRequest) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool) {
	if o == nil || IsNil(o.AllowedMentions.Get()) {
		return nil, false
	}
	return o.AllowedMentions.Get(), o.AllowedMentions.IsSet()
}

// HasAllowedMentions returns a boolean if a field has been set.
func (o *IncomingWebhookInteractionRequest) HasAllowedMentions() bool {
	if o != nil && o.AllowedMentions.IsSet() {
		return true
	}

	return false
}

// SetAllowedMentions gets a reference to the given NullableMessageAllowedMentionsRequest and assigns it to the AllowedMentions field.
func (o *IncomingWebhookInteractionRequest) SetAllowedMentions(v MessageAllowedMentionsRequest) {
	o.AllowedMentions.Set(&v)
}

// SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil
func (o *IncomingWebhookInteractionRequest) SetAllowedMentionsNil() {
	o.AllowedMentions.Set(nil)
}

// UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
func (o *IncomingWebhookInteractionRequest) UnsetAllowedMentions() {
	o.AllowedMentions.Unset()
}

// GetComponents returns the Components field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookInteractionRequest) GetComponents() []BaseCreateMessageCreateRequestComponentsInner {
	if o == nil {
		var ret []BaseCreateMessageCreateRequestComponentsInner
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookInteractionRequest) GetComponentsOk() ([]BaseCreateMessageCreateRequestComponentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *IncomingWebhookInteractionRequest) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []BaseCreateMessageCreateRequestComponentsInner and assigns it to the Components field.
func (o *IncomingWebhookInteractionRequest) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner) {
	o.Components = v
}


// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookInteractionRequest) GetAttachments() []MessageAttachmentRequest {
	if o == nil {
		var ret []MessageAttachmentRequest
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookInteractionRequest) GetAttachmentsOk() ([]MessageAttachmentRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *IncomingWebhookInteractionRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []MessageAttachmentRequest and assigns it to the Attachments field.
func (o *IncomingWebhookInteractionRequest) SetAttachments(v []MessageAttachmentRequest) {
	o.Attachments = v
}


// GetPoll returns the Poll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookInteractionRequest) GetPoll() PollCreateRequest {
	if o == nil || IsNil(o.Poll.Get()) {
		var ret PollCreateRequest
		return ret
	}
	return *o.Poll.Get()
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookInteractionRequest) GetPollOk() (*PollCreateRequest, bool) {
	if o == nil || IsNil(o.Poll.Get()) {
		return nil, false
	}
	return o.Poll.Get(), o.Poll.IsSet()
}

// HasPoll returns a boolean if a field has been set.
func (o *IncomingWebhookInteractionRequest) HasPoll() bool {
	if o != nil && o.Poll.IsSet() {
		return true
	}

	return false
}

// SetPoll gets a reference to the given NullablePollCreateRequest and assigns it to the Poll field.
func (o *IncomingWebhookInteractionRequest) SetPoll(v PollCreateRequest) {
	o.Poll.Set(&v)
}

// SetPollNil sets the value for Poll to be an explicit nil
func (o *IncomingWebhookInteractionRequest) SetPollNil() {
	o.Poll.Set(nil)
}

// UnsetPoll ensures that no value is present for Poll, not even an explicit nil
func (o *IncomingWebhookInteractionRequest) UnsetPoll() {
	o.Poll.Unset()
}

// GetTts returns the Tts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookInteractionRequest) GetTts() bool {
	if o == nil || IsNil(o.Tts.Get()) {
		var ret bool
		return ret
	}
	return *o.Tts.Get()
}

// GetTtsOk returns a tuple with the Tts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookInteractionRequest) GetTtsOk() (*bool, bool) {
	if o == nil || IsNil(o.Tts.Get()) {
		return nil, false
	}
	return o.Tts.Get(), o.Tts.IsSet()
}

// HasTts returns a boolean if a field has been set.
func (o *IncomingWebhookInteractionRequest) HasTts() bool {
	if o != nil && o.Tts.IsSet() {
		return true
	}

	return false
}

// SetTts gets a reference to the given NullableBool and assigns it to the Tts field.
func (o *IncomingWebhookInteractionRequest) SetTts(v bool) {
	o.Tts.Set(&v)
}

// SetTtsNil sets the value for Tts to be an explicit nil
func (o *IncomingWebhookInteractionRequest) SetTtsNil() {
	o.Tts.Set(nil)
}

// UnsetTts ensures that no value is present for Tts, not even an explicit nil
func (o *IncomingWebhookInteractionRequest) UnsetTts() {
	o.Tts.Unset()
}

// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookInteractionRequest) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookInteractionRequest) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags.Get()) {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *IncomingWebhookInteractionRequest) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *IncomingWebhookInteractionRequest) SetFlags(v int32) {
	o.Flags.Set(&v)
}

// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *IncomingWebhookInteractionRequest) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *IncomingWebhookInteractionRequest) UnsetFlags() {
	o.Flags.Unset()
}

func (o IncomingWebhookInteractionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomingWebhookInteractionRequest) ToMap() (map[string]interface{}, error) {
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
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Poll.IsSet() {
		toSerialize["poll"] = o.Poll.Get()
	}
	if o.Tts.IsSet() {
		toSerialize["tts"] = o.Tts.Get()
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	return toSerialize, nil
}

type NullableIncomingWebhookInteractionRequest struct {
	value *IncomingWebhookInteractionRequest
	isSet bool
}

func (v NullableIncomingWebhookInteractionRequest) Get() *IncomingWebhookInteractionRequest {
	return v.value
}

func (v *NullableIncomingWebhookInteractionRequest) Set(val *IncomingWebhookInteractionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomingWebhookInteractionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomingWebhookInteractionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomingWebhookInteractionRequest(val *IncomingWebhookInteractionRequest) *NullableIncomingWebhookInteractionRequest {
	return &NullableIncomingWebhookInteractionRequest{value: val, isSet: true}
}

func (v NullableIncomingWebhookInteractionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomingWebhookInteractionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


