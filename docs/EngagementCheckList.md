# EngagementCheckList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**SessionManagement** | Pointer to **string** |  | [optional] 
**EncryptionCrypto** | Pointer to **string** |  | [optional] 
**ConfigurationManagement** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** |  | [optional] 
**AuthorizationAndAccessControl** | Pointer to **string** |  | [optional] 
**DataInputSanitizationValidation** | Pointer to **string** |  | [optional] 
**SensitiveData** | Pointer to **string** |  | [optional] 
**Other** | Pointer to **string** |  | [optional] 
**Engagement** | **int32** |  | [readonly] 
**SessionIssues** | Pointer to **[]int32** |  | [optional] 
**CryptoIssues** | Pointer to **[]int32** |  | [optional] 
**ConfigIssues** | Pointer to **[]int32** |  | [optional] 
**AuthIssues** | Pointer to **[]int32** |  | [optional] 
**AuthorIssues** | Pointer to **[]int32** |  | [optional] 
**DataIssues** | Pointer to **[]int32** |  | [optional] 
**SensitiveIssues** | Pointer to **[]int32** |  | [optional] 
**OtherIssues** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewEngagementCheckList

`func NewEngagementCheckList(id int32, engagement int32, ) *EngagementCheckList`

NewEngagementCheckList instantiates a new EngagementCheckList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementCheckListWithDefaults

`func NewEngagementCheckListWithDefaults() *EngagementCheckList`

NewEngagementCheckListWithDefaults instantiates a new EngagementCheckList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EngagementCheckList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngagementCheckList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngagementCheckList) SetId(v int32)`

SetId sets Id field to given value.


### GetSessionManagement

`func (o *EngagementCheckList) GetSessionManagement() string`

GetSessionManagement returns the SessionManagement field if non-nil, zero value otherwise.

### GetSessionManagementOk

`func (o *EngagementCheckList) GetSessionManagementOk() (*string, bool)`

GetSessionManagementOk returns a tuple with the SessionManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionManagement

`func (o *EngagementCheckList) SetSessionManagement(v string)`

SetSessionManagement sets SessionManagement field to given value.

### HasSessionManagement

`func (o *EngagementCheckList) HasSessionManagement() bool`

HasSessionManagement returns a boolean if a field has been set.

### GetEncryptionCrypto

`func (o *EngagementCheckList) GetEncryptionCrypto() string`

GetEncryptionCrypto returns the EncryptionCrypto field if non-nil, zero value otherwise.

### GetEncryptionCryptoOk

`func (o *EngagementCheckList) GetEncryptionCryptoOk() (*string, bool)`

GetEncryptionCryptoOk returns a tuple with the EncryptionCrypto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionCrypto

`func (o *EngagementCheckList) SetEncryptionCrypto(v string)`

SetEncryptionCrypto sets EncryptionCrypto field to given value.

### HasEncryptionCrypto

`func (o *EngagementCheckList) HasEncryptionCrypto() bool`

HasEncryptionCrypto returns a boolean if a field has been set.

### GetConfigurationManagement

`func (o *EngagementCheckList) GetConfigurationManagement() string`

GetConfigurationManagement returns the ConfigurationManagement field if non-nil, zero value otherwise.

### GetConfigurationManagementOk

`func (o *EngagementCheckList) GetConfigurationManagementOk() (*string, bool)`

GetConfigurationManagementOk returns a tuple with the ConfigurationManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationManagement

`func (o *EngagementCheckList) SetConfigurationManagement(v string)`

SetConfigurationManagement sets ConfigurationManagement field to given value.

### HasConfigurationManagement

`func (o *EngagementCheckList) HasConfigurationManagement() bool`

HasConfigurationManagement returns a boolean if a field has been set.

### GetAuthentication

`func (o *EngagementCheckList) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *EngagementCheckList) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *EngagementCheckList) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *EngagementCheckList) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAuthorizationAndAccessControl

`func (o *EngagementCheckList) GetAuthorizationAndAccessControl() string`

GetAuthorizationAndAccessControl returns the AuthorizationAndAccessControl field if non-nil, zero value otherwise.

### GetAuthorizationAndAccessControlOk

`func (o *EngagementCheckList) GetAuthorizationAndAccessControlOk() (*string, bool)`

GetAuthorizationAndAccessControlOk returns a tuple with the AuthorizationAndAccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAndAccessControl

`func (o *EngagementCheckList) SetAuthorizationAndAccessControl(v string)`

SetAuthorizationAndAccessControl sets AuthorizationAndAccessControl field to given value.

### HasAuthorizationAndAccessControl

`func (o *EngagementCheckList) HasAuthorizationAndAccessControl() bool`

HasAuthorizationAndAccessControl returns a boolean if a field has been set.

### GetDataInputSanitizationValidation

`func (o *EngagementCheckList) GetDataInputSanitizationValidation() string`

GetDataInputSanitizationValidation returns the DataInputSanitizationValidation field if non-nil, zero value otherwise.

### GetDataInputSanitizationValidationOk

`func (o *EngagementCheckList) GetDataInputSanitizationValidationOk() (*string, bool)`

GetDataInputSanitizationValidationOk returns a tuple with the DataInputSanitizationValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInputSanitizationValidation

`func (o *EngagementCheckList) SetDataInputSanitizationValidation(v string)`

SetDataInputSanitizationValidation sets DataInputSanitizationValidation field to given value.

### HasDataInputSanitizationValidation

`func (o *EngagementCheckList) HasDataInputSanitizationValidation() bool`

HasDataInputSanitizationValidation returns a boolean if a field has been set.

### GetSensitiveData

`func (o *EngagementCheckList) GetSensitiveData() string`

GetSensitiveData returns the SensitiveData field if non-nil, zero value otherwise.

### GetSensitiveDataOk

`func (o *EngagementCheckList) GetSensitiveDataOk() (*string, bool)`

GetSensitiveDataOk returns a tuple with the SensitiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveData

`func (o *EngagementCheckList) SetSensitiveData(v string)`

SetSensitiveData sets SensitiveData field to given value.

### HasSensitiveData

`func (o *EngagementCheckList) HasSensitiveData() bool`

HasSensitiveData returns a boolean if a field has been set.

### GetOther

`func (o *EngagementCheckList) GetOther() string`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *EngagementCheckList) GetOtherOk() (*string, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *EngagementCheckList) SetOther(v string)`

SetOther sets Other field to given value.

### HasOther

`func (o *EngagementCheckList) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetEngagement

`func (o *EngagementCheckList) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *EngagementCheckList) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *EngagementCheckList) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.


### GetSessionIssues

`func (o *EngagementCheckList) GetSessionIssues() []int32`

GetSessionIssues returns the SessionIssues field if non-nil, zero value otherwise.

### GetSessionIssuesOk

`func (o *EngagementCheckList) GetSessionIssuesOk() (*[]int32, bool)`

GetSessionIssuesOk returns a tuple with the SessionIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIssues

`func (o *EngagementCheckList) SetSessionIssues(v []int32)`

SetSessionIssues sets SessionIssues field to given value.

### HasSessionIssues

`func (o *EngagementCheckList) HasSessionIssues() bool`

HasSessionIssues returns a boolean if a field has been set.

### GetCryptoIssues

`func (o *EngagementCheckList) GetCryptoIssues() []int32`

GetCryptoIssues returns the CryptoIssues field if non-nil, zero value otherwise.

### GetCryptoIssuesOk

`func (o *EngagementCheckList) GetCryptoIssuesOk() (*[]int32, bool)`

GetCryptoIssuesOk returns a tuple with the CryptoIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoIssues

`func (o *EngagementCheckList) SetCryptoIssues(v []int32)`

SetCryptoIssues sets CryptoIssues field to given value.

### HasCryptoIssues

`func (o *EngagementCheckList) HasCryptoIssues() bool`

HasCryptoIssues returns a boolean if a field has been set.

### GetConfigIssues

`func (o *EngagementCheckList) GetConfigIssues() []int32`

GetConfigIssues returns the ConfigIssues field if non-nil, zero value otherwise.

### GetConfigIssuesOk

`func (o *EngagementCheckList) GetConfigIssuesOk() (*[]int32, bool)`

GetConfigIssuesOk returns a tuple with the ConfigIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIssues

`func (o *EngagementCheckList) SetConfigIssues(v []int32)`

SetConfigIssues sets ConfigIssues field to given value.

### HasConfigIssues

`func (o *EngagementCheckList) HasConfigIssues() bool`

HasConfigIssues returns a boolean if a field has been set.

### GetAuthIssues

`func (o *EngagementCheckList) GetAuthIssues() []int32`

GetAuthIssues returns the AuthIssues field if non-nil, zero value otherwise.

### GetAuthIssuesOk

`func (o *EngagementCheckList) GetAuthIssuesOk() (*[]int32, bool)`

GetAuthIssuesOk returns a tuple with the AuthIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIssues

`func (o *EngagementCheckList) SetAuthIssues(v []int32)`

SetAuthIssues sets AuthIssues field to given value.

### HasAuthIssues

`func (o *EngagementCheckList) HasAuthIssues() bool`

HasAuthIssues returns a boolean if a field has been set.

### GetAuthorIssues

`func (o *EngagementCheckList) GetAuthorIssues() []int32`

GetAuthorIssues returns the AuthorIssues field if non-nil, zero value otherwise.

### GetAuthorIssuesOk

`func (o *EngagementCheckList) GetAuthorIssuesOk() (*[]int32, bool)`

GetAuthorIssuesOk returns a tuple with the AuthorIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIssues

`func (o *EngagementCheckList) SetAuthorIssues(v []int32)`

SetAuthorIssues sets AuthorIssues field to given value.

### HasAuthorIssues

`func (o *EngagementCheckList) HasAuthorIssues() bool`

HasAuthorIssues returns a boolean if a field has been set.

### GetDataIssues

`func (o *EngagementCheckList) GetDataIssues() []int32`

GetDataIssues returns the DataIssues field if non-nil, zero value otherwise.

### GetDataIssuesOk

`func (o *EngagementCheckList) GetDataIssuesOk() (*[]int32, bool)`

GetDataIssuesOk returns a tuple with the DataIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIssues

`func (o *EngagementCheckList) SetDataIssues(v []int32)`

SetDataIssues sets DataIssues field to given value.

### HasDataIssues

`func (o *EngagementCheckList) HasDataIssues() bool`

HasDataIssues returns a boolean if a field has been set.

### GetSensitiveIssues

`func (o *EngagementCheckList) GetSensitiveIssues() []int32`

GetSensitiveIssues returns the SensitiveIssues field if non-nil, zero value otherwise.

### GetSensitiveIssuesOk

`func (o *EngagementCheckList) GetSensitiveIssuesOk() (*[]int32, bool)`

GetSensitiveIssuesOk returns a tuple with the SensitiveIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveIssues

`func (o *EngagementCheckList) SetSensitiveIssues(v []int32)`

SetSensitiveIssues sets SensitiveIssues field to given value.

### HasSensitiveIssues

`func (o *EngagementCheckList) HasSensitiveIssues() bool`

HasSensitiveIssues returns a boolean if a field has been set.

### GetOtherIssues

`func (o *EngagementCheckList) GetOtherIssues() []int32`

GetOtherIssues returns the OtherIssues field if non-nil, zero value otherwise.

### GetOtherIssuesOk

`func (o *EngagementCheckList) GetOtherIssuesOk() (*[]int32, bool)`

GetOtherIssuesOk returns a tuple with the OtherIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherIssues

`func (o *EngagementCheckList) SetOtherIssues(v []int32)`

SetOtherIssues sets OtherIssues field to given value.

### HasOtherIssues

`func (o *EngagementCheckList) HasOtherIssues() bool`

HasOtherIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


