package main

type Retiever interface {
	Get(url string) string
}

func Download(r Retiever) string {
	return r.Get("adsfasdfasdf")
}

func main() {

}
