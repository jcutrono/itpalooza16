# HELLO ITPalooza 2016

----
## Lets get started - sign into AWS
![Dashboard](https://github.com/jcutrono/itpalooza16/blob/master/pics/dashboard.png)

## Choose an OS
![OS](https://github.com/jcutrono/itpalooza16/blob/master/pics/os.png)

## Select an instance size
![Instance](https://github.com/jcutrono/itpalooza16/blob/master/pics/instance.png)

## Set the instance details
![Instance Details](https://github.com/jcutrono/itpalooza16/blob/master/pics/instance_details.png)

## Set any tags
![Tags](https://github.com/jcutrono/itpalooza16/blob/master/pics/tags.png)

## Choose ports for the firewall
![Firewall](https://github.com/jcutrono/itpalooza16/blob/master/pics/firewall.png)

## Launch the instance
![Launch](https://github.com/jcutrono/itpalooza16/blob/master/pics/launch.png)

## Download pem
![Pem](https://github.com/jcutrono/itpalooza16/blob/master/pics/pem.png)

## Connect to the instance
![Connect](https://github.com/jcutrono/itpalooza16/blob/master/pics/connect.png)

## Copy ssh string
![ssh](https://github.com/jcutrono/itpalooza16/blob/master/pics/ssh.png)

## Run the following from the command line to download/setup dokku
    wget https://raw.githubusercontent.com/dokku/dokku/v0.7.2/bootstrap.sh;
   sudo DOKKU_TAG=v0.7.2 bash bootstrap.sh

## (Errors around permission denied)
*coming soon*

## Exit out back to your repo
    git remote add dokku dokku@{instance}:{app}

## DEPLOY
    git push dokku master

## Install Mongo
    sudo dokku plugin:install https://github.com/dokku/dokku-mongo.git mongo

## Create database
    dokku mongo:create {dbname}
