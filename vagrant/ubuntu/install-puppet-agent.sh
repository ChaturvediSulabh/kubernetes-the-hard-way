#¡/bin/sh
cd /tmp
wget https://apt.puppetlabs.com/puppet5-release-bionic.deb
sudo dpkg -i puppet5-release-bionic.deb
sudo apt update
sudo apt-get install puppet-agent -y
sudo systemctl start puppet
sudo systemctl enable puppet
sudo systemctl status puppet
sudo /opt/puppetlabs/bin/puppet agent --fingerprint
