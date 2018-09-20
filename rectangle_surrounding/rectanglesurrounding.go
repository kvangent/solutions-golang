package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// [Metadata]
// Title: The Rectangles Are Surrounding Us!
// URL: https://open.kattis.com/problems/rectanglesurrounding
// Categories: geometry
// Difficulty: 3.6

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)
func scanf(f string, a ...interface{}) { fmt.Fscanf(in, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(out, f, a...) }

// Represents of start/stop of rectangle from left to right
type Seg struct {
	y, left, right int
	isStart bool
}

type Node struct {
	start, end int
	next *Node
}

func main(){
	defer out.Flush()

	for {
		// Parse input
		var n int
		scanf("%d\n", &n)
		 if n == 0 { break }
		segments := make([]Seg, 0, n+n)
		for i := 0; i < n; i++ {
			var left, bot, right, top int
			scanf("%d %d %d %d\n", &left, &bot, &right, &top)
			segments = append(segments, Seg{bot, left, right, true})  // Rectangle starts
			segments = append(segments, Seg{top, left, right, false}) // Rectangle ends
		}

		// Sort segments by y-axis (ascending)
		sort.Slice(segments, func(i, j int) bool {
			if segments[i].y == segments[j].y {
				// (starts first in case of 0 length rectangles)
				return segments[i].isStart
			}
			return segments[i].y < segments[j].y
		})

		// Process Segments
		area, lastLoc := 0, 0
		var covered *Node // Linked List representing coverage of segments
		for len(segments) > 0 {
			curLoc := segments[0].y
			// Add area for lastLoc to curLoc
			l, w := curLoc - lastLoc, calcWidth(covered)
			area += l*w
			// Add or remove any changes in coverage
			for len(segments) > 0 && curLoc == segments[0].y {
				var seg *Seg
				seg, segments = &segments[0], segments[1:]
				s := &Node{start: seg.left, end: seg.right}
				if seg.isStart { // mark segment as covered
					covered = listAdd(covered, s)
				} else { // unmark segment as covered
					covered = listRmv(covered, s)
				}
			}
			lastLoc = curLoc
		}
		printf("%d\n", area)
	}
}

// Adds Node{start, end} to list in sorted order
func listAdd(cur* Node, target *Node) *Node {
	if cur == nil ||  target.start < cur.start {
		// Insert target before cur
		target.next = cur
		return target
	} // otherwise Node is added farther down the list
	cur.next = listAdd(cur.next, target)
	return cur
}

// Removes Node{start, end} from list
func listRmv(cur *Node, target *Node) *Node {
	if cur == nil || cur.equals(target){
		return cur.next // remove current node
	} // otherwise Node must be removed farther down the list
	cur.next = listRmv(cur.next, target)
	return cur
}

// Returns true if i == j, otherwise false
func (i *Node) equals(j *Node) bool{
	return i.start == j.start && i.end == j.end
}

// Returns the total width covered by the list
func calcWidth(cur *Node) int {
	if cur == nil {
		return 0 // empty list
	}
	start, end := cur.start, cur.end
	// If next segment overlaps current segment, combine
	for cur.next != nil && cur.next.start < end {
		if end < cur.next.end {
			end = cur.next.end // take farthest end
		}
		cur = cur.next
	}
	// Return this segment length + next segment length
	return end-start + calcWidth(cur.next)
}
