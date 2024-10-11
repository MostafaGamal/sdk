# Transactions
(*Transactions*)

## Overview

Endpoints for managing financial transactions

### Available Operations

* [GetAccountTransactions](#getaccounttransactions) - List transactions
* [GetTransactions](#gettransactions) - List transactions

## GetAccountTransactions

Retrieve a list of transactions for the authenticated user.

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
    s := sdk.New(
        sdk.WithSecurity("PFM_OPEN_ID"),
    )

    ctx := context.Background()
    res, err := s.Transactions.GetAccountTransactions(ctx, operations.GetAccountTransactionsRequest{
        ConnectionID: 1,
        AccountID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAccountTransactionsRequest](../../models/operations/getaccounttransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAccountTransactionsResponse](../../models/operations/getaccounttransactionsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.Error      | 401, 500             | application/json     |
| sdkerrors.Validation | 422                  | application/json     |
| sdkerrors.SDKError   | 4XX, 5XX             | \*/\*                |

## GetTransactions

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
    res, err := s.Transactions.GetTransactions(ctx, nil, nil)
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