# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"

  config.vm.synced_folder ENV['GOPATH'], "ENV["HOME"] + "/go"

  config.vm.provider "virtualbox" do |vb|
    vb.memory = "4096"
  end

  config.vm.provision "shell", inline: <<-SHELL
    sudo apt-get remove docker docker-engine docker.io
    sudo apt-get update
    sudo apt-get install -y \
        make \
        git \
        apt-transport-https \
        ca-certificates \
        curl \
        software-properties-common
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
    sudo add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
       $(lsb_release -cs) \
       stable"
    sudo apt-get update
    sudo apt-get install -y docker-ce
    curl https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz > /tmp/golang.tar.gz
    sudo tar -C /usr/local -xzf /tmp/golang.tar.gz
    echo "export PATH=$PATH:/usr/local/go/bin:/home/ubuntu/go/bin" >> /home/ubuntu/.bashrc
    echo "export GOPATH=/home/ubuntu/go" >> /home/ubuntu/.bashrc
    sudo chown ubuntu:ubuntu -R /home/ubuntu/go
    sudo usermod -aG docker ubuntu
  SHELL
end