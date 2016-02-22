package runeTree

import(
	"testing"
)

func TestRuneTreeInsert(t *testing.T){

	cases := []struct {
		Word string
		ReturnEarly bool
		Expect bool
		Description string
	}{
		{
			Word: "testing",
			Description: "1st insert 'testing' into empty tree",
			ReturnEarly: false,
			Expect: false,
		},{
			Word: "test",
			Description: "1st insert 'test' into existing",
			ReturnEarly: false,
			Expect: false,
		},{
			Word: "test",
			Description: "2nd insert 'test' into existing",
			ReturnEarly: false,
			Expect: true,
		},{
			Word: "testing",
			Description: "2nd insert 'testing' into existing",
			ReturnEarly: false,
			Expect: true,
		},{
			Word: "zzztesting",
			Description: "search for 'zzztesting' in existing tree",
			ReturnEarly: true,
			Expect: false,
		},{
			Word: "zzztesting",
			Description: "2nd search for 'zzztesting' in existing tree",
			ReturnEarly: true,
			Expect: false,
		},
	}

	//make a new tree
	tree := &RuneTree{'-', &[]RuneTree{}}

	for cIndex, cItem := range cases {
		returned := RuneTreeInsert(tree, cItem.Word, cItem.ReturnEarly)

		if(returned != cItem.Expect){
			t.Errorf("RuneTree Test %[1]v: %[2]s. Expected %[3]t, but got: %[4]t", cIndex, cItem.Description, cItem.Expect, returned)
		}
	}

}