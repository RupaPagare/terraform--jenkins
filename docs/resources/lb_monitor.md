---
subcategory: "Elastic Load Balance (ELB)"
---

# huaweicloud\_lb\_monitor

Manages an ELB monitor resource within HuaweiCloud.
This is an alternative to `huaweicloud_lb_monitor_v2`

## Example Usage

```hcl
resource "huaweicloud_lb_monitor" "monitor_1" {
  pool_id     = huaweicloud_lb_pool_v2.pool_1.id
  type        = "PING"
  delay       = 20
  timeout     = 10
  max_retries = 5
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to create the ELB monitor resource.
    If omitted, the provider-level region will be used.
    Changing this creates a new monitor.

* `pool_id` - (Required) The id of the pool that this monitor will be assigned to.

* `name` - (Optional) The Name of the Monitor.

* `tenant_id` - (Optional) Required for admins. The UUID of the tenant who owns
    the monitor.  Only administrative users can specify a tenant UUID
    other than their own. Changing this creates a new monitor.

* `type` - (Required) The type of probe, which is PING, TCP, HTTP, or HTTPS,
    that is sent by the load balancer to verify the member state. Changing this
    creates a new monitor.

* `delay` - (Required) The time, in seconds, between sending probes to members.

* `timeout` - (Required) Maximum number of seconds for a monitor to wait for a
    ping reply before it times out. The value must be less than the delay
    value.

* `max_retries` - (Required) Number of permissible ping failures before
    changing the member's status to INACTIVE. Must be a number between 1
    and 10.

* `url_path` - (Optional) Required for HTTP(S) types. URI path that will be
    accessed if monitor type is HTTP or HTTPS.

*  `http_method` - (Optional) Required for HTTP(S) types. The HTTP method used
    for requests by the monitor. If this attribute is not specified, it
    defaults to "GET".

* `expected_codes` - (Optional) Required for HTTP(S) types. Expected HTTP codes
    for a passing HTTP(S) monitor. You can either specify a single status like
    "200", or a range like "200-202".

* `admin_state_up` - (Optional) The administrative state of the monitor.
    A valid value is true (UP) or false (DOWN).

## Attributes Reference

The following attributes are exported:

* `id` - The unique ID for the monitor.

## Timeouts
This resource provides the following timeouts configuration options:
- `create` - Default is 10 minute.
- `update` - Default is 10 minute.
- `delete` - Default is 10 minute.