# Device Farm SKU watcher

Sometimes devices get added or removed. This provides a dumb, but working way to monitor when SKUs go offline.

## IAM

Make a keypair with the following permissions:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Stmt1456443193000",
            "Effect": "Allow",
            "Action": [
                "devicefarm:ListDevices"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
}
```