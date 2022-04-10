# アクセスログの管理
## ログの保管場所設定
/etc/nginx/nginx.conf
```
access_log /var/log/nginx/access.log; # アクセスログの出力先
error_log /var/log/nginx/error.log; # エラーログの出力先
```
## ログフォーマット
デフォルト
```
log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';
```