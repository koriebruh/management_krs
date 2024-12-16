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


select * from daftar_nilai WHERE nim_dinus ='8f218f15d69c6178c2b2a6ccd975ba28';


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


# delete from sesi_kuliah_bentrok;
SHOW TABLES ;
delete from jadwal_tawar;
select * from jadwal_tawar;


ALTER TABLE `jadwal_tawar`
    MODIFY `id_sesi1` INT NULL,
    MODIFY `id_sesi2` INT NULL,
    MODIFY `id_sesi3` INT NULL,
    MODIFY `id_ruang1` INT NULL,
    MODIFY `id_ruang2` INT NULL,
    MODIFY `id_ruang3` INT NULL;


#
