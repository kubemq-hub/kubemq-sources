# Kubemq sqs source Connector

Kubemq aws-sqs source connector allows services using kubemq server to access aws sqs service.

## Prerequisites
The following required to run the aws-sqs source connector:

- kubemq cluster
- aws account with sqs active service
- kubemq-source deployment

## Configuration

sqs source connector configuration properties:

| Properties Key                 | Required | Description                                                       | Example                     |
|:-------------------------------|:---------|:------------------------------------------------------------------|:----------------------------|
| aws_key                        | yes      | aws key                                                           | aws key supplied by aws         |
| aws_secret_key                 | yes      | aws secret key                                                    | aws secret key supplied by aws  |
| region                         | yes      | region                                                            | aws region                      |
| retries                        | no       | number of retries on send                                         | 1 (default 0)                   |
| token                          | no       | aws token ("default" empty string                                 | "my token"                      |
| queue                          | yes      | queue name                                                        | "my_queue_name"          |
| max_number_of_messages         | no       | max messages receive per call                                     | "1" (default 1)                      |
| visibility_timeout             | no       | max messages receive per call                                     | "1" (default 0)                      |
| pullDelay                      | no       | wait time between calls (milliseconds)                            | "1" (default 60)                      |
 

Example:

```yaml
bindings:
  - name: kubemq-query-aws-sqs
    source:
      kind: source.query
      name: kubemq-query
      properties:
        host: "localhost"
        port: "50000"
        client_id: "kubemq-query-aws-sqs-connector"
        auth_token: ""
        channel: "query.aws.sqs"
        group:   ""
        concurrency: "1"
        auto_reconnect: "true"
        reconnect_interval_seconds: "1"
        max_reconnects: "0"
    target:
      kind: source.aws.sqs
      name: source-aws-sqs
      properties:
        aws_key: "id"
        aws_secret_key: 'json'
        region:  "instance"
        queue : "my_queue"
```

