package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-test/deep"
	"github.com/int128/ghcp/adaptors"
)

type singleNameFilter struct {
	t    *testing.T
	dir  string // name of directory to skip (if empty, do nothing)
	file string // name of file to exclude (if empty, do nothing)
}

func (f *singleNameFilter) SkipDir(path string) bool {
	f.t.Logf("visiting dir %s", path)
	base := filepath.Base(path)
	return f.dir == base
}

func (f *singleNameFilter) ExcludeFile(path string) bool {
	f.t.Logf("visiting file %s", path)
	base := filepath.Base(path)
	return f.file == base
}

func TestFileSystem_FindFiles(t *testing.T) {
	fs := &FileSystem{}
	tempDir, err := ioutil.TempDir("", "FindFiles")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)
	if err := os.Chdir(tempDir); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir("dir1", 0755); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile("dir1/a.jpg", []byte{}, 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir("dir2", 0755); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile("dir2/b.jpg", []byte{}, 0644); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile("dir2/c.jpg", []byte{}, 0755); err != nil {
		t.Fatal(err)
	}

	t.Run("FindDirectory", func(t *testing.T) {
		files, err := fs.FindFiles([]string{"."}, &singleNameFilter{t: t})
		if err != nil {
			t.Fatalf("FindFiles returned error: %+v", err)
		}
		wants := []adaptors.File{
			{Path: "dir1/a.jpg"},
			{Path: "dir2/b.jpg"},
			{Path: "dir2/c.jpg", Executable: true},
		}
		if diff := deep.Equal(wants, files); diff != nil {
			t.Error(diff)
		}
	})
	t.Run("FindFiles", func(t *testing.T) {
		files, err := fs.FindFiles([]string{"dir1/a.jpg", "dir2/c.jpg"}, &singleNameFilter{t: t})
		if err != nil {
			t.Fatalf("FindFiles returned error: %+v", err)
		}
		wants := []adaptors.File{
			{Path: "dir1/a.jpg"},
			{Path: "dir2/c.jpg", Executable: true},
		}
		if diff := deep.Equal(wants, files); diff != nil {
			t.Error(diff)
		}
	})
	t.Run("NoSuchFile", func(t *testing.T) {
		files, err := fs.FindFiles([]string{"dir3"}, &singleNameFilter{t: t})
		if files != nil {
			t.Errorf("files wants nil but %+v", files)
		}
		if err == nil {
			t.Fatalf("err wants non-nil but nil")
		}
	})
	t.Run("ExcludeDirectory", func(t *testing.T) {
		files, err := fs.FindFiles([]string{"."}, &singleNameFilter{t: t, dir: "dir2"})
		if err != nil {
			t.Fatalf("FindFiles returned error: %+v", err)
		}
		wants := []adaptors.File{
			{Path: "dir1/a.jpg"},
		}
		if diff := deep.Equal(wants, files); diff != nil {
			t.Error(diff)
		}
	})
	t.Run("SkipFile", func(t *testing.T) {
		files, err := fs.FindFiles([]string{"."}, &singleNameFilter{t: t, file: "b.jpg"})
		if err != nil {
			t.Fatalf("FindFiles returned error: %+v", err)
		}
		wants := []adaptors.File{
			{Path: "dir1/a.jpg"},
			{Path: "dir2/c.jpg", Executable: true},
		}
		if diff := deep.Equal(wants, files); diff != nil {
			t.Error(diff)
		}
	})
}

func TestFileSystem_ReadAsBase64EncodedContent(t *testing.T) {
	fs := &FileSystem{}
	tempFile := makeTempFile(t, "hello\nworld")
	defer os.RemoveAll(tempFile)
	content, err := fs.ReadAsBase64EncodedContent(tempFile)
	if err != nil {
		t.Fatalf("ReadAsBase64EncodedContent returned error: %+v", err)
	}
	want := "aGVsbG8Kd29ybGQ="
	if want != content {
		t.Errorf("content wants %s but %s", want, content)
	}
}

func makeTempFile(t *testing.T, content string) string {
	tempFile, err := ioutil.TempFile("", "fs_test")
	if err != nil {
		t.Fatalf("error while creating a temp file: %s", err)
	}
	defer tempFile.Close()
	if _, err := fmt.Fprint(tempFile, content); err != nil {
		t.Fatalf("error while writing content: %s", err)
	}
	return tempFile.Name()
}
