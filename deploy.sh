git add .
git commit -m "Commit"
git push

set GOOS=linux
set GOARCH=amd64

go build -o bootstrap -tags lambda.norpc main.go
rm bootstrap.zip
zip bootstrap.zip bootstrap

# chmod +x deploy.sh
# ./deploy.sh