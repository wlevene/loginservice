package vectordb

type (
	QueryMatch struct{}

	VectorDB interface {
		UpsertEmbeddings() error
		Retrieve() ([]QueryMatch, error)
	}
)
