package main

import "testing"

func Test_eval(t *testing.T) {
	tests := []struct {
		name      string
		expected  node
		totalSize int
		input     []string
	}{
		{
			"golden",
			node{
				children: map[string]node{
					"a": {children: map[string]node{
						"e": {children: map[string]node{
							"i": {fileSize: 584}}},
						"f":     {fileSize: 29116},
						"g":     {fileSize: 2557},
						"h.lst": {fileSize: 62596},
					}},
					"b.txt": {fileSize: 14848514},
					"c.dat": {fileSize: 8504156},
					"d": {children: map[string]node{
						"j":     {fileSize: 4060174},
						"d.log": {fileSize: 8033020},
						"d.ext": {fileSize: 5626152},
						"k":     {fileSize: 7214296}}},
				},
			},
			25,
			[]string{
				"$ ls",
				"dir a",
				"14848514 b.txt",
				"8504156 c.dat",
				"dir d",
				"$ cd a",
				"$ ls",
				"dir e",
				"29116 f",
				"2557 g",
				"62596 h.lst",
				"$ cd e",
				"$ ls",
				"584 i",
				"$ cd ..",
				"$ cd ..",
				"$ cd d",
				"$ ls",
				"4060174 j",
				"8033020 d.log",
				"5626152 d.ext",
				"7214296 k",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stream := make(chan string)
			go func() {
				for _, line := range test.input {
					stream <- line
				}
				close(stream)
			}()

			n := node{children: make(map[string]node)}
			n.eval(stream)
			testChild(n, test.expected, t)
		})
	}
}

func Test_size(t *testing.T) {
	tests := []struct {
		name      string
		expected  node
		totalSize int
		input     []string
	}{
		{
			"single child",
			node{
				children: map[string]node{
					"a": {fileSize: 10},
				},
			},
			10,
			[]string{"10 a"},
		},
		{
			"two children",
			node{
				children: map[string]node{
					"a": {fileSize: 10},
					"b": {fileSize: 15},
				},
			},
			25,
			[]string{
				"10 a",
				"15 b",
			},
		},
		{
			"dupes",
			node{
				children: map[string]node{
					"a": {fileSize: 10},
				},
			},
			10,
			[]string{
				"10 a",
				"10 a",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stream := make(chan string)
			go func() {
				for _, line := range test.input {
					stream <- line
				}
				close(stream)
			}()

			n := node{children: make(map[string]node)}
			n.eval(stream)
			testChild(n, test.expected, t)
			if n.size() != test.totalSize {
				t.Errorf("expected: %+v\ngot: %+v", test.expected, n)
			}
		})
	}
}

func testChild(n node, expected node, t *testing.T) {
	for name, child := range n.children {
		expectedChild, ok := expected.children[name]
		if !ok {
			t.Errorf("child %v not found", name)
		}
		if child.fileSize != expectedChild.fileSize {
			t.Errorf("file %v expected: %v got:%v", name, expectedChild.fileSize, child.fileSize)
		}
		testChild(child, expectedChild, t)
	}
}

func Test_flatten(t *testing.T) {
	tests := []struct {
		name     string
		input    node
		expected map[string]node
	}{
		{
			name: "golden",
			input: node{
				children: map[string]node{
					"a": {children: map[string]node{
						"e": {children: map[string]node{
							"i": {fileSize: 584}}},
						"f":     {fileSize: 29116},
						"g":     {fileSize: 2557},
						"h.lst": {fileSize: 62596},
					}},
					"b.txt": {fileSize: 14848514},
					"c.dat": {fileSize: 8504156},
					"d": {children: map[string]node{
						"j":     {fileSize: 4060174},
						"d.log": {fileSize: 8033020},
						"d.ext": {fileSize: 5626152},
						"k":     {fileSize: 7214296}}},
				},
			},
			expected: map[string]node{
				"/": {children: map[string]node{
					"a": {children: map[string]node{
						"e": {children: map[string]node{
							"i": {fileSize: 584}}},
						"f":     {fileSize: 29116},
						"g":     {fileSize: 2557},
						"h.lst": {fileSize: 62596},
					}},
					"b.txt": {fileSize: 14848514},
					"c.dat": {fileSize: 8504156},
					"d": {children: map[string]node{
						"j":     {fileSize: 4060174},
						"d.log": {fileSize: 8033020},
						"d.ext": {fileSize: 5626152},
						"k":     {fileSize: 7214296}}},
				}},
				"/a": {children: map[string]node{
					"e": {children: map[string]node{
						"i": {fileSize: 584}}},
					"f":     {fileSize: 29116},
					"g":     {fileSize: 2557},
					"h.lst": {fileSize: 62596},
				}},
				"/a/e": {children: map[string]node{
					"i": {fileSize: 584}}},
				"/d": {children: map[string]node{
					"j":     {fileSize: 4060174},
					"d.log": {fileSize: 8033020},
					"d.ext": {fileSize: 5626152},
					"k":     {fileSize: 7214296}}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := test.input.flatten("/")
			if len(output) != len(test.expected) {
				t.Errorf("expected len: %v actual: %v", len(test.expected), len(output))
			}
		})
	}
}

func Test_sum(t *testing.T) {
	tests := []struct {
		name     string
		input    node
		expected int
	}{
		{
			name: "golden",
			input: node{
				children: map[string]node{
					"a": {children: map[string]node{
						"e": {children: map[string]node{
							"i": {fileSize: 584}}},
						"f":     {fileSize: 29116},
						"g":     {fileSize: 2557},
						"h.lst": {fileSize: 62596},
					}},
					"b.txt": {fileSize: 14848514},
					"c.dat": {fileSize: 8504156},
					"d": {children: map[string]node{
						"j":     {fileSize: 4060174},
						"d.log": {fileSize: 8033020},
						"d.ext": {fileSize: 5626152},
						"k":     {fileSize: 7214296}}},
				},
			},
			expected: 95437,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := test.input.sum(100000)
			if output != test.expected {
				t.Errorf("expected len: %v actual: %v", output, test.expected)
			}
		})
	}
}
