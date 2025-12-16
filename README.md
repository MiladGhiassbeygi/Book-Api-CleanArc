# Book-Api-CleanArc

یک API ساده مدیریت کتاب و نویسنده با Go، بر اساس **Clean Architecture** و الگوی **CQRS** با مستندسازی **Swagger**.

---

## ساختار پروژه

C:.
│ go.mod
│ go.sum
│
├───cmd
│ └───server # نقطه ورود برنامه
├───docs # مستندات Swagger
├───infrastructure # پیاده‌سازی پایگاه داده و Repository
├───interfaces # Handlerهای HTTP و gRPC
└───internal # لایه داخلی: Domain, App, Command, Query, Repository


---

## نیازمندی‌ها

- Go 1.21+
- PostgreSQL
- Gin
- GORM
- Swaggo

---

## نصب و اجرای پروژه

1. کلون کردن پروژه:
```bash
git clone <your-repo-url>
cd Book-Api-CleanArc

    نصب وابستگی‌ها:

go mod tidy

    تنظیم اتصال به پایگاه داده در database/database.go:

dsn := "host=localhost user=postgres password=123456 dbname=bookdb port=5432 sslmode=disable"

    اجرای مهاجرت‌ها:

go run cmd/server/main.go

    دسترسی به Swagger:

http://localhost:8080/swagger/index.html

    اجرای سرور:

go run cmd/server/main.go

API
Authors

    GET /authors : دریافت همه نویسندگان

    GET /authors/:id : دریافت نویسنده با ID

    POST /authors : ایجاد نویسنده جدید

    DELETE /authors/:id : حذف نویسنده

Books

    GET /books : دریافت همه کتاب‌ها

    GET /books/:id : دریافت کتاب با ID

    GET /books/author/:authorID : دریافت کتاب‌های یک نویسنده

    POST /books : ایجاد کتاب جدید

ویژگی‌ها

    Clean Architecture

    CQRS (Command & Query)

    Repository Pattern

    مستندسازی با Swagger

    مدیریت روابط Author ↔ Book
