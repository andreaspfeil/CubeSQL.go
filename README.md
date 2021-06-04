# cubeSQL golang client

Bug free, feature complete, uses Marco Bambinis nativ C SDK.

## Usage example

```golang
package main

import "cubesql"

func main() {
	cube := cubesql.New()
	result := cube.connect( "localhost", 4430, "loginname", "password", 10, 0 );
	println( result )
	time.Sleep( 10 * time.Second )
	cube.disconnect( 0 );
}
```

## Installation

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
