# AisConnectionsResponse

List of the user's created connections (accounts)


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Data`                                                                   | [][components.AisConnections](../../models/components/aisconnections.md) | :heavy_minus_sign:                                                       | N/A                                                                      |
| `PerPage`                                                                | **int64*                                                                 | :heavy_minus_sign:                                                       | The maximum number of items that can be returned per page.               |
| `NextCursor`                                                             | **string*                                                                | :heavy_minus_sign:                                                       | Could be used to retrieve the next page.                                 |
| `PrevCursor`                                                             | **string*                                                                | :heavy_minus_sign:                                                       | Could be used to retrieve the previous page.                             |