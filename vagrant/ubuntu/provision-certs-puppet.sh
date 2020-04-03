#!/bin/sh
sudo /opt/puppetlabs/bin/puppetserver ca list
sudo /opt/puppetlabs/bin/puppetserver ca sign --certname master-1
sudo /opt/puppetlabs/bin/puppetserver ca sign --certname master-2
sudo /opt/puppetlabs/bin/puppetserver ca sign --certname worker-1
sudo /opt/puppetlabs/bin/puppetserver ca sign --certname worker-2
sudo /opt/puppetlabs/bin/puppetserver ca sign --certname loablancer