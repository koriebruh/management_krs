use krs_management;
show tables ;

select * from mahasiswa_dinus;

show databases ;

drop database krs_management;
create database krs_management;
use krs_management;
show tables ;

SELECT
    TABLE_NAME,
    COLUMN_NAME,
    CONSTRAINT_NAME,
    REFERENCED_TABLE_NAME,
    REFERENCED_COLUMN_NAME
FROM
    INFORMATION_SCHEMA.KEY_COLUMN_USAGE
WHERE
    REFERENCED_TABLE_NAME = 'matkul_kurikulum' AND
    REFERENCED_COLUMN_NAME = 'kdmk';


SHOW CREATE TABLE matkul_kurikulum;

select * from matkul_kurikulum;
select * from krs_record;
select * from krs_record_log;
select * from daftar_nilai;

desc  tagihan_mhs;
desc  ip_semester;

select * from mahasiswa_dinus;

#DUNGU DATA NIM DI TAGIAN != DAFTAR MAHASIWA
SELECT * FROM mahasiswa_dinus WHERE nim_dinus = '5c2cfdac12475311e3bcbd51121ff877';
SELECT * FROM mahasiswa_dinus WHERE nim_dinus = '381bbb519c5a3287251420a5f338736e';

#
SELECT * FROM mahasiswa_dinus WHERE nim_dinus = '6b2a2a2b68e3b0252ffd026fb0c21666';
SELECT * FROM mahasiswa_dinus WHERE nim_dinus = '95cedfff35b96393991dd55e982120c2';


SELECT * FROM tahun_ajaran WHERE kode = 20201;
select * from tahun_ajaran;

select * from tagihan_mhs;

select * from tahun_ajaran;
select  * from  ip_semester;

select * from mhs_dipaketkan;
select * from ip_semester;

select * from sesi_kuliah;
select * from sesi_kuliah_bentrok;


select * from daftar_nilai WHERE nim_dinus ='6f0581e506eda3d4cd846ef75f5165d2';
4003953,"20232",P31330442,283282,"6f0581e506eda3d4cd846ef75f5165d2",B,6,0,0

select * from matkul_kurikulum;
select * from matkul_kurikulum where kdmk = 'P31330442';

select * from jadwal_tawar;
desc jadwal_tawar;
SHOW CREATE TABLE jadwal_tawar;

desc hari;
select * from hari;
select * from ruang;

desc jadwal_tawar;

select * from jadwal_tawar;

SELECT * FROM sesi_kuliah;
desc jadwal_tawar;


select * from krs_record_log;
DESC krs_record_log;

# delete from sesi_kuliah_bentrok;
SHOW TABLES ;
# delete from krs_record;
select * from krs_record;





ALTER TABLE `jadwal_tawar`
    MODIFY `id_sesi1` INT NULL,
    MODIFY `id_sesi2` INT NULL,
    MODIFY `id_sesi3` INT NULL,
    MODIFY `id_ruang1` INT NULL,
    MODIFY `id_ruang2` INT NULL,
    MODIFY `id_ruang3` INT NULL;


#
desc matkul_kurikulum;

ALTER TABLE `matkul_kurikulum`
    MODIFY `kur_aktif` TINYINT(1);

desc krs_record_log;
ALTER TABLE krs_record_log RENAME COLUMN lastUpdate TO last_update;


select * from mahasiswa_dinus where nim_dinus = '1a4421a533b58bb95212ca38610c13de';

select * from herregist_mahasiswa;

SELECT mahasiswa_dinus.nim_dinus, mahasiswa_dinus.ta_masuk, mahasiswa_dinus.prodi, mahasiswa_dinus.akdm_stat, mahasiswa_dinus.kelas, herregist_mahasiswa.date_reg, tagihan_mhs.spp_bayar, tagihan_mhs.spp_status, tagihan_mhs.spp_transaksi
FROM mahasiswa_dinus
         JOIN herregist_mahasiswa ON mahasiswa_dinus.nim_dinus = herregist_mahasiswa.nim_dinus
         JOIN krs_management.tagihan_mhs ON herregist_mahasiswa.nim_dinus = tagihan_mhs.nim_dinus
WHERE mahasiswa_dinus.nim_dinus = '006f92df0e8baf555ce525162f681678';

SELECT DISTINCT md.nim_dinus
FROM mahasiswa_dinus md
         JOIN herregist_mahasiswa hm ON md.nim_dinus = hm.nim_dinus
         JOIN krs_management.tagihan_mhs tm ON hm.nim_dinus = tm.nim_dinus;

select * from mhs_ijin_krs;
