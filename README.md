# nginx-log-parser

Log parser on Go

Run:
```
go run main.go your_file.log
```

Log format:
127.0.0.1 - - [01/Apr/2026:12:00:00 +0000] "GET /login HTTP/1.1" 200 123 "-" "curl/8.0"


Output: 

```
Total lines: 120
Status codes:
200 -> 95
404 -> 20
500 -> 5

Top IPs:
127.0.0.1 -> 30
10.0.0.5 -> 18

Top paths:
/login -> 22
/api/users -> 15
/favicon.ico -> 10
```
