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

# Requirements
- MySQL 15.1 
- Go 1.9.1

<sup>If you haven't done so already and it is safe, run `mysql_secure_installation` to get an interactive prompt where you'll be able to set the `root` user and password. </sup>

#### Configure MySQL

```bash
./provision/mysql-config.sh
```
#### Seed database w/ test data

```bash
cd provision
go run seeder.go
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

