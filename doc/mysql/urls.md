# Table Schema: urls

| Column     | Type           | Default Value | Nullable | Character Set | Collation          | Privileges                         | Extra          | Comments |
| ---------- | -------------- | ------------- | -------- | ------------- | ------------------ | ---------------------------------- | -------------- | -------- |
| id         | uint(11)       |               | NO       |               |                    | select, insert, update, references | auto_increment | URL UUID |
| tiny       | varchar(8)     |               | NO       | utf8mb4       | utf8mb4_general_ci | select, insert, update, references |                | 短網址 |
| origin     | varchar(220)   |               | NO       | utf8mb4       | utf8mb4_general_ci | select, insert, update, references |                | 原始網址 |
| created_at | datetime       |               | NO       |               |                    | select, insert, update, references |                | 短網址產生時間 |
| expires_at | datetime       |               | NO       |               |                    | select, insert, update, references |                | 短網址失效時間 |

## tiny encoding format

```
+---+---+---+---+---+---+
| 0 | 1 | 2 | 3 | 4 | 5 |
+---+---+---+---+---+---+
|     mermer3 hash      |
+---+---+---+---+---+---+
```
