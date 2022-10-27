# Database Schema

## URL

| Column     | Type           | Comments |
| ---------- | -------------- | -------- |
| id         | uint(11)       | UUID |
| tiny       | varchar(8)     | 短網址 |
| origin     | varchar(220)   | 原始網址 |
| created_at | datetime       | 短網址產生時間 |
| expires_at | datetime       | 短網址失效時間 |

**tiny encoding format**

```
+---+---+---+---+---+---+
| 0 | 1 | 2 | 3 | 4 | 5 |
+---+---+---+---+---+---+
|     mermer3 hash      |
+---+---+---+---+---+---+
```
