use krs_management;
show tables ;

select * from mahasiswa_dinus;

show databases ;

drop database krs_management;
create database krs_management;
use krs_management;
show tables ;



SHOW TABLES LIKE 'sesi_kuliah_bentroks';

DESCRIBE sesi_kuliah;
DESCRIBE jatwal_tawar;

SHOW TABLES LIKE 'mhs_ijin_krs';


desc tahun_ajaran;
DROP TABLE IF EXISTS mahasiswa_dinus;


desc mahasiswa_dinus;
desc mhs_dipaketkan ;


desc herregis_mahasiswa;
desc mhs_ijin_krs ;

SELECT TABLE_NAME, COLUMN_NAME, DATA_TYPE
FROM INFORMATION_SCHEMA.COLUMNS
WHERE COLUMN_NAME = 'nim_dinus'
  AND TABLE_SCHEMA = 'krs_management';
