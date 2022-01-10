# Setup - Ubuntu
https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-20-04-ja
## Nginxのインストール。
```
sudo apt update
sudo apt install nginx
```
## ファイアウォールの調整
```
sudo ufw app list
```
出力
```
Available applications:
  Nginx Full - 80ポート、443ポートの両方を開きます
  Nginx HTTP - 80ポートのみ開きます
  Nginx HTTPS - 443ポートのみ開きます
  OpenSSH
```
設定
```
sudo ufw allow 'Nginx HTTP'
```
変更の確認
```
sudo ufw status
```
### Webサーバーを確認する
インストール後に起動されてWebサーバーは稼働状態。
```
systemctl status nginx
```
出力
```
● nginx.service - A high performance web server and a reverse proxy server
     Loaded: loaded (/lib/systemd/system/nginx.service; enabled; vendor preset: enabled)
     Active: active (running) since Mon 2022-01-10 21:02:25 JST; 12min ago
       Docs: man:nginx(8)
   Main PID: 539071 (nginx)
      Tasks: 3 (limit: 1071)
     Memory: 7.0M
     CGroup: /system.slice/nginx.service
             ├─539071 nginx: master process /usr/sbin/nginx -g daemon on; master_process on;
             ├─539072 nginx: worker process
             └─539073 nginx: worker process
```
パブリックIPにアクセスするとNginxのデフォルトページが表示されます。
もしパブリックIPが分からない場合は、icanhazip.comツールを使用して見つけられる。
```
curl -4 icanhazip.com
```
## Nginxのプロセス管理
停止
```
sudo systemctl stop ngingx
```
停止後に起動する
```
sudo systemctl start nginx
```
停止後に再開する
```
sudo systemctl restart nginx
```
構成変更時に接続を切らずにリロードする
```
sudo systemctl reload nginx
```
サービスの有効化
```
sudo systemctl enable nginx
```
## サーバーブロックの構成
新しいファイルの作成
```
sudo nano /etc/nginx/sites-available/your_domain
```
/etc/nginx/sites-available/your_domain
```
server {
    listen 80;
    listen [::]:80;

    root /var/www/your_domain/html;
    index index.html index.htm index.nginx-debian.html;

    server_name your_domain www.your_domain;

    location / {
        try_files $uri $uri/ =404;
    }
}
```
Nginxが起動時に読み取るsites-enabledディレクトリにリンクを作成して、ファイルを有効にする
```
sudo ln -s /etc/nginx/sites-available/your_domain /etc/nginx/sites-enabled/
```
ファイルに構文エラーがないかテストする
```
sudo nginx -t
```
エラーがなければNginxをリロードして変更を有効にする
```
sudo systemctl reload nginx
```