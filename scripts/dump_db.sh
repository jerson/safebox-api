#!/bin/sh
mysqldump -u safebox --password=safebox --no-data --no-create-db --databases safebox > database/02_structure.sql
mysqldump -u safebox --password=safebox --no-create-info --no-create-db --skip-triggers --where="true LIMIT 50" --databases safebox > database/03_data.sql

