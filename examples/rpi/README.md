# Setup aws-r53-ddns with rpi (arm v7)

Simple Setup for automatic updating your Record with your current IP

## Edit update_ip.sh

If you want, copy the update_ip.sh from the examples/script folder.
Edit the update_ip.sh with your settings.

## Download source and compile


```
git clone https://github.com/andybroger/aws-r53-ddns
cd aws-r53-ddns
GOOS=linux GOARCH=arm GOARM=7 go build
```

## cronjob

This cronjob runs every 5 minutes, checks the current ip and if it differs from the current record, change the record defined, use`crontab -e` to add the cronjob:

```shell
*/5 * * * * AWS_REGION=<region> AWS_ACCESS_KEY_ID=<key_id> AWS_SECRET_ACCESS_KEY=<secret_key> aws-r53-ddns -d <record> -z <zoneid> >/dev/null 2>&1
```