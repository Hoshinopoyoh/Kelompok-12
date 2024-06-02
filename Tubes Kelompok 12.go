package main

import "fmt"

type atribut struct {
	nama   string
	no_id  int
	stock  int
	harga  int
	terjual int // Tambahkan atribut ini
}

const NMAX int = 100

type tabInt [NMAX]atribut

func menu() {
	fmt.Println("||------------------------------MENU-------------------------------||")
	fmt.Println("01. Tambah Data")
	fmt.Println("02. Hapus Data")
	fmt.Println("03. Cari Data")
	fmt.Println("04. Tampilkan Data")
	fmt.Println("05. Penjualan")
	fmt.Println("06. Edit Data")
	fmt.Println("07. Urutkan Data")
	fmt.Println("08. Mencari data dengan stok terbanyak")
	fmt.Println("09. Mencari data dengan stok terendah")
	fmt.Println("10. Mencari data dengan harga tertinggi")
	fmt.Println("11. Mencari data dengan harga terendah")
	fmt.Println("12. Top 5 Barang dengan Penjualan Terbanyak")
	fmt.Println("13. Keluar")
	fmt.Println("||----------------------------------------------------------------||")
	fmt.Println(" ")
	fmt.Println("Pilih menu yang ingin dibuka :")
}

func switchcase(a int, A *tabInt, n *int) {

	switch a {
	case 1:
		*n = tambah_data(A, *n)
	case 2:
		hapus_data(A, n)
	case 3:
		search(*A, *n)
	case 4:
		tampilkan_data(*A, *n)
	case 5:
		penjualan(A, *n)
	case 6:
		edit_data(A, *n)
	case 7:
		sorting(A, *n)
	case 8:
		max_stok(*A, *n)
	case 9:
		min_stock(*A, *n)
	case 10:
		max_price(*A, *n)
	case 11:
		min_price(*A, *n)
	case 12:
		top_5_penjualan(*A, *n)
	case 13:
		fmt.Println("Terimakasih :)")
		tampilkan_data_akhir(*A, *n)
	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}

func tambah_data(A *tabInt, n int) int {
	var nama string
	var id, stock, harga int
	var i, tambah int
	fmt.Println("Berapa data yang ingin dimasukkan?")
	fmt.Scan(&tambah)
	fmt.Println("Masukkan id, nama, stock, harga (dipisahkan dengan spasi):")
	for i = 0; i < tambah && n < NMAX; i++ {
		fmt.Scan(&id, &nama, &stock, &harga)
		A[n].no_id = id
		A[n].nama = nama
		A[n].stock = stock
		A[n].harga = harga
		A[n].terjual = 0 // Inisialisasi terjual dengan 0
		n++
	}
	return n
}

func tampilkan_data(A tabInt, n int) {
	fmt.Println("Berikut tampilan data :")
	for i := 0; i < n; i++ {
		fmt.Println(A[i].no_id, A[i].nama, A[i].stock, A[i].harga, A[i].terjual)
	}
}

func search(A tabInt, n int) int {
	var id_cari int
	fmt.Println("Masukkan id barang yang ingin dicari:")
	fmt.Scan(&id_cari)
	for i := 0; i < n; i++ {
		if A[i].no_id == id_cari {
			fmt.Println("Barang ditemukan:", A[i].no_id, A[i].nama, A[i].stock, A[i].harga, A[i].terjual)
			return i
		}
	}
	fmt.Println("Barang tidak ditemukan")
	return -1
}

func hapus_data(A *tabInt, n *int) {
	x := search(*A, *n)
	if x != -1 {
		for i := x; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func penjualan(A *tabInt, n int) {
	var id_jual, qty_jual int

	fmt.Println("Masukkan ID barang yang ingin dijual:")
	fmt.Scan(&id_jual)

	index := -1
	for i := 0; i < n; i++ {
		if A[i].no_id == id_jual {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Barang tidak ditemukan")
		return
	}
	fmt.Println("Masukkan jumlah barang yang ingin dijual:")
	fmt.Scan(&qty_jual)

	if A[index].stock < qty_jual {
		fmt.Println("Stok tidak mencukupi")
		return
	}
	// mengupdate stock nya
	A[index].stock -= qty_jual

	// mengupdate jumlah terjual
	A[index].terjual += qty_jual

	// mengtotalkan harga nya
	total_price := A[index].harga * qty_jual

	fmt.Println("Barang berhasil dijual")
	fmt.Printf("Total harga: %d\n", total_price)
}

func edit_data(A *tabInt, n int) {
	var id_edit int
	var found bool
	id_edit = search_for_edit(*A, n)

	// Mencari item yang diberikan
	for i := 0; i < n; i++ {
		if A[i].no_id == A[id_edit].no_id {
			found = true
			var new_name string
			var new_stock, new_harga int
			fmt.Println("Masukkan nama baru:")
			fmt.Scan(&new_name)
			fmt.Println("Masukkan stock baru:")
			fmt.Scan(&new_stock)
			fmt.Println("Masukkan harga baru:")
			fmt.Scan(&new_harga)

			// Mengupdate attribute item
			A[i].nama = new_name
			A[i].stock = new_stock
			A[i].harga = new_harga
			fmt.Println("Data berhasil diupdate")
			break
		}
	}
	if !found {
		fmt.Println("Barang tidak ditemukan")
	}
}

func search_for_edit(A tabInt, n int) int {
	var id_cari int
	fmt.Println("Masukkan id barang yang ingin di edit:")
	fmt.Scan(&id_cari)
	for i := 0; i < n; i++ {
		if A[i].no_id == id_cari {
			fmt.Println("Barang ditemukan:", A[i].no_id, A[i].nama, A[i].stock, A[i].harga, A[i].terjual)
			return i
		}
	}
	fmt.Println("Barang tidak ditemukan")
	return -1
}

func sorting(A *tabInt, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if A[j].no_id > A[j+1].no_id {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}
	fmt.Println("Data berhasil diurutkan")
}

func max_stok(A tabInt, n int) {
	var max int
	var nama_produk string
	max = A[0].stock
	for i := 0; i < n; i++ {
		if A[i].stock > max {
			max = A[i].stock
			nama_produk = A[i].nama
		}
	}
	fmt.Println("Stock terbanyak : ", nama_produk, max)
}

func min_stock(A tabInt, n int) {
	var min int
	var nama_produk string
	min = A[0].stock
	for i := 0; i < n; i++ {
		if A[i].stock < min {
			min = A[i].stock
			nama_produk = A[i].nama
		}
	}
	fmt.Println("Stock terendah :", nama_produk, min)
}

func max_price(A tabInt, n int) {
	var max int
	var nama_produk string
	max = A[0].harga
	for i := 0; i < n; i++ {
		if A[i].harga > max {
			max = A[i].harga
			nama_produk = A[i].nama
		}
	}
	fmt.Println("Harga tertinggi : ", nama_produk, max)
}

func min_price(A tabInt, n int) {
	var min int
	var nama_produk string
	min = A[0].harga
	for i := 0; i < n; i++ {
		if A[i].harga < min {
			min = A[i].harga
			nama_produk = A[i].nama
		}
	}
	fmt.Println("Harga terendah : ", nama_produk, min)
}

func top_5_penjualan(A tabInt, n int) {
	if n < 5 {
		fmt.Println("Data tidak mencukupi untuk menampilkan top 5 penjualan")
		return
	}

	// Lakukan Bubble Sort untuk mengurutkan berdasarkan jumlah terjual
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if A[j].terjual < A[j+1].terjual {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}

	fmt.Println("Top 5 Barang dengan Penjualan Terbanyak:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d. %s - Terjual: %d\n", i+1, A[i].nama, A[i].terjual)
	}
}

func tampilkan_data_akhir(A tabInt, n int) {
	fmt.Println("Berikut tampilan data terakhir:")
	for i := 0; i < n; i++ {
		fmt.Println(A[i].no_id, A[i].nama, A[i].stock, A[i].harga, A[i].terjual)
	}
}

func main() {
	var pilihan int
	var A tabInt
	var n int
	n = 0
	for pilihan != 13 {
		menu()
		fmt.Scan(&pilihan)
		switchcase(pilihan, &A, &n)
	}
}
