package logger

// Logger is the interface that should be implemented by a logging library, or that should be accepted when receiving a
// logger in your code, (usually a library).
//
// There are implementations of this such as StdLogger and NilLogger within this github.com/gologger project.
type Logger interface {
	// Fatal is equivalent to Print() followed by a call to os.Exit(1).
	Fatal(v ...interface{})
	// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
	Fatalf(format string, v ...interface{})
	// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
	Fatalln(v ...interface{})
	// Panic is equivalent to Print() followed by a call to panic().
	Panic(v ...interface{})
	// Panicf is equivalent to Printf() followed by a call to panic().
	Panicf(format string, v ...interface{})
	// Panicln is equivalent to Println() followed by a call to panic().
	Panicln(v ...interface{})
	// Print outputs to the underlying concrete logger. Arguments are handled as specified by that logger.
	Print(v ...interface{})
	// Printf outputs to the underlying concrete logger. Arguments are handled as specified by that logger.
	Printf(format string, v ...interface{})
	// Println outputs to the underlying concrete logger. Arguments are handled as specified by that logger.
	Println(v ...interface{})
}
