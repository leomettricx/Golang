#!/bin/bash
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Usage:
#
# ./prereqs-ubuntu.sh
#
# User must then logout and login upon completion of script
#

# Exit on any failure
set -e

# Array of supported versions
declare -a versions=('trusty' 'xenial' 'yakkety', 'bionic');

# check the version and extract codename of ubuntu if release codename not provided by user
if [ -z "$1" ]; then
    source /etc/lsb-release || \
        (echo "Error: Release information not found, run script passing Ubuntu version codename as a parameter"; exit 1)
    CODENAME=${DISTRIB_CODENAME}
else
    CODENAME=${1}
fi

# check version is supported
if echo ${versions[@]} | grep -q -w ${CODENAME}; then
    echo "Installing Hyperledger Composer prereqs for Ubuntu ${CODENAME}"
else
    echo "Error: Ubuntu ${CODENAME} is not supported"
    exit 1
fi

# Install curl
sudo apt-get install curl

# Mudar o local de onde se atualiza upgrade
software-properties-gtk

# VS Code
curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg
sudo install -o root -g root -m 644 microsoft.gpg /etc/apt/trusted.gpg.d/
sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list'


# Update package lists
echo "# Updating package lists"
sudo apt-add-repository -y ppa:git-core/ppa
sudo apt-get update

# Install Git
echo "# Installing Git"
sudo apt-get install -y git

# Install Go language
wget https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.12.6.linux-amd64.tar.gz

#GO VARIABLES
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH:

# Editar ./profile

sudo groupadd docker
sudo usermod -aG docker $USER

#END GO VARIABLES

restart

# Install nvm dependencies
echo "# Installing nvm dependencies"
sudo apt-get -y install build-essential libssl-dev

# Execute nvm installation script
echo "# Executing nvm installation script"
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.2/install.sh | bash



# Set up nvm environment without restarting the shell
export NVM_DIR="${HOME}/.nvm"
[ -s "${NVM_DIR}/nvm.sh" ] && . "${NVM_DIR}/nvm.sh"
[ -s "${NVM_DIR}/bash_completion" ] && . "${NVM_DIR}/bash_completion"

# Install node
echo "# Installing nodeJS"
nvm install 8
nvm use 8
node --version



# Ensure that CA certificates are installed
sudo apt-get -y install apt-transport-https ca-certificates

sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
	
# Add Docker repository key to APT keychain
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# Update where APT will search for Docker Packages
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"


# Update package lists
sudo apt-get update



# Install kernel packages which allows us to use aufs storage driver if V14 (trusty/utopic)
if [ "${CODENAME}" == "trusty" ]; then
    echo "# Installing required kernel packages"
    sudo apt-get -y install linux-image-extra-$(uname -r) linux-image-extra-virtual
fi

# Install Docker
echo "# Installing Docker"
sudo apt-get -y install docker-ce // Não funconu
sudo snap install docker


# Verifies APT is pulling from the correct Repository
sudo apt-cache policy docker-ce

# Add user account to the docker group
sudo usermod -aG docker $(whoami)

# Install docker compose
echo "# Installing Docker-Compose"
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Install python v2 if required
set +e
COUNT="$(python -V 2>&1 | grep -c 2.)"
if [ ${COUNT} -ne 1 ]
then
   sudo apt-get install -y python-minimal
fi

# Install unzip, required to install hyperledger fabric.
sudo apt-get -y install unzip

# Print installation details for user
echo ''
echo 'Installation completed, versions installed are:'
echo ''
echo -n 'Node:           '
node --version
echo -n 'npm:            '
npm --version
echo -n 'Docker:         '
docker --version
echo -n 'Docker Compose: '
docker-compose --version
echo -n 'Python:         '
python -V
echo -n 'Go Language'
go version
go env

# Print reminder of need to logout in order for these changes to take effect!
echo ''
echo "Please logout then login before continuing."


# Fabric-Samples
curl -sSL http://bit.ly/2ysbOFE | bash -s


# Trade-Finance-Logistics
cd $GOPATH/src
git clone https://github.com/leomettricx/trade-finance-logistics



docker info
docker version
docker images
docker ps

# colocar /bin no PATH e copiar /fabric-samples/bin
#ir para dir Networks e rodar

# Criar primeiro o diretório crypto-config

cryptogen generate --config=./crypto-config.yaml




echo "######################################################"
echo "#######    Generating Orderer Genesis block    #######"
echo "######################################################"

# No arquivo CONFIGTX.YAML no paragrafo Profiles: tem este label -> MTTXOrdererGenesis:

configtxgen -profile MTTXOrdererGenesis -outputBlock ./channel-artifacts/genesis.block

export COMPOSE_PROJECT_NAME=Mettricx-Networks
export IMAGE_TAG=latest

docker-compose -f docker-compose-e2e.yaml up

docker-compose -f docker-compose-e2e.yaml down



