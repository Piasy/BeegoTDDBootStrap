#!/bin/sh
mysql -uroot -Nse 'show tables' beego_unit_test | while read table; do echo "drop table $table;"; done | mysql -uroot beego_unit_test
