package main

import "fmt"

type ERC20 interface {
	totalSupply() int
	balanceOf (tokenOwner string) int
	allowance (tokenOwner string, spender string ) int
	transfer (tostring, tokens int) int
	approve (spender string, tokens int) int
	transferFrom (from string, to string, tokens int ) int
}

type contract struct {
 Owner string
 balance map[string] int
 allowed map[string]map[string] int
 totalsupply int
}



func (c contract) totalSupply() int {


return c.totalsupply

}

func (c contract) balanceOf (tokenOwner string) int {


return c.balance[tokenOwner]

}

func (c contract) allowance (tokenOwner string, spender string) int{


return c.allowed[tokenOwner][spender]

}

func (c contract) transfer (to string, tokens int) bool{

  	c.balance[to]  = c.balance[to] + tokens
  	c.balance[c.Owner] = c.balance[c.Owner] - tokens
	return true

}

func (c contract) transferFrom (from string, to string, tokens int) bool{


 	c.balance[from] = c.balance[from] - tokens;
        c.balance[to] = c.balance[to] +  tokens;
	return true

}

func main() {

 	a := contract {}
	a.balance = make(map[string]int)
	a.allowed = make(map[string]map[string]int)

	a.Owner = "0x7e2E7d9c5917D9399101054Ec69f5Ed19f256b19"
	
	a.totalsupply  = 1321131
	
	a.balance[a.Owner] = a.totalsupply
	
	a.balance["0x7e2E7d9c5917D9399101054Ec69f5Ed19f2563123213"] = 0
	
	a.transfer("0x7e2E7d9c5917D9399101054Ec69f5Ed19f2563123213",1213)
	
	a.balanceOf("0x7e2E7d9c5917D9399101054Ec69f5Ed19f2563123213")
	
	fmt.Println(a.balanceOf("0x7e2E7d9c5917D9399101054Ec69f5Ed19f2563123213"))

}
