package main

// go build -o cubeSQL main.go

////////////// alles weitere da unten geht nicht/will man nicht = dynamic

// gcc -fPIC -shared -Isdk -Isdk/zlib -Isdk/crypt cubesql.c -o cubesql.so
// gcc -fPIC -shared -Isdk -Isdk/zlib -Isdk/crypt sdk/zlib/zlib.c sdk/crypt/*.c cubesql.c -o libcubesql.so

import "cubesql"


func main() {
  println( "Hallo" )

	db := newCubeSQL()
	result := db.connect( "127.0.0.1", 4430, "demo", "demo", 10, 0 );
	println( result )
	time.Sleep( 10 * time.Second )
	db.disconnect( 0 );

}

