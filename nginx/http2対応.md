# HTTP2対応
- ssl化は必須。
- 以下を参考にconfを編集。
```
listen 443 ssl http2;
ssl_certificate 証明書パス;
ssl_certificate_key キーパス;
...
```