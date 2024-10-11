# Users
(*Users*)

## Overview

Endpoints related to user management

### Available Operations

* [GetUserInfo](#getuserinfo) - Endpoint to provides details about current logged in user

## GetUserInfo

Provides details about current logged in user

### Example Usage

```go
package main

import(
	"github.com/MostafaGamal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity("PFM_OPEN_ID"),
    )

    ctx := context.Background()
    res, err := s.Users.GetUserInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetUserInfoResponse](../../models/operations/getuserinforesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401, 500           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |