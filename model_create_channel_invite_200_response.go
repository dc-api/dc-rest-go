/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-28T06:15:04.313783603Z[Etc/UTC]
 *    * - **Generator Version**: 7.19.0
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

// CreateChannelInvite200Response - struct for CreateChannelInvite200Response
type CreateChannelInvite200Response struct {
	FriendInviteResponse *FriendInviteResponse
	GroupDMInviteResponse *GroupDMInviteResponse
	GuildInviteResponse *GuildInviteResponse
}

// FriendInviteResponseAsCreateChannelInvite200Response is a convenience function that returns FriendInviteResponse wrapped in CreateChannelInvite200Response
func FriendInviteResponseAsCreateChannelInvite200Response(v *FriendInviteResponse) CreateChannelInvite200Response {
	return CreateChannelInvite200Response{
		FriendInviteResponse: v,
	}
}

// GroupDMInviteResponseAsCreateChannelInvite200Response is a convenience function that returns GroupDMInviteResponse wrapped in CreateChannelInvite200Response
func GroupDMInviteResponseAsCreateChannelInvite200Response(v *GroupDMInviteResponse) CreateChannelInvite200Response {
	return CreateChannelInvite200Response{
		GroupDMInviteResponse: v,
	}
}

// GuildInviteResponseAsCreateChannelInvite200Response is a convenience function that returns GuildInviteResponse wrapped in CreateChannelInvite200Response
func GuildInviteResponseAsCreateChannelInvite200Response(v *GuildInviteResponse) CreateChannelInvite200Response {
	return CreateChannelInvite200Response{
		GuildInviteResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateChannelInvite200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FriendInviteResponse
	err = newStrictDecoder(data).Decode(&dst.FriendInviteResponse)
	if err == nil {
		jsonFriendInviteResponse, _ := json.Marshal(dst.FriendInviteResponse)
		if string(jsonFriendInviteResponse) == "{}" { // empty struct
			dst.FriendInviteResponse = nil
		} else {
			if err = validator.Validate(dst.FriendInviteResponse); err != nil {
				dst.FriendInviteResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.FriendInviteResponse = nil
	}

	// try to unmarshal data into GroupDMInviteResponse
	err = newStrictDecoder(data).Decode(&dst.GroupDMInviteResponse)
	if err == nil {
		jsonGroupDMInviteResponse, _ := json.Marshal(dst.GroupDMInviteResponse)
		if string(jsonGroupDMInviteResponse) == "{}" { // empty struct
			dst.GroupDMInviteResponse = nil
		} else {
			if err = validator.Validate(dst.GroupDMInviteResponse); err != nil {
				dst.GroupDMInviteResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.GroupDMInviteResponse = nil
	}

	// try to unmarshal data into GuildInviteResponse
	err = newStrictDecoder(data).Decode(&dst.GuildInviteResponse)
	if err == nil {
		jsonGuildInviteResponse, _ := json.Marshal(dst.GuildInviteResponse)
		if string(jsonGuildInviteResponse) == "{}" { // empty struct
			dst.GuildInviteResponse = nil
		} else {
			if err = validator.Validate(dst.GuildInviteResponse); err != nil {
				dst.GuildInviteResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.GuildInviteResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FriendInviteResponse = nil
		dst.GroupDMInviteResponse = nil
		dst.GuildInviteResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateChannelInvite200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateChannelInvite200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateChannelInvite200Response) MarshalJSON() ([]byte, error) {
	if src.FriendInviteResponse != nil {
		return json.Marshal(&src.FriendInviteResponse)
	}

	if src.GroupDMInviteResponse != nil {
		return json.Marshal(&src.GroupDMInviteResponse)
	}

	if src.GuildInviteResponse != nil {
		return json.Marshal(&src.GuildInviteResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateChannelInvite200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.FriendInviteResponse != nil {
		return obj.FriendInviteResponse
	}

	if obj.GroupDMInviteResponse != nil {
		return obj.GroupDMInviteResponse
	}

	if obj.GuildInviteResponse != nil {
		return obj.GuildInviteResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CreateChannelInvite200Response) GetActualInstanceValue() (interface{}) {
	if obj.FriendInviteResponse != nil {
		return *obj.FriendInviteResponse
	}

	if obj.GroupDMInviteResponse != nil {
		return *obj.GroupDMInviteResponse
	}

	if obj.GuildInviteResponse != nil {
		return *obj.GuildInviteResponse
	}

	// all schemas are nil
	return nil
}

type NullableCreateChannelInvite200Response struct {
	value *CreateChannelInvite200Response
	isSet bool
}

func (v NullableCreateChannelInvite200Response) Get() *CreateChannelInvite200Response {
	return v.value
}

func (v *NullableCreateChannelInvite200Response) Set(val *CreateChannelInvite200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChannelInvite200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChannelInvite200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChannelInvite200Response(val *CreateChannelInvite200Response) *NullableCreateChannelInvite200Response {
	return &NullableCreateChannelInvite200Response{value: val, isSet: true}
}

func (v NullableCreateChannelInvite200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChannelInvite200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

