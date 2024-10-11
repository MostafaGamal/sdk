# github.com/MostafaGamal/sdk

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/MostafaGamal/sdk* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/mostafa-gamal/sdk&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/fintech-galaxy/fintech-galaxy). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Personal Finance Management APIs: APIs for managing business finances, including accounts transaction aggregation, budgets, expenses, and reports.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Special Types](#special-types)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/MostafaGamal/sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"github.com/MostafaGamal/sdk"
	"github.com/MostafaGamal/sdk/models/operations"
	"log"
)

func main() {
	s := sdk.New()

	ctx := context.Background()
	res, err := s.AccessToken.GetAccessToken(ctx, &operations.GetAccessTokenRequestBody{
		ClientID:     "demo",
		Code:         "42a7913c-62b4-40cc-a1aa-66286c56a7d5.9dacc7ad-6ca9-4f1b-badb-bbb55ec67323.b913eb22-9a40-464f-bc5a-0c17589214c2",
		RedirectURI:  "https://localhost:8080",
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccessToken](docs/sdks/accesstoken/README.md)

* [GetAccessToken](docs/sdks/accesstoken/README.md#getaccesstoken) - Get an Access Token

### [Accounts](docs/sdks/accounts/README.md)

* [GetUserAccounts](docs/sdks/accounts/README.md#getuseraccounts) - Get accounts
* [GetUserAccountByID](docs/sdks/accounts/README.md#getuseraccountbyid) - Get account details
* [GetUserAccountBalances](docs/sdks/accounts/README.md#getuseraccountbalances) - Get account balances

### [Connections](docs/sdks/connections/README.md)

* [GetUserConnections](docs/sdks/connections/README.md#getuserconnections) - Get user's ais connections
* [CreateAisConnection](docs/sdks/connections/README.md#createaisconnection) - Create a new bank ais connection
* [DeleteAisConnection](docs/sdks/connections/README.md#deleteaisconnection) - Delete Ais Connection by ID

### [Helpdesk](docs/sdks/helpdesk/README.md)

* [CreateInquiry](docs/sdks/helpdesk/README.md#createinquiry) - Create Helpdesk Ticket

### [Kyc](docs/sdks/kyc/README.md)

* [SubmitKyc](docs/sdks/kyc/README.md#submitkyc) - Submit KYC information
* [GetKycStatus](docs/sdks/kyc/README.md#getkycstatus) - Get KYC verification status


### [Transactions](docs/sdks/transactions/README.md)

* [GetAccountTransactions](docs/sdks/transactions/README.md#getaccounttransactions) - List transactions
* [GetTransactions](docs/sdks/transactions/README.md#gettransactions) - List transactions

### [Users](docs/sdks/users/README.md)

* [GetUserInfo](docs/sdks/users/README.md#getuserinfo) - Endpoint to provides details about current logged in user

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"github.com/MostafaGamal/sdk"
	"github.com/MostafaGamal/sdk/models/operations"
	"github.com/MostafaGamal/sdk/retry"
	"log"
	"models/operations"
)

func main() {
	s := sdk.New()

	ctx := context.Background()
	res, err := s.AccessToken.GetAccessToken(ctx, &operations.GetAccessTokenRequestBody{
		ClientID:     "demo",
		Code:         "42a7913c-62b4-40cc-a1aa-66286c56a7d5.9dacc7ad-6ca9-4f1b-badb-bbb55ec67323.b913eb22-9a40-464f-bc5a-0c17589214c2",
		RedirectURI:  "https://localhost:8080",
		CodeVerifier: "ZefY3-yNqKev8BbgtuudLnGRInI70uSt3-s5rr7x5Cc",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"github.com/MostafaGamal/sdk"
	"github.com/MostafaGamal/sdk/models/operations"
	"github.com/MostafaGamal/sdk/retry"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
	)

	ctx := context.Background()
	res, err := s.AccessToken.GetAccessToken(ctx, &operations.GetAccessTokenRequestBody{
		ClientID:     "demo",
		Code:         "42a7913c-62b4-40cc-a1aa-66286c56a7d5.9dacc7ad-6ca9-4f1b-badb-bbb55ec67323.b913eb22-9a40-464f-bc5a-0c17589214c2",
		RedirectURI:  "https://localhost:8080",
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetUserInfo` function may return the following errors:

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401, 500           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/MostafaGamal/sdk"
	"github.com/MostafaGamal/sdk/models/sdkerrors"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithSecurity("PFM_OPEN_ID"),
	)

	ctx := context.Background()
	res, err := s.Users.GetUserInfo(ctx)
	if err != nil {

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://pfm.thefinx.io/api/v1` | None |
| 1 | `https://pfm.thefinxdev.io/api/v1` | None |

#### Example

```go
package main

import (
	"context"
	"github.com/MostafaGamal/sdk"
	"github.com/MostafaGamal/sdk/models/operations"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithServerIndex(1),
	)

	ctx := context.Background()
	res, err := s.AccessToken.GetAccessToken(ctx, &operations.GetAccessTokenRequestBody{
		ClientID:     "demo",
		Code:         "42a7913c-62b4-40cc-a1aa-66286c56a7d5.9dacc7ad-6ca9-4f1b-badb-bbb55ec67323.b913eb22-9a40-464f-bc5a-0c17589214c2",
		RedirectURI:  "https://localhost:8080",
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/MostafaGamal/sdk"
	"github.com/MostafaGamal/sdk/models/operations"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithServerURL("https://pfm.thefinx.io/api/v1"),
	)

	ctx := context.Background()
	res, err := s.AccessToken.GetAccessToken(ctx, &operations.GetAccessTokenRequestBody{
		ClientID:     "demo",
		Code:         "42a7913c-62b4-40cc-a1aa-66286c56a7d5.9dacc7ad-6ca9-4f1b-badb-bbb55ec67323.b913eb22-9a40-464f-bc5a-0c17589214c2",
		RedirectURI:  "https://localhost:8080",
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

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
```go
package main

import (
	"context"
	"github.com/MostafaGamal/sdk"
	"github.com/MostafaGamal/sdk/models/operations"
	"log"
)

func main() {
	s := sdk.New()

	ctx := context.Background()
	res, err := s.AccessToken.GetAccessToken(ctx, &operations.GetAccessTokenRequestBody{
		ClientID:     "demo",
		Code:         "42a7913c-62b4-40cc-a1aa-66286c56a7d5.9dacc7ad-6ca9-4f1b-badb-bbb55ec67323.b913eb22-9a40-464f-bc5a-0c17589214c2",
		RedirectURI:  "https://localhost:8080",
		CodeVerifier: "ZefY3-yNqKev8BbgtuudLnGRInI70uSt3-s5rr7x5Cc",
	}, operations.WithServerURL("https://auth.thefinxdev.io"))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name                     | Type                     | Scheme                   | Environment Variable     |
| ------------------------ | ------------------------ | ------------------------ | ------------------------ |
| `OpenID`                 | openIdConnect            | OpenID Connect Discovery | `PFM_OPEN_ID`            |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/MostafaGamal/sdk"
	"github.com/MostafaGamal/sdk/models/operations"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithSecurity("PFM_OPEN_ID"),
	)

	ctx := context.Background()
	res, err := s.AccessToken.GetAccessToken(ctx, &operations.GetAccessTokenRequestBody{
		ClientID:     "demo",
		Code:         "42a7913c-62b4-40cc-a1aa-66286c56a7d5.9dacc7ad-6ca9-4f1b-badb-bbb55ec67323.b913eb22-9a40-464f-bc5a-0c17589214c2",
		RedirectURI:  "https://localhost:8080",
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/mostafa-gamal/sdk&utm_campaign=go)
