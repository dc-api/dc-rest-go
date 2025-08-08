/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-08-08T14:09:23.736426080Z[Etc/UTC]
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
)

// checks if the GuildRoleTagsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildRoleTagsResponse{}

// GuildRoleTagsResponse struct for GuildRoleTagsResponse
type GuildRoleTagsResponse struct {
	PremiumSubscriber interface{} `json:"premium_subscriber,omitempty"`
	BotId *string `json:"bot_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	IntegrationId *string `json:"integration_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SubscriptionListingId *string `json:"subscription_listing_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	AvailableForPurchase interface{} `json:"available_for_purchase,omitempty"`
	GuildConnections interface{} `json:"guild_connections,omitempty"`
}

// NewGuildRoleTagsResponse instantiates a new GuildRoleTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildRoleTagsResponse() *GuildRoleTagsResponse {
	this := GuildRoleTagsResponse{}
	return &this
}

// NewGuildRoleTagsResponseWithDefaults instantiates a new GuildRoleTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildRoleTagsResponseWithDefaults() *GuildRoleTagsResponse {
	this := GuildRoleTagsResponse{}
	return &this
}

// GetPremiumSubscriber returns the PremiumSubscriber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleTagsResponse) GetPremiumSubscriber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PremiumSubscriber
}

// GetPremiumSubscriberOk returns a tuple with the PremiumSubscriber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleTagsResponse) GetPremiumSubscriberOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PremiumSubscriber, true
}

// HasPremiumSubscriber returns a boolean if a field has been set.
func (o *GuildRoleTagsResponse) HasPremiumSubscriber() bool {
	if o != nil && !IsNil(o.PremiumSubscriber) {
		return true
	}

	return false
}

// SetPremiumSubscriber gets a reference to the given interface{} and assigns it to the PremiumSubscriber field.
func (o *GuildRoleTagsResponse) SetPremiumSubscriber(v interface{}) {
	o.PremiumSubscriber = v
}


// GetBotId returns the BotId field value if set, zero value otherwise.
func (o *GuildRoleTagsResponse) GetBotId() string {
	if o == nil || IsNil(o.BotId) {
		var ret string
		return ret
	}
	return *o.BotId
}

// GetBotIdOk returns a tuple with the BotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildRoleTagsResponse) GetBotIdOk() (*string, bool) {
	if o == nil || IsNil(o.BotId) {
		return nil, false
	}
	return o.BotId, true
}

// HasBotId returns a boolean if a field has been set.
func (o *GuildRoleTagsResponse) HasBotId() bool {
	if o != nil && !IsNil(o.BotId) {
		return true
	}

	return false
}

// SetBotId gets a reference to the given string and assigns it to the BotId field.
func (o *GuildRoleTagsResponse) SetBotId(v string) {
	o.BotId = &v
}


// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *GuildRoleTagsResponse) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildRoleTagsResponse) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *GuildRoleTagsResponse) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *GuildRoleTagsResponse) SetIntegrationId(v string) {
	o.IntegrationId = &v
}


// GetSubscriptionListingId returns the SubscriptionListingId field value if set, zero value otherwise.
func (o *GuildRoleTagsResponse) GetSubscriptionListingId() string {
	if o == nil || IsNil(o.SubscriptionListingId) {
		var ret string
		return ret
	}
	return *o.SubscriptionListingId
}

// GetSubscriptionListingIdOk returns a tuple with the SubscriptionListingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildRoleTagsResponse) GetSubscriptionListingIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionListingId) {
		return nil, false
	}
	return o.SubscriptionListingId, true
}

// HasSubscriptionListingId returns a boolean if a field has been set.
func (o *GuildRoleTagsResponse) HasSubscriptionListingId() bool {
	if o != nil && !IsNil(o.SubscriptionListingId) {
		return true
	}

	return false
}

// SetSubscriptionListingId gets a reference to the given string and assigns it to the SubscriptionListingId field.
func (o *GuildRoleTagsResponse) SetSubscriptionListingId(v string) {
	o.SubscriptionListingId = &v
}


// GetAvailableForPurchase returns the AvailableForPurchase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleTagsResponse) GetAvailableForPurchase() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AvailableForPurchase
}

// GetAvailableForPurchaseOk returns a tuple with the AvailableForPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleTagsResponse) GetAvailableForPurchaseOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableForPurchase, true
}

// HasAvailableForPurchase returns a boolean if a field has been set.
func (o *GuildRoleTagsResponse) HasAvailableForPurchase() bool {
	if o != nil && !IsNil(o.AvailableForPurchase) {
		return true
	}

	return false
}

// SetAvailableForPurchase gets a reference to the given interface{} and assigns it to the AvailableForPurchase field.
func (o *GuildRoleTagsResponse) SetAvailableForPurchase(v interface{}) {
	o.AvailableForPurchase = v
}


// GetGuildConnections returns the GuildConnections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleTagsResponse) GetGuildConnections() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GuildConnections
}

// GetGuildConnectionsOk returns a tuple with the GuildConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleTagsResponse) GetGuildConnectionsOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildConnections, true
}

// HasGuildConnections returns a boolean if a field has been set.
func (o *GuildRoleTagsResponse) HasGuildConnections() bool {
	if o != nil && !IsNil(o.GuildConnections) {
		return true
	}

	return false
}

// SetGuildConnections gets a reference to the given interface{} and assigns it to the GuildConnections field.
func (o *GuildRoleTagsResponse) SetGuildConnections(v interface{}) {
	o.GuildConnections = v
}


func (o GuildRoleTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildRoleTagsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PremiumSubscriber != nil {
		toSerialize["premium_subscriber"] = o.PremiumSubscriber
	}
	if !IsNil(o.BotId) {
		toSerialize["bot_id"] = o.BotId
	}
	if !IsNil(o.IntegrationId) {
		toSerialize["integration_id"] = o.IntegrationId
	}
	if !IsNil(o.SubscriptionListingId) {
		toSerialize["subscription_listing_id"] = o.SubscriptionListingId
	}
	if o.AvailableForPurchase != nil {
		toSerialize["available_for_purchase"] = o.AvailableForPurchase
	}
	if o.GuildConnections != nil {
		toSerialize["guild_connections"] = o.GuildConnections
	}
	return toSerialize, nil
}

type NullableGuildRoleTagsResponse struct {
	value *GuildRoleTagsResponse
	isSet bool
}

func (v NullableGuildRoleTagsResponse) Get() *GuildRoleTagsResponse {
	return v.value
}

func (v *NullableGuildRoleTagsResponse) Set(val *GuildRoleTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildRoleTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildRoleTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildRoleTagsResponse(val *GuildRoleTagsResponse) *NullableGuildRoleTagsResponse {
	return &NullableGuildRoleTagsResponse{value: val, isSet: true}
}

func (v NullableGuildRoleTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildRoleTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


