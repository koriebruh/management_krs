package domain

//type Hari struct {
//	ID     uint8  `gorm:"primaryKey"`
//	Nama   string `gorm:"size:6;not null"`
//	NamaEn string `gorm:"size:20;not null"`
//}

//type SesiKuliah struct {
//	ID         uint16 `gorm:"primaryKey;autoIncrement"`
//	Jam        string `gorm:"size:11;not null;default:''"`
//	SKS        uint8  `gorm:"not null;default:0"`
//	JamMulai   *time.Time
//	JamSelesai *time.Time
//	Status     *int `gorm:"default:1;comment:0=tidak valid, 1= jam valid(kelipatan 50menit), 2 = jam yang harusnya tisak di pakai(jam istirahat)"`
//}

//type Ruang struct {
//	ID           uint    `gorm:"primaryKey;autoIncrement"`
//	Nama         string  `gorm:"size:250;not null"`
//	Nama2        string  `gorm:"size:250;default:'-'"`
//	IDJenisMakul *uint   `gorm:"size:11"`
//	IDFakultas   *string `gorm:"size:5"`
//	Kapasitas    *int    `gorm:"size:3;default:0"`
//	KapUjian     *int    `gorm:"size:3;default:0"`
//	Status       *int8   `gorm:"default:1;comment:1: buka 0: tutup 2: hapus"`
//	Luas         string  `gorm:"size:5;default:'0';comment:meter persegi"`
//	Kondisi      *string `gorm:"size:50"`
//	Jumlah       *int    `gorm:"size:5"`
//}

//type TahunAjaran struct {
//	ID             uint64 `gorm:"primaryKey;autoIncrement"`
//	Kode           string `gorm:"not null;uniqueIndex"`
//	TahunAkhir     string `gorm:"not null"`
//	TahunAwal      string `gorm:"not null"`
//	JnsSmt         int8   `gorm:"not null;comment:1 = reg ganjil, 2 = reg genap, 3 = sp ganjil, 4 = sp genap"`
//	SetAktif       int8   `gorm:"not null"`
//	BikuTagihJenis *int8  `gorm:"default:0;comment:1 = spp; 2 = sks; 3 = kekurangan"`
//	UpdateTime     *time.Time
//	UpdateID       *string `gorm:"size:18"`
//	UpdateHost     *string `gorm:"size:18"`
//	AddedTime      *time.Time
//	AddedID        *string `gorm:"size:18"`
//	AddedHost      *string `gorm:"size:18"`
//	TglMasuk       *time.Time
//	TahunAjarans   []TahunAjaran `gorm:"foreignkey:Kode"`
//}

//type MahasiswaDinus struct {
//	ID       uint   `gorm:"primaryKey;autoIncrement"`
//	NimDinus string `gorm:"size:50;not null;default:'';uniqueIndex"`
//	TaMasuk  *int
//	Prodi    *string `gorm:"size:5"`
//	PassMhs  *string `gorm:"size:128"`
//	Kelas    int8    `gorm:"not null;comment:0 = not choose,1 = pagi,2 = malam,3 = karyawan/pegawai"`
//	AkdmStat string  `gorm:"size:2;not null;comment:1 = aktif, 2 = cuti, 3 = keluar, 4 = lulus, 5 = mangkir, 6 = meninggal, 7 = DO, 8 = Aktif Keuangan;index:STS_AKD_MHS"`
//}

//type MatkulKurikulum struct {
//	KurID         *uint   `gorm:"size:11"`
//	Kdmk          string  `gorm:"primaryKey;size:255"`
//	Nmmk          *string `gorm:"size:255"`
//	Nmen          *string `gorm:"size:255"`
//	TP            *string `gorm:"type:enum('T', 'P', 'TP')"`
//	SKS           *int    `gorm:"size:7"`
//	SKST          *int16  `gorm:"size:3"`
//	SKSP          *int16  `gorm:"size:3"`
//	Smt           *uint
//	JnsSmt        *int8
//	Aktif         *int8
//	KurNama       *string `gorm:"size:255"`
//	KelompokMakul *string `gorm:"type:enum('MPK', 'MKK', 'MKB', 'MKD', 'MBB', 'MPB')"`
//	KurAktif      *bool
//	JenisMatkul   *string `gorm:"type:enum('wajib', 'pilihan')"`
//}

//type JadwalTawar struct {
//	ID              uint    `gorm:"primaryKey;autoIncrement"`
//	TA              int     `gorm:"not null;default:0"`
//	Kdmk            string  `gorm:"size:15;not null"`
//	Klpk            string  `gorm:"size:15;not null"`
//	Klpk2           *string `gorm:"size:15"`
//	Kdds            int     `gorm:"not null"`
//	Kdds2           *int
//	Jmax            *int            `gorm:"size:3;default:0"`
//	Jsisa           *int            `gorm:"size:3;default:0"`
//	IDHari1         uint8           `gorm:"not null"`
//	IDHari2         uint8           `gorm:"not null"`
//	IDHari3         uint8           `gorm:"not null"`
//	IDSesi1         int             `gorm:"size:3;not null"`
//	IDSesi2         int             `gorm:"size:3;not null"`
//	IDSesi3         int             `gorm:"size:3;not null"`
//	IDRuang1        int             `gorm:"size:3;not null"`
//	IDRuang2        int             `gorm:"size:3;not null"`
//	IDRuang3        int             `gorm:"size:3;not null"`
//	JnsJam          uint8           `gorm:"not null;comment:1=pagi, 2=malam, 3=pagi-malam"`
//	OpenClass       uint8           `gorm:"not null;default:1;comment:kelas dibuka utk KRS : 1 = open; 0 = close"`
//	Hari1           Hari            `gorm:"foreignKey:IDHari1"`
//	Hari2           Hari            `gorm:"foreignKey:IDHari2"`
//	Hari3           Hari            `gorm:"foreignKey:IDHari3"`
//	Sesi1           SesiKuliah      `gorm:"foreignKey:IDSesi1"`
//	Sesi2           SesiKuliah      `gorm:"foreignKey:IDSesi2"`
//	Sesi3           SesiKuliah      `gorm:"foreignKey:IDSesi3"`
//	Ruang1          Ruang           `gorm:"foreignKey:IDRuang1"`
//	Ruang2          Ruang           `gorm:"foreignKey:IDRuang2"`
//	Ruang3          Ruang           `gorm:"foreignKey:IDRuang3"`
//	MatkulKurikulum MatkulKurikulum `gorm:"foreignKey:Kdmk"`
//}

//type TagihanMhs struct {
//	ID             uint       `gorm:"primaryKey;autoIncrement;comment:id biasa"`
//	TA             int        `gorm:"not null;comment:tahun ajaran"`
//	NimDinus       string     `gorm:"size:50;not null;comment:nim mahasiswa;uniqueIndex:nim"`
//	SPPBank        *string    `gorm:"size:11"`
//	SPPBayar       int8       `gorm:"not null;default:0;comment:status bayar spp 1: bayar 0: belum bayar"`
//	SPPBayarDate   *time.Time `gorm:"comment:tanggal pada saat operator input pembayaran"`
//	SPPDispensasi  *int
//	SPPHost        *string        `gorm:"size:25;comment:ip/host operator"`
//	SPPStatus      int8           `gorm:"not null;comment:1 : full payment 0 : dispensasi"`
//	SPPTransaksi   *string        `gorm:"size:20;comment:jenis pembayaran : langsung, transfer"`
//	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
//	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
//}

//type IPSemester struct {
//	ID             uint   `gorm:"primaryKey;autoIncrement"`
//	TA             int    `gorm:"not null;default:0;uniqueIndex:nim"`
//	NimDinus       string `gorm:"size:50;not null;uniqueIndex:nim"`
//	SKS            int    `gorm:"not null"`
//	IPS            string `gorm:"size:5;not null"`
//	LastUpdate     *time.Time
//	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
//	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
//}

//type KRSRecord struct {
//	ID              uint64          `gorm:"primaryKey;autoIncrement"`
//	TA              string          `gorm:"not null;index:PERIODE"`
//	Kdmk            string          `gorm:"size:255;not null"`
//	IDJadwal        uint            `gorm:"not null"`
//	NimDinus        string          `gorm:"size:50;not null;index:MAHASISWA"`
//	Sts             string          `gorm:"size:1;not null"`
//	SKS             int             `gorm:"not null"`
//	Modul           int8            `gorm:"not null;default:0"`
//	TahunAjaran     TahunAjaran     `gorm:"foreignkey:TA"`
//	MatkulKurikulum MatkulKurikulum `gorm:"foreignkey:Kdmk"`
//	JadwalTawar     JadwalTawar     `gorm:"foreignkey:IDJadwal"`
//	MahasiswaDinus  MahasiswaDinus  `gorm:"foreignkey:NimDinus"`
//}

//type KRSRecordLog struct {
//	IDKRS           *uint
//	NimDinus        *string `gorm:"size:50"`
//	Kdmk            *string `gorm:"size:255"`
//	Aksi            *int8   `gorm:"comment:1=insert,2=delete"`
//	IDJadwal        *uint
//	IPAddr          *string          `gorm:"size:50"`
//	LastUpdate      time.Time        `gorm:"autoCreateTime"`
//	KRSRecord       *KRSRecord       `gorm:"foreignkey:IDKRS"`
//	MahasiswaDinus  *MahasiswaDinus  `gorm:"foreignkey:NimDinus"`
//	MatkulKurikulum *MatkulKurikulum `gorm:"foreignkey:Kdmk"`
//}

//type MhsIjinKRS struct {
//	ID             uint    `gorm:"primaryKey;autoIncrement"`
//	TA             *int    `gorm:"uniqueIndex:nim"`
//	NimDinus       *string `gorm:"size:50;uniqueIndex:nim"`
//	Ijinkan        *int8
//	Time           time.Time      `gorm:"autoCreateTime"`
//	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
//	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
//}

//type HerregistMahasiswa struct {
//	ID             uint    `gorm:"primaryKey;autoIncrement"`
//	NimDinus       *string `gorm:"size:50"`
//	TA             *string `gorm:"size:5"`
//	DateReg        *time.Time
//	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
//	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
//}

//type MhsDipaketkan struct {
//	NimDinus       string `gorm:"primaryKey;size:50"`
//	TAMasukMhs     *int
//	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
//}

//type SesiKuliahBentrok struct {
//	ID        uint16     `gorm:"primaryKey"`
//	IDBentrok uint16     `gorm:"primaryKey"`
//	Sesi      SesiKuliah `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
//	Bentrok   SesiKuliah `gorm:"foreignKey:IDBentrok;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;index:FK_sesi_kuliah_bentrok2"`
//}

//RANCU NI ?

//type DaftarNilai struct {
//	ID              uint            `gorm:"primaryKey;autoIncrement"`
//	NimDinus        *string         `gorm:"size:50;index:nim"`
//	Kdmk            *string         `gorm:"size:20;index:nim"`
//	Nl              *string         `gorm:"size:2"`
//	Hide            *int8           `gorm:"default:0;comment:0 = nilai muncul; 1 = nilai disembunyikan (utk keperluan spt hapus mata kuliah)"`
//	MahasiswaDinus  MahasiswaDinus  `gorm:"foreignkey:NimDinus"`
//	MatkulKurikulum MatkulKurikulum `gorm:"foreignkey:Kdmk"`
//}

//type ValidasiKRSMhs struct {
//	ID             uint   `gorm:"primaryKey;autoIncrement"`
//	NimDinus       string `gorm:"size:50;not null"`
//	JobDate        *time.Time
//	JobHost        *string `gorm:"size:255"`
//	JobAgent       *string `gorm:"size:255"`
//	TA             *int
//	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
//	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
//}
