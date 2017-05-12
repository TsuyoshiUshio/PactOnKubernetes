// Package types contains types to use across the Consumer/Provider tests.
package types

// UserRepository is an in-memory user database.
type ProductRepository struct {
	Product map[string]*Product
}

// ByUsername finds a user by their username.
func (p *ProductRepository) ByKeyword(keyword string) (*Product, error) {
	if product, ok := p.Product[keyword]; ok {
		return product, nil
	}
	return nil, ErrNotFound
}
