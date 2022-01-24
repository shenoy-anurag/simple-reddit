# Setup Database
homebrew = Homebrew/homebrew-core
check_brew := $(shell brew -v)

setup-mongo-macos:
	@echo "Installing XCode Command Line Utilities..."
	xcode-select --install
ifneq (,$(findstring $(homebrew), $(check_brew)))
	@echo "Found Homebrew, skipping install..."
else
	/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
endif
	@echo "Installing MongoDB..."
	brew tap mongodb/brew
	brew install mongodb-community@5.0
	@echo "Installing MongoDB Database Tools..."
	brew install mongodb-database-tools
	@echo "Now we will run the mongod daemon using 'mongod --dbpath Path/to/data' command, and you might be required to enter your sudo password"
	mkdir data
	mongod --dbpath data
	@echo "Extracting and restoring database..."
	unzip db.zip
	mongorestore db 
	@echo "All Done!"

setup-mongo-ubuntu:
# This setup is for Ubuntu 20.04 Focal. You can check your distro using 'lsb_release -dc' command.
	@echo "Installing MongoDB..."
# Get mongodb pgp key and create mongodb list
	wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -
	echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/5.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list
	sudo apt-get update
	sudo apt-get install -y mongodb-org
	@echo "Starting mongod service..."
	sudo systemctl daemon-reload
	sudo systemctl start mongod
	@echo "Installing MongoDB Database Tools..."
	wget -O mongodb-database-tools-ubuntu2004-x86_64-100.5.1.tgz https://fastdl.mongodb.org/tools/db/mongodb-database-tools-ubuntu2004-x86_64-100.5.1.tgz
	tar -zxvf mongodb-database-tools-ubuntu2004-x86_64-100.5.1.tgz
	cd mongodb-database-tools-ubuntu2004-x86_64-100.5.1.tgz; sudo cp -r * /usr/local/bin/
	@echo "Installing unzip to extract db, you might need to enter your sudo password..."
	sudo apt install unzip
	@echo "Extracting and restoring database..."
	unzip db.zip
	mongorestore db
	@echo "All Done!"


# Build/Serve Back-end
serve-backend:
	# cd backend; $(MAKE) serve-backend
	$(MAKE) -C backend serve-backend

# Build/Serve Front-end
setup-frontend:
	cd frontend/forum-app; npm install; npm install -g @angular/cli

serve-frontend:
	cd frontend/forum-app; ng serve
