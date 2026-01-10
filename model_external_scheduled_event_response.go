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
	"time"
	"bytes"
	"fmt"
)

// checks if the ExternalScheduledEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalScheduledEventResponse{}

// ExternalScheduledEventResponse struct for ExternalScheduledEventResponse
type ExternalScheduledEventResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	ChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"channel_id,omitempty"`
	CreatorId *GetEntitlementsSkuIdsParameterOneOfInner `json:"creator_id,omitempty"`
	Creator *UserResponse `json:"creator,omitempty"`
	Image *string `json:"image,omitempty"`
	ScheduledStartTime time.Time `json:"scheduled_start_time"`
	ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
	Status GuildScheduledEventStatuses `json:"status"`
	EntityType GuildScheduledEventEntityTypes `json:"entity_type"`
	EntityId *GetEntitlementsSkuIdsParameterOneOfInner `json:"entity_id,omitempty"`
	UserCount *int32 `json:"user_count,omitempty"`
	PrivacyLevel GuildScheduledEventPrivacyLevels `json:"privacy_level"`
	UserRsvp *ExternalScheduledEventResponseUserRsvp `json:"user_rsvp,omitempty"`
	EntityMetadata EntityMetadataExternalResponse `json:"entity_metadata"`
}

type _ExternalScheduledEventResponse ExternalScheduledEventResponse

// NewExternalScheduledEventResponse instantiates a new ExternalScheduledEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalScheduledEventResponse(id string, guildId string, name string, scheduledStartTime time.Time, status GuildScheduledEventStatuses, entityType GuildScheduledEventEntityTypes, privacyLevel GuildScheduledEventPrivacyLevels, entityMetadata EntityMetadataExternalResponse) *ExternalScheduledEventResponse {
	this := ExternalScheduledEventResponse{}
	this.Id = id
	this.GuildId = guildId
	this.Name = name
	this.ScheduledStartTime = scheduledStartTime
	this.Status = status
	this.EntityType = entityType
	this.PrivacyLevel = privacyLevel
	this.EntityMetadata = entityMetadata
	return &this
}

// NewExternalScheduledEventResponseWithDefaults instantiates a new ExternalScheduledEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalScheduledEventResponseWithDefaults() *ExternalScheduledEventResponse {
	this := ExternalScheduledEventResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalScheduledEventResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalScheduledEventResponse) SetId(v string) {
	o.Id = v
}

// GetGuildId returns the GuildId field value
func (o *ExternalScheduledEventResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *ExternalScheduledEventResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetName returns the Name field value
func (o *ExternalScheduledEventResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalScheduledEventResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExternalScheduledEventResponse) SetDescription(v string) {
	o.Description = &v
}


// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ChannelId field.
func (o *ExternalScheduledEventResponse) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ChannelId = &v
}


// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetCreatorId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.CreatorId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetCreatorIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the CreatorId field.
func (o *ExternalScheduledEventResponse) SetCreatorId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.CreatorId = &v
}


// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetCreator() UserResponse {
	if o == nil || IsNil(o.Creator) {
		var ret UserResponse
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetCreatorOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given UserResponse and assigns it to the Creator field.
func (o *ExternalScheduledEventResponse) SetCreator(v UserResponse) {
	o.Creator = &v
}


// GetImage returns the Image field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ExternalScheduledEventResponse) SetImage(v string) {
	o.Image = &v
}


// GetScheduledStartTime returns the ScheduledStartTime field value
func (o *ExternalScheduledEventResponse) GetScheduledStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ScheduledStartTime
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetScheduledStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledStartTime, true
}

// SetScheduledStartTime sets field value
func (o *ExternalScheduledEventResponse) SetScheduledStartTime(v time.Time) {
	o.ScheduledStartTime = v
}

// GetScheduledEndTime returns the ScheduledEndTime field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetScheduledEndTime() time.Time {
	if o == nil || IsNil(o.ScheduledEndTime) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndTime
}

// GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetScheduledEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledEndTime) {
		return nil, false
	}
	return o.ScheduledEndTime, true
}

// HasScheduledEndTime returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasScheduledEndTime() bool {
	if o != nil && !IsNil(o.ScheduledEndTime) {
		return true
	}

	return false
}

// SetScheduledEndTime gets a reference to the given time.Time and assigns it to the ScheduledEndTime field.
func (o *ExternalScheduledEventResponse) SetScheduledEndTime(v time.Time) {
	o.ScheduledEndTime = &v
}


// GetStatus returns the Status field value
func (o *ExternalScheduledEventResponse) GetStatus() GuildScheduledEventStatuses {
	if o == nil {
		var ret GuildScheduledEventStatuses
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetStatusOk() (*GuildScheduledEventStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExternalScheduledEventResponse) SetStatus(v GuildScheduledEventStatuses) {
	o.Status = v
}

// GetEntityType returns the EntityType field value
func (o *ExternalScheduledEventResponse) GetEntityType() GuildScheduledEventEntityTypes {
	if o == nil {
		var ret GuildScheduledEventEntityTypes
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetEntityTypeOk() (*GuildScheduledEventEntityTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *ExternalScheduledEventResponse) SetEntityType(v GuildScheduledEventEntityTypes) {
	o.EntityType = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetEntityId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.EntityId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetEntityIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the EntityId field.
func (o *ExternalScheduledEventResponse) SetEntityId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.EntityId = &v
}


// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetUserCount() int32 {
	if o == nil || IsNil(o.UserCount) {
		var ret int32
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UserCount) {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasUserCount() bool {
	if o != nil && !IsNil(o.UserCount) {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.
func (o *ExternalScheduledEventResponse) SetUserCount(v int32) {
	o.UserCount = &v
}


// GetPrivacyLevel returns the PrivacyLevel field value
func (o *ExternalScheduledEventResponse) GetPrivacyLevel() GuildScheduledEventPrivacyLevels {
	if o == nil {
		var ret GuildScheduledEventPrivacyLevels
		return ret
	}

	return o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyLevel, true
}

// SetPrivacyLevel sets field value
func (o *ExternalScheduledEventResponse) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels) {
	o.PrivacyLevel = v
}

// GetUserRsvp returns the UserRsvp field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetUserRsvp() ExternalScheduledEventResponseUserRsvp {
	if o == nil || IsNil(o.UserRsvp) {
		var ret ExternalScheduledEventResponseUserRsvp
		return ret
	}
	return *o.UserRsvp
}

// GetUserRsvpOk returns a tuple with the UserRsvp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetUserRsvpOk() (*ExternalScheduledEventResponseUserRsvp, bool) {
	if o == nil || IsNil(o.UserRsvp) {
		return nil, false
	}
	return o.UserRsvp, true
}

// HasUserRsvp returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasUserRsvp() bool {
	if o != nil && !IsNil(o.UserRsvp) {
		return true
	}

	return false
}

// SetUserRsvp gets a reference to the given ExternalScheduledEventResponseUserRsvp and assigns it to the UserRsvp field.
func (o *ExternalScheduledEventResponse) SetUserRsvp(v ExternalScheduledEventResponseUserRsvp) {
	o.UserRsvp = &v
}


// GetEntityMetadata returns the EntityMetadata field value
func (o *ExternalScheduledEventResponse) GetEntityMetadata() EntityMetadataExternalResponse {
	if o == nil {
		var ret EntityMetadataExternalResponse
		return ret
	}

	return o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetEntityMetadataOk() (*EntityMetadataExternalResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityMetadata, true
}

// SetEntityMetadata sets field value
func (o *ExternalScheduledEventResponse) SetEntityMetadata(v EntityMetadataExternalResponse) {
	o.EntityMetadata = v
}

func (o ExternalScheduledEventResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalScheduledEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["guild_id"] = o.GuildId
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creator_id"] = o.CreatorId
	}
	if !IsNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	toSerialize["scheduled_start_time"] = o.ScheduledStartTime
	if !IsNil(o.ScheduledEndTime) {
		toSerialize["scheduled_end_time"] = o.ScheduledEndTime
	}
	toSerialize["status"] = o.Status
	toSerialize["entity_type"] = o.EntityType
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.UserCount) {
		toSerialize["user_count"] = o.UserCount
	}
	toSerialize["privacy_level"] = o.PrivacyLevel
	if !IsNil(o.UserRsvp) {
		toSerialize["user_rsvp"] = o.UserRsvp
	}
	toSerialize["entity_metadata"] = o.EntityMetadata
	return toSerialize, nil
}

func (o *ExternalScheduledEventResponse) UnmarshalJSON(data []byte) (err error) {
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
		"entity_metadata",
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

	varExternalScheduledEventResponse := _ExternalScheduledEventResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalScheduledEventResponse)

	if err != nil {
		return err
	}

	*o = ExternalScheduledEventResponse(varExternalScheduledEventResponse)

	return err
}

type NullableExternalScheduledEventResponse struct {
	value *ExternalScheduledEventResponse
	isSet bool
}

func (v NullableExternalScheduledEventResponse) Get() *ExternalScheduledEventResponse {
	return v.value
}

func (v *NullableExternalScheduledEventResponse) Set(val *ExternalScheduledEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalScheduledEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalScheduledEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalScheduledEventResponse(val *ExternalScheduledEventResponse) *NullableExternalScheduledEventResponse {
	return &NullableExternalScheduledEventResponse{value: val, isSet: true}
}

func (v NullableExternalScheduledEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalScheduledEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

