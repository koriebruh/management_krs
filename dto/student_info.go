package dto

import "time"

//MANTI LAH SKIP DULU DTO

type StudentInfo struct {
	NimDinus         string
	TahunMasuk       string
	AkademikStat     string
	Bayar            bool
	Bank             string
	BayarDate        time.Time
	TransactionId    string
	RegistrationDate time.Time
	Kelas            string
}
