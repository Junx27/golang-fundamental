package main

import (
	"fmt"
	"strconv"
)

// membuat biodata, menentukan tipe data variabel
func soal1() {
	type biodata struct {
		nama, pekerjaan string
		umur            int
	}
	saya := biodata{
		nama:      "Tri Saptono",
		pekerjaan: "Freelancer Web Development",
		umur:      33,
	}
	info := fmt.Sprintf("Nama: %s\nPekerjaan: %s\nUmur: %d\n", saya.nama, saya.pekerjaan, saya.umur)
	fmt.Println(info)
}

// fungsi operasi hitungan
func hasilOpearsi(a, b int, op string) int {
	var res int
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "/":
		res = a / b
	case "*":
		res = a * b
	case "%":
		res = a % b
	default:
		res = 0
	}
	return res
}

// fungsi keterangan operator
func keterangan(op string) string {
	var keteranganOperator string
	if op == "+" {
		keteranganOperator = "penjumlahan"
	} else if op == "-" {
		keteranganOperator = "pengurangan"
	} else if op == "/" {
		keteranganOperator = "pembagian"
	} else if op == "*" {
		keteranganOperator = "perkalian"
	} else if op == "%" {
		keteranganOperator = "sisa bagi"
	} else {
		keteranganOperator = "operator tidak diketahui"
	}
	return keteranganOperator
}

// membuat hasil dari opeartor (+,-,/,*,%)
func soal2() {
	var operators = []string{"+", "-", "/", "*", "%"}
	a, b := 3, 3

	for _, operator := range operators {
		hasilAkhir := hasilOpearsi(a, b, operator)
		namaOperator := keterangan(operator)
		fmt.Println("hasil", namaOperator, "dari bilangan", a, "dengan", b, "adalah", hasilAkhir)
	}
}

// membuat keterangan nilai dengan menggunakan if-else
func soal3() {
	for {
		var angka string
		fmt.Println("Masukan bilnagan bulat:")
		fmt.Scanln(&angka)
		number, err := strconv.Atoi(angka)
		if err != nil {
			fmt.Println("angka yang dimasukan salah")
			continue
		} else {
			if number > 85 {
				fmt.Println("A")
			} else if number >= 70 && number <= 84 {
				fmt.Println("B")
			} else if number >= 50 && number <= 69 {
				fmt.Println("C")
			} else {
				fmt.Println("D")
			}
		}
		break
	}
}

// membuat perulangan bilangan genap 1 sampai 20 dengan for loop
func soal4() {
	fmt.Println("Bilangan bulat 1 sampai 20")
	for i := 0; i <= 20; i++ {
		if i > 0 && i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// membuat perulangan huruf di dalam variabel dengan for in range
func soal5() {
	nama := "Tri Saptono"
	for _, char := range nama {
		fmt.Println(string(char))
	}
}

func main() {
	//memanggil fungsi soal1
	soal1()
	//memanggil fungsi soal2
	soal2()
	//memanggil fungsi soal3
	soal3()
	//memanggil fungsi soal4
	soal4()
	//memanggil fungsi soal5
	soal5()
}
