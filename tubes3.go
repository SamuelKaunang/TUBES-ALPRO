package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Flight struct {
	Airline    string
	Price      int
	Departure  string
	Arrival    string
	FlightTime string
}

type Booking struct {
	Name        string
	Flight      Flight
	Destination string
}

var bookings []Booking

func readString(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInt(reader *bufio.Reader) (int, error) {
	strInput := readString(reader)
	return strconv.Atoi(strInput)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Penerbangan Domestik")
		fmt.Println("2. Penerbangan Internasional")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")

		choice, err := readInt(reader)
		if err != nil {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		if choice == 3 {
			fmt.Println("Terima kasih telah menggunakan layanan kami.")
			break
		}

		if choice != 1 && choice != 2 {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		var destination string
		if choice == 1 {
			fmt.Print("Masukkan tujuan penerbangan domestik: ")
		} else if choice == 2 {
			fmt.Print("Masukkan tujuan penerbangan internasional: ")
		}
		destination = readString(reader)

		flights := getFlights(destination)
		if len(flights) == 0 {
			fmt.Println("Tidak ada penerbangan yang tersedia untuk tujuan ini.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Println("\nDaftar penerbangan:")
		printFlights(flights)

		fmt.Println("\nMenu Sorting:")
		fmt.Println("1. Termurah ke Termahal")
		fmt.Println("2. Termahal ke Termurah")
		fmt.Println("3. Berdasarkan Waktu Keberangkatan")
		fmt.Print("Pilih menu sorting: ")

		sortChoice, err := readInt(reader)
		if err != nil {
			sortChoice = 0 
		}

		switch sortChoice {
		case 1:
			sort.Slice(flights, func(i, j int) bool {
				return flights[i].Price < flights[j].Price
			})
		case 2:
			sort.Slice(flights, func(i, j int) bool {
				return flights[i].Price > flights[j].Price
			})
		case 3:
			sort.Slice(flights, func(i, j int) bool {
				return flights[i].Departure < flights[j].Departure
			})
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Println("\nDaftar penerbangan setelah sorting:")
		printFlights(flights)

		fmt.Print("Pilih penerbangan (nomor): ")
		flightChoice, err := readInt(reader)
		if err != nil || flightChoice < 1 || flightChoice > len(flights) {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		flight := flights[flightChoice-1]
		fmt.Printf("\nDetail penerbangan yang dipilih:\nMaskapai: %s\nHarga: Rp%d\nKeberangkatan: %s\nKedatangan: %s\nDurasi: %s\n",
			flight.Airline, flight.Price, flight.Departure, flight.Arrival, flight.FlightTime)

		fmt.Print("\nApakah Anda ingin melanjutkan dengan penerbangan ini? (yes/no): ")
		confirm := readString(reader)
		if strings.ToLower(confirm) != "yes" {
			fmt.Println("Pemesanan dibatalkan.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Print("Masukkan nama Anda: ")
		name := readString(reader)

		booking := Booking{
			Name:        name,
			Flight:      flight,
			Destination: destination,
		}
		bookings = append(bookings, booking)

		fmt.Printf("Anda telah sukses terdaftar pada penerbangan %s dengan tujuan %s.\n", flight.Airline, destination)
		time.Sleep(3 * time.Second)
	}
}


func getFlights(destination string) []Flight {
	return []Flight{
		{"Garuda Indonesia", 1500000, "08:00", "10:00", "2h"},
		{"Lion Air", 900000, "09:00", "11:00", "2h"},
		{"AirAsia", 1200000, "10:00", "12:00", "2h"},
		{"Singapore Airlines", 2500000, "11:00", "13:00", "2h"},
	}
}

func printFlights(flights []Flight) {
	for i, flight := range flights {
		fmt.Printf("%d. %s - Rp%d (Keberangkatan: %s, Kedatangan: %s)\n", i+1, flight.Airline, flight.Price, flight.Departure, flight.Arrival)
	}
}
