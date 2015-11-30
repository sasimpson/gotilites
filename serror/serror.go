package serror

import "log"

//CheckError - does the normal go error checking, but has a
// test for failing if the error occurs.
func CheckError(fail bool, err error) {
	if err != nil {
		if fail {
			log.Fatalln(err)
		} else {
			log.Println(err)
		}
	}
}

//FailError - checks for error and fails if there is an error
func FailError(err error) {
	CheckError(true, err)
}

//LogError - checks for error, logs it, continues on
func LogError(err error) {
	CheckError(false, err)
}
