/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"os"
	"fmt"
)

// checks if the FileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileRequest{}

// FileRequest struct for FileRequest
type FileRequest struct {
	File *os.File `json:"file"`
	Title string `json:"title"`
	AdditionalProperties map[string]interface{}
}

type _FileRequest FileRequest

// NewFileRequest instantiates a new FileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileRequest(file *os.File, title string) *FileRequest {
	this := FileRequest{}
	this.File = file
	this.Title = title
	return &this
}

// NewFileRequestWithDefaults instantiates a new FileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileRequestWithDefaults() *FileRequest {
	this := FileRequest{}
	return &this
}

// GetFile returns the File field value
func (o *FileRequest) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *FileRequest) GetFileOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *FileRequest) SetFile(v *os.File) {
	o.File = v
}


// GetTitle returns the Title field value
func (o *FileRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *FileRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *FileRequest) SetTitle(v string) {
	o.Title = v
}


func (o FileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file",
		"title",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varFileRequest := _FileRequest{}

	err = json.Unmarshal(data, &varFileRequest)

	if err != nil {
		return err
	}

	*o = FileRequest(varFileRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileRequest struct {
	value *FileRequest
	isSet bool
}

func (v NullableFileRequest) Get() *FileRequest {
	return v.value
}

func (v *NullableFileRequest) Set(val *FileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileRequest(val *FileRequest) *NullableFileRequest {
	return &NullableFileRequest{value: val, isSet: true}
}

func (v NullableFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


