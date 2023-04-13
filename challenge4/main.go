package main

import (
	"fmt"
	"os"
	"strconv"
)

type friend struct {
	Name       string
	Address    string
	Job        string
	Motivation string
}

func (f *friend) print() {
	fmt.Println("Nama:", f.Name)
	fmt.Println("Alamat:", f.Address)
	fmt.Println("Pekerjaan:", f.Job)
	fmt.Println("Alasan memilih kelas Golang:", f.Motivation)
}

func getData() []friend {
	return []friend{
		{
			Name:       "John",
			Address:    "Jl. Kebon Jeruk",
			Job:        "Programmer",
			Motivation: "Mau jadi programmer golang",
		},
		{
			Name:       "Richard",
			Address:    "Jl. Pegangsaan",
			Job:        "UI/UX Designer",
			Motivation: "Coba career switch",
		},
		{
			Name:       "Yuhu",
			Address:    "Jl. Jalan",
			Job:        "Business Analyst",
			Motivation: "Penasaran aja sih yang tim tech lakuin selama ini tu apa aja sebenarnya",
		},
		{
			Name:       "Brick",
			Address:    "Jl. Turunan",
			Job:        "Wiraswasta",
			Motivation: "Mau go digital dengan go lang",
		},
		{
			Name:       "Rethorn",
			Address:    "Jl. Tunjungan",
			Job:        "Pengangguran",
			Motivation: "Ngalir ajalah",
		},
		{
			Name:       "Erick",
			Address:    "Jl. Tanjakan",
			Job:        "Content Creator",
			Motivation: "Biar kelihatan keren aja gitu bisa ngoding go lang",
		},
		{
			Name:       "Calvin",
			Address:    "Jl. Kebon Jeruk",
			Job:        "Student",
			Motivation: "Disuruh join",
		},
		{
			Name:       "Media",
			Address:    "Jl. Hunian",
			Job:        "PHP Developer",
			Motivation: "New language, new income",
		},
		{
			Name:       "Zuru",
			Address:    "Jl. Mangga",
			Job:        "PNS",
			Motivation: "Bosen",
		},
		{
			Name:       "Brodi",
			Address:    "Jl. Kenangan",
			Job:        "CEO",
			Motivation: "Belajar ngoding sendiri, jadi gak perlu hire programmer",
		},
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please add number after running command. Example: go run main.go {number}")
		return
	}

	number, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("[ERROR]:", err.Error())
		return
	}

	number--
	data := getData()
	if number < 0 || number >= int64(len(data)) {
		fmt.Println("[ERROR]: Number out of range. Please try again between 1 to 10")
		return
	}

	data[number].print()
}
