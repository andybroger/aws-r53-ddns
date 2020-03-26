# aws-r53-ddns

Cross-Compilable Script to update a A Record in a Route 53 Hosted Zone

If you leave -ip param empty, it will grab your current public IP!

usage:

```shell
$ AWS_REGION=<aws_region> AWS_ACCESS_KEY_ID=<aws_key_id> AWS_SECRET_ACCESS_KEY=<secret_key> aws-r53-ddns -d <record> -z <zoneid>

$ AWS_REGION=eu-central-1 AWS_ACCESS_KEY_ID=XXXXX AWS_SECRET_ACCESS_KEY=XXXXX aws-r53-ddns -d test.example.com -z XXXXX
$ 2020/03/26 10:07:36 updating record test.example.com to new value 1.2.3.5
```
