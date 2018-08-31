#!/bin/bash
#Update Script for Updating Multiple Records

# Just copy this line to update multiple records
AWS_REGION=<region> AWS_ACCESS_KEY_ID=<key_id> AWS_SECRET_ACCESS_KEY=<secret_key> /config/scripts/aws-r53-ddns/aws-r53-ddns -d <record> -z <zoneid>