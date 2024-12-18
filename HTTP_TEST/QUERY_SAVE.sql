use krs_management;

SELECT * FROM jadwal_input_krs;
SELECT * FROM jadwal_input_krs;


# INPUT KRS HAL AWAL
# PRELOAD KE domian.heregist_mahasiswa

select * from  mahasiswa_dinus; # AMBIL akdm_start
desc mahasiswa_dinus;
select * from  tagihan_mhs; #AMBIL FIELD ssp_bayar, ssp_bank, ssp_bayar_date DAN ssp_transaction
desc tagihan_mhs;
select * from  herregist_mahasiswa; #AMBIL FIELD date_reg
# WHERE mahasiswa_dinus = d263aa7f29800ba04bc6e8ad674b58cd

SELECT
    md.akdm_stat AS akademik_stat,
    tm.spp_bayar AS bayar,
    tm.spp_bank AS bank,
    tm.spp_bayar_date AS bayar_date,
    tm.spp_transaksi AS transaction_id,
    hr.date_reg AS registrasi_date
FROM
    mahasiswa_dinus md
        INNER JOIN
    tagihan_mhs tm ON md.nim_dinus = tm.nim_dinus
        INNER JOIN
    herregist_mahasiswa hr ON md.nim_dinus = hr.nim_dinus
WHERE
    md.nim_dinus = 'd263aa7f29800ba04bc6e8ad674b58cd';

SELECT * FROM mahasiswa_dinus WHERE nim_dinus = 'd263aa7f29800ba04bc6e8ad674b58cd';
