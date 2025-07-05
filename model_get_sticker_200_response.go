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

// GetSticker200Response - struct for GetSticker200Response
type GetSticker200Response struct {
	GuildStickerResponse *GuildStickerResponse
	StandardStickerResponse *StandardStickerResponse
}

// GuildStickerResponseAsGetSticker200Response is a convenience function that returns GuildStickerResponse wrapped in GetSticker200Response
func GuildStickerResponseAsGetSticker200Response(v *GuildStickerResponse) GetSticker200Response {
	return GetSticker200Response{
		GuildStickerResponse: v,
	}
}

// StandardStickerResponseAsGetSticker200Response is a convenience function that returns StandardStickerResponse wrapped in GetSticker200Response
func StandardStickerResponseAsGetSticker200Response(v *StandardStickerResponse) GetSticker200Response {
	return GetSticker200Response{
		StandardStickerResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSticker200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GuildStickerResponse
	err = newStrictDecoder(data).Decode(&dst.GuildStickerResponse)
	if err == nil {
		jsonGuildStickerResponse, _ := json.Marshal(dst.GuildStickerResponse)
		if string(jsonGuildStickerResponse) == "{}" { // empty struct
			dst.GuildStickerResponse = nil
		} else {
			if err = validator.Validate(dst.GuildStickerResponse); err != nil {
				dst.GuildStickerResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.GuildStickerResponse = nil
	}

	// try to unmarshal data into StandardStickerResponse
	err = newStrictDecoder(data).Decode(&dst.StandardStickerResponse)
	if err == nil {
		jsonStandardStickerResponse, _ := json.Marshal(dst.StandardStickerResponse)
		if string(jsonStandardStickerResponse) == "{}" { // empty struct
			dst.StandardStickerResponse = nil
		} else {
			if err = validator.Validate(dst.StandardStickerResponse); err != nil {
				dst.StandardStickerResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.StandardStickerResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GuildStickerResponse = nil
		dst.StandardStickerResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSticker200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSticker200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSticker200Response) MarshalJSON() ([]byte, error) {
	if src.GuildStickerResponse != nil {
		return json.Marshal(&src.GuildStickerResponse)
	}

	if src.StandardStickerResponse != nil {
		return json.Marshal(&src.StandardStickerResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSticker200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GuildStickerResponse != nil {
		return obj.GuildStickerResponse
	}

	if obj.StandardStickerResponse != nil {
		return obj.StandardStickerResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetSticker200Response) GetActualInstanceValue() (interface{}) {
	if obj.GuildStickerResponse != nil {
		return *obj.GuildStickerResponse
	}

	if obj.StandardStickerResponse != nil {
		return *obj.StandardStickerResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetSticker200Response struct {
	value *GetSticker200Response
	isSet bool
}

func (v NullableGetSticker200Response) Get() *GetSticker200Response {
	return v.value
}

func (v *NullableGetSticker200Response) Set(val *GetSticker200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSticker200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSticker200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSticker200Response(val *GetSticker200Response) *NullableGetSticker200Response {
	return &NullableGetSticker200Response{value: val, isSet: true}
}

func (v NullableGetSticker200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSticker200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


