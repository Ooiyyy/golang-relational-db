package model

type StudentDetail struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Kelas     string `json:"kelas"`
	WaliKelas string `json:"wali_kelas"`
}

type StudentsGrade struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Nilai int    `json:"nilai"`
	Mapel string `json:"mapel"`
}

type StudentsAvg struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Avg  float64 `json:"rata_rata"`
}
