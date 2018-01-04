API
===

## Parameters
### Transaction

|Property|Type|Content|Default|Example Value|
|--------|--------------|----------|--------|--------| 
|Description|String|A short description of the transaction|None|Dozen eggs|
|Purchaser|uint|User ID for user who purchased|None|4|
|Amount|uint|Cost of purchase in pennies|0|500|

## Endpoints

### Transaction
| HTTP Method | Endpoint          | Method       | 
|:------------|:------------------|:------------------|
| POST        | /transactions/    | createTransaction |
| GET         | /transactions/    | listTransaction   |
| GET         | /transactions/:id | showTransaction   |
| PUT         | /transactions/:id | updateTransaction |
| DELETE      | /transactions/:id | deleteTransaction |

createTransaction
