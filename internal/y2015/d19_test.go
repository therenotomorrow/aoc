package y2015

import "testing"

func TestDay19Part1(t *testing.T) {
	type args struct {
		replacements []string
		molecule     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{replacements: []string{"H => HO", "H => OH", "O => HH"}, molecule: "HOH"}, want: 4},
		{name: "smoke 2", args: args{replacements: []string{"H => HO", "H => OH", "O => HH"}, molecule: "HOHOHO"}, want: 7},
		{name: "answer", args: args{replacements: []string{"Al => ThF", "Al => ThRnFAr", "B => BCa", "B => TiB", "B => TiRnFAr", "Ca => CaCa", "Ca => PB", "Ca => PRnFAr", "Ca => SiRnFYFAr", "Ca => SiRnMgAr", "Ca => SiTh", "F => CaF", "F => PMg", "F => SiAl", "H => CRnAlAr", "H => CRnFYFYFAr", "H => CRnFYMgAr", "H => CRnMgYFAr", "H => HCa", "H => NRnFYFAr", "H => NRnMgAr", "H => NTh", "H => OB", "H => ORnFAr", "Mg => BF", "Mg => TiMg", "N => CRnFAr", "N => HSi", "O => CRnFYFAr", "O => CRnMgAr", "O => HP", "O => NRnFAr", "O => OTi", "P => CaP", "P => PTi", "P => SiRnFAr", "Si => CaSi", "Th => ThCa", "Ti => BP", "Ti => TiTi", "e => HF", "e => NAl", "e => OMg"}, molecule: "ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF"}, want: 576},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day19Part1(tt.args.replacements, tt.args.molecule); got != tt.want {
				t.Errorf("Day19Part1() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDay19Part2(t *testing.T) {
	type args struct {
		replacements []string
		molecule     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{replacements: []string{"H => HO", "H => OH", "O => HH", "e => H", "e => O"}, molecule: "HOH"}, want: 3},
		{name: "smoke 2", args: args{replacements: []string{"H => HO", "H => OH", "O => HH", "e => H", "e => O"}, molecule: "HOHOHO"}, want: 6},
		{name: "answer", args: args{replacements: []string{"Al => ThF", "Al => ThRnFAr", "B => BCa", "B => TiB", "B => TiRnFAr", "Ca => CaCa", "Ca => PB", "Ca => PRnFAr", "Ca => SiRnFYFAr", "Ca => SiRnMgAr", "Ca => SiTh", "F => CaF", "F => PMg", "F => SiAl", "H => CRnAlAr", "H => CRnFYFYFAr", "H => CRnFYMgAr", "H => CRnMgYFAr", "H => HCa", "H => NRnFYFAr", "H => NRnMgAr", "H => NTh", "H => OB", "H => ORnFAr", "Mg => BF", "Mg => TiMg", "N => CRnFAr", "N => HSi", "O => CRnFYFAr", "O => CRnMgAr", "O => HP", "O => NRnFAr", "O => OTi", "P => CaP", "P => PTi", "P => SiRnFAr", "Si => CaSi", "Th => ThCa", "Ti => BP", "Ti => TiTi", "e => HF", "e => NAl", "e => OMg"}, molecule: "ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF"}, want: 207},
		{name: "impossible", args: args{replacements: []string{"e => H"}, molecule: "e"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day19Part2(tt.args.replacements, tt.args.molecule); got != tt.want {
				t.Errorf("Day19Part2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
