# Setup aws-r53-ddns with ubiquiti edgerouter-x (sfp)

Simple Setup for automatic updating your Record with your current IP

## Edit update_ip.sh (client)

Edit the update_ip.sh with your settings

## Download source and compile


```
git clone https://github.com/andybroger/aws-r53-ddns
cd aws-r53-ddns
GOOS=linux GOARCH=mipsle go build
```

## Configure Task-Scheduler to run script every hour (router)

```sh
configure
set system task-scheduler task update_ip executable path /config/scripts/aws-r53-ddns/update_ip.sh
set system task-scheduler task update_ip interval 5m
commit
save
exit
```

### to build from source

```
GOOS=linux GOARCH=mipsle go build
```
