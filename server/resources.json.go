package main

var jsonResources = []byte(`
[
  {
    "name": "zoneLoadBalancer",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer",
    "location": "centralus",
    "associations": [
      {
        "name": "backendPool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
        "associationType": "Contains"
      },
      {
        "name": "healthProbe",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/probes/healthProbe",
        "associationType": "Contains"
      },
      {
        "name": "frontBool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool",
        "associationType": "Contains"
      }
    ],
    "type": "Microsoft.Network/loadBalancers",
    "azDetails": {
      "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer",
      "location": "centralus",
      "properties": {
        "frontendIPConfigurations": [
          {
            "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool",
            "name": "frontBool",
            "properties": {
              "loadBalancingRules": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/loadBalancingRules/basicRule"
                }
              ],
              "privateIPAllocationMethod": "Dynamic",
              "privateIPAddressVersion": "IPv4",
              "publicIPAddress": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/publicIPAddresses/zoneIP"
              },
              "provisioningState": "Succeeded"
            }
          }
        ],
        "backendAddressPools": [
          {
            "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
            "name": "backendPool",
            "properties": {
              "backendIPConfigurations": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1/ipConfigurations/ipconfig1"
                },
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2/ipConfigurations/ipconfig1"
                },
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3/ipConfigurations/ipconfig1"
                }
              ],
              "loadBalancingRules": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/loadBalancingRules/basicRule"
                }
              ],
              "provisioningState": "Succeeded"
            }
          }
        ],
        "loadBalancingRules": [
          {
            "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/loadBalancingRules/basicRule",
            "name": "basicRule",
            "properties": {
              "frontendIPConfiguration": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool"
              },
              "backendAddressPool": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool"
              },
              "probe": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/probes/healthProbe"
              },
              "protocol": "Tcp",
              "loadDistribution": "Default",
              "frontendPort": 80,
              "backendPort": 80,
              "idleTimeoutInMinutes": 4,
              "enableFloatingIP": false,
              "enableTcpReset": false,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "probes": [
          {
            "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/probes/healthProbe",
            "name": "healthProbe",
            "properties": {
              "loadBalancingRules": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/loadBalancingRules/basicRule"
                }
              ],
              "protocol": "Tcp",
              "port": 80,
              "intervalInSeconds": 15,
              "numberOfProbes": 2,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "inboundNatRules": [],
        "inboundNatPools": [],
        "resourceGuid": "36e5ffe7-ac76-4707-9332-41e5f5f88890",
        "provisioningState": "Succeeded"
      },
      "sku": {
        "name": "Basic"
      },
      "tags": {}
    }
  },
  {
    "name": "backendPool",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
    "location": "centralus",
    "associations": [
      {
        "name": "myNic1",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1",
        "associationType": "Contains"
      },
      {
        "name": "myNic2",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2",
        "associationType": "Contains"
      },
      {
        "name": "myNic3",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3",
        "associationType": "Contains"
      }
    ]
  },
  {
    "name": "healthProbe",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/probes/healthProbe",
    "location": "centralus",
    "associations": []
  },
  {
    "name": "frontBool",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool",
    "location": "centralus",
    "associations": []
  },
  {
    "name": "myNic1",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1",
    "location": "centralus",
    "associations": [
      {
        "name": "veneraNSG",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG",
        "associationType": "Associated"
      },
      {
        "name": "veneraSubnet",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
        "associationType": "Associated"
      },
      {
        "name": "backendPool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
        "associationType": "Associated"
      }
    ],
    "type": "Microsoft.Network/networkInterfaces",
    "azDetails": {
      "etag": "W/\"7ec73e7a-57a1-42dd-951f-2916b958a4bf\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1",
      "location": "centralus",
      "properties": {
        "networkSecurityGroup": {
          "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG"
        },
        "ipConfigurations": [
          {
            "etag": "W/\"7ec73e7a-57a1-42dd-951f-2916b958a4bf\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1/ipConfigurations/ipconfig1",
            "name": "ipconfig1",
            "properties": {
              "loadBalancerBackendAddressPools": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool"
                }
              ],
              "privateIPAddress": "10.0.0.4",
              "privateIPAllocationMethod": "Dynamic",
              "privateIPAddressVersion": "IPv4",
              "subnet": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet"
              },
              "primary": true,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "tapConfigurations": [],
        "dnsSettings": {
          "dnsServers": [],
          "appliedDnsServers": [],
          "internalDomainNameSuffix": "vgxoqu0elhzuhgvqwklgknekzh.gx.internal.cloudapp.net"
        },
        "enableAcceleratedNetworking": false,
        "enableIPForwarding": false,
        "hostedWorkloads": [],
        "resourceGuid": "50aba3bf-ef3f-47a6-9502-821c15da40e7",
        "provisioningState": "Succeeded"
      }
    }
  },
  {
    "name": "myNic2",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2",
    "location": "centralus",
    "associations": [
      {
        "name": "veneraNSG",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG",
        "associationType": "Associated"
      },
      {
        "name": "veneraSubnet",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
        "associationType": "Associated"
      },
      {
        "name": "backendPool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
        "associationType": "Associated"
      }
    ],
    "type": "Microsoft.Network/networkInterfaces",
    "azDetails": {
      "etag": "W/\"ff4780c0-3540-4be4-9b4a-44bae873707b\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2",
      "location": "centralus",
      "properties": {
        "networkSecurityGroup": {
          "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG"
        },
        "ipConfigurations": [
          {
            "etag": "W/\"ff4780c0-3540-4be4-9b4a-44bae873707b\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2/ipConfigurations/ipconfig1",
            "name": "ipconfig1",
            "properties": {
              "loadBalancerBackendAddressPools": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool"
                }
              ],
              "privateIPAddress": "10.0.0.5",
              "privateIPAllocationMethod": "Dynamic",
              "privateIPAddressVersion": "IPv4",
              "subnet": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet"
              },
              "primary": true,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "tapConfigurations": [],
        "dnsSettings": {
          "dnsServers": [],
          "appliedDnsServers": [],
          "internalDomainNameSuffix": "vgxoqu0elhzuhgvqwklgknekzh.gx.internal.cloudapp.net"
        },
        "enableAcceleratedNetworking": false,
        "enableIPForwarding": false,
        "hostedWorkloads": [],
        "resourceGuid": "0c2ceb16-08cc-44a7-919d-877a14ea37da",
        "provisioningState": "Succeeded"
      }
    }
  },
  {
    "name": "myNic3",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3",
    "location": "centralus",
    "associations": [
      {
        "name": "veneraNSG",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG",
        "associationType": "Associated"
      },
      {
        "name": "veneraSubnet",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
        "associationType": "Associated"
      },
      {
        "name": "backendPool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
        "associationType": "Associated"
      }
    ],
    "type": "Microsoft.Network/networkInterfaces",
    "azDetails": {
      "etag": "W/\"9452cd85-d146-4498-8714-cda7d31364b4\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3",
      "location": "centralus",
      "properties": {
        "networkSecurityGroup": {
          "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG"
        },
        "ipConfigurations": [
          {
            "etag": "W/\"9452cd85-d146-4498-8714-cda7d31364b4\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3/ipConfigurations/ipconfig1",
            "name": "ipconfig1",
            "properties": {
              "loadBalancerBackendAddressPools": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool"
                }
              ],
              "privateIPAddress": "10.0.0.6",
              "privateIPAllocationMethod": "Dynamic",
              "privateIPAddressVersion": "IPv4",
              "subnet": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet"
              },
              "primary": true,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "tapConfigurations": [],
        "dnsSettings": {
          "dnsServers": [],
          "appliedDnsServers": [],
          "internalDomainNameSuffix": "vgxoqu0elhzuhgvqwklgknekzh.gx.internal.cloudapp.net"
        },
        "enableAcceleratedNetworking": false,
        "enableIPForwarding": false,
        "hostedWorkloads": [],
        "resourceGuid": "fa6e6650-fef0-4bd9-93dd-01fb254bdc3c",
        "provisioningState": "Succeeded"
      }
    }
  },
  {
    "name": "veneraNSG",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG",
    "location": "centralus",
    "associations": []
  },
  {
    "name": "zoneIP",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/publicIPAddresses/zoneIP",
    "location": "centralus",
    "associations": [
      {
        "name": "zoneLoadBalancer",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer",
        "associationType": "Associated"
      }
    ],
    "type": "Microsoft.Network/publicIPAddresses",
    "azDetails": {
      "etag": "W/\"6a2ee7be-1810-496c-8e3b-4c0b0f6a6007\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/publicIPAddresses/zoneIP",
      "location": "centralus",
      "properties": {
        "publicIPAllocationMethod": "Dynamic",
        "publicIPAddressVersion": "IPv4",
        "ipConfiguration": {
          "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
          "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool",
          "name": "frontBool",
          "properties": {
            "privateIPAllocationMethod": "Dynamic",
            "publicIPAddress": {
              "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/publicIPAddresses/zoneIP"
            },
            "provisioningState": "Succeeded"
          }
        },
        "ipTags": [],
        "idleTimeoutInMinutes": 4,
        "resourceGuid": "bf7e89d7-5183-42c5-9f0b-30f5faf70a24",
        "provisioningState": "Succeeded"
      },
      "sku": {
        "name": "Basic"
      }
    }
  },
  {
    "name": "veneraVNET",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET",
    "location": "centralus",
    "associations": [
      {
        "name": "veneraSubnet",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
        "associationType": "Contains"
      }
    ],
    "type": "Microsoft.Network/virtualNetworks",
    "azDetails": {
      "etag": "W/\"8ccc318b-8bd6-4117-8b09-6c23f01ef739\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET",
      "location": "centralus",
      "properties": {
        "addressSpace": {
          "addressPrefixes": [
            "10.0.0.0/16"
          ]
        },
        "dhcpOptions": {
          "dnsServers": []
        },
        "subnets": [
          {
            "etag": "W/\"8ccc318b-8bd6-4117-8b09-6c23f01ef739\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
            "name": "veneraSubnet",
            "properties": {
              "addressPrefix": "10.0.0.0/24",
              "ipConfigurations": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1/ipConfigurations/ipconfig1"
                },
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2/ipConfigurations/ipconfig1"
                },
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3/ipConfigurations/ipconfig1"
                }
              ],
              "delegations": [],
              "provisioningState": "Succeeded",
              "privateEndpointNetworkPolicies": "Enabled",
              "privateLinkServiceNetworkPolicies": "Enabled"
            }
          }
        ],
        "virtualNetworkPeerings": [],
        "resourceGuid": "53e8aea9-5944-43f3-9ab0-b29665348acf",
        "provisioningState": "Succeeded",
        "enableDdosProtection": false,
        "enableVmProtection": false
      },
      "tags": {}
    }
  },
  {
    "name": "veneraSubnet",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
    "location": "centralus",
    "associations": []
  }
]
`)

var jsonResourcesSlice = [][]byte{
	[]byte(`{
    "name": "zoneLoadBalancer",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer",
    "location": "centralus",
    "associations": [
      {
        "name": "backendPool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
        "associationType": "Contains"
      },
      {
        "name": "healthProbe",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/probes/healthProbe",
        "associationType": "Contains"
      },
      {
        "name": "frontBool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool",
        "associationType": "Contains"
      }
    ],
    "type": "Microsoft.Network/loadBalancers",
    "azDetails": {
      "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer",
      "location": "centralus",
      "properties": {
        "frontendIPConfigurations": [
          {
            "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool",
            "name": "frontBool",
            "properties": {
              "loadBalancingRules": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/loadBalancingRules/basicRule"
                }
              ],
              "privateIPAllocationMethod": "Dynamic",
              "privateIPAddressVersion": "IPv4",
              "publicIPAddress": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/publicIPAddresses/zoneIP"
              },
              "provisioningState": "Succeeded"
            }
          }
        ],
        "backendAddressPools": [
          {
            "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
            "name": "backendPool",
            "properties": {
              "backendIPConfigurations": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1/ipConfigurations/ipconfig1"
                },
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2/ipConfigurations/ipconfig1"
                },
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3/ipConfigurations/ipconfig1"
                }
              ],
              "loadBalancingRules": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/loadBalancingRules/basicRule"
                }
              ],
              "provisioningState": "Succeeded"
            }
          }
        ],
        "loadBalancingRules": [
          {
            "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/loadBalancingRules/basicRule",
            "name": "basicRule",
            "properties": {
              "frontendIPConfiguration": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool"
              },
              "backendAddressPool": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool"
              },
              "probe": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/probes/healthProbe"
              },
              "protocol": "Tcp",
              "loadDistribution": "Default",
              "frontendPort": 80,
              "backendPort": 80,
              "idleTimeoutInMinutes": 4,
              "enableFloatingIP": false,
              "enableTcpReset": false,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "probes": [
          {
            "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/probes/healthProbe",
            "name": "healthProbe",
            "properties": {
              "loadBalancingRules": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/loadBalancingRules/basicRule"
                }
              ],
              "protocol": "Tcp",
              "port": 80,
              "intervalInSeconds": 15,
              "numberOfProbes": 2,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "inboundNatRules": [],
        "inboundNatPools": [],
        "resourceGuid": "36e5ffe7-ac76-4707-9332-41e5f5f88890",
        "provisioningState": "Succeeded"
      },
      "sku": {
        "name": "Basic"
      },
      "tags": {}
    }
  }`),
	[]byte(`{
    "name": "backendPool",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
    "location": "centralus",
    "associations": [
      {
        "name": "myNic1",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1",
        "associationType": "Contains"
      },
      {
        "name": "myNic2",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2",
        "associationType": "Contains"
      },
      {
        "name": "myNic3",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3",
        "associationType": "Contains"
      }
    ]
  }`),
	[]byte(`{
    "name": "healthProbe",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/probes/healthProbe",
    "location": "centralus",
    "associations": []
  }`),
	[]byte(`{
    "name": "frontBool",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool",
    "location": "centralus",
    "associations": []
  }`),
	[]byte(`{
    "name": "myNic1",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1",
    "location": "centralus",
    "associations": [
      {
        "name": "veneraNSG",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG",
        "associationType": "Associated"
      },
      {
        "name": "veneraSubnet",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
        "associationType": "Associated"
      },
      {
        "name": "backendPool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
        "associationType": "Associated"
      }
    ],
    "type": "Microsoft.Network/networkInterfaces",
    "azDetails": {
      "etag": "W/\"7ec73e7a-57a1-42dd-951f-2916b958a4bf\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1",
      "location": "centralus",
      "properties": {
        "networkSecurityGroup": {
          "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG"
        },
        "ipConfigurations": [
          {
            "etag": "W/\"7ec73e7a-57a1-42dd-951f-2916b958a4bf\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1/ipConfigurations/ipconfig1",
            "name": "ipconfig1",
            "properties": {
              "loadBalancerBackendAddressPools": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool"
                }
              ],
              "privateIPAddress": "10.0.0.4",
              "privateIPAllocationMethod": "Dynamic",
              "privateIPAddressVersion": "IPv4",
              "subnet": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet"
              },
              "primary": true,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "tapConfigurations": [],
        "dnsSettings": {
          "dnsServers": [],
          "appliedDnsServers": [],
          "internalDomainNameSuffix": "vgxoqu0elhzuhgvqwklgknekzh.gx.internal.cloudapp.net"
        },
        "enableAcceleratedNetworking": false,
        "enableIPForwarding": false,
        "hostedWorkloads": [],
        "resourceGuid": "50aba3bf-ef3f-47a6-9502-821c15da40e7",
        "provisioningState": "Succeeded"
      }
    }
  }`),
	[]byte(`{
    "name": "myNic2",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2",
    "location": "centralus",
    "associations": [
      {
        "name": "veneraNSG",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG",
        "associationType": "Associated"
      },
      {
        "name": "veneraSubnet",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
        "associationType": "Associated"
      },
      {
        "name": "backendPool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
        "associationType": "Associated"
      }
    ],
    "type": "Microsoft.Network/networkInterfaces",
    "azDetails": {
      "etag": "W/\"ff4780c0-3540-4be4-9b4a-44bae873707b\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2",
      "location": "centralus",
      "properties": {
        "networkSecurityGroup": {
          "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG"
        },
        "ipConfigurations": [
          {
            "etag": "W/\"ff4780c0-3540-4be4-9b4a-44bae873707b\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2/ipConfigurations/ipconfig1",
            "name": "ipconfig1",
            "properties": {
              "loadBalancerBackendAddressPools": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool"
                }
              ],
              "privateIPAddress": "10.0.0.5",
              "privateIPAllocationMethod": "Dynamic",
              "privateIPAddressVersion": "IPv4",
              "subnet": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet"
              },
              "primary": true,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "tapConfigurations": [],
        "dnsSettings": {
          "dnsServers": [],
          "appliedDnsServers": [],
          "internalDomainNameSuffix": "vgxoqu0elhzuhgvqwklgknekzh.gx.internal.cloudapp.net"
        },
        "enableAcceleratedNetworking": false,
        "enableIPForwarding": false,
        "hostedWorkloads": [],
        "resourceGuid": "0c2ceb16-08cc-44a7-919d-877a14ea37da",
        "provisioningState": "Succeeded"
      }
    }
  }`),
	[]byte(`{
    "name": "myNic3",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3",
    "location": "centralus",
    "associations": [
      {
        "name": "veneraNSG",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG",
        "associationType": "Associated"
      },
      {
        "name": "veneraSubnet",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
        "associationType": "Associated"
      },
      {
        "name": "backendPool",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool",
        "associationType": "Associated"
      }
    ],
    "type": "Microsoft.Network/networkInterfaces",
    "azDetails": {
      "etag": "W/\"9452cd85-d146-4498-8714-cda7d31364b4\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3",
      "location": "centralus",
      "properties": {
        "networkSecurityGroup": {
          "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG"
        },
        "ipConfigurations": [
          {
            "etag": "W/\"9452cd85-d146-4498-8714-cda7d31364b4\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3/ipConfigurations/ipconfig1",
            "name": "ipconfig1",
            "properties": {
              "loadBalancerBackendAddressPools": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/backendAddressPools/backendPool"
                }
              ],
              "privateIPAddress": "10.0.0.6",
              "privateIPAllocationMethod": "Dynamic",
              "privateIPAddressVersion": "IPv4",
              "subnet": {
                "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet"
              },
              "primary": true,
              "provisioningState": "Succeeded"
            }
          }
        ],
        "tapConfigurations": [],
        "dnsSettings": {
          "dnsServers": [],
          "appliedDnsServers": [],
          "internalDomainNameSuffix": "vgxoqu0elhzuhgvqwklgknekzh.gx.internal.cloudapp.net"
        },
        "enableAcceleratedNetworking": false,
        "enableIPForwarding": false,
        "hostedWorkloads": [],
        "resourceGuid": "fa6e6650-fef0-4bd9-93dd-01fb254bdc3c",
        "provisioningState": "Succeeded"
      }
    }
  }`),
	[]byte(`{
    "name": "veneraNSG",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkSecurityGroups/veneraNSG",
    "location": "centralus",
    "associations": []
  }`),
	[]byte(`{
    "name": "zoneIP",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/publicIPAddresses/zoneIP",
    "location": "centralus",
    "associations": [
      {
        "name": "zoneLoadBalancer",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer",
        "associationType": "Associated"
      }
    ],
    "type": "Microsoft.Network/publicIPAddresses",
    "azDetails": {
      "etag": "W/\"6a2ee7be-1810-496c-8e3b-4c0b0f6a6007\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/publicIPAddresses/zoneIP",
      "location": "centralus",
      "properties": {
        "publicIPAllocationMethod": "Dynamic",
        "publicIPAddressVersion": "IPv4",
        "ipConfiguration": {
          "etag": "W/\"4e76e58a-3f08-4f87-9cfe-fcd8c53c3834\"",
          "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/loadBalancers/zoneLoadBalancer/frontendIPConfigurations/frontBool",
          "name": "frontBool",
          "properties": {
            "privateIPAllocationMethod": "Dynamic",
            "publicIPAddress": {
              "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/publicIPAddresses/zoneIP"
            },
            "provisioningState": "Succeeded"
          }
        },
        "ipTags": [],
        "idleTimeoutInMinutes": 4,
        "resourceGuid": "bf7e89d7-5183-42c5-9f0b-30f5faf70a24",
        "provisioningState": "Succeeded"
      },
      "sku": {
        "name": "Basic"
      }
    }
  }`),
	[]byte(`{
    "name": "veneraVNET",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET",
    "location": "centralus",
    "associations": [
      {
        "name": "veneraSubnet",
        "resourceId": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
        "associationType": "Contains"
      }
    ],
    "type": "Microsoft.Network/virtualNetworks",
    "azDetails": {
      "etag": "W/\"8ccc318b-8bd6-4117-8b09-6c23f01ef739\"",
      "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET",
      "location": "centralus",
      "properties": {
        "addressSpace": {
          "addressPrefixes": [
            "10.0.0.0/16"
          ]
        },
        "dhcpOptions": {
          "dnsServers": []
        },
        "subnets": [
          {
            "etag": "W/\"8ccc318b-8bd6-4117-8b09-6c23f01ef739\"",
            "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
            "name": "veneraSubnet",
            "properties": {
              "addressPrefix": "10.0.0.0/24",
              "ipConfigurations": [
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic1/ipConfigurations/ipconfig1"
                },
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic2/ipConfigurations/ipconfig1"
                },
                {
                  "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/networkInterfaces/myNic3/ipConfigurations/ipconfig1"
                }
              ],
              "delegations": [],
              "provisioningState": "Succeeded",
              "privateEndpointNetworkPolicies": "Enabled",
              "privateLinkServiceNetworkPolicies": "Enabled"
            }
          }
        ],
        "virtualNetworkPeerings": [],
        "resourceGuid": "53e8aea9-5944-43f3-9ab0-b29665348acf",
        "provisioningState": "Succeeded",
        "enableDdosProtection": false,
        "enableVmProtection": false
      },
      "tags": {}
    }
  }`),
	[]byte(`{
    "name": "veneraSubnet",
    "id": "/subscriptions/fcc18659-b65b-4b94-90a1-3ca49d3a3601/resourceGroups/veneraGroup/providers/Microsoft.Network/virtualNetworks/veneraVNET/subnets/veneraSubnet",
    "location": "centralus",
    "associations": []
  }`),
}
