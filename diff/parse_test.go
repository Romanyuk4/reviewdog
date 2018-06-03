	"encoding/json"
//go:generate go run testdata/gen.go


		wantfile, err := os.Open(fname + ".json")
		if err != nil {
			t.Fatal(err)
		}
		dec := json.NewDecoder(wantfile)
		var want []*FileDiff
		if err := dec.Decode(&want); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(difffiles, want) {
			l := len(difffiles)
			if len(want) > l {
				l = len(want)
			}
			for i := 0; i < l; i++ {
				var a, b *FileDiff
				if i < len(want) {
					a = want[i]
				}
				if i < len(difffiles) {
					b = difffiles[i]
				}
				t.Errorf("want %#v, got %#v", a, b)
			}

		wantfile.Close()