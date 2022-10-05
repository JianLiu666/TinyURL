# Database Schema

## URL

| Column     | Type           | Comments |
| ---------- | -------------- | -------- |
| id         | varchar(16)    | tiny url |
| original   | varchar(65535) | original url |
| md5        | char(32)       | orignial url encoded to md5 |
| created_at | datetime       | tiny url created time |
| expired_at | datetime       | tiny url expreid time |