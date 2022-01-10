# 查看局域网IP

## WIN 10

```cmd
for /L %i IN (1,1,254) DO ping -w 2 -n 1 192.168.1.%i
arp -a
```

