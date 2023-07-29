## 各種コマンド

### 環境構築

```
make init
```

コンテナ停止後に、コンテナを起動する場合

```
make up
```

### マイグレーション

dry-run(適用される DDL を確認)

```
make migrate-dry-run
```

apply(実際に適用)

```
make migrate-apply
```

### テスト

```
make test
```

### API Schema の確認

```
make gen-swagger
```

URL: http://localhost:8002

## directory structure

### project 全体

```bash
.
├── Makefile
├── README.md
├── app # アプリケーションの実装
├── docker-compose.yml
├── ops # ops周り
└── pkg # ドメインロジックとは関係ない汎用的な処理
```

- application 内

```bash
.
├── application # アプリケーションサービス層
├── cli # cliにて処理する際に利用
├── cmd # アプリケーションのmain.go
├── config # 各種設定値
├── docs #swagger関連
├── domain #domain層
├── domain_service # domainService層（消すかも）
├── go.mod
├── go.sum
├── infrastructure # infrastructre層
└── presentation # presentationの層
```
