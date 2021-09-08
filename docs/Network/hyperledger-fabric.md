# Hyperledger Fabric

## Install

```zsh
export http_proxy=localhost:8889
export https_proxy=localhost:8889
sudo export http_proxy=localhost:8889
sudo export https_proxy=localhost:8889
sudo mkdir /src/github.com/hyperledger/fabric
cd /src/github.com/hyperledger/fabric
sudo git clone https://github.com/hyperledger/fabric.git
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin/orderer
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin/configtxgen
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin/configtxlator
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin/cryptogen
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin/discover
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin/idemixgen
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin/ledgerutil
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin/osnadmin
sudo mkdir -p /src/github.com/hyperledger/fabric/release/linux-amd64/bin/peer
yay -S hyperledger-fabric
```

