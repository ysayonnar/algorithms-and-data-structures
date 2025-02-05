package leetcode

//следующий больший элемент слева!!!, включая нынешний

type StockSpanner struct {
	Stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{Stack: make([][2]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	if len(this.Stack) == 0 || this.Stack[len(this.Stack)-1][0] > price {
		this.Stack = append(this.Stack, [2]int{price, 1})
		return 1
	}

	span := 0
	for len(this.Stack) > 0 && this.Stack[len(this.Stack)-1][0] <= price {
		span += this.Stack[len(this.Stack)-1][1]
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
	this.Stack = append(this.Stack, [2]int{price, span + 1})
	return span + 1
}
