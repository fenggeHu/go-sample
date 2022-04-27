package main

type Greet struct {
	p []byte
}

func (g *Greet) Read(p []byte) (n int, err error) {
	g.p = p
	return len(g.p), nil
}

func (g *Greet) Write(p []byte) (n int, err error) {
	p = g.p
	return len(g.p), nil
}

func (g *Greet) String() string {
	return string(g.p)
}
