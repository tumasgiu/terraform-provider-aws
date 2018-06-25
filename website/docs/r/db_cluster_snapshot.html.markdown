---
layout: "aws"
page_title: "AWS: aws_db_cluster_snapshot"
sidebar_current: "docs-aws-resource-db-cluster-snapshot"
description: |-
  Creates a Snapshot of an RDS Aurora Cluster.
---

# aws_db_cluster_snapshot

Creates a Snapshot of an RDS Aurora Cluster.

## Example Usage

```hcl
data "aws_availability_zones" "available" {}

resource "aws_rds_cluster" "test" {
  cluster_identifier      = "test-aurora-cluster"
  availability_zones      = ["${data.aws_availability_zones.available.names}"]
  database_name           = "mydb"
  master_username         = "foo"
  master_password         = "barbarbarbar"
  skip_final_snapshot     = true
}

resource "aws_db_cluster_snapshot" "test" {
 db_cluster_identifier = "${aws_rds_cluster.test.cluster_identifier}"
 db_cluster_snapshot_identifier = "testsnapshot1234"
}
```

## Argument Reference

The following arguments are supported:

* `db_cluster_snapshot_identifier` - (Required) The DB Cluster Identifier from which to take the snapshot.
* `db_cluster_identifier` - (Required) The Identifier for the snapshot.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allocated_storage` - Specifies the allocated storage size in gigabytes (GB).
* `availability_zones` - Provides the list of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
* `db_cluster_snapshot_arn` - The Amazon Resource Name (ARN) for the DB cluster snapshot.
* `engine` - Specifies the name of the database engine.
* `engine_version` - Specifies the version of the database engine.
* `kms_key_id` - The ARN for the KMS encryption key.
* `license_model` - License model information for the restored DB instance.
* `iam_database_authentication_enabled` - True if mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false. 
* `source_db_snapshot_identifier` - The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
* `master_username` - Provides the master username for the DB cluster snapshot.
* `port` - Specifies the port that the DB cluster was listening on at the time of the snapshot.
* `snapshot_type` - Provides the type of the DB cluster snapshot.
* `source_db_cluster_snapshot_arn` - If the DB cluster snapshot was copied from a source DB cluster snapshot, the Amazon Resource Name (ARN) for the source DB cluster snapshot
* `status` - Specifies the status of this DB cluster snapshot.
* `storage_type` - Specifies the storage type associated with DB snapshot.
* `storage_encrypted` - Specifies whether the DB snapshot is encrypted.
* `vpc_id` - Provides the VPC ID associated with the DB cluster snapshot.