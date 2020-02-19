# Simple Deploy

Program ini di buat untuk mempelajari tingkat awal membangun sebuah simple API dengan menggunakan bahasa pemrograma golang atau go language. Dengan adanya contoh ini semoga kita dapat lebih mengerti, paham, dan tahu bagaimana membangun API menggunakan bahasa pemrograman **go**, dimulai dari hello world!!!

## Instalasi

Silahkan kunjungi link di bawah ini untuk menginstall docker

`https://docs.docker.com/install/linux/docker-ce/ubuntu/`

Silahkan kunjungi link di bawah ini untuk menginstall docker-compose

`https://docs.docker.com/compose/install/`

Silahkan kunjungi link di bawah ini untuk menginstall go language

`https://medium.com/better-programming/install-go-1-11-on-ubuntu-18-04-16-04-lts-8c098c503c5f`

## Run

### Docker

Untuk menjalankan program menggunakan docker kita pertama-tama harus build image nya terlebih dahulu menggunakan command dibawah ini

`docker build -t simple_server .`

selanjutnya kita bisa menggunakan docker run untuk menjalankannya

`docker run -p 1234:1234 simple_server`

atau apabila teman-teman sudah mempunyai docker-compose bisa menggunakan command di bawah ini.

`docker-compose up -d`

### go

sebelum kita run program ini kita bisa melakukan test program dulu dengan menggunakan command dibawah ini

`go test`

untuk menambahkan packages yang di perlukan kita menggunakan command **go get -v** dan lanju kita jalan kan program ini dengan menggunakan command 

`go run server.go`

## Testing

Coming Soon
