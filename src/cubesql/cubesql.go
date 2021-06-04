package cubesql

//// #cgo CFLAGS: -Wno-multichar -Isdk -Isdk/zlib -Isdk/crypt
//// #cgo LDFLAGS: -L. -lcubesql -ldl

// #cgo CFLAGS: -Wno-multichar -Isdk -Isdk/zlib -Isdk/crypt
// #cgo LDFLAGS: -L./sdk -l:libcubesql.a -ldl
// #include <stdlib.h>
// #include "cubesql.h"
import "C"
import (
  "unsafe"
	"time"
)

type CubeSQL struct {
	db 			*C.struct_csqldb
}

type CubeSQLCursor 							*C.struct_csqlc
type CubeSQLPrepairedStatement 	C.struct_csqlvm


func newCubeSQL() *CubeSQL {
	this := CubeSQL { db: nil }
	return &this
} 

func (this *CubeSQL ) connect( host string, port int, userName string, password string, timeout int, encryption int ) int {
	Host := C.CString( host )
	defer C.free( unsafe.Pointer( Host ) )

	Login := C.CString( userName )
	defer C.free( unsafe.Pointer( Login ) )
	
	Password := C.CString( password )
	defer C.free( unsafe.Pointer( Password ) )
	
	res := int( C.cubesql_connect( &this.db, Host, C.int( port ), Login, Password, C.int( timeout ), C.int( encryption ) ) )
	if res == 0 {
		this.Execute( "SET CLIENT TYPE TO 'GoLang 1.0.0';" );
	}
	return res;
}
func (this *CubeSQL ) disconnect( gracefully int ) {
	C.cubesql_disconnect( this.db, C.int( gracefully ) )
}
func (this *CubeSQL ) Execute( sql string ) int {
 // CUBESQL_APIEXPORT int		cubesql_execute (csqldb *db, const char *sql);
	SQL := C.CString( sql )
	defer C.free( unsafe.Pointer( SQL ) )
	return int( C.cubesql_execute( this.db, SQL ) )
}
func (this *CubeSQL ) commit() int {
 	// CUBESQL_APIEXPORT int		cubesql_commit (csqldb *db);
	return int( C.cubesql_commit( this.db ) )
}
func (this *CubeSQL ) rollback() int {
 	// CUBESQL_APIEXPORT int		cubesql_rollback (csqldb *db);
	return int( C.cubesql_rollback( this.db ) )
}
func (this *CubeSQL ) ping() int {
 	// CUBESQL_APIEXPORT int		cubesql_ping (csqldb *db);
	return int( C.cubesql_ping( this.db ) )
}
func (this *CubeSQL ) error() int {
 	// CUBESQL_APIEXPORT int		cubesql_errcode (csqldb *db);
	return int( C.cubesql_errcode( this.db ) )
}
func (this *CubeSQL ) changes() int64 {
 	// CUBESQL_APIEXPORT int64		cubesql_changes (csqldb *db);
	return int64( C.cubesql_changes( this.db ) )
}
func (this *CubeSQL ) cancel() {
 	// CUBESQL_APIEXPORT void		cubesql_cancel (csqldb *db);
	C.cubesql_cancel( this.db )
}
// func (this *CubeSQL ) cancel() {
// https://www.lobaro.com/embedded-development-with-c-and-golang-cgo/
//  	// CUBESQL_APIEXPORT void		cubesql_trace (csqldb *db, trace_function trace, void *arg);
// 	C.cubesql_trace( this.db )
// }
func (this *CubeSQL ) errmsg() string {
	// https://gist.github.com/zchee/b9c99695463d8902cd33
 	// CUBESQL_APIEXPORT char		*cubesql_errmsg (csqldb *db);
	msg := C.cubesql_errmsg( this.db )
	return C.GoString(msg)
}
//func (this *CubeSQL ) bind( sql string, colvalue string, colsize int, coltype int, ncols int ) int {
//	// https://gist.github.com/zchee/b9c99695463d8902cd33
// 	// CUBESQL_APIEXPORT int		cubesql_bind (csqldb *db, const char *sql, char **colvalue, int *colsize, int *coltype, int ncols);
//
//	SQL := C.CString( sql )
//	defer C.free( unsafe.Pointer( SQL ) )
//	
//	return int( C.cubesql_bind( this.db, SQL,  ) )
//}
func (this *CubeSQL ) Select(sql string, unused int ) CubeSQLCursor {
	// https://gist.github.com/zchee/b9c99695463d8902cd33
 	// CUBESQL_APIEXPORT csqlc		*cubesql_select (csqldb *db, const char *sql, int unused);
	SQL := C.CString( sql )
	defer C.free( unsafe.Pointer( SQL ) )
	
	return C.cubesql_select( this.db, SQL, C.int( unused ) )
}

////////// Prepare Statements CubeSQLPrepairedStatement

//func (this *CubeSQL ) Prepare( sql string ) CubeSQLPrepairedStatement {
//	// CUBESQL_APIEXPORT csqlvm	*cubesql_vmprepare (csqldb *db, const char *sql);
//	SQL := C.CString( sql )
//	defer C.free( unsafe.Pointer( SQL ) )
//	res := C.cubesql_vmprepare( this.db, SQL )
//	return CubeSQLPrepairedStatement(res)
//}	
//func (this *CubeSQLPrepairedStatement ) Select() CubeSQLCursor {
//	// CUBESQL_APIEXPORT csqlc		*cubesql_vmselect (csqlvm *vm);
//	return C.cubesql_vmselect( *this )
//}
//func (this *CubeSQLPrepairedStatement ) BindInt( index int, value int ) int {
//	// CUBESQL_APIEXPORT int		cubesql_vmbind_int (csqlvm *vm, int index, int value);
//
//	return int( C.cubesql_vmbind_int( this, C.int( index ), C.int( value ) ) )
//}



// CUBESQL_APIEXPORT int		cubesql_vmbind_double (csqlvm *vm, int index, double value);
// CUBESQL_APIEXPORT int		cubesql_vmbind_text (csqlvm *vm, int index, char *value, int len);
// CUBESQL_APIEXPORT int		cubesql_vmbind_blob (csqlvm *vm, int index, void *value, int len);
// CUBESQL_APIEXPORT int		cubesql_vmbind_null (csqlvm *vm, int index);
// CUBESQL_APIEXPORT int		cubesql_vmbind_int64 (csqlvm *vm, int index, int64 value);
// CUBESQL_APIEXPORT int		cubesql_vmbind_zeroblob (csqlvm *vm, int index, int len);
// CUBESQL_APIEXPORT int		cubesql_vmexecute (csqlvm *vm);

// CUBESQL_APIEXPORT int		cubesql_vmclose (csqlvm *vm);
