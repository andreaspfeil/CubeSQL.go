//
//         ___              __                        ____  ____     _ __
//        /   |  ____  ____/ /_______  ____ ______   / __ \/ __/__  (_) /
//       / /| | / __ \/ __  / ___/ _ \/ __ `/ ___/  / /_/ / /_/ _ \/ / / 
//      / ___ |/ / / / /_/ / /  /  __/ /_/ (__  )  / ____/ __/  __/ / /  
//     /_/  |_/_/ /_/\__,_/_/   \___/\__,_/____/  /_/   /_/  \___/_/_/
//                                                          
//  Product:     cubeSQL.go - Wrapper for the cubeSQL C SDK database driver
//  Version:     Revision: 1.0.0, Build: 1
//  Date:        2021/06/29 21:01:33
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

// #cgo CFLAGS: -Wno-multichar -Isdk/C_SDK
// #include "cubesql.h"
import "C"

const TRUE = 1 // custom boolean values
const FALSE = 0

const DEFAULT_PORT = 4430 // default values
const DEFAULT_TIMEOUT =	12
const NOERR	= 0
const ERR	= -1

const MEMORY_ERROR	=	-2			// client side errors
const PARAMETER_ERROR =	-3
const PROTOCOL_ERROR	= -4
const ZLIB_ERROR	= -5
const SSL_ERROR = -6
const SSL_CERT_ERROR =	-7

const AESNONE = 0 // encryption flags used in cubesql_connect
const AES128	=	2
const AES192	=	3
const AES256	=	4
const SSL = 8
const SSL_AES128	=	SSL + AES128
const SSL_AES192	=	SSL + AES192
const SSL_AES256	=	SSL + AES256

const CUBESQL_COLNAME = 0 // flag used in cubesql_cursor_getfield
const CUBESQL_CURROW	= -1
const CUBESQL_COLTABLE	= -2
const CUBESQL_ROWID = -665 - 1

const CUBESQL_SEEKNEXT	= -2 // flag used in cubesql_cursor_seek
const CUBESQL_SEEKFIRST	= -3
const CUBESQL_SEEKLAST	= -4
const CUBESQL_SEEKPREV	= -5

const SSL_LIBRARY_PATH	=	1 // SSL
const CRYPTO_LIBRARY_PATH = 2

const TYPE_None		= 0 // column types coming from the server
const TYPE_Integer	= 1
const TYPE_Float		= 2
const TYPE_Text		= 3
const TYPE_Blob		= 4
const TYPE_Boolean	= 5
const TYPE_Date		= 6
const TYPE_Time		= 7
const TYPE_Timestamp	= 8
const TYPE_Currency	= 9

const BIND_INTEGER	= 1 // column types to specify in the cubesql_bind command (coltype)
const BIND_DOUBLE	= 2
const BIND_TEXT	= 3
const BIND_BLOB	= 4
const BIND_NULL	= 5
const BIND_INT64	= 8
const BIND_ZEROBLOB	= 9