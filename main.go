package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
)

func main() {
	var nik int

	fmt.Print("Input NIK: ")
	_, err := fmt.Scanf("%d", &nik)

	if err != nil {
		log.Fatal("NIK Tidak Valid")
	}

	nikstr := strconv.Itoa(nik)

	if len(nikstr) != 16 {
		log.Fatal("NIK Tidak Valid")
	}

	tanggal := nikstr[6:8]
	bulan := nikstr[8:10]
	tahun := nikstr[10:12]
	provinsi := nikstr[0:2]
	kabkot := nikstr[0:4]
	kecamatan := nikstr[0:6]
	uniqcode := nikstr[12:16]
	jeniskelamin := "LAKI-LAKI"
	cekjk, _ := strconv.Atoi(nikstr[6:8])

	if cekjk > 40 {
		jeniskelamin = "PEREMPUAN"
	}

	data, err := os.Open("data.json")
	defer data.Close()

	if err != nil {
		log.Fatal("Data Tidak Valid")
	}

	byteValue, _ := ioutil.ReadAll(data)
	provinsi, _ = jsonparser.GetString(byteValue, "provinsi", provinsi)
	kabkot, _ = jsonparser.GetString(byteValue, "kabkot", kabkot)
	kecamatan, _ = jsonparser.GetString(byteValue, "kecamatan", kecamatan)

	fmt.Printf("Tanggal Lahir: %s/%s/%s\n", tanggal, bulan, tahun)
	fmt.Printf("Jenis Kelamin: %s\n", jeniskelamin)
	fmt.Printf("Provinsi: %s\n", provinsi)
	fmt.Printf("Kab/Kota: %s\n", kabkot)
	fmt.Printf("Kecamatan: %s\n", strings.Split(kecamatan, "--")[0])
	fmt.Printf("Kode Pos: %s\n", strings.Split(kecamatan, "--")[1])
	fmt.Printf("Uniqcode: %s\n", uniqcode)
}
