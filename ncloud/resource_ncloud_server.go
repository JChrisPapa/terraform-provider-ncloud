package ncloud

import (
	"fmt"
	"log"
	"time"

	"github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNcloudServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceNcloudServerCreate,
		Read:   resourceNcloudServerRead,
		Delete: resourceNcloudServerDelete,
		Update: resourceNcloudServerUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(DefaultCreateTimeout),
			Delete: schema.DefaultTimeout(DefaultTimeout),
		},
		Schema: map[string]*schema.Schema{
			"server_image_product_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Server image product code to determine which server image to create. It can be obtained through getServerImageProductList. You are required to select one among two parameters: server image product code (server_image_product_code) and member server image number(member_server_image_no).",
			},
			"server_product_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Server product code to determine the server specification to create. It can be obtained through the getServerProductList action. Default : Selected as minimum specification. The minimum standards are 1. memory 2. CPU 3. basic block storage size 4. disk type (NET,LOCAL)",
			},
			"member_server_image_no": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Required value when creating a server from a manually created server image. It can be obtained through the getMemberServerImageList action.",
			},
			"server_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateServerName,
				Description:  "Server name to create. default: Assigned by ncloud",
			},
			"server_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Server description to create",
			},
			"login_key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The login key name to encrypt with the public key. Default : Uses the most recently created login key name",
			},
			"is_protect_server_termination": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateBoolValue,
				Description:  "You can set whether or not to protect return when creating. default : false",
			},
			"server_create_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "Number of servers that can be created at a time, and not more than 20 servers can be created at a time. default: 1",
			},
			"server_create_start_no": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "If you create multiple servers at once, the server name will be serialized. You can set the starting number of the serial numbers. The total number of servers created and server starting number cannot exceed 1000. Default : If number of servers created(serverCreateCount) is greater than 1, and if there is no corresponding parameter value, the default will start from 001",
			},
			"internet_line_type_code": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateInternetLineTypeCode,
				Description:  "Internet line identification code. PUBLC(Public), GLBL(Global). default : PUBLC(Public)",
			},
			"fee_system_type_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A rate system identification code. There are time plan(MTRAT) and flat rate (FXSUM). Default : Time plan(MTRAT)",
			},
			"zone_code": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Zone code. You can determine the ZONE where the server will be created. It can be obtained through the getZoneList action. Default : Assigned by NAVER Cloud Platform.",
				ConflictsWith: []string{"zone_no"},
			},
			"zone_no": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Zone number. You can determine the ZONE where the server will be created. It can be obtained through the getZoneList action. Default : Assigned by NAVER Cloud Platform.",
				ConflictsWith: []string{"zone_code"},
			},

			"access_control_group_configuration_no_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				MinItems:    1,
				Description: "You can set the ACG created when creating the server. ACG setting number can be obtained through the getAccessControlGroupList action. Default : Default ACG number",
			},
			"user_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The server will execute the user data script set by the user at first boot. To view the column, it is returned only when viewing the server instance. You must need base64 Encoding, URL Encoding before put in value of userData. If you don't URL Encoding again it occurs signature invalid error.",
			},
			"raid_type_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Raid Type Name",
			},

			"server_instance_no": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"base_block_storage_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"platform_type": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     commonCodeSchemaResource,
			},
			"is_fee_charging_monitoring": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"public_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_image_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_instance_status": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     commonCodeSchemaResource,
			},
			"server_instance_operation": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     commonCodeSchemaResource,
			},
			"server_instance_status_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uptime": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_forwarding_public_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_forwarding_external_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_forwarding_internal_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     zoneSchemaResource,
			},
			"region": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     regionSchemaResource,
			},
			"base_block_storage_disk_type": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     commonCodeSchemaResource,
			},
			"base_block_storage_disk_detail_type": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     commonCodeSchemaResource,
			},
			"internet_line_type": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     commonCodeSchemaResource,
			},
		},
	}
}

func resourceNcloudServerCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*NcloudSdk).conn

	reqParams := buildCreateServerInstanceReqParams(conn, d)

	var resp *sdk.ServerInstanceList
	err := resource.Retry(10*time.Second, func() *resource.RetryError {
		var err error
		resp, err = conn.CreateServerInstances(reqParams)

		if err != nil && resp != nil && isRetryableErr(&resp.CommonResponse, 23006) {
			return resource.RetryableError(err)
		}
		return resource.NonRetryableError(err)
	})

	if err != nil {
		logErrorResponse("CreateServerInstances", err, reqParams)
		return err
	}
	logCommonResponse("CreateServerInstances", reqParams, resp.CommonResponse)

	serverInstance := &resp.ServerInstanceList[0]
	d.SetId(serverInstance.ServerInstanceNo)

	if err := waitForServerInstance(conn, serverInstance.ServerInstanceNo, "RUN"); err != nil {
		return err
	}
	return resourceNcloudServerRead(d, meta)
}

func resourceNcloudServerRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*NcloudSdk).conn

	instance, err := getServerInstance(conn, d.Id())
	if err != nil {
		return err
	}

	if instance != nil {
		d.Set("server_instance_no", instance.ServerInstanceNo)
		d.Set("server_name", instance.ServerName)
		d.Set("server_image_product_code", instance.ServerImageProductCode)
		d.Set("server_instance_status", setCommonCode(instance.ServerInstanceStatus))
		d.Set("server_instance_status_name", instance.ServerInstanceStatusName)
		d.Set("uptime", instance.Uptime)
		d.Set("server_image_name", instance.ServerImageName)
		d.Set("private_ip", instance.PrivateIP)
		d.Set("cpu_count", instance.CPUCount)
		d.Set("memory_size", instance.MemorySize)
		d.Set("base_block_storage_size", instance.BaseBlockStorageSize)
		d.Set("platform_type", setCommonCode(instance.PlatformType))
		d.Set("is_fee_charging_monitoring", instance.IsFeeChargingMonitoring)
		d.Set("public_ip", instance.PublicIP)
		d.Set("private_ip", instance.PrivateIP)
		d.Set("server_instance_operation", setCommonCode(instance.ServerInstanceOperation))
		d.Set("create_date", instance.CreateDate)
		d.Set("uptime", instance.Uptime)
		d.Set("port_forwarding_public_ip", instance.PortForwardingPublicIP)
		d.Set("port_forwarding_external_port", instance.PortForwardingExternalPort)
		d.Set("port_forwarding_internal_port", instance.PortForwardingInternalPort)
		d.Set("zone", setZone(instance.Zone))
		d.Set("region", setRegion(instance.Region))
		d.Set("base_block_storage_disk_type", setCommonCode(instance.BaseBlockStorageDiskType))
		d.Set("base_block_storage_disk_detail_type", setCommonCode(instance.BaseBlockStroageDiskDetailType))
		d.Set("internet_line_type", setCommonCode(instance.InternetLineType))
		d.Set("user_data", d.Get("user_data").(string))

	}

	return nil
}

func resourceNcloudServerDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*NcloudSdk).conn
	serverInstance, err := getServerInstance(conn, d.Id())
	if err != nil {
		return err
	}

	if serverInstance.ServerInstanceStatus.Code != "NSTOP" {
		if err := stopServerInstance(conn, d.Id()); err != nil {
			return err
		}
		if err := waitForServerInstance(conn, serverInstance.ServerInstanceNo, "NSTOP"); err != nil {
			return err
		}
	}

	err = detachBlockStorageByServerInstanceNo(conn, d.Id())
	if err != nil {
		log.Printf("[ERROR] detachBlockStorageByServerInstanceNo err: %s", err)
		return err
	}

	return terminateServerInstance(conn, d.Id())
}

func resourceNcloudServerUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*NcloudSdk).conn

	if d.HasChange("server_product_code") {
		reqParams := &sdk.RequestChangeServerInstanceSpec{
			ServerInstanceNo:  d.Get("server_instance_no").(string),
			ServerProductCode: d.Get("server_product_code").(string),
		}

		resp, err := conn.ChangeServerInstanceSpec(reqParams)
		if err != nil {
			logErrorResponse("ChangeServerInstanceSpec", err, reqParams)
			return err
		}
		logCommonResponse("ChangeServerInstanceSpec", reqParams, resp.CommonResponse)
	}

	return resourceNcloudServerRead(d, meta)
}

func buildCreateServerInstanceReqParams(conn *sdk.Conn, d *schema.ResourceData) *sdk.RequestCreateServerInstance {

	var paramAccessControlGroupConfigurationNoList []string
	if param, ok := d.GetOk("access_control_group_configuration_no_list"); ok {
		paramAccessControlGroupConfigurationNoList = StringList(param.([]interface{}))
	}

	reqParams := &sdk.RequestCreateServerInstance{
		ServerImageProductCode:     d.Get("server_image_product_code").(string),
		ServerProductCode:          d.Get("server_product_code").(string),
		MemberServerImageNo:        d.Get("member_server_image_no").(string),
		ServerName:                 d.Get("server_name").(string),
		ServerDescription:          d.Get("server_description").(string),
		LoginKeyName:               d.Get("login_key_name").(string),
		IsProtectServerTermination: d.Get("is_protect_server_termination").(string),
		ServerCreateCount:          d.Get("server_create_count").(int),
		ServerCreateStartNo:        d.Get("server_create_start_no").(int),
		InternetLineTypeCode:       d.Get("internet_line_type_code").(string),
		FeeSystemTypeCode:          d.Get("fee_system_type_code").(string),
		ZoneNo:                     parseZoneNoParameter(conn, d),
		AccessControlGroupConfigurationNoList: paramAccessControlGroupConfigurationNoList,
		UserData:     d.Get("user_data").(string),
		RaidTypeName: d.Get("raid_type_name").(string),
	}
	return reqParams
}

func getServerInstance(conn *sdk.Conn, serverInstanceNo string) (*sdk.ServerInstance, error) {
	reqParams := new(sdk.RequestGetServerInstanceList)
	reqParams.ServerInstanceNoList = []string{serverInstanceNo}
	resp, err := conn.GetServerInstanceList(reqParams)

	if err != nil {
		logErrorResponse("GetServerInstanceList", err, reqParams)
		return nil, err
	}
	logCommonResponse("GetServerInstanceList", reqParams, resp.CommonResponse)
	if len(resp.ServerInstanceList) > 0 {
		inst := &resp.ServerInstanceList[0]
		return inst, nil
	}
	return nil, nil
}

func stopServerInstance(conn *sdk.Conn, serverInstanceNo string) error {
	reqParams := &sdk.RequestStopServerInstances{
		ServerInstanceNoList: []string{serverInstanceNo},
	}
	resp, err := conn.StopServerInstances(reqParams)
	if err != nil {
		logErrorResponse("StopServerInstances", err, reqParams)
		return err
	}
	logCommonResponse("StopServerInstances", reqParams, resp.CommonResponse)

	return nil
}

func terminateServerInstance(conn *sdk.Conn, serverInstanceNo string) error {
	reqParams := &sdk.RequestTerminateServerInstances{
		ServerInstanceNoList: []string{serverInstanceNo},
	}
	resp, err := conn.TerminateServerInstances(reqParams)
	if err != nil {
		logErrorResponse("TerminateServerInstances", err, reqParams)
		// TODO: check 502 Bad Gateway error
		// return err
		return nil
	}
	logCommonResponse("TerminateServerInstances", reqParams, resp.CommonResponse)
	return nil
}

func waitForServerInstance(conn *sdk.Conn, instanceId string, status string) error {

	c1 := make(chan error, 1)

	go func() {
		for {
			instance, err := getServerInstance(conn, instanceId)

			if err != nil {
				c1 <- err
				return
			}
			if instance == nil || instance.ServerInstanceStatus.Code == status {
				c1 <- nil
				return
			}
			log.Printf("[DEBUG] Wait to server instance (%s)", instanceId)
			time.Sleep(time.Second * 1)
		}
	}()

	select {
	case res := <-c1:
		return res
	case <-time.After(DefaultCreateTimeout):
		return fmt.Errorf("TIMEOUT : Wait to server instance  (%s)", instanceId)
	}

}