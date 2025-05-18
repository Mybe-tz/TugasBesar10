package main
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)
const NMAX int = 1000000
type infoKredit struct {
	No			int
	idKredit       string
	JumlahK			float64
	TenorK			float64
	StatusK  	string
}
type arrKredit [NMAX] infoKredit
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}
func InsertionSortJumlahPNaik(A *arrKredit, N int){
	var i, pass int
	var temp infoKredit
	pass = 1
	for pass <= N-1 {
	i = pass
	temp.No = A[pass].No
	temp.idKredit =A[pass].idKredit
	temp.JumlahK = A[pass].JumlahK
	temp.TenorK = A[pass].TenorK
	temp.StatusK = A[pass].StatusK
	for i > 0 && temp.JumlahK < A[i-1].JumlahK {
		A[i].No=A[i-1].No
		A[i].idKredit=A[i-1].idKredit
		A[i].JumlahK=A[i-1].JumlahK
		A[i].TenorK=A[i-1].TenorK
		A[i].StatusK=A[i-1].StatusK
		i = i - 1
	}
	A[i].No=temp.No
	A[i].idKredit=temp.idKredit
	A[i].JumlahK=temp.JumlahK
	A[i].TenorK=temp.TenorK
	A[i].StatusK=temp.StatusK
	pass = pass + 1
	}
}
func InsertionSortJumlahPNurun(A *arrKredit, N int){
	var i, pass int
	var temp infoKredit
	pass = 1
	for pass <= N-1 {
	i = pass
	temp.No = A[pass].No
	temp.idKredit =A[pass].idKredit
	temp.JumlahK = A[pass].JumlahK
	temp.TenorK = A[pass].TenorK
	temp.StatusK = A[pass].StatusK
	for i > 0 && temp.JumlahK > A[i-1].JumlahK {
		A[i].No=A[i-1].No
		A[i].idKredit=A[i-1].idKredit
		A[i].JumlahK=A[i-1].JumlahK
		A[i].TenorK=A[i-1].TenorK
		A[i].StatusK=A[i-1].StatusK
		i = i - 1
	}
	A[i].No=temp.No
	A[i].idKredit=temp.idKredit
	A[i].JumlahK=temp.JumlahK
	A[i].TenorK=temp.TenorK
	A[i].StatusK=temp.StatusK
	pass = pass + 1
	}
}
func SelectionSortNoNaik(A *arrKredit, N int){
	var i, idx, pass int
	var temp infoKredit
	pass = 1
	for pass < N {
	idx = pass - 1
	i = pass
	for i < N {
		if A[i].No < A[idx].No {
		idx = i
		}
	i = i + 1
	}
	temp.No = A[pass-1].No
	temp.idKredit = A[pass-1].idKredit
	temp.JumlahK = A[pass-1].JumlahK
	temp.TenorK = A[pass-1].TenorK
	temp.StatusK = A[pass-1].StatusK
	A[pass-1].No = A[idx].No
	A[pass-1].idKredit = A[idx].idKredit
	A[pass-1].JumlahK = A[idx].JumlahK
	A[pass-1].TenorK = A[idx].TenorK
	A[pass-1].StatusK = A[idx].StatusK
	A[idx].No = temp.No
	A[idx].idKredit = temp.idKredit
	A[idx].JumlahK = temp.JumlahK
	A[idx].TenorK = temp.TenorK
	A[idx].StatusK = temp.StatusK
	pass = pass + 1
	}
}
func SelectionSortNoNurun(A *arrKredit, N int){
	var i, idx, pass int
	var temp infoKredit
	pass = 1
	for pass < N {
	idx = pass - 1
	i = pass
	for i < N {
		if A[i].No > A[idx].No {
		idx = i
		}
	i = i + 1
	}
	temp.No = A[pass-1].No
	temp.idKredit = A[pass-1].idKredit
	temp.JumlahK = A[pass-1].JumlahK
	temp.TenorK = A[pass-1].TenorK
	temp.StatusK = A[pass-1].StatusK
	A[pass-1].No = A[idx].No
	A[pass-1].idKredit = A[idx].idKredit
	A[pass-1].JumlahK = A[idx].JumlahK
	A[pass-1].TenorK = A[idx].TenorK
	A[pass-1].StatusK = A[idx].StatusK
	A[idx].No = temp.No
	A[idx].idKredit = temp.idKredit
	A[idx].JumlahK = temp.JumlahK
	A[idx].TenorK = temp.TenorK
	A[idx].StatusK = temp.StatusK
	pass = pass + 1
	}
}
func SelectionSortTenNaik(A *arrKredit, N int){
	var i, idx, pass int
	var temp infoKredit
	pass = 1
	for pass < N {
	idx = pass - 1
	i = pass
	for i < N {
		if A[i].TenorK < A[idx].TenorK {
		idx = i
		}
	i = i + 1
	}
	temp.No = A[pass-1].No
	temp.idKredit = A[pass-1].idKredit
	temp.JumlahK = A[pass-1].JumlahK
	temp.TenorK = A[pass-1].TenorK
	temp.StatusK = A[pass-1].StatusK
	A[pass-1].No = A[idx].No
	A[pass-1].idKredit = A[idx].idKredit
	A[pass-1].JumlahK = A[idx].JumlahK
	A[pass-1].TenorK = A[idx].TenorK
	A[pass-1].StatusK = A[idx].StatusK
	A[idx].No = temp.No
	A[idx].idKredit = temp.idKredit
	A[idx].JumlahK = temp.JumlahK
	A[idx].TenorK = temp.TenorK
	A[idx].StatusK = temp.StatusK
	pass = pass + 1
	}
}
func SelectionSortTenNurun(A *arrKredit, N int){
	var i, idx, pass int
	var temp infoKredit
	pass = 1
	for pass < N {
	idx = pass - 1
	i = pass
	for i < N {
		if A[i].TenorK > A[idx].TenorK {
		idx = i
		}
	i = i + 1
	}
	temp.No = A[pass-1].No
	temp.idKredit = A[pass-1].idKredit
	temp.JumlahK = A[pass-1].JumlahK
	temp.TenorK = A[pass-1].TenorK
	temp.StatusK = A[pass-1].StatusK
	A[pass-1].No = A[idx].No
	A[pass-1].idKredit = A[idx].idKredit
	A[pass-1].JumlahK = A[idx].JumlahK
	A[pass-1].TenorK = A[idx].TenorK
	A[pass-1].StatusK = A[idx].StatusK
	A[idx].No = temp.No
	A[idx].idKredit = temp.idKredit
	A[idx].JumlahK = temp.JumlahK
	A[idx].TenorK = temp.TenorK
	A[idx].StatusK = temp.StatusK
	pass = pass + 1
	}
}
func BinarySearch(A arrKredit, n int,x int)int{
	var left, right, mid int
	var idx int
	right = n - 1
	idx = -1
	for left <= right && idx == -1{
	mid = (right + left) / 2
	if x < A[mid].No {
	right = mid - 1
	}else if x > A[mid].No {
	left = mid + 1
	}else if x == A[mid].No {
	idx = mid
	}
	}
	return idx 
}
func Sequential(data arrKredit, nData int,idx *int, searchKey string) {
	var i int
	*idx = -1
	for i = 0; i < nData; i++ {
		if data[i].idKredit == searchKey {
			*idx = i
		}
	}
}
func halamanSatu(){
	var masukan string
	ClearScreen()
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Println("==                Aplikasi Pinjaman dan Kredit Sederhana           ==")
	fmt.Println("==                           Created by                            ==")
	fmt.Println("==                   Muhammad Ihsan Abdurrasyid                    ==")
	fmt.Println("==             Sianturi Matthew Juan Rizky Hasiholan               ==")
	fmt.Println("==                   Algoritma Pemrograman 2025                    ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Press x to continue.....")
	for masukan != "x"{
		fmt.Scan(&masukan)
	}
}
func halamanDua(data *arrKredit){
	var nData int
	var pilih,i int
	i = 0
	ClearScreen()
	for pilih != 8{
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Println("==                              MENU                               ==")
	fmt.Println("==                      1. Menambahkan Data                        ==")
	fmt.Println("==                      2. Bunga dan Cicilan                       ==")
	fmt.Println("==                      3. Mengubah Data                           ==")
	fmt.Println("==                      4. Mencari Data                            ==")
	fmt.Println("==                      5. Sortir Data                             ==")
	fmt.Println("==                      6. Menghapus Data                          ==")
	fmt.Println("==                      7. Data yang Tersedia                      ==")
	fmt.Println("==                      8. Keluar                                  ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	switch pilih{
	case 1:
			ClearScreen()
			MenambahkanData(data,&i, &nData)
			ClearScreen()
	case 2:
			ClearScreen()
			BungaKredit(data, nData)
			ClearScreen()
	case 3:
			ClearScreen()
			MengubahData(data, &nData)
			ClearScreen()
	case 4:
			ClearScreen()
			MencariData(data, &nData)
			ClearScreen()
	case 5:
			ClearScreen()
			SortirData(data, &nData)
			ClearScreen()
	case 6:
			ClearScreen()
			MenghapusData(data, &nData, &i)
			ClearScreen()
	case 7:
			ClearScreen()
			DataTersedia(data, &nData)
			ClearScreen()
	}
	}
}
func MenambahkanData(A *arrKredit,i *int, nData *int){
	var n,j,idx,pilih int
	var x string
	idx = -1
	if *i > 0 {
	Data(A, nData)
	}
	fmt.Println("== ---------------------- Penambahan Data ------------------------ ==")
	fmt.Print("== data yang akan ditambahkan       : ")
	fmt.Scan(&n)
	*nData = n + *nData
	for j < n{
		A[*i].No = *i+1
		j++
		fmt.Print("== id peminjaman                    : ")
		fmt.Scan(&x)
		Sequential(*A, *nData,&idx,x)
			for idx != -1{
				idx = -1
				fmt.Println("== id tidak boleh sama ")
				fmt.Println("== jika id berupa nama")
				fmt.Println("== tambahkan angka dibelakangnya ")
				fmt.Println("== contoh (mamat1)")
				fmt.Print("== id peminjaman(ubah)              : ")
				fmt.Scan(&x)
				Sequential(*A, *nData,&idx,x)
				A[*i].idKredit = x
			}
		A[*i].idKredit = x
		fmt.Print("== jumlah yang akan dipinjam(Rp)    : ")
		fmt.Scan(&A[*i].JumlahK)
		fmt.Print("== batas waktu peminjaman(bulan)    : ")
		fmt.Scan(&A[*i].TenorK)
		fmt.Println("== status Pembayaran                : ")
		pilih = 0
		for pilih != 1 && pilih != 2{
		fmt.Println("==                                                                 ==")
		fmt.Println("==                      1. Lunas                                   ==")
		fmt.Println("==                      2. Belum                                   ==")
		fmt.Scan(&pilih)
		switch pilih{
		case 1:
				A[*i].StatusK = "Lunas"
		case 2:
				A[*i].StatusK = "Belum"
		}
		}
		fmt.Println("== --------------------------------------------------------------- ==")
		*i++
	}
}
func BungaKredit(A *arrKredit,nData int){
	var i,pilih int
	var hasilBunga,bunga,jumlahKredit float64
	if nData > 0{
	fmt.Println("== ------------------------- Menu Bunga -------------------------- ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Print("== Bunga yang akan ditambahkan(%)              : ")
	fmt.Scan(&bunga)
	bunga = bunga / 100
	for i < nData {
	jumlahKredit = A[i].JumlahK
	hasilBunga = menghitungBunga(jumlahKredit,bunga)
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Printf("== untuk id %-4s dengan jumlah Rp.%.0f dan tenor %-1.0f bulan           \n", A[i].idKredit, A[i].JumlahK, A[i].TenorK)
	fmt.Println("== Total pinjaman yang harus dibayar           :", hasilBunga)
	fmt.Printf("== Total pinjaman yang harus dicicil per bulan : %.2f\n", hasilBunga / A[i].TenorK)
	fmt.Println("== --------------------------------------------------------------- ==")
	i++
	}
	fmt.Println("==                        1. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	for pilih != 1{
	fmt.Scan(&pilih)
	}
	}else{
	for pilih != 1{
	fmt.Println("== -------------------------- Data ------------------------------- ==")
	fmt.Println("==                      Data not found :(                          ==")
	fmt.Println("==                        1. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}
	}
}
func menghitungBunga(jumlahkredit float64,bunga float64)float64{
	return (bunga * jumlahkredit)+jumlahkredit
}
func UbahJumlahPinjaman(A *arrKredit, nData *int){
	var idx int
	var pengubah float64
	var x string
	Data(A, nData)
	fmt.Println("== --------------------- Mengubah data --------------------------- ==")
	fmt.Print("== id data yang akan diubah                  : ")
	fmt.Scan(&x)
	Sequential(*A, *nData, &idx, x)
	if idx != -1{
	fmt.Print("== data ditemukan dan akan diubah menjadi    : ")
	fmt.Scan(&pengubah)
	A[idx].JumlahK = pengubah
	}else{
	fmt.Println("==                      Data not found                             ==")
	}
	fmt.Println("== --------------------------------------------------------------- ==")
}
func UbahTenor(A *arrKredit, nData *int){
	var idx int
	var pengubah float64
	var x string
	Data(A, nData)
	fmt.Println("== --------------------- Mengubah data --------------------------- ==")
	fmt.Print("== id data yang akan diubah                  : ")
	fmt.Scan(&x)
	Sequential(*A, *nData, &idx, x)
	if idx != -1{
	fmt.Print("== data ditemukan dan akan diubah menjadi    : ")
	fmt.Scan(&pengubah)
	A[idx].TenorK = pengubah
	}else{
	fmt.Println("==                      Data not found                             ==")
	}
	fmt.Println("== --------------------------------------------------------------- ==")
}
func UbahStatus(A *arrKredit, nData *int) {
	var idx int
	var pilih int
	var x string
	Data(A, nData)
	fmt.Println("== --------------------- Mengubah data --------------------------- ==")
	fmt.Print("== id data yang akan diubah                  : ")
	fmt.Scan(&x)
	Sequential(*A, *nData, &idx, x)
	if idx != -1{
	pilih = 0
		for pilih != 1 && pilih != 2{
		fmt.Println("==                                                                 ==")
		fmt.Println("==                      1. Lunas                                   ==")
		fmt.Println("==                      2. Belum                                   ==")
		fmt.Scan(&pilih)
		switch pilih{
		case 1:
				A[idx].StatusK = "Lunas"
		case 2:
				A[idx].StatusK = "Belum"
		}
		}
	}else{
	fmt.Println("==                      Data not found                             ==")
	}
	fmt.Println("== --------------------------------------------------------------- ==")
}
func MengubahData(A *arrKredit, nData *int){
	var pilih int
	for pilih != 4{
	Data(A, nData)
	fmt.Println("== --------------------- Mengubah data --------------------------- ==")
	fmt.Println("==                    1. Jumlah Pinjaman                           ==")
	fmt.Println("==                    2. Tenor                                     ==")
	fmt.Println("==                    3. Status                                    ==")
	fmt.Println("==                    4. Kembali                                   ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	switch pilih{
	case 1:
			ClearScreen()
			UbahJumlahPinjaman(A, nData)
			ClearScreen()
	case 2:
			ClearScreen()
			UbahTenor(A, nData)
			ClearScreen()
	case 3:
			ClearScreen()
			UbahStatus(A, nData)
			ClearScreen()
	}
	}
}
func CariJumlahP(A *arrKredit, nData *int){
	var idx,pilih,i,j int
	var masukan float64
	fmt.Println("== ---------------------- Mencari data --------------------------- ==")
	for pilih != 2 {
	idx = -1
	fmt.Print("== jumlah pinjaman yang akan dicari      : ")
	fmt.Scan(&masukan)
	fmt.Println("== hasil pencarian                       :")
	for i = 0; i < *nData; i++ {
		if A[i].JumlahK == masukan {
			if j == 0 &&  A[i].JumlahK == masukan {
				fmt.Println("== No   ID Peminjam    Jumlah Pinjaman    Tenor (bulan)    Status  ==")
			}
			fmt.Printf("== %.0d     %-14s %-18.0f %-16.0f %-6s \n",A[i].No, A[i].idKredit, A[i].JumlahK, A[i].TenorK, A[i].StatusK)
			idx = j
			j++
		}
	}
	if idx != -1 {
	fmt.Println("==                        1. Cari lagi                             ==")
	fmt.Println("==                        2. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}else{
	fmt.Println("==                  hasil tidak ditemukan :(                       ==")
	fmt.Println("==                        1. Cari lagi                             ==")
	fmt.Println("==                        2. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}
	}
}
func CariId(A *arrKredit, nData *int){
	var idx,pilih int
	var masukan string
	fmt.Println("== ---------------------- Mencari data --------------------------- ==")
	for pilih != 2 {
	idx = -1
	fmt.Print("== id data yang akan dicari              : ")
	fmt.Scan(&masukan)
	Sequential(*A,*nData,&idx,masukan)
	fmt.Println("== hasil pencarian                       :")
	if idx != -1 {
	fmt.Println("== No   ID Peminjam    Jumlah Pinjaman    Tenor (bulan)    Status  ==")
    fmt.Printf("== %.0d     %-14s %-18.0f %-16.0f %-6s \n",A[idx].No, A[idx].idKredit, A[idx].JumlahK, A[idx].TenorK, A[idx].StatusK)
	fmt.Println("==                        1. Cari lagi                             ==")
	fmt.Println("==                        2. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}else{
	fmt.Println("==                  hasil tidak ditemukan :(                       ==")
	fmt.Println("==                        1. Cari lagi                             ==")
	fmt.Println("==                        2. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}
	}
}
func CariTenor(A *arrKredit, nData *int){
	var idx,pilih,i,j int
	var masukan float64
	fmt.Println("== ---------------------- Mencari data --------------------------- ==")
	for pilih != 2 {
	idx = -1
	fmt.Print("== Tenor yang akan dicari(bulan)         : ")
	fmt.Scan(&masukan)
	fmt.Println("== hasil pencarian                       :")
	for i = 0; i < *nData; i++ {
		if A[i].JumlahK == masukan {
			if j == 0 &&  A[i].TenorK == masukan {
				fmt.Println("== No   ID Peminjam    Jumlah Pinjaman    Tenor (bulan)    Status  ==")
			}
			fmt.Printf("== %.0d     %-14s %-18.0f %-16.0f %-6s \n",A[i].No, A[i].idKredit, A[i].JumlahK, A[i].TenorK, A[i].StatusK)
			idx = j
			j++
		}
	}
	if idx != -1 {
	fmt.Println("==                        1. Cari lagi                             ==")
	fmt.Println("==                        2. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}else{
	fmt.Println("==                  hasil tidak ditemukan :(                       ==")
	fmt.Println("==                        1. Cari lagi                             ==")
	fmt.Println("==                        2. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}
	}
}
func CariStatus(A *arrKredit, nData *int){
	var idx,pilih,i,j int
	var masukan string
	fmt.Println("== ---------------------- Mencari data --------------------------- ==")
	for pilih != 2 {
	idx = -1
	fmt.Print("== Status yang akan dicari                 : ")
	fmt.Scan(&masukan)
	fmt.Println("== hasil pencarian                       :")
	for i = 0; i < *nData; i++ {
		if A[i].StatusK == masukan {
			if j == 0 &&  A[i].StatusK == masukan {
				fmt.Println("== No   ID Peminjam    Jumlah Pinjaman    Tenor (bulan)    Status  ==")
			}
			fmt.Printf("== %.0d     %-14s %-18.0f %-16.0f %-6s \n",A[i].No, A[i].idKredit, A[i].JumlahK, A[i].TenorK, A[i].StatusK)
			idx = j
			j++
		}
	}
	if idx != -1 {
	fmt.Println("==                        1. Cari lagi                             ==")
	fmt.Println("==                        2. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}else{
	fmt.Println("==                  hasil tidak ditemukan :(                       ==")
	fmt.Println("==                        1. Cari lagi                             ==")
	fmt.Println("==                        2. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}
	}
}
func MencariData(A *arrKredit, nData *int){
	var pilih int
	for pilih != 5{
	Data(A, nData)
	fmt.Println("== ---------------------- Mencari data --------------------------- ==")
	fmt.Println("==                    1. Id data                                   ==")
	fmt.Println("==                    2. Jumlah Pinjaman                           ==")
	fmt.Println("==                    3. Tenor                                     ==")
	fmt.Println("==                    4. Status                                    ==")
	fmt.Println("==                    5. Kembali                                   ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	switch pilih{
	case 1:
			ClearScreen()
			CariId(A, nData)
			ClearScreen()
	case 2:
			ClearScreen()
			CariJumlahP(A, nData)
			ClearScreen()
	case 3:
			ClearScreen()
			CariTenor(A, nData)
			ClearScreen()
	case 4:
			ClearScreen()
			CariStatus(A, nData)
			ClearScreen()
	}
	}
}
func HapusData(A *arrKredit, nData *int,k *int){
	var idx,i,j int
	var x int
	j = 1
	Data(A, nData)
	fmt.Println("== -------------------- Menghapus data --------------------------- ==")
	fmt.Print("== No data yang akan dihapus                 : ")
	fmt.Scan(&x)
	idx = BinarySearch(*A, *nData, x)
	if idx != -1{
	for i<*nData{
	A[idx+i] = A[idx+j]
	i++
	j++
	}
	*k--
	*nData--
	}else{
	fmt.Println("==                      Data not found                             ==")
	}
	fmt.Println("== --------------------------------------------------------------- ==")
}
func MenghapusData(A *arrKredit, nData *int, k *int){
	var pilih int
	for pilih != 2{
	Data(A, nData)
	fmt.Println("== -------------------- Menghapus data --------------------------- ==")
	fmt.Println("==                    1. No data                                   ==")
	fmt.Println("==                    2. Kembali                                   ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	switch pilih{
	case 1:
			ClearScreen()
			HapusData(A, nData, k)
			ClearScreen()
	}
	}
}
func SortirJumlahP(A *arrKredit, nData *int){
	var pilih int
	for pilih != 4{
	fmt.Println("== --------------------- Mengurutkan data ------------------------ ==")
	fmt.Println("==                       1. Menurun                                ==")
	fmt.Println("==                       2. Menaik                                 ==")
	fmt.Println("==                       3. default                                ==")
	fmt.Println("==                       4. Kembali                                ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	switch pilih{
	case 1:
			ClearScreen()
			InsertionSortJumlahPNurun(&*A,*nData)
			Data(A, nData)
	case 2:
			ClearScreen()
			InsertionSortJumlahPNaik(&*A,*nData)
			Data(A, nData)
	case 3:
			ClearScreen()
			SelectionSortNoNaik(&*A,*nData)
			Data(A, nData)
	}
	}
}
func SortirNo(A *arrKredit, nData *int){
	var pilih int
	Data(A, nData)
	for pilih != 4{
	fmt.Println("== --------------------- Mengurutkan data ------------------------ ==")
	fmt.Println("==                       1. Menurun                                ==")
	fmt.Println("==                       2. Menaik                                 ==")
	fmt.Println("==                       3. default                                ==")
	fmt.Println("==                       4. Kembali                                ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	switch pilih{
	case 1:
			ClearScreen()
			SelectionSortNoNurun(&*A,*nData)
			Data(A, nData)
	case 2:
			ClearScreen()
			SelectionSortNoNaik(&*A,*nData)
			Data(A, nData)
	case 3:
			ClearScreen()
			SelectionSortNoNaik(&*A,*nData)
			Data(A, nData)
	}
	}
}
func SortirTenor(A *arrKredit, nData *int){
	var pilih int
	Data(A, nData)
	for pilih != 4{
	fmt.Println("== --------------------- Mengurutkan data ------------------------ ==")
	fmt.Println("==                       1. Menurun                                ==")
	fmt.Println("==                       2. Menaik                                 ==")
	fmt.Println("==                       3. default                                ==")
	fmt.Println("==                       4. Kembali                                ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	switch pilih{
	case 1:
			ClearScreen()
			SelectionSortTenNurun(&*A,*nData)
			Data(A, nData)
	case 2:
			ClearScreen()
			SelectionSortTenNaik(&*A,*nData)
			Data(A, nData)
	case 3:
			ClearScreen()
			SelectionSortNoNaik(&*A,*nData)
			Data(A, nData)
	}
	}
}
func SortirStatusLun(A *arrKredit, N *int){
	var i, idx, pass int
	var temp infoKredit
	pass = 1
	for pass < *N {
	idx = pass - 1
	i = pass
	for i < *N {
		if A[i].StatusK == "Lunas" {
		idx = i
		}
	i = i + 1
	}
	temp.No = A[pass-1].No
	temp.idKredit = A[pass-1].idKredit
	temp.JumlahK = A[pass-1].JumlahK
	temp.TenorK = A[pass-1].TenorK
	temp.StatusK = A[pass-1].StatusK
	A[pass-1].No = A[idx].No
	A[pass-1].idKredit = A[idx].idKredit
	A[pass-1].JumlahK = A[idx].JumlahK
	A[pass-1].TenorK = A[idx].TenorK
	A[pass-1].StatusK = A[idx].StatusK
	A[idx].No = temp.No
	A[idx].idKredit = temp.idKredit
	A[idx].JumlahK = temp.JumlahK
	A[idx].TenorK = temp.TenorK
	A[idx].StatusK = temp.StatusK
	pass = pass + 1
	}
}
func SortirStatusBel(A *arrKredit, N *int){
	var i, idx, pass int
	var temp infoKredit
	pass = 1
	for pass < *N {
	idx = pass - 1
	i = pass
	for i < *N {
		if A[i].StatusK == "Belum" {
		idx = i
		}
	i = i + 1
	}
	temp.No = A[pass-1].No
	temp.idKredit = A[pass-1].idKredit
	temp.JumlahK = A[pass-1].JumlahK
	temp.TenorK = A[pass-1].TenorK
	temp.StatusK = A[pass-1].StatusK
	A[pass-1].No = A[idx].No
	A[pass-1].idKredit = A[idx].idKredit
	A[pass-1].JumlahK = A[idx].JumlahK
	A[pass-1].TenorK = A[idx].TenorK
	A[pass-1].StatusK = A[idx].StatusK
	A[idx].No = temp.No
	A[idx].idKredit = temp.idKredit
	A[idx].JumlahK = temp.JumlahK
	A[idx].TenorK = temp.TenorK
	A[idx].StatusK = temp.StatusK
	pass = pass + 1
	}
}
func SortirStatus(A *arrKredit, nData *int){
	var pilih int
	Data(A, nData)
	for pilih != 4{
	fmt.Println("== ------------------- Mengurutkan data -------------------------- ==")
	fmt.Println("==                     1. Lunas ke Belum                           ==")
	fmt.Println("==                     2. Belum ke Lunas                           ==")
	fmt.Println("==                     3. default                                  ==")
	fmt.Println("==                     4. Kembali                                  ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	switch pilih{
	case 1:
			ClearScreen()
			SortirStatusLun(A, nData)
			Data(A, nData)
	case 2:
			ClearScreen()
			SortirStatusBel(A, nData)
			Data(A, nData)
	case 3:
			ClearScreen()
			SelectionSortNoNaik(&*A,*nData)
			Data(A, nData)
	}
	}
}
func SortirData(A *arrKredit, nData *int){
	var pilih int
	for pilih != 5{
	Data(A, nData)
	fmt.Println("== ------------------- Mengurutkan data -------------------------- ==")
	fmt.Println("==                       Berdasarkan                               ==")
	fmt.Println("==                    1. No data                                   ==")
	fmt.Println("==                    2. Jumlah Pinjaman                           ==")
	fmt.Println("==                    3. Tenor                                     ==")
	fmt.Println("==                    4. Status                                    ==")
	fmt.Println("==                    5. Kembali                                   ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	switch pilih{
	case 1:
			ClearScreen()
			SortirNo(A, nData)
			ClearScreen()
	case 2:
			ClearScreen()
			SortirJumlahP(A, nData)
			ClearScreen()
	case 3:
			ClearScreen()
			SortirTenor(A, nData)
			ClearScreen()
	case 4:
			ClearScreen()
			SortirStatus(A, nData)
			ClearScreen()
	}
	}
}
func DataTersedia(A *arrKredit, n *int){
	var i,pilih int
	for pilih != 1{
	if *n == 0{
	fmt.Println("== -------------------------- Data ------------------------------- ==")
	fmt.Println("==                      Data not found :(                          ==")
	fmt.Println("==                        1. Kembali                               ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}else{
	fmt.Println("== -------------------------- Data ------------------------------- ==")
	fmt.Println("== No   ID Peminjam    Jumlah Pinjaman    Tenor (bulan)    Status  ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	for i = 0; i < *n; i++ {
    fmt.Printf("== %.0d     %-14s %-18.0f %-16.0f %-6s \n",A[i].No, A[i].idKredit, A[i].JumlahK, A[i].TenorK, A[i].StatusK)
	}
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Println("==                                                                 ==")
	fmt.Println("==                       1. Kembali                                ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Scan(&pilih)
	}
	}
}
func Data(A *arrKredit, n *int){
	var i int
	if *n == 0{
	fmt.Println("== -------------------------- Data ------------------------------- ==")
	fmt.Println("==                      Data not found                             ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	}else{
	fmt.Println("== -------------------------- Data ------------------------------- ==")
	fmt.Println("== No   ID Peminjam    Jumlah Pinjaman    Tenor (bulan)    Status  ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	for i = 0; i < *n; i++ {
    fmt.Printf("== %.0d     %-14s %-18.0f %-16.0f %-6s \n",A[i].No, A[i].idKredit, A[i].JumlahK, A[i].TenorK, A[i].StatusK)
	}
	fmt.Println("== --------------------------------------------------------------- ==")
	fmt.Println("==                                                                 ==")
	fmt.Println("== --------------------------------------------------------------- ==")
	}
}
func main(){
	var data arrKredit
	halamanSatu()
	halamanDua(&data)
}