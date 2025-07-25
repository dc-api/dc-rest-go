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
	"fmt"
	"gopkg.in/validator.v2"
)

// ListGuildIntegrations200ResponseInner - struct for ListGuildIntegrations200ResponseInner
type ListGuildIntegrations200ResponseInner struct {
	DiscordIntegrationResponse *DiscordIntegrationResponse
	ExternalConnectionIntegrationResponse *ExternalConnectionIntegrationResponse
	GuildSubscriptionIntegrationResponse *GuildSubscriptionIntegrationResponse
}

// DiscordIntegrationResponseAsListGuildIntegrations200ResponseInner is a convenience function that returns DiscordIntegrationResponse wrapped in ListGuildIntegrations200ResponseInner
func DiscordIntegrationResponseAsListGuildIntegrations200ResponseInner(v *DiscordIntegrationResponse) ListGuildIntegrations200ResponseInner {
	return ListGuildIntegrations200ResponseInner{
		DiscordIntegrationResponse: v,
	}
}

// ExternalConnectionIntegrationResponseAsListGuildIntegrations200ResponseInner is a convenience function that returns ExternalConnectionIntegrationResponse wrapped in ListGuildIntegrations200ResponseInner
func ExternalConnectionIntegrationResponseAsListGuildIntegrations200ResponseInner(v *ExternalConnectionIntegrationResponse) ListGuildIntegrations200ResponseInner {
	return ListGuildIntegrations200ResponseInner{
		ExternalConnectionIntegrationResponse: v,
	}
}

// GuildSubscriptionIntegrationResponseAsListGuildIntegrations200ResponseInner is a convenience function that returns GuildSubscriptionIntegrationResponse wrapped in ListGuildIntegrations200ResponseInner
func GuildSubscriptionIntegrationResponseAsListGuildIntegrations200ResponseInner(v *GuildSubscriptionIntegrationResponse) ListGuildIntegrations200ResponseInner {
	return ListGuildIntegrations200ResponseInner{
		GuildSubscriptionIntegrationResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListGuildIntegrations200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DiscordIntegrationResponse
	err = newStrictDecoder(data).Decode(&dst.DiscordIntegrationResponse)
	if err == nil {
		jsonDiscordIntegrationResponse, _ := json.Marshal(dst.DiscordIntegrationResponse)
		if string(jsonDiscordIntegrationResponse) == "{}" { // empty struct
			dst.DiscordIntegrationResponse = nil
		} else {
			if err = validator.Validate(dst.DiscordIntegrationResponse); err != nil {
				dst.DiscordIntegrationResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.DiscordIntegrationResponse = nil
	}

	// try to unmarshal data into ExternalConnectionIntegrationResponse
	err = newStrictDecoder(data).Decode(&dst.ExternalConnectionIntegrationResponse)
	if err == nil {
		jsonExternalConnectionIntegrationResponse, _ := json.Marshal(dst.ExternalConnectionIntegrationResponse)
		if string(jsonExternalConnectionIntegrationResponse) == "{}" { // empty struct
			dst.ExternalConnectionIntegrationResponse = nil
		} else {
			if err = validator.Validate(dst.ExternalConnectionIntegrationResponse); err != nil {
				dst.ExternalConnectionIntegrationResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ExternalConnectionIntegrationResponse = nil
	}

	// try to unmarshal data into GuildSubscriptionIntegrationResponse
	err = newStrictDecoder(data).Decode(&dst.GuildSubscriptionIntegrationResponse)
	if err == nil {
		jsonGuildSubscriptionIntegrationResponse, _ := json.Marshal(dst.GuildSubscriptionIntegrationResponse)
		if string(jsonGuildSubscriptionIntegrationResponse) == "{}" { // empty struct
			dst.GuildSubscriptionIntegrationResponse = nil
		} else {
			if err = validator.Validate(dst.GuildSubscriptionIntegrationResponse); err != nil {
				dst.GuildSubscriptionIntegrationResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.GuildSubscriptionIntegrationResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DiscordIntegrationResponse = nil
		dst.ExternalConnectionIntegrationResponse = nil
		dst.GuildSubscriptionIntegrationResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListGuildIntegrations200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListGuildIntegrations200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListGuildIntegrations200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.DiscordIntegrationResponse != nil {
		return json.Marshal(&src.DiscordIntegrationResponse)
	}

	if src.ExternalConnectionIntegrationResponse != nil {
		return json.Marshal(&src.ExternalConnectionIntegrationResponse)
	}

	if src.GuildSubscriptionIntegrationResponse != nil {
		return json.Marshal(&src.GuildSubscriptionIntegrationResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListGuildIntegrations200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DiscordIntegrationResponse != nil {
		return obj.DiscordIntegrationResponse
	}

	if obj.ExternalConnectionIntegrationResponse != nil {
		return obj.ExternalConnectionIntegrationResponse
	}

	if obj.GuildSubscriptionIntegrationResponse != nil {
		return obj.GuildSubscriptionIntegrationResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListGuildIntegrations200ResponseInner) GetActualInstanceValue() (interface{}) {
	if obj.DiscordIntegrationResponse != nil {
		return *obj.DiscordIntegrationResponse
	}

	if obj.ExternalConnectionIntegrationResponse != nil {
		return *obj.ExternalConnectionIntegrationResponse
	}

	if obj.GuildSubscriptionIntegrationResponse != nil {
		return *obj.GuildSubscriptionIntegrationResponse
	}

	// all schemas are nil
	return nil
}

type NullableListGuildIntegrations200ResponseInner struct {
	value *ListGuildIntegrations200ResponseInner
	isSet bool
}

func (v NullableListGuildIntegrations200ResponseInner) Get() *ListGuildIntegrations200ResponseInner {
	return v.value
}

func (v *NullableListGuildIntegrations200ResponseInner) Set(val *ListGuildIntegrations200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListGuildIntegrations200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListGuildIntegrations200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGuildIntegrations200ResponseInner(val *ListGuildIntegrations200ResponseInner) *NullableListGuildIntegrations200ResponseInner {
	return &NullableListGuildIntegrations200ResponseInner{value: val, isSet: true}
}

func (v NullableListGuildIntegrations200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGuildIntegrations200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


