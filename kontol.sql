# QUERY SAVE

SELECT *
FROM matkul_kurikulum
where kdmk = 'A15.19406';
SELECT *
FROM jadwal_tawar;
select *
from matkul_kurikulum;

desc matkul_kurikulum;


select *
from krs_record;
select *
from krs_record_log;


use krs_management;

select *
from tagihan_mhs
where nim_dinus = '4cd489ef6697a6b3cdf790f8474ec443';

select *
from mahasiswa_dinus;
select *
from mahasiswa_dinus
where kelas = 2;


select *
from krs_record;
select *
from krs_record
where nim_dinus = '1a4421a533b58bb95212ca38610c13de';

#84869e359674baf5c549cf07f5264884

select *
from krs_record;

select *
from mahasiswa_dinus
where akdm_stat = 1;
select *
from mahasiswa_dinus
where nim_dinus = '0024fc24720ea15ae00c812a6eb8e866';
select *
from krs_record
where nim_dinus = '00435514a5421daa10aa604fb04134ed';

#MENGECEK NIM YG EXIST DI KEDUA TABEL
SELECT DISTINCT m.nim_dinus
FROM mahasiswa_dinus m
         INNER JOIN krs_record k ON m.nim_dinus = k.nim_dinus;

select *
from krs_record;


select *
from krs_record
where nim_dinus = '007f2b7dfa36210fb07bd7fe847f5bf3';
select *
from mahasiswa_dinus
where nim_dinus = '1a4421a533b58bb95212ca38610c13de';
select *
from krs_record
where nim_dinus = '007f2b7dfa36210fb07bd7fe847f5bf3';

SELECT *
from krs_record
where nim_dinus = '1a4421a533b58bb95212ca38610c13de';

SELECT id,
       ta,
       kdmk,
       id_jadwal,
       nim_dinus,
       sts,
       sks,
       modul
FROM krs_record
WHERE nim_dinus = '1a4421a533b58bb95212ca38610c13de';

#SAVE FOR 2
SELECT mk.nmmk         AS NamaMatkul,
       mk.nmen         AS NamaMatkulEN,
       mk.tp           AS Tipe,
       mk.smt          AS Semester,
       mk.jenis_matkul AS JenisMatkul,
       h1.nama         AS Hari1,
       h2.nama         AS Hari2,
       h3.nama         AS Hari3
FROM krs_record kr
         JOIN
     matkul_kurikulum mk ON mk.kdmk = kr.kdmk
         JOIN
     jadwal_tawar jt ON jt.id = kr.id_jadwal
         LEFT JOIN
     hari h1 ON h1.id = jt.id_hari1
         LEFT JOIN
     hari h2 ON h2.id = jt.id_hari2
         LEFT JOIN
     hari h3 ON h3.id = jt.id_hari3
WHERE kr.nim_dinus = '1a4421a533b58bb95212ca38610c13de';


SELECT *
from mahasiswa_dinus;
SELECT *
from herregist_mahasiswa;
SELECT *
from tagihan_mhs;

#SAVE QUERY
SELECT md.nim_dinus,
       md.ta_masuk,
       md.prodi,
       md.akdm_stat,
       md.kelas,
       he.date_reg,
       tm.spp_bayar,
       tm.spp_status,
       tm.spp_transaksi
FROM mahasiswa_dinus md
         JOIN herregist_mahasiswa he ON md.nim_dinus = he.nim_dinus
         JOIN krs_management.tagihan_mhs tm on he.nim_dinus = tm.nim_dinus
    WHERE he.nim_dinus = '262019ecd15e0169f7efdea9a64ad30e'
;


select * from validasi_krs_mhs;
select * from mhs_dipaketkan;
select * from ip_semester;

SELECT *,
       CASE
           WHEN job_date <= NOW() THEN 'Validated'
           ELSE 'Not Validated'
           END AS validation_status
FROM validasi_krs_mhs;

# ta
# ketika ada di mhs di paketkan berarti sudah di paketkan
# sks, ips
select * from ip_semester;

select ip_s.ta,
       ip_s.sks,
       ip_s.ips,
       md.ta_masuk
from ip_semester ip_s  join mahasiswa_dinus md where ip_s.nim_dinus = 'cc9d8a25e4226f36a0e6f30abe3420f1'


select * from matkul_kurikulum; # wehere aktif = 1
select * from jadwal_tawar;
select * from ruang;
select * from hari;
select * from sesi_kuliah;
select * from tahun_ajaran;
select * from mahasiswa_dinus;

SELECT
    jt.ta AS tahun_ajaran,
    jt.klpk AS kelompok,
    mk.nmmk AS nama_mata_kuliah,
    mk.sks AS jumlah_sks,
    h.nama AS hari,
    sk.jam_mulai,
    sk.jam_selesai,
    r.nama AS ruang
FROM jadwal_tawar jt
         JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
         JOIN hari h ON jt.id_hari1 = h.id
         JOIN sesi_kuliah sk ON jt.id_sesi1 = sk.id
         JOIN ruang r ON jt.id_ruang1 = r.id
WHERE
    mk.kur_aktif = 1 AND -- Hanya kurikulum aktif
    jt.ta = '20232'; -- Semester aktif (contoh tahun ajaran)




desc tahun_ajaran;
desc sesi_kuliah;

select * from jadwal_input_krs where prodi = 'D22';

select * from  mhs_ijin_krs where nim_dinus = '6f41ddf2e566f37089dd0e2f5fdbeca1'