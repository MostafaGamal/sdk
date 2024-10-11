# Error

Unauthorized access to the resource.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ID`                                                                      | *string*                                                                  | :heavy_check_mark:                                                        | Unique identifier for the error                                           |
| `Code`                                                                    | *int64*                                                                   | :heavy_check_mark:                                                        | Error code is to identify the type of error.                              |
| `Message`                                                                 | *string*                                                                  | :heavy_check_mark:                                                        | Friendly error message to show to the user                                |
| `Detail`                                                                  | **string*                                                                 | :heavy_minus_sign:                                                        | Detailed technical error description used by system admins and developers |