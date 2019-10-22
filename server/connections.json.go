package main

type Connections struct {
	Edges []Edge          `json:"edges"`
	Nodes map[string]Node `json:"nodes"`
}

type Edge struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type Node struct {
	Hostname  string   `json:"hostname"`
	Ips       []string `json:"ips"`
	Processes []string `json:"processes"`
	Openports []int    `json:"openports"`
}

var jsonConnections = []byte(`{
  "edges": [
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "2606a1fc-77a1-45d1-b69a-db9d6d11adcc"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "a2ebface-b962-4f45-80a6-e0e4c49d2e55"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "373e25ac-17d9-48c2-b7b0-6a611ad3c5bf"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "7c507314-4946-4b61-8399-c240fb078dfe"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "0bc8b68f-f4a4-4f86-9c5f-39d22b756c2e"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "70c32c47-0b0e-4c77-9881-a01d9358fe10"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "10c367c8-18bf-40bf-89c0-a94a908a7294"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "04557516-3a3f-472f-9c78-62c1265672ba"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "1a567d97-2069-40a5-a43e-4b7e1c47ca90"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "84e7f50c-e6fa-4485-8a75-6b891dc989af"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "0fbd1b70-6430-4ef7-9b88-3b7bb1ea09ab"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "7ef99a53-94c7-4842-a019-e7cf5e1eb5a9"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "efc434ce-edaf-45c2-af19-42a6f2831ce3"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "2db19aa0-3c45-4f6c-ab1e-d82276e9ee4b"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "c890f767-2443-4dd3-a7e4-8e56a78a0e98"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "cc4d73cc-5887-4c22-8177-e870a6b32ac5"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "7cd82f73-d0ef-4867-970b-d5a7e471471d"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "38ba1af4-0029-4e3e-9887-68f48ef99a56"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "286d1591-4131-4d5a-8340-2612c982ec19"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "3808df69-60c6-4406-87f9-184748984bf7"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "1ad9a79d-e9c7-4401-b549-66a726cb9034"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "82701892-0db0-4afb-8bf1-ca788d29f123"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "588c2abc-626b-4ecc-86d8-437f66def6c3"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "b23b9fa4-92bc-4c2e-956c-9538abae6115"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "f1048856-0200-43ff-921e-21e5a850952a"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "ddb8c160-dead-4353-9b49-3f3d84153f60"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "e39cc0d2-9ca0-4d80-91c6-5651706ddaf6"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "b59fdcff-9841-444a-8400-c34fd7321cd4"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "5a4253ac-caaa-4520-af8a-6491c07a7d10"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "b23b9fa4-92bc-4c2e-956c-9538abae6115"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "6ac3e1bf-61b3-4a17-81da-6bba1ad799e0"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "c1089a71-d84e-44f2-9809-5cbabc5feed4"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "1723dd46-a836-4063-bbab-42ee80fe4c89"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "059ec6b3-b4d0-4ebe-a5bf-f96354ad0bb5"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "fa6d20ee-eff6-4248-9dbc-eeb473b1cca0"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "df026d83-f477-4cce-9267-4ecf2b00db8b"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "acbf2faa-1bf5-463f-9c7a-7b3f4f800ad9"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "a20d23c8-a56d-4cfa-8971-9d9a7ca0bb66"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "d3cf75b2-66d7-4969-86a5-532b6c02c2d2"
    },
    {
      "source": "114ed2be-4ddb-47b0-a597-41f7d240280d",
      "destination": "7d48a5d1-024d-42f9-82a5-3f2328d8e2e4"
    }
  ],
  "nodes": {
    "04557516-3a3f-472f-9c78-62c1265672ba": {
      "hostname": "",
      "ips": [
        "40.97.92.34"
      ],
      "processes": [
        "http",
        "pop3",
        "https",
        "imap",
        "submission",
        "imaps",
        "pop3s",
        "smtp"
      ],
      "openports": [
        110,
        143,
        587,
        993,
        995,
        443,
        25,
        80
      ]
    },
    "059ec6b3-b4d0-4ebe-a5bf-f96354ad0bb5": {
      "hostname": "server-13-249-50-31.iah50.r.cloudfront.net",
      "ips": [
        "13.249.50.31"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "0bc8b68f-f4a4-4f86-9c5f-39d22b756c2e": {
      "hostname": "",
      "ips": [
        "::"
      ],
      "processes": [
        "Unknown"
      ],
      "openports": [

      ]
    },
    "0fbd1b70-6430-4ef7-9b88-3b7bb1ea09ab": {
      "hostname": "edge-star-mini-shv-02-dfw5.facebook.com",
      "ips": [
        "31.13.93.35"
      ],
      "processes": [
        "https",
        "unknown",
        "xmpp-client",
        "http"
      ],
      "openports": [
        443,
        80,
        843,
        5222
      ]
    },
    "10c367c8-18bf-40bf-89c0-a94a908a7294": {
      "hostname": "ec2-18-211-118-21.compute-1.amazonaws.com",
      "ips": [
        "18.211.118.21"
      ],
      "processes": [
        "http",
        "https"
      ],
      "openports": [
        80,
        443
      ]
    },
    "114ed2be-4ddb-47b0-a597-41f7d240280d": {
      "hostname": "DESKTOP-BS53P6P",
      "ips": [
        "169.254.47.139/16",
        "169.254.139.177/16",
        "169.254.212.105/16",
        "fe80::58ab:5035:3750:4e9d/64",
        "10.102.37.150/30",
        "fe80::908c:e0fe:2eb0:bf19/64",
        "fe80::385c:aec0:776e:8bb1/64",
        "fe80::2594:e28e:b3f2:d469/64",
        "fe80::f5a0:535b:e573:305d/64",
        "169.254.48.93/16",
        "::1/128",
        "fe80::a9f1:b44f:d622:8785/64",
        "169.254.135.133/16",
        "fe80::5cf2:9aa8:63b4:2f8b/64",
        "169.254.78.157/16",
        "192.168.50.160/24",
        "127.0.0.1/8"
      ],
      "processes": [
        "1576/svchost.exe",
        "4740/svchost.exe",
        "3492/svchost.exe",
        "4304/svchost.exe",
        "12072/AppleMobileDeviceProcess.exe",
        "17472/SearchUI.exe",
        "7444/svchost.exe",
        "10880/NVIDIA Web Helper.exe",
        "25096/Code.exe",
        "18580/nvcontainer.exe",
        "3284/svchost.exe",
        "932/lsass.exe",
        "840/wininit.exe",
        "1540/svchost.exe",
        "23836/chrome.exe",
        "13360/chrome.exe",
        "6772/nvcontainer.exe",
        "18616/slack.exe",
        "23800/MsMpEng.exe",
        "4640/nvcontainer.exe",
        "12096/OneDrive.exe",
        "20688/NVIDIA Web Helper.exe",
        "1160/svchost.exe",
        "17632/OneDrive.exe",
        "912/services.exe",
        "7784/ksde.exe",
        "3420/Telegram.exe",
        "4/System",
        "1868/svchost.exe",
        "3768/spoolsv.exe",
        "16648/Code.exe",
        "676/msedge.exe",
        "2252/svchost.exe",
        "13748/msedge.exe",
        "19236/tester.exe"
      ],
      "openports": [
        7680,
        49669,
        58496,
        58778,
        58799,
        10020,
        546,
        58745,
        49666,
        138,
        54395,
        56297,
        54398,
        54396,
        58700,
        58749,
        58796,
        58520,
        58759,
        60953,
        58782,
        58787,
        58792,
        5050,
        64265,
        7001,
        58798,
        58405,
        58775,
        58788,
        58791,
        58776,
        5353,
        54393,
        58494,
        58686,
        58747,
        58756,
        5355,
        58530,
        58754,
        58757,
        1900,
        53111,
        58556,
        58735,
        58783,
        58794,
        54397,
        58733,
        58743,
        58793,
        137,
        58526,
        58734,
        58761,
        58764,
        58768,
        49744,
        65001,
        58753,
        58773,
        58780,
        58751,
        58762,
        58769,
        52365,
        53734,
        60956,
        445,
        58748,
        58760,
        45759,
        58480,
        58693,
        58744,
        135,
        49665,
        58687,
        58771,
        58779,
        49664,
        58495,
        58732,
        58740,
        58750,
        5040,
        139,
        58767,
        58790,
        58800,
        49668,
        68,
        58497,
        58731,
        58772,
        58777,
        58746,
        58765,
        58766,
        49667,
        22910,
        58758,
        58781,
        58797,
        58755,
        60955,
        54394,
        27015,
        58737,
        58774,
        52370,
        58493,
        52367,
        58528,
        58727,
        58785,
        10010,
        60954,
        58770,
        58795,
        58739,
        58741,
        58784,
        56299,
        58742,
        58789
      ]
    },
    "1723dd46-a836-4063-bbab-42ee80fe4c89": {
      "hostname": "",
      "ips": [
        "151.101.184.134"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "1a567d97-2069-40a5-a43e-4b7e1c47ca90": {
      "hostname": "www-ubuntu-com.avocado.canonical.com",
      "ips": [
        "91.189.90.58"
      ],
      "processes": [
        "http",
        "https"
      ],
      "openports": [
        80,
        443
      ]
    },
    "1ad9a79d-e9c7-4401-b549-66a726cb9034": {
      "hostname": "",
      "ips": [
        "0.0.0.0"
      ],
      "processes": [
        "Unknown"
      ],
      "openports": [

      ]
    },
    "2606a1fc-77a1-45d1-b69a-db9d6d11adcc": {
      "hostname": "",
      "ips": [
        "172.217.195.188"
      ],
      "processes": [
        "http",
        "https",
        "hpvroom"
      ],
      "openports": [
        80,
        443,
        5228
      ]
    },
    "286d1591-4131-4d5a-8340-2612c982ec19": {
      "hostname": "108-174-10-10.fwd.linkedin.com",
      "ips": [
        "108.174.10.10"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "2db19aa0-3c45-4f6c-ab1e-d82276e9ee4b": {
      "hostname": "ec2-34-196-111-83.compute-1.amazonaws.com",
      "ips": [
        "34.196.111.83"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "373e25ac-17d9-48c2-b7b0-6a611ad3c5bf": {
      "hostname": "acai.canonical.com",
      "ips": [
        "91.189.88.22"
      ],
      "processes": [
        "https",
        "http",
        "rsync",
        "nrpe",
        "jetdirect",
        "ftp",
        "ssh",
        "smtp"
      ],
      "openports": [
        25,
        443,
        873,
        5666,
        9103,
        80,
        21,
        22
      ]
    },
    "3808df69-60c6-4406-87f9-184748984bf7": {
      "hostname": "dfw06s48-in-f14.1e100.net",
      "ips": [
        "216.58.194.110"
      ],
      "processes": [
        "http",
        "https"
      ],
      "openports": [
        80,
        443
      ]
    },
    "38ba1af4-0029-4e3e-9887-68f48ef99a56": {
      "hostname": "a104-73-165-57.deploy.static.akamaitechnologies.com",
      "ips": [
        "104.73.165.57"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        80,
        443
      ]
    },
    "588c2abc-626b-4ecc-86d8-437f66def6c3": {
      "hostname": "",
      "ips": [
        "20.36.219.28"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "5a4253ac-caaa-4520-af8a-6491c07a7d10": {
      "hostname": "108-174-10-14.fwd.linkedin.com",
      "ips": [
        "108.174.10.14"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        80,
        443
      ]
    },
    "6ac3e1bf-61b3-4a17-81da-6bba1ad799e0": {
      "hostname": "",
      "ips": [
        "192.0.76.3"
      ],
      "processes": [
        "http",
        "https"
      ],
      "openports": [
        443,
        80
      ]
    },
    "70c32c47-0b0e-4c77-9881-a01d9358fe10": {
      "hostname": "",
      "ips": [
        "195.122.177.227"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "7c507314-4946-4b61-8399-c240fb078dfe": {
      "hostname": "",
      "ips": [
        "199.16.172.121"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "7cd82f73-d0ef-4867-970b-d5a7e471471d": {
      "hostname": "",
      "ips": [
        "149.154.167.50"
      ],
      "processes": [
        "xmpp-client",
        "sun-answerbook",
        "http",
        "https"
      ],
      "openports": [
        443,
        80,
        5222,
        8888
      ]
    },
    "7d48a5d1-024d-42f9-82a5-3f2328d8e2e4": {
      "hostname": "",
      "ips": [
        "151.101.192.134"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "7ef99a53-94c7-4842-a019-e7cf5e1eb5a9": {
      "hostname": "",
      "ips": [
        "40.121.3.131"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "82701892-0db0-4afb-8bf1-ca788d29f123": {
      "hostname": "",
      "ips": [
        "52.230.222.68"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "84e7f50c-e6fa-4485-8a75-6b891dc989af": {
      "hostname": "a23-200-43-16.deploy.static.akamaitechnologies.com",
      "ips": [
        "23.200.43.16"
      ],
      "processes": [
        "http",
        "https"
      ],
      "openports": [
        443,
        80
      ]
    },
    "a20d23c8-a56d-4cfa-8971-9d9a7ca0bb66": {
      "hostname": "assets.ubuntu.com",
      "ips": [
        "162.213.33.101"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        80,
        443
      ]
    },
    "a2ebface-b962-4f45-80a6-e0e4c49d2e55": {
      "hostname": "",
      "ips": [
        "151.101.184.157"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "acbf2faa-1bf5-463f-9c7a-7b3f4f800ad9": {
      "hostname": "",
      "ips": [
        "149.154.175.53"
      ],
      "processes": [
        "https",
        "http",
        "xmpp-client",
        "sun-answerbook"
      ],
      "openports": [
        5222,
        8888,
        443,
        80
      ]
    },
    "b23b9fa4-92bc-4c2e-956c-9538abae6115": {
      "hostname": "wordpress.com",
      "ips": [
        "192.0.77.32"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "b59fdcff-9841-444a-8400-c34fd7321cd4": {
      "hostname": "xx-fbcdn-shv-02-dfw5.fbcdn.net",
      "ips": [
        "31.13.93.26"
      ],
      "processes": [
        "unknown",
        "xmpp-client",
        "http",
        "https"
      ],
      "openports": [
        80,
        843,
        5222,
        443
      ]
    },
    "c1089a71-d84e-44f2-9809-5cbabc5feed4": {
      "hostname": "",
      "ips": [
        "104.16.76.166"
      ],
      "processes": [
        "https",
        "http-proxy",
        "https-alt",
        "http"
      ],
      "openports": [
        8080,
        8443,
        443,
        80
      ]
    },
    "c890f767-2443-4dd3-a7e4-8e56a78a0e98": {
      "hostname": "e8.cc.36a9.ip4.static.sl-reverse.com",
      "ips": [
        "169.54.204.232"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "cc4d73cc-5887-4c22-8177-e870a6b32ac5": {
      "hostname": "a-0001.a-msedge.net",
      "ips": [
        "204.79.197.200"
      ],
      "processes": [
        "domain",
        "http",
        "https"
      ],
      "openports": [
        443,
        53,
        80
      ]
    },
    "d3cf75b2-66d7-4969-86a5-532b6c02c2d2": {
      "hostname": "",
      "ips": [
        "104.244.42.5"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "ddb8c160-dead-4353-9b49-3f3d84153f60": {
      "hostname": "dfw06s49-in-f19.1e100.net",
      "ips": [
        "216.58.194.147"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        80,
        443
      ]
    },
    "df026d83-f477-4cce-9267-4ecf2b00db8b": {
      "hostname": "",
      "ips": [
        "104.19.198.151"
      ],
      "processes": [
        "https",
        "http",
        "http-proxy",
        "https-alt"
      ],
      "openports": [
        443,
        80,
        8080,
        8443
      ]
    },
    "e39cc0d2-9ca0-4d80-91c6-5651706ddaf6": {
      "hostname": "ec2-18-233-219-152.compute-1.amazonaws.com",
      "ips": [
        "18.233.219.152"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "efc434ce-edaf-45c2-af19-42a6f2831ce3": {
      "hostname": "",
      "ips": [
        "13.107.136.254"
      ],
      "processes": [
        "https",
        "domain",
        "http"
      ],
      "openports": [
        443,
        53,
        80
      ]
    },
    "f1048856-0200-43ff-921e-21e5a850952a": {
      "hostname": "",
      "ips": [
        "192.0.73.2"
      ],
      "processes": [
        "https",
        "http"
      ],
      "openports": [
        443,
        80
      ]
    },
    "fa6d20ee-eff6-4248-9dbc-eeb473b1cca0": {
      "hostname": "",
      "ips": [
        "104.19.148.8"
      ],
      "processes": [
        "https",
        "http",
        "http-proxy",
        "https-alt"
      ],
      "openports": [
        8080,
        8443,
        443,
        80
      ]
    }
  }
}`)
