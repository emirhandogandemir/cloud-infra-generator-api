# cloud-infra-generator-api

git clone https://github.com/emirhandogandemir/cloud-infra-generator-api.git

go mod download

go run main.go

`localhost:7070`

#Goals
- Tüm kaynakları görebilecegimiz bir dashboard vm sayımızızı vmIdlerimizi, kubernetes clusterımızın bilgilerini gibi ana ana ana sayfada tüm cloud providerlar için All resources sayfası koyulması
- /createaws endpointinde imageId alanı inputuna ihtiyacımız var
- /getinstancetypesaws array dönüyor onun içindekileri virtualMachines tabi koyup hem vmlistemizi hemde seçilebilecek ilerisi için imageTypesları alırız. İleride ImageIdlist diye de ekliyor olacagız
- /getbillingaws = buradada billgindeki start ve end tarihleri için boş olsada 2 alan girilecek bir alttaki method ile ay ay fatura tutarlarını dönüyor Billing sekmesi koyup oradan amazon ve azure fatura tutarlarının    grafiklerini görüyor olacagız
- /getbillingazure => yukarıdaki ile aynı
- /createeks = 2 tane string input alanı gereklidir
- /createnodegroupaws endpointi için => 6 tane alana ihtiyacımız var
- /createvmazure => 6 tane string deger alacak alana ihtiyacımı var
- azure-aws ortamlarının kubernetes cluster monitoringini saglamak pod sayıları nod sayıları gibi verileri
- secret tarafındaki strateji geliştirilecek = user bazlı secret credentialsların dbde tutulması
- Bucket S3 and blob storage endpoints=>

User Json
`{
  "username": "John Doe",
  "email": "johndoe@example.com",
  "password": "mypassword",
  "aws_accesses": [
    {
      "accessKey": "accesskey1",
      "secretKey": "secretkey1"
    },
    {
      "accessKey": "accesskey2",
      "secretKey": "secretkey2"
    }
  ]
}`

![image](https://github.com/emirhandogandemir/cloud-infra-generator-api/assets/74687192/b3581b88-691a-42c1-9f29-41728a109c1f)

