# Setup aws-r53-ddns with ubiquiti edgerouter-x (sfp)

Simple Setup for automatic updating your Record with your current IP

## Edit update_ip.sh

Edit the update_ip.sh with your settings

## Copy files to router

```sh
scp -r aws-r53-ddns/ <user>@<router_ip>:/config/scripts/
```

## Configure Task-Scheduler to run script every hour

```sh
configure
set system task-scheduler task update_ip executable path /config/scripts/aws-r53-ddns/update_ip.sh
set system task-scheduler task update_ip interval 1h
commit
exit
```
