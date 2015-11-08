# BeegoTDDBootStrap
A bootstrap project for TDD using beego framework.

## Dev tips
+  modify the `.ci/prepare_db.sh`, `.ci/clean_db.sh`, `conf/app.conf` files, to configure your DB.
+  TDD steps
  +  dnderstand the demand
  +  design API, write API docs, design test cases
  +  write API test in `tests` package
  +  design and implement API, modify test cases if need
  +  if other classes/functions are needed during implementing API, follow the same step as above
+  CI support
  +  install mysql
  +  modify the `.ci/prepare_db.sh`, `.ci/clean_db.sh`, `conf/app.conf` files, to configure your DB.
  +  execute `.ci/ci.sh` to check