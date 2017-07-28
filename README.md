# Go Webhooks Server
###### Designed to be a lightweight and easily configurable and serve multiple Webhooks

The config.json contains the following
* `Secretkey`: random string to obfuscate URL.
* `Port`: Used to define port.
* `hooks`: List of hooks
* `Name`: Arbitrary name for the hook
* `URL`: Last part of URL
* `Command`: Command to be run


The url is generated in the follow format `/${Secretkey}/${hook.URL}`

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
