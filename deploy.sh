aws_user="user"
aws_address="127.0.0.1"
db_address="127.0.0.1"
db_user="user"
db_pass="pass"

echo "UPLOAD"
scp -i key.pem /zus.zip $aws_user@$aws_address:~/


echo "INSTALL GO"
ssh -i key.pem $aws_user@$aws_address
cd ~/
sudo snap install --classic go
mkdir ~/go
echo "export GOPATH=\$HOME/go" >> ~/.bash_profile
source ~/.bash_profile


echo "UNPACK"
rm -rf zus || true
mkdir -p zus
unzip zus*.zip -d zus
rm zus*.zip
mkdir -p ~/go/src/github.com/ildomm/zus
cp -R zus/ ~/go/src/github.com/ildomm/zus


echo "BUILD"
export GOROOT=/usr/local/go
export GOPATH=~/go/
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
cd ~/go/src/github.com/ildomm/zus
go -o zus build cmd/zus-server/main.go


echo "MIGRATE DB"
wget https://github.com/golang-migrate/migrate/releases/download/v4.4.0/migrate.linux-amd64.tar.gz
tar -xzf migrate.linux-amd64.tar.gz
migrate.linux-amd64 -source file://db/migrations -database mysql://$db_user:$db_pass@tcp($db_address:3306)/zus goto 0202


echo "START"
sudo start-stop-daemon -b -m --start --name "zus" --pidfile "zus.pid" --exec "zus" --chdir "/home/ubuntu/go/src/github.com/ildomm/zus"