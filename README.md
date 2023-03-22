## 部署

```bash
nohup /home/netsvr-linux-amd64.bin -config /home/configs/netsvr.toml 1>/home/log/netsvr-stdout.log 2>/home/log/netsvr-stderr.log &
nohup /home/lottery-linux-amd64.bin -config /home/configs/lottery.toml 1>/home/log/lottery-stdout.log 2>/home/log/lottery-stderr.log &
```

systemctl start lottery.service


使用方式：
```bash
systemctl start lottery.service
systemctl status lottery.service
systemctl stop lottery.service
# 查看某个 Unit 的日志
journalctl -u lottery.service
journalctl -u lottery.service --since today
# 持续打印某个 Unit 的最新日志
journalctl  -u lottery.service  -f
```