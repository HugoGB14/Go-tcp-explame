go build .
GOOS=darwin go build .
mv Client Client-Darwin
GOOS=linux go build .
mv Client Client-Linux-GNU