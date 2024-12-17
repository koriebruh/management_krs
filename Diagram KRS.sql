drop database krs_management;
create database krs_management;
USE krs_management;
show tables;

CREATE TABLE `tagihan_mhs`
(
    `id`             int(50) PRIMARY KEY NOT NULL AUTO_INCREMENT COMMENT 'id biasa',
    `ta`             int(10)             NOT NULL COMMENT 'tahun ajaran',
    `nim_dinus`      varchar(50)         NOT NULL COMMENT 'nim mahasiswa',
    `spp_bank`       varchar(11),
    `spp_bayar`      int(1)              NOT NULL DEFAULT '0' COMMENT 'status bayar spp 1: bayar 0: belum bayar',
    `spp_bayar_date` datetime COMMENT 'tanggal pada saat operator input pembayaran',
    `spp_dispensasi` int(11),
    `spp_host`       varchar(25) COMMENT 'ip/host operator',
    `spp_status`     int(1)              NOT NULL COMMENT '1 : full payment 0 : dispensasi',
    `spp_transaksi`  varchar(20) COMMENT 'jenis pembayaran : langsung, transfer'
);

CREATE TABLE `mahasiswa_dinus`
(
    `id`        int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `nim_dinus` varchar(50)         NOT NULL DEFAULT '',
    `ta_masuk`  int(10),
    `prodi`     varchar(5),
    `pass_mhs`  varchar(128)                 DEFAULT null,
    `kelas`     int(1)              NOT NULL COMMENT '0 = not choose,\r\n1 = pagi,\r\n2 = malam,\r\n3 = karyawan/pegawai',
    `akdm_stat` char(2)             NOT NULL COMMENT '1 = aktif, 2 = cuti, 3 = keluar, 4 = lulus, 5 = mangkir, 6 = meninggal, 7 = DO, 8 = Aktif Keuangan'
);

#BODOAMAD TAMBAH ID BUAT RELASI DATA NYA TOLOL ASU
CREATE TABLE `matkul_kurikulum`
(
    `kur_id`         int(11),
    `kdmk`           varchar(255) NOT NULL,
    `nmmk`           varchar(255),
    `nmen`           varchar(255),
    `tp`             ENUM ('T', 'P', 'TP'),
    `sks`            int(7),
    `sks_t`          smallint(3),
    `sks_p`          smallint(3),
    `smt`            int(11),
    `jns_smt`        int(1),
    `aktif`          tinyint(1),
    `kur_nama`       varchar(255),
    `kelompok_makul` ENUM ('MPK', 'MKK', 'MKB', 'MKD', 'MBB', 'MPB'),
    `kur_aktif`      bit(1),
    `jenis_matkul`   ENUM ('wajib', 'pilihan'),
    PRIMARY KEY (`kur_id`, `kdmk`) -- Gabungan PRIMARY KEY
);

CREATE TABLE `hari`
(
    `id`      tinyint(1)  NOT NULL UNIQUE,
    `nama`    varchar(6)  NOT NULL,
    `nama_en` varchar(20) NOT NULL
);

CREATE TABLE `jadwal_input_krs`
(
    `id`          int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `ta`          int(10)             NOT NULL DEFAULT '0',
    `prodi`       char(3),
    `tgl_mulai`   datetime,
    `tgl_selesai` datetime
);

CREATE TABLE `ip_semester`
(
    `id`          int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `ta`          int(10)             NOT NULL DEFAULT '0',
    `nim_dinus`   varchar(50)         NOT NULL,
    `sks`         int(4)              NOT NULL,
    `ips`         varchar(5)          NOT NULL,
    `last_update` datetime                     DEFAULT null
);

CREATE TABLE `sesi_kuliah`
(
    `id`          int(5) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `jam`         varchar(11)        NOT NULL DEFAULT '',
    `sks`         smallint(1)        NOT NULL DEFAULT '0',
    `jam_mulai`   time                        DEFAULT null,
    `jam_selesai` time                        DEFAULT null,
    `status`      int(11)                     DEFAULT '1' COMMENT '0=tidak valid, 1= jam valid(kelipatan 50menit), 2 = jam yang harusnya tisak di pakai(jam istirahat)'
);

CREATE TABLE `sesi_kuliah_bentrok`
(
    `id`         int(5) NOT NULL,
    `id_bentrok` int(5) NOT NULL,
    PRIMARY KEY (`id`, `id_bentrok`)
);

CREATE TABLE `krs_record`
(
    `id`        int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `ta`        int(10)             NOT NULL DEFAULT '0',
    `kur_id`         int(11),
    `kdmk`      varchar(255)        NOT NULL,
    `id_jadwal` int(11)             NOT NULL,
    `nim_dinus` varchar(50)         NOT NULL,
    `sts`       char(1)             NOT NULL,
    `sks`       int(11)             NOT NULL,
    `modul`     int(1)              NOT NULL DEFAULT '0'
);

CREATE TABLE `krs_record_log`
(
    `id_krs`     int(11)            DEFAULT null,
    `nim_dinus`  varchar(50)        DEFAULT null,
    `kur_id`         int(11),
    `kdmk`       varchar(255)       DEFAULT null,
    `aksi`       tinyint(3)         DEFAULT null COMMENT '1=insert,2=delete',
    `id_jadwal`  int(11)            DEFAULT null,
    `ip_addr`    varchar(50)        DEFAULT null,
    `lastUpdate` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `jadwal_tawar`
(
    `id`         int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `ta`         int(10)             NOT NULL DEFAULT '0',
    `kur_id`         int(11),
    `kdmk`       varchar(15)         NOT NULL,
    `klpk`       varchar(15)         NOT NULL,
    `klpk_2`     varchar(15)                  DEFAULT null,
    `kdds`       int(10)             NOT NULL,
    `kdds2`      int(10)                      DEFAULT null,
    `jmax`       int(3)                       DEFAULT '0',
    `jsisa`      int(3)                       DEFAULT '0',
    `id_hari1`   tinyint(1)          NOT NULL,
    `id_hari2`   tinyint(1)          NOT NULL,
    `id_hari3`   tinyint(1)          NOT NULL,
    `id_sesi1`   int(3)              NOT NULL,
    `id_sesi2`   int(3)              NOT NULL,
    `id_sesi3`   int(3)              NOT NULL,
    `id_ruang1`  int(3)              NOT NULL,
    `id_ruang2`  int(3)              NOT NULL,
    `id_ruang3`  int(3)              NOT NULL,
    `jns_jam`    tinyint(1)          NOT NULL COMMENT '1=pagi, 2=malam, 3=pagi-malam',
    `open_class` tinyint(1)          NOT NULL DEFAULT '1' COMMENT 'kelas dibuka utk KRS : 1 = open; 0 = close'
);

CREATE TABLE `mhs_ijin_krs`
(
    `id`        int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `ta`        int(10)             NOT NULL DEFAULT '0',
    `nim_dinus` varchar(50)                  DEFAULT null,
    `ijinkan`   tinyint(1)                   DEFAULT null,
    `time`      timestamp           NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `herregist_mahasiswa`
(
    `id`        int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `nim_dinus` varchar(50)                  DEFAULT null,
    `ta`        int(10)             NOT NULL DEFAULT '0',
    `date_reg`  datetime                     DEFAULT null
);

CREATE TABLE `mhs_dipaketkan`
(
    `nim_dinus`    varchar(50) PRIMARY KEY NOT NULL,
    `ta_masuk_mhs` int(10)
);

CREATE TABLE `ruang`
(
    `id`             int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `nama`           varchar(250)        NOT NULL,
    `nama2`          varchar(250) DEFAULT '-',
    `id_jenis_makul` int(11)      DEFAULT null,
    `id_fakultas`    varchar(5)   DEFAULT null,
    `kapasitas`      int(3)       DEFAULT '0',
    `kap_ujian`      int(3)       DEFAULT '0',
    `status`         smallint(1)  DEFAULT '1' COMMENT '1: buka 0: tutup 2: hapus',
    `luas`           varchar(5)   DEFAULT '0' COMMENT 'meter persegi',
    `kondisi`        varchar(50)  DEFAULT null,
    `jumlah`         int(5)       DEFAULT null
);

CREATE TABLE `tahun_ajaran`
(
    `id`               bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `kode`             int(10)                NOT NULL,
    `tahun_akhir`      varchar(255)           NOT NULL,
    `tahun_awal`       varchar(255)           NOT NULL,
    `jns_smt`          int(1)                 NOT NULL COMMENT '1 = reg ganjil, 2 = reg genap, 3 = sp ganjil, 4 = sp genap',
    `set_aktif`        tinyint(1)             NOT NULL,
    `biku_tagih_jenis` tinyint(1)  DEFAULT '0' COMMENT '1 = spp; 2 = sks; 3 = kekurangan',
    `update_time`      datetime    DEFAULT null,
    `update_id`        varchar(18) DEFAULT null,
    `update_host`      varchar(18) DEFAULT null,
    `added_time`       datetime    DEFAULT null,
    `added_id`         varchar(18) DEFAULT null,
    `added_host`       varchar(18) DEFAULT null,
    `tgl_masuk`        date        DEFAULT null
);

CREATE TABLE `daftar_nilai`
(
    `_id`       int(10) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `nim_dinus` varchar(50) DEFAULT null,
    `kur_id`         int(11),
    `kdmk`      varchar(20) DEFAULT null,
    `nl`        char(2)     DEFAULT null,
    `hide`      smallint(1) DEFAULT '0' COMMENT '0 = nilai muncul;\r\n1 = nilai disembunyikan (utk keperluan spt hapus mata kuliah)'
);

CREATE TABLE `validasi_krs_mhs`
(
    `id`        int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `nim_dinus` varchar(50)         NOT NULL,
    `job_date`  datetime                     DEFAULT null,
    `job_host`  varchar(255)                 DEFAULT null,
    `job_agent` varchar(255)                 DEFAULT null,
    `ta`        int(10)             NOT NULL DEFAULT '0'
);

CREATE UNIQUE INDEX `nim` ON `tagihan_mhs` (`nim_dinus`, `ta`);

CREATE UNIQUE INDEX `nim_dinus` ON `mahasiswa_dinus` (`nim_dinus`);

CREATE INDEX `STS_AKD_MHS` ON `mahasiswa_dinus` (`akdm_stat`);

CREATE UNIQUE INDEX `nim` ON `ip_semester` (`ta`, `nim_dinus`);

CREATE UNIQUE INDEX `jam_unik` ON `sesi_kuliah` (`jam_mulai`, `jam_selesai`);

CREATE INDEX `FK_sesi_kuliah_bentrok2` ON `sesi_kuliah_bentrok` (`id_bentrok`);

CREATE INDEX `PERIODE` ON `krs_record` (`ta`);

CREATE INDEX `MAHASISWA` ON `krs_record` (`nim_dinus`);

CREATE UNIQUE INDEX `nim` ON `mhs_ijin_krs` (`ta`, `nim_dinus`);

CREATE UNIQUE INDEX `kode` ON `tahun_ajaran` (`kode`);

CREATE INDEX `nim` ON `daftar_nilai` (`nim_dinus`, `kdmk`);

ALTER TABLE `tagihan_mhs`
    ADD FOREIGN KEY (`ta`) REFERENCES `tahun_ajaran` (`kode`);

ALTER TABLE `tagihan_mhs`
    ADD FOREIGN KEY (`nim_dinus`) REFERENCES `mahasiswa_dinus` (`nim_dinus`);

ALTER TABLE `ip_semester`
    ADD FOREIGN KEY (`ta`) REFERENCES `tahun_ajaran` (`kode`);

ALTER TABLE `ip_semester`
    ADD FOREIGN KEY (`nim_dinus`) REFERENCES `mahasiswa_dinus` (`nim_dinus`);

ALTER TABLE `krs_record`
    ADD FOREIGN KEY (`ta`) REFERENCES `tahun_ajaran` (`kode`);
#
ALTER TABLE `krs_record`
    ADD FOREIGN KEY (`kur_id`, `kdmk`) REFERENCES `matkul_kurikulum` (`kur_id`, `kdmk`);

ALTER TABLE `krs_record`
    ADD FOREIGN KEY (`id_jadwal`) REFERENCES `jadwal_tawar` (`id`);

ALTER TABLE `krs_record`
    ADD FOREIGN KEY (`nim_dinus`) REFERENCES `mahasiswa_dinus` (`nim_dinus`);

ALTER TABLE `krs_record_log`
    ADD FOREIGN KEY (`id_krs`) REFERENCES `krs_record` (`id`); #donn harusnya

ALTER TABLE `krs_record_log`
    ADD FOREIGN KEY (`nim_dinus`) REFERENCES `mahasiswa_dinus` (`nim_dinus`);

ALTER TABLE `krs_record_log`
    ADD FOREIGN KEY (`kur_id`, `kdmk`) REFERENCES `matkul_kurikulum` (`kur_id`, `kdmk`);

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`kur_id`, `kdmk`) REFERENCES `matkul_kurikulum` (`kur_id`, `kdmk`);

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`id_hari1`) REFERENCES `hari` (`id`);

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`id_hari2`) REFERENCES `hari` (`id`); # solve harusnya soalnya sudah saya ganti hari id jadi unique

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`id_hari3`) REFERENCES `hari` (`id`);

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`id_sesi1`) REFERENCES `sesi_kuliah` (`id`); #

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`id_sesi2`) REFERENCES `sesi_kuliah` (`id`);

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`id_sesi3`) REFERENCES `sesi_kuliah` (`id`);

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`id_ruang1`) REFERENCES `ruang` (`id`);

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`id_ruang2`) REFERENCES `ruang` (`id`);

ALTER TABLE `jadwal_tawar`
    ADD FOREIGN KEY (`id_ruang3`) REFERENCES `ruang` (`id`);

ALTER TABLE `mhs_ijin_krs`
    ADD FOREIGN KEY (`ta`) REFERENCES `tahun_ajaran` (`kode`);

ALTER TABLE `mhs_ijin_krs`
    ADD FOREIGN KEY (`nim_dinus`) REFERENCES `mahasiswa_dinus` (`nim_dinus`);

ALTER TABLE `herregist_mahasiswa`
    ADD FOREIGN KEY (`nim_dinus`) REFERENCES `mahasiswa_dinus` (`nim_dinus`);

ALTER TABLE `herregist_mahasiswa`
    ADD FOREIGN KEY (`ta`) REFERENCES `tahun_ajaran` (`kode`);

ALTER TABLE `mhs_dipaketkan`
    ADD FOREIGN KEY (`nim_dinus`) REFERENCES `mahasiswa_dinus` (`nim_dinus`);

ALTER TABLE `sesi_kuliah_bentrok`
    ADD CONSTRAINT `FK_sesi_kuliah_bentrok` FOREIGN KEY (`id`) REFERENCES `sesi_kuliah` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `sesi_kuliah_bentrok`
    ADD CONSTRAINT `FK_sesi_kuliah_bentrok2` FOREIGN KEY (`id_bentrok`) REFERENCES `sesi_kuliah` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

# ALTER TABLE `tahun_ajaran`
#     ADD FOREIGN KEY (`kode`) REFERENCES `tahun_ajaran` (`id`);

ALTER TABLE `daftar_nilai`
    ADD FOREIGN KEY (`nim_dinus`) REFERENCES `mahasiswa_dinus` (`nim_dinus`);

ALTER TABLE `daftar_nilai`
    ADD FOREIGN KEY (`kur_id`, `kdmk`) REFERENCES `matkul_kurikulum` (`kur_id`, `kdmk`);

ALTER TABLE `validasi_krs_mhs`
    ADD FOREIGN KEY (`nim_dinus`) REFERENCES `mahasiswa_dinus` (`nim_dinus`);

ALTER TABLE `validasi_krs_mhs`
    ADD FOREIGN KEY (`ta`) REFERENCES `tahun_ajaran` (`kode`);


#MENGATASI DATA SESI DAN RUANG YG DI DB 0 BIAR BISA DI BUAT NULL
ALTER TABLE `jadwal_tawar`
    MODIFY `id_sesi1` INT NULL,
    MODIFY `id_sesi2` INT NULL,
    MODIFY `id_sesi3` INT NULL,
    MODIFY `id_ruang1` INT NULL,
    MODIFY `id_ruang2` INT NULL,
    MODIFY `id_ruang3` INT NULL;

ALTER TABLE `matkul_kurikulum`
    MODIFY `kur_aktif` TINYINT(1);

ALTER TABLE krs_record_log RENAME COLUMN lastUpdate TO last_update;
