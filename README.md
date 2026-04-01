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
2xx -> 95
4xx -> 20
5xx -> 5

Top IPs:
127.0.0.1 -> 30
10.0.0.5 -> 18

Top paths:
/login -> 22
/api/users -> 15
/favicon.ico -> 10
```



Features:

1. Show number of lines in log [X]
2. Show top 3 messages (errors) that we have and count of appears [X]
3. Show top IPs and count of appears in log []
4. Show top paths and count of appears []
