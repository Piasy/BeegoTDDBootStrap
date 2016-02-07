# BeegoTDDBootStrap
A bootstrap project for TDD using beego framework.

## Dev tips
+  modify the `.ci/prepare_db.sh`, `.ci/clean_db.sh`, `conf/app.conf` files, to configure your DB.
+  create `logs` dir, and touch `orm.log`, then run `bee run`.
+  TDD steps
  +  understand your demand
  +  design API, write API docs, design test cases
  +  write API test in `tests` package
  +  design and implement API, modify test cases if need
  +  if other classes/functions are needed during implementing API, follow the same step as above
+  CI support (manually)
  +  install mysql
  +  modify the `.ci/prepare_db.sh`, `.ci/clean_db.sh`, `conf/app.conf` files, to configure your DB.
  +  execute `.ci/ci.sh` to check
+  Change namespace
  +  Open this project with IntelliJ IDEA
  +  Use global search/replace, substitute `github.com/Piasy/BeegoTDDBootStrap` into the proper gopath string, e.g. `github.com/Piasy/HabitsAPI`
