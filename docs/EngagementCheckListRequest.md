# EngagementCheckListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionManagement** | Pointer to **string** |  | [optional] 
**EncryptionCrypto** | Pointer to **string** |  | [optional] 
**ConfigurationManagement** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** |  | [optional] 
**AuthorizationAndAccessControl** | Pointer to **string** |  | [optional] 
**DataInputSanitizationValidation** | Pointer to **string** |  | [optional] 
**SensitiveData** | Pointer to **string** |  | [optional] 
**Other** | Pointer to **string** |  | [optional] 
**SessionIssues** | Pointer to **[]int32** |  | [optional] 
**CryptoIssues** | Pointer to **[]int32** |  | [optional] 
**ConfigIssues** | Pointer to **[]int32** |  | [optional] 
**AuthIssues** | Pointer to **[]int32** |  | [optional] 
**AuthorIssues** | Pointer to **[]int32** |  | [optional] 
**DataIssues** | Pointer to **[]int32** |  | [optional] 
**SensitiveIssues** | Pointer to **[]int32** |  | [optional] 
**OtherIssues** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewEngagementCheckListRequest

`func NewEngagementCheckListRequest() *EngagementCheckListRequest`

NewEngagementCheckListRequest instantiates a new EngagementCheckListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementCheckListRequestWithDefaults

`func NewEngagementCheckListRequestWithDefaults() *EngagementCheckListRequest`

NewEngagementCheckListRequestWithDefaults instantiates a new EngagementCheckListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionManagement

`func (o *EngagementCheckListRequest) GetSessionManagement() string`

GetSessionManagement returns the SessionManagement field if non-nil, zero value otherwise.

### GetSessionManagementOk

`func (o *EngagementCheckListRequest) GetSessionManagementOk() (*string, bool)`

GetSessionManagementOk returns a tuple with the SessionManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionManagement

`func (o *EngagementCheckListRequest) SetSessionManagement(v string)`

SetSessionManagement sets SessionManagement field to given value.

### HasSessionManagement

`func (o *EngagementCheckListRequest) HasSessionManagement() bool`

HasSessionManagement returns a boolean if a field has been set.

### GetEncryptionCrypto

`func (o *EngagementCheckListRequest) GetEncryptionCrypto() string`

GetEncryptionCrypto returns the EncryptionCrypto field if non-nil, zero value otherwise.

### GetEncryptionCryptoOk

`func (o *EngagementCheckListRequest) GetEncryptionCryptoOk() (*string, bool)`

GetEncryptionCryptoOk returns a tuple with the EncryptionCrypto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionCrypto

`func (o *EngagementCheckListRequest) SetEncryptionCrypto(v string)`

SetEncryptionCrypto sets EncryptionCrypto field to given value.

### HasEncryptionCrypto

`func (o *EngagementCheckListRequest) HasEncryptionCrypto() bool`

HasEncryptionCrypto returns a boolean if a field has been set.

### GetConfigurationManagement

`func (o *EngagementCheckListRequest) GetConfigurationManagement() string`

GetConfigurationManagement returns the ConfigurationManagement field if non-nil, zero value otherwise.

### GetConfigurationManagementOk

`func (o *EngagementCheckListRequest) GetConfigurationManagementOk() (*string, bool)`

GetConfigurationManagementOk returns a tuple with the ConfigurationManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationManagement

`func (o *EngagementCheckListRequest) SetConfigurationManagement(v string)`

SetConfigurationManagement sets ConfigurationManagement field to given value.

### HasConfigurationManagement

`func (o *EngagementCheckListRequest) HasConfigurationManagement() bool`

HasConfigurationManagement returns a boolean if a field has been set.

### GetAuthentication

`func (o *EngagementCheckListRequest) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *EngagementCheckListRequest) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *EngagementCheckListRequest) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *EngagementCheckListRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAuthorizationAndAccessControl

`func (o *EngagementCheckListRequest) GetAuthorizationAndAccessControl() string`

GetAuthorizationAndAccessControl returns the AuthorizationAndAccessControl field if non-nil, zero value otherwise.

### GetAuthorizationAndAccessControlOk

`func (o *EngagementCheckListRequest) GetAuthorizationAndAccessControlOk() (*string, bool)`

GetAuthorizationAndAccessControlOk returns a tuple with the AuthorizationAndAccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAndAccessControl

`func (o *EngagementCheckListRequest) SetAuthorizationAndAccessControl(v string)`

SetAuthorizationAndAccessControl sets AuthorizationAndAccessControl field to given value.

### HasAuthorizationAndAccessControl

`func (o *EngagementCheckListRequest) HasAuthorizationAndAccessControl() bool`

HasAuthorizationAndAccessControl returns a boolean if a field has been set.

### GetDataInputSanitizationValidation

`func (o *EngagementCheckListRequest) GetDataInputSanitizationValidation() string`

GetDataInputSanitizationValidation returns the DataInputSanitizationValidation field if non-nil, zero value otherwise.

### GetDataInputSanitizationValidationOk

`func (o *EngagementCheckListRequest) GetDataInputSanitizationValidationOk() (*string, bool)`

GetDataInputSanitizationValidationOk returns a tuple with the DataInputSanitizationValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInputSanitizationValidation

`func (o *EngagementCheckListRequest) SetDataInputSanitizationValidation(v string)`

SetDataInputSanitizationValidation sets DataInputSanitizationValidation field to given value.

### HasDataInputSanitizationValidation

`func (o *EngagementCheckListRequest) HasDataInputSanitizationValidation() bool`

HasDataInputSanitizationValidation returns a boolean if a field has been set.

### GetSensitiveData

`func (o *EngagementCheckListRequest) GetSensitiveData() string`

GetSensitiveData returns the SensitiveData field if non-nil, zero value otherwise.

### GetSensitiveDataOk

`func (o *EngagementCheckListRequest) GetSensitiveDataOk() (*string, bool)`

GetSensitiveDataOk returns a tuple with the SensitiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveData

`func (o *EngagementCheckListRequest) SetSensitiveData(v string)`

SetSensitiveData sets SensitiveData field to given value.

### HasSensitiveData

`func (o *EngagementCheckListRequest) HasSensitiveData() bool`

HasSensitiveData returns a boolean if a field has been set.

### GetOther

`func (o *EngagementCheckListRequest) GetOther() string`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *EngagementCheckListRequest) GetOtherOk() (*string, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *EngagementCheckListRequest) SetOther(v string)`

SetOther sets Other field to given value.

### HasOther

`func (o *EngagementCheckListRequest) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetSessionIssues

`func (o *EngagementCheckListRequest) GetSessionIssues() []int32`

GetSessionIssues returns the SessionIssues field if non-nil, zero value otherwise.

### GetSessionIssuesOk

`func (o *EngagementCheckListRequest) GetSessionIssuesOk() (*[]int32, bool)`

GetSessionIssuesOk returns a tuple with the SessionIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIssues

`func (o *EngagementCheckListRequest) SetSessionIssues(v []int32)`

SetSessionIssues sets SessionIssues field to given value.

### HasSessionIssues

`func (o *EngagementCheckListRequest) HasSessionIssues() bool`

HasSessionIssues returns a boolean if a field has been set.

### GetCryptoIssues

`func (o *EngagementCheckListRequest) GetCryptoIssues() []int32`

GetCryptoIssues returns the CryptoIssues field if non-nil, zero value otherwise.

### GetCryptoIssuesOk

`func (o *EngagementCheckListRequest) GetCryptoIssuesOk() (*[]int32, bool)`

GetCryptoIssuesOk returns a tuple with the CryptoIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoIssues

`func (o *EngagementCheckListRequest) SetCryptoIssues(v []int32)`

SetCryptoIssues sets CryptoIssues field to given value.

### HasCryptoIssues

`func (o *EngagementCheckListRequest) HasCryptoIssues() bool`

HasCryptoIssues returns a boolean if a field has been set.

### GetConfigIssues

`func (o *EngagementCheckListRequest) GetConfigIssues() []int32`

GetConfigIssues returns the ConfigIssues field if non-nil, zero value otherwise.

### GetConfigIssuesOk

`func (o *EngagementCheckListRequest) GetConfigIssuesOk() (*[]int32, bool)`

GetConfigIssuesOk returns a tuple with the ConfigIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIssues

`func (o *EngagementCheckListRequest) SetConfigIssues(v []int32)`

SetConfigIssues sets ConfigIssues field to given value.

### HasConfigIssues

`func (o *EngagementCheckListRequest) HasConfigIssues() bool`

HasConfigIssues returns a boolean if a field has been set.

### GetAuthIssues

`func (o *EngagementCheckListRequest) GetAuthIssues() []int32`

GetAuthIssues returns the AuthIssues field if non-nil, zero value otherwise.

### GetAuthIssuesOk

`func (o *EngagementCheckListRequest) GetAuthIssuesOk() (*[]int32, bool)`

GetAuthIssuesOk returns a tuple with the AuthIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIssues

`func (o *EngagementCheckListRequest) SetAuthIssues(v []int32)`

SetAuthIssues sets AuthIssues field to given value.

### HasAuthIssues

`func (o *EngagementCheckListRequest) HasAuthIssues() bool`

HasAuthIssues returns a boolean if a field has been set.

### GetAuthorIssues

`func (o *EngagementCheckListRequest) GetAuthorIssues() []int32`

GetAuthorIssues returns the AuthorIssues field if non-nil, zero value otherwise.

### GetAuthorIssuesOk

`func (o *EngagementCheckListRequest) GetAuthorIssuesOk() (*[]int32, bool)`

GetAuthorIssuesOk returns a tuple with the AuthorIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIssues

`func (o *EngagementCheckListRequest) SetAuthorIssues(v []int32)`

SetAuthorIssues sets AuthorIssues field to given value.

### HasAuthorIssues

`func (o *EngagementCheckListRequest) HasAuthorIssues() bool`

HasAuthorIssues returns a boolean if a field has been set.

### GetDataIssues

`func (o *EngagementCheckListRequest) GetDataIssues() []int32`

GetDataIssues returns the DataIssues field if non-nil, zero value otherwise.

### GetDataIssuesOk

`func (o *EngagementCheckListRequest) GetDataIssuesOk() (*[]int32, bool)`

GetDataIssuesOk returns a tuple with the DataIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIssues

`func (o *EngagementCheckListRequest) SetDataIssues(v []int32)`

SetDataIssues sets DataIssues field to given value.

### HasDataIssues

`func (o *EngagementCheckListRequest) HasDataIssues() bool`

HasDataIssues returns a boolean if a field has been set.

### GetSensitiveIssues

`func (o *EngagementCheckListRequest) GetSensitiveIssues() []int32`

GetSensitiveIssues returns the SensitiveIssues field if non-nil, zero value otherwise.

### GetSensitiveIssuesOk

`func (o *EngagementCheckListRequest) GetSensitiveIssuesOk() (*[]int32, bool)`

GetSensitiveIssuesOk returns a tuple with the SensitiveIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveIssues

`func (o *EngagementCheckListRequest) SetSensitiveIssues(v []int32)`

SetSensitiveIssues sets SensitiveIssues field to given value.

### HasSensitiveIssues

`func (o *EngagementCheckListRequest) HasSensitiveIssues() bool`

HasSensitiveIssues returns a boolean if a field has been set.

### GetOtherIssues

`func (o *EngagementCheckListRequest) GetOtherIssues() []int32`

GetOtherIssues returns the OtherIssues field if non-nil, zero value otherwise.

### GetOtherIssuesOk

`func (o *EngagementCheckListRequest) GetOtherIssuesOk() (*[]int32, bool)`

GetOtherIssuesOk returns a tuple with the OtherIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherIssues

`func (o *EngagementCheckListRequest) SetOtherIssues(v []int32)`

SetOtherIssues sets OtherIssues field to given value.

### HasOtherIssues

`func (o *EngagementCheckListRequest) HasOtherIssues() bool`

HasOtherIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


