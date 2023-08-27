git add .
git commit -m "Commit"
git push

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

go build -o main main.go
rm main.zip
zip main.zip main

# chmod +x deploy.sh
# ./deploy.sh