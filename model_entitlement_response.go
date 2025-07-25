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
	"time"
	"bytes"
	"fmt"
)

// checks if the EntitlementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementResponse{}

// EntitlementResponse struct for EntitlementResponse
type EntitlementResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SkuId string `json:"sku_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ApplicationId string `json:"application_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	UserId string `json:"user_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Deleted bool `json:"deleted"`
	StartsAt NullableTime `json:"starts_at,omitempty"`
	EndsAt NullableTime `json:"ends_at,omitempty"`
	Type NullableInt32 `json:"type"`
	FulfilledAt NullableTime `json:"fulfilled_at,omitempty"`
	FulfillmentStatus NullableInt32 `json:"fulfillment_status,omitempty"`
	Consumed NullableBool `json:"consumed,omitempty"`
}

type _EntitlementResponse EntitlementResponse

// NewEntitlementResponse instantiates a new EntitlementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementResponse(id string, skuId string, applicationId string, userId string, deleted bool, type_ NullableInt32) *EntitlementResponse {
	this := EntitlementResponse{}
	this.Id = id
	this.SkuId = skuId
	this.ApplicationId = applicationId
	this.UserId = userId
	this.Deleted = deleted
	this.Type = type_
	return &this
}

// NewEntitlementResponseWithDefaults instantiates a new EntitlementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementResponseWithDefaults() *EntitlementResponse {
	this := EntitlementResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementResponse) SetId(v string) {
	o.Id = v
}

// GetSkuId returns the SkuId field value
func (o *EntitlementResponse) GetSkuId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value
// and a boolean to check if the value has been set.
func (o *EntitlementResponse) GetSkuIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkuId, true
}

// SetSkuId sets field value
func (o *EntitlementResponse) SetSkuId(v string) {
	o.SkuId = v
}

// GetApplicationId returns the ApplicationId field value
func (o *EntitlementResponse) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *EntitlementResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *EntitlementResponse) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetUserId returns the UserId field value
func (o *EntitlementResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *EntitlementResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *EntitlementResponse) SetUserId(v string) {
	o.UserId = v
}

// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *EntitlementResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *EntitlementResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *EntitlementResponse) SetGuildId(v string) {
	o.GuildId = &v
}


// GetDeleted returns the Deleted field value
func (o *EntitlementResponse) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *EntitlementResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *EntitlementResponse) SetDeleted(v bool) {
	o.Deleted = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementResponse) GetStartsAt() time.Time {
	if o == nil || IsNil(o.StartsAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartsAt.Get()
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementResponse) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartsAt.Get()) {
		return nil, false
	}
	return o.StartsAt.Get(), o.StartsAt.IsSet()
}

// HasStartsAt returns a boolean if a field has been set.
func (o *EntitlementResponse) HasStartsAt() bool {
	if o != nil && o.StartsAt.IsSet() {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given NullableTime and assigns it to the StartsAt field.
func (o *EntitlementResponse) SetStartsAt(v time.Time) {
	o.StartsAt.Set(&v)
}

// SetStartsAtNil sets the value for StartsAt to be an explicit nil
func (o *EntitlementResponse) SetStartsAtNil() {
	o.StartsAt.Set(nil)
}

// UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
func (o *EntitlementResponse) UnsetStartsAt() {
	o.StartsAt.Unset()
}

// GetEndsAt returns the EndsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementResponse) GetEndsAt() time.Time {
	if o == nil || IsNil(o.EndsAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndsAt.Get()
}

// GetEndsAtOk returns a tuple with the EndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementResponse) GetEndsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndsAt.Get()) {
		return nil, false
	}
	return o.EndsAt.Get(), o.EndsAt.IsSet()
}

// HasEndsAt returns a boolean if a field has been set.
func (o *EntitlementResponse) HasEndsAt() bool {
	if o != nil && o.EndsAt.IsSet() {
		return true
	}

	return false
}

// SetEndsAt gets a reference to the given NullableTime and assigns it to the EndsAt field.
func (o *EntitlementResponse) SetEndsAt(v time.Time) {
	o.EndsAt.Set(&v)
}

// SetEndsAtNil sets the value for EndsAt to be an explicit nil
func (o *EntitlementResponse) SetEndsAtNil() {
	o.EndsAt.Set(nil)
}

// UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
func (o *EntitlementResponse) UnsetEndsAt() {
	o.EndsAt.Unset()
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *EntitlementResponse) GetType() int32 {
	if o == nil || o.Type.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *EntitlementResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// GetFulfilledAt returns the FulfilledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementResponse) GetFulfilledAt() time.Time {
	if o == nil || IsNil(o.FulfilledAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.FulfilledAt.Get()
}

// GetFulfilledAtOk returns a tuple with the FulfilledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementResponse) GetFulfilledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FulfilledAt.Get()) {
		return nil, false
	}
	return o.FulfilledAt.Get(), o.FulfilledAt.IsSet()
}

// HasFulfilledAt returns a boolean if a field has been set.
func (o *EntitlementResponse) HasFulfilledAt() bool {
	if o != nil && o.FulfilledAt.IsSet() {
		return true
	}

	return false
}

// SetFulfilledAt gets a reference to the given NullableTime and assigns it to the FulfilledAt field.
func (o *EntitlementResponse) SetFulfilledAt(v time.Time) {
	o.FulfilledAt.Set(&v)
}

// SetFulfilledAtNil sets the value for FulfilledAt to be an explicit nil
func (o *EntitlementResponse) SetFulfilledAtNil() {
	o.FulfilledAt.Set(nil)
}

// UnsetFulfilledAt ensures that no value is present for FulfilledAt, not even an explicit nil
func (o *EntitlementResponse) UnsetFulfilledAt() {
	o.FulfilledAt.Unset()
}

// GetFulfillmentStatus returns the FulfillmentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementResponse) GetFulfillmentStatus() int32 {
	if o == nil || IsNil(o.FulfillmentStatus.Get()) {
		var ret int32
		return ret
	}
	return *o.FulfillmentStatus.Get()
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementResponse) GetFulfillmentStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.FulfillmentStatus.Get()) {
		return nil, false
	}
	return o.FulfillmentStatus.Get(), o.FulfillmentStatus.IsSet()
}

// HasFulfillmentStatus returns a boolean if a field has been set.
func (o *EntitlementResponse) HasFulfillmentStatus() bool {
	if o != nil && o.FulfillmentStatus.IsSet() {
		return true
	}

	return false
}

// SetFulfillmentStatus gets a reference to the given NullableInt32 and assigns it to the FulfillmentStatus field.
func (o *EntitlementResponse) SetFulfillmentStatus(v int32) {
	o.FulfillmentStatus.Set(&v)
}

// SetFulfillmentStatusNil sets the value for FulfillmentStatus to be an explicit nil
func (o *EntitlementResponse) SetFulfillmentStatusNil() {
	o.FulfillmentStatus.Set(nil)
}

// UnsetFulfillmentStatus ensures that no value is present for FulfillmentStatus, not even an explicit nil
func (o *EntitlementResponse) UnsetFulfillmentStatus() {
	o.FulfillmentStatus.Unset()
}

// GetConsumed returns the Consumed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementResponse) GetConsumed() bool {
	if o == nil || IsNil(o.Consumed.Get()) {
		var ret bool
		return ret
	}
	return *o.Consumed.Get()
}

// GetConsumedOk returns a tuple with the Consumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementResponse) GetConsumedOk() (*bool, bool) {
	if o == nil || IsNil(o.Consumed.Get()) {
		return nil, false
	}
	return o.Consumed.Get(), o.Consumed.IsSet()
}

// HasConsumed returns a boolean if a field has been set.
func (o *EntitlementResponse) HasConsumed() bool {
	if o != nil && o.Consumed.IsSet() {
		return true
	}

	return false
}

// SetConsumed gets a reference to the given NullableBool and assigns it to the Consumed field.
func (o *EntitlementResponse) SetConsumed(v bool) {
	o.Consumed.Set(&v)
}

// SetConsumedNil sets the value for Consumed to be an explicit nil
func (o *EntitlementResponse) SetConsumedNil() {
	o.Consumed.Set(nil)
}

// UnsetConsumed ensures that no value is present for Consumed, not even an explicit nil
func (o *EntitlementResponse) UnsetConsumed() {
	o.Consumed.Unset()
}

func (o EntitlementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["sku_id"] = o.SkuId
	toSerialize["application_id"] = o.ApplicationId
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	toSerialize["deleted"] = o.Deleted
	if o.StartsAt.IsSet() {
		toSerialize["starts_at"] = o.StartsAt.Get()
	}
	if o.EndsAt.IsSet() {
		toSerialize["ends_at"] = o.EndsAt.Get()
	}
	toSerialize["type"] = o.Type.Get()
	if o.FulfilledAt.IsSet() {
		toSerialize["fulfilled_at"] = o.FulfilledAt.Get()
	}
	if o.FulfillmentStatus.IsSet() {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus.Get()
	}
	if o.Consumed.IsSet() {
		toSerialize["consumed"] = o.Consumed.Get()
	}
	return toSerialize, nil
}

func (o *EntitlementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"sku_id",
		"application_id",
		"user_id",
		"deleted",
		"type",
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

	varEntitlementResponse := _EntitlementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntitlementResponse)

	if err != nil {
		return err
	}

	*o = EntitlementResponse(varEntitlementResponse)

	return err
}

type NullableEntitlementResponse struct {
	value *EntitlementResponse
	isSet bool
}

func (v NullableEntitlementResponse) Get() *EntitlementResponse {
	return v.value
}

func (v *NullableEntitlementResponse) Set(val *EntitlementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementResponse(val *EntitlementResponse) *NullableEntitlementResponse {
	return &NullableEntitlementResponse{value: val, isSet: true}
}

func (v NullableEntitlementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


