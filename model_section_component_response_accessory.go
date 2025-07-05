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

// SectionComponentResponseAccessory - struct for SectionComponentResponseAccessory
type SectionComponentResponseAccessory struct {
	ButtonComponentResponse *ButtonComponentResponse
	ThumbnailComponentResponse *ThumbnailComponentResponse
}

// ButtonComponentResponseAsSectionComponentResponseAccessory is a convenience function that returns ButtonComponentResponse wrapped in SectionComponentResponseAccessory
func ButtonComponentResponseAsSectionComponentResponseAccessory(v *ButtonComponentResponse) SectionComponentResponseAccessory {
	return SectionComponentResponseAccessory{
		ButtonComponentResponse: v,
	}
}

// ThumbnailComponentResponseAsSectionComponentResponseAccessory is a convenience function that returns ThumbnailComponentResponse wrapped in SectionComponentResponseAccessory
func ThumbnailComponentResponseAsSectionComponentResponseAccessory(v *ThumbnailComponentResponse) SectionComponentResponseAccessory {
	return SectionComponentResponseAccessory{
		ThumbnailComponentResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SectionComponentResponseAccessory) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ButtonComponentResponse
	err = newStrictDecoder(data).Decode(&dst.ButtonComponentResponse)
	if err == nil {
		jsonButtonComponentResponse, _ := json.Marshal(dst.ButtonComponentResponse)
		if string(jsonButtonComponentResponse) == "{}" { // empty struct
			dst.ButtonComponentResponse = nil
		} else {
			if err = validator.Validate(dst.ButtonComponentResponse); err != nil {
				dst.ButtonComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ButtonComponentResponse = nil
	}

	// try to unmarshal data into ThumbnailComponentResponse
	err = newStrictDecoder(data).Decode(&dst.ThumbnailComponentResponse)
	if err == nil {
		jsonThumbnailComponentResponse, _ := json.Marshal(dst.ThumbnailComponentResponse)
		if string(jsonThumbnailComponentResponse) == "{}" { // empty struct
			dst.ThumbnailComponentResponse = nil
		} else {
			if err = validator.Validate(dst.ThumbnailComponentResponse); err != nil {
				dst.ThumbnailComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ThumbnailComponentResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ButtonComponentResponse = nil
		dst.ThumbnailComponentResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SectionComponentResponseAccessory)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SectionComponentResponseAccessory)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SectionComponentResponseAccessory) MarshalJSON() ([]byte, error) {
	if src.ButtonComponentResponse != nil {
		return json.Marshal(&src.ButtonComponentResponse)
	}

	if src.ThumbnailComponentResponse != nil {
		return json.Marshal(&src.ThumbnailComponentResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SectionComponentResponseAccessory) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ButtonComponentResponse != nil {
		return obj.ButtonComponentResponse
	}

	if obj.ThumbnailComponentResponse != nil {
		return obj.ThumbnailComponentResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SectionComponentResponseAccessory) GetActualInstanceValue() (interface{}) {
	if obj.ButtonComponentResponse != nil {
		return *obj.ButtonComponentResponse
	}

	if obj.ThumbnailComponentResponse != nil {
		return *obj.ThumbnailComponentResponse
	}

	// all schemas are nil
	return nil
}

type NullableSectionComponentResponseAccessory struct {
	value *SectionComponentResponseAccessory
	isSet bool
}

func (v NullableSectionComponentResponseAccessory) Get() *SectionComponentResponseAccessory {
	return v.value
}

func (v *NullableSectionComponentResponseAccessory) Set(val *SectionComponentResponseAccessory) {
	v.value = val
	v.isSet = true
}

func (v NullableSectionComponentResponseAccessory) IsSet() bool {
	return v.isSet
}

func (v *NullableSectionComponentResponseAccessory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectionComponentResponseAccessory(val *SectionComponentResponseAccessory) *NullableSectionComponentResponseAccessory {
	return &NullableSectionComponentResponseAccessory{value: val, isSet: true}
}

func (v NullableSectionComponentResponseAccessory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectionComponentResponseAccessory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


