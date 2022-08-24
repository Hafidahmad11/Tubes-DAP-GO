package main

import (
	"fmt"
	"os"
	"os/exec"
)

var pil2 int
var all int

type (
	Camaba struct {
		nama        string
		kelamin     string
		prodiTek    [3]ProdiTeknik
		prodiNonTek [3]ProdiNonTeknik
		nilai       Nilaicalon
	}
	Nilaicalon struct {
		mtk, ipa, ips, eng float64
		tpa, tbi           float64
	}
	ProdiTeknik struct {
		nameJur, akreditasi string
		num                 int
	}
	ProdiNonTeknik struct {
		nameJur, akreditasi string
		num                 int
	}
)
type ArrCamaba [100]Camaba
type ProdTek [100]ProdiTeknik
type ProdNonTek [100]ProdiNonTeknik

//BATAS TIPE BENTUKAN//
var data ArrCamaba
var PrTek ProdTek
var PrNon ProdNonTek
var numCamaba int
var numTek int
var numSos int
var numPemilih int

func Create(data *ArrCamaba) {
	var prodi string
	fmt.Print("Nama	: ")
	fmt.Scan(&data[all].nama)
	fmt.Print("jenis Kelamin Laki-Laki[L] / Perempuan[P] : ")
	fmt.Scan(&data[all].kelamin)
	fmt.Print("Nilai Matematika : ")
	fmt.Scan(&data[all].nilai.mtk)
	fmt.Print("Nilai Bahasa Inggris : ")
	fmt.Scan(&data[all].nilai.eng)
	fmt.Print("Nilai Tes Potensi Akademik : ")
	fmt.Scan(&data[all].nilai.tpa)
	fmt.Print("Nilai Tes Bahasa Inggris : ")
	fmt.Scan(&data[all].nilai.tbi)
	fmt.Print("Pilih [Teknik] / [NonTek] (Perhatikan Besar Kecil Huruf ) : ")
	fmt.Scan(&prodi)
	if prodi == "Teknik" {
		fmt.Print("Nilai IPA : ")
		fmt.Scan(&data[all].nilai.ipa)
		//syarat LULUS ITU LEBIH DARI 75 UNTUK SETIAP MATA PEL TEKNIK//
		for j := 0; j < 3; j++ {
			fmt.Print("Pilihan Program Studi ", j+1, " : ")
			fmt.Scan(&data[all].prodiTek[j].nameJur)
			PrTek[j].nameJur = data[all].prodiTek[j].nameJur
			PrTek[j].num++
			if data[all].nilai.mtk >= 75 && data[all].nilai.ipa >= 75 && data[all].nilai.eng >= 75 {
				fmt.Println("LULUS")
			} else if data[all].nilai.mtk <= 75 && data[all].nilai.ipa <= 75 && data[all].nilai.eng <= 75 {
				fmt.Println("TIDAK LULUS")
			}
		}
		numTek++
	} else if prodi == "NonTek" {
		fmt.Print("Nilai IPS : ")
		fmt.Scan(&data[all].nilai.ips)
		for j := 0; j < 3; j++ {
			fmt.Print("Pilihan Program Studi ", j+1, " : ")
			fmt.Scan(&data[all].prodiNonTek[j].nameJur)
			PrNon[j].nameJur = data[all].prodiNonTek[j].nameJur
			PrNon[j].num++
			if data[all].nilai.mtk >= 75 && data[all].nilai.ipa >= 75 && data[all].nilai.eng >= 75 {
				fmt.Println("LULUS")
			} else if data[all].nilai.mtk <= 75 && data[all].nilai.ips <= 70 && data[all].nilai.eng <= 75 {
				fmt.Println("TIDAK LULUS")
			}

		}
		numSos++
	}
	numCamaba++
	all++
}

//MENAMPILKAN JUMLAH PRODI DARI SETIAP CAMABA,
//MAENAMPILKAN INI BERDASARKAN JUMLAH INPUTAN CAMABA
func tampilPRODI() {
	for i := 0; i < numTek; i++ {
		fmt.Println("Nama jurusan  :", PrTek[i].nameJur)
		fmt.Println("Jumlah Camaba :", PrTek[i].num)
	}

	for i := 0; i < numSos; i++ {
		fmt.Println("Nama jurusan  :", PrNon[i].nameJur)
		fmt.Println("Jumlah Camaba :", PrNon[i].num)
	}
}

func editProdi(data *ArrCamaba, hasil int) {
	//JIKA INGIN MERUBAH , UBAH SEMUANYA JIKA INGIN dan menampilkan ke menu inputan awal//
	// DI SEARCHING TEMBAK NAMANYA LALU UBAH SEMUA//
	var tanya, Prd string

	fmt.Println("Apakah ingin Merubah Prodi ?(Y/T)", data[hasil].prodiTek[0].nameJur)
	fmt.Scan(&tanya)
	if tanya == "Y" {
		fmt.Println("Prodi yang Di Ubah :", data[hasil].prodiTek[0].nameJur)
		fmt.Print("Prodi Baru :")
		fmt.Scan(&Prd)
		data[hasil].prodiTek[0].nameJur = Prd

	}
	fmt.Println("Apakah ingin Merubah Prodi ?(Y/T)", data[hasil].prodiTek[1].nameJur)
	fmt.Scan(&tanya)
	if tanya == "Y" {
		fmt.Println("Prodi yang Di Ubah :", data[hasil].prodiTek[1].nameJur)
		fmt.Print("Prodi Baru :")
		fmt.Scan(&Prd)
		data[hasil].prodiTek[1].nameJur = Prd
	}
	fmt.Println("Apakah ingin Merubah Prodi ?(Y/T)", data[hasil].prodiTek[2].nameJur)
	fmt.Scan(&tanya)
	if tanya == "Y" {
		fmt.Println("Prodi yang Di Ubah :", data[hasil].prodiTek[2].nameJur)
		fmt.Print("Prodi Baru :")
		fmt.Scan(&Prd)
		data[hasil].prodiTek[2].nameJur = Prd
	}

}

//edit biodata mahasiswa
func editBiodata(data *ArrCamaba) {
	var cari string
	var ketemu bool
	var prodi string

	fmt.Print("Nama Calon yang ingin diubah ? : ")
	fmt.Scan(&cari)
	for i := 0; i < numCamaba; i++ {
		if cari == data[i].nama {
			fmt.Println("Data Ditemukan")
			ketemu = true
			fmt.Print("Nama	: ")
			fmt.Scan(&data[i].nama)
			fmt.Print("jenis Kelamin Laki-Laki[L] / Perempuan[P] : ")
			fmt.Scan(&data[i].kelamin)
			fmt.Print("Nilai Matematika : ")
			fmt.Scan(&data[i].nilai.mtk)
			fmt.Print("Nilai Bahasa Inggris : ")
			fmt.Scan(&data[i].nilai.eng)
			fmt.Print("Nilai Tes Potensi Akademik : ")
			fmt.Scan(&data[i].nilai.tpa)
			fmt.Print("Nilai Tes Bahasa Inggris : ")
			fmt.Scan(&data[i].nilai.tbi)
			fmt.Print("Pilih Teknik / NonTek : ")
			fmt.Scan(&prodi)
			if prodi == "Teknik" {
				fmt.Print("Nilai IPA : ")
				fmt.Scan(&data[i].nilai.ipa)
				fmt.Print("Pilihan Program Studi ", 1, " : ")
				fmt.Scan(&data[i].prodiTek[0].nameJur)
				fmt.Print("Pilihan Program Studi ", 2, " : ")
				fmt.Scan(&data[i].prodiTek[1].nameJur)
				fmt.Print("Pilihan Program Studi ", 3, " : ")
				fmt.Scan(&data[i].prodiTek[2].nameJur)

			} else if prodi == "NonTek" {
				fmt.Print("Nilai IPS : ")
				fmt.Scan(&data[i].nilai.ips)
				fmt.Print("Pilihan Program Studi ", 1, " : ")
				fmt.Scan(&data[i].prodiNonTek[0].nameJur)
				fmt.Print("Pilihan Program Studi ", 2, " : ")
				fmt.Scan(&data[i].prodiNonTek[1].nameJur)
				fmt.Print("Pilihan Program Studi ", 3, " : ")
				fmt.Scan(&data[i].prodiNonTek[2].nameJur)
			}
		} else if ketemu != false {
			fmt.Println("Data Tidak Ditemukan")
			fmt.Println("")
		}
	}
}

func Searcha(data *ArrCamaba, cari string) int {
	fmt.Printf("mulai searching \n")
	fmt.Printf("----------------\n")

	var hasil, awal, tgh, akhir int

	awal = 0
	akhir = all - 1
	tgh = (awal + akhir) / 2
	for akhir >= tgh && cari != data[tgh].nama {
		if cari > data[tgh].nama {
			awal = tgh + 1
		} else {
			akhir = tgh - 1
		}
		tgh = (awal + akhir) / 2
	}
	if cari == data[tgh].nama {
		hasil = tgh
	}
	return hasil
}

func MenampilkanProdTek(data ArrCamaba) {
	for i := 0; i < numCamaba; i++ {
		fmt.Println("nama :", data[i].nama)
		fmt.Println("jenis kelamin :", data[i].kelamin)
		fmt.Println("Nilai MTK :", data[i].nilai.mtk)
		fmt.Println("Nilai B.Ing:", data[i].nilai.eng)
		fmt.Println("Nilai Tpa :", data[i].nilai.tpa)
		fmt.Println("Nilai Tbi :", data[i].nilai.tbi)

		if data[i].prodiTek[0].nameJur != "" {
			fmt.Println("Nilai IPA :", data[i].nilai.ipa)
			fmt.Println("PRODI 1 : ", data[i].prodiTek[0].nameJur)
			fmt.Println("PRODI 2 : ", data[i].prodiTek[1].nameJur)
			fmt.Println("PRODI 3 :", data[i].prodiTek[2].nameJur)
			fmt.Println("---------------")
		} else if data[i].prodiNonTek[0].nameJur != "" {
			fmt.Println("Nilai IPS :", data[i].nilai.ips)
			fmt.Println("PRODI 1 :", data[i].prodiNonTek[0].nameJur)
			fmt.Println("PRODI 2 :", data[i].prodiNonTek[1].nameJur)
			fmt.Println("PRODI 3 :", data[i].prodiNonTek[2].nameJur)
		}

	}
}

//mengurutkan data Mahasiswa dan Prodi//mengurutkan berdasarkan nilai TES POTENSI AKADEMIK//
func Sorting(data *ArrCamaba, all int) {
	fmt.Println("Proses Mengurutkan")
	for i := 0; i < all-1; i++ {
		for j := i + 1; j > 0 && data[j-1].nilai.tpa < data[j].nilai.tpa; j-- {
			temp := data[j-1]
			data[j-1] = data[j]
			data[j] = temp
		}
	}

	for i := 0; i < all; i++ {
		fmt.Println("Mahasiswa ke ", i+1)
		fmt.Println("Nama 	: ", data[i].nama)
		fmt.Println("Nilai Tes Potensi Akademik : ", data[i].nilai.tpa)
		fmt.Println("Nilai Tes Bahasa Inggris :", data[i].nilai.tbi)
		fmt.Println("======================")

	}
}

//menghapus Prodi, Data Mahasiswa, Nilai
func delete(data *ArrCamaba, hasil int) {
	var tanya, blank, prodi, prodi1 string

	blank = " "
	fmt.Println("Apakah ingin menghapus nama maba ? (Y/T)")
	fmt.Scan(&tanya)
	if tanya == "Y" {
		fmt.Println("nama yang dihapus :", data[hasil].nama)
		data[hasil].nama = blank
	}
	fmt.Println("Apakah ingin menghapus jenis kelamin ?(Y/T)")
	fmt.Scan(&tanya)
	if tanya == "Y" {
		fmt.Println("jenis kelamin yang di hapus: ", data[hasil].kelamin)
		data[hasil].kelamin = blank
	}
	fmt.Println("Apakah ingin menghapus nilai MTK ?(Y/T)")
	fmt.Scan(&tanya)
	if tanya == "Y" {
		fmt.Println("nilai yang di hapus :", data[hasil].nilai.mtk)
		data[hasil].nilai.mtk = 0
	}
	fmt.Println("Apakah ingin menghapus nilai B.Ing ?(Y/T)")
	fmt.Scan(&tanya)
	if tanya == "Y" {
		fmt.Println("nilai yang di hapus :", data[hasil].nilai.eng)
		data[hasil].nilai.eng = 0
	}
	fmt.Println("Apakah ingin menghapus nilai Tes Potensi Akademik ?(Y/T)")
	fmt.Scan(&tanya)
	if tanya == "Y" {
		fmt.Println("nilai yang di hapus :", data[hasil].nilai.tpa)
		data[hasil].nilai.tpa = 0
	}
	fmt.Println("Apakah ingin menghapus nilai Tes Bahasa Inggris ?(Y/T)")
	fmt.Scan(&tanya)
	if tanya == "Y" {
		fmt.Println("nilai yang di hapus :", data[hasil].nilai.tbi)
		data[hasil].nilai.tbi = 0
	}
	fmt.Print("Pilih [Teknik] / [NonTek] (Perhatikan Besar Kecil Huruf ) : ")
	fmt.Scan(&prodi)
	if prodi == "Teknik" {
		fmt.Println("nilai yang di hapus :", data[hasil].nilai.ipa)
		data[hasil].nilai.ipa = 0
	} else if prodi == "NonTek" {
		fmt.Println("nilai yang di hapus :", data[hasil].nilai.ips)
		data[hasil].nilai.ips = 0
	}
	fmt.Print("Pilih ingin hapus Prodi ?(Teknik/NonTek)")
	fmt.Scan(&prodi1)
	if prodi1 == "Teknik" {
		fmt.Print(data[hasil].prodiTek[0].nameJur, " , ")
		fmt.Println("Prodi yang di hapus :", data[hasil].prodiTek[0].nameJur)
		data[hasil].prodiTek[0].nameJur = blank

		fmt.Print(data[hasil].prodiTek[1].nameJur, " , ")
		fmt.Println("Prodi yang di hapus :", data[hasil].prodiTek[1].nameJur)
		data[hasil].prodiTek[1].nameJur = blank

		fmt.Print(data[hasil].prodiTek[2].nameJur, " , ")
		fmt.Println("Prodi yang di hapus :", data[hasil].prodiTek[2].nameJur)
		data[hasil].prodiTek[2].nameJur = blank
	} else if prodi1 == "NonTek" {
		fmt.Print(data[hasil].prodiNonTek[0].nameJur, " , ")
		fmt.Println("Prodi yang di hapus :", data[hasil].prodiNonTek[0].nameJur)
		data[hasil].prodiNonTek[0].nameJur = blank

		fmt.Print(data[hasil].prodiTek[1].nameJur, " , ")
		fmt.Println("Prodi yang di hapus :", data[hasil].prodiNonTek[1].nameJur)
		data[hasil].prodiNonTek[1].nameJur = blank

		fmt.Print(data[hasil].prodiTek[2].nameJur, " , ")
		fmt.Println("Prodi yang di hapus :", data[hasil].prodiNonTek[2].nameJur)
		data[hasil].prodiNonTek[2].nameJur = blank
	}

}

func menu() {
	fmt.Println("  =========================================================  ")
	fmt.Println("||    SELAMAT DATANG DI PENDAFTARAN UNIVERSITAS ONLINE     ||")
	fmt.Println("  =========================================================  ")
	fmt.Println("|| Pilih menu :                                            ||")
	fmt.Println("|| 1.Input Data CalonMABA                                  ||")
	fmt.Println("|| 2.Edit Data CalonMABA                                   ||")
	fmt.Println("|| 3.Delete data CalonMABA			   	   ||")
	fmt.Println("|| 4.Tampilkan Data Mahasiswa yang Mendaftar               ||")
	fmt.Println("|| 5.Tampilkan Data yang Terurut                           ||")
	fmt.Println("|| 6.Menampilkan jumCalonMABA yang mendaftar pada prodi    || ")
	fmt.Println("|| 7.Exit Menu 						   ||")
	fmt.Println("|| 8. Cari maBA 											||")
	fmt.Println("  =========================================================  ")
}

func clscr() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func main() {
	var (
		pil, pildat int
		hasil       int
		cari        string
	)

	clscr()

	menuawal := "Y"
	for menuawal == "Y" {
		menu()
		fmt.Print("Pilihan menu	: ")
		fmt.Scan(&pil)
		switch pil {
		case 1:
			fmt.Println("  =========================================================  ")
			fmt.Println("|| 1. Tambah Data Calon Mahasiswa                          ||")
			fmt.Println("  =========================================================  ")
			fmt.Print("Pilih menu : ")
			fmt.Scan(&pildat)
			if pildat == 1 {
				Create(&data)
			} else if pildat == 0 {

			}

		case 2:
			fmt.Println("  =======================================================  ")
			fmt.Println("|| 1. Edit Data Calon Mahasiswa                          ||")
			fmt.Println("|| 2. Edit Data Program Studi                            ||")
			fmt.Println("  =======================================================  ")
			fmt.Print("Pilih menu : ")
			fmt.Scan(&pildat)
			if pildat == 1 {
				editBiodata(&data)
			} else if pildat == 2 {
				Sorting(&data, hasil)
				fmt.Print("Maba mana yang ingin Di Edit :")
				fmt.Scan(&cari)
				all = Searcha(&data, cari)
				editProdi(&data, hasil)
				Sorting(&data, hasil)
				fmt.Println(data[hasil])
			} else if pildat == 0 {

			}
		case 3:
			Sorting(&data, all)
			fmt.Println("nama yang ingin di delete")
			fmt.Scan(&cari)
			hasil = Searcha(&data, cari)
			delete(&data, hasil)
			j := 0
			for j < all {
				fmt.Println(data[j])
				j++
			}
		case 4:
			MenampilkanProdTek(data)

		case 5:
			Sorting(&data, all)

		case 6:
			tampilPRODI()

		case 7:
			fmt.Println("==================")
			fmt.Println("||Terimakasih :)||")
			fmt.Println("==================")
		case 8:
			fmt.Println("1. cari berdasarkan nama maba")
			fmt.Print("pilihan :")
			fmt.Scan(&pildat)

			if pildat == 1 {
				//cari nama
				fmt.Print("nama maba :")
				fmt.Scan(&cari)
				all := Searcha(&data, cari)
				if all != -1 {
					MenampilkanProdTek(data)
				} else {
					fmt.Println("data tidak ditemukan ", cari)
				}

			}

		default:
			fmt.Println("MAAF INPUTAN ANDA SALAH")

		}
		fmt.Print("Kembali ke Menu?(Y/T)")
		fmt.Scan(&menuawal)
	}
}
