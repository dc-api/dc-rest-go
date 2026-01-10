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

// checks if the StageScheduledEventCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StageScheduledEventCreateRequest{}

// StageScheduledEventCreateRequest struct for StageScheduledEventCreateRequest
type StageScheduledEventCreateRequest struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Image *string `json:"image,omitempty"`
	ScheduledStartTime time.Time `json:"scheduled_start_time"`
	ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
	PrivacyLevel GuildScheduledEventPrivacyLevels `json:"privacy_level"`
	EntityType GuildScheduledEventEntityTypes `json:"entity_type"`
	ChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"channel_id,omitempty"`
	EntityMetadata *StageScheduledEventCreateRequestEntityMetadata `json:"entity_metadata,omitempty"`
}

type _StageScheduledEventCreateRequest StageScheduledEventCreateRequest

// NewStageScheduledEventCreateRequest instantiates a new StageScheduledEventCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStageScheduledEventCreateRequest(name string, scheduledStartTime time.Time, privacyLevel GuildScheduledEventPrivacyLevels, entityType GuildScheduledEventEntityTypes) *StageScheduledEventCreateRequest {
	this := StageScheduledEventCreateRequest{}
	this.Name = name
	this.ScheduledStartTime = scheduledStartTime
	this.PrivacyLevel = privacyLevel
	this.EntityType = entityType
	return &this
}

// NewStageScheduledEventCreateRequestWithDefaults instantiates a new StageScheduledEventCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStageScheduledEventCreateRequestWithDefaults() *StageScheduledEventCreateRequest {
	this := StageScheduledEventCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *StageScheduledEventCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StageScheduledEventCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StageScheduledEventCreateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StageScheduledEventCreateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageScheduledEventCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StageScheduledEventCreateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StageScheduledEventCreateRequest) SetDescription(v string) {
	o.Description = &v
}


// GetImage returns the Image field value if set, zero value otherwise.
func (o *StageScheduledEventCreateRequest) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageScheduledEventCreateRequest) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *StageScheduledEventCreateRequest) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *StageScheduledEventCreateRequest) SetImage(v string) {
	o.Image = &v
}


// GetScheduledStartTime returns the ScheduledStartTime field value
func (o *StageScheduledEventCreateRequest) GetScheduledStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ScheduledStartTime
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value
// and a boolean to check if the value has been set.
func (o *StageScheduledEventCreateRequest) GetScheduledStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledStartTime, true
}

// SetScheduledStartTime sets field value
func (o *StageScheduledEventCreateRequest) SetScheduledStartTime(v time.Time) {
	o.ScheduledStartTime = v
}

// GetScheduledEndTime returns the ScheduledEndTime field value if set, zero value otherwise.
func (o *StageScheduledEventCreateRequest) GetScheduledEndTime() time.Time {
	if o == nil || IsNil(o.ScheduledEndTime) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndTime
}

// GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageScheduledEventCreateRequest) GetScheduledEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledEndTime) {
		return nil, false
	}
	return o.ScheduledEndTime, true
}

// HasScheduledEndTime returns a boolean if a field has been set.
func (o *StageScheduledEventCreateRequest) HasScheduledEndTime() bool {
	if o != nil && !IsNil(o.ScheduledEndTime) {
		return true
	}

	return false
}

// SetScheduledEndTime gets a reference to the given time.Time and assigns it to the ScheduledEndTime field.
func (o *StageScheduledEventCreateRequest) SetScheduledEndTime(v time.Time) {
	o.ScheduledEndTime = &v
}


// GetPrivacyLevel returns the PrivacyLevel field value
func (o *StageScheduledEventCreateRequest) GetPrivacyLevel() GuildScheduledEventPrivacyLevels {
	if o == nil {
		var ret GuildScheduledEventPrivacyLevels
		return ret
	}

	return o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value
// and a boolean to check if the value has been set.
func (o *StageScheduledEventCreateRequest) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyLevel, true
}

// SetPrivacyLevel sets field value
func (o *StageScheduledEventCreateRequest) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels) {
	o.PrivacyLevel = v
}

// GetEntityType returns the EntityType field value
func (o *StageScheduledEventCreateRequest) GetEntityType() GuildScheduledEventEntityTypes {
	if o == nil {
		var ret GuildScheduledEventEntityTypes
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *StageScheduledEventCreateRequest) GetEntityTypeOk() (*GuildScheduledEventEntityTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *StageScheduledEventCreateRequest) SetEntityType(v GuildScheduledEventEntityTypes) {
	o.EntityType = v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *StageScheduledEventCreateRequest) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageScheduledEventCreateRequest) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *StageScheduledEventCreateRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ChannelId field.
func (o *StageScheduledEventCreateRequest) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ChannelId = &v
}


// GetEntityMetadata returns the EntityMetadata field value if set, zero value otherwise.
func (o *StageScheduledEventCreateRequest) GetEntityMetadata() StageScheduledEventCreateRequestEntityMetadata {
	if o == nil || IsNil(o.EntityMetadata) {
		var ret StageScheduledEventCreateRequestEntityMetadata
		return ret
	}
	return *o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageScheduledEventCreateRequest) GetEntityMetadataOk() (*StageScheduledEventCreateRequestEntityMetadata, bool) {
	if o == nil || IsNil(o.EntityMetadata) {
		return nil, false
	}
	return o.EntityMetadata, true
}

// HasEntityMetadata returns a boolean if a field has been set.
func (o *StageScheduledEventCreateRequest) HasEntityMetadata() bool {
	if o != nil && !IsNil(o.EntityMetadata) {
		return true
	}

	return false
}

// SetEntityMetadata gets a reference to the given StageScheduledEventCreateRequestEntityMetadata and assigns it to the EntityMetadata field.
func (o *StageScheduledEventCreateRequest) SetEntityMetadata(v StageScheduledEventCreateRequestEntityMetadata) {
	o.EntityMetadata = &v
}


func (o StageScheduledEventCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StageScheduledEventCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	toSerialize["scheduled_start_time"] = o.ScheduledStartTime
	if !IsNil(o.ScheduledEndTime) {
		toSerialize["scheduled_end_time"] = o.ScheduledEndTime
	}
	toSerialize["privacy_level"] = o.PrivacyLevel
	toSerialize["entity_type"] = o.EntityType
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.EntityMetadata) {
		toSerialize["entity_metadata"] = o.EntityMetadata
	}
	return toSerialize, nil
}

func (o *StageScheduledEventCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"scheduled_start_time",
		"privacy_level",
		"entity_type",
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

	varStageScheduledEventCreateRequest := _StageScheduledEventCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStageScheduledEventCreateRequest)

	if err != nil {
		return err
	}

	*o = StageScheduledEventCreateRequest(varStageScheduledEventCreateRequest)

	return err
}

type NullableStageScheduledEventCreateRequest struct {
	value *StageScheduledEventCreateRequest
	isSet bool
}

func (v NullableStageScheduledEventCreateRequest) Get() *StageScheduledEventCreateRequest {
	return v.value
}

func (v *NullableStageScheduledEventCreateRequest) Set(val *StageScheduledEventCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStageScheduledEventCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStageScheduledEventCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStageScheduledEventCreateRequest(val *StageScheduledEventCreateRequest) *NullableStageScheduledEventCreateRequest {
	return &NullableStageScheduledEventCreateRequest{value: val, isSet: true}
}

func (v NullableStageScheduledEventCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStageScheduledEventCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

