// package main

// import "database/sql"

// // Select 'ohyou_api' database
// USE ohyou_api;

// // Reset
// DROP TABLE IF EXISTS transactions;

// CREATE TABLE transactions (
//     `id` INTEGER PRIMARY KEY,
//     `description` VARCHAR(140),
//     `purchaser`   INTEGER,
//     `amount`      INTEGER
// );

// INSERT INTO transactions
// VALUES (
//     1,
//     "Costco Eggs",
//     1,
//     140
// );

// // Select 'transactions' table
// SELECT * FROM transactions;