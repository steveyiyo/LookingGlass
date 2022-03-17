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
    "PoP": "",             # The POP Name that you want to name.
    "MGMT_IP": ""           # The IP address of the POP.
}
```