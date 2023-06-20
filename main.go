package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Field struct {
	OrangTua string
	Nama     string
	Pasangan string
}

type Tree struct {
	Data Field
	Anak []*Tree
}

var Halilintar Tree

func buatPerintis(nama string) string {
	Halilintar.Data.Nama = nama
	return "Si " + nama + " telah dibuat"
}

func cariKeluarga(nama string) *Tree {
	if Halilintar.Data.Nama == nama {
		return &Halilintar
	}

	return cariAnak(Halilintar.Anak, nama)
}

func cariAnak(anak []*Tree, nama string) *Tree {
	if anak == nil {
		return nil
	}

	for _, val := range anak {
		if val.Data.Nama == nama {
			return val
		}

		found := cariAnak(val.Anak, nama)
		if found != nil {
			return found
		}
	}

	return nil
}

func nikahkan(nama string, pasangan string) string {
	res := cariKeluarga(nama)

	if cariKeluarga(pasangan) != nil {
		return "Nama pasangan tidak boleh " + pasangan + " karena nama sudah di pakai"
	}

	if res != nil {
		if res.Data.Pasangan != "" {
			return "Eitts " + nama + " hayooo, ngga boleh poligami"
		}
		res.Data.Pasangan = pasangan
		return "Si " + nama + " telah dinikahkan dengan " + pasangan
	}

	return "Keluarga tidak ditemukan"
}

func buatAnak(namaOrtu string, namaAnak string) string {
	if cariKeluarga(namaAnak) != nil {
		return "Nama anak tidak boleh " + namaAnak + " karena nama sudah di pakai"
	}

	if namaOrtu == namaAnak {
		return "Nama anak dan ortu tidak boleh sama"
	}

	temp := Tree{
		Data: Field{
			OrangTua: namaOrtu,
			Nama:     namaAnak,
		},
	}

	res := cariKeluarga(namaOrtu)

	if res != nil {
		if res.Data.Pasangan == "" {
			return "Si " + namaOrtu + " belum nikah, gimana mau punya anak"
		}
		res.Anak = append(res.Anak, &temp)
		return "Si " + namaOrtu + " telah mempunyai anak baru yaitu " + namaAnak
	}

	return "Data orang tua tidak ditemukan"
}

func lihatSemuaKeluarga(orang *Tree, spasi1 string, spasi2 string) {
	if orang.Data.Nama == "" {
		fmt.Println("Belum ada keluarga")
		return
	}
	fmt.Print(spasi1 + spasi2 + "|_ ", orang.Data.Nama)

	if orang.Data.Pasangan != "" {
		fmt.Println(" -> " + orang.Data.Pasangan)
	}else{
		fmt.Println()
	}

	temp := spasi1 + spasi2

	for _, val := range orang.Anak {
		lihatSemuaKeluarga(val, temp, "    ")
	}
}

func lihatOrang(nama string) {
	res := cariKeluarga(nama)

	if res == nil {
		fmt.Println("Nama orang tidak di temukan")
		return
	}

	fmt.Print("Si " + res.Data.Nama)
	if res.Data.OrangTua != "" {
		fmt.Println(" mempunyai orang tua ++ " + res.Data.OrangTua)

		ortu := cariKeluarga(res.Data.OrangTua)
		fmt.Print("Si " + res.Data.Nama)
		if len(ortu.Anak) > 1 {
			fmt.Print(" mempuyai saudara")
			for idx, val := range ortu.Anak {
				if idx == 0 {
					continue
				} else {
					fmt.Print(" // " + val.Data.Nama + " ")
				}
			}
		} else {
			fmt.Print(" tidak mempuyai saudara")
		}
		fmt.Println()

	} else {
		fmt.Println(" tidak mempunyai orang tua")
	}

	fmt.Print("Si " + res.Data.Nama)
	if res.Data.Pasangan != "" {
		fmt.Println(" mempunyai Pasangan si -> " + res.Data.Pasangan)

		fmt.Print("Si " + res.Data.Nama)
		if res.Anak != nil {
			fmt.Print(" mempunyai anak")
			for _, val := range res.Anak {
				fmt.Print(" |_ " + val.Data.Nama)
			}
			fmt.Println()
		} else {
			fmt.Println(" tidak mempunyai anak")
		}

	} else {
		fmt.Println(" tidak mempunyai pasangan")
	}

	ada := false

	fmt.Print("Cucu si " + res.Data.Nama)
	for idx := range res.Anak {
		for _, cucu := range res.Anak[idx].Anak {
			fmt.Print(" && ", cucu.Data.Nama)
			ada = true
		}
	}
	if !ada {
		fmt.Print(" tidak mempunyai cucu")
	}
	fmt.Println()
}

func strip() {
	fmt.Println("==================================================")
}

func slash(value string) {
	strip()
	fmt.Println("                 " + value + "                    ")
	strip()
}

func main() {
	var menu int
	var nama string
	var pasangan string
	var anak string
	next := true

	for {
		if next {
			clear()
		}
		strip()
		lihatSemuaKeluarga(&Halilintar, "", "")
		strip()
		fmt.Println("Selamat datang di keluarga Halilintar")
		fmt.Println("1. Buat perintis")
		fmt.Println("2. Tambah pasangan")
		fmt.Println("3. Buat anak")
		fmt.Println("4. Lihat satu orang")
		fmt.Println("5. exit")
		fmt.Println("Pilih salah satu")
		fmt.Scan(&menu)

		if menu == 1 {
			clear()
			slash("Buat Perintis")
			if Halilintar.Data.Nama != "" {
				fmt.Println("Perintis telah ada yaitu si " + Halilintar.Data.Nama)
				continue
			}
			fmt.Println("Masukan nama perintis")
			fmt.Scan(&nama)
			fmt.Println(buatPerintis(nama))
			time.Sleep(1 * time.Second)
		} else if menu == 2 {
			clear()
			slash("Tambah pasangan")
			lihatSemuaKeluarga(&Halilintar, "", "")
			strip()
			fmt.Println("Masukan nama yang mau di nikahkan")
			fmt.Scan(&nama)
			fmt.Println("Masukan nama pasangan")
			fmt.Scan(&pasangan)
			fmt.Println(nikahkan(nama, pasangan))
			time.Sleep(1 * time.Second)
		} else if menu == 3 {
			clear()
			slash("Buat anak")
			lihatSemuaKeluarga(&Halilintar, "", "")
			strip()
			fmt.Println("Masukan nama orang tua nya")
			fmt.Scan(&nama)
			fmt.Println("Masukan nama anaknya")
			fmt.Scan(&anak)
			fmt.Println(buatAnak(nama, anak))
			time.Sleep(1 * time.Second)
		} else if menu == 4 {
			clear()
			slash("Lihat orang")
			fmt.Println("Masukan nama orangnya")
			fmt.Scan(&nama)
			strip()
			lihatOrang(nama)
			next = false
		} else if menu == 5 {
			break
		} else {
			fmt.Println("Format yang anda masukan salah")
		}
	}
}