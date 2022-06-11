REST API GOLANG CRUD - ANIMAL
- Storage system menggunakan automigrate, maka tidak perlu membuat tahap inisialisasi database seperti mengeksekusi query, membuat Table, dsb. Hal itu dikarenakan method bernama AutoMigrate akan melakukan proses migration terhadap semua entitas yang akan digunakan pada database ketika kode pertama kali dijalankan.

TOOLS
# REST API GOLANG CRUD - ANIMAL
### REST-API-CRUD-GO (GIN FRAMEWORK)-GORM-MYSQL

Storage system menggunakan automigrate, maka tidak perlu membuat tahap inisialisasi database seperti mengeksekusi query, membuat Table, dsb. Hal itu dikarenakan method bernama AutoMigrate akan melakukan proses migration terhadap semua entitas yang akan digunakan pada database ketika kode pertama kali dijalankan.

## TOOLS
1. Go v1.18.3 (Gin Framework)
2. GORM
3. MySQL
4. Postman

## RUN PROGRAM
cd testgo
go run main.go

## 1. API MENYIMPAN ANIMAL (POST)
### Input :
```
{
    "name": "lion",
    "class": "mammal",
    "legs" : "4"
}
```
### Request :

http://localhost:8080/animals

### Response : (Status Code : 200)
```
{
    "data": {
        "id": 3,
        "name": "lion",
        "class": "mammal",
        "legs": "4"
    }
}
```
### Image :

![POST](https://user-images.githubusercontent.com/71954147/173164714-30de04bf-f403-4161-9f9f-6059a0629ac4.png)

## 2. API MEMPERBARUI ANIMAL (PUT)
### Input :
```
{
    "name": "harimau",
    "class": "mammal",
    "legs" : "4"
}
```
### Request :

http://localhost:8080/animals/6

### Response : (Status Code : 200)
```
{
    "data": {
        "id": 6,
        "name": "harimau",
        "class": "mammal",
        "legs": "4"
    }
}
```
### Image :

![PUT](https://user-images.githubusercontent.com/71954147/173164727-4f4285e7-1cc2-47c0-9d17-29f4ca9fb0a2.png)

## 3. API MENGHAPUS ANIMAL (DELETE)
### Request :

http://localhost:8080/animals/7

### Response : (Status Code : 200)
```
{
    "data": true
}
```
### Image :

![DELETE](https://user-images.githubusercontent.com/71954147/173164737-d1726340-efae-40bc-ae03-9b664337e2a3.png)

## 4. API MENAMPILKAN SEMUA ANIMAL (GET)
### Request :

http://localhost:8080/animals

### Response : (Status Code : 200)
```
{
    "data": [
        {
            "id": 4,
            "name": "rusa",
            "class": "mammal",
            "legs": "4"
        },
        {
            "id": 5,
            "name": "lion",
            "class": "mammal",
            "legs": "4"
        },
        {
            "id": 6,
            "name": "harimau",
            "class": "mammal",
            "legs": "4"
        }
    ]
}
```
### Image :

![GET all](https://user-images.githubusercontent.com/71954147/173164765-d07621df-4cdd-4ee8-aaec-f1cd7fdc5937.png)

## 5. API MENAMPILKAN ANIMAL (GET)
### Request :

http://localhost:8080/animals/4

### Response : (Status Code : 200)
```
{
    "data": {
        "id": 4,
        "name": "rusa",
        "class": "mammal",
        "legs": "4"
    }
}
```
### Image :

![GET id](https://user-images.githubusercontent.com/71954147/173164773-1c6cb723-eaa9-47d7-a55c-d1777b61fdaf.png)
