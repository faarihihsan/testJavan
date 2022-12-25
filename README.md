## What's inside
- assessment for number 1 is in file ```keluarga.sql```
- assessment for number 2 & 3 can be tested by running the app

## How to Run
1. initiate database by running ```docker compose up```
2. create table using file ```keluarga.sql```
3. run the app by running ```go build && ./testJavan```

## HTTP protocol endpoints
### Get list keluarga
- endpoint: ```localhost:8089/get-list-keluarga``` 

### Add keluarga 
- endpoint: ```localhost:8089/add-keluarga```
- example body message
```json
{
    "nama": "Budi",
    "parent": 1

}
```

### Update keluarga
- endpoint: ```localhost:8089/update-keluarga```
- example body message
```json
{
  "id" : 2,
  "nama": "Budi",
  "parent": 10

}
```

### Delete keluarga
- endpoint: ```localhost:8089/delete-keluarga```
- example body message
```json
{
  "id" : 2

}
```
### Add aset
- endpoint: ```localhost:8089/add-aset```
- example body message
```json
{
  "nama": "Huawei P30"
}
```
### Update aset
### Add aset
- endpoint: ```localhost:8089/update-aset```
- example body message
```json
{
  "id": 1,
  "nama": "Huawei P29"
}
```
### Delete aset
- endpoint: ```localhost:8089/delete-aset```
- example body message
```json
{
  "id": 1
}
```
### Add aset keluarga
- endpoint: ```localhost:8089/add-aset-keluarga```
- example body message
```json
{
  "idKeluarga": 1,
  "idAset": 2
}
```
### Delete aset keluarga
- endpoint: ```localhost:8089/delete-aset-keluarga```
- example body message
```json
{
  "id": 1
}
```

## Hot to use TCP protocol
1. run ```tcpClient``` in terminal by running ```go run tcpClient.go```
2. run command using below format
   1. Add keluarga ```ADD KELUARGA {{nama}} {{id_parent}}```
   2. Update keluarga ```UPDATE KELUARGA {{id}} {{nama}} {{id_parent}}```
   3. Delete keluarga ```DELETE KELUARGA {{id}}```
   4. Add aset ```ADD ASET {{nama}}```
   5. Update aset ```UPDATE ASET {{id}} {{nama}}```
   6. Delete aset ```DELETE ASET {{id}}```
   7. Add aset keluarga ```ADD ASET_KELUARGA {{id_keluarga}} {{id_aset}}```
   8. Delete aset keluarga ```DELETE ASET_KELUARGA {{id}}```
   9. Get list keluarga ```GET LIST KELUARGA```
