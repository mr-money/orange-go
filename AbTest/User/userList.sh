#!/bin/bash
host=http://127.0.0.1
port=8080
path=/api/user/userList
header=Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3LCJhdWQiOiLljrvnjqnlhL8iLCJleHAiOjE2NjU0NzEzNTksImp0aSI6IjEyYTMzMTI0LWJiYTctNGYxYi1iOTQ2LWZjY2Y2NDI3MjFjYiIsImlhdCI6MTY2NTM4NDk1OSwiaXNzIjoiYWRtaW4iLCJuYmYiOjE2NjUzODQ5NTksInN1YiI6ImxvZ2luIn0.cLoXo_sUTdP8vLWCzsM-npPM2YsZXgICBKp71RdZGg4

ab -n 10000 -c 1000 -H $header $host:$port$path