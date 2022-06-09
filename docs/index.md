# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Key.proto](#Key-proto)
    - [AddKeyRequest](#com-example-AddKeyRequest)
    - [AddingKeyStatus](#com-example-AddingKeyStatus)
    - [GetKeyRequest](#com-example-GetKeyRequest)
    - [GetPageRequest](#com-example-GetPageRequest)
    - [KeyPageResponse](#com-example-KeyPageResponse)
    - [ListKeysRequest](#com-example-ListKeysRequest)
    - [ListKeysResponse](#com-example-ListKeysResponse)
    - [V1RfidDto](#com-example-V1RfidDto)
  
    - [KeyService](#com-example-KeyService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="Key-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## Key.proto



<a name="com-example-AddKeyRequest"></a>

### AddKeyRequest
Request message for AddKey method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [V1RfidDto](#com-example-V1RfidDto) |  | The key resource to create. |






<a name="com-example-AddingKeyStatus"></a>

### AddingKeyStatus
Represents the status of key adding


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Unique key ID. |
| status | [bool](#bool) |  | Status of operation |






<a name="com-example-GetKeyRequest"></a>

### GetKeyRequest
Request message for GetKey, DeleteKey methods.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Unique key ID |






<a name="com-example-GetPageRequest"></a>

### GetPageRequest
Request message for GetPage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_number | [int32](#int32) |  | Page number to get |
| page_size | [int32](#int32) |  | Wanted keys count for page |






<a name="com-example-KeyPageResponse"></a>

### KeyPageResponse
Response for GetPage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keys | [V1RfidDto](#com-example-V1RfidDto) | repeated | List of keys for requested page |
| count | [int32](#int32) |  | Result Keys count for requested page |
| total_count | [int32](#int32) |  | Total keys count |
| page_number | [int32](#int32) |  | Requested page number to get |
| page_size | [int32](#int32) |  | Wanted keys count for page |
| is_first | [bool](#bool) |  | First page flag |
| is_last | [bool](#bool) |  | Last page flag |






<a name="com-example-ListKeysRequest"></a>

### ListKeysRequest
Request message for GetPageContinued


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_size | [int32](#int32) |  | The maximum number of results |
| page_token | [string](#string) |  | Specific page of the list results. The next_page_token value returned from a previous List request |






<a name="com-example-ListKeysResponse"></a>

### ListKeysResponse
Response for GetPageContinued


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keys | [V1RfidDto](#com-example-V1RfidDto) | repeated | Current page key list |
| next_page_token | [string](#string) |  | Represents the pagination token to retrieve the next page of results. If the value is &#34;&#34;, it means no further results for the request |






<a name="com-example-V1RfidDto"></a>

### V1RfidDto
Represents the key


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Unique key ID |
| key | [string](#string) |  | Key body |





 

 

 


<a name="com-example-KeyService"></a>

### KeyService
Service for keys exchanging

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddKey | [AddKeyRequest](#com-example-AddKeyRequest) | [V1RfidDto](#com-example-V1RfidDto) | Adding a new key |
| GetKey | [GetKeyRequest](#com-example-GetKeyRequest) | [V1RfidDto](#com-example-V1RfidDto) | Get key by ID |
| DeleteKey | [GetKeyRequest](#com-example-GetKeyRequest) | [V1RfidDto](#com-example-V1RfidDto) | Delete key by ID |
| StreamAddKey | [AddKeyRequest](#com-example-AddKeyRequest) stream | [AddingKeyStatus](#com-example-AddingKeyStatus) stream | Used to bidirectional stream adding. |
| GetPage | [GetPageRequest](#com-example-GetPageRequest) | [KeyPageResponse](#com-example-KeyPageResponse) | Get key list (with explicit page number) |
| GetPageContinued | [ListKeysRequest](#com-example-ListKeysRequest) | [ListKeysResponse](#com-example-ListKeysResponse) | Get continued key list by token (another GetPage variant) |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

