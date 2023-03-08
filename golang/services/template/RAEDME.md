# Template
以下を用いた一般的なHTTPサーバーのサンプルテンプレート
- go 1.20.1
- gin (http framework)
- gorm (ORM library)

### development

- サービス立ち上げ

```shell
docker-compose up -d
```

- DBに入る

```shell
docker exec -it template_service_database mysql -u root -p
```

passwordは`root`