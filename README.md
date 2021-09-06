## MLSQL Lang Cli

mlsql lang command line.

## Build

```
make all
```

## Run

Copy to MLSQL lang distribution bin directory. 

First, set MLSQL_HOME

> Download mlsql lang:
> [mlsql lang mac](http://download.mlsql.tech/mlsql-lang-darwin-amd64.tar.gz)
> [mlsql lang linux](http://download.mlsql.tech/mlsql-lang-linux-amd64.tar.gz)

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
