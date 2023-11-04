package support

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"text/tabwriter"

	"vaughany.com/advent_of_code_go/internal/output"
)

// https://gist.github.com/clarkmcc/1fdab4472283bb68464d066d6b4169bc?permalink_comment_id=4405804#gistcomment-4405804
func GetAllFilenames(efs *embed.FS) error {
	var files []string

	err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		files = append(files, fmt.Sprintf("%s\t%d", path[7:], info.Size()))

		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("Files:", len(files))

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	for _, file := range files {
		fmt.Fprintln(tw, file)
	}
	tw.Flush()

	return nil
}

func GetPuzzleInput(year, day int, session string) error {
	client := &http.Client{}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	output.Info(fmt.Sprintf("Fetching %s...", url))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("expected 200 OK, received %d %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Make the inputs folder for the specified year, if it doesn't exist.
	err = os.MkdirAll(fmt.Sprintf("./cmd/%d/inputs", year), os.ModePerm)
	if err != nil {
		return err
	}

	// This path references the path within the code repository and will not work if using a compiled binary.
	filename := fmt.Sprintf("./cmd/%d/inputs/%.2d.txt", year, day)
	err = os.WriteFile(filename, bodyBytes, 0664)
	if err != nil {
		return err
	}

	output.Info(fmt.Sprintf("%d bytes written to %s.", len(bodyBytes), filename))

	return nil
}

// ServeInputs creates a http web server on the supplied address, and serves the inputs from the supplied embedded file system.
func ServeInputs(efs *embed.FS, address string) error {
	serverRoot, err := fs.Sub(efs, "inputs")
	http.Handle("/", http.FileServer(http.FS(serverRoot)))

	output.Info("Running web server on http://" + address)

	err = http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}

	return nil
}
