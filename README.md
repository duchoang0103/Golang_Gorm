# Golang_Gorm
#Phần migrate table template
1. go run main.go 
2. dùng công cụ kiểm tra api vd như postman để kiểm tra 
3. tạo bảng templates với api phương thức POST http://localhost:8080/templates như hình 1 "pic_create"
4. kiểm tra db với GET http://localhost:8080/templates như hình 2 "pic_getall"

#Phần Sử dụng temple lấy từ DB cho mail được gửi
5. api denemail: http://localhost:8080/sendmail
6. EX: body post request:
{
    "id": 3,
    "array":
    [
        {
            "link": "tthttps://www.w3schools.com",
            "string": "Visit W3Schools.com!"
        }
    ]
}

7. Database EX: nhu trong pic3

8. ! EX: Bỏ * ở yourDomain và yourAPIKey nếu dùng
