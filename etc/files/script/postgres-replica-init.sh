#!/bin/bash

echo "[${SLAVE_NAME}] Initializing replication..."

REPLICATION_SLOT_NAME=$(echo "${SLAVE_NAME}" | tr -cd 'a-z0-9_' | tr '-' '_')
echo "[${SLAVE_NAME}] Sanitized replication slot name: ${REPLICATION_SLOT_NAME}"

# Function to log and exit on failure
function log_and_exit {
    echo "[${SLAVE_NAME}] $1"
    exit 1
}

# Wait for master database to be ready
until psql -h "${POSTGRES_MASTER_HOST}" -p "${POSTGRES_MASTER_PORT}" -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -c '\l' > /dev/null 2>&1; do
    echo "[${SLAVE_NAME}] Waiting for the master database to be ready..."
    sleep 1
done
echo "[${SLAVE_NAME}] Master database is ready."

# Check if replication slot exists
if psql -h "${POSTGRES_MASTER_HOST}" -p "${POSTGRES_MASTER_PORT}" -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -c "SELECT * FROM pg_replication_slots WHERE slot_name = '${REPLICATION_SLOT_NAME}';" | grep "${REPLICATION_SLOT_NAME}" > /dev/null; then
    echo "[${SLAVE_NAME}] Replication slot already exists."
else
    echo "[${SLAVE_NAME}] Replication slot does not exist, creating new one..."
    psql -h "${POSTGRES_MASTER_HOST}" -p "${POSTGRES_MASTER_PORT}" -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -c "SELECT pg_create_physical_replication_slot('${REPLICATION_SLOT_NAME}');" || log_and_exit "Failed to create replication slot."
fi

# Check if data directory is empty, then clone master
if [ -z "$(ls -A /var/lib/postgresql/data)" ]; then
    echo "[${SLAVE_NAME}] Data directory is empty, cloning master..."
    until pg_basebackup --slot="${REPLICATION_SLOT_NAME}" --host="${POSTGRES_MASTER_HOST}" --port="${POSTGRES_MASTER_PORT}" --username="${POSTGRES_USER}" --pgdata=/var/lib/postgresql/data -R; do
        echo "[${SLAVE_NAME}] Waiting for connection from master db..."
        sleep 1
    done || log_and_exit "Failed to clone master data."
fi

# Finalizing and starting replication
echo "[${SLAVE_NAME}] Replication successfully initialized."
chmod 0700 /var/lib/postgresql/data || log_and_exit "Failed to set permissions on data directory."
echo "[${SLAVE_NAME}] Starting replication..."
exec "$@"
