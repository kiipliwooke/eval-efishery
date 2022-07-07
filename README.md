# eval-efishery
Repository ini dibuat untuk memenuhi evaluasi eFishery Aqua-Developer Batch 1
Algoritma yang dipilih yaitu Palindrome dengan source code pada algorithm/
ERD berada pada docs/
Terdapat daftar test menggunakan curl pada docs/

## Asumsi
* status verifikasi dokumen telah dilakukan pada proses sebelumnya
* akun telah dibuat pada proses sebelumnya
* value loan telah ditetapkan pada proses sebelumnya
* status loan telah ditetapkan pada proses sebelumnya

## Daftar Fitur
* User dapat login (1 akun dalam satu waktu)
* User dapat logout (jika sebelumnya telah melakukan login)
* Aplikasi dapat membedakan user biasa dan admin ketika seseorang login
* User biasa dapat melihat seluruh daftar loan miliknya
* User biasa dapat melihat detil loan miliknya
* Admin dapat melihat seluruh daftar loan
* Admin dapat melihat detil seluruh loan
* User biasa dapat mencantumkan dokumen baru pada loan miliknya
* Aplikasi dapat mengubah status loan menjadi review setelah user menambahkan dokumen
* Admin dapat menghapus loan
* User biasa dapat menambahkan (mengupdate) approval setelah status loan approved
* Admin dapat menambahkan (mengupdate) approval setelah user meng-approve loan

## Daftar Bug yang terdetect
* foreignKey
* handling file nama sama
* backtracking (mereverse) ketika ada proses yang gagal
* besar file tidak ada batas
* cascade saat delete (sekarang manual)
* tidak ada hashing password