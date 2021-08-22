package main

type Chain struct {
	stones              map[Position]Position
	liberties           uint
	amountOfRivalStones uint
}

func (chain *Chain) AddPos(pos Position) {
	chain.stones[pos] = pos
}

func (chain *Chain) AddLiberty() {
	chain.liberties++
}

func (chain *Chain) AddRival() {
	chain.amountOfRivalStones++
}

func (chain *Chain) HasPos(pos Position) bool {
	_, visited := chain.stones[pos]
	return visited
}

func NewChain(initialPos Position) Chain {
	stones := make(map[Position]Position)
	chain := Chain{stones, 0, 0}
	chain.stones[initialPos] = initialPos
	return chain
}
