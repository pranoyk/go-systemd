## go-systemd

- this application runs a binary of another go application as a systemd process
- in this repo we have a binary `main` which has a urlshortner code
- the `main` binary will be used to run the systemd process


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

