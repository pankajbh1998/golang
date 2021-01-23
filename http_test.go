package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestH(t *testing.T){
	testCases:=[]struct{
		id string
		expectedOutput string
		statusCodes int
	}{
		{"-1","Hello Negative",100},
		{"0","Hello",100},
		{"1","Hello Positive",100},
		{"2","Not Found",500},
		{"100","Not Found",500},
		{"abc","Invalid",400},
		{"abc","Invalid",400},
	}
	for i,tc:=range testCases {
		w:=httptest.NewRecorder()
		//link:="/new?id=%v"
		//r:=httptest.NewRequest("GET",fmt.Sprintf(link,tc.id),nil)
		r:=httptest.NewRequest("GET","/new?id="+tc.id,nil)
		SayHello(w,r)
		res:=w.Result()
		readbyte,err:=ioutil.ReadAll(res.Body)
		//t.Logf(string(readbyte))
		if err!=nil {
			t.Fatal(err)
		}else if tc.expectedOutput!=string(readbyte){
			t.Fatalf("Failed at %d\nExpected : %s\nFound : %s\n",i+1,tc.expectedOutput,string(readbyte))
		} else if tc.statusCodes!=res.StatusCode {
			t.Fatalf("Status Codes are not matching at TestCase : %d\n",i+1)
		}
		t.Logf("Passed at %d",i+1)
	}
}