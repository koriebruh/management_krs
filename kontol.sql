# QUERY SAVE

SELECT * FROM matkul_kurikulum where kdmk = 'A15.19406';
SELECT * FROM jadwal_tawar;
select * from matkul_kurikulum;

desc matkul_kurikulum;


select * from  krs_record;
select * from  krs_record_log;


use krs_management;

select * from tagihan_mhs where nim_dinus = '4cd489ef6697a6b3cdf790f8474ec443';

select * from mahasiswa_dinus;
select * from mahasiswa_dinus where kelas = 2;


select * from krs_record ;
select * from krs_record where nim_dinus = '1a4421a533b58bb95212ca38610c13de';

#84869e359674baf5c549cf07f5264884

select * from krs_record;

select * from mahasiswa_dinus where akdm_stat=1;
select * from mahasiswa_dinus where nim_dinus='0024fc24720ea15ae00c812a6eb8e866';
select * from krs_record where nim_dinus = '00435514a5421daa10aa604fb04134ed';

#MENGECEK NIM YG EXIST DI KEDUA TABEL
SELECT DISTINCT m.nim_dinus
FROM mahasiswa_dinus m
         INNER JOIN krs_record k ON m.nim_dinus = k.nim_dinus;

select * from krs_record;


select * from krs_record where nim_dinus='007f2b7dfa36210fb07bd7fe847f5bf3';
select * from mahasiswa_dinus where nim_dinus='1a4421a533b58bb95212ca38610c13de';
select * from krs_record where nim_dinus='007f2b7dfa36210fb07bd7fe847f5bf3';

SELECT * from krs_record where nim_dinus ='1a4421a533b58bb95212ca38610c13de';