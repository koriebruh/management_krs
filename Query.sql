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
