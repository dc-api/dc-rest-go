/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// ListChannelWebhooks200ResponseInner - struct for ListChannelWebhooks200ResponseInner
type ListChannelWebhooks200ResponseInner struct {
	ApplicationIncomingWebhookResponse *ApplicationIncomingWebhookResponse
	ChannelFollowerWebhookResponse *ChannelFollowerWebhookResponse
	GuildIncomingWebhookResponse *GuildIncomingWebhookResponse
}

// ApplicationIncomingWebhookResponseAsListChannelWebhooks200ResponseInner is a convenience function that returns ApplicationIncomingWebhookResponse wrapped in ListChannelWebhooks200ResponseInner
func ApplicationIncomingWebhookResponseAsListChannelWebhooks200ResponseInner(v *ApplicationIncomingWebhookResponse) ListChannelWebhooks200ResponseInner {
	return ListChannelWebhooks200ResponseInner{
		ApplicationIncomingWebhookResponse: v,
	}
}

// ChannelFollowerWebhookResponseAsListChannelWebhooks200ResponseInner is a convenience function that returns ChannelFollowerWebhookResponse wrapped in ListChannelWebhooks200ResponseInner
func ChannelFollowerWebhookResponseAsListChannelWebhooks200ResponseInner(v *ChannelFollowerWebhookResponse) ListChannelWebhooks200ResponseInner {
	return ListChannelWebhooks200ResponseInner{
		ChannelFollowerWebhookResponse: v,
	}
}

// GuildIncomingWebhookResponseAsListChannelWebhooks200ResponseInner is a convenience function that returns GuildIncomingWebhookResponse wrapped in ListChannelWebhooks200ResponseInner
func GuildIncomingWebhookResponseAsListChannelWebhooks200ResponseInner(v *GuildIncomingWebhookResponse) ListChannelWebhooks200ResponseInner {
	return ListChannelWebhooks200ResponseInner{
		GuildIncomingWebhookResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListChannelWebhooks200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApplicationIncomingWebhookResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationIncomingWebhookResponse)
	if err == nil {
		jsonApplicationIncomingWebhookResponse, _ := json.Marshal(dst.ApplicationIncomingWebhookResponse)
		if string(jsonApplicationIncomingWebhookResponse) == "{}" { // empty struct
			dst.ApplicationIncomingWebhookResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationIncomingWebhookResponse); err != nil {
				dst.ApplicationIncomingWebhookResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationIncomingWebhookResponse = nil
	}

	// try to unmarshal data into ChannelFollowerWebhookResponse
	err = newStrictDecoder(data).Decode(&dst.ChannelFollowerWebhookResponse)
	if err == nil {
		jsonChannelFollowerWebhookResponse, _ := json.Marshal(dst.ChannelFollowerWebhookResponse)
		if string(jsonChannelFollowerWebhookResponse) == "{}" { // empty struct
			dst.ChannelFollowerWebhookResponse = nil
		} else {
			if err = validator.Validate(dst.ChannelFollowerWebhookResponse); err != nil {
				dst.ChannelFollowerWebhookResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChannelFollowerWebhookResponse = nil
	}

	// try to unmarshal data into GuildIncomingWebhookResponse
	err = newStrictDecoder(data).Decode(&dst.GuildIncomingWebhookResponse)
	if err == nil {
		jsonGuildIncomingWebhookResponse, _ := json.Marshal(dst.GuildIncomingWebhookResponse)
		if string(jsonGuildIncomingWebhookResponse) == "{}" { // empty struct
			dst.GuildIncomingWebhookResponse = nil
		} else {
			if err = validator.Validate(dst.GuildIncomingWebhookResponse); err != nil {
				dst.GuildIncomingWebhookResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.GuildIncomingWebhookResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApplicationIncomingWebhookResponse = nil
		dst.ChannelFollowerWebhookResponse = nil
		dst.GuildIncomingWebhookResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListChannelWebhooks200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListChannelWebhooks200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListChannelWebhooks200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.ApplicationIncomingWebhookResponse != nil {
		return json.Marshal(&src.ApplicationIncomingWebhookResponse)
	}

	if src.ChannelFollowerWebhookResponse != nil {
		return json.Marshal(&src.ChannelFollowerWebhookResponse)
	}

	if src.GuildIncomingWebhookResponse != nil {
		return json.Marshal(&src.GuildIncomingWebhookResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListChannelWebhooks200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApplicationIncomingWebhookResponse != nil {
		return obj.ApplicationIncomingWebhookResponse
	}

	if obj.ChannelFollowerWebhookResponse != nil {
		return obj.ChannelFollowerWebhookResponse
	}

	if obj.GuildIncomingWebhookResponse != nil {
		return obj.GuildIncomingWebhookResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListChannelWebhooks200ResponseInner) GetActualInstanceValue() (interface{}) {
	if obj.ApplicationIncomingWebhookResponse != nil {
		return *obj.ApplicationIncomingWebhookResponse
	}

	if obj.ChannelFollowerWebhookResponse != nil {
		return *obj.ChannelFollowerWebhookResponse
	}

	if obj.GuildIncomingWebhookResponse != nil {
		return *obj.GuildIncomingWebhookResponse
	}

	// all schemas are nil
	return nil
}

type NullableListChannelWebhooks200ResponseInner struct {
	value *ListChannelWebhooks200ResponseInner
	isSet bool
}

func (v NullableListChannelWebhooks200ResponseInner) Get() *ListChannelWebhooks200ResponseInner {
	return v.value
}

func (v *NullableListChannelWebhooks200ResponseInner) Set(val *ListChannelWebhooks200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListChannelWebhooks200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListChannelWebhooks200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListChannelWebhooks200ResponseInner(val *ListChannelWebhooks200ResponseInner) *NullableListChannelWebhooks200ResponseInner {
	return &NullableListChannelWebhooks200ResponseInner{value: val, isSet: true}
}

func (v NullableListChannelWebhooks200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListChannelWebhooks200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


