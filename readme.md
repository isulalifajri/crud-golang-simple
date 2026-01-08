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
	fmt.Println("Hello Golang")
}

```

jalankan: `go run main.go`
output : `Hello Golang`

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

## web server paling sederhana

Sekarang kita bikin server HTTP tanpa database dulu

ganti code `main.go` jadi seperti ini:

```
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Golang API")
	})

	// === Jalankan server ===
	fmt.Println("Server berjalan di http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", router))
}


```

karena kita pake mux, maka jalankan ini:
`go get github.com/gorilla/mux`

jalankan: `go run .`

akses di browser: `http://localhost:2000`
output: `Hello API Golang`

