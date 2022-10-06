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
            server ->> server : encode origin url
            server ->> mysql : 檢查是否存在相同網址<br>table: urls
            rect rgb(136,186,186)
            alt
            Note over server, mysql : 已經存在相同網址且短網址仍有效
            server ->> client: reponse 400: Bad Request
            else
            autonumber 4
            Note over server, mysql : 網址不存在或短網址已失效
            server ->> mysql : 寫入短網址 (InsertUpdate)<br>table: urls
            server ->> client: reponse 200: OK
            end
            end
        end
    end
```
