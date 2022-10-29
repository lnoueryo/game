
## environment
makeコマンドがない場合はインストール
windows: https://gnuwin32.sourceforge.net/packages/make.htm

.envを作成
```
DB_HOST=
DB_PORT=
DB_DATABASE=
DB_USER=
DB_PASSWORD=
TABLE_SERVER_PORT=
```

dockerを起動し下記コマンドを入力
```
$ make init
```
Dockerのネットワーク、イメージ、コンテナが作成される。
以降は下記コマンドで立ち上げ可能。
```
$ make start
```
もしイメージも作成しなおしたい場合は下記コマンド。
```
$ make restart
```

## migrate
```
$ make seed
```
## test
不整合の確認を考慮してテスト回数を数十回行う。
```
$ go test ./interface/controllers/... -count 30
```

## error
###### Pool overlaps with other one on this address space
利用されていないネットワーク設定を削除することで解消。使用していないネットワークを全削除するコマンドは下記の通り。
```
$ docker system prune
```