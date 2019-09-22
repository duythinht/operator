### Basic usage

```
you[at]ci-runner$ operator release gitlab workloads/develop -a https://git.teko.vn -t [your_pat_token]
## Changelog(time="2019-09-23 00:59", ancestor="1.0.0")

* Missing Defination:
  * 93591e3b Teko auto release hub.teko.vn/oms/cqrs:m-67778c8167
  * cfa4db66 Teko auto release hub.teko.vn/oms/cqrs:m-85c967a47a
  * d9c21661 set app for upgrate db api v2
  * c58c6504 Update warehouse-api-v2.yaml
  * 6e582e6a Update warehouse-api-v2.yaml
  * fc80c16f Teko auto release hub.teko.vn/warehouse/apiv2:m-ea51dfd0
  * dff3319b Teko auto release hub.teko.vn/warehouse/apiv2:m-8e81c6ff
  * 3f8c4fa0 Teko auto release hub.teko.vn/warehouse/connector:m-a6f75d77
  * d1492cd4 add db config warehouse-connector
  * 7a0739e1 Teko auto release hub.teko.vn/warehouse/apiv2:m-3daf20a0
  * ef2380b1 edit db name api-v2
  * bebc438e Teko auto release hub.teko.vn/warehouse/connector:m-eeaa5219
  * 244006ae Teko auto release hub.teko.vn/oms/opb:m-98d229ce
  * b6e674d7 Teko auto release hub.teko.vn/warehouse/odoo:m-6f22ed2d
  * 541cf975 update warehouse-stack
  * 2ebd8347 add domain
  * 12df4cf8 Teko auto release hub.teko.vn/warehouse/odoo:m-00aa8fd3
  * bcf2a85e Teko auto release hub.teko.vn/adidi/sites/admin:m-8b9a59a1
  * 7b99eed0 Change adidi port
  * 6b5744b8 Add adidi helmfile to this stack
  * 1084e2ed Add adidi site admin
  * de125763 Teko auto release hub.teko.vn/warehouse/odoo:m-87611b51
  * ad049cd7 update config pg odoo
  * ffe2bd6e edit postgresql host
  * 6d0cacd2 Teko auto release hub.teko.vn/oms/cqrs:m-104314d00f
  * ed0123af Add redis host
  * 74f59368 Teko auto release hub.teko.vn/oms/cqrs:m-5014d63237
  * c80f4cb0 edit job warehouse-odoo
  * 56134c24 Update chart version
  * 4d10f68b Update args
  * 9886bfc3 Change resource
  * 516b6784 update job init db
  * 92c50b44 Teko auto release hub.teko.vn/oms/cqrs:m-714579836b
  * c56ba349 Update config for teko-secret due to health check fail

you[at]ci-runner$ operator release gitlab workloads/develop -a https://git.teko.vn -t [your_pat_token] --submit
A release was submmited!!!
```
