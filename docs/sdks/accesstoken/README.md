# AccessToken
(*AccessToken*)

## Overview

### Available Operations

* [Create](#create) - Get an Access Token

## Create

Get an access token from auth server to authenticate your requests to FINX PFM APIs.

### Example Usage

```go
package main

import(
	"github.com/MostafaGamal/sdk"
	"context"
	"github.com/MostafaGamal/sdk/models/operations"
	"log"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.AccessToken.Create(ctx, &operations.GetAccessTokenRequestBody{
        ClientID: "demo",
        Code: "42a7913c-62b4-40cc-a1aa-66286c56a7d5.9dacc7ad-6ca9-4f1b-badb-bbb55ec67323.b913eb22-9a40-464f-bc5a-0c17589214c2",
        RedirectURI: "https://localhost:8080",
        CodeVerifier: "ZefY3-yNqKev8BbgtuudLnGRInI70uSt3-s5rr7x5Cc",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetAccessTokenRequestBody](../../models/operations/getaccesstokenrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetAccessTokenResponse](../../models/operations/getaccesstokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |