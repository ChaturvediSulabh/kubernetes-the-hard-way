#¡/bin/sh
cd /tmp
wget https://apt.puppetlabs.com/puppet5-release-bionic.deb
sudo dpkg -i puppet5-release-bionic.deb
sudo apt update
sudo apt-get install puppetserver -y
sudo cat > /etc/puppetlabs/puppet/puppet.conf << EOF
[main]
certname = puppet-master

[master]
certname = puppet-master
vardir = /opt/puppetlabs/server/data/puppetserver
logdir = /var/log/puppetlabs/puppetserver
rundir = /var/run/puppetlabs/puppetserver
pidfile = /var/run/puppetlabs/puppetserver/puppetserver.pid
codedir = /etc/puppetlabs/code
EOF
sudo sed -i 's/-Xms[0-9]g/-Xms512m/g' /etc/default/puppetserver
sudo sed -i 's/-Xmx[0-9]g/-Xmx512m/g' /etc/default/puppetserver
sudo /opt/puppetlabs/bin/puppetserver ca setup
sudo systemctl start puppetserver
sudo systemctl enable puppetserver
sudo systemctl status puppetserver
sudo /opt/puppetlabs/bin/puppetserver ca list