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

// checks if the StageScheduledEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StageScheduledEventResponse{}

// StageScheduledEventResponse struct for StageScheduledEventResponse
type StageScheduledEventResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	CreatorId *string `json:"creator_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Creator NullableUserResponse `json:"creator,omitempty"`
	Image NullableString `json:"image,omitempty"`
	ScheduledStartTime time.Time `json:"scheduled_start_time"`
	ScheduledEndTime NullableTime `json:"scheduled_end_time,omitempty"`
	Status NullableInt32 `json:"status"`
	EntityType NullableInt32 `json:"entity_type"`
	EntityId *string `json:"entity_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	UserCount NullableInt32 `json:"user_count,omitempty"`
	PrivacyLevel interface{} `json:"privacy_level"`
	UserRsvp NullableScheduledEventUserResponse `json:"user_rsvp,omitempty"`
	EntityMetadata map[string]interface{} `json:"entity_metadata,omitempty"`
}

type _StageScheduledEventResponse StageScheduledEventResponse

// NewStageScheduledEventResponse instantiates a new StageScheduledEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStageScheduledEventResponse(id string, guildId string, name string, scheduledStartTime time.Time, status NullableInt32, entityType NullableInt32, privacyLevel interface{}) *StageScheduledEventResponse {
	this := StageScheduledEventResponse{}
	this.Id = id
	this.GuildId = guildId
	this.Name = name
	this.ScheduledStartTime = scheduledStartTime
	this.Status = status
	this.EntityType = entityType
	this.PrivacyLevel = privacyLevel
	return &this
}

// NewStageScheduledEventResponseWithDefaults instantiates a new StageScheduledEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStageScheduledEventResponseWithDefaults() *StageScheduledEventResponse {
	this := StageScheduledEventResponse{}
	return &this
}

// GetId returns the Id field value
func (o *StageScheduledEventResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StageScheduledEventResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StageScheduledEventResponse) SetId(v string) {
	o.Id = v
}

// GetGuildId returns the GuildId field value
func (o *StageScheduledEventResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *StageScheduledEventResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *StageScheduledEventResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetName returns the Name field value
func (o *StageScheduledEventResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StageScheduledEventResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StageScheduledEventResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StageScheduledEventResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageScheduledEventResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StageScheduledEventResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StageScheduledEventResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StageScheduledEventResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *StageScheduledEventResponse) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageScheduledEventResponse) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *StageScheduledEventResponse) SetChannelId(v string) {
	o.ChannelId = &v
}


// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *StageScheduledEventResponse) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageScheduledEventResponse) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *StageScheduledEventResponse) SetCreatorId(v string) {
	o.CreatorId = &v
}


// GetCreator returns the Creator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StageScheduledEventResponse) GetCreator() UserResponse {
	if o == nil || IsNil(o.Creator.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.Creator.Get()
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageScheduledEventResponse) GetCreatorOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Creator.Get()) {
		return nil, false
	}
	return o.Creator.Get(), o.Creator.IsSet()
}

// HasCreator returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasCreator() bool {
	if o != nil && o.Creator.IsSet() {
		return true
	}

	return false
}

// SetCreator gets a reference to the given NullableUserResponse and assigns it to the Creator field.
func (o *StageScheduledEventResponse) SetCreator(v UserResponse) {
	o.Creator.Set(&v)
}

// SetCreatorNil sets the value for Creator to be an explicit nil
func (o *StageScheduledEventResponse) SetCreatorNil() {
	o.Creator.Set(nil)
}

// UnsetCreator ensures that no value is present for Creator, not even an explicit nil
func (o *StageScheduledEventResponse) UnsetCreator() {
	o.Creator.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StageScheduledEventResponse) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageScheduledEventResponse) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image.Get()) {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *StageScheduledEventResponse) SetImage(v string) {
	o.Image.Set(&v)
}

// SetImageNil sets the value for Image to be an explicit nil
func (o *StageScheduledEventResponse) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *StageScheduledEventResponse) UnsetImage() {
	o.Image.Unset()
}

// GetScheduledStartTime returns the ScheduledStartTime field value
func (o *StageScheduledEventResponse) GetScheduledStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ScheduledStartTime
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value
// and a boolean to check if the value has been set.
func (o *StageScheduledEventResponse) GetScheduledStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledStartTime, true
}

// SetScheduledStartTime sets field value
func (o *StageScheduledEventResponse) SetScheduledStartTime(v time.Time) {
	o.ScheduledStartTime = v
}

// GetScheduledEndTime returns the ScheduledEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StageScheduledEventResponse) GetScheduledEndTime() time.Time {
	if o == nil || IsNil(o.ScheduledEndTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndTime.Get()
}

// GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageScheduledEventResponse) GetScheduledEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledEndTime.Get()) {
		return nil, false
	}
	return o.ScheduledEndTime.Get(), o.ScheduledEndTime.IsSet()
}

// HasScheduledEndTime returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasScheduledEndTime() bool {
	if o != nil && o.ScheduledEndTime.IsSet() {
		return true
	}

	return false
}

// SetScheduledEndTime gets a reference to the given NullableTime and assigns it to the ScheduledEndTime field.
func (o *StageScheduledEventResponse) SetScheduledEndTime(v time.Time) {
	o.ScheduledEndTime.Set(&v)
}

// SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil
func (o *StageScheduledEventResponse) SetScheduledEndTimeNil() {
	o.ScheduledEndTime.Set(nil)
}

// UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
func (o *StageScheduledEventResponse) UnsetScheduledEndTime() {
	o.ScheduledEndTime.Unset()
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *StageScheduledEventResponse) GetStatus() int32 {
	if o == nil || o.Status.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageScheduledEventResponse) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *StageScheduledEventResponse) SetStatus(v int32) {
	o.Status.Set(&v)
}

// GetEntityType returns the EntityType field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *StageScheduledEventResponse) GetEntityType() int32 {
	if o == nil || o.EntityType.Get() == nil {
		var ret int32
		return ret
	}

	return *o.EntityType.Get()
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageScheduledEventResponse) GetEntityTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityType.Get(), o.EntityType.IsSet()
}

// SetEntityType sets field value
func (o *StageScheduledEventResponse) SetEntityType(v int32) {
	o.EntityType.Set(&v)
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *StageScheduledEventResponse) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageScheduledEventResponse) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *StageScheduledEventResponse) SetEntityId(v string) {
	o.EntityId = &v
}


// GetUserCount returns the UserCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StageScheduledEventResponse) GetUserCount() int32 {
	if o == nil || IsNil(o.UserCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UserCount.Get()
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageScheduledEventResponse) GetUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UserCount.Get()) {
		return nil, false
	}
	return o.UserCount.Get(), o.UserCount.IsSet()
}

// HasUserCount returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasUserCount() bool {
	if o != nil && o.UserCount.IsSet() {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given NullableInt32 and assigns it to the UserCount field.
func (o *StageScheduledEventResponse) SetUserCount(v int32) {
	o.UserCount.Set(&v)
}

// SetUserCountNil sets the value for UserCount to be an explicit nil
func (o *StageScheduledEventResponse) SetUserCountNil() {
	o.UserCount.Set(nil)
}

// UnsetUserCount ensures that no value is present for UserCount, not even an explicit nil
func (o *StageScheduledEventResponse) UnsetUserCount() {
	o.UserCount.Unset()
}

// GetPrivacyLevel returns the PrivacyLevel field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *StageScheduledEventResponse) GetPrivacyLevel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageScheduledEventResponse) GetPrivacyLevelOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PrivacyLevel) {
		return nil, false
	}
	return &o.PrivacyLevel, true
}

// SetPrivacyLevel sets field value
func (o *StageScheduledEventResponse) SetPrivacyLevel(v interface{}) {
	o.PrivacyLevel = v
}

// GetUserRsvp returns the UserRsvp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StageScheduledEventResponse) GetUserRsvp() ScheduledEventUserResponse {
	if o == nil || IsNil(o.UserRsvp.Get()) {
		var ret ScheduledEventUserResponse
		return ret
	}
	return *o.UserRsvp.Get()
}

// GetUserRsvpOk returns a tuple with the UserRsvp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageScheduledEventResponse) GetUserRsvpOk() (*ScheduledEventUserResponse, bool) {
	if o == nil || IsNil(o.UserRsvp.Get()) {
		return nil, false
	}
	return o.UserRsvp.Get(), o.UserRsvp.IsSet()
}

// HasUserRsvp returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasUserRsvp() bool {
	if o != nil && o.UserRsvp.IsSet() {
		return true
	}

	return false
}

// SetUserRsvp gets a reference to the given NullableScheduledEventUserResponse and assigns it to the UserRsvp field.
func (o *StageScheduledEventResponse) SetUserRsvp(v ScheduledEventUserResponse) {
	o.UserRsvp.Set(&v)
}

// SetUserRsvpNil sets the value for UserRsvp to be an explicit nil
func (o *StageScheduledEventResponse) SetUserRsvpNil() {
	o.UserRsvp.Set(nil)
}

// UnsetUserRsvp ensures that no value is present for UserRsvp, not even an explicit nil
func (o *StageScheduledEventResponse) UnsetUserRsvp() {
	o.UserRsvp.Unset()
}

// GetEntityMetadata returns the EntityMetadata field value if set, zero value otherwise.
func (o *StageScheduledEventResponse) GetEntityMetadata() map[string]interface{} {
	if o == nil || IsNil(o.EntityMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageScheduledEventResponse) GetEntityMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EntityMetadata) {
		return map[string]interface{}{}, false
	}
	return o.EntityMetadata, true
}

// HasEntityMetadata returns a boolean if a field has been set.
func (o *StageScheduledEventResponse) HasEntityMetadata() bool {
	if o != nil && !IsNil(o.EntityMetadata) {
		return true
	}

	return false
}

// SetEntityMetadata gets a reference to the given map[string]interface{} and assigns it to the EntityMetadata field.
func (o *StageScheduledEventResponse) SetEntityMetadata(v map[string]interface{}) {
	o.EntityMetadata = v
}


func (o StageScheduledEventResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StageScheduledEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["guild_id"] = o.GuildId
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creator_id"] = o.CreatorId
	}
	if o.Creator.IsSet() {
		toSerialize["creator"] = o.Creator.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	toSerialize["scheduled_start_time"] = o.ScheduledStartTime
	if o.ScheduledEndTime.IsSet() {
		toSerialize["scheduled_end_time"] = o.ScheduledEndTime.Get()
	}
	toSerialize["status"] = o.Status.Get()
	toSerialize["entity_type"] = o.EntityType.Get()
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if o.UserCount.IsSet() {
		toSerialize["user_count"] = o.UserCount.Get()
	}
	if o.PrivacyLevel != nil {
		toSerialize["privacy_level"] = o.PrivacyLevel
	}
	if o.UserRsvp.IsSet() {
		toSerialize["user_rsvp"] = o.UserRsvp.Get()
	}
	if !IsNil(o.EntityMetadata) {
		toSerialize["entity_metadata"] = o.EntityMetadata
	}
	return toSerialize, nil
}

func (o *StageScheduledEventResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"guild_id",
		"name",
		"scheduled_start_time",
		"status",
		"entity_type",
		"privacy_level",
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

	varStageScheduledEventResponse := _StageScheduledEventResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStageScheduledEventResponse)

	if err != nil {
		return err
	}

	*o = StageScheduledEventResponse(varStageScheduledEventResponse)

	return err
}

type NullableStageScheduledEventResponse struct {
	value *StageScheduledEventResponse
	isSet bool
}

func (v NullableStageScheduledEventResponse) Get() *StageScheduledEventResponse {
	return v.value
}

func (v *NullableStageScheduledEventResponse) Set(val *StageScheduledEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStageScheduledEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStageScheduledEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStageScheduledEventResponse(val *StageScheduledEventResponse) *NullableStageScheduledEventResponse {
	return &NullableStageScheduledEventResponse{value: val, isSet: true}
}

func (v NullableStageScheduledEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStageScheduledEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


