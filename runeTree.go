package runeTree

//no imports!

//Stores words in a branching RuneTree, similar to a binary search tree, except there is one branch per rune
//Returns true if the word exists in the tree, false if it does not
//For most efficient results words should be lowercase and contain no symbols
func RuneTreeInsert(tree *RuneTree, word string, returnEarly bool) bool{
    
    word = word + "\n"
    var currentBranch *RuneTree = tree

    if currentBranch == nil {
        currentBranch = &RuneTree{'-', &[]RuneTree{}}
    }
    
    //loop through the word and travel down the tree branches
    foundRune := false
    for _, thisRune := range word {
        foundRune = false

        for _, searchBranch := range *currentBranch.Branches {
            if searchBranch.Value == thisRune{
                foundRune = true
                currentBranch = &searchBranch
                break;
            }
        }

        //if the last branch was not found, return early, or start adding a new branch
        if(!foundRune){
            if returnEarly{
                return false;
            } else {
                newBanch := RuneTree{thisRune, &[]RuneTree{}}
                *currentBranch.Branches = append(*currentBranch.Branches, newBanch)
                currentBranch = &newBanch
            }

            foundRune = false
        }
    }

    //finally return the result of the last search
    return foundRune
}

type RuneTree struct{
    Value rune
    Branches *[]RuneTree
}
