package fsvenues

import (
  "fmt"
	"os"
	"testing"
)

//func Test_GetVenue(t *testing.T) {
//	fs := NewFSVenuesClient(os.Getenv("FOURSQUARE_ID"), os.Getenv("FOURSQUARE_SECRET"))
//	params := make(map[string]string)
//	params["ll"] = "40.7,-74"
//	if v, e := fs.GetVenue(params); e == nil {
//		t.Log("Passed", v)
//	} else {
//		t.Error("Failed", e)
//	}
//}
//
//func Test_ExploreVenues(t *testing.T) {
//	fs := NewFSVenuesClient(os.Getenv("FOURSQUARE_ID"), os.Getenv("FOURSQUARE_SECRET"))
//	params := make(map[string]string)
//	params["ll"] = "40.7,-74"
//	if v, e := fs.ExploreVenues(params); e == nil {
//		t.Log("Passed", v)
//	} else {
//		t.Error("Failed", e)
//	}
//}

func Test_GetVenues(t *testing.T) {
	fs := NewFSVenuesClient(os.Getenv("FOURSQUARE_ID"), os.Getenv("FOURSQUARE_SECRET"))
	params := make(map[string]string)
	params["ll"] = "40.7,-74"
	v, e := fs.GetVenues(params)

  fmt.Println(v)

  if e != nil {
		t.Error("Failed", e)
	} else {
    if len(v) < 1 {
      t.Error("Call returned amount: ", len(v))
    } else{
		  t.Log("Passed", v)
    }
	}
}
