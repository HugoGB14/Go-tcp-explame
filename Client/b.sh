GOOS=windows go build .
GOOS=darwin go build .
mv Client Client-Darwin
go build .
mv Client Client-Linux-GNU