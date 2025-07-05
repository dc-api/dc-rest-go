/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
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

// checks if the GuildInviteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildInviteResponse{}

// GuildInviteResponse struct for GuildInviteResponse
type GuildInviteResponse struct {
	Type NullableInt32 `json:"type,omitempty"`
	Code string `json:"code"`
	Inviter NullableUserResponse `json:"inviter,omitempty"`
	MaxAge NullableInt32 `json:"max_age,omitempty"`
	CreatedAt NullableTime `json:"created_at,omitempty"`
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	IsContact NullableBool `json:"is_contact,omitempty"`
	Flags NullableInt32 `json:"flags,omitempty"`
	Guild NullableInviteGuildResponse `json:"guild,omitempty"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Channel NullableInviteChannelResponse `json:"channel,omitempty"`
	TargetType NullableInt32 `json:"target_type,omitempty"`
	TargetUser NullableUserResponse `json:"target_user,omitempty"`
	TargetApplication NullableInviteApplicationResponse `json:"target_application,omitempty"`
	GuildScheduledEvent NullableScheduledEventResponse `json:"guild_scheduled_event,omitempty"`
	Uses NullableInt32 `json:"uses,omitempty"`
	MaxUses NullableInt32 `json:"max_uses,omitempty"`
	Temporary NullableBool `json:"temporary,omitempty"`
	ApproximateMemberCount NullableInt32 `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount NullableInt32 `json:"approximate_presence_count,omitempty"`
	IsNicknameChangeable NullableBool `json:"is_nickname_changeable,omitempty"`
}

type _GuildInviteResponse GuildInviteResponse

// NewGuildInviteResponse instantiates a new GuildInviteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildInviteResponse(code string) *GuildInviteResponse {
	this := GuildInviteResponse{}
	this.Code = code
	return &this
}

// NewGuildInviteResponseWithDefaults instantiates a new GuildInviteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildInviteResponseWithDefaults() *GuildInviteResponse {
	this := GuildInviteResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *GuildInviteResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *GuildInviteResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *GuildInviteResponse) UnsetType() {
	o.Type.Unset()
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

// GetInviter returns the Inviter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetInviter() UserResponse {
	if o == nil || IsNil(o.Inviter.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.Inviter.Get()
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetInviterOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Inviter.Get()) {
		return nil, false
	}
	return o.Inviter.Get(), o.Inviter.IsSet()
}

// HasInviter returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasInviter() bool {
	if o != nil && o.Inviter.IsSet() {
		return true
	}

	return false
}

// SetInviter gets a reference to the given NullableUserResponse and assigns it to the Inviter field.
func (o *GuildInviteResponse) SetInviter(v UserResponse) {
	o.Inviter.Set(&v)
}

// SetInviterNil sets the value for Inviter to be an explicit nil
func (o *GuildInviteResponse) SetInviterNil() {
	o.Inviter.Set(nil)
}

// UnsetInviter ensures that no value is present for Inviter, not even an explicit nil
func (o *GuildInviteResponse) UnsetInviter() {
	o.Inviter.Unset()
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAge.Get()
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge.Get()) {
		return nil, false
	}
	return o.MaxAge.Get(), o.MaxAge.IsSet()
}

// HasMaxAge returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasMaxAge() bool {
	if o != nil && o.MaxAge.IsSet() {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given NullableInt32 and assigns it to the MaxAge field.
func (o *GuildInviteResponse) SetMaxAge(v int32) {
	o.MaxAge.Set(&v)
}

// SetMaxAgeNil sets the value for MaxAge to be an explicit nil
func (o *GuildInviteResponse) SetMaxAgeNil() {
	o.MaxAge.Set(nil)
}

// UnsetMaxAge ensures that no value is present for MaxAge, not even an explicit nil
func (o *GuildInviteResponse) UnsetMaxAge() {
	o.MaxAge.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *GuildInviteResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GuildInviteResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GuildInviteResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
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

// GetIsContact returns the IsContact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetIsContact() bool {
	if o == nil || IsNil(o.IsContact.Get()) {
		var ret bool
		return ret
	}
	return *o.IsContact.Get()
}

// GetIsContactOk returns a tuple with the IsContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetIsContactOk() (*bool, bool) {
	if o == nil || IsNil(o.IsContact.Get()) {
		return nil, false
	}
	return o.IsContact.Get(), o.IsContact.IsSet()
}

// HasIsContact returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasIsContact() bool {
	if o != nil && o.IsContact.IsSet() {
		return true
	}

	return false
}

// SetIsContact gets a reference to the given NullableBool and assigns it to the IsContact field.
func (o *GuildInviteResponse) SetIsContact(v bool) {
	o.IsContact.Set(&v)
}

// SetIsContactNil sets the value for IsContact to be an explicit nil
func (o *GuildInviteResponse) SetIsContactNil() {
	o.IsContact.Set(nil)
}

// UnsetIsContact ensures that no value is present for IsContact, not even an explicit nil
func (o *GuildInviteResponse) UnsetIsContact() {
	o.IsContact.Unset()
}

// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags.Get()) {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *GuildInviteResponse) SetFlags(v int32) {
	o.Flags.Set(&v)
}

// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *GuildInviteResponse) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *GuildInviteResponse) UnsetFlags() {
	o.Flags.Unset()
}

// GetGuild returns the Guild field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetGuild() InviteGuildResponse {
	if o == nil || IsNil(o.Guild.Get()) {
		var ret InviteGuildResponse
		return ret
	}
	return *o.Guild.Get()
}

// GetGuildOk returns a tuple with the Guild field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetGuildOk() (*InviteGuildResponse, bool) {
	if o == nil || IsNil(o.Guild.Get()) {
		return nil, false
	}
	return o.Guild.Get(), o.Guild.IsSet()
}

// HasGuild returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasGuild() bool {
	if o != nil && o.Guild.IsSet() {
		return true
	}

	return false
}

// SetGuild gets a reference to the given NullableInviteGuildResponse and assigns it to the Guild field.
func (o *GuildInviteResponse) SetGuild(v InviteGuildResponse) {
	o.Guild.Set(&v)
}

// SetGuildNil sets the value for Guild to be an explicit nil
func (o *GuildInviteResponse) SetGuildNil() {
	o.Guild.Set(nil)
}

// UnsetGuild ensures that no value is present for Guild, not even an explicit nil
func (o *GuildInviteResponse) UnsetGuild() {
	o.Guild.Unset()
}

// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *GuildInviteResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildInviteResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *GuildInviteResponse) SetGuildId(v string) {
	o.GuildId = &v
}


// GetChannel returns the Channel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetChannel() InviteChannelResponse {
	if o == nil || IsNil(o.Channel.Get()) {
		var ret InviteChannelResponse
		return ret
	}
	return *o.Channel.Get()
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetChannelOk() (*InviteChannelResponse, bool) {
	if o == nil || IsNil(o.Channel.Get()) {
		return nil, false
	}
	return o.Channel.Get(), o.Channel.IsSet()
}

// HasChannel returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasChannel() bool {
	if o != nil && o.Channel.IsSet() {
		return true
	}

	return false
}

// SetChannel gets a reference to the given NullableInviteChannelResponse and assigns it to the Channel field.
func (o *GuildInviteResponse) SetChannel(v InviteChannelResponse) {
	o.Channel.Set(&v)
}

// SetChannelNil sets the value for Channel to be an explicit nil
func (o *GuildInviteResponse) SetChannelNil() {
	o.Channel.Set(nil)
}

// UnsetChannel ensures that no value is present for Channel, not even an explicit nil
func (o *GuildInviteResponse) UnsetChannel() {
	o.Channel.Unset()
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

// GetTargetUser returns the TargetUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetTargetUser() UserResponse {
	if o == nil || IsNil(o.TargetUser.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.TargetUser.Get()
}

// GetTargetUserOk returns a tuple with the TargetUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetTargetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.TargetUser.Get()) {
		return nil, false
	}
	return o.TargetUser.Get(), o.TargetUser.IsSet()
}

// HasTargetUser returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasTargetUser() bool {
	if o != nil && o.TargetUser.IsSet() {
		return true
	}

	return false
}

// SetTargetUser gets a reference to the given NullableUserResponse and assigns it to the TargetUser field.
func (o *GuildInviteResponse) SetTargetUser(v UserResponse) {
	o.TargetUser.Set(&v)
}

// SetTargetUserNil sets the value for TargetUser to be an explicit nil
func (o *GuildInviteResponse) SetTargetUserNil() {
	o.TargetUser.Set(nil)
}

// UnsetTargetUser ensures that no value is present for TargetUser, not even an explicit nil
func (o *GuildInviteResponse) UnsetTargetUser() {
	o.TargetUser.Unset()
}

// GetTargetApplication returns the TargetApplication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetTargetApplication() InviteApplicationResponse {
	if o == nil || IsNil(o.TargetApplication.Get()) {
		var ret InviteApplicationResponse
		return ret
	}
	return *o.TargetApplication.Get()
}

// GetTargetApplicationOk returns a tuple with the TargetApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetTargetApplicationOk() (*InviteApplicationResponse, bool) {
	if o == nil || IsNil(o.TargetApplication.Get()) {
		return nil, false
	}
	return o.TargetApplication.Get(), o.TargetApplication.IsSet()
}

// HasTargetApplication returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasTargetApplication() bool {
	if o != nil && o.TargetApplication.IsSet() {
		return true
	}

	return false
}

// SetTargetApplication gets a reference to the given NullableInviteApplicationResponse and assigns it to the TargetApplication field.
func (o *GuildInviteResponse) SetTargetApplication(v InviteApplicationResponse) {
	o.TargetApplication.Set(&v)
}

// SetTargetApplicationNil sets the value for TargetApplication to be an explicit nil
func (o *GuildInviteResponse) SetTargetApplicationNil() {
	o.TargetApplication.Set(nil)
}

// UnsetTargetApplication ensures that no value is present for TargetApplication, not even an explicit nil
func (o *GuildInviteResponse) UnsetTargetApplication() {
	o.TargetApplication.Unset()
}

// GetGuildScheduledEvent returns the GuildScheduledEvent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetGuildScheduledEvent() ScheduledEventResponse {
	if o == nil || IsNil(o.GuildScheduledEvent.Get()) {
		var ret ScheduledEventResponse
		return ret
	}
	return *o.GuildScheduledEvent.Get()
}

// GetGuildScheduledEventOk returns a tuple with the GuildScheduledEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetGuildScheduledEventOk() (*ScheduledEventResponse, bool) {
	if o == nil || IsNil(o.GuildScheduledEvent.Get()) {
		return nil, false
	}
	return o.GuildScheduledEvent.Get(), o.GuildScheduledEvent.IsSet()
}

// HasGuildScheduledEvent returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasGuildScheduledEvent() bool {
	if o != nil && o.GuildScheduledEvent.IsSet() {
		return true
	}

	return false
}

// SetGuildScheduledEvent gets a reference to the given NullableScheduledEventResponse and assigns it to the GuildScheduledEvent field.
func (o *GuildInviteResponse) SetGuildScheduledEvent(v ScheduledEventResponse) {
	o.GuildScheduledEvent.Set(&v)
}

// SetGuildScheduledEventNil sets the value for GuildScheduledEvent to be an explicit nil
func (o *GuildInviteResponse) SetGuildScheduledEventNil() {
	o.GuildScheduledEvent.Set(nil)
}

// UnsetGuildScheduledEvent ensures that no value is present for GuildScheduledEvent, not even an explicit nil
func (o *GuildInviteResponse) UnsetGuildScheduledEvent() {
	o.GuildScheduledEvent.Unset()
}

// GetUses returns the Uses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetUses() int32 {
	if o == nil || IsNil(o.Uses.Get()) {
		var ret int32
		return ret
	}
	return *o.Uses.Get()
}

// GetUsesOk returns a tuple with the Uses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.Uses.Get()) {
		return nil, false
	}
	return o.Uses.Get(), o.Uses.IsSet()
}

// HasUses returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasUses() bool {
	if o != nil && o.Uses.IsSet() {
		return true
	}

	return false
}

// SetUses gets a reference to the given NullableInt32 and assigns it to the Uses field.
func (o *GuildInviteResponse) SetUses(v int32) {
	o.Uses.Set(&v)
}

// SetUsesNil sets the value for Uses to be an explicit nil
func (o *GuildInviteResponse) SetUsesNil() {
	o.Uses.Set(nil)
}

// UnsetUses ensures that no value is present for Uses, not even an explicit nil
func (o *GuildInviteResponse) UnsetUses() {
	o.Uses.Unset()
}

// GetMaxUses returns the MaxUses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetMaxUses() int32 {
	if o == nil || IsNil(o.MaxUses.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxUses.Get()
}

// GetMaxUsesOk returns a tuple with the MaxUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetMaxUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUses.Get()) {
		return nil, false
	}
	return o.MaxUses.Get(), o.MaxUses.IsSet()
}

// HasMaxUses returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasMaxUses() bool {
	if o != nil && o.MaxUses.IsSet() {
		return true
	}

	return false
}

// SetMaxUses gets a reference to the given NullableInt32 and assigns it to the MaxUses field.
func (o *GuildInviteResponse) SetMaxUses(v int32) {
	o.MaxUses.Set(&v)
}

// SetMaxUsesNil sets the value for MaxUses to be an explicit nil
func (o *GuildInviteResponse) SetMaxUsesNil() {
	o.MaxUses.Set(nil)
}

// UnsetMaxUses ensures that no value is present for MaxUses, not even an explicit nil
func (o *GuildInviteResponse) UnsetMaxUses() {
	o.MaxUses.Unset()
}

// GetTemporary returns the Temporary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetTemporary() bool {
	if o == nil || IsNil(o.Temporary.Get()) {
		var ret bool
		return ret
	}
	return *o.Temporary.Get()
}

// GetTemporaryOk returns a tuple with the Temporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetTemporaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Temporary.Get()) {
		return nil, false
	}
	return o.Temporary.Get(), o.Temporary.IsSet()
}

// HasTemporary returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasTemporary() bool {
	if o != nil && o.Temporary.IsSet() {
		return true
	}

	return false
}

// SetTemporary gets a reference to the given NullableBool and assigns it to the Temporary field.
func (o *GuildInviteResponse) SetTemporary(v bool) {
	o.Temporary.Set(&v)
}

// SetTemporaryNil sets the value for Temporary to be an explicit nil
func (o *GuildInviteResponse) SetTemporaryNil() {
	o.Temporary.Set(nil)
}

// UnsetTemporary ensures that no value is present for Temporary, not even an explicit nil
func (o *GuildInviteResponse) UnsetTemporary() {
	o.Temporary.Unset()
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

// GetIsNicknameChangeable returns the IsNicknameChangeable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildInviteResponse) GetIsNicknameChangeable() bool {
	if o == nil || IsNil(o.IsNicknameChangeable.Get()) {
		var ret bool
		return ret
	}
	return *o.IsNicknameChangeable.Get()
}

// GetIsNicknameChangeableOk returns a tuple with the IsNicknameChangeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildInviteResponse) GetIsNicknameChangeableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNicknameChangeable.Get()) {
		return nil, false
	}
	return o.IsNicknameChangeable.Get(), o.IsNicknameChangeable.IsSet()
}

// HasIsNicknameChangeable returns a boolean if a field has been set.
func (o *GuildInviteResponse) HasIsNicknameChangeable() bool {
	if o != nil && o.IsNicknameChangeable.IsSet() {
		return true
	}

	return false
}

// SetIsNicknameChangeable gets a reference to the given NullableBool and assigns it to the IsNicknameChangeable field.
func (o *GuildInviteResponse) SetIsNicknameChangeable(v bool) {
	o.IsNicknameChangeable.Set(&v)
}

// SetIsNicknameChangeableNil sets the value for IsNicknameChangeable to be an explicit nil
func (o *GuildInviteResponse) SetIsNicknameChangeableNil() {
	o.IsNicknameChangeable.Set(nil)
}

// UnsetIsNicknameChangeable ensures that no value is present for IsNicknameChangeable, not even an explicit nil
func (o *GuildInviteResponse) UnsetIsNicknameChangeable() {
	o.IsNicknameChangeable.Unset()
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
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	toSerialize["code"] = o.Code
	if o.Inviter.IsSet() {
		toSerialize["inviter"] = o.Inviter.Get()
	}
	if o.MaxAge.IsSet() {
		toSerialize["max_age"] = o.MaxAge.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.IsContact.IsSet() {
		toSerialize["is_contact"] = o.IsContact.Get()
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	if o.Guild.IsSet() {
		toSerialize["guild"] = o.Guild.Get()
	}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	if o.Channel.IsSet() {
		toSerialize["channel"] = o.Channel.Get()
	}
	if o.TargetType.IsSet() {
		toSerialize["target_type"] = o.TargetType.Get()
	}
	if o.TargetUser.IsSet() {
		toSerialize["target_user"] = o.TargetUser.Get()
	}
	if o.TargetApplication.IsSet() {
		toSerialize["target_application"] = o.TargetApplication.Get()
	}
	if o.GuildScheduledEvent.IsSet() {
		toSerialize["guild_scheduled_event"] = o.GuildScheduledEvent.Get()
	}
	if o.Uses.IsSet() {
		toSerialize["uses"] = o.Uses.Get()
	}
	if o.MaxUses.IsSet() {
		toSerialize["max_uses"] = o.MaxUses.Get()
	}
	if o.Temporary.IsSet() {
		toSerialize["temporary"] = o.Temporary.Get()
	}
	if o.ApproximateMemberCount.IsSet() {
		toSerialize["approximate_member_count"] = o.ApproximateMemberCount.Get()
	}
	if o.ApproximatePresenceCount.IsSet() {
		toSerialize["approximate_presence_count"] = o.ApproximatePresenceCount.Get()
	}
	if o.IsNicknameChangeable.IsSet() {
		toSerialize["is_nickname_changeable"] = o.IsNicknameChangeable.Get()
	}
	return toSerialize, nil
}

func (o *GuildInviteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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


