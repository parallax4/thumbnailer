package thumbnailer

import "testing"

func TestMatchMP4(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		data         []byte
		expectedMime string
		expectedExt  string
	}{
		{
			name: "mp41 brand",
			// ftyp box with mp41 brand (like cat.mp4)
			data:         []byte{0x00, 0x00, 0x00, 0x20, 'f', 't', 'y', 'p', 'i', 's', 'o', 'm', 0x00, 0x00, 0x02, 0x00, 'i', 's', 'o', 'm', 'i', 's', 'o', '2', 'a', 'v', 'c', '1', 'm', 'p', '4', '1'},
			expectedMime: "video/mp4",
			expectedExt:  "mp4",
		},
		{
			name: "dash brand",
			// ftyp box with dash brand (like rare_brand.mp4)
			data:         []byte{0x00, 0x00, 0x00, 0x18, 'f', 't', 'y', 'p', 'i', 's', 'o', '5', 0x00, 0x00, 0x00, 0x01, 'a', 'v', 'c', '1', 'd', 'a', 's', 'h'},
			expectedMime: "video/mp4",
			expectedExt:  "mp4",
		},
		{
			name: "isom only - Twitter style",
			// ftyp box with only isom/iso4 brands (like thiel.mp4 from Twitter)
			data:         []byte{0x00, 0x00, 0x00, 0x18, 'f', 't', 'y', 'p', 'i', 's', 'o', 'm', 0x00, 0x00, 0x00, 0x01, 'i', 's', 'o', 'm', 'i', 's', 'o', '4'},
			expectedMime: "video/mp4",
			expectedExt:  "mp4",
		},
		{
			name: "iso5 only",
			// ftyp box with only iso5 brand
			data:         []byte{0x00, 0x00, 0x00, 0x14, 'f', 't', 'y', 'p', 'i', 's', 'o', '5', 0x00, 0x00, 0x00, 0x01, 'i', 's', 'o', '5'},
			expectedMime: "video/mp4",
			expectedExt:  "mp4",
		},
		{
			name:         "not mp4 - no ftyp",
			data:         []byte{0x00, 0x00, 0x00, 0x18, 'n', 'o', 't', 'f', 'i', 's', 'o', 'm'},
			expectedMime: "",
			expectedExt:  "",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mime, ext := matchMP4(tc.data)
			if mime != tc.expectedMime {
				t.Errorf("expected mime %q, got %q", tc.expectedMime, mime)
			}
			if ext != tc.expectedExt {
				t.Errorf("expected ext %q, got %q", tc.expectedExt, ext)
			}
		})
	}
}
