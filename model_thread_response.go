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
	"time"
	"bytes"
	"fmt"
)

// checks if the ThreadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadResponse{}

// ThreadResponse struct for ThreadResponse
type ThreadResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type ChannelTypes `json:"type"`
	LastMessageId *GetEntitlementsSkuIdsParameterOneOfInner `json:"last_message_id,omitempty"`
	Flags int32 `json:"flags"`
	LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	ParentId *GetEntitlementsSkuIdsParameterOneOfInner `json:"parent_id,omitempty"`
	RateLimitPerUser *int32 `json:"rate_limit_per_user,omitempty"`
	Bitrate *int32 `json:"bitrate,omitempty"`
	UserLimit *int32 `json:"user_limit,omitempty"`
	RtcRegion *string `json:"rtc_region,omitempty"`
	VideoQualityMode *VideoQualityModes `json:"video_quality_mode,omitempty"`
	Permissions *string `json:"permissions,omitempty"`
	OwnerId string `json:"owner_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ThreadMetadata ThreadMetadataResponse `json:"thread_metadata"`
	MessageCount int32 `json:"message_count"`
	MemberCount int32 `json:"member_count"`
	TotalMessageSent int32 `json:"total_message_sent"`
	AppliedTags []string `json:"applied_tags,omitempty"`
	Member *ThreadMemberResponse `json:"member,omitempty"`
}

type _ThreadResponse ThreadResponse

// NewThreadResponse instantiates a new ThreadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadResponse(id string, type_ ChannelTypes, flags int32, guildId string, name string, ownerId string, threadMetadata ThreadMetadataResponse, messageCount int32, memberCount int32, totalMessageSent int32) *ThreadResponse {
	this := ThreadResponse{}
	this.Id = id
	this.Type = type_
	this.Flags = flags
	this.GuildId = guildId
	this.Name = name
	this.OwnerId = ownerId
	this.ThreadMetadata = threadMetadata
	this.MessageCount = messageCount
	this.MemberCount = memberCount
	this.TotalMessageSent = totalMessageSent
	return &this
}

// NewThreadResponseWithDefaults instantiates a new ThreadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadResponseWithDefaults() *ThreadResponse {
	this := ThreadResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ThreadResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThreadResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ThreadResponse) GetType() ChannelTypes {
	if o == nil {
		var ret ChannelTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetTypeOk() (*ChannelTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ThreadResponse) SetType(v ChannelTypes) {
	o.Type = v
}

// GetLastMessageId returns the LastMessageId field value if set, zero value otherwise.
func (o *ThreadResponse) GetLastMessageId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.LastMessageId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.LastMessageId
}

// GetLastMessageIdOk returns a tuple with the LastMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetLastMessageIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.LastMessageId) {
		return nil, false
	}
	return o.LastMessageId, true
}

// HasLastMessageId returns a boolean if a field has been set.
func (o *ThreadResponse) HasLastMessageId() bool {
	if o != nil && !IsNil(o.LastMessageId) {
		return true
	}

	return false
}

// SetLastMessageId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the LastMessageId field.
func (o *ThreadResponse) SetLastMessageId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.LastMessageId = &v
}


// GetFlags returns the Flags field value
func (o *ThreadResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *ThreadResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetLastPinTimestamp returns the LastPinTimestamp field value if set, zero value otherwise.
func (o *ThreadResponse) GetLastPinTimestamp() time.Time {
	if o == nil || IsNil(o.LastPinTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastPinTimestamp
}

// GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetLastPinTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPinTimestamp) {
		return nil, false
	}
	return o.LastPinTimestamp, true
}

// HasLastPinTimestamp returns a boolean if a field has been set.
func (o *ThreadResponse) HasLastPinTimestamp() bool {
	if o != nil && !IsNil(o.LastPinTimestamp) {
		return true
	}

	return false
}

// SetLastPinTimestamp gets a reference to the given time.Time and assigns it to the LastPinTimestamp field.
func (o *ThreadResponse) SetLastPinTimestamp(v time.Time) {
	o.LastPinTimestamp = &v
}


// GetGuildId returns the GuildId field value
func (o *ThreadResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *ThreadResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetName returns the Name field value
func (o *ThreadResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ThreadResponse) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ThreadResponse) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ParentId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ThreadResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ParentId field.
func (o *ThreadResponse) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ParentId = &v
}


// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise.
func (o *ThreadResponse) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimitPerUser) {
		return nil, false
	}
	return o.RateLimitPerUser, true
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *ThreadResponse) HasRateLimitPerUser() bool {
	if o != nil && !IsNil(o.RateLimitPerUser) {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given int32 and assigns it to the RateLimitPerUser field.
func (o *ThreadResponse) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser = &v
}


// GetBitrate returns the Bitrate field value if set, zero value otherwise.
func (o *ThreadResponse) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate) {
		var ret int32
		return ret
	}
	return *o.Bitrate
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.Bitrate) {
		return nil, false
	}
	return o.Bitrate, true
}

// HasBitrate returns a boolean if a field has been set.
func (o *ThreadResponse) HasBitrate() bool {
	if o != nil && !IsNil(o.Bitrate) {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given int32 and assigns it to the Bitrate field.
func (o *ThreadResponse) SetBitrate(v int32) {
	o.Bitrate = &v
}


// GetUserLimit returns the UserLimit field value if set, zero value otherwise.
func (o *ThreadResponse) GetUserLimit() int32 {
	if o == nil || IsNil(o.UserLimit) {
		var ret int32
		return ret
	}
	return *o.UserLimit
}

// GetUserLimitOk returns a tuple with the UserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetUserLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UserLimit) {
		return nil, false
	}
	return o.UserLimit, true
}

// HasUserLimit returns a boolean if a field has been set.
func (o *ThreadResponse) HasUserLimit() bool {
	if o != nil && !IsNil(o.UserLimit) {
		return true
	}

	return false
}

// SetUserLimit gets a reference to the given int32 and assigns it to the UserLimit field.
func (o *ThreadResponse) SetUserLimit(v int32) {
	o.UserLimit = &v
}


// GetRtcRegion returns the RtcRegion field value if set, zero value otherwise.
func (o *ThreadResponse) GetRtcRegion() string {
	if o == nil || IsNil(o.RtcRegion) {
		var ret string
		return ret
	}
	return *o.RtcRegion
}

// GetRtcRegionOk returns a tuple with the RtcRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetRtcRegionOk() (*string, bool) {
	if o == nil || IsNil(o.RtcRegion) {
		return nil, false
	}
	return o.RtcRegion, true
}

// HasRtcRegion returns a boolean if a field has been set.
func (o *ThreadResponse) HasRtcRegion() bool {
	if o != nil && !IsNil(o.RtcRegion) {
		return true
	}

	return false
}

// SetRtcRegion gets a reference to the given string and assigns it to the RtcRegion field.
func (o *ThreadResponse) SetRtcRegion(v string) {
	o.RtcRegion = &v
}


// GetVideoQualityMode returns the VideoQualityMode field value if set, zero value otherwise.
func (o *ThreadResponse) GetVideoQualityMode() VideoQualityModes {
	if o == nil || IsNil(o.VideoQualityMode) {
		var ret VideoQualityModes
		return ret
	}
	return *o.VideoQualityMode
}

// GetVideoQualityModeOk returns a tuple with the VideoQualityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetVideoQualityModeOk() (*VideoQualityModes, bool) {
	if o == nil || IsNil(o.VideoQualityMode) {
		return nil, false
	}
	return o.VideoQualityMode, true
}

// HasVideoQualityMode returns a boolean if a field has been set.
func (o *ThreadResponse) HasVideoQualityMode() bool {
	if o != nil && !IsNil(o.VideoQualityMode) {
		return true
	}

	return false
}

// SetVideoQualityMode gets a reference to the given VideoQualityModes and assigns it to the VideoQualityMode field.
func (o *ThreadResponse) SetVideoQualityMode(v VideoQualityModes) {
	o.VideoQualityMode = &v
}


// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ThreadResponse) GetPermissions() string {
	if o == nil || IsNil(o.Permissions) {
		var ret string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetPermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ThreadResponse) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given string and assigns it to the Permissions field.
func (o *ThreadResponse) SetPermissions(v string) {
	o.Permissions = &v
}


// GetOwnerId returns the OwnerId field value
func (o *ThreadResponse) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *ThreadResponse) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetThreadMetadata returns the ThreadMetadata field value
func (o *ThreadResponse) GetThreadMetadata() ThreadMetadataResponse {
	if o == nil {
		var ret ThreadMetadataResponse
		return ret
	}

	return o.ThreadMetadata
}

// GetThreadMetadataOk returns a tuple with the ThreadMetadata field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetThreadMetadataOk() (*ThreadMetadataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadMetadata, true
}

// SetThreadMetadata sets field value
func (o *ThreadResponse) SetThreadMetadata(v ThreadMetadataResponse) {
	o.ThreadMetadata = v
}

// GetMessageCount returns the MessageCount field value
func (o *ThreadResponse) GetMessageCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetMessageCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageCount, true
}

// SetMessageCount sets field value
func (o *ThreadResponse) SetMessageCount(v int32) {
	o.MessageCount = v
}

// GetMemberCount returns the MemberCount field value
func (o *ThreadResponse) GetMemberCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetMemberCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberCount, true
}

// SetMemberCount sets field value
func (o *ThreadResponse) SetMemberCount(v int32) {
	o.MemberCount = v
}

// GetTotalMessageSent returns the TotalMessageSent field value
func (o *ThreadResponse) GetTotalMessageSent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalMessageSent
}

// GetTotalMessageSentOk returns a tuple with the TotalMessageSent field value
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetTotalMessageSentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMessageSent, true
}

// SetTotalMessageSent sets field value
func (o *ThreadResponse) SetTotalMessageSent(v int32) {
	o.TotalMessageSent = v
}

// GetAppliedTags returns the AppliedTags field value if set, zero value otherwise.
func (o *ThreadResponse) GetAppliedTags() []string {
	if o == nil || IsNil(o.AppliedTags) {
		var ret []string
		return ret
	}
	return o.AppliedTags
}

// GetAppliedTagsOk returns a tuple with the AppliedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetAppliedTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppliedTags) {
		return nil, false
	}
	return o.AppliedTags, true
}

// HasAppliedTags returns a boolean if a field has been set.
func (o *ThreadResponse) HasAppliedTags() bool {
	if o != nil && !IsNil(o.AppliedTags) {
		return true
	}

	return false
}

// SetAppliedTags gets a reference to the given []string and assigns it to the AppliedTags field.
func (o *ThreadResponse) SetAppliedTags(v []string) {
	o.AppliedTags = v
}


// GetMember returns the Member field value if set, zero value otherwise.
func (o *ThreadResponse) GetMember() ThreadMemberResponse {
	if o == nil || IsNil(o.Member) {
		var ret ThreadMemberResponse
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadResponse) GetMemberOk() (*ThreadMemberResponse, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *ThreadResponse) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given ThreadMemberResponse and assigns it to the Member field.
func (o *ThreadResponse) SetMember(v ThreadMemberResponse) {
	o.Member = &v
}


func (o ThreadResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.LastMessageId) {
		toSerialize["last_message_id"] = o.LastMessageId
	}
	toSerialize["flags"] = o.Flags
	if !IsNil(o.LastPinTimestamp) {
		toSerialize["last_pin_timestamp"] = o.LastPinTimestamp
	}
	toSerialize["guild_id"] = o.GuildId
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.RateLimitPerUser) {
		toSerialize["rate_limit_per_user"] = o.RateLimitPerUser
	}
	if !IsNil(o.Bitrate) {
		toSerialize["bitrate"] = o.Bitrate
	}
	if !IsNil(o.UserLimit) {
		toSerialize["user_limit"] = o.UserLimit
	}
	if !IsNil(o.RtcRegion) {
		toSerialize["rtc_region"] = o.RtcRegion
	}
	if !IsNil(o.VideoQualityMode) {
		toSerialize["video_quality_mode"] = o.VideoQualityMode
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["thread_metadata"] = o.ThreadMetadata
	toSerialize["message_count"] = o.MessageCount
	toSerialize["member_count"] = o.MemberCount
	toSerialize["total_message_sent"] = o.TotalMessageSent
	if !IsNil(o.AppliedTags) {
		toSerialize["applied_tags"] = o.AppliedTags
	}
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	return toSerialize, nil
}

func (o *ThreadResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"flags",
		"guild_id",
		"name",
		"owner_id",
		"thread_metadata",
		"message_count",
		"member_count",
		"total_message_sent",
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

	varThreadResponse := _ThreadResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThreadResponse)

	if err != nil {
		return err
	}

	*o = ThreadResponse(varThreadResponse)

	return err
}

type NullableThreadResponse struct {
	value *ThreadResponse
	isSet bool
}

func (v NullableThreadResponse) Get() *ThreadResponse {
	return v.value
}

func (v *NullableThreadResponse) Set(val *ThreadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadResponse(val *ThreadResponse) *NullableThreadResponse {
	return &NullableThreadResponse{value: val, isSet: true}
}

func (v NullableThreadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

