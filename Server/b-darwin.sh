GOOS=windows go build .
go build .
mv Server Server-Darwin
GOOS=linux go build .
mv Server Server-Linux-GNU