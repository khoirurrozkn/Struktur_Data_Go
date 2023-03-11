package main
import ("fmt")

type Table struct{
	bulan, tujuan, status string
}

func show(data[100]Table){
	for idx, val := range data{
		if( val.bulan == ""){
			break;
		}
		fmt.Println(" Nomer ", (idx + 1))
		fmt.Println(" Bulan : ", val.bulan)
		fmt.Println(" Tujuan : ", val.tujuan)
		fmt.Println(" Status : ", val.status, "\n")
	}
}

func main(){
	data := [100]Table{}
	index := 0
	var inputTemp string
	var inputGanti, inputPilihan, inputNilai int

	for {
		fmt.Println("-------------------------------------------------")
		fmt.Println(" Tujuan  Wisata ")
		fmt.Println("-------------------------------------------------")

		show(data)

		fmt.Println("-------------------------------------------------")

		fmt.Println(" 1. Edit Rencana")
		fmt.Println(" 2. Tambah Rencana")
		fmt.Println("-------------------------------------------------")

		fmt.Scan(&inputPilihan)

		fmt.Println("-------------------------------------------------")
		if(inputPilihan == 1){
			fmt.Println(" Masukan nomor wisata ")

			fmt.Scan(&inputGanti)
			inputGanti--
			fmt.Println("-------------------------------------------------")

			fmt.Println(" Apa yang mau di ganti ")
			fmt.Println(" 1. Bulan")
			fmt.Println(" 2. Tujuan")
			fmt.Println(" 3. Status")

			fmt.Scan(&inputNilai)
			fmt.Println("-------------------------------------------------")

			fmt.Println(" Input : ")
			fmt.Scan(&inputTemp)

			if(inputNilai == 1){

				data[inputGanti].bulan = inputTemp
			}else if(inputNilai == 2){
				
				data[inputGanti].tujuan = inputTemp
			}else{

				data[inputGanti].status = inputTemp
			}
		}else if(inputPilihan == 2){
			fmt.Println(" Masukan Bulan : ")
			fmt.Scan(&data[index].bulan)
			fmt.Println(" Masukan Tujuan : ")
			fmt.Scan(&data[index].tujuan)
			fmt.Println(" Masukan Status : ")
			fmt.Scan(&data[index].status)
			index++
		}
	}
}