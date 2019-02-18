package PriorityQueue

type item struct {
	prior int
	val   int
}

type Queue struct {
	list []*item
}

func (p *Queue) Enqueue(value, prior int) {
	p.list = append(p.list, &item{prior, value})
	i := len(p.list) - 1

	var ppos int
	var cur, parent *item
	for {
		ppos = (i-1)/2
		cur = p.list[i]
		parent = p.list[ppos]
		if cur.prior > parent.prior {
			p.list[i] = parent
			p.list[ppos] = cur
			i = ppos
		}else{
			break
		}
	}
}

func (p *Queue) Dequeue() int {
	n := len(p.list)
	if n < 1 {
		return -1
	}
	result := p.list[0]
	p.list[0] = p.list[len(p.list)-1]
	p.list = p.list[:len(p.list)-1]
	i, maxPos := 0, 0
	var tmp *item
	for {
		if 2*i+1 < n && p.list[2*i+1].prior > p.list[maxPos].prior {
			maxPos = 2*i+1
		}
		if 2*i+2 < n && p.list[2*i+2].prior > p.list[maxPos].prior {
			maxPos = 2*i+2
		}
		if i == maxPos { break }
		tmp = p.list[i]
		p.list[i] = p.list[maxPos]
		p.list[maxPos] = tmp
		i = maxPos
	}
	return result.val
}
