/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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
	"bytes"
	"fmt"
)

// checks if the MessageRoleSubscriptionDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageRoleSubscriptionDataResponse{}

// MessageRoleSubscriptionDataResponse struct for MessageRoleSubscriptionDataResponse
type MessageRoleSubscriptionDataResponse struct {
	RoleSubscriptionListingId string `json:"role_subscription_listing_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	TierName string `json:"tier_name"`
	TotalMonthsSubscribed int32 `json:"total_months_subscribed"`
	IsRenewal bool `json:"is_renewal"`
}

type _MessageRoleSubscriptionDataResponse MessageRoleSubscriptionDataResponse

// NewMessageRoleSubscriptionDataResponse instantiates a new MessageRoleSubscriptionDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageRoleSubscriptionDataResponse(roleSubscriptionListingId string, tierName string, totalMonthsSubscribed int32, isRenewal bool) *MessageRoleSubscriptionDataResponse {
	this := MessageRoleSubscriptionDataResponse{}
	this.RoleSubscriptionListingId = roleSubscriptionListingId
	this.TierName = tierName
	this.TotalMonthsSubscribed = totalMonthsSubscribed
	this.IsRenewal = isRenewal
	return &this
}

// NewMessageRoleSubscriptionDataResponseWithDefaults instantiates a new MessageRoleSubscriptionDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageRoleSubscriptionDataResponseWithDefaults() *MessageRoleSubscriptionDataResponse {
	this := MessageRoleSubscriptionDataResponse{}
	return &this
}

// GetRoleSubscriptionListingId returns the RoleSubscriptionListingId field value
func (o *MessageRoleSubscriptionDataResponse) GetRoleSubscriptionListingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleSubscriptionListingId
}

// GetRoleSubscriptionListingIdOk returns a tuple with the RoleSubscriptionListingId field value
// and a boolean to check if the value has been set.
func (o *MessageRoleSubscriptionDataResponse) GetRoleSubscriptionListingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleSubscriptionListingId, true
}

// SetRoleSubscriptionListingId sets field value
func (o *MessageRoleSubscriptionDataResponse) SetRoleSubscriptionListingId(v string) {
	o.RoleSubscriptionListingId = v
}

// GetTierName returns the TierName field value
func (o *MessageRoleSubscriptionDataResponse) GetTierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TierName
}

// GetTierNameOk returns a tuple with the TierName field value
// and a boolean to check if the value has been set.
func (o *MessageRoleSubscriptionDataResponse) GetTierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TierName, true
}

// SetTierName sets field value
func (o *MessageRoleSubscriptionDataResponse) SetTierName(v string) {
	o.TierName = v
}

// GetTotalMonthsSubscribed returns the TotalMonthsSubscribed field value
func (o *MessageRoleSubscriptionDataResponse) GetTotalMonthsSubscribed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalMonthsSubscribed
}

// GetTotalMonthsSubscribedOk returns a tuple with the TotalMonthsSubscribed field value
// and a boolean to check if the value has been set.
func (o *MessageRoleSubscriptionDataResponse) GetTotalMonthsSubscribedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMonthsSubscribed, true
}

// SetTotalMonthsSubscribed sets field value
func (o *MessageRoleSubscriptionDataResponse) SetTotalMonthsSubscribed(v int32) {
	o.TotalMonthsSubscribed = v
}

// GetIsRenewal returns the IsRenewal field value
func (o *MessageRoleSubscriptionDataResponse) GetIsRenewal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRenewal
}

// GetIsRenewalOk returns a tuple with the IsRenewal field value
// and a boolean to check if the value has been set.
func (o *MessageRoleSubscriptionDataResponse) GetIsRenewalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRenewal, true
}

// SetIsRenewal sets field value
func (o *MessageRoleSubscriptionDataResponse) SetIsRenewal(v bool) {
	o.IsRenewal = v
}

func (o MessageRoleSubscriptionDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageRoleSubscriptionDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role_subscription_listing_id"] = o.RoleSubscriptionListingId
	toSerialize["tier_name"] = o.TierName
	toSerialize["total_months_subscribed"] = o.TotalMonthsSubscribed
	toSerialize["is_renewal"] = o.IsRenewal
	return toSerialize, nil
}

func (o *MessageRoleSubscriptionDataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role_subscription_listing_id",
		"tier_name",
		"total_months_subscribed",
		"is_renewal",
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

	varMessageRoleSubscriptionDataResponse := _MessageRoleSubscriptionDataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageRoleSubscriptionDataResponse)

	if err != nil {
		return err
	}

	*o = MessageRoleSubscriptionDataResponse(varMessageRoleSubscriptionDataResponse)

	return err
}

type NullableMessageRoleSubscriptionDataResponse struct {
	value *MessageRoleSubscriptionDataResponse
	isSet bool
}

func (v NullableMessageRoleSubscriptionDataResponse) Get() *MessageRoleSubscriptionDataResponse {
	return v.value
}

func (v *NullableMessageRoleSubscriptionDataResponse) Set(val *MessageRoleSubscriptionDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageRoleSubscriptionDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageRoleSubscriptionDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageRoleSubscriptionDataResponse(val *MessageRoleSubscriptionDataResponse) *NullableMessageRoleSubscriptionDataResponse {
	return &NullableMessageRoleSubscriptionDataResponse{value: val, isSet: true}
}

func (v NullableMessageRoleSubscriptionDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageRoleSubscriptionDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


