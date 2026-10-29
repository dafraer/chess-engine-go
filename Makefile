run:
	go run cmd/main.go
push:
	git add . && git commit -m "$(m)" && git push
