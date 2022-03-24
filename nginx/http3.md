# 参考記事
https://qiita.com/keys/items/bf492ef543fde3d8d822
# 環境
AWS EC2上のUbuntu20.04で構築
# 準備
http2とhttp3の設定は[共存可能](https://qiita.com/keys/items/bf492ef543fde3d8d822)だがアンインストールからが楽
アンインストール
```
sudo apt purge nginx nginx-full && sudo apt autoremove --purge
```
QUIC版nginxのビルドに必要なパッケージインストール
```
sudo apt update && sudo apt upgrade
sudo apt install git mercurial cmake make golang libunwind-dev libpcre3-dev zlib1g-dev gcc g++
```
# Boringssl
適当なフォルダを作成して`Google`からOpenSSLのフォークをインストール。
```
# ディレクトリを作る
sudo mkdir /opt/nginx
# /opt.nginxに移動
cd /opt/nginx
# clone
sudo git clone https://boringssl.googlesource.com/boringssl
# インストール
cd boringssl
sudo mkdir build
cd build
sudo cmake ..
sudo make
# テスト
cd ..
go run util/all_tests.go
cd ssl/test/runner
go test
```
# Nginx
```
cd /opt/nginx
# git clone的な
sudo hg clone -b quic https://hg.nginx.org/nginx-quic
```
aptでインストールした時と同じようになるようにインストール  
/opt/nginx/configure.shを作成
```
#!/usr/bin/bash

cd nginx-quic
./auto/configure \
    --with-debug \
    --with-http_v2_module \
    --with-http_v3_module \
    --with-stream \
    --with-stream_quic_module \
    --with-http_ssl_module \
    --with-cc-opt="-I../boringssl/include"   \
    --with-ld-opt="-L../boringssl/build/ssl  \
                   -L../boringssl/build/crypto" \
    --prefix=/etc/nginx \
    --sbin-path=/usr/sbin/nginx \
    --conf-path=/etc/nginx/nginx.conf \
    --error-log-path=/var/log/nginx/error.log \
    --http-log-path=/var/log/nginx/access.log \
    --pid-path=/run/nginx.pid \
    --lock-path=/run/nginx.lock \
    --http-client-body-temp-path=/var/cache/nginx/client_temp \
    --http-proxy-temp-path=/var/cache/nginx/proxy_temp \
    --http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp \
    --http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp \
    --http-scgi-temp-path=/var/cache/nginx/scgi_temp \
    --user=www-data \
    --group=www-data
```
インストール実行
```
# 実行権付与
sudo chmod 744 configure.sh
# configure.sh実行
sudo -s ./configure.sh
cd nginx-quic
# インストール
# ./objs配下にファイルができます
sudo make
# ファイルを各所にコピーします（既に別バージョンがインストールされている場合は注意）
sudo make install
# 確認
sudo -s /usr/sbin/nginx -V
```
ディレクトリ作成
```
sudo mkdir /var/cache/nginx
```
# systemd管理
/lib/systemd/system/nginx.service を作成
```
[Unit]
Description=The NGINX HTTP and reverse proxy server
After=syslog.target network-online.target remote-fs.target nss-lookup.target
Wants=network-online.target

[Service]
Type=forking
#PIDFile=/run/nginx.pid
ExecStartPre=/usr/sbin/nginx -t
ExecStart=/usr/sbin/nginx
ExecReload=/usr/sbin/nginx -s reload
ExecStop=/bin/kill -s QUIT $MAINPID
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```
sytemd
```
sudo systemctl daemon-reload
sudo systemctl start nginx
```
自動起動するようにする
```
sudo systemctl enable nginx
```
# SSL証明書取得

/etc/letsencrypt/live/sbym.tech/fullchain.pem
/etc/letsencrypt/live/sbym.tech/privkey.pem