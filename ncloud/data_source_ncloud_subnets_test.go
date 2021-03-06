package ncloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNcloudSubnetsBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNcloudSubnetsConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceID("data.ncloud_subnets.all"),
				),
			},
		},
	})
}

func TestAccDataSourceNcloudSubnetsName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNcloudSubnetsConfigSubnet("10.2.1.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceID("data.ncloud_subnets.by_cidr"),
				),
			},
		},
	})
}

func TestAccDataSourceNcloudSubnetsVpcNo(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNcloudSubnetsConfigVpcNo("502"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceID("data.ncloud_subnets.by_vpc_no"),
				),
			},
		},
	})
}

func testAccDataSourceNcloudSubnetsConfig() string {
	return `
data "ncloud_subnets" "all" {}
`
}

func testAccDataSourceNcloudSubnetsConfigSubnet(cidr string) string {
	return fmt.Sprintf(`
data "ncloud_subnets" "by_cidr" {
	subnet = "%s"
}
`, cidr)
}

func testAccDataSourceNcloudSubnetsConfigVpcNo(vpcNo string) string {
	return fmt.Sprintf(`
data "ncloud_subnets" "by_vpc_no" {
	vpc_no = "%s"
}
`, vpcNo)
}
