DB_NAME="b3bas_framework"
YOUR_USER="zeroc0d3"
YOUR_PASSWORD="password"
TEMP_FILE="temp.sql"

printf "USE ${DB_NAME}\n"
cat 000-set-timezone.sql > $TEMP_FILE

# cat $TEMP_FILE
# psql -h localhost -U ${YOUR_USER} -W ${YOUR_PASSWORD} $DB_NAME < $TEMP_FILE
# rm $TEMP_FILE

