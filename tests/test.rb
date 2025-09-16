require 'net/http'
require 'uri'
require 'json'

uri = URI('http://localhost:3000/add_message')
body = JSON.generate({ message: 'Message', user_id: 1, chat_id: 100 })

10.times do
  res = Net::HTTP.post(uri, body, "Content-Type" => "application/json")
  pp res
end
