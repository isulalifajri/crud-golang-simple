## membuat program crud sederhana dengan golang

structur folder nya nanti:
```
golang-crud/
│
├── config/
│   └── database.go
│
├── models/
│   └── user.go
│
├── handlers/
│   └── user_handler.go
│
├── seeders/
│   └── user_seeder.go
│
├── routes/
│   └── routes.go
│
├── main.go
└── go.mod

```

# langkah pertama

jalankan :
1. `mkdir golang-crud-simple` -> buat file dengan terminal
2. `cd golang-crud-simple` -> masuk kedalam filenya
3. `go mod init golang-crud-simple` -> init module

step 3 ini nanti menghasilkan file dengan nama `go.md` di dlm folder `golang-crud-simple`

file go pertama:
buat file `main.go`
```
package main

import "fmt"

func main() {
	fmt.Println("Hello Golang 🚀")
}

```

jalankan: `go run main.go`
output : `Hello Golang 🚀`

## konsep dasar wajib paham

Konsep dasar (wajib paham)

Sedikit teori dulu (nggak bisa diskip):

🔹 package main

Entry point aplikasi

Kayak `public/index.php` di Laravel

🔹 func main()

Fungsi pertama yang dijalankan

🔹 import

Mirip `use / require`

