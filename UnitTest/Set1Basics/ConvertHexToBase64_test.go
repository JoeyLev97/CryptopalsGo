package Test

import (
	 "testing"
	 set1Basics "../../set1Basics" 
)


func TestSum(t *testing.T) {
	total := set1Basics.DecodeHex("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if(total != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"){
		t.Errorf("Converted hex is incorrect")
	}
}