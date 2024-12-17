package dto

import "time"

type StudentInfo struct {
	AkademikStat     string
	Bayar            bool
	Bank             string
	BayarDate        time.Time
	TransactionId    string
	RegistrationDate time.Time
	Kelas            string
}
