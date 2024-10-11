# Connections
(*Connections*)

## Overview

Endpoints related to user connections with the banks

### Available Operations

* [List](#list) - Get user's ais connections
* [Create](#create) - Create a new bank ais connection
* [Delete](#delete) - Delete Ais Connection by ID

## List

Retrieve the list of user ais connections.

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
    res, err := s.Connections.List(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AisConnectionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `limit`                                                              | **int64*                                                             | :heavy_minus_sign:                                                   | N/A                                                                  |
| `offset`                                                             | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `sort`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | Key to Sort by followed by order (- for descending)                  |
| `filter`                                                             | **string*                                                            | :heavy_minus_sign:                                                   | Filter by resources having fields starting with this parameter value |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.GetUserConnectionsResponse](../../models/operations/getuserconnectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400, 401, 500      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create a new user bank ais connection (account).

### Example Usage

```go
package main

import(
	"github.com/MostafaGamal/sdk"
	"context"
	"github.com/MostafaGamal/sdk/types"
	"github.com/MostafaGamal/sdk/models/components"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity("PFM_OPEN_ID"),
    )

    ctx := context.Background()
    res, err := s.Connections.Create(ctx, components.CreateAisConnection{
        ConsentDetails: &components.ConsentDetails{
            FetchStartDate: types.MustNewDateFromString("2020-01-26"),
            ExpiresAt: types.MustTimeFromString("2024-01-26T02:25:34.569+00:00"),
            Scopes: []components.Scopes{
                components.ScopesStandingOrders,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NewAisConnectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.CreateAisConnection](../../models/components/createaisconnection.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateAisConnectionResponse](../../models/operations/createaisconnectionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.Error      | 400, 401, 500        | application/json     |
| sdkerrors.Validation | 422                  | application/json     |
| sdkerrors.SDKError   | 4XX, 5XX             | \*/\*                |

## Delete

Delete Ais Connection by ID

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
    res, err := s.Connections.Delete(ctx, 1)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `connectionID`                                           | *int64*                                                  | :heavy_check_mark:                                       | The `id` of the connection for which will be deleted.    | 1                                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAisConnectionResponse](../../models/operations/deleteaisconnectionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.Error      | 401, 404, 500        | application/json     |
| sdkerrors.Validation | 422                  | application/json     |
| sdkerrors.SDKError   | 4XX, 5XX             | \*/\*                |