git add .
git commit -m "Commit"
git push

set GOOS=linux
set GOARCH=amd64

go build -tags lambda.norpc -o bootstrap main.go
# rm main.zip
# zip main.zip main

# zip myFunction.zip bootstrap


# chmod +x deploy.sh
# ./deploy.sh