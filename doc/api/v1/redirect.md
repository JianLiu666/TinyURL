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

    autonumber 1
    client ->> server: [GET] /api/v1/{tiny_uuid}
    server ->> mysql: 取得短網址對應的原始網址<br>table: urls

    rect rgb(242, 238, 229)
    alt
        Note over server: origin url not found
        server ->> client: response 400: Bad Request
    else
        autonumber 3
        Note over server: origin url exists
        server ->> client: response 302: Found
    end
    end
```