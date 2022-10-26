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
            server ->> client: reponse 400: field invalid.
        else
            autonumber 2
            Note over server : parse request body successfully
            server ->> server : 產生短網址 (MurmurHash)
            
            rect rgb(136,186,186)
                Note over client, mysql: perform multiple activities by transaction 
                server ->> mysql : 檢查是否存在相同的短網址<br>table: urls
                rect rgb(242, 238, 229)
                    alt
                        Note over server: 存在相同短網址 AND (短網址為用戶客製化 OR 原始網址相同) 時
                    server ->> client : response 400: alias dunplicated.
                    else
                        Note over server: 僅只有短網址相同時
                        server ->> server : 短網址增加 timestamp 的後綴
                    end
                end
                server ->> mysql : 寫入短網址 (InsertUpdate)<br>table: urls
            end
            server ->> client: reponse 200: OK
        end
    end
```
