//
//         ___              __                        ____  ____     _ __
//        /   |  ____  ____/ /_______  ____ ______   / __ \/ __/__  (_) /
//       / /| | / __ \/ __  / ___/ _ \/ __ `/ ___/  / /_/ / /_/ _ \/ / / 
//      / ___ |/ / / / /_/ / /  /  __/ /_/ (__  )  / ____/ __/  __/ / /  
//     /_/  |_/_/ /_/\__,_/_/   \___/\__,_/____/  /_/   /_/  \___/_/_/
//                                                          
//  Product:     cubeSQL.go - Wrapper for the cubeSQL C SDK database driver
//  Version:     Revision: 1.0.0, Build: 1
//  Date:        2021/06/29 21:04:11
//  Author:      Andreas Pfeil <patreon@familie-pfeil.com>
//
//  Description: golang wrapper for the cubeSQL database client driver based 
//               on Marco Bambini's C SDK.
//
//  Usage:       import "cubesql"
//
//  License:     BEER license / MIT license
//
//  Copyright (C) 2021 by Andreas Pfeil
//
// -----------------------------------------------------------------------TAB=2

package cubesql

import "fmt"

func (this *CubeSQL ) Use( database string ) int {
  return this.Execute( fmt.Sprintf( "USE DATABASE %s;", database ) )
}

func (this *CubeSQL ) AutoCommit( enabled bool ) int {
  commit := "OFF"
  if enabled {
    commit = "ON"
  }
  return this.Execute( fmt.Sprintf( "SET AUTOCOMMIT TO %s;", commit ) )
}