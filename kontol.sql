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


select *
from validasi_krs_mhs;
select *
from mhs_dipaketkan;
select *
from ip_semester;

SELECT *,
       CASE
           WHEN job_date <= NOW() THEN 'Validated'
           ELSE 'Not Validated'
           END AS validation_status
FROM validasi_krs_mhs;

# ta
# ketika ada di mhs di paketkan berarti sudah di paketkan
# sks, ips
select *
from ip_semester;

select ip_s.ta,
       ip_s.sks,
       ip_s.ips,
       md.ta_masuk
from ip_semester ip_s
         join mahasiswa_dinus md
where ip_s.nim_dinus = 'cc9d8a25e4226f36a0e6f30abe3420f1';


select *
from matkul_kurikulum; # wehere aktif = 1
select *
from jadwal_tawar;
select *
from ruang;
select *
from hari;
select *
from sesi_kuliah;
select *
from tahun_ajaran;
select *
from mahasiswa_dinus;

SELECT distinct jt.ta   AS tahun_ajaran,
                jt.klpk AS kelompok,
                mk.nmmk AS nama_mata_kuliah,
                mk.sks  AS jumlah_sks,
                h.nama  AS hari,
                sk.jam_mulai,
                sk.jam_selesai,
                r.nama  AS ruang
FROM jadwal_tawar jt
         JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
         JOIN hari h ON jt.id_hari1 = h.id
         JOIN sesi_kuliah sk ON jt.id_sesi1 = sk.id
         JOIN ruang r ON jt.id_ruang1 = r.id;
WHERE mk.kur_aktif = 1
  AND -- Hanya kurikulum aktif
    jt.ta = '20232'; -- Semester aktif (contoh tahun ajaran)

SELECT DISTINCT jt.ta   AS tahun_ajaran,
                jt.klpk AS kelompok,
                mk.nmmk AS nama_mata_kuliah,
                mk.sks  AS jumlah_sks,
                h.nama  AS hari,
                sk.jam_mulai,
                sk.jam_selesai,
                r.nama  AS ruang
FROM jadwal_tawar jt
         LEFT JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
         LEFT JOIN hari h ON jt.id_hari1 = h.id
         LEFT JOIN sesi_kuliah sk ON jt.id_sesi1 = sk.id
         LEFT JOIN ruang r ON jt.id_ruang1 = r.id
WHERE jt.ta = '20232';



select *
from hari;

select *
from tahun_ajaran;
select *
from jadwal_tawar
where ta = '20232';
desc sesi_kuliah;

select *
from jadwal_input_krs
where prodi = 'D22';

select *
from mhs_ijin_krs
where nim_dinus = '6f41ddf2e566f37089dd0e2f5fdbeca1';


SELECT *
FROM daftar_nilai;
SELECT *
FROM matkul_kurikulum;

SELECT mk.kdmk         AS kode_matkul,
       mk.nmen         AS matakuliah,
       mk.sks          AS sks,
       mk.tp           AS kategory,
       mk.jenis_matkul AS jenis_maktkul,
       dn.nl           AS nilai
FROM matkul_kurikulum mk
         JOIN daftar_nilai dn ON mk.kdmk = dn.kdmk
WHERE nim_dinus = '6f41ddf2e566f37089dd0e2f5fdbeca1'
  AND dn.hide = 0;

select *
from matkul_kurikulum;
select *
from jadwal_tawar
where id_hari2; #where open_class = 1 AND where sisa < jmax AND where jns_jam in (1,2)
select *
from ruang;
select *
from hari;
select *
from sesi_kuliah;
select *
from krs_record;

#SAVE DULU
#UNUTK CEK DATA APAAH ADA TEELAH DI TAMBHAKAN BY ID
SELECT kr.kdmk,
       mk.nmmk AS nama_mata_kuliah,
       h.nama  AS hari,
       sk.jam_mulai,
       sk.jam_selesai
FROM krs_record kr
         JOIN jadwal_tawar jt ON kr.id_jadwal = jt.id
         JOIN hari h ON jt.id_hari1 = h.id
         JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
         JOIN matkul_kurikulum mk ON kr.kdmk = mk.kdmk
WHERE kr.nim_dinus = '17d3264f6edf07182311d5cd19e1cd0a';

#FINAL QUERY
SELECT DISTINCT jt.ta   AS tahun_ajaran,
                jt.klpk AS kelompok,
                mk.nmmk AS nama_mata_kuliah,
                mk.sks  AS jumlah_sks,
                h.nama  AS hari,
                sk.jam_mulai,
                sk.jam_selesai,
                r.nama  AS ruang,
                CASE
                    WHEN EXISTS (SELECT 1
                                 FROM krs_record kr
                                          JOIN jadwal_tawar jt_inner ON kr.id_jadwal = jt_inner.id
                                          JOIN sesi_kuliah sk_inner ON sk_inner.id = jt_inner.id_sesi1
                                 WHERE kr.nim_dinus = '17d3264f6edf07182311d5cd19e1cd0a'
                                   AND jt.id_hari1 = jt_inner.id_hari1 -- Hari yang sama
                                   AND (
                                     (sk.jam_mulai < sk_inner.jam_selesai AND
                                      sk.jam_selesai > sk_inner.jam_mulai) -- Jam BENTROK
                                     )) THEN 'BENTROK'
                    ELSE NULL
                    END AS status_bentrok,
                CASE
                    WHEN jt.jsisa = jt.jmax THEN CONCAT(jt.jsisa, '/', jt.jmax, ' SLOT PENUH')
                    ELSE CONCAT(jt.jsisa, '/', jt.jmax)
                    END AS keterangan_slot
FROM jadwal_tawar jt
         JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
         JOIN hari h ON jt.id_hari1 = h.id
         JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
         JOIN ruang r ON jt.id_ruang1 = r.id
WHERE mk.kur_aktif = 1     -- Hanya kurikulum aktif
  AND jt.ta = '20232'      -- Kode Tahun Ajaran
  AND jt.jns_jam IN (1, 2) -- Untuk kelas pagi atau malam
  AND jt.jsisa <= jt.jmax;
-- Memastikan kuotanya kosong atau penuh

SELECT DISTINCT mk.nmmk
FROM matkul_kurikulum mk;


#CHECK
SELECT jt.ta
FROM jadwal_tawar jt
         JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
         JOIN mahasiswa_dinus
WHERE nim_dinus = '647e27c32c8935273e876a457b81b186';


SELECT *
FROM daftar_nilai
where nim_dinus = '647e27c32c8935273e876a457b81b186';

SELECT DISTINCT jt.ta   AS tahun_ajaran,
                jt.kdmk AS kode_mata_kuiah,
                jt.klpk AS kelompok,
                mk.nmmk AS nama_mata_kuliah,
                mk.sks  AS jumlah_sks,
                h.nama  AS hari,
                sk.jam_mulai,
                sk.jam_selesai,
                r.nama  AS ruang,
                CASE
                    WHEN EXISTS (SELECT 1
                                 FROM daftar_nilai dn
                                 WHERE dn.kdmk = jt.kdmk
                                   AND dn.nl = 'A'
                                   AND dn.nim_dinus = '647e27c32c8935273e876a457b81b186') THEN 'Tidak Bisa'
                    ELSE 'Bisa'
                    END AS status_pemilihan
FROM jadwal_tawar jt
         JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
         JOIN hari h ON jt.id_hari1 = h.id
         JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
         JOIN ruang r ON jt.id_ruang1 = r.id
WHERE jt.ta IS NOT NULL -- Pastikan hanya menampilkan data valid
ORDER BY jt.ta, mk.nmmk;


SELECT DISTINCT jt.ta   AS tahun_ajaran,
                jt.kdmk AS kode_mata_kuiah,
                jt.klpk AS kelompok,
                mk.nmmk AS nama_mata_kuliah,
                mk.sks  AS jumlah_sks,
                h.nama  AS hari,
                sk.jam_mulai,
                sk.jam_selesai,
                r.nama  AS ruang,
                CASE
                    WHEN EXISTS (SELECT 1
                                 FROM daftar_nilai dn
                                 WHERE dn.kdmk = jt.kdmk
                                   AND dn.nl = 'A'
                                   AND dn.nim_dinus = '647e27c32c8935273e876a457b81b186') THEN 'Tidak Bisa'
                    ELSE 'Bisa'
                    END AS status_pemilihan
FROM jadwal_tawar jt
         LEFT JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
         LEFT JOIN hari h ON jt.id_hari1 = h.id
         LEFT JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
         LEFT JOIN ruang r ON jt.id_ruang1 = r.id
WHERE jt.ta IS NOT NULL   -- Pastikan hanya menampilkan data valid
  AND jt.klpk LIKE 'B11%' -- Hanya tampilkan kelompok yang dimulai dengan 'B11'
  AND jt.ta = '20232'
ORDER BY jt.ta, mk.nmmk;


select *
from ip_semester
where nim_dinus = '647e27c32c8935273e876a457b81b186';

select *
from mahasiswa_dinus
where nim_dinus = '647e27c32c8935273e876a457b81b186';


select *
from krs_record;
# KETIKA total sementara + sks matakuliah_kurikulum > sks_max_user maka tambahkan field krs  tidak mencukupi
select sum(sks) as total_sks_sementara
from krs_record
where nim_dinus = '647e27c32c8935273e876a457b81b186';
select sum(sks) as sks_max_user
from ip_semester
where nim_dinus = '647e27c32c8935273e876a457b81b186'
ORDER BY last_update
limit 1;
select sks
from ip_semester
where nim_dinus = '647e27c32c8935273e876a457b81b186'
ORDER BY last_update
limit 1;
select *
from ip_semester;

SELECT DISTINCT jt.ta   AS tahun_ajaran,
                jt.kdmk AS kode_mata_kuliah,
                jt.klpk AS kelompok,
                mk.nmmk AS nama_mata_kuliah,
                mk.sks  AS jumlah_sks,
                h.nama  AS hari,
                sk.jam_mulai,
                sk.jam_selesai,
                r.nama  AS ruang,
                CASE
                    WHEN EXISTS (SELECT 1
                                 FROM daftar_nilai dn
                                 WHERE dn.kdmk = jt.kdmk
                                   AND dn.nl = 'A'
                                   AND dn.nim_dinus = '647e27c32c8935273e876a457b81b186') THEN 'Tidak Bisa'
                    ELSE 'Bisa'
                    END AS status_pemilihan,
                CASE
                    WHEN (
                             (SELECT COALESCE(SUM(sks), 0)
                              FROM krs_record
                              WHERE nim_dinus = '647e27c32c8935273e876a457b81b186')
                                 + mk.sks
                             ) > (SELECT COALESCE(MAX(sks), 0)
                                  FROM ip_semester
                                  WHERE nim_dinus = '647e27c32c8935273e876a457b81b186'
                                  ORDER BY last_update
                                  limit 1) THEN 'Tidak Mencukupi'
                    ELSE CONCAT(
                            'Jika di ambil Sisa ',
                            (SELECT COALESCE(MAX(sks), 0)
                             FROM ip_semester
                             WHERE nim_dinus = '647e27c32c8935273e876a457b81b186'
                             ORDER BY last_update
                             limit 1)
                                - (SELECT COALESCE(SUM(sks), 0)
                                   FROM krs_record
                                   WHERE nim_dinus = '647e27c32c8935273e876a457b81b186')
                                - mk.sks
                         )
                    END AS status_krs
FROM jadwal_tawar jt
         LEFT JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
         LEFT JOIN hari h ON jt.id_hari1 = h.id
         LEFT JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
         LEFT JOIN ruang r ON jt.id_ruang1 = r.id
WHERE jt.ta IS NOT NULL   -- Pastikan hanya menampilkan data valid
  AND jt.klpk LIKE 'B11%' -- Hanya tampilkan kelompok yang dimulai dengan 'B11'
  AND jt.ta = '20232'
ORDER BY jt.ta, mk.nmmk;


select *
from matkul_kurikulum;

#jt.jns_jam IN (1, 2)

SELECT *
FROM krs_record;
SELECT *
FROM jadwal_tawar;
SELECT DISTINCT *
from krs_record
         JOIN jadwal_tawar
where nim_dinus = '647e27c32c8935273e876a457b81b186'

SELECT jns_jam
FROM jadwal_tawar
WHERE id = '276209';

select *
from jadwal_tawar;


select *
from krs_record;
DESC krs_record;

select jsisa
from jadwal_tawar
where id = '276209';

SELECT *
FROM mhs_ijin_krs;
SELECT *
from ip_semester
where nim_dinus = '00ce7b9909293860b12f4bd86ab26d0e';

#find u ser bisa insert test
SELECT m.nim_dinus
FROM mhs_ijin_krs m
         JOIN ip_semester i ON m.nim_dinus = i.nim_dinus
         JOIN mahasiswa_dinus md ON md.nim_dinus = m.nim_dinus
WHERE akdm_stat = 1;


SELECT mahasiswa_dinus.nim_dinus
FROM mahasiswa_dinus
WHERE mahasiswa_dinus.nim_dinus = '020c6355071b91f8d3eea6442d968525';

SELECT *
FROM jadwal_tawar
         JOIN sesi_kuliah on jadwal_tawar.id_sesi1 = sesi_kuliah.id
WHERE jadwal_tawar.id = 275486;



SELECT DISTINCT jt.id,
                jt.ta   AS tahun_ajaran,
                jt.kdmk AS kode_mata_kuiah,
                jt.klpk AS kelompok,
                mk.nmmk AS nama_mata_kuliah,
                mk.sks  AS jumlah_sks,
                h.nama  AS hari,
                sk.jam_mulai,
                sk.jam_selesai,
                r.nama  AS ruang
FROM jadwal_tawar jt
         LEFT JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
         LEFT JOIN hari h ON jt.id_hari1 = h.id
         LEFT JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
         LEFT JOIN ruang r ON jt.id_ruang1 = r.id
         JOIN krs_record kr On kr.kdmk = mk.kdmk
WHERE jt.ta IS NOT NULL -- Pastikan hanya menampilkan data valid
  AND nim_dinus = '560b4d78fc163d57774e045317be842f';


SELECT krl.*
FROM krs_record_log krl
         left JOIN
     krs_record kr ON krl.id_krs = kr.id
WHERE kr.ta = '20232'
  AND kr.nim_dinus = '9b77bd5b68ed9c7887a81905016731d2';

select * from krs_record_log where nim_dinus = '9b77bd5b68ed9c7887a81905016731d2';


SELECT * FROM validasi_krs_mhs;
SELECT * FROM krs_record where nim_dinus = '9b77bd5b68ed9c7887a81905016731d2';
SELECT * FROM krs_record_log where nim_dinus = '9b77bd5b68ed9c7887a81905016731d2';


select * from validasi_krs_mhs where job_host = 'haha';
desc validasi_krs_mhs;

SELECT * FROM `tahun_ajaran` WHERE kode ='20232' ORDER BY `tahun_ajaran`.`id` LIMIT 1
;
2024-02-20 15:01:19 validate
2024-12-25 10:34:43 not validate
select * from validasi_krs_mhs where job_host = 'haha';

SELECT * FROM validasi_krs_mhs;
SELECT * FROM krs_record;

SELECT DISTINCT a.nim_dinus
FROM validasi_krs_mhs a
         INNER JOIN krs_record b
                    ON a.nim_dinus = b.nim_dinus INNER JOIN mhs_ijin_krs c ON a.nim_dinus = c.nim_dinus;


select * from mhs_ijin_krs where nim_dinus = '176bd4688305a3ae546b35b52aed75c8';

select * from mhs_ijin_krs where ijinkan ;
select * from validasi_krs_mhs where ta != 20232;
select * from krs_record;

SHOW TABLES LIKE 'tagihan_mhs';
select * from tagihan_mhs;

select * from krs_management.tagihan_mhs;
desc krs_management.tagihan_mhs;
desc tagihan_mhs;
select * from tahun_ajaran;

select * from ip_semester where nim_dinus = 'b7ecdd849e627833e97716bb4cfcaed8';

SELECT COUNT(*) AS total
FROM tagihan_mhs;

SELECT COUNT(*) AS total
FROM ip_semester;
SELECT COUNT(*) AS total
FROM mahasiswa_dinus;

select * from herregist_mahasiswa where nim_dinus = '2ef5428b895c0bf6a76f808a7b04baf4';


SELECT mahasiswa_dinus.nim_dinus,
       mahasiswa_dinus.ta_masuk,
       mahasiswa_dinus.prodi,
       mahasiswa_dinus.akdm_stat,
       mahasiswa_dinus.kelas,
       herregist_mahasiswa.date_reg,
       tagihan_mhs.spp_bayar,
       tagihan_mhs.spp_status,
       tagihan_mhs.spp_transaksi
FROM mahasiswa_dinus
         LEFT JOIN herregist_mahasiswa
                   ON mahasiswa_dinus.nim_dinus = herregist_mahasiswa.nim_dinus
         LEFT JOIN tagihan_mhs
                   ON herregist_mahasiswa.nim_dinus = tagihan_mhs.nim_dinus
WHERE mahasiswa_dinus.nim_dinus = '2ef5428b895c0bf6a76f808a7b04baf4';

SELECT krl.*
FROM krs_record_log krl
         LEFT JOIN krs_record kr ON krl.id_krs = kr.id
WHERE kr.ta = '20232' ;

select * from krs_record_log where nim_dinus = '2ef5428b895c0bf6a76f808a7b04baf4';

SELECT krl.*
FROM krs_record_log krl
         LEFT JOIN krs_record kr
                   ON krl.id_krs = kr.id AND kr.ta = '20232' where kr.nim_dinus = '2ef5428b895c0bf6a76f808a7b04baf4';

SELECT *
FROM krs_record
WHERE ta = '20232' AND nim_dinus = '2ef5428b895c0bf6a76f808a7b04baf4';

SELECT *
FROM krs_record_log
WHERE id_krs NOT IN (SELECT id FROM krs_record);
SELECT *
FROM krs_record_log
WHERE id_krs NOT IN (SELECT id FROM krs_record);

desc krs_record de;

select * from krs_record_log where nim_dinus = '2ef5428b895c0bf6a76f808a7b04baf4';
select * from krs_record where nim_dinus = '2ef5428b895c0bf6a76f808a7b04baf4';

select * from krs_record ;
select * from mahasiswa_dinus;


SELECT DISTINCT nim_dinus
FROM (
         SELECT nim_dinus FROM mahasiswa_dinus
         UNION
         SELECT nim_dinus FROM herregist_mahasiswa
         UNION
         SELECT nim_dinus FROM krs_record
         UNION
         SELECT nim_dinus FROM ip_semester
         UNION
         SELECT nim_dinus FROM tagihan_mhs
     ) AS combined_nim_dinus;


SELECT COUNT(*) AS total
FROM mahasiswa_dinus;

drop database krs_management;
create database krs_management;
use krs_management;
show tables ;

select * from ip_semester;

select nim_dinus as nd from mahasiswa_dinus limit 1200;
select * from mahasiswa_dinus where nim_dinus = '0438b811b2c935373dd425540c388ec1';

SHOW INDEX FROM mahasiswa_dinus;
EXPLAIN SELECT * FROM mahasiswa_dinus WHERE nim_dinus = '0386d19d2c53289bd518c30e1bd0fc42' ORDER BY id LIMIT 1;
SHOW VARIABLES LIKE 'thread%';


SHOW VARIABLES LIKE 'max_connections';

select * from herregist_mahasiswa where nim_dinus = '00b4f80268c6cc31f651f31186ee67d9';


desc mahasiswa_dinus;


#NIM UNTUK UJI COBA yang data nya semua ada
SELECT md.nim_dinus
FROM mahasiswa_dinus md
    LEFT JOIN herregist_mahasiswa he ON he.nim_dinus = md.nim_dinus
    RIGHT JOIN krs_record kr ON kr.nim_dinus = md.nim_dinus
    WHERE md.akdm_stat = 1;

