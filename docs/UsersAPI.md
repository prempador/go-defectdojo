# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreate**](UsersAPI.md#UsersCreate) | **Post** /api/v2/users/ | 
[**UsersDeletePreviewList**](UsersAPI.md#UsersDeletePreviewList) | **Get** /api/v2/users/{id}/delete_preview/ | 
[**UsersDestroy**](UsersAPI.md#UsersDestroy) | **Delete** /api/v2/users/{id}/ | 
[**UsersList**](UsersAPI.md#UsersList) | **Get** /api/v2/users/ | 
[**UsersPartialUpdate**](UsersAPI.md#UsersPartialUpdate) | **Patch** /api/v2/users/{id}/ | 
[**UsersRetrieve**](UsersAPI.md#UsersRetrieve) | **Get** /api/v2/users/{id}/ | 
[**UsersUpdate**](UsersAPI.md#UsersUpdate) | **Put** /api/v2/users/{id}/ | 



## UsersCreate

> User UsersCreate(ctx).UserRequest(userRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	userRequest := *openapiclient.NewUserRequest("Username_example", "Email_example") // UserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersCreate(context.Background()).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersCreate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRequest** | [**UserRequest**](UserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeletePreviewList

> PaginatedDeletePreviewList UsersDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this user.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeletePreviewListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedDeletePreviewList**](PaginatedDeletePreviewList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDestroy

> UsersDestroy(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersList

> PaginatedUserList UsersList(ctx).DateJoinedAfter(dateJoinedAfter).DateJoinedBefore(dateJoinedBefore).Email(email).FirstName(firstName).Id(id).IsActive(isActive).IsSuperuser(isSuperuser).LastLoginAfter(lastLoginAfter).LastLoginBefore(lastLoginBefore).LastName(lastName).Limit(limit).O(o).Offset(offset).Username(username).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	dateJoinedAfter := time.Now() // string |  (optional)
	dateJoinedBefore := time.Now() // string |  (optional)
	email := "email_example" // string |  (optional)
	firstName := "firstName_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	isActive := true // bool |  (optional)
	isSuperuser := true // bool |  (optional)
	lastLoginAfter := time.Now() // string |  (optional)
	lastLoginBefore := time.Now() // string |  (optional)
	lastName := "lastName_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	o := []string{"O_example"} // []string | Ordering  * `username` - Username * `-username` - Username (descending) * `last_name` - Last name * `-last_name` - Last name (descending) * `first_name` - First name * `-first_name` - First name (descending) * `email` - Email * `-email` - Email (descending) * `is_active` - Is active * `-is_active` - Is active (descending) * `is_superuser` - Is superuser * `-is_superuser` - Is superuser (descending) * `date_joined` - Date joined * `-date_joined` - Date joined (descending) * `last_login` - Last login * `-last_login` - Last login (descending) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	username := "username_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersList(context.Background()).DateJoinedAfter(dateJoinedAfter).DateJoinedBefore(dateJoinedBefore).Email(email).FirstName(firstName).Id(id).IsActive(isActive).IsSuperuser(isSuperuser).LastLoginAfter(lastLoginAfter).LastLoginBefore(lastLoginBefore).LastName(lastName).Limit(limit).O(o).Offset(offset).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersList`: PaginatedUserList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateJoinedAfter** | **string** |  | 
 **dateJoinedBefore** | **string** |  | 
 **email** | **string** |  | 
 **firstName** | **string** |  | 
 **id** | **int32** |  | 
 **isActive** | **bool** |  | 
 **isSuperuser** | **bool** |  | 
 **lastLoginAfter** | **string** |  | 
 **lastLoginBefore** | **string** |  | 
 **lastName** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **o** | **[]string** | Ordering  * &#x60;username&#x60; - Username * &#x60;-username&#x60; - Username (descending) * &#x60;last_name&#x60; - Last name * &#x60;-last_name&#x60; - Last name (descending) * &#x60;first_name&#x60; - First name * &#x60;-first_name&#x60; - First name (descending) * &#x60;email&#x60; - Email * &#x60;-email&#x60; - Email (descending) * &#x60;is_active&#x60; - Is active * &#x60;-is_active&#x60; - Is active (descending) * &#x60;is_superuser&#x60; - Is superuser * &#x60;-is_superuser&#x60; - Is superuser (descending) * &#x60;date_joined&#x60; - Date joined * &#x60;-date_joined&#x60; - Date joined (descending) * &#x60;last_login&#x60; - Last login * &#x60;-last_login&#x60; - Last login (descending) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **username** | **string** |  | 

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPartialUpdate

> User UsersPartialUpdate(ctx, id).PatchedUserRequest(patchedUserRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this user.
	patchedUserRequest := *openapiclient.NewPatchedUserRequest() // PatchedUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPartialUpdate(context.Background(), id).PatchedUserRequest(patchedUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPartialUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUserRequest** | [**PatchedUserRequest**](PatchedUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRetrieve

> User UsersRetrieve(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersRetrieve`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdate

> User UsersUpdate(ctx, id).UserRequest(userRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this user.
	userRequest := *openapiclient.NewUserRequest("Username_example", "Email_example") // UserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUpdate(context.Background(), id).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRequest** | [**UserRequest**](UserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

