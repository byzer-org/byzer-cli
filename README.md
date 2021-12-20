## Byzer Cli

Command line interface of Byzer-lang

## Build

```
make all
```

## Run

Copy to Byzer-lang distribution bin directory. 

First, set MLSQL_HOME

Download mlsql lang:
1. [mlsql lang mac](https://mlsql-downloads.kyligence.io/2.1.0/mlsql-app_2.4-2.1.0-darwin-amd64.tar.gz)
2. [mlsql lang linux](https://mlsql-downloads.kyligence.io/2.1.0/mlsql-app_2.4-2.1.0-linux-amd64.tar.gz)

```
export MLSQL_HOME=....
```

Second, set PATH

```
export PATH=${MLSQL_HOME}/bin:$PATH
```

Third, run mlsql script

```
mlsql run ./src/common/hello.mlsql
```
