# sqliteの使い方

keinos/sqlite3にて、、、

自分的解釈  
SQLiteは、ファイルでDBを簡易的に管理している。  
今回のmain.goではdb/example.sqlが１つのデータベースになっている。  
その中にusersテーブルを作成している。  

コンテナ立ち上げ
```bash
$ docker-compose up -d
```

sqliteコンテナに入る
```bash
$ docker-compose exec sqlite3 bash
```

コンテナ内コマンド(ここでのコマンド、都度ファイル名は変更してください)
```bash
$ sqlite3 ./workspace/example.sql
```
あとは自由にsqlite3コマンド実行出来る。
