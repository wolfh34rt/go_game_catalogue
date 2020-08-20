package string_generator

import(
    "math/rand"
    _ "fmt"
    _ "strings"
    _ "io/ioutil"
    _ "sort"
    "bytes"
    "time"
)

const(
    DEFAULT_RANDOM_ID_LENGTH = 30
    ASCII_SUBTRACTOR = 47
    ASCII_CHAR_LIMIT = 123
    ASCII_SPECIAL_CHARS_SET_1_LOWER_LIMIT = 58
    ASCII_SPECIAL_CHARS_SET_1_UPPER_LIMIT = 64
    ASCII_SPECIAL_CHARS_SET_2_LOWER_LIMIT = 91
    ASCII_SPECIAL_CHARS_SET_2_UPPER_LIMIT = 96
)

func RandomId() string {
    var buffer bytes.Buffer

    for i := 0 ; i < ( DEFAULT_RANDOM_ID_LENGTH + 1 ) ; i++ {
        for {
            random_number := random_int()

            if random_number > ASCII_SUBTRACTOR && random_number < ASCII_CHAR_LIMIT {
                if ( random_number < ASCII_SPECIAL_CHARS_SET_1_LOWER_LIMIT || random_number > ASCII_SPECIAL_CHARS_SET_1_UPPER_LIMIT ) && ( random_number < ASCII_SPECIAL_CHARS_SET_2_LOWER_LIMIT || random_number > ASCII_SPECIAL_CHARS_SET_2_UPPER_LIMIT ) {
                    buffer.WriteString( string( random_number ) )
                
                    break;
                }
            } 
        }
    }

    return buffer.String()
}

func random_int() int {
    rand.Seed( time.Now().UnixNano() )
    return rand.Intn( ( ASCII_CHAR_LIMIT - 1 ) - ( ASCII_SUBTRACTOR + 1 ) ) + ( ASCII_SUBTRACTOR + 1 )
}