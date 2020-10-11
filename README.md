# a


参考: https://github.com/nakabonne/cleanarchitecture-sample

> curl.exe http://localhost:8081/user -X POST

mysql.server start --skip-grant-tables

root saijopw

http://www.gitignore.io/api/list
http://www.gitignore.io/api/go,windows,macos,visualstudiocode,vim

power shellさんやおかしくないか（はまった）
`> curl.exe -X POST http://localhost:8081/user  --header 'Content-Type: application/json' -d '"{\"name\":\"saijo\",\"email\":\"hogehoge\"}"'`

テスト
- request
- usecase

階層

|階層|説明|
|:-:|:-:|
|router|requestのルーティング|
|controller|request,responseの処理、逐次フロー|
|usecase|modelの処理、バリデーション|
|model|データの結合など|
|db|db_instanceとのやり取り|

検討

- サブディレクトリ(model/domain etc)のパッケージはほかのルートディレクトリ（usecase, controller etc）で利用しないようにする？
- interfaceとmodelパッケージとかは、外部から呼ばれるので制約としては厳しい？
