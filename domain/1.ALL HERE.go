package domain

//TagihanMhs represents the tagihan_mhs table
//type TagihanMhs struct {
//	ID            int       `gorm:"primaryKey;autoIncrement;comment:id biasa"`
//	TA            int       `gorm:"not null;comment:tahun ajaran"`
//	NimDinus      string    `gorm:"not null;comment:nim mahasiswa"`
//	SppBank       string
//	SppBayar      int       `gorm:"not null;default:0;comment:status bayar spp 1: bayar 0: belum bayar"`
//	SppBayarDate  time.Time `gorm:"comment:tanggal pada saat operator input pembayaran"`
//	SppDispensasi int
//	SppHost       string    `gorm:"comment:ip/host operator"`
//	SppStatus     int       `gorm:"not null;comment:1 : full payment 0 : dispensasi"`
//	SppTransaksi  string    `gorm:"comment:jenis pembayaran : langsung, transfer"`
//	TahunAjaran   TahunAjaran `gorm:"foreignKey:TA;references:Kode"`
//	Mahasiswa     MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
//}
//
//MahasiswaDinus represents the mahasiswa_dinus table
//type MahasiswaDinus struct {
//	ID       int    `gorm:"primaryKey;autoIncrement"`
//	NimDinus string `gorm:"unique;not null;default:''"`
//	TaMasuk  int
//	Prodi    string
//	PassMhs  string    `gorm:"default:null"`
//	Kelas    int       `gorm:"not null;comment:0 = not choose,1 = pagi,2 = malam,3 = karyawan/pegawai"`
//	AkdmStat string    `gorm:"type:char(2);not null;comment:1 = aktif, 2 = cuti, 3 = keluar, 4 = lulus, 5 = mangkir, 6 = meninggal, 7 = DO, 8 = Aktif Keuangan"`
//}
//
//MatkulKurikulum represents the matkul_kurikulum table
//type MatkulKurikulum struct {
//	KurID         int    `gorm:"column:kur_id"`
//	Kdmk          string `gorm:"unique;not null"`
//	Nmmk          string
//	Nmen          string
//	Tp            string `gorm:"type:enum('T','P','TP')"`
//	Sks           int
//	SksT          int16
//	SksP          int16
//	Smt           int
//	JnsSmt        int
//	Aktif         bool
//	KurNama       string
//	KelompokMakul string `gorm:"type:enum('MPK','MKK','MKB','MKD','MBB','MPB')"`
//	KurAktif      bool   `gorm:"type:bit(1)"`
//	JenisMatkul   string `gorm:"type:enum('wajib','pilihan')"`
//}
//
//Hari represents the hari table
//type Hari struct {
//	ID     int8   `gorm:"unique;not null"`
//	Nama   string `gorm:"type:varchar(6);not null"`
//	NamaEn string `gorm:"type:varchar(20);not null"`
//}
//
//JadwalInputKrs represents the jadwal_input_krs table
//type JadwalInputKrs struct {
//	ID          int       `gorm:"primaryKey;autoIncrement"`
//	TA          int       `gorm:"not null;default:0"`
//	Prodi       string    `gorm:"type:char(3)"`
//	TglMulai    time.Time
//	TglSelesai  time.Time
//}
//
//IpSemester represents the ip_semester table
//type IpSemester struct {
//	ID         int       `gorm:"primaryKey;autoIncrement"`
//	TA         int       `gorm:"not null;default:0"`
//	NimDinus   string    `gorm:"not null"`
//	Sks        int       `gorm:"not null"`
//	Ips        string    `gorm:"not null"`
//	LastUpdate time.Time `gorm:"default:null"`
//	TahunAjaran TahunAjaran `gorm:"foreignKey:TA;references:Kode"`
//	Mahasiswa   MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
//}
//
//SesiKuliah represents the sesi_kuliah table
//type SesiKuliah struct {
//	ID          int       `gorm:"primaryKey;autoIncrement"`
//	Jam         string    `gorm:"not null;default:''"`
//	Sks         int16     `gorm:"not null;default:0"`
//	JamMulai    time.Time `gorm:"type:time;default:null"`
//	JamSelesai  time.Time `gorm:"type:time;default:null"`
//	Status      int       `gorm:"default:1;comment:0=tidak valid, 1= jam valid(kelipatan 50menit), 2 = jam yang harusnya tisak di pakai(jam istirahat)"`
//}
//
//SesiKuliahBentrok represents the sesi_kuliah_bentrok table
//type SesiKuliahBentrok struct {
//	ID        int `gorm:"primaryKey"`
//	IDBentrok int `gorm:"primaryKey"`
//	Sesi      SesiKuliah `gorm:"foreignKey:ID;references:ID"`
//	SesiBentrok SesiKuliah `gorm:"foreignKey:IDBentrok;references:ID"`
//}
//
//KrsRecord represents the krs_record table
//type KrsRecord struct {
//	ID        int    `gorm:"primaryKey;autoIncrement"`
//	TA        int    `gorm:"not null;default:0"`
//	Kdmk      string `gorm:"not null"`
//	IDJadwal  int    `gorm:"not null"`
//	NimDinus  string `gorm:"not null"`
//	Sts       string `gorm:"type:char(1);not null"`
//	Sks       int    `gorm:"not null"`
//	Modul     int    `gorm:"not null;default:0"`
//	TahunAjaran TahunAjaran `gorm:"foreignKey:TA;references:Kode"`
//	MataKuliah MatkulKurikulum `gorm:"foreignKey:Kdmk;references:Kdmk"`
//	Jadwal     JadwalTawar `gorm:"foreignKey:IDJadwal;references:ID"`
//	Mahasiswa  MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
//}
//
//KrsRecordLog represents the krs_record_log table
//type KrsRecordLog struct {
//	IDKrs      int       `gorm:"default:null"`
//	NimDinus   string    `gorm:"default:null"`
//	Kdmk       string    `gorm:"default:null"`
//	Aksi       int8      `gorm:"default:null;comment:1=insert,2=delete"`
//	IDJadwal   int       `gorm:"default:null"`
//	IpAddr     string    `gorm:"default:null"`
//	LastUpdate time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
//	KrsRecord  KrsRecord `gorm:"foreignKey:IDKrs;references:ID"`
//	Mahasiswa  MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
//	MataKuliah MatkulKurikulum `gorm:"foreignKey:Kdmk;references:Kdmk"`
//}
//
//// JadwalTawar represents the jadwal_tawar table
//type JadwalTawar struct {
//	ID         int    `gorm:"primaryKey;autoIncrement"`
//	TA         int    `gorm:"not null;default:0"`
//	Kdmk       string `gorm:"not null"`
//	Klpk       string `gorm:"not null"`
//	Klpk2      string `gorm:"default:null"`
//	Kdds       int    `gorm:"not null"`
//	Kdds2      int    `gorm:"default:null"`
//	Jmax       int    `gorm:"default:0"`
//	Jsisa      int    `gorm:"default:0"`
//	IDHari1    int8   `gorm:"not null"`
//	IDHari2    int8   `gorm:"not null"`
//	IDHari3    int8   `gorm:"not null"`
//	IDSesi1    int    `gorm:"not null"`
//	IDSesi2    int    `gorm:"not null"`
//	IDSesi3    int    `gorm:"not null"`
//	IDRuang1   int    `gorm:"not null"`
//	IDRuang2   int    `gorm:"not null"`
//	IDRuang3   int    `gorm:"not null"`
//	JnsJam     int8   `gorm:"not null;comment:1=pagi, 2=malam, 3=pagi-malam"`
//	OpenClass  bool   `gorm:"not null;default:1;comment:kelas dibuka utk KRS : 1 = open; 0 = close"`
//	MataKuliah MatkulKurikulum `gorm:"foreignKey:Kdmk;references:Kdmk"`
//	Hari1      Hari `gorm:"foreignKey:IDHari1;references:ID"`
//	Hari2      Hari `gorm:"foreignKey:IDHari2;references:ID"`
//	Hari3      Hari `gorm:"foreignKey:IDHari3;references:ID"`
//	Sesi1      SesiKuliah `gorm:"foreignKey:IDSesi1;references:ID"`
//	Sesi2      SesiKuliah `gorm:"foreignKey:IDSesi2;references:ID"`
//	Sesi3      SesiKuliah `gorm:"foreignKey:IDSesi3;references:ID"`
//	Ruang1     Ruang `gorm:"foreignKey:IDRuang1;references:ID"`
//	Ruang2     Ruang `gorm:"foreignKey:IDRuang2;references:ID"`
//	Ruang3     Ruang `gorm:"foreignKey:IDRuang3;references:ID"`
//}
//
//MhsIjinKrs represents the mhs_ijin_krs table
//type MhsIjinKrs struct {
//	ID        int       `gorm:"primaryKey;autoIncrement"`
//	TA        int       `gorm:"not null;default:0"`
//	NimDinus  string    `gorm:"default:null"`
//	Ijinkan   bool      `gorm:"default:null"`
//	Time      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
//	TahunAjaran TahunAjaran `gorm:"foreignKey:TA;references:Kode"`
//	Mahasiswa MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
//}
//
//HerregistMahasiswa represents the herregist_mahasiswa table
//type HerregistMahasiswa struct {
//	ID        int       `gorm:"primaryKey;autoIncrement"`
//	NimDinus  string    `gorm:"default:null"`
//	TA        int       `gorm:"not null;default:0"`
//	DateReg   time.Time `gorm:"default:null"`
//	Mahasiswa MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
//	TahunAjaran TahunAjaran `gorm:"foreignKey:TA;references:Kode"`
//}
//
//MhsDipaketkan represents the mhs_dipaketkan table
//type MhsDipaketkan struct {
//	NimDinus    string `gorm:"primaryKey"`
//	TaMasukMhs  int
//	Mahasiswa   MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
//}
//
//Ruang represents the ruang table
//type Ruang struct {
//	ID            int     `gorm:"primaryKey;autoIncrement"`
//	Nama          string  `gorm:"not null"`
//	Nama2         string  `gorm:"default:'-'"`
//	IDJenisMakul  int     `gorm:"default:null"`
//	IDFakultas    string  `gorm:"default:null"`
//	Kapasitas     int     `gorm:"default:0"`
//	KapUjian      int     `gorm:"default:0"`
//	Status        int16   `gorm:"default:1;comment:1: buka 0: tutup 2: hapus"`
//	Luas          string  `gorm:"default:'0';comment:meter persegi"`
//	Kondisi       string  `gorm:"default:null"`
//	Jumlah        int     `gorm:"default:null"`
//}
//
//TahunAjaran represents the tahun_ajaran table
//type TahunAjaran struct {
//	ID              int64     `gorm:"primaryKey;autoIncrement"`
//	Kode            int       `gorm:"unique;not null"`
//	TahunAkhir      string    `gorm:"not null"`
//	TahunAwal       string    `gorm:"not null"`
//	JnsSmt          int       `gorm:"not null;comment:1 = reg ganjil, 2 = reg genap, 3 = sp ganjil, 4 = sp genap"`
//	SetAktif        bool      `gorm:"not null"`
//	BikuTagihJenis  int8      `gorm:"default:0;comment:1 = spp; 2 = sks; 3 = kekurangan"`
//	UpdateTime      time.Time `gorm:"default:null"`
//	UpdateID        string    `gorm:"default:null"`
//	UpdateHost      string    `gorm:"default:null"`
//	AddedTime       time.Time `gorm:"default:null"`
//	AddedID         string    `gorm:"default:null"`
//	AddedHost       string    `gorm:"default:null"`
//	TglMasuk        time.Time `gorm:"type:date;default:null"`
//}
//
//DaftarNilai represents the daftar_nilai table
//type DaftarNilai struct {
//	ID       int    `gorm:"column:_id;primaryKey;autoIncrement"`
//	NimDinus string `gorm:"default:null"`
//	Kdmk     string `gorm:"default:null"`
//	Nl       string `gorm:"type:char(2);default:null"`
//	Hide     int16  `gorm:"default:0;comment:0 = nilai muncul;1 = nilai disembunyikan (utk keperluan spt hapus mata kuliah)"`
//	Mahasiswa MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
//	MataKuliah MatkulKurikulum `gorm:"foreignKey:Kdmk;references:Kdmk"`
//}
//
//ValidasiKrsMhs represents the validasi_krs_mhs table
//type ValidasiKrsMhs struct {
//	ID        int       `gorm:"primaryKey;autoIncrement"`
//	NimDinus  string    `gorm:"not null"`
//	JobDate   time.Time `gorm:"default:null"`
//	JobHost   string    `gorm:"default:null"`
//	JobAgent  string    `gorm:"default:null"`
//	TA        int       `gorm:"not null;default:0"`
//	Mahasiswa MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
//	TahunAjaran TahunAjaran `gorm:"foreignKey:TA;references:Kode"`
//}
