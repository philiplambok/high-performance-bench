install:
	cd rails-api && rails db:setup
	cd ..
rails-server:
	cd rails-api && rails server
go-server:
	cd go-api && go run main.go
roda-server:
	cd roda-api && bundle exec puma