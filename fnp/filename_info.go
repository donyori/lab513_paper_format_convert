package fnp

type FilenameInfo interface {
	Directory() string
	Number() (number int32, isSupported bool)
	NumberRange() (numberRange [2]int32, isSupported bool)
	Title() (title string, isSupported bool)
}
