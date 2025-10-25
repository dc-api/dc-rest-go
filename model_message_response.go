/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-10-25T01:29:55.152026602Z[Etc/UTC]
 *    * - **Generator Version**: 7.16.0
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

// checks if the MessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageResponse{}

// MessageResponse struct for MessageResponse
type MessageResponse struct {
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
	Stickers []GetSticker200Response `json:"stickers,omitempty"`
	StickerItems []MessageStickerItemResponse `json:"sticker_items,omitempty"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Author UserResponse `json:"author"`
	Pinned bool `json:"pinned"`
	MentionEveryone bool `json:"mention_everyone"`
	Tts bool `json:"tts"`
	Call *MessageCallResponse `json:"call,omitempty"`
	Activity map[string]interface{} `json:"activity,omitempty"`
	Application *BasicApplicationResponse `json:"application,omitempty"`
	ApplicationId *string `json:"application_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Interaction *MessageInteractionResponse `json:"interaction,omitempty"`
	Nonce NullableBasicMessageResponseNonce `json:"nonce,omitempty"`
	WebhookId *string `json:"webhook_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	MessageReference *MessageReferenceResponse `json:"message_reference,omitempty"`
	Thread *ThreadResponse `json:"thread,omitempty"`
	MentionChannels []*MessageMentionChannelResponse `json:"mention_channels,omitempty"`
	RoleSubscriptionData *MessageRoleSubscriptionDataResponse `json:"role_subscription_data,omitempty"`
	PurchaseNotification *PurchaseNotificationResponse `json:"purchase_notification,omitempty"`
	Position *int32 `json:"position,omitempty"`
	Resolved *ResolvedObjectsResponse `json:"resolved,omitempty"`
	Poll *PollResponse `json:"poll,omitempty"`
	SharedClientTheme NullableCustomClientThemeResponse `json:"shared_client_theme,omitempty"`
	InteractionMetadata *BasicMessageResponseInteractionMetadata `json:"interaction_metadata,omitempty"`
	MessageSnapshots []MessageSnapshotResponse `json:"message_snapshots,omitempty"`
	Reactions []MessageReactionResponse `json:"reactions,omitempty"`
	ReferencedMessage NullableBasicMessageResponse `json:"referenced_message,omitempty"`
}

type _MessageResponse MessageResponse

// NewMessageResponse instantiates a new MessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageResponse(type_ int32, content string, mentions []UserResponse, mentionRoles []string, attachments []MessageAttachmentResponse, embeds []MessageEmbedResponse, timestamp time.Time, flags int32, components []BasicMessageResponseComponentsInner, id string, channelId string, author UserResponse, pinned bool, mentionEveryone bool, tts bool) *MessageResponse {
	this := MessageResponse{}
	this.Type = type_
	this.Content = content
	this.Mentions = mentions
	this.MentionRoles = mentionRoles
	this.Attachments = attachments
	this.Embeds = embeds
	this.Timestamp = timestamp
	this.Flags = flags
	this.Components = components
	this.Id = id
	this.ChannelId = channelId
	this.Author = author
	this.Pinned = pinned
	this.MentionEveryone = mentionEveryone
	this.Tts = tts
	return &this
}

// NewMessageResponseWithDefaults instantiates a new MessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageResponseWithDefaults() *MessageResponse {
	this := MessageResponse{}
	return &this
}

// GetType returns the Type field value
func (o *MessageResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageResponse) SetType(v int32) {
	o.Type = v
}

// GetContent returns the Content field value
func (o *MessageResponse) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *MessageResponse) SetContent(v string) {
	o.Content = v
}

// GetMentions returns the Mentions field value
func (o *MessageResponse) GetMentions() []UserResponse {
	if o == nil {
		var ret []UserResponse
		return ret
	}

	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetMentionsOk() ([]UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mentions, true
}

// SetMentions sets field value
func (o *MessageResponse) SetMentions(v []UserResponse) {
	o.Mentions = v
}

// GetMentionRoles returns the MentionRoles field value
func (o *MessageResponse) GetMentionRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MentionRoles
}

// GetMentionRolesOk returns a tuple with the MentionRoles field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetMentionRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionRoles, true
}

// SetMentionRoles sets field value
func (o *MessageResponse) SetMentionRoles(v []string) {
	o.MentionRoles = v
}

// GetAttachments returns the Attachments field value
func (o *MessageResponse) GetAttachments() []MessageAttachmentResponse {
	if o == nil {
		var ret []MessageAttachmentResponse
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetAttachmentsOk() ([]MessageAttachmentResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// SetAttachments sets field value
func (o *MessageResponse) SetAttachments(v []MessageAttachmentResponse) {
	o.Attachments = v
}

// GetEmbeds returns the Embeds field value
func (o *MessageResponse) GetEmbeds() []MessageEmbedResponse {
	if o == nil {
		var ret []MessageEmbedResponse
		return ret
	}

	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetEmbedsOk() ([]MessageEmbedResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embeds, true
}

// SetEmbeds sets field value
func (o *MessageResponse) SetEmbeds(v []MessageEmbedResponse) {
	o.Embeds = v
}

// GetTimestamp returns the Timestamp field value
func (o *MessageResponse) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MessageResponse) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetEditedTimestamp returns the EditedTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageResponse) GetEditedTimestamp() time.Time {
	if o == nil || IsNil(o.EditedTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EditedTimestamp.Get()
}

// GetEditedTimestampOk returns a tuple with the EditedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageResponse) GetEditedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EditedTimestamp.Get()) {
		return nil, false
	}
	return o.EditedTimestamp.Get(), o.EditedTimestamp.IsSet()
}

// HasEditedTimestamp returns a boolean if a field has been set.
func (o *MessageResponse) HasEditedTimestamp() bool {
	if o != nil && o.EditedTimestamp.IsSet() {
		return true
	}

	return false
}

// SetEditedTimestamp gets a reference to the given NullableTime and assigns it to the EditedTimestamp field.
func (o *MessageResponse) SetEditedTimestamp(v time.Time) {
	o.EditedTimestamp.Set(&v)
}

// SetEditedTimestampNil sets the value for EditedTimestamp to be an explicit nil
func (o *MessageResponse) SetEditedTimestampNil() {
	o.EditedTimestamp.Set(nil)
}

// UnsetEditedTimestamp ensures that no value is present for EditedTimestamp, not even an explicit nil
func (o *MessageResponse) UnsetEditedTimestamp() {
	o.EditedTimestamp.Unset()
}

// GetFlags returns the Flags field value
func (o *MessageResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *MessageResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetComponents returns the Components field value
func (o *MessageResponse) GetComponents() []BasicMessageResponseComponentsInner {
	if o == nil {
		var ret []BasicMessageResponseComponentsInner
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetComponentsOk() ([]BasicMessageResponseComponentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *MessageResponse) SetComponents(v []BasicMessageResponseComponentsInner) {
	o.Components = v
}

// GetStickers returns the Stickers field value if set, zero value otherwise.
func (o *MessageResponse) GetStickers() []GetSticker200Response {
	if o == nil || IsNil(o.Stickers) {
		var ret []GetSticker200Response
		return ret
	}
	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetStickersOk() ([]GetSticker200Response, bool) {
	if o == nil || IsNil(o.Stickers) {
		return nil, false
	}
	return o.Stickers, true
}

// HasStickers returns a boolean if a field has been set.
func (o *MessageResponse) HasStickers() bool {
	if o != nil && !IsNil(o.Stickers) {
		return true
	}

	return false
}

// SetStickers gets a reference to the given []GetSticker200Response and assigns it to the Stickers field.
func (o *MessageResponse) SetStickers(v []GetSticker200Response) {
	o.Stickers = v
}


// GetStickerItems returns the StickerItems field value if set, zero value otherwise.
func (o *MessageResponse) GetStickerItems() []MessageStickerItemResponse {
	if o == nil || IsNil(o.StickerItems) {
		var ret []MessageStickerItemResponse
		return ret
	}
	return o.StickerItems
}

// GetStickerItemsOk returns a tuple with the StickerItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetStickerItemsOk() ([]MessageStickerItemResponse, bool) {
	if o == nil || IsNil(o.StickerItems) {
		return nil, false
	}
	return o.StickerItems, true
}

// HasStickerItems returns a boolean if a field has been set.
func (o *MessageResponse) HasStickerItems() bool {
	if o != nil && !IsNil(o.StickerItems) {
		return true
	}

	return false
}

// SetStickerItems gets a reference to the given []MessageStickerItemResponse and assigns it to the StickerItems field.
func (o *MessageResponse) SetStickerItems(v []MessageStickerItemResponse) {
	o.StickerItems = v
}


// GetId returns the Id field value
func (o *MessageResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageResponse) SetId(v string) {
	o.Id = v
}

// GetChannelId returns the ChannelId field value
func (o *MessageResponse) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *MessageResponse) SetChannelId(v string) {
	o.ChannelId = v
}

// GetAuthor returns the Author field value
func (o *MessageResponse) GetAuthor() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetAuthorOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *MessageResponse) SetAuthor(v UserResponse) {
	o.Author = v
}

// GetPinned returns the Pinned field value
func (o *MessageResponse) GetPinned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pinned
}

// GetPinnedOk returns a tuple with the Pinned field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pinned, true
}

// SetPinned sets field value
func (o *MessageResponse) SetPinned(v bool) {
	o.Pinned = v
}

// GetMentionEveryone returns the MentionEveryone field value
func (o *MessageResponse) GetMentionEveryone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MentionEveryone
}

// GetMentionEveryoneOk returns a tuple with the MentionEveryone field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetMentionEveryoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MentionEveryone, true
}

// SetMentionEveryone sets field value
func (o *MessageResponse) SetMentionEveryone(v bool) {
	o.MentionEveryone = v
}

// GetTts returns the Tts field value
func (o *MessageResponse) GetTts() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Tts
}

// GetTtsOk returns a tuple with the Tts field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetTtsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tts, true
}

// SetTts sets field value
func (o *MessageResponse) SetTts(v bool) {
	o.Tts = v
}

// GetCall returns the Call field value if set, zero value otherwise.
func (o *MessageResponse) GetCall() MessageCallResponse {
	if o == nil || IsNil(o.Call) {
		var ret MessageCallResponse
		return ret
	}
	return *o.Call
}

// GetCallOk returns a tuple with the Call field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetCallOk() (*MessageCallResponse, bool) {
	if o == nil || IsNil(o.Call) {
		return nil, false
	}
	return o.Call, true
}

// HasCall returns a boolean if a field has been set.
func (o *MessageResponse) HasCall() bool {
	if o != nil && !IsNil(o.Call) {
		return true
	}

	return false
}

// SetCall gets a reference to the given MessageCallResponse and assigns it to the Call field.
func (o *MessageResponse) SetCall(v MessageCallResponse) {
	o.Call = &v
}


// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *MessageResponse) GetActivity() map[string]interface{} {
	if o == nil || IsNil(o.Activity) {
		var ret map[string]interface{}
		return ret
	}
	return o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetActivityOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Activity) {
		return map[string]interface{}{}, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *MessageResponse) HasActivity() bool {
	if o != nil && !IsNil(o.Activity) {
		return true
	}

	return false
}

// SetActivity gets a reference to the given map[string]interface{} and assigns it to the Activity field.
func (o *MessageResponse) SetActivity(v map[string]interface{}) {
	o.Activity = v
}


// GetApplication returns the Application field value if set, zero value otherwise.
func (o *MessageResponse) GetApplication() BasicApplicationResponse {
	if o == nil || IsNil(o.Application) {
		var ret BasicApplicationResponse
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetApplicationOk() (*BasicApplicationResponse, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *MessageResponse) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given BasicApplicationResponse and assigns it to the Application field.
func (o *MessageResponse) SetApplication(v BasicApplicationResponse) {
	o.Application = &v
}


// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *MessageResponse) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *MessageResponse) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *MessageResponse) SetApplicationId(v string) {
	o.ApplicationId = &v
}


// GetInteraction returns the Interaction field value if set, zero value otherwise.
func (o *MessageResponse) GetInteraction() MessageInteractionResponse {
	if o == nil || IsNil(o.Interaction) {
		var ret MessageInteractionResponse
		return ret
	}
	return *o.Interaction
}

// GetInteractionOk returns a tuple with the Interaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetInteractionOk() (*MessageInteractionResponse, bool) {
	if o == nil || IsNil(o.Interaction) {
		return nil, false
	}
	return o.Interaction, true
}

// HasInteraction returns a boolean if a field has been set.
func (o *MessageResponse) HasInteraction() bool {
	if o != nil && !IsNil(o.Interaction) {
		return true
	}

	return false
}

// SetInteraction gets a reference to the given MessageInteractionResponse and assigns it to the Interaction field.
func (o *MessageResponse) SetInteraction(v MessageInteractionResponse) {
	o.Interaction = &v
}


// GetNonce returns the Nonce field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageResponse) GetNonce() BasicMessageResponseNonce {
	if o == nil || IsNil(o.Nonce.Get()) {
		var ret BasicMessageResponseNonce
		return ret
	}
	return *o.Nonce.Get()
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageResponse) GetNonceOk() (*BasicMessageResponseNonce, bool) {
	if o == nil || IsNil(o.Nonce.Get()) {
		return nil, false
	}
	return o.Nonce.Get(), o.Nonce.IsSet()
}

// HasNonce returns a boolean if a field has been set.
func (o *MessageResponse) HasNonce() bool {
	if o != nil && o.Nonce.IsSet() {
		return true
	}

	return false
}

// SetNonce gets a reference to the given NullableBasicMessageResponseNonce and assigns it to the Nonce field.
func (o *MessageResponse) SetNonce(v BasicMessageResponseNonce) {
	o.Nonce.Set(&v)
}

// SetNonceNil sets the value for Nonce to be an explicit nil
func (o *MessageResponse) SetNonceNil() {
	o.Nonce.Set(nil)
}

// UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
func (o *MessageResponse) UnsetNonce() {
	o.Nonce.Unset()
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *MessageResponse) GetWebhookId() string {
	if o == nil || IsNil(o.WebhookId) {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetWebhookIdOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookId) {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *MessageResponse) HasWebhookId() bool {
	if o != nil && !IsNil(o.WebhookId) {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *MessageResponse) SetWebhookId(v string) {
	o.WebhookId = &v
}


// GetMessageReference returns the MessageReference field value if set, zero value otherwise.
func (o *MessageResponse) GetMessageReference() MessageReferenceResponse {
	if o == nil || IsNil(o.MessageReference) {
		var ret MessageReferenceResponse
		return ret
	}
	return *o.MessageReference
}

// GetMessageReferenceOk returns a tuple with the MessageReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetMessageReferenceOk() (*MessageReferenceResponse, bool) {
	if o == nil || IsNil(o.MessageReference) {
		return nil, false
	}
	return o.MessageReference, true
}

// HasMessageReference returns a boolean if a field has been set.
func (o *MessageResponse) HasMessageReference() bool {
	if o != nil && !IsNil(o.MessageReference) {
		return true
	}

	return false
}

// SetMessageReference gets a reference to the given MessageReferenceResponse and assigns it to the MessageReference field.
func (o *MessageResponse) SetMessageReference(v MessageReferenceResponse) {
	o.MessageReference = &v
}


// GetThread returns the Thread field value if set, zero value otherwise.
func (o *MessageResponse) GetThread() ThreadResponse {
	if o == nil || IsNil(o.Thread) {
		var ret ThreadResponse
		return ret
	}
	return *o.Thread
}

// GetThreadOk returns a tuple with the Thread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetThreadOk() (*ThreadResponse, bool) {
	if o == nil || IsNil(o.Thread) {
		return nil, false
	}
	return o.Thread, true
}

// HasThread returns a boolean if a field has been set.
func (o *MessageResponse) HasThread() bool {
	if o != nil && !IsNil(o.Thread) {
		return true
	}

	return false
}

// SetThread gets a reference to the given ThreadResponse and assigns it to the Thread field.
func (o *MessageResponse) SetThread(v ThreadResponse) {
	o.Thread = &v
}


// GetMentionChannels returns the MentionChannels field value if set, zero value otherwise.
func (o *MessageResponse) GetMentionChannels() []*MessageMentionChannelResponse {
	if o == nil || IsNil(o.MentionChannels) {
		var ret []*MessageMentionChannelResponse
		return ret
	}
	return o.MentionChannels
}

// GetMentionChannelsOk returns a tuple with the MentionChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetMentionChannelsOk() ([]*MessageMentionChannelResponse, bool) {
	if o == nil || IsNil(o.MentionChannels) {
		return nil, false
	}
	return o.MentionChannels, true
}

// HasMentionChannels returns a boolean if a field has been set.
func (o *MessageResponse) HasMentionChannels() bool {
	if o != nil && !IsNil(o.MentionChannels) {
		return true
	}

	return false
}

// SetMentionChannels gets a reference to the given []*MessageMentionChannelResponse and assigns it to the MentionChannels field.
func (o *MessageResponse) SetMentionChannels(v []*MessageMentionChannelResponse) {
	o.MentionChannels = v
}


// GetRoleSubscriptionData returns the RoleSubscriptionData field value if set, zero value otherwise.
func (o *MessageResponse) GetRoleSubscriptionData() MessageRoleSubscriptionDataResponse {
	if o == nil || IsNil(o.RoleSubscriptionData) {
		var ret MessageRoleSubscriptionDataResponse
		return ret
	}
	return *o.RoleSubscriptionData
}

// GetRoleSubscriptionDataOk returns a tuple with the RoleSubscriptionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetRoleSubscriptionDataOk() (*MessageRoleSubscriptionDataResponse, bool) {
	if o == nil || IsNil(o.RoleSubscriptionData) {
		return nil, false
	}
	return o.RoleSubscriptionData, true
}

// HasRoleSubscriptionData returns a boolean if a field has been set.
func (o *MessageResponse) HasRoleSubscriptionData() bool {
	if o != nil && !IsNil(o.RoleSubscriptionData) {
		return true
	}

	return false
}

// SetRoleSubscriptionData gets a reference to the given MessageRoleSubscriptionDataResponse and assigns it to the RoleSubscriptionData field.
func (o *MessageResponse) SetRoleSubscriptionData(v MessageRoleSubscriptionDataResponse) {
	o.RoleSubscriptionData = &v
}


// GetPurchaseNotification returns the PurchaseNotification field value if set, zero value otherwise.
func (o *MessageResponse) GetPurchaseNotification() PurchaseNotificationResponse {
	if o == nil || IsNil(o.PurchaseNotification) {
		var ret PurchaseNotificationResponse
		return ret
	}
	return *o.PurchaseNotification
}

// GetPurchaseNotificationOk returns a tuple with the PurchaseNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetPurchaseNotificationOk() (*PurchaseNotificationResponse, bool) {
	if o == nil || IsNil(o.PurchaseNotification) {
		return nil, false
	}
	return o.PurchaseNotification, true
}

// HasPurchaseNotification returns a boolean if a field has been set.
func (o *MessageResponse) HasPurchaseNotification() bool {
	if o != nil && !IsNil(o.PurchaseNotification) {
		return true
	}

	return false
}

// SetPurchaseNotification gets a reference to the given PurchaseNotificationResponse and assigns it to the PurchaseNotification field.
func (o *MessageResponse) SetPurchaseNotification(v PurchaseNotificationResponse) {
	o.PurchaseNotification = &v
}


// GetPosition returns the Position field value if set, zero value otherwise.
func (o *MessageResponse) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *MessageResponse) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *MessageResponse) SetPosition(v int32) {
	o.Position = &v
}


// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *MessageResponse) GetResolved() ResolvedObjectsResponse {
	if o == nil || IsNil(o.Resolved) {
		var ret ResolvedObjectsResponse
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetResolvedOk() (*ResolvedObjectsResponse, bool) {
	if o == nil || IsNil(o.Resolved) {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *MessageResponse) HasResolved() bool {
	if o != nil && !IsNil(o.Resolved) {
		return true
	}

	return false
}

// SetResolved gets a reference to the given ResolvedObjectsResponse and assigns it to the Resolved field.
func (o *MessageResponse) SetResolved(v ResolvedObjectsResponse) {
	o.Resolved = &v
}


// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *MessageResponse) GetPoll() PollResponse {
	if o == nil || IsNil(o.Poll) {
		var ret PollResponse
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetPollOk() (*PollResponse, bool) {
	if o == nil || IsNil(o.Poll) {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *MessageResponse) HasPoll() bool {
	if o != nil && !IsNil(o.Poll) {
		return true
	}

	return false
}

// SetPoll gets a reference to the given PollResponse and assigns it to the Poll field.
func (o *MessageResponse) SetPoll(v PollResponse) {
	o.Poll = &v
}


// GetSharedClientTheme returns the SharedClientTheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageResponse) GetSharedClientTheme() CustomClientThemeResponse {
	if o == nil || IsNil(o.SharedClientTheme.Get()) {
		var ret CustomClientThemeResponse
		return ret
	}
	return *o.SharedClientTheme.Get()
}

// GetSharedClientThemeOk returns a tuple with the SharedClientTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageResponse) GetSharedClientThemeOk() (*CustomClientThemeResponse, bool) {
	if o == nil || IsNil(o.SharedClientTheme.Get()) {
		return nil, false
	}
	return o.SharedClientTheme.Get(), o.SharedClientTheme.IsSet()
}

// HasSharedClientTheme returns a boolean if a field has been set.
func (o *MessageResponse) HasSharedClientTheme() bool {
	if o != nil && o.SharedClientTheme.IsSet() {
		return true
	}

	return false
}

// SetSharedClientTheme gets a reference to the given NullableCustomClientThemeResponse and assigns it to the SharedClientTheme field.
func (o *MessageResponse) SetSharedClientTheme(v CustomClientThemeResponse) {
	o.SharedClientTheme.Set(&v)
}

// SetSharedClientThemeNil sets the value for SharedClientTheme to be an explicit nil
func (o *MessageResponse) SetSharedClientThemeNil() {
	o.SharedClientTheme.Set(nil)
}

// UnsetSharedClientTheme ensures that no value is present for SharedClientTheme, not even an explicit nil
func (o *MessageResponse) UnsetSharedClientTheme() {
	o.SharedClientTheme.Unset()
}

// GetInteractionMetadata returns the InteractionMetadata field value if set, zero value otherwise.
func (o *MessageResponse) GetInteractionMetadata() BasicMessageResponseInteractionMetadata {
	if o == nil || IsNil(o.InteractionMetadata) {
		var ret BasicMessageResponseInteractionMetadata
		return ret
	}
	return *o.InteractionMetadata
}

// GetInteractionMetadataOk returns a tuple with the InteractionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetInteractionMetadataOk() (*BasicMessageResponseInteractionMetadata, bool) {
	if o == nil || IsNil(o.InteractionMetadata) {
		return nil, false
	}
	return o.InteractionMetadata, true
}

// HasInteractionMetadata returns a boolean if a field has been set.
func (o *MessageResponse) HasInteractionMetadata() bool {
	if o != nil && !IsNil(o.InteractionMetadata) {
		return true
	}

	return false
}

// SetInteractionMetadata gets a reference to the given BasicMessageResponseInteractionMetadata and assigns it to the InteractionMetadata field.
func (o *MessageResponse) SetInteractionMetadata(v BasicMessageResponseInteractionMetadata) {
	o.InteractionMetadata = &v
}


// GetMessageSnapshots returns the MessageSnapshots field value if set, zero value otherwise.
func (o *MessageResponse) GetMessageSnapshots() []MessageSnapshotResponse {
	if o == nil || IsNil(o.MessageSnapshots) {
		var ret []MessageSnapshotResponse
		return ret
	}
	return o.MessageSnapshots
}

// GetMessageSnapshotsOk returns a tuple with the MessageSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetMessageSnapshotsOk() ([]MessageSnapshotResponse, bool) {
	if o == nil || IsNil(o.MessageSnapshots) {
		return nil, false
	}
	return o.MessageSnapshots, true
}

// HasMessageSnapshots returns a boolean if a field has been set.
func (o *MessageResponse) HasMessageSnapshots() bool {
	if o != nil && !IsNil(o.MessageSnapshots) {
		return true
	}

	return false
}

// SetMessageSnapshots gets a reference to the given []MessageSnapshotResponse and assigns it to the MessageSnapshots field.
func (o *MessageResponse) SetMessageSnapshots(v []MessageSnapshotResponse) {
	o.MessageSnapshots = v
}


// GetReactions returns the Reactions field value if set, zero value otherwise.
func (o *MessageResponse) GetReactions() []MessageReactionResponse {
	if o == nil || IsNil(o.Reactions) {
		var ret []MessageReactionResponse
		return ret
	}
	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetReactionsOk() ([]MessageReactionResponse, bool) {
	if o == nil || IsNil(o.Reactions) {
		return nil, false
	}
	return o.Reactions, true
}

// HasReactions returns a boolean if a field has been set.
func (o *MessageResponse) HasReactions() bool {
	if o != nil && !IsNil(o.Reactions) {
		return true
	}

	return false
}

// SetReactions gets a reference to the given []MessageReactionResponse and assigns it to the Reactions field.
func (o *MessageResponse) SetReactions(v []MessageReactionResponse) {
	o.Reactions = v
}


// GetReferencedMessage returns the ReferencedMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageResponse) GetReferencedMessage() BasicMessageResponse {
	if o == nil || IsNil(o.ReferencedMessage.Get()) {
		var ret BasicMessageResponse
		return ret
	}
	return *o.ReferencedMessage.Get()
}

// GetReferencedMessageOk returns a tuple with the ReferencedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageResponse) GetReferencedMessageOk() (*BasicMessageResponse, bool) {
	if o == nil || IsNil(o.ReferencedMessage.Get()) {
		return nil, false
	}
	return o.ReferencedMessage.Get(), o.ReferencedMessage.IsSet()
}

// HasReferencedMessage returns a boolean if a field has been set.
func (o *MessageResponse) HasReferencedMessage() bool {
	if o != nil && o.ReferencedMessage.IsSet() {
		return true
	}

	return false
}

// SetReferencedMessage gets a reference to the given NullableBasicMessageResponse and assigns it to the ReferencedMessage field.
func (o *MessageResponse) SetReferencedMessage(v BasicMessageResponse) {
	o.ReferencedMessage.Set(&v)
}

// SetReferencedMessageNil sets the value for ReferencedMessage to be an explicit nil
func (o *MessageResponse) SetReferencedMessageNil() {
	o.ReferencedMessage.Set(nil)
}

// UnsetReferencedMessage ensures that no value is present for ReferencedMessage, not even an explicit nil
func (o *MessageResponse) UnsetReferencedMessage() {
	o.ReferencedMessage.Unset()
}

func (o MessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Stickers) {
		toSerialize["stickers"] = o.Stickers
	}
	if !IsNil(o.StickerItems) {
		toSerialize["sticker_items"] = o.StickerItems
	}
	toSerialize["id"] = o.Id
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["author"] = o.Author
	toSerialize["pinned"] = o.Pinned
	toSerialize["mention_everyone"] = o.MentionEveryone
	toSerialize["tts"] = o.Tts
	if !IsNil(o.Call) {
		toSerialize["call"] = o.Call
	}
	if !IsNil(o.Activity) {
		toSerialize["activity"] = o.Activity
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	if !IsNil(o.Interaction) {
		toSerialize["interaction"] = o.Interaction
	}
	if o.Nonce.IsSet() {
		toSerialize["nonce"] = o.Nonce.Get()
	}
	if !IsNil(o.WebhookId) {
		toSerialize["webhook_id"] = o.WebhookId
	}
	if !IsNil(o.MessageReference) {
		toSerialize["message_reference"] = o.MessageReference
	}
	if !IsNil(o.Thread) {
		toSerialize["thread"] = o.Thread
	}
	if !IsNil(o.MentionChannels) {
		toSerialize["mention_channels"] = o.MentionChannels
	}
	if !IsNil(o.RoleSubscriptionData) {
		toSerialize["role_subscription_data"] = o.RoleSubscriptionData
	}
	if !IsNil(o.PurchaseNotification) {
		toSerialize["purchase_notification"] = o.PurchaseNotification
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Resolved) {
		toSerialize["resolved"] = o.Resolved
	}
	if !IsNil(o.Poll) {
		toSerialize["poll"] = o.Poll
	}
	if o.SharedClientTheme.IsSet() {
		toSerialize["shared_client_theme"] = o.SharedClientTheme.Get()
	}
	if !IsNil(o.InteractionMetadata) {
		toSerialize["interaction_metadata"] = o.InteractionMetadata
	}
	if !IsNil(o.MessageSnapshots) {
		toSerialize["message_snapshots"] = o.MessageSnapshots
	}
	if !IsNil(o.Reactions) {
		toSerialize["reactions"] = o.Reactions
	}
	if o.ReferencedMessage.IsSet() {
		toSerialize["referenced_message"] = o.ReferencedMessage.Get()
	}
	return toSerialize, nil
}

func (o *MessageResponse) UnmarshalJSON(data []byte) (err error) {
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
		"id",
		"channel_id",
		"author",
		"pinned",
		"mention_everyone",
		"tts",
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

	varMessageResponse := _MessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageResponse)

	if err != nil {
		return err
	}

	*o = MessageResponse(varMessageResponse)

	return err
}

type NullableMessageResponse struct {
	value *MessageResponse
	isSet bool
}

func (v NullableMessageResponse) Get() *MessageResponse {
	return v.value
}

func (v *NullableMessageResponse) Set(val *MessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageResponse(val *MessageResponse) *NullableMessageResponse {
	return &NullableMessageResponse{value: val, isSet: true}
}

func (v NullableMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

