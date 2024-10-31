package algo

type Node struct {
	Val rune
	Child map[rune]*Node
}

func BuildTree(node *Node, content []rune, chooseIndex int) {
	if node == nil || len(content) == 0 || chooseIndex >= len(content) {
		return
	}
	if node.Child == nil {
		node.Child = make(map[rune]*Node)
	}
	letter := content[chooseIndex]
	nextNode, exist := node.Child[letter]
	if exist {
		BuildTree(nextNode, content, chooseIndex+1)
	} else {
		node.Child[letter] = &Node{Val: letter}
		BuildTree(node.Child[letter], content, chooseIndex+1)
	}
}