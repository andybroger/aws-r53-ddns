# Setup aws-r53-ddns with rpi (arm v7)

Simple Setup for automatic updating your Record with your current IP

## Edit update_ip.sh

If you want, copy the update_ip.sh from the examples/script folder.
Edit the update_ip.sh with your settings

## Download the script

```sh
curl https://github.com/andybroger/aws-r53-ddns/raw/master/examples/rpi/aws-r53-ddns/aws-r53-ddns -O
```

### to build from source

```
GOOS=linux GOARCH=arm GOARM=7 go build
```
