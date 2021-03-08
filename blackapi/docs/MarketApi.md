# \MarketApi

All URIs are relative to *https://black.betinasia.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents**](MarketApi.md#GetEvents) | **Post** /s/prices/eventdata1/ | 



## GetEvents

> GetEventsResponse GetEvents(ctx).GetEventsRequest(getEventsRequest).Execute()



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
    getEventsRequest := *openapiclient.NewGetEventsRequest(false, false) // GetEventsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.GetEvents(context.Background()).GetEventsRequest(getEventsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: GetEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEventsRequest** | [**GetEventsRequest**](GetEventsRequest.md) |  | 

### Return type

[**GetEventsResponse**](GetEventsResponse.md)

### Authorization

[session](../README.md#session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

