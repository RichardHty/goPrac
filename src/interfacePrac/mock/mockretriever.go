package mock

type Retriever struct {
	Contents string
}

func (ret Retriever) Get(url string) string {
	return ret.Contents
}



