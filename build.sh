# https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
GOOS=linux
GOARCH=amd64
GOOS=$GOOS GOARCH=$GOARCH go build -o go-backend-$GOOS-$GOARCH