package copy

import (
	"io"
	"os"
)

// File copies the contents of the file named sourceFile to the file named
// by dstFile. The file will be created if it does not already exist. If the
// destination file exists, all its contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
func File(sourceFile, dstFile string) (err error) {
	inFilePointer, err := os.Open(sourceFile)
	if err != nil {
		return
	}
	defer inFilePointer.Close()

	outFilePointer, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer func() {
		if e := outFilePointer.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(outFilePointer, inFilePointer)
	if err != nil {
		return
	}

	err = outFilePointer.Sync()
	if err != nil {
		return
	}

	fileInfo, err := os.Stat(sourceFile)
	if err != nil {
		return
	}
	err = os.Chmod(dstFile, fileInfo.Mode())
	if err != nil {
		return
	}

	return
}
