install:
	cd rails-api && rails db:setup
	cd ..
rails-server:
	cd rails-api && rails server
go-server:
	cd go-api && go run main.go
roda-server:
	cd roda-api && bundle exec puma
on-rails:
	sh bench.sh $(R) http://localhost:3000/messages
on-go:
	sh bench.sh $(R) http://localhost:8080/messages
on-roda:
	sh bench.sh $(R) http://localhost:9292/messages