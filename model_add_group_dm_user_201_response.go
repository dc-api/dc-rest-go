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
	"fmt"
	"gopkg.in/validator.v2"
)

// AddGroupDmUser201Response - struct for AddGroupDmUser201Response
type AddGroupDmUser201Response struct {
	PrivateChannelResponse *PrivateChannelResponse
	PrivateGroupChannelResponse *PrivateGroupChannelResponse
}

// PrivateChannelResponseAsAddGroupDmUser201Response is a convenience function that returns PrivateChannelResponse wrapped in AddGroupDmUser201Response
func PrivateChannelResponseAsAddGroupDmUser201Response(v *PrivateChannelResponse) AddGroupDmUser201Response {
	return AddGroupDmUser201Response{
		PrivateChannelResponse: v,
	}
}

// PrivateGroupChannelResponseAsAddGroupDmUser201Response is a convenience function that returns PrivateGroupChannelResponse wrapped in AddGroupDmUser201Response
func PrivateGroupChannelResponseAsAddGroupDmUser201Response(v *PrivateGroupChannelResponse) AddGroupDmUser201Response {
	return AddGroupDmUser201Response{
		PrivateGroupChannelResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddGroupDmUser201Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PrivateChannelResponse = nil
		dst.PrivateGroupChannelResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddGroupDmUser201Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddGroupDmUser201Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddGroupDmUser201Response) MarshalJSON() ([]byte, error) {
	if src.PrivateChannelResponse != nil {
		return json.Marshal(&src.PrivateChannelResponse)
	}

	if src.PrivateGroupChannelResponse != nil {
		return json.Marshal(&src.PrivateGroupChannelResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddGroupDmUser201Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PrivateChannelResponse != nil {
		return obj.PrivateChannelResponse
	}

	if obj.PrivateGroupChannelResponse != nil {
		return obj.PrivateGroupChannelResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AddGroupDmUser201Response) GetActualInstanceValue() (interface{}) {
	if obj.PrivateChannelResponse != nil {
		return *obj.PrivateChannelResponse
	}

	if obj.PrivateGroupChannelResponse != nil {
		return *obj.PrivateGroupChannelResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddGroupDmUser201Response struct {
	value *AddGroupDmUser201Response
	isSet bool
}

func (v NullableAddGroupDmUser201Response) Get() *AddGroupDmUser201Response {
	return v.value
}

func (v *NullableAddGroupDmUser201Response) Set(val *AddGroupDmUser201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGroupDmUser201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGroupDmUser201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGroupDmUser201Response(val *AddGroupDmUser201Response) *NullableAddGroupDmUser201Response {
	return &NullableAddGroupDmUser201Response{value: val, isSet: true}
}

func (v NullableAddGroupDmUser201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGroupDmUser201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


