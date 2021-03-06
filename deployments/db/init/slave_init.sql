CREATE USER IF NOT EXISTS 'exporter'@'%' IDENTIFIED BY 'password' WITH MAX_USER_CONNECTIONS 3;
GRANT PROCESS, REPLICATION CLIENT, SELECT ON *.* TO 'exporter'@'%';

FLUSH PRIVILEGES;

/* start replication */
CHANGE MASTER TO MASTER_HOST='db-master',MASTER_USER='otus_social_slave', MASTER_PASSWORD='otus_social_passwd', MASTER_LOG_FILE='mysql-bin.000001', MASTER_LOG_POS=1;
START SLAVE;