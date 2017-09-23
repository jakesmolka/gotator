package gotator

import (
	"encoding/json"
	"testing"
)

type newClientTest struct {
	url      string
	accepted bool
}

var newClientCases = []newClientTest{
	{ // empty url, shall not be accepted
		"",
		false,
	},
	{ // valid url, shall be accepted
		"https://www.ncbi.nlm.nih.gov/CBBresearch/Lu/Demo/RESTful/tmTool.cgi",
		true,
	},
}

func TestNewClient(t *testing.T) {
	for _, tc := range newClientCases {
		c := new(Client)
		c, err := c.NewClient(tc.url)
		if err != nil && tc.accepted == true {
			t.Fatalf("expected url to be accepted: %v, err: %v", tc.url, err)
		}

		if err == nil && tc.accepted == false {
			t.Fatalf("expected url not to be accepted: %v, err: %v", tc.url, err)
		}
	}
}

type getAllAnnotationsTest struct {
	url string
	id  string
	res []byte
}

var gAACases = []getAllAnnotationsTest{
	{ // no id, no result
		"https://www.ncbi.nlm.nih.gov/CBBresearch/Lu/Demo/RESTful/tmTool.cgi",
		"",
		[]byte(""),
	}, { // valid id, valid result
		"https://www.ncbi.nlm.nih.gov/CBBresearch/Lu/Demo/RESTful/tmTool.cgi",
		"19894120",
		[]byte(`{
			
				"sourcedb": "PubMed",
				"sourceid": "19894120",
				"text": "Lipopolysaccharide increases the expression of multidrug resistance-associated protein 1 (MRP1) in RAW 264.7 macrophages. Multidrug resistance-associated protein 1 (MRP-1) is a ubiquitously expressed member of the ATP-binding cassette transporter family. MRP-1 is one of the primary transporters of glutathione and glutathione conjugates. This protein also transports antiretroviral therapeutics, such as HIV-1 protease inhibitors (PI). We hypothesized that inflammatory mediators that activate macrophages would modify the expression and activity of MRP-1 in macrophages. Real-time PCR assays, western blots, and calcein efflux assays were used to show that exposure of macrophage cell line RAW 264.7 to lipopolysaccharide (LPS) increased expression of MRP-1 at the levels of mRNA, protein, and functional activity. Treatment of macrophages with LPS resulted in 2-fold increases of MRP-1 expression or functional activity. LPS-mediated increases in calcein efflux were repressed by the MRP-specific inhibitor MK-571. These results suggest that the effectiveness of HIV-1 PI therapy may be compromised by the presence of opportunistic infections.",
				"denotations": [
					{
						"obj": "Chemical:C059141",
						"span": {
							"begin": 1010,
							"end": 1016
						}
					},
					{
						"obj": "Chemical:C007740",
						"span": {
							"begin": 614,
							"end": 621
						}
					},
					{
						"obj": "Chemical:D005978",
						"span": {
							"begin": 315,
							"end": 326
						}
					},
					{
						"obj": "Chemical:D005978",
						"span": {
							"begin": 299,
							"end": 310
						}
					},
					{
						"obj": "Chemical:D000255",
						"span": {
							"begin": 214,
							"end": 217
						}
					},
					{
						"obj": "Chemical:CHEBI:16412",
						"span": {
							"begin": 0,
							"end": 18
						}
					}
				]
			
			}`),
	}, //{ // TODO: add failing test cases for coverage. right now everything passes.
	//"https://www.ncbi.nlm.nih.gov/CBBresearch/Lu/Demo/RESTful/tmTool.cgi",
	//"www.gibberish.com", //should make Parse(), NewRequest() or Do() fail
	//[]byte(""),
	//},
}

func TestGetAllAnnotations(t *testing.T) {
	for _, tc := range gAACases {
		c := new(Client)
		c, _ = c.NewClient(tc.url)

		art, err := c.GetAllAnnotations(tc.id)
		if err != nil {
			t.Logf("got error that's tolerable in some cases for testing: %v, with url: %v", err, tc.url)
		}

		//log.Printf("article: %v", art)
		//if art == nil {
		//	t.Fatal("expected response, got none")
		//}

		// test if both are empty, if not test equalness
		if len(tc.res) == 0 {
			if art != nil {
				t.Fatal("tc.res is empty and so should art")
			}
		} else {
			var testArt Article
			err = json.Unmarshal(tc.res, &testArt)
			if err != nil {
				t.Fatalf("test case res isn't valid: %v", tc.res)
			}
			// TODO DeepEquals seems not to work when order of json objects are different
			//		so just .Text gets compared right now
			//if !reflect.DeepEqual(art, &testArt) {
			//	t.Fatalf("expected %v, \n got %v", art, &testArt)
			//}
			if art.Text != testArt.Text {
				t.Fatalf("expected same attribute content")
			}
		}
	}
}
