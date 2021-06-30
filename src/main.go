//
//         ___              __                        ____  ____     _ __
//        /   |  ____  ____/ /_______  ____ ______   / __ \/ __/__  (_) /
//       / /| | / __ \/ __  / ___/ _ \/ __ `/ ___/  / /_/ / /_/ _ \/ / / 
//      / ___ |/ / / / /_/ / /  /  __/ /_/ (__  )  / ____/ __/  __/ / /  
//     /_/  |_/_/ /_/\__,_/_/   \___/\__,_/____/  /_/   /_/  \___/_/_/
//                                                          
//  Product:     cubeSQL.go - Demo App for the C SDK golang wrapper
//  Version:     Revision: 1.0.0, Build: 1
//  Date:        2021/06/03 21:58:48
//  Author:      Andreas Pfeil <patreon@familie-pfeil.com>
//
//  Description: Opens a cubeSQL database connection with the help of
//               Marco Bambini's native C SDK driver, selects a database 
//               and selects some rows. Outputs info and the rows.
//
//  Usage:       import "cubesql"
//
//  License:     BEER license / MIT license
//
//  Copyright (C) 2021 by Andreas Pfeil
//
// -----------------------------------------------------------------------TAB=2

package main

// go run .
// go build -o demo main.go

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