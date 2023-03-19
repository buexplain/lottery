```bash
nohup /home/netsvr-linux-amd64.bin -config /home/configs/netsvr.toml 1>/home/log/netsvr-stdout.log 2>/home/log/netsvr-stderr.log &
nohup /home/lottery-linux-amd64.bin -config /home/configs/lottery.toml 1>/home/log/lottery-stdout.log 2>/home/log/lottery-stderr.log &
```
