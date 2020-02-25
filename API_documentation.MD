FORMAT: 1A
HOST: https://polls.apiblueprint.org/

# Server2

Fetch backend server

## Storages [/storages/list]

### List all storage [GET]

+ Request

    + Headers

            Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJ1c2VybmFtZTEiLCJuYW1lIjoiYW5nbGluZyB6Iiwicm9sZSI6Im93bmVyIiwidGltZXN0YW1wIjoiMTU4MjU0MDczNSIsImV4cCI6MTYxNDA3NzM3NH0.SNyTFamGc_82QZ7D3gky-pmKqXOwdGxEd4SBYIpCqzQ
        
+ Response 200 (application/json)

        [
            {
                "uuid":null,
                "komoditas":null,
                "area_provinsi":null,
                "area_kota":null,
                "size":null,"price":null,
                "tgl_parsed":null,"timestamp":null
            },
            {
                "uuid":"8a23fcab-ef67-48b8-8ba1-7055ea91ea3b",
                "komoditas":"Penaeus Vannamei",
                "area_provinsi":"BANTEN",
                "area_kota":"PANDEGLANG",
                "size":"40",
                "price":"72000",
                "tgl_parsed":"2019-11-11T17:00:00.000Z",
                "timestamp":"1573491600"
            },
            {
                "uuid":"8a23fcab-ef67-48b8-8ba1-7055ea91ea3b",
                "komoditas":"Penaeus Vannamei",
                "area_provinsi":"BANTEN",
                "area_kota":"PANDEGLANG",
                "size":"40",
                "price":"72000",
                "tgl_parsed":"2019-11-11T17:00:00.000Z",
                "timestamp":"1573491600"
            },
            ...
         ]

## Storage area [/storages/area]

### List all storage area option [GET]

+ Response 200 (application/json)

        [
            {
                "province": "BANTEN",
                "city": "PANDEGLANG"
            },
            {
                "province": "DKI JAKARTA",
                "city": null
            },
            {
                "province": "JAWA BARAT",
                "city": null
            },
            ...
         ]

## Storage size [/storages/size]

### List all storage size option [GET]

+ Response 200 (application/json)

        [
            {
                "size": "30"
            },
            {
                "size": "40"
            },
            {
                "size": "50"
            },
            ...
         ]

## JWT [/token/check]

### CLAIMS [GET]

Accept jwt token and return private claims if token is valid 

+ Request

    + Headers

            Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJ1c2VybmFtZTEiLCJuYW1lIjoiYW5nbGluZyB6Iiwicm9sZSI6Im93bmVyIiwidGltZXN0YW1wIjoiMTU4MjU0MDczNSIsImV4cCI6MTYxNDA3NzM3NH0.SNyTFamGc_82QZ7D3gky-pmKqXOwdGxEd4SBYIpCqzQ
        
+ Response 200 (application/json)

        {
            "private claims": {
                "exp": 1614077374,
                "name": "angling z",
                "role": "owner",
                "timestamp": "1582540735",
                "user_name": "username1"
            }
        }
