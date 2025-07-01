/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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
)


// ListChannelInvites200ResponseInner struct for ListChannelInvites200ResponseInner
type ListChannelInvites200ResponseInner struct {
	FriendInviteResponse *FriendInviteResponse
	GroupDMInviteResponse *GroupDMInviteResponse
	GuildInviteResponse *GuildInviteResponse
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ListChannelInvites200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FriendInviteResponse
	err = json.Unmarshal(data, &dst.FriendInviteResponse);
	if err == nil {
		jsonFriendInviteResponse, _ := json.Marshal(dst.FriendInviteResponse)
		if string(jsonFriendInviteResponse) == "{}" { // empty struct
			dst.FriendInviteResponse = nil
		} else {
			return nil // data stored in dst.FriendInviteResponse, return on the first match
		}
	} else {
		dst.FriendInviteResponse = nil
	}

	// try to unmarshal JSON data into GroupDMInviteResponse
	err = json.Unmarshal(data, &dst.GroupDMInviteResponse);
	if err == nil {
		jsonGroupDMInviteResponse, _ := json.Marshal(dst.GroupDMInviteResponse)
		if string(jsonGroupDMInviteResponse) == "{}" { // empty struct
			dst.GroupDMInviteResponse = nil
		} else {
			return nil // data stored in dst.GroupDMInviteResponse, return on the first match
		}
	} else {
		dst.GroupDMInviteResponse = nil
	}

	// try to unmarshal JSON data into GuildInviteResponse
	err = json.Unmarshal(data, &dst.GuildInviteResponse);
	if err == nil {
		jsonGuildInviteResponse, _ := json.Marshal(dst.GuildInviteResponse)
		if string(jsonGuildInviteResponse) == "{}" { // empty struct
			dst.GuildInviteResponse = nil
		} else {
			return nil // data stored in dst.GuildInviteResponse, return on the first match
		}
	} else {
		dst.GuildInviteResponse = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ListChannelInvites200ResponseInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListChannelInvites200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.FriendInviteResponse != nil {
		return json.Marshal(&src.FriendInviteResponse)
	}

	if src.GroupDMInviteResponse != nil {
		return json.Marshal(&src.GroupDMInviteResponse)
	}

	if src.GuildInviteResponse != nil {
		return json.Marshal(&src.GuildInviteResponse)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableListChannelInvites200ResponseInner struct {
	value *ListChannelInvites200ResponseInner
	isSet bool
}

func (v NullableListChannelInvites200ResponseInner) Get() *ListChannelInvites200ResponseInner {
	return v.value
}

func (v *NullableListChannelInvites200ResponseInner) Set(val *ListChannelInvites200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListChannelInvites200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListChannelInvites200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListChannelInvites200ResponseInner(val *ListChannelInvites200ResponseInner) *NullableListChannelInvites200ResponseInner {
	return &NullableListChannelInvites200ResponseInner{value: val, isSet: true}
}

func (v NullableListChannelInvites200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListChannelInvites200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


