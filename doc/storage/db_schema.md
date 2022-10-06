# Database Schema

## URL

| Column     | Type           | Comments |
| ---------- | -------------- | -------- |
| tiny       | varchar(11)    | 短網址 |
| origin     | varchar(220)   | 原始網址 |
| created_at | datetime       | 短網址產生時間 |
| expires_at | datetime       | 短網址失效時間 |