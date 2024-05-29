# CandyCapade
 
### О проекте

CandyCapade - это учебный проект по языку Go. Включает в себя задачи по созданию сервера для обработки заказов конфет, реализацию аутентификации по сертификатам TLS, а также интеграцию с программой, создающей ASCII-изображения животных.

### Как запустить 


```bash
./go run cmd/server/main.go
./candy-client -k AA -c 2 -m 50 --cert cert/client/cert.pem --key cert/client/key.pem --cacert cert/minica.pem
```
