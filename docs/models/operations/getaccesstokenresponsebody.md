# GetAccessTokenResponseBody

Created


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AccessToken`                                                 | **string*                                                     | :heavy_minus_sign:                                            | Access Token to be used in requests to FINX PFM APIs.         | eyJhbGciOiJIUzI1.SflKxwRJSMeKKF2QT4fwpMeJf36P.Ok6yJV_adQssw5c |
| `Scope`                                                       | **string*                                                     | :heavy_minus_sign:                                            | Application scopes on requested API, separated by spaces.     | ais                                                           |
| `ExpiresIn`                                                   | **int64*                                                      | :heavy_minus_sign:                                            | Number of seconds after which an Access Token will expire.    | 86400                                                         |
| `TokenType`                                                   | **string*                                                     | :heavy_minus_sign:                                            | Your Access Token type.                                       | Bearer                                                        |