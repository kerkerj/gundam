#!/usr/bin/env ruby

while(true) do 
  r = rand(255)
  g = rand(255)
  b = rand(255)
  puts "(r, g, b) = (#{r}, #{g}, #{b})"
  `curl -X PUT localhost:3000/rgb/#{r},#{g},#{b}`
  sleep 1
end
