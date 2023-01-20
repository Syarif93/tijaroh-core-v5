# Tijaroh Core v5

Aplikasi ini berjalan di go versi 1.19^

**Development**

Untuk menjalankan development mode dengan menjalankan file dev.sh pada root folder project. Pastikan anda sudah menginstall nodejs + npm lalu install nodemon untuk merestart server setiap ada perubahan pada file .go.

    sudo npm install -g nodemon

Jalankan perintah berikut untuk menginstall package yang dibutuhkan :

    go mod download

Jalankan file dev.sh

    chmod +x ./dev.sh
    ./dev.sh
