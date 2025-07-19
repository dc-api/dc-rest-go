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

// checks if the FollowChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowChannelRequest{}

// FollowChannelRequest struct for FollowChannelRequest
type FollowChannelRequest struct {
	WebhookChannelId string `json:"webhook_channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _FollowChannelRequest FollowChannelRequest

// NewFollowChannelRequest instantiates a new FollowChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowChannelRequest(webhookChannelId string) *FollowChannelRequest {
	this := FollowChannelRequest{}
	this.WebhookChannelId = webhookChannelId
	return &this
}

// NewFollowChannelRequestWithDefaults instantiates a new FollowChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowChannelRequestWithDefaults() *FollowChannelRequest {
	this := FollowChannelRequest{}
	return &this
}

// GetWebhookChannelId returns the WebhookChannelId field value
func (o *FollowChannelRequest) GetWebhookChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookChannelId
}

// GetWebhookChannelIdOk returns a tuple with the WebhookChannelId field value
// and a boolean to check if the value has been set.
func (o *FollowChannelRequest) GetWebhookChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookChannelId, true
}

// SetWebhookChannelId sets field value
func (o *FollowChannelRequest) SetWebhookChannelId(v string) {
	o.WebhookChannelId = v
}

func (o FollowChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["webhook_channel_id"] = o.WebhookChannelId
	return toSerialize, nil
}

func (o *FollowChannelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"webhook_channel_id",
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

	varFollowChannelRequest := _FollowChannelRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFollowChannelRequest)

	if err != nil {
		return err
	}

	*o = FollowChannelRequest(varFollowChannelRequest)

	return err
}

type NullableFollowChannelRequest struct {
	value *FollowChannelRequest
	isSet bool
}

func (v NullableFollowChannelRequest) Get() *FollowChannelRequest {
	return v.value
}

func (v *NullableFollowChannelRequest) Set(val *FollowChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowChannelRequest(val *FollowChannelRequest) *NullableFollowChannelRequest {
	return &NullableFollowChannelRequest{value: val, isSet: true}
}

func (v NullableFollowChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


