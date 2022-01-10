# Let's Encryptを使ったSSL証明書の発行
https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-20-04-ja
## Cert botのインストール
```
sudo apt install certbot python3-certbot-nginx
```
## SSL証明書の取得
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