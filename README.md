# cubeSQL golang client

Bug free, feature complete, uses Marco Bambinis nativ C SDK.

## Usage example

```golang
package main

import "cubesql"
import "fmt"

func main() {
  cube := cubesql.New()
  if cube.Connect( "dbhost", 4430, "loginname", "password", 10, 0 ) == cubesql.NOERR {
    defer cube.Disconnect( 0 )

    fmt.Printf( "PING... " )
    if cube.Ping() == cubesql.ERR {
      fmt.Printf( "ERROR %d (%s).\r\n", cube.ErrorCode(), cube.ErrorMessage() )
    } else {
      println( "SUCCESS." )

      cube.Use( "demo" )
      cube.AutoCommit( true )

      cube.Execute( `CREATE TABLE IF NOT EXISTS "Friends" ("Name" TEXT PRIMARY KEY NOT NULL UNIQUE, "Birthday" DATE);` )

      // Inserting values...
      cube.Execute( `INSERT OR IGNORE INTO Friends VALUES ( "Buddy1", "1974-05-14" )` )
      cube.Execute( fmt.Sprintf( `INSERT OR IGNORE INTO Friends VALUES ( "%s", "%s" )`, "Buddy2", "2009-01-14" ) )

      var values []interface{}
      cube.Bind( `INSERT OR IGNORE INTO Friends VALUES ( ?1, ?2 )`, append( values, "Buddy3", "2007-07-18" ) )

      statement := cube.Prepare( `INSERT OR IGNORE INTO Friends VALUES ( ?1, ?2 )` )
      statement.BindText( 1, "Buddy4" )
      statement.BindText( 2, "1974-07-26" )
      statement.Execute()
      statement.Close()

      result := cube.Select( "SELECT * FROM Friends;" )
      defer result.Free()

      fmt.Printf( "%d Rows with %d Coulmns found in Table 'Friends':\r\n", result.NumRows(), result.NumColumns() )

      for col:= 1; col <= result.NumColumns(); col++ {
        fmt.Printf( "Column %d: Name = %-10s, Type = %2d\r\n", col, result.GetField( cubesql.COLNAME, col ), result.ColumnType( col ) )
      }

      for {
        fmt.Printf( "Row %02d: ", result.CurrentRow() )
        for col:= 1; col <= result.NumColumns(); col++ {
          fmt.Printf( "%s | ", result.String( cubesql.CURROW, col ) )
        }
        println( "" )

        result.Seek( cubesql.SEEKNEXT )
        if result.IsEOF() == cubesql.TRUE {
          break
        }
      }

      // More compact... (but wrong! - Do you know why? HINT: Free your mind to figuer it out ;-)
      for result := cube.Select( "SELECT * FROM Friends ORDER BY Birthday;" ); result.IsEOF() == cubesql.FALSE; result.Seek( cubesql.SEEKNEXT ) {
        println( result.String( cubesql.CURROW, 1 ) )
      }
    }
  }    
}
```

## Installation
### MacOS, Linux, ...
1. Set up your go development environment:
```console
go env -w GO111MODULE=off
mkdir ~/MyProject
cd ~/MyProject

```
2. Download CubeSQL.go
```console
git clone https://github.com/andreaspfeil/CubeSQL.go.git
```
3. Enter CubeSQL.go directory and compile the native database driver
```console
cd CubeSQL.go/src/cubesql/
make
```
4. Test the example programm:
```console
cd ../..
export GOPATH=`pwd`
echo $GOPATH
cd src
go run .
```
5. Why does the example programm not work?\
Make sure that your database is running and can be reached from your workstation. Also change your credentials in:
```golang
cube.Connect( "dbhost", 4430, "loginname", "password", 10, 0 ) 
```
and use the correct database:
```golang
cube.Use( "demo" )
```
### Windows
If you are interersted in how to install the CubeSQL.go driver on Windows, drop me a line.

## Future Ideas
- [ ] Windows Installer 
- [ ] More handy methods
- [ ] Better Error handling
- [ ] Make a go module for even easier integration into your project

If you are interersted in any of this, drop me a line or please consider buying me a beer...

## Documentation

- [Wiki](https://github.com/andreaspfeil/CubeSQL.go/wiki)

## Video Tutorials

- [YouTube](https://www.youtube.com/channel/UCQF_wTmbR5aJZUcb7U1_0Fw)

## Donate

- [github](https://github.com/sponsors/andreaspfeil)
- [Patreon](https://www.patreon.com/andreas_pfeil)
- [PayPal](https://www.paypal.com/paypalme/PfeilAndreas/10.00EUR)

## Contributors

- [Marco Bambini](https://github.com/marcobambini) (Author of cubeSQL and the original nativ client SDK)

## Acknowledgments

- [cubeSQL](https://sqlabs.com/cubesql)

## See also

- [cubeSQL.Python2](https://github.com/andreaspfeil/CubeSQL.Python2)
- [cubeSQL.Python3](https://github.com/andreaspfeil/CubeSQL.Python3)
- [cubeSQL.NET](https://github.com/andreaspfeil/CubeSQL.NET)

## License

[BEER license / MIT license](https://github.com/andreaspfeil/CubeSQL.go/blob/main/LICENSE) 

The BEER license is basically the same as the MIT license [(see link)](https://github.com/andreaspfeil/CubeSQL.go/blob/main/LICENSE), except 
that you should buy the author a beer [(see Donate)](https://github.com/andreaspfeil/CubeSQL.go#donate) if you use this software.

## Sponsors

none yet - YOU can still be number one in this list!!!
