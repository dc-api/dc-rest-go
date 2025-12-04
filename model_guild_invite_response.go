/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-12-04T03:54:15.906780824Z[Etc/UTC]
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

// checks if the GuildInviteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildInviteResponse{}

// GuildInviteResponse struct for GuildInviteResponse
type GuildInviteResponse struct {
	Type NullableInt32 `json:"type"`
	Code string `json:"code"`
	Inviter *UserResponse `json:"inviter,omitempty"`
	MaxAge *int32 `json:"max_age,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	IsContact *bool `json:"is_contact,omitempty"`
	Flags *int32 `json:"flags,omitempty"`
	Guild InviteGuildResponse `json:"guild"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Channel InviteChannelResponse `json:"channel"`
	TargetType NullableInt32 `json:"target_type,omitempty"`
	TargetUser *UserResponse `json:"target_user,omitempty"`
	TargetApplication *InviteApplicationResponse `json:"target_application,omitempty"`
	GuildScheduledEvent *ScheduledEventResponse `json:"guild_scheduled_event,omitempty"`
	Uses *int32 `json:"uses,omitempty"`
	MaxUses *int32 `json:"max_uses,omitempty"`
	Temporary *bool `json:"temporary,omitempty"`
	ApproximateMemberCount NullableInt32 `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount NullableInt32 `json:"approximate_presence_count,omitempty"`
	IsNicknameChangeable *bool `json:"is_nickname_changeable,omitempty"`
}

type _GuildInviteResponse GuildInviteResponse

// NewGuildInviteResponse instantiates a new GuildInviteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildInviteResponse(type_ NullableInt32, code string, guild InviteGuildResponse, guildId string, channel InviteChannelResponse) *GuildInviteResponse {
	this := GuildInviteResponse{}
	this.Type = type_
	this.Code = code
	this.Guild = guild
	this.GuildId = guildId
	this.Channel = channel
	return &this
}

// NewGuildInviteResponseWithDefaults instantiates a new GuildInviteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildInviteResponseWithDefaults() *GuildInviteResponse {
	this := GuildInviteResponse{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *GuildInviteResponse) GetType() int32 {
	if o == nil || o.Type.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *GuildInviteResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// GetCode returns the Code field value
func (o *GuildInviteResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GuildInviteResponse) SetCode(v string) {
	o.Code = v
}

// GetInviter returns the Inviter field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetInviter() UserResponse {
	if o == nil || IsNil(o.Inviter) {
		var ret UserResponse
		return ret
	}
	return *o.Inviter
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetInviterOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Inviter) {
		return nil, false
	}
	return o.Inviter, true
}

// HasInviter returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasInviter() bool {
	if o != nil && !IsNil(o.Inviter) {
		return true
	}

	return false
}

// SetInviter gets a reference to the given UserResponse and assigns it to the Inviter field.
func (o *GuildInviteResponse) SetInviter(v UserResponse) {
	o.Inviter = &v
}


// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge) {
		var ret int32
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int32 and assigns it to the MaxAge field.
func (o *GuildInviteResponse) SetMaxAge(v int32) {
	o.MaxAge = &v
}


// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GuildInviteResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}


// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *GuildInviteResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *GuildInviteResponse) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *GuildInviteResponse) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetIsContact returns the IsContact field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetIsContact() bool {
	if o == nil || IsNil(o.IsContact) {
		var ret bool
		return ret
	}
	return *o.IsContact
}

// GetIsContactOk returns a tuple with the IsContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetIsContactOk() (*bool, bool) {
	if o == nil || IsNil(o.IsContact) {
		return nil, false
	}
	return o.IsContact, true
}

// HasIsContact returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasIsContact() bool {
	if o != nil && !IsNil(o.IsContact) {
		return true
	}

	return false
}

// SetIsContact gets a reference to the given bool and assigns it to the IsContact field.
func (o *GuildInviteResponse) SetIsContact(v bool) {
	o.IsContact = &v
}


// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetFlags() int32 {
	if o == nil || IsNil(o.Flags) {
		var ret int32
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int32 and assigns it to the Flags field.
func (o *GuildInviteResponse) SetFlags(v int32) {
	o.Flags = &v
}


// GetGuild returns the Guild field value
func (o *GuildInviteResponse) GetGuild() InviteGuildResponse {
	if o == nil {
		var ret InviteGuildResponse
		return ret
	}

	return o.Guild
}

// GetGuildOk returns a tuple with the Guild field value
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetGuildOk() (*InviteGuildResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guild, true
}

// SetGuild sets field value
func (o *GuildInviteResponse) SetGuild(v InviteGuildResponse) {
	o.Guild = v
}

// GetGuildId returns the GuildId field value
func (o *GuildInviteResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *GuildInviteResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetChannel returns the Channel field value
func (o *GuildInviteResponse) GetChannel() InviteChannelResponse {
	if o == nil {
		var ret InviteChannelResponse
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetChannelOk() (*InviteChannelResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *GuildInviteResponse) SetChannel(v InviteChannelResponse) {
	o.Channel = v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetTargetType() int32 {
	if o == nil || IsNil(o.TargetType.Get()) {
		var ret int32
		return ret
	}
	return *o.TargetType.Get()
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetTargetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetType.Get()) {
		return nil, false
	}
	return o.TargetType.Get(), o.TargetType.IsSet()
}

// HasTargetType returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasTargetType() bool {
	if o != nil && o.TargetType.IsSet() {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given NullableInt32 and assigns it to the TargetType field.
func (o *GuildInviteResponse) SetTargetType(v int32) {
	o.TargetType.Set(&v)
}

// SetTargetTypeNil sets the value for TargetType to be an explicit nil
func (o *GuildInviteResponse) SetTargetTypeNil() {
	o.TargetType.Set(nil)
}

// UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
func (o *GuildInviteResponse) UnsetTargetType() {
	o.TargetType.Unset()
}

// GetTargetUser returns the TargetUser field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetTargetUser() UserResponse {
	if o == nil || IsNil(o.TargetUser) {
		var ret UserResponse
		return ret
	}
	return *o.TargetUser
}

// GetTargetUserOk returns a tuple with the TargetUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetTargetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.TargetUser) {
		return nil, false
	}
	return o.TargetUser, true
}

// HasTargetUser returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasTargetUser() bool {
	if o != nil && !IsNil(o.TargetUser) {
		return true
	}

	return false
}

// SetTargetUser gets a reference to the given UserResponse and assigns it to the TargetUser field.
func (o *GuildInviteResponse) SetTargetUser(v UserResponse) {
	o.TargetUser = &v
}


// GetTargetApplication returns the TargetApplication field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetTargetApplication() InviteApplicationResponse {
	if o == nil || IsNil(o.TargetApplication) {
		var ret InviteApplicationResponse
		return ret
	}
	return *o.TargetApplication
}

// GetTargetApplicationOk returns a tuple with the TargetApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetTargetApplicationOk() (*InviteApplicationResponse, bool) {
	if o == nil || IsNil(o.TargetApplication) {
		return nil, false
	}
	return o.TargetApplication, true
}

// HasTargetApplication returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasTargetApplication() bool {
	if o != nil && !IsNil(o.TargetApplication) {
		return true
	}

	return false
}

// SetTargetApplication gets a reference to the given InviteApplicationResponse and assigns it to the TargetApplication field.
func (o *GuildInviteResponse) SetTargetApplication(v InviteApplicationResponse) {
	o.TargetApplication = &v
}


// GetGuildScheduledEvent returns the GuildScheduledEvent field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetGuildScheduledEvent() ScheduledEventResponse {
	if o == nil || IsNil(o.GuildScheduledEvent) {
		var ret ScheduledEventResponse
		return ret
	}
	return *o.GuildScheduledEvent
}

// GetGuildScheduledEventOk returns a tuple with the GuildScheduledEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetGuildScheduledEventOk() (*ScheduledEventResponse, bool) {
	if o == nil || IsNil(o.GuildScheduledEvent) {
		return nil, false
	}
	return o.GuildScheduledEvent, true
}

// HasGuildScheduledEvent returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasGuildScheduledEvent() bool {
	if o != nil && !IsNil(o.GuildScheduledEvent) {
		return true
	}

	return false
}

// SetGuildScheduledEvent gets a reference to the given ScheduledEventResponse and assigns it to the GuildScheduledEvent field.
func (o *GuildInviteResponse) SetGuildScheduledEvent(v ScheduledEventResponse) {
	o.GuildScheduledEvent = &v
}


// GetUses returns the Uses field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetUses() int32 {
	if o == nil || IsNil(o.Uses) {
		var ret int32
		return ret
	}
	return *o.Uses
}

// GetUsesOk returns a tuple with the Uses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.Uses) {
		return nil, false
	}
	return o.Uses, true
}

// HasUses returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasUses() bool {
	if o != nil && !IsNil(o.Uses) {
		return true
	}

	return false
}

// SetUses gets a reference to the given int32 and assigns it to the Uses field.
func (o *GuildInviteResponse) SetUses(v int32) {
	o.Uses = &v
}


// GetMaxUses returns the MaxUses field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetMaxUses() int32 {
	if o == nil || IsNil(o.MaxUses) {
		var ret int32
		return ret
	}
	return *o.MaxUses
}

// GetMaxUsesOk returns a tuple with the MaxUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetMaxUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUses) {
		return nil, false
	}
	return o.MaxUses, true
}

// HasMaxUses returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasMaxUses() bool {
	if o != nil && !IsNil(o.MaxUses) {
		return true
	}

	return false
}

// SetMaxUses gets a reference to the given int32 and assigns it to the MaxUses field.
func (o *GuildInviteResponse) SetMaxUses(v int32) {
	o.MaxUses = &v
}


// GetTemporary returns the Temporary field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetTemporary() bool {
	if o == nil || IsNil(o.Temporary) {
		var ret bool
		return ret
	}
	return *o.Temporary
}

// GetTemporaryOk returns a tuple with the Temporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetTemporaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Temporary) {
		return nil, false
	}
	return o.Temporary, true
}

// HasTemporary returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasTemporary() bool {
	if o != nil && !IsNil(o.Temporary) {
		return true
	}

	return false
}

// SetTemporary gets a reference to the given bool and assigns it to the Temporary field.
func (o *GuildInviteResponse) SetTemporary(v bool) {
	o.Temporary = &v
}


// GetApproximateMemberCount returns the ApproximateMemberCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetApproximateMemberCount() int32 {
	if o == nil || IsNil(o.ApproximateMemberCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ApproximateMemberCount.Get()
}

// GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetApproximateMemberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproximateMemberCount.Get()) {
		return nil, false
	}
	return o.ApproximateMemberCount.Get(), o.ApproximateMemberCount.IsSet()
}

// HasApproximateMemberCount returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasApproximateMemberCount() bool {
	if o != nil && o.ApproximateMemberCount.IsSet() {
		return true
	}

	return false
}

// SetApproximateMemberCount gets a reference to the given NullableInt32 and assigns it to the ApproximateMemberCount field.
func (o *GuildInviteResponse) SetApproximateMemberCount(v int32) {
	o.ApproximateMemberCount.Set(&v)
}

// SetApproximateMemberCountNil sets the value for ApproximateMemberCount to be an explicit nil
func (o *GuildInviteResponse) SetApproximateMemberCountNil() {
	o.ApproximateMemberCount.Set(nil)
}

// UnsetApproximateMemberCount ensures that no value is present for ApproximateMemberCount, not even an explicit nil
func (o *GuildInviteResponse) UnsetApproximateMemberCount() {
	o.ApproximateMemberCount.Unset()
}

// GetApproximatePresenceCount returns the ApproximatePresenceCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetApproximatePresenceCount() int32 {
	if o == nil || IsNil(o.ApproximatePresenceCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ApproximatePresenceCount.Get()
}

// GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetApproximatePresenceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproximatePresenceCount.Get()) {
		return nil, false
	}
	return o.ApproximatePresenceCount.Get(), o.ApproximatePresenceCount.IsSet()
}

// HasApproximatePresenceCount returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasApproximatePresenceCount() bool {
	if o != nil && o.ApproximatePresenceCount.IsSet() {
		return true
	}

	return false
}

// SetApproximatePresenceCount gets a reference to the given NullableInt32 and assigns it to the ApproximatePresenceCount field.
func (o *GuildInviteResponse) SetApproximatePresenceCount(v int32) {
	o.ApproximatePresenceCount.Set(&v)
}

// SetApproximatePresenceCountNil sets the value for ApproximatePresenceCount to be an explicit nil
func (o *GuildInviteResponse) SetApproximatePresenceCountNil() {
	o.ApproximatePresenceCount.Set(nil)
}

// UnsetApproximatePresenceCount ensures that no value is present for ApproximatePresenceCount, not even an explicit nil
func (o *GuildInviteResponse) UnsetApproximatePresenceCount() {
	o.ApproximatePresenceCount.Unset()
}

// GetIsNicknameChangeable returns the IsNicknameChangeable field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetIsNicknameChangeable() bool {
	if o == nil || IsNil(o.IsNicknameChangeable) {
		var ret bool
		return ret
	}
	return *o.IsNicknameChangeable
}

// GetIsNicknameChangeableOk returns a tuple with the IsNicknameChangeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetIsNicknameChangeableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNicknameChangeable) {
		return nil, false
	}
	return o.IsNicknameChangeable, true
}

// HasIsNicknameChangeable returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasIsNicknameChangeable() bool {
	if o != nil && !IsNil(o.IsNicknameChangeable) {
		return true
	}

	return false
}

// SetIsNicknameChangeable gets a reference to the given bool and assigns it to the IsNicknameChangeable field.
func (o *GuildInviteResponse) SetIsNicknameChangeable(v bool) {
	o.IsNicknameChangeable = &v
}


func (o GuildInviteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildInviteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type.Get()
	toSerialize["code"] = o.Code
	if !IsNil(o.Inviter) {
		toSerialize["inviter"] = o.Inviter
	}
	if !IsNil(o.MaxAge) {
		toSerialize["max_age"] = o.MaxAge
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if !IsNil(o.IsContact) {
		toSerialize["is_contact"] = o.IsContact
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	toSerialize["guild"] = o.Guild
	toSerialize["guild_id"] = o.GuildId
	toSerialize["channel"] = o.Channel
	if o.TargetType.IsSet() {
		toSerialize["target_type"] = o.TargetType.Get()
	}
	if !IsNil(o.TargetUser) {
		toSerialize["target_user"] = o.TargetUser
	}
	if !IsNil(o.TargetApplication) {
		toSerialize["target_application"] = o.TargetApplication
	}
	if !IsNil(o.GuildScheduledEvent) {
		toSerialize["guild_scheduled_event"] = o.GuildScheduledEvent
	}
	if !IsNil(o.Uses) {
		toSerialize["uses"] = o.Uses
	}
	if !IsNil(o.MaxUses) {
		toSerialize["max_uses"] = o.MaxUses
	}
	if !IsNil(o.Temporary) {
		toSerialize["temporary"] = o.Temporary
	}
	if o.ApproximateMemberCount.IsSet() {
		toSerialize["approximate_member_count"] = o.ApproximateMemberCount.Get()
	}
	if o.ApproximatePresenceCount.IsSet() {
		toSerialize["approximate_presence_count"] = o.ApproximatePresenceCount.Get()
	}
	if !IsNil(o.IsNicknameChangeable) {
		toSerialize["is_nickname_changeable"] = o.IsNicknameChangeable
	}
	return toSerialize, nil
}

func (o *GuildInviteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"code",
		"guild",
		"guild_id",
		"channel",
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

	varGuildInviteResponse := _GuildInviteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildInviteResponse)

	if err != nil {
		return err
	}

	*o = GuildInviteResponse(varGuildInviteResponse)

	return err
}

type NullableGuildInviteResponse struct {
	value *GuildInviteResponse
	isSet bool
}

func (v NullableGuildInviteResponse) Get() *GuildInviteResponse {
	return v.value
}

func (v *NullableGuildInviteResponse) Set(val *GuildInviteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildInviteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildInviteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildInviteResponse(val *GuildInviteResponse) *NullableGuildInviteResponse {
	return &NullableGuildInviteResponse{value: val, isSet: true}
}

func (v NullableGuildInviteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildInviteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

