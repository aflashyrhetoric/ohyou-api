#!/bin/bash

# mysql-config.sh
## configure mysql database
## this file is idempotent and can be repeatedly run safely 

# echo "Starting Mariadb service..."
# brew services start mariadb

# echo "Logging into mysql client..."
#mysql -uroot --password="password"

echo "Create database payup_api..."
mysql --defaults-extra-file=mysql.cnf -e \
  "CREATE DATABASE IF NOT EXISTS payup_api;"
echo "...success!"

echo "Create expenses table..."
mysql --defaults-extra-file=mysql.cnf -e \
  "USE payup_api; CREATE TABLE IF NOT EXISTS expenses (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(200),    
    purchaser INTEGER,
    amount INTEGER,
    receipt_id INTEGER);"
echo "...success!"

echo "Create junction table expenses_beneficiaries..."
mysql --defaults-extra-file=mysql.cnf -e \
  "USE payup_api; CREATE TABLE IF NOT EXISTS expenses_beneficiaries (
    expense_id INTEGER,
    beneficiary_id INTEGER);"

echo "...success!"

echo "Create receipts table..."
mysql --defaults-extra-file=mysql.cnf -e \
  "USE payup_api; CREATE TABLE IF NOT EXISTS receipts (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    merchant VARCHAR(100),    
    total INTEGER);"
echo "...success!"

# Run the seeder
# go run seeder.go