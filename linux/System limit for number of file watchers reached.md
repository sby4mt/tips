# 現在のinotifyの上限値を調べる
```
$ cat /proc/sys/fs/inotify/max_user_watches
8192
```
# inotifyの上限値を永続的に増加させる
```
echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
```