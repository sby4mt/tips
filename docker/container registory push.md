PATを環境変数に設定。
```
export CR_PAT=YOUR_TOKEN
```
ログイン。
```
echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
```
コンテナイメージをプッシュ。
```
docker push ghcr.io/OWNER/IMAGE_NAME:latest
```
