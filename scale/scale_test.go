package scale

import "testing"

func TestAllScales(t *testing.T) {
	if len(allScales()) != 29 {
		t.Error("Wrong number of all scales.")
	}

	for _, v := range AllScales {
		if v == "mixolydian" {
			return
		}
	}
	t.Error("AllChords doesn't have `mixolydian`.")
}

func TestGetScale(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "ionian",
			want: []int{0, 2, 4, 5, 7, 9, 11},
		},
		{
			name: "Ionian",
			want: []int{0, 2, 4, 5, 7, 9, 11},
		},
		{
			name: "super-locrianb7",
			want: []int{0, 1, 3, 4, 6, 8, 9},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, _ := GetScale(test.name);
			if len(test.want) != len(actual) {
				t.Errorf(`GetScale("%v"), actual:"%v", want:"%v"`, test.name, actual, test.want)
			}
			for i, v := range actual {
				if test.want[i] != v {
					t.Errorf(`GetScale("%v"), actual:"%v", want:"%v"`, test.name, actual, test.want)
				}
			}
		})
	}
}

func TestGetScaleError(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{
			name: "notfoundian",
			want: ErrorNotFoundScale("notfoundian"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetScale(test.name);
			if len(got) != 0 {
				t.Errorf(`GetScale("%v") wants empty result. But got (%v).`, test.name, got)
			}
			if err.Error() != test.want.Error() {
				t.Errorf(`GetScale("%v") wants Error(%v). but it's wrong. "%v"`, test.name, err, test.want)
			}
		})
	}
}

func TestGetScaleFromRoot(t *testing.T) {
	tests := []struct {
		name string
		root string
		want []int
	}{
		{
			name: "ionian",
			root: "D4",
			want: []int{62, 64, 66, 67, 69, 71, 73},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, _ := GetScaleFromRoot(test.name, test.root);
			if len(test.want) != len(actual) {
				t.Errorf(`GetScaleFromRoot("%v"), actual:"%v", want:"%v"`, test.name, actual, test.want)
			}
			for i, v := range actual {
				if test.want[i] != v {
					t.Errorf(`GetScaleFromRoot("%v"), actual:"%v", want:"%v"`, test.name, actual, test.want)
				}
			}
		})
	}
}
