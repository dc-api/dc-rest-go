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
)

// checks if the ExternalScheduledEventPatchRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalScheduledEventPatchRequestPartial{}

// ExternalScheduledEventPatchRequestPartial struct for ExternalScheduledEventPatchRequestPartial
type ExternalScheduledEventPatchRequestPartial struct {
	Status *ExternalScheduledEventPatchRequestPartialStatus `json:"status,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Image *string `json:"image,omitempty"`
	ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
	EntityType *ExternalScheduledEventPatchRequestPartialEntityType `json:"entity_type,omitempty"`
	PrivacyLevel *GuildScheduledEventPrivacyLevels `json:"privacy_level,omitempty"`
	ChannelId *GetEntitlementsSkuIdsParameterOneOfInner `json:"channel_id,omitempty"`
	EntityMetadata *EntityMetadataExternal `json:"entity_metadata,omitempty"`
}

// NewExternalScheduledEventPatchRequestPartial instantiates a new ExternalScheduledEventPatchRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalScheduledEventPatchRequestPartial() *ExternalScheduledEventPatchRequestPartial {
	this := ExternalScheduledEventPatchRequestPartial{}
	return &this
}

// NewExternalScheduledEventPatchRequestPartialWithDefaults instantiates a new ExternalScheduledEventPatchRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalScheduledEventPatchRequestPartialWithDefaults() *ExternalScheduledEventPatchRequestPartial {
	this := ExternalScheduledEventPatchRequestPartial{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetStatus() ExternalScheduledEventPatchRequestPartialStatus {
	if o == nil || IsNil(o.Status) {
		var ret ExternalScheduledEventPatchRequestPartialStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetStatusOk() (*ExternalScheduledEventPatchRequestPartialStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ExternalScheduledEventPatchRequestPartialStatus and assigns it to the Status field.
func (o *ExternalScheduledEventPatchRequestPartial) SetStatus(v ExternalScheduledEventPatchRequestPartialStatus) {
	o.Status = &v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalScheduledEventPatchRequestPartial) SetName(v string) {
	o.Name = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExternalScheduledEventPatchRequestPartial) SetDescription(v string) {
	o.Description = &v
}


// GetImage returns the Image field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ExternalScheduledEventPatchRequestPartial) SetImage(v string) {
	o.Image = &v
}


// GetScheduledStartTime returns the ScheduledStartTime field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetScheduledStartTime() time.Time {
	if o == nil || IsNil(o.ScheduledStartTime) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledStartTime
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetScheduledStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledStartTime) {
		return nil, false
	}
	return o.ScheduledStartTime, true
}

// HasScheduledStartTime returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasScheduledStartTime() bool {
	if o != nil && !IsNil(o.ScheduledStartTime) {
		return true
	}

	return false
}

// SetScheduledStartTime gets a reference to the given time.Time and assigns it to the ScheduledStartTime field.
func (o *ExternalScheduledEventPatchRequestPartial) SetScheduledStartTime(v time.Time) {
	o.ScheduledStartTime = &v
}


// GetScheduledEndTime returns the ScheduledEndTime field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetScheduledEndTime() time.Time {
	if o == nil || IsNil(o.ScheduledEndTime) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndTime
}

// GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetScheduledEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledEndTime) {
		return nil, false
	}
	return o.ScheduledEndTime, true
}

// HasScheduledEndTime returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasScheduledEndTime() bool {
	if o != nil && !IsNil(o.ScheduledEndTime) {
		return true
	}

	return false
}

// SetScheduledEndTime gets a reference to the given time.Time and assigns it to the ScheduledEndTime field.
func (o *ExternalScheduledEventPatchRequestPartial) SetScheduledEndTime(v time.Time) {
	o.ScheduledEndTime = &v
}


// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetEntityType() ExternalScheduledEventPatchRequestPartialEntityType {
	if o == nil || IsNil(o.EntityType) {
		var ret ExternalScheduledEventPatchRequestPartialEntityType
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetEntityTypeOk() (*ExternalScheduledEventPatchRequestPartialEntityType, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given ExternalScheduledEventPatchRequestPartialEntityType and assigns it to the EntityType field.
func (o *ExternalScheduledEventPatchRequestPartial) SetEntityType(v ExternalScheduledEventPatchRequestPartialEntityType) {
	o.EntityType = &v
}


// GetPrivacyLevel returns the PrivacyLevel field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetPrivacyLevel() GuildScheduledEventPrivacyLevels {
	if o == nil || IsNil(o.PrivacyLevel) {
		var ret GuildScheduledEventPrivacyLevels
		return ret
	}
	return *o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool) {
	if o == nil || IsNil(o.PrivacyLevel) {
		return nil, false
	}
	return o.PrivacyLevel, true
}

// HasPrivacyLevel returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasPrivacyLevel() bool {
	if o != nil && !IsNil(o.PrivacyLevel) {
		return true
	}

	return false
}

// SetPrivacyLevel gets a reference to the given GuildScheduledEventPrivacyLevels and assigns it to the PrivacyLevel field.
func (o *ExternalScheduledEventPatchRequestPartial) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels) {
	o.PrivacyLevel = &v
}


// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ChannelId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ChannelId field.
func (o *ExternalScheduledEventPatchRequestPartial) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ChannelId = &v
}


// GetEntityMetadata returns the EntityMetadata field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetEntityMetadata() EntityMetadataExternal {
	if o == nil || IsNil(o.EntityMetadata) {
		var ret EntityMetadataExternal
		return ret
	}
	return *o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetEntityMetadataOk() (*EntityMetadataExternal, bool) {
	if o == nil || IsNil(o.EntityMetadata) {
		return nil, false
	}
	return o.EntityMetadata, true
}

// HasEntityMetadata returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasEntityMetadata() bool {
	if o != nil && !IsNil(o.EntityMetadata) {
		return true
	}

	return false
}

// SetEntityMetadata gets a reference to the given EntityMetadataExternal and assigns it to the EntityMetadata field.
func (o *ExternalScheduledEventPatchRequestPartial) SetEntityMetadata(v EntityMetadataExternal) {
	o.EntityMetadata = &v
}


func (o ExternalScheduledEventPatchRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalScheduledEventPatchRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.ScheduledStartTime) {
		toSerialize["scheduled_start_time"] = o.ScheduledStartTime
	}
	if !IsNil(o.ScheduledEndTime) {
		toSerialize["scheduled_end_time"] = o.ScheduledEndTime
	}
	if !IsNil(o.EntityType) {
		toSerialize["entity_type"] = o.EntityType
	}
	if !IsNil(o.PrivacyLevel) {
		toSerialize["privacy_level"] = o.PrivacyLevel
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.EntityMetadata) {
		toSerialize["entity_metadata"] = o.EntityMetadata
	}
	return toSerialize, nil
}

type NullableExternalScheduledEventPatchRequestPartial struct {
	value *ExternalScheduledEventPatchRequestPartial
	isSet bool
}

func (v NullableExternalScheduledEventPatchRequestPartial) Get() *ExternalScheduledEventPatchRequestPartial {
	return v.value
}

func (v *NullableExternalScheduledEventPatchRequestPartial) Set(val *ExternalScheduledEventPatchRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalScheduledEventPatchRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalScheduledEventPatchRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalScheduledEventPatchRequestPartial(val *ExternalScheduledEventPatchRequestPartial) *NullableExternalScheduledEventPatchRequestPartial {
	return &NullableExternalScheduledEventPatchRequestPartial{value: val, isSet: true}
}

func (v NullableExternalScheduledEventPatchRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalScheduledEventPatchRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

