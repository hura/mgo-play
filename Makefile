# vim: ts=4 sta ai si noci 
#
#
# All version at: http://dl.mongodb.org/dl/linux/x86_64
#MGO_BIN := http://fastdl.mongodb.org/linux/mongodb-linux-x86_64-2.4.6.tgz
#MGO_BIN := http://fastdl.mongodb.org/linux/mongodb-linux-x86_64-latest.tgz
MGO_BIN := http://fastdl.mongodb.org/linux/mongodb-linux-x86_64-v2.4-latest.tgz
MGO_DEST := mongodb
MGO_CONF := $(MGO_DEST)/mongodb.conf

# s/debian/rpm/ for RH/centos (used pid file which debian doesnt)
MGO_CONF_SRC := https://raw.github.com/mongodb/mongo/master/debian/mongodb.conf

def_target: none

install: install-mongodb

install-mongodb :
	mkdir -p $(MGO_DEST)
	curl $(MGO_BIN) | tar xz --strip-components=1 -C $(MGO_DEST)
	mkdir -p $(MGO_DEST)/data
	curl $(MGO_CONF_SRC) > $(MGO_CONF).orig
	awk '/dbpath=/{print"dbpath=$(MGO_DEST)/data";next}{print}' $(MGO_CONF).orig > $(MGO_CONF).orig2
	awk '/logpath=/{print"logpath=$(MGO_DEST)/data/mongo.log";next}{print}' $(MGO_CONF).orig2 > $(MGO_CONF)
	rm $(MGO_CONF).orig2
	echo "fork = True" >> $(MGO_CONF)
	echo "pidfilepath = $(MGO_DEST)/mongo.pid" >> $(MGO_CONF)


start : start-mongo

start-mongo :
	@echo "Please wait until all children are running"
	$(MGO_DEST)/bin/mongod --config $(MGO_CONF)

stop : stop-mongo

stop-mongo :
	# Doesnt seem to work since mongo doesnt create pid file:
	#kill `cat $(MGO_DEST)/mongo.pid`
	kill `pidof mongod`

restart-mongo : stop-mongo start-mongo
restart : restart-mongo

