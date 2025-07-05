/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
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

// GetChannel200Response - struct for GetChannel200Response
type GetChannel200Response struct {
	GuildChannelResponse *GuildChannelResponse
	PrivateChannelResponse *PrivateChannelResponse
	PrivateGroupChannelResponse *PrivateGroupChannelResponse
	ThreadResponse *ThreadResponse
}

// GuildChannelResponseAsGetChannel200Response is a convenience function that returns GuildChannelResponse wrapped in GetChannel200Response
func GuildChannelResponseAsGetChannel200Response(v *GuildChannelResponse) GetChannel200Response {
	return GetChannel200Response{
		GuildChannelResponse: v,
	}
}

// PrivateChannelResponseAsGetChannel200Response is a convenience function that returns PrivateChannelResponse wrapped in GetChannel200Response
func PrivateChannelResponseAsGetChannel200Response(v *PrivateChannelResponse) GetChannel200Response {
	return GetChannel200Response{
		PrivateChannelResponse: v,
	}
}

// PrivateGroupChannelResponseAsGetChannel200Response is a convenience function that returns PrivateGroupChannelResponse wrapped in GetChannel200Response
func PrivateGroupChannelResponseAsGetChannel200Response(v *PrivateGroupChannelResponse) GetChannel200Response {
	return GetChannel200Response{
		PrivateGroupChannelResponse: v,
	}
}

// ThreadResponseAsGetChannel200Response is a convenience function that returns ThreadResponse wrapped in GetChannel200Response
func ThreadResponseAsGetChannel200Response(v *ThreadResponse) GetChannel200Response {
	return GetChannel200Response{
		ThreadResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetChannel200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GuildChannelResponse
	err = newStrictDecoder(data).Decode(&dst.GuildChannelResponse)
	if err == nil {
		jsonGuildChannelResponse, _ := json.Marshal(dst.GuildChannelResponse)
		if string(jsonGuildChannelResponse) == "{}" { // empty struct
			dst.GuildChannelResponse = nil
		} else {
			if err = validator.Validate(dst.GuildChannelResponse); err != nil {
				dst.GuildChannelResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.GuildChannelResponse = nil
	}

	// try to unmarshal data into PrivateChannelResponse
	err = newStrictDecoder(data).Decode(&dst.PrivateChannelResponse)
	if err == nil {
		jsonPrivateChannelResponse, _ := json.Marshal(dst.PrivateChannelResponse)
		if string(jsonPrivateChannelResponse) == "{}" { // empty struct
			dst.PrivateChannelResponse = nil
		} else {
			if err = validator.Validate(dst.PrivateChannelResponse); err != nil {
				dst.PrivateChannelResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.PrivateChannelResponse = nil
	}

	// try to unmarshal data into PrivateGroupChannelResponse
	err = newStrictDecoder(data).Decode(&dst.PrivateGroupChannelResponse)
	if err == nil {
		jsonPrivateGroupChannelResponse, _ := json.Marshal(dst.PrivateGroupChannelResponse)
		if string(jsonPrivateGroupChannelResponse) == "{}" { // empty struct
			dst.PrivateGroupChannelResponse = nil
		} else {
			if err = validator.Validate(dst.PrivateGroupChannelResponse); err != nil {
				dst.PrivateGroupChannelResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.PrivateGroupChannelResponse = nil
	}

	// try to unmarshal data into ThreadResponse
	err = newStrictDecoder(data).Decode(&dst.ThreadResponse)
	if err == nil {
		jsonThreadResponse, _ := json.Marshal(dst.ThreadResponse)
		if string(jsonThreadResponse) == "{}" { // empty struct
			dst.ThreadResponse = nil
		} else {
			if err = validator.Validate(dst.ThreadResponse); err != nil {
				dst.ThreadResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ThreadResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GuildChannelResponse = nil
		dst.PrivateChannelResponse = nil
		dst.PrivateGroupChannelResponse = nil
		dst.ThreadResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetChannel200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetChannel200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetChannel200Response) MarshalJSON() ([]byte, error) {
	if src.GuildChannelResponse != nil {
		return json.Marshal(&src.GuildChannelResponse)
	}

	if src.PrivateChannelResponse != nil {
		return json.Marshal(&src.PrivateChannelResponse)
	}

	if src.PrivateGroupChannelResponse != nil {
		return json.Marshal(&src.PrivateGroupChannelResponse)
	}

	if src.ThreadResponse != nil {
		return json.Marshal(&src.ThreadResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetChannel200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GuildChannelResponse != nil {
		return obj.GuildChannelResponse
	}

	if obj.PrivateChannelResponse != nil {
		return obj.PrivateChannelResponse
	}

	if obj.PrivateGroupChannelResponse != nil {
		return obj.PrivateGroupChannelResponse
	}

	if obj.ThreadResponse != nil {
		return obj.ThreadResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetChannel200Response) GetActualInstanceValue() (interface{}) {
	if obj.GuildChannelResponse != nil {
		return *obj.GuildChannelResponse
	}

	if obj.PrivateChannelResponse != nil {
		return *obj.PrivateChannelResponse
	}

	if obj.PrivateGroupChannelResponse != nil {
		return *obj.PrivateGroupChannelResponse
	}

	if obj.ThreadResponse != nil {
		return *obj.ThreadResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetChannel200Response struct {
	value *GetChannel200Response
	isSet bool
}

func (v NullableGetChannel200Response) Get() *GetChannel200Response {
	return v.value
}

func (v *NullableGetChannel200Response) Set(val *GetChannel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChannel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChannel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChannel200Response(val *GetChannel200Response) *NullableGetChannel200Response {
	return &NullableGetChannel200Response{value: val, isSet: true}
}

func (v NullableGetChannel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChannel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


