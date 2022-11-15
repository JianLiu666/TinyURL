## 跳轉網址

**GET** `{domain}/api/v1/{tiny_uuid}`

### Authorization

- none

### Request

Path Prameters

| Parameter  | Description |
| ---------- | ----------- |
| tiny_uuid  | 短網址 UUID |

### Response

- none

### Flow

```mermaid
sequenceDiagram
    participant client
    participant server
    participant mysql
    participant redis

    autonumber 1
    client ->> server: [GET] /api/v1/{tiny_uuid}
    server ->> redis: [GET] 取回短網址對應的原始網址<br>key: tiny:{tiny}

    alt
        Note over server, redis: 短網址命中時
        alt
            Note over client, server: value 有資料時 (i.e. 存在對應的原始網址)
            server ->> client: response 302: Found
        else
            autonumber 3
            Note over client, server: value 無資料時 (緩存短網址阻擋重複請求)
            server ->> client: response 400: Bad Request
        end
    else
        autonumber 4
        Note over server, redis: 短網址不存在時
        server ->> mysql: 取得短網址對應的原始網址<br>table: urls
        alt
            Note over server, mysql: 短網址不存在於 mysql 時
            server ->> redis: [SET] 寫入短網址供往後檢查使用<br>key: tiny:{tiny}, value: "", expired: 1小時
            server ->> client: response 400: Bad Request
        else
            autonumber 5
            Note over server, mysql: 短網址存在於 mysql 時
            server ->> redis: [SET] 寫入短網址供往後檢查使用<br>key: tiny:{tiny}, value: {origin}, expired: 1小時
            server ->> client: response 302: Found
        end
    end
```