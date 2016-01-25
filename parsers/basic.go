package parsers

import (
	"bufio"
	. "github.com/ruxton/mix_cover_builder/data"
	"github.com/ruxton/term"
	"io"
	"os"
  "regexp"
)

func ParseBasicTracklist(bufReader *bufio.Reader) []Track {
	var list []Track

  regex,err := regexp.Compile(`([0-9]+)(?:.{1}?\s)(.+)(?:\s-\s)(.+)`)
  if(err != nil) {
    term.OutputError("Error matching strings")
    os.Exit(2)
  }

	for line, _, err := bufReader.ReadLine(); err != io.EOF; line, _, err = bufReader.ReadLine() {

    trackdata := regex.FindAllStringSubmatch(string(line),3)[0]
		// trackdata := strings.SplitN(data[1], " - ", 2)
    // t
		if len(trackdata) != 4 {
			term.OutputError("Error parsing track " + string(line))
			term.OutputMessage("Please enter an artist for this track: ")
			artist, err := term.STD_IN.ReadString('\n')
			if err != nil {
				term.OutputError("Incorrect artist entry.")
				os.Exit(2)
			}
			term.OutputMessage("Please enter a name for this track: ")
			track, err := term.STD_IN.ReadString('\n')
			if err != nil {
				term.OutputError("Incorrect track name entry.")
				os.Exit(2)
			}

			trackdata = []string{"0", artist, track}
		}

		thistrack := new(Track)
		thistrack.Artist = trackdata[2]
		thistrack.Song = trackdata[3]

		list = append(list, *thistrack)
	}

	return list
}
