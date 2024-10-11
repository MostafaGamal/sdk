# Transactions
(*Transactions*)

## Overview

Endpoints for managing financial transactions

### Available Operations

* [ListForAccount](#listforaccount) - List transactions
* [List](#list) - List transactions

## ListForAccount

Retrieve a list of transactions for the authenticated user.

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
    res, err := s.Transactions.ListForAccount(ctx, 1, 1, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `connectionID`                                                                 | *int64*                                                                        | :heavy_check_mark:                                                             | N/A                                                                            | 1                                                                              |
| `accountID`                                                                    | *int64*                                                                        | :heavy_check_mark:                                                             | The `id` of the account for which to retrieve the account balance information. | 1                                                                              |
| `limit`                                                                        | **int64*                                                                       | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |
| `offset`                                                                       | **string*                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.GetAccountTransactionsResponse](../../models/operations/getaccounttransactionsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.Error      | 401, 500             | application/json     |
| sdkerrors.Validation | 422                  | application/json     |
| sdkerrors.SDKError   | 4XX, 5XX             | \*/\*                |

## List

Retrieve a list of transactions for the authenticated user.

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
    res, err := s.Transactions.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `offset`                                                 | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTransactionsResponse](../../models/operations/gettransactionsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.Error      | 401, 500             | application/json     |
| sdkerrors.Validation | 422                  | application/json     |
| sdkerrors.SDKError   | 4XX, 5XX             | \*/\*                |