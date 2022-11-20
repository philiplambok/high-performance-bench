# cat config.ru
require "roda"
require 'sequel'
require 'debug'

Sequel.single_threaded = true
DB = Sequel.connect('mysql2://localhost/golang_test?user=root&password=', max_connections: 10)

class App < Roda
  plugin :json

  route do |r|
    r.is('messages') do
      r.post do
        DB[:messages].insert(message: 'Hello posting from Roda')
        {message: 'success'}
      end
    end

    r.get('ping') do
      { message: 'pong' }
    end
  end
end

run App