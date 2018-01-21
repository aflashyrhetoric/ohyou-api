# mysql-config.sh
## configure mysql database
## this file is idempotent and can be repeatedly run safely 

echo "Starting MySQL service..."
brew services start mysql

# echo "Logging into mysql client..."
#mysql -uroot --password="password"

echo "Create database ohyou_api..."
mysql --defaults-extra-file=mysql.cnf -e \
  "CREATE DATABASE IF NOT EXISTS ohyou_api;"
echo "...success!"

echo "Create transactions table..."
mysql --defaults-extra-file=mysql.cnf -e \
  "USE ohyou_api; CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(100) NOT NULL,    
    purchaser INTEGER NOT NULL);"
echo "...success!"

echo "Create junction table transactions_beneficiaries..."
mysql --defaults-extra-file=mysql.cnf -e \
  "USE ohyou_api; CREATE TABLE IF NOT EXISTS transactions_beneficiaries (
    transaction_id INTEGER NOT NULL,
    beneficiary_id INTEGER NOT NULL);"

echo "...success!"
