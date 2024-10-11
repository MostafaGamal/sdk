# AccountBalanceOptional

Optional data associated with the balance.<br /><br /><strong>Note:</strong> `optional` parameter will always be present, but the array may be empty, or it may contain all or some of the parameters listed below.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `BalanceType`                                            | **string*                                                | :heavy_minus_sign:                                       | Type of the balance.                                     |
| `SourceDate`                                             | [*types.Date](../../types/date.md)                       | :heavy_minus_sign:                                       | The last date when the bank updated the balance.         |
| `SourceTime`                                             | **string*                                                | :heavy_minus_sign:                                       | The last time when the bank updated the balance.         |
| `AvailableAmount`                                        | **float32*                                               | :heavy_minus_sign:                                       | The available amount of funds in the account's currency. |
| `BlockedAmount`                                          | **float32*                                               | :heavy_minus_sign:                                       | The blocked amount of funds in the account's currency.   |
| `CreditLimit`                                            | **float32*                                               | :heavy_minus_sign:                                       | The maximum amount of funds that the user can spend.     |
| `ClosingBalance`                                         | **float32*                                               | :heavy_minus_sign:                                       | Closing balance.                                         |
| `OpeningBalance`                                         | **float32*                                               | :heavy_minus_sign:                                       | Opening balance.                                         |