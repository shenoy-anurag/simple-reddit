# Backend

The backend of this project is written in Go.

## Contributors:
- Anurag Shenoy
- Ganesan Santhanam

## Overview:

### Design Pattern: 
We will try to adhere to Domain Driven Design (DDD) pattern.
See more here: <https://www.citerus.se/go-ddd/>

### Pre-requisites:
1. Go version `1.17.x` (`1.17.6` is used for this project)
2. MongoDB version `5.x.x` (`5.0.4` to be exact)
3. MongoDB Database Tools version `100.5.1`

#### Installing Go:
Install version `1.17.6` from <https://go.dev/> using the steps given here: <https://go.dev/doc/install>.

It is recommended to use a package manager such as `homebrew` (for MacOS <https://brew.sh/>) or the relevant package manager for your OS.

For MacOS: If using brew, run the command `brew install go@1.17` to install version 1.17.x. More on how to install Go using brew here: <https://formulae.brew.sh/formula/go>.

#### Installing MongoDB:
Install version `5.0.4` of MongoDB using the steps given here: <https://docs.mongodb.com/manual/installation/> or you can download the correct file from <https://www.mongodb.com/download-center/community/releases/archive> if the version is an older one.

#### Installing MongoDB Database Tools:
Install the mongo database tools from <https://docs.mongodb.com/database-tools/installation/installation/>. This project uses version `100.5.1` of MongoDB Database Tools.

### Setting up MongoDB and the Database:
1. After installing MongoDB, run the `mongod` daemon/service using the instructions from here <https://docs.mongodb.com/manual/installation/>.
2. De-compress the db from the `db.zip` file provided in the repo.
3. Restore the db using the `mongorestore` command like so `mongorestore db/`. An example using options: `mongorestore --host <host> --port <port number> <path to the backup>` and replace the values within `<>` appropriately. For eg. `mongorestore --host 127.0.0.1 --port 27017 /Path/to/database/folder`.

More information on `mongorestore` command can be found here: <https://docs.mongodb.com/database-tools/mongorestore/>.

And that's it! You have setup MongoDB service and restored the database.

### Running Server:
`main.go` file (within backend folder) contains the code that will start the server. The server can be started using these commands: 
1. `cd backend/`
2. `go run main.go`.

Now, you can follow the steps to get the frontend up and running and check out the web-app.