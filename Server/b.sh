GOOS=windows go build .
GOOS=darwin go build .
mv Server Server-Darwin
go build .
mv Server Server-Linux-GNU