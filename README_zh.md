# Looking Glass

[English](README.md)

請至 [Releases](https://github.com/steveyiyo/LookingGlass/releases) 下載對應的二進制文件。

```
./LookingGlass {server, agent}
```

## Agent

需要設定以下資訊，在 Agent 啟動後會與 Server 自動建立 Socket 連線。
- Server EndPoint
- Server Authorized Key
- POP Name

## Server

需要設定以下資訊，如果沒有指定將會自動產生。
- Listen Port
- Authorized Key

## API EndPoint

POST    /api/v1/mtr/  
POST    /api/v1/ping/  
POST	/api/v1/bgp_route/  

```
{
    "PoP": "",              # POP Name.
    "Dst_IP": ""            # IP Address for enddevices.
}
```

POST /api/v1/admin/AddPoP/

```
{
    "Authorized_Key": "",   # Which you write in .env file.
    "PoP": "",              # The POP Name that you want to name.
    "MGMT_IP": ""           # The IP address of the POP.
}
```