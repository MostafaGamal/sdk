<!-- Start SDK Example Usage [usage] -->
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
	res, err := s.AccessToken.Create(ctx, &operations.GetAccessTokenRequestBody{
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