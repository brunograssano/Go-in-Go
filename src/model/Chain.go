package model

type Chain struct {
	stones      map[Position]Position
	liberties   map[Position]Position
	rivalStones map[Position]Position
}

func (chain *Chain) AddPos(pos Position) {
	chain.stones[pos] = pos
}

func (chain *Chain) AddLiberty(pos Position) {
	chain.liberties[pos] = pos
}

func (chain *Chain) AddRival(pos Position) {
	chain.rivalStones[pos] = pos
}

func (chain *Chain) HasPos(pos Position) bool {
	_, visited := chain.stones[pos]
	return visited
}

func (chain *Chain) GetAmountOfRivalStones() uint {
	return uint(len(chain.rivalStones))
}

func (chain *Chain) GetLiberties() uint {
	return uint(len(chain.liberties))
}

func (chain *Chain) HasAnyLiberties() bool {
	return uint(len(chain.liberties)) > 0
}

func NewEmptyChain() Chain {
	stones := make(map[Position]Position)
	liberties := make(map[Position]Position)
	rivalStones := make(map[Position]Position)
	chain := Chain{stones, liberties, rivalStones}
	return chain
}

func NewChain(initialPos Position) Chain {
	stones := make(map[Position]Position)
	liberties := make(map[Position]Position)
	rivalStones := make(map[Position]Position)
	chain := Chain{stones, liberties, rivalStones}
	chain.stones[initialPos] = initialPos
	return chain
}
