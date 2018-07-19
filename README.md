# aws-r53-ddns

Cross-Compilable Script to update a A Record in a Route 53 Hosted Zone

If you leave -ip param empty, it will grab your current public IP!

usage:

```shell
AWS_REGION=<aws_region> AWS_ACCESS_KEY_ID=<aws_key_id> AWS_SECRET_ACCESS_KEY=<secret_key> aws-r53-ddns -d <record> -z <zoneid>
```

## Ubiquit edgerouter-X SFP

See /docs/ubnt-er-x for a precompiled script for the edgerouter series, checkout the README.md
