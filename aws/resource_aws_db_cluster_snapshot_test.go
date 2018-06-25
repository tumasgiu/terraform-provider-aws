package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSDBClusterSnapshot_basic(t *testing.T) {
	var v rds.DBClusterSnapshot
	rInt := acctest.RandInt()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsDbClusterSnapshotConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbClusterSnapshotExists("aws_db_cluster_snapshot.test", &v),
				),
			},
		},
	})
}

func testAccCheckDbClusterSnapshotExists(n string, v *rds.DBClusterSnapshot) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		conn := testAccProvider.Meta().(*AWSClient).rdsconn

		params := &rds.DescribeDBClusterSnapshotsInput{
			DBClusterSnapshotIdentifier: aws.String(rs.Primary.ID),
		}

		response, err := conn.DescribeDBClusterSnapshots(params)
		if err == nil {
			if response.DBClusterSnapshots != nil && len(response.DBClusterSnapshots) > 0 {
				*v = *response.DBClusterSnapshots[0]
				return nil
			}
		}
		return fmt.Errorf("Error finding RDS DB Cluster Snapshot %s", rs.Primary.ID)
	}
}

func testAccAwsDbClusterSnapshotConfig(rInt int) string {
	return fmt.Sprintf(`
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
 db_cluster_snapshot_identifier = "testsnapshot%d"
}`, rInt)
}
