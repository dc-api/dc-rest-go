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

// BasicMessageResponseComponentsInner - struct for BasicMessageResponseComponentsInner
type BasicMessageResponseComponentsInner struct {
	ActionRowComponentResponse *ActionRowComponentResponse
	ContainerComponentResponse *ContainerComponentResponse
	FileComponentResponse *FileComponentResponse
	MediaGalleryComponentResponse *MediaGalleryComponentResponse
	SectionComponentResponse *SectionComponentResponse
	SeparatorComponentResponse *SeparatorComponentResponse
	TextDisplayComponentResponse *TextDisplayComponentResponse
}

// ActionRowComponentResponseAsBasicMessageResponseComponentsInner is a convenience function that returns ActionRowComponentResponse wrapped in BasicMessageResponseComponentsInner
func ActionRowComponentResponseAsBasicMessageResponseComponentsInner(v *ActionRowComponentResponse) BasicMessageResponseComponentsInner {
	return BasicMessageResponseComponentsInner{
		ActionRowComponentResponse: v,
	}
}

// ContainerComponentResponseAsBasicMessageResponseComponentsInner is a convenience function that returns ContainerComponentResponse wrapped in BasicMessageResponseComponentsInner
func ContainerComponentResponseAsBasicMessageResponseComponentsInner(v *ContainerComponentResponse) BasicMessageResponseComponentsInner {
	return BasicMessageResponseComponentsInner{
		ContainerComponentResponse: v,
	}
}

// FileComponentResponseAsBasicMessageResponseComponentsInner is a convenience function that returns FileComponentResponse wrapped in BasicMessageResponseComponentsInner
func FileComponentResponseAsBasicMessageResponseComponentsInner(v *FileComponentResponse) BasicMessageResponseComponentsInner {
	return BasicMessageResponseComponentsInner{
		FileComponentResponse: v,
	}
}

// MediaGalleryComponentResponseAsBasicMessageResponseComponentsInner is a convenience function that returns MediaGalleryComponentResponse wrapped in BasicMessageResponseComponentsInner
func MediaGalleryComponentResponseAsBasicMessageResponseComponentsInner(v *MediaGalleryComponentResponse) BasicMessageResponseComponentsInner {
	return BasicMessageResponseComponentsInner{
		MediaGalleryComponentResponse: v,
	}
}

// SectionComponentResponseAsBasicMessageResponseComponentsInner is a convenience function that returns SectionComponentResponse wrapped in BasicMessageResponseComponentsInner
func SectionComponentResponseAsBasicMessageResponseComponentsInner(v *SectionComponentResponse) BasicMessageResponseComponentsInner {
	return BasicMessageResponseComponentsInner{
		SectionComponentResponse: v,
	}
}

// SeparatorComponentResponseAsBasicMessageResponseComponentsInner is a convenience function that returns SeparatorComponentResponse wrapped in BasicMessageResponseComponentsInner
func SeparatorComponentResponseAsBasicMessageResponseComponentsInner(v *SeparatorComponentResponse) BasicMessageResponseComponentsInner {
	return BasicMessageResponseComponentsInner{
		SeparatorComponentResponse: v,
	}
}

// TextDisplayComponentResponseAsBasicMessageResponseComponentsInner is a convenience function that returns TextDisplayComponentResponse wrapped in BasicMessageResponseComponentsInner
func TextDisplayComponentResponseAsBasicMessageResponseComponentsInner(v *TextDisplayComponentResponse) BasicMessageResponseComponentsInner {
	return BasicMessageResponseComponentsInner{
		TextDisplayComponentResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BasicMessageResponseComponentsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActionRowComponentResponse
	err = newStrictDecoder(data).Decode(&dst.ActionRowComponentResponse)
	if err == nil {
		jsonActionRowComponentResponse, _ := json.Marshal(dst.ActionRowComponentResponse)
		if string(jsonActionRowComponentResponse) == "{}" { // empty struct
			dst.ActionRowComponentResponse = nil
		} else {
			if err = validator.Validate(dst.ActionRowComponentResponse); err != nil {
				dst.ActionRowComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActionRowComponentResponse = nil
	}

	// try to unmarshal data into ContainerComponentResponse
	err = newStrictDecoder(data).Decode(&dst.ContainerComponentResponse)
	if err == nil {
		jsonContainerComponentResponse, _ := json.Marshal(dst.ContainerComponentResponse)
		if string(jsonContainerComponentResponse) == "{}" { // empty struct
			dst.ContainerComponentResponse = nil
		} else {
			if err = validator.Validate(dst.ContainerComponentResponse); err != nil {
				dst.ContainerComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ContainerComponentResponse = nil
	}

	// try to unmarshal data into FileComponentResponse
	err = newStrictDecoder(data).Decode(&dst.FileComponentResponse)
	if err == nil {
		jsonFileComponentResponse, _ := json.Marshal(dst.FileComponentResponse)
		if string(jsonFileComponentResponse) == "{}" { // empty struct
			dst.FileComponentResponse = nil
		} else {
			if err = validator.Validate(dst.FileComponentResponse); err != nil {
				dst.FileComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.FileComponentResponse = nil
	}

	// try to unmarshal data into MediaGalleryComponentResponse
	err = newStrictDecoder(data).Decode(&dst.MediaGalleryComponentResponse)
	if err == nil {
		jsonMediaGalleryComponentResponse, _ := json.Marshal(dst.MediaGalleryComponentResponse)
		if string(jsonMediaGalleryComponentResponse) == "{}" { // empty struct
			dst.MediaGalleryComponentResponse = nil
		} else {
			if err = validator.Validate(dst.MediaGalleryComponentResponse); err != nil {
				dst.MediaGalleryComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MediaGalleryComponentResponse = nil
	}

	// try to unmarshal data into SectionComponentResponse
	err = newStrictDecoder(data).Decode(&dst.SectionComponentResponse)
	if err == nil {
		jsonSectionComponentResponse, _ := json.Marshal(dst.SectionComponentResponse)
		if string(jsonSectionComponentResponse) == "{}" { // empty struct
			dst.SectionComponentResponse = nil
		} else {
			if err = validator.Validate(dst.SectionComponentResponse); err != nil {
				dst.SectionComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.SectionComponentResponse = nil
	}

	// try to unmarshal data into SeparatorComponentResponse
	err = newStrictDecoder(data).Decode(&dst.SeparatorComponentResponse)
	if err == nil {
		jsonSeparatorComponentResponse, _ := json.Marshal(dst.SeparatorComponentResponse)
		if string(jsonSeparatorComponentResponse) == "{}" { // empty struct
			dst.SeparatorComponentResponse = nil
		} else {
			if err = validator.Validate(dst.SeparatorComponentResponse); err != nil {
				dst.SeparatorComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.SeparatorComponentResponse = nil
	}

	// try to unmarshal data into TextDisplayComponentResponse
	err = newStrictDecoder(data).Decode(&dst.TextDisplayComponentResponse)
	if err == nil {
		jsonTextDisplayComponentResponse, _ := json.Marshal(dst.TextDisplayComponentResponse)
		if string(jsonTextDisplayComponentResponse) == "{}" { // empty struct
			dst.TextDisplayComponentResponse = nil
		} else {
			if err = validator.Validate(dst.TextDisplayComponentResponse); err != nil {
				dst.TextDisplayComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.TextDisplayComponentResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActionRowComponentResponse = nil
		dst.ContainerComponentResponse = nil
		dst.FileComponentResponse = nil
		dst.MediaGalleryComponentResponse = nil
		dst.SectionComponentResponse = nil
		dst.SeparatorComponentResponse = nil
		dst.TextDisplayComponentResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BasicMessageResponseComponentsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BasicMessageResponseComponentsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BasicMessageResponseComponentsInner) MarshalJSON() ([]byte, error) {
	if src.ActionRowComponentResponse != nil {
		return json.Marshal(&src.ActionRowComponentResponse)
	}

	if src.ContainerComponentResponse != nil {
		return json.Marshal(&src.ContainerComponentResponse)
	}

	if src.FileComponentResponse != nil {
		return json.Marshal(&src.FileComponentResponse)
	}

	if src.MediaGalleryComponentResponse != nil {
		return json.Marshal(&src.MediaGalleryComponentResponse)
	}

	if src.SectionComponentResponse != nil {
		return json.Marshal(&src.SectionComponentResponse)
	}

	if src.SeparatorComponentResponse != nil {
		return json.Marshal(&src.SeparatorComponentResponse)
	}

	if src.TextDisplayComponentResponse != nil {
		return json.Marshal(&src.TextDisplayComponentResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BasicMessageResponseComponentsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ActionRowComponentResponse != nil {
		return obj.ActionRowComponentResponse
	}

	if obj.ContainerComponentResponse != nil {
		return obj.ContainerComponentResponse
	}

	if obj.FileComponentResponse != nil {
		return obj.FileComponentResponse
	}

	if obj.MediaGalleryComponentResponse != nil {
		return obj.MediaGalleryComponentResponse
	}

	if obj.SectionComponentResponse != nil {
		return obj.SectionComponentResponse
	}

	if obj.SeparatorComponentResponse != nil {
		return obj.SeparatorComponentResponse
	}

	if obj.TextDisplayComponentResponse != nil {
		return obj.TextDisplayComponentResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj BasicMessageResponseComponentsInner) GetActualInstanceValue() (interface{}) {
	if obj.ActionRowComponentResponse != nil {
		return *obj.ActionRowComponentResponse
	}

	if obj.ContainerComponentResponse != nil {
		return *obj.ContainerComponentResponse
	}

	if obj.FileComponentResponse != nil {
		return *obj.FileComponentResponse
	}

	if obj.MediaGalleryComponentResponse != nil {
		return *obj.MediaGalleryComponentResponse
	}

	if obj.SectionComponentResponse != nil {
		return *obj.SectionComponentResponse
	}

	if obj.SeparatorComponentResponse != nil {
		return *obj.SeparatorComponentResponse
	}

	if obj.TextDisplayComponentResponse != nil {
		return *obj.TextDisplayComponentResponse
	}

	// all schemas are nil
	return nil
}

type NullableBasicMessageResponseComponentsInner struct {
	value *BasicMessageResponseComponentsInner
	isSet bool
}

func (v NullableBasicMessageResponseComponentsInner) Get() *BasicMessageResponseComponentsInner {
	return v.value
}

func (v *NullableBasicMessageResponseComponentsInner) Set(val *BasicMessageResponseComponentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicMessageResponseComponentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicMessageResponseComponentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicMessageResponseComponentsInner(val *BasicMessageResponseComponentsInner) *NullableBasicMessageResponseComponentsInner {
	return &NullableBasicMessageResponseComponentsInner{value: val, isSet: true}
}

func (v NullableBasicMessageResponseComponentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicMessageResponseComponentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


