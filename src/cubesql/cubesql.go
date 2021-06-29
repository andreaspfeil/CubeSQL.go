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

// #cgo CFLAGS: -Wno-multichar -Isdk/C_SDK
// #cgo LDFLAGS: -L. -lcubesql -ldl
// #include <stdlib.h>
// #include "cubesql.h"
import "C"

import (
  "unsafe"
	"reflect"
)

type CubeSQL struct {
	db 			*C.struct_csqldb
}
type CubeSQLCursor struct {
	cursor 	*C.struct_csqlc
}				
type CubeSQLPrepairedStatement struct {
  statement *C.struct_csqlvm
}

func New() *CubeSQL {
	return new( CubeSQL )
} 

// CUBESQL_APIEXPORT int		cubesql_connect (csqldb **db, const char *host, int port, const char *username, const char *password, int timeout, int encryption);
func (this *CubeSQL ) Connect( host string, port int, userName string, password string, timeout int, encryption int ) int {
	cHost := C.CString( host )
	defer C.free( unsafe.Pointer( cHost ) )

	cLogin := C.CString( userName )
	defer C.free( unsafe.Pointer( cLogin ) )
	
	cPassword := C.CString( password )
	defer C.free( unsafe.Pointer( cPassword ) )
	
	res := int( C.cubesql_connect( &this.db, cHost, C.int( port ), cLogin, cPassword, C.int( timeout ), C.int( encryption ) ) )
	if res == 0 {
		this.Execute( "SET CLIENT TYPE TO 'GoLang 1.0.0';" );
	}
	return res;
}
// CUBESQL_APIEXPORT int		cubesql_connect_ssl (csqldb **db, const char *host, int port, const char *username, const char *password, int timeout, char *ssl_certificate_path);
func (this *CubeSQL ) ConnectSSL( host string, port int, userName string, password string, timeout int, sslCertificatePath string ) int {
	cHost := C.CString( host )
	defer C.free( unsafe.Pointer( cHost ) )

	cLogin := C.CString( userName )
	defer C.free( unsafe.Pointer( cLogin ) )
	
	cPassword := C.CString( password )
	defer C.free( unsafe.Pointer( cPassword ) )

	cSSLCertPath := C.CString( sslCertificatePath )
	defer C.free( unsafe.Pointer( cSSLCertPath ) )

	res := int( C.cubesql_connect_ssl( &this.db, cHost, C.int( port ), cLogin, cPassword, C.int( timeout ), cSSLCertPath ) )
	if res == 0 {
		this.Execute( "SET CLIENT TYPE TO 'GoLang 1.0.0';" );
	}
	return res;
}
// CUBESQL_APIEXPORT void		cubesql_disconnect (csqldb *db, int gracefully);
func (this *CubeSQL ) Disconnect( gracefully int ) {
	C.cubesql_disconnect( this.db, C.int( gracefully ) )
}

 // CUBESQL_APIEXPORT int		cubesql_execute (csqldb *db, const char *sql);
func (this *CubeSQL ) Execute( sql string ) int {
	cSQL := C.CString( sql )
	defer C.free( unsafe.Pointer( cSQL ) )

	return int( C.cubesql_execute( this.db, cSQL ) )
}
// CUBESQL_APIEXPORT csqlc		*cubesql_select (csqldb *db, const char *sql, int unused);
func (this *CubeSQL ) Select( sql string, unused int ) *CubeSQLCursor {
	cSQL := C.CString( sql )
	defer C.free( unsafe.Pointer( cSQL ) )

	// cursor := new( CubeSQLCursor )
	// cursor.cursor = C.cubesql_select( this.db, cSQL, C.int( unused ) )
	cursor := CubeSQLCursor { cursor: C.cubesql_select( this.db, cSQL, C.int( unused ) ) }
	return &cursor
}

// CUBESQL_APIEXPORT int		cubesql_commit (csqldb *db);
func (this *CubeSQL ) Commit() int {
	return int( C.cubesql_commit( this.db ) )
}
// CUBESQL_APIEXPORT int		cubesql_rollback (csqldb *db);
func (this *CubeSQL ) Rollback() int {
	return int( C.cubesql_rollback( this.db ) )
}
// CUBESQL_APIEXPORT int		cubesql_bind (csqldb *db, const char *sql, char **colvalue, int *colsize, int *coltype, int ncols);
func (this *CubeSQL ) Bind( sql string, values []interface{} ) int {
	var ncols int = len( values )
	colval  := make( []*C.char, ncols)
	colsize := make( []C.int, ncols )
	coltype := make( []C.int, ncols )

	cSQL := C.CString( sql )
	defer C.free( unsafe.Pointer( cSQL ) )

	for index, interfaceValue := range values {
		switch value := interfaceValue.(type) {
		case nil:
			coltype[ index ] = BIND_NULL
			colsize[ index ] = 1

			colval[ index ] = C.CString( "" )
			defer C.free( unsafe.Pointer( colval[ index ] ) )
			
		case int:
			coltype[ index ] = BIND_INTEGER
			colsize[ index ] = C.int( 1 + unsafe.Sizeof( C.int(0) ) )

			i := C.int( int( value ) )
			colval[ index ] = (*C.char)(unsafe.Pointer( &i ) )
			defer C.free( unsafe.Pointer( &i ) )
			
		case int64:
			coltype[ index ] = BIND_INT64
			colsize[ index ] = C.int( 1 + unsafe.Sizeof( C.int64(0) ) )

			i := C.int64( int( value ) )
			colval[ index ] = (*C.char)(unsafe.Pointer( &i ) )
			defer C.free( unsafe.Pointer( &i ) )
			
		case float32:
			coltype[ index ] = BIND_DOUBLE
			colsize[ index ] = C.int( 1 + unsafe.Sizeof( C.double(0) ) )

			f := C.double( float32( value ) )
			colval[ index ] = (*C.char)(unsafe.Pointer( &f ) )
			defer C.free( unsafe.Pointer( &f ) )
			
		case float64:
			coltype[ index ] = BIND_DOUBLE
			colsize[ index ] = C.int( 1 + unsafe.Sizeof( C.double(0) ) )

			f := C.double( float64( value ) )
			colval[ index ] = (*C.char)(unsafe.Pointer( &f ) )
			defer C.free( unsafe.Pointer( &f ) )
			
		case string:
			coltype[ index ] = BIND_TEXT
			colsize[ index ] = C.int( 1 + len( value ) )

			colval[ index ] = C.CString( value )
			defer C.free( unsafe.Pointer( colval[ index ] ) )
			
		case []byte:
			colsize[ index ] = C.int( len( value ) )
			if colsize[ index ] == 0 {
				colval[ index ] = C.CString( "" )
				coltype[ index ] = BIND_ZEROBLOB
			} else {
				colval[ index ] = (*C.char)(C.CBytes( value ))
				coltype[ index ] = BIND_BLOB
			}
			defer C.free( unsafe.Pointer( colval[ index ] ) )
		default:
			return ERR
		}
	}

	return int( C.cubesql_bind( this.db, cSQL, (**C.char)(unsafe.Pointer( colval[ 0 ] ) ), (*C.int)(unsafe.Pointer( &colsize[ 0 ] ) ), (*C.int)(unsafe.Pointer( &coltype[ 0 ] ) ), C.int( ncols ) ) )
}

// CUBESQL_APIEXPORT int		cubesql_ping (csqldb *db);
func (this *CubeSQL ) Ping() int {
	return int( C.cubesql_ping( this.db ) )
}
// CUBESQL_APIEXPORT void		cubesql_cancel (csqldb *db);
func (this *CubeSQL ) Cancel() {	
 C.cubesql_cancel( this.db )
}
// CUBESQL_APIEXPORT int		cubesql_errcode (csqldb *db);
func (this *CubeSQL ) ErrorCode() int {
	return int( C.cubesql_errcode( this.db ) )
}
// CUBESQL_APIEXPORT char		*cubesql_errmsg (csqldb *db);
func (this *CubeSQL ) ErrorMessage() string {
	// https://gist.github.com/zchee/b9c99695463d8902cd33
	return C.GoString( C.cubesql_errmsg( this.db ) )
}
func (this *CubeSQL ) Error() ( int, string ) {
	return this.ErrorCode(), this.ErrorMessage()
}

// CUBESQL_APIEXPORT int64		cubesql_changes (csqldb *db);
func (this *CubeSQL ) Changes() int64 {
	return int64( C.cubesql_changes( this.db ) )
}
// CUBESQL_APIEXPORT void		cubesql_trace (csqldb *db, trace_function trace, void *arg);
// typedef void (*trace_function) (const char*, void*);
func (this *CubeSQL ) Trace( function func(), arg string ) {

	// typedef void (*trace_function) (const char*, void*);

	// https://www.lobaro.com/embedded-development-with-c-and-golang-cgo/
	//C.cubesql_trace( this.db, function, 1 )
}

////////// CubeSQL Cursor Functions

// CUBESQL_APIEXPORT int		cubesql_cursor_numrows (csqlc *c);
func( this *CubeSQLCursor ) NumRows() int {
	return int( C.cubesql_cursor_numrows( this.cursor ) )
}
// CUBESQL_APIEXPORT int		cubesql_cursor_numcolumns (csqlc *c);
func( this *CubeSQLCursor ) NumColumns() int {
	return int( C.cubesql_cursor_numcolumns( this.cursor ) )
}
// CUBESQL_APIEXPORT int		cubesql_cursor_currentrow (csqlc *c);
func( this *CubeSQLCursor ) CurrentRow() int {
	return int( C.cubesql_cursor_currentrow( this.cursor ) )
}
// CUBESQL_APIEXPORT int		cubesql_cursor_seek (csqlc *c, int index);
func( this *CubeSQLCursor ) Seek( index int ) int {
	return int( C.cubesql_cursor_seek( this.cursor, C.int( index ) ) )
}
// CUBESQL_APIEXPORT int		cubesql_cursor_iseof (csqlc *c);
func( this *CubeSQLCursor ) IsEOF() int {
	return int( C.cubesql_cursor_iseof( this.cursor ) )
}
// CUBESQL_APIEXPORT int		cubesql_cursor_columntype (csqlc *c, int index);
func( this *CubeSQLCursor ) ColumnType( index int ) int {
	return int( C.cubesql_cursor_columntype( this.cursor, C.int( index ) ) )
}
// CUBESQL_APIEXPORT char		*cubesql_cursor_field (csqlc *c, int row, int column, int *len);
func( this *CubeSQLCursor ) GetField( row int, column int ) ( string, int ) {
	var len C.int = 0
	return C.GoString( C.cubesql_cursor_field( this.cursor, C.int( row ), C.int( column ), &len ) ), int( len ) // Problem: NULL Pointer in return
}
// CUBESQL_APIEXPORT int64		cubesql_cursor_rowid (csqlc *c, int row);
func( this *CubeSQLCursor ) RowID( row int ) int64 {
	return int64( C.cubesql_cursor_rowid( this.cursor, C.int( row ) ) )
}
// CUBESQL_APIEXPORT int64		cubesql_cursor_int64 (csqlc *c, int row, int column, int64 default_value);
func( this *CubeSQLCursor ) Int64( row int, column int, defaultValue int ) int64 {
	return int64( C.cubesql_cursor_int64( this.cursor, C.int( row ), C.int( column ), C.int64( defaultValue ) ) )
}
// CUBESQL_APIEXPORT int		cubesql_cursor_int (csqlc *c, int row, int column, int default_value);
func( this *CubeSQLCursor ) Int( row int, column int, defaultValue int ) int {
	return int( C.cubesql_cursor_int( this.cursor, C.int( row ), C.int( column ), C.int( defaultValue ) ) )
}
// CUBESQL_APIEXPORT double	cubesql_cursor_double (csqlc *c, int row, int column, double default_value);
func( this *CubeSQLCursor ) Float32( row int, column int, defaultValue float32 ) float32 {
	return float32( C.cubesql_cursor_double( this.cursor, C.int( row ), C.int( column ), C.double( defaultValue ) ) )
}
func( this *CubeSQLCursor ) Float64( row int, column int, defaultValue float64 ) float64 {
	return float64( C.cubesql_cursor_double( this.cursor, C.int( row ), C.int( column ), C.double( defaultValue ) ) )
}
// CUBESQL_APIEXPORT char		*cubesql_cursor_cstring (csqlc *c, int row, int column);
func( this *CubeSQLCursor ) String( row int, column int ) string {
	return C.GoString( C.cubesql_cursor_cstring( this.cursor, C.int( row ), C.int( column ) ) )
}
// CUBESQL_APIEXPORT char		*cubesql_cursor_cstring_static (csqlc *c, int row, int column, char *static_buffer, int bufferlen);
func( this *CubeSQLCursor ) Bytes( row int, column int, buffer *[]byte ) *[]byte {
	b := reflect.ValueOf( buffer )
	C.cubesql_cursor_cstring_static( this.cursor, C.int( row ), C.int( column ), (*C.char)(unsafe.Pointer( b.Pointer() )), C.int( b.Len() ) )
	return buffer
}
// CUBESQL_APIEXPORT void		cubesql_cursor_free (csqlc *c);
func( this *CubeSQLCursor ) Free()  {
	C.cubesql_cursor_free( this.cursor )
}

////////// Prepare Statements CubeSQLPrepairedStatement

// CUBESQL_APIEXPORT csqlvm	*cubesql_vmprepare (csqldb *db, const char *sql);
func (this *CubeSQL ) Prepare( sql string ) *CubeSQLPrepairedStatement {
	cSQL := C.CString( sql )
	defer C.free( unsafe.Pointer( cSQL ) )

	statement := CubeSQLPrepairedStatement { statement: C.cubesql_vmprepare( this.db, cSQL ) }
	return &statement
}
// CUBESQL_APIEXPORT int		cubesql_vmbind_int (csqlvm *vm, int index, int value);
func (this *CubeSQLPrepairedStatement ) BindInt( index int, value int ) int {
	return int( C.cubesql_vmbind_int( this.statement, C.int( index ), C.int( value ) ) )
}
// CUBESQL_APIEXPORT int		cubesql_vmbind_double (csqlvm *vm, int index, double value);
func (this *CubeSQLPrepairedStatement ) BindDouble( index int, value float32 ) int {
	return int( C.cubesql_vmbind_double( this.statement, C.int( index ), C.double( value ) ) )
}
// CUBESQL_APIEXPORT int		cubesql_vmbind_text (csqlvm *vm, int index, char *value, int len);
func (this *CubeSQLPrepairedStatement ) BindText( index int, value string ) int {
	cValue := C.CString( value )
	defer C.free( unsafe.Pointer( cValue ) )
	return int( C.cubesql_vmbind_text( this.statement, C.int( index ), cValue, C.int( len( value ) ) ) )
}
// CUBESQL_APIEXPORT int		cubesql_vmbind_blob (csqlvm *vm, int index, void *value, int len);
func (this *CubeSQLPrepairedStatement ) BindBlob( index int, value []byte ) int {
	cValue := C.CBytes( value )
	defer C.free( unsafe.Pointer( cValue ) )
	return int( C.cubesql_vmbind_blob( this.statement, C.int( index ), cValue, C.int( len( value ) ) ) )
}
// CUBESQL_APIEXPORT int		cubesql_vmbind_null (csqlvm *vm, int index);
func (this *CubeSQLPrepairedStatement ) BindNull( index int ) int {
	return int( C.cubesql_vmbind_null( this.statement, C.int( index ) ) )
}
// CUBESQL_APIEXPORT int		cubesql_vmbind_int64 (csqlvm *vm, int index, int64 value);
func (this *CubeSQLPrepairedStatement ) BindInt64( index int, value int64 ) int {
	return int( C.cubesql_vmbind_int64( this.statement, C.int( index ), C.int64( value ) ) )
}
// CUBESQL_APIEXPORT int		cubesql_vmbind_zeroblob (csqlvm *vm, int index, int len);
func (this *CubeSQLPrepairedStatement ) BindZeroBlob( index int, len int ) int {
	return int( C.cubesql_vmbind_zeroblob( this.statement, C.int( index ), C.int( len ) ) )
}
// CUBESQL_APIEXPORT int		cubesql_vmexecute (csqlvm *vm);
func (this *CubeSQLPrepairedStatement ) Execute() int {
	return int( C.cubesql_vmexecute( this.statement ) )
}
// CUBESQL_APIEXPORT csqlc		*cubesql_vmselect (csqlvm *vm);
func (this *CubeSQLPrepairedStatement ) Select() *CubeSQLCursor {
	cursor := CubeSQLCursor { cursor: C.cubesql_vmselect( this.statement ) }
	return &cursor
}
// CUBESQL_APIEXPORT int		cubesql_vmclose (csqlvm *vm);
func (this *CubeSQLPrepairedStatement ) Close() int {
	return int( C.cubesql_vmclose( this.statement ) )
}