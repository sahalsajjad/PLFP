if [ ! -d "Sources" ]; then
	mkdir Sources && cd Sources  
else
	cd Sources
fi

wget https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
GOROOT=$HOME/Sources/go
export PATH=$PATH:$GOROOT/bin
sudo apt-get update
sudo apt-get install gcc-go
sudo apt-get install erlang erlang-doc

