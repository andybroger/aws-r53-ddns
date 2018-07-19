# aws-r53-ddns

Cross-Compilable Script to update a A Record in a Route 53 Hosted Zone

If you leave -ip param empty, it will grab your current public IP!

usage:

```shell
AWS_REGION=<aws_region> AWS_ACCESS_KEY_ID=<aws_key_id> AWS_SECRET_ACCESS_KEY=<secret_key> go run main.go -d <record> -z <zoneid>
```

to Build:

```sh
GOOS=linux GOARCH=mipsle go build
GOOS=Darwin GOARCH=i386 go build
```
