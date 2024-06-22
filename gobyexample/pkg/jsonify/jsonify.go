package jsonify

// what can I do and test
// .. json vs bson ->
// .. marshal and unmarshall a json -> accessing values of a json object

// .. create a json store
//    input a json -> validate json -> append in jsonstore
// 	  input a json -> segregate json based on validation
// .. get input of json and validate it with a predefined json structure and datatype check
/*
	JSON VALIDATION
	...
	A JSON validator parses the JSON text and checks for the following:

	Syntax Errors: Ensures the JSON structure is valid, such as proper use of braces {}, brackets [], commas ,, colons :, and quotes ".

	Structural Integrity: Ensures that arrays and objects are properly formed and nested.

	Value Types: Checks that values are of acceptable types, like strings, numbers, objects, arrays, true, false, or null.

	Key-Value Pairs: Ensures that object keys are strings enclosed in double quotes and each key has an associated value.
*/

// what more can we add
// .. custom json validation and check
// .. reference json and predefined data type check

// what should a json look like
/*
 * JSON Definition Rules:
 * 1. Enclosing Structure:
 *    - JSON must start and end with either curly braces `{}` for an object or square brackets `[]` for an array.
 *    - Each opening bracket or brace must have a corresponding closing bracket or brace.
 * 2. Key-Value Pairs:
 *    - Objects are collections of key-value pairs.
 *    - Keys must be strings enclosed in double quotes `""`.
 *    - Values can be strings, numbers, arrays, objects, `true`, `false`, or `null`.
 * 3. Strings:
 *    - String values must be enclosed in double quotes `""`.
 *    - Keys in JSON must also be strings, enclosed in double quotes `""`.
 * 4. Commas:
 *    - Key-value pairs within an object are separated by commas `,`.
 *    - Elements within an array are separated by commas `,`.
 *    - The last element in an array or the last key-value pair in an object should not have a trailing comma.
 * 5. Colon Separation:
 *    - Keys and values within an object are separated by a colon `:`.
 * 6. Array and Object Structure:
 *    - Arrays are ordered collections of values, enclosed in square brackets `[]`.
 *    - Objects are unordered collections of key-value pairs, enclosed in curly braces `{}`.
 * 7. Whitespace:
 *    - JSON can contain whitespace (spaces, tabs, and newlines), which is ignored by parsers and can be used to make the JSON more readable.
 * 8. Special Characters:
 *    - JSON strings must not contain unescaped special characters. Common escapes include `\n` for newline, `\t` for tab, `\"` for double quotes, and `\\` for backslash.
 * 9. Numeric Values:
 *    - Numbers in JSON must not be quoted. They can be integers or floating-point numbers.
 * 10. Boolean and Null:
 *    - Boolean values (`true` and `false`) and `null` must not be quoted.
 */
