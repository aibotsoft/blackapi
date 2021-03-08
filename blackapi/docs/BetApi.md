# \BetApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetById**](BetApi.md#BetById) | **Get** /v1/orders/{order_id}/ | 
[**BetList**](BetApi.md#BetList) | **Get** /v1/orders/ | 
[**BetLog**](BetApi.md#BetLog) | **Get** /v1/orders/{order_id}/log/ | 
[**BetSlip**](BetApi.md#BetSlip) | **Post** /v1/betslips/ | 
[**PlaceBet**](BetApi.md#PlaceBet) | **Post** /v1/orders/ | 
[**RefreshBetSlip**](BetApi.md#RefreshBetSlip) | **Post** /v1/betslips/{betslip_id}/refresh/ | 



## BetById

> PlaceBetResponse BetById(ctx, orderId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := int64(56) // int64 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetApi.BetById(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetApi.BetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetById`: PlaceBetResponse
    fmt.Fprintf(os.Stdout, "Response from `BetApi.BetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlaceBetResponse**](PlaceBetResponse.md)

### Authorization

[session](../README.md#session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetList

> BetListResponse BetList(ctx).OrderBy(orderBy).PageSize(pageSize).Page(page).DateFrom(dateFrom).Sport(sport).EventId(eventId).Status(status).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderBy := "placement_time desc" // string |  (optional)
    pageSize := int64(56) // int64 |  (optional)
    page := int64(56) // int64 |  (optional)
    dateFrom := "2020-05-25T21:00" // string |  (optional)
    sport := "sport_example" // string |  (optional)
    eventId := "eventId_example" // string |  (optional)
    status := "status_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetApi.BetList(context.Background()).OrderBy(orderBy).PageSize(pageSize).Page(page).DateFrom(dateFrom).Sport(sport).EventId(eventId).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetApi.BetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetList`: BetListResponse
    fmt.Fprintf(os.Stdout, "Response from `BetApi.BetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** |  | 
 **pageSize** | **int64** |  | 
 **page** | **int64** |  | 
 **dateFrom** | **string** |  | 
 **sport** | **string** |  | 
 **eventId** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**BetListResponse**](BetListResponse.md)

### Authorization

[session](../README.md#session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetLog

> BetLogResponse BetLog(ctx, orderId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := int64(56) // int64 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetApi.BetLog(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetApi.BetLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetLog`: BetLogResponse
    fmt.Fprintf(os.Stdout, "Response from `BetApi.BetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BetLogResponse**](BetLogResponse.md)

### Authorization

[session](../README.md#session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetSlip

> BetSlipResponse BetSlip(ctx).BetSlipRequest(betSlipRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    betSlipRequest := *openapiclient.NewBetSlipRequest("Sport_example", "EventId_example", "BetType_example", false, false) // BetSlipRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetApi.BetSlip(context.Background()).BetSlipRequest(betSlipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetApi.BetSlip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetSlip`: BetSlipResponse
    fmt.Fprintf(os.Stdout, "Response from `BetApi.BetSlip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betSlipRequest** | [**BetSlipRequest**](BetSlipRequest.md) |  | 

### Return type

[**BetSlipResponse**](BetSlipResponse.md)

### Authorization

[session](../README.md#session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceBet

> PlaceBetResponse PlaceBet(ctx).PlaceBetRequest(placeBetRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    placeBetRequest := *openapiclient.NewPlaceBetRequest("BetslipId_example", float64(123), []interface{}{interface{}(123)}, false, false, "RequestUuid_example", int64(123)) // PlaceBetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetApi.PlaceBet(context.Background()).PlaceBetRequest(placeBetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetApi.PlaceBet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceBet`: PlaceBetResponse
    fmt.Fprintf(os.Stdout, "Response from `BetApi.PlaceBet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaceBetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placeBetRequest** | [**PlaceBetRequest**](PlaceBetRequest.md) |  | 

### Return type

[**PlaceBetResponse**](PlaceBetResponse.md)

### Authorization

[session](../README.md#session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshBetSlip

> RefreshBetSlipResponse RefreshBetSlip(ctx, betslipId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    betslipId := "betslipId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetApi.RefreshBetSlip(context.Background(), betslipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetApi.RefreshBetSlip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshBetSlip`: RefreshBetSlipResponse
    fmt.Fprintf(os.Stdout, "Response from `BetApi.RefreshBetSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**betslipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshBetSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefreshBetSlipResponse**](RefreshBetSlipResponse.md)

### Authorization

[session](../README.md#session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

