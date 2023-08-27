git add .
git commit -m "Commit"
git push

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0


go build -o Handler main.go



rm Handler.zip
zip Handler.zip Handler

# chmod +x deploy.sh
# ./deploy.sh
