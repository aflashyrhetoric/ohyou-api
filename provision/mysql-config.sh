# mysql-config.sh
## configure mysql database
## this file is idempotent and can be repeatedly run safely 

echo "Starting Mariadb service..."
brew services start mariadb

# echo "Logging into mysql client..."
#mysql -uroot --password="password"

echo "Create database payup_api..."
mysql --defaults-extra-file=mysql.cnf -e \
  "CREATE DATABASE IF NOT EXISTS payup_api;"
echo "...success!"

echo "Create transactions table..."
mysql --defaults-extra-file=mysql.cnf -e \
  "USE payup_api; CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(200),    
    purchaser INTEGER,
    amount INTEGER);"
echo "...success!"

echo "Create junction table transactions_beneficiaries..."
mysql --defaults-extra-file=mysql.cnf -e \
  "USE payup_api; CREATE TABLE IF NOT EXISTS transactions_beneficiaries (
    transaction_id INTEGER,
    beneficiary_id INTEGER);"

echo "...success!"

# Run the seeder
go run seeder.go