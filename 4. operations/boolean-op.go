package main

import "fmt"

// main is the entry point of the program.
// It evaluates whether a student passes based on their final score and attendance.
func main() {
	// nilaiAkhir represents the final score of the student.
	var nilaiAkhir = 90
	// absensi represents the attendance percentage of the student.
	var absensi = 80

	// lulusNilaiAkhir checks if the final score is greater than 80.
	var lulusNilaiAkhir = nilaiAkhir > 80
	// lulusAbsensi checks if the attendance is greater than 80.
	var lulusAbsensi = absensi > 80

	// lulus determines if the student passes based on both final score and attendance.
	var lulus = lulusNilaiAkhir && lulusAbsensi

	// Print the result of lulus.
	fmt.Println(lulus)
	// Print the result of the logical AND operation between final score > 80 and attendance > 70.
	fmt.Println(nilaiAkhir > 80 && absensi > 70)
}
