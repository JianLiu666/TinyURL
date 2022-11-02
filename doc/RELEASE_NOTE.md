# RELEASE NOTE

- [RELEASE NOTE](#release-note)
  - [v0.2.0 continuing ...](#v020-continuing-)
  - [v0.1.0 at 2022/10/27](#v010-at-20221027)

---

## v0.2.0 continuing ...

- `fix` 修正 `api/v1/create` 業務邏輯, 使其符合使用情境
- `fix` 修正 `api/v1/{tiny_url}` 業務邏輯, 使其符合使用情境
- `optimize` 調整 `api/v1/create` 執行效能
- `optimize` 調整 `api/v1/{tiny_url}` 執行效能
- `feat` 增加壓力測試情境
- `feat` 加入 Redis 轉移 MySQL 負載
- `feat` 加入效能監控與資料視覺化工具: Prometheus & Grafana
- `feat` 加入日誌管理工具: Graylog

---

## v0.1.0 at 2022/10/27

- `feat` 完成基礎建設
- `feat` 完成壓力測試框架
- `feat` 完成整合測試框架
- `feat` 透過 MySQL 處理所有請求
