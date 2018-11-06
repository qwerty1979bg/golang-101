Vagrant.configure("2") do |config|

  config.vm.provider "virtualbox" do |vb|
    vb.memory = "256"
  end

  config.vm.define name="dev1" do |node|
    node.vm.box = "ubuntu/xenial64"
    node.vm.hostname = "go101"
    node.vm.provision "shell", path: "provision_go.sh"
  end

end
