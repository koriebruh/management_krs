package dto

type SelectedKrs struct {
	ID             int    `json:"id_schedule"`      // ID Jadwal Tawar
	KrsRecordId    int    `json:"krs_record_id"`    // ID Jadwal Tawar
	TahunAjaran    string `json:"tahun_ajaran"`     // Tahun Ajaran
	KodeMataKuliah string `json:"kode_mata_kuliah"` // Kode Mata Kuliah
	Kelompok       string `json:"kelompok"`         // Kelompok
	NamaMataKuliah string `json:"nama_mata_kuliah"` // Nama Mata Kuliah
	JumlahSKS      int    `json:"jumlah_sks"`       // Jumlah SKS
	Hari           string `json:"hari"`             // Nama Hari
	JamMulai       string `json:"jam_mulai"`        // Jam Mulai Kuliah
	JamSelesai     string `json:"jam_selesai"`      // Jam Selesai Kuliah
	Ruang          string `json:"ruang"`            // Nama Ruang
	JnsJam         string `json:"jns_jam"`          //
}
