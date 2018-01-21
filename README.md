Ohyou API
===

Ohyou API is the back-end to Ohyou. Ohyou (temp) is an app to help you manage debts between friends. Existing solutions felt like they subject to feature creep, and like they were not optimized for my particular use-case.

## Getting Started

### Clone & Install Dependencies

```bash
# clone repo
git@gitlab.com:aflashyrhetoric/ohyou-api.git

# Install dependencies
go get
```

#### Ensure MySQL settings are as follows:

|||
|---|---|
|version|`mysql  Ver 15.1`|
|database name|`ohyou_api`|
|host|`localhost` (or `tcp(localhost)`)|
|user|`root`|
|password|`password`|

<sup>If you haven't done so already and it is safe, run `mysql_secure_installation` to get an interactive prompt where you'll be able to set the `root` user and password. </sup>

#### Connect to MySQL.

<!-- TODO: Create bash script for initialization of MySQL -->

```bash
# ensure mysql is running
mysqld

# connect to mysql
mysql -uroot --password="password"
```

#### Configure Database.

```SQL
-- Type the following when connected to mysql

-- create database
CREATE DATABASE ohyou_api;

-- create transaction table
CREATE TABLE transactions (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(100) NOT NULL,    
    purchaser INTEGER NOT NULL);
);

-- create transactions__beneficiaries table
CREATE TABLE transactions__beneficiaries (
    transaction_id INTEGER NOT NULL,
    beneficiary_id INTEGER NOT NULL);
```

## Run

In project directory: `go run *.go`. 

For now, `main.go` will connect to your database via the hardcoded parameter string: 

<!-- TODO: Retrieve database connections from an environment file -->

```golang
db, err = sql.Open("mysql", "root:password@tcp(localhost)/ohyou_api")
```

The app should now be running the API on localhost on port 8080. :smile_cat:

## Testing & Development Workflow

I use [Insomnia](https://insomnia.rest) for testing our API. It's a beautifully designed, intuitive app that's easy to get started with. 

I use [Visual Studio Code](https://code.visualstudio.com) with Go language support.

#### Seeding

For now, issue a **GET** request to the following endpoint to seed the database: 

`http://localhost:8080/api/v1/seed`

<!-- TODO: Create a bash file (or some other solution) to seed the database instead of using an endpoint -->

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

