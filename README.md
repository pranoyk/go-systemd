## go-systemd

- this application runs a binary of another go application as a systemd process


### service file

add the below file in `/etc/systemd/system/urlshortner.service` path

```
[Unit]
Description=UrlShortner go service

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=<path to binary>

[Install]
WantedBy=multi-user.target
```

