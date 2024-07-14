package main

// Nama	 : Denis Lizard Sambawo Dimara
// NPM	 : 21081010159
// House : 8
// ID	 : BE9067925
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

// Mendefinisikan interface untuk manajemen mahasiswa
type StudentManager interface {
	// Method untuk proses login
	Login(id string, name string) error
	// Method untuk proses registrasi mahasiswa baru
	Register(id string, name string, studyProgram string) error
	// Method untuk mendapatkan nama program studi berdasarkan kode program studi
	GetStudyProgram(code string) (string, error)
	// Method untuk memodifikasi data mahasiswa
	ModifyStudent(name string, fn model.StudentModifier) error
}

// Implementasi dari interface StudentManager yang menggunakan penyimpanan data di dalam memori
type InMemoryStudentManager struct {
	students             []model.Student   // Slice untuk menyimpan data mahasiswa
	studentStudyPrograms map[string]string // Map untuk menyimpan program studi berdasarkan kode program studi
}

// Fungsi untuk membuat objek InMemoryStudentManager baru
func NewInMemoryStudentManager() *InMemoryStudentManager {
	// Inisialisasi objek InMemoryStudentManager dengan data mahasiswa awal dan program studi yang tersedia
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
	}
}

// Method untuk mendapatkan data mahasiswa
// sm receiver dari Method
// ini method signature
func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	return sm.students
}

// Method untuk proses login mahasiswa
func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	// Validasi input ID dan Nama
	// UI Login Processing
	fmt.Println("------------------------------------------------")
	fmt.Println("===============Login Processing=================")
	fmt.Println("------------------------------------------------")
	if id == "" || name == "" {
		return "", fmt.Errorf("%s", "ID or Name is undefined!") // Mengembalikan error jika ID atau Nama kosong
	}

	// Loop melalui data mahasiswa
	for _, student := range sm.students {
		// Jika ID dan Nama cocok dengan data mahasiswa
		fmt.Println("Cek Data Student : ", student)
		// {A12345 Aditira TI} , {B21313 Dito TK} , {A34555 Afis MI}
		if student.ID == id && student.Name == name { // cek apakah id dan name terdapat
			return fmt.Sprintf("Login berhasil: %s", name), nil // Mengembalikan pesan berhasil jika autentikasi sukses
		}
	}

	return "", fmt.Errorf("Login gagal: data mahasiswa tidak ditemukan") // Mengembalikan error jika data mahasiswa tidak ditemukan
}

// Method untuk proses registrasi mahasiswa baru
func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	//UI Register Processing
	fmt.Println("------------------------------------------------")
	fmt.Println("===============Register Processing==============")
	fmt.Println("------------------------------------------------")
	// Validasi input ID, Nama, dan Program Studi
	if id == "" || name == "" || studyProgram == "" {
		return "", fmt.Errorf("%s", "ID, Name or StudyProgram is undefined!") // Mengembalikan error jika ID, Nama, atau Program Studi kosong
	}

	// Validasi apakah Program Studi tersedia
	if _, exists := sm.studentStudyPrograms[studyProgram]; !exists {
		// fmt.Println("Cek Exists :", exists)                                  // Exists true / false
		return "", fmt.Errorf("Study program %s is not found", studyProgram) // Mengembalikan error jika program studi tidak ditemukan
	}

	// Validasi apakah ID sudah digunakan
	for _, student := range sm.students {
		if student.ID == id {
			return "", fmt.Errorf("%s", "Registrasi gagal: id sudah digunakan") // Mengembalikan error jika ID sudah digunakan
		}
	}

	// Menambahkan mahasiswa baru ke dalam slice students
	sm.students = append(sm.students, model.Student{
		ID:           id,           // tambahkan id baru
		Name:         name,         // tambahkan nama baru
		StudyProgram: studyProgram, // tambahkan studyProgram
	})

	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram), nil // Mengembalikan pesan berhasil jika registrasi sukses
}

// Method untuk mendapatkan nama program studi berdasarkan kode program studi
func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	fmt.Println("---------------------------------------------------")
	fmt.Println("============= GET STUDY PROGRAM CEK ===============")
	fmt.Println("---------------------------------------------------")
	// Validasi input kode program studi
	if code == "" {
		return "", fmt.Errorf("%s", "Code is undefined!") // Mengembalikan error jika kode program studi kosong
	}

	// Mengambil nama program studi dari map jika kode program studi ditemukan
	if program, exists := sm.studentStudyPrograms[code]; exists {
		return program, nil // Mengembalikan nama program studi jika kode program studi ditemukan
	}

	return "", fmt.Errorf("%s", "Kode program studi tidak ditemukan") // Mengembalikan error jika kode program studi tidak ditemukan
}

// Method untuk memodifikasi data mahasiswa
func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	fmt.Println("------------------------------------------------")
	fmt.Println("============= MODIFY DATA STUDENT ==============")
	fmt.Println("------------------------------------------------")
	// Validasi input nama mahasiswa
	if name == "" {
		return "", fmt.Errorf("%s", "Mahasiswa tidak ditemukan.") // Mengembalikan error jika nama mahasiswa kosong
	}

	// Loop melalui data mahasiswa
	for _, student := range sm.students {
		// Jika nama mahasiswa ditemukan
		if student.Name == name {
			// Melakukan modifikasi data mahasiswa menggunakan fungsi modifier
			if err := fn(&student); err != nil {
				return "", err // Mengembalikan error jika terjadi kesalahan dalam memodifikasi mahasiswa
			}
			return "Program studi mahasiswa berhasil diubah.", nil // Mengembalikan pesan berhasil jika modifikasi sukses
		}
	}

	return "", fmt.Errorf("%s", "Mahasiswa tidak ditemukan.") // Mengembalikan error jika mahasiswa tidak ditemukan
}

// Method untuk membuat fungsi modifier program studi mahasiswa
func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	return func(s *model.Student) error {
		// Validasi apakah kode program studi tersedia
		if _, exists := sm.studentStudyPrograms[programStudi]; !exists {
			return fmt.Errorf("%s", "Kode program studi tidak ditemukan") // Mengembalikan error jika kode program studi tidak ditemukan
		}
		s.StudyProgram = programStudi // Mengubah program studi mahasiswa
		return nil
	}
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "5":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
