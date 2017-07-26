# Go Webhooks Server
###### Design to be a lightweight and easily configurable and serve multiple Webhooks

The config.json contains a `Secretkey` field which is a random string to obfuscate all webhook links and a list of `hooks`. Each `hook` contains a `Name`, `URL`, and `Command`. The url is generated in the follow format `/${Secretkey}/${hook.URL}`

#### config.json
```json
{
  "Secretkey": "as0fn912g9ag9-9bq2g9afw0",
  "hooks": [
    {
      "Name": "docker",
      "Url": "docker-v1",
      "Command": "docker run -p80:8080"
    },
    {
      "Name": "restart",
      "Url": "restart-v1",
      "Command": "systemctl reboot -i"
    }
  ]
}
```
