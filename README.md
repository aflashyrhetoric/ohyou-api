Payup API
===

Payup is an app to help you manage debts between friends. 

## Why?

Existing solutions felt like they were subject to feature creep; they also felt like they weren't optimized for my particular use-case.

## Getting Started

To install Payup API locally on your machine, follow the steps below.

### Clone & Install Dependencies

```bash
# clone repo
git@gitlab.com:aflashyrhetoric/payup-api.git

# Install dependencies
go get
dep ensure
```

# Requirements
- MariaDB  `brew install mariadb`
- Go 1.9.1 or later [Download Link](https://golang.org/dl/)

<sup>If you haven't done so already and it is safe, run `mysql_secure_installation` to get an interactive prompt where you'll be able to set settings according to the table below.</sup>

|Prompt begins with...|Value|
|---------------------|-----|
|Enter current password for root|<Enter for none>|
|Remove anonymous users|Y|
|Disallow root login|Y|
|Remove test database (...)|Y|
|Reload privilege tables|Y|


#### Configure MySQL

**Ensure that Docker, or other applications that may affect your ports, is NOT running**

Run the provisioning script, which should populate the database and run the seeder (which adds 25 fake records)

```bash
./provision/mysql-config.sh

# You may need to add execution permissions to script with `chmod +x ./provision/mysql-config.sh`
```
#### Seed database w/ test data

To run the seeder ONLY, or to add more records to the database, run the following:

```bash
cd provision
go run seeder.go
# or
go run provision/seeder.go
```

## Run

In project directory: `go run *.go`. 

For now, `main.go` will connect to your database via the hardcoded parameter string: 

<!-- TODO: Retrieve database connections from an environment file -->

```golang
db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/ohyou_api")
```

The API should now be running the API on localhost on port 8114. :smile_cat:

_NOTE: Visiting `http://localhost:8114` will do nothing! The endpoints are available, but no pages!_

## Testing & Development Workflow

I use [Insomnia](https://insomnia.rest) for testing our API. It's a beautifully designed, intuitive app that's easy to get started with. 

I use [Visual Studio Code](https://code.visualstudio.com) with Go language support.

---
## API Overview

## Parameters
### Transaction

|Property|Type|Content|Default|Example Value|
|--------|--------------|----------|--------|--------| 
|Description|String|A short description of the transaction|N/A|Dozen eggs|
|Purchaser|uint|User ID for user who purchased|N/A|4|
|Amount|uint|Cost of purchase in pennies (in USD for now)|0|500|
|Beneficiaries|[]int|User IDs of users who benefitted from transaction|0|500|

## Endpoints

### Transaction
| HTTP Method | Endpoint          | Method       | 
|:------------|:------------------|:------------------|
| GET         | /transactions/    | listTransaction   |
| POST        | /transactions/    | createTransaction |
| GET         | /transactions/:id | showTransaction   |
| PUT         | /transactions/:id | updateTransaction |
| DELETE      | /transactions/:id | deleteTransaction |

