package model

import (
	"restoran/conn"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Pesanan structure
type Pesanan struct {
	IDPesanan  int64     `bson:"id_pesanan"`
	Detail     []detail  `bson:"detail"`
	TotalHarga int64     `bson:"total_harga"`
	CreatedAt  time.Time `bson:"created_at"`
}

type detail struct {
	Makanan string `bson:"makanan"`
	Jumlah  int    `bson:"jumlah"`
	Harga   int    `bson:"harga"`
}

//ListPesanan declaration
type ListPesanan []Pesanan

//PesananInfo function
func PesananInfo(id int64, pesananCollection string) (Pesanan, error) {
	db := conn.GetMongoDB()
	pesanan := Pesanan{}
	err := db.C(pesananCollection).Find(bson.M{"id_pesanan": &id}).One(&pesanan)

	return pesanan, err
}
