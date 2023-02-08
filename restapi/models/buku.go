package models

import (
	"html"
	"strings"
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"restapi/utils/token"
)

type Buku struct {
	gorm.Model
	JudulBuku string `gorm:"size:255;not null;unique" json:"judul_buku"`
	Deskripsi string `gorm:"size:255;not null;" json:"deskripsi"`
	IDPenulis string `gorm:"size:255;not null;" json:"id_penulis"`
	IDKategori string `gorm:"size:255;not null;" json:"id_kategori"`
	TahunTerbit string `gorm:"size:255;not null;" json:"tahun_terbit"`
}

type Penulis struct {
	gorm.Model
	NamaPenulis string `gorm:"size:100;not null;unique" json:"nama_penulis"`
	Alamat string `gorm:"size:255;not null;" json:"alamat"`
}

type Kategori struct {
	gorm.Model
	Kategori string `gorm:"size:50;not null;unique" json:"kategori"`
	Deskripsi string `gorm:"size:255;not null;" json:"deskripsi"`
}

func GetBukuByJudulBuku(judulbuku string) (Buku,error) {

	var buku []Buku

	if err = db.Joins("LEFT JOIN penulis ON buku.id_penulis = penulis.id").
		Joins("LEFT JOIN kategori ON buku.id_kategori = kategori.id").
		Where("buku.judul_buku=?", judulbuku)..
		Find(&buku).Error; err != nil {
		log.Fatal(err)
	}
	
	return buku,nil

}

func GetBukuByNamaPenulis(namapenulis string) (Buku,error) {

	var buku []Buku

	if err = db.Joins("LEFT JOIN penulis ON buku.id_penulis = penulis.id").
		Joins("LEFT JOIN kategori ON buku.id_kategori = kategori.id").
		Where("penulis.nama_penulis=?", namapenulis)..
		Find(&buku).Error; err != nil {
		log.Fatal(err)
	}
	
	return buku,nil

}

func GetBukuByKategori(kategori string) (Buku,error) {

	var buku []Buku

	if err = db.Joins("LEFT JOIN penulis ON buku.id_penulis = penulis.id").
		Joins("LEFT JOIN kategori ON buku.id_kategori = kategori.id").
		Where("kategori.kategori=?", kategori)..
		Find(&buku).Error; err != nil {
		log.Fatal(err)
	}
	
	return buku,nil

}

func (u *Buku) SaveBuku() (*Buku, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Buku{}, err
	}
	return u, nil
}