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

// go build -o cubeSQL main.go

import "cubesql"
import "time"


func main() {
  cube := cubesql.New()
  result := cube.Connect( "localhost", 4430, "loginname", "password", 10, 0 );
  println( result )
  time.Sleep( 10 * time.Second )
  cube.Disconnect( 0 );
}