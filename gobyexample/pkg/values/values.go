// values : more about data types and features
package values

// Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
// strings are immutable, made of bytes / runes
// datatypes : int32, int64, float32, float64, strings, bytes and runes

/* no need to implement and check these
Size		Minimum Value				Maximum Value
8-bits	-	(2^7) = 128					2^7 - 1 = 127
16-bits	-	(2^15) = 32,767				2^15 - 1 = 32,767
32-bits	-	(2^31) = -2,147,483,648		2^31 - 1 = 2,147,483,647
64-bits	-	(2^63) ~= -(9 x 10^19)		2^63 - 1 ~= 9 x 10^19
*/

type GrpDatatype struct {
	S string
	I int32
	F float32
}

func Values_method() GrpDatatype {
	var dataset GrpDatatype = GrpDatatype{
		"hello",
		32,
		42.5,
	}

	return dataset
}
