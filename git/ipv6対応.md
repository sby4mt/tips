# ipv6使用時の注意
pushが出来なくなった場合、gitへのssh接続をipv4に固定する。  
~/.ssh/configに`AddressFamily inet`を追記する。
```
Host github.com
  Hostname github.com
  Post 443
  AddressFamily inet
```