package main

import "fmt"

type Trie struct {
	// 是否为空
	fill bool
	// 是否是单词
	word bool
	// 子节点
	nodes []Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var t Trie
	t.fill = true
	t.word = false
	t.nodes = make([]Trie, 26)
	return t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	t := this
	// 循环每个字母
	for _, x := range []byte(word) {
		// 索引是该字母的ascii-'a'
		idx := int(x) - int('a')
		// 如果不存在，则创建节点
		if !t.nodes[idx].fill {
			var n Trie
			n.fill = true
			n.word = false
			n.nodes = make([]Trie, 26)
			t.nodes[idx] = n
		}
		// 移动到该节点
		t = &t.nodes[idx]
	}
	// 到达末尾，把单词标志位置true
	t.word = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this
	for _, x := range []byte(word) {
		idx := int(x) - int('a')
		if !t.nodes[idx].fill {
			return false
		}
		t = &t.nodes[idx]
	}
	return t.word
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for _, x := range []byte(prefix) {
		idx := int(x) - int('a')
		if !t.nodes[idx].fill {
			return false
		}
		t = &t.nodes[idx]
	}
	return t.fill
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	obj := Constructor()
	obj.Insert("apple")
	// fmt.Printf("%#v\n", obj)
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("pple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	fmt.Println(obj.StartsWith("pple"))
}
