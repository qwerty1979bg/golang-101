Vagrant.configure("2") do |config|

  config.vm.provider "virtualbox" do |vb|
    vb.memory = "256"
  end

#config.ssh.pty = true

  config.vm.define name="dev1" do |node|
#    node.vm.box = "qwerty1979/mysql"
    node.vm.box = "ubuntu/xenial64"
    node.vm.hostname = "go101"
    node.vm.provision "shell", path: "provision_go.sh"
#    node.vm.provision "shell", path: "provision_mysql.sh"
#    node.vm.provision "shell", path: "provision_vault.sh"
  end

end
