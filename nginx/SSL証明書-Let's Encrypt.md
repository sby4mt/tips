# Let's Encryptを使ったSSL証明書の発行
https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-20-04-ja
## Cert botのインストール(sanp)推奨
### snapのインストール
```
sudo apt update
sudo apt install snapd
# 更新
sudo snap install core; sudo snap refresh core
```
## certbotインストール
```
sudo snap install --classic certbot
# コマンド実行できるようにシンボリックリンクを張る
sudo ln -s /snap/bin/certbot /usr/bin/certbot
```
## Cert botのインストール(apt)
```
sudo apt install certbot python3-certbot-nginx
```
## SSL証明書の取得
発行のみ
```
sudo certbot certonly --nginx
```
自動設定
```
sudo certbot --nginx -d example.com -d www.example.com
```
certbotを実行するとメールアドレスの入力と利用規約への同意を求められます。

証明書の発行に成功するとHttpsのリダイレクト設定を聞かれます。
どちらかを選択すると設定が更新されNginxがリロードされます。
## Certbotの自動更新の検証
タイマーの確認
```
sudo systemctl status certbot.timer
```
更新プロセスのテスト
```
sudo certbot renew --dry-run
```