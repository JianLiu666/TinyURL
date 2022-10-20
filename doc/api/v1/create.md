## 縮網址

**POST** `{domain}/api/v1/create`

### Authorization

- none

### Request 

Request Body

| Field  | Type   | Required | Description |
| ------ | ------ | :------: | ----------- |
| url    | string | Yes      | 原始網址 |
| alias  | string | No       | 指定短網址替代碼 |

### Response

Schema

| Field   | Type   | Description |
| ------- | ------ | ----------- |
| origin  | string | 原始網址 |
| tiny    | string | 短網址 |
| created_at | string | 短網址產生時間 |
| expires_at | string | 短網址失效時間 |

### Flow

```mermaid
sequenceDiagram
    participant client
    participant server
    participant mysql
    
    autonumber 1
    client ->> server: [POST] /api/v1/create
    
    rect rgb(242, 238, 229)
    alt
        Note over server: failed to parse request body
        server ->> client: reponse 400: Bad Request
    else
        autonumber 2
        Note over server : parse request body successfully
        server ->> server : 產生短網址 (MurmurHash)
        
        server ->> mysql : 檢查是否存在相同的短網址<br>table: urls
        server ->> mysql : 寫入短網址 (InsertUpdate)<br>table: urls
        Note over server, mysql : 如果短網址發生碰撞, 加入 timestamp 作為後綴
        
        server ->> client: reponse 200: OK
    end
    end
```
