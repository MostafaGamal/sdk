# Accounts
(*Accounts*)

## Overview

Endpoints related to user accounts

### Available Operations

* [GetUserAccounts](#getuseraccounts) - Get accounts
* [GetUserAccountByID](#getuseraccountbyid) - Get account details
* [GetUserAccountBalances](#getuseraccountbalances) - Get account balances

## GetUserAccounts

Retrieve all the accounts associated with a particular connection.

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
    res, err := s.Accounts.GetUserAccounts(ctx, 1)
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `connectionID`                                           | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      | 1                                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetUserAccountsResponse](../../models/operations/getuseraccountsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.Error      | 401, 500             | application/json     |
| sdkerrors.Validation | 422                  | application/json     |
| sdkerrors.SDKError   | 4XX, 5XX             | \*/\*                |

## GetUserAccountByID

Retrieve the details of a specific user account.

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
    res, err := s.Accounts.GetUserAccountByID(ctx, 1, 1)
    if err != nil {
        log.Fatal(err)
    }
    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `accountID`                                                                    | *int64*                                                                        | :heavy_check_mark:                                                             | The `id` of the account for which to retrieve the account balance information. | 1                                                                              |
| `connectionID`                                                                 | *int64*                                                                        | :heavy_check_mark:                                                             | N/A                                                                            | 1                                                                              |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.GetUserAccountByIDResponse](../../models/operations/getuseraccountbyidresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.Error      | 401, 404, 500        | application/json     |
| sdkerrors.Validation | 422                  | application/json     |
| sdkerrors.SDKError   | 4XX, 5XX             | \*/\*                |

## GetUserAccountBalances

Returns the account balance information for a particular account/connection.

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
    res, err := s.Accounts.GetUserAccountBalances(ctx, operations.GetUserAccountBalancesRequest{
        ConnectionID: 1,
        AccountID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountBalancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetUserAccountBalancesRequest](../../models/operations/getuseraccountbalancesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetUserAccountBalancesResponse](../../models/operations/getuseraccountbalancesresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.Error      | 401, 500             | application/json     |
| sdkerrors.Validation | 422                  | application/json     |
| sdkerrors.SDKError   | 4XX, 5XX             | \*/\*                |