# cloud-infra-generator-api

git clone https://github.com/emirhandogandemir/cloud-infra-generator-api.git

go mod download

go run main.go

`localhost:7070`

![image](https://github.com/emirhandogandemir/cloud-infra-generator-api/assets/74687192/b91d1b32-fff3-4f1f-9247-516c54cea731)

![image](https://github.com/emirhandogandemir/cloud-infra-generator-api/assets/74687192/b59b7d63-4def-486a-b3d5-50601d97d192)
![image](https://github.com/emirhandogandemir/cloud-infra-generator-api/assets/74687192/8933554c-a7e2-42b1-a735-ef2d04370606)

![image](https://github.com/emirhandogandemir/cloud-infra-generator-api/assets/74687192/93cfc381-91d3-4943-8822-f554f6b1340f)

![image](https://github.com/emirhandogandemir/cloud-infra-generator-api/assets/74687192/ad034d23-b6e3-4bb2-89d0-e18fcb6ce3a7)


#Goals
Tüm kaynakları görebilecegimiz bir dashboard vm sayımızızı vmIdlerimizi, kubernetes clusterımızın bilgilerini gibi ana ana ana sayfada tüm cloud providerlar için All resources sayfası koyulması
/createaws endpointinde imageId alanı inputuna ihtiyacımız var
/getinstancetypesaws array dönüyor onun içindekileri virtualMachines tabi koyup hem vmlistemizi hemde seçilebilecek ilerisi için imageTypesları alırız. İleride ImageIdlist diye de ekliyor olacagız
/getbillingaws = buradada billgindeki start ve end tarihleri için boş olsada 2 alan girilecek bir alttaki method ile ay ay fatura tutarlarını dönüyor Billing sekmesi koyup oradan amazon ve azure fatura tutarlarının grafiklerini görüyor olacagız
/getbillingazure => yukarıdaki ile aynı
/createeks = 2 tane string input alanı gereklidir
/createnodegroupaws endpointi için => 6 tane alana ihtiyacımız var
/createvmazure => 6 tane string deger alacak alana ihtiyacımı var
